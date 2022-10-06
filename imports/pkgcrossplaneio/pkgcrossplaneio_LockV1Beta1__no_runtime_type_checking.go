//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// pkgcrossplaneio
package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateLockV1Beta1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLockV1Beta1_ManifestParameters(props *LockV1Beta1Props) error {
	return nil
}

func validateLockV1Beta1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewLockV1Beta1Parameters(scope constructs.Construct, id *string, props *LockV1Beta1Props) error {
	return nil
}

