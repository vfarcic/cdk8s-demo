// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution string

const (
	// Required.
	SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution_REQUIRED SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution = "REQUIRED"
	// Optional.
	SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution_OPTIONAL SecurityGroupSpecForProviderIngressUserIdGroupPairsVpcIdSelectorPolicyResolution = "OPTIONAL"
)

