# GetThreatprotectionRulecategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**IsFactoryResetEnabled** | Pointer to **bool** | Determines if factory reset is enabled for this rule category. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the rule category. | [optional] [readonly] 
**Ruleset** | Pointer to **string** | The version of the ruleset the category assigned to. | [optional] [readonly] 
**Result** | Pointer to [**ThreatprotectionRulecategory**](ThreatprotectionRulecategory.md) |  | [optional] 

## Methods

### NewGetThreatprotectionRulecategoryResponse

`func NewGetThreatprotectionRulecategoryResponse() *GetThreatprotectionRulecategoryResponse`

NewGetThreatprotectionRulecategoryResponse instantiates a new GetThreatprotectionRulecategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionRulecategoryResponseWithDefaults

`func NewGetThreatprotectionRulecategoryResponseWithDefaults() *GetThreatprotectionRulecategoryResponse`

NewGetThreatprotectionRulecategoryResponseWithDefaults instantiates a new GetThreatprotectionRulecategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionRulecategoryResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionRulecategoryResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionRulecategoryResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionRulecategoryResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIsFactoryResetEnabled

`func (o *GetThreatprotectionRulecategoryResponse) GetIsFactoryResetEnabled() bool`

GetIsFactoryResetEnabled returns the IsFactoryResetEnabled field if non-nil, zero value otherwise.

### GetIsFactoryResetEnabledOk

`func (o *GetThreatprotectionRulecategoryResponse) GetIsFactoryResetEnabledOk() (*bool, bool)`

GetIsFactoryResetEnabledOk returns a tuple with the IsFactoryResetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactoryResetEnabled

`func (o *GetThreatprotectionRulecategoryResponse) SetIsFactoryResetEnabled(v bool)`

SetIsFactoryResetEnabled sets IsFactoryResetEnabled field to given value.

### HasIsFactoryResetEnabled

`func (o *GetThreatprotectionRulecategoryResponse) HasIsFactoryResetEnabled() bool`

HasIsFactoryResetEnabled returns a boolean if a field has been set.

### GetName

`func (o *GetThreatprotectionRulecategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetThreatprotectionRulecategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetThreatprotectionRulecategoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetThreatprotectionRulecategoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleset

`func (o *GetThreatprotectionRulecategoryResponse) GetRuleset() string`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *GetThreatprotectionRulecategoryResponse) GetRulesetOk() (*string, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *GetThreatprotectionRulecategoryResponse) SetRuleset(v string)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *GetThreatprotectionRulecategoryResponse) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionRulecategoryResponse) GetResult() ThreatprotectionRulecategory`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionRulecategoryResponse) GetResultOk() (*ThreatprotectionRulecategory, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionRulecategoryResponse) SetResult(v ThreatprotectionRulecategory)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionRulecategoryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


