// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve_ALWAYS RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicyResolve = "IF_NOT_PRESENT"
)

