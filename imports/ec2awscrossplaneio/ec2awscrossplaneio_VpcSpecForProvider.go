// ec2awscrossplaneio
package ec2awscrossplaneio


// VPCParameters define the desired state of an AWS Virtual Private Cloud.
type VpcSpecForProvider struct {
	// CIDRBlock is the IPv4 network range for the VPC, in CIDR notation.
	//
	// For example, 10.0.0.0/16.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.
	//
	// You cannot specify the range of IP addresses, or the size of the CIDR block.
	AmazonProvidedIpv6CidrBlock *bool `field:"optional" json:"amazonProvidedIpv6CidrBlock" yaml:"amazonProvidedIpv6CidrBlock"`
	// Indicates whether the instances launched in the VPC get DNS hostnames.
	EnableDnsHostNames *bool `field:"optional" json:"enableDnsHostNames" yaml:"enableDnsHostNames"`
	// A boolean flag to enable/disable DNS support in the VPC.
	EnableDnsSupport *bool `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	// The allowed tenancy of instances launched into the VPC.
	InstanceTenancy *string `field:"optional" json:"instanceTenancy" yaml:"instanceTenancy"`
	// The IPv6 CIDR block from the IPv6 address pool.
	//
	// You must also specify Ipv6Pool in the request. To let Amazon choose the IPv6 CIDR block for you, omit this parameter.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.
	Ipv6Pool *string `field:"optional" json:"ipv6Pool" yaml:"ipv6Pool"`
	// Region is the region you'd like your VPC to be created in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags are used as identification helpers between AWS resources.
	Tags *[]*VpcSpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
}

