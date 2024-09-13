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

func CloudProjectNetworkPrivateSubnetsDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"subnets": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"cidr": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Subnet CIDR",
						MarkdownDescription: "Subnet CIDR",
					},
					"dhcp_enabled": schema.BoolAttribute{
						CustomType:          ovhtypes.TfBoolType{},
						Computed:            true,
						Description:         "Is DHCP enabled for the subnet",
						MarkdownDescription: "Is DHCP enabled for the subnet",
					},
					"gateway_ip": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Gateway IP in the subnet",
						MarkdownDescription: "Gateway IP in the subnet",
					},
					"id": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Subnet id",
						MarkdownDescription: "Subnet id",
					},
					"ip_pools": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dhcp": schema.BoolAttribute{
									CustomType:          ovhtypes.TfBoolType{},
									Computed:            true,
									Description:         "Enable DHCP",
									MarkdownDescription: "Enable DHCP",
								},
								"end": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "Last IP for this region (eg: 192.168.1.24)",
									MarkdownDescription: "Last IP for this region (eg: 192.168.1.24)",
								},
								"network": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "Global network with cidr (eg: 192.168.1.0/24)",
									MarkdownDescription: "Global network with cidr (eg: 192.168.1.0/24)",
								},
								"region": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "Region of the subnet",
									MarkdownDescription: "Region of the subnet",
								},
								"start": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "First IP for this region (eg: 192.168.1.12)",
									MarkdownDescription: "First IP for this region (eg: 192.168.1.12)",
								},
							},
							CustomType: CloudProjectNetworkPrivateSubnetsIpPoolsType{
								ObjectType: types.ObjectType{
									AttrTypes: CloudProjectNetworkPrivateSubnetsIpPoolsValue{}.AttributeTypes(ctx),
								},
							},
						},
						CustomType:          ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivateSubnetsIpPoolsValue](ctx),
						Computed:            true,
						Description:         "List of ip pools allocated in subnet",
						MarkdownDescription: "List of ip pools allocated in subnet",
					},
				},
				CustomType: CloudProjectNetworkPrivateSubnetsType{
					ObjectType: types.ObjectType{
						AttrTypes: CloudProjectNetworkPrivateSubnetsValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType: ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivateSubnetsValue](ctx),
			Computed:   true,
		},
		"network_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Network ID",
			MarkdownDescription: "Network ID",
		},
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type CloudProjectNetworkPrivateSubnetsModel struct {
	Subnets     ovhtypes.TfListNestedValue[CloudProjectNetworkPrivateSubnetsValue] `tfsdk:"subnets" json:"cloudProjectNetworkPrivateSubnets"`
	NetworkId   ovhtypes.TfStringValue                                             `tfsdk:"network_id" json:"networkId"`
	ServiceName ovhtypes.TfStringValue                                             `tfsdk:"service_name" json:"serviceName"`
}

func (v *CloudProjectNetworkPrivateSubnetsModel) MergeWith(other *CloudProjectNetworkPrivateSubnetsModel) {

	if (v.Subnets.IsUnknown() || v.Subnets.IsNull()) && !other.Subnets.IsUnknown() {
		v.Subnets = other.Subnets
	}

	if (v.NetworkId.IsUnknown() || v.NetworkId.IsNull()) && !other.NetworkId.IsUnknown() {
		v.NetworkId = other.NetworkId
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}

}

var _ basetypes.ObjectTypable = CloudProjectNetworkPrivateSubnetsType{}

type CloudProjectNetworkPrivateSubnetsType struct {
	basetypes.ObjectType
}

func (t CloudProjectNetworkPrivateSubnetsType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectNetworkPrivateSubnetsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectNetworkPrivateSubnetsType) String() string {
	return "CloudProjectNetworkPrivateSubnetsType"
}

func (t CloudProjectNetworkPrivateSubnetsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	cidrAttribute, ok := attributes["cidr"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cidr is missing from object`)

		return nil, diags
	}

	cidrVal, ok := cidrAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cidr expected to be ovhtypes.TfStringValue, was: %T`, cidrAttribute))
	}

	dhcpEnabledAttribute, ok := attributes["dhcp_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dhcp_enabled is missing from object`)

		return nil, diags
	}

	dhcpEnabledVal, ok := dhcpEnabledAttribute.(ovhtypes.TfBoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dhcp_enabled expected to be ovhtypes.TfBoolValue, was: %T`, dhcpEnabledAttribute))
	}

	gatewayIpAttribute, ok := attributes["gateway_ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`gateway_ip is missing from object`)

		return nil, diags
	}

	gatewayIpVal, ok := gatewayIpAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`gateway_ip expected to be ovhtypes.TfStringValue, was: %T`, gatewayIpAttribute))
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

	ipPoolsAttribute, ok := attributes["ip_pools"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip_pools is missing from object`)

		return nil, diags
	}

	ipPoolsVal, ok := ipPoolsAttribute.(ovhtypes.TfListNestedValue[CloudProjectNetworkPrivateSubnetsIpPoolsValue])

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip_pools expected to be ovhtypes.TfListNestedValue[CloudProjectNetworkPrivateSubnetsIpPoolsValue], was: %T`, ipPoolsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectNetworkPrivateSubnetsValue{
		Cidr:        cidrVal,
		DhcpEnabled: dhcpEnabledVal,
		GatewayIp:   gatewayIpVal,
		Id:          idVal,
		IpPools:     ipPoolsVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivateSubnetsValueNull() CloudProjectNetworkPrivateSubnetsValue {
	return CloudProjectNetworkPrivateSubnetsValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectNetworkPrivateSubnetsValueUnknown() CloudProjectNetworkPrivateSubnetsValue {
	return CloudProjectNetworkPrivateSubnetsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectNetworkPrivateSubnetsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectNetworkPrivateSubnetsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectNetworkPrivateSubnetsValue Attribute Value",
				"While creating a CloudProjectNetworkPrivateSubnetsValue value, a missing attribute value was detected. "+
					"A CloudProjectNetworkPrivateSubnetsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivateSubnetsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectNetworkPrivateSubnetsValue Attribute Type",
				"While creating a CloudProjectNetworkPrivateSubnetsValue value, an invalid attribute value was detected. "+
					"A CloudProjectNetworkPrivateSubnetsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivateSubnetsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectNetworkPrivateSubnetsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectNetworkPrivateSubnetsValue Attribute Value",
				"While creating a CloudProjectNetworkPrivateSubnetsValue value, an extra attribute value was detected. "+
					"A CloudProjectNetworkPrivateSubnetsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectNetworkPrivateSubnetsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	cidrAttribute, ok := attributes["cidr"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cidr is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	cidrVal, ok := cidrAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cidr expected to be ovhtypes.TfStringValue, was: %T`, cidrAttribute))
	}

	dhcpEnabledAttribute, ok := attributes["dhcp_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dhcp_enabled is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	dhcpEnabledVal, ok := dhcpEnabledAttribute.(ovhtypes.TfBoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dhcp_enabled expected to be ovhtypes.TfBoolValue, was: %T`, dhcpEnabledAttribute))
	}

	gatewayIpAttribute, ok := attributes["gateway_ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`gateway_ip is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	gatewayIpVal, ok := gatewayIpAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`gateway_ip expected to be ovhtypes.TfStringValue, was: %T`, gatewayIpAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	ipPoolsAttribute, ok := attributes["ip_pools"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip_pools is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	ipPoolsVal, ok := ipPoolsAttribute.(ovhtypes.TfListNestedValue[CloudProjectNetworkPrivateSubnetsIpPoolsValue])

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip_pools expected to be ovhtypes.TfListNestedValue[CloudProjectNetworkPrivateSubnetsIpPoolsValue], was: %T`, ipPoolsAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), diags
	}

	return CloudProjectNetworkPrivateSubnetsValue{
		Cidr:        cidrVal,
		DhcpEnabled: dhcpEnabledVal,
		GatewayIp:   gatewayIpVal,
		Id:          idVal,
		IpPools:     ipPoolsVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivateSubnetsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectNetworkPrivateSubnetsValue {
	object, diags := NewCloudProjectNetworkPrivateSubnetsValue(attributeTypes, attributes)

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

		panic("NewCloudProjectNetworkPrivateSubnetsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectNetworkPrivateSubnetsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectNetworkPrivateSubnetsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectNetworkPrivateSubnetsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectNetworkPrivateSubnetsValueNull(), nil
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

	return NewCloudProjectNetworkPrivateSubnetsValueMust(CloudProjectNetworkPrivateSubnetsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectNetworkPrivateSubnetsType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectNetworkPrivateSubnetsValue{}
}

var _ basetypes.ObjectValuable = CloudProjectNetworkPrivateSubnetsValue{}

type CloudProjectNetworkPrivateSubnetsValue struct {
	Cidr        ovhtypes.TfStringValue                                                    `tfsdk:"cidr" json:"cidr"`
	DhcpEnabled ovhtypes.TfBoolValue                                                      `tfsdk:"dhcp_enabled" json:"dhcpEnabled"`
	GatewayIp   ovhtypes.TfStringValue                                                    `tfsdk:"gateway_ip" json:"gatewayIp"`
	Id          ovhtypes.TfStringValue                                                    `tfsdk:"id" json:"id"`
	IpPools     ovhtypes.TfListNestedValue[CloudProjectNetworkPrivateSubnetsIpPoolsValue] `tfsdk:"ip_pools" json:"ipPools"`
	state       attr.ValueState
}

func (v *CloudProjectNetworkPrivateSubnetsValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectNetworkPrivateSubnetsValue CloudProjectNetworkPrivateSubnetsValue

	var tmp JsonCloudProjectNetworkPrivateSubnetsValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Cidr = tmp.Cidr
	v.DhcpEnabled = tmp.DhcpEnabled
	v.GatewayIp = tmp.GatewayIp
	v.Id = tmp.Id
	v.IpPools = tmp.IpPools

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectNetworkPrivateSubnetsValue) MergeWith(other *CloudProjectNetworkPrivateSubnetsValue) {

	if (v.Cidr.IsUnknown() || v.Cidr.IsNull()) && !other.Cidr.IsUnknown() {
		v.Cidr = other.Cidr
	}

	if (v.DhcpEnabled.IsUnknown() || v.DhcpEnabled.IsNull()) && !other.DhcpEnabled.IsUnknown() {
		v.DhcpEnabled = other.DhcpEnabled
	}

	if (v.GatewayIp.IsUnknown() || v.GatewayIp.IsNull()) && !other.GatewayIp.IsUnknown() {
		v.GatewayIp = other.GatewayIp
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.IpPools.IsUnknown() || v.IpPools.IsNull()) && !other.IpPools.IsUnknown() {
		v.IpPools = other.IpPools
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectNetworkPrivateSubnetsValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"cidr":        v.Cidr,
		"dhcpEnabled": v.DhcpEnabled,
		"gatewayIp":   v.GatewayIp,
		"id":          v.Id,
		"ipPools":     v.IpPools,
	}
}
func (v CloudProjectNetworkPrivateSubnetsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 5)

	var val tftypes.Value
	var err error

	attrTypes["cidr"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["dhcp_enabled"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["gateway_ip"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["ip_pools"] = basetypes.ListType{
		ElemType: CloudProjectNetworkPrivateSubnetsIpPoolsValue{}.Type(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 5)

		val, err = v.Cidr.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["cidr"] = val

		val, err = v.DhcpEnabled.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["dhcp_enabled"] = val

		val, err = v.GatewayIp.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["gateway_ip"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.IpPools.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ip_pools"] = val

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

func (v CloudProjectNetworkPrivateSubnetsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectNetworkPrivateSubnetsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectNetworkPrivateSubnetsValue) String() string {
	return "CloudProjectNetworkPrivateSubnetsValue"
}

func (v CloudProjectNetworkPrivateSubnetsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"cidr":         ovhtypes.TfStringType{},
			"dhcp_enabled": ovhtypes.TfBoolType{},
			"gateway_ip":   ovhtypes.TfStringType{},
			"id":           ovhtypes.TfStringType{},
			"ip_pools":     ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivateSubnetsIpPoolsValue](ctx),
		},
		map[string]attr.Value{
			"cidr":         v.Cidr,
			"dhcp_enabled": v.DhcpEnabled,
			"gateway_ip":   v.GatewayIp,
			"id":           v.Id,
			"ip_pools":     v.IpPools,
		})

	return objVal, diags
}

func (v CloudProjectNetworkPrivateSubnetsValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectNetworkPrivateSubnetsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Cidr.Equal(other.Cidr) {
		return false
	}

	if !v.DhcpEnabled.Equal(other.DhcpEnabled) {
		return false
	}

	if !v.GatewayIp.Equal(other.GatewayIp) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.IpPools.Equal(other.IpPools) {
		return false
	}

	return true
}

func (v CloudProjectNetworkPrivateSubnetsValue) Type(ctx context.Context) attr.Type {
	return CloudProjectNetworkPrivateSubnetsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectNetworkPrivateSubnetsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"cidr":         ovhtypes.TfStringType{},
		"dhcp_enabled": ovhtypes.TfBoolType{},
		"gateway_ip":   ovhtypes.TfStringType{},
		"id":           ovhtypes.TfStringType{},
		"ip_pools":     ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivateSubnetsIpPoolsValue](ctx),
	}
}

var _ basetypes.ObjectTypable = CloudProjectNetworkPrivateSubnetsIpPoolsType{}

type CloudProjectNetworkPrivateSubnetsIpPoolsType struct {
	basetypes.ObjectType
}

func (t CloudProjectNetworkPrivateSubnetsIpPoolsType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectNetworkPrivateSubnetsIpPoolsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectNetworkPrivateSubnetsIpPoolsType) String() string {
	return "CloudProjectNetworkPrivateSubnetsIpPoolsType"
}

func (t CloudProjectNetworkPrivateSubnetsIpPoolsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	dhcpAttribute, ok := attributes["dhcp"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dhcp is missing from object`)

		return nil, diags
	}

	dhcpVal, ok := dhcpAttribute.(ovhtypes.TfBoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dhcp expected to be ovhtypes.TfBoolValue, was: %T`, dhcpAttribute))
	}

	endAttribute, ok := attributes["end"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`end is missing from object`)

		return nil, diags
	}

	endVal, ok := endAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`end expected to be ovhtypes.TfStringValue, was: %T`, endAttribute))
	}

	networkAttribute, ok := attributes["network"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`network is missing from object`)

		return nil, diags
	}

	networkVal, ok := networkAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`network expected to be ovhtypes.TfStringValue, was: %T`, networkAttribute))
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

	startAttribute, ok := attributes["start"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`start is missing from object`)

		return nil, diags
	}

	startVal, ok := startAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`start expected to be ovhtypes.TfStringValue, was: %T`, startAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectNetworkPrivateSubnetsIpPoolsValue{
		Dhcp:    dhcpVal,
		End:     endVal,
		Network: networkVal,
		Region:  regionVal,
		Start:   startVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivateSubnetsIpPoolsValueNull() CloudProjectNetworkPrivateSubnetsIpPoolsValue {
	return CloudProjectNetworkPrivateSubnetsIpPoolsValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown() CloudProjectNetworkPrivateSubnetsIpPoolsValue {
	return CloudProjectNetworkPrivateSubnetsIpPoolsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectNetworkPrivateSubnetsIpPoolsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectNetworkPrivateSubnetsIpPoolsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Value",
				"While creating a CloudProjectNetworkPrivateSubnetsIpPoolsValue value, a missing attribute value was detected. "+
					"A CloudProjectNetworkPrivateSubnetsIpPoolsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Type",
				"While creating a CloudProjectNetworkPrivateSubnetsIpPoolsValue value, an invalid attribute value was detected. "+
					"A CloudProjectNetworkPrivateSubnetsIpPoolsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Value",
				"While creating a CloudProjectNetworkPrivateSubnetsIpPoolsValue value, an extra attribute value was detected. "+
					"A CloudProjectNetworkPrivateSubnetsIpPoolsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectNetworkPrivateSubnetsIpPoolsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	dhcpAttribute, ok := attributes["dhcp"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dhcp is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	dhcpVal, ok := dhcpAttribute.(ovhtypes.TfBoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dhcp expected to be ovhtypes.TfBoolValue, was: %T`, dhcpAttribute))
	}

	endAttribute, ok := attributes["end"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`end is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	endVal, ok := endAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`end expected to be ovhtypes.TfStringValue, was: %T`, endAttribute))
	}

	networkAttribute, ok := attributes["network"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`network is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	networkVal, ok := networkAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`network expected to be ovhtypes.TfStringValue, was: %T`, networkAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
	}

	startAttribute, ok := attributes["start"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`start is missing from object`)

		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	startVal, ok := startAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`start expected to be ovhtypes.TfStringValue, was: %T`, startAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), diags
	}

	return CloudProjectNetworkPrivateSubnetsIpPoolsValue{
		Dhcp:    dhcpVal,
		End:     endVal,
		Network: networkVal,
		Region:  regionVal,
		Start:   startVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivateSubnetsIpPoolsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectNetworkPrivateSubnetsIpPoolsValue {
	object, diags := NewCloudProjectNetworkPrivateSubnetsIpPoolsValue(attributeTypes, attributes)

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

		panic("NewCloudProjectNetworkPrivateSubnetsIpPoolsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectNetworkPrivateSubnetsIpPoolsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueNull(), nil
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

	return NewCloudProjectNetworkPrivateSubnetsIpPoolsValueMust(CloudProjectNetworkPrivateSubnetsIpPoolsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectNetworkPrivateSubnetsIpPoolsType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectNetworkPrivateSubnetsIpPoolsValue{}
}

var _ basetypes.ObjectValuable = CloudProjectNetworkPrivateSubnetsIpPoolsValue{}

type CloudProjectNetworkPrivateSubnetsIpPoolsValue struct {
	Dhcp    ovhtypes.TfBoolValue   `tfsdk:"dhcp" json:"dhcp"`
	End     ovhtypes.TfStringValue `tfsdk:"end" json:"end"`
	Network ovhtypes.TfStringValue `tfsdk:"network" json:"network"`
	Region  ovhtypes.TfStringValue `tfsdk:"region" json:"region"`
	Start   ovhtypes.TfStringValue `tfsdk:"start" json:"start"`
	state   attr.ValueState
}

func (v *CloudProjectNetworkPrivateSubnetsIpPoolsValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectNetworkPrivateSubnetsIpPoolsValue CloudProjectNetworkPrivateSubnetsIpPoolsValue

	var tmp JsonCloudProjectNetworkPrivateSubnetsIpPoolsValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Dhcp = tmp.Dhcp
	v.End = tmp.End
	v.Network = tmp.Network
	v.Region = tmp.Region
	v.Start = tmp.Start

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectNetworkPrivateSubnetsIpPoolsValue) MergeWith(other *CloudProjectNetworkPrivateSubnetsIpPoolsValue) {

	if (v.Dhcp.IsUnknown() || v.Dhcp.IsNull()) && !other.Dhcp.IsUnknown() {
		v.Dhcp = other.Dhcp
	}

	if (v.End.IsUnknown() || v.End.IsNull()) && !other.End.IsUnknown() {
		v.End = other.End
	}

	if (v.Network.IsUnknown() || v.Network.IsNull()) && !other.Network.IsUnknown() {
		v.Network = other.Network
	}

	if (v.Region.IsUnknown() || v.Region.IsNull()) && !other.Region.IsUnknown() {
		v.Region = other.Region
	}

	if (v.Start.IsUnknown() || v.Start.IsNull()) && !other.Start.IsUnknown() {
		v.Start = other.Start
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"dhcp":    v.Dhcp,
		"end":     v.End,
		"network": v.Network,
		"region":  v.Region,
		"start":   v.Start,
	}
}
func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 5)

	var val tftypes.Value
	var err error

	attrTypes["dhcp"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["end"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["network"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["region"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["start"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 5)

		val, err = v.Dhcp.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["dhcp"] = val

		val, err = v.End.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["end"] = val

		val, err = v.Network.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["network"] = val

		val, err = v.Region.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["region"] = val

		val, err = v.Start.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["start"] = val

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

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) String() string {
	return "CloudProjectNetworkPrivateSubnetsIpPoolsValue"
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"dhcp":    ovhtypes.TfBoolType{},
			"end":     ovhtypes.TfStringType{},
			"network": ovhtypes.TfStringType{},
			"region":  ovhtypes.TfStringType{},
			"start":   ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"dhcp":    v.Dhcp,
			"end":     v.End,
			"network": v.Network,
			"region":  v.Region,
			"start":   v.Start,
		})

	return objVal, diags
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectNetworkPrivateSubnetsIpPoolsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Dhcp.Equal(other.Dhcp) {
		return false
	}

	if !v.End.Equal(other.End) {
		return false
	}

	if !v.Network.Equal(other.Network) {
		return false
	}

	if !v.Region.Equal(other.Region) {
		return false
	}

	if !v.Start.Equal(other.Start) {
		return false
	}

	return true
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) Type(ctx context.Context) attr.Type {
	return CloudProjectNetworkPrivateSubnetsIpPoolsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectNetworkPrivateSubnetsIpPoolsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"dhcp":    ovhtypes.TfBoolType{},
		"end":     ovhtypes.TfStringType{},
		"network": ovhtypes.TfStringType{},
		"region":  ovhtypes.TfStringType{},
		"start":   ovhtypes.TfStringType{},
	}
}
