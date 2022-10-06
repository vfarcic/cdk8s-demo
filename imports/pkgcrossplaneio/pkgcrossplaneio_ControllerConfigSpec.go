// pkgcrossplaneio
package pkgcrossplaneio


// ControllerConfigSpec specifies the configuration for a packaged controller.
//
// Values provided will override package manager defaults. Labels and annotations are passed to both the controller Deployment and ServiceAccount.
type ControllerConfigSpec struct {
	// If specified, the pod's scheduling constraints.
	Affinity *ControllerConfigSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	// Arguments to the entrypoint.
	//
	// The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// List of environment variables to set in the container.
	//
	// Cannot be updated.
	Env *[]*ControllerConfigSpecEnv `field:"optional" json:"env" yaml:"env"`
	// List of sources to populate environment variables in the container.
	//
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	EnvFrom *[]*ControllerConfigSpecEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	// Docker image name.
	//
	// More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Image pull policy.
	//
	// One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	//
	// If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod Setting ImagePullSecrets will replace any secrets that have been propagated to a controller Deployment, typically via packagePullSecrets.
	ImagePullSecrets *[]*ControllerConfigSpecImagePullSecrets `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// Metadata that will be added to the provider Pod.
	Metadata *ControllerConfigSpecMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// NodeName is a request to schedule this pod onto a specific node.
	//
	// If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	//
	// Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	//
	// Optional: Defaults to empty.  See type description for default values of each field.
	PodSecurityContext *ControllerConfigSpecPodSecurityContext `field:"optional" json:"podSecurityContext" yaml:"podSecurityContext"`
	// List of container ports to expose on the container.
	Ports *[]*ControllerConfigSpecPorts `field:"optional" json:"ports" yaml:"ports"`
	// If specified, indicates the pod's priority.
	//
	// "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// Number of desired pods.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 1. Note: If more than 1 replica is set and leader election is not enabled then controllers could conflict. Environment variable "LEADER_ELECTION" can be used to enable leader election process.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// Compute Resources required by this container.
	//
	// Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Resources *ControllerConfigSpecResources `field:"optional" json:"resources" yaml:"resources"`
	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of Kubernetes v1.14.
	RuntimeClassName *string `field:"optional" json:"runtimeClassName" yaml:"runtimeClassName"`
	// SecurityContext holds container-level security attributes and common container settings.
	//
	// Optional: Defaults to empty.  See type description for default values of each field.
	SecurityContext *ControllerConfigSpecSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	// ServiceAccountName is the name of the ServiceAccount to use to run this pod.
	//
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// If specified, the pod's tolerations.
	Tolerations *[]*ControllerConfigSpecTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
}

