// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve_ALWAYS RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolve = "IF_NOT_PRESENT"
)

