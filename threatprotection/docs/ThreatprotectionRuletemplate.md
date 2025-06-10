# ThreatprotectionRuletemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowedActions** | Pointer to **[]string** | The list of allowed actions of rhe rule template. | [optional] [readonly] 
**Category** | Pointer to **string** | The rule category this template assigned to. | [optional] [readonly] 
**DefaultConfig** | Pointer to [**ThreatprotectionRuletemplateDefaultConfig**](ThreatprotectionRuletemplateDefaultConfig.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the rule template. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the rule template. | [optional] [readonly] 
**Ruleset** | Pointer to **string** | The version of the ruleset the template assigned to. | [optional] [readonly] 
**Sid** | Pointer to **int64** | The Rule ID. | [optional] [readonly] 

## Methods

### NewThreatprotectionRuletemplate

`func NewThreatprotectionRuletemplate() *ThreatprotectionRuletemplate`

NewThreatprotectionRuletemplate instantiates a new ThreatprotectionRuletemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionRuletemplateWithDefaults

`func NewThreatprotectionRuletemplateWithDefaults() *ThreatprotectionRuletemplate`

NewThreatprotectionRuletemplateWithDefaults instantiates a new ThreatprotectionRuletemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatprotectionRuletemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatprotectionRuletemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatprotectionRuletemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatprotectionRuletemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowedActions

`func (o *ThreatprotectionRuletemplate) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ThreatprotectionRuletemplate) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ThreatprotectionRuletemplate) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ThreatprotectionRuletemplate) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetCategory

`func (o *ThreatprotectionRuletemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ThreatprotectionRuletemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ThreatprotectionRuletemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ThreatprotectionRuletemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultConfig

`func (o *ThreatprotectionRuletemplate) GetDefaultConfig() ThreatprotectionRuletemplateDefaultConfig`

GetDefaultConfig returns the DefaultConfig field if non-nil, zero value otherwise.

### GetDefaultConfigOk

`func (o *ThreatprotectionRuletemplate) GetDefaultConfigOk() (*ThreatprotectionRuletemplateDefaultConfig, bool)`

GetDefaultConfigOk returns a tuple with the DefaultConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfig

`func (o *ThreatprotectionRuletemplate) SetDefaultConfig(v ThreatprotectionRuletemplateDefaultConfig)`

SetDefaultConfig sets DefaultConfig field to given value.

### HasDefaultConfig

`func (o *ThreatprotectionRuletemplate) HasDefaultConfig() bool`

HasDefaultConfig returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatprotectionRuletemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatprotectionRuletemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatprotectionRuletemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatprotectionRuletemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ThreatprotectionRuletemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatprotectionRuletemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatprotectionRuletemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreatprotectionRuletemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleset

`func (o *ThreatprotectionRuletemplate) GetRuleset() string`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *ThreatprotectionRuletemplate) GetRulesetOk() (*string, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *ThreatprotectionRuletemplate) SetRuleset(v string)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *ThreatprotectionRuletemplate) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSid

`func (o *ThreatprotectionRuletemplate) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ThreatprotectionRuletemplate) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ThreatprotectionRuletemplate) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ThreatprotectionRuletemplate) HasSid() bool`

HasSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


