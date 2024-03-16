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

type CephACLInitParameters struct {
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type CephACLObservation struct {
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type CephACLParameters struct {

	// +kubebuilder:validation:Optional
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// CephACLSpec defines the desired state of CephACL
type CephACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CephACLParameters `json:"forProvider"`
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
	InitProvider CephACLInitParameters `json:"initProvider,omitempty"`
}

// CephACLStatus defines the observed state of CephACL.
type CephACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CephACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CephACL is the Schema for the CephACLs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type CephACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.netmask) || (has(self.initProvider) && has(self.initProvider.netmask))",message="spec.forProvider.netmask is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.network) || (has(self.initProvider) && has(self.initProvider.network))",message="spec.forProvider.network is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   CephACLSpec   `json:"spec"`
	Status CephACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CephACLList contains a list of CephACLs
type CephACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CephACL `json:"items"`
}

// Repository type metadata.
var (
	CephACL_Kind             = "CephACL"
	CephACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CephACL_Kind}.String()
	CephACL_KindAPIVersion   = CephACL_Kind + "." + CRDGroupVersion.String()
	CephACL_GroupVersionKind = CRDGroupVersion.WithKind(CephACL_Kind)
)

func init() {
	SchemeBuilder.Register(&CephACL{}, &CephACLList{})
}
