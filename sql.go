package main

import (
	composition "example.com/cdk8s-demo/imports/apiextensionscrossplaneio"
	pg "example.com/cdk8s-demo/imports/postgresqlsqlcrossplaneio"
	"github.com/aws/jsii-runtime-go"
)

func getPgProviderConfig() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("sql-config"),
		Base: pg.ProviderConfig_Manifest(&pg.ProviderConfigProps{
			Spec: &pg.ProviderConfigSpec{
				Credentials: &pg.ProviderConfigSpecCredentials{
					Source: pg.ProviderConfigSpecCredentialsSource_POSTGRE_SQL_CONNECTION_SECRET,
				},
				SslMode: pg.ProviderConfigSpecSslMode_REQUIRE,
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("metadata.name"),
			}, {
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("spec.credentials.connectionSecretRef.name"),
			}, {
				FromFieldPath: jsii.String("spec.claimRef.namespace"),
				ToFieldPath:   jsii.String("spec.credentials.connectionSecretRef.namespace"),
			},
		},
		ReadinessChecks: &[]*composition.CompositionSpecResourcesReadinessChecks{
			{Type: composition.CompositionSpecResourcesReadinessChecksType_NONE},
		},
	}
}

func getPgDatabase() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("sql-db"),
		Base: pg.Database_Manifest(&pg.DatabaseProps{
			Spec: &pg.DatabaseSpec{
				ForProvider: &pg.DatabaseSpecForProvider{},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("metadata.name"),
			}, {
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("spec.providerConfigRef.name"),
			},
		},
		ReadinessChecks: &[]*composition.CompositionSpecResourcesReadinessChecks{
			{Type: composition.CompositionSpecResourcesReadinessChecksType_NONE},
		},
	}
}
