# GridServicerestartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Error** | Pointer to **string** | The error message if restart has failed. | [optional] [readonly] 
**Forced** | Pointer to **bool** | Indicates if this is a force restart. | [optional] [readonly] 
**Group** | Pointer to **string** | The name of the Restart Group associated with the request. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **int64** | The timestamp when the status of the request has changed. | [optional] [readonly] 
**Member** | Pointer to **string** | The member to restart. | [optional] [readonly] 
**Needed** | Pointer to **string** | Indicates if restart is needed. | [optional] [readonly] 
**Order** | Pointer to **int32** | The order to restart. | [optional] [readonly] 
**Result** | Pointer to **string** | The result of the restart operation. | [optional] [readonly] 
**Service** | Pointer to **string** | The service to restart. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the request. | [optional] [readonly] 

## Methods

### NewGridServicerestartRequest

`func NewGridServicerestartRequest() *GridServicerestartRequest`

NewGridServicerestartRequest instantiates a new GridServicerestartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridServicerestartRequestWithDefaults

`func NewGridServicerestartRequestWithDefaults() *GridServicerestartRequest`

NewGridServicerestartRequestWithDefaults instantiates a new GridServicerestartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridServicerestartRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridServicerestartRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridServicerestartRequest) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridServicerestartRequest) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetError

`func (o *GridServicerestartRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GridServicerestartRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GridServicerestartRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GridServicerestartRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetForced

`func (o *GridServicerestartRequest) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *GridServicerestartRequest) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *GridServicerestartRequest) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *GridServicerestartRequest) HasForced() bool`

HasForced returns a boolean if a field has been set.

### GetGroup

`func (o *GridServicerestartRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GridServicerestartRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GridServicerestartRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GridServicerestartRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *GridServicerestartRequest) GetLastUpdatedTime() int64`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GridServicerestartRequest) GetLastUpdatedTimeOk() (*int64, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GridServicerestartRequest) SetLastUpdatedTime(v int64)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *GridServicerestartRequest) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMember

`func (o *GridServicerestartRequest) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GridServicerestartRequest) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GridServicerestartRequest) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GridServicerestartRequest) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNeeded

`func (o *GridServicerestartRequest) GetNeeded() string`

GetNeeded returns the Needed field if non-nil, zero value otherwise.

### GetNeededOk

`func (o *GridServicerestartRequest) GetNeededOk() (*string, bool)`

GetNeededOk returns a tuple with the Needed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeeded

`func (o *GridServicerestartRequest) SetNeeded(v string)`

SetNeeded sets Needed field to given value.

### HasNeeded

`func (o *GridServicerestartRequest) HasNeeded() bool`

HasNeeded returns a boolean if a field has been set.

### GetOrder

`func (o *GridServicerestartRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GridServicerestartRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GridServicerestartRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GridServicerestartRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetResult

`func (o *GridServicerestartRequest) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GridServicerestartRequest) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GridServicerestartRequest) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *GridServicerestartRequest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetService

`func (o *GridServicerestartRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GridServicerestartRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GridServicerestartRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GridServicerestartRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetState

`func (o *GridServicerestartRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GridServicerestartRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GridServicerestartRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GridServicerestartRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


