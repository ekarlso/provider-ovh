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

type ConfigurationInitParameters struct {

	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationObservation struct {

	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationParameters struct {

	// Identifier of the resource
	// +kubebuilder:validation:Optional
	Label *string `json:"label" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DetailsInitParameters struct {
}

type DetailsObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	OrderDetailID *float64 `json:"orderDetailId,omitempty" tf:"order_detail_id,omitempty"`

	Quantity *string `json:"quantity,omitempty" tf:"quantity,omitempty"`
}

type DetailsParameters struct {
}

type OrderInitParameters struct {
}

type OrderObservation struct {

	// date
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Information about a Bill entry
	Details []DetailsObservation `json:"details,omitempty" tf:"details,omitempty"`

	// expiration date
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// order id
	OrderID *float64 `json:"orderId,omitempty" tf:"order_id,omitempty"`
}

type OrderParameters struct {
}

type PlanInitParameters struct {

	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanObservation struct {

	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanOptionConfigurationInitParameters struct {

	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PlanOptionConfigurationObservation struct {

	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PlanOptionConfigurationParameters struct {

	// Identifier of the resource
	// +kubebuilder:validation:Optional
	Label *string `json:"label" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type PlanOptionInitParameters struct {

	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	Configuration []PlanOptionConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanOptionObservation struct {

	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	Configuration []PlanOptionConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanOptionParameters struct {

	// Catalog name
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// +kubebuilder:validation:Optional
	Configuration []PlanOptionConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// Plan code
	// +kubebuilder:validation:Optional
	PlanCode *string `json:"planCode" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// +kubebuilder:validation:Optional
	PricingMode *string `json:"pricingMode" tf:"pricing_mode,omitempty"`
}

type PlanParameters struct {

	// Catalog name
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// Plan code
	// +kubebuilder:validation:Optional
	PlanCode *string `json:"planCode" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// +kubebuilder:validation:Optional
	PricingMode *string `json:"pricingMode" tf:"pricing_mode,omitempty"`
}

type VrackInitParameters struct {

	// yourvrackdescription
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// yourvrackname
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Details about an Order
	Order []OrderInitParameters `json:"order,omitempty" tf:"order,omitempty"`

	// Ovh Subsidiary
	OvhSubsidiary *string `json:"ovhSubsidiary,omitempty" tf:"ovh_subsidiary,omitempty"`

	// Ovh payment mode
	PaymentMean *string `json:"paymentMean,omitempty" tf:"payment_mean,omitempty"`

	// Product Plan to order
	Plan []PlanInitParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Product Plan to order
	PlanOption []PlanOptionInitParameters `json:"planOption,omitempty" tf:"plan_option,omitempty"`
}

type VrackObservation struct {

	// yourvrackdescription
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// yourvrackname
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Details about an Order
	Order []OrderObservation `json:"order,omitempty" tf:"order,omitempty"`

	// Ovh Subsidiary
	OvhSubsidiary *string `json:"ovhSubsidiary,omitempty" tf:"ovh_subsidiary,omitempty"`

	// Ovh payment mode
	PaymentMean *string `json:"paymentMean,omitempty" tf:"payment_mean,omitempty"`

	// Product Plan to order
	Plan []PlanObservation `json:"plan,omitempty" tf:"plan,omitempty"`

	// Product Plan to order
	PlanOption []PlanOptionObservation `json:"planOption,omitempty" tf:"plan_option,omitempty"`

	// The internal name of your vrack
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type VrackParameters struct {

	// yourvrackdescription
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// yourvrackname
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Details about an Order
	// +kubebuilder:validation:Optional
	Order []OrderParameters `json:"order,omitempty" tf:"order,omitempty"`

	// Ovh Subsidiary
	// +kubebuilder:validation:Optional
	OvhSubsidiary *string `json:"ovhSubsidiary,omitempty" tf:"ovh_subsidiary,omitempty"`

	// Ovh payment mode
	// +kubebuilder:validation:Optional
	PaymentMean *string `json:"paymentMean,omitempty" tf:"payment_mean,omitempty"`

	// Product Plan to order
	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Product Plan to order
	// +kubebuilder:validation:Optional
	PlanOption []PlanOptionParameters `json:"planOption,omitempty" tf:"plan_option,omitempty"`
}

// VrackSpec defines the desired state of Vrack
type VrackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VrackParameters `json:"forProvider"`
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
	InitProvider VrackInitParameters `json:"initProvider,omitempty"`
}

// VrackStatus defines the observed state of Vrack.
type VrackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VrackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Vrack is the Schema for the Vracks API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type Vrack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ovhSubsidiary) || (has(self.initProvider) && has(self.initProvider.ovhSubsidiary))",message="spec.forProvider.ovhSubsidiary is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plan) || (has(self.initProvider) && has(self.initProvider.plan))",message="spec.forProvider.plan is a required parameter"
	Spec   VrackSpec   `json:"spec"`
	Status VrackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VrackList contains a list of Vracks
type VrackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vrack `json:"items"`
}

// Repository type metadata.
var (
	Vrack_Kind             = "Vrack"
	Vrack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vrack_Kind}.String()
	Vrack_KindAPIVersion   = Vrack_Kind + "." + CRDGroupVersion.String()
	Vrack_GroupVersionKind = CRDGroupVersion.WithKind(Vrack_Kind)
)

func init() {
	SchemeBuilder.Register(&Vrack{}, &VrackList{})
}
