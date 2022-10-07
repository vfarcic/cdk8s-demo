// databaseawscrossplaneio
package databaseawscrossplaneio


// ScalingConfiguration is the scaling properties of the DB cluster.
//
// You can only modify scaling properties for DB clusters in serverless DB engine mode.
type RdsInstanceSpecForProviderScalingConfiguration struct {
	// AutoPause specifies whether to allow or disallow automatic pause for an Aurora DB cluster in serverless DB engine mode.
	//
	// A DB cluster can be paused only when it's idle (it has no connections). If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it.
	AutoPause *bool `field:"optional" json:"autoPause" yaml:"autoPause"`
	// MaxCapacity is the maximum capacity for an Aurora DB cluster in serverless DB engine mode.
	//
	// Valid capacity values are 2, 4, 8, 16, 32, 64, 128, and 256. The maximum capacity must be greater than or equal to the minimum capacity.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// MinCapacity is the minimum capacity for an Aurora DB cluster in serverless DB engine mode.
	//
	// Valid capacity values are 2, 4, 8, 16, 32, 64, 128, and 256. The minimum capacity must be less than or equal to the maximum capacity.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// SecondsUntilAutoPause is the time, in seconds, before an Aurora DB cluster in serverless mode is paused.
	SecondsUntilAutoPause *float64 `field:"optional" json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
}

