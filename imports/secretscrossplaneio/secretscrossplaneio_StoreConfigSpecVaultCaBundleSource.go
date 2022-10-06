// secretscrossplaneio
package secretscrossplaneio


// Source of the credentials.
type StoreConfigSpecVaultCaBundleSource string

const (
	// None.
	StoreConfigSpecVaultCaBundleSource_NONE StoreConfigSpecVaultCaBundleSource = "NONE"
	// Secret.
	StoreConfigSpecVaultCaBundleSource_SECRET StoreConfigSpecVaultCaBundleSource = "SECRET"
	// Environment.
	StoreConfigSpecVaultCaBundleSource_ENVIRONMENT StoreConfigSpecVaultCaBundleSource = "ENVIRONMENT"
	// Filesystem.
	StoreConfigSpecVaultCaBundleSource_FILESYSTEM StoreConfigSpecVaultCaBundleSource = "FILESYSTEM"
)

