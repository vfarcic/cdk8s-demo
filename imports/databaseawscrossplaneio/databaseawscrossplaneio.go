package databaseawscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"databaseawscrossplaneio.DbSubnetGroup",
		reflect.TypeOf((*DbSubnetGroup)(nil)).Elem(),
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
			j := jsiiProxy_DbSubnetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupProps",
		reflect.TypeOf((*DbSubnetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpec",
		reflect.TypeOf((*DbSubnetGroupSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecDeletionPolicy",
		reflect.TypeOf((*DbSubnetGroupSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": DbSubnetGroupSpecDeletionPolicy_ORPHAN,
			"DELETE": DbSubnetGroupSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProvider",
		reflect.TypeOf((*DbSubnetGroupSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdRefs",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdRefsPolicy",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdRefsPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolution",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolution_REQUIRED,
			"OPTIONAL": DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdSelector",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdSelectorPolicy",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolution",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecForProviderTags",
		reflect.TypeOf((*DbSubnetGroupSpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderConfigRef",
		reflect.TypeOf((*DbSubnetGroupSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderConfigRefPolicy",
		reflect.TypeOf((*DbSubnetGroupSpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*DbSubnetGroupSpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": DbSubnetGroupSpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": DbSubnetGroupSpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*DbSubnetGroupSpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": DbSubnetGroupSpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": DbSubnetGroupSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderRef",
		reflect.TypeOf((*DbSubnetGroupSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderRefPolicy",
		reflect.TypeOf((*DbSubnetGroupSpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderRefPolicyResolution",
		reflect.TypeOf((*DbSubnetGroupSpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": DbSubnetGroupSpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": DbSubnetGroupSpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecProviderRefPolicyResolve",
		reflect.TypeOf((*DbSubnetGroupSpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": DbSubnetGroupSpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": DbSubnetGroupSpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecPublishConnectionDetailsTo",
		reflect.TypeOf((*DbSubnetGroupSpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*DbSubnetGroupSpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*DbSubnetGroupSpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.DbSubnetGroupSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*DbSubnetGroupSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"databaseawscrossplaneio.RdsInstance",
		reflect.TypeOf((*RdsInstance)(nil)).Elem(),
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
			j := jsiiProxy_RdsInstance{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceProps",
		reflect.TypeOf((*RdsInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpec",
		reflect.TypeOf((*RdsInstanceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecDeletionPolicy",
		reflect.TypeOf((*RdsInstanceSpecDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORPHAN": RdsInstanceSpecDeletionPolicy_ORPHAN,
			"DELETE": RdsInstanceSpecDeletionPolicy_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProvider",
		reflect.TypeOf((*RdsInstanceSpecForProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderCloudwatchLogsExportConfiguration",
		reflect.TypeOf((*RdsInstanceSpecForProviderCloudwatchLogsExportConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameRef",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameSelector",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameRef",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameSelector",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMasterPasswordSecretRef",
		reflect.TypeOf((*RdsInstanceSpecForProviderMasterPasswordSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnRef",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnSelector",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderProcessorFeatures",
		reflect.TypeOf((*RdsInstanceSpecForProviderProcessorFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFrom",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromPointInTime",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromPointInTime)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameRef",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameSelector",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRef",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelector",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromSnapshot",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromSnapshot)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderRestoreFromSource",
		reflect.TypeOf((*RdsInstanceSpecForProviderRestoreFromSource)(nil)).Elem(),
		map[string]interface{}{
			"S3": RdsInstanceSpecForProviderRestoreFromSource_S3,
			"SNAPSHOT": RdsInstanceSpecForProviderRestoreFromSource_SNAPSHOT,
			"POINT_IN_TIME": RdsInstanceSpecForProviderRestoreFromSource_POINT_IN_TIME,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderScalingConfiguration",
		reflect.TypeOf((*RdsInstanceSpecForProviderScalingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderTags",
		reflect.TypeOf((*RdsInstanceSpecForProviderTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdRefs",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdSelector",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicy",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecProviderConfigRef",
		reflect.TypeOf((*RdsInstanceSpecProviderConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecProviderConfigRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecProviderConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecProviderConfigRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecProviderConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecProviderConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecProviderConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecProviderConfigRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecProviderConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecProviderConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecProviderRef",
		reflect.TypeOf((*RdsInstanceSpecProviderRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecProviderRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecProviderRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecProviderRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecProviderRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecProviderRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecProviderRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecProviderRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecProviderRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecProviderRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecProviderRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecPublishConnectionDetailsTo",
		reflect.TypeOf((*RdsInstanceSpecPublishConnectionDetailsTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecPublishConnectionDetailsToConfigRef",
		reflect.TypeOf((*RdsInstanceSpecPublishConnectionDetailsToConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicy",
		reflect.TypeOf((*RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution",
		reflect.TypeOf((*RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED,
			"OPTIONAL": RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL,
		},
	)
	_jsii_.RegisterEnum(
		"databaseawscrossplaneio.RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve",
		reflect.TypeOf((*RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS,
			"IF_NOT_PRESENT": RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecPublishConnectionDetailsToMetadata",
		reflect.TypeOf((*RdsInstanceSpecPublishConnectionDetailsToMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"databaseawscrossplaneio.RdsInstanceSpecWriteConnectionSecretToRef",
		reflect.TypeOf((*RdsInstanceSpecWriteConnectionSecretToRef)(nil)).Elem(),
	)
}
