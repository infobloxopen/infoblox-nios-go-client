# GetNsgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the name server group; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalPrimaries** | Pointer to [**[]NsgroupExternalPrimaries**](NsgroupExternalPrimaries.md) | The list of external primary servers. | [optional] 
**ExternalSecondaries** | Pointer to [**[]NsgroupExternalSecondaries**](NsgroupExternalSecondaries.md) | The list of external secondary servers. | [optional] 
**GridPrimary** | Pointer to [**[]NsgroupGridPrimary**](NsgroupGridPrimary.md) | The grid primary servers for this group. | [optional] 
**GridSecondaries** | Pointer to [**[]NsgroupGridSecondaries**](NsgroupGridSecondaries.md) | The list with Grid members that are secondary servers for this group. | [optional] 
**IsGridDefault** | Pointer to **bool** | Determines if this name server group is the Grid default. | [optional] 
**IsMultimaster** | Pointer to **bool** | Determines if the \&quot;multiple DNS primaries\&quot; feature is enabled for the group. | [optional] 
**Name** | Pointer to **string** | The name of this name server group. | [optional] 
**UseExternalPrimary** | Pointer to **bool** | This flag controls whether the group is using an external primary. Note that modification of this field requires passing values for \&quot;grid_secondaries\&quot; and \&quot;external_primaries\&quot;. | [optional] 
**Result** | Pointer to [**Nsgroup**](Nsgroup.md) |  | [optional] 

## Methods

### NewGetNsgroupResponse

`func NewGetNsgroupResponse() *GetNsgroupResponse`

NewGetNsgroupResponse instantiates a new GetNsgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNsgroupResponseWithDefaults

`func NewGetNsgroupResponseWithDefaults() *GetNsgroupResponse`

NewGetNsgroupResponseWithDefaults instantiates a new GetNsgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNsgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNsgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNsgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNsgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetNsgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNsgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNsgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNsgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetNsgroupResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetNsgroupResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetNsgroupResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetNsgroupResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *GetNsgroupResponse) GetExternalPrimaries() []NsgroupExternalPrimaries`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *GetNsgroupResponse) GetExternalPrimariesOk() (*[]NsgroupExternalPrimaries, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *GetNsgroupResponse) SetExternalPrimaries(v []NsgroupExternalPrimaries)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *GetNsgroupResponse) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *GetNsgroupResponse) GetExternalSecondaries() []NsgroupExternalSecondaries`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *GetNsgroupResponse) GetExternalSecondariesOk() (*[]NsgroupExternalSecondaries, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *GetNsgroupResponse) SetExternalSecondaries(v []NsgroupExternalSecondaries)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *GetNsgroupResponse) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetGridPrimary

`func (o *GetNsgroupResponse) GetGridPrimary() []NsgroupGridPrimary`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *GetNsgroupResponse) GetGridPrimaryOk() (*[]NsgroupGridPrimary, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *GetNsgroupResponse) SetGridPrimary(v []NsgroupGridPrimary)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *GetNsgroupResponse) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *GetNsgroupResponse) GetGridSecondaries() []NsgroupGridSecondaries`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *GetNsgroupResponse) GetGridSecondariesOk() (*[]NsgroupGridSecondaries, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *GetNsgroupResponse) SetGridSecondaries(v []NsgroupGridSecondaries)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *GetNsgroupResponse) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetIsGridDefault

`func (o *GetNsgroupResponse) GetIsGridDefault() bool`

GetIsGridDefault returns the IsGridDefault field if non-nil, zero value otherwise.

### GetIsGridDefaultOk

`func (o *GetNsgroupResponse) GetIsGridDefaultOk() (*bool, bool)`

GetIsGridDefaultOk returns a tuple with the IsGridDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGridDefault

`func (o *GetNsgroupResponse) SetIsGridDefault(v bool)`

SetIsGridDefault sets IsGridDefault field to given value.

### HasIsGridDefault

`func (o *GetNsgroupResponse) HasIsGridDefault() bool`

HasIsGridDefault returns a boolean if a field has been set.

### GetIsMultimaster

`func (o *GetNsgroupResponse) GetIsMultimaster() bool`

GetIsMultimaster returns the IsMultimaster field if non-nil, zero value otherwise.

### GetIsMultimasterOk

`func (o *GetNsgroupResponse) GetIsMultimasterOk() (*bool, bool)`

GetIsMultimasterOk returns a tuple with the IsMultimaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultimaster

`func (o *GetNsgroupResponse) SetIsMultimaster(v bool)`

SetIsMultimaster sets IsMultimaster field to given value.

### HasIsMultimaster

`func (o *GetNsgroupResponse) HasIsMultimaster() bool`

HasIsMultimaster returns a boolean if a field has been set.

### GetName

`func (o *GetNsgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNsgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNsgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNsgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUseExternalPrimary

`func (o *GetNsgroupResponse) GetUseExternalPrimary() bool`

GetUseExternalPrimary returns the UseExternalPrimary field if non-nil, zero value otherwise.

### GetUseExternalPrimaryOk

`func (o *GetNsgroupResponse) GetUseExternalPrimaryOk() (*bool, bool)`

GetUseExternalPrimaryOk returns a tuple with the UseExternalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalPrimary

`func (o *GetNsgroupResponse) SetUseExternalPrimary(v bool)`

SetUseExternalPrimary sets UseExternalPrimary field to given value.

### HasUseExternalPrimary

`func (o *GetNsgroupResponse) HasUseExternalPrimary() bool`

HasUseExternalPrimary returns a boolean if a field has been set.

### GetResult

`func (o *GetNsgroupResponse) GetResult() Nsgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNsgroupResponse) GetResultOk() (*Nsgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNsgroupResponse) SetResult(v Nsgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNsgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


