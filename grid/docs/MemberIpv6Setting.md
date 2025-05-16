# MemberIpv6Setting

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

### NewMemberIpv6Setting

`func NewMemberIpv6Setting() *MemberIpv6Setting`

NewMemberIpv6Setting instantiates a new MemberIpv6Setting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberIpv6SettingWithDefaults

`func NewMemberIpv6SettingWithDefaults() *MemberIpv6Setting`

NewMemberIpv6SettingWithDefaults instantiates a new MemberIpv6Setting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MemberIpv6Setting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberIpv6Setting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberIpv6Setting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberIpv6Setting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVirtualIp

`func (o *MemberIpv6Setting) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *MemberIpv6Setting) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *MemberIpv6Setting) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *MemberIpv6Setting) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetCidrPrefix

`func (o *MemberIpv6Setting) GetCidrPrefix() int64`

GetCidrPrefix returns the CidrPrefix field if non-nil, zero value otherwise.

### GetCidrPrefixOk

`func (o *MemberIpv6Setting) GetCidrPrefixOk() (*int64, bool)`

GetCidrPrefixOk returns a tuple with the CidrPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrPrefix

`func (o *MemberIpv6Setting) SetCidrPrefix(v int64)`

SetCidrPrefix sets CidrPrefix field to given value.

### HasCidrPrefix

`func (o *MemberIpv6Setting) HasCidrPrefix() bool`

HasCidrPrefix returns a boolean if a field has been set.

### GetGateway

`func (o *MemberIpv6Setting) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MemberIpv6Setting) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MemberIpv6Setting) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MemberIpv6Setting) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetAutoRouterConfigEnabled

`func (o *MemberIpv6Setting) GetAutoRouterConfigEnabled() bool`

GetAutoRouterConfigEnabled returns the AutoRouterConfigEnabled field if non-nil, zero value otherwise.

### GetAutoRouterConfigEnabledOk

`func (o *MemberIpv6Setting) GetAutoRouterConfigEnabledOk() (*bool, bool)`

GetAutoRouterConfigEnabledOk returns a tuple with the AutoRouterConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterConfigEnabled

`func (o *MemberIpv6Setting) SetAutoRouterConfigEnabled(v bool)`

SetAutoRouterConfigEnabled sets AutoRouterConfigEnabled field to given value.

### HasAutoRouterConfigEnabled

`func (o *MemberIpv6Setting) HasAutoRouterConfigEnabled() bool`

HasAutoRouterConfigEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *MemberIpv6Setting) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *MemberIpv6Setting) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *MemberIpv6Setting) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *MemberIpv6Setting) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPrimary

`func (o *MemberIpv6Setting) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MemberIpv6Setting) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MemberIpv6Setting) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MemberIpv6Setting) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetDscp

`func (o *MemberIpv6Setting) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *MemberIpv6Setting) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *MemberIpv6Setting) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *MemberIpv6Setting) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetUseDscp

`func (o *MemberIpv6Setting) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *MemberIpv6Setting) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *MemberIpv6Setting) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *MemberIpv6Setting) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


