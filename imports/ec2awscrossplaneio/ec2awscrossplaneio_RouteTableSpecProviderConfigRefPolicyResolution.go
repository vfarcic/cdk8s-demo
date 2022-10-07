// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RouteTableSpecProviderConfigRefPolicyResolution string

const (
	// Required.
	RouteTableSpecProviderConfigRefPolicyResolution_REQUIRED RouteTableSpecProviderConfigRefPolicyResolution = "REQUIRED"
	// Optional.
	RouteTableSpecProviderConfigRefPolicyResolution_OPTIONAL RouteTableSpecProviderConfigRefPolicyResolution = "OPTIONAL"
)

