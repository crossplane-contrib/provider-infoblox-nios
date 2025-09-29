/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	aaaarecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/aaaarecord"
	arecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/arecord"
	cnamerecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/cnamerecord"
	mxrecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/mxrecord"
	ptrrecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/ptrrecord"
	srvrecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/srvrecord"
	txtrecord "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/dns/txtrecord"
	allocation "github.com/crossplane-contrib/provider-infoblox-nios/internal/controller/ip/allocation"
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
		aaaarecord.Setup,
		arecord.Setup,
		cnamerecord.Setup,
		mxrecord.Setup,
		ptrrecord.Setup,
		srvrecord.Setup,
		txtrecord.Setup,
		allocation.Setup,
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
