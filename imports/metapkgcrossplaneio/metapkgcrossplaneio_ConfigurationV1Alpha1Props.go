// metapkgcrossplaneio
package metapkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Configuration is the description of a Crossplane Configuration package.
type ConfigurationV1Alpha1Props struct {
	// ConfigurationSpec specifies the configuration of a Configuration.
	Spec *ConfigurationV1Alpha1Spec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

