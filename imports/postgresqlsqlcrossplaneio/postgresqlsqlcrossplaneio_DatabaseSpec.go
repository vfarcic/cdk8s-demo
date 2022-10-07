// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// A DatabaseSpec defines the desired state of a Database.
type DatabaseSpec struct {
	// DatabaseParameters are the configurable fields of a Database.
	ForProvider *DatabaseSpecForProvider `field:"required" json:"forProvider" yaml:"forProvider"`
	// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
	//
	// The "Delete" policy is the default when no policy is specified.
	DeletionPolicy DatabaseSpecDeletionPolicy `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// ProviderConfigReference specifies how the provider that will be used to create, observe, update, and delete this managed resource should be configured.
	ProviderConfigRef *DatabaseSpecProviderConfigRef `field:"optional" json:"providerConfigRef" yaml:"providerConfigRef"`
	// ProviderReference specifies the provider that will be used to create, observe, update, and delete this managed resource.
	//
	// Deprecated: Please use ProviderConfigReference, i.e. `providerConfigRef`
	ProviderRef *DatabaseSpecProviderRef `field:"optional" json:"providerRef" yaml:"providerRef"`
	// WriteConnectionSecretToReference specifies the namespace and name of a Secret to which any connection details for this managed resource should be written.
	//
	// Connection details frequently include the endpoint, username, and password required to connect to the managed resource.
	WriteConnectionSecretToRef *DatabaseSpecWriteConnectionSecretToRef `field:"optional" json:"writeConnectionSecretToRef" yaml:"writeConnectionSecretToRef"`
}

