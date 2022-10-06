// secretscrossplaneio
package secretscrossplaneio


// Source of the credentials.
type StoreConfigSpecKubernetesAuthSource string

const (
	// None.
	StoreConfigSpecKubernetesAuthSource_NONE StoreConfigSpecKubernetesAuthSource = "NONE"
	// Secret.
	StoreConfigSpecKubernetesAuthSource_SECRET StoreConfigSpecKubernetesAuthSource = "SECRET"
	// Environment.
	StoreConfigSpecKubernetesAuthSource_ENVIRONMENT StoreConfigSpecKubernetesAuthSource = "ENVIRONMENT"
	// Filesystem.
	StoreConfigSpecKubernetesAuthSource_FILESYSTEM StoreConfigSpecKubernetesAuthSource = "FILESYSTEM"
)

