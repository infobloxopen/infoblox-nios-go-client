# GetThreatprotectionRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Config** | Pointer to [**ThreatprotectionRuleConfig**](ThreatprotectionRuleConfig.md) |  | [optional] 
**Disable** | Pointer to **bool** | Determines if the rule is enabled or not for the member. | [optional] 
**Member** | Pointer to **string** | The name of the Threat protection member. | [optional] [readonly] 
**Rule** | Pointer to **string** | The rule object name. | [optional] [readonly] 
**Sid** | Pointer to **int64** | The Rule ID. | [optional] [readonly] 
**UseConfig** | Pointer to **bool** | Use flag for: config | [optional] 
**UseDisable** | Pointer to **bool** | Use flag for: disable | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**ThreatprotectionRule**](ThreatprotectionRule.md) |  | [optional] 

## Methods

### NewGetThreatprotectionRuleResponse

`func NewGetThreatprotectionRuleResponse() *GetThreatprotectionRuleResponse`

NewGetThreatprotectionRuleResponse instantiates a new GetThreatprotectionRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionRuleResponseWithDefaults

`func NewGetThreatprotectionRuleResponseWithDefaults() *GetThreatprotectionRuleResponse`

NewGetThreatprotectionRuleResponseWithDefaults instantiates a new GetThreatprotectionRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionRuleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionRuleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionRuleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionRuleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConfig

`func (o *GetThreatprotectionRuleResponse) GetConfig() ThreatprotectionRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetThreatprotectionRuleResponse) GetConfigOk() (*ThreatprotectionRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetThreatprotectionRuleResponse) SetConfig(v ThreatprotectionRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetThreatprotectionRuleResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisable

`func (o *GetThreatprotectionRuleResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetThreatprotectionRuleResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetThreatprotectionRuleResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetThreatprotectionRuleResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetMember

`func (o *GetThreatprotectionRuleResponse) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetThreatprotectionRuleResponse) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetThreatprotectionRuleResponse) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetThreatprotectionRuleResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetRule

`func (o *GetThreatprotectionRuleResponse) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GetThreatprotectionRuleResponse) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GetThreatprotectionRuleResponse) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *GetThreatprotectionRuleResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSid

`func (o *GetThreatprotectionRuleResponse) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *GetThreatprotectionRuleResponse) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *GetThreatprotectionRuleResponse) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *GetThreatprotectionRuleResponse) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUseConfig

`func (o *GetThreatprotectionRuleResponse) GetUseConfig() bool`

GetUseConfig returns the UseConfig field if non-nil, zero value otherwise.

### GetUseConfigOk

`func (o *GetThreatprotectionRuleResponse) GetUseConfigOk() (*bool, bool)`

GetUseConfigOk returns a tuple with the UseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseConfig

`func (o *GetThreatprotectionRuleResponse) SetUseConfig(v bool)`

SetUseConfig sets UseConfig field to given value.

### HasUseConfig

`func (o *GetThreatprotectionRuleResponse) HasUseConfig() bool`

HasUseConfig returns a boolean if a field has been set.

### GetUseDisable

`func (o *GetThreatprotectionRuleResponse) GetUseDisable() bool`

GetUseDisable returns the UseDisable field if non-nil, zero value otherwise.

### GetUseDisableOk

`func (o *GetThreatprotectionRuleResponse) GetUseDisableOk() (*bool, bool)`

GetUseDisableOk returns a tuple with the UseDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisable

`func (o *GetThreatprotectionRuleResponse) SetUseDisable(v bool)`

SetUseDisable sets UseDisable field to given value.

### HasUseDisable

`func (o *GetThreatprotectionRuleResponse) HasUseDisable() bool`

HasUseDisable returns a boolean if a field has been set.

### GetUuid

`func (o *GetThreatprotectionRuleResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetThreatprotectionRuleResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetThreatprotectionRuleResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetThreatprotectionRuleResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionRuleResponse) GetResult() ThreatprotectionRule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionRuleResponse) GetResultOk() (*ThreatprotectionRule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionRuleResponse) SetResult(v ThreatprotectionRule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionRuleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


