// pkgcrossplaneio
package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A ConfigurationRevision that has been added to Crossplane.
type ConfigurationRevisionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PackageRevisionSpec specifies the desired state of a PackageRevision.
	Spec *ConfigurationRevisionSpec `field:"optional" json:"spec" yaml:"spec"`
}

