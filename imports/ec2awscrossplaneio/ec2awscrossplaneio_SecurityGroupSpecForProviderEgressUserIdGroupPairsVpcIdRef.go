// ec2awscrossplaneio
package ec2awscrossplaneio


// VPCIDRef reference a VPC to retrieve its vpcId.
type SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

