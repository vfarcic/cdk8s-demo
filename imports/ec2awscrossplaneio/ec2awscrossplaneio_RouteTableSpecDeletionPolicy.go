// ec2awscrossplaneio
package ec2awscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type RouteTableSpecDeletionPolicy string

const (
	// Orphan.
	RouteTableSpecDeletionPolicy_ORPHAN RouteTableSpecDeletionPolicy = "ORPHAN"
	// Delete.
	RouteTableSpecDeletionPolicy_DELETE RouteTableSpecDeletionPolicy = "DELETE"
)

