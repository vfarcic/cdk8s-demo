// ec2awscrossplaneio
package ec2awscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution string

const (
	// Required.
	SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution_REQUIRED SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution = "REQUIRED"
	// Optional.
	SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution_OPTIONAL SecurityGroupSpecForProviderIngressUserIdGroupPairsGroupIdRefPolicyResolution = "OPTIONAL"
)

