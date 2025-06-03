# Ipv6fixedaddresstemplateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the DHCP option. | [optional] 
**Num** | Pointer to **int64** | The code of the DHCP option. | [optional] 
**VendorClass** | Pointer to **string** | The name of the space this DHCP option is associated to. | [optional] 
**Value** | Pointer to **string** | Value of the DHCP option | [optional] 
**UseOption** | Pointer to **bool** | Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers | [optional] 

## Methods

### NewIpv6fixedaddresstemplateOptions

`func NewIpv6fixedaddresstemplateOptions() *Ipv6fixedaddresstemplateOptions`

NewIpv6fixedaddresstemplateOptions instantiates a new Ipv6fixedaddresstemplateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6fixedaddresstemplateOptionsWithDefaults

`func NewIpv6fixedaddresstemplateOptionsWithDefaults() *Ipv6fixedaddresstemplateOptions`

NewIpv6fixedaddresstemplateOptionsWithDefaults instantiates a new Ipv6fixedaddresstemplateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Ipv6fixedaddresstemplateOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6fixedaddresstemplateOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6fixedaddresstemplateOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6fixedaddresstemplateOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNum

`func (o *Ipv6fixedaddresstemplateOptions) GetNum() int64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *Ipv6fixedaddresstemplateOptions) GetNumOk() (*int64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *Ipv6fixedaddresstemplateOptions) SetNum(v int64)`

SetNum sets Num field to given value.

### HasNum

`func (o *Ipv6fixedaddresstemplateOptions) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetVendorClass

`func (o *Ipv6fixedaddresstemplateOptions) GetVendorClass() string`

GetVendorClass returns the VendorClass field if non-nil, zero value otherwise.

### GetVendorClassOk

`func (o *Ipv6fixedaddresstemplateOptions) GetVendorClassOk() (*string, bool)`

GetVendorClassOk returns a tuple with the VendorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorClass

`func (o *Ipv6fixedaddresstemplateOptions) SetVendorClass(v string)`

SetVendorClass sets VendorClass field to given value.

### HasVendorClass

`func (o *Ipv6fixedaddresstemplateOptions) HasVendorClass() bool`

HasVendorClass returns a boolean if a field has been set.

### GetValue

`func (o *Ipv6fixedaddresstemplateOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Ipv6fixedaddresstemplateOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Ipv6fixedaddresstemplateOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Ipv6fixedaddresstemplateOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUseOption

`func (o *Ipv6fixedaddresstemplateOptions) GetUseOption() bool`

GetUseOption returns the UseOption field if non-nil, zero value otherwise.

### GetUseOptionOk

`func (o *Ipv6fixedaddresstemplateOptions) GetUseOptionOk() (*bool, bool)`

GetUseOptionOk returns a tuple with the UseOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOption

`func (o *Ipv6fixedaddresstemplateOptions) SetUseOption(v bool)`

SetUseOption sets UseOption field to given value.

### HasUseOption

`func (o *Ipv6fixedaddresstemplateOptions) HasUseOption() bool`

HasUseOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


