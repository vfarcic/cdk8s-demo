// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Type sets the patching behaviour to be used.
//
// Each patch type may require its' own fields to be set on the Patch object.
type CompositionRevisionSpecPatchSetsPatchesType string

const (
	// FromCompositeFieldPath.
	CompositionRevisionSpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH CompositionRevisionSpecPatchSetsPatchesType = "FROM_COMPOSITE_FIELD_PATH"
	// PatchSet.
	CompositionRevisionSpecPatchSetsPatchesType_PATCH_SET CompositionRevisionSpecPatchSetsPatchesType = "PATCH_SET"
	// ToCompositeFieldPath.
	CompositionRevisionSpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH CompositionRevisionSpecPatchSetsPatchesType = "TO_COMPOSITE_FIELD_PATH"
	// CombineFromComposite.
	CompositionRevisionSpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE CompositionRevisionSpecPatchSetsPatchesType = "COMBINE_FROM_COMPOSITE"
	// CombineToComposite.
	CompositionRevisionSpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE CompositionRevisionSpecPatchSetsPatchesType = "COMBINE_TO_COMPOSITE"
)

