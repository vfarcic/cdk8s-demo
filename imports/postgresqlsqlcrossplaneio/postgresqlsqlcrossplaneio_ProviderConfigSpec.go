// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials *ProviderConfigSpecCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// Defines the database name used to set up a connection to the provided PostgreSQL instance.
	//
	// Same as PGDATABASE environment variable.
	DefaultDatabase *string `field:"optional" json:"defaultDatabase" yaml:"defaultDatabase"`
	// Defines the SSL mode used to set up a connection to the provided PostgreSQL instance.
	SslMode ProviderConfigSpecSslMode `field:"optional" json:"sslMode" yaml:"sslMode"`
}

