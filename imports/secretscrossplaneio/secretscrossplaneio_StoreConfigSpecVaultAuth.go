// secretscrossplaneio
package secretscrossplaneio


// Auth configures an authentication method for Vault.
type StoreConfigSpecVaultAuth struct {
	// Method configures which auth method will be used.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Token configures Token Auth for Vault.
	Token *StoreConfigSpecVaultAuthToken `field:"optional" json:"token" yaml:"token"`
}

