// apiextensionscrossplaneio
package apiextensionscrossplaneio


// EnforcedCompositionRef refers to the Composition resource that will be used by all composite instances whose schema is defined by this definition.
type CompositeResourceDefinitionSpecEnforcedCompositionRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
}

