# NsgroupForwardingmember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Forwarding Member Name Server Group; maximum 256 characters. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForwardingServers** | Pointer to [**[]NsgroupForwardingmemberForwardingServers**](NsgroupForwardingmemberForwardingServers.md) | The list of forwarding member servers. | [optional] 
**Name** | Pointer to **string** | The name of the Forwarding Member Name Server Group. | [optional] 

## Methods

### NewNsgroupForwardingmember

`func NewNsgroupForwardingmember() *NsgroupForwardingmember`

NewNsgroupForwardingmember instantiates a new NsgroupForwardingmember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupForwardingmemberWithDefaults

`func NewNsgroupForwardingmemberWithDefaults() *NsgroupForwardingmember`

NewNsgroupForwardingmemberWithDefaults instantiates a new NsgroupForwardingmember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *NsgroupForwardingmember) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *NsgroupForwardingmember) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *NsgroupForwardingmember) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *NsgroupForwardingmember) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *NsgroupForwardingmember) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NsgroupForwardingmember) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NsgroupForwardingmember) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NsgroupForwardingmember) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *NsgroupForwardingmember) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *NsgroupForwardingmember) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *NsgroupForwardingmember) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *NsgroupForwardingmember) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *NsgroupForwardingmember) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *NsgroupForwardingmember) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *NsgroupForwardingmember) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *NsgroupForwardingmember) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *NsgroupForwardingmember) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *NsgroupForwardingmember) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *NsgroupForwardingmember) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *NsgroupForwardingmember) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForwardingServers

`func (o *NsgroupForwardingmember) GetForwardingServers() []NsgroupForwardingmemberForwardingServers`

GetForwardingServers returns the ForwardingServers field if non-nil, zero value otherwise.

### GetForwardingServersOk

`func (o *NsgroupForwardingmember) GetForwardingServersOk() (*[]NsgroupForwardingmemberForwardingServers, bool)`

GetForwardingServersOk returns a tuple with the ForwardingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingServers

`func (o *NsgroupForwardingmember) SetForwardingServers(v []NsgroupForwardingmemberForwardingServers)`

SetForwardingServers sets ForwardingServers field to given value.

### HasForwardingServers

`func (o *NsgroupForwardingmember) HasForwardingServers() bool`

HasForwardingServers returns a boolean if a field has been set.

### GetName

`func (o *NsgroupForwardingmember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NsgroupForwardingmember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NsgroupForwardingmember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NsgroupForwardingmember) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


