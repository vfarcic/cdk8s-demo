// apiextensionscrossplaneio
package apiextensionscrossplaneio


// CompositeResourceDefinitionSpec specifies the desired state of the definition.
type CompositeResourceDefinitionSpec struct {
	// Group specifies the API group of the defined composite resource.
	//
	// Composite resources are served under `/apis/<group>/...`. Must match the name of the XRD (in the form `<names.plural>.<group>`).
	Group *string `field:"required" json:"group" yaml:"group"`
	// Names specifies the resource and kind names of the defined composite resource.
	Names *CompositeResourceDefinitionSpecNames `field:"required" json:"names" yaml:"names"`
	// Versions is the list of all API versions of the defined composite resource.
	//
	// Version names are used to compute the order in which served versions are listed in API discovery. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10. Note that all versions must have identical schemas; Crossplane does not currently support conversion between different version schemas.
	Versions *[]*CompositeResourceDefinitionSpecVersions `field:"required" json:"versions" yaml:"versions"`
	// ClaimNames specifies the names of an optional composite resource claim.
	//
	// When claim names are specified Crossplane will create a namespaced 'composite resource claim' CRD that corresponds to the defined composite resource. This composite resource claim acts as a namespaced proxy for the composite resource; creating, updating, or deleting the claim will create, update, or delete a corresponding composite resource. You may add claim names to an existing CompositeResourceDefinition, but they cannot be changed or removed once they have been set.
	ClaimNames *CompositeResourceDefinitionSpecClaimNames `field:"optional" json:"claimNames" yaml:"claimNames"`
	// ConnectionSecretKeys is the list of keys that will be exposed to the end user of the defined kind.
	//
	// If the list is empty, all keys will be published.
	ConnectionSecretKeys *[]*string `field:"optional" json:"connectionSecretKeys" yaml:"connectionSecretKeys"`
	// DefaultCompositionRef refers to the Composition resource that will be used in case no composition selector is given.
	DefaultCompositionRef *CompositeResourceDefinitionSpecDefaultCompositionRef `field:"optional" json:"defaultCompositionRef" yaml:"defaultCompositionRef"`
	// EnforcedCompositionRef refers to the Composition resource that will be used by all composite instances whose schema is defined by this definition.
	EnforcedCompositionRef *CompositeResourceDefinitionSpecEnforcedCompositionRef `field:"optional" json:"enforcedCompositionRef" yaml:"enforcedCompositionRef"`
}

