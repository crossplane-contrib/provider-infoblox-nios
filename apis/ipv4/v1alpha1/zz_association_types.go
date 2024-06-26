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

type AssociationInitParameters struct {

	// The address in cidr format.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description of the IP association.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// view in which record has to be created.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// DHCP unique identifier for IPv6.
	Duid *string `json:"duid,omitempty" tf:"duid,omitempty"`

	// flag that defines if the host record is to be used for IPAM Purposes.
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// flag that defines if the host record is to be used for DNS Purposes
	EnableDNS *bool `json:"enableDns,omitempty" tf:"enable_dns,omitempty"`

	// The Extensible attributes for IP Association, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// The host name for Host Record in FQDN format.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// IP address of cloud instance.
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// mac address of cloud instance.
	MacAddr *string `json:"macAddr,omitempty" tf:"mac_addr,omitempty"`

	// Network view name of NIOS server.
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// TTL attribute value for the record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type AssociationObservation struct {

	// The address in cidr format.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description of the IP association.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// view in which record has to be created.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// DHCP unique identifier for IPv6.
	Duid *string `json:"duid,omitempty" tf:"duid,omitempty"`

	// flag that defines if the host record is to be used for IPAM Purposes.
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// flag that defines if the host record is to be used for DNS Purposes
	EnableDNS *bool `json:"enableDns,omitempty" tf:"enable_dns,omitempty"`

	// The Extensible attributes for IP Association, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// The host name for Host Record in FQDN format.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of cloud instance.
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// mac address of cloud instance.
	MacAddr *string `json:"macAddr,omitempty" tf:"mac_addr,omitempty"`

	// Network view name of NIOS server.
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// TTL attribute value for the record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type AssociationParameters struct {

	// The address in cidr format.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description of the IP association.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// view in which record has to be created.
	// +kubebuilder:validation:Optional
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// DHCP unique identifier for IPv6.
	// +kubebuilder:validation:Optional
	Duid *string `json:"duid,omitempty" tf:"duid,omitempty"`

	// flag that defines if the host record is to be used for IPAM Purposes.
	// +kubebuilder:validation:Optional
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// flag that defines if the host record is to be used for DNS Purposes
	// +kubebuilder:validation:Optional
	EnableDNS *bool `json:"enableDns,omitempty" tf:"enable_dns,omitempty"`

	// The Extensible attributes for IP Association, as a map in JSON format
	// +kubebuilder:validation:Optional
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// The host name for Host Record in FQDN format.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// IP address of cloud instance.
	// +kubebuilder:validation:Optional
	IPAddr *string `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	// mac address of cloud instance.
	// +kubebuilder:validation:Optional
	MacAddr *string `json:"macAddr,omitempty" tf:"mac_addr,omitempty"`

	// Network view name of NIOS server.
	// +kubebuilder:validation:Optional
	NetworkView *string `json:"networkView,omitempty" tf:"network_view,omitempty"`

	// TTL attribute value for the record.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// AssociationSpec defines the desired state of Association
type AssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssociationParameters `json:"forProvider"`
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
	InitProvider AssociationInitParameters `json:"initProvider,omitempty"`
}

// AssociationStatus defines the observed state of Association.
type AssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Association is the Schema for the Associations API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,infoblox-nios}
type Association struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fqdn) || (has(self.initProvider) && has(self.initProvider.fqdn))",message="spec.forProvider.fqdn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddr) || (has(self.initProvider) && has(self.initProvider.ipAddr))",message="spec.forProvider.ipAddr is a required parameter"
	Spec   AssociationSpec   `json:"spec"`
	Status AssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssociationList contains a list of Associations
type AssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Association `json:"items"`
}

// Repository type metadata.
var (
	Association_Kind             = "Association"
	Association_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Association_Kind}.String()
	Association_KindAPIVersion   = Association_Kind + "." + CRDGroupVersion.String()
	Association_GroupVersionKind = CRDGroupVersion.WithKind(Association_Kind)
)

func init() {
	SchemeBuilder.Register(&Association{}, &AssociationList{})
}
