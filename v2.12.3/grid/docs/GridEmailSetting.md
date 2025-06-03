# GridEmailSetting

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

### NewGridEmailSetting

`func NewGridEmailSetting() *GridEmailSetting`

NewGridEmailSetting instantiates a new GridEmailSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridEmailSettingWithDefaults

`func NewGridEmailSettingWithDefaults() *GridEmailSetting`

NewGridEmailSettingWithDefaults instantiates a new GridEmailSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GridEmailSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GridEmailSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GridEmailSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GridEmailSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFromAddress

`func (o *GridEmailSetting) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GridEmailSetting) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GridEmailSetting) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *GridEmailSetting) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetAddress

`func (o *GridEmailSetting) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GridEmailSetting) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GridEmailSetting) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GridEmailSetting) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRelayEnabled

`func (o *GridEmailSetting) GetRelayEnabled() bool`

GetRelayEnabled returns the RelayEnabled field if non-nil, zero value otherwise.

### GetRelayEnabledOk

`func (o *GridEmailSetting) GetRelayEnabledOk() (*bool, bool)`

GetRelayEnabledOk returns a tuple with the RelayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayEnabled

`func (o *GridEmailSetting) SetRelayEnabled(v bool)`

SetRelayEnabled sets RelayEnabled field to given value.

### HasRelayEnabled

`func (o *GridEmailSetting) HasRelayEnabled() bool`

HasRelayEnabled returns a boolean if a field has been set.

### GetRelay

`func (o *GridEmailSetting) GetRelay() string`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *GridEmailSetting) GetRelayOk() (*string, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *GridEmailSetting) SetRelay(v string)`

SetRelay sets Relay field to given value.

### HasRelay

`func (o *GridEmailSetting) HasRelay() bool`

HasRelay returns a boolean if a field has been set.

### GetPassword

`func (o *GridEmailSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GridEmailSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GridEmailSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GridEmailSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSmtps

`func (o *GridEmailSetting) GetSmtps() bool`

GetSmtps returns the Smtps field if non-nil, zero value otherwise.

### GetSmtpsOk

`func (o *GridEmailSetting) GetSmtpsOk() (*bool, bool)`

GetSmtpsOk returns a tuple with the Smtps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtps

`func (o *GridEmailSetting) SetSmtps(v bool)`

SetSmtps sets Smtps field to given value.

### HasSmtps

`func (o *GridEmailSetting) HasSmtps() bool`

HasSmtps returns a boolean if a field has been set.

### GetPortNumber

`func (o *GridEmailSetting) GetPortNumber() int64`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *GridEmailSetting) GetPortNumberOk() (*int64, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *GridEmailSetting) SetPortNumber(v int64)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *GridEmailSetting) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.

### GetUseAuthentication

`func (o *GridEmailSetting) GetUseAuthentication() bool`

GetUseAuthentication returns the UseAuthentication field if non-nil, zero value otherwise.

### GetUseAuthenticationOk

`func (o *GridEmailSetting) GetUseAuthenticationOk() (*bool, bool)`

GetUseAuthenticationOk returns a tuple with the UseAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthentication

`func (o *GridEmailSetting) SetUseAuthentication(v bool)`

SetUseAuthentication sets UseAuthentication field to given value.

### HasUseAuthentication

`func (o *GridEmailSetting) HasUseAuthentication() bool`

HasUseAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


