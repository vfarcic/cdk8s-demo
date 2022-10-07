package ec2awscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"ec2awscrossplaneio.Subnet",
		reflect.TypeOf((*Subnet)(nil)).Elem(),
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
			j := jsiiProxy_Subnet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetProps",
		reflect.TypeOf((*SubnetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpec",
		reflect.TypeOf((*SubnetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecDeletionPolicy",
		reflect.TypeOf((*SubnetSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": SubnetSpecDeletionPolicy_ORPHAN,
			"DELETE": SubnetSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecForProvider",
		reflect.TypeOf((*SubnetSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecForProviderTags",
		reflect.TypeOf((*SubnetSpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdRef",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdRefPolicy",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdRefPolicyResolution",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SubnetSpecForProviderVpcIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": SubnetSpecForProviderVpcIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdRefPolicyResolve",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SubnetSpecForProviderVpcIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SubnetSpecForProviderVpcIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdSelector",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdSelectorPolicy",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdSelectorPolicyResolution",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SubnetSpecForProviderVpcIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": SubnetSpecForProviderVpcIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecForProviderVpcIdSelectorPolicyResolve",
		reflect.TypeOf((*SubnetSpecForProviderVpcIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SubnetSpecForProviderVpcIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SubnetSpecForProviderVpcIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecProviderConfigRef",
		reflect.TypeOf((*SubnetSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecProviderConfigRefPolicy",
		reflect.TypeOf((*SubnetSpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*SubnetSpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SubnetSpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": SubnetSpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*SubnetSpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SubnetSpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SubnetSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecProviderRef",
		reflect.TypeOf((*SubnetSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecProviderRefPolicy",
		reflect.TypeOf((*SubnetSpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecProviderRefPolicyResolution",
		reflect.TypeOf((*SubnetSpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SubnetSpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": SubnetSpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecProviderRefPolicyResolve",
		reflect.TypeOf((*SubnetSpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SubnetSpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SubnetSpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecPublishConnectionDetailsTo",
		reflect.TypeOf((*SubnetSpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*SubnetSpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*SubnetSpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*SubnetSpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SubnetSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*SubnetSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
}
