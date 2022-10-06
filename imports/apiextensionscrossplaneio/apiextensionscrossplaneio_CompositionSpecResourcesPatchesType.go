// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Type sets the patching behaviour to be used.
//
// Each patch type may require its' own fields to be set on the Patch object.
type CompositionSpecResourcesPatchesType string

const (
	// FromCompositeFieldPath.
	CompositionSpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH CompositionSpecResourcesPatchesType = "FROM_COMPOSITE_FIELD_PATH"
	// PatchSet.
	CompositionSpecResourcesPatchesType_PATCH_SET CompositionSpecResourcesPatchesType = "PATCH_SET"
	// ToCompositeFieldPath.
	CompositionSpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH CompositionSpecResourcesPatchesType = "TO_COMPOSITE_FIELD_PATH"
	// CombineFromComposite.
	CompositionSpecResourcesPatchesType_COMBINE_FROM_COMPOSITE CompositionSpecResourcesPatchesType = "COMBINE_FROM_COMPOSITE"
	// CombineToComposite.
	CompositionSpecResourcesPatchesType_COMBINE_TO_COMPOSITE CompositionSpecResourcesPatchesType = "COMBINE_TO_COMPOSITE"
)

