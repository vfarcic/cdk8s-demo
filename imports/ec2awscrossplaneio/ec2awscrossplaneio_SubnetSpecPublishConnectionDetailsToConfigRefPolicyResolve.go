// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve string

const (
	// Always.
	SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve_ALWAYS SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve_IF_NOT_PRESENT SubnetSpecPublishConnectionDetailsToConfigRefPolicyResolve = "IF_NOT_PRESENT"
)

