# GetGridServicerestartStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Failures** | Pointer to **int64** | The number of failed requests. | [optional] [readonly] 
**Finished** | Pointer to **int64** | The number of finished requests. | [optional] [readonly] 
**Grouped** | Pointer to **string** | The type of grouping. | [optional] [readonly] 
**NeededRestart** | Pointer to **int64** | The number of created yet unprocessed requests for restart. | [optional] [readonly] 
**NoRestart** | Pointer to **int64** | The number of requests that did not require a restart. | [optional] [readonly] 
**Parent** | Pointer to **string** | A reference to the grid or grid:servicerestart:group object associated with the request. | [optional] [readonly] 
**Pending** | Pointer to **int64** | The number of requests that are pending a restart. | [optional] [readonly] 
**PendingRestart** | Pointer to **int64** | The number of forced or needed requests pending for restart. | [optional] [readonly] 
**Processing** | Pointer to **int64** | The number of not forced and not needed requests pending for restart. | [optional] [readonly] 
**Restarting** | Pointer to **int64** | The number of service restarts for corresponding members. | [optional] [readonly] 
**Success** | Pointer to **int64** | The number of requests associated with successful restarts. | [optional] [readonly] 
**Timeouts** | Pointer to **int64** | The number of timeout requests. | [optional] [readonly] 
**Result** | Pointer to [**GridServicerestartStatus**](GridServicerestartStatus.md) |  | [optional] 

## Methods

### NewGetGridServicerestartStatusResponse

`func NewGetGridServicerestartStatusResponse() *GetGridServicerestartStatusResponse`

NewGetGridServicerestartStatusResponse instantiates a new GetGridServicerestartStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridServicerestartStatusResponseWithDefaults

`func NewGetGridServicerestartStatusResponseWithDefaults() *GetGridServicerestartStatusResponse`

NewGetGridServicerestartStatusResponseWithDefaults instantiates a new GetGridServicerestartStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridServicerestartStatusResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridServicerestartStatusResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridServicerestartStatusResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridServicerestartStatusResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetFailures

`func (o *GetGridServicerestartStatusResponse) GetFailures() int64`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *GetGridServicerestartStatusResponse) GetFailuresOk() (*int64, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *GetGridServicerestartStatusResponse) SetFailures(v int64)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *GetGridServicerestartStatusResponse) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetFinished

`func (o *GetGridServicerestartStatusResponse) GetFinished() int64`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetGridServicerestartStatusResponse) GetFinishedOk() (*int64, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetGridServicerestartStatusResponse) SetFinished(v int64)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetGridServicerestartStatusResponse) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetGrouped

`func (o *GetGridServicerestartStatusResponse) GetGrouped() string`

GetGrouped returns the Grouped field if non-nil, zero value otherwise.

### GetGroupedOk

`func (o *GetGridServicerestartStatusResponse) GetGroupedOk() (*string, bool)`

GetGroupedOk returns a tuple with the Grouped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouped

`func (o *GetGridServicerestartStatusResponse) SetGrouped(v string)`

SetGrouped sets Grouped field to given value.

### HasGrouped

`func (o *GetGridServicerestartStatusResponse) HasGrouped() bool`

HasGrouped returns a boolean if a field has been set.

### GetNeededRestart

`func (o *GetGridServicerestartStatusResponse) GetNeededRestart() int64`

GetNeededRestart returns the NeededRestart field if non-nil, zero value otherwise.

### GetNeededRestartOk

`func (o *GetGridServicerestartStatusResponse) GetNeededRestartOk() (*int64, bool)`

GetNeededRestartOk returns a tuple with the NeededRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededRestart

`func (o *GetGridServicerestartStatusResponse) SetNeededRestart(v int64)`

SetNeededRestart sets NeededRestart field to given value.

### HasNeededRestart

`func (o *GetGridServicerestartStatusResponse) HasNeededRestart() bool`

HasNeededRestart returns a boolean if a field has been set.

### GetNoRestart

`func (o *GetGridServicerestartStatusResponse) GetNoRestart() int64`

GetNoRestart returns the NoRestart field if non-nil, zero value otherwise.

### GetNoRestartOk

`func (o *GetGridServicerestartStatusResponse) GetNoRestartOk() (*int64, bool)`

GetNoRestartOk returns a tuple with the NoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRestart

`func (o *GetGridServicerestartStatusResponse) SetNoRestart(v int64)`

SetNoRestart sets NoRestart field to given value.

### HasNoRestart

`func (o *GetGridServicerestartStatusResponse) HasNoRestart() bool`

HasNoRestart returns a boolean if a field has been set.

### GetParent

`func (o *GetGridServicerestartStatusResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetGridServicerestartStatusResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetGridServicerestartStatusResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetGridServicerestartStatusResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPending

`func (o *GetGridServicerestartStatusResponse) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GetGridServicerestartStatusResponse) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GetGridServicerestartStatusResponse) SetPending(v int64)`

SetPending sets Pending field to given value.

### HasPending

`func (o *GetGridServicerestartStatusResponse) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetPendingRestart

`func (o *GetGridServicerestartStatusResponse) GetPendingRestart() int64`

GetPendingRestart returns the PendingRestart field if non-nil, zero value otherwise.

### GetPendingRestartOk

`func (o *GetGridServicerestartStatusResponse) GetPendingRestartOk() (*int64, bool)`

GetPendingRestartOk returns a tuple with the PendingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRestart

`func (o *GetGridServicerestartStatusResponse) SetPendingRestart(v int64)`

SetPendingRestart sets PendingRestart field to given value.

### HasPendingRestart

`func (o *GetGridServicerestartStatusResponse) HasPendingRestart() bool`

HasPendingRestart returns a boolean if a field has been set.

### GetProcessing

`func (o *GetGridServicerestartStatusResponse) GetProcessing() int64`

GetProcessing returns the Processing field if non-nil, zero value otherwise.

### GetProcessingOk

`func (o *GetGridServicerestartStatusResponse) GetProcessingOk() (*int64, bool)`

GetProcessingOk returns a tuple with the Processing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessing

`func (o *GetGridServicerestartStatusResponse) SetProcessing(v int64)`

SetProcessing sets Processing field to given value.

### HasProcessing

`func (o *GetGridServicerestartStatusResponse) HasProcessing() bool`

HasProcessing returns a boolean if a field has been set.

### GetRestarting

`func (o *GetGridServicerestartStatusResponse) GetRestarting() int64`

GetRestarting returns the Restarting field if non-nil, zero value otherwise.

### GetRestartingOk

`func (o *GetGridServicerestartStatusResponse) GetRestartingOk() (*int64, bool)`

GetRestartingOk returns a tuple with the Restarting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarting

`func (o *GetGridServicerestartStatusResponse) SetRestarting(v int64)`

SetRestarting sets Restarting field to given value.

### HasRestarting

`func (o *GetGridServicerestartStatusResponse) HasRestarting() bool`

HasRestarting returns a boolean if a field has been set.

### GetSuccess

`func (o *GetGridServicerestartStatusResponse) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGridServicerestartStatusResponse) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGridServicerestartStatusResponse) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGridServicerestartStatusResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTimeouts

`func (o *GetGridServicerestartStatusResponse) GetTimeouts() int64`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *GetGridServicerestartStatusResponse) GetTimeoutsOk() (*int64, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *GetGridServicerestartStatusResponse) SetTimeouts(v int64)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *GetGridServicerestartStatusResponse) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetResult

`func (o *GetGridServicerestartStatusResponse) GetResult() GridServicerestartStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridServicerestartStatusResponse) GetResultOk() (*GridServicerestartStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridServicerestartStatusResponse) SetResult(v GridServicerestartStatus)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridServicerestartStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


