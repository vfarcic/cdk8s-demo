//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// apiextensionscrossplaneio
package apiextensionscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateCompositeResourceDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCompositeResourceDefinition_ManifestParameters(props *CompositeResourceDefinitionProps) error {
	return nil
}

func validateCompositeResourceDefinition_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewCompositeResourceDefinitionParameters(scope constructs.Construct, id *string, props *CompositeResourceDefinitionProps) error {
	return nil
}

