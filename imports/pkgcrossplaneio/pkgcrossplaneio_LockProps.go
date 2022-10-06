// pkgcrossplaneio
package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Lock is the CRD type that tracks package dependencies.
//
// [DEPRECATED]: Please use the identical v1beta1 API instead. The v1alpha1 API is scheduled to be removed in Crossplane v1.7.
type LockProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Packages *[]*LockPackages `field:"optional" json:"packages" yaml:"packages"`
}

