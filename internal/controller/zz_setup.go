/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	aaaa "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/aaaa"
	cname "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/cname"
	mx "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/mx"
	ptr "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/ptr"
	record "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/record"
	srv "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/srv"
	txt "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/txt"
	allocation "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ip/allocation"
	association "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ip/association"
	allocationipv4 "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv4/allocation"
	associationipv4 "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv4/association"
	network "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv4/network"
	networkcontainer "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv4/networkcontainer"
	allocationipv6 "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv6/allocation"
	associationipv6 "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv6/association"
	networkipv6 "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv6/network"
	networkcontaineripv6 "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv6/networkcontainer"
	view "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/network/view"
	providerconfig "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aaaa.Setup,
		cname.Setup,
		mx.Setup,
		ptr.Setup,
		record.Setup,
		srv.Setup,
		txt.Setup,
		allocation.Setup,
		association.Setup,
		allocationipv4.Setup,
		associationipv4.Setup,
		network.Setup,
		networkcontainer.Setup,
		allocationipv6.Setup,
		associationipv6.Setup,
		networkipv6.Setup,
		networkcontaineripv6.Setup,
		view.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
