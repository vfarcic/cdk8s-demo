// metapkgcrossplaneio
package metapkgcrossplaneio


// Configuration for the packaged Provider's controller.
type ProviderSpecController struct {
	// Image is the packaged Provider controller image.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// PermissionRequests for RBAC rules required for this provider's controller to function.
	//
	// The RBAC manager is responsible for assessing the requested permissions.
	PermissionRequests *[]*ProviderSpecControllerPermissionRequests `field:"optional" json:"permissionRequests" yaml:"permissionRequests"`
}

