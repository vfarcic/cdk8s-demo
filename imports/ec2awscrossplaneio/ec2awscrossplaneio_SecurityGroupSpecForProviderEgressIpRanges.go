// ec2awscrossplaneio
package ec2awscrossplaneio


// IPRange describes an IPv4 range.
type SecurityGroupSpecForProviderEgressIpRanges struct {
	// The IPv4 CIDR range.
	//
	// You can either specify a CIDR range or a source security group, not both. To specify a single IPv4 address, use the /32 prefix length.
	CidrIp *string `field:"required" json:"cidrIp" yaml:"cidrIp"`
	// A description for the security group rule that references this IPv4 address range.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	Description *string `field:"optional" json:"description" yaml:"description"`
}

