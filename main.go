package main

import (
//	"github.com/matteoredaelli/steampipe-plugin-talend/talend"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: talend.Plugin})
}
