# MemberMgmtPortSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines if MGMT port settings should be enabled. | [optional] 
**VpnEnabled** | Pointer to **bool** | Determines if VPN on the MGMT port is enabled or not. | [optional] 
**SecurityAccessEnabled** | Pointer to **bool** | Determines if security access on the MGMT port is enabled or not. | [optional] 

## Methods

### NewMemberMgmtPortSetting

`func NewMemberMgmtPortSetting() *MemberMgmtPortSetting`

NewMemberMgmtPortSetting instantiates a new MemberMgmtPortSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberMgmtPortSettingWithDefaults

`func NewMemberMgmtPortSettingWithDefaults() *MemberMgmtPortSetting`

NewMemberMgmtPortSettingWithDefaults instantiates a new MemberMgmtPortSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MemberMgmtPortSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberMgmtPortSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberMgmtPortSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberMgmtPortSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVpnEnabled

`func (o *MemberMgmtPortSetting) GetVpnEnabled() bool`

GetVpnEnabled returns the VpnEnabled field if non-nil, zero value otherwise.

### GetVpnEnabledOk

`func (o *MemberMgmtPortSetting) GetVpnEnabledOk() (*bool, bool)`

GetVpnEnabledOk returns a tuple with the VpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnEnabled

`func (o *MemberMgmtPortSetting) SetVpnEnabled(v bool)`

SetVpnEnabled sets VpnEnabled field to given value.

### HasVpnEnabled

`func (o *MemberMgmtPortSetting) HasVpnEnabled() bool`

HasVpnEnabled returns a boolean if a field has been set.

### GetSecurityAccessEnabled

`func (o *MemberMgmtPortSetting) GetSecurityAccessEnabled() bool`

GetSecurityAccessEnabled returns the SecurityAccessEnabled field if non-nil, zero value otherwise.

### GetSecurityAccessEnabledOk

`func (o *MemberMgmtPortSetting) GetSecurityAccessEnabledOk() (*bool, bool)`

GetSecurityAccessEnabledOk returns a tuple with the SecurityAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAccessEnabled

`func (o *MemberMgmtPortSetting) SetSecurityAccessEnabled(v bool)`

SetSecurityAccessEnabled sets SecurityAccessEnabled field to given value.

### HasSecurityAccessEnabled

`func (o *MemberMgmtPortSetting) HasSecurityAccessEnabled() bool`

HasSecurityAccessEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


