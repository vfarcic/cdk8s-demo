package secretscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"secretscrossplaneio.StoreConfig",
		reflect.TypeOf((*StoreConfig)(nil)).Elem(),
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
			j := jsiiProxy_StoreConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigProps",
		reflect.TypeOf((*StoreConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpec",
		reflect.TypeOf((*StoreConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecKubernetes",
		reflect.TypeOf((*StoreConfigSpecKubernetes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecKubernetesAuth",
		reflect.TypeOf((*StoreConfigSpecKubernetesAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecKubernetesAuthEnv",
		reflect.TypeOf((*StoreConfigSpecKubernetesAuthEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecKubernetesAuthFs",
		reflect.TypeOf((*StoreConfigSpecKubernetesAuthFs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecKubernetesAuthSecretRef",
		reflect.TypeOf((*StoreConfigSpecKubernetesAuthSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"secretscrossplaneio.StoreConfigSpecKubernetesAuthSource",
		reflect.TypeOf((*StoreConfigSpecKubernetesAuthSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": StoreConfigSpecKubernetesAuthSource_NONE,
			"SECRET": StoreConfigSpecKubernetesAuthSource_SECRET,
			"ENVIRONMENT": StoreConfigSpecKubernetesAuthSource_ENVIRONMENT,
			"FILESYSTEM": StoreConfigSpecKubernetesAuthSource_FILESYSTEM,
		},
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVault",
		reflect.TypeOf((*StoreConfigSpecVault)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultAuth",
		reflect.TypeOf((*StoreConfigSpecVaultAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultAuthToken",
		reflect.TypeOf((*StoreConfigSpecVaultAuthToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultAuthTokenEnv",
		reflect.TypeOf((*StoreConfigSpecVaultAuthTokenEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultAuthTokenFs",
		reflect.TypeOf((*StoreConfigSpecVaultAuthTokenFs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultAuthTokenSecretRef",
		reflect.TypeOf((*StoreConfigSpecVaultAuthTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"secretscrossplaneio.StoreConfigSpecVaultAuthTokenSource",
		reflect.TypeOf((*StoreConfigSpecVaultAuthTokenSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": StoreConfigSpecVaultAuthTokenSource_NONE,
			"SECRET": StoreConfigSpecVaultAuthTokenSource_SECRET,
			"ENVIRONMENT": StoreConfigSpecVaultAuthTokenSource_ENVIRONMENT,
			"FILESYSTEM": StoreConfigSpecVaultAuthTokenSource_FILESYSTEM,
		},
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultCaBundle",
		reflect.TypeOf((*StoreConfigSpecVaultCaBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultCaBundleEnv",
		reflect.TypeOf((*StoreConfigSpecVaultCaBundleEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultCaBundleFs",
		reflect.TypeOf((*StoreConfigSpecVaultCaBundleFs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"secretscrossplaneio.StoreConfigSpecVaultCaBundleSecretRef",
		reflect.TypeOf((*StoreConfigSpecVaultCaBundleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"secretscrossplaneio.StoreConfigSpecVaultCaBundleSource",
		reflect.TypeOf((*StoreConfigSpecVaultCaBundleSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": StoreConfigSpecVaultCaBundleSource_NONE,
			"SECRET": StoreConfigSpecVaultCaBundleSource_SECRET,
			"ENVIRONMENT": StoreConfigSpecVaultCaBundleSource_ENVIRONMENT,
			"FILESYSTEM": StoreConfigSpecVaultCaBundleSource_FILESYSTEM,
		},
	)
}
