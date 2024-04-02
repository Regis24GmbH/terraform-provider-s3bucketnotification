package main

import (
	"github.com/Regis24GmbH/terraform-provider-s3bucketnotification/internal/pkg/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
