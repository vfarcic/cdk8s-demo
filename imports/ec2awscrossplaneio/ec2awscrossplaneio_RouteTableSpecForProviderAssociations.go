// ec2awscrossplaneio
package ec2awscrossplaneio


// Association describes an association between a route table and a subnet.
type RouteTableSpecForProviderAssociations struct {
	// The ID of the subnet.
	//
	// A subnet ID is not returned for an implicit association.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// A referencer to retrieve the ID of a subnet.
	SubnetIdRef *RouteTableSpecForProviderAssociationsSubnetIdRef `field:"optional" json:"subnetIdRef" yaml:"subnetIdRef"`
	// A selector to select a referencer to retrieve the ID of a subnet.
	SubnetIdSelector *RouteTableSpecForProviderAssociationsSubnetIdSelector `field:"optional" json:"subnetIdSelector" yaml:"subnetIdSelector"`
}

