// pkgcrossplaneio
package pkgcrossplaneio


// Source for the environment variable's value.
//
// Cannot be used if value is not empty.
type ControllerConfigSpecEnvValueFrom struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *ControllerConfigSpecEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	FieldRef *ControllerConfigSpecEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	ResourceFieldRef *ControllerConfigSpecEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// Selects a key of a secret in the pod's namespace.
	SecretKeyRef *ControllerConfigSpecEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

