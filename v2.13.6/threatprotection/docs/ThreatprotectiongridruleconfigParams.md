# ThreatprotectiongridruleconfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The rule parameter name. | [optional] 
**Description** | Pointer to **string** | The rule parameter description. | [optional] [readonly] 
**Syntax** | Pointer to **string** | The rule parameter syntax. | [optional] [readonly] 
**Value** | Pointer to **string** | The rule parameter value. | [optional] 
**Min** | Pointer to **int64** | The rule parameter minimum. | [optional] [readonly] 
**Max** | Pointer to **int64** | The rule parameter maximum. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | Determines if parameter value is editable at member level. | [optional] [readonly] 
**EnumValues** | Pointer to **[]string** | The rule parameter enum values. | [optional] [readonly] 

## Methods

### NewThreatprotectiongridruleconfigParams

`func NewThreatprotectiongridruleconfigParams() *ThreatprotectiongridruleconfigParams`

NewThreatprotectiongridruleconfigParams instantiates a new ThreatprotectiongridruleconfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectiongridruleconfigParamsWithDefaults

`func NewThreatprotectiongridruleconfigParamsWithDefaults() *ThreatprotectiongridruleconfigParams`

NewThreatprotectiongridruleconfigParamsWithDefaults instantiates a new ThreatprotectiongridruleconfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ThreatprotectiongridruleconfigParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatprotectiongridruleconfigParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatprotectiongridruleconfigParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreatprotectiongridruleconfigParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatprotectiongridruleconfigParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatprotectiongridruleconfigParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatprotectiongridruleconfigParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatprotectiongridruleconfigParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSyntax

`func (o *ThreatprotectiongridruleconfigParams) GetSyntax() string`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *ThreatprotectiongridruleconfigParams) GetSyntaxOk() (*string, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *ThreatprotectiongridruleconfigParams) SetSyntax(v string)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *ThreatprotectiongridruleconfigParams) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.

### GetValue

`func (o *ThreatprotectiongridruleconfigParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ThreatprotectiongridruleconfigParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ThreatprotectiongridruleconfigParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ThreatprotectiongridruleconfigParams) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMin

`func (o *ThreatprotectiongridruleconfigParams) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ThreatprotectiongridruleconfigParams) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ThreatprotectiongridruleconfigParams) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *ThreatprotectiongridruleconfigParams) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *ThreatprotectiongridruleconfigParams) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ThreatprotectiongridruleconfigParams) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ThreatprotectiongridruleconfigParams) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ThreatprotectiongridruleconfigParams) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetReadOnly

`func (o *ThreatprotectiongridruleconfigParams) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ThreatprotectiongridruleconfigParams) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ThreatprotectiongridruleconfigParams) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ThreatprotectiongridruleconfigParams) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetEnumValues

`func (o *ThreatprotectiongridruleconfigParams) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *ThreatprotectiongridruleconfigParams) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *ThreatprotectiongridruleconfigParams) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *ThreatprotectiongridruleconfigParams) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


