// pkgcrossplaneio
package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type ControllerConfigSpecEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for ControllerConfigSpecEnvValueFromResourceFieldRefDivisor
type jsiiProxy_ControllerConfigSpecEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_ControllerConfigSpecEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControllerConfigSpecEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) ControllerConfigSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateControllerConfigSpecEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControllerConfigSpecEnvValueFromResourceFieldRefDivisor_FromString(value *string) ControllerConfigSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateControllerConfigSpecEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

