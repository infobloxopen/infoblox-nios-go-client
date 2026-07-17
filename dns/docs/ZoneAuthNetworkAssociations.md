# ZoneAuthNetworkAssociations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the network or network container object. | [optional] [readonly] 
**Network** | Pointer to **string** | The network address in CIDR notation. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view. | [optional] [readonly] 

## Methods

### NewZoneAuthNetworkAssociations

`func NewZoneAuthNetworkAssociations() *ZoneAuthNetworkAssociations`

NewZoneAuthNetworkAssociations instantiates a new ZoneAuthNetworkAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthNetworkAssociationsWithDefaults

`func NewZoneAuthNetworkAssociationsWithDefaults() *ZoneAuthNetworkAssociations`

NewZoneAuthNetworkAssociationsWithDefaults instantiates a new ZoneAuthNetworkAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ZoneAuthNetworkAssociations) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ZoneAuthNetworkAssociations) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ZoneAuthNetworkAssociations) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ZoneAuthNetworkAssociations) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetNetwork

`func (o *ZoneAuthNetworkAssociations) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ZoneAuthNetworkAssociations) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ZoneAuthNetworkAssociations) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ZoneAuthNetworkAssociations) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *ZoneAuthNetworkAssociations) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *ZoneAuthNetworkAssociations) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *ZoneAuthNetworkAssociations) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *ZoneAuthNetworkAssociations) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


