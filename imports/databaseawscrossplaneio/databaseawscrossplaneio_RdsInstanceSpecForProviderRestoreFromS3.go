// databaseawscrossplaneio
package databaseawscrossplaneio


// S3 specifies the details of the S3 backup to restore from.
type RdsInstanceSpecForProviderRestoreFromS3 struct {
	// SourceEngine is the engine used to create the backup.
	//
	// Must be "mysql".
	SourceEngine *string `field:"required" json:"sourceEngine" yaml:"sourceEngine"`
	// SourceEngineVersion is the version of the engine used to create the backup.
	//
	// Example: "5.7.30"
	SourceEngineVersion *string `field:"required" json:"sourceEngineVersion" yaml:"sourceEngineVersion"`
	// BucketName is the name of the S3 bucket containing the backup to restore.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// BucketNameRef is a reference to a Bucket used to set BucketName.
	BucketNameRef *RdsInstanceSpecForProviderRestoreFromS3BucketNameRef `field:"optional" json:"bucketNameRef" yaml:"bucketNameRef"`
	// BucketNameSelector selects a reference to a Bucket used to set BucketName.
	BucketNameSelector *RdsInstanceSpecForProviderRestoreFromS3BucketNameSelector `field:"optional" json:"bucketNameSelector" yaml:"bucketNameSelector"`
	// IngestionRoleARN is the IAM role RDS can assume that will allow it to access the contents of the S3 bucket.
	IngestionRoleArn *string `field:"optional" json:"ingestionRoleArn" yaml:"ingestionRoleArn"`
	// IngestionRoleARNRef is a reference to a IAM Role used to set IngestionRoleARN.
	IngestionRoleArnRef *RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRef `field:"optional" json:"ingestionRoleArnRef" yaml:"ingestionRoleArnRef"`
	// IngestionRoleARNSelector selects a reference to a IAM Role used to set IngestionRoleARN.
	IngestionRoleArnSelector *RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelector `field:"optional" json:"ingestionRoleArnSelector" yaml:"ingestionRoleArnSelector"`
	// Prefix is the path prefix of the S3 bucket within which the backup to restore is located.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

