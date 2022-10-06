// pkgcrossplaneio
package pkgcrossplaneio


// A null or empty node selector term matches no objects.
//
// The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms struct {
	// A list of node selector requirements by node's labels.
	MatchExpressions *[]*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// A list of node selector requirements by node's fields.
	MatchFields *[]*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields `field:"optional" json:"matchFields" yaml:"matchFields"`
}

