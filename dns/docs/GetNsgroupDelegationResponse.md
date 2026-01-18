# GetNsgroupDelegationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the delegated NS group. | [optional] 
**DelegateTo** | Pointer to [**[]NsgroupDelegationDelegateTo**](NsgroupDelegationDelegateTo.md) | The list of delegated servers for the delegated NS group. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the delegated NS group. | [optional] 
**Result** | Pointer to [**NsgroupDelegation**](NsgroupDelegation.md) |  | [optional] 

## Methods

### NewGetNsgroupDelegationResponse

`func NewGetNsgroupDelegationResponse() *GetNsgroupDelegationResponse`

NewGetNsgroupDelegationResponse instantiates a new GetNsgroupDelegationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNsgroupDelegationResponseWithDefaults

`func NewGetNsgroupDelegationResponseWithDefaults() *GetNsgroupDelegationResponse`

NewGetNsgroupDelegationResponseWithDefaults instantiates a new GetNsgroupDelegationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNsgroupDelegationResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNsgroupDelegationResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNsgroupDelegationResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNsgroupDelegationResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetNsgroupDelegationResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNsgroupDelegationResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNsgroupDelegationResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNsgroupDelegationResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegateTo

`func (o *GetNsgroupDelegationResponse) GetDelegateTo() []NsgroupDelegationDelegateTo`

GetDelegateTo returns the DelegateTo field if non-nil, zero value otherwise.

### GetDelegateToOk

`func (o *GetNsgroupDelegationResponse) GetDelegateToOk() (*[]NsgroupDelegationDelegateTo, bool)`

GetDelegateToOk returns a tuple with the DelegateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateTo

`func (o *GetNsgroupDelegationResponse) SetDelegateTo(v []NsgroupDelegationDelegateTo)`

SetDelegateTo sets DelegateTo field to given value.

### HasDelegateTo

`func (o *GetNsgroupDelegationResponse) HasDelegateTo() bool`

HasDelegateTo returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetNsgroupDelegationResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetNsgroupDelegationResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetNsgroupDelegationResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetNsgroupDelegationResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetNsgroupDelegationResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetNsgroupDelegationResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetNsgroupDelegationResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetNsgroupDelegationResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetNsgroupDelegationResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetNsgroupDelegationResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetNsgroupDelegationResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetNsgroupDelegationResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GetNsgroupDelegationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNsgroupDelegationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNsgroupDelegationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNsgroupDelegationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetNsgroupDelegationResponse) GetResult() NsgroupDelegation`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNsgroupDelegationResponse) GetResultOk() (*NsgroupDelegation, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNsgroupDelegationResponse) SetResult(v NsgroupDelegation)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNsgroupDelegationResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


