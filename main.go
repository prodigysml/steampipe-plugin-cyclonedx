package main

import (
	"github.com/prodigysml/steampipe-plugin-cyclonedx/cyclonedx"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: cyclonedx.Plugin,
	})
}
