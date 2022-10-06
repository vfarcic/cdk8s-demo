//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// pkgcrossplaneio
package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateControllerConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateControllerConfig_ManifestParameters(props *ControllerConfigProps) error {
	return nil
}

func validateControllerConfig_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewControllerConfigParameters(scope constructs.Construct, id *string, props *ControllerConfigProps) error {
	return nil
}

