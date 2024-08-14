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

type PrivatedatabaseUserGrantInitParameters struct {

	// Database name where add grant
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Database name where add grant
	Grant *string `json:"grant,omitempty" tf:"grant,omitempty"`

	// The internal name of your private database
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// User name used to connect on your databases
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type PrivatedatabaseUserGrantObservation struct {

	// Database name where add grant
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Database name where add grant
	Grant *string `json:"grant,omitempty" tf:"grant,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The internal name of your private database
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// User name used to connect on your databases
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type PrivatedatabaseUserGrantParameters struct {

	// Database name where add grant
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Database name where add grant
	// +kubebuilder:validation:Optional
	Grant *string `json:"grant,omitempty" tf:"grant,omitempty"`

	// The internal name of your private database
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// User name used to connect on your databases
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// PrivatedatabaseUserGrantSpec defines the desired state of PrivatedatabaseUserGrant
type PrivatedatabaseUserGrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivatedatabaseUserGrantParameters `json:"forProvider"`
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
	InitProvider PrivatedatabaseUserGrantInitParameters `json:"initProvider,omitempty"`
}

// PrivatedatabaseUserGrantStatus defines the observed state of PrivatedatabaseUserGrant.
type PrivatedatabaseUserGrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivatedatabaseUserGrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivatedatabaseUserGrant is the Schema for the PrivatedatabaseUserGrants API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type PrivatedatabaseUserGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseName) || (has(self.initProvider) && has(self.initProvider.databaseName))",message="spec.forProvider.databaseName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.grant) || (has(self.initProvider) && has(self.initProvider.grant))",message="spec.forProvider.grant is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || (has(self.initProvider) && has(self.initProvider.userName))",message="spec.forProvider.userName is a required parameter"
	Spec   PrivatedatabaseUserGrantSpec   `json:"spec"`
	Status PrivatedatabaseUserGrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivatedatabaseUserGrantList contains a list of PrivatedatabaseUserGrants
type PrivatedatabaseUserGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivatedatabaseUserGrant `json:"items"`
}

// Repository type metadata.
var (
	PrivatedatabaseUserGrant_Kind             = "PrivatedatabaseUserGrant"
	PrivatedatabaseUserGrant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivatedatabaseUserGrant_Kind}.String()
	PrivatedatabaseUserGrant_KindAPIVersion   = PrivatedatabaseUserGrant_Kind + "." + CRDGroupVersion.String()
	PrivatedatabaseUserGrant_GroupVersionKind = CRDGroupVersion.WithKind(PrivatedatabaseUserGrant_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivatedatabaseUserGrant{}, &PrivatedatabaseUserGrantList{})
}
