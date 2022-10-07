// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RouteTableSpecProviderConfigRefPolicyResolve string

const (
	// Always.
	RouteTableSpecProviderConfigRefPolicyResolve_ALWAYS RouteTableSpecProviderConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RouteTableSpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT RouteTableSpecProviderConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

