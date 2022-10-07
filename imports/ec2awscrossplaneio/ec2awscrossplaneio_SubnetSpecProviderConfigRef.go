// ec2awscrossplaneio
package ec2awscrossplaneio


// ProviderConfigReference specifies how the provider that will be used to create, observe, update, and delete this managed resource should be configured.
type SubnetSpecProviderConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *SubnetSpecProviderConfigRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

