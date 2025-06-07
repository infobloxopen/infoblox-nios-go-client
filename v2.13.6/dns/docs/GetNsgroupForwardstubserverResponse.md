# GetNsgroupForwardstubserverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Forward Stub Server Name Server Group; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalServers** | Pointer to [**[]NsgroupForwardstubserverExternalServers**](NsgroupForwardstubserverExternalServers.md) | The list of external servers. | [optional] 
**Name** | Pointer to **string** | The name of this Forward Stub Server Name Server Group. | [optional] 
**Result** | Pointer to [**NsgroupForwardstubserver**](NsgroupForwardstubserver.md) |  | [optional] 

## Methods

### NewGetNsgroupForwardstubserverResponse

`func NewGetNsgroupForwardstubserverResponse() *GetNsgroupForwardstubserverResponse`

NewGetNsgroupForwardstubserverResponse instantiates a new GetNsgroupForwardstubserverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNsgroupForwardstubserverResponseWithDefaults

`func NewGetNsgroupForwardstubserverResponseWithDefaults() *GetNsgroupForwardstubserverResponse`

NewGetNsgroupForwardstubserverResponseWithDefaults instantiates a new GetNsgroupForwardstubserverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNsgroupForwardstubserverResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNsgroupForwardstubserverResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNsgroupForwardstubserverResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNsgroupForwardstubserverResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetNsgroupForwardstubserverResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNsgroupForwardstubserverResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNsgroupForwardstubserverResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNsgroupForwardstubserverResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetNsgroupForwardstubserverResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetNsgroupForwardstubserverResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetNsgroupForwardstubserverResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetNsgroupForwardstubserverResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetExternalServers

`func (o *GetNsgroupForwardstubserverResponse) GetExternalServers() []NsgroupForwardstubserverExternalServers`

GetExternalServers returns the ExternalServers field if non-nil, zero value otherwise.

### GetExternalServersOk

`func (o *GetNsgroupForwardstubserverResponse) GetExternalServersOk() (*[]NsgroupForwardstubserverExternalServers, bool)`

GetExternalServersOk returns a tuple with the ExternalServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServers

`func (o *GetNsgroupForwardstubserverResponse) SetExternalServers(v []NsgroupForwardstubserverExternalServers)`

SetExternalServers sets ExternalServers field to given value.

### HasExternalServers

`func (o *GetNsgroupForwardstubserverResponse) HasExternalServers() bool`

HasExternalServers returns a boolean if a field has been set.

### GetName

`func (o *GetNsgroupForwardstubserverResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNsgroupForwardstubserverResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNsgroupForwardstubserverResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNsgroupForwardstubserverResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetNsgroupForwardstubserverResponse) GetResult() NsgroupForwardstubserver`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNsgroupForwardstubserverResponse) GetResultOk() (*NsgroupForwardstubserver, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNsgroupForwardstubserverResponse) SetResult(v NsgroupForwardstubserver)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNsgroupForwardstubserverResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


