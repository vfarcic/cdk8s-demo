//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// metapkgcrossplaneio
package metapkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateConfigurationV1Alpha1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigurationV1Alpha1_ManifestParameters(props *ConfigurationV1Alpha1Props) error {
	return nil
}

func validateConfigurationV1Alpha1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewConfigurationV1Alpha1Parameters(scope constructs.Construct, id *string, props *ConfigurationV1Alpha1Props) error {
	return nil
}

