# ThreatprotectionruleconfigParams

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

### NewThreatprotectionruleconfigParams

`func NewThreatprotectionruleconfigParams() *ThreatprotectionruleconfigParams`

NewThreatprotectionruleconfigParams instantiates a new ThreatprotectionruleconfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionruleconfigParamsWithDefaults

`func NewThreatprotectionruleconfigParamsWithDefaults() *ThreatprotectionruleconfigParams`

NewThreatprotectionruleconfigParamsWithDefaults instantiates a new ThreatprotectionruleconfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ThreatprotectionruleconfigParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatprotectionruleconfigParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatprotectionruleconfigParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreatprotectionruleconfigParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatprotectionruleconfigParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatprotectionruleconfigParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatprotectionruleconfigParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatprotectionruleconfigParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSyntax

`func (o *ThreatprotectionruleconfigParams) GetSyntax() string`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *ThreatprotectionruleconfigParams) GetSyntaxOk() (*string, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *ThreatprotectionruleconfigParams) SetSyntax(v string)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *ThreatprotectionruleconfigParams) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.

### GetValue

`func (o *ThreatprotectionruleconfigParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ThreatprotectionruleconfigParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ThreatprotectionruleconfigParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ThreatprotectionruleconfigParams) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMin

`func (o *ThreatprotectionruleconfigParams) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ThreatprotectionruleconfigParams) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ThreatprotectionruleconfigParams) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *ThreatprotectionruleconfigParams) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *ThreatprotectionruleconfigParams) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ThreatprotectionruleconfigParams) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ThreatprotectionruleconfigParams) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ThreatprotectionruleconfigParams) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetReadOnly

`func (o *ThreatprotectionruleconfigParams) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ThreatprotectionruleconfigParams) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ThreatprotectionruleconfigParams) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ThreatprotectionruleconfigParams) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetEnumValues

`func (o *ThreatprotectionruleconfigParams) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *ThreatprotectionruleconfigParams) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *ThreatprotectionruleconfigParams) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *ThreatprotectionruleconfigParams) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


