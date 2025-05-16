# MemberLan2PortSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualRouterId** | Pointer to **int64** | If the &#39;enabled&#39; field is set to True, this defines the virtual router ID for the LAN2 port. | [optional] 
**Enabled** | Pointer to **bool** | If this field is set to True, then it has its own IP settings. Otherwise, port redundancy mechanism is used, in which the LAN1 and LAN2 ports share the same IP settings for failover purposes. | [optional] 
**NetworkSetting** | Pointer to [**Memberlan2portsettingNetworkSetting**](Memberlan2portsettingNetworkSetting.md) |  | [optional] 
**V6NetworkSetting** | Pointer to [**Memberlan2portsettingV6NetworkSetting**](Memberlan2portsettingV6NetworkSetting.md) |  | [optional] 
**NicFailoverEnabled** | Pointer to **bool** | Determines if NIC failover is enabled or not. | [optional] 
**NicFailoverEnablePrimary** | Pointer to **bool** | Prefer LAN1 when available. | [optional] 
**DefaultRouteFailoverEnabled** | Pointer to **bool** | Default route failover for LAN1 and LAN2. | [optional] 

## Methods

### NewMemberLan2PortSetting

`func NewMemberLan2PortSetting() *MemberLan2PortSetting`

NewMemberLan2PortSetting instantiates a new MemberLan2PortSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberLan2PortSettingWithDefaults

`func NewMemberLan2PortSettingWithDefaults() *MemberLan2PortSetting`

NewMemberLan2PortSettingWithDefaults instantiates a new MemberLan2PortSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualRouterId

`func (o *MemberLan2PortSetting) GetVirtualRouterId() int64`

GetVirtualRouterId returns the VirtualRouterId field if non-nil, zero value otherwise.

### GetVirtualRouterIdOk

`func (o *MemberLan2PortSetting) GetVirtualRouterIdOk() (*int64, bool)`

GetVirtualRouterIdOk returns a tuple with the VirtualRouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouterId

`func (o *MemberLan2PortSetting) SetVirtualRouterId(v int64)`

SetVirtualRouterId sets VirtualRouterId field to given value.

### HasVirtualRouterId

`func (o *MemberLan2PortSetting) HasVirtualRouterId() bool`

HasVirtualRouterId returns a boolean if a field has been set.

### GetEnabled

`func (o *MemberLan2PortSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberLan2PortSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberLan2PortSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberLan2PortSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetworkSetting

`func (o *MemberLan2PortSetting) GetNetworkSetting() Memberlan2portsettingNetworkSetting`

GetNetworkSetting returns the NetworkSetting field if non-nil, zero value otherwise.

### GetNetworkSettingOk

`func (o *MemberLan2PortSetting) GetNetworkSettingOk() (*Memberlan2portsettingNetworkSetting, bool)`

GetNetworkSettingOk returns a tuple with the NetworkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSetting

`func (o *MemberLan2PortSetting) SetNetworkSetting(v Memberlan2portsettingNetworkSetting)`

SetNetworkSetting sets NetworkSetting field to given value.

### HasNetworkSetting

`func (o *MemberLan2PortSetting) HasNetworkSetting() bool`

HasNetworkSetting returns a boolean if a field has been set.

### GetV6NetworkSetting

`func (o *MemberLan2PortSetting) GetV6NetworkSetting() Memberlan2portsettingV6NetworkSetting`

GetV6NetworkSetting returns the V6NetworkSetting field if non-nil, zero value otherwise.

### GetV6NetworkSettingOk

`func (o *MemberLan2PortSetting) GetV6NetworkSettingOk() (*Memberlan2portsettingV6NetworkSetting, bool)`

GetV6NetworkSettingOk returns a tuple with the V6NetworkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6NetworkSetting

`func (o *MemberLan2PortSetting) SetV6NetworkSetting(v Memberlan2portsettingV6NetworkSetting)`

SetV6NetworkSetting sets V6NetworkSetting field to given value.

### HasV6NetworkSetting

`func (o *MemberLan2PortSetting) HasV6NetworkSetting() bool`

HasV6NetworkSetting returns a boolean if a field has been set.

### GetNicFailoverEnabled

`func (o *MemberLan2PortSetting) GetNicFailoverEnabled() bool`

GetNicFailoverEnabled returns the NicFailoverEnabled field if non-nil, zero value otherwise.

### GetNicFailoverEnabledOk

`func (o *MemberLan2PortSetting) GetNicFailoverEnabledOk() (*bool, bool)`

GetNicFailoverEnabledOk returns a tuple with the NicFailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicFailoverEnabled

`func (o *MemberLan2PortSetting) SetNicFailoverEnabled(v bool)`

SetNicFailoverEnabled sets NicFailoverEnabled field to given value.

### HasNicFailoverEnabled

`func (o *MemberLan2PortSetting) HasNicFailoverEnabled() bool`

HasNicFailoverEnabled returns a boolean if a field has been set.

### GetNicFailoverEnablePrimary

`func (o *MemberLan2PortSetting) GetNicFailoverEnablePrimary() bool`

GetNicFailoverEnablePrimary returns the NicFailoverEnablePrimary field if non-nil, zero value otherwise.

### GetNicFailoverEnablePrimaryOk

`func (o *MemberLan2PortSetting) GetNicFailoverEnablePrimaryOk() (*bool, bool)`

GetNicFailoverEnablePrimaryOk returns a tuple with the NicFailoverEnablePrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicFailoverEnablePrimary

`func (o *MemberLan2PortSetting) SetNicFailoverEnablePrimary(v bool)`

SetNicFailoverEnablePrimary sets NicFailoverEnablePrimary field to given value.

### HasNicFailoverEnablePrimary

`func (o *MemberLan2PortSetting) HasNicFailoverEnablePrimary() bool`

HasNicFailoverEnablePrimary returns a boolean if a field has been set.

### GetDefaultRouteFailoverEnabled

`func (o *MemberLan2PortSetting) GetDefaultRouteFailoverEnabled() bool`

GetDefaultRouteFailoverEnabled returns the DefaultRouteFailoverEnabled field if non-nil, zero value otherwise.

### GetDefaultRouteFailoverEnabledOk

`func (o *MemberLan2PortSetting) GetDefaultRouteFailoverEnabledOk() (*bool, bool)`

GetDefaultRouteFailoverEnabledOk returns a tuple with the DefaultRouteFailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRouteFailoverEnabled

`func (o *MemberLan2PortSetting) SetDefaultRouteFailoverEnabled(v bool)`

SetDefaultRouteFailoverEnabled sets DefaultRouteFailoverEnabled field to given value.

### HasDefaultRouteFailoverEnabled

`func (o *MemberLan2PortSetting) HasDefaultRouteFailoverEnabled() bool`

HasDefaultRouteFailoverEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


