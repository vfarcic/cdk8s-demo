// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve_ALWAYS RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolve = "IF_NOT_PRESENT"
)

