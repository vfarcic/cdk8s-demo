// pkgcrossplaneio
package pkgcrossplaneio


// Metadata that will be added to the provider Pod.
type ControllerConfigSpecMetadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// This will only affect labels on the pod, not the pod selector. Labels will be merged with internal labels used by crossplane, and labels with a crossplane.io key might be overwritten. More info: http://kubernetes.io/docs/user-guide/labels
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

