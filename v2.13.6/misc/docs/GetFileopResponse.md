# GetFileopResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Result** | Pointer to [**Fileop**](Fileop.md) |  | [optional] 

## Methods

### NewGetFileopResponse

`func NewGetFileopResponse() *GetFileopResponse`

NewGetFileopResponse instantiates a new GetFileopResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFileopResponseWithDefaults

`func NewGetFileopResponseWithDefaults() *GetFileopResponse`

NewGetFileopResponseWithDefaults instantiates a new GetFileopResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFileopResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFileopResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFileopResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFileopResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetResult

`func (o *GetFileopResponse) GetResult() Fileop`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFileopResponse) GetResultOk() (*Fileop, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFileopResponse) SetResult(v Fileop)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFileopResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


