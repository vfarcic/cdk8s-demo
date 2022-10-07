// ec2awscrossplaneio
package ec2awscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A VPC is a managed resource that represents an AWS Virtual Private Cloud.
type VpcProps struct {
	// A VPCSpec defines the desired state of a VPC.
	Spec *VpcSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

