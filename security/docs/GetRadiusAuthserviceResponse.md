# GetRadiusAuthserviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Result** | Pointer to [**RadiusAuthservice**](RadiusAuthservice.md) |  | [optional] 

## Methods

### NewGetRadiusAuthserviceResponse

`func NewGetRadiusAuthserviceResponse() *GetRadiusAuthserviceResponse`

NewGetRadiusAuthserviceResponse instantiates a new GetRadiusAuthserviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRadiusAuthserviceResponseWithDefaults

`func NewGetRadiusAuthserviceResponseWithDefaults() *GetRadiusAuthserviceResponse`

NewGetRadiusAuthserviceResponseWithDefaults instantiates a new GetRadiusAuthserviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRadiusAuthserviceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRadiusAuthserviceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRadiusAuthserviceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRadiusAuthserviceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAcctRetries

`func (o *GetRadiusAuthserviceResponse) GetAcctRetries() int64`

GetAcctRetries returns the AcctRetries field if non-nil, zero value otherwise.

### GetAcctRetriesOk

`func (o *GetRadiusAuthserviceResponse) GetAcctRetriesOk() (*int64, bool)`

GetAcctRetriesOk returns a tuple with the AcctRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctRetries

`func (o *GetRadiusAuthserviceResponse) SetAcctRetries(v int64)`

SetAcctRetries sets AcctRetries field to given value.

### HasAcctRetries

`func (o *GetRadiusAuthserviceResponse) HasAcctRetries() bool`

HasAcctRetries returns a boolean if a field has been set.

### GetAcctTimeout

`func (o *GetRadiusAuthserviceResponse) GetAcctTimeout() int64`

GetAcctTimeout returns the AcctTimeout field if non-nil, zero value otherwise.

### GetAcctTimeoutOk

`func (o *GetRadiusAuthserviceResponse) GetAcctTimeoutOk() (*int64, bool)`

GetAcctTimeoutOk returns a tuple with the AcctTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctTimeout

`func (o *GetRadiusAuthserviceResponse) SetAcctTimeout(v int64)`

SetAcctTimeout sets AcctTimeout field to given value.

### HasAcctTimeout

`func (o *GetRadiusAuthserviceResponse) HasAcctTimeout() bool`

HasAcctTimeout returns a boolean if a field has been set.

### GetAuthRetries

`func (o *GetRadiusAuthserviceResponse) GetAuthRetries() int64`

GetAuthRetries returns the AuthRetries field if non-nil, zero value otherwise.

### GetAuthRetriesOk

`func (o *GetRadiusAuthserviceResponse) GetAuthRetriesOk() (*int64, bool)`

GetAuthRetriesOk returns a tuple with the AuthRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRetries

`func (o *GetRadiusAuthserviceResponse) SetAuthRetries(v int64)`

SetAuthRetries sets AuthRetries field to given value.

### HasAuthRetries

`func (o *GetRadiusAuthserviceResponse) HasAuthRetries() bool`

HasAuthRetries returns a boolean if a field has been set.

### GetAuthTimeout

`func (o *GetRadiusAuthserviceResponse) GetAuthTimeout() int64`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *GetRadiusAuthserviceResponse) GetAuthTimeoutOk() (*int64, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *GetRadiusAuthserviceResponse) SetAuthTimeout(v int64)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *GetRadiusAuthserviceResponse) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetCacheTtl

`func (o *GetRadiusAuthserviceResponse) GetCacheTtl() int64`

GetCacheTtl returns the CacheTtl field if non-nil, zero value otherwise.

### GetCacheTtlOk

`func (o *GetRadiusAuthserviceResponse) GetCacheTtlOk() (*int64, bool)`

GetCacheTtlOk returns a tuple with the CacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtl

`func (o *GetRadiusAuthserviceResponse) SetCacheTtl(v int64)`

SetCacheTtl sets CacheTtl field to given value.

### HasCacheTtl

`func (o *GetRadiusAuthserviceResponse) HasCacheTtl() bool`

HasCacheTtl returns a boolean if a field has been set.

### GetComment

`func (o *GetRadiusAuthserviceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRadiusAuthserviceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRadiusAuthserviceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRadiusAuthserviceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetRadiusAuthserviceResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRadiusAuthserviceResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRadiusAuthserviceResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRadiusAuthserviceResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableCache

`func (o *GetRadiusAuthserviceResponse) GetEnableCache() bool`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *GetRadiusAuthserviceResponse) GetEnableCacheOk() (*bool, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *GetRadiusAuthserviceResponse) SetEnableCache(v bool)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *GetRadiusAuthserviceResponse) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetMode

`func (o *GetRadiusAuthserviceResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetRadiusAuthserviceResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetRadiusAuthserviceResponse) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetRadiusAuthserviceResponse) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GetRadiusAuthserviceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRadiusAuthserviceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRadiusAuthserviceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRadiusAuthserviceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecoveryInterval

`func (o *GetRadiusAuthserviceResponse) GetRecoveryInterval() int64`

GetRecoveryInterval returns the RecoveryInterval field if non-nil, zero value otherwise.

### GetRecoveryIntervalOk

`func (o *GetRadiusAuthserviceResponse) GetRecoveryIntervalOk() (*int64, bool)`

GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInterval

`func (o *GetRadiusAuthserviceResponse) SetRecoveryInterval(v int64)`

SetRecoveryInterval sets RecoveryInterval field to given value.

### HasRecoveryInterval

`func (o *GetRadiusAuthserviceResponse) HasRecoveryInterval() bool`

HasRecoveryInterval returns a boolean if a field has been set.

### GetServers

`func (o *GetRadiusAuthserviceResponse) GetServers() []RadiusAuthserviceServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetRadiusAuthserviceResponse) GetServersOk() (*[]RadiusAuthserviceServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetRadiusAuthserviceResponse) SetServers(v []RadiusAuthserviceServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetRadiusAuthserviceResponse) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetResult

`func (o *GetRadiusAuthserviceResponse) GetResult() RadiusAuthservice`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRadiusAuthserviceResponse) GetResultOk() (*RadiusAuthservice, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRadiusAuthserviceResponse) SetResult(v RadiusAuthservice)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRadiusAuthserviceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


