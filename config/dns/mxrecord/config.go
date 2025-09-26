package mxrecord

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_mx_record", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "MX"
	})
}
