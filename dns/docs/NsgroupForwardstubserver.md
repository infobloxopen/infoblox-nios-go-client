# NsgroupForwardstubserver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the Forward Stub Server Name Server Group; maximum 256 characters. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalServers** | Pointer to [**[]NsgroupForwardstubserverExternalServers**](NsgroupForwardstubserverExternalServers.md) | The list of external servers. | [optional] 
**Name** | Pointer to **string** | The name of this Forward Stub Server Name Server Group. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewNsgroupForwardstubserver

`func NewNsgroupForwardstubserver() *NsgroupForwardstubserver`

NewNsgroupForwardstubserver instantiates a new NsgroupForwardstubserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupForwardstubserverWithDefaults

`func NewNsgroupForwardstubserverWithDefaults() *NsgroupForwardstubserver`

NewNsgroupForwardstubserverWithDefaults instantiates a new NsgroupForwardstubserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *NsgroupForwardstubserver) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *NsgroupForwardstubserver) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *NsgroupForwardstubserver) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *NsgroupForwardstubserver) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *NsgroupForwardstubserver) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NsgroupForwardstubserver) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NsgroupForwardstubserver) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NsgroupForwardstubserver) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *NsgroupForwardstubserver) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *NsgroupForwardstubserver) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *NsgroupForwardstubserver) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *NsgroupForwardstubserver) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *NsgroupForwardstubserver) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *NsgroupForwardstubserver) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *NsgroupForwardstubserver) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *NsgroupForwardstubserver) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *NsgroupForwardstubserver) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *NsgroupForwardstubserver) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *NsgroupForwardstubserver) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *NsgroupForwardstubserver) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetExternalServers

`func (o *NsgroupForwardstubserver) GetExternalServers() []NsgroupForwardstubserverExternalServers`

GetExternalServers returns the ExternalServers field if non-nil, zero value otherwise.

### GetExternalServersOk

`func (o *NsgroupForwardstubserver) GetExternalServersOk() (*[]NsgroupForwardstubserverExternalServers, bool)`

GetExternalServersOk returns a tuple with the ExternalServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServers

`func (o *NsgroupForwardstubserver) SetExternalServers(v []NsgroupForwardstubserverExternalServers)`

SetExternalServers sets ExternalServers field to given value.

### HasExternalServers

`func (o *NsgroupForwardstubserver) HasExternalServers() bool`

HasExternalServers returns a boolean if a field has been set.

### GetName

`func (o *NsgroupForwardstubserver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NsgroupForwardstubserver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NsgroupForwardstubserver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NsgroupForwardstubserver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *NsgroupForwardstubserver) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NsgroupForwardstubserver) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NsgroupForwardstubserver) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NsgroupForwardstubserver) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


