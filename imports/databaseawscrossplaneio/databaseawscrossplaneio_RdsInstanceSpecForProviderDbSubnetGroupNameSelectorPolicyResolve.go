// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve_ALWAYS RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderDbSubnetGroupNameSelectorPolicyResolve = "IF_NOT_PRESENT"
)

