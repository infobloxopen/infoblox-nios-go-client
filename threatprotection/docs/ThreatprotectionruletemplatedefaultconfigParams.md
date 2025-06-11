# ThreatprotectionruletemplatedefaultconfigParams

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

### NewThreatprotectionruletemplatedefaultconfigParams

`func NewThreatprotectionruletemplatedefaultconfigParams() *ThreatprotectionruletemplatedefaultconfigParams`

NewThreatprotectionruletemplatedefaultconfigParams instantiates a new ThreatprotectionruletemplatedefaultconfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionruletemplatedefaultconfigParamsWithDefaults

`func NewThreatprotectionruletemplatedefaultconfigParamsWithDefaults() *ThreatprotectionruletemplatedefaultconfigParams`

NewThreatprotectionruletemplatedefaultconfigParamsWithDefaults instantiates a new ThreatprotectionruletemplatedefaultconfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSyntax

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetSyntax() string`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetSyntaxOk() (*string, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetSyntax(v string)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.

### GetValue

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMin

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetReadOnly

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetEnumValues

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *ThreatprotectionruletemplatedefaultconfigParams) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *ThreatprotectionruletemplatedefaultconfigParams) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *ThreatprotectionruletemplatedefaultconfigParams) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


