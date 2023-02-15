/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ContainerNamespaceObservation struct {

	// The ID of the namespace
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization ID the namespace is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The registry endpoint of the namespace.
	// The endpoint reachable by docker
	RegistryEndpoint *string `json:"registryEndpoint,omitempty" tf:"registry_endpoint,omitempty"`

	// The registry namespace ID of the namespace.
	// The ID of the registry namespace
	RegistryNamespaceID *string `json:"registryNamespaceId,omitempty" tf:"registry_namespace_id,omitempty"`
}

type ContainerNamespaceParameters struct {

	// The description of the namespace.
	// The description of the container namespace
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Defaults to false). Destroy linked container registry on deletion.
	// Destroy registry on deletion
	// +kubebuilder:validation:Optional
	DestroyRegistry *bool `json:"destroyRegistry,omitempty" tf:"destroy_registry,omitempty"`

	// The environment variables of the namespace.
	// The environment variables of the container namespace
	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The unique name of the container namespace.
	// The name of the container namespace
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the namespace is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the namespace should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The secret environment variables of the namespace.
	// The secret environment variables of the container namespace
	// +kubebuilder:validation:Optional
	SecretEnvironmentVariablesSecretRef *v1.SecretReference `json:"secretEnvironmentVariablesSecretRef,omitempty" tf:"-"`
}

// ContainerNamespaceSpec defines the desired state of ContainerNamespace
type ContainerNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerNamespaceParameters `json:"forProvider"`
}

// ContainerNamespaceStatus defines the observed state of ContainerNamespace.
type ContainerNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerNamespace is the Schema for the ContainerNamespaces API. Manages Scaleway Container Namespaces.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type ContainerNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerNamespaceSpec   `json:"spec"`
	Status            ContainerNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerNamespaceList contains a list of ContainerNamespaces
type ContainerNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerNamespace `json:"items"`
}

// Repository type metadata.
var (
	ContainerNamespace_Kind             = "ContainerNamespace"
	ContainerNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerNamespace_Kind}.String()
	ContainerNamespace_KindAPIVersion   = ContainerNamespace_Kind + "." + CRDGroupVersion.String()
	ContainerNamespace_GroupVersionKind = CRDGroupVersion.WithKind(ContainerNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerNamespace{}, &ContainerNamespaceList{})
}
