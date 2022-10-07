// databaseawscrossplaneio
package databaseawscrossplaneio


// Source is the type of the backup to restore when creating a new RDS instance.
//
// S3, Snapshot and PointInTime are supported.
type RdsInstanceSpecForProviderRestoreFromSource string

const (
	// S3.
	RdsInstanceSpecForProviderRestoreFromSource_S3 RdsInstanceSpecForProviderRestoreFromSource = "S3"
	// Snapshot.
	RdsInstanceSpecForProviderRestoreFromSource_SNAPSHOT RdsInstanceSpecForProviderRestoreFromSource = "SNAPSHOT"
	// PointInTime.
	RdsInstanceSpecForProviderRestoreFromSource_POINT_IN_TIME RdsInstanceSpecForProviderRestoreFromSource = "POINT_IN_TIME"
)

