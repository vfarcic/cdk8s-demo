// ec2awscrossplaneio
package ec2awscrossplaneio


// A referencer to retrieve the ID of a NAT gateway.
type RouteTableSpecForProviderRoutesNatGatewayIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RouteTableSpecForProviderRoutesNatGatewayIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

