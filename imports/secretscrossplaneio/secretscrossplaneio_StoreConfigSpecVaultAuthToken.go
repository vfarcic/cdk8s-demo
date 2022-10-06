// secretscrossplaneio
package secretscrossplaneio


// Token configures Token Auth for Vault.
type StoreConfigSpecVaultAuthToken struct {
	// Source of the credentials.
	Source StoreConfigSpecVaultAuthTokenSource `field:"required" json:"source" yaml:"source"`
	// Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.
	Env *StoreConfigSpecVaultAuthTokenEnv `field:"optional" json:"env" yaml:"env"`
	// Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.
	Fs *StoreConfigSpecVaultAuthTokenFs `field:"optional" json:"fs" yaml:"fs"`
	// A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.
	SecretRef *StoreConfigSpecVaultAuthTokenSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

