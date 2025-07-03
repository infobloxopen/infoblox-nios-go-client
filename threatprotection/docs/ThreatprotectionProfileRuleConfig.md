# ThreatprotectionProfileRuleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The rule action. | [optional] 
**LogSeverity** | Pointer to **string** | The rule log severity. | [optional] 
**Params** | Pointer to [**[]ThreatprotectionprofileruleconfigParams**](ThreatprotectionprofileruleconfigParams.md) | The threat protection rule parameters. | [optional] 

## Methods

### NewThreatprotectionProfileRuleConfig

`func NewThreatprotectionProfileRuleConfig() *ThreatprotectionProfileRuleConfig`

NewThreatprotectionProfileRuleConfig instantiates a new ThreatprotectionProfileRuleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionProfileRuleConfigWithDefaults

`func NewThreatprotectionProfileRuleConfigWithDefaults() *ThreatprotectionProfileRuleConfig`

NewThreatprotectionProfileRuleConfigWithDefaults instantiates a new ThreatprotectionProfileRuleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ThreatprotectionProfileRuleConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ThreatprotectionProfileRuleConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ThreatprotectionProfileRuleConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ThreatprotectionProfileRuleConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetLogSeverity

`func (o *ThreatprotectionProfileRuleConfig) GetLogSeverity() string`

GetLogSeverity returns the LogSeverity field if non-nil, zero value otherwise.

### GetLogSeverityOk

`func (o *ThreatprotectionProfileRuleConfig) GetLogSeverityOk() (*string, bool)`

GetLogSeverityOk returns a tuple with the LogSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeverity

`func (o *ThreatprotectionProfileRuleConfig) SetLogSeverity(v string)`

SetLogSeverity sets LogSeverity field to given value.

### HasLogSeverity

`func (o *ThreatprotectionProfileRuleConfig) HasLogSeverity() bool`

HasLogSeverity returns a boolean if a field has been set.

### GetParams

`func (o *ThreatprotectionProfileRuleConfig) GetParams() []ThreatprotectionprofileruleconfigParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ThreatprotectionProfileRuleConfig) GetParamsOk() (*[]ThreatprotectionprofileruleconfigParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ThreatprotectionProfileRuleConfig) SetParams(v []ThreatprotectionprofileruleconfigParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *ThreatprotectionProfileRuleConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


