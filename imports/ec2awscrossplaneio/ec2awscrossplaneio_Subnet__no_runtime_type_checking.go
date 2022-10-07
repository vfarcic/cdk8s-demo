//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// ec2awscrossplaneio
package ec2awscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateSubnet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubnet_ManifestParameters(props *SubnetProps) error {
	return nil
}

func validateSubnet_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewSubnetParameters(scope constructs.Construct, id *string, props *SubnetProps) error {
	return nil
}

