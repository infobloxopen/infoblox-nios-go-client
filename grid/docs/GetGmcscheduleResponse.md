# GetGmcscheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ActivateGmcGroupSchedule** | Pointer to **bool** | Determines whether the gmc schedule is active. | [optional] 
**GmcGroups** | Pointer to **[]string** | Object array of gmc groups | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Gmcschedule**](Gmcschedule.md) |  | [optional] 

## Methods

### NewGetGmcscheduleResponse

`func NewGetGmcscheduleResponse() *GetGmcscheduleResponse`

NewGetGmcscheduleResponse instantiates a new GetGmcscheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGmcscheduleResponseWithDefaults

`func NewGetGmcscheduleResponseWithDefaults() *GetGmcscheduleResponse`

NewGetGmcscheduleResponseWithDefaults instantiates a new GetGmcscheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGmcscheduleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGmcscheduleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGmcscheduleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGmcscheduleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActivateGmcGroupSchedule

`func (o *GetGmcscheduleResponse) GetActivateGmcGroupSchedule() bool`

GetActivateGmcGroupSchedule returns the ActivateGmcGroupSchedule field if non-nil, zero value otherwise.

### GetActivateGmcGroupScheduleOk

`func (o *GetGmcscheduleResponse) GetActivateGmcGroupScheduleOk() (*bool, bool)`

GetActivateGmcGroupScheduleOk returns a tuple with the ActivateGmcGroupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateGmcGroupSchedule

`func (o *GetGmcscheduleResponse) SetActivateGmcGroupSchedule(v bool)`

SetActivateGmcGroupSchedule sets ActivateGmcGroupSchedule field to given value.

### HasActivateGmcGroupSchedule

`func (o *GetGmcscheduleResponse) HasActivateGmcGroupSchedule() bool`

HasActivateGmcGroupSchedule returns a boolean if a field has been set.

### GetGmcGroups

`func (o *GetGmcscheduleResponse) GetGmcGroups() []string`

GetGmcGroups returns the GmcGroups field if non-nil, zero value otherwise.

### GetGmcGroupsOk

`func (o *GetGmcscheduleResponse) GetGmcGroupsOk() (*[]string, bool)`

GetGmcGroupsOk returns a tuple with the GmcGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmcGroups

`func (o *GetGmcscheduleResponse) SetGmcGroups(v []string)`

SetGmcGroups sets GmcGroups field to given value.

### HasGmcGroups

`func (o *GetGmcscheduleResponse) HasGmcGroups() bool`

HasGmcGroups returns a boolean if a field has been set.

### GetUuid

`func (o *GetGmcscheduleResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGmcscheduleResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGmcscheduleResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGmcscheduleResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGmcscheduleResponse) GetResult() Gmcschedule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGmcscheduleResponse) GetResultOk() (*Gmcschedule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGmcscheduleResponse) SetResult(v Gmcschedule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGmcscheduleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


