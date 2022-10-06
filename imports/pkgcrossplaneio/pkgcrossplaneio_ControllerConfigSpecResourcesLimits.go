// pkgcrossplaneio
package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ControllerConfigSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ControllerConfigSpecResourcesLimits
type jsiiProxy_ControllerConfigSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ControllerConfigSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControllerConfigSpecResourcesLimits_FromNumber(value *float64) ControllerConfigSpecResourcesLimits {
	_init_.Initialize()

	if err := validateControllerConfigSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecResourcesLimits

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControllerConfigSpecResourcesLimits_FromString(value *string) ControllerConfigSpecResourcesLimits {
	_init_.Initialize()

	if err := validateControllerConfigSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecResourcesLimits

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

