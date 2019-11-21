package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/vvb/terraform-provider-intersight/intersight"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: intersight.Provider,
	})
}
