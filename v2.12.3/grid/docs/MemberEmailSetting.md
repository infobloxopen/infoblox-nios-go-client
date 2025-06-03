# MemberEmailSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines if email notification is enabled or not. | [optional] 
**FromAddress** | Pointer to **string** | The email address of a Grid Member for &#39;from&#39; field in notification. | [optional] 
**Address** | Pointer to **string** | The notification email address of a Grid member. | [optional] 
**RelayEnabled** | Pointer to **bool** | Determines if email relay is enabled or not. | [optional] 
**Relay** | Pointer to **string** | The relay name or IP address. | [optional] 
**Password** | Pointer to **string** | Password to validate from address | [optional] 
**Smtps** | Pointer to **bool** | SMTP over TLS | [optional] 
**PortNumber** | Pointer to **int64** | SMTP port number | [optional] 
**UseAuthentication** | Pointer to **bool** | Enable or disable SMTP auth | [optional] 

## Methods

### NewMemberEmailSetting

`func NewMemberEmailSetting() *MemberEmailSetting`

NewMemberEmailSetting instantiates a new MemberEmailSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberEmailSettingWithDefaults

`func NewMemberEmailSettingWithDefaults() *MemberEmailSetting`

NewMemberEmailSettingWithDefaults instantiates a new MemberEmailSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MemberEmailSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberEmailSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberEmailSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberEmailSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFromAddress

`func (o *MemberEmailSetting) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MemberEmailSetting) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MemberEmailSetting) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MemberEmailSetting) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetAddress

`func (o *MemberEmailSetting) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberEmailSetting) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberEmailSetting) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberEmailSetting) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRelayEnabled

`func (o *MemberEmailSetting) GetRelayEnabled() bool`

GetRelayEnabled returns the RelayEnabled field if non-nil, zero value otherwise.

### GetRelayEnabledOk

`func (o *MemberEmailSetting) GetRelayEnabledOk() (*bool, bool)`

GetRelayEnabledOk returns a tuple with the RelayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayEnabled

`func (o *MemberEmailSetting) SetRelayEnabled(v bool)`

SetRelayEnabled sets RelayEnabled field to given value.

### HasRelayEnabled

`func (o *MemberEmailSetting) HasRelayEnabled() bool`

HasRelayEnabled returns a boolean if a field has been set.

### GetRelay

`func (o *MemberEmailSetting) GetRelay() string`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *MemberEmailSetting) GetRelayOk() (*string, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *MemberEmailSetting) SetRelay(v string)`

SetRelay sets Relay field to given value.

### HasRelay

`func (o *MemberEmailSetting) HasRelay() bool`

HasRelay returns a boolean if a field has been set.

### GetPassword

`func (o *MemberEmailSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MemberEmailSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MemberEmailSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MemberEmailSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSmtps

`func (o *MemberEmailSetting) GetSmtps() bool`

GetSmtps returns the Smtps field if non-nil, zero value otherwise.

### GetSmtpsOk

`func (o *MemberEmailSetting) GetSmtpsOk() (*bool, bool)`

GetSmtpsOk returns a tuple with the Smtps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtps

`func (o *MemberEmailSetting) SetSmtps(v bool)`

SetSmtps sets Smtps field to given value.

### HasSmtps

`func (o *MemberEmailSetting) HasSmtps() bool`

HasSmtps returns a boolean if a field has been set.

### GetPortNumber

`func (o *MemberEmailSetting) GetPortNumber() int64`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *MemberEmailSetting) GetPortNumberOk() (*int64, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *MemberEmailSetting) SetPortNumber(v int64)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *MemberEmailSetting) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.

### GetUseAuthentication

`func (o *MemberEmailSetting) GetUseAuthentication() bool`

GetUseAuthentication returns the UseAuthentication field if non-nil, zero value otherwise.

### GetUseAuthenticationOk

`func (o *MemberEmailSetting) GetUseAuthenticationOk() (*bool, bool)`

GetUseAuthenticationOk returns a tuple with the UseAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthentication

`func (o *MemberEmailSetting) SetUseAuthentication(v bool)`

SetUseAuthentication sets UseAuthentication field to given value.

### HasUseAuthentication

`func (o *MemberEmailSetting) HasUseAuthentication() bool`

HasUseAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


