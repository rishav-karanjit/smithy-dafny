package software.amazon.polymorph.smithyjava.generator.awssdk;

import com.squareup.javapoet.CodeBlock;
import com.squareup.javapoet.MethodSpec;

import org.junit.Before;
import org.junit.Test;

import java.nio.file.Path;
import java.util.Map;

import software.amazon.polymorph.smithyjava.MethodReference;
import software.amazon.polymorph.smithyjava.ModelConstants;
import software.amazon.polymorph.utils.TokenTree;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.model.shapes.StructureShape;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertThrows;
import static software.amazon.polymorph.util.Tokenizer.tokenizeAndAssertEqual;

@SuppressWarnings("OptionalGetWithoutIsPresent")
public class ToDafnyTest {
    protected ToDafny underTest;
    protected Model model;

    @Before
    public void setup() {
        model = TestSetupUtils.setupTwoLocalModel(ModelConstants.KMS_KITCHEN, ModelConstants.OTHER_NAMESPACE);
        underTest  = new ToDafny(TestSetupUtils.setupAwsSdk(model, "kms"));
    }

    @Test
    public void memberConversionMethodReference() {
        ShapeId structureId = ShapeId.fromParts("com.amazonaws.kms", "Kitchen");
        StructureShape structureShape = model.expectShape(structureId, StructureShape.class);
        // If the target is simple, use SIMPLE_CONVERSION_METHOD_FROM_SHAPE_TYPE
        MemberShape stringMember = structureShape.getMember("name").get();
        MethodReference simpleActual = underTest.memberConversionMethodReference(stringMember);
        String simpleExpected = ToDafnyConstants.STRING_CONVERSION;
        tokenizeAndAssertEqual(simpleExpected, simpleActual.asNormalReference().toString());
        // if in namespace reference created converter
        MemberShape enumMember = structureShape.getMember("keyUsage").get();
        MethodReference enumActual = underTest.memberConversionMethodReference(enumMember);
        String enumExpected = ToDafnyConstants.KEY_USAGE_TYPE_CONVERSION;
        tokenizeAndAssertEqual(enumExpected, enumActual.asNormalReference().toString());
        // Otherwise, this target must be in another namespace
        MemberShape otherNamespaceMember = structureShape.getMember("otherNamespace").get();
        MethodReference otherNamespaceActual = underTest.memberConversionMethodReference(otherNamespaceMember);
        String otherNamespaceExpected = ToDafnyConstants.OTHER_NAMESPACE_CONVERSION;
        tokenizeAndAssertEqual(otherNamespaceExpected, otherNamespaceActual.asNormalReference().toString());
    }

    @Test
    public void memberAssignment() {
        ShapeId structureId = ShapeId.fromParts("com.amazonaws.kms", "Kitchen");
        StructureShape structureShape = model.expectShape(structureId, StructureShape.class);
        // if required
        MemberShape memberRequired = structureShape.getMember("name").get();
        CodeBlock actualRequired = underTest.memberAssignment(memberRequired, CodeBlock.of("name"));
        String expectedRequired = ToDafnyConstants.MEMBER_ASSIGNMENT_REQUIRED;
        tokenizeAndAssertEqual(expectedRequired, actualRequired.toString());
        // if optional
        MemberShape memberOptional = structureShape.getMember("message").get();
        CodeBlock actualOptional = underTest.memberAssignment(memberOptional, CodeBlock.of("message"));
        String expectedOptional = ToDafnyConstants.MEMBER_ASSIGNMENT_OPTIONAL;
        tokenizeAndAssertEqual(expectedOptional, actualOptional.toString());
    }

    @Test
    public void memberDeclaration() {
        ShapeId structureId = ShapeId.fromParts("com.amazonaws.kms", "Kitchen");
        StructureShape structureShape = model.expectShape(structureId, StructureShape.class);
        // if required
        MemberShape memberRequired = structureShape.getMember("name").get();
        CodeBlock actualRequired = underTest.memberDeclaration(memberRequired, CodeBlock.of("name"));
        String expectedRequired = ToDafnyConstants.MEMBER_DECLARATION_REQUIRED;
        tokenizeAndAssertEqual(expectedRequired, actualRequired.toString());
        // optional
        MemberShape memberOptional = structureShape.getMember("message").get();
        CodeBlock actualOptional = underTest.memberDeclaration(memberOptional, CodeBlock.of("message"));
        String expectedOptional = ToDafnyConstants.MEMBER_DECLARATION_OPTIONAL;
        tokenizeAndAssertEqual(expectedOptional, actualOptional.toString());
    }

    @Test
    public void generateConvertStructure() {
        // structureShape.members().size() == 0
        ShapeId simpleId = ShapeId.fromParts("com.amazonaws.kms", "Simple");
        MethodSpec simpleActual = underTest.generateConvertStructure(simpleId);
        String simpleExpected = ToDafnyConstants.SIMPLE_STRUCTURE;
        tokenizeAndAssertEqual(simpleExpected, simpleActual.toString());
        // structureShape.members().size() != 0
        ShapeId aOptionalId = ShapeId.fromParts("com.amazonaws.kms", "AOptional");
        MethodSpec aOptionalActual = underTest.generateConvert(aOptionalId);
        tokenizeAndAssertEqual(ToDafnyConstants.A_OPTIONAL_STRUCTURE, aOptionalActual.toString());
    }

    @Test
    public void generateConvertResponseV1() {
        Model localModel = TestSetupUtils.setupLocalModel(ModelConstants.KMS_A_STRING_OPERATION);
        ToDafny localUnderTest = new ToDafny(TestSetupUtils.setupAwsSdk(localModel, "kms"));
        ShapeId responseId = ShapeId.fromParts("com.amazonaws.kms", "DoSomethingResponse");
        MethodSpec actual = localUnderTest.generateConvertResponseV1(responseId);
        tokenizeAndAssertEqual(ToDafnyConstants.DO_SOMETHING_RESPONSE, actual.toString());
    }

    @Test
    public void generateConvertEnumEnum() {
        ShapeId enumId = ShapeId.fromParts("com.amazonaws.kms", "KeyUsageType");
        MethodSpec actual = underTest.generateConvertEnumEnum(enumId);
        tokenizeAndAssertEqual(ToDafnyConstants.KEY_USAGE_TYPE, actual.toString());
    }

    @Test
    public void generateConvertEnumString() {
        ShapeId enumId = ShapeId.fromParts("com.amazonaws.kms", "KeyUsageType");
        MethodSpec actual = underTest.generateConvertEnumString(enumId);
        tokenizeAndAssertEqual(ToDafnyConstants.KEY_USAGE_TYPE_STRING, actual.toString());
    }

    @Test
    public void generateConvert() {
        // case Simple
        ShapeId CiphertextTypeId = ShapeId.fromParts("com.amazonaws.kms", "CiphertextType");
        assertNull(underTest.generateConvert(CiphertextTypeId));
        // case LIST of Enums (which will take a list of Strings for AWS SDK for Java V1)
        ShapeId listEnumId = ShapeId.fromParts("com.amazonaws.kms", "KeyUsageTypes");
        tokenizeAndAssertEqual(ToDafnyConstants.GENERATE_CONVERT_LIST, underTest.generateConvert(listEnumId).toString());
        // case LIST of Structures from other AWS SDK namespace
        ShapeId listStructureId = ShapeId.fromParts("com.amazonaws.kms", "OtherNamespaces");
        tokenizeAndAssertEqual(ToDafnyConstants.GENERATE_CONVERT_LIST_STRUCTURES, underTest.generateConvert(listStructureId).toString());
        // case MAP
        ShapeId mapId = ShapeId.fromParts("com.amazonaws.kms", "EncryptionContextType");
        tokenizeAndAssertEqual(ToDafnyConstants.GENERATE_CONVERT_MAP_STRING, underTest.generateConvert(mapId).toString());
        // case SET
        ShapeId setId = ShapeId.fromParts("com.amazonaws.kms", "Names");
        tokenizeAndAssertEqual(ToDafnyConstants.GENERATE_CONVERT_SET_STRING, underTest.generateConvert(setId).toString());
        // case Structure
        ShapeId simpleId = ShapeId.fromParts("com.amazonaws.kms", "Simple");
        tokenizeAndAssertEqual(ToDafnyConstants.SIMPLE_STRUCTURE, underTest.generateConvert(simpleId).toString());
        // default
        ShapeId doubleId = ShapeId.fromParts("com.amazonaws.kms", "NotSupported");
        assertThrows(UnsupportedOperationException.class, () -> underTest.generateConvert(doubleId));
    }

    @Test
    public void generateConvertError() {
        Model localModel = TestSetupUtils.setupLocalModel(ModelConstants.KMS_A_STRING_OPERATION);
        ToDafny localUnderTest = new ToDafny(TestSetupUtils.setupAwsSdk(localModel, "kms"));
        ShapeId errorId = ShapeId.fromParts("com.amazonaws.kms", "DependencyTimeoutException");
        StructureShape errorShape = localModel.expectShape(errorId, StructureShape.class);
        MethodSpec errorActual = localUnderTest.generateConvertError(errorShape);
        String errorExpected = ToDafnyConstants.GENERATE_CONVERT_ERROR;
        tokenizeAndAssertEqual(errorExpected, errorActual.toString());
    }

    @Test
    public void generateConvertOpaqueError() {
        tokenizeAndAssertEqual(ToDafnyConstants.GENERATE_CONVERT_OPAQUE_ERROR, underTest.generateConvertOpaqueError().toString());
    }

    @Test
    public void generate() {
        Model localModel = TestSetupUtils.setupLocalModel(ModelConstants.KMS_A_STRING_OPERATION);
        ToDafny localUnderTest = new ToDafny(TestSetupUtils.setupAwsSdk(localModel, "kms"));
        final Map<Path, TokenTree> actual = localUnderTest.generate();
        final Path expectedPath = Path.of("Dafny/Com/Amazonaws/Kms/ToDafny.java");
        Path[] temp = new Path[1];
        final Path actualPath = actual.keySet().toArray(temp)[0];
        assertEquals(expectedPath, actualPath);
        final String actualSource = actual.get(actualPath).toString();
        tokenizeAndAssertEqual(ToDafnyConstants.KMS_A_STRING_OPERATION_JAVA_FILE, actualSource);
    }
}