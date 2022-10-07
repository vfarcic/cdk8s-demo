package ec2awscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"ec2awscrossplaneio.InternetGateway",
		reflect.TypeOf((*InternetGateway)(nil)).Elem(),
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
			j := jsiiProxy_InternetGateway{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewayProps",
		reflect.TypeOf((*InternetGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpec",
		reflect.TypeOf((*InternetGatewaySpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecDeletionPolicy",
		reflect.TypeOf((*InternetGatewaySpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": InternetGatewaySpecDeletionPolicy_ORPHAN,
			"DELETE": InternetGatewaySpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecForProvider",
		reflect.TypeOf((*InternetGatewaySpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderTags",
		reflect.TypeOf((*InternetGatewaySpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdRef",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdRefPolicy",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdRefPolicyResolution",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": InternetGatewaySpecForProviderVpcIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": InternetGatewaySpecForProviderVpcIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdRefPolicyResolve",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": InternetGatewaySpecForProviderVpcIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": InternetGatewaySpecForProviderVpcIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdSelector",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdSelectorPolicy",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdSelectorPolicyResolution",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": InternetGatewaySpecForProviderVpcIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": InternetGatewaySpecForProviderVpcIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve",
		reflect.TypeOf((*InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecProviderConfigRef",
		reflect.TypeOf((*InternetGatewaySpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecProviderConfigRefPolicy",
		reflect.TypeOf((*InternetGatewaySpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*InternetGatewaySpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": InternetGatewaySpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": InternetGatewaySpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*InternetGatewaySpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": InternetGatewaySpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": InternetGatewaySpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecProviderRef",
		reflect.TypeOf((*InternetGatewaySpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecProviderRefPolicy",
		reflect.TypeOf((*InternetGatewaySpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecProviderRefPolicyResolution",
		reflect.TypeOf((*InternetGatewaySpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": InternetGatewaySpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": InternetGatewaySpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecProviderRefPolicyResolve",
		reflect.TypeOf((*InternetGatewaySpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": InternetGatewaySpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": InternetGatewaySpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecPublishConnectionDetailsTo",
		reflect.TypeOf((*InternetGatewaySpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*InternetGatewaySpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*InternetGatewaySpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.InternetGatewaySpecWriteConnectionSecretToRef",
		reflect.TypeOf((*InternetGatewaySpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"ec2awscrossplaneio.RouteTable",
		reflect.TypeOf((*RouteTable)(nil)).Elem(),
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
			j := jsiiProxy_RouteTable{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableProps",
		reflect.TypeOf((*RouteTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpec",
		reflect.TypeOf((*RouteTableSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecDeletionPolicy",
		reflect.TypeOf((*RouteTableSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": RouteTableSpecDeletionPolicy_ORPHAN,
			"DELETE": RouteTableSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProvider",
		reflect.TypeOf((*RouteTableSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociations",
		reflect.TypeOf((*RouteTableSpecForProviderAssociations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdRef",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdRefPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderAssociationsSubnetIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdSelector",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderAssociationsSubnetIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutes",
		reflect.TypeOf((*RouteTableSpecForProviderRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdRef",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdRefPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdSelector",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdSelectorPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderRoutesGatewayIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdRef",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdRefPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderRoutesNatGatewayIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdSelector",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderTags",
		reflect.TypeOf((*RouteTableSpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdRef",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdRefPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderVpcIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderVpcIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderVpcIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderVpcIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdSelector",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdSelectorPolicy",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdSelectorPolicyResolution",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecForProviderVpcIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecForProviderVpcIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecForProviderVpcIdSelectorPolicyResolve",
		reflect.TypeOf((*RouteTableSpecForProviderVpcIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecForProviderVpcIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecForProviderVpcIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecProviderConfigRef",
		reflect.TypeOf((*RouteTableSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecProviderConfigRefPolicy",
		reflect.TypeOf((*RouteTableSpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecProviderRef",
		reflect.TypeOf((*RouteTableSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecProviderRefPolicy",
		reflect.TypeOf((*RouteTableSpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecProviderRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecProviderRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecPublishConnectionDetailsTo",
		reflect.TypeOf((*RouteTableSpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*RouteTableSpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*RouteTableSpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RouteTableSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*RouteTableSpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.RouteTableSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*RouteTableSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"ec2awscrossplaneio.SecurityGroup",
		reflect.TypeOf((*SecurityGroup)(nil)).Elem(),
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
			j := jsiiProxy_SecurityGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupProps",
		reflect.TypeOf((*SecurityGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpec",
		reflect.TypeOf((*SecurityGroupSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecDeletionPolicy",
		reflect.TypeOf((*SecurityGroupSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": SecurityGroupSpecDeletionPolicy_ORPHAN,
			"DELETE": SecurityGroupSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProvider",
		reflect.TypeOf((*SecurityGroupSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgress",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressIpRanges",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressIpv6Ranges",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressIpv6Ranges)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressPrefixListIds",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressPrefixListIds)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairs",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRef",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelector",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRef",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelector",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngress",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressIpRanges",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressIpv6Ranges",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressIpv6Ranges)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressPrefixListIds",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressPrefixListIds)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairs",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRef",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelector",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRef",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelector",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderTags",
		reflect.TypeOf((*SecurityGroupSpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdRef",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderVpcIdRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderVpcIdRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderVpcIdRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderVpcIdRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdSelector",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdSelectorPolicy",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdSelectorPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecForProviderVpcIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecForProviderVpcIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecForProviderVpcIdSelectorPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecForProviderVpcIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecForProviderVpcIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecForProviderVpcIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecProviderConfigRef",
		reflect.TypeOf((*SecurityGroupSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecProviderConfigRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecProviderRef",
		reflect.TypeOf((*SecurityGroupSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecProviderRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecProviderRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecProviderRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecPublishConnectionDetailsTo",
		reflect.TypeOf((*SecurityGroupSpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*SecurityGroupSpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"ec2awscrossplaneio.SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*SecurityGroupSpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ec2awscrossplaneio.SecurityGroupSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*SecurityGroupSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
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
