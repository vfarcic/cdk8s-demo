// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve_ALWAYS RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnSelectorPolicyResolve = "IF_NOT_PRESENT"
)

