// pkgcrossplaneio
package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Lock is the CRD type that tracks package dependencies.
type LockV1Beta1Props struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Packages *[]*LockV1Beta1Packages `field:"optional" json:"packages" yaml:"packages"`
}

