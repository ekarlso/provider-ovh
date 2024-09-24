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

type TCPFarmServerInitParameters struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	FarmID *float64 `json:"farmId,omitempty" tf:"farm_id,omitempty"`

	OnMarkedDown *string `json:"onMarkedDown,omitempty" tf:"on_marked_down,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Probe *bool `json:"probe,omitempty" tf:"probe,omitempty"`

	ProxyProtocolVersion *string `json:"proxyProtocolVersion,omitempty" tf:"proxy_protocol_version,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TCPFarmServerObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	FarmID *float64 `json:"farmId,omitempty" tf:"farm_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OnMarkedDown *string `json:"onMarkedDown,omitempty" tf:"on_marked_down,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Probe *bool `json:"probe,omitempty" tf:"probe,omitempty"`

	ProxyProtocolVersion *string `json:"proxyProtocolVersion,omitempty" tf:"proxy_protocol_version,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TCPFarmServerParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FarmID *float64 `json:"farmId,omitempty" tf:"farm_id,omitempty"`

	// +kubebuilder:validation:Optional
	OnMarkedDown *string `json:"onMarkedDown,omitempty" tf:"on_marked_down,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Probe *bool `json:"probe,omitempty" tf:"probe,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyProtocolVersion *string `json:"proxyProtocolVersion,omitempty" tf:"proxy_protocol_version,omitempty"`

	// +kubebuilder:validation:Optional
	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// TCPFarmServerSpec defines the desired state of TCPFarmServer
type TCPFarmServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TCPFarmServerParameters `json:"forProvider"`
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
	InitProvider TCPFarmServerInitParameters `json:"initProvider,omitempty"`
}

// TCPFarmServerStatus defines the observed state of TCPFarmServer.
type TCPFarmServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TCPFarmServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TCPFarmServer is the Schema for the TCPFarmServers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type TCPFarmServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.address) || (has(self.initProvider) && has(self.initProvider.address))",message="spec.forProvider.address is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.farmId) || (has(self.initProvider) && has(self.initProvider.farmId))",message="spec.forProvider.farmId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || (has(self.initProvider) && has(self.initProvider.status))",message="spec.forProvider.status is a required parameter"
	Spec   TCPFarmServerSpec   `json:"spec"`
	Status TCPFarmServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TCPFarmServerList contains a list of TCPFarmServers
type TCPFarmServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPFarmServer `json:"items"`
}

// Repository type metadata.
var (
	TCPFarmServer_Kind             = "TCPFarmServer"
	TCPFarmServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TCPFarmServer_Kind}.String()
	TCPFarmServer_KindAPIVersion   = TCPFarmServer_Kind + "." + CRDGroupVersion.String()
	TCPFarmServer_GroupVersionKind = CRDGroupVersion.WithKind(TCPFarmServer_Kind)
)

func init() {
	SchemeBuilder.Register(&TCPFarmServer{}, &TCPFarmServerList{})
}
