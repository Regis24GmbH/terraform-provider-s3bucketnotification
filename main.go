package main

import (
	"github.com/Regis24GmbH/terraform-s3bucket-notifications/internal/pkg/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
