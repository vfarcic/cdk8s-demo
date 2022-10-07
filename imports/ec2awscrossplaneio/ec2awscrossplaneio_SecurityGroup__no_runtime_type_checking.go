//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// ec2awscrossplaneio
package ec2awscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateSecurityGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityGroup_ManifestParameters(props *SecurityGroupProps) error {
	return nil
}

func validateSecurityGroup_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewSecurityGroupParameters(scope constructs.Construct, id *string, props *SecurityGroupProps) error {
	return nil
}

