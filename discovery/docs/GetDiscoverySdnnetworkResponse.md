# GetDiscoverySdnnetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**FirstSeen** | Pointer to **int64** | Timestamp when this SDN network was first discovered. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the SDN network. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view assigned to this SDN network. | [optional] [readonly] 
**SourceSdnConfig** | Pointer to **string** | Name of SDN configuration this network belongs to. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DiscoverySdnnetwork**](DiscoverySdnnetwork.md) |  | [optional] 

## Methods

### NewGetDiscoverySdnnetworkResponse

`func NewGetDiscoverySdnnetworkResponse() *GetDiscoverySdnnetworkResponse`

NewGetDiscoverySdnnetworkResponse instantiates a new GetDiscoverySdnnetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoverySdnnetworkResponseWithDefaults

`func NewGetDiscoverySdnnetworkResponseWithDefaults() *GetDiscoverySdnnetworkResponse`

NewGetDiscoverySdnnetworkResponseWithDefaults instantiates a new GetDiscoverySdnnetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoverySdnnetworkResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoverySdnnetworkResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoverySdnnetworkResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoverySdnnetworkResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetFirstSeen

`func (o *GetDiscoverySdnnetworkResponse) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetDiscoverySdnnetworkResponse) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetDiscoverySdnnetworkResponse) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GetDiscoverySdnnetworkResponse) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetName

`func (o *GetDiscoverySdnnetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoverySdnnetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoverySdnnetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoverySdnnetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetDiscoverySdnnetworkResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetDiscoverySdnnetworkResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetDiscoverySdnnetworkResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetDiscoverySdnnetworkResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetSourceSdnConfig

`func (o *GetDiscoverySdnnetworkResponse) GetSourceSdnConfig() string`

GetSourceSdnConfig returns the SourceSdnConfig field if non-nil, zero value otherwise.

### GetSourceSdnConfigOk

`func (o *GetDiscoverySdnnetworkResponse) GetSourceSdnConfigOk() (*string, bool)`

GetSourceSdnConfigOk returns a tuple with the SourceSdnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSdnConfig

`func (o *GetDiscoverySdnnetworkResponse) SetSourceSdnConfig(v string)`

SetSourceSdnConfig sets SourceSdnConfig field to given value.

### HasSourceSdnConfig

`func (o *GetDiscoverySdnnetworkResponse) HasSourceSdnConfig() bool`

HasSourceSdnConfig returns a boolean if a field has been set.

### GetUuid

`func (o *GetDiscoverySdnnetworkResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDiscoverySdnnetworkResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDiscoverySdnnetworkResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDiscoverySdnnetworkResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoverySdnnetworkResponse) GetResult() DiscoverySdnnetwork`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoverySdnnetworkResponse) GetResultOk() (*DiscoverySdnnetwork, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoverySdnnetworkResponse) SetResult(v DiscoverySdnnetwork)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoverySdnnetworkResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


