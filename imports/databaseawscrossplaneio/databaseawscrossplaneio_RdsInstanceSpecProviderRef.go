// databaseawscrossplaneio
package databaseawscrossplaneio


// ProviderReference specifies the provider that will be used to create, observe, update, and delete this managed resource.
//
// Deprecated: Please use ProviderConfigReference, i.e. `providerConfigRef`
type RdsInstanceSpecProviderRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecProviderRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

