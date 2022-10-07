// databaseawscrossplaneio
package databaseawscrossplaneio


// Snapshot specifies the details of the database snapshot to restore from.
type RdsInstanceSpecForProviderRestoreFromSnapshot struct {
	// SnapshotIdentifier is the identifier of the database snapshot to restore.
	SnapshotIdentifier *string `field:"required" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
}

