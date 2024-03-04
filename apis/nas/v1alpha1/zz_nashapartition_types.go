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

type NashaPartitionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type NashaPartitionObservation struct {
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	UsedBySnapshots *float64 `json:"usedBySnapshots,omitempty" tf:"used_by_snapshots,omitempty"`
}

type NashaPartitionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

// NashaPartitionSpec defines the desired state of NashaPartition
type NashaPartitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NashaPartitionParameters `json:"forProvider"`
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
	InitProvider NashaPartitionInitParameters `json:"initProvider,omitempty"`
}

// NashaPartitionStatus defines the observed state of NashaPartition.
type NashaPartitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NashaPartitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NashaPartition is the Schema for the NashaPartitions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type NashaPartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	Spec   NashaPartitionSpec   `json:"spec"`
	Status NashaPartitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NashaPartitionList contains a list of NashaPartitions
type NashaPartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NashaPartition `json:"items"`
}

// Repository type metadata.
var (
	NashaPartition_Kind             = "NashaPartition"
	NashaPartition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NashaPartition_Kind}.String()
	NashaPartition_KindAPIVersion   = NashaPartition_Kind + "." + CRDGroupVersion.String()
	NashaPartition_GroupVersionKind = CRDGroupVersion.WithKind(NashaPartition_Kind)
)

func init() {
	SchemeBuilder.Register(&NashaPartition{}, &NashaPartitionList{})
}
