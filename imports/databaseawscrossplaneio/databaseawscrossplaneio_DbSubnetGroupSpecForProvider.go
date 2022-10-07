// databaseawscrossplaneio
package databaseawscrossplaneio


// DBSubnetGroupParameters define the desired state of an AWS VPC Database Subnet Group.
type DbSubnetGroupSpecForProvider struct {
	// The description for the DB subnet group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Region is the region you'd like your DBSubnetGroup to be created in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// SubnetIDRefs is a set of references that each retrieve the subnetID from the referenced Subnet.
	SubnetIdRefs *[]*DbSubnetGroupSpecForProviderSubnetIdRefs `field:"optional" json:"subnetIdRefs" yaml:"subnetIdRefs"`
	// The EC2 Subnet IDs for the DB subnet group.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// SubnetIDSelector selects a set of references that each retrieve the subnetID from the referenced Subnet.
	SubnetIdSelector *DbSubnetGroupSpecForProviderSubnetIdSelector `field:"optional" json:"subnetIdSelector" yaml:"subnetIdSelector"`
	// A list of tags.
	//
	// For more information, see Tagging Amazon RDS Resources (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the Amazon RDS User Guide.
	Tags *[]*DbSubnetGroupSpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
}

