// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecProviderRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecProviderRefPolicyResolve_ALWAYS RdsInstanceSpecProviderRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecProviderRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecProviderRefPolicyResolve = "IF_NOT_PRESENT"
)

