//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// pkgcrossplaneio
package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateConfigurationRevision_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigurationRevision_ManifestParameters(props *ConfigurationRevisionProps) error {
	return nil
}

func validateConfigurationRevision_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewConfigurationRevisionParameters(scope constructs.Construct, id *string, props *ConfigurationRevisionProps) error {
	return nil
}

