# ThreatprotectionGridRuleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The rule action. | [optional] 
**LogSeverity** | Pointer to **string** | The rule log severity. | [optional] 
**Params** | Pointer to [**ThreatprotectiongridruleconfigParams**](ThreatprotectiongridruleconfigParams.md) |  | [optional] 

## Methods

### NewThreatprotectionGridRuleConfig

`func NewThreatprotectionGridRuleConfig() *ThreatprotectionGridRuleConfig`

NewThreatprotectionGridRuleConfig instantiates a new ThreatprotectionGridRuleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionGridRuleConfigWithDefaults

`func NewThreatprotectionGridRuleConfigWithDefaults() *ThreatprotectionGridRuleConfig`

NewThreatprotectionGridRuleConfigWithDefaults instantiates a new ThreatprotectionGridRuleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ThreatprotectionGridRuleConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ThreatprotectionGridRuleConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ThreatprotectionGridRuleConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ThreatprotectionGridRuleConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetLogSeverity

`func (o *ThreatprotectionGridRuleConfig) GetLogSeverity() string`

GetLogSeverity returns the LogSeverity field if non-nil, zero value otherwise.

### GetLogSeverityOk

`func (o *ThreatprotectionGridRuleConfig) GetLogSeverityOk() (*string, bool)`

GetLogSeverityOk returns a tuple with the LogSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeverity

`func (o *ThreatprotectionGridRuleConfig) SetLogSeverity(v string)`

SetLogSeverity sets LogSeverity field to given value.

### HasLogSeverity

`func (o *ThreatprotectionGridRuleConfig) HasLogSeverity() bool`

HasLogSeverity returns a boolean if a field has been set.

### GetParams

`func (o *ThreatprotectionGridRuleConfig) GetParams() ThreatprotectiongridruleconfigParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ThreatprotectionGridRuleConfig) GetParamsOk() (*ThreatprotectiongridruleconfigParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ThreatprotectionGridRuleConfig) SetParams(v ThreatprotectiongridruleconfigParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *ThreatprotectionGridRuleConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


