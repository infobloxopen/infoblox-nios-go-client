# GetThreatprotectionProfileRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Config** | Pointer to [**ThreatprotectionProfileRuleConfig**](ThreatprotectionProfileRuleConfig.md) |  | [optional] 
**Disable** | Pointer to **bool** | Determines if the rule is enabled or not for the profile. | [optional] 
**Profile** | Pointer to **string** | The name of the Threat protection profile. | [optional] [readonly] 
**Rule** | Pointer to **string** | The rule object name. | [optional] [readonly] 
**Sid** | Pointer to **int64** | The snort rule ID. | [optional] [readonly] 
**UseConfig** | Pointer to **bool** | Use flag for: config | [optional] 
**UseDisable** | Pointer to **bool** | Use flag for: disable | [optional] 
**Result** | Pointer to [**ThreatprotectionProfileRule**](ThreatprotectionProfileRule.md) |  | [optional] 

## Methods

### NewGetThreatprotectionProfileRuleResponse

`func NewGetThreatprotectionProfileRuleResponse() *GetThreatprotectionProfileRuleResponse`

NewGetThreatprotectionProfileRuleResponse instantiates a new GetThreatprotectionProfileRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionProfileRuleResponseWithDefaults

`func NewGetThreatprotectionProfileRuleResponseWithDefaults() *GetThreatprotectionProfileRuleResponse`

NewGetThreatprotectionProfileRuleResponseWithDefaults instantiates a new GetThreatprotectionProfileRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionProfileRuleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionProfileRuleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionProfileRuleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionProfileRuleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConfig

`func (o *GetThreatprotectionProfileRuleResponse) GetConfig() ThreatprotectionProfileRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetThreatprotectionProfileRuleResponse) GetConfigOk() (*ThreatprotectionProfileRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetThreatprotectionProfileRuleResponse) SetConfig(v ThreatprotectionProfileRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetThreatprotectionProfileRuleResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisable

`func (o *GetThreatprotectionProfileRuleResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetThreatprotectionProfileRuleResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetThreatprotectionProfileRuleResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetThreatprotectionProfileRuleResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetProfile

`func (o *GetThreatprotectionProfileRuleResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetThreatprotectionProfileRuleResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetThreatprotectionProfileRuleResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetThreatprotectionProfileRuleResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRule

`func (o *GetThreatprotectionProfileRuleResponse) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GetThreatprotectionProfileRuleResponse) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GetThreatprotectionProfileRuleResponse) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *GetThreatprotectionProfileRuleResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSid

`func (o *GetThreatprotectionProfileRuleResponse) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *GetThreatprotectionProfileRuleResponse) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *GetThreatprotectionProfileRuleResponse) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *GetThreatprotectionProfileRuleResponse) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUseConfig

`func (o *GetThreatprotectionProfileRuleResponse) GetUseConfig() bool`

GetUseConfig returns the UseConfig field if non-nil, zero value otherwise.

### GetUseConfigOk

`func (o *GetThreatprotectionProfileRuleResponse) GetUseConfigOk() (*bool, bool)`

GetUseConfigOk returns a tuple with the UseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseConfig

`func (o *GetThreatprotectionProfileRuleResponse) SetUseConfig(v bool)`

SetUseConfig sets UseConfig field to given value.

### HasUseConfig

`func (o *GetThreatprotectionProfileRuleResponse) HasUseConfig() bool`

HasUseConfig returns a boolean if a field has been set.

### GetUseDisable

`func (o *GetThreatprotectionProfileRuleResponse) GetUseDisable() bool`

GetUseDisable returns the UseDisable field if non-nil, zero value otherwise.

### GetUseDisableOk

`func (o *GetThreatprotectionProfileRuleResponse) GetUseDisableOk() (*bool, bool)`

GetUseDisableOk returns a tuple with the UseDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisable

`func (o *GetThreatprotectionProfileRuleResponse) SetUseDisable(v bool)`

SetUseDisable sets UseDisable field to given value.

### HasUseDisable

`func (o *GetThreatprotectionProfileRuleResponse) HasUseDisable() bool`

HasUseDisable returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionProfileRuleResponse) GetResult() ThreatprotectionProfileRule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionProfileRuleResponse) GetResultOk() (*ThreatprotectionProfileRule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionProfileRuleResponse) SetResult(v ThreatprotectionProfileRule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionProfileRuleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


