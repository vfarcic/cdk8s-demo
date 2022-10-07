// ec2awscrossplaneio
package ec2awscrossplaneio


// SecurityGroupParameters define the desired state of an AWS VPC Security Group.
type SecurityGroupSpecForProvider struct {
	// A description of the security group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the security group.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// [EC2-VPC] One or more outbound rules associated with the security group.
	Egress *[]*SecurityGroupSpecForProviderEgress `field:"optional" json:"egress" yaml:"egress"`
	// One or more inbound rules associated with the security group.
	Ingress *[]*SecurityGroupSpecForProviderIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// Region is the region you'd like your SecurityGroup to be created in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags represents to current ec2 tags.
	Tags *[]*SecurityGroupSpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
	// VPCID is the ID of the VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// VPCIDRef references a VPC to and retrieves its vpcId.
	VpcIdRef *SecurityGroupSpecForProviderVpcIdRef `field:"optional" json:"vpcIdRef" yaml:"vpcIdRef"`
	// VPCIDSelector selects a reference to a VPC to and retrieves its vpcId.
	VpcIdSelector *SecurityGroupSpecForProviderVpcIdSelector `field:"optional" json:"vpcIdSelector" yaml:"vpcIdSelector"`
}

