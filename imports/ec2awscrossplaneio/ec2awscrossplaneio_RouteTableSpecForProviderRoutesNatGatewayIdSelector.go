// ec2awscrossplaneio
package ec2awscrossplaneio


// A selector to select a referencer to retrieve the ID of a NAT gateway.
type RouteTableSpecForProviderRoutesNatGatewayIdSelector struct {
	// MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.
	MatchControllerRef *bool `field:"optional" json:"matchControllerRef" yaml:"matchControllerRef"`
	// MatchLabels ensures an object with matching labels is selected.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
	// Policies for selection.
	Policy *RouteTableSpecForProviderRoutesNatGatewayIdSelectorPolicy `field:"optional" json:"policy" yaml:"policy"`
}

