package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func main() {
	app := cdk8s.NewApp(nil)
	NewSillyDemo(app, "silly-demo")
	NewAws(app, "aws")
	app.Synth()
}

func NewSillyDemo(scope constructs.Construct, id string) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(id), nil)
	NewApp(
		chart,
		jsii.String("sql"),
		&AppProps{
			Image:         jsii.String("vfarcic/silly-demo"),
			Replicas:      jsii.Number(3),
			Port:          jsii.Number(8080),
			ContainerPort: jsii.Number(8080),
		},
	)
	return chart
}

func NewAws(scope constructs.Construct, id string) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(id), nil)
	NewComposition(
		chart,
		jsii.String("sql"),
		&CompositionProps{
			Provider: jsii.String("sql"),
			DB:       jsii.String("postgresql"),
			Kind:     jsii.String("SQL"),
		},
		GetAwsResources(),
	)
	return chart
}
