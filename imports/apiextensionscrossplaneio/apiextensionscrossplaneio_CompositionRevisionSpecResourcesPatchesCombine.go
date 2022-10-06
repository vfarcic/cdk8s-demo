// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.
type CompositionRevisionSpecResourcesPatchesCombine struct {
	// Strategy defines the strategy to use to combine the input variable values.
	//
	// Currently only string is supported.
	Strategy CompositionRevisionSpecResourcesPatchesCombineStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// Variables are the list of variables whose values will be retrieved and combined.
	Variables *[]*CompositionRevisionSpecResourcesPatchesCombineVariables `field:"required" json:"variables" yaml:"variables"`
	// String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.
	String *CompositionRevisionSpecResourcesPatchesCombineString `field:"optional" json:"string" yaml:"string"`
}

