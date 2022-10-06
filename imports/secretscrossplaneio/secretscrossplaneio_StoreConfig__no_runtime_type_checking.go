//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// secretscrossplaneio
package secretscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateStoreConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStoreConfig_ManifestParameters(props *StoreConfigProps) error {
	return nil
}

func validateStoreConfig_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewStoreConfigParameters(scope constructs.Construct, id *string, props *StoreConfigProps) error {
	return nil
}

