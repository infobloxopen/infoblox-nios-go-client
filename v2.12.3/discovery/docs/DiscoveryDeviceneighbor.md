# DiscoveryDeviceneighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the device neighbor. | [optional] [readonly] 
**AddressRef** | Pointer to **string** | The ref to the management IP address of the device neighbor. | [optional] [readonly] 
**Device** | Pointer to **string** | The ref to the device to which the device neighbor belongs. | [optional] [readonly] 
**Interface** | Pointer to **string** | The ref to the interface to which the device neighbor belongs. | [optional] [readonly] 
**Mac** | Pointer to **string** | The MAC address of the device neighbor. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the device neighbor. | [optional] [readonly] 
**VlanInfos** | Pointer to [**[]DiscoveryDeviceneighborVlanInfos**](DiscoveryDeviceneighborVlanInfos.md) | The list of VLAN information associated with the device neighbor. | [optional] [readonly] 

## Methods

### NewDiscoveryDeviceneighbor

`func NewDiscoveryDeviceneighbor() *DiscoveryDeviceneighbor`

NewDiscoveryDeviceneighbor instantiates a new DiscoveryDeviceneighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDeviceneighborWithDefaults

`func NewDiscoveryDeviceneighborWithDefaults() *DiscoveryDeviceneighbor`

NewDiscoveryDeviceneighborWithDefaults instantiates a new DiscoveryDeviceneighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryDeviceneighbor) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryDeviceneighbor) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryDeviceneighbor) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryDeviceneighbor) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveryDeviceneighbor) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveryDeviceneighbor) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveryDeviceneighbor) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveryDeviceneighbor) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressRef

`func (o *DiscoveryDeviceneighbor) GetAddressRef() string`

GetAddressRef returns the AddressRef field if non-nil, zero value otherwise.

### GetAddressRefOk

`func (o *DiscoveryDeviceneighbor) GetAddressRefOk() (*string, bool)`

GetAddressRefOk returns a tuple with the AddressRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRef

`func (o *DiscoveryDeviceneighbor) SetAddressRef(v string)`

SetAddressRef sets AddressRef field to given value.

### HasAddressRef

`func (o *DiscoveryDeviceneighbor) HasAddressRef() bool`

HasAddressRef returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveryDeviceneighbor) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveryDeviceneighbor) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveryDeviceneighbor) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveryDeviceneighbor) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetInterface

`func (o *DiscoveryDeviceneighbor) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DiscoveryDeviceneighbor) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DiscoveryDeviceneighbor) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DiscoveryDeviceneighbor) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetMac

`func (o *DiscoveryDeviceneighbor) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DiscoveryDeviceneighbor) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DiscoveryDeviceneighbor) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DiscoveryDeviceneighbor) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *DiscoveryDeviceneighbor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryDeviceneighbor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryDeviceneighbor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveryDeviceneighbor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanInfos

`func (o *DiscoveryDeviceneighbor) GetVlanInfos() []DiscoveryDeviceneighborVlanInfos`

GetVlanInfos returns the VlanInfos field if non-nil, zero value otherwise.

### GetVlanInfosOk

`func (o *DiscoveryDeviceneighbor) GetVlanInfosOk() (*[]DiscoveryDeviceneighborVlanInfos, bool)`

GetVlanInfosOk returns a tuple with the VlanInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfos

`func (o *DiscoveryDeviceneighbor) SetVlanInfos(v []DiscoveryDeviceneighborVlanInfos)`

SetVlanInfos sets VlanInfos field to given value.

### HasVlanInfos

`func (o *DiscoveryDeviceneighbor) HasVlanInfos() bool`

HasVlanInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


