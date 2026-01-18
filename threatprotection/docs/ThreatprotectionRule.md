# ThreatprotectionRule

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

## Methods

### NewThreatprotectionRule

`func NewThreatprotectionRule() *ThreatprotectionRule`

NewThreatprotectionRule instantiates a new ThreatprotectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionRuleWithDefaults

`func NewThreatprotectionRuleWithDefaults() *ThreatprotectionRule`

NewThreatprotectionRuleWithDefaults instantiates a new ThreatprotectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatprotectionRule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatprotectionRule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatprotectionRule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatprotectionRule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConfig

`func (o *ThreatprotectionRule) GetConfig() ThreatprotectionRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ThreatprotectionRule) GetConfigOk() (*ThreatprotectionRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ThreatprotectionRule) SetConfig(v ThreatprotectionRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ThreatprotectionRule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisable

`func (o *ThreatprotectionRule) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ThreatprotectionRule) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ThreatprotectionRule) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ThreatprotectionRule) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetMember

`func (o *ThreatprotectionRule) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ThreatprotectionRule) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ThreatprotectionRule) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *ThreatprotectionRule) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetRule

`func (o *ThreatprotectionRule) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ThreatprotectionRule) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ThreatprotectionRule) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *ThreatprotectionRule) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSid

`func (o *ThreatprotectionRule) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ThreatprotectionRule) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ThreatprotectionRule) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ThreatprotectionRule) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUseConfig

`func (o *ThreatprotectionRule) GetUseConfig() bool`

GetUseConfig returns the UseConfig field if non-nil, zero value otherwise.

### GetUseConfigOk

`func (o *ThreatprotectionRule) GetUseConfigOk() (*bool, bool)`

GetUseConfigOk returns a tuple with the UseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseConfig

`func (o *ThreatprotectionRule) SetUseConfig(v bool)`

SetUseConfig sets UseConfig field to given value.

### HasUseConfig

`func (o *ThreatprotectionRule) HasUseConfig() bool`

HasUseConfig returns a boolean if a field has been set.

### GetUseDisable

`func (o *ThreatprotectionRule) GetUseDisable() bool`

GetUseDisable returns the UseDisable field if non-nil, zero value otherwise.

### GetUseDisableOk

`func (o *ThreatprotectionRule) GetUseDisableOk() (*bool, bool)`

GetUseDisableOk returns a tuple with the UseDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisable

`func (o *ThreatprotectionRule) SetUseDisable(v bool)`

SetUseDisable sets UseDisable field to given value.

### HasUseDisable

`func (o *ThreatprotectionRule) HasUseDisable() bool`

HasUseDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


