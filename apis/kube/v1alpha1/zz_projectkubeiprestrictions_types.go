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

type ProjectKubeIprestrictionsInitParameters struct {

	// List of IP restrictions for the cluster
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// Service name
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectKubeIprestrictionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of IP restrictions for the cluster
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// Kube ID
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// Service name
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectKubeIprestrictionsParameters struct {

	// List of IP restrictions for the cluster
	// +kubebuilder:validation:Optional
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// Kube ID
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/kube/v1alpha1.ProjectKube
	// +kubebuilder:validation:Optional
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// Reference to a ProjectKube in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDRef *v1.Reference `json:"kubeIdRef,omitempty" tf:"-"`

	// Selector for a ProjectKube in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDSelector *v1.Selector `json:"kubeIdSelector,omitempty" tf:"-"`

	// Service name
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectKubeIprestrictionsSpec defines the desired state of ProjectKubeIprestrictions
type ProjectKubeIprestrictionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectKubeIprestrictionsParameters `json:"forProvider"`
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
	InitProvider ProjectKubeIprestrictionsInitParameters `json:"initProvider,omitempty"`
}

// ProjectKubeIprestrictionsStatus defines the observed state of ProjectKubeIprestrictions.
type ProjectKubeIprestrictionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectKubeIprestrictionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectKubeIprestrictions is the Schema for the ProjectKubeIprestrictionss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectKubeIprestrictions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ips) || (has(self.initProvider) && has(self.initProvider.ips))",message="spec.forProvider.ips is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectKubeIprestrictionsSpec   `json:"spec"`
	Status ProjectKubeIprestrictionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectKubeIprestrictionsList contains a list of ProjectKubeIprestrictionss
type ProjectKubeIprestrictionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectKubeIprestrictions `json:"items"`
}

// Repository type metadata.
var (
	ProjectKubeIprestrictions_Kind             = "ProjectKubeIprestrictions"
	ProjectKubeIprestrictions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectKubeIprestrictions_Kind}.String()
	ProjectKubeIprestrictions_KindAPIVersion   = ProjectKubeIprestrictions_Kind + "." + CRDGroupVersion.String()
	ProjectKubeIprestrictions_GroupVersionKind = CRDGroupVersion.WithKind(ProjectKubeIprestrictions_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectKubeIprestrictions{}, &ProjectKubeIprestrictionsList{})
}
