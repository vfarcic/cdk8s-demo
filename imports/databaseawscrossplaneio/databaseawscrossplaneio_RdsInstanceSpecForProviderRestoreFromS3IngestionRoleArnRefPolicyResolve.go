// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve_ALWAYS RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecForProviderRestoreFromS3IngestionRoleArnRefPolicyResolve = "IF_NOT_PRESENT"
)

