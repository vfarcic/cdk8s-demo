// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution string

const (
	// Required.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution_REQUIRED RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution_OPTIONAL RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolution = "OPTIONAL"
)

