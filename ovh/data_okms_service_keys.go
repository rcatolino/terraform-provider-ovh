package ovh

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigure = (*okmsServiceKeysDataSource)(nil)

func NewOkmsServiceKeysDataSource() datasource.DataSource {
	return &okmsServiceKeysDataSource{}
}

type okmsServiceKeysDataSource struct {
	config *Config
}

func (d *okmsServiceKeysDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_okms_service_keys"
}

func (d *okmsServiceKeysDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*Config)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Config, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.config = config
}

func (d *okmsServiceKeysDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = OkmsServiceKeysDataSourceSchema(ctx)
}

func (d *okmsServiceKeysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OkmsServiceKeysModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "/okms/resource/" + url.PathEscape(data.OkmsId.ValueString()) + "/serviceKey"

	var arr []OkmsServiceKeysValue

	if err := d.config.OVHClient.Get(endpoint, &arr); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error calling Get %s", endpoint),
			err.Error(),
		)
		return
	}

	var b []attr.Value
	for _, a := range arr {
		b = append(b, a)
	}

	data.OkmsServiceKeys = ovhtypes.TfListNestedValue[OkmsServiceKeysValue]{
		ListValue: basetypes.NewListValueMust(OkmsServiceKeysValue{}.Type(ctx), b),
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
