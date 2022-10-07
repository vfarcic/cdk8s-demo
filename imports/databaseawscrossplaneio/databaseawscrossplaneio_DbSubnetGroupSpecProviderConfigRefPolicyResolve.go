// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type DbSubnetGroupSpecProviderConfigRefPolicyResolve string

const (
	// Always.
	DbSubnetGroupSpecProviderConfigRefPolicyResolve_ALWAYS DbSubnetGroupSpecProviderConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	DbSubnetGroupSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT DbSubnetGroupSpecProviderConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

