// databaseawscrossplaneio
package databaseawscrossplaneio


// DBSubnetGroupNameRef is a reference to a DBSubnetGroup used to set DBSubnetGroupName.
type RdsInstanceSpecForProviderDbSubnetGroupNameRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecForProviderDbSubnetGroupNameRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

