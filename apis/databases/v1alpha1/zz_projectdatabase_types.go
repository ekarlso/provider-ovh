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

type EndpointsInitParameters struct {
}

type EndpointsObservation struct {
	Component *string `json:"component,omitempty" tf:"component,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`

	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type EndpointsParameters struct {
}

type IPRestrictionsInitParameters struct {

	// Description of the IP restriction
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Authorized IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type IPRestrictionsObservation struct {

	// Description of the IP restriction
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Authorized IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Current status of the IP restriction
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type IPRestrictionsParameters struct {

	// Description of the IP restriction
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Authorized IP
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type NodesInitParameters struct {

	// Private network ID in which the node is. It's the regional openstackId of the private network.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Region of the node
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Private subnet ID in which the node is
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NodesObservation struct {

	// Private network ID in which the node is. It's the regional openstackId of the private network.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Region of the node
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Private subnet ID in which the node is
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NodesParameters struct {

	// Private network ID in which the node is. It's the regional openstackId of the private network.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Region of the node
	// +kubebuilder:validation:Optional
	Region *string `json:"region" tf:"region,omitempty"`

	// Private subnet ID in which the node is
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type ProjectDatabaseInitParameters struct {

	// Advanced configuration key / value
	// +mapType=granular
	AdvancedConfiguration map[string]*string `json:"advancedConfiguration,omitempty" tf:"advanced_configuration,omitempty"`

	// List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
	BackupRegions []*string `json:"backupRegions,omitempty" tf:"backup_regions,omitempty"`

	// Time on which backups start every day
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	// Description of the cluster
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disk size attributes of the cluster
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Name of the engine of the service
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The node flavor used for this cluster
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// IP Blocks authorized to access to the cluster
	IPRestrictions []IPRestrictionsInitParameters `json:"ipRestrictions,omitempty" tf:"ip_restrictions,omitempty"`

	// Defines whether the REST API is enabled on a Kafka cluster
	KafkaRestAPI *bool `json:"kafkaRestApi,omitempty" tf:"kafka_rest_api,omitempty"`

	// Defines whether the schema registry is enabled on a Kafka cluster
	KafkaSchemaRegistry *bool `json:"kafkaSchemaRegistry,omitempty" tf:"kafka_schema_registry,omitempty"`

	// List of nodes composing the service
	Nodes []NodesInitParameters `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Defines whether the ACLs are enabled on an Opensearch cluster
	OpensearchAclsEnabled *bool `json:"opensearchAclsEnabled,omitempty" tf:"opensearch_acls_enabled,omitempty"`

	// Plan of the cluster
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Version of the engine deployed on the cluster
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ProjectDatabaseObservation struct {

	// Advanced configuration key / value
	// +mapType=granular
	AdvancedConfiguration map[string]*string `json:"advancedConfiguration,omitempty" tf:"advanced_configuration,omitempty"`

	// List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
	BackupRegions []*string `json:"backupRegions,omitempty" tf:"backup_regions,omitempty"`

	// Time on which backups start every day
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	// Date of the creation of the cluster
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the cluster
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disk size attributes of the cluster
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Disk type attributes of the cluster
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// List of all endpoints of the service
	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// Name of the engine of the service
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The node flavor used for this cluster
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP Blocks authorized to access to the cluster
	IPRestrictions []IPRestrictionsObservation `json:"ipRestrictions,omitempty" tf:"ip_restrictions,omitempty"`

	// Defines whether the REST API is enabled on a Kafka cluster
	KafkaRestAPI *bool `json:"kafkaRestApi,omitempty" tf:"kafka_rest_api,omitempty"`

	// Defines whether the schema registry is enabled on a Kafka cluster
	KafkaSchemaRegistry *bool `json:"kafkaSchemaRegistry,omitempty" tf:"kafka_schema_registry,omitempty"`

	// Time on which maintenances can start every day
	MaintenanceTime *string `json:"maintenanceTime,omitempty" tf:"maintenance_time,omitempty"`

	// Type of network of the cluster
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// List of nodes composing the service
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Defines whether the ACLs are enabled on an Opensearch cluster
	OpensearchAclsEnabled *bool `json:"opensearchAclsEnabled,omitempty" tf:"opensearch_acls_enabled,omitempty"`

	// Plan of the cluster
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Current status of the cluster
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Version of the engine deployed on the cluster
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ProjectDatabaseParameters struct {

	// Advanced configuration key / value
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdvancedConfiguration map[string]*string `json:"advancedConfiguration,omitempty" tf:"advanced_configuration,omitempty"`

	// List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
	// +kubebuilder:validation:Optional
	BackupRegions []*string `json:"backupRegions,omitempty" tf:"backup_regions,omitempty"`

	// Time on which backups start every day
	// +kubebuilder:validation:Optional
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	// Description of the cluster
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disk size attributes of the cluster
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Name of the engine of the service
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The node flavor used for this cluster
	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// IP Blocks authorized to access to the cluster
	// +kubebuilder:validation:Optional
	IPRestrictions []IPRestrictionsParameters `json:"ipRestrictions,omitempty" tf:"ip_restrictions,omitempty"`

	// Defines whether the REST API is enabled on a Kafka cluster
	// +kubebuilder:validation:Optional
	KafkaRestAPI *bool `json:"kafkaRestApi,omitempty" tf:"kafka_rest_api,omitempty"`

	// Defines whether the schema registry is enabled on a Kafka cluster
	// +kubebuilder:validation:Optional
	KafkaSchemaRegistry *bool `json:"kafkaSchemaRegistry,omitempty" tf:"kafka_schema_registry,omitempty"`

	// List of nodes composing the service
	// +kubebuilder:validation:Optional
	Nodes []NodesParameters `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Defines whether the ACLs are enabled on an Opensearch cluster
	// +kubebuilder:validation:Optional
	OpensearchAclsEnabled *bool `json:"opensearchAclsEnabled,omitempty" tf:"opensearch_acls_enabled,omitempty"`

	// Plan of the cluster
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Version of the engine deployed on the cluster
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// ProjectDatabaseSpec defines the desired state of ProjectDatabase
type ProjectDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseStatus defines the observed state of ProjectDatabase.
type ProjectDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabase is the Schema for the ProjectDatabases API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type ProjectDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavor) || (has(self.initProvider) && has(self.initProvider.flavor))",message="spec.forProvider.flavor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodes) || (has(self.initProvider) && has(self.initProvider.nodes))",message="spec.forProvider.nodes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plan) || (has(self.initProvider) && has(self.initProvider.plan))",message="spec.forProvider.plan is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   ProjectDatabaseSpec   `json:"spec"`
	Status ProjectDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseList contains a list of ProjectDatabases
type ProjectDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabase `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabase_Kind             = "ProjectDatabase"
	ProjectDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabase_Kind}.String()
	ProjectDatabase_KindAPIVersion   = ProjectDatabase_Kind + "." + CRDGroupVersion.String()
	ProjectDatabase_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabase{}, &ProjectDatabaseList{})
}
