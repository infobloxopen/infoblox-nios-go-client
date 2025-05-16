# MembernodeinfoV6MgmtNetworkSetting

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

### NewMembernodeinfoV6MgmtNetworkSetting

`func NewMembernodeinfoV6MgmtNetworkSetting() *MembernodeinfoV6MgmtNetworkSetting`

NewMembernodeinfoV6MgmtNetworkSetting instantiates a new MembernodeinfoV6MgmtNetworkSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembernodeinfoV6MgmtNetworkSettingWithDefaults

`func NewMembernodeinfoV6MgmtNetworkSettingWithDefaults() *MembernodeinfoV6MgmtNetworkSetting`

NewMembernodeinfoV6MgmtNetworkSettingWithDefaults instantiates a new MembernodeinfoV6MgmtNetworkSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVirtualIp

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetCidrPrefix

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetCidrPrefix() int64`

GetCidrPrefix returns the CidrPrefix field if non-nil, zero value otherwise.

### GetCidrPrefixOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetCidrPrefixOk() (*int64, bool)`

GetCidrPrefixOk returns a tuple with the CidrPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrPrefix

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetCidrPrefix(v int64)`

SetCidrPrefix sets CidrPrefix field to given value.

### HasCidrPrefix

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasCidrPrefix() bool`

HasCidrPrefix returns a boolean if a field has been set.

### GetGateway

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetAutoRouterConfigEnabled

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetAutoRouterConfigEnabled() bool`

GetAutoRouterConfigEnabled returns the AutoRouterConfigEnabled field if non-nil, zero value otherwise.

### GetAutoRouterConfigEnabledOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetAutoRouterConfigEnabledOk() (*bool, bool)`

GetAutoRouterConfigEnabledOk returns a tuple with the AutoRouterConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterConfigEnabled

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetAutoRouterConfigEnabled(v bool)`

SetAutoRouterConfigEnabled sets AutoRouterConfigEnabled field to given value.

### HasAutoRouterConfigEnabled

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasAutoRouterConfigEnabled() bool`

HasAutoRouterConfigEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPrimary

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetDscp

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetUseDscp

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *MembernodeinfoV6MgmtNetworkSetting) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *MembernodeinfoV6MgmtNetworkSetting) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *MembernodeinfoV6MgmtNetworkSetting) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


