package app

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("cloudfoundry_app", func(r *config.Resource) {
        r.ShortGroup = "app"
    })
}
