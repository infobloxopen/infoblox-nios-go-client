# GetDistributionscheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Active** | Pointer to **bool** | Determines whether the distribution schedule is active. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the distribution. | [optional] 
**TimeZone** | Pointer to **string** | Time zone of the distribution start time. | [optional] [readonly] 
**UpgradeGroups** | Pointer to [**[]DistributionscheduleUpgradeGroups**](DistributionscheduleUpgradeGroups.md) | The upgrade groups scheduling settings. | [optional] 
**Result** | Pointer to [**Distributionschedule**](Distributionschedule.md) |  | [optional] 

## Methods

### NewGetDistributionscheduleResponse

`func NewGetDistributionscheduleResponse() *GetDistributionscheduleResponse`

NewGetDistributionscheduleResponse instantiates a new GetDistributionscheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDistributionscheduleResponseWithDefaults

`func NewGetDistributionscheduleResponseWithDefaults() *GetDistributionscheduleResponse`

NewGetDistributionscheduleResponseWithDefaults instantiates a new GetDistributionscheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDistributionscheduleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDistributionscheduleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDistributionscheduleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDistributionscheduleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActive

`func (o *GetDistributionscheduleResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetDistributionscheduleResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetDistributionscheduleResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetDistributionscheduleResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartTime

`func (o *GetDistributionscheduleResponse) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetDistributionscheduleResponse) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetDistributionscheduleResponse) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetDistributionscheduleResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetDistributionscheduleResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetDistributionscheduleResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetDistributionscheduleResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetDistributionscheduleResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUpgradeGroups

`func (o *GetDistributionscheduleResponse) GetUpgradeGroups() []DistributionscheduleUpgradeGroups`

GetUpgradeGroups returns the UpgradeGroups field if non-nil, zero value otherwise.

### GetUpgradeGroupsOk

`func (o *GetDistributionscheduleResponse) GetUpgradeGroupsOk() (*[]DistributionscheduleUpgradeGroups, bool)`

GetUpgradeGroupsOk returns a tuple with the UpgradeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGroups

`func (o *GetDistributionscheduleResponse) SetUpgradeGroups(v []DistributionscheduleUpgradeGroups)`

SetUpgradeGroups sets UpgradeGroups field to given value.

### HasUpgradeGroups

`func (o *GetDistributionscheduleResponse) HasUpgradeGroups() bool`

HasUpgradeGroups returns a boolean if a field has been set.

### GetResult

`func (o *GetDistributionscheduleResponse) GetResult() Distributionschedule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDistributionscheduleResponse) GetResultOk() (*Distributionschedule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDistributionscheduleResponse) SetResult(v Distributionschedule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDistributionscheduleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


