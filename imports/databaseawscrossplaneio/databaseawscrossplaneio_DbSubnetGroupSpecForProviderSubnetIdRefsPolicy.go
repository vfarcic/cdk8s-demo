// databaseawscrossplaneio
package databaseawscrossplaneio


// Policies for referencing.
type DbSubnetGroupSpecForProviderSubnetIdRefsPolicy struct {
	// Resolution specifies whether resolution of this reference is required.
	//
	// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
	Resolution DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolution `field:"optional" json:"resolution" yaml:"resolution"`
	// Resolve specifies when this reference should be resolved.
	//
	// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
	Resolve DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve `field:"optional" json:"resolve" yaml:"resolve"`
}

