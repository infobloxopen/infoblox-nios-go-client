# GetNatgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The NAT group descriptive comment. | [optional] 
**Name** | Pointer to **string** | The name of a NAT group object. | [optional] 
**Result** | Pointer to [**Natgroup**](Natgroup.md) |  | [optional] 

## Methods

### NewGetNatgroupResponse

`func NewGetNatgroupResponse() *GetNatgroupResponse`

NewGetNatgroupResponse instantiates a new GetNatgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNatgroupResponseWithDefaults

`func NewGetNatgroupResponseWithDefaults() *GetNatgroupResponse`

NewGetNatgroupResponseWithDefaults instantiates a new GetNatgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNatgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNatgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNatgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNatgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetNatgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNatgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNatgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNatgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *GetNatgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNatgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNatgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNatgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetNatgroupResponse) GetResult() Natgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNatgroupResponse) GetResultOk() (*Natgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNatgroupResponse) SetResult(v Natgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNatgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


