# ListDiscoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]Discovery**](Discovery.md) |  | [optional] 

## Methods

### NewListDiscoveryResponse

`func NewListDiscoveryResponse() *ListDiscoveryResponse`

NewListDiscoveryResponse instantiates a new ListDiscoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscoveryResponseWithDefaults

`func NewListDiscoveryResponseWithDefaults() *ListDiscoveryResponse`

NewListDiscoveryResponseWithDefaults instantiates a new ListDiscoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDiscoveryResponse) GetResult() []Discovery`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDiscoveryResponse) GetResultOk() (*[]Discovery, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDiscoveryResponse) SetResult(v []Discovery)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDiscoveryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


