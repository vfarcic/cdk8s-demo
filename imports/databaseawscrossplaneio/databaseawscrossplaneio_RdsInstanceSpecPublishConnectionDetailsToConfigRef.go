// databaseawscrossplaneio
package databaseawscrossplaneio


// SecretStoreConfigRef specifies which secret store config should be used for this ConnectionSecret.
type RdsInstanceSpecPublishConnectionDetailsToConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

