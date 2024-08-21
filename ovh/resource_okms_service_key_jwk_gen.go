package ovh

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OkmsServiceJwkResourceSchema(ctx context.Context) schema.Schema {
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
		"jwk": schema.SingleNestedAttribute{
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
			Required:            true,
			Description:         "JSON Web Keys in case of import, incompatible with type,size,curve and operations",
			MarkdownDescription: "Set of JSON Web Keys in case of import, incompatible with type,size,curve and operations",
		},
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
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type OkmsServiceJwkResourceModel struct {
	Context            ovhtypes.TfStringValue `tfsdk:"context" json:"context"`
	CreatedAt          ovhtypes.TfStringValue `tfsdk:"created_at" json:"createdAt"`
	DeactivationReason ovhtypes.TfStringValue `tfsdk:"deactivation_reason" json:"deactivationReason"`
	Id                 ovhtypes.TfStringValue `tfsdk:"id" json:"id"`
	Jwk                JwkModel               `tfsdk:"jwk" json:"keys"`
	Name               ovhtypes.TfStringValue `tfsdk:"name" json:"name"`
	OkmsId             ovhtypes.TfStringValue `tfsdk:"okms_id" json:"okmsId"`
	State              ovhtypes.TfStringValue `tfsdk:"state" json:"state"`
}

type JwkModel struct {
	Alg    ovhtypes.TfStringValue                             `tfsdk:"alg" json:"alg"`
	Crv    ovhtypes.TfStringValue                             `tfsdk:"crv" json:"crv"`
	E      ovhtypes.TfStringValue                             `tfsdk:"e" json:"e"`
	D      ovhtypes.TfStringValue                             `tfsdk:"d" json:"d"`
	Dp     ovhtypes.TfStringValue                             `tfsdk:"dp" json:"dp"`
	Dq     ovhtypes.TfStringValue                             `tfsdk:"dq" json:"dq"`
	KeyOps ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"key_ops" json:"key_ops"`
	K      ovhtypes.TfStringValue                             `tfsdk:"k" json:"k"`
	Kid    ovhtypes.TfStringValue                             `tfsdk:"kid" json:"kid"`
	Kty    ovhtypes.TfStringValue                             `tfsdk:"kty" json:"kty"`
	N      ovhtypes.TfStringValue                             `tfsdk:"n" json:"n"`
	P      ovhtypes.TfStringValue                             `tfsdk:"p" json:"p"`
	Q      ovhtypes.TfStringValue                             `tfsdk:"q" json:"q"`
	Qi     ovhtypes.TfStringValue                             `tfsdk:"qi" json:"qi"`
	Use    ovhtypes.TfStringValue                             `tfsdk:"use" json:"use"`
	X      ovhtypes.TfStringValue                             `tfsdk:"x" json:"x"`
	Y      ovhtypes.TfStringValue                             `tfsdk:"y" json:"y"`
}

func (v *JwkModel) UnmarshalJSON(data []byte) error {
	type JsonKeys JwkModel
	var tmp []JsonKeys
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	v.Alg = tmp[0].Alg
	v.Crv = tmp[0].Crv
	v.E = tmp[0].E
	v.D = tmp[0].D
	v.Dp = tmp[0].Dp
	// TODO: finish
	return nil
}

func (v *JwkModel) MergeWith(other *JwkModel) {
	if (v.Alg.IsUnknown() || v.Alg.IsNull()) && !other.Alg.IsUnknown() {
		v.Alg = other.Alg
	}

	if (v.Crv.IsUnknown() || v.Crv.IsNull()) && !other.Crv.IsUnknown() {
		v.Crv = other.Crv
	}

	if (v.E.IsUnknown() || v.E.IsNull()) && !other.E.IsUnknown() {
		v.E = other.E
	}

	if (v.D.IsUnknown() || v.D.IsNull()) && !other.D.IsUnknown() {
		v.D = other.D
	}

	if (v.Dp.IsUnknown() || v.Dp.IsNull()) && !other.Dp.IsUnknown() {
		v.Dp = other.Dp
	}

	if (v.Dq.IsUnknown() || v.Dq.IsNull()) && !other.Dq.IsUnknown() {
		v.Dq = other.Dq
	}

	if (v.KeyOps.IsUnknown() || v.KeyOps.IsNull()) && !other.KeyOps.IsUnknown() {
		v.KeyOps = other.KeyOps
	}

	if (v.K.IsUnknown() || v.K.IsNull()) && !other.K.IsUnknown() {
		v.K = other.K
	}

	if (v.Kid.IsUnknown() || v.Kid.IsNull()) && !other.Kid.IsUnknown() {
		v.Kid = other.Kid
	}

	if (v.Kty.IsUnknown() || v.Kty.IsNull()) && !other.Kty.IsUnknown() {
		v.Kty = other.Kty
	}

	if (v.N.IsUnknown() || v.N.IsNull()) && !other.N.IsUnknown() {
		v.N = other.N
	}

	if (v.P.IsUnknown() || v.P.IsNull()) && !other.P.IsUnknown() {
		v.P = other.P
	}

	if (v.Q.IsUnknown() || v.Q.IsNull()) && !other.Q.IsUnknown() {
		v.Q = other.Q
	}

	if (v.Qi.IsUnknown() || v.Qi.IsNull()) && !other.Qi.IsUnknown() {
		v.Qi = other.Qi
	}

	if (v.Use.IsUnknown() || v.Use.IsNull()) && !other.Use.IsUnknown() {
		v.Use = other.Use
	}

	if (v.X.IsUnknown() || v.X.IsNull()) && !other.X.IsUnknown() {
		v.X = other.X
	}

	if (v.Y.IsUnknown() || v.Y.IsNull()) && !other.Y.IsUnknown() {
		v.Y = other.Y
	}

}

func (v *OkmsServiceJwkResourceModel) MergeWith(other *OkmsServiceJwkResourceModel) {
	if (v.Context.IsUnknown() || v.Context.IsNull()) && !other.Context.IsUnknown() {
		v.Context = other.Context
	}

	if (v.CreatedAt.IsUnknown() || v.CreatedAt.IsNull()) && !other.CreatedAt.IsUnknown() {
		v.CreatedAt = other.CreatedAt
	}

	if (v.DeactivationReason.IsUnknown() || v.DeactivationReason.IsNull()) && !other.DeactivationReason.IsUnknown() {
		v.DeactivationReason = other.DeactivationReason
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	v.Jwk.MergeWith(&other.Jwk)

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.OkmsId.IsUnknown() || v.OkmsId.IsNull()) && !other.OkmsId.IsUnknown() {
		v.OkmsId = other.OkmsId
	}

	if (v.State.IsUnknown() || v.State.IsNull()) && !other.State.IsUnknown() {
		v.State = other.State
	}
}

type OkmsServiceJwkWritableModel struct {
	Context            *ovhtypes.TfStringValue `tfsdk:"context" json:"context,omitempty"`
	DeactivationReason *ovhtypes.TfStringValue `tfsdk:"deactivation_reason" json:"deactivationReason,omitempty"`
	Keys               []JwkModel              `json:"keys"`
	Name               *ovhtypes.TfStringValue `tfsdk:"name" json:"name,omitempty"`
	State              *ovhtypes.TfStringValue `tfsdk:"state" json:"state,omitempty"`
}

func (v OkmsServiceJwkResourceModel) ToCreate() *OkmsServiceJwkWritableModel {
	res := &OkmsServiceJwkWritableModel{}

	if !v.Context.IsUnknown() {
		res.Context = &v.Context
	}

	if !v.Name.IsUnknown() {
		res.Name = &v.Name
	}

	res.Keys = []JwkModel{v.Jwk}

	return res
}

func (v OkmsServiceJwkResourceModel) ToUpdate() *OkmsServiceJwkWritableModel {
	res := &OkmsServiceJwkWritableModel{}

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
