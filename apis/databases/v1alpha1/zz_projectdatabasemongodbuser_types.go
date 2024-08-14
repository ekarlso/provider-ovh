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

type ProjectDatabaseMongodbUserInitParameters struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Name of the user
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	// Roles the user belongs to with the authentication database
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectDatabaseMongodbUserObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Date of the creation of the user
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the user
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	// Roles the user belongs to with the authentication database
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Current status of the user
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ProjectDatabaseMongodbUserParameters struct {

	// Id of the database cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Name of the user
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	// +kubebuilder:validation:Optional
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	// Roles the user belongs to with the authentication database
	// +kubebuilder:validation:Optional
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectDatabaseMongodbUserSpec defines the desired state of ProjectDatabaseMongodbUser
type ProjectDatabaseMongodbUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseMongodbUserParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseMongodbUserInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseMongodbUserStatus defines the observed state of ProjectDatabaseMongodbUser.
type ProjectDatabaseMongodbUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseMongodbUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseMongodbUser is the Schema for the ProjectDatabaseMongodbUsers API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type ProjectDatabaseMongodbUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterId) || (has(self.initProvider) && has(self.initProvider.clusterId))",message="spec.forProvider.clusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectDatabaseMongodbUserSpec   `json:"spec"`
	Status ProjectDatabaseMongodbUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseMongodbUserList contains a list of ProjectDatabaseMongodbUsers
type ProjectDatabaseMongodbUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseMongodbUser `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseMongodbUser_Kind             = "ProjectDatabaseMongodbUser"
	ProjectDatabaseMongodbUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseMongodbUser_Kind}.String()
	ProjectDatabaseMongodbUser_KindAPIVersion   = ProjectDatabaseMongodbUser_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseMongodbUser_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseMongodbUser_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseMongodbUser{}, &ProjectDatabaseMongodbUserList{})
}
