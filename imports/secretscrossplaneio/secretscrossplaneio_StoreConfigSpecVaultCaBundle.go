// secretscrossplaneio
package secretscrossplaneio


// CABundle configures CA bundle for Vault Server.
type StoreConfigSpecVaultCaBundle struct {
	// Source of the credentials.
	Source StoreConfigSpecVaultCaBundleSource `field:"required" json:"source" yaml:"source"`
	// Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.
	Env *StoreConfigSpecVaultCaBundleEnv `field:"optional" json:"env" yaml:"env"`
	// Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.
	Fs *StoreConfigSpecVaultCaBundleFs `field:"optional" json:"fs" yaml:"fs"`
	// A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.
	SecretRef *StoreConfigSpecVaultCaBundleSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

