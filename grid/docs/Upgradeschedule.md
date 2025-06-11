# Upgradeschedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Active** | Pointer to **bool** | Determines whether the upgrade schedule is active. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the upgrade. | [optional] 
**TimeZone** | Pointer to **string** | The time zone for upgrade start time. | [optional] [readonly] 
**UpgradeGroups** | Pointer to [**[]UpgradescheduleUpgradeGroups**](UpgradescheduleUpgradeGroups.md) | The upgrade groups scheduling settings. | [optional] 

## Methods

### NewUpgradeschedule

`func NewUpgradeschedule() *Upgradeschedule`

NewUpgradeschedule instantiates a new Upgradeschedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradescheduleWithDefaults

`func NewUpgradescheduleWithDefaults() *Upgradeschedule`

NewUpgradescheduleWithDefaults instantiates a new Upgradeschedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Upgradeschedule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Upgradeschedule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Upgradeschedule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Upgradeschedule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActive

`func (o *Upgradeschedule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Upgradeschedule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Upgradeschedule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Upgradeschedule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartTime

`func (o *Upgradeschedule) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Upgradeschedule) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Upgradeschedule) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Upgradeschedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *Upgradeschedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Upgradeschedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Upgradeschedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Upgradeschedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUpgradeGroups

`func (o *Upgradeschedule) GetUpgradeGroups() []UpgradescheduleUpgradeGroups`

GetUpgradeGroups returns the UpgradeGroups field if non-nil, zero value otherwise.

### GetUpgradeGroupsOk

`func (o *Upgradeschedule) GetUpgradeGroupsOk() (*[]UpgradescheduleUpgradeGroups, bool)`

GetUpgradeGroupsOk returns a tuple with the UpgradeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGroups

`func (o *Upgradeschedule) SetUpgradeGroups(v []UpgradescheduleUpgradeGroups)`

SetUpgradeGroups sets UpgradeGroups field to given value.

### HasUpgradeGroups

`func (o *Upgradeschedule) HasUpgradeGroups() bool`

HasUpgradeGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


