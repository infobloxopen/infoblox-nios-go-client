# DiscoveryDeviceNetworkInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **string** | The ref to the network to which the management IP address belongs. | [optional] [readonly] 
**NetworkStr** | Pointer to **string** | The Network address in format address/cidr. | [optional] [readonly] 

## Methods

### NewDiscoveryDeviceNetworkInfos

`func NewDiscoveryDeviceNetworkInfos() *DiscoveryDeviceNetworkInfos`

NewDiscoveryDeviceNetworkInfos instantiates a new DiscoveryDeviceNetworkInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDeviceNetworkInfosWithDefaults

`func NewDiscoveryDeviceNetworkInfosWithDefaults() *DiscoveryDeviceNetworkInfos`

NewDiscoveryDeviceNetworkInfosWithDefaults instantiates a new DiscoveryDeviceNetworkInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *DiscoveryDeviceNetworkInfos) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DiscoveryDeviceNetworkInfos) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DiscoveryDeviceNetworkInfos) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DiscoveryDeviceNetworkInfos) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkStr

`func (o *DiscoveryDeviceNetworkInfos) GetNetworkStr() string`

GetNetworkStr returns the NetworkStr field if non-nil, zero value otherwise.

### GetNetworkStrOk

`func (o *DiscoveryDeviceNetworkInfos) GetNetworkStrOk() (*string, bool)`

GetNetworkStrOk returns a tuple with the NetworkStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStr

`func (o *DiscoveryDeviceNetworkInfos) SetNetworkStr(v string)`

SetNetworkStr sets NetworkStr field to given value.

### HasNetworkStr

`func (o *DiscoveryDeviceNetworkInfos) HasNetworkStr() bool`

HasNetworkStr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


