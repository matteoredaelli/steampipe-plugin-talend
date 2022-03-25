package talend

import (
	"context"
	"errors"
	"os"

	"github.com/matteoredaelli/talend-go"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*talend.Client, error) {

	base_url := os.Getenv("TALEND_BASE_URL")
	api_key := os.Getenv("TALEND_API_KEY")

	talendConfig := GetConfig(d.Connection)
	if &talendConfig != nil {
		if talendConfig.BaseUrl != nil {
			base_url = *talendConfig.BaseUrl
		}
		if talendConfig.Email != nil {
			uapi_key = *talendConfig.ApiKey
		}
	}

	if base_url == "" {
		return nil, errors.New("'base_url' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if api_key == "" {
		return nil, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// You can set custom *http.Client here
	client, err := talend.NewClient(base_url, api_key)
	if err != nil {
		return nil, err
	}

	return client, nil
}
