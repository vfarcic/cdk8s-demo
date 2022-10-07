// databaseawscrossplaneio
package databaseawscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type DbSubnetGroupSpecDeletionPolicy string

const (
	// Orphan.
	DbSubnetGroupSpecDeletionPolicy_ORPHAN DbSubnetGroupSpecDeletionPolicy = "ORPHAN"
	// Delete.
	DbSubnetGroupSpecDeletionPolicy_DELETE DbSubnetGroupSpecDeletionPolicy = "DELETE"
)

