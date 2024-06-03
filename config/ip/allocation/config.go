package allocation

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_ip_allocation", func(r *config.Resource) {

		r.ShortGroup = "ip"
		r.ExternalName = config.NameAsIdentifier
	})
}
