// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve_ALWAYS RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve = "IF_NOT_PRESENT"
)

