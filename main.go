package main

import (
	"github.com/cisco-intersight/terraform-provider-intersight/intersight"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: intersight.Provider,
	})
}
