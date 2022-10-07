// databaseawscrossplaneio
package databaseawscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type RdsInstanceSpecDeletionPolicy string

const (
	// Orphan.
	RdsInstanceSpecDeletionPolicy_ORPHAN RdsInstanceSpecDeletionPolicy = "ORPHAN"
	// Delete.
	RdsInstanceSpecDeletionPolicy_DELETE RdsInstanceSpecDeletionPolicy = "DELETE"
)

