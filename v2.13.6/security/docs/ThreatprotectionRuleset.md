# ThreatprotectionRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddType** | Pointer to **string** | Determines the way the ruleset was added. | [optional] [readonly] 
**AddedTime** | Pointer to **int64** | The time when the ruleset was added. | [optional] [readonly] 
**Comment** | Pointer to **string** | The human readable comment for the ruleset. | [optional] 
**DoNotDelete** | Pointer to **bool** | Determines if the ruleset will not be deleted during upgrade. | [optional] 
**IsFactoryResetEnabled** | Pointer to **bool** | Determines if factory reset is enabled for this ruleset. | [optional] [readonly] 
**UsedBy** | Pointer to **[]string** | The users of the ruleset. | [optional] [readonly] 
**Version** | Pointer to **string** | The ruleset version. | [optional] [readonly] 

## Methods

### NewThreatprotectionRuleset

`func NewThreatprotectionRuleset() *ThreatprotectionRuleset`

NewThreatprotectionRuleset instantiates a new ThreatprotectionRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionRulesetWithDefaults

`func NewThreatprotectionRulesetWithDefaults() *ThreatprotectionRuleset`

NewThreatprotectionRulesetWithDefaults instantiates a new ThreatprotectionRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatprotectionRuleset) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatprotectionRuleset) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatprotectionRuleset) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatprotectionRuleset) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddType

`func (o *ThreatprotectionRuleset) GetAddType() string`

GetAddType returns the AddType field if non-nil, zero value otherwise.

### GetAddTypeOk

`func (o *ThreatprotectionRuleset) GetAddTypeOk() (*string, bool)`

GetAddTypeOk returns a tuple with the AddType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddType

`func (o *ThreatprotectionRuleset) SetAddType(v string)`

SetAddType sets AddType field to given value.

### HasAddType

`func (o *ThreatprotectionRuleset) HasAddType() bool`

HasAddType returns a boolean if a field has been set.

### GetAddedTime

`func (o *ThreatprotectionRuleset) GetAddedTime() int64`

GetAddedTime returns the AddedTime field if non-nil, zero value otherwise.

### GetAddedTimeOk

`func (o *ThreatprotectionRuleset) GetAddedTimeOk() (*int64, bool)`

GetAddedTimeOk returns a tuple with the AddedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTime

`func (o *ThreatprotectionRuleset) SetAddedTime(v int64)`

SetAddedTime sets AddedTime field to given value.

### HasAddedTime

`func (o *ThreatprotectionRuleset) HasAddedTime() bool`

HasAddedTime returns a boolean if a field has been set.

### GetComment

`func (o *ThreatprotectionRuleset) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ThreatprotectionRuleset) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ThreatprotectionRuleset) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ThreatprotectionRuleset) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDoNotDelete

`func (o *ThreatprotectionRuleset) GetDoNotDelete() bool`

GetDoNotDelete returns the DoNotDelete field if non-nil, zero value otherwise.

### GetDoNotDeleteOk

`func (o *ThreatprotectionRuleset) GetDoNotDeleteOk() (*bool, bool)`

GetDoNotDeleteOk returns a tuple with the DoNotDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDelete

`func (o *ThreatprotectionRuleset) SetDoNotDelete(v bool)`

SetDoNotDelete sets DoNotDelete field to given value.

### HasDoNotDelete

`func (o *ThreatprotectionRuleset) HasDoNotDelete() bool`

HasDoNotDelete returns a boolean if a field has been set.

### GetIsFactoryResetEnabled

`func (o *ThreatprotectionRuleset) GetIsFactoryResetEnabled() bool`

GetIsFactoryResetEnabled returns the IsFactoryResetEnabled field if non-nil, zero value otherwise.

### GetIsFactoryResetEnabledOk

`func (o *ThreatprotectionRuleset) GetIsFactoryResetEnabledOk() (*bool, bool)`

GetIsFactoryResetEnabledOk returns a tuple with the IsFactoryResetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactoryResetEnabled

`func (o *ThreatprotectionRuleset) SetIsFactoryResetEnabled(v bool)`

SetIsFactoryResetEnabled sets IsFactoryResetEnabled field to given value.

### HasIsFactoryResetEnabled

`func (o *ThreatprotectionRuleset) HasIsFactoryResetEnabled() bool`

HasIsFactoryResetEnabled returns a boolean if a field has been set.

### GetUsedBy

`func (o *ThreatprotectionRuleset) GetUsedBy() []string`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *ThreatprotectionRuleset) GetUsedByOk() (*[]string, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *ThreatprotectionRuleset) SetUsedBy(v []string)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *ThreatprotectionRuleset) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.

### GetVersion

`func (o *ThreatprotectionRuleset) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreatprotectionRuleset) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreatprotectionRuleset) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreatprotectionRuleset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


