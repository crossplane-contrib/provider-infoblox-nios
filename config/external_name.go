/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"infoblox_ip_allocation":          config.IdentifierFromProvider,
	"infoblox_ip_association":         config.IdentifierFromProvider,
	"infoblox_ipv4_allocation":        config.IdentifierFromProvider,
	"infoblox_ipv4_association":       config.IdentifierFromProvider,
	"infoblox_ipv4_network":           config.IdentifierFromProvider,
	"infoblox_ipv4_network_container": config.IdentifierFromProvider,
	"infoblox_ipv6_allocation":        config.IdentifierFromProvider,
	"infoblox_ipv6_association":       config.IdentifierFromProvider,
	"infoblox_ipv6_network":           config.IdentifierFromProvider,
	"infoblox_ipv6_network_container": config.IdentifierFromProvider,
	"infoblox_a_record":               config.IdentifierFromProvider,
	"infoblox_aaaa_record":            config.IdentifierFromProvider,
	"infoblox_cname_record":           config.IdentifierFromProvider,
	"infoblox_srv_record":             config.IdentifierFromProvider,
	"infoblox_txt_record":             config.IdentifierFromProvider,
	"infoblox_ptr_record":             config.IdentifierFromProvider,
	"infoblox_mx_record":              config.IdentifierFromProvider,
	"infoblox_network_view":           config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
