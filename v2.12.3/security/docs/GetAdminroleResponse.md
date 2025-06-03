# GetAdminroleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of the Admin Role object. | [optional] 
**Disable** | Pointer to **bool** | The disable flag. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of an admin role. | [optional] 
**Result** | Pointer to [**Adminrole**](Adminrole.md) |  | [optional] 

## Methods

### NewGetAdminroleResponse

`func NewGetAdminroleResponse() *GetAdminroleResponse`

NewGetAdminroleResponse instantiates a new GetAdminroleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdminroleResponseWithDefaults

`func NewGetAdminroleResponseWithDefaults() *GetAdminroleResponse`

NewGetAdminroleResponseWithDefaults instantiates a new GetAdminroleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAdminroleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAdminroleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAdminroleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAdminroleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetAdminroleResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAdminroleResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAdminroleResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAdminroleResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetAdminroleResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetAdminroleResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetAdminroleResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetAdminroleResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetAdminroleResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetAdminroleResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetAdminroleResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetAdminroleResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *GetAdminroleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdminroleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdminroleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdminroleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetAdminroleResponse) GetResult() Adminrole`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAdminroleResponse) GetResultOk() (*Adminrole, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAdminroleResponse) SetResult(v Adminrole)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAdminroleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


