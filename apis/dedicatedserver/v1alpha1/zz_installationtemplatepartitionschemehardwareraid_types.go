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

type InstallationTemplatePartitionSchemeHardwareRaidInitParameters struct {

	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks []*string `json:"disks,omitempty" tf:"disks,omitempty"`

	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Hardware RAID name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// name of this partitioning scheme
	SchemeName *string `json:"schemeName,omitempty" tf:"scheme_name,omitempty"`

	// Specifies the creation order of the hardware RAID
	Step *float64 `json:"step,omitempty" tf:"step,omitempty"`

	// Template name
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

type InstallationTemplatePartitionSchemeHardwareRaidObservation struct {

	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks []*string `json:"disks,omitempty" tf:"disks,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Hardware RAID name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// name of this partitioning scheme
	SchemeName *string `json:"schemeName,omitempty" tf:"scheme_name,omitempty"`

	// Specifies the creation order of the hardware RAID
	Step *float64 `json:"step,omitempty" tf:"step,omitempty"`

	// Template name
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

type InstallationTemplatePartitionSchemeHardwareRaidParameters struct {

	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	// +kubebuilder:validation:Optional
	Disks []*string `json:"disks,omitempty" tf:"disks,omitempty"`

	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Hardware RAID name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// name of this partitioning scheme
	// +kubebuilder:validation:Optional
	SchemeName *string `json:"schemeName,omitempty" tf:"scheme_name,omitempty"`

	// Specifies the creation order of the hardware RAID
	// +kubebuilder:validation:Optional
	Step *float64 `json:"step,omitempty" tf:"step,omitempty"`

	// Template name
	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

// InstallationTemplatePartitionSchemeHardwareRaidSpec defines the desired state of InstallationTemplatePartitionSchemeHardwareRaid
type InstallationTemplatePartitionSchemeHardwareRaidSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstallationTemplatePartitionSchemeHardwareRaidParameters `json:"forProvider"`
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
	InitProvider InstallationTemplatePartitionSchemeHardwareRaidInitParameters `json:"initProvider,omitempty"`
}

// InstallationTemplatePartitionSchemeHardwareRaidStatus defines the observed state of InstallationTemplatePartitionSchemeHardwareRaid.
type InstallationTemplatePartitionSchemeHardwareRaidStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstallationTemplatePartitionSchemeHardwareRaidObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstallationTemplatePartitionSchemeHardwareRaid is the Schema for the InstallationTemplatePartitionSchemeHardwareRaids API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type InstallationTemplatePartitionSchemeHardwareRaid struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.disks) || (has(self.initProvider) && has(self.initProvider.disks))",message="spec.forProvider.disks is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemeName) || (has(self.initProvider) && has(self.initProvider.schemeName))",message="spec.forProvider.schemeName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.step) || (has(self.initProvider) && has(self.initProvider.step))",message="spec.forProvider.step is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateName) || (has(self.initProvider) && has(self.initProvider.templateName))",message="spec.forProvider.templateName is a required parameter"
	Spec   InstallationTemplatePartitionSchemeHardwareRaidSpec   `json:"spec"`
	Status InstallationTemplatePartitionSchemeHardwareRaidStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstallationTemplatePartitionSchemeHardwareRaidList contains a list of InstallationTemplatePartitionSchemeHardwareRaids
type InstallationTemplatePartitionSchemeHardwareRaidList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstallationTemplatePartitionSchemeHardwareRaid `json:"items"`
}

// Repository type metadata.
var (
	InstallationTemplatePartitionSchemeHardwareRaid_Kind             = "InstallationTemplatePartitionSchemeHardwareRaid"
	InstallationTemplatePartitionSchemeHardwareRaid_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstallationTemplatePartitionSchemeHardwareRaid_Kind}.String()
	InstallationTemplatePartitionSchemeHardwareRaid_KindAPIVersion   = InstallationTemplatePartitionSchemeHardwareRaid_Kind + "." + CRDGroupVersion.String()
	InstallationTemplatePartitionSchemeHardwareRaid_GroupVersionKind = CRDGroupVersion.WithKind(InstallationTemplatePartitionSchemeHardwareRaid_Kind)
)

func init() {
	SchemeBuilder.Register(&InstallationTemplatePartitionSchemeHardwareRaid{}, &InstallationTemplatePartitionSchemeHardwareRaidList{})
}
