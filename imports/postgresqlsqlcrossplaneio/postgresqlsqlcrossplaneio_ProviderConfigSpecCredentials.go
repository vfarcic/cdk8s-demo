// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// Credentials required to authenticate to this provider.
type ProviderConfigSpecCredentials struct {
	// Source of the provider credentials.
	Source ProviderConfigSpecCredentialsSource `field:"required" json:"source" yaml:"source"`
	// A CredentialsSecretRef is a reference to a PostgreSQL connection secret that contains the credentials that must be used to connect to the provider.
	//
	// +optional.
	ConnectionSecretRef *ProviderConfigSpecCredentialsConnectionSecretRef `field:"optional" json:"connectionSecretRef" yaml:"connectionSecretRef"`
}

