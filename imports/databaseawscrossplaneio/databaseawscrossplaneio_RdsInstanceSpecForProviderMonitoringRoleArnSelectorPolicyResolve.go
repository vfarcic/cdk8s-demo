// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve_ALWAYS RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderMonitoringRoleArnSelectorPolicyResolve = "IF_NOT_PRESENT"
)

