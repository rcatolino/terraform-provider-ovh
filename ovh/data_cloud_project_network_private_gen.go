// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudProjectNetworkPrivateDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
		"network_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Network ID",
			MarkdownDescription: "Network ID",
		},
		"name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Network name",
			MarkdownDescription: "Network name",
		},
		"regions": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"openstack_id": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network id on openstack region",
						MarkdownDescription: "Network id on openstack region",
					},
					"region": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network region",
						MarkdownDescription: "Network region",
					},
					"status": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network region status",
						MarkdownDescription: "Network region status",
					},
				},
				CustomType: RegionsType{
					ObjectType: types.ObjectType{
						AttrTypes: RegionsValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType:          ovhtypes.NewTfListNestedType[RegionsValue](ctx),
			Computed:            true,
			Description:         "Details about private network in region",
			MarkdownDescription: "Details about private network in region",
		},
		"status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Network status",
			MarkdownDescription: "Network status",
		},
		"type": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Network type",
			MarkdownDescription: "Network type",
		},
		"vlan_id": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Computed:            true,
			Description:         "Network VLAN id",
			MarkdownDescription: "Network VLAN id",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type CloudProjectNetworkPrivateModel struct {
	Name        ovhtypes.TfStringValue                   `tfsdk:"name" json:"name"`
	NetworkId   ovhtypes.TfStringValue                   `tfsdk:"network_id" json:"id"`
	Regions     ovhtypes.TfListNestedValue[RegionsValue] `tfsdk:"regions" json:"regions"`
	ServiceName ovhtypes.TfStringValue                   `tfsdk:"service_name" json:"serviceName"`
	Status      ovhtypes.TfStringValue                   `tfsdk:"status" json:"status"`
	Type        ovhtypes.TfStringValue                   `tfsdk:"type" json:"type"`
	VlanId      ovhtypes.TfInt64Value                    `tfsdk:"vlan_id" json:"vlanId"`
}

func (v *CloudProjectNetworkPrivateModel) MergeWith(other *CloudProjectNetworkPrivateModel) {

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.NetworkId.IsUnknown() || v.NetworkId.IsNull()) && !other.NetworkId.IsUnknown() {
		v.NetworkId = other.NetworkId
	}

	if (v.Regions.IsUnknown() || v.Regions.IsNull()) && !other.Regions.IsUnknown() {
		v.Regions = other.Regions
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.Type.IsUnknown() || v.Type.IsNull()) && !other.Type.IsUnknown() {
		v.Type = other.Type
	}

	if (v.VlanId.IsUnknown() || v.VlanId.IsNull()) && !other.VlanId.IsUnknown() {
		v.VlanId = other.VlanId
	}

}

var _ basetypes.ObjectTypable = RegionsType{}

type RegionsType struct {
	basetypes.ObjectType
}

func (t RegionsType) Equal(o attr.Type) bool {
	other, ok := o.(RegionsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t RegionsType) String() string {
	return "RegionsType"
}

func (t RegionsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	openstackIdAttribute, ok := attributes["openstack_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`openstack_id is missing from object`)

		return nil, diags
	}

	openstackIdVal, ok := openstackIdAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`openstack_id expected to be ovhtypes.TfStringValue, was: %T`, openstackIdAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return nil, diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
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

	return RegionsValue{
		OpenstackId: openstackIdVal,
		Region:      regionVal,
		Status:      statusVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewRegionsValueNull() RegionsValue {
	return RegionsValue{
		state: attr.ValueStateNull,
	}
}

func NewRegionsValueUnknown() RegionsValue {
	return RegionsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewRegionsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (RegionsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing RegionsValue Attribute Value",
				"While creating a RegionsValue value, a missing attribute value was detected. "+
					"A RegionsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("RegionsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid RegionsValue Attribute Type",
				"While creating a RegionsValue value, an invalid attribute value was detected. "+
					"A RegionsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("RegionsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("RegionsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra RegionsValue Attribute Value",
				"While creating a RegionsValue value, an extra attribute value was detected. "+
					"A RegionsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra RegionsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewRegionsValueUnknown(), diags
	}

	openstackIdAttribute, ok := attributes["openstack_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`openstack_id is missing from object`)

		return NewRegionsValueUnknown(), diags
	}

	openstackIdVal, ok := openstackIdAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`openstack_id expected to be ovhtypes.TfStringValue, was: %T`, openstackIdAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return NewRegionsValueUnknown(), diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewRegionsValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return NewRegionsValueUnknown(), diags
	}

	return RegionsValue{
		OpenstackId: openstackIdVal,
		Region:      regionVal,
		Status:      statusVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewRegionsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) RegionsValue {
	object, diags := NewRegionsValue(attributeTypes, attributes)

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

		panic("NewRegionsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t RegionsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewRegionsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewRegionsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewRegionsValueNull(), nil
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

	return NewRegionsValueMust(RegionsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t RegionsType) ValueType(ctx context.Context) attr.Value {
	return RegionsValue{}
}

var _ basetypes.ObjectValuable = RegionsValue{}

type RegionsValue struct {
	OpenstackId ovhtypes.TfStringValue `tfsdk:"openstack_id" json:"openstackId"`
	Region      ovhtypes.TfStringValue `tfsdk:"region" json:"region"`
	Status      ovhtypes.TfStringValue `tfsdk:"status" json:"status"`
	state       attr.ValueState
}

func (v *RegionsValue) UnmarshalJSON(data []byte) error {
	type JsonRegionsValue RegionsValue

	var tmp JsonRegionsValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.OpenstackId = tmp.OpenstackId
	v.Region = tmp.Region
	v.Status = tmp.Status

	v.state = attr.ValueStateKnown

	return nil
}

func (v *RegionsValue) MergeWith(other *RegionsValue) {

	if (v.OpenstackId.IsUnknown() || v.OpenstackId.IsNull()) && !other.OpenstackId.IsUnknown() {
		v.OpenstackId = other.OpenstackId
	}

	if (v.Region.IsUnknown() || v.Region.IsNull()) && !other.Region.IsUnknown() {
		v.Region = other.Region
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v RegionsValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"openstackId": v.OpenstackId,
		"region":      v.Region,
		"status":      v.Status,
	}
}
func (v RegionsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["openstack_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["region"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.OpenstackId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["openstack_id"] = val

		val, err = v.Region.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["region"] = val

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

func (v RegionsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v RegionsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v RegionsValue) String() string {
	return "RegionsValue"
}

func (v RegionsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"openstack_id": ovhtypes.TfStringType{},
			"region":       ovhtypes.TfStringType{},
			"status":       ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"openstack_id": v.OpenstackId,
			"region":       v.Region,
			"status":       v.Status,
		})

	return objVal, diags
}

func (v RegionsValue) Equal(o attr.Value) bool {
	other, ok := o.(RegionsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.OpenstackId.Equal(other.OpenstackId) {
		return false
	}

	if !v.Region.Equal(other.Region) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	return true
}

func (v RegionsValue) Type(ctx context.Context) attr.Type {
	return RegionsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v RegionsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"openstack_id": ovhtypes.TfStringType{},
		"region":       ovhtypes.TfStringType{},
		"status":       ovhtypes.TfStringType{},
	}
}