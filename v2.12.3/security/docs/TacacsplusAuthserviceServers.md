# TacacsplusAuthserviceServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The valid IP address or FQDN of the TACACS+ server. | [optional] 
**Port** | Pointer to **int64** | The TACACS+ server port. | [optional] 
**SharedSecret** | Pointer to **string** | The secret key with which to connect to the TACACS+ server. | [optional] 
**AuthType** | Pointer to **string** | The authentication protocol. | [optional] 
**Comment** | Pointer to **string** | The TACACS+ descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the TACACS+ server is disabled. | [optional] 
**UseMgmtPort** | Pointer to **bool** | Determines whether the TACACS+ server is connected via the management interface. | [optional] 
**UseAccounting** | Pointer to **bool** | Determines whether the TACACS+ accounting server is used. | [optional] 

## Methods

### NewTacacsplusAuthserviceServers

`func NewTacacsplusAuthserviceServers() *TacacsplusAuthserviceServers`

NewTacacsplusAuthserviceServers instantiates a new TacacsplusAuthserviceServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsplusAuthserviceServersWithDefaults

`func NewTacacsplusAuthserviceServersWithDefaults() *TacacsplusAuthserviceServers`

NewTacacsplusAuthserviceServersWithDefaults instantiates a new TacacsplusAuthserviceServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TacacsplusAuthserviceServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TacacsplusAuthserviceServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TacacsplusAuthserviceServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TacacsplusAuthserviceServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *TacacsplusAuthserviceServers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TacacsplusAuthserviceServers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TacacsplusAuthserviceServers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *TacacsplusAuthserviceServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSharedSecret

`func (o *TacacsplusAuthserviceServers) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *TacacsplusAuthserviceServers) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *TacacsplusAuthserviceServers) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *TacacsplusAuthserviceServers) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthType

`func (o *TacacsplusAuthserviceServers) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *TacacsplusAuthserviceServers) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *TacacsplusAuthserviceServers) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *TacacsplusAuthserviceServers) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetComment

`func (o *TacacsplusAuthserviceServers) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TacacsplusAuthserviceServers) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TacacsplusAuthserviceServers) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TacacsplusAuthserviceServers) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *TacacsplusAuthserviceServers) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *TacacsplusAuthserviceServers) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *TacacsplusAuthserviceServers) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *TacacsplusAuthserviceServers) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *TacacsplusAuthserviceServers) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *TacacsplusAuthserviceServers) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *TacacsplusAuthserviceServers) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *TacacsplusAuthserviceServers) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.

### GetUseAccounting

`func (o *TacacsplusAuthserviceServers) GetUseAccounting() bool`

GetUseAccounting returns the UseAccounting field if non-nil, zero value otherwise.

### GetUseAccountingOk

`func (o *TacacsplusAuthserviceServers) GetUseAccountingOk() (*bool, bool)`

GetUseAccountingOk returns a tuple with the UseAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAccounting

`func (o *TacacsplusAuthserviceServers) SetUseAccounting(v bool)`

SetUseAccounting sets UseAccounting field to given value.

### HasUseAccounting

`func (o *TacacsplusAuthserviceServers) HasUseAccounting() bool`

HasUseAccounting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


