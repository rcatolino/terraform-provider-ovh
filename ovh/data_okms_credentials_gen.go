package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OkmsCredentialsDataSourceSchema(ctx context.Context) schema.Schema {
	credAttrs := okmsCredentialAttributes(ctx)
	credAttrs["id"] = schema.StringAttribute{
		CustomType:          ovhtypes.TfStringType{},
		Computed:            true,
		Description:         "ID of the credential",
		MarkdownDescription: "ID of the credential",
	}

	attrs := map[string]schema.Attribute{
		"okms_credentials": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: credAttrs,
				CustomType: OkmsCredentialsType{
					ObjectType: types.ObjectType{
						AttrTypes: OkmsCredentialsValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType: ovhtypes.NewTfListNestedType[OkmsCredentialsValue](ctx),
			Computed:   true,
		},
		"okms_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Okms ID",
			MarkdownDescription: "Okms ID",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type OkmsCredentialsModel struct {
	OkmsCredentials ovhtypes.TfListNestedValue[OkmsCredentialsValue] `tfsdk:"okms_credentials" json:"okmsCredentials"`
	OkmsId          ovhtypes.TfStringValue                           `tfsdk:"okms_id" json:"okmsId"`
}

func (v *OkmsCredentialsModel) MergeWith(other *OkmsCredentialsModel) {

	if (v.OkmsCredentials.IsUnknown() || v.OkmsCredentials.IsNull()) && !other.OkmsCredentials.IsUnknown() {
		v.OkmsCredentials = other.OkmsCredentials
	}

	if (v.OkmsId.IsUnknown() || v.OkmsId.IsNull()) && !other.OkmsId.IsUnknown() {
		v.OkmsId = other.OkmsId
	}

}

var _ basetypes.ObjectTypable = OkmsCredentialsType{}

type OkmsCredentialsType struct {
	basetypes.ObjectType
}

func (t OkmsCredentialsType) Equal(o attr.Type) bool {
	other, ok := o.(OkmsCredentialsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t OkmsCredentialsType) String() string {
	return "OkmsCredentialsType"
}

func (t OkmsCredentialsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	certificatePemAttribute, ok := attributes["certificate_pem"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`certificate_pem is missing from object`)

		return nil, diags
	}

	certificatePemVal, ok := certificatePemAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`certificate_pem expected to be ovhtypes.TfStringValue, was: %T`, certificatePemAttribute))
	}

	createdAtAttribute, ok := attributes["created_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_at is missing from object`)

		return nil, diags
	}

	createdAtVal, ok := createdAtAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_at expected to be ovhtypes.TfStringValue, was: %T`, createdAtAttribute))
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return nil, diags
	}

	descriptionVal, ok := descriptionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be ovhtypes.TfStringValue, was: %T`, descriptionAttribute))
	}

	expiredAtAttribute, ok := attributes["expired_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`expired_at is missing from object`)

		return nil, diags
	}

	expiredAtVal, ok := expiredAtAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`expired_at expected to be ovhtypes.TfStringValue, was: %T`, expiredAtAttribute))
	}

	fromCsrAttribute, ok := attributes["from_csr"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`from_csr is missing from object`)

		return nil, diags
	}

	fromCsrVal, ok := fromCsrAttribute.(ovhtypes.TfBoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`from_csr expected to be ovhtypes.TfBoolValue, was: %T`, fromCsrAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	identityUrnsAttribute, ok := attributes["identity_urns"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`identity_urns is missing from object`)

		return nil, diags
	}

	identityUrnsVal, ok := identityUrnsAttribute.(ovhtypes.TfListNestedValue[ovhtypes.TfStringValue])

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`identity_urns expected to be ovhtypes.TfListNestedValue[ovhtypes.TfStringValue], was: %T`, identityUrnsAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return nil, diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return OkmsCredentialsValue{
		CertificatePem: certificatePemVal,
		CreatedAt:      createdAtVal,
		Description:    descriptionVal,
		ExpiredAt:      expiredAtVal,
		FromCsr:        fromCsrVal,
		Id:             idVal,
		IdentityUrns:   identityUrnsVal,
		Name:           nameVal,
		Status:         statusVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewOkmsCredentialsValueNull() OkmsCredentialsValue {
	return OkmsCredentialsValue{
		state: attr.ValueStateNull,
	}
}

func NewOkmsCredentialsValueUnknown() OkmsCredentialsValue {
	return OkmsCredentialsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewOkmsCredentialsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (OkmsCredentialsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing OkmsCredentialsValue Attribute Value",
				"While creating a OkmsCredentialsValue value, a missing attribute value was detected. "+
					"A OkmsCredentialsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OkmsCredentialsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid OkmsCredentialsValue Attribute Type",
				"While creating a OkmsCredentialsValue value, an invalid attribute value was detected. "+
					"A OkmsCredentialsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OkmsCredentialsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("OkmsCredentialsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra OkmsCredentialsValue Attribute Value",
				"While creating a OkmsCredentialsValue value, an extra attribute value was detected. "+
					"A OkmsCredentialsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra OkmsCredentialsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewOkmsCredentialsValueUnknown(), diags
	}

	certificatePemAttribute, ok := attributes["certificate_pem"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`certificate_pem is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	certificatePemVal, ok := certificatePemAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`certificate_pem expected to be ovhtypes.TfStringValue, was: %T`, certificatePemAttribute))
	}

	createdAtAttribute, ok := attributes["created_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_at is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	createdAtVal, ok := createdAtAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_at expected to be ovhtypes.TfStringValue, was: %T`, createdAtAttribute))
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	descriptionVal, ok := descriptionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be ovhtypes.TfStringValue, was: %T`, descriptionAttribute))
	}

	expiredAtAttribute, ok := attributes["expired_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`expired_at is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	expiredAtVal, ok := expiredAtAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`expired_at expected to be ovhtypes.TfStringValue, was: %T`, expiredAtAttribute))
	}

	fromCsrAttribute, ok := attributes["from_csr"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`from_csr is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	fromCsrVal, ok := fromCsrAttribute.(ovhtypes.TfBoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`from_csr expected to be ovhtypes.TfBoolValue, was: %T`, fromCsrAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	identityUrnsAttribute, ok := attributes["identity_urns"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`identity_urns is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	identityUrnsVal, ok := identityUrnsAttribute.(ovhtypes.TfListNestedValue[ovhtypes.TfStringValue])

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`identity_urns expected to be ovhtypes.TfListNestedValue[ovhtypes.TfStringValue], was: %T`, identityUrnsAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewOkmsCredentialsValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return NewOkmsCredentialsValueUnknown(), diags
	}

	return OkmsCredentialsValue{
		CertificatePem: certificatePemVal,
		CreatedAt:      createdAtVal,
		Description:    descriptionVal,
		ExpiredAt:      expiredAtVal,
		FromCsr:        fromCsrVal,
		Id:             idVal,
		IdentityUrns:   identityUrnsVal,
		Name:           nameVal,
		Status:         statusVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewOkmsCredentialsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) OkmsCredentialsValue {
	object, diags := NewOkmsCredentialsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewOkmsCredentialsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t OkmsCredentialsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewOkmsCredentialsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewOkmsCredentialsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewOkmsCredentialsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewOkmsCredentialsValueMust(OkmsCredentialsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t OkmsCredentialsType) ValueType(ctx context.Context) attr.Value {
	return OkmsCredentialsValue{}
}

var _ basetypes.ObjectValuable = OkmsCredentialsValue{}

type OkmsCredentialsValue struct {
	CertificatePem ovhtypes.TfStringValue                             `tfsdk:"certificate_pem" json:"certificatePem"`
	CreatedAt      ovhtypes.TfStringValue                             `tfsdk:"created_at" json:"createdAt"`
	Description    ovhtypes.TfStringValue                             `tfsdk:"description" json:"description"`
	ExpiredAt      ovhtypes.TfStringValue                             `tfsdk:"expired_at" json:"expiredAt"`
	FromCsr        ovhtypes.TfBoolValue                               `tfsdk:"from_csr" json:"fromCsr"`
	Id             ovhtypes.TfStringValue                             `tfsdk:"id" json:"id"`
	IdentityUrns   ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"identity_urns" json:"identityURNs"`
	Name           ovhtypes.TfStringValue                             `tfsdk:"name" json:"name"`
	Status         ovhtypes.TfStringValue                             `tfsdk:"status" json:"status"`
	state          attr.ValueState
}

func (v *OkmsCredentialsValue) UnmarshalJSON(data []byte) error {
	type JsonOkmsCredentialsValue OkmsCredentialsValue

	var tmp JsonOkmsCredentialsValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.CertificatePem = tmp.CertificatePem
	v.CreatedAt = tmp.CreatedAt
	v.Description = tmp.Description
	v.ExpiredAt = tmp.ExpiredAt
	v.FromCsr = tmp.FromCsr
	v.Id = tmp.Id
	v.IdentityUrns = tmp.IdentityUrns
	v.Name = tmp.Name
	v.Status = tmp.Status

	v.state = attr.ValueStateKnown

	return nil
}

func (v *OkmsCredentialsValue) MergeWith(other *OkmsCredentialsValue) {

	if (v.CertificatePem.IsUnknown() || v.CertificatePem.IsNull()) && !other.CertificatePem.IsUnknown() {
		v.CertificatePem = other.CertificatePem
	}

	if (v.CreatedAt.IsUnknown() || v.CreatedAt.IsNull()) && !other.CreatedAt.IsUnknown() {
		v.CreatedAt = other.CreatedAt
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

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v OkmsCredentialsValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"certificatePem": v.CertificatePem,
		"createdAt":      v.CreatedAt,
		"description":    v.Description,
		"expiredAt":      v.ExpiredAt,
		"fromCsr":        v.FromCsr,
		"id":             v.Id,
		"identityUrns":   v.IdentityUrns,
		"name":           v.Name,
		"status":         v.Status,
	}
}
func (v OkmsCredentialsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 9)

	var val tftypes.Value
	var err error

	attrTypes["certificate_pem"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["created_at"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["description"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["expired_at"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["from_csr"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["identity_urns"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 9)

		val, err = v.CertificatePem.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["certificate_pem"] = val

		val, err = v.CreatedAt.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["created_at"] = val

		val, err = v.Description.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["description"] = val

		val, err = v.ExpiredAt.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["expired_at"] = val

		val, err = v.FromCsr.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["from_csr"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.IdentityUrns.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["identity_urns"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Status.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["status"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v OkmsCredentialsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v OkmsCredentialsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v OkmsCredentialsValue) String() string {
	return "OkmsCredentialsValue"
}

func (v OkmsCredentialsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"certificate_pem": ovhtypes.TfStringType{},
			"created_at":      ovhtypes.TfStringType{},
			"description":     ovhtypes.TfStringType{},
			"expired_at":      ovhtypes.TfStringType{},
			"from_csr":        ovhtypes.TfBoolType{},
			"id":              ovhtypes.TfStringType{},
			"identity_urns":   ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
			"name":            ovhtypes.TfStringType{},
			"status":          ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"certificate_pem": v.CertificatePem,
			"created_at":      v.CreatedAt,
			"description":     v.Description,
			"expired_at":      v.ExpiredAt,
			"from_csr":        v.FromCsr,
			"id":              v.Id,
			"identity_urns":   v.IdentityUrns,
			"name":            v.Name,
			"status":          v.Status,
		})

	return objVal, diags
}

func (v OkmsCredentialsValue) Equal(o attr.Value) bool {
	other, ok := o.(OkmsCredentialsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.CertificatePem.Equal(other.CertificatePem) {
		return false
	}

	if !v.CreatedAt.Equal(other.CreatedAt) {
		return false
	}

	if !v.Description.Equal(other.Description) {
		return false
	}

	if !v.ExpiredAt.Equal(other.ExpiredAt) {
		return false
	}

	if !v.FromCsr.Equal(other.FromCsr) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.IdentityUrns.Equal(other.IdentityUrns) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	return true
}

func (v OkmsCredentialsValue) Type(ctx context.Context) attr.Type {
	return OkmsCredentialsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v OkmsCredentialsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"certificate_pem": ovhtypes.TfStringType{},
		"created_at":      ovhtypes.TfStringType{},
		"description":     ovhtypes.TfStringType{},
		"expired_at":      ovhtypes.TfStringType{},
		"from_csr":        ovhtypes.TfBoolType{},
		"id":              ovhtypes.TfStringType{},
		"identity_urns":   ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
		"name":            ovhtypes.TfStringType{},
		"status":          ovhtypes.TfStringType{},
	}
}
