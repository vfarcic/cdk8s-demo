// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Type indicates the type of probe you'd like to use.
type CompositionRevisionSpecResourcesReadinessChecksType string

const (
	// MatchString.
	CompositionRevisionSpecResourcesReadinessChecksType_MATCH_STRING CompositionRevisionSpecResourcesReadinessChecksType = "MATCH_STRING"
	// MatchInteger.
	CompositionRevisionSpecResourcesReadinessChecksType_MATCH_INTEGER CompositionRevisionSpecResourcesReadinessChecksType = "MATCH_INTEGER"
	// NonEmpty.
	CompositionRevisionSpecResourcesReadinessChecksType_NON_EMPTY CompositionRevisionSpecResourcesReadinessChecksType = "NON_EMPTY"
	// None.
	CompositionRevisionSpecResourcesReadinessChecksType_NONE CompositionRevisionSpecResourcesReadinessChecksType = "NONE"
)

