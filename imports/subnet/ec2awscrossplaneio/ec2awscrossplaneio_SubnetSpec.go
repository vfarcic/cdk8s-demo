// ec2awscrossplaneio
package ec2awscrossplaneio


// A SubnetSpec defines the desired state of a Subnet.
type SubnetSpec struct {
	// SubnetParameters define the desired state of an AWS VPC Subnet.
	ForProvider *SubnetSpecForProvider `field:"required" json:"forProvider" yaml:"forProvider"`
	// DeletionPolicy specifies what will happen to the underlying external when this managed resource is deleted - either "Delete" or "Orphan" the external resource.
	DeletionPolicy SubnetSpecDeletionPolicy `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// ProviderConfigReference specifies how the provider that will be used to create, observe, update, and delete this managed resource should be configured.
	ProviderConfigRef *SubnetSpecProviderConfigRef `field:"optional" json:"providerConfigRef" yaml:"providerConfigRef"`
	// ProviderReference specifies the provider that will be used to create, observe, update, and delete this managed resource.
	//
	// Deprecated: Please use ProviderConfigReference, i.e. `providerConfigRef`
	ProviderRef *SubnetSpecProviderRef `field:"optional" json:"providerRef" yaml:"providerRef"`
	// PublishConnectionDetailsTo specifies the connection secret config which contains a name, metadata and a reference to secret store config to which any connection details for this managed resource should be written.
	//
	// Connection details frequently include the endpoint, username, and password required to connect to the managed resource.
	PublishConnectionDetailsTo *SubnetSpecPublishConnectionDetailsTo `field:"optional" json:"publishConnectionDetailsTo" yaml:"publishConnectionDetailsTo"`
	// WriteConnectionSecretToReference specifies the namespace and name of a Secret to which any connection details for this managed resource should be written.
	//
	// Connection details frequently include the endpoint, username, and password required to connect to the managed resource. This field is planned to be replaced in a future release in favor of PublishConnectionDetailsTo. Currently, both could be set independently and connection details would be published to both without affecting each other.
	WriteConnectionSecretToRef *SubnetSpecWriteConnectionSecretToRef `field:"optional" json:"writeConnectionSecretToRef" yaml:"writeConnectionSecretToRef"`
}

