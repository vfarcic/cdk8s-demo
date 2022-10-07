// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve_ALWAYS RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicyResolve = "IF_NOT_PRESENT"
)

