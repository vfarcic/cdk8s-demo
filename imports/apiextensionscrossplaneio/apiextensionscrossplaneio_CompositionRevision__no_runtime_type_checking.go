//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// apiextensionscrossplaneio
package apiextensionscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateCompositionRevision_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCompositionRevision_ManifestParameters(props *CompositionRevisionProps) error {
	return nil
}

func validateCompositionRevision_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewCompositionRevisionParameters(scope constructs.Construct, id *string, props *CompositionRevisionProps) error {
	return nil
}

