// databaseawscrossplaneio
package databaseawscrossplaneio


// Resolve specifies when this reference should be resolved.
//
// The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.
type DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve string

const (
	// Always.
	DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve_ALWAYS DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve = "ALWAYS"
	// IfNotPresent.
	DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve_IF_NOT_PRESENT DbSubnetGroupSpecForProviderSubnetIdSelectorPolicyResolve = "IF_NOT_PRESENT"
)

