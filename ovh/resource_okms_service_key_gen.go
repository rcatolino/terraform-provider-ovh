package ovh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OkmsServiceKeyResourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"context": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Optional:            true,
			Computed:            true,
			Description:         "Context of the key",
			MarkdownDescription: "Context of the key",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplaceIfConfigured(),
			},
		},
		"created_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Creation time of the key",
			MarkdownDescription: "Creation time of the key",
		},
		"curve": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Optional:            true,
			Computed:            true,
			Description:         "Curve type for Elliptic Curve (EC) keys",
			MarkdownDescription: "Curve type for Elliptic Curve (EC) keys",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplaceIfConfigured(),
			},
			Validators: []validator.String{
				stringvalidator.OneOf(
					"P-256",
					"P-384",
					"P-521",
				),
			},
		},
		"deactivation_reason": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Key deactivation reason",
			MarkdownDescription: "Key deactivation reason",
			Validators: []validator.String{
				stringvalidator.OneOf(
					"AFFILIATION_CHANGED",
					"CA_COMPROMISE",
					"CESSATION_OF_OPERATION",
					"KEY_COMPROMISE",
					"PRIVILEGE_WITHDRAWN",
					"SUPERSEDED",
					"UNSPECIFIED",
				),
			},
		},
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Key ID",
			MarkdownDescription: "Key ID",
		},
		/*
			"keys": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"alg": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The algorithm intended to be used with the key",
							MarkdownDescription: "The algorithm intended to be used with the key",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ES256",
									"ES384",
									"ES512",
									"PS256",
									"PS384",
									"PS512",
									"RS256",
									"RS384",
									"RS512",
								),
							},
						},
						"crv": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The cryptographic curve used with the key",
							MarkdownDescription: "The cryptographic curve used with the key",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"P-256",
									"P-384",
									"P-521",
								),
							},
						},
						"d": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The RSA or EC private exponent",
							MarkdownDescription: "The RSA or EC private exponent",
						},
						"dp": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The RSA private key's first factor CRT exponent",
							MarkdownDescription: "The RSA private key's first factor CRT exponent",
						},
						"dq": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The RSA private key's second factor CRT exponent",
							MarkdownDescription: "The RSA private key's second factor CRT exponent",
						},
						"e": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The exponent value for the RSA public key",
							MarkdownDescription: "The exponent value for the RSA public key",
						},
						"k": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The value of the symmetric (or other single-valued) key",
							MarkdownDescription: "The value of the symmetric (or other single-valued) key",
						},
						"key_ops": schema.ListAttribute{
							CustomType:          ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
							Optional:            true,
							Computed:            true,
							Description:         "The operation for which the key is intended to be used",
							MarkdownDescription: "The operation for which the key is intended to be used",
						},
						"kid": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "key ID parameter used to match a specific key",
							MarkdownDescription: "key ID parameter used to match a specific key",
						},
						"kty": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC",
							MarkdownDescription: "Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"EC",
									"RSA",
									"oct",
								),
							},
						},
						"n": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The modulus value for the RSA public key",
							MarkdownDescription: "The modulus value for the RSA public key",
						},
						"p": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The first prime factor of the RSA private key",
							MarkdownDescription: "The first prime factor of the RSA private key",
						},
						"q": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The second prime factor of the RSA private key",
							MarkdownDescription: "The second prime factor of the RSA private key",
						},
						"qi": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
							Description:         "The CRT coefficient of the second factor of the RSA private key",
							MarkdownDescription: "The CRT coefficient of the second factor of the RSA private key",
						},
						"use": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The intended use of the public key",
							MarkdownDescription: "The intended use of the public key",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"enc",
									"sig",
								),
							},
						},
						"x": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The x coordinate for the Elliptic Curve point",
							MarkdownDescription: "The x coordinate for the Elliptic Curve point",
						},
						"y": schema.StringAttribute{
							CustomType:          ovhtypes.TfStringType{},
							Optional:            true,
							Computed:            true,
							Description:         "The y coordinate for the Elliptic Curve point",
							MarkdownDescription: "The y coordinate for the Elliptic Curve point",
						},
					},
					CustomType: KeysType{
						ObjectType: types.ObjectType{
							AttrTypes: KeysValue{}.AttributeTypes(ctx),
						},
					},
				},
				CustomType:          ovhtypes.NewTfListNestedType[KeysValue](ctx),
				Optional:            true,
				Computed:            true,
				Description:         "Set of JSON Web Keys in case of import, incompatible with type,size,curve and operations",
				MarkdownDescription: "Set of JSON Web Keys in case of import, incompatible with type,size,curve and operations",
			},
		*/
		"name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Key name",
			MarkdownDescription: "Key name",
		},
		"okms_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Okms ID",
			MarkdownDescription: "Okms ID",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"operations": schema.ListAttribute{
			CustomType:          ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
			Required:            true,
			Description:         "The operations for which the key is intended to be used",
			MarkdownDescription: "The operations for which the key is intended to be used",
			PlanModifiers: []planmodifier.List{
				listplanmodifier.RequiresReplace(),
			},
		},
		"size": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Optional:            true,
			Computed:            true,
			Description:         "Size of the key to be created",
			MarkdownDescription: "Size of the key to be created",
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.RequiresReplaceIfConfigured(),
			},
			Validators: []validator.Int64{
				int64validator.OneOf(
					128,
					192,
					256,
					2048,
					3072,
					4096,
				),
			},
		},
		"state": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "State of the key",
			MarkdownDescription: "State of the key",
			Validators: []validator.String{
				stringvalidator.OneOf(
					"ACTIVE",
					"COMPROMISED",
					"DEACTIVATED",
				),
			},
		},
		"type": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Type of the key to be created",
			MarkdownDescription: "Type of the key to be created",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Validators: []validator.String{
				stringvalidator.OneOf(
					"EC",
					"RSA",
					"oct",
				),
			},
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type OkmsServiceKeyResourceModel struct {
	Context            ovhtypes.TfStringValue                             `tfsdk:"context" json:"context"`
	CreatedAt          ovhtypes.TfStringValue                             `tfsdk:"created_at" json:"createdAt"`
	Curve              ovhtypes.TfStringValue                             `tfsdk:"curve" json:"curve"`
	DeactivationReason ovhtypes.TfStringValue                             `tfsdk:"deactivation_reason" json:"deactivationReason"`
	Id                 ovhtypes.TfStringValue                             `tfsdk:"id" json:"id"`
	Name               ovhtypes.TfStringValue                             `tfsdk:"name" json:"name"`
	OkmsId             ovhtypes.TfStringValue                             `tfsdk:"okms_id" json:"okmsId"`
	Operations         ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"operations" json:"operations"`
	Size               ovhtypes.TfInt64Value                              `tfsdk:"size" json:"size"`
	State              ovhtypes.TfStringValue                             `tfsdk:"state" json:"state"`
	Type               ovhtypes.TfStringValue                             `tfsdk:"type" json:"type"`
}

func (v *OkmsServiceKeyResourceModel) MergeWith(other *OkmsServiceKeyResourceModel) {

	if (v.Context.IsUnknown() || v.Context.IsNull()) && !other.Context.IsUnknown() {
		v.Context = other.Context
	}

	if (v.CreatedAt.IsUnknown() || v.CreatedAt.IsNull()) && !other.CreatedAt.IsUnknown() {
		v.CreatedAt = other.CreatedAt
	}

	if (v.Curve.IsUnknown() || v.Curve.IsNull()) && !other.Curve.IsUnknown() {
		v.Curve = other.Curve
	}

	if (v.DeactivationReason.IsUnknown() || v.DeactivationReason.IsNull()) && !other.DeactivationReason.IsUnknown() {
		v.DeactivationReason = other.DeactivationReason
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.OkmsId.IsUnknown() || v.OkmsId.IsNull()) && !other.OkmsId.IsUnknown() {
		v.OkmsId = other.OkmsId
	}

	if (v.Operations.IsUnknown() || v.Operations.IsNull()) && !other.Operations.IsUnknown() {
		v.Operations = other.Operations
	}

	if (v.Size.IsUnknown() || v.Size.IsNull()) && !other.Size.IsUnknown() {
		v.Size = other.Size
	}

	if (v.State.IsUnknown() || v.State.IsNull()) && !other.State.IsUnknown() {
		v.State = other.State
	}

	if (v.Type.IsUnknown() || v.Type.IsNull()) && !other.Type.IsUnknown() {
		v.Type = other.Type
	}
}

type OkmsServiceKeyWritableModel struct {
	Context            *ovhtypes.TfStringValue                             `tfsdk:"context" json:"context,omitempty"`
	Curve              *ovhtypes.TfStringValue                             `tfsdk:"curve" json:"curve,omitempty"`
	DeactivationReason *ovhtypes.TfStringValue                             `tfsdk:"deactivation_reason" json:"deactivationReason,omitempty"`
	Name               *ovhtypes.TfStringValue                             `tfsdk:"name" json:"name,omitempty"`
	Operations         *ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"operations" json:"operations,omitempty"`
	Size               *ovhtypes.TfInt64Value                              `tfsdk:"size" json:"size,omitempty"`
	State              *ovhtypes.TfStringValue                             `tfsdk:"state" json:"state,omitempty"`
	Type               *ovhtypes.TfStringValue                             `tfsdk:"type" json:"type,omitempty"`
}

func (v OkmsServiceKeyResourceModel) ToCreate() *OkmsServiceKeyWritableModel {
	res := &OkmsServiceKeyWritableModel{}

	if !v.Context.IsUnknown() {
		res.Context = &v.Context
	}

	if !v.Curve.IsUnknown() {
		res.Curve = &v.Curve
	}

	if !v.Name.IsUnknown() {
		res.Name = &v.Name
	}

	if !v.Operations.IsUnknown() {
		res.Operations = &v.Operations
	}

	if !v.Size.IsUnknown() {
		res.Size = &v.Size
	}

	if !v.Type.IsUnknown() {
		res.Type = &v.Type
	}

	return res
}

func (v OkmsServiceKeyResourceModel) ToUpdate() *OkmsServiceKeyWritableModel {
	res := &OkmsServiceKeyWritableModel{}

	if !v.DeactivationReason.IsUnknown() {
		res.DeactivationReason = &v.DeactivationReason
	}

	if !v.Name.IsUnknown() {
		res.Name = &v.Name
	}

	if !v.State.IsUnknown() {
		res.State = &v.State
	}

	return res
}
