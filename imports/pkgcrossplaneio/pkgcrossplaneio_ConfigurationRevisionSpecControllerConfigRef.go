// pkgcrossplaneio
package pkgcrossplaneio


// ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.
type ConfigurationRevisionSpecControllerConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
}

