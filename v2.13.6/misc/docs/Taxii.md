# Taxii

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**EnableService** | Pointer to **bool** | Indicates whether the Taxii service is running on the given member or not. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the Grid member. | [optional] [readonly] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Grid member. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Taxii Member. | [optional] [readonly] 
**TaxiiRpzConfig** | Pointer to [**[]TaxiiTaxiiRpzConfig**](TaxiiTaxiiRpzConfig.md) | Taxii service RPZ configuration list. | [optional] 

## Methods

### NewTaxii

`func NewTaxii() *Taxii`

NewTaxii instantiates a new Taxii object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxiiWithDefaults

`func NewTaxiiWithDefaults() *Taxii`

NewTaxiiWithDefaults instantiates a new Taxii object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Taxii) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Taxii) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Taxii) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Taxii) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnableService

`func (o *Taxii) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *Taxii) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *Taxii) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *Taxii) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetIpv4addr

`func (o *Taxii) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *Taxii) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *Taxii) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *Taxii) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6addr

`func (o *Taxii) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *Taxii) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *Taxii) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *Taxii) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *Taxii) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Taxii) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Taxii) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Taxii) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxiiRpzConfig

`func (o *Taxii) GetTaxiiRpzConfig() []TaxiiTaxiiRpzConfig`

GetTaxiiRpzConfig returns the TaxiiRpzConfig field if non-nil, zero value otherwise.

### GetTaxiiRpzConfigOk

`func (o *Taxii) GetTaxiiRpzConfigOk() (*[]TaxiiTaxiiRpzConfig, bool)`

GetTaxiiRpzConfigOk returns a tuple with the TaxiiRpzConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxiiRpzConfig

`func (o *Taxii) SetTaxiiRpzConfig(v []TaxiiTaxiiRpzConfig)`

SetTaxiiRpzConfig sets TaxiiRpzConfig field to given value.

### HasTaxiiRpzConfig

`func (o *Taxii) HasTaxiiRpzConfig() bool`

HasTaxiiRpzConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


