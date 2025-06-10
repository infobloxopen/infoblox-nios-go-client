# GetUpgradegroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Result** | Pointer to [**Upgradegroup**](Upgradegroup.md) |  | [optional] 

## Methods

### NewGetUpgradegroupResponse

`func NewGetUpgradegroupResponse() *GetUpgradegroupResponse`

NewGetUpgradegroupResponse instantiates a new GetUpgradegroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUpgradegroupResponseWithDefaults

`func NewGetUpgradegroupResponseWithDefaults() *GetUpgradegroupResponse`

NewGetUpgradegroupResponseWithDefaults instantiates a new GetUpgradegroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetUpgradegroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetUpgradegroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetUpgradegroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetUpgradegroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetUpgradegroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetUpgradegroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetUpgradegroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetUpgradegroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDistributionDependentGroup

`func (o *GetUpgradegroupResponse) GetDistributionDependentGroup() string`

GetDistributionDependentGroup returns the DistributionDependentGroup field if non-nil, zero value otherwise.

### GetDistributionDependentGroupOk

`func (o *GetUpgradegroupResponse) GetDistributionDependentGroupOk() (*string, bool)`

GetDistributionDependentGroupOk returns a tuple with the DistributionDependentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionDependentGroup

`func (o *GetUpgradegroupResponse) SetDistributionDependentGroup(v string)`

SetDistributionDependentGroup sets DistributionDependentGroup field to given value.

### HasDistributionDependentGroup

`func (o *GetUpgradegroupResponse) HasDistributionDependentGroup() bool`

HasDistributionDependentGroup returns a boolean if a field has been set.

### GetDistributionPolicy

`func (o *GetUpgradegroupResponse) GetDistributionPolicy() string`

GetDistributionPolicy returns the DistributionPolicy field if non-nil, zero value otherwise.

### GetDistributionPolicyOk

`func (o *GetUpgradegroupResponse) GetDistributionPolicyOk() (*string, bool)`

GetDistributionPolicyOk returns a tuple with the DistributionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPolicy

`func (o *GetUpgradegroupResponse) SetDistributionPolicy(v string)`

SetDistributionPolicy sets DistributionPolicy field to given value.

### HasDistributionPolicy

`func (o *GetUpgradegroupResponse) HasDistributionPolicy() bool`

HasDistributionPolicy returns a boolean if a field has been set.

### GetDistributionTime

`func (o *GetUpgradegroupResponse) GetDistributionTime() int64`

GetDistributionTime returns the DistributionTime field if non-nil, zero value otherwise.

### GetDistributionTimeOk

`func (o *GetUpgradegroupResponse) GetDistributionTimeOk() (*int64, bool)`

GetDistributionTimeOk returns a tuple with the DistributionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionTime

`func (o *GetUpgradegroupResponse) SetDistributionTime(v int64)`

SetDistributionTime sets DistributionTime field to given value.

### HasDistributionTime

`func (o *GetUpgradegroupResponse) HasDistributionTime() bool`

HasDistributionTime returns a boolean if a field has been set.

### GetMembers

`func (o *GetUpgradegroupResponse) GetMembers() []UpgradegroupMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetUpgradegroupResponse) GetMembersOk() (*[]UpgradegroupMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetUpgradegroupResponse) SetMembers(v []UpgradegroupMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetUpgradegroupResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *GetUpgradegroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUpgradegroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUpgradegroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUpgradegroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetUpgradegroupResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetUpgradegroupResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetUpgradegroupResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetUpgradegroupResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUpgradeDependentGroup

`func (o *GetUpgradegroupResponse) GetUpgradeDependentGroup() string`

GetUpgradeDependentGroup returns the UpgradeDependentGroup field if non-nil, zero value otherwise.

### GetUpgradeDependentGroupOk

`func (o *GetUpgradegroupResponse) GetUpgradeDependentGroupOk() (*string, bool)`

GetUpgradeDependentGroupOk returns a tuple with the UpgradeDependentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDependentGroup

`func (o *GetUpgradegroupResponse) SetUpgradeDependentGroup(v string)`

SetUpgradeDependentGroup sets UpgradeDependentGroup field to given value.

### HasUpgradeDependentGroup

`func (o *GetUpgradegroupResponse) HasUpgradeDependentGroup() bool`

HasUpgradeDependentGroup returns a boolean if a field has been set.

### GetUpgradePolicy

`func (o *GetUpgradegroupResponse) GetUpgradePolicy() string`

GetUpgradePolicy returns the UpgradePolicy field if non-nil, zero value otherwise.

### GetUpgradePolicyOk

`func (o *GetUpgradegroupResponse) GetUpgradePolicyOk() (*string, bool)`

GetUpgradePolicyOk returns a tuple with the UpgradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePolicy

`func (o *GetUpgradegroupResponse) SetUpgradePolicy(v string)`

SetUpgradePolicy sets UpgradePolicy field to given value.

### HasUpgradePolicy

`func (o *GetUpgradegroupResponse) HasUpgradePolicy() bool`

HasUpgradePolicy returns a boolean if a field has been set.

### GetUpgradeTime

`func (o *GetUpgradegroupResponse) GetUpgradeTime() int64`

GetUpgradeTime returns the UpgradeTime field if non-nil, zero value otherwise.

### GetUpgradeTimeOk

`func (o *GetUpgradegroupResponse) GetUpgradeTimeOk() (*int64, bool)`

GetUpgradeTimeOk returns a tuple with the UpgradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTime

`func (o *GetUpgradegroupResponse) SetUpgradeTime(v int64)`

SetUpgradeTime sets UpgradeTime field to given value.

### HasUpgradeTime

`func (o *GetUpgradegroupResponse) HasUpgradeTime() bool`

HasUpgradeTime returns a boolean if a field has been set.

### GetResult

`func (o *GetUpgradegroupResponse) GetResult() Upgradegroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetUpgradegroupResponse) GetResultOk() (*Upgradegroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetUpgradegroupResponse) SetResult(v Upgradegroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetUpgradegroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


