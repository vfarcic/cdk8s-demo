// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution string

const (
	// Required.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution_REQUIRED RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution_OPTIONAL RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolution = "OPTIONAL"
)

