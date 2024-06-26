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

type PTRInitParameters struct {

	// The network address in cidr format under which record has to be created.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description about PTR record.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Dns View under which the zone has been created.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// The Extensible attributes of PTR record to be added/updated, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// Network view name of NIOS server.
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// The domain name in FQDN to which the record should point to.
	Ptrdname *string `json:"ptrdname,omitempty" tf:"ptrdname,omitempty"`

	// The name of the DNS PTR record in FQDN format
	RecordName *string `json:"recordName,omitempty" tf:"record_name,omitempty"`

	// TTL attribute value for the record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type PTRObservation struct {

	// The network address in cidr format under which record has to be created.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description about PTR record.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Dns View under which the zone has been created.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// The Extensible attributes of PTR record to be added/updated, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// Network view name of NIOS server.
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// The domain name in FQDN to which the record should point to.
	Ptrdname *string `json:"ptrdname,omitempty" tf:"ptrdname,omitempty"`

	// The name of the DNS PTR record in FQDN format
	RecordName *string `json:"recordName,omitempty" tf:"record_name,omitempty"`

	// TTL attribute value for the record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type PTRParameters struct {

	// The network address in cidr format under which record has to be created.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description about PTR record.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Dns View under which the zone has been created.
	// +kubebuilder:validation:Optional
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// The Extensible attributes of PTR record to be added/updated, as a map in JSON format
	// +kubebuilder:validation:Optional
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field
	// +kubebuilder:validation:Optional
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// Network view name of NIOS server.
	// +kubebuilder:validation:Optional
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// The domain name in FQDN to which the record should point to.
	// +kubebuilder:validation:Optional
	Ptrdname *string `json:"ptrdname,omitempty" tf:"ptrdname,omitempty"`

	// The name of the DNS PTR record in FQDN format
	// +kubebuilder:validation:Optional
	RecordName *string `json:"recordName,omitempty" tf:"record_name,omitempty"`

	// TTL attribute value for the record.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// PTRSpec defines the desired state of PTR
type PTRSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PTRParameters `json:"forProvider"`
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
	InitProvider PTRInitParameters `json:"initProvider,omitempty"`
}

// PTRStatus defines the observed state of PTR.
type PTRStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PTRObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PTR is the Schema for the PTRs API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,infoblox-nios}
type PTR struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ptrdname) || (has(self.initProvider) && has(self.initProvider.ptrdname))",message="spec.forProvider.ptrdname is a required parameter"
	Spec   PTRSpec   `json:"spec"`
	Status PTRStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PTRList contains a list of PTRs
type PTRList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PTR `json:"items"`
}

// Repository type metadata.
var (
	PTR_Kind             = "PTR"
	PTR_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PTR_Kind}.String()
	PTR_KindAPIVersion   = PTR_Kind + "." + CRDGroupVersion.String()
	PTR_GroupVersionKind = CRDGroupVersion.WithKind(PTR_Kind)
)

func init() {
	SchemeBuilder.Register(&PTR{}, &PTRList{})
}
