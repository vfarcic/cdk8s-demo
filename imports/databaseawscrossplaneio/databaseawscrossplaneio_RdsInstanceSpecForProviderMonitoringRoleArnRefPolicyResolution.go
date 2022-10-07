// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution string

const (
	// Required.
	RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution_REQUIRED RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution_OPTIONAL RdsInstanceSpecForProviderMonitoringRoleArnRefPolicyResolution = "OPTIONAL"
)

