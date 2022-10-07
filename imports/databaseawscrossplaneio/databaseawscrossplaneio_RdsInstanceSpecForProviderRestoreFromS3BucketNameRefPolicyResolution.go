// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution string

const (
	// Required.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution_REQUIRED RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution_OPTIONAL RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolution = "OPTIONAL"
)

