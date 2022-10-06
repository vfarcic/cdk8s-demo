package pkgcrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Configuration",
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
		"pkgcrossplaneio.ConfigurationProps",
		reflect.TypeOf((*ConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ConfigurationRevision",
		reflect.TypeOf((*ConfigurationRevision)(nil)).Elem(),
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
			j := jsiiProxy_ConfigurationRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionProps",
		reflect.TypeOf((*ConfigurationRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionSpec",
		reflect.TypeOf((*ConfigurationRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionSpecControllerConfigRef",
		reflect.TypeOf((*ConfigurationRevisionSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionSpecPackagePullSecrets",
		reflect.TypeOf((*ConfigurationRevisionSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationSpec",
		reflect.TypeOf((*ConfigurationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationSpecPackagePullSecrets",
		reflect.TypeOf((*ConfigurationSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfig",
		reflect.TypeOf((*ControllerConfig)(nil)).Elem(),
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
			j := jsiiProxy_ControllerConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigProps",
		reflect.TypeOf((*ControllerConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpec",
		reflect.TypeOf((*ControllerConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnv",
		reflect.TypeOf((*ControllerConfigSpecEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvFrom",
		reflect.TypeOf((*ControllerConfigSpecEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvFromConfigMapRef",
		reflect.TypeOf((*ControllerConfigSpecEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvFromSecretRef",
		reflect.TypeOf((*ControllerConfigSpecEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFrom",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromFieldRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromResourceFieldRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromSecretKeyRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecImagePullSecrets",
		reflect.TypeOf((*ControllerConfigSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecMetadata",
		reflect.TypeOf((*ControllerConfigSpecMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContext",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextSeLinuxOptions",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextSeccompProfile",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextSysctls",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextWindowsOptions",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPorts",
		reflect.TypeOf((*ControllerConfigSpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecResources",
		reflect.TypeOf((*ControllerConfigSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecResourcesLimits",
		reflect.TypeOf((*ControllerConfigSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecResourcesRequests",
		reflect.TypeOf((*ControllerConfigSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContext",
		reflect.TypeOf((*ControllerConfigSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextCapabilities",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecTolerations",
		reflect.TypeOf((*ControllerConfigSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Lock",
		reflect.TypeOf((*Lock)(nil)).Elem(),
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
			j := jsiiProxy_Lock{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockPackages",
		reflect.TypeOf((*LockPackages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockPackagesDependencies",
		reflect.TypeOf((*LockPackagesDependencies)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockProps",
		reflect.TypeOf((*LockProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.LockV1Beta1",
		reflect.TypeOf((*LockV1Beta1)(nil)).Elem(),
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
			j := jsiiProxy_LockV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockV1Beta1Packages",
		reflect.TypeOf((*LockV1Beta1Packages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockV1Beta1PackagesDependencies",
		reflect.TypeOf((*LockV1Beta1PackagesDependencies)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockV1Beta1Props",
		reflect.TypeOf((*LockV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Provider",
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
		"pkgcrossplaneio.ProviderProps",
		reflect.TypeOf((*ProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ProviderRevision",
		reflect.TypeOf((*ProviderRevision)(nil)).Elem(),
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
			j := jsiiProxy_ProviderRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionProps",
		reflect.TypeOf((*ProviderRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpec",
		reflect.TypeOf((*ProviderRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpecControllerConfigRef",
		reflect.TypeOf((*ProviderRevisionSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpecPackagePullSecrets",
		reflect.TypeOf((*ProviderRevisionSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpec",
		reflect.TypeOf((*ProviderSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpecControllerConfigRef",
		reflect.TypeOf((*ProviderSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpecPackagePullSecrets",
		reflect.TypeOf((*ProviderSpecPackagePullSecrets)(nil)).Elem(),
	)
}
