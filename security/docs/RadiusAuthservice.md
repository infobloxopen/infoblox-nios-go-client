# RadiusAuthservice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AcctRetries** | Pointer to **int64** | The number of times to attempt to contact an accounting RADIUS server. | [optional] 
**AcctTimeout** | Pointer to **int64** | The number of seconds to wait for a response from the RADIUS server. | [optional] 
**AuthRetries** | Pointer to **int64** | The number of times to attempt to contact an authentication RADIUS server. | [optional] 
**AuthTimeout** | Pointer to **int64** | The number of seconds to wait for a response from the RADIUS server. | [optional] 
**CacheTtl** | Pointer to **int64** | The TTL of cached authentication data in seconds. | [optional] 
**Comment** | Pointer to **string** | The RADIUS descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the RADIUS authentication service is disabled. | [optional] 
**EnableCache** | Pointer to **bool** | Determines whether the authentication cache is enabled. | [optional] 
**Mode** | Pointer to **string** | The way to contact the RADIUS server. | [optional] 
**Name** | Pointer to **string** | The RADIUS authentication service name. | [optional] 
**RecoveryInterval** | Pointer to **int64** | The time period to wait before retrying a server that has been marked as down. | [optional] 
**Servers** | Pointer to [**[]RadiusAuthserviceServers**](RadiusAuthserviceServers.md) | The ordered list of RADIUS authentication servers. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewRadiusAuthservice

`func NewRadiusAuthservice() *RadiusAuthservice`

NewRadiusAuthservice instantiates a new RadiusAuthservice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAuthserviceWithDefaults

`func NewRadiusAuthserviceWithDefaults() *RadiusAuthservice`

NewRadiusAuthserviceWithDefaults instantiates a new RadiusAuthservice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RadiusAuthservice) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RadiusAuthservice) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RadiusAuthservice) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RadiusAuthservice) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAcctRetries

`func (o *RadiusAuthservice) GetAcctRetries() int64`

GetAcctRetries returns the AcctRetries field if non-nil, zero value otherwise.

### GetAcctRetriesOk

`func (o *RadiusAuthservice) GetAcctRetriesOk() (*int64, bool)`

GetAcctRetriesOk returns a tuple with the AcctRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctRetries

`func (o *RadiusAuthservice) SetAcctRetries(v int64)`

SetAcctRetries sets AcctRetries field to given value.

### HasAcctRetries

`func (o *RadiusAuthservice) HasAcctRetries() bool`

HasAcctRetries returns a boolean if a field has been set.

### GetAcctTimeout

`func (o *RadiusAuthservice) GetAcctTimeout() int64`

GetAcctTimeout returns the AcctTimeout field if non-nil, zero value otherwise.

### GetAcctTimeoutOk

`func (o *RadiusAuthservice) GetAcctTimeoutOk() (*int64, bool)`

GetAcctTimeoutOk returns a tuple with the AcctTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctTimeout

`func (o *RadiusAuthservice) SetAcctTimeout(v int64)`

SetAcctTimeout sets AcctTimeout field to given value.

### HasAcctTimeout

`func (o *RadiusAuthservice) HasAcctTimeout() bool`

HasAcctTimeout returns a boolean if a field has been set.

### GetAuthRetries

`func (o *RadiusAuthservice) GetAuthRetries() int64`

GetAuthRetries returns the AuthRetries field if non-nil, zero value otherwise.

### GetAuthRetriesOk

`func (o *RadiusAuthservice) GetAuthRetriesOk() (*int64, bool)`

GetAuthRetriesOk returns a tuple with the AuthRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRetries

`func (o *RadiusAuthservice) SetAuthRetries(v int64)`

SetAuthRetries sets AuthRetries field to given value.

### HasAuthRetries

`func (o *RadiusAuthservice) HasAuthRetries() bool`

HasAuthRetries returns a boolean if a field has been set.

### GetAuthTimeout

`func (o *RadiusAuthservice) GetAuthTimeout() int64`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *RadiusAuthservice) GetAuthTimeoutOk() (*int64, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *RadiusAuthservice) SetAuthTimeout(v int64)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *RadiusAuthservice) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetCacheTtl

`func (o *RadiusAuthservice) GetCacheTtl() int64`

GetCacheTtl returns the CacheTtl field if non-nil, zero value otherwise.

### GetCacheTtlOk

`func (o *RadiusAuthservice) GetCacheTtlOk() (*int64, bool)`

GetCacheTtlOk returns a tuple with the CacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtl

`func (o *RadiusAuthservice) SetCacheTtl(v int64)`

SetCacheTtl sets CacheTtl field to given value.

### HasCacheTtl

`func (o *RadiusAuthservice) HasCacheTtl() bool`

HasCacheTtl returns a boolean if a field has been set.

### GetComment

`func (o *RadiusAuthservice) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RadiusAuthservice) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RadiusAuthservice) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RadiusAuthservice) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RadiusAuthservice) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RadiusAuthservice) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RadiusAuthservice) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RadiusAuthservice) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableCache

`func (o *RadiusAuthservice) GetEnableCache() bool`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *RadiusAuthservice) GetEnableCacheOk() (*bool, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *RadiusAuthservice) SetEnableCache(v bool)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *RadiusAuthservice) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetMode

`func (o *RadiusAuthservice) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RadiusAuthservice) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RadiusAuthservice) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RadiusAuthservice) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *RadiusAuthservice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusAuthservice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusAuthservice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RadiusAuthservice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecoveryInterval

`func (o *RadiusAuthservice) GetRecoveryInterval() int64`

GetRecoveryInterval returns the RecoveryInterval field if non-nil, zero value otherwise.

### GetRecoveryIntervalOk

`func (o *RadiusAuthservice) GetRecoveryIntervalOk() (*int64, bool)`

GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInterval

`func (o *RadiusAuthservice) SetRecoveryInterval(v int64)`

SetRecoveryInterval sets RecoveryInterval field to given value.

### HasRecoveryInterval

`func (o *RadiusAuthservice) HasRecoveryInterval() bool`

HasRecoveryInterval returns a boolean if a field has been set.

### GetServers

`func (o *RadiusAuthservice) GetServers() []RadiusAuthserviceServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RadiusAuthservice) GetServersOk() (*[]RadiusAuthserviceServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RadiusAuthservice) SetServers(v []RadiusAuthserviceServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *RadiusAuthservice) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetUuid

`func (o *RadiusAuthservice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RadiusAuthservice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RadiusAuthservice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RadiusAuthservice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


