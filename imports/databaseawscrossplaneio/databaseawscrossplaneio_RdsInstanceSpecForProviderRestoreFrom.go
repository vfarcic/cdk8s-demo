// databaseawscrossplaneio
package databaseawscrossplaneio


// RestoreFrom specifies the details of the backup to restore when creating a new RDS instance.
//
// (If the RDS instance already exists, this property will be ignored.)
type RdsInstanceSpecForProviderRestoreFrom struct {
	// Source is the type of the backup to restore when creating a new RDS instance.
	//
	// S3, Snapshot and PointInTime are supported.
	Source RdsInstanceSpecForProviderRestoreFromSource `field:"required" json:"source" yaml:"source"`
	// PointInTime specifies the details of the point in time restore.
	PointInTime *RdsInstanceSpecForProviderRestoreFromPointInTime `field:"optional" json:"pointInTime" yaml:"pointInTime"`
	// S3 specifies the details of the S3 backup to restore from.
	S3 *RdsInstanceSpecForProviderRestoreFromS3 `field:"optional" json:"s3" yaml:"s3"`
	// Snapshot specifies the details of the database snapshot to restore from.
	Snapshot *RdsInstanceSpecForProviderRestoreFromSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
}

