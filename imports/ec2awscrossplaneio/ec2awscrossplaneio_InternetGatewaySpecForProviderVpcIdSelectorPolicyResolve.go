// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve string

const (
	// Always.
	InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve_ALWAYS InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve_IF_NOT_PRESENT InternetGatewaySpecForProviderVpcIdSelectorPolicyResolve = "IF_NOT_PRESENT"
)

