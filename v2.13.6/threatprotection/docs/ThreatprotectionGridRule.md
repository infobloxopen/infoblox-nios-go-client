# ThreatprotectionGridRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowedActions** | Pointer to **[]string** | The list of allowed actions of the custom rule. | [optional] [readonly] 
**Category** | Pointer to **string** | The rule category the custom rule assigned to. | [optional] [readonly] 
**Comment** | Pointer to **string** | The human readable comment for the custom rule. | [optional] 
**Config** | Pointer to [**ThreatprotectionGridRuleConfig**](ThreatprotectionGridRuleConfig.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the custom rule. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Determines if the custom rule is disabled. | [optional] 
**IsFactoryResetEnabled** | Pointer to **bool** | Determines if factory reset is enabled for the custom rule. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the rule custom rule concatenated with its rule config parameters. | [optional] [readonly] 
**Ruleset** | Pointer to **string** | The version of the ruleset the custom rule assigned to. | [optional] [readonly] 
**Sid** | Pointer to **int64** | The Rule ID. | [optional] [readonly] 
**Template** | Pointer to **string** | The threat protection rule template used to create this rule. | [optional] 
**Type** | Pointer to **string** | The type of the custom rule. | [optional] [readonly] 

## Methods

### NewThreatprotectionGridRule

`func NewThreatprotectionGridRule() *ThreatprotectionGridRule`

NewThreatprotectionGridRule instantiates a new ThreatprotectionGridRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionGridRuleWithDefaults

`func NewThreatprotectionGridRuleWithDefaults() *ThreatprotectionGridRule`

NewThreatprotectionGridRuleWithDefaults instantiates a new ThreatprotectionGridRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatprotectionGridRule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatprotectionGridRule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatprotectionGridRule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatprotectionGridRule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowedActions

`func (o *ThreatprotectionGridRule) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ThreatprotectionGridRule) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ThreatprotectionGridRule) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ThreatprotectionGridRule) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetCategory

`func (o *ThreatprotectionGridRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ThreatprotectionGridRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ThreatprotectionGridRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ThreatprotectionGridRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetComment

`func (o *ThreatprotectionGridRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ThreatprotectionGridRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ThreatprotectionGridRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ThreatprotectionGridRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConfig

`func (o *ThreatprotectionGridRule) GetConfig() ThreatprotectionGridRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ThreatprotectionGridRule) GetConfigOk() (*ThreatprotectionGridRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ThreatprotectionGridRule) SetConfig(v ThreatprotectionGridRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ThreatprotectionGridRule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatprotectionGridRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatprotectionGridRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatprotectionGridRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatprotectionGridRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *ThreatprotectionGridRule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ThreatprotectionGridRule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ThreatprotectionGridRule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ThreatprotectionGridRule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetIsFactoryResetEnabled

`func (o *ThreatprotectionGridRule) GetIsFactoryResetEnabled() bool`

GetIsFactoryResetEnabled returns the IsFactoryResetEnabled field if non-nil, zero value otherwise.

### GetIsFactoryResetEnabledOk

`func (o *ThreatprotectionGridRule) GetIsFactoryResetEnabledOk() (*bool, bool)`

GetIsFactoryResetEnabledOk returns a tuple with the IsFactoryResetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactoryResetEnabled

`func (o *ThreatprotectionGridRule) SetIsFactoryResetEnabled(v bool)`

SetIsFactoryResetEnabled sets IsFactoryResetEnabled field to given value.

### HasIsFactoryResetEnabled

`func (o *ThreatprotectionGridRule) HasIsFactoryResetEnabled() bool`

HasIsFactoryResetEnabled returns a boolean if a field has been set.

### GetName

`func (o *ThreatprotectionGridRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatprotectionGridRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatprotectionGridRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreatprotectionGridRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleset

`func (o *ThreatprotectionGridRule) GetRuleset() string`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *ThreatprotectionGridRule) GetRulesetOk() (*string, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *ThreatprotectionGridRule) SetRuleset(v string)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *ThreatprotectionGridRule) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSid

`func (o *ThreatprotectionGridRule) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ThreatprotectionGridRule) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ThreatprotectionGridRule) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ThreatprotectionGridRule) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetTemplate

`func (o *ThreatprotectionGridRule) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ThreatprotectionGridRule) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ThreatprotectionGridRule) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ThreatprotectionGridRule) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *ThreatprotectionGridRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreatprotectionGridRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreatprotectionGridRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreatprotectionGridRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


