// databaseawscrossplaneio
package databaseawscrossplaneio


// MonitoringRoleARNRef is a reference to an IAMRole used to set MonitoringRoleARN.
type RdsInstanceSpecForProviderMonitoringRoleArnRef struct {
	// Name of the referenced object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policies for referencing.
	Policy *RdsInstanceSpecForProviderMonitoringRoleArnRefPolicy `field:"optional" json:"policy" yaml:"policy"`
}

