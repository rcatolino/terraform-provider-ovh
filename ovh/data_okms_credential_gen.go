// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OkmsCredentialAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
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
		"description": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Description of the credential",
			MarkdownDescription: "Description of the credential",
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
			Computed:            true,
			Description:         "List of identity URNs associated with the credential",
			MarkdownDescription: "List of identity URNs associated with the credential",
		},
		"name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Name of the credential",
			MarkdownDescription: "Name of the credential",
		},
		"status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Status of the credential",
			MarkdownDescription: "Status of the credential",
		},
	}
}

func OkmsCredentialDataSourceSchema(ctx context.Context) schema.Schema {
	credAttrs := OkmsCredentialAttributes(ctx)
	credAttrs["credential_id"] = schema.StringAttribute{
		CustomType:          ovhtypes.TfStringType{},
		Required:            true,
		Description:         "Credential ID",
		MarkdownDescription: "Credential ID",
	}
	credAttrs["okms_id"] = schema.StringAttribute{
		CustomType:          ovhtypes.TfStringType{},
		Required:            true,
		Description:         "Okms ID",
		MarkdownDescription: "Okms ID",
	}

	return schema.Schema{
		Attributes: credAttrs,
	}
}

type OkmsCredentialModel struct {
	CertificatePem ovhtypes.TfStringValue                             `tfsdk:"certificate_pem" json:"certificatePem"`
	CreatedAt      ovhtypes.TfStringValue                             `tfsdk:"created_at" json:"createdAt"`
	CredentialId   ovhtypes.TfStringValue                             `tfsdk:"credential_id" json:"credentialId"`
	Description    ovhtypes.TfStringValue                             `tfsdk:"description" json:"description"`
	ExpiredAt      ovhtypes.TfStringValue                             `tfsdk:"expired_at" json:"expiredAt"`
	FromCsr        ovhtypes.TfBoolValue                               `tfsdk:"from_csr" json:"fromCsr"`
	Id             ovhtypes.TfStringValue                             `tfsdk:"id" json:"id"`
	IdentityUrns   ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"identity_urns" json:"identityUrns"`
	Name           ovhtypes.TfStringValue                             `tfsdk:"name" json:"name"`
	OkmsId         ovhtypes.TfStringValue                             `tfsdk:"okms_id" json:"okmsId"`
	Status         ovhtypes.TfStringValue                             `tfsdk:"status" json:"status"`
}

func (v *OkmsCredentialModel) MergeWith(other *OkmsCredentialModel) {

	if (v.CertificatePem.IsUnknown() || v.CertificatePem.IsNull()) && !other.CertificatePem.IsUnknown() {
		v.CertificatePem = other.CertificatePem
	}

	if (v.CreatedAt.IsUnknown() || v.CreatedAt.IsNull()) && !other.CreatedAt.IsUnknown() {
		v.CreatedAt = other.CreatedAt
	}

	if (v.CredentialId.IsUnknown() || v.CredentialId.IsNull()) && !other.CredentialId.IsUnknown() {
		v.CredentialId = other.CredentialId
	}

	if (v.Description.IsUnknown() || v.Description.IsNull()) && !other.Description.IsUnknown() {
		v.Description = other.Description
	}

	if (v.ExpiredAt.IsUnknown() || v.ExpiredAt.IsNull()) && !other.ExpiredAt.IsUnknown() {
		v.ExpiredAt = other.ExpiredAt
	}

	if (v.FromCsr.IsUnknown() || v.FromCsr.IsNull()) && !other.FromCsr.IsUnknown() {
		v.FromCsr = other.FromCsr
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.IdentityUrns.IsUnknown() || v.IdentityUrns.IsNull()) && !other.IdentityUrns.IsUnknown() {
		v.IdentityUrns = other.IdentityUrns
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.OkmsId.IsUnknown() || v.OkmsId.IsNull()) && !other.OkmsId.IsUnknown() {
		v.OkmsId = other.OkmsId
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

}
