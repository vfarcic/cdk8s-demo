// databaseawscrossplaneio
package databaseawscrossplaneio


// IngestionRoleARNRef is a reference to a IAM Role used to set IngestionRoleARN.
type RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

