// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Type sets the patching behaviour to be used.
//
// Each patch type may require its' own fields to be set on the Patch object.
type CompositionRevisionSpecResourcesPatchesType string

const (
	// FromCompositeFieldPath.
	CompositionRevisionSpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH CompositionRevisionSpecResourcesPatchesType = "FROM_COMPOSITE_FIELD_PATH"
	// PatchSet.
	CompositionRevisionSpecResourcesPatchesType_PATCH_SET CompositionRevisionSpecResourcesPatchesType = "PATCH_SET"
	// ToCompositeFieldPath.
	CompositionRevisionSpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH CompositionRevisionSpecResourcesPatchesType = "TO_COMPOSITE_FIELD_PATH"
	// CombineFromComposite.
	CompositionRevisionSpecResourcesPatchesType_COMBINE_FROM_COMPOSITE CompositionRevisionSpecResourcesPatchesType = "COMBINE_FROM_COMPOSITE"
	// CombineToComposite.
	CompositionRevisionSpecResourcesPatchesType_COMBINE_TO_COMPOSITE CompositionRevisionSpecResourcesPatchesType = "COMBINE_TO_COMPOSITE"
)

