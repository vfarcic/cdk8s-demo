// ec2awscrossplaneio
package ec2awscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type VpcSpecDeletionPolicy string

const (
	// Orphan.
	VpcSpecDeletionPolicy_ORPHAN VpcSpecDeletionPolicy = "ORPHAN"
	// Delete.
	VpcSpecDeletionPolicy_DELETE VpcSpecDeletionPolicy = "DELETE"
)

