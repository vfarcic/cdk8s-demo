// pkgcrossplaneio
package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ControllerConfigSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for ControllerConfigSpecResourcesRequests
type jsiiProxy_ControllerConfigSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_ControllerConfigSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControllerConfigSpecResourcesRequests_FromNumber(value *float64) ControllerConfigSpecResourcesRequests {
	_init_.Initialize()

	if err := validateControllerConfigSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecResourcesRequests

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControllerConfigSpecResourcesRequests_FromString(value *string) ControllerConfigSpecResourcesRequests {
	_init_.Initialize()

	if err := validateControllerConfigSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecResourcesRequests

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

