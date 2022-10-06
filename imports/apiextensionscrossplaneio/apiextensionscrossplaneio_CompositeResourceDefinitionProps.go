// apiextensionscrossplaneio
package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A CompositeResourceDefinition defines a new kind of composite infrastructure resource.
//
// The new resource is composed of other composite or managed infrastructure resources.
type CompositeResourceDefinitionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// CompositeResourceDefinitionSpec specifies the desired state of the definition.
	Spec *CompositeResourceDefinitionSpec `field:"optional" json:"spec" yaml:"spec"`
}

