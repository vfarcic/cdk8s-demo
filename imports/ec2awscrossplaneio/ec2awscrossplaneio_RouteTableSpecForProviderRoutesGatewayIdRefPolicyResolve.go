// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve string

const (
	// Always.
	RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve_ALWAYS RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve_IF_NOT_PRESENT RouteTableSpecForProviderRoutesGatewayIdRefPolicyResolve = "IF_NOT_PRESENT"
)

