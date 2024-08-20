package ovh

import (
	"context"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OkmsCredentialResourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"certificate_pem": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Certificate PEM of the credential",
			MarkdownDescription: "Certificate PEM of the credential",
		},
		"created_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Creation time of the credential",
			MarkdownDescription: "Creation time of the credential",
		},
		/*
			"credential_id": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "ID of the credential",
				MarkdownDescription: "ID of the credential",
			},
		*/
		"csr": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Optional:            true,
			Computed:            true,
			Description:         "Valid Certificate Signing Request",
			MarkdownDescription: "Valid Certificate Signing Request",
		},
		"description": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Optional:            true,
			Computed:            true,
			Description:         "Description of the credential (max 200)",
			MarkdownDescription: "Description of the credential (max 200)",
		},
		"expired_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Expiration time of the credential",
			MarkdownDescription: "Expiration time of the credential",
		},
		"from_csr": schema.BoolAttribute{
			CustomType:          ovhtypes.TfBoolType{},
			Computed:            true,
			Description:         "Is the credential generated from CSR",
			MarkdownDescription: "Is the credential generated from CSR",
		},
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "ID of the credential",
			MarkdownDescription: "ID of the credential",
		},
		"identity_urns": schema.ListAttribute{
			CustomType:          ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
			Required:            true,
			Description:         "List of identity URNs associated with the credential (max 25)",
			MarkdownDescription: "List of identity URNs associated with the credential (max 25)",
		},
		"name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Name of the credential (max 50)",
			MarkdownDescription: "Name of the credential (max 50)",
		},
		"okms_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Okms ID",
			MarkdownDescription: "Okms ID",
		},
		"private_key_pem": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Sensitive:           true,
			Description:         "Private Key PEM of the credential if no CSR is provided (cannot be retrieve later)",
			MarkdownDescription: "Private Key PEM of the credential if no CSR is provided (cannot be retrieve later)",
		},
		"status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Status of the credential",
			MarkdownDescription: "Status of the credential",
		},
		"validity": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Optional:            true,
			Computed:            true,
			Description:         "Validity in days (default 365, max 365)",
			MarkdownDescription: "Validity in days (default 365, max 365)",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type OkmsCredentialWritableModel struct {
	Csr          *ovhtypes.TfStringValue                             `tfsdk:"csr" json:"csr,omitempty"`
	Description  *ovhtypes.TfStringValue                             `tfsdk:"description" json:"description,omitempty"`
	IdentityUrns *ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"identity_urns" json:"identityURNs,omitempty"`
	Name         *ovhtypes.TfStringValue                             `tfsdk:"name" json:"name,omitempty"`
	Validity     *ovhtypes.TfInt64Value                              `tfsdk:"validity" json:"validity,omitempty"`
}

func (v OkmsCredentialModel) ToCreate() *OkmsCredentialWritableModel {
	res := &OkmsCredentialWritableModel{}

	if !v.Csr.IsUnknown() {
		res.Csr = &v.Csr
	}

	if !v.Description.IsUnknown() {
		res.Description = &v.Description
	}

	if !v.IdentityUrns.IsUnknown() {
		res.IdentityUrns = &v.IdentityUrns
	}

	if !v.Name.IsUnknown() {
		res.Name = &v.Name
	}

	if !v.Validity.IsUnknown() {
		res.Validity = &v.Validity
	}

	return res
}

func (v OkmsCredentialModel) ToUpdate() *OkmsCredentialWritableModel {
	res := &OkmsCredentialWritableModel{}

	if !v.Csr.IsUnknown() {
		res.Csr = &v.Csr
	}

	if !v.Description.IsUnknown() {
		res.Description = &v.Description
	}

	if !v.IdentityUrns.IsUnknown() {
		res.IdentityUrns = &v.IdentityUrns
	}

	if !v.Name.IsUnknown() {
		res.Name = &v.Name
	}

	if !v.Validity.IsUnknown() {
		res.Validity = &v.Validity
	}

	return res
}
