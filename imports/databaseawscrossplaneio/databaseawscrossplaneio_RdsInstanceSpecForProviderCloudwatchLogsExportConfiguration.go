// databaseawscrossplaneio
package databaseawscrossplaneio


// CloudwatchLogsExportConfiguration is the configuration setting for the log types to be enabled for export to CloudWatch Logs for a specific DB instance.
type RdsInstanceSpecForProviderCloudwatchLogsExportConfiguration struct {
	// DisableLogTypes is the list of log types to disable.
	DisableLogTypes *[]*string `field:"optional" json:"disableLogTypes" yaml:"disableLogTypes"`
	// EnableLogTypes is the list of log types to enable.
	EnableLogTypes *[]*string `field:"optional" json:"enableLogTypes" yaml:"enableLogTypes"`
}

