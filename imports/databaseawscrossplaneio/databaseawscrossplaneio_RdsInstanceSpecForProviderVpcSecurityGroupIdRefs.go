// databaseawscrossplaneio
package databaseawscrossplaneio


// A Reference to a named object.
type RdsInstanceSpecForProviderVpcSecurityGroupIdRefs struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecForProviderVpcSecurityGroupIdRefsPolicy `field:"optional" json:"policy" yaml:"policy"`
}

