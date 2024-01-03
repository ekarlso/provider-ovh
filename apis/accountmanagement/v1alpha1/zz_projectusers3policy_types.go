// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ProjectUserS3PolicyInitParameters struct {

	// The policy document. This is a JSON formatted string.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The user ID
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type ProjectUserS3PolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy document. This is a JSON formatted string.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The user ID
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type ProjectUserS3PolicyParameters struct {

	// The policy document. This is a JSON formatted string.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The user ID
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// ProjectUserS3PolicySpec defines the desired state of ProjectUserS3Policy
type ProjectUserS3PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectUserS3PolicyParameters `json:"forProvider"`
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
	InitProvider ProjectUserS3PolicyInitParameters `json:"initProvider,omitempty"`
}

// ProjectUserS3PolicyStatus defines the observed state of ProjectUserS3Policy.
type ProjectUserS3PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectUserS3PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectUserS3Policy is the Schema for the ProjectUserS3Policys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectUserS3Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userId) || (has(self.initProvider) && has(self.initProvider.userId))",message="spec.forProvider.userId is a required parameter"
	Spec   ProjectUserS3PolicySpec   `json:"spec"`
	Status ProjectUserS3PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectUserS3PolicyList contains a list of ProjectUserS3Policys
type ProjectUserS3PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectUserS3Policy `json:"items"`
}

// Repository type metadata.
var (
	ProjectUserS3Policy_Kind             = "ProjectUserS3Policy"
	ProjectUserS3Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectUserS3Policy_Kind}.String()
	ProjectUserS3Policy_KindAPIVersion   = ProjectUserS3Policy_Kind + "." + CRDGroupVersion.String()
	ProjectUserS3Policy_GroupVersionKind = CRDGroupVersion.WithKind(ProjectUserS3Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectUserS3Policy{}, &ProjectUserS3PolicyList{})
}
