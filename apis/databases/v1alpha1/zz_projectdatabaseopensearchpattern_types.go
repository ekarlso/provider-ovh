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

type ProjectDatabaseOpensearchPatternInitParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Maximum number of index for this pattern
	MaxIndexCount *float64 `json:"maxIndexCount,omitempty" tf:"max_index_count,omitempty"`

	// Pattern format
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectDatabaseOpensearchPatternObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number of index for this pattern
	MaxIndexCount *float64 `json:"maxIndexCount,omitempty" tf:"max_index_count,omitempty"`

	// Pattern format
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectDatabaseOpensearchPatternParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Maximum number of index for this pattern
	// +kubebuilder:validation:Optional
	MaxIndexCount *float64 `json:"maxIndexCount,omitempty" tf:"max_index_count,omitempty"`

	// Pattern format
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectDatabaseOpensearchPatternSpec defines the desired state of ProjectDatabaseOpensearchPattern
type ProjectDatabaseOpensearchPatternSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseOpensearchPatternParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProjectDatabaseOpensearchPatternInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseOpensearchPatternStatus defines the observed state of ProjectDatabaseOpensearchPattern.
type ProjectDatabaseOpensearchPatternStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseOpensearchPatternObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseOpensearchPattern is the Schema for the ProjectDatabaseOpensearchPatterns API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type ProjectDatabaseOpensearchPattern struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pattern) || (has(self.initProvider) && has(self.initProvider.pattern))",message="spec.forProvider.pattern is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectDatabaseOpensearchPatternSpec   `json:"spec"`
	Status ProjectDatabaseOpensearchPatternStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseOpensearchPatternList contains a list of ProjectDatabaseOpensearchPatterns
type ProjectDatabaseOpensearchPatternList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseOpensearchPattern `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseOpensearchPattern_Kind             = "ProjectDatabaseOpensearchPattern"
	ProjectDatabaseOpensearchPattern_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseOpensearchPattern_Kind}.String()
	ProjectDatabaseOpensearchPattern_KindAPIVersion   = ProjectDatabaseOpensearchPattern_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseOpensearchPattern_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseOpensearchPattern_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseOpensearchPattern{}, &ProjectDatabaseOpensearchPatternList{})
}
