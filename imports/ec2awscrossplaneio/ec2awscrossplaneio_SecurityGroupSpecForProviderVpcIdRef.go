// ec2awscrossplaneio
package ec2awscrossplaneio


// VPCIDRef references a VPC to and retrieves its vpcId.
type SecurityGroupSpecForProviderVpcIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *SecurityGroupSpecForProviderVpcIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

