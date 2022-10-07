// ec2awscrossplaneio
package ec2awscrossplaneio


// IPPermission Describes a set of permissions for a security group rule.
type SecurityGroupSpecForProviderEgress struct {
	// The IP protocol name (tcp, udp, icmp, icmpv6) or number (see Protocol Numbers (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)). [VPC only] Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.
	//
	// A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The IPv4 ranges.
	IpRanges *[]*SecurityGroupSpecForProviderEgressIpRanges `field:"optional" json:"ipRanges" yaml:"ipRanges"`
	// The IPv6 ranges.
	//
	// [VPC only].
	Ipv6Ranges *[]*SecurityGroupSpecForProviderEgressIpv6Ranges `field:"optional" json:"ipv6Ranges" yaml:"ipv6Ranges"`
	// PrefixListIDs for an AWS service.
	//
	// With outbound rules, this is the AWS service to access through a VPC endpoint from instances associated with the security group.
	// [VPC only].
	PrefixListIds *[]*SecurityGroupSpecForProviderEgressPrefixListIds `field:"optional" json:"prefixListIds" yaml:"prefixListIds"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	//
	// A value of -1 indicates all ICMP/ICMPv6 codes. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
	// UserIDGroupPairs are the source security group and AWS account ID pairs.
	//
	// It contains one or more accounts and security groups to allow flows from security groups of other accounts.
	UserIdGroupPairs *[]*SecurityGroupSpecForProviderEgressUserIdGroupPairs `field:"optional" json:"userIdGroupPairs" yaml:"userIdGroupPairs"`
}

