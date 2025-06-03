# DiscoveryDeviceinterfaceIfaddrInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the device. | [optional] [readonly] 
**AddressObject** | Pointer to **string** | The ref to IPv4/Ipv6 Address. | [optional] [readonly] 
**Network** | Pointer to **string** | The network to which this device belongs, in IPv4 Address/CIDR format. | [optional] [readonly] 

## Methods

### NewDiscoveryDeviceinterfaceIfaddrInfos

`func NewDiscoveryDeviceinterfaceIfaddrInfos() *DiscoveryDeviceinterfaceIfaddrInfos`

NewDiscoveryDeviceinterfaceIfaddrInfos instantiates a new DiscoveryDeviceinterfaceIfaddrInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDeviceinterfaceIfaddrInfosWithDefaults

`func NewDiscoveryDeviceinterfaceIfaddrInfosWithDefaults() *DiscoveryDeviceinterfaceIfaddrInfos`

NewDiscoveryDeviceinterfaceIfaddrInfosWithDefaults instantiates a new DiscoveryDeviceinterfaceIfaddrInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressObject

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) GetAddressObject() string`

GetAddressObject returns the AddressObject field if non-nil, zero value otherwise.

### GetAddressObjectOk

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) GetAddressObjectOk() (*string, bool)`

GetAddressObjectOk returns a tuple with the AddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressObject

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) SetAddressObject(v string)`

SetAddressObject sets AddressObject field to given value.

### HasAddressObject

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) HasAddressObject() bool`

HasAddressObject returns a boolean if a field has been set.

### GetNetwork

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DiscoveryDeviceinterfaceIfaddrInfos) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


