/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/aaaarecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/arecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/cnamerecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/mxrecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/ptrrecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/srvrecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/dns/txtrecord"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ip/allocation"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ip/association"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv4/ipv4allocation"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv4/ipv4association"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv4/ipv4network"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv4/ipv4networkcontainer"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv6/ipv6allocation"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv6/ipv6association"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv6/ipv6network"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/ipv6/ipv6networkcontainer"
	"github.com/crossplane-contrib/provider-infoblox-nios/config/network/view"
)

const (
	resourcePrefix = "infoblox-nios"
	rootGroup      = resourcePrefix + "." + "crossplane.io"
	modulePath     = "github.com/crossplane-contrib/provider-infoblox-nios"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom resource configurators
		allocation.Configure,
		association.Configure,
		ipv4allocation.Configure,
		ipv4association.Configure,
		ipv4network.Configure,
		ipv4networkcontainer.Configure,
		ipv6allocation.Configure,
		ipv6association.Configure,
		ipv6network.Configure,
		ipv6networkcontainer.Configure,
		arecord.Configure,
		aaaarecord.Configure,
		cnamerecord.Configure,
		mxrecord.Configure,
		ptrrecord.Configure,
		srvrecord.Configure,
		txtrecord.Configure,
		view.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
