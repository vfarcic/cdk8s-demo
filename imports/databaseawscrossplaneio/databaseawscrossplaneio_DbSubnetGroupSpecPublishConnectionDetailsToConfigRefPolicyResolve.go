// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve string

const (
	// Always.
	DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT DbSubnetGroupSpecPublishConnectionDetailsToConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

