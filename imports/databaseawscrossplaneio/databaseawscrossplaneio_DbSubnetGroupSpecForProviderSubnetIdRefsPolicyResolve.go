// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve string

const (
	// Always.
	DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve_ALWAYS DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve = "ALWAYS"
	// IfNotPresent.
	DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve_IF_NOT_PRESENT DbSubnetGroupSpecForProviderSubnetIdRefsPolicyResolve = "IF_NOT_PRESENT"
)

