// pkgcrossplaneio
package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Provider is the CRD type for a request to add a provider to Crossplane.
type ProviderProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ProviderSpec specifies details about a request to install a provider to Crossplane.
	Spec *ProviderSpec `field:"optional" json:"spec" yaml:"spec"`
}

