# TacacsplusAuthservice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AcctRetries** | Pointer to **int64** | The number of the accounting retries before giving up and moving on to the next server. | [optional] 
**AcctTimeout** | Pointer to **int64** | The accounting retry period in milliseconds. | [optional] 
**AuthRetries** | Pointer to **int64** | The number of the authentication/authorization retries before giving up and moving on to the next server. | [optional] 
**AuthTimeout** | Pointer to **int64** | The authentication/authorization timeout period in milliseconds. | [optional] 
**Comment** | Pointer to **string** | The TACACS+ authentication service descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the TACACS+ authentication service object is disabled. | [optional] 
**Name** | Pointer to **string** | The TACACS+ authentication service name. | [optional] 
**Servers** | Pointer to [**[]TacacsplusAuthserviceServers**](TacacsplusAuthserviceServers.md) | The list of the TACACS+ servers used for authentication. | [optional] 

## Methods

### NewTacacsplusAuthservice

`func NewTacacsplusAuthservice() *TacacsplusAuthservice`

NewTacacsplusAuthservice instantiates a new TacacsplusAuthservice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsplusAuthserviceWithDefaults

`func NewTacacsplusAuthserviceWithDefaults() *TacacsplusAuthservice`

NewTacacsplusAuthserviceWithDefaults instantiates a new TacacsplusAuthservice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *TacacsplusAuthservice) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *TacacsplusAuthservice) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *TacacsplusAuthservice) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *TacacsplusAuthservice) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAcctRetries

`func (o *TacacsplusAuthservice) GetAcctRetries() int64`

GetAcctRetries returns the AcctRetries field if non-nil, zero value otherwise.

### GetAcctRetriesOk

`func (o *TacacsplusAuthservice) GetAcctRetriesOk() (*int64, bool)`

GetAcctRetriesOk returns a tuple with the AcctRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctRetries

`func (o *TacacsplusAuthservice) SetAcctRetries(v int64)`

SetAcctRetries sets AcctRetries field to given value.

### HasAcctRetries

`func (o *TacacsplusAuthservice) HasAcctRetries() bool`

HasAcctRetries returns a boolean if a field has been set.

### GetAcctTimeout

`func (o *TacacsplusAuthservice) GetAcctTimeout() int64`

GetAcctTimeout returns the AcctTimeout field if non-nil, zero value otherwise.

### GetAcctTimeoutOk

`func (o *TacacsplusAuthservice) GetAcctTimeoutOk() (*int64, bool)`

GetAcctTimeoutOk returns a tuple with the AcctTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctTimeout

`func (o *TacacsplusAuthservice) SetAcctTimeout(v int64)`

SetAcctTimeout sets AcctTimeout field to given value.

### HasAcctTimeout

`func (o *TacacsplusAuthservice) HasAcctTimeout() bool`

HasAcctTimeout returns a boolean if a field has been set.

### GetAuthRetries

`func (o *TacacsplusAuthservice) GetAuthRetries() int64`

GetAuthRetries returns the AuthRetries field if non-nil, zero value otherwise.

### GetAuthRetriesOk

`func (o *TacacsplusAuthservice) GetAuthRetriesOk() (*int64, bool)`

GetAuthRetriesOk returns a tuple with the AuthRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRetries

`func (o *TacacsplusAuthservice) SetAuthRetries(v int64)`

SetAuthRetries sets AuthRetries field to given value.

### HasAuthRetries

`func (o *TacacsplusAuthservice) HasAuthRetries() bool`

HasAuthRetries returns a boolean if a field has been set.

### GetAuthTimeout

`func (o *TacacsplusAuthservice) GetAuthTimeout() int64`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *TacacsplusAuthservice) GetAuthTimeoutOk() (*int64, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *TacacsplusAuthservice) SetAuthTimeout(v int64)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *TacacsplusAuthservice) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetComment

`func (o *TacacsplusAuthservice) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TacacsplusAuthservice) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TacacsplusAuthservice) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TacacsplusAuthservice) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *TacacsplusAuthservice) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *TacacsplusAuthservice) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *TacacsplusAuthservice) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *TacacsplusAuthservice) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetName

`func (o *TacacsplusAuthservice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TacacsplusAuthservice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TacacsplusAuthservice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TacacsplusAuthservice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServers

`func (o *TacacsplusAuthservice) GetServers() []TacacsplusAuthserviceServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *TacacsplusAuthservice) GetServersOk() (*[]TacacsplusAuthserviceServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *TacacsplusAuthservice) SetServers(v []TacacsplusAuthserviceServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *TacacsplusAuthservice) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


