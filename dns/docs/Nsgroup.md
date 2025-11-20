# Nsgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the name server group; maximum 256 characters. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalPrimaries** | Pointer to [**[]NsgroupExternalPrimaries**](NsgroupExternalPrimaries.md) | The list of external primary servers. | [optional] 
**ExternalSecondaries** | Pointer to [**[]NsgroupExternalSecondaries**](NsgroupExternalSecondaries.md) | The list of external secondary servers. | [optional] 
**GridPrimary** | Pointer to [**[]NsgroupGridPrimary**](NsgroupGridPrimary.md) | The grid primary servers for this group. | [optional] 
**GridSecondaries** | Pointer to [**[]NsgroupGridSecondaries**](NsgroupGridSecondaries.md) | The list with Grid members that are secondary servers for this group. | [optional] 
**IsGridDefault** | Pointer to **bool** | Determines if this name server group is the Grid default. | [optional] 
**IsMultimaster** | Pointer to **bool** | Determines if the \&quot;multiple DNS primaries\&quot; feature is enabled for the group. | [optional] 
**Name** | Pointer to **string** | The name of this name server group. | [optional] 
**UseExternalPrimary** | Pointer to **bool** | This flag controls whether the group is using an external primary. Note that modification of this field requires passing values for \&quot;grid_secondaries\&quot; and \&quot;external_primaries\&quot;. | [optional] 

## Methods

### NewNsgroup

`func NewNsgroup() *Nsgroup`

NewNsgroup instantiates a new Nsgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupWithDefaults

`func NewNsgroupWithDefaults() *Nsgroup`

NewNsgroupWithDefaults instantiates a new Nsgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Nsgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Nsgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Nsgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Nsgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Nsgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Nsgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Nsgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Nsgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Nsgroup) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Nsgroup) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Nsgroup) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Nsgroup) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Nsgroup) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Nsgroup) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Nsgroup) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Nsgroup) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Nsgroup) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Nsgroup) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Nsgroup) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Nsgroup) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *Nsgroup) GetExternalPrimaries() []NsgroupExternalPrimaries`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *Nsgroup) GetExternalPrimariesOk() (*[]NsgroupExternalPrimaries, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *Nsgroup) SetExternalPrimaries(v []NsgroupExternalPrimaries)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *Nsgroup) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *Nsgroup) GetExternalSecondaries() []NsgroupExternalSecondaries`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *Nsgroup) GetExternalSecondariesOk() (*[]NsgroupExternalSecondaries, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *Nsgroup) SetExternalSecondaries(v []NsgroupExternalSecondaries)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *Nsgroup) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetGridPrimary

`func (o *Nsgroup) GetGridPrimary() []NsgroupGridPrimary`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *Nsgroup) GetGridPrimaryOk() (*[]NsgroupGridPrimary, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *Nsgroup) SetGridPrimary(v []NsgroupGridPrimary)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *Nsgroup) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *Nsgroup) GetGridSecondaries() []NsgroupGridSecondaries`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *Nsgroup) GetGridSecondariesOk() (*[]NsgroupGridSecondaries, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *Nsgroup) SetGridSecondaries(v []NsgroupGridSecondaries)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *Nsgroup) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetIsGridDefault

`func (o *Nsgroup) GetIsGridDefault() bool`

GetIsGridDefault returns the IsGridDefault field if non-nil, zero value otherwise.

### GetIsGridDefaultOk

`func (o *Nsgroup) GetIsGridDefaultOk() (*bool, bool)`

GetIsGridDefaultOk returns a tuple with the IsGridDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGridDefault

`func (o *Nsgroup) SetIsGridDefault(v bool)`

SetIsGridDefault sets IsGridDefault field to given value.

### HasIsGridDefault

`func (o *Nsgroup) HasIsGridDefault() bool`

HasIsGridDefault returns a boolean if a field has been set.

### GetIsMultimaster

`func (o *Nsgroup) GetIsMultimaster() bool`

GetIsMultimaster returns the IsMultimaster field if non-nil, zero value otherwise.

### GetIsMultimasterOk

`func (o *Nsgroup) GetIsMultimasterOk() (*bool, bool)`

GetIsMultimasterOk returns a tuple with the IsMultimaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultimaster

`func (o *Nsgroup) SetIsMultimaster(v bool)`

SetIsMultimaster sets IsMultimaster field to given value.

### HasIsMultimaster

`func (o *Nsgroup) HasIsMultimaster() bool`

HasIsMultimaster returns a boolean if a field has been set.

### GetName

`func (o *Nsgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Nsgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Nsgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Nsgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUseExternalPrimary

`func (o *Nsgroup) GetUseExternalPrimary() bool`

GetUseExternalPrimary returns the UseExternalPrimary field if non-nil, zero value otherwise.

### GetUseExternalPrimaryOk

`func (o *Nsgroup) GetUseExternalPrimaryOk() (*bool, bool)`

GetUseExternalPrimaryOk returns a tuple with the UseExternalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalPrimary

`func (o *Nsgroup) SetUseExternalPrimary(v bool)`

SetUseExternalPrimary sets UseExternalPrimary field to given value.

### HasUseExternalPrimary

`func (o *Nsgroup) HasUseExternalPrimary() bool`

HasUseExternalPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


