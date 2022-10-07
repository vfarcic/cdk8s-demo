// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve_ALWAYS RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicyResolve = "IF_NOT_PRESENT"
)

