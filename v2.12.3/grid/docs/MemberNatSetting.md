# MemberNatSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines if NAT should be enabled. | [optional] 
**ExternalVirtualIp** | Pointer to **string** | External IP address for NAT. | [optional] 
**Group** | Pointer to **string** | The NAT group. | [optional] 

## Methods

### NewMemberNatSetting

`func NewMemberNatSetting() *MemberNatSetting`

NewMemberNatSetting instantiates a new MemberNatSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberNatSettingWithDefaults

`func NewMemberNatSettingWithDefaults() *MemberNatSetting`

NewMemberNatSettingWithDefaults instantiates a new MemberNatSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MemberNatSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberNatSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberNatSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberNatSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalVirtualIp

`func (o *MemberNatSetting) GetExternalVirtualIp() string`

GetExternalVirtualIp returns the ExternalVirtualIp field if non-nil, zero value otherwise.

### GetExternalVirtualIpOk

`func (o *MemberNatSetting) GetExternalVirtualIpOk() (*string, bool)`

GetExternalVirtualIpOk returns a tuple with the ExternalVirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVirtualIp

`func (o *MemberNatSetting) SetExternalVirtualIp(v string)`

SetExternalVirtualIp sets ExternalVirtualIp field to given value.

### HasExternalVirtualIp

`func (o *MemberNatSetting) HasExternalVirtualIp() bool`

HasExternalVirtualIp returns a boolean if a field has been set.

### GetGroup

`func (o *MemberNatSetting) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MemberNatSetting) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MemberNatSetting) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MemberNatSetting) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


