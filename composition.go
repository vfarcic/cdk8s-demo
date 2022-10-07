package main

import (
	composition "example.com/cdk8s-demo/imports/apiextensionscrossplaneio"

	// subnet "example.com/cdk8s-demo/imports/subnet/ec2awscrossplaneio"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CompositionProps struct {
	Provider *string
	DB       *string
	Kind     *string
}

func NewComposition(scope constructs.Construct, id *string, props *CompositionProps, resources *[]*composition.CompositionSpecResources) constructs.Construct {
	construct := constructs.NewConstruct(scope, id)
	label := map[string]*string{
		"provider": props.Provider,
		"db":       props.DB,
	}
	composition.NewComposition(construct, jsii.String("composition"), &composition.CompositionProps{
		Metadata: &cdk8s.ApiObjectMetadata{Labels: &label},
		Spec: &composition.CompositionSpec{
			WriteConnectionSecretsToNamespace: jsii.String("crossplane-system"),
			CompositeTypeRef: &composition.CompositionSpecCompositeTypeRef{
				ApiVersion: jsii.String("devopstoolkitseries.com/v1alpha1"),
				Kind:       jsii.String(*props.Kind),
			},
			PatchSets: &[]*composition.CompositionSpecPatchSets{
				{
					Name: jsii.String("name"),
					Patches: &[]*composition.CompositionSpecPatchSetsPatches{
						{FromFieldPath: jsii.String("metadata.labels")},
					},
				},
			},
			Resources: resources,
		},
	})

	return construct
}
