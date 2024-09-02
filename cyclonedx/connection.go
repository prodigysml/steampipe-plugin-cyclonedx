package cyclonedx

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type cyclonedxConfig struct{}

var ConfigSchema = map[string]*schema.Attribute{}

func ConfigInstance() interface{} {
	return &cyclonedxConfig{}
}

func GetConfig(connection *plugin.Connection) cyclonedxConfig {
	if connection == nil || connection.Config == nil {
		return cyclonedxConfig{}
	}
	config, _ := connection.Config.(cyclonedxConfig)
	return config
}
