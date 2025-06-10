# GetScavengingtaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Action** | Pointer to **string** | The scavenging action. | [optional] [readonly] 
**AssociatedObject** | Pointer to **string** | The reference to the object associated with the scavenging task. | [optional] [readonly] 
**EndTime** | Pointer to **int64** | The scavenging process end time. | [optional] [readonly] 
**ProcessedRecords** | Pointer to **int64** | The number of processed during scavenging resource records. | [optional] [readonly] 
**ReclaimableRecords** | Pointer to **int64** | The number of resource records that are allowed to be reclaimed during the scavenging process. | [optional] [readonly] 
**ReclaimedRecords** | Pointer to **int64** | The number of reclaimed during the scavenging process resource records. | [optional] [readonly] 
**StartTime** | Pointer to **int64** | The scavenging process start time. | [optional] [readonly] 
**Status** | Pointer to **string** | The scavenging process status. This is a read-only attribute. | [optional] [readonly] 
**Result** | Pointer to [**Scavengingtask**](Scavengingtask.md) |  | [optional] 

## Methods

### NewGetScavengingtaskResponse

`func NewGetScavengingtaskResponse() *GetScavengingtaskResponse`

NewGetScavengingtaskResponse instantiates a new GetScavengingtaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetScavengingtaskResponseWithDefaults

`func NewGetScavengingtaskResponseWithDefaults() *GetScavengingtaskResponse`

NewGetScavengingtaskResponseWithDefaults instantiates a new GetScavengingtaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetScavengingtaskResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetScavengingtaskResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetScavengingtaskResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetScavengingtaskResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAction

`func (o *GetScavengingtaskResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetScavengingtaskResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetScavengingtaskResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetScavengingtaskResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *GetScavengingtaskResponse) GetAssociatedObject() string`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *GetScavengingtaskResponse) GetAssociatedObjectOk() (*string, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *GetScavengingtaskResponse) SetAssociatedObject(v string)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *GetScavengingtaskResponse) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetEndTime

`func (o *GetScavengingtaskResponse) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetScavengingtaskResponse) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetScavengingtaskResponse) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetScavengingtaskResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetProcessedRecords

`func (o *GetScavengingtaskResponse) GetProcessedRecords() int64`

GetProcessedRecords returns the ProcessedRecords field if non-nil, zero value otherwise.

### GetProcessedRecordsOk

`func (o *GetScavengingtaskResponse) GetProcessedRecordsOk() (*int64, bool)`

GetProcessedRecordsOk returns a tuple with the ProcessedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRecords

`func (o *GetScavengingtaskResponse) SetProcessedRecords(v int64)`

SetProcessedRecords sets ProcessedRecords field to given value.

### HasProcessedRecords

`func (o *GetScavengingtaskResponse) HasProcessedRecords() bool`

HasProcessedRecords returns a boolean if a field has been set.

### GetReclaimableRecords

`func (o *GetScavengingtaskResponse) GetReclaimableRecords() int64`

GetReclaimableRecords returns the ReclaimableRecords field if non-nil, zero value otherwise.

### GetReclaimableRecordsOk

`func (o *GetScavengingtaskResponse) GetReclaimableRecordsOk() (*int64, bool)`

GetReclaimableRecordsOk returns a tuple with the ReclaimableRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimableRecords

`func (o *GetScavengingtaskResponse) SetReclaimableRecords(v int64)`

SetReclaimableRecords sets ReclaimableRecords field to given value.

### HasReclaimableRecords

`func (o *GetScavengingtaskResponse) HasReclaimableRecords() bool`

HasReclaimableRecords returns a boolean if a field has been set.

### GetReclaimedRecords

`func (o *GetScavengingtaskResponse) GetReclaimedRecords() int64`

GetReclaimedRecords returns the ReclaimedRecords field if non-nil, zero value otherwise.

### GetReclaimedRecordsOk

`func (o *GetScavengingtaskResponse) GetReclaimedRecordsOk() (*int64, bool)`

GetReclaimedRecordsOk returns a tuple with the ReclaimedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimedRecords

`func (o *GetScavengingtaskResponse) SetReclaimedRecords(v int64)`

SetReclaimedRecords sets ReclaimedRecords field to given value.

### HasReclaimedRecords

`func (o *GetScavengingtaskResponse) HasReclaimedRecords() bool`

HasReclaimedRecords returns a boolean if a field has been set.

### GetStartTime

`func (o *GetScavengingtaskResponse) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetScavengingtaskResponse) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetScavengingtaskResponse) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetScavengingtaskResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *GetScavengingtaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetScavengingtaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetScavengingtaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetScavengingtaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetScavengingtaskResponse) GetResult() Scavengingtask`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetScavengingtaskResponse) GetResultOk() (*Scavengingtask, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetScavengingtaskResponse) SetResult(v Scavengingtask)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetScavengingtaskResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


