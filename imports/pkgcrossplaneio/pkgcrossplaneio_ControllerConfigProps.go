// pkgcrossplaneio
package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ControllerConfig is the CRD type for a packaged controller configuration.
type ControllerConfigProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ControllerConfigSpec specifies the configuration for a packaged controller.
	//
	// Values provided will override package manager defaults. Labels and annotations are passed to both the controller Deployment and ServiceAccount.
	Spec *ControllerConfigSpec `field:"optional" json:"spec" yaml:"spec"`
}

