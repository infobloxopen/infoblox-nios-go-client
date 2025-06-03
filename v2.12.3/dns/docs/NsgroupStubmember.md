# NsgroupStubmember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Stub Member Name Server Group; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the Stub Member Name Server Group. | [optional] 
**StubMembers** | Pointer to [**[]NsgroupStubmemberStubMembers**](NsgroupStubmemberStubMembers.md) | The Grid member servers of this stub zone. Note that the lead/stealth/grid_replicate/ preferred_primaries/override_preferred_primaries fields of the struct will be ignored when set in this field. | [optional] 

## Methods

### NewNsgroupStubmember

`func NewNsgroupStubmember() *NsgroupStubmember`

NewNsgroupStubmember instantiates a new NsgroupStubmember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupStubmemberWithDefaults

`func NewNsgroupStubmemberWithDefaults() *NsgroupStubmember`

NewNsgroupStubmemberWithDefaults instantiates a new NsgroupStubmember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *NsgroupStubmember) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *NsgroupStubmember) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *NsgroupStubmember) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *NsgroupStubmember) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *NsgroupStubmember) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NsgroupStubmember) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NsgroupStubmember) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NsgroupStubmember) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *NsgroupStubmember) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *NsgroupStubmember) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *NsgroupStubmember) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *NsgroupStubmember) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *NsgroupStubmember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NsgroupStubmember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NsgroupStubmember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NsgroupStubmember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStubMembers

`func (o *NsgroupStubmember) GetStubMembers() []NsgroupStubmemberStubMembers`

GetStubMembers returns the StubMembers field if non-nil, zero value otherwise.

### GetStubMembersOk

`func (o *NsgroupStubmember) GetStubMembersOk() (*[]NsgroupStubmemberStubMembers, bool)`

GetStubMembersOk returns a tuple with the StubMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubMembers

`func (o *NsgroupStubmember) SetStubMembers(v []NsgroupStubmemberStubMembers)`

SetStubMembers sets StubMembers field to given value.

### HasStubMembers

`func (o *NsgroupStubmember) HasStubMembers() bool`

HasStubMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


