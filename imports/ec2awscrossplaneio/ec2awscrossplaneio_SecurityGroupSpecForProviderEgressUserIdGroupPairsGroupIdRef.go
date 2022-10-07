// ec2awscrossplaneio
package ec2awscrossplaneio


// GroupIDRef reference a security group to retrieve its GroupID.
type SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

