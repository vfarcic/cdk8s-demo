// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution string

const (
	// Required.
	RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution_REQUIRED RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution_OPTIONAL RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolution = "OPTIONAL"
)

