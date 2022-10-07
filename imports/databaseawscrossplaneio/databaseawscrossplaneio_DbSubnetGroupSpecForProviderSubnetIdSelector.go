// databaseawscrossplaneio
package databaseawscrossplaneio


// SubnetIDSelector selects a set of references that each retrieve the subnetID from the referenced Subnet.
type DbSubnetGroupSpecForProviderSubnetIdSelector struct {
	// MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.
	MatchControllerRef *bool `field:"optional" json:"matchControllerRef" yaml:"matchControllerRef"`
	// MatchLabels ensures an object with matching labels is selected.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
	// Policies for selection.
	Policy *DbSubnetGroupSpecForProviderSubnetIdSelectorPolicy `field:"optional" json:"policy" yaml:"policy"`
}

