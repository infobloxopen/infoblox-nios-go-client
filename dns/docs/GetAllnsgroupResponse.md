# GetAllnsgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for the name server group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the name server group. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the name server group. | [optional] [readonly] 
**Result** | Pointer to [**Allnsgroup**](Allnsgroup.md) |  | [optional] 

## Methods

### NewGetAllnsgroupResponse

`func NewGetAllnsgroupResponse() *GetAllnsgroupResponse`

NewGetAllnsgroupResponse instantiates a new GetAllnsgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllnsgroupResponseWithDefaults

`func NewGetAllnsgroupResponseWithDefaults() *GetAllnsgroupResponse`

NewGetAllnsgroupResponseWithDefaults instantiates a new GetAllnsgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAllnsgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAllnsgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAllnsgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAllnsgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetAllnsgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAllnsgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAllnsgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAllnsgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *GetAllnsgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllnsgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllnsgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllnsgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetAllnsgroupResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAllnsgroupResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAllnsgroupResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAllnsgroupResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetAllnsgroupResponse) GetResult() Allnsgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllnsgroupResponse) GetResultOk() (*Allnsgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllnsgroupResponse) SetResult(v Allnsgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAllnsgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


