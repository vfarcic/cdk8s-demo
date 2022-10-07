//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// databaseawscrossplaneio
package databaseawscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateDbSubnetGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDbSubnetGroup_ManifestParameters(props *DbSubnetGroupProps) error {
	return nil
}

func validateDbSubnetGroup_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewDbSubnetGroupParameters(scope constructs.Construct, id *string, props *DbSubnetGroupProps) error {
	return nil
}

