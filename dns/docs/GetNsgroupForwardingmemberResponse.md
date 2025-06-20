# GetNsgroupForwardingmemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Forwarding Member Name Server Group; maximum 256 characters. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForwardingServers** | Pointer to [**[]NsgroupForwardingmemberForwardingServers**](NsgroupForwardingmemberForwardingServers.md) | The list of forwarding member servers. | [optional] 
**Name** | Pointer to **string** | The name of the Forwarding Member Name Server Group. | [optional] 
**Result** | Pointer to [**NsgroupForwardingmember**](NsgroupForwardingmember.md) |  | [optional] 

## Methods

### NewGetNsgroupForwardingmemberResponse

`func NewGetNsgroupForwardingmemberResponse() *GetNsgroupForwardingmemberResponse`

NewGetNsgroupForwardingmemberResponse instantiates a new GetNsgroupForwardingmemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNsgroupForwardingmemberResponseWithDefaults

`func NewGetNsgroupForwardingmemberResponseWithDefaults() *GetNsgroupForwardingmemberResponse`

NewGetNsgroupForwardingmemberResponseWithDefaults instantiates a new GetNsgroupForwardingmemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNsgroupForwardingmemberResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNsgroupForwardingmemberResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNsgroupForwardingmemberResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNsgroupForwardingmemberResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetNsgroupForwardingmemberResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNsgroupForwardingmemberResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNsgroupForwardingmemberResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNsgroupForwardingmemberResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetNsgroupForwardingmemberResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetNsgroupForwardingmemberResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetNsgroupForwardingmemberResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetNsgroupForwardingmemberResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForwardingServers

`func (o *GetNsgroupForwardingmemberResponse) GetForwardingServers() []NsgroupForwardingmemberForwardingServers`

GetForwardingServers returns the ForwardingServers field if non-nil, zero value otherwise.

### GetForwardingServersOk

`func (o *GetNsgroupForwardingmemberResponse) GetForwardingServersOk() (*[]NsgroupForwardingmemberForwardingServers, bool)`

GetForwardingServersOk returns a tuple with the ForwardingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingServers

`func (o *GetNsgroupForwardingmemberResponse) SetForwardingServers(v []NsgroupForwardingmemberForwardingServers)`

SetForwardingServers sets ForwardingServers field to given value.

### HasForwardingServers

`func (o *GetNsgroupForwardingmemberResponse) HasForwardingServers() bool`

HasForwardingServers returns a boolean if a field has been set.

### GetName

`func (o *GetNsgroupForwardingmemberResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNsgroupForwardingmemberResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNsgroupForwardingmemberResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNsgroupForwardingmemberResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetNsgroupForwardingmemberResponse) GetResult() NsgroupForwardingmember`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNsgroupForwardingmemberResponse) GetResultOk() (*NsgroupForwardingmember, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNsgroupForwardingmemberResponse) SetResult(v NsgroupForwardingmember)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNsgroupForwardingmemberResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


