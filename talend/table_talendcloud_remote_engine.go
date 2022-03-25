package talend

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableTalendRemoteEngine() *plugin.Table {
	return &plugin.Table{
		Name:        "talend_remote_engine",
		Description: "When support requests arrive in Talend Support, they can be assigned to a Remote_engine. Remote_engines serve as the core element of ticket workflow; support agents are organized into Remote_engines and tickets can be assigned to a Remote_engine only, or to an assigned agent within a Remote_engine. A ticket can never be assigned to an agent without also being assigned to a Remote_engine.",
		List: &plugin.ListConfig{
			Hydrate: listRemoteEngine,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRemoteEngine,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the remote_engine"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the remote_engine"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the remote_engine"},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the remote_engine has been deleted"},
			{Name: "createDate", Type: proto.ColumnType_TIMESTAMP, Description: "The time the remote_engine was created"},
			{Name: "updateDate", Type: proto.ColumnType_TIMESTAMP, Description: "The time of the last update of the remote_engine"},
			{Name: "workspace", Type: proto.ColumnType_JSON, Description: "workspace "},
			{Name: "runtimeId", Type: proto.ColumnType_STRING, Description: "runtimeId"},
			{Name: "availability", Type: proto.ColumnType_STRING, Description: "availability"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "status"},
			{Name: "runProfiles", Type: proto.ColumnType_JSON, Description: "runProfiles "},
			{Name: "clusterId", Type: proto.ColumnType_STRING, Description: "clusterId"},
			{Name: "esbCompatibilityVersion", Type: proto.ColumnType_STRING, Description: "esbCompatibilityVersion"},
		},
	}
}

func listRemoteEngine(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	// TODO - The library doesn't support paging for input?
	remote_engines, err := conn.Get("runtimes/remote-engines", nil)
	if err != nil {
		return nil, err
	}
	for _, t := range remote_engines {
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getRemoteEngine(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := conn.Get("runtimes/remote-engines/" + string(id), nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
