//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubePersistentVolumeClaimList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubePersistentVolumeClaimList_ManifestParameters(props *KubePersistentVolumeClaimListProps) error {
	return nil
}

func validateKubePersistentVolumeClaimList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubePersistentVolumeClaimListParameters(scope constructs.Construct, id *string, props *KubePersistentVolumeClaimListProps) error {
	return nil
}

