# Scavengingtask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Action** | Pointer to **string** | The scavenging action. | [optional] [readonly] 
**AssociatedObject** | Pointer to **string** | The reference to the object associated with the scavenging task. | [optional] [readonly] 
**EndTime** | Pointer to **int64** | The scavenging process end time. | [optional] [readonly] 
**ProcessedRecords** | Pointer to **int64** | The number of processed during scavenging resource records. | [optional] [readonly] 
**ReclaimableRecords** | Pointer to **int64** | The number of resource records that are allowed to be reclaimed during the scavenging process. | [optional] [readonly] 
**ReclaimedRecords** | Pointer to **int64** | The number of reclaimed during the scavenging process resource records. | [optional] [readonly] 
**StartTime** | Pointer to **int64** | The scavenging process start time. | [optional] [readonly] 
**Status** | Pointer to **string** | The scavenging process status. This is a read-only attribute. | [optional] [readonly] 

## Methods

### NewScavengingtask

`func NewScavengingtask() *Scavengingtask`

NewScavengingtask instantiates a new Scavengingtask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScavengingtaskWithDefaults

`func NewScavengingtaskWithDefaults() *Scavengingtask`

NewScavengingtaskWithDefaults instantiates a new Scavengingtask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Scavengingtask) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Scavengingtask) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Scavengingtask) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Scavengingtask) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAction

`func (o *Scavengingtask) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Scavengingtask) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Scavengingtask) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Scavengingtask) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *Scavengingtask) GetAssociatedObject() string`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *Scavengingtask) GetAssociatedObjectOk() (*string, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *Scavengingtask) SetAssociatedObject(v string)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *Scavengingtask) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetEndTime

`func (o *Scavengingtask) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Scavengingtask) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Scavengingtask) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Scavengingtask) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetProcessedRecords

`func (o *Scavengingtask) GetProcessedRecords() int64`

GetProcessedRecords returns the ProcessedRecords field if non-nil, zero value otherwise.

### GetProcessedRecordsOk

`func (o *Scavengingtask) GetProcessedRecordsOk() (*int64, bool)`

GetProcessedRecordsOk returns a tuple with the ProcessedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRecords

`func (o *Scavengingtask) SetProcessedRecords(v int64)`

SetProcessedRecords sets ProcessedRecords field to given value.

### HasProcessedRecords

`func (o *Scavengingtask) HasProcessedRecords() bool`

HasProcessedRecords returns a boolean if a field has been set.

### GetReclaimableRecords

`func (o *Scavengingtask) GetReclaimableRecords() int64`

GetReclaimableRecords returns the ReclaimableRecords field if non-nil, zero value otherwise.

### GetReclaimableRecordsOk

`func (o *Scavengingtask) GetReclaimableRecordsOk() (*int64, bool)`

GetReclaimableRecordsOk returns a tuple with the ReclaimableRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimableRecords

`func (o *Scavengingtask) SetReclaimableRecords(v int64)`

SetReclaimableRecords sets ReclaimableRecords field to given value.

### HasReclaimableRecords

`func (o *Scavengingtask) HasReclaimableRecords() bool`

HasReclaimableRecords returns a boolean if a field has been set.

### GetReclaimedRecords

`func (o *Scavengingtask) GetReclaimedRecords() int64`

GetReclaimedRecords returns the ReclaimedRecords field if non-nil, zero value otherwise.

### GetReclaimedRecordsOk

`func (o *Scavengingtask) GetReclaimedRecordsOk() (*int64, bool)`

GetReclaimedRecordsOk returns a tuple with the ReclaimedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimedRecords

`func (o *Scavengingtask) SetReclaimedRecords(v int64)`

SetReclaimedRecords sets ReclaimedRecords field to given value.

### HasReclaimedRecords

`func (o *Scavengingtask) HasReclaimedRecords() bool`

HasReclaimedRecords returns a boolean if a field has been set.

### GetStartTime

`func (o *Scavengingtask) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Scavengingtask) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Scavengingtask) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Scavengingtask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *Scavengingtask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Scavengingtask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Scavengingtask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Scavengingtask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


