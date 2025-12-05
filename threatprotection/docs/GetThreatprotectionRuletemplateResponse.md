# GetThreatprotectionRuletemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowedActions** | Pointer to **[]string** | The list of allowed actions of rhe rule template. | [optional] [readonly] 
**Category** | Pointer to **string** | The rule category this template assigned to. | [optional] [readonly] 
**DefaultConfig** | Pointer to [**ThreatprotectionRuletemplateDefaultConfig**](ThreatprotectionRuletemplateDefaultConfig.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the rule template. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the rule template. | [optional] [readonly] 
**Ruleset** | Pointer to **string** | The version of the ruleset the template assigned to. | [optional] [readonly] 
**Sid** | Pointer to **int64** | The Rule ID. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**ThreatprotectionRuletemplate**](ThreatprotectionRuletemplate.md) |  | [optional] 

## Methods

### NewGetThreatprotectionRuletemplateResponse

`func NewGetThreatprotectionRuletemplateResponse() *GetThreatprotectionRuletemplateResponse`

NewGetThreatprotectionRuletemplateResponse instantiates a new GetThreatprotectionRuletemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionRuletemplateResponseWithDefaults

`func NewGetThreatprotectionRuletemplateResponseWithDefaults() *GetThreatprotectionRuletemplateResponse`

NewGetThreatprotectionRuletemplateResponseWithDefaults instantiates a new GetThreatprotectionRuletemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionRuletemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionRuletemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionRuletemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionRuletemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowedActions

`func (o *GetThreatprotectionRuletemplateResponse) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *GetThreatprotectionRuletemplateResponse) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *GetThreatprotectionRuletemplateResponse) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *GetThreatprotectionRuletemplateResponse) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetCategory

`func (o *GetThreatprotectionRuletemplateResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetThreatprotectionRuletemplateResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetThreatprotectionRuletemplateResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetThreatprotectionRuletemplateResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultConfig

`func (o *GetThreatprotectionRuletemplateResponse) GetDefaultConfig() ThreatprotectionRuletemplateDefaultConfig`

GetDefaultConfig returns the DefaultConfig field if non-nil, zero value otherwise.

### GetDefaultConfigOk

`func (o *GetThreatprotectionRuletemplateResponse) GetDefaultConfigOk() (*ThreatprotectionRuletemplateDefaultConfig, bool)`

GetDefaultConfigOk returns a tuple with the DefaultConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfig

`func (o *GetThreatprotectionRuletemplateResponse) SetDefaultConfig(v ThreatprotectionRuletemplateDefaultConfig)`

SetDefaultConfig sets DefaultConfig field to given value.

### HasDefaultConfig

`func (o *GetThreatprotectionRuletemplateResponse) HasDefaultConfig() bool`

HasDefaultConfig returns a boolean if a field has been set.

### GetDescription

`func (o *GetThreatprotectionRuletemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetThreatprotectionRuletemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetThreatprotectionRuletemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetThreatprotectionRuletemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GetThreatprotectionRuletemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetThreatprotectionRuletemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetThreatprotectionRuletemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetThreatprotectionRuletemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleset

`func (o *GetThreatprotectionRuletemplateResponse) GetRuleset() string`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *GetThreatprotectionRuletemplateResponse) GetRulesetOk() (*string, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *GetThreatprotectionRuletemplateResponse) SetRuleset(v string)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *GetThreatprotectionRuletemplateResponse) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSid

`func (o *GetThreatprotectionRuletemplateResponse) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *GetThreatprotectionRuletemplateResponse) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *GetThreatprotectionRuletemplateResponse) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *GetThreatprotectionRuletemplateResponse) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUuid

`func (o *GetThreatprotectionRuletemplateResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetThreatprotectionRuletemplateResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetThreatprotectionRuletemplateResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetThreatprotectionRuletemplateResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionRuletemplateResponse) GetResult() ThreatprotectionRuletemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionRuletemplateResponse) GetResultOk() (*ThreatprotectionRuletemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionRuletemplateResponse) SetResult(v ThreatprotectionRuletemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionRuletemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


