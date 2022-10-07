// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
//
// The "Delete" policy is the default when no policy is specified.
type DatabaseSpecDeletionPolicy string

const (
	// Orphan.
	DatabaseSpecDeletionPolicy_ORPHAN DatabaseSpecDeletionPolicy = "ORPHAN"
	// Delete.
	DatabaseSpecDeletionPolicy_DELETE DatabaseSpecDeletionPolicy = "DELETE"
)

