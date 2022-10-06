// metapkgcrossplaneio
package metapkgcrossplaneio


// ProviderSpec specifies the configuration of a Provider.
type ProviderSpec struct {
	// Configuration for the packaged Provider's controller.
	Controller *ProviderSpecController `field:"required" json:"controller" yaml:"controller"`
	// Semantic version constraints of Crossplane that package is compatible with.
	Crossplane *ProviderSpecCrossplane `field:"optional" json:"crossplane" yaml:"crossplane"`
	// Dependencies on other packages.
	DependsOn *[]*ProviderSpecDependsOn `field:"optional" json:"dependsOn" yaml:"dependsOn"`
}

