// apiextensionscrossplaneio
package apiextensionscrossplaneio


// String is used to transform the input into a string or a different kind of string.
//
// Note that the input does not necessarily need to be a string.
type CompositionRevisionSpecPatchSetsPatchesTransformsString struct {
	// Convert the type of conversion to Upper/Lower case.
	Convert CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert `field:"optional" json:"convert" yaml:"convert"`
	// Format the input using a Go format string.
	//
	// See https://golang.org/pkg/fmt/ for details.
	Fmt *string `field:"optional" json:"fmt" yaml:"fmt"`
	// Extract a match from the input using a regular expression.
	Regexp *CompositionRevisionSpecPatchSetsPatchesTransformsStringRegexp `field:"optional" json:"regexp" yaml:"regexp"`
	// Trim the prefix or suffix from the input.
	Trim *string `field:"optional" json:"trim" yaml:"trim"`
	// Type of the string transform to be run.
	Type CompositionRevisionSpecPatchSetsPatchesTransformsStringType `field:"optional" json:"type" yaml:"type"`
}

