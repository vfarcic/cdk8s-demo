// databaseawscrossplaneio
package databaseawscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// An RDSInstance is a managed resource that represents an AWS Relational Database Service instance.
type RdsInstanceProps struct {
	// An RDSInstanceSpec defines the desired state of an RDSInstance.
	Spec *RdsInstanceSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

