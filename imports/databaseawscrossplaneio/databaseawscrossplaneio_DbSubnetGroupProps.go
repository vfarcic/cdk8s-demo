// databaseawscrossplaneio
package databaseawscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A DBSubnetGroup is a managed resource that represents an AWS VPC Database Subnet Group.
type DbSubnetGroupProps struct {
	// A DBSubnetGroupSpec defines the desired state of a DBSubnetGroup.
	Spec *DbSubnetGroupSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

