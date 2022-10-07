// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve_ALWAYS RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicyResolve = "IF_NOT_PRESENT"
)

