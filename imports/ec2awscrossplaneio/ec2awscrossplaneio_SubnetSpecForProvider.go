// ec2awscrossplaneio
package ec2awscrossplaneio


// SubnetParameters define the desired state of an AWS VPC Subnet.
type SubnetSpecForProvider struct {
	// CIDRBlock is the IPv4 network range for the Subnet, in CIDR notation.
	//
	// For example, 10.0.0.0/18.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives an IPv6 address.
	AssignIpv6AddressOnCreation *bool `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// The Availability Zone for the subnet.
	//
	// Default: AWS selects one for you. If you create more than one subnet in your VPC, we may not necessarily select a different zone for each subnet.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID or the Local Zone ID of the subnet.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The IPv6 network range for the subnet, in CIDR notation.
	//
	// The subnet size must use a /64 prefix length.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Indicates whether instances launched in this subnet receive a public IPv4 address.
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Region is the region you'd like your Subnet to be created in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags represents to current ec2 tags.
	Tags *[]*SubnetSpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
	// VPCID is the ID of the VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// VPCIDRef reference a VPC to retrieve its vpcId.
	VpcIdRef *SubnetSpecForProviderVpcIdRef `field:"optional" json:"vpcIdRef" yaml:"vpcIdRef"`
	// VPCIDSelector selects reference to a VPC to retrieve its vpcId.
	VpcIdSelector *SubnetSpecForProviderVpcIdSelector `field:"optional" json:"vpcIdSelector" yaml:"vpcIdSelector"`
}

