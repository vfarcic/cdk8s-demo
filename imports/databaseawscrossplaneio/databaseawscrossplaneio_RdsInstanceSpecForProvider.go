// databaseawscrossplaneio
package databaseawscrossplaneio


// RDSInstanceParameters define the desired state of an AWS Relational Database Service instance.
type RdsInstanceSpecForProvider struct {
	// DBInstanceClass is the compute and memory capacity of the DB instance, for example, db.m4.large. Not all DB instance classes are available in all AWS Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see DB Instance Class (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide.
	DbInstanceClass *string `field:"required" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// Engine is the name of the database engine to be used for this instance.
	//
	// Not every database engine is available for every AWS Region. Valid Values: * aurora (for MySQL 5.6-compatible Aurora) * aurora-mysql (for MySQL 5.7-compatible Aurora) * aurora-postgresql * mariadb * mysql * oracle-ee * oracle-se2 * oracle-se1 * oracle-se * postgres * sqlserver-ee * sqlserver-se * sqlserver-ex * sqlserver-web Engine is a required field
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// AllocatedStorage is the amount of storage (in gibibytes) to allocate for the DB instance.
	//
	// Type: Integer Amazon Aurora Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume. MySQL Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1): Must be an integer from 100 to 16384. * Magnetic storage (standard): Must be an integer from 5 to 3072. MariaDB Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1): Must be an integer from 100 to 16384. * Magnetic storage (standard): Must be an integer from 5 to 3072. PostgreSQL Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1): Must be an integer from 100 to 16384. * Magnetic storage (standard): Must be an integer from 5 to 3072. Oracle Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1): Must be an integer from 100 to 16384. * Magnetic storage (standard): Must be an integer from 10 to 3072. SQL Server Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2): Enterprise and Standard editions: Must be an integer from 200 to 16384. Web and Express editions: Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must be an integer from 200 to 16384. Web and Express editions: Must be an integer from 100 to 16384. * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 200 to 1024. Web and Express editions: Must be an integer from 20 to 1024.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// AllowMajorVersionUpgrade indicates that major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible. Constraints: This parameter must be set to true when specifying a value for the EngineVersion parameter that is a different major version than the DB instance's current version.
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// ApplyModificationsImmediately specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the PreferredMaintenanceWindow setting for the DB instance.
	//
	// If this parameter is set to false, changes to the DB instance are applied during the next maintenance window. Some parameter changes can cause an outage and are applied on the next call to RebootDBInstance, or the next failure reboot. Review the table of parameters in Modifying a DB Instance and Using the Apply Immediately Parameter (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html) in the Amazon RDS User Guide. to see the impact that setting ApplyImmediately to true or false has for each modified parameter and to determine when the changes are applied. Default: false
	ApplyModificationsImmediately *bool `field:"optional" json:"applyModificationsImmediately" yaml:"applyModificationsImmediately"`
	// AutoMinorVersionUpgrade indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	//
	// Default: true.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// AvailabilityZone is the EC2 Availability Zone that the DB instance is created in.
	//
	// For information on AWS Regions and Availability Zones, see Regions and Availability Zones (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html). Default: A random, system-chosen Availability Zone in the endpoint's AWS Region. Example: us-east-1d Constraint: The AvailabilityZone parameter can't be specified if the MultiAZ parameter is set to true. The specified Availability Zone must be in the same AWS Region as the current endpoint.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// BackupRetentionPeriod is the number of days for which automated backups are retained.
	//
	// Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. Amazon Aurora Not applicable. The retention period for automated backups is managed by the DB cluster. For more information, see CreateDBCluster. Default: 1 Constraints: * Must be a value from 0 to 35 * Cannot be set to 0 if the DB instance is a source to Read Replicas
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// CACertificateIdentifier indicates the certificate that needs to be associated with the instance.
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// CharacterSetName indicates that the DB instance should be associated with the specified CharacterSet for supported engines, Amazon Aurora Not applicable.
	//
	// The character set is managed by the DB cluster. For more information, see CreateDBCluster.
	CharacterSetName *string `field:"optional" json:"characterSetName" yaml:"characterSetName"`
	// CloudwatchLogsExportConfiguration is the configuration setting for the log types to be enabled for export to CloudWatch Logs for a specific DB instance.
	CloudwatchLogsExportConfiguration *RdsInstanceSpecForProviderCloudwatchLogsExportConfiguration `field:"optional" json:"cloudwatchLogsExportConfiguration" yaml:"cloudwatchLogsExportConfiguration"`
	// CopyTagsToSnapshot should be true to copy all tags from the DB instance to snapshots of the DB instance, and otherwise false.
	//
	// The default is false.
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// DBClusterIdentifier is the identifier of the DB cluster that the instance will belong to.
	//
	// For information on creating a DB cluster, see CreateDBCluster. Type: String
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// DBName is the meaning of this parameter differs according to the database engine you use.
	//
	// Type: String MySQL The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Cannot be a word reserved by the specified database engine MariaDB The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Cannot be a word reserved by the specified database engine PostgreSQL The name of the database to create when the DB instance is created. If this parameter is not specified, the default "postgres" database is created in the DB instance. Constraints: * Must contain 1 to 63 letters, numbers, or underscores. * Must begin with a letter or an underscore. Subsequent characters can be letters, underscores, or digits (0-9). * Cannot be a word reserved by the specified database engine Oracle The Oracle System ID (SID) of the created DB instance. If you specify null, the default value ORCL is used. You can't specify the string NULL, or any other reserved word, for DBName. Default: ORCL Constraints: * Cannot be longer than 8 characters SQL Server Not applicable. Must be null. Amazon Aurora The name of the database to create when the primary instance of the DB cluster is created. If this parameter is not specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Cannot be a word reserved by the specified database engine
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// DBParameterGroupName is the name of the DB parameter group to associate with this DB instance.
	//
	// If this argument is omitted, the default DBParameterGroup for the specified engine is used. Constraints: * Must be 1 to 255 letters, numbers, or hyphens. * First character must be a letter * Cannot end with a hyphen or contain two consecutive hyphens
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// DBSecurityGroups is a list of DB security groups to associate with this DB instance.
	//
	// Default: The default DB security group for the database engine.
	DbSecurityGroups *[]*string `field:"optional" json:"dbSecurityGroups" yaml:"dbSecurityGroups"`
	// DBSubnetGroupName is a DB subnet group to associate with this DB instance.
	//
	// If there is no DB subnet group, then it is a non-VPC DB instance.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// DBSubnetGroupNameRef is a reference to a DBSubnetGroup used to set DBSubnetGroupName.
	DbSubnetGroupNameRef *RdsInstanceSpecForProviderDbSubnetGroupNameRef `field:"optional" json:"dbSubnetGroupNameRef" yaml:"dbSubnetGroupNameRef"`
	// DBSubnetGroupNameSelector selects a reference to a DBSubnetGroup used to set DBSubnetGroupName.
	DbSubnetGroupNameSelector *RdsInstanceSpecForProviderDbSubnetGroupNameSelector `field:"optional" json:"dbSubnetGroupNameSelector" yaml:"dbSubnetGroupNameSelector"`
	// DeleteAutomatedBackups indicates whether to remove automated backups immediately after the DB instance is deleted.
	//
	// The default is to remove automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups *bool `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// DeletionProtection indicates if the DB instance should have deletion protection enabled.
	//
	// The database can't be deleted when this value is set to true. The default is false. For more information, see  Deleting a DB Instance (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Domain specifies the Active Directory Domain to create the instance in.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// DomainIAMRoleName specifies the name of the IAM role to be used when making API calls to the Directory Service.
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// DomainIAMRoleNameRef is a reference to an IAMRole used to set DomainIAMRoleName.
	DomainIamRoleNameRef *RdsInstanceSpecForProviderDomainIamRoleNameRef `field:"optional" json:"domainIamRoleNameRef" yaml:"domainIamRoleNameRef"`
	// DomainIAMRoleNameSelector selects a reference to an IAMRole used to set DomainIAMRoleName.
	DomainIamRoleNameSelector *RdsInstanceSpecForProviderDomainIamRoleNameSelector `field:"optional" json:"domainIamRoleNameSelector" yaml:"domainIamRoleNameSelector"`
	// EnableCloudwatchLogsExports is the list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see Publishing Database Logs to Amazon CloudWatch Logs  (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon Relational Database Service User Guide.
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// EnableIAMDatabaseAuthentication should be true to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts, and otherwise false.
	//
	// You can enable IAM database authentication for the following database engines: Amazon Aurora Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster. For more information, see CreateDBCluster. MySQL * For MySQL 5.6, minor version 5.6.34 or higher * For MySQL 5.7, minor version 5.7.16 or higher Default: false
	EnableIamDatabaseAuthentication *bool `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// EnablePerformanceInsights should be true to enable Performance Insights for the DB instance, and otherwise false.
	//
	// For more information, see Using Amazon Performance Insights (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon Relational Database Service User Guide.
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// EngineVersion is the version number of the database engine to use.
	//
	// For a list of valid engine versions, call DescribeDBEngineVersions. The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every AWS Region. Amazon Aurora Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster. For more information, see CreateDBCluster. MariaDB See MariaDB on Amazon RDS Versions (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide. Microsoft SQL Server See Version and Feature Support on Amazon RDS (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.FeatureSupport) in the Amazon RDS User Guide. MySQL See MySQL on Amazon RDS Versions (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide. Oracle See Oracle Database Engine Release Notes (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide. PostgreSQL See Supported PostgreSQL Database Versions (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts.General.DBVersions) in the Amazon RDS User Guide.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The DBSnapshotIdentifier of the new DBSnapshot created when SkipFinalSnapshot is set to false.
	//
	// Specifying this parameter and also setting the SkipFinalShapshot parameter to true results in an error. Constraints: * Must be 1 to 255 letters or numbers. * First character must be a letter * Cannot end with a hyphen or contain two consecutive hyphens * Cannot be specified when deleting a Read Replica.
	FinalDbSnapshotIdentifier *string `field:"optional" json:"finalDbSnapshotIdentifier" yaml:"finalDbSnapshotIdentifier"`
	// IOPS is the amount of Provisioned IOPS (input/output operations per second) to be initially allocated for the DB instance.
	//
	// For information about valid IOPS values, see see Amazon RDS Provisioned IOPS Storage to Improve Performance (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the Amazon RDS User Guide. Constraints: Must be a multiple between 1 and 50 of the storage amount for the DB instance. Must also be an integer multiple of 1000. For example, if the size of your DB instance is 500 GiB, then your IOPS value can be 2000, 3000, 4000, or 5000.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// KMSKeyID for an encrypted DB instance.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption key. If you are creating a DB instance with the same AWS account that owns the KMS encryption key used to encrypt the new DB instance, then you can use the KMS key alias instead of the ARN for the KM encryption key. Amazon Aurora Not applicable. The KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster. If the StorageEncrypted parameter is true, and you do not specify a value for the KMSKeyID parameter, then Amazon RDS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS Region.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// LicenseModel information for this DB instance.
	//
	// Valid values: license-included | bring-your-own-license | general-public-license.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// MasterPasswordSecretRef references the secret that contains the password used in the creation of this RDS instance.
	//
	// If no reference is given, a password will be auto-generated.
	MasterPasswordSecretRef *RdsInstanceSpecForProviderMasterPasswordSecretRef `field:"optional" json:"masterPasswordSecretRef" yaml:"masterPasswordSecretRef"`
	// MasterUsername is the name for the master user.
	//
	// Amazon Aurora Not applicable. The name for the master user is managed by the DB cluster. For more information, see CreateDBCluster. MariaDB Constraints: * Required for MariaDB. * Must be 1 to 16 letters or numbers. * Cannot be a reserved word for the chosen database engine. Microsoft SQL Server Constraints: * Required for SQL Server. * Must be 1 to 128 letters or numbers. * The first character must be a letter. * Cannot be a reserved word for the chosen database engine. MySQL Constraints: * Required for MySQL. * Must be 1 to 16 letters or numbers. * First character must be a letter. * Cannot be a reserved word for the chosen database engine. Oracle Constraints: * Required for Oracle. * Must be 1 to 30 letters or numbers. * First character must be a letter. * Cannot be a reserved word for the chosen database engine. PostgreSQL Constraints: * Required for PostgreSQL. * Must be 1 to 63 letters or numbers. * First character must be a letter. * Cannot be a reserved word for the chosen database engine.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.
	//
	// For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide.
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// MonitoringInterval is the interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	//
	// To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. If MonitoringRoleARN is specified, then you must also set MonitoringInterval to a value other than 0. Valid Values: 0, 1, 5, 10, 15, 30, 60
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// MonitoringRoleARN is the ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
	//
	// For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, go to Setting Up and Enabling Enhanced Monitoring (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleARN value.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// MonitoringRoleARNRef is a reference to an IAMRole used to set MonitoringRoleARN.
	MonitoringRoleArnRef *RdsInstanceSpecForProviderMonitoringRoleArnRef `field:"optional" json:"monitoringRoleArnRef" yaml:"monitoringRoleArnRef"`
	// MonitoringRoleARNSelector selects a reference to an IAMRole used to set MonitoringRoleARN.
	MonitoringRoleArnSelector *RdsInstanceSpecForProviderMonitoringRoleArnSelector `field:"optional" json:"monitoringRoleArnSelector" yaml:"monitoringRoleArnSelector"`
	// MultiAZ specifies if the DB instance is a Multi-AZ deployment.
	//
	// You can't set the AvailabilityZone parameter if the MultiAZ parameter is set to true.
	MultiAz *bool `field:"optional" json:"multiAz" yaml:"multiAz"`
	// OptionGroupName indicates that the DB instance should be associated with the specified option group.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group, and that option group can't be removed from a DB instance once it is associated with a DB instance.
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// PerformanceInsightsKMSKeyID is the AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS encryption key.
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// PerformanceInsightsRetentionPeriod is the amount of time, in days, to retain Performance Insights data.
	//
	// Valid values are 7 or 731 (2 years).
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// Port number on which the database accepts connections.
	//
	// MySQL Default: 3306 Valid Values: 1150-65535 Type: Integer MariaDB Default: 3306 Valid Values: 1150-65535 Type: Integer PostgreSQL Default: 5432 Valid Values: 1150-65535 Type: Integer Oracle Default: 1521 Valid Values: 1150-65535 SQL Server Default: 1433 Valid Values: 1150-65535 except for 1434, 3389, 47001, 49152, and 49152 through 49156. Amazon Aurora Default: 3306 Valid Values: 1150-65535 Type: Integer
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// PreferredBackupWindow is the daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter.
	//
	// For more information, see The Backup Window (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide. Amazon Aurora Not applicable. The daily time range for creating automated backups is managed by the DB cluster. For more information, see CreateDBCluster. The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region. To see the time blocks available, see  Adjusting the Preferred DB Instance Maintenance Window (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow) in the Amazon RDS User Guide. Constraints: * Must be in the format hh24:mi-hh24:mi. * Must be in Universal Coordinated Time (UTC). * Must not conflict with the preferred maintenance window. * Must be at least 30 minutes.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// PreferredMaintenanceWindow is the time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// For more information, see Amazon RDS Maintenance Window (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance). Format: ddd:hh24:mi-ddd:hh24:mi The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun. Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// ProcessorFeatures is the number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
	ProcessorFeatures *[]*RdsInstanceSpecForProviderProcessorFeatures `field:"optional" json:"processorFeatures" yaml:"processorFeatures"`
	// PromotionTier specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
	//
	// For more information, see  Fault Tolerance for an Aurora DB Cluster (http://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide. Default: 1 Valid Values: 0 - 15
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
	// PubliclyAccessible specifies the accessibility options for the DB instance.
	//
	// A value of true specifies an Internet-facing instance with a publicly resolvable DNS name, which resolves to a public IP address. A value of false specifies an internal instance with a DNS name that resolves to a private IP address. Default: The default behavior varies depending on whether DBSubnetGroupName is specified. If DBSubnetGroupName is not specified, and PubliclyAccessible is not specified, the following applies: * If the default VPC in the target region doesn’t have an Internet gateway attached to it, the DB instance is private. * If the default VPC in the target region has an Internet gateway attached to it, the DB instance is public. If DBSubnetGroupName is specified, and PubliclyAccessible is not specified, the following applies: * If the subnets are part of a VPC that doesn’t have an Internet gateway attached to it, the DB instance is private. * If the subnets are part of a VPC that has an Internet gateway attached to it, the DB instance is public.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region is the region you'd like your RDSInstance to be created in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// RestoreFrom specifies the details of the backup to restore when creating a new RDS instance.
	//
	// (If the RDS instance already exists, this property will be ignored.)
	RestoreFrom *RdsInstanceSpecForProviderRestoreFrom `field:"optional" json:"restoreFrom" yaml:"restoreFrom"`
	// ScalingConfiguration is the scaling properties of the DB cluster.
	//
	// You can only modify scaling properties for DB clusters in serverless DB engine mode.
	ScalingConfiguration *RdsInstanceSpecForProviderScalingConfiguration `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// Determines whether a final DB snapshot is created before the DB instance is deleted.
	//
	// If true is specified, no DBSnapshot is created. If false is specified, a DB snapshot is created before the DB instance is deleted. Note that when a DB instance is in a failure state and has a status of 'failed', 'incompatible-restore', or 'incompatible-network', it can only be deleted when the SkipFinalSnapshotBeforeDeletion parameter is set to "true". Specify true when deleting a Read Replica. The FinalDBSnapshotIdentifier parameter must be specified if SkipFinalSnapshotBeforeDeletion is false. Default: false
	SkipFinalSnapshotBeforeDeletion *bool `field:"optional" json:"skipFinalSnapshotBeforeDeletion" yaml:"skipFinalSnapshotBeforeDeletion"`
	// StorageEncrypted specifies whether the DB instance is encrypted.
	//
	// Amazon Aurora Not applicable. The encryption for DB instances is managed by the DB cluster. For more information, see CreateDBCluster. Default: false
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// StorageType specifies the storage type to be associated with the DB instance.
	//
	// Valid values: standard | gp2 | io1 If you specify io1, you must also include a value for the IOPS parameter. Default: io1 if the IOPS parameter is specified, otherwise standard
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Tags.
	//
	// For more information, see Tagging Amazon RDS Resources (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the Amazon RDS User Guide.
	Tags *[]*RdsInstanceSpecForProviderTags `field:"optional" json:"tags" yaml:"tags"`
	// Timezone of the DB instance.
	//
	// The time zone parameter is currently supported only by Microsoft SQL Server (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// A value that specifies that the DB instance class of the DB instance uses its default processor features.
	UseDefaultProcessorFeatures *bool `field:"optional" json:"useDefaultProcessorFeatures" yaml:"useDefaultProcessorFeatures"`
	// VPCSecurityGroupIDRefs are references to VPCSecurityGroups used to set the VPCSecurityGroupIDs.
	VpcSecurityGroupIdRefs *[]*RdsInstanceSpecForProviderVpcSecurityGroupIdRefs `field:"optional" json:"vpcSecurityGroupIdRefs" yaml:"vpcSecurityGroupIdRefs"`
	// VPCSecurityGroupIDs is a list of EC2 VPC security groups to associate with this DB instance.
	//
	// Amazon Aurora Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster. For more information, see CreateDBCluster. Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// VPCSecurityGroupIDSelector selects references to VPCSecurityGroups used to set the VPCSecurityGroupIDs.
	VpcSecurityGroupIdSelector *RdsInstanceSpecForProviderVpcSecurityGroupIdSelector `field:"optional" json:"vpcSecurityGroupIdSelector" yaml:"vpcSecurityGroupIdSelector"`
}

