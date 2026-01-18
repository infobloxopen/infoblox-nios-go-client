# GetDtcResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Result** | Pointer to [**Dtc**](Dtc.md) |  | [optional] 

## Methods

### NewGetDtcResponse

`func NewGetDtcResponse() *GetDtcResponse`

NewGetDtcResponse instantiates a new GetDtcResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcResponseWithDefaults

`func NewGetDtcResponseWithDefaults() *GetDtcResponse`

NewGetDtcResponseWithDefaults instantiates a new GetDtcResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcResponse) GetResult() Dtc`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcResponse) GetResultOk() (*Dtc, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcResponse) SetResult(v Dtc)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


