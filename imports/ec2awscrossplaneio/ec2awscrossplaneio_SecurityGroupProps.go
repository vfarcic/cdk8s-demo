// ec2awscrossplaneio
package ec2awscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A SecurityGroup is a managed resource that represents an AWS VPC Security Group.
type SecurityGroupProps struct {
	// A SecurityGroupSpec defines the desired state of a SecurityGroup.
	Spec *SecurityGroupSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

