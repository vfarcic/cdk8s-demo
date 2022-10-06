// secretscrossplaneio
package secretscrossplaneio


// Source of the credentials.
type StoreConfigSpecVaultAuthTokenSource string

const (
	// None.
	StoreConfigSpecVaultAuthTokenSource_NONE StoreConfigSpecVaultAuthTokenSource = "NONE"
	// Secret.
	StoreConfigSpecVaultAuthTokenSource_SECRET StoreConfigSpecVaultAuthTokenSource = "SECRET"
	// Environment.
	StoreConfigSpecVaultAuthTokenSource_ENVIRONMENT StoreConfigSpecVaultAuthTokenSource = "ENVIRONMENT"
	// Filesystem.
	StoreConfigSpecVaultAuthTokenSource_FILESYSTEM StoreConfigSpecVaultAuthTokenSource = "FILESYSTEM"
)

