# MemberadditionaliplistIpv6NetworkSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines if IPv6 networking should be enabled. | [optional] 
**VirtualIp** | Pointer to **string** | IPv6 address. | [optional] 
**CidrPrefix** | Pointer to **int64** | IPv6 cidr prefix | [optional] 
**Gateway** | Pointer to **string** | Gateway address. | [optional] 
**AutoRouterConfigEnabled** | Pointer to **bool** | Determines if automatic router configuration should be enabled. | [optional] 
**VlanId** | Pointer to **int64** | The identifier for the VLAN. Valid values are from 1 to 4096. | [optional] 
**Primary** | Pointer to **bool** | Determines if the current address is the primary VLAN address or not. | [optional] 
**Dscp** | Pointer to **int64** | The DSCP (Differentiated Services Code Point) value determines relative priorities for the type of services on your network. The appliance implements QoS (Quality of Service) rules based on this configuration. Valid values are from 0 to 63. | [optional] 
**UseDscp** | Pointer to **bool** | Use flag for: dscp | [optional] 

## Methods

### NewMemberadditionaliplistIpv6NetworkSetting

`func NewMemberadditionaliplistIpv6NetworkSetting() *MemberadditionaliplistIpv6NetworkSetting`

NewMemberadditionaliplistIpv6NetworkSetting instantiates a new MemberadditionaliplistIpv6NetworkSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberadditionaliplistIpv6NetworkSettingWithDefaults

`func NewMemberadditionaliplistIpv6NetworkSettingWithDefaults() *MemberadditionaliplistIpv6NetworkSetting`

NewMemberadditionaliplistIpv6NetworkSettingWithDefaults instantiates a new MemberadditionaliplistIpv6NetworkSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVirtualIp

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetCidrPrefix

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetCidrPrefix() int64`

GetCidrPrefix returns the CidrPrefix field if non-nil, zero value otherwise.

### GetCidrPrefixOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetCidrPrefixOk() (*int64, bool)`

GetCidrPrefixOk returns a tuple with the CidrPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrPrefix

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetCidrPrefix(v int64)`

SetCidrPrefix sets CidrPrefix field to given value.

### HasCidrPrefix

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasCidrPrefix() bool`

HasCidrPrefix returns a boolean if a field has been set.

### GetGateway

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetAutoRouterConfigEnabled

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetAutoRouterConfigEnabled() bool`

GetAutoRouterConfigEnabled returns the AutoRouterConfigEnabled field if non-nil, zero value otherwise.

### GetAutoRouterConfigEnabledOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetAutoRouterConfigEnabledOk() (*bool, bool)`

GetAutoRouterConfigEnabledOk returns a tuple with the AutoRouterConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterConfigEnabled

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetAutoRouterConfigEnabled(v bool)`

SetAutoRouterConfigEnabled sets AutoRouterConfigEnabled field to given value.

### HasAutoRouterConfigEnabled

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasAutoRouterConfigEnabled() bool`

HasAutoRouterConfigEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPrimary

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetDscp

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetUseDscp

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *MemberadditionaliplistIpv6NetworkSetting) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *MemberadditionaliplistIpv6NetworkSetting) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *MemberadditionaliplistIpv6NetworkSetting) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


