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

type OIDCConfigurationInitParameters struct {
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	IssuerURL *string `json:"issuerUrl,omitempty" tf:"issuer_url,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/kube/v1alpha1.Cluster
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// Reference to a Cluster in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDRef *v1.Reference `json:"kubeIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDSelector *v1.Selector `json:"kubeIdSelector,omitempty" tf:"-"`

	OidcCAContent *string `json:"oidcCaContent,omitempty" tf:"oidc_ca_content,omitempty"`

	OidcGroupsClaim []*string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	OidcGroupsPrefix *string `json:"oidcGroupsPrefix,omitempty" tf:"oidc_groups_prefix,omitempty"`

	OidcRequiredClaim []*string `json:"oidcRequiredClaim,omitempty" tf:"oidc_required_claim,omitempty"`

	OidcSigningAlgs []*string `json:"oidcSigningAlgs,omitempty" tf:"oidc_signing_algs,omitempty"`

	OidcUsernameClaim *string `json:"oidcUsernameClaim,omitempty" tf:"oidc_username_claim,omitempty"`

	OidcUsernamePrefix *string `json:"oidcUsernamePrefix,omitempty" tf:"oidc_username_prefix,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type OIDCConfigurationObservation struct {
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IssuerURL *string `json:"issuerUrl,omitempty" tf:"issuer_url,omitempty"`

	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	OidcCAContent *string `json:"oidcCaContent,omitempty" tf:"oidc_ca_content,omitempty"`

	OidcGroupsClaim []*string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	OidcGroupsPrefix *string `json:"oidcGroupsPrefix,omitempty" tf:"oidc_groups_prefix,omitempty"`

	OidcRequiredClaim []*string `json:"oidcRequiredClaim,omitempty" tf:"oidc_required_claim,omitempty"`

	OidcSigningAlgs []*string `json:"oidcSigningAlgs,omitempty" tf:"oidc_signing_algs,omitempty"`

	OidcUsernameClaim *string `json:"oidcUsernameClaim,omitempty" tf:"oidc_username_claim,omitempty"`

	OidcUsernamePrefix *string `json:"oidcUsernamePrefix,omitempty" tf:"oidc_username_prefix,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type OIDCConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	IssuerURL *string `json:"issuerUrl,omitempty" tf:"issuer_url,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/kube/v1alpha1.Cluster
	// +kubebuilder:validation:Optional
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// Reference to a Cluster in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDRef *v1.Reference `json:"kubeIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDSelector *v1.Selector `json:"kubeIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OidcCAContent *string `json:"oidcCaContent,omitempty" tf:"oidc_ca_content,omitempty"`

	// +kubebuilder:validation:Optional
	OidcGroupsClaim []*string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// +kubebuilder:validation:Optional
	OidcGroupsPrefix *string `json:"oidcGroupsPrefix,omitempty" tf:"oidc_groups_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	OidcRequiredClaim []*string `json:"oidcRequiredClaim,omitempty" tf:"oidc_required_claim,omitempty"`

	// +kubebuilder:validation:Optional
	OidcSigningAlgs []*string `json:"oidcSigningAlgs,omitempty" tf:"oidc_signing_algs,omitempty"`

	// +kubebuilder:validation:Optional
	OidcUsernameClaim *string `json:"oidcUsernameClaim,omitempty" tf:"oidc_username_claim,omitempty"`

	// +kubebuilder:validation:Optional
	OidcUsernamePrefix *string `json:"oidcUsernamePrefix,omitempty" tf:"oidc_username_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// OIDCConfigurationSpec defines the desired state of OIDCConfiguration
type OIDCConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OIDCConfigurationParameters `json:"forProvider"`
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
	InitProvider OIDCConfigurationInitParameters `json:"initProvider,omitempty"`
}

// OIDCConfigurationStatus defines the observed state of OIDCConfiguration.
type OIDCConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OIDCConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OIDCConfiguration is the Schema for the OIDCConfigurations API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type OIDCConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuerUrl) || (has(self.initProvider) && has(self.initProvider.issuerUrl))",message="spec.forProvider.issuerUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   OIDCConfigurationSpec   `json:"spec"`
	Status OIDCConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OIDCConfigurationList contains a list of OIDCConfigurations
type OIDCConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OIDCConfiguration `json:"items"`
}

// Repository type metadata.
var (
	OIDCConfiguration_Kind             = "OIDCConfiguration"
	OIDCConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OIDCConfiguration_Kind}.String()
	OIDCConfiguration_KindAPIVersion   = OIDCConfiguration_Kind + "." + CRDGroupVersion.String()
	OIDCConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(OIDCConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&OIDCConfiguration{}, &OIDCConfigurationList{})
}
