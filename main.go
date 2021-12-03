package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/sh-miyoshi/terraform-provider-sample/sample"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sample.Provider,
	})
}
