//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allocation) DeepCopyInto(out *Allocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allocation.
func (in *Allocation) DeepCopy() *Allocation {
	if in == nil {
		return nil
	}
	out := new(Allocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Allocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationInitParameters) DeepCopyInto(out *AllocationInitParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.DNSView != nil {
		in, out := &in.DNSView, &out.DNSView
		*out = new(string)
		**out = **in
	}
	if in.EnableDNS != nil {
		in, out := &in.EnableDNS, &out.EnableDNS
		*out = new(bool)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.IPAddr != nil {
		in, out := &in.IPAddr, &out.IPAddr
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationInitParameters.
func (in *AllocationInitParameters) DeepCopy() *AllocationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AllocationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationList) DeepCopyInto(out *AllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Allocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationList.
func (in *AllocationList) DeepCopy() *AllocationList {
	if in == nil {
		return nil
	}
	out := new(AllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationObservation) DeepCopyInto(out *AllocationObservation) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.DNSView != nil {
		in, out := &in.DNSView, &out.DNSView
		*out = new(string)
		**out = **in
	}
	if in.EnableDNS != nil {
		in, out := &in.EnableDNS, &out.EnableDNS
		*out = new(bool)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPAddr != nil {
		in, out := &in.IPAddr, &out.IPAddr
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationObservation.
func (in *AllocationObservation) DeepCopy() *AllocationObservation {
	if in == nil {
		return nil
	}
	out := new(AllocationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationParameters) DeepCopyInto(out *AllocationParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.DNSView != nil {
		in, out := &in.DNSView, &out.DNSView
		*out = new(string)
		**out = **in
	}
	if in.EnableDNS != nil {
		in, out := &in.EnableDNS, &out.EnableDNS
		*out = new(bool)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.IPAddr != nil {
		in, out := &in.IPAddr, &out.IPAddr
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationParameters.
func (in *AllocationParameters) DeepCopy() *AllocationParameters {
	if in == nil {
		return nil
	}
	out := new(AllocationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationSpec) DeepCopyInto(out *AllocationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationSpec.
func (in *AllocationSpec) DeepCopy() *AllocationSpec {
	if in == nil {
		return nil
	}
	out := new(AllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationStatus) DeepCopyInto(out *AllocationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationStatus.
func (in *AllocationStatus) DeepCopy() *AllocationStatus {
	if in == nil {
		return nil
	}
	out := new(AllocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Association) DeepCopyInto(out *Association) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Association.
func (in *Association) DeepCopy() *Association {
	if in == nil {
		return nil
	}
	out := new(Association)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Association) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationInitParameters) DeepCopyInto(out *AssociationInitParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.DNSView != nil {
		in, out := &in.DNSView, &out.DNSView
		*out = new(string)
		**out = **in
	}
	if in.Duid != nil {
		in, out := &in.Duid, &out.Duid
		*out = new(string)
		**out = **in
	}
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNS != nil {
		in, out := &in.EnableDNS, &out.EnableDNS
		*out = new(bool)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.IPAddr != nil {
		in, out := &in.IPAddr, &out.IPAddr
		*out = new(string)
		**out = **in
	}
	if in.MacAddr != nil {
		in, out := &in.MacAddr, &out.MacAddr
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationInitParameters.
func (in *AssociationInitParameters) DeepCopy() *AssociationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AssociationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationList) DeepCopyInto(out *AssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Association, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationList.
func (in *AssociationList) DeepCopy() *AssociationList {
	if in == nil {
		return nil
	}
	out := new(AssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationObservation) DeepCopyInto(out *AssociationObservation) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.DNSView != nil {
		in, out := &in.DNSView, &out.DNSView
		*out = new(string)
		**out = **in
	}
	if in.Duid != nil {
		in, out := &in.Duid, &out.Duid
		*out = new(string)
		**out = **in
	}
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNS != nil {
		in, out := &in.EnableDNS, &out.EnableDNS
		*out = new(bool)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPAddr != nil {
		in, out := &in.IPAddr, &out.IPAddr
		*out = new(string)
		**out = **in
	}
	if in.MacAddr != nil {
		in, out := &in.MacAddr, &out.MacAddr
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationObservation.
func (in *AssociationObservation) DeepCopy() *AssociationObservation {
	if in == nil {
		return nil
	}
	out := new(AssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationParameters) DeepCopyInto(out *AssociationParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.DNSView != nil {
		in, out := &in.DNSView, &out.DNSView
		*out = new(string)
		**out = **in
	}
	if in.Duid != nil {
		in, out := &in.Duid, &out.Duid
		*out = new(string)
		**out = **in
	}
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNS != nil {
		in, out := &in.EnableDNS, &out.EnableDNS
		*out = new(bool)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.IPAddr != nil {
		in, out := &in.IPAddr, &out.IPAddr
		*out = new(string)
		**out = **in
	}
	if in.MacAddr != nil {
		in, out := &in.MacAddr, &out.MacAddr
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationParameters.
func (in *AssociationParameters) DeepCopy() *AssociationParameters {
	if in == nil {
		return nil
	}
	out := new(AssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationSpec) DeepCopyInto(out *AssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationSpec.
func (in *AssociationSpec) DeepCopy() *AssociationSpec {
	if in == nil {
		return nil
	}
	out := new(AssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationStatus) DeepCopyInto(out *AssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationStatus.
func (in *AssociationStatus) DeepCopy() *AssociationStatus {
	if in == nil {
		return nil
	}
	out := new(AssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Network) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainer) DeepCopyInto(out *NetworkContainer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainer.
func (in *NetworkContainer) DeepCopy() *NetworkContainer {
	if in == nil {
		return nil
	}
	out := new(NetworkContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkContainer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerInitParameters) DeepCopyInto(out *NetworkContainerInitParameters) {
	*out = *in
	if in.AllocatePrefixLen != nil {
		in, out := &in.AllocatePrefixLen, &out.AllocatePrefixLen
		*out = new(float64)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.ParentCidr != nil {
		in, out := &in.ParentCidr, &out.ParentCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerInitParameters.
func (in *NetworkContainerInitParameters) DeepCopy() *NetworkContainerInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerList) DeepCopyInto(out *NetworkContainerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerList.
func (in *NetworkContainerList) DeepCopy() *NetworkContainerList {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkContainerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerObservation) DeepCopyInto(out *NetworkContainerObservation) {
	*out = *in
	if in.AllocatePrefixLen != nil {
		in, out := &in.AllocatePrefixLen, &out.AllocatePrefixLen
		*out = new(float64)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.ParentCidr != nil {
		in, out := &in.ParentCidr, &out.ParentCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerObservation.
func (in *NetworkContainerObservation) DeepCopy() *NetworkContainerObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerParameters) DeepCopyInto(out *NetworkContainerParameters) {
	*out = *in
	if in.AllocatePrefixLen != nil {
		in, out := &in.AllocatePrefixLen, &out.AllocatePrefixLen
		*out = new(float64)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.ParentCidr != nil {
		in, out := &in.ParentCidr, &out.ParentCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerParameters.
func (in *NetworkContainerParameters) DeepCopy() *NetworkContainerParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerSpec) DeepCopyInto(out *NetworkContainerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerSpec.
func (in *NetworkContainerSpec) DeepCopy() *NetworkContainerSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerStatus) DeepCopyInto(out *NetworkContainerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerStatus.
func (in *NetworkContainerStatus) DeepCopy() *NetworkContainerStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInitParameters) DeepCopyInto(out *NetworkInitParameters) {
	*out = *in
	if in.AllocatePrefixLen != nil {
		in, out := &in.AllocatePrefixLen, &out.AllocatePrefixLen
		*out = new(float64)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.ParentCidr != nil {
		in, out := &in.ParentCidr, &out.ParentCidr
		*out = new(string)
		**out = **in
	}
	if in.ReserveIP != nil {
		in, out := &in.ReserveIP, &out.ReserveIP
		*out = new(float64)
		**out = **in
	}
	if in.ReserveIPv6 != nil {
		in, out := &in.ReserveIPv6, &out.ReserveIPv6
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInitParameters.
func (in *NetworkInitParameters) DeepCopy() *NetworkInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkList) DeepCopyInto(out *NetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkList.
func (in *NetworkList) DeepCopy() *NetworkList {
	if in == nil {
		return nil
	}
	out := new(NetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.AllocatePrefixLen != nil {
		in, out := &in.AllocatePrefixLen, &out.AllocatePrefixLen
		*out = new(float64)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.ParentCidr != nil {
		in, out := &in.ParentCidr, &out.ParentCidr
		*out = new(string)
		**out = **in
	}
	if in.ReserveIP != nil {
		in, out := &in.ReserveIP, &out.ReserveIP
		*out = new(float64)
		**out = **in
	}
	if in.ReserveIPv6 != nil {
		in, out := &in.ReserveIPv6, &out.ReserveIPv6
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.AllocatePrefixLen != nil {
		in, out := &in.AllocatePrefixLen, &out.AllocatePrefixLen
		*out = new(float64)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ExtAttrs != nil {
		in, out := &in.ExtAttrs, &out.ExtAttrs
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.NetworkView != nil {
		in, out := &in.NetworkView, &out.NetworkView
		*out = new(string)
		**out = **in
	}
	if in.ParentCidr != nil {
		in, out := &in.ParentCidr, &out.ParentCidr
		*out = new(string)
		**out = **in
	}
	if in.ReserveIP != nil {
		in, out := &in.ReserveIP, &out.ReserveIP
		*out = new(float64)
		**out = **in
	}
	if in.ReserveIPv6 != nil {
		in, out := &in.ReserveIPv6, &out.ReserveIPv6
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}
