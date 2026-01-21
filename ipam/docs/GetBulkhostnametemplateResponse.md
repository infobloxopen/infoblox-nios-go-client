# GetBulkhostnametemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**IsGridDefault** | Pointer to **bool** | True if this template is Grid default. | [optional] [readonly] 
**PreDefined** | Pointer to **bool** | True if this is a pre-defined template, False otherwise. | [optional] [readonly] 
**TemplateFormat** | Pointer to **string** | The format of bulk host name template. It should follow certain rules (please use Administration Guide as reference). | [optional] 
**TemplateName** | Pointer to **string** | The name of bulk host name template. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Bulkhostnametemplate**](Bulkhostnametemplate.md) |  | [optional] 

## Methods

### NewGetBulkhostnametemplateResponse

`func NewGetBulkhostnametemplateResponse() *GetBulkhostnametemplateResponse`

NewGetBulkhostnametemplateResponse instantiates a new GetBulkhostnametemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBulkhostnametemplateResponseWithDefaults

`func NewGetBulkhostnametemplateResponseWithDefaults() *GetBulkhostnametemplateResponse`

NewGetBulkhostnametemplateResponseWithDefaults instantiates a new GetBulkhostnametemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetBulkhostnametemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetBulkhostnametemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetBulkhostnametemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetBulkhostnametemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIsGridDefault

`func (o *GetBulkhostnametemplateResponse) GetIsGridDefault() bool`

GetIsGridDefault returns the IsGridDefault field if non-nil, zero value otherwise.

### GetIsGridDefaultOk

`func (o *GetBulkhostnametemplateResponse) GetIsGridDefaultOk() (*bool, bool)`

GetIsGridDefaultOk returns a tuple with the IsGridDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGridDefault

`func (o *GetBulkhostnametemplateResponse) SetIsGridDefault(v bool)`

SetIsGridDefault sets IsGridDefault field to given value.

### HasIsGridDefault

`func (o *GetBulkhostnametemplateResponse) HasIsGridDefault() bool`

HasIsGridDefault returns a boolean if a field has been set.

### GetPreDefined

`func (o *GetBulkhostnametemplateResponse) GetPreDefined() bool`

GetPreDefined returns the PreDefined field if non-nil, zero value otherwise.

### GetPreDefinedOk

`func (o *GetBulkhostnametemplateResponse) GetPreDefinedOk() (*bool, bool)`

GetPreDefinedOk returns a tuple with the PreDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDefined

`func (o *GetBulkhostnametemplateResponse) SetPreDefined(v bool)`

SetPreDefined sets PreDefined field to given value.

### HasPreDefined

`func (o *GetBulkhostnametemplateResponse) HasPreDefined() bool`

HasPreDefined returns a boolean if a field has been set.

### GetTemplateFormat

`func (o *GetBulkhostnametemplateResponse) GetTemplateFormat() string`

GetTemplateFormat returns the TemplateFormat field if non-nil, zero value otherwise.

### GetTemplateFormatOk

`func (o *GetBulkhostnametemplateResponse) GetTemplateFormatOk() (*string, bool)`

GetTemplateFormatOk returns a tuple with the TemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFormat

`func (o *GetBulkhostnametemplateResponse) SetTemplateFormat(v string)`

SetTemplateFormat sets TemplateFormat field to given value.

### HasTemplateFormat

`func (o *GetBulkhostnametemplateResponse) HasTemplateFormat() bool`

HasTemplateFormat returns a boolean if a field has been set.

### GetTemplateName

`func (o *GetBulkhostnametemplateResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *GetBulkhostnametemplateResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *GetBulkhostnametemplateResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *GetBulkhostnametemplateResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetUuid

`func (o *GetBulkhostnametemplateResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetBulkhostnametemplateResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetBulkhostnametemplateResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetBulkhostnametemplateResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetBulkhostnametemplateResponse) GetResult() Bulkhostnametemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetBulkhostnametemplateResponse) GetResultOk() (*Bulkhostnametemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetBulkhostnametemplateResponse) SetResult(v Bulkhostnametemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetBulkhostnametemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


