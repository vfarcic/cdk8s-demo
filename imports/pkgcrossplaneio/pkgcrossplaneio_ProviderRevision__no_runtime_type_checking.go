//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// pkgcrossplaneio
package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateProviderRevision_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProviderRevision_ManifestParameters(props *ProviderRevisionProps) error {
	return nil
}

func validateProviderRevision_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewProviderRevisionParameters(scope constructs.Construct, id *string, props *ProviderRevisionProps) error {
	return nil
}

