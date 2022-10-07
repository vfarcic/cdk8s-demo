// ec2awscrossplaneio
package ec2awscrossplaneio


// A referencer to retrieve the ID of a subnet.
type RouteTableSpecForProviderAssociationsSubnetIdRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RouteTableSpecForProviderAssociationsSubnetIdRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

