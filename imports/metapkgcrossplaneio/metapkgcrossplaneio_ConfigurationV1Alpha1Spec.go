// metapkgcrossplaneio
package metapkgcrossplaneio


// ConfigurationSpec specifies the configuration of a Configuration.
type ConfigurationV1Alpha1Spec struct {
	// Semantic version constraints of Crossplane that package is compatible with.
	Crossplane *ConfigurationV1Alpha1SpecCrossplane `field:"optional" json:"crossplane" yaml:"crossplane"`
	// Dependencies on other packages.
	DependsOn *[]*ConfigurationV1Alpha1SpecDependsOn `field:"optional" json:"dependsOn" yaml:"dependsOn"`
}

