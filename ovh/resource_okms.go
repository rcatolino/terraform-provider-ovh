package ovh

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/ovh/go-ovh/ovh"
	"github.com/ovh/terraform-provider-ovh/ovh/types"
)

var (
	_ resource.ResourceWithConfigure   = (*okmsResource)(nil)
	_ resource.ResourceWithImportState = (*okmsResource)(nil)
)

func NewOkmsResource() resource.Resource {
	return &okmsResource{}
}

type okmsResource struct {
	config *Config
}

func (r *okmsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_okms"
}

func (d *okmsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*Config)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Config, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.config = config
}

func (d *okmsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = OkmsResourceSchema(ctx)
}

func (r *okmsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)

	var data OkmsModel
	endpoint := "/v2/okms/resource/" + url.PathEscape(req.ID)

	if err := r.config.OVHClient.Get(endpoint, &data); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error calling Get %s", endpoint),
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *okmsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OkmsModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create order and wait for service to be delivered
	order := &OrderModel{
		Order:         data.Order,
		OvhSubsidiary: data.OvhSubsidiary,
		Plan:          data.Plan,
		PlanOption:    data.PlanOption,
	}

	if err := orderCreate(order, r.config, "okms"); err != nil {
		resp.Diagnostics.AddError("failed to create order", err.Error())
	}

	// Find service name from order
	orderID := order.Order.OrderId.ValueInt64()
	plans := []PlanValue{}
	resp.Diagnostics.Append(data.Plan.ElementsAs(ctx, &plans, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := serviceNameFromOrder(r.config.OVHClient, orderID, plans[0].PlanCode.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to retrieve service name", err.Error())
	}

	data.Id = types.TfStringValue{
		StringValue: basetypes.NewStringValue(id),
	}

	var responseData OkmsModel
	// Read updated resource & update corresponding tf resource state
	err = retry.RetryContext(ctx, 2*time.Minute, func() *retry.RetryError {
		// Read updated resource
		endpoint := "/v2/okms/resource" + url.PathEscape(id)
		if err := r.config.OVHClient.Get(endpoint, &responseData); err != nil {
			return retry.NonRetryableError(fmt.Errorf("error calling GET %s", endpoint))
		}

		resp.Diagnostics.AddWarning("Get "+endpoint+"failed", err.Error())
		return retry.RetryableError(errors.New("waiting for KMS creation timeout"))
	})

	if err != nil {
		resp.Diagnostics.AddError("Failed to create KMS, creation timeout", err.Error())
	}

	data.MergeWith(&responseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *okmsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OkmsModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "/v2/okms/resource/" + url.PathEscape(data.Id.ValueString())

	if err := r.config.OVHClient.Get(endpoint, &data); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error calling Get %s", endpoint),
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *okmsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, planData OkmsModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: there's nothing to update, we always have to recreate.
}

func (r *okmsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OkmsModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	serviceId, err := serviceIdFromResourceName(r.config.OVHClient, id)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error locating service okms %s", id),
			err.Error(),
		)
		return
	}

	type TerminateServiceOption struct {
		ServiceId int `json:"serviceId"`
	}

	terminate := func() (string, error) {
		log.Printf("[DEBUG] Will terminate okms %s (service %d)", id, serviceId)
		endpoint := fmt.Sprintf("/services/%d/terminate", serviceId)
		if err := r.config.OVHClient.Post(endpoint, nil, nil); err != nil {
			if errOvh, ok := err.(*ovh.APIError); ok && (errOvh.Code == 404 || errOvh.Code == 460) {
				return "", nil
			}
			return "", fmt.Errorf("calling Post %s:\n\t %q", endpoint, err)
		}
		return id, nil
	}

	confirmTerminate := func(token string) error {
		log.Printf("[DEBUG] Will confirm termination of okms %s", id)
		endpoint := fmt.Sprintf("/services/%d/terminate/confirm", serviceId)
		if err := r.config.OVHClient.Post(endpoint, &ConfirmTerminationOpts{Token: token}, nil); err != nil {
			return fmt.Errorf("calling Post %s:\n\t %q", endpoint, err)
		}
		return nil
	}

	if err := orderDelete(r.config, terminate, confirmTerminate); err != nil {
		resp.Diagnostics.AddError("failed to delete resource", err.Error())
		return
	}
}
