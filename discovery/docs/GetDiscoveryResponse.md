# GetDiscoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Result** | Pointer to [**Discovery**](Discovery.md) |  | [optional] 

## Methods

### NewGetDiscoveryResponse

`func NewGetDiscoveryResponse() *GetDiscoveryResponse`

NewGetDiscoveryResponse instantiates a new GetDiscoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryResponseWithDefaults

`func NewGetDiscoveryResponseWithDefaults() *GetDiscoveryResponse`

NewGetDiscoveryResponseWithDefaults instantiates a new GetDiscoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryResponse) GetResult() Discovery`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryResponse) GetResultOk() (*Discovery, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryResponse) SetResult(v Discovery)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


