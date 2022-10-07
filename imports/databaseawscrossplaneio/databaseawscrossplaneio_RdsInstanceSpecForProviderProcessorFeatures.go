// databaseawscrossplaneio
package databaseawscrossplaneio


// ProcessorFeature is a processor feature entry.
//
// For more information, see Configuring the Processor of the DB Instance Class (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html#USER_ConfigureProcessor) in the Amazon RDS User Guide. Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ProcessorFeature
type RdsInstanceSpecForProviderProcessorFeatures struct {
	// Name of the processor feature.
	//
	// Valid names are coreCount and threadsPerCore.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a processor feature name.
	Value *string `field:"required" json:"value" yaml:"value"`
}

