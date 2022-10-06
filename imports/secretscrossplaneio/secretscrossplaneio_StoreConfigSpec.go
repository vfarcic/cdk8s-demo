// secretscrossplaneio
package secretscrossplaneio


// A StoreConfigSpec defines the desired state of a StoreConfig.
type StoreConfigSpec struct {
	// DefaultScope used for scoping secrets for "cluster-scoped" resources.
	//
	// If store type is "Kubernetes", this would mean the default namespace to store connection secrets for cluster scoped resources. In case of "Vault", this would be used as the default parent path. Typically, should be set as Crossplane installation namespace.
	DefaultScope *string `field:"required" json:"defaultScope" yaml:"defaultScope"`
	// Kubernetes configures a Kubernetes secret store.
	//
	// If the "type" is "Kubernetes" but no config provided, in cluster config will be used.
	Kubernetes *StoreConfigSpecKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Type configures which secret store to be used.
	//
	// Only the configuration block for this store will be used and others will be ignored if provided. Default is Kubernetes.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Vault configures a Vault secret store.
	Vault *StoreConfigSpecVault `field:"optional" json:"vault" yaml:"vault"`
}

