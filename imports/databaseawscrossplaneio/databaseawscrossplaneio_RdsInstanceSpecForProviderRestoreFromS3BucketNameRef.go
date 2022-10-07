// databaseawscrossplaneio
package databaseawscrossplaneio


// BucketNameRef is a reference to a Bucket used to set BucketName.
type RdsInstanceSpecForProviderRestoreFromS3BucketNameRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

