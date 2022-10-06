package metapkgcrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"metapkgcrossplaneio.Configuration",
		reflect.TypeOf((*Configuration)(nil)).Elem(),
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
			j := jsiiProxy_Configuration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationProps",
		reflect.TypeOf((*ConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationSpec",
		reflect.TypeOf((*ConfigurationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationSpecCrossplane",
		reflect.TypeOf((*ConfigurationSpecCrossplane)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationSpecDependsOn",
		reflect.TypeOf((*ConfigurationSpecDependsOn)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"metapkgcrossplaneio.ConfigurationV1Alpha1",
		reflect.TypeOf((*ConfigurationV1Alpha1)(nil)).Elem(),
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
			j := jsiiProxy_ConfigurationV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationV1Alpha1Props",
		reflect.TypeOf((*ConfigurationV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationV1Alpha1Spec",
		reflect.TypeOf((*ConfigurationV1Alpha1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationV1Alpha1SpecCrossplane",
		reflect.TypeOf((*ConfigurationV1Alpha1SpecCrossplane)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ConfigurationV1Alpha1SpecDependsOn",
		reflect.TypeOf((*ConfigurationV1Alpha1SpecDependsOn)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"metapkgcrossplaneio.Provider",
		reflect.TypeOf((*Provider)(nil)).Elem(),
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
			j := jsiiProxy_Provider{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderProps",
		reflect.TypeOf((*ProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderSpec",
		reflect.TypeOf((*ProviderSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderSpecController",
		reflect.TypeOf((*ProviderSpecController)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderSpecControllerPermissionRequests",
		reflect.TypeOf((*ProviderSpecControllerPermissionRequests)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderSpecCrossplane",
		reflect.TypeOf((*ProviderSpecCrossplane)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderSpecDependsOn",
		reflect.TypeOf((*ProviderSpecDependsOn)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"metapkgcrossplaneio.ProviderV1Alpha1",
		reflect.TypeOf((*ProviderV1Alpha1)(nil)).Elem(),
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
			j := jsiiProxy_ProviderV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderV1Alpha1Props",
		reflect.TypeOf((*ProviderV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderV1Alpha1Spec",
		reflect.TypeOf((*ProviderV1Alpha1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderV1Alpha1SpecController",
		reflect.TypeOf((*ProviderV1Alpha1SpecController)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderV1Alpha1SpecControllerPermissionRequests",
		reflect.TypeOf((*ProviderV1Alpha1SpecControllerPermissionRequests)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderV1Alpha1SpecCrossplane",
		reflect.TypeOf((*ProviderV1Alpha1SpecCrossplane)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"metapkgcrossplaneio.ProviderV1Alpha1SpecDependsOn",
		reflect.TypeOf((*ProviderV1Alpha1SpecDependsOn)(nil)).Elem(),
	)
}
