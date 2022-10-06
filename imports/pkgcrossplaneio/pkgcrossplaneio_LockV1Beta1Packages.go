// pkgcrossplaneio
package pkgcrossplaneio


// LockPackage is a package that is in the lock.
type LockV1Beta1Packages struct {
	// Dependencies are the list of dependencies of this package.
	//
	// The order of the dependencies will dictate the order in which they are resolved.
	Dependencies *[]*LockV1Beta1PackagesDependencies `field:"required" json:"dependencies" yaml:"dependencies"`
	// Name corresponds to the name of the package revision for this package.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source is the OCI image name without a tag or digest.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Type is the type of package.
	//
	// Can be either Configuration or Provider.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Version is the tag or digest of the OCI image.
	Version *string `field:"required" json:"version" yaml:"version"`
}

