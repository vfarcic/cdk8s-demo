// ec2awscrossplaneio
package ec2awscrossplaneio


// IPv6Range describes an IPv6 range.
type SecurityGroupSpecForProviderIngressIpv6Ranges struct {
	// The IPv6 CIDR range.
	//
	// You can either specify a CIDR range or a source security group, not both. To specify a single IPv6 address, use the /128 prefix length.
	CidrIPv6 *string `field:"required" json:"cidrIPv6" yaml:"cidrIPv6"`
	// A description for the security group rule that references this IPv6 address range.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	Description *string `field:"optional" json:"description" yaml:"description"`
}

