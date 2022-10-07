package main

import (
	composition "example.com/cdk8s-demo/imports/apiextensionscrossplaneio"
	"github.com/aws/jsii-runtime-go"

	subnet "example.com/cdk8s-demo/imports/subnet/ec2awscrossplaneio"
	vpc "example.com/cdk8s-demo/imports/vpc/ec2awscrossplaneio"
)

func GetAwsResources() *[]*composition.CompositionSpecResources {
	return &[]*composition.CompositionSpecResources{
		getVpc(),
		getSubnet(),
	}
}

func getVpc() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("vpc"),
		Base: vpc.Vpc_Manifest(&vpc.VpcProps{
			Spec: &vpc.VpcSpec{
				ForProvider: &vpc.VpcSpecForProvider{
					Region:             jsii.String("us-east-1"),
					CidrBlock:          jsii.String("10.0.0.0/16"),
					EnableDnsSupport:   jsii.Bool(true),
					EnableDnsHostNames: jsii.Bool(true),
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("metadata.name")},
		},
	}
}

func getSubnet() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("subnet-a"),
		Base: subnet.Subnet_Manifest(&subnet.SubnetProps{
			Spec: &subnet.SubnetSpec{
				ForProvider: &subnet.SubnetSpecForProvider{
					Region:    jsii.String("us-east-1"),
					CidrBlock: jsii.String("10.0.0.0/24"),
				},
			},
		}),
	}
}
