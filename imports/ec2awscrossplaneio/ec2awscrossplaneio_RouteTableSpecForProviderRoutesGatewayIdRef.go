// ec2awscrossplaneio
package ec2awscrossplaneio


// A referencer to retrieve the ID of a gateway.
type RouteTableSpecForProviderRoutesGatewayIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RouteTableSpecForProviderRoutesGatewayIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

