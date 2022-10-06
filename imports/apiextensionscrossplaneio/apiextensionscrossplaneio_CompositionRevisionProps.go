// apiextensionscrossplaneio
package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A CompositionRevision represents a revision in time of a Composition.
//
// Revisions are created by Crossplane; they should be treated as immutable.
type CompositionRevisionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// CompositionRevisionSpec specifies the desired state of the composition revision.
	Spec *CompositionRevisionSpec `field:"optional" json:"spec" yaml:"spec"`
}

