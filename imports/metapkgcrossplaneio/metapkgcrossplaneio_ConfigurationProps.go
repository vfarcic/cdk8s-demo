// metapkgcrossplaneio
package metapkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Configuration is the description of a Crossplane Configuration package.
type ConfigurationProps struct {
	// ConfigurationSpec specifies the configuration of a Configuration.
	Spec *ConfigurationSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

