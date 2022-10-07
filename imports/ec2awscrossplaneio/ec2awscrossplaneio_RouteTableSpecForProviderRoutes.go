// ec2awscrossplaneio
package ec2awscrossplaneio


// RouteBeta describes a route in a route table.
//
// provider-aws currently provides both a standalone Route resource and a RouteTable resource with routes defined in-line. At this time you cannot use a Route Table with in-line routes in conjunction with any Route resources. Doing so will cause a conflict of rule settings and will overwrite rules.
type RouteTableSpecForProviderRoutes struct {
	// The IPv4 CIDR address block used for the destination match.
	//
	// Routing decisions are based on the most specific match.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The IPv6 CIDR address block used for the destination match.
	//
	// Routing decisions are based on the most specific match.
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayId *string `field:"optional" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
	// The ID of an internet gateway or virtual private gateway attached to your VPC.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// A referencer to retrieve the ID of a gateway.
	GatewayIdRef *RouteTableSpecForProviderRoutesGatewayIdRef `field:"optional" json:"gatewayIdRef" yaml:"gatewayIdRef"`
	// A selector to select a referencer to retrieve the ID of a gateway.
	GatewayIdSelector *RouteTableSpecForProviderRoutesGatewayIdSelector `field:"optional" json:"gatewayIdSelector" yaml:"gatewayIdSelector"`
	// The ID of a NAT instance in your VPC.
	//
	// The operation fails if you specify an instance ID unless exactly one network interface is attached.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the local gateway.
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// [IPv4 traffic only] The ID of a NAT gateway.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// A referencer to retrieve the ID of a NAT gateway.
	NatGatewayIdRef *RouteTableSpecForProviderRoutesNatGatewayIdRef `field:"optional" json:"natGatewayIdRef" yaml:"natGatewayIdRef"`
	// A selector to select a referencer to retrieve the ID of a NAT gateway.
	NatGatewayIdSelector *RouteTableSpecForProviderRoutesNatGatewayIdSelector `field:"optional" json:"natGatewayIdSelector" yaml:"natGatewayIdSelector"`
	// The ID of a network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of a transit gateway.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

