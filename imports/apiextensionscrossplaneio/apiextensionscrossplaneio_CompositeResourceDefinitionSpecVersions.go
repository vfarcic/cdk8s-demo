// apiextensionscrossplaneio
package apiextensionscrossplaneio


// CompositeResourceDefinitionVersion describes a version of an XR.
type CompositeResourceDefinitionSpecVersions struct {
	// Name of this version, e.g. “v1”, “v2beta1”, etc. Composite resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Referenceable specifies that this version may be referenced by a Composition in order to configure which resources an XR may be composed of.
	//
	// Exactly one version must be marked as referenceable; all Compositions must target only the referenceable version. The referenceable version must be served.
	Referenceable *bool `field:"required" json:"referenceable" yaml:"referenceable"`
	// Served specifies that this version should be served via REST APIs.
	Served *bool `field:"required" json:"served" yaml:"served"`
	// AdditionalPrinterColumns specifies additional columns returned in Table output.
	//
	// If no columns are specified, a single column displaying the age of the custom resource is used. See the following link for details: https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables
	AdditionalPrinterColumns *[]*CompositeResourceDefinitionSpecVersionsAdditionalPrinterColumns `field:"optional" json:"additionalPrinterColumns" yaml:"additionalPrinterColumns"`
	// The deprecated field specifies that this version is deprecated and should not be used.
	Deprecated *bool `field:"optional" json:"deprecated" yaml:"deprecated"`
	// DeprecationWarning specifies the message that should be shown to the user when using this version.
	DeprecationWarning *string `field:"optional" json:"deprecationWarning" yaml:"deprecationWarning"`
	// Schema describes the schema used for validation, pruning, and defaulting of this version of the defined composite resource.
	//
	// Fields required by all composite resources will be injected into this schema automatically, and will override equivalently named fields in this schema. Omitting this schema results in a schema that contains only the fields required by all composite resources.
	Schema *CompositeResourceDefinitionSpecVersionsSchema `field:"optional" json:"schema" yaml:"schema"`
}

