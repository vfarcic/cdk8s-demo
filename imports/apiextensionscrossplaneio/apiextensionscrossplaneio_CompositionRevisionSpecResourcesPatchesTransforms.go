// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Transform is a unit of process whose input is transformed into an output with the supplied configuration.
type CompositionRevisionSpecResourcesPatchesTransforms struct {
	// Type of the transform to be run.
	Type CompositionRevisionSpecResourcesPatchesTransformsType `field:"required" json:"type" yaml:"type"`
	// Convert is used to cast the input into the given output type.
	Convert *CompositionRevisionSpecResourcesPatchesTransformsConvert `field:"optional" json:"convert" yaml:"convert"`
	// Map uses the input as a key in the given map and returns the value.
	Map *map[string]*string `field:"optional" json:"map" yaml:"map"`
	// Math is used to transform the input via mathematical operations such as multiplication.
	Math *CompositionRevisionSpecResourcesPatchesTransformsMath `field:"optional" json:"math" yaml:"math"`
	// String is used to transform the input into a string or a different kind of string.
	//
	// Note that the input does not necessarily need to be a string.
	String *CompositionRevisionSpecResourcesPatchesTransformsString `field:"optional" json:"string" yaml:"string"`
}

