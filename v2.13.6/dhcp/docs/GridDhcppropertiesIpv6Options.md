# GridDhcppropertiesIpv6Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the DHCP option. | [optional] 
**Num** | Pointer to **int64** | The code of the DHCP option. | [optional] 
**VendorClass** | Pointer to **string** | The name of the space this DHCP option is associated to. | [optional] 
**Value** | Pointer to **string** | Value of the DHCP option | [optional] 
**UseOption** | Pointer to **bool** | Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers | [optional] 

## Methods

### NewGridDhcppropertiesIpv6Options

`func NewGridDhcppropertiesIpv6Options() *GridDhcppropertiesIpv6Options`

NewGridDhcppropertiesIpv6Options instantiates a new GridDhcppropertiesIpv6Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDhcppropertiesIpv6OptionsWithDefaults

`func NewGridDhcppropertiesIpv6OptionsWithDefaults() *GridDhcppropertiesIpv6Options`

NewGridDhcppropertiesIpv6OptionsWithDefaults instantiates a new GridDhcppropertiesIpv6Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GridDhcppropertiesIpv6Options) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridDhcppropertiesIpv6Options) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridDhcppropertiesIpv6Options) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridDhcppropertiesIpv6Options) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNum

`func (o *GridDhcppropertiesIpv6Options) GetNum() int64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *GridDhcppropertiesIpv6Options) GetNumOk() (*int64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *GridDhcppropertiesIpv6Options) SetNum(v int64)`

SetNum sets Num field to given value.

### HasNum

`func (o *GridDhcppropertiesIpv6Options) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetVendorClass

`func (o *GridDhcppropertiesIpv6Options) GetVendorClass() string`

GetVendorClass returns the VendorClass field if non-nil, zero value otherwise.

### GetVendorClassOk

`func (o *GridDhcppropertiesIpv6Options) GetVendorClassOk() (*string, bool)`

GetVendorClassOk returns a tuple with the VendorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorClass

`func (o *GridDhcppropertiesIpv6Options) SetVendorClass(v string)`

SetVendorClass sets VendorClass field to given value.

### HasVendorClass

`func (o *GridDhcppropertiesIpv6Options) HasVendorClass() bool`

HasVendorClass returns a boolean if a field has been set.

### GetValue

`func (o *GridDhcppropertiesIpv6Options) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GridDhcppropertiesIpv6Options) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GridDhcppropertiesIpv6Options) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GridDhcppropertiesIpv6Options) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUseOption

`func (o *GridDhcppropertiesIpv6Options) GetUseOption() bool`

GetUseOption returns the UseOption field if non-nil, zero value otherwise.

### GetUseOptionOk

`func (o *GridDhcppropertiesIpv6Options) GetUseOptionOk() (*bool, bool)`

GetUseOptionOk returns a tuple with the UseOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOption

`func (o *GridDhcppropertiesIpv6Options) SetUseOption(v bool)`

SetUseOption sets UseOption field to given value.

### HasUseOption

`func (o *GridDhcppropertiesIpv6Options) HasUseOption() bool`

HasUseOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


