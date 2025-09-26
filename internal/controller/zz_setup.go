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
	association "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ip/association"
	network "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv4/network"
	networkcontainer "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ipv4/networkcontainer"
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
		association.Setup,
		network.Setup,
		networkcontainer.Setup,
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
