// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type InternetGatewaySpecProviderConfigRefPolicyResolution string

const (
	// Required.
	InternetGatewaySpecProviderConfigRefPolicyResolution_REQUIRED InternetGatewaySpecProviderConfigRefPolicyResolution = "REQUIRED"
	// Optional.
	InternetGatewaySpecProviderConfigRefPolicyResolution_OPTIONAL InternetGatewaySpecProviderConfigRefPolicyResolution = "OPTIONAL"
)

