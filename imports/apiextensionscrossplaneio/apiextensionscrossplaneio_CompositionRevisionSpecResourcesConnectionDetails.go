// apiextensionscrossplaneio
package apiextensionscrossplaneio


// ConnectionDetail includes the information about the propagation of the connection information from one secret to another.
type CompositionRevisionSpecResourcesConnectionDetails struct {
	// FromConnectionSecretKey is the key that will be used to fetch the value from the given target resource's secret.
	FromConnectionSecretKey *string `field:"optional" json:"fromConnectionSecretKey" yaml:"fromConnectionSecretKey"`
	// FromFieldPath is the path of the field on the composed resource whose value to be used as input.
	//
	// Name must be specified if the type is FromFieldPath is specified.
	FromFieldPath *string `field:"optional" json:"fromFieldPath" yaml:"fromFieldPath"`
	// Name of the connection secret key that will be propagated to the connection secret of the composition instance.
	//
	// Leave empty if you'd like to use the same key name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type sets the connection detail fetching behaviour to be used.
	//
	// Each connection detail type may require its own fields to be set on the ConnectionDetail object. If the type is omitted Crossplane will attempt to infer it based on which other fields were specified.
	Type CompositionRevisionSpecResourcesConnectionDetailsType `field:"optional" json:"type" yaml:"type"`
	// Value that will be propagated to the connection secret of the composition instance.
	//
	// Typically you should use FromConnectionSecretKey instead, but an explicit value may be set to inject a fixed, non-sensitive connection secret values, for example a well-known port. Supercedes FromConnectionSecretKey when set.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

