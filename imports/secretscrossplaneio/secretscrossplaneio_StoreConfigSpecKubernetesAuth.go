// secretscrossplaneio
package secretscrossplaneio


// Credentials used to connect to the Kubernetes API.
type StoreConfigSpecKubernetesAuth struct {
	// Source of the credentials.
	Source StoreConfigSpecKubernetesAuthSource `field:"required" json:"source" yaml:"source"`
	// Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.
	Env *StoreConfigSpecKubernetesAuthEnv `field:"optional" json:"env" yaml:"env"`
	// Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.
	Fs *StoreConfigSpecKubernetesAuthFs `field:"optional" json:"fs" yaml:"fs"`
	// A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.
	SecretRef *StoreConfigSpecKubernetesAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

