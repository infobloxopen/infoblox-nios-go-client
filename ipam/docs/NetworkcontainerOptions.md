# NetworkcontainerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the DHCP option. | [optional] 
**Num** | Pointer to **int64** | The code of the DHCP option. | [optional] 
**VendorClass** | Pointer to **string** | The name of the space this DHCP option is associated to. | [optional] 
**Value** | Pointer to **string** | Value of the DHCP option | [optional] 
**UseOption** | Pointer to **bool** | Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers | [optional] 

## Methods

### NewNetworkcontainerOptions

`func NewNetworkcontainerOptions() *NetworkcontainerOptions`

NewNetworkcontainerOptions instantiates a new NetworkcontainerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkcontainerOptionsWithDefaults

`func NewNetworkcontainerOptionsWithDefaults() *NetworkcontainerOptions`

NewNetworkcontainerOptionsWithDefaults instantiates a new NetworkcontainerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkcontainerOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkcontainerOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkcontainerOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkcontainerOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNum

`func (o *NetworkcontainerOptions) GetNum() int64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *NetworkcontainerOptions) GetNumOk() (*int64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *NetworkcontainerOptions) SetNum(v int64)`

SetNum sets Num field to given value.

### HasNum

`func (o *NetworkcontainerOptions) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetVendorClass

`func (o *NetworkcontainerOptions) GetVendorClass() string`

GetVendorClass returns the VendorClass field if non-nil, zero value otherwise.

### GetVendorClassOk

`func (o *NetworkcontainerOptions) GetVendorClassOk() (*string, bool)`

GetVendorClassOk returns a tuple with the VendorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorClass

`func (o *NetworkcontainerOptions) SetVendorClass(v string)`

SetVendorClass sets VendorClass field to given value.

### HasVendorClass

`func (o *NetworkcontainerOptions) HasVendorClass() bool`

HasVendorClass returns a boolean if a field has been set.

### GetValue

`func (o *NetworkcontainerOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NetworkcontainerOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NetworkcontainerOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NetworkcontainerOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUseOption

`func (o *NetworkcontainerOptions) GetUseOption() bool`

GetUseOption returns the UseOption field if non-nil, zero value otherwise.

### GetUseOptionOk

`func (o *NetworkcontainerOptions) GetUseOptionOk() (*bool, bool)`

GetUseOptionOk returns a tuple with the UseOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOption

`func (o *NetworkcontainerOptions) SetUseOption(v bool)`

SetUseOption sets UseOption field to given value.

### HasUseOption

`func (o *NetworkcontainerOptions) HasUseOption() bool`

HasUseOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


