// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// A CredentialsSecretRef is a reference to a PostgreSQL connection secret that contains the credentials that must be used to connect to the provider.
//
// +optional.
type ProviderConfigSpecCredentialsConnectionSecretRef struct {
	// Name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the secret.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

