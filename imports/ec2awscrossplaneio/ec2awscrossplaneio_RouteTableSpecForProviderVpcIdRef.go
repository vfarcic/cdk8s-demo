// ec2awscrossplaneio
package ec2awscrossplaneio


// VPCIDRef references a VPC to retrieve its vpcId.
type RouteTableSpecForProviderVpcIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RouteTableSpecForProviderVpcIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

