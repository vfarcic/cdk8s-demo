// metapkgcrossplaneio
package metapkgcrossplaneio


// ConfigurationSpec specifies the configuration of a Configuration.
type ConfigurationSpec struct {
	// Semantic version constraints of Crossplane that package is compatible with.
	Crossplane *ConfigurationSpecCrossplane `field:"optional" json:"crossplane" yaml:"crossplane"`
	// Dependencies on other packages.
	DependsOn *[]*ConfigurationSpecDependsOn `field:"optional" json:"dependsOn" yaml:"dependsOn"`
}

