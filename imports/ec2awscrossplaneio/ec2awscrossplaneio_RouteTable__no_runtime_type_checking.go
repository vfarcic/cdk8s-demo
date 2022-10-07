//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// ec2awscrossplaneio
package ec2awscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateRouteTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRouteTable_ManifestParameters(props *RouteTableProps) error {
	return nil
}

func validateRouteTable_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewRouteTableParameters(scope constructs.Construct, id *string, props *RouteTableProps) error {
	return nil
}

