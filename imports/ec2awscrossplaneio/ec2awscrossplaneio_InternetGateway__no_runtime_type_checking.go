//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// ec2awscrossplaneio
package ec2awscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateInternetGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInternetGateway_ManifestParameters(props *InternetGatewayProps) error {
	return nil
}

func validateInternetGateway_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewInternetGatewayParameters(scope constructs.Construct, id *string, props *InternetGatewayProps) error {
	return nil
}

