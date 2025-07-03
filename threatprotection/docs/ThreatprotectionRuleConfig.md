# ThreatprotectionRuleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The rule action. | [optional] 
**LogSeverity** | Pointer to **string** | The rule log severity. | [optional] 
**Params** | Pointer to [**[]ThreatprotectionruleconfigParams**](ThreatprotectionruleconfigParams.md) | The threat protection rule parameters. | [optional] 

## Methods

### NewThreatprotectionRuleConfig

`func NewThreatprotectionRuleConfig() *ThreatprotectionRuleConfig`

NewThreatprotectionRuleConfig instantiates a new ThreatprotectionRuleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionRuleConfigWithDefaults

`func NewThreatprotectionRuleConfigWithDefaults() *ThreatprotectionRuleConfig`

NewThreatprotectionRuleConfigWithDefaults instantiates a new ThreatprotectionRuleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ThreatprotectionRuleConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ThreatprotectionRuleConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ThreatprotectionRuleConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ThreatprotectionRuleConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetLogSeverity

`func (o *ThreatprotectionRuleConfig) GetLogSeverity() string`

GetLogSeverity returns the LogSeverity field if non-nil, zero value otherwise.

### GetLogSeverityOk

`func (o *ThreatprotectionRuleConfig) GetLogSeverityOk() (*string, bool)`

GetLogSeverityOk returns a tuple with the LogSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeverity

`func (o *ThreatprotectionRuleConfig) SetLogSeverity(v string)`

SetLogSeverity sets LogSeverity field to given value.

### HasLogSeverity

`func (o *ThreatprotectionRuleConfig) HasLogSeverity() bool`

HasLogSeverity returns a boolean if a field has been set.

### GetParams

`func (o *ThreatprotectionRuleConfig) GetParams() []ThreatprotectionruleconfigParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ThreatprotectionRuleConfig) GetParamsOk() (*[]ThreatprotectionruleconfigParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ThreatprotectionRuleConfig) SetParams(v []ThreatprotectionruleconfigParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *ThreatprotectionRuleConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


