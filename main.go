package main

import (
	"github.com/hashicorp/terraform/plugin"
	statuspage "github.com/nagelflorian/terraform-provider-statuspage/statuspage"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: statuspage.Provider,
	})
}
