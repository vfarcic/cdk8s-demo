// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type DbSubnetGroupSpecProviderRefPolicyResolve string

const (
	// Always.
	DbSubnetGroupSpecProviderRefPolicyResolve_ALWAYS DbSubnetGroupSpecProviderRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	DbSubnetGroupSpecProviderRefPolicyResolve_IF_NOT_PRESENT DbSubnetGroupSpecProviderRefPolicyResolve = "IF_NOT_PRESENT"
)

