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

type ProjectDatabaseM3DbNamespaceInitParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the namespace
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resolution for an aggregated namespace
	Resolution *string `json:"resolution,omitempty" tf:"resolution,omitempty"`

	// Controls how long we wait before expiring stale data
	RetentionBlockDataExpirationDuration *string `json:"retentionBlockDataExpirationDuration,omitempty" tf:"retention_block_data_expiration_duration,omitempty"`

	// Controls how long to keep a block in memory before flushing to a fileset on disk
	RetentionBlockSizeDuration *string `json:"retentionBlockSizeDuration,omitempty" tf:"retention_block_size_duration,omitempty"`

	// Controls how far into the future writes to the namespace will be accepted
	RetentionBufferFutureDuration *string `json:"retentionBufferFutureDuration,omitempty" tf:"retention_buffer_future_duration,omitempty"`

	// Controls how far into the past writes to the namespace will be accepted
	RetentionBufferPastDuration *string `json:"retentionBufferPastDuration,omitempty" tf:"retention_buffer_past_duration,omitempty"`

	// Controls the duration of time that M3DB will retain data for the namespace
	RetentionPeriodDuration *string `json:"retentionPeriodDuration,omitempty" tf:"retention_period_duration,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Defines whether M3db will create snapshot files for this namespace
	SnapshotEnabled *bool `json:"snapshotEnabled,omitempty" tf:"snapshot_enabled,omitempty"`

	// Defines whether M3db will include writes to this namespace in the commit log
	WritesToCommitLogEnabled *bool `json:"writesToCommitLogEnabled,omitempty" tf:"writes_to_commit_log_enabled,omitempty"`
}

type ProjectDatabaseM3DbNamespaceObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the namespace
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resolution for an aggregated namespace
	Resolution *string `json:"resolution,omitempty" tf:"resolution,omitempty"`

	// Controls how long we wait before expiring stale data
	RetentionBlockDataExpirationDuration *string `json:"retentionBlockDataExpirationDuration,omitempty" tf:"retention_block_data_expiration_duration,omitempty"`

	// Controls how long to keep a block in memory before flushing to a fileset on disk
	RetentionBlockSizeDuration *string `json:"retentionBlockSizeDuration,omitempty" tf:"retention_block_size_duration,omitempty"`

	// Controls how far into the future writes to the namespace will be accepted
	RetentionBufferFutureDuration *string `json:"retentionBufferFutureDuration,omitempty" tf:"retention_buffer_future_duration,omitempty"`

	// Controls how far into the past writes to the namespace will be accepted
	RetentionBufferPastDuration *string `json:"retentionBufferPastDuration,omitempty" tf:"retention_buffer_past_duration,omitempty"`

	// Controls the duration of time that M3DB will retain data for the namespace
	RetentionPeriodDuration *string `json:"retentionPeriodDuration,omitempty" tf:"retention_period_duration,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Defines whether M3db will create snapshot files for this namespace
	SnapshotEnabled *bool `json:"snapshotEnabled,omitempty" tf:"snapshot_enabled,omitempty"`

	// Type of namespace
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines whether M3db will include writes to this namespace in the commit log
	WritesToCommitLogEnabled *bool `json:"writesToCommitLogEnabled,omitempty" tf:"writes_to_commit_log_enabled,omitempty"`
}

type ProjectDatabaseM3DbNamespaceParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the namespace
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resolution for an aggregated namespace
	// +kubebuilder:validation:Optional
	Resolution *string `json:"resolution,omitempty" tf:"resolution,omitempty"`

	// Controls how long we wait before expiring stale data
	// +kubebuilder:validation:Optional
	RetentionBlockDataExpirationDuration *string `json:"retentionBlockDataExpirationDuration,omitempty" tf:"retention_block_data_expiration_duration,omitempty"`

	// Controls how long to keep a block in memory before flushing to a fileset on disk
	// +kubebuilder:validation:Optional
	RetentionBlockSizeDuration *string `json:"retentionBlockSizeDuration,omitempty" tf:"retention_block_size_duration,omitempty"`

	// Controls how far into the future writes to the namespace will be accepted
	// +kubebuilder:validation:Optional
	RetentionBufferFutureDuration *string `json:"retentionBufferFutureDuration,omitempty" tf:"retention_buffer_future_duration,omitempty"`

	// Controls how far into the past writes to the namespace will be accepted
	// +kubebuilder:validation:Optional
	RetentionBufferPastDuration *string `json:"retentionBufferPastDuration,omitempty" tf:"retention_buffer_past_duration,omitempty"`

	// Controls the duration of time that M3DB will retain data for the namespace
	// +kubebuilder:validation:Optional
	RetentionPeriodDuration *string `json:"retentionPeriodDuration,omitempty" tf:"retention_period_duration,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Defines whether M3db will create snapshot files for this namespace
	// +kubebuilder:validation:Optional
	SnapshotEnabled *bool `json:"snapshotEnabled,omitempty" tf:"snapshot_enabled,omitempty"`

	// Defines whether M3db will include writes to this namespace in the commit log
	// +kubebuilder:validation:Optional
	WritesToCommitLogEnabled *bool `json:"writesToCommitLogEnabled,omitempty" tf:"writes_to_commit_log_enabled,omitempty"`
}

// ProjectDatabaseM3DbNamespaceSpec defines the desired state of ProjectDatabaseM3DbNamespace
type ProjectDatabaseM3DbNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseM3DbNamespaceParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseM3DbNamespaceInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseM3DbNamespaceStatus defines the observed state of ProjectDatabaseM3DbNamespace.
type ProjectDatabaseM3DbNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseM3DbNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseM3DbNamespace is the Schema for the ProjectDatabaseM3DbNamespaces API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type ProjectDatabaseM3DbNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resolution) || (has(self.initProvider) && has(self.initProvider.resolution))",message="spec.forProvider.resolution is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionPeriodDuration) || (has(self.initProvider) && has(self.initProvider.retentionPeriodDuration))",message="spec.forProvider.retentionPeriodDuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectDatabaseM3DbNamespaceSpec   `json:"spec"`
	Status ProjectDatabaseM3DbNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseM3DbNamespaceList contains a list of ProjectDatabaseM3DbNamespaces
type ProjectDatabaseM3DbNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseM3DbNamespace `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseM3DbNamespace_Kind             = "ProjectDatabaseM3DbNamespace"
	ProjectDatabaseM3DbNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseM3DbNamespace_Kind}.String()
	ProjectDatabaseM3DbNamespace_KindAPIVersion   = ProjectDatabaseM3DbNamespace_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseM3DbNamespace_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseM3DbNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseM3DbNamespace{}, &ProjectDatabaseM3DbNamespaceList{})
}
