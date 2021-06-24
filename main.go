package main

import (
	"github.com/AliAllomani/terraform-provider-hue/hue"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: hue.Provider})
}
