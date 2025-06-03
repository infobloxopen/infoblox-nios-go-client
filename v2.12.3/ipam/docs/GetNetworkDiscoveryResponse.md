# GetNetworkDiscoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ClearDiscoveryData** | Pointer to **map[string]interface{}** |  | [optional] 
**Result** | Pointer to [**NetworkDiscovery**](NetworkDiscovery.md) |  | [optional] 

## Methods

### NewGetNetworkDiscoveryResponse

`func NewGetNetworkDiscoveryResponse() *GetNetworkDiscoveryResponse`

NewGetNetworkDiscoveryResponse instantiates a new GetNetworkDiscoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDiscoveryResponseWithDefaults

`func NewGetNetworkDiscoveryResponseWithDefaults() *GetNetworkDiscoveryResponse`

NewGetNetworkDiscoveryResponseWithDefaults instantiates a new GetNetworkDiscoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNetworkDiscoveryResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNetworkDiscoveryResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNetworkDiscoveryResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNetworkDiscoveryResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClearDiscoveryData

`func (o *GetNetworkDiscoveryResponse) GetClearDiscoveryData() map[string]interface{}`

GetClearDiscoveryData returns the ClearDiscoveryData field if non-nil, zero value otherwise.

### GetClearDiscoveryDataOk

`func (o *GetNetworkDiscoveryResponse) GetClearDiscoveryDataOk() (*map[string]interface{}, bool)`

GetClearDiscoveryDataOk returns a tuple with the ClearDiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearDiscoveryData

`func (o *GetNetworkDiscoveryResponse) SetClearDiscoveryData(v map[string]interface{})`

SetClearDiscoveryData sets ClearDiscoveryData field to given value.

### HasClearDiscoveryData

`func (o *GetNetworkDiscoveryResponse) HasClearDiscoveryData() bool`

HasClearDiscoveryData returns a boolean if a field has been set.

### GetResult

`func (o *GetNetworkDiscoveryResponse) GetResult() NetworkDiscovery`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNetworkDiscoveryResponse) GetResultOk() (*NetworkDiscovery, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNetworkDiscoveryResponse) SetResult(v NetworkDiscovery)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNetworkDiscoveryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


