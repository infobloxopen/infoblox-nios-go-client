# Upgradegroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The upgrade group descriptive comment. | [optional] 
**DistributionDependentGroup** | Pointer to **string** | The distribution dependent group name. | [optional] 
**DistributionPolicy** | Pointer to **string** | The distribution scheduling policy. | [optional] 
**DistributionTime** | Pointer to **int64** | The time of the next scheduled distribution. | [optional] 
**Members** | Pointer to [**[]UpgradegroupMembers**](UpgradegroupMembers.md) | The upgrade group members. | [optional] 
**Name** | Pointer to **string** | The upgrade group name. | [optional] 
**TimeZone** | Pointer to **string** | The time zone for scheduling operations. | [optional] [readonly] 
**UpgradeDependentGroup** | Pointer to **string** | The upgrade dependent group name. | [optional] 
**UpgradePolicy** | Pointer to **string** | The upgrade scheduling policy. | [optional] 
**UpgradeTime** | Pointer to **int64** | The time of the next scheduled upgrade. | [optional] 

## Methods

### NewUpgradegroup

`func NewUpgradegroup() *Upgradegroup`

NewUpgradegroup instantiates a new Upgradegroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradegroupWithDefaults

`func NewUpgradegroupWithDefaults() *Upgradegroup`

NewUpgradegroupWithDefaults instantiates a new Upgradegroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Upgradegroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Upgradegroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Upgradegroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Upgradegroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Upgradegroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Upgradegroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Upgradegroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Upgradegroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDistributionDependentGroup

`func (o *Upgradegroup) GetDistributionDependentGroup() string`

GetDistributionDependentGroup returns the DistributionDependentGroup field if non-nil, zero value otherwise.

### GetDistributionDependentGroupOk

`func (o *Upgradegroup) GetDistributionDependentGroupOk() (*string, bool)`

GetDistributionDependentGroupOk returns a tuple with the DistributionDependentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionDependentGroup

`func (o *Upgradegroup) SetDistributionDependentGroup(v string)`

SetDistributionDependentGroup sets DistributionDependentGroup field to given value.

### HasDistributionDependentGroup

`func (o *Upgradegroup) HasDistributionDependentGroup() bool`

HasDistributionDependentGroup returns a boolean if a field has been set.

### GetDistributionPolicy

`func (o *Upgradegroup) GetDistributionPolicy() string`

GetDistributionPolicy returns the DistributionPolicy field if non-nil, zero value otherwise.

### GetDistributionPolicyOk

`func (o *Upgradegroup) GetDistributionPolicyOk() (*string, bool)`

GetDistributionPolicyOk returns a tuple with the DistributionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPolicy

`func (o *Upgradegroup) SetDistributionPolicy(v string)`

SetDistributionPolicy sets DistributionPolicy field to given value.

### HasDistributionPolicy

`func (o *Upgradegroup) HasDistributionPolicy() bool`

HasDistributionPolicy returns a boolean if a field has been set.

### GetDistributionTime

`func (o *Upgradegroup) GetDistributionTime() int64`

GetDistributionTime returns the DistributionTime field if non-nil, zero value otherwise.

### GetDistributionTimeOk

`func (o *Upgradegroup) GetDistributionTimeOk() (*int64, bool)`

GetDistributionTimeOk returns a tuple with the DistributionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionTime

`func (o *Upgradegroup) SetDistributionTime(v int64)`

SetDistributionTime sets DistributionTime field to given value.

### HasDistributionTime

`func (o *Upgradegroup) HasDistributionTime() bool`

HasDistributionTime returns a boolean if a field has been set.

### GetMembers

`func (o *Upgradegroup) GetMembers() []UpgradegroupMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Upgradegroup) GetMembersOk() (*[]UpgradegroupMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Upgradegroup) SetMembers(v []UpgradegroupMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Upgradegroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *Upgradegroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Upgradegroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Upgradegroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Upgradegroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *Upgradegroup) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Upgradegroup) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Upgradegroup) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Upgradegroup) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUpgradeDependentGroup

`func (o *Upgradegroup) GetUpgradeDependentGroup() string`

GetUpgradeDependentGroup returns the UpgradeDependentGroup field if non-nil, zero value otherwise.

### GetUpgradeDependentGroupOk

`func (o *Upgradegroup) GetUpgradeDependentGroupOk() (*string, bool)`

GetUpgradeDependentGroupOk returns a tuple with the UpgradeDependentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDependentGroup

`func (o *Upgradegroup) SetUpgradeDependentGroup(v string)`

SetUpgradeDependentGroup sets UpgradeDependentGroup field to given value.

### HasUpgradeDependentGroup

`func (o *Upgradegroup) HasUpgradeDependentGroup() bool`

HasUpgradeDependentGroup returns a boolean if a field has been set.

### GetUpgradePolicy

`func (o *Upgradegroup) GetUpgradePolicy() string`

GetUpgradePolicy returns the UpgradePolicy field if non-nil, zero value otherwise.

### GetUpgradePolicyOk

`func (o *Upgradegroup) GetUpgradePolicyOk() (*string, bool)`

GetUpgradePolicyOk returns a tuple with the UpgradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePolicy

`func (o *Upgradegroup) SetUpgradePolicy(v string)`

SetUpgradePolicy sets UpgradePolicy field to given value.

### HasUpgradePolicy

`func (o *Upgradegroup) HasUpgradePolicy() bool`

HasUpgradePolicy returns a boolean if a field has been set.

### GetUpgradeTime

`func (o *Upgradegroup) GetUpgradeTime() int64`

GetUpgradeTime returns the UpgradeTime field if non-nil, zero value otherwise.

### GetUpgradeTimeOk

`func (o *Upgradegroup) GetUpgradeTimeOk() (*int64, bool)`

GetUpgradeTimeOk returns a tuple with the UpgradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTime

`func (o *Upgradegroup) SetUpgradeTime(v int64)`

SetUpgradeTime sets UpgradeTime field to given value.

### HasUpgradeTime

`func (o *Upgradegroup) HasUpgradeTime() bool`

HasUpgradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


