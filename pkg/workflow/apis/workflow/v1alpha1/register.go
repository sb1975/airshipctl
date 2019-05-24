package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Workflow constants
const (
	Kind      string = "Workflow"
	Group     string = "argoproj.io"
	Singular  string = "workflow"
	Plural    string = "workflows"
	ShortName string = "wf"
	FullName  string = Plural + "." + Group
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion     = schema.GroupVersion{Group: Group, Version: "v1alpha1"}
	SchemaGroupVersionKind = schema.GroupVersionKind{Group: Group, Version: "v1alpha1", Kind: Kind}
)

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// addKnownTypes adds the set of types defined in this package to the supplied scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Workflow{},
		&WorkflowList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
