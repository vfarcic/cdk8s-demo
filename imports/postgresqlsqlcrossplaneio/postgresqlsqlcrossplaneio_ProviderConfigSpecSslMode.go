// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// Defines the SSL mode used to set up a connection to the provided PostgreSQL instance.
type ProviderConfigSpecSslMode string

const (
	// disable.
	ProviderConfigSpecSslMode_DISABLE ProviderConfigSpecSslMode = "DISABLE"
	// require.
	ProviderConfigSpecSslMode_REQUIRE ProviderConfigSpecSslMode = "REQUIRE"
	// verify-ca.
	ProviderConfigSpecSslMode_VERIFY_CA ProviderConfigSpecSslMode = "VERIFY_CA"
	// verify-full.
	ProviderConfigSpecSslMode_VERIFY_FULL ProviderConfigSpecSslMode = "VERIFY_FULL"
)

