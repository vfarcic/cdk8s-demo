// ec2awscrossplaneio
package ec2awscrossplaneio


// PrefixListID describes a prefix list ID.
type SecurityGroupSpecForProviderEgressPrefixListIds struct {
	// The ID of the prefix.
	PrefixListId *string `field:"required" json:"prefixListId" yaml:"prefixListId"`
	// A description for the security group rule that references this prefix list ID.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description *string `field:"optional" json:"description" yaml:"description"`
}

