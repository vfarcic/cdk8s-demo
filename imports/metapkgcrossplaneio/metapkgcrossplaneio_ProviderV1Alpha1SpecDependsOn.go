// metapkgcrossplaneio
package metapkgcrossplaneio


// Dependency is a dependency on another package.
//
// One of Provider or Configuration may be supplied.
type ProviderV1Alpha1SpecDependsOn struct {
	// Version is the semantic version constraints of the dependency image.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration is the name of a Configuration package image.
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// Provider is the name of a Provider package image.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

