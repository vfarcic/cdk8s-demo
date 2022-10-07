// ec2awscrossplaneio
package ec2awscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type SubnetSpecDeletionPolicy string

const (
	// Orphan.
	SubnetSpecDeletionPolicy_ORPHAN SubnetSpecDeletionPolicy = "ORPHAN"
	// Delete.
	SubnetSpecDeletionPolicy_DELETE SubnetSpecDeletionPolicy = "DELETE"
)

