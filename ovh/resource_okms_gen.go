package ovh

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OkmsResourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"iam": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"display_name": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Resource display name",
					MarkdownDescription: "Resource display name",
				},
				"id": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Unique identifier of the resource",
					MarkdownDescription: "Unique identifier of the resource",
				},
				"tags": schema.MapAttribute{
					CustomType:          ovhtypes.NewTfMapNestedType[ovhtypes.TfStringValue](ctx),
					Computed:            true,
					Description:         "Resource tags. Tags that were internally computed are prefixed with ovh:",
					MarkdownDescription: "Resource tags. Tags that were internally computed are prefixed with ovh:",
				},
				"urn": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "Unique resource name used in policies",
					MarkdownDescription: "Unique resource name used in policies",
				},
			},
			CustomType: IamType{
				ObjectType: types.ObjectType{
					AttrTypes: IamValue{}.AttributeTypes(ctx),
				},
			},
			Computed:            true,
			Description:         "IAM resource metadata",
			MarkdownDescription: "IAM resource metadata",
		},
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "OKMS ID",
			MarkdownDescription: "OKMS ID",
		},
		"kmip_endpoint": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS kmip API endpoint",
			MarkdownDescription: "KMS kmip API endpoint",
		},
		"public_ca": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Optional:            true,
			Computed:            true,
			Description:         "KMS public CA (Certificate Authority)",
			MarkdownDescription: "KMS public CA (Certificate Authority)",
		},
		// TODO: add region as a required first level parameter instead of a plan option
		"rest_endpoint": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS rest API endpoint",
			MarkdownDescription: "KMS rest API endpoint",
		},
		"swagger_endpoint": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS rest API swagger UI",
			MarkdownDescription: "KMS rest API swagger UI",
		},
	}

	// Add order attributes
	for k, v := range OrderResourceSchema(ctx).Attributes {
		attrs[k] = v
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type OkmsModel struct {
	Iam             IamValue                                    `tfsdk:"iam" json:"iam"`
	Id              ovhtypes.TfStringValue                      `tfsdk:"id" json:"id"`
	KmipEndpoint    ovhtypes.TfStringValue                      `tfsdk:"kmip_endpoint" json:"kmipEndpoint"`
	PublicCa        ovhtypes.TfStringValue                      `tfsdk:"public_ca" json:"publicCa"`
	RestEndpoint    ovhtypes.TfStringValue                      `tfsdk:"rest_endpoint" json:"restEndpoint"`
	SwaggerEndpoint ovhtypes.TfStringValue                      `tfsdk:"swagger_endpoint" json:"swaggerEndpoint"`
	Order           OrderValue                                  `tfsdk:"order" json:"order"`
	OvhSubsidiary   ovhtypes.TfStringValue                      `tfsdk:"ovh_subsidiary" json:"ovhSubsidiary"`
	Plan            ovhtypes.TfListNestedValue[PlanValue]       `tfsdk:"plan" json:"plan"`
	PlanOption      ovhtypes.TfListNestedValue[PlanOptionValue] `tfsdk:"plan_option" json:"planOption"`
}

func (v *OkmsModel) MergeWith(other *OkmsModel) {

	if (v.Iam.IsUnknown() || v.Iam.IsNull()) && !other.Iam.IsUnknown() {
		v.Iam = other.Iam
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.KmipEndpoint.IsUnknown() || v.KmipEndpoint.IsNull()) && !other.KmipEndpoint.IsUnknown() {
		v.KmipEndpoint = other.KmipEndpoint
	}

	if (v.PublicCa.IsUnknown() || v.PublicCa.IsNull()) && !other.PublicCa.IsUnknown() {
		v.PublicCa = other.PublicCa
	}

	if (v.RestEndpoint.IsUnknown() || v.RestEndpoint.IsNull()) && !other.RestEndpoint.IsUnknown() {
		v.RestEndpoint = other.RestEndpoint
	}

	if (v.SwaggerEndpoint.IsUnknown() || v.SwaggerEndpoint.IsNull()) && !other.SwaggerEndpoint.IsUnknown() {
		v.SwaggerEndpoint = other.SwaggerEndpoint
	}

	if v.Order.IsUnknown() && !other.Order.IsUnknown() {
		v.Order = other.Order
	} else if !other.Order.IsUnknown() {
		v.Order.MergeWith(&other.Order)
	}

	if (v.OvhSubsidiary.IsUnknown() || v.OvhSubsidiary.IsNull()) && !other.OvhSubsidiary.IsUnknown() {
		v.OvhSubsidiary = other.OvhSubsidiary
	}

	if (v.Plan.IsUnknown() || v.Plan.IsNull()) && !other.Plan.IsUnknown() {
		v.Plan = other.Plan
	}

	if (v.PlanOption.IsUnknown() || v.PlanOption.IsNull()) && !other.PlanOption.IsUnknown() {
		v.PlanOption = other.PlanOption
	}

}
