package apiextensionscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.CompositeResourceDefinition",
		reflect.TypeOf((*CompositeResourceDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CompositeResourceDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionProps",
		reflect.TypeOf((*CompositeResourceDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpec",
		reflect.TypeOf((*CompositeResourceDefinitionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecClaimNames",
		reflect.TypeOf((*CompositeResourceDefinitionSpecClaimNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecDefaultCompositionRef",
		reflect.TypeOf((*CompositeResourceDefinitionSpecDefaultCompositionRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecEnforcedCompositionRef",
		reflect.TypeOf((*CompositeResourceDefinitionSpecEnforcedCompositionRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecNames",
		reflect.TypeOf((*CompositeResourceDefinitionSpecNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecVersions",
		reflect.TypeOf((*CompositeResourceDefinitionSpecVersions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecVersionsAdditionalPrinterColumns",
		reflect.TypeOf((*CompositeResourceDefinitionSpecVersionsAdditionalPrinterColumns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecVersionsSchema",
		reflect.TypeOf((*CompositeResourceDefinitionSpecVersionsSchema)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.Composition",
		reflect.TypeOf((*Composition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Composition{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionProps",
		reflect.TypeOf((*CompositionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.CompositionRevision",
		reflect.TypeOf((*CompositionRevision)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CompositionRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionProps",
		reflect.TypeOf((*CompositionRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpec",
		reflect.TypeOf((*CompositionRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecCompositeTypeRef",
		reflect.TypeOf((*CompositionRevisionSpecCompositeTypeRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSets",
		reflect.TypeOf((*CompositionRevisionSpecPatchSets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatches",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombine",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombineStrategy",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecPatchSetsPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombineString",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombineVariables",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesPolicy",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransforms",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsConvert",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_STRING,
			"INT": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_INT,
			"INT64": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_FLOAT64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMath",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsString",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_REGEXP,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionRevisionSpecPatchSetsPatchesTransformsType_MAP,
			"MATH": CompositionRevisionSpecPatchSetsPatchesTransformsType_MATH,
			"STRING": CompositionRevisionSpecPatchSetsPatchesTransformsType_STRING,
			"CONVERT": CompositionRevisionSpecPatchSetsPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionRevisionSpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionRevisionSpecPatchSetsPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionRevisionSpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionRevisionSpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionRevisionSpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPublishConnectionDetailsWithStoreConfigRef",
		reflect.TypeOf((*CompositionRevisionSpecPublishConnectionDetailsWithStoreConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResources",
		reflect.TypeOf((*CompositionRevisionSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesConnectionDetails",
		reflect.TypeOf((*CompositionRevisionSpecResourcesConnectionDetails)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesConnectionDetailsType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesConnectionDetailsType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_CONNECTION_SECRET_KEY": CompositionRevisionSpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY,
			"FROM_FIELD_PATH": CompositionRevisionSpecResourcesConnectionDetailsType_FROM_FIELD_PATH,
			"FROM_VALUE": CompositionRevisionSpecResourcesConnectionDetailsType_FROM_VALUE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatches",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombine",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombineStrategy",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecResourcesPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombineString",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombineVariables",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesPolicy",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransforms",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsConvert",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_STRING,
			"INT": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_INT,
			"INT64": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_FLOAT64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMath",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsString",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_FROM_BASE64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionRevisionSpecResourcesPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionRevisionSpecResourcesPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionRevisionSpecResourcesPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionRevisionSpecResourcesPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionRevisionSpecResourcesPatchesTransformsStringType_REGEXP,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionRevisionSpecResourcesPatchesTransformsType_MAP,
			"MATH": CompositionRevisionSpecResourcesPatchesTransformsType_MATH,
			"STRING": CompositionRevisionSpecResourcesPatchesTransformsType_STRING,
			"CONVERT": CompositionRevisionSpecResourcesPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionRevisionSpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionRevisionSpecResourcesPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionRevisionSpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionRevisionSpecResourcesPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionRevisionSpecResourcesPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesReadinessChecks",
		reflect.TypeOf((*CompositionRevisionSpecResourcesReadinessChecks)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesReadinessChecksType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesReadinessChecksType)(nil)).Elem(),
		map[string]interface{}{
			"MATCH_STRING": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_STRING,
			"MATCH_INTEGER": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_INTEGER,
			"NON_EMPTY": CompositionRevisionSpecResourcesReadinessChecksType_NON_EMPTY,
			"NONE": CompositionRevisionSpecResourcesReadinessChecksType_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpec",
		reflect.TypeOf((*CompositionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecCompositeTypeRef",
		reflect.TypeOf((*CompositionSpecCompositeTypeRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSets",
		reflect.TypeOf((*CompositionSpecPatchSets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatches",
		reflect.TypeOf((*CompositionSpecPatchSetsPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombine",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombineStrategy",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecPatchSetsPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombineString",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombineVariables",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesPolicy",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionSpecPatchSetsPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionSpecPatchSetsPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransforms",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsConvert",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecPatchSetsPatchesTransformsConvertToType_STRING,
			"INT": CompositionSpecPatchSetsPatchesTransformsConvertToType_INT,
			"INT64": CompositionSpecPatchSetsPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionSpecPatchSetsPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionSpecPatchSetsPatchesTransformsConvertToType_FLOAT64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMath",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsString",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionSpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionSpecPatchSetsPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionSpecPatchSetsPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionSpecPatchSetsPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionSpecPatchSetsPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionSpecPatchSetsPatchesTransformsStringType_REGEXP,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionSpecPatchSetsPatchesTransformsType_MAP,
			"MATH": CompositionSpecPatchSetsPatchesTransformsType_MATH,
			"STRING": CompositionSpecPatchSetsPatchesTransformsType_STRING,
			"CONVERT": CompositionSpecPatchSetsPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionSpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionSpecPatchSetsPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionSpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionSpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionSpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPublishConnectionDetailsWithStoreConfigRef",
		reflect.TypeOf((*CompositionSpecPublishConnectionDetailsWithStoreConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResources",
		reflect.TypeOf((*CompositionSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesConnectionDetails",
		reflect.TypeOf((*CompositionSpecResourcesConnectionDetails)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesConnectionDetailsType",
		reflect.TypeOf((*CompositionSpecResourcesConnectionDetailsType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_CONNECTION_SECRET_KEY": CompositionSpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY,
			"FROM_FIELD_PATH": CompositionSpecResourcesConnectionDetailsType_FROM_FIELD_PATH,
			"FROM_VALUE": CompositionSpecResourcesConnectionDetailsType_FROM_VALUE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatches",
		reflect.TypeOf((*CompositionSpecResourcesPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombine",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombineStrategy",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecResourcesPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombineString",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombineVariables",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesPolicy",
		reflect.TypeOf((*CompositionSpecResourcesPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionSpecResourcesPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionSpecResourcesPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionSpecResourcesPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionSpecResourcesPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransforms",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsConvert",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecResourcesPatchesTransformsConvertToType_STRING,
			"INT": CompositionSpecResourcesPatchesTransformsConvertToType_INT,
			"INT64": CompositionSpecResourcesPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionSpecResourcesPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionSpecResourcesPatchesTransformsConvertToType_FLOAT64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMath",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsString",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionSpecResourcesPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionSpecResourcesPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionSpecResourcesPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionSpecResourcesPatchesTransformsStringConvert_FROM_BASE64,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionSpecResourcesPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionSpecResourcesPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionSpecResourcesPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionSpecResourcesPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionSpecResourcesPatchesTransformsStringType_REGEXP,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionSpecResourcesPatchesTransformsType_MAP,
			"MATH": CompositionSpecResourcesPatchesTransformsType_MATH,
			"STRING": CompositionSpecResourcesPatchesTransformsType_STRING,
			"CONVERT": CompositionSpecResourcesPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionSpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionSpecResourcesPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionSpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionSpecResourcesPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionSpecResourcesPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesReadinessChecks",
		reflect.TypeOf((*CompositionSpecResourcesReadinessChecks)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesReadinessChecksType",
		reflect.TypeOf((*CompositionSpecResourcesReadinessChecksType)(nil)).Elem(),
		map[string]interface{}{
			"MATCH_STRING": CompositionSpecResourcesReadinessChecksType_MATCH_STRING,
			"MATCH_INTEGER": CompositionSpecResourcesReadinessChecksType_MATCH_INTEGER,
			"NON_EMPTY": CompositionSpecResourcesReadinessChecksType_NON_EMPTY,
			"NONE": CompositionSpecResourcesReadinessChecksType_NONE,
		},
	)
}
