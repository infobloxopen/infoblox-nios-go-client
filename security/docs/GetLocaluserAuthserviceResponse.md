# GetLocaluserAuthserviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The local user authentication service comment. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Flag that indicates whether the local user authentication service is enabled or not. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the local user authentication service. | [optional] [readonly] 
**Result** | Pointer to [**LocaluserAuthservice**](LocaluserAuthservice.md) |  | [optional] 

## Methods

### NewGetLocaluserAuthserviceResponse

`func NewGetLocaluserAuthserviceResponse() *GetLocaluserAuthserviceResponse`

NewGetLocaluserAuthserviceResponse instantiates a new GetLocaluserAuthserviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLocaluserAuthserviceResponseWithDefaults

`func NewGetLocaluserAuthserviceResponseWithDefaults() *GetLocaluserAuthserviceResponse`

NewGetLocaluserAuthserviceResponseWithDefaults instantiates a new GetLocaluserAuthserviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetLocaluserAuthserviceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetLocaluserAuthserviceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetLocaluserAuthserviceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetLocaluserAuthserviceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetLocaluserAuthserviceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetLocaluserAuthserviceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetLocaluserAuthserviceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetLocaluserAuthserviceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *GetLocaluserAuthserviceResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetLocaluserAuthserviceResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetLocaluserAuthserviceResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetLocaluserAuthserviceResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *GetLocaluserAuthserviceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLocaluserAuthserviceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLocaluserAuthserviceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLocaluserAuthserviceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetLocaluserAuthserviceResponse) GetResult() LocaluserAuthservice`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetLocaluserAuthserviceResponse) GetResultOk() (*LocaluserAuthservice, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetLocaluserAuthserviceResponse) SetResult(v LocaluserAuthservice)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetLocaluserAuthserviceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


