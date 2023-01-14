package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/nfagerlund/terraform-provider-fakerancher/fakerancher"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fakerancher.Provider})
}
