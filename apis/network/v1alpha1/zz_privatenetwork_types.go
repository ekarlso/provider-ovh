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

type PrivateNetworkInitParameters struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type PrivateNetworkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	RegionsAttributes []RegionsAttributesObservation `json:"regionsAttributes,omitempty" tf:"regions_attributes,omitempty"`

	RegionsStatus []RegionsStatusObservation `json:"regionsStatus,omitempty" tf:"regions_status,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type PrivateNetworkParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type RegionsAttributesInitParameters struct {
}

type RegionsAttributesObservation struct {
	Openstackid *string `json:"openstackid,omitempty" tf:"openstackid,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegionsAttributesParameters struct {
}

type RegionsStatusInitParameters struct {
}

type RegionsStatusObservation struct {
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegionsStatusParameters struct {
}

// PrivateNetworkSpec defines the desired state of PrivateNetwork
type PrivateNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateNetworkParameters `json:"forProvider"`
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
	InitProvider PrivateNetworkInitParameters `json:"initProvider,omitempty"`
}

// PrivateNetworkStatus defines the observed state of PrivateNetwork.
type PrivateNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateNetwork is the Schema for the PrivateNetworks API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type PrivateNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   PrivateNetworkSpec   `json:"spec"`
	Status PrivateNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateNetworkList contains a list of PrivateNetworks
type PrivateNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateNetwork `json:"items"`
}

// Repository type metadata.
var (
	PrivateNetwork_Kind             = "PrivateNetwork"
	PrivateNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateNetwork_Kind}.String()
	PrivateNetwork_KindAPIVersion   = PrivateNetwork_Kind + "." + CRDGroupVersion.String()
	PrivateNetwork_GroupVersionKind = CRDGroupVersion.WithKind(PrivateNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateNetwork{}, &PrivateNetworkList{})
}
