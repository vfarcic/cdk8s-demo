// databaseawscrossplaneio
package databaseawscrossplaneio


// ProviderConfigReference specifies how the provider that will be used to create, observe, update, and delete this managed resource should be configured.
type RdsInstanceSpecProviderConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecProviderConfigRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

