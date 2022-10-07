// ec2awscrossplaneio
package ec2awscrossplaneio


// RouteTableParameters define the desired state of an AWS VPC Route Table.
type RouteTableSpecForProvider struct {
	// The associations between the route table and one or more subnets.
	Associations *[]*RouteTableSpecForProviderAssociations `field:"required" json:"associations" yaml:"associations"`
	// Region is the region you'd like your VPC to be created in.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Indicates whether we reconcile inline routes.
	IgnoreRoutes *bool `field:"optional" json:"ignoreRoutes" yaml:"ignoreRoutes"`
	// inline routes in the route table Deprecated: Routes inline exists for historical compatibility and should not be used.
	//
	// Please use separate route resource.
	Routes *[]*RouteTableSpecForProviderRoutes `field:"optional" json:"routes" yaml:"routes"`
	// Tags represents to current ec2 tags.
	Tags *[]*RouteTableSpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
	// VPCID is the ID of the VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// VPCIDRef references a VPC to retrieve its vpcId.
	VpcIdRef *RouteTableSpecForProviderVpcIdRef `field:"optional" json:"vpcIdRef" yaml:"vpcIdRef"`
	// VPCIDSelector selects a reference to a VPC to retrieve its vpcId.
	VpcIdSelector *RouteTableSpecForProviderVpcIdSelector `field:"optional" json:"vpcIdSelector" yaml:"vpcIdSelector"`
}

