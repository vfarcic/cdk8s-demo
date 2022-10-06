//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// ec2awscrossplaneio
package ec2awscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateVpc_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpc_ManifestParameters(props *VpcProps) error {
	return nil
}

func validateVpc_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewVpcParameters(scope constructs.Construct, id *string, props *VpcProps) error {
	return nil
}

