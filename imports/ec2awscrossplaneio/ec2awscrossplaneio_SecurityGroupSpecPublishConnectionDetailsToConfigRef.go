// ec2awscrossplaneio
package ec2awscrossplaneio


// SecretStoreConfigRef specifies which secret store config should be used for this ConnectionSecret.
type SecurityGroupSpecPublishConnectionDetailsToConfigRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *SecurityGroupSpecPublishConnectionDetailsToConfigRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

