# Distributionschedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Active** | Pointer to **bool** | Determines whether the distribution schedule is active. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the distribution. | [optional] 
**TimeZone** | Pointer to **string** | Time zone of the distribution start time. | [optional] [readonly] 
**UpgradeGroups** | Pointer to [**[]DistributionscheduleUpgradeGroups**](DistributionscheduleUpgradeGroups.md) | The upgrade groups scheduling settings. | [optional] 

## Methods

### NewDistributionschedule

`func NewDistributionschedule() *Distributionschedule`

NewDistributionschedule instantiates a new Distributionschedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionscheduleWithDefaults

`func NewDistributionscheduleWithDefaults() *Distributionschedule`

NewDistributionscheduleWithDefaults instantiates a new Distributionschedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Distributionschedule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Distributionschedule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Distributionschedule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Distributionschedule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActive

`func (o *Distributionschedule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Distributionschedule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Distributionschedule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Distributionschedule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartTime

`func (o *Distributionschedule) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Distributionschedule) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Distributionschedule) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Distributionschedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *Distributionschedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Distributionschedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Distributionschedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Distributionschedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUpgradeGroups

`func (o *Distributionschedule) GetUpgradeGroups() []DistributionscheduleUpgradeGroups`

GetUpgradeGroups returns the UpgradeGroups field if non-nil, zero value otherwise.

### GetUpgradeGroupsOk

`func (o *Distributionschedule) GetUpgradeGroupsOk() (*[]DistributionscheduleUpgradeGroups, bool)`

GetUpgradeGroupsOk returns a tuple with the UpgradeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGroups

`func (o *Distributionschedule) SetUpgradeGroups(v []DistributionscheduleUpgradeGroups)`

SetUpgradeGroups sets UpgradeGroups field to given value.

### HasUpgradeGroups

`func (o *Distributionschedule) HasUpgradeGroups() bool`

HasUpgradeGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


