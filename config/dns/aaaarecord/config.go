package aaaarecord

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_aaaa_record", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "AAAA"
	})
}
