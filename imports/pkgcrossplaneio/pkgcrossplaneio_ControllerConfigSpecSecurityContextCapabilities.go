// pkgcrossplaneio
package pkgcrossplaneio


// The capabilities to add/drop when running containers.
//
// Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.
type ControllerConfigSpecSecurityContextCapabilities struct {
	// Added capabilities.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

