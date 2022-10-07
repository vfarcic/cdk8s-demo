// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve string

const (
	// Always.
	InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT InternetGatewaySpecPublishConnectionDetailsToConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

