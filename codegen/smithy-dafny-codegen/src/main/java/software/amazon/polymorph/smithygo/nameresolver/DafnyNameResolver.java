package software.amazon.polymorph.smithygo.nameresolver;

import software.amazon.polymorph.smithygo.codegen.GenerationContext;
import software.amazon.polymorph.smithygo.codegen.GoSettings;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.model.shapes.ToShapeId;

import static software.amazon.polymorph.smithygo.nameresolver.Constants.BLANK;
import static software.amazon.polymorph.smithygo.nameresolver.Constants.DOT;
import static software.amazon.polymorph.smithygo.nameresolver.Constants.INTERNAL_DAFNY;
import static software.amazon.polymorph.smithygo.nameresolver.Constants.INTERNAL_DAFNY_TYPES;

public class DafnyNameResolver {

    public static String dafnyTypesNamespace(final GoSettings settings) {
        return settings.getModuleName().replace(DOT, BLANK).toLowerCase() + INTERNAL_DAFNY_TYPES;
    }

    public static String dafnyNamespace(final GoSettings settings) {
        return settings.getModuleName().replace(DOT, BLANK).toLowerCase() + INTERNAL_DAFNY;
    }

    public static String serviceNamespace(final ServiceShape serviceShape) {
        return serviceShape.toShapeId().getNamespace().replace(DOT, BLANK).toLowerCase();
    }

    public static String getDafnyType(final GoSettings settings, final Symbol symbol) {
        return DafnyNameResolver.dafnyTypesNamespace(settings).concat(DOT).concat(symbol.getName());
    }
    public static String getDafnySubErrorType(final GoSettings settings, final Symbol symbol) {
        return DafnyNameResolver.getDafnyBaseErrorType(settings).concat("_").concat(symbol.getName());
    }

    public static String getDafnyBaseErrorType(final GoSettings settings) {
        return DafnyNameResolver.dafnyTypesNamespace(settings).concat(DOT).concat("Error");
    }


    public static String getDafnyCompanionType(final GoSettings settings, final Symbol symbol) {
        return DafnyNameResolver.dafnyTypesNamespace(settings).concat(DOT).concat("Companion_%s_".formatted(symbol.getName()));
    }

    public static String getDafnyErrorCompanion(final GoSettings settings) {
        return DafnyNameResolver.dafnyTypesNamespace(settings).concat(DOT).concat("Companion_Error_");
    }

    public static String getDafnyErrorCompanionCreate(final GoSettings settings, final Symbol symbol) {
        return DafnyNameResolver.getDafnyErrorCompanion(settings).concat(DOT).concat("Create_%s_".formatted(symbol.getName()));
    }

    public static String getDafnyCompanionStructType(final GoSettings settings, final Symbol symbol) {
        return DafnyNameResolver.dafnyTypesNamespace(settings).concat(DOT).concat("CompanionStruct_%s_".formatted(symbol.getName()));
    }

    public static String getDafnyCompanionTypeCreate(final GoSettings settings, final Symbol symbol) {
        return DafnyNameResolver.getDafnyCompanionType(settings, symbol).concat(DOT).concat("Create_%s_".formatted(symbol.getName()));
    }

    public static String getInputToDafnyMethodName(final GenerationContext context, final ToShapeId shapeId) {
        return getToDafnyMethodName(context, shapeId, "_Input");
    }

    public static String getOutputToDafnyMethodName(final GenerationContext context, final ToShapeId shapeId) {
        return getToDafnyMethodName(context, shapeId, "_Output");
    }

    public static String getToDafnyMethodName(final GenerationContext context, final ToShapeId shapeId, final String disambiguator) {
        var name = context.settings().getService(context.model()).getContextualName(shapeId);
        return name.concat(disambiguator).concat("_ToDafny");
    }

    public static String getFromDafnyMethodName(final GenerationContext context, final ToShapeId shapeId, final String disambiguator) {
        var name = context.settings().getService(context.model()).getContextualName(shapeId);
        return name.concat(disambiguator).concat("_FromDafny");
    }

    public static String getOutputFromDafnyMethodName(final GenerationContext context, final ToShapeId shapeId) {
        var name = context.settings().getService(context.model()).getContextualName(shapeId);
        return getFromDafnyMethodName(context, shapeId, "_Output");
    }

    public static String getInputFromDafnyMethodName(final GenerationContext context, final ToShapeId shapeId) {
        var name = context.settings().getService(context.model()).getContextualName(shapeId);
        return getFromDafnyMethodName(context, shapeId, "_Input");
    }

    public static String getDafnyClient(final GoSettings settings, final String sdkId) {
        return DafnyNameResolver.dafnyNamespace(settings).concat(DOT).concat(sdkId).concat("Client");
    }
    public static String getDafnyInterfaceClient(final GoSettings settings, final String sdkId) {
        return DafnyNameResolver.dafnyTypesNamespace(settings).concat(DOT).concat("I").concat(settings.getService().getName()).concat("Client");
    }


    public static String createDafnyClient(final GoSettings settings, final String sdkId) {
        return DafnyNameResolver.dafnyNamespace(settings).concat(".Companion_Default___").concat(DOT).concat(sdkId);
    }

}