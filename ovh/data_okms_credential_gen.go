package ovh

import (
	"context"
	// "encoding/json"

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
	/*
		credAttrs["credential_id"] = schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Credential ID",
			MarkdownDescription: "Credential ID",
		}
	*/
	credAttrs["id"] = schema.StringAttribute{
		CustomType:          ovhtypes.TfStringType{},
		Required:            true,
		Description:         "ID of the credential",
		MarkdownDescription: "ID of the credential",
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

// TODO: duplicate for resources
type OkmsCredentialModel struct {
	CertificatePem ovhtypes.TfStringValue `tfsdk:"certificate_pem" json:"certificatePem"`
	CreatedAt      ovhtypes.TfStringValue `tfsdk:"created_at" json:"createdAt"`
	//CredentialId   ovhtypes.TfStringValue                             `tfsdk:"credential_id" json:"credentialId"`
	Csr           ovhtypes.TfStringValue                             `tfsdk:"csr" json:"csr"`
	Description   ovhtypes.TfStringValue                             `tfsdk:"description" json:"description"`
	ExpiredAt     ovhtypes.TfStringValue                             `tfsdk:"expired_at" json:"expiredAt"`
	FromCsr       ovhtypes.TfBoolValue                               `tfsdk:"from_csr" json:"fromCsr"`
	Id            ovhtypes.TfStringValue                             `tfsdk:"id" json:"id"`
	IdentityUrns  ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"identity_urns" json:"identityURNs"`
	Name          ovhtypes.TfStringValue                             `tfsdk:"name" json:"name"`
	OkmsId        ovhtypes.TfStringValue                             `tfsdk:"okms_id" json:"okmsId"`
	PrivateKeyPem ovhtypes.TfStringValue                             `tfsdk:"private_key_pem" json:"privateKeyPem"`
	Status        ovhtypes.TfStringValue                             `tfsdk:"status" json:"status"`
	Validity      ovhtypes.TfInt64Value                              `tfsdk:"validity" json:"validity"`
}

/*
func (o *OkmsCredentialModel) UnmarshalJSON(data []byte) error {
	type JsonOkmsCredential OkmsCredentialModel

	var tmp JsonOkmsCredential
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	output := OkmsCredentialModel(tmp)
	o.MergeWith(&output)
	return nil
}
*/

func (o *OkmsCredentialModel) MergeWith(other *OkmsCredentialModel) {

	if (o.CertificatePem.IsUnknown() || o.CertificatePem.IsNull()) && !other.CertificatePem.IsUnknown() {
		o.CertificatePem = other.CertificatePem
	}

	if (o.CreatedAt.IsUnknown() || o.CreatedAt.IsNull()) && !other.CreatedAt.IsUnknown() {
		o.CreatedAt = other.CreatedAt
	}

	/*
		if (o.CredentialId.IsUnknown() || o.CredentialId.IsNull()) && !other.CredentialId.IsUnknown() {
			o.CredentialId = other.CredentialId
		}
	*/

	if (o.Csr.IsUnknown() || o.Csr.IsNull()) && !other.Csr.IsUnknown() {
		o.Csr = other.Csr
	}

	if (o.Description.IsUnknown() || o.Description.IsNull()) && !other.Description.IsUnknown() {
		o.Description = other.Description
	}

	if (o.ExpiredAt.IsUnknown() || o.ExpiredAt.IsNull()) && !other.ExpiredAt.IsUnknown() {
		o.ExpiredAt = other.ExpiredAt
	}

	if (o.FromCsr.IsUnknown() || o.FromCsr.IsNull()) && !other.FromCsr.IsUnknown() {
		o.FromCsr = other.FromCsr
	}

	if (o.Id.IsUnknown() || o.Id.IsNull()) && !other.Id.IsUnknown() {
		o.Id = other.Id
	}

	if (o.IdentityUrns.IsUnknown() || o.IdentityUrns.IsNull()) && !other.IdentityUrns.IsUnknown() {
		o.IdentityUrns = other.IdentityUrns
	}

	if (o.Name.IsUnknown() || o.Name.IsNull()) && !other.Name.IsUnknown() {
		o.Name = other.Name
	}

	if (o.OkmsId.IsUnknown() || o.OkmsId.IsNull()) && !other.OkmsId.IsUnknown() {
		o.OkmsId = other.OkmsId
	}

	if (o.PrivateKeyPem.IsUnknown() || o.PrivateKeyPem.IsNull()) && !other.PrivateKeyPem.IsUnknown() {
		o.PrivateKeyPem = other.PrivateKeyPem
	}

	if (o.Status.IsUnknown() || o.Status.IsNull()) && !other.Status.IsUnknown() {
		o.Status = other.Status
	}

	if (o.Validity.IsUnknown() || o.Validity.IsNull()) && !other.Validity.IsUnknown() {
		o.Validity = other.Validity
	}

}
