package main

import (
	"bitbucket-eng-sjc1.cisco.com/an/terraform-provider-intersight/intersight"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: intersight.Provider,
	})
}
