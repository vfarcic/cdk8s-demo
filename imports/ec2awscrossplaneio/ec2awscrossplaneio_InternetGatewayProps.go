// ec2awscrossplaneio
package ec2awscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// An InternetGateway is a managed resource that represents an AWS VPC Internet Gateway.
type InternetGatewayProps struct {
	// An InternetGatewaySpec defines the desired state of an InternetGateway.
	Spec *InternetGatewaySpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

