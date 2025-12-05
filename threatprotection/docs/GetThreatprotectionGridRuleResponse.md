# GetThreatprotectionGridRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**ThreatprotectionGridRule**](ThreatprotectionGridRule.md) |  | [optional] 

## Methods

### NewGetThreatprotectionGridRuleResponse

`func NewGetThreatprotectionGridRuleResponse() *GetThreatprotectionGridRuleResponse`

NewGetThreatprotectionGridRuleResponse instantiates a new GetThreatprotectionGridRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionGridRuleResponseWithDefaults

`func NewGetThreatprotectionGridRuleResponseWithDefaults() *GetThreatprotectionGridRuleResponse`

NewGetThreatprotectionGridRuleResponseWithDefaults instantiates a new GetThreatprotectionGridRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionGridRuleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionGridRuleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionGridRuleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionGridRuleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowedActions

`func (o *GetThreatprotectionGridRuleResponse) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *GetThreatprotectionGridRuleResponse) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *GetThreatprotectionGridRuleResponse) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *GetThreatprotectionGridRuleResponse) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetCategory

`func (o *GetThreatprotectionGridRuleResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetThreatprotectionGridRuleResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetThreatprotectionGridRuleResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetThreatprotectionGridRuleResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetComment

`func (o *GetThreatprotectionGridRuleResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetThreatprotectionGridRuleResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetThreatprotectionGridRuleResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetThreatprotectionGridRuleResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConfig

`func (o *GetThreatprotectionGridRuleResponse) GetConfig() ThreatprotectionGridRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetThreatprotectionGridRuleResponse) GetConfigOk() (*ThreatprotectionGridRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetThreatprotectionGridRuleResponse) SetConfig(v ThreatprotectionGridRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetThreatprotectionGridRuleResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *GetThreatprotectionGridRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetThreatprotectionGridRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetThreatprotectionGridRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetThreatprotectionGridRuleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *GetThreatprotectionGridRuleResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetThreatprotectionGridRuleResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetThreatprotectionGridRuleResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetThreatprotectionGridRuleResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetIsFactoryResetEnabled

`func (o *GetThreatprotectionGridRuleResponse) GetIsFactoryResetEnabled() bool`

GetIsFactoryResetEnabled returns the IsFactoryResetEnabled field if non-nil, zero value otherwise.

### GetIsFactoryResetEnabledOk

`func (o *GetThreatprotectionGridRuleResponse) GetIsFactoryResetEnabledOk() (*bool, bool)`

GetIsFactoryResetEnabledOk returns a tuple with the IsFactoryResetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactoryResetEnabled

`func (o *GetThreatprotectionGridRuleResponse) SetIsFactoryResetEnabled(v bool)`

SetIsFactoryResetEnabled sets IsFactoryResetEnabled field to given value.

### HasIsFactoryResetEnabled

`func (o *GetThreatprotectionGridRuleResponse) HasIsFactoryResetEnabled() bool`

HasIsFactoryResetEnabled returns a boolean if a field has been set.

### GetName

`func (o *GetThreatprotectionGridRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetThreatprotectionGridRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetThreatprotectionGridRuleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetThreatprotectionGridRuleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleset

`func (o *GetThreatprotectionGridRuleResponse) GetRuleset() string`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *GetThreatprotectionGridRuleResponse) GetRulesetOk() (*string, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *GetThreatprotectionGridRuleResponse) SetRuleset(v string)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *GetThreatprotectionGridRuleResponse) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSid

`func (o *GetThreatprotectionGridRuleResponse) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *GetThreatprotectionGridRuleResponse) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *GetThreatprotectionGridRuleResponse) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *GetThreatprotectionGridRuleResponse) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetTemplate

`func (o *GetThreatprotectionGridRuleResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetThreatprotectionGridRuleResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetThreatprotectionGridRuleResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetThreatprotectionGridRuleResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *GetThreatprotectionGridRuleResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetThreatprotectionGridRuleResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetThreatprotectionGridRuleResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetThreatprotectionGridRuleResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *GetThreatprotectionGridRuleResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetThreatprotectionGridRuleResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetThreatprotectionGridRuleResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetThreatprotectionGridRuleResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionGridRuleResponse) GetResult() ThreatprotectionGridRule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionGridRuleResponse) GetResultOk() (*ThreatprotectionGridRule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionGridRuleResponse) SetResult(v ThreatprotectionGridRule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionGridRuleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


