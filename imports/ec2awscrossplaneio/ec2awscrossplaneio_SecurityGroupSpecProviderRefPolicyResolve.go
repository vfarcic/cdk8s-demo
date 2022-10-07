// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type SecurityGroupSpecProviderRefPolicyResolve string

const (
	// Always.
	SecurityGroupSpecProviderRefPolicyResolve_ALWAYS SecurityGroupSpecProviderRefPolicyResolve = "ALWAYS"
	// IfNotPresent.
	SecurityGroupSpecProviderRefPolicyResolve_IF_NOT_PRESENT SecurityGroupSpecProviderRefPolicyResolve = "IF_NOT_PRESENT"
)

