package talend

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type talendConfig struct {
	BaseUrl *string `cty:"base_url"`
	ApiKey  *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"base_url": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &talendConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) talendConfig {
	if connection == nil || connection.Config == nil {
		return talendConfig{}
	}
	config, _ := connection.Config.(talendConfig)
	return config
}
