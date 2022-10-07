// ec2awscrossplaneio
package ec2awscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A RouteTable is a managed resource that represents an AWS VPC Route Table.
type RouteTableProps struct {
	// A RouteTableSpec defines the desired state of a RouteTable.
	Spec *RouteTableSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

