// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve string

const (
	// Always.
	RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT RdsInstanceSpecPublishConnectionDetailsToConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

