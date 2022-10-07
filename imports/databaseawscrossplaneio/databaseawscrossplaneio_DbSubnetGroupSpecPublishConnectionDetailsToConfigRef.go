// databaseawscrossplaneio
package databaseawscrossplaneio


// SecretStoreConfigRef specifies which secret store config should be used for this ConnectionSecret.
type DbSubnetGroupSpecPublishConnectionDetailsToConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

