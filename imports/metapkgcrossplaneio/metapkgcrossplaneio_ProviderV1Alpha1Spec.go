// metapkgcrossplaneio
package metapkgcrossplaneio


// ProviderSpec specifies the configuration of a Provider.
type ProviderV1Alpha1Spec struct {
	// Configuration for the packaged Provider's controller.
	Controller *ProviderV1Alpha1SpecController `field:"required" json:"controller" yaml:"controller"`
	// Semantic version constraints of Crossplane that package is compatible with.
	Crossplane *ProviderV1Alpha1SpecCrossplane `field:"optional" json:"crossplane" yaml:"crossplane"`
	// Dependencies on other packages.
	DependsOn *[]*ProviderV1Alpha1SpecDependsOn `field:"optional" json:"dependsOn" yaml:"dependsOn"`
}

