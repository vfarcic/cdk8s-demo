// databaseawscrossplaneio
package databaseawscrossplaneio


// VPCSecurityGroupIDSelector selects references to VPCSecurityGroups used to set the VPCSecurityGroupIDs.
type RdsInstanceSpecForProviderVpcSecurityGroupIdSelector struct {
	// MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.
	MatchControllerRef *bool `field:"optional" json:"matchControllerRef" yaml:"matchControllerRef"`
	// MatchLabels ensures an object with matching labels is selected.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
	// Policies for selection.
	Policy *RdsInstanceSpecForProviderVpcSecurityGroupIdSelectorPolicy `field:"optional" json:"policy" yaml:"policy"`
}

