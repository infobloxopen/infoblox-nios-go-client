# GetSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Result** | Pointer to [**Search**](Search.md) |  | [optional] 

## Methods

### NewGetSearchResponse

`func NewGetSearchResponse() *GetSearchResponse`

NewGetSearchResponse instantiates a new GetSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSearchResponseWithDefaults

`func NewGetSearchResponseWithDefaults() *GetSearchResponse`

NewGetSearchResponseWithDefaults instantiates a new GetSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSearchResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSearchResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSearchResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSearchResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetResult

`func (o *GetSearchResponse) GetResult() Search`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSearchResponse) GetResultOk() (*Search, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSearchResponse) SetResult(v Search)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSearchResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


