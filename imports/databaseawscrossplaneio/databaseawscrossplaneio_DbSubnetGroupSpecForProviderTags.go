// databaseawscrossplaneio
package databaseawscrossplaneio


// Tag is a metadata assigned to an Amazon RDS resource consisting of a key-value pair.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/Tag
type DbSubnetGroupSpecForProviderTags struct {
	// A key is the required name of the tag.
	//
	// The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with "aws:" or "rds:". The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value is the optional value of the tag.
	//
	// The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with "aws:" or "rds:". The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Value *string `field:"optional" json:"value" yaml:"value"`
}

