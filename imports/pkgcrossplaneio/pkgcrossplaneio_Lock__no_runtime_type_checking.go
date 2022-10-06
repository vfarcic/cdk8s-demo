//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// pkgcrossplaneio
package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateLock_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLock_ManifestParameters(props *LockProps) error {
	return nil
}

func validateLock_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewLockParameters(scope constructs.Construct, id *string, props *LockProps) error {
	return nil
}

