// databaseawscrossplaneio
package databaseawscrossplaneio


// DomainIAMRoleNameRef is a reference to an IAMRole used to set DomainIAMRoleName.
type RdsInstanceSpecForProviderDomainIamRoleNameRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecForProviderDomainIamRoleNameRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

