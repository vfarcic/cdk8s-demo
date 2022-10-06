// pkgcrossplaneio
package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Configuration is the CRD type for a request to add a configuration to Crossplane.
type ConfigurationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ConfigurationSpec specifies details about a request to install a configuration to Crossplane.
	Spec *ConfigurationSpec `field:"optional" json:"spec" yaml:"spec"`
}

