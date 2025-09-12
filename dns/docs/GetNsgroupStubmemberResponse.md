# GetNsgroupStubmemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Stub Member Name Server Group; maximum 256 characters. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the Stub Member Name Server Group. | [optional] 
**StubMembers** | Pointer to [**[]NsgroupStubmemberStubMembers**](NsgroupStubmemberStubMembers.md) | The Grid member servers of this stub zone. Note that the lead/stealth/grid_replicate/ preferred_primaries/override_preferred_primaries fields of the struct will be ignored when set in this field. | [optional] 
**Result** | Pointer to [**NsgroupStubmember**](NsgroupStubmember.md) |  | [optional] 

## Methods

### NewGetNsgroupStubmemberResponse

`func NewGetNsgroupStubmemberResponse() *GetNsgroupStubmemberResponse`

NewGetNsgroupStubmemberResponse instantiates a new GetNsgroupStubmemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNsgroupStubmemberResponseWithDefaults

`func NewGetNsgroupStubmemberResponseWithDefaults() *GetNsgroupStubmemberResponse`

NewGetNsgroupStubmemberResponseWithDefaults instantiates a new GetNsgroupStubmemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNsgroupStubmemberResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNsgroupStubmemberResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNsgroupStubmemberResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNsgroupStubmemberResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetNsgroupStubmemberResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNsgroupStubmemberResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNsgroupStubmemberResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNsgroupStubmemberResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetNsgroupStubmemberResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetNsgroupStubmemberResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetNsgroupStubmemberResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetNsgroupStubmemberResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetNsgroupStubmemberResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetNsgroupStubmemberResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetNsgroupStubmemberResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetNsgroupStubmemberResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetNsgroupStubmemberResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetNsgroupStubmemberResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetNsgroupStubmemberResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetNsgroupStubmemberResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GetNsgroupStubmemberResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNsgroupStubmemberResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNsgroupStubmemberResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNsgroupStubmemberResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStubMembers

`func (o *GetNsgroupStubmemberResponse) GetStubMembers() []NsgroupStubmemberStubMembers`

GetStubMembers returns the StubMembers field if non-nil, zero value otherwise.

### GetStubMembersOk

`func (o *GetNsgroupStubmemberResponse) GetStubMembersOk() (*[]NsgroupStubmemberStubMembers, bool)`

GetStubMembersOk returns a tuple with the StubMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubMembers

`func (o *GetNsgroupStubmemberResponse) SetStubMembers(v []NsgroupStubmemberStubMembers)`

SetStubMembers sets StubMembers field to given value.

### HasStubMembers

`func (o *GetNsgroupStubmemberResponse) HasStubMembers() bool`

HasStubMembers returns a boolean if a field has been set.

### GetResult

`func (o *GetNsgroupStubmemberResponse) GetResult() NsgroupStubmember`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNsgroupStubmemberResponse) GetResultOk() (*NsgroupStubmember, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNsgroupStubmemberResponse) SetResult(v NsgroupStubmember)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNsgroupStubmemberResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


