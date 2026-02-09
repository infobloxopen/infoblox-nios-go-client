# FixedaddressMsServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address or FQDN of the Microsoft server. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the Microsoft server. | [optional] [readonly] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Microsoft server. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Microsoft server. | [optional] [readonly] 

## Methods

### NewFixedaddressMsServer

`func NewFixedaddressMsServer() *FixedaddressMsServer`

NewFixedaddressMsServer instantiates a new FixedaddressMsServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedaddressMsServerWithDefaults

`func NewFixedaddressMsServerWithDefaults() *FixedaddressMsServer`

NewFixedaddressMsServerWithDefaults instantiates a new FixedaddressMsServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FixedaddressMsServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FixedaddressMsServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FixedaddressMsServer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FixedaddressMsServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIpv4addr

`func (o *FixedaddressMsServer) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *FixedaddressMsServer) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *FixedaddressMsServer) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *FixedaddressMsServer) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6addr

`func (o *FixedaddressMsServer) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *FixedaddressMsServer) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *FixedaddressMsServer) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *FixedaddressMsServer) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *FixedaddressMsServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixedaddressMsServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FixedaddressMsServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FixedaddressMsServer) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


