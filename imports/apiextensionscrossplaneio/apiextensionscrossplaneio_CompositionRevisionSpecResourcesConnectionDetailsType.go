// apiextensionscrossplaneio
package apiextensionscrossplaneio


// Type sets the connection detail fetching behaviour to be used.
//
// Each connection detail type may require its own fields to be set on the ConnectionDetail object. If the type is omitted Crossplane will attempt to infer it based on which other fields were specified.
type CompositionRevisionSpecResourcesConnectionDetailsType string

const (
	// FromConnectionSecretKey.
	CompositionRevisionSpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY CompositionRevisionSpecResourcesConnectionDetailsType = "FROM_CONNECTION_SECRET_KEY"
	// FromFieldPath.
	CompositionRevisionSpecResourcesConnectionDetailsType_FROM_FIELD_PATH CompositionRevisionSpecResourcesConnectionDetailsType = "FROM_FIELD_PATH"
	// FromValue.
	CompositionRevisionSpecResourcesConnectionDetailsType_FROM_VALUE CompositionRevisionSpecResourcesConnectionDetailsType = "FROM_VALUE"
)

