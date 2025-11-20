# ThreatprotectionProfileRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Config** | Pointer to [**ThreatprotectionProfileRuleConfig**](ThreatprotectionProfileRuleConfig.md) |  | [optional] 
**Disable** | Pointer to **bool** | Determines if the rule is enabled or not for the profile. | [optional] 
**Profile** | Pointer to **string** | The name of the Threat protection profile. | [optional] [readonly] 
**Rule** | Pointer to **string** | The rule object name. | [optional] [readonly] 
**Sid** | Pointer to **int64** | The snort rule ID. | [optional] [readonly] 
**UseConfig** | Pointer to **bool** | Use flag for: config | [optional] 
**UseDisable** | Pointer to **bool** | Use flag for: disable | [optional] 

## Methods

### NewThreatprotectionProfileRule

`func NewThreatprotectionProfileRule() *ThreatprotectionProfileRule`

NewThreatprotectionProfileRule instantiates a new ThreatprotectionProfileRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionProfileRuleWithDefaults

`func NewThreatprotectionProfileRuleWithDefaults() *ThreatprotectionProfileRule`

NewThreatprotectionProfileRuleWithDefaults instantiates a new ThreatprotectionProfileRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatprotectionProfileRule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatprotectionProfileRule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatprotectionProfileRule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatprotectionProfileRule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConfig

`func (o *ThreatprotectionProfileRule) GetConfig() ThreatprotectionProfileRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ThreatprotectionProfileRule) GetConfigOk() (*ThreatprotectionProfileRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ThreatprotectionProfileRule) SetConfig(v ThreatprotectionProfileRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ThreatprotectionProfileRule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisable

`func (o *ThreatprotectionProfileRule) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ThreatprotectionProfileRule) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ThreatprotectionProfileRule) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ThreatprotectionProfileRule) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetProfile

`func (o *ThreatprotectionProfileRule) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ThreatprotectionProfileRule) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ThreatprotectionProfileRule) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ThreatprotectionProfileRule) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRule

`func (o *ThreatprotectionProfileRule) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ThreatprotectionProfileRule) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ThreatprotectionProfileRule) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *ThreatprotectionProfileRule) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSid

`func (o *ThreatprotectionProfileRule) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ThreatprotectionProfileRule) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ThreatprotectionProfileRule) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ThreatprotectionProfileRule) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUseConfig

`func (o *ThreatprotectionProfileRule) GetUseConfig() bool`

GetUseConfig returns the UseConfig field if non-nil, zero value otherwise.

### GetUseConfigOk

`func (o *ThreatprotectionProfileRule) GetUseConfigOk() (*bool, bool)`

GetUseConfigOk returns a tuple with the UseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseConfig

`func (o *ThreatprotectionProfileRule) SetUseConfig(v bool)`

SetUseConfig sets UseConfig field to given value.

### HasUseConfig

`func (o *ThreatprotectionProfileRule) HasUseConfig() bool`

HasUseConfig returns a boolean if a field has been set.

### GetUseDisable

`func (o *ThreatprotectionProfileRule) GetUseDisable() bool`

GetUseDisable returns the UseDisable field if non-nil, zero value otherwise.

### GetUseDisableOk

`func (o *ThreatprotectionProfileRule) GetUseDisableOk() (*bool, bool)`

GetUseDisableOk returns a tuple with the UseDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisable

`func (o *ThreatprotectionProfileRule) SetUseDisable(v bool)`

SetUseDisable sets UseDisable field to given value.

### HasUseDisable

`func (o *ThreatprotectionProfileRule) HasUseDisable() bool`

HasUseDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


