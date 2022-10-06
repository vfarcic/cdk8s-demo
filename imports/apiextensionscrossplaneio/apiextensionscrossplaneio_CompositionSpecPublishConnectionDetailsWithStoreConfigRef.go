// apiextensionscrossplaneio
package apiextensionscrossplaneio


// PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection secrets of composite resource dynamically provisioned using this composition will be published.
type CompositionSpecPublishConnectionDetailsWithStoreConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
}

