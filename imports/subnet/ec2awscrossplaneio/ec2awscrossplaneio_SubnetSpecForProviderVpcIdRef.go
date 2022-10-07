// ec2awscrossplaneio
package ec2awscrossplaneio


// VPCIDRef reference a VPC to retrieve its vpcId.
type SubnetSpecForProviderVpcIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *SubnetSpecForProviderVpcIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

