// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve_ALWAYS RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderRestoreFromS3BucketNameRefPolicyResolve = "IF_NOT_PRESENT"
)

