// ec2awscrossplaneio
package ec2awscrossplaneio


// UserIDGroupPair describes a security group and AWS account ID pair.
type SecurityGroupSpecForProviderEgressUserIdGroupPairs struct {
	// A description for the security group rule that references this user ID group pair.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the security group.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// GroupIDRef reference a security group to retrieve its GroupID.
	GroupIdRef *SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdRef `field:"optional" json:"groupIdRef" yaml:"groupIdRef"`
	// GroupIDSelector selects reference to a security group to retrieve its GroupID.
	GroupIdSelector *SecurityGroupSpecForProviderEgressUserIdGroupPairsGroupIdSelector `field:"optional" json:"groupIdSelector" yaml:"groupIdSelector"`
	// The name of the security group.
	//
	// In a request, use this parameter for a security group in EC2-Classic or a default VPC only. For a security group in a nondefault VPC, use the security group ID.
	// For a referenced security group in another VPC, this value is not returned if the referenced security group is deleted.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The ID of an AWS account.
	//
	// For a referenced security group in another VPC, the account ID of the referenced security group is returned in the response. If the referenced security group is deleted, this value is not returned.
	// [EC2-Classic] Required when adding or removing rules that reference a security group in another AWS account.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// The ID of the VPC for the referenced security group, if applicable.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// VPCIDRef reference a VPC to retrieve its vpcId.
	VpcIdRef *SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdRef `field:"optional" json:"vpcIdRef" yaml:"vpcIdRef"`
	// VPCIDSelector selects reference to a VPC to retrieve its vpcId.
	VpcIdSelector *SecurityGroupSpecForProviderEgressUserIdGroupPairsVpcIdSelector `field:"optional" json:"vpcIdSelector" yaml:"vpcIdSelector"`
	// The ID of the VPC peering connection, if applicable.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

