// ec2awscrossplaneio
package ec2awscrossplaneio


// InternetGatewayParameters define the desired state of an AWS VPC Internet Gateway.
type InternetGatewaySpecForProvider struct {
	// Region is the region you'd like your VPC to be created in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags represents to current ec2 tags.
	Tags *[]*InternetGatewaySpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
	// VPCID is the ID of the VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// VPCIDRef references a VPC to and retrieves its vpcId.
	VpcIdRef *InternetGatewaySpecForProviderVpcIdRef `field:"optional" json:"vpcIdRef" yaml:"vpcIdRef"`
	// VPCIDSelector selects a reference to a VPC to and retrieves its vpcId.
	VpcIdSelector *InternetGatewaySpecForProviderVpcIdSelector `field:"optional" json:"vpcIdSelector" yaml:"vpcIdSelector"`
}

