# GetGridServicerestartRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Error** | Pointer to **string** | The error message if restart has failed. | [optional] [readonly] 
**Forced** | Pointer to **bool** | Indicates if this is a force restart. | [optional] [readonly] 
**Group** | Pointer to **string** | The name of the Restart Group associated with the request. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **int64** | The timestamp when the status of the request has changed. | [optional] [readonly] 
**Member** | Pointer to **string** | The member to restart. | [optional] [readonly] 
**Needed** | Pointer to **string** | Indicates if restart is needed. | [optional] [readonly] 
**Order** | Pointer to **int64** | The order to restart. | [optional] [readonly] 
**Result** | Pointer to [**GridServicerestartRequest**](GridServicerestartRequest.md) |  | [optional] 
**Service** | Pointer to **string** | The service to restart. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the request. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewGetGridServicerestartRequestResponse

`func NewGetGridServicerestartRequestResponse() *GetGridServicerestartRequestResponse`

NewGetGridServicerestartRequestResponse instantiates a new GetGridServicerestartRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridServicerestartRequestResponseWithDefaults

`func NewGetGridServicerestartRequestResponseWithDefaults() *GetGridServicerestartRequestResponse`

NewGetGridServicerestartRequestResponseWithDefaults instantiates a new GetGridServicerestartRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridServicerestartRequestResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridServicerestartRequestResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridServicerestartRequestResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridServicerestartRequestResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetError

`func (o *GetGridServicerestartRequestResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetGridServicerestartRequestResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetGridServicerestartRequestResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetGridServicerestartRequestResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetForced

`func (o *GetGridServicerestartRequestResponse) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *GetGridServicerestartRequestResponse) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *GetGridServicerestartRequestResponse) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *GetGridServicerestartRequestResponse) HasForced() bool`

HasForced returns a boolean if a field has been set.

### GetGroup

`func (o *GetGridServicerestartRequestResponse) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetGridServicerestartRequestResponse) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetGridServicerestartRequestResponse) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetGridServicerestartRequestResponse) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *GetGridServicerestartRequestResponse) GetLastUpdatedTime() int64`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GetGridServicerestartRequestResponse) GetLastUpdatedTimeOk() (*int64, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GetGridServicerestartRequestResponse) SetLastUpdatedTime(v int64)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *GetGridServicerestartRequestResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMember

`func (o *GetGridServicerestartRequestResponse) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetGridServicerestartRequestResponse) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetGridServicerestartRequestResponse) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetGridServicerestartRequestResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNeeded

`func (o *GetGridServicerestartRequestResponse) GetNeeded() string`

GetNeeded returns the Needed field if non-nil, zero value otherwise.

### GetNeededOk

`func (o *GetGridServicerestartRequestResponse) GetNeededOk() (*string, bool)`

GetNeededOk returns a tuple with the Needed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeeded

`func (o *GetGridServicerestartRequestResponse) SetNeeded(v string)`

SetNeeded sets Needed field to given value.

### HasNeeded

`func (o *GetGridServicerestartRequestResponse) HasNeeded() bool`

HasNeeded returns a boolean if a field has been set.

### GetOrder

`func (o *GetGridServicerestartRequestResponse) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetGridServicerestartRequestResponse) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetGridServicerestartRequestResponse) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetGridServicerestartRequestResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetResult

`func (o *GetGridServicerestartRequestResponse) GetResult() GridServicerestartRequest`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridServicerestartRequestResponse) GetResultOk() (*GridServicerestartRequest, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridServicerestartRequestResponse) SetResult(v GridServicerestartRequest)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridServicerestartRequestResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetService

`func (o *GetGridServicerestartRequestResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetGridServicerestartRequestResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetGridServicerestartRequestResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetGridServicerestartRequestResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetState

`func (o *GetGridServicerestartRequestResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetGridServicerestartRequestResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetGridServicerestartRequestResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetGridServicerestartRequestResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridServicerestartRequestResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridServicerestartRequestResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridServicerestartRequestResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridServicerestartRequestResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


