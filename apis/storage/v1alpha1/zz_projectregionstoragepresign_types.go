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

type ProjectRegionStoragePresignInitParameters struct {

	// How long (in seconds) the URL will be valid.
	Expire *float64 `json:"expire,omitempty" tf:"expire,omitempty"`

	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The S3 storage container's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the object to download or upload.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Region name.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectRegionStoragePresignObservation struct {

	// How long (in seconds) the URL will be valid.
	Expire *float64 `json:"expire,omitempty" tf:"expire,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The S3 storage container's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the object to download or upload.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Region name.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Presigned URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ProjectRegionStoragePresignParameters struct {

	// How long (in seconds) the URL will be valid.
	// +kubebuilder:validation:Optional
	Expire *float64 `json:"expire,omitempty" tf:"expire,omitempty"`

	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The S3 storage container's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the object to download or upload.
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Region name.
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectRegionStoragePresignSpec defines the desired state of ProjectRegionStoragePresign
type ProjectRegionStoragePresignSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectRegionStoragePresignParameters `json:"forProvider"`
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
	InitProvider ProjectRegionStoragePresignInitParameters `json:"initProvider,omitempty"`
}

// ProjectRegionStoragePresignStatus defines the observed state of ProjectRegionStoragePresign.
type ProjectRegionStoragePresignStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectRegionStoragePresignObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectRegionStoragePresign is the Schema for the ProjectRegionStoragePresigns API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectRegionStoragePresign struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.expire) || (has(self.initProvider) && has(self.initProvider.expire))",message="spec.forProvider.expire is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.method) || (has(self.initProvider) && has(self.initProvider.method))",message="spec.forProvider.method is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.object) || (has(self.initProvider) && has(self.initProvider.object))",message="spec.forProvider.object is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionName) || (has(self.initProvider) && has(self.initProvider.regionName))",message="spec.forProvider.regionName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectRegionStoragePresignSpec   `json:"spec"`
	Status ProjectRegionStoragePresignStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectRegionStoragePresignList contains a list of ProjectRegionStoragePresigns
type ProjectRegionStoragePresignList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRegionStoragePresign `json:"items"`
}

// Repository type metadata.
var (
	ProjectRegionStoragePresign_Kind             = "ProjectRegionStoragePresign"
	ProjectRegionStoragePresign_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectRegionStoragePresign_Kind}.String()
	ProjectRegionStoragePresign_KindAPIVersion   = ProjectRegionStoragePresign_Kind + "." + CRDGroupVersion.String()
	ProjectRegionStoragePresign_GroupVersionKind = CRDGroupVersion.WithKind(ProjectRegionStoragePresign_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectRegionStoragePresign{}, &ProjectRegionStoragePresignList{})
}
