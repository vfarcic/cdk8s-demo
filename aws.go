package main

import (
	composition "example.com/cdk8s-demo/imports/apiextensionscrossplaneio"
	db "example.com/cdk8s-demo/imports/databaseawscrossplaneio"
	ec2 "example.com/cdk8s-demo/imports/ec2awscrossplaneio"
	"github.com/aws/jsii-runtime-go"
)

const region string = "us-east-1"

func GetAwsResources() *[]*composition.CompositionSpecResources {
	return &[]*composition.CompositionSpecResources{
		getAwsVpc(),
		getAwsSubnet("a", "10.0.0.0/24"),
		getAwsSubnet("b", "10.0.1.0/24"),
		getAwsSubnet("c", "10.0.2.0/24"),
		getAwsDbSubnetGroup(),
		getAwsInternetGateway(),
		getAwsRouteTable(),
		getAwsSecurityGroup(),
		getAwsRdsInstance(),
		getPgProviderConfig(),
		getPgDatabase(),
	}
}

func getAwsVpc() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("vpc"),
		Base: ec2.Vpc_Manifest(&ec2.VpcProps{
			Spec: &ec2.VpcSpec{
				ForProvider: &ec2.VpcSpecForProvider{
					Region:             jsii.String(region),
					CidrBlock:          jsii.String("10.0.0.0/16"),
					EnableDnsSupport:   jsii.Bool(true),
					EnableDnsHostNames: jsii.Bool(true),
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("metadata.name"),
			},
		},
	}
}

func getAwsSubnet(zone, cidrBlock string) *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("subnet-" + zone),
		Base: ec2.Subnet_Manifest(&ec2.SubnetProps{
			Spec: &ec2.SubnetSpec{
				ForProvider: &ec2.SubnetSpecForProvider{
					Region:           jsii.String(region),
					CidrBlock:        jsii.String(cidrBlock),
					AvailabilityZone: jsii.String(region + zone),
					VpcIdSelector: &ec2.SubnetSpecForProviderVpcIdSelector{
						MatchControllerRef: jsii.Bool(true),
					},
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("metadata.name"),
				Transforms: &[]*composition.CompositionSpecResourcesPatchesTransforms{
					{
						Type:   composition.CompositionSpecResourcesPatchesTransformsType_STRING,
						String: &composition.CompositionSpecResourcesPatchesTransformsString{Fmt: jsii.String("%s-" + zone)},
					},
				},
			},
			{
				FromFieldPath: jsii.String("spec.forProvider.availabilityZone"),
				ToFieldPath:   jsii.String("metadata.labels.zone"),
			},
		},
	}
}

func getAwsDbSubnetGroup() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("dbsubnetgroup"),
		Base: db.DbSubnetGroup_Manifest(&db.DbSubnetGroupProps{
			Spec: &db.DbSubnetGroupSpec{
				ForProvider: &db.DbSubnetGroupSpecForProvider{
					Region:      jsii.String(region),
					Description: jsii.String("I'm too lazy to write a good description"),
					SubnetIdSelector: &db.DbSubnetGroupSpecForProviderSubnetIdSelector{
						MatchControllerRef: jsii.Bool(true),
					},
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("metadata.name")},
		},
	}
}

func getAwsInternetGateway() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("gateway"),
		Base: ec2.InternetGateway_Manifest(&ec2.InternetGatewayProps{
			Spec: &ec2.InternetGatewaySpec{
				ForProvider: &ec2.InternetGatewaySpecForProvider{
					Region: jsii.String(region),
					VpcIdSelector: &ec2.InternetGatewaySpecForProviderVpcIdSelector{
						MatchControllerRef: jsii.Bool(true),
					},
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("metadata.name")},
		},
	}
}

func getAwsRouteTable() *composition.CompositionSpecResources {
	associations := []*ec2.RouteTableSpecForProviderAssociations{}
	for _, zone := range []string{"a", "b", "c"} {
		associations = append(associations, &ec2.RouteTableSpecForProviderAssociations{
			SubnetIdSelector: &ec2.RouteTableSpecForProviderAssociationsSubnetIdSelector{
				MatchControllerRef: jsii.Bool(true),
				MatchLabels: &map[string]*string{
					"zone": jsii.String(region + zone),
				},
			},
		})
	}
	return &composition.CompositionSpecResources{
		Name: jsii.String("routetable"),
		Base: ec2.RouteTable_Manifest(&ec2.RouteTableProps{
			Spec: &ec2.RouteTableSpec{
				ForProvider: &ec2.RouteTableSpecForProvider{
					Region: jsii.String(region),
					VpcIdSelector: &ec2.RouteTableSpecForProviderVpcIdSelector{
						MatchControllerRef: jsii.Bool(true),
					},
					Routes: &[]*ec2.RouteTableSpecForProviderRoutes{
						{
							DestinationCidrBlock: jsii.String("0.0.0.0/0"),
							GatewayIdSelector: &ec2.RouteTableSpecForProviderRoutesGatewayIdSelector{
								MatchControllerRef: jsii.Bool(true),
							},
						},
					},
					Associations: &associations,
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("metadata.name")},
		},
	}
}

func getAwsSecurityGroup() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("securitygroup"),
		Base: ec2.SecurityGroup_Manifest(&ec2.SecurityGroupProps{
			Spec: &ec2.SecurityGroupSpec{
				ForProvider: &ec2.SecurityGroupSpecForProvider{
					Description: jsii.String("I'm too lazy to write a good description"),
					Region:      jsii.String(region),
					VpcIdSelector: &ec2.SecurityGroupSpecForProviderVpcIdSelector{
						MatchControllerRef: jsii.Bool(true),
					},
					GroupName: jsii.String("PATCHED!"),
					Ingress: &[]*ec2.SecurityGroupSpecForProviderIngress{
						{
							FromPort:   jsii.Number(5432),
							ToPort:     jsii.Number(5432),
							IpProtocol: jsii.String("tcp"),
							IpRanges: &[]*ec2.SecurityGroupSpecForProviderIngressIpRanges{
								{
									CidrIp: jsii.String("0.0.0.0/0"),
								},
							},
						},
					},
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("metadata.name")},
			{FromFieldPath: jsii.String("spec.id"), ToFieldPath: jsii.String("spec.forProvider.groupName")},
		},
	}
}

func getAwsRdsInstance() *composition.CompositionSpecResources {
	return &composition.CompositionSpecResources{
		Name: jsii.String("rdsinstance"),
		Base: db.RdsInstance_Manifest(&db.RdsInstanceProps{
			Spec: &db.RdsInstanceSpec{
				ForProvider: &db.RdsInstanceSpecForProvider{
					Region: jsii.String(region),
					DbSubnetGroupNameSelector: &db.RdsInstanceSpecForProviderDbSubnetGroupNameSelector{
						MatchControllerRef: jsii.Bool(true),
					},
					VpcSecurityGroupIdSelector: &db.RdsInstanceSpecForProviderVpcSecurityGroupIdSelector{
						MatchControllerRef: jsii.Bool(true),
					},
					DbInstanceClass:                 jsii.String("db.m5.large"),
					MasterUsername:                  jsii.String("masteruser"),
					Engine:                          jsii.String("postgres"),
					EngineVersion:                   jsii.String("14.1"),
					SkipFinalSnapshotBeforeDeletion: jsii.Bool(true),
					PubliclyAccessible:              jsii.Bool(true),
					AllocatedStorage:                jsii.Number(200),
				},
				WriteConnectionSecretToRef: &db.RdsInstanceSpecWriteConnectionSecretToRef{
					Namespace: jsii.String("crossplane-system"),
					Name:      jsii.String("PATCHED!"),
				},
			},
		}),
		Patches: &[]*composition.CompositionSpecResourcesPatches{
			{
				FromFieldPath: jsii.String("spec.id"),
				ToFieldPath:   jsii.String("metadata.name"),
			}, {
				FromFieldPath: jsii.String("metadata.uid"),
				ToFieldPath:   jsii.String("spec.writeConnectionSecretToRef.name"),
				Transforms: &[]*composition.CompositionSpecResourcesPatchesTransforms{
					{
						Type:   composition.CompositionSpecResourcesPatchesTransformsType_STRING,
						String: &composition.CompositionSpecResourcesPatchesTransformsString{Fmt: jsii.String("%s-sql")},
					},
				},
			}, {
				FromFieldPath: jsii.String("spec.parameters.size"),
				ToFieldPath:   jsii.String("spec.forProvider.dbInstanceClass"),
				Transforms: &[]*composition.CompositionSpecResourcesPatchesTransforms{
					{
						Type: composition.CompositionSpecResourcesPatchesTransformsType_MAP,
						Map: &map[string]*string{
							"small":  jsii.String("db.m5.large"),
							"medium": jsii.String("db.m5.large"),
							"large":  jsii.String("db.m5.8xlarge"),
						},
					},
				},
			}, {
				FromFieldPath: jsii.String("spec.parameters.version"),
				ToFieldPath:   jsii.String("spec.forProvider.engineVersion"),
			},
		},
		ConnectionDetails: &[]*composition.CompositionSpecResourcesConnectionDetails{
			{FromConnectionSecretKey: jsii.String("username")},
			{FromConnectionSecretKey: jsii.String("password")},
			{FromConnectionSecretKey: jsii.String("endpoint")},
			{FromConnectionSecretKey: jsii.String("port")},
		},
	}
}
