package postgresqlsqlcrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"postgresqlsqlcrossplaneio.Database",
		reflect.TypeOf((*Database)(nil)).Elem(),
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
			j := jsiiProxy_Database{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.DatabaseProps",
		reflect.TypeOf((*DatabaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.DatabaseSpec",
		reflect.TypeOf((*DatabaseSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlsqlcrossplaneio.DatabaseSpecDeletionPolicy",
		reflect.TypeOf((*DatabaseSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": DatabaseSpecDeletionPolicy_ORPHAN,
			"DELETE": DatabaseSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.DatabaseSpecForProvider",
		reflect.TypeOf((*DatabaseSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.DatabaseSpecProviderConfigRef",
		reflect.TypeOf((*DatabaseSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.DatabaseSpecProviderRef",
		reflect.TypeOf((*DatabaseSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.DatabaseSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*DatabaseSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlsqlcrossplaneio.ProviderConfig",
		reflect.TypeOf((*ProviderConfig)(nil)).Elem(),
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
			j := jsiiProxy_ProviderConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.ProviderConfigProps",
		reflect.TypeOf((*ProviderConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.ProviderConfigSpec",
		reflect.TypeOf((*ProviderConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.ProviderConfigSpecCredentials",
		reflect.TypeOf((*ProviderConfigSpecCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlsqlcrossplaneio.ProviderConfigSpecCredentialsConnectionSecretRef",
		reflect.TypeOf((*ProviderConfigSpecCredentialsConnectionSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlsqlcrossplaneio.ProviderConfigSpecCredentialsSource",
		reflect.TypeOf((*ProviderConfigSpecCredentialsSource)(nil)).Elem(),
		map[string]interface{}{
			"POSTGRE_SQL_CONNECTION_SECRET": ProviderConfigSpecCredentialsSource_POSTGRE_SQL_CONNECTION_SECRET,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlsqlcrossplaneio.ProviderConfigSpecSslMode",
		reflect.TypeOf((*ProviderConfigSpecSslMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": ProviderConfigSpecSslMode_DISABLE,
			"REQUIRE": ProviderConfigSpecSslMode_REQUIRE,
			"VERIFY_CA": ProviderConfigSpecSslMode_VERIFY_CA,
			"VERIFY_FULL": ProviderConfigSpecSslMode_VERIFY_FULL,
		},
	)
}
