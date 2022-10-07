// databaseawscrossplaneio
package databaseawscrossplaneio


// Policies for referencing.
type RdsInstanceSpecForProviderDomainIamRoleNameRefPolicy struct {
	// Resolution specifies whether resolution of this reference is required.
	//
	// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
	Resolution RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolution `field:"optional" json:"resolution" yaml:"resolution"`
	// Resolve specifies when this reference should be resolved.
	//
	// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
	Resolve RdsInstanceSpecForProviderDomainIamRoleNameRefPolicyResolve `field:"optional" json:"resolve" yaml:"resolve"`
}
