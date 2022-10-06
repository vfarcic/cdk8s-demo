// secretscrossplaneio
package secretscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A StoreConfig configures how Crossplane controllers should store connection details.
type StoreConfigProps struct {
	// A StoreConfigSpec defines the desired state of a StoreConfig.
	Spec *StoreConfigSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

