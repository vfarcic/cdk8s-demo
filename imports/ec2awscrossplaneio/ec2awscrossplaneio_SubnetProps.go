// ec2awscrossplaneio
package ec2awscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Subnet is a managed resource that represents an AWS VPC Subnet.
type SubnetProps struct {
	// A SubnetSpec defines the desired state of a Subnet.
	Spec *SubnetSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

