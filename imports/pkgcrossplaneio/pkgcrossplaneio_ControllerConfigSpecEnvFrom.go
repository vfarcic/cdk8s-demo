// pkgcrossplaneio
package pkgcrossplaneio


// EnvFromSource represents the source of a set of ConfigMaps.
type ControllerConfigSpecEnvFrom struct {
	// The ConfigMap to select from.
	ConfigMapRef *ControllerConfigSpecEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifier to prepend to each key in the ConfigMap.
	//
	// Must be a C_IDENTIFIER.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The Secret to select from.
	SecretRef *ControllerConfigSpecEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

