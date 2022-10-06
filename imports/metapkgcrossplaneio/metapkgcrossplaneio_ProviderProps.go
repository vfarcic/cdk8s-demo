// metapkgcrossplaneio
package metapkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Provider is the description of a Crossplane Provider package.
type ProviderProps struct {
	// ProviderSpec specifies the configuration of a Provider.
	Spec *ProviderSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

