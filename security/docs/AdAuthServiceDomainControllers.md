# AdAuthServiceDomainControllers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FqdnOrIp** | Pointer to **string** | The FQDN (Fully Qualified Domain Name) or IP address of the server. | [optional] 
**AuthPort** | Pointer to **int64** | The authentication port. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the AD authentication server. | [optional] 
**Disabled** | Pointer to **bool** | Determines if the AD authorization server is disabled. | [optional] 
**Encryption** | Pointer to **string** | The type of encryption to use. | [optional] 
**MgmtPort** | Pointer to **bool** | Determine if the MGMT port is enabled for the AD authentication server. | [optional] 
**UseMgmtPort** | Pointer to **bool** | Use flag for: mgmt_port | [optional] 

## Methods

### NewAdAuthServiceDomainControllers

`func NewAdAuthServiceDomainControllers() *AdAuthServiceDomainControllers`

NewAdAuthServiceDomainControllers instantiates a new AdAuthServiceDomainControllers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdAuthServiceDomainControllersWithDefaults

`func NewAdAuthServiceDomainControllersWithDefaults() *AdAuthServiceDomainControllers`

NewAdAuthServiceDomainControllersWithDefaults instantiates a new AdAuthServiceDomainControllers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdnOrIp

`func (o *AdAuthServiceDomainControllers) GetFqdnOrIp() string`

GetFqdnOrIp returns the FqdnOrIp field if non-nil, zero value otherwise.

### GetFqdnOrIpOk

`func (o *AdAuthServiceDomainControllers) GetFqdnOrIpOk() (*string, bool)`

GetFqdnOrIpOk returns a tuple with the FqdnOrIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnOrIp

`func (o *AdAuthServiceDomainControllers) SetFqdnOrIp(v string)`

SetFqdnOrIp sets FqdnOrIp field to given value.

### HasFqdnOrIp

`func (o *AdAuthServiceDomainControllers) HasFqdnOrIp() bool`

HasFqdnOrIp returns a boolean if a field has been set.

### GetAuthPort

`func (o *AdAuthServiceDomainControllers) GetAuthPort() int64`

GetAuthPort returns the AuthPort field if non-nil, zero value otherwise.

### GetAuthPortOk

`func (o *AdAuthServiceDomainControllers) GetAuthPortOk() (*int64, bool)`

GetAuthPortOk returns a tuple with the AuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPort

`func (o *AdAuthServiceDomainControllers) SetAuthPort(v int64)`

SetAuthPort sets AuthPort field to given value.

### HasAuthPort

`func (o *AdAuthServiceDomainControllers) HasAuthPort() bool`

HasAuthPort returns a boolean if a field has been set.

### GetComment

`func (o *AdAuthServiceDomainControllers) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AdAuthServiceDomainControllers) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AdAuthServiceDomainControllers) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AdAuthServiceDomainControllers) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *AdAuthServiceDomainControllers) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AdAuthServiceDomainControllers) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AdAuthServiceDomainControllers) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AdAuthServiceDomainControllers) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEncryption

`func (o *AdAuthServiceDomainControllers) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *AdAuthServiceDomainControllers) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *AdAuthServiceDomainControllers) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *AdAuthServiceDomainControllers) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetMgmtPort

`func (o *AdAuthServiceDomainControllers) GetMgmtPort() bool`

GetMgmtPort returns the MgmtPort field if non-nil, zero value otherwise.

### GetMgmtPortOk

`func (o *AdAuthServiceDomainControllers) GetMgmtPortOk() (*bool, bool)`

GetMgmtPortOk returns a tuple with the MgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPort

`func (o *AdAuthServiceDomainControllers) SetMgmtPort(v bool)`

SetMgmtPort sets MgmtPort field to given value.

### HasMgmtPort

`func (o *AdAuthServiceDomainControllers) HasMgmtPort() bool`

HasMgmtPort returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *AdAuthServiceDomainControllers) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *AdAuthServiceDomainControllers) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *AdAuthServiceDomainControllers) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *AdAuthServiceDomainControllers) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


