// databaseawscrossplaneio
package databaseawscrossplaneio


// BucketNameSelector selects a reference to a Bucket used to set BucketName.
type RdsInstanceSpecForProviderRestoreFromS3BucketNameSelector struct {
	// MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.
	MatchControllerRef *bool `field:"optional" json:"matchControllerRef" yaml:"matchControllerRef"`
	// MatchLabels ensures an object with matching labels is selected.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
	// Policies for selection.
	Policy *RdsInstanceSpecForProviderRestoreFromS3BucketNameSelectorPolicy `field:"optional" json:"policy" yaml:"policy"`
}

