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

type RecordInitParameters struct {

	// Network to allocate an IP address from, when the 'ip_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of the A-record.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// DNS view which the zone does exist within.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// Extensible attributes of the A-record to be added/updated, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// FQDN for the A-record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields.
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty.
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// TTL value for the A-record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type RecordObservation struct {

	// Network to allocate an IP address from, when the 'ip_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of the A-record.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// DNS view which the zone does exist within.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// Extensible attributes of the A-record to be added/updated, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// FQDN for the A-record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields.
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty.
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// TTL value for the A-record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type RecordParameters struct {

	// Network to allocate an IP address from, when the 'ip_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of the A-record.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// DNS view which the zone does exist within.
	// +kubebuilder:validation:Optional
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// Extensible attributes of the A-record to be added/updated, as a map in JSON format
	// +kubebuilder:validation:Optional
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// FQDN for the A-record.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields.
	// +kubebuilder:validation:Optional
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty.
	// +kubebuilder:validation:Optional
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// TTL value for the A-record.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
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
	InitProvider RecordInitParameters `json:"initProvider,omitempty"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Record is the Schema for the Records API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,infoblox-nios}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fqdn) || (has(self.initProvider) && has(self.initProvider.fqdn))",message="spec.forProvider.fqdn is a required parameter"
	Spec   RecordSpec   `json:"spec"`
	Status RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
