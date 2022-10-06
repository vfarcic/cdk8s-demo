// apiextensionscrossplaneio
package apiextensionscrossplaneio


// A PatchSet is a set of patches that can be reused from all resources within a Composition.
type CompositionRevisionSpecPatchSets struct {
	// Name of this PatchSet.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Patches will be applied as an overlay to the base resource.
	Patches *[]*CompositionRevisionSpecPatchSetsPatches `field:"required" json:"patches" yaml:"patches"`
}

