package view

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_network_view", func(r *config.Resource) {

		r.ShortGroup = "network"
	})
}
