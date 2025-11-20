# GetThreatprotectionRulesetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AddType** | Pointer to **string** | Determines the way the ruleset was added. | [optional] [readonly] 
**AddedTime** | Pointer to **int64** | The time when the ruleset was added. | [optional] [readonly] 
**Comment** | Pointer to **string** | The human readable comment for the ruleset. | [optional] 
**DoNotDelete** | Pointer to **bool** | Determines if the ruleset will not be deleted during upgrade. | [optional] 
**IsFactoryResetEnabled** | Pointer to **bool** | Determines if factory reset is enabled for this ruleset. | [optional] [readonly] 
**UsedBy** | Pointer to **[]string** | The users of the ruleset. | [optional] [readonly] 
**Version** | Pointer to **string** | The ruleset version. | [optional] [readonly] 
**Result** | Pointer to [**ThreatprotectionRuleset**](ThreatprotectionRuleset.md) |  | [optional] 

## Methods

### NewGetThreatprotectionRulesetResponse

`func NewGetThreatprotectionRulesetResponse() *GetThreatprotectionRulesetResponse`

NewGetThreatprotectionRulesetResponse instantiates a new GetThreatprotectionRulesetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionRulesetResponseWithDefaults

`func NewGetThreatprotectionRulesetResponseWithDefaults() *GetThreatprotectionRulesetResponse`

NewGetThreatprotectionRulesetResponseWithDefaults instantiates a new GetThreatprotectionRulesetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionRulesetResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionRulesetResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionRulesetResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionRulesetResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddType

`func (o *GetThreatprotectionRulesetResponse) GetAddType() string`

GetAddType returns the AddType field if non-nil, zero value otherwise.

### GetAddTypeOk

`func (o *GetThreatprotectionRulesetResponse) GetAddTypeOk() (*string, bool)`

GetAddTypeOk returns a tuple with the AddType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddType

`func (o *GetThreatprotectionRulesetResponse) SetAddType(v string)`

SetAddType sets AddType field to given value.

### HasAddType

`func (o *GetThreatprotectionRulesetResponse) HasAddType() bool`

HasAddType returns a boolean if a field has been set.

### GetAddedTime

`func (o *GetThreatprotectionRulesetResponse) GetAddedTime() int64`

GetAddedTime returns the AddedTime field if non-nil, zero value otherwise.

### GetAddedTimeOk

`func (o *GetThreatprotectionRulesetResponse) GetAddedTimeOk() (*int64, bool)`

GetAddedTimeOk returns a tuple with the AddedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTime

`func (o *GetThreatprotectionRulesetResponse) SetAddedTime(v int64)`

SetAddedTime sets AddedTime field to given value.

### HasAddedTime

`func (o *GetThreatprotectionRulesetResponse) HasAddedTime() bool`

HasAddedTime returns a boolean if a field has been set.

### GetComment

`func (o *GetThreatprotectionRulesetResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetThreatprotectionRulesetResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetThreatprotectionRulesetResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetThreatprotectionRulesetResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDoNotDelete

`func (o *GetThreatprotectionRulesetResponse) GetDoNotDelete() bool`

GetDoNotDelete returns the DoNotDelete field if non-nil, zero value otherwise.

### GetDoNotDeleteOk

`func (o *GetThreatprotectionRulesetResponse) GetDoNotDeleteOk() (*bool, bool)`

GetDoNotDeleteOk returns a tuple with the DoNotDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDelete

`func (o *GetThreatprotectionRulesetResponse) SetDoNotDelete(v bool)`

SetDoNotDelete sets DoNotDelete field to given value.

### HasDoNotDelete

`func (o *GetThreatprotectionRulesetResponse) HasDoNotDelete() bool`

HasDoNotDelete returns a boolean if a field has been set.

### GetIsFactoryResetEnabled

`func (o *GetThreatprotectionRulesetResponse) GetIsFactoryResetEnabled() bool`

GetIsFactoryResetEnabled returns the IsFactoryResetEnabled field if non-nil, zero value otherwise.

### GetIsFactoryResetEnabledOk

`func (o *GetThreatprotectionRulesetResponse) GetIsFactoryResetEnabledOk() (*bool, bool)`

GetIsFactoryResetEnabledOk returns a tuple with the IsFactoryResetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactoryResetEnabled

`func (o *GetThreatprotectionRulesetResponse) SetIsFactoryResetEnabled(v bool)`

SetIsFactoryResetEnabled sets IsFactoryResetEnabled field to given value.

### HasIsFactoryResetEnabled

`func (o *GetThreatprotectionRulesetResponse) HasIsFactoryResetEnabled() bool`

HasIsFactoryResetEnabled returns a boolean if a field has been set.

### GetUsedBy

`func (o *GetThreatprotectionRulesetResponse) GetUsedBy() []string`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *GetThreatprotectionRulesetResponse) GetUsedByOk() (*[]string, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *GetThreatprotectionRulesetResponse) SetUsedBy(v []string)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *GetThreatprotectionRulesetResponse) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.

### GetVersion

`func (o *GetThreatprotectionRulesetResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetThreatprotectionRulesetResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetThreatprotectionRulesetResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetThreatprotectionRulesetResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionRulesetResponse) GetResult() ThreatprotectionRuleset`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionRulesetResponse) GetResultOk() (*ThreatprotectionRuleset, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionRulesetResponse) SetResult(v ThreatprotectionRuleset)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionRulesetResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


