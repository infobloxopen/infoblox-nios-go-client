# DistributionscheduleUpgradeGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The upgrade group name. | [optional] 
**TimeZone** | Pointer to **string** | The time zone for scheduling operations. | [optional] [readonly] 
**DistributionDependentGroup** | Pointer to **string** | The distribution dependent group name. | [optional] 
**UpgradeDependentGroup** | Pointer to **string** | The upgrade dependent group name. | [optional] 
**DistributionTime** | Pointer to **int64** | The time of the next scheduled distribution. | [optional] 
**UpgradeTime** | Pointer to **int64** | The time of the next scheduled upgrade. | [optional] 

## Methods

### NewDistributionscheduleUpgradeGroups

`func NewDistributionscheduleUpgradeGroups() *DistributionscheduleUpgradeGroups`

NewDistributionscheduleUpgradeGroups instantiates a new DistributionscheduleUpgradeGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionscheduleUpgradeGroupsWithDefaults

`func NewDistributionscheduleUpgradeGroupsWithDefaults() *DistributionscheduleUpgradeGroups`

NewDistributionscheduleUpgradeGroupsWithDefaults instantiates a new DistributionscheduleUpgradeGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DistributionscheduleUpgradeGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributionscheduleUpgradeGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributionscheduleUpgradeGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DistributionscheduleUpgradeGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *DistributionscheduleUpgradeGroups) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DistributionscheduleUpgradeGroups) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DistributionscheduleUpgradeGroups) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DistributionscheduleUpgradeGroups) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetDistributionDependentGroup

`func (o *DistributionscheduleUpgradeGroups) GetDistributionDependentGroup() string`

GetDistributionDependentGroup returns the DistributionDependentGroup field if non-nil, zero value otherwise.

### GetDistributionDependentGroupOk

`func (o *DistributionscheduleUpgradeGroups) GetDistributionDependentGroupOk() (*string, bool)`

GetDistributionDependentGroupOk returns a tuple with the DistributionDependentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionDependentGroup

`func (o *DistributionscheduleUpgradeGroups) SetDistributionDependentGroup(v string)`

SetDistributionDependentGroup sets DistributionDependentGroup field to given value.

### HasDistributionDependentGroup

`func (o *DistributionscheduleUpgradeGroups) HasDistributionDependentGroup() bool`

HasDistributionDependentGroup returns a boolean if a field has been set.

### GetUpgradeDependentGroup

`func (o *DistributionscheduleUpgradeGroups) GetUpgradeDependentGroup() string`

GetUpgradeDependentGroup returns the UpgradeDependentGroup field if non-nil, zero value otherwise.

### GetUpgradeDependentGroupOk

`func (o *DistributionscheduleUpgradeGroups) GetUpgradeDependentGroupOk() (*string, bool)`

GetUpgradeDependentGroupOk returns a tuple with the UpgradeDependentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDependentGroup

`func (o *DistributionscheduleUpgradeGroups) SetUpgradeDependentGroup(v string)`

SetUpgradeDependentGroup sets UpgradeDependentGroup field to given value.

### HasUpgradeDependentGroup

`func (o *DistributionscheduleUpgradeGroups) HasUpgradeDependentGroup() bool`

HasUpgradeDependentGroup returns a boolean if a field has been set.

### GetDistributionTime

`func (o *DistributionscheduleUpgradeGroups) GetDistributionTime() int64`

GetDistributionTime returns the DistributionTime field if non-nil, zero value otherwise.

### GetDistributionTimeOk

`func (o *DistributionscheduleUpgradeGroups) GetDistributionTimeOk() (*int64, bool)`

GetDistributionTimeOk returns a tuple with the DistributionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionTime

`func (o *DistributionscheduleUpgradeGroups) SetDistributionTime(v int64)`

SetDistributionTime sets DistributionTime field to given value.

### HasDistributionTime

`func (o *DistributionscheduleUpgradeGroups) HasDistributionTime() bool`

HasDistributionTime returns a boolean if a field has been set.

### GetUpgradeTime

`func (o *DistributionscheduleUpgradeGroups) GetUpgradeTime() int64`

GetUpgradeTime returns the UpgradeTime field if non-nil, zero value otherwise.

### GetUpgradeTimeOk

`func (o *DistributionscheduleUpgradeGroups) GetUpgradeTimeOk() (*int64, bool)`

GetUpgradeTimeOk returns a tuple with the UpgradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTime

`func (o *DistributionscheduleUpgradeGroups) SetUpgradeTime(v int64)`

SetUpgradeTime sets UpgradeTime field to given value.

### HasUpgradeTime

`func (o *DistributionscheduleUpgradeGroups) HasUpgradeTime() bool`

HasUpgradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


