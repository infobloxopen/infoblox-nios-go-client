# MembernodeinfoLanHaPortSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MgmtLan** | Pointer to **string** | Public IPv4 address for the LAN1 interface. | [optional] 
**MgmtIpv6addr** | Pointer to **string** | Public IPv6 address for the LAN1 interface. | [optional] 
**HaIpAddress** | Pointer to **string** | HA IP address. | [optional] 
**LanPortSetting** | Pointer to [**MembernodeinfolanhaportsettingLanPortSetting**](MembernodeinfolanhaportsettingLanPortSetting.md) |  | [optional] 
**HaPortSetting** | Pointer to [**MembernodeinfolanhaportsettingHaPortSetting**](MembernodeinfolanhaportsettingHaPortSetting.md) |  | [optional] 
**HaCloudAttribute** | Pointer to **string** | HA cloud interface from cloud platform side. | [optional] 

## Methods

### NewMembernodeinfoLanHaPortSetting

`func NewMembernodeinfoLanHaPortSetting() *MembernodeinfoLanHaPortSetting`

NewMembernodeinfoLanHaPortSetting instantiates a new MembernodeinfoLanHaPortSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembernodeinfoLanHaPortSettingWithDefaults

`func NewMembernodeinfoLanHaPortSettingWithDefaults() *MembernodeinfoLanHaPortSetting`

NewMembernodeinfoLanHaPortSettingWithDefaults instantiates a new MembernodeinfoLanHaPortSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMgmtLan

`func (o *MembernodeinfoLanHaPortSetting) GetMgmtLan() string`

GetMgmtLan returns the MgmtLan field if non-nil, zero value otherwise.

### GetMgmtLanOk

`func (o *MembernodeinfoLanHaPortSetting) GetMgmtLanOk() (*string, bool)`

GetMgmtLanOk returns a tuple with the MgmtLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtLan

`func (o *MembernodeinfoLanHaPortSetting) SetMgmtLan(v string)`

SetMgmtLan sets MgmtLan field to given value.

### HasMgmtLan

`func (o *MembernodeinfoLanHaPortSetting) HasMgmtLan() bool`

HasMgmtLan returns a boolean if a field has been set.

### GetMgmtIpv6addr

`func (o *MembernodeinfoLanHaPortSetting) GetMgmtIpv6addr() string`

GetMgmtIpv6addr returns the MgmtIpv6addr field if non-nil, zero value otherwise.

### GetMgmtIpv6addrOk

`func (o *MembernodeinfoLanHaPortSetting) GetMgmtIpv6addrOk() (*string, bool)`

GetMgmtIpv6addrOk returns a tuple with the MgmtIpv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpv6addr

`func (o *MembernodeinfoLanHaPortSetting) SetMgmtIpv6addr(v string)`

SetMgmtIpv6addr sets MgmtIpv6addr field to given value.

### HasMgmtIpv6addr

`func (o *MembernodeinfoLanHaPortSetting) HasMgmtIpv6addr() bool`

HasMgmtIpv6addr returns a boolean if a field has been set.

### GetHaIpAddress

`func (o *MembernodeinfoLanHaPortSetting) GetHaIpAddress() string`

GetHaIpAddress returns the HaIpAddress field if non-nil, zero value otherwise.

### GetHaIpAddressOk

`func (o *MembernodeinfoLanHaPortSetting) GetHaIpAddressOk() (*string, bool)`

GetHaIpAddressOk returns a tuple with the HaIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIpAddress

`func (o *MembernodeinfoLanHaPortSetting) SetHaIpAddress(v string)`

SetHaIpAddress sets HaIpAddress field to given value.

### HasHaIpAddress

`func (o *MembernodeinfoLanHaPortSetting) HasHaIpAddress() bool`

HasHaIpAddress returns a boolean if a field has been set.

### GetLanPortSetting

`func (o *MembernodeinfoLanHaPortSetting) GetLanPortSetting() MembernodeinfolanhaportsettingLanPortSetting`

GetLanPortSetting returns the LanPortSetting field if non-nil, zero value otherwise.

### GetLanPortSettingOk

`func (o *MembernodeinfoLanHaPortSetting) GetLanPortSettingOk() (*MembernodeinfolanhaportsettingLanPortSetting, bool)`

GetLanPortSettingOk returns a tuple with the LanPortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPortSetting

`func (o *MembernodeinfoLanHaPortSetting) SetLanPortSetting(v MembernodeinfolanhaportsettingLanPortSetting)`

SetLanPortSetting sets LanPortSetting field to given value.

### HasLanPortSetting

`func (o *MembernodeinfoLanHaPortSetting) HasLanPortSetting() bool`

HasLanPortSetting returns a boolean if a field has been set.

### GetHaPortSetting

`func (o *MembernodeinfoLanHaPortSetting) GetHaPortSetting() MembernodeinfolanhaportsettingHaPortSetting`

GetHaPortSetting returns the HaPortSetting field if non-nil, zero value otherwise.

### GetHaPortSettingOk

`func (o *MembernodeinfoLanHaPortSetting) GetHaPortSettingOk() (*MembernodeinfolanhaportsettingHaPortSetting, bool)`

GetHaPortSettingOk returns a tuple with the HaPortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPortSetting

`func (o *MembernodeinfoLanHaPortSetting) SetHaPortSetting(v MembernodeinfolanhaportsettingHaPortSetting)`

SetHaPortSetting sets HaPortSetting field to given value.

### HasHaPortSetting

`func (o *MembernodeinfoLanHaPortSetting) HasHaPortSetting() bool`

HasHaPortSetting returns a boolean if a field has been set.

### GetHaCloudAttribute

`func (o *MembernodeinfoLanHaPortSetting) GetHaCloudAttribute() string`

GetHaCloudAttribute returns the HaCloudAttribute field if non-nil, zero value otherwise.

### GetHaCloudAttributeOk

`func (o *MembernodeinfoLanHaPortSetting) GetHaCloudAttributeOk() (*string, bool)`

GetHaCloudAttributeOk returns a tuple with the HaCloudAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaCloudAttribute

`func (o *MembernodeinfoLanHaPortSetting) SetHaCloudAttribute(v string)`

SetHaCloudAttribute sets HaCloudAttribute field to given value.

### HasHaCloudAttribute

`func (o *MembernodeinfoLanHaPortSetting) HasHaCloudAttribute() bool`

HasHaCloudAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


