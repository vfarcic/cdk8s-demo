// secretscrossplaneio
package secretscrossplaneio


// Vault configures a Vault secret store.
type StoreConfigSpecVault struct {
	// Auth configures an authentication method for Vault.
	Auth *StoreConfigSpecVaultAuth `field:"required" json:"auth" yaml:"auth"`
	// MountPath is the mount path of the KV secrets engine.
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Server is the url of the Vault server, e.g. "https://vault.acme.org".
	Server *string `field:"required" json:"server" yaml:"server"`
	// CABundle configures CA bundle for Vault Server.
	CaBundle *StoreConfigSpecVaultCaBundle `field:"optional" json:"caBundle" yaml:"caBundle"`
	// Version of the KV Secrets engine of Vault.
	//
	// https://www.vaultproject.io/docs/secrets/kv
	Version *string `field:"optional" json:"version" yaml:"version"`
}

