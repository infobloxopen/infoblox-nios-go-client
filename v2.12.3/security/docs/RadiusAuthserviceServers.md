# RadiusAuthserviceServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctPort** | Pointer to **int64** | The accounting port. | [optional] 
**AuthPort** | Pointer to **int64** | The authorization port. | [optional] 
**AuthType** | Pointer to **string** | The authentication protocol. | [optional] 
**Comment** | Pointer to **string** | The RADIUS descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the RADIUS server is disabled. | [optional] 
**Address** | Pointer to **string** | The FQDN or the IP address of the RADIUS server that is used for authentication. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret that the NIOS appliance and the RADIUS server use to encrypt and decrypt their messages. | [optional] 
**UseAccounting** | Pointer to **bool** | Determines whether RADIUS accounting is enabled. | [optional] 
**UseMgmtPort** | Pointer to **bool** | Determines whether connection via the management interface is allowed. | [optional] 

## Methods

### NewRadiusAuthserviceServers

`func NewRadiusAuthserviceServers() *RadiusAuthserviceServers`

NewRadiusAuthserviceServers instantiates a new RadiusAuthserviceServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAuthserviceServersWithDefaults

`func NewRadiusAuthserviceServersWithDefaults() *RadiusAuthserviceServers`

NewRadiusAuthserviceServersWithDefaults instantiates a new RadiusAuthserviceServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctPort

`func (o *RadiusAuthserviceServers) GetAcctPort() int64`

GetAcctPort returns the AcctPort field if non-nil, zero value otherwise.

### GetAcctPortOk

`func (o *RadiusAuthserviceServers) GetAcctPortOk() (*int64, bool)`

GetAcctPortOk returns a tuple with the AcctPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctPort

`func (o *RadiusAuthserviceServers) SetAcctPort(v int64)`

SetAcctPort sets AcctPort field to given value.

### HasAcctPort

`func (o *RadiusAuthserviceServers) HasAcctPort() bool`

HasAcctPort returns a boolean if a field has been set.

### GetAuthPort

`func (o *RadiusAuthserviceServers) GetAuthPort() int64`

GetAuthPort returns the AuthPort field if non-nil, zero value otherwise.

### GetAuthPortOk

`func (o *RadiusAuthserviceServers) GetAuthPortOk() (*int64, bool)`

GetAuthPortOk returns a tuple with the AuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPort

`func (o *RadiusAuthserviceServers) SetAuthPort(v int64)`

SetAuthPort sets AuthPort field to given value.

### HasAuthPort

`func (o *RadiusAuthserviceServers) HasAuthPort() bool`

HasAuthPort returns a boolean if a field has been set.

### GetAuthType

`func (o *RadiusAuthserviceServers) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RadiusAuthserviceServers) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RadiusAuthserviceServers) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *RadiusAuthserviceServers) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetComment

`func (o *RadiusAuthserviceServers) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RadiusAuthserviceServers) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RadiusAuthserviceServers) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RadiusAuthserviceServers) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RadiusAuthserviceServers) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RadiusAuthserviceServers) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RadiusAuthserviceServers) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RadiusAuthserviceServers) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetAddress

`func (o *RadiusAuthserviceServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RadiusAuthserviceServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RadiusAuthserviceServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RadiusAuthserviceServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSharedSecret

`func (o *RadiusAuthserviceServers) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *RadiusAuthserviceServers) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *RadiusAuthserviceServers) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *RadiusAuthserviceServers) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetUseAccounting

`func (o *RadiusAuthserviceServers) GetUseAccounting() bool`

GetUseAccounting returns the UseAccounting field if non-nil, zero value otherwise.

### GetUseAccountingOk

`func (o *RadiusAuthserviceServers) GetUseAccountingOk() (*bool, bool)`

GetUseAccountingOk returns a tuple with the UseAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAccounting

`func (o *RadiusAuthserviceServers) SetUseAccounting(v bool)`

SetUseAccounting sets UseAccounting field to given value.

### HasUseAccounting

`func (o *RadiusAuthserviceServers) HasUseAccounting() bool`

HasUseAccounting returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *RadiusAuthserviceServers) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *RadiusAuthserviceServers) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *RadiusAuthserviceServers) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *RadiusAuthserviceServers) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


