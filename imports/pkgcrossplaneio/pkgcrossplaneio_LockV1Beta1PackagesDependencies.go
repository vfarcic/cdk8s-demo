// pkgcrossplaneio
package pkgcrossplaneio


// A Dependency is a dependency of a package in the lock.
type LockV1Beta1PackagesDependencies struct {
	// Constraints is a valid semver range, which will be used to select a valid dependency version.
	Constraints *string `field:"required" json:"constraints" yaml:"constraints"`
	// Package is the OCI image name without a tag or digest.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Type is the type of package.
	//
	// Can be either Configuration or Provider.
	Type *string `field:"required" json:"type" yaml:"type"`
}

