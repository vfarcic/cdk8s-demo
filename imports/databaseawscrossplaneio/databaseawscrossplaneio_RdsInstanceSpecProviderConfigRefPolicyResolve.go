// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecProviderConfigRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecProviderConfigRefPolicyResolve_ALWAYS RdsInstanceSpecProviderConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecProviderConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

