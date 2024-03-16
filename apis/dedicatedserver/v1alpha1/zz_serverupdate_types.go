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

type ServerUpdateInitParameters struct {

	// The boot id of your dedicated server.
	BootID *float64 `json:"bootId,omitempty" tf:"boot_id,omitempty"`

	// The boot script of your dedicated server.
	BootScript *string `json:"bootScript,omitempty" tf:"boot_script,omitempty"`

	// Icmp monitoring state
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// The internal name of your dedicated server.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// error, hacked, hackedBlocked, ok
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ServerUpdateObservation struct {

	// The boot id of your dedicated server.
	BootID *float64 `json:"bootId,omitempty" tf:"boot_id,omitempty"`

	// The boot script of your dedicated server.
	BootScript *string `json:"bootScript,omitempty" tf:"boot_script,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Icmp monitoring state
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// The internal name of your dedicated server.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// error, hacked, hackedBlocked, ok
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ServerUpdateParameters struct {

	// The boot id of your dedicated server.
	// +kubebuilder:validation:Optional
	BootID *float64 `json:"bootId,omitempty" tf:"boot_id,omitempty"`

	// The boot script of your dedicated server.
	// +kubebuilder:validation:Optional
	BootScript *string `json:"bootScript,omitempty" tf:"boot_script,omitempty"`

	// Icmp monitoring state
	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// The internal name of your dedicated server.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// error, hacked, hackedBlocked, ok
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

// ServerUpdateSpec defines the desired state of ServerUpdate
type ServerUpdateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerUpdateParameters `json:"forProvider"`
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
	InitProvider ServerUpdateInitParameters `json:"initProvider,omitempty"`
}

// ServerUpdateStatus defines the observed state of ServerUpdate.
type ServerUpdateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerUpdateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerUpdate is the Schema for the ServerUpdates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ServerUpdate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ServerUpdateSpec   `json:"spec"`
	Status ServerUpdateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerUpdateList contains a list of ServerUpdates
type ServerUpdateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerUpdate `json:"items"`
}

// Repository type metadata.
var (
	ServerUpdate_Kind             = "ServerUpdate"
	ServerUpdate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerUpdate_Kind}.String()
	ServerUpdate_KindAPIVersion   = ServerUpdate_Kind + "." + CRDGroupVersion.String()
	ServerUpdate_GroupVersionKind = CRDGroupVersion.WithKind(ServerUpdate_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerUpdate{}, &ServerUpdateList{})
}
