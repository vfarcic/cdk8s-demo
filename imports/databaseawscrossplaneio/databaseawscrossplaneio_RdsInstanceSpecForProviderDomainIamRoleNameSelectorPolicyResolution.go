// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution string

const (
	// Required.
	RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution_REQUIRED RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution = "REQUIRED"
	// Optional.
	RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution_OPTIONAL RdsInstanceSpecForProviderDomainIamRoleNameSelectorPolicyResolution = "OPTIONAL"
)

