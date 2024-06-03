package ipv4allocation

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_ipv4_allocation", func(r *config.Resource) {

		r.ShortGroup = "ipv4"
		r.ExternalName = config.NameAsIdentifier
	})
}
