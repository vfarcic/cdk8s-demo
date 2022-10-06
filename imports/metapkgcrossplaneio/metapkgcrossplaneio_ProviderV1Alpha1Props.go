// metapkgcrossplaneio
package metapkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Provider is the description of a Crossplane Provider package.
type ProviderV1Alpha1Props struct {
	// ProviderSpec specifies the configuration of a Provider.
	Spec *ProviderV1Alpha1Spec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

