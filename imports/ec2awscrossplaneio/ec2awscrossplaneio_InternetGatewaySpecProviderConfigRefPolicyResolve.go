// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type InternetGatewaySpecProviderConfigRefPolicyResolve string

const (
	// Always.
	InternetGatewaySpecProviderConfigRefPolicyResolve_ALWAYS InternetGatewaySpecProviderConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	InternetGatewaySpecProviderConfigRefPolicyResolve_IF_NOT_PRESENT InternetGatewaySpecProviderConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

