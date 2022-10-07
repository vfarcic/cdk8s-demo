// databaseawscrossplaneio
package databaseawscrossplaneio

import (
	"time"
)

// PointInTime specifies the details of the point in time restore.
type RdsInstanceSpecForProviderRestoreFromPointInTime struct {
	// RestoreTime is the date and time (UTC) to restore from.
	//
	// Must be before the latest restorable time for the DB instance. Can't be specified if the useLatestRestorableTime parameter is enabled. Example: 2011-09-07T23:45:00Z
	RestoreTime *time.Time `field:"optional" json:"restoreTime" yaml:"restoreTime"`
	// SourceDBInstanceAutomatedBackupsArn specifies the Amazon Resource Name (ARN) of the replicated automated backups from which to restore.
	//
	// Example: arn:aws:rds:useast-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE.
	SourceDbInstanceAutomatedBackupsArn *string `field:"optional" json:"sourceDbInstanceAutomatedBackupsArn" yaml:"sourceDbInstanceAutomatedBackupsArn"`
	// SourceDBInstanceIdentifier specifies the identifier of the source DB instance from which to restore.
	//
	// Constraints: Must match the identifier of an existing DB instance.
	SourceDbInstanceIdentifier *string `field:"optional" json:"sourceDbInstanceIdentifier" yaml:"sourceDbInstanceIdentifier"`
	// SourceDbiResourceID specifies the resource ID of the source DB instance from which to restore.
	SourceDbiResourceId *string `field:"optional" json:"sourceDbiResourceId" yaml:"sourceDbiResourceId"`
	// UseLatestRestorableTime indicates that the DB instance is restored from the latest backup Can't be specified if the restoreTime parameter is provided.
	UseLatestRestorableTime *bool `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}

