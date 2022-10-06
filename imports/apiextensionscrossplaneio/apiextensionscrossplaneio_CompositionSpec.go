// apiextensionscrossplaneio
package apiextensionscrossplaneio


// CompositionSpec specifies desired state of a composition.
type CompositionSpec struct {
	// CompositeTypeRef specifies the type of composite resource that this composition is compatible with.
	CompositeTypeRef *CompositionSpecCompositeTypeRef `field:"required" json:"compositeTypeRef" yaml:"compositeTypeRef"`
	// Resources is the list of resource templates that will be used when a composite resource referring to this composition is created.
	Resources *[]*CompositionSpecResources `field:"required" json:"resources" yaml:"resources"`
	// PatchSets define a named set of patches that may be included by any resource in this Composition.
	//
	// PatchSets cannot themselves refer to other PatchSets.
	PatchSets *[]*CompositionSpecPatchSets `field:"optional" json:"patchSets" yaml:"patchSets"`
	// PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection secrets of composite resource dynamically provisioned using this composition will be published.
	PublishConnectionDetailsWithStoreConfigRef *CompositionSpecPublishConnectionDetailsWithStoreConfigRef `field:"optional" json:"publishConnectionDetailsWithStoreConfigRef" yaml:"publishConnectionDetailsWithStoreConfigRef"`
	// WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created.
	//
	// This field is planned to be replaced in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.
	WriteConnectionSecretsToNamespace *string `field:"optional" json:"writeConnectionSecretsToNamespace" yaml:"writeConnectionSecretsToNamespace"`
}

