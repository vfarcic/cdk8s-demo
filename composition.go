package main

import (
	composition "example.com/cdk8s-demo/imports/apiextensionscrossplaneio"
	ec2 "example.com/cdk8s-demo/imports/ec2awscrossplaneio"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CompositionProps struct {
	Provider *string
	DB       *string
	// Image         *string
	// Replicas      *float64
	// Port          *float64
	// ContainerPort *float64
}

type CompositionBase struct {
	ApiVersion *string
}

func NewComposition(scope constructs.Construct, id *string, props *CompositionProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, id)

	// replicas := props.Replicas
	// if replicas == nil {
	// 	replicas = jsii.Number(1)
	// }

	// port := props.Port
	// if port == nil {
	// 	port = jsii.Number(80)
	// }

	// containerPort := props.ContainerPort
	// if containerPort == nil {
	// 	containerPort = jsii.Number(8080)
	// }
	label := map[string]*string{
		"provider": props.Provider,
		"db":       props.DB,
	}
	vpc := ec2.NewVpc(construct, jsii.String("vpc"), &ec2.VpcProps{
		Spec: &ec2.VpcSpec{
			ForProvider: &ec2.VpcSpecForProvider{
				Region:             jsii.String("us-east-1"),
				CidrBlock:          jsii.String("10.0.0.0/16"),
				EnableDnsSupport:   jsii.Bool(true),
				EnableDnsHostNames: jsii.Bool(true),
			},
		},
	})
	composition.NewComposition(construct, jsii.String("composition"), &composition.CompositionProps{
		Metadata: &cdk8s.ApiObjectMetadata{Labels: &label},
		Spec: &composition.CompositionSpec{
			WriteConnectionSecretsToNamespace: jsii.String("crossplane-system"),
			CompositeTypeRef: &composition.CompositionSpecCompositeTypeRef{
				ApiVersion: jsii.String("devopstoolkitseries.com/v1alpha1"),
				Kind:       jsii.String("SQL"),
			},
			PatchSets: &[]*composition.CompositionSpecPatchSets{
				{
					Name: jsii.String("name"),
					Patches: &[]*composition.CompositionSpecPatchSetsPatches{
						{FromFieldPath: jsii.String("metadata.labels")},
					},
				},
			},
			// TODO: Move out
			Resources: &[]*composition.CompositionSpecResources{
				{
					Name: jsii.String("vpc"),
					Base: vpc.ToJson(),
					Patches: &[]*composition.CompositionSpecResourcesPatches{
						{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("metadata.name")},
					},
				},
			},
			// - name: vpc
			// base:
			//   apiVersion: ec2.aws.crossplane.io/v1beta1
			//   kind: VPC
			//   spec:
			// 	forProvider:
			// 	  region: us-east-1
			// 	  cidrBlock: 10.0.0.0/16
			// 	  enableDnsSupport: true
			// 	  enableDnsHostNames: true
			// patches:
			// - fromFieldPath: spec.id
			//   toFieldPath: metadata.name
		},
	})

	// k8s.NewKubeService(construct, jsii.String("service"), &k8s.KubeServiceProps{
	// 	Metadata: &k8s.ObjectMeta{Labels: &label},
	// 	Spec: &k8s.ServiceSpec{
	// 		Type: jsii.String("LoadBalancer"),
	// 		Ports: &[]*k8s.ServicePort{{
	// 			Port:       port,
	// 			TargetPort: k8s.IntOrString_FromNumber(containerPort),
	// 		}},
	// 		Selector: &label,
	// 	},
	// })

	// k8s.NewKubeDeployment(construct, jsii.String("deployment"), &k8s.KubeDeploymentProps{
	// 	Metadata: &k8s.ObjectMeta{Labels: &label},
	// 	Spec: &k8s.DeploymentSpec{
	// 		Replicas: replicas,
	// 		Selector: &k8s.LabelSelector{MatchLabels: &label},
	// 		Template: &k8s.PodTemplateSpec{
	// 			Metadata: &k8s.ObjectMeta{Labels: &label},
	// 			Spec: &k8s.PodSpec{
	// 				Containers: &[]*k8s.Container{{
	// 					Name:  jsii.String("web"),
	// 					Image: props.Image,
	// 					Ports: &[]*k8s.ContainerPort{{ContainerPort: containerPort}},
	// 				}},
	// 			},
	// 		},
	// 	},
	// })

	return construct
}
