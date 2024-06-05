// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package software.amazon.polymorph.smithyjava.generator.library;

import static javax.lang.model.element.Modifier.FINAL;
import static javax.lang.model.element.Modifier.PUBLIC;
import static software.amazon.polymorph.smithyjava.generator.Generator.Constants.TESTNG_ASSERT;
import static software.amazon.polymorph.smithyjava.generator.Generator.Constants.TESTNG_TEST;

import com.squareup.javapoet.CodeBlock;
import com.squareup.javapoet.JavaFile;
import com.squareup.javapoet.MethodSpec;
import com.squareup.javapoet.TypeName;
import com.squareup.javapoet.TypeSpec;
import java.util.LinkedHashSet;
import java.util.Set;
import software.amazon.polymorph.smithyjava.generator.Generator;
import software.amazon.polymorph.smithyjava.modeled.ModeledShapeValue;
import software.amazon.polymorph.traits.LocalServiceTrait;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.smoketests.traits.SmokeTestCase;
import software.amazon.smithy.smoketests.traits.SmokeTestsTrait;

/**
 * ModelTestCodegen generates tests for the content of the Subject's Model package.
 * e.g. generating JUnit tests from traits like @smithy.test#smokeTests.
 * This is distinct from the Dafny testing code
 * and the testing wrapper to support it generated by TestJavaLibrary.
 */
public class ModelTestCodegen extends Generator {

  final JavaLibrary subject;

  public ModelTestCodegen(JavaLibrary subject) {
    super(subject);
    this.subject = subject;
  }

  @Override
  public Set<JavaFile> javaFiles() {
    final LinkedHashSet<JavaFile> rtn = new LinkedHashSet<>();
    subject.model
      .getOperationShapesWithTrait(SmokeTestsTrait.class)
      .stream()
      .map(this::smokeTestsClass)
      .forEachOrdered(rtn::add);
    return rtn;
  }

  private JavaFile smokeTestsClass(OperationShape shape) {
    final TypeSpec.Builder spec = TypeSpec
      .classBuilder(shape.getId().getName() + "SmokeTests")
      .addModifiers(PUBLIC, FINAL);
    final SmokeTestsTrait smokeTests = shape.expectTrait(SmokeTestsTrait.class);
    smokeTests
      .getTestCases()
      .stream()
      .map(testCase -> smokeTest(shape, testCase))
      .forEachOrdered(spec::addMethod);
    final TypeSpec classType = spec.build();
    return JavaFile.builder(subject.modelPackageName, classType).build();
  }

  private MethodSpec smokeTest(
    final OperationShape operationShape,
    SmokeTestCase testCase
  ) {
    final String methodName = testCase.getId();
    final MethodSpec.Builder method = MethodSpec
      .methodBuilder(methodName)
      .addAnnotation(TESTNG_TEST)
      .addModifiers(PUBLIC)
      .returns(TypeName.VOID);

    final TypeName clientType = subject.nativeNameResolver.typeForShape(
      subject.serviceShape.toShapeId()
    );
    final String operationName = operationShape.toShapeId().getName();
    final ShapeId configShapeId = subject.serviceShape
      .expectTrait(LocalServiceTrait.class)
      .getConfigId();
    final TypeName configType = subject.nativeNameResolver.typeForShape(
      configShapeId
    );

    // SimpleConstraintsConfig config = SimpleConstraintsConfig.builder().build();
    // SimpleConstraints client = SimpleConstraints.builder()
    //         .SimpleConstraintsConfig(config)
    //         .build();
    method.addStatement(
      "$T config = $T.builder().build()",
      configType,
      configType
    );
    method.addStatement(
      "$T client = $T.builder().$N(config).build()",
      clientType,
      clientType,
      configShapeId.getName()
    );

    // GetConstraintsInput input = GetConstraintsInput.builder()
    // ...
    // (multiple .foo(...) calls to populate builder)
    // ...
    // .build();
    final Shape inputShape = subject.model.expectShape(
      operationShape.getInput().get()
    );
    final TypeName inputType = subject.nativeNameResolver.typeForShape(
      inputShape.getId()
    );
    final CodeBlock inputValue = ModeledShapeValue.shapeValue(
      subject,
      false,
      inputShape,
      testCase.getParams().get()
    );
    final CodeBlock inputAndClientCall = CodeBlock
      .builder()
      .addStatement("$T input = $L", inputType, inputValue)
      .addStatement("client.$L(input)", operationName)
      .build();

    if (testCase.getExpectation().isSuccess()) {
      method.addCode(inputAndClientCall);
    } else {
      // We're not specific about what kind of exception for now.
      // If the smokeTests trait gets more specific we can be too.
      // The inputAndClientCall.toString() is necessary because otherwise we get nested
      // $[ ... $] statements, which JavaPoet doesn't support.
      method.addStatement(
        "$T.assertThrows(Exception.class, () -> {\n$L\n})",
        TESTNG_ASSERT,
        inputAndClientCall.toString()
      );
    }

    return method.build();
  }
}