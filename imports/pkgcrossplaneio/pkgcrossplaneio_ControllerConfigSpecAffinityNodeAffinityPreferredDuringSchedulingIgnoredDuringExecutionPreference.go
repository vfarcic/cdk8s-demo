// pkgcrossplaneio
package pkgcrossplaneio


// A node selector term, associated with the corresponding weight.
type ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference struct {
	// A list of node selector requirements by node's labels.
	MatchExpressions *[]*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// A list of node selector requirements by node's fields.
	MatchFields *[]*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields `field:"optional" json:"matchFields" yaml:"matchFields"`
}

