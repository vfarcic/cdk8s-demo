// databaseawscrossplaneio
package databaseawscrossplaneio


// A Reference to a named object.
type DbSubnetGroupSpecForProviderSubnetIdRefs struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *DbSubnetGroupSpecForProviderSubnetIdRefsPolicy `field:"optional" json:"policy" yaml:"policy"`
}

