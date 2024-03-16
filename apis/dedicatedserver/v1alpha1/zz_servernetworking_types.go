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

type InterfacesInitParameters struct {

	// Interface Mac address
	// +listType=set
	Macs []*string `json:"macs,omitempty" tf:"macs,omitempty"`

	// Interface type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InterfacesObservation struct {

	// Interface Mac address
	// +listType=set
	Macs []*string `json:"macs,omitempty" tf:"macs,omitempty"`

	// Interface type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InterfacesParameters struct {

	// Interface Mac address
	// +kubebuilder:validation:Optional
	// +listType=set
	Macs []*string `json:"macs" tf:"macs,omitempty"`

	// Interface type
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ServerNetworkingInitParameters struct {

	// Interface or interfaces aggregation.
	Interfaces []InterfacesInitParameters `json:"interfaces,omitempty" tf:"interfaces,omitempty"`

	// The internal name of your dedicated server.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ServerNetworkingObservation struct {

	// Operation description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Interface or interfaces aggregation.
	Interfaces []InterfacesObservation `json:"interfaces,omitempty" tf:"interfaces,omitempty"`

	// The internal name of your dedicated server.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Operation status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServerNetworkingParameters struct {

	// Interface or interfaces aggregation.
	// +kubebuilder:validation:Optional
	Interfaces []InterfacesParameters `json:"interfaces,omitempty" tf:"interfaces,omitempty"`

	// The internal name of your dedicated server.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ServerNetworkingSpec defines the desired state of ServerNetworking
type ServerNetworkingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerNetworkingParameters `json:"forProvider"`
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
	InitProvider ServerNetworkingInitParameters `json:"initProvider,omitempty"`
}

// ServerNetworkingStatus defines the observed state of ServerNetworking.
type ServerNetworkingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerNetworkingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerNetworking is the Schema for the ServerNetworkings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ServerNetworking struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interfaces) || (has(self.initProvider) && has(self.initProvider.interfaces))",message="spec.forProvider.interfaces is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ServerNetworkingSpec   `json:"spec"`
	Status ServerNetworkingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerNetworkingList contains a list of ServerNetworkings
type ServerNetworkingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerNetworking `json:"items"`
}

// Repository type metadata.
var (
	ServerNetworking_Kind             = "ServerNetworking"
	ServerNetworking_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerNetworking_Kind}.String()
	ServerNetworking_KindAPIVersion   = ServerNetworking_Kind + "." + CRDGroupVersion.String()
	ServerNetworking_GroupVersionKind = CRDGroupVersion.WithKind(ServerNetworking_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerNetworking{}, &ServerNetworkingList{})
}
