package software.amazon.polymorph.smithygo.shapevisitor;

import java.util.Map;
import java.util.Map.Entry;

import software.amazon.polymorph.smithygo.codegen.GenerationContext;
import software.amazon.polymorph.smithygo.codegen.GoWriter;
import software.amazon.polymorph.smithygo.codegen.knowledge.GoPointableIndex;
import software.amazon.polymorph.smithygo.nameresolver.DafnyNameResolver;
import software.amazon.polymorph.traits.ReferenceTrait;
import software.amazon.smithy.codegen.core.CodegenException;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.BigDecimalShape;
import software.amazon.smithy.model.shapes.BigIntegerShape;
import software.amazon.smithy.model.shapes.BlobShape;
import software.amazon.smithy.model.shapes.BooleanShape;
import software.amazon.smithy.model.shapes.ByteShape;
import software.amazon.smithy.model.shapes.DoubleShape;
import software.amazon.smithy.model.shapes.EnumShape;
import software.amazon.smithy.model.shapes.FloatShape;
import software.amazon.smithy.model.shapes.IntegerShape;
import software.amazon.smithy.model.shapes.ListShape;
import software.amazon.smithy.model.shapes.LongShape;
import software.amazon.smithy.model.shapes.MapShape;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.ResourceShape;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeVisitor;
import software.amazon.smithy.model.shapes.ShortShape;
import software.amazon.smithy.model.shapes.StringShape;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.shapes.TimestampShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.model.traits.EnumTrait;
import software.amazon.smithy.model.traits.ErrorTrait;
import software.amazon.smithy.model.traits.RangeTrait;
import software.amazon.smithy.utils.CaseUtils;
import software.amazon.smithy.utils.StringUtils;

public class SmithyToDafnyShapeVisitor extends ShapeVisitor.Default<String> {
    private final GenerationContext context;
    private final String dataSource;
    private final GoWriter writer;
    private final boolean isConfigShape;

    private final boolean isOptional;
    private final boolean isPointerType;

    public SmithyToDafnyShapeVisitor(
            final GenerationContext context,
            final String dataSource,
            final GoWriter writer,
            final boolean isConfigShape,
            final boolean isOptional,
            final boolean isPointerType
    ) {
        this.context = context;
        this.dataSource = dataSource;
        this.writer = writer;
        this.isConfigShape = isConfigShape;
        this.isOptional = isOptional;
        this.isPointerType = isPointerType;
    }

    protected String referenceStructureShape(StructureShape shape) {
        ReferenceTrait referenceTrait = shape.expectTrait(ReferenceTrait.class);
        Shape resourceOrService = context.model().expectShape(referenceTrait.getReferentId());

        if (resourceOrService.asResourceShape().isPresent()) {
            ResourceShape resourceShape = resourceOrService.asResourceShape().get();
            return dataSource + ".(*%s).impl".formatted(resourceShape.toShapeId().getName());
        }

        if (resourceOrService.asServiceShape().isPresent()) {
            ServiceShape resourceShape = resourceOrService.asServiceShape().get();
            return dataSource + ".(*%s).impl".formatted(resourceShape.toShapeId().getName());
        }

        throw new UnsupportedOperationException("Unknown referenceStructureShape type: " + shape);
    }

    @Override
    protected String getDefault(Shape shape) {
        throw new CodegenException(String.format(
                "Unsupported conversion of %s to %s using the %s protocol",
                shape, shape.getType(), context.protocolGenerator().getName()));
    }

    @Override
    public String blobShape(BlobShape shape) {
        writer.addImport("dafny");
        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s";
        String returnType = "dafny.Sequence";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
            returnType = "Wrappers.Option";
        }
        return """
                func () %s {
                    var v []interface{}
                    if %s == nil {return %s}
                    for _, e := range %s {
                    	v = append(v, e)
                    }
                    return %s;
                }()""".formatted(returnType, dataSource, nilWrapIfRequired, dataSource, someWrapIfRequired.formatted("dafny.SeqOf(v...)"));
    }

    @Override
    public String structureShape(final StructureShape shape) {
        if (shape.hasTrait(ReferenceTrait.class)) {
            return referenceStructureShape(shape);
        }
        final var builder = new StringBuilder();
        writer.addImport("Wrappers");
        writer.addImport(DafnyNameResolver.dafnyTypesNamespace(context.settings()));

        String someWrapIfRequired = "%s";

        String companionStruct;
        String returnType;
        if (shape.hasTrait(ErrorTrait.class)) {
            companionStruct = DafnyNameResolver.getDafnyErrorCompanionCreate(context.settings(), context.symbolProvider().toSymbol(shape));
            returnType = DafnyNameResolver.getDafnyBaseErrorType(context.settings());
        } else {
            companionStruct = DafnyNameResolver.getDafnyCompanionTypeCreate(context.settings(), context.symbolProvider().toSymbol(shape));
            returnType = DafnyNameResolver.getDafnyType(context.settings(), context.symbolProvider().toSymbol(shape));
        }
        String nilWrapIfRequired = returnType.concat("{}");

        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
            returnType = "Wrappers.Option";
        }

        var nilCheck = "";
        if (isPointerType) {
            nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
        }
        var goCodeBlock = """
                               func () %s {
                                   %s
                                   return %s
                               }()""";


        builder.append("%1$s(".formatted(companionStruct));
        String fieldSeparator = ",";
        for (final var memberShapeEntry : shape.getAllMembers().entrySet()) {
            // System.out.println("Pointer:"+isPointerType);
            final var memberName = memberShapeEntry.getKey();
            final var memberShape = memberShapeEntry.getValue();
            final var targetShape = context.model().expectShape(memberShape.getTarget());
            // System.out.println("\n\n\n\n Smithy to Dafny Shape Visitor:" + targetShape + targetShape.hasTrait(RangeTrait.class));
            if(targetShape.hasTrait(RangeTrait.class))
                builder.append("%1$s%2$s".formatted(
                    targetShape.accept(
                            new SmithyToDafnyShapeVisitor(context, dataSource + "." + StringUtils.capitalize(memberName),
                                                          writer, isConfigShape, memberShape.isOptional(), false
                                                          )), fieldSeparator
                ));
            else
                builder.append("%1$s%2$s".formatted(
                        targetShape.accept(
                                new SmithyToDafnyShapeVisitor(context, dataSource + "." + StringUtils.capitalize(memberName),
                                                            writer, isConfigShape, memberShape.isOptional(), true
                                                            )), fieldSeparator
                ));
        }


        return goCodeBlock.formatted(returnType, nilCheck, someWrapIfRequired.formatted(builder.append(")").toString()));
    }

    @Override
    public String mapShape(MapShape shape) {
        StringBuilder builder = new StringBuilder();

        writer.addImport("dafny");

        MemberShape keyMemberShape = shape.getKey();
        final Shape keyTargetShape = context.model().expectShape(keyMemberShape.getTarget());
        MemberShape valueMemberShape = shape.getValue();
        final Shape valueTargetShape = context.model().expectShape(valueMemberShape.getTarget());
        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s";
        String returnType = "dafny.Map";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
            returnType = "Wrappers.Option";
        }
        builder.append("""
                               func () %s {
                                   if %s == nil { return %s }
                               	   fieldValue := dafny.NewMapBuilder()
                               	   for key, val := range %s {
                               		    fieldValue.Add(%s, %s)
                               	   }
                               	   return %s
                               }()""".formatted(returnType, dataSource, nilWrapIfRequired, dataSource,
                                             keyTargetShape.accept(
                                                     new SmithyToDafnyShapeVisitor(context, "key", writer, isConfigShape, false, false)),
                                             valueTargetShape.accept(
                                                     new SmithyToDafnyShapeVisitor(context, "val", writer, isConfigShape, false, false)),
                                             someWrapIfRequired.formatted("fieldValue.ToMap()")
                       )
        );

        // Close structure
        return builder.toString();

    }

    @Override
    public String listShape(ListShape shape) {
        writer.addImport("dafny");

        StringBuilder builder = new StringBuilder();

        MemberShape memberShape = shape.getMember();
        final Shape targetShape = context.model().expectShape(memberShape.getTarget());

        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s";
        String returnType = "dafny.Sequence";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
            returnType = "Wrappers.Option";
        }

        builder.append("""
                               func () %s {
                                      if %s == nil { return %s }
                                      var fieldValue []interface{} = make([]interface{}, 0)
                                      for _, val := range %s {
                                          element := %s
                                          fieldValue = append(fieldValue, element)
                                      }
                                      return %s
                               }()""".formatted(returnType, dataSource, nilWrapIfRequired, dataSource,
                                             targetShape.accept(
                                                     new SmithyToDafnyShapeVisitor(context, "val", writer, isConfigShape, false, false)
                                             ), someWrapIfRequired.formatted("dafny.SeqOf(fieldValue...)")));

        // Close structure
        return builder.toString();
    }

    @Override
    public String booleanShape(BooleanShape shape) {
        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s%s";
        String returnType = "interface {}";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s%s)";
            returnType = "Wrappers.Option";
        }

        var dereferenceIfRequired = isPointerType ? "*" : "";
        var nilCheck = "";
        if (isPointerType) {
            nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
        }

        return """
                    func () %s {
                        %s
                        return %s
                    }()""".formatted(returnType, nilCheck,  someWrapIfRequired.formatted(dereferenceIfRequired, dataSource));
    }

    @Override
    public String stringShape(StringShape shape) {
        writer.addImport("dafny");
        if (shape.hasTrait(EnumTrait.class)) {
            String nilWrapIfRequired = "nil";
            String someWrapIfRequired = "%s";
            String returnType = DafnyNameResolver.getDafnyType(context.settings(), context.symbolProvider().toSymbol(shape));
            if (this.isOptional) {
                nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
                someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
                returnType = "Wrappers.Option";
            }

            var nilCheck = "";
            var dereferenceIfRequired = isPointerType ? "*" : "";
            if (isPointerType) {
                nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
            }
            return """
        func () %s {
        %s
		var index int
		for _, enumVal := range %s.Values() {
			index++
			if enumVal == %s%s{
				break;
			}
		}
		var enum interface{}
		for allEnums, i := dafny.Iterate(%s{}.AllSingletonConstructors()), 0; i < index; i++ {
			var ok bool
			enum, ok = allEnums()
			if !ok {
				break;
			}
		}
		return %s
	}()""".formatted(returnType, nilCheck, dataSource, dereferenceIfRequired, dataSource, DafnyNameResolver.getDafnyCompanionStructType(context.settings(), context.symbolProvider().toSymbol(shape)), someWrapIfRequired.formatted("enum"));
        } else {
            String nilWrapIfRequired = "nil";
            String someWrapIfRequired = "%s";
            String returnType = "dafny.Sequence";
            if (this.isOptional) {
                nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
                someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
                returnType = "Wrappers.Option";
            }

            var nilCheck = "";
            var dereferenceIfRequired = isPointerType ? "*" : "";
            if (isPointerType) {
                nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
            }

            return """
                    func () %s {
                        %s
                        return %s
                    }()""".formatted(returnType, nilCheck, someWrapIfRequired.formatted("dafny.SeqOfChars([]dafny.Char(%s%s)...)".formatted(dereferenceIfRequired, dataSource)));
        }
    }

    @Override
    public String integerShape(IntegerShape shape) {
        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s%s";
        String returnType = "interface {}";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s%s)";
            returnType = "Wrappers.Option";
        }

        // System.out.println("\n\n\n\n Smithy to Dafny Shape Visitor:" + shape.getId());
        // System.out.println(isPointerType);

        var dereferenceIfRequired = isPointerType ? "*" : "";
        var nilCheck = "";
        if (isPointerType) {
            nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
        }

        return """
                    func () %s {
                        %s
                        return %s
                    }()""".formatted(returnType, nilCheck,  someWrapIfRequired.formatted(dereferenceIfRequired, dataSource));

    }

    @Override
    public String longShape(LongShape shape) {
        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s%s";
        String returnType = "interface {}";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s%s)";
            returnType = "Wrappers.Option";
        }

        var dereferenceIfRequired = isPointerType ? "*" : "";
        var nilCheck = "";
        if (isPointerType) {
            nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
        }

        return """
                    func () %s {
                        %s
                        return %s
                    }()""".formatted(returnType, nilCheck,  someWrapIfRequired.formatted(dereferenceIfRequired, dataSource));
    }

    @Override
    public String doubleShape(DoubleShape shape) {
        writer.addImport("dafny");
        writer.addImport("encoding/binary");
        writer.addImport("math");

        String nilWrapIfRequired = "nil";
        String someWrapIfRequired = "%s";
        String returnType = "interface {}";
        if (this.isOptional) {
            nilWrapIfRequired = "Wrappers.Companion_Option_.Create_None_()";
            someWrapIfRequired = "Wrappers.Companion_Option_.Create_Some_(%s)";
            returnType = "Wrappers.Option";
        }

        var dereferenceIfRequired = isPointerType ? "*" : "";
        var nilCheck = "";
        if (isPointerType) {
            nilCheck = "if %s == nil {return %s}".formatted(dataSource, nilWrapIfRequired);
        }

        return """
                func () %s {
                    %s
    	            var bits = math.Float64bits(%s%s)
                    var bytes = make([]byte, 8)
                    binary.LittleEndian.PutUint64(bytes, bits)
    	            var v []interface{}
    	            for _, e := range bytes {
    		            v = append(v, e)
    	            }
    	            return %s;
                }()""".formatted(returnType, nilCheck, dereferenceIfRequired, dataSource, someWrapIfRequired.formatted("dafny.SeqOf(v...)"));
    }

    @Override
    public String unionShape(UnionShape shape) {
        return "Unionnnnnn";
    }

    @Override
    public String timestampShape(TimestampShape shape) {
        return "Timestampppppp";
    }
}
