package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)
	NewComposition(
		chart,
		jsii.String("sql"), &CompositionProps{
			Provider: jsii.String("sql"),
			DB:       jsii.String("postgresql"),
			Kind:     jsii.String("SQL"),
		},
		GetAwsResources(),
	)
	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "cdk8s-demo", nil)
	app.Synth()
}
