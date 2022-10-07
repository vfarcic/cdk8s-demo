// ec2awscrossplaneio
package ec2awscrossplaneio


// Tag defines a tag.
type InternetGatewaySpecForProviderTags struct {
	// Key is the name of the tag.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value is the value of the tag.
	Value *string `field:"required" json:"value" yaml:"value"`
}

