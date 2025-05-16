# Bulkhostnametemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**IsGridDefault** | Pointer to **bool** | True if this template is Grid default. | [optional] [readonly] 
**PreDefined** | Pointer to **bool** | True if this is a pre-defined template, False otherwise. | [optional] [readonly] 
**TemplateFormat** | Pointer to **string** | The format of bulk host name template. It should follow certain rules (please use Administration Guide as reference). | [optional] 
**TemplateName** | Pointer to **string** | The name of bulk host name template. | [optional] 

## Methods

### NewBulkhostnametemplate

`func NewBulkhostnametemplate() *Bulkhostnametemplate`

NewBulkhostnametemplate instantiates a new Bulkhostnametemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkhostnametemplateWithDefaults

`func NewBulkhostnametemplateWithDefaults() *Bulkhostnametemplate`

NewBulkhostnametemplateWithDefaults instantiates a new Bulkhostnametemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Bulkhostnametemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Bulkhostnametemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Bulkhostnametemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Bulkhostnametemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIsGridDefault

`func (o *Bulkhostnametemplate) GetIsGridDefault() bool`

GetIsGridDefault returns the IsGridDefault field if non-nil, zero value otherwise.

### GetIsGridDefaultOk

`func (o *Bulkhostnametemplate) GetIsGridDefaultOk() (*bool, bool)`

GetIsGridDefaultOk returns a tuple with the IsGridDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGridDefault

`func (o *Bulkhostnametemplate) SetIsGridDefault(v bool)`

SetIsGridDefault sets IsGridDefault field to given value.

### HasIsGridDefault

`func (o *Bulkhostnametemplate) HasIsGridDefault() bool`

HasIsGridDefault returns a boolean if a field has been set.

### GetPreDefined

`func (o *Bulkhostnametemplate) GetPreDefined() bool`

GetPreDefined returns the PreDefined field if non-nil, zero value otherwise.

### GetPreDefinedOk

`func (o *Bulkhostnametemplate) GetPreDefinedOk() (*bool, bool)`

GetPreDefinedOk returns a tuple with the PreDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDefined

`func (o *Bulkhostnametemplate) SetPreDefined(v bool)`

SetPreDefined sets PreDefined field to given value.

### HasPreDefined

`func (o *Bulkhostnametemplate) HasPreDefined() bool`

HasPreDefined returns a boolean if a field has been set.

### GetTemplateFormat

`func (o *Bulkhostnametemplate) GetTemplateFormat() string`

GetTemplateFormat returns the TemplateFormat field if non-nil, zero value otherwise.

### GetTemplateFormatOk

`func (o *Bulkhostnametemplate) GetTemplateFormatOk() (*string, bool)`

GetTemplateFormatOk returns a tuple with the TemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFormat

`func (o *Bulkhostnametemplate) SetTemplateFormat(v string)`

SetTemplateFormat sets TemplateFormat field to given value.

### HasTemplateFormat

`func (o *Bulkhostnametemplate) HasTemplateFormat() bool`

HasTemplateFormat returns a boolean if a field has been set.

### GetTemplateName

`func (o *Bulkhostnametemplate) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *Bulkhostnametemplate) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *Bulkhostnametemplate) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *Bulkhostnametemplate) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


