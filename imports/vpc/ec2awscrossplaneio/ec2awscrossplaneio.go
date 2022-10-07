package ec2awscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"ec2awscrossplaneio.Vpc",
		reflect.TypeOf((*Vpc)(nil)).Elem(),
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
			j := jsiiProxy_Vpc{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcProps",
		reflect.TypeOf((*VpcProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpec",
		reflect.TypeOf((*VpcSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecDeletionPolicy",
		reflect.TypeOf((*VpcSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": VpcSpecDeletionPolicy_ORPHAN,
			"DELETE": VpcSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecForProvider",
		reflect.TypeOf((*VpcSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecForProviderTags",
		reflect.TypeOf((*VpcSpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecProviderConfigRef",
		reflect.TypeOf((*VpcSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecProviderConfigRefPolicy",
		reflect.TypeOf((*VpcSpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*VpcSpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": VpcSpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": VpcSpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*VpcSpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": VpcSpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": VpcSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecProviderRef",
		reflect.TypeOf((*VpcSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecProviderRefPolicy",
		reflect.TypeOf((*VpcSpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecProviderRefPolicyResolution",
		reflect.TypeOf((*VpcSpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": VpcSpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": VpcSpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecProviderRefPolicyResolve",
		reflect.TypeOf((*VpcSpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": VpcSpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": VpcSpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecPublishConnectionDetailsTo",
		reflect.TypeOf((*VpcSpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*VpcSpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*VpcSpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*VpcSpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": VpcSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": VpcSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.VpcSpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*VpcSpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": VpcSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": VpcSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*VpcSpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.VpcSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*VpcSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
}
