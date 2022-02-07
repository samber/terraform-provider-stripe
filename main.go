package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/samber/terraform-provider-stripe/stripe"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: stripe.Provider})
}
