// ec2awscrossplaneio
package ec2awscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type SecurityGroupSpecDeletionPolicy string

const (
	// Orphan.
	SecurityGroupSpecDeletionPolicy_ORPHAN SecurityGroupSpecDeletionPolicy = "ORPHAN"
	// Delete.
	SecurityGroupSpecDeletionPolicy_DELETE SecurityGroupSpecDeletionPolicy = "DELETE"
)

