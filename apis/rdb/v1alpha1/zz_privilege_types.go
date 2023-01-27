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

type PrivilegeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivilegeParameters struct {

	// Database name
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Instance on which the database is created
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Privilege
	// +kubebuilder:validation:Required
	Permission *string `json:"permission" tf:"permission,omitempty"`

	// User name
	// +kubebuilder:validation:Required
	UserName *string `json:"userName" tf:"user_name,omitempty"`
}

// PrivilegeSpec defines the desired state of Privilege
type PrivilegeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivilegeParameters `json:"forProvider"`
}

// PrivilegeStatus defines the observed state of Privilege.
type PrivilegeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivilegeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Privilege is the Schema for the Privileges API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Privilege struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivilegeSpec   `json:"spec"`
	Status            PrivilegeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivilegeList contains a list of Privileges
type PrivilegeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Privilege `json:"items"`
}

// Repository type metadata.
var (
	Privilege_Kind             = "Privilege"
	Privilege_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Privilege_Kind}.String()
	Privilege_KindAPIVersion   = Privilege_Kind + "." + CRDGroupVersion.String()
	Privilege_GroupVersionKind = CRDGroupVersion.WithKind(Privilege_Kind)
)

func init() {
	SchemeBuilder.Register(&Privilege{}, &PrivilegeList{})
}
