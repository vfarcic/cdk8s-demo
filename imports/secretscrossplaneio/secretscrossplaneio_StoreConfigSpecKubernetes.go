// secretscrossplaneio
package secretscrossplaneio


// Kubernetes configures a Kubernetes secret store.
//
// If the "type" is "Kubernetes" but no config provided, in cluster config will be used.
type StoreConfigSpecKubernetes struct {
	// Credentials used to connect to the Kubernetes API.
	Auth *StoreConfigSpecKubernetesAuth `field:"required" json:"auth" yaml:"auth"`
}

