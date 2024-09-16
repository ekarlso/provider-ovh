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

type ProjectDatabaseIPRestrictionInitParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Description of the IP restriction
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the engine of the service
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Authorized IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectDatabaseIPRestrictionObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Description of the IP restriction
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the engine of the service
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Authorized IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Current status of the IP restriction
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ProjectDatabaseIPRestrictionParameters struct {

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

	// Description of the IP restriction
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the engine of the service
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Authorized IP
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectDatabaseIPRestrictionSpec defines the desired state of ProjectDatabaseIPRestriction
type ProjectDatabaseIPRestrictionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseIPRestrictionParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseIPRestrictionInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseIPRestrictionStatus defines the observed state of ProjectDatabaseIPRestriction.
type ProjectDatabaseIPRestrictionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseIPRestrictionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseIPRestriction is the Schema for the ProjectDatabaseIPRestrictions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type ProjectDatabaseIPRestriction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ip) || (has(self.initProvider) && has(self.initProvider.ip))",message="spec.forProvider.ip is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectDatabaseIPRestrictionSpec   `json:"spec"`
	Status ProjectDatabaseIPRestrictionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseIPRestrictionList contains a list of ProjectDatabaseIPRestrictions
type ProjectDatabaseIPRestrictionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseIPRestriction `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseIPRestriction_Kind             = "ProjectDatabaseIPRestriction"
	ProjectDatabaseIPRestriction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseIPRestriction_Kind}.String()
	ProjectDatabaseIPRestriction_KindAPIVersion   = ProjectDatabaseIPRestriction_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseIPRestriction_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseIPRestriction_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseIPRestriction{}, &ProjectDatabaseIPRestrictionList{})
}
