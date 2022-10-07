// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution string

const (
	// Required.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution_REQUIRED RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution_OPTIONAL RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolution = "OPTIONAL"
)

