// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Database represents the declarative state of a PostgreSQL database.
type DatabaseProps struct {
	// A DatabaseSpec defines the desired state of a Database.
	Spec *DatabaseSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

