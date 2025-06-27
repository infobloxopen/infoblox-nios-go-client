# DxlEndpointBrokers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **string** | The FQDN for the DXL endpoint broker. | [optional] 
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address for the DXL endpoint broker. | [optional] 
**Port** | Pointer to **int64** | The communication port for the DXL endpoint broker. | [optional] 
**UniqueId** | Pointer to **string** | The unique identifier for the DXL endpoint. | [optional] 

## Methods

### NewDxlEndpointBrokers

`func NewDxlEndpointBrokers() *DxlEndpointBrokers`

NewDxlEndpointBrokers instantiates a new DxlEndpointBrokers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDxlEndpointBrokersWithDefaults

`func NewDxlEndpointBrokersWithDefaults() *DxlEndpointBrokers`

NewDxlEndpointBrokersWithDefaults instantiates a new DxlEndpointBrokers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *DxlEndpointBrokers) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *DxlEndpointBrokers) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *DxlEndpointBrokers) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *DxlEndpointBrokers) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetAddress

`func (o *DxlEndpointBrokers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DxlEndpointBrokers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DxlEndpointBrokers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DxlEndpointBrokers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *DxlEndpointBrokers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DxlEndpointBrokers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DxlEndpointBrokers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DxlEndpointBrokers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUniqueId

`func (o *DxlEndpointBrokers) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *DxlEndpointBrokers) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *DxlEndpointBrokers) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *DxlEndpointBrokers) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


