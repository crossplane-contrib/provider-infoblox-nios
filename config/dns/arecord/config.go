package arecord

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_a_record", func(r *config.Resource) {
		r.ShortGroup = "dns"
		// r.Kind = "A"
		r.ExternalName = config.NameAsIdentifier
	})
}
