// databaseawscrossplaneio
package databaseawscrossplaneio


// MasterPasswordSecretRef references the secret that contains the password used in the creation of this RDS instance.
//
// If no reference is given, a password will be auto-generated.
type RdsInstanceSpecForProviderMasterPasswordSecretRef struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the secret.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

