// ec2awscrossplaneio
package ec2awscrossplaneio


// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
type InternetGatewaySpecDeletionPolicy string

const (
	// Orphan.
	InternetGatewaySpecDeletionPolicy_ORPHAN InternetGatewaySpecDeletionPolicy = "ORPHAN"
	// Delete.
	InternetGatewaySpecDeletionPolicy_DELETE InternetGatewaySpecDeletionPolicy = "DELETE"
)

