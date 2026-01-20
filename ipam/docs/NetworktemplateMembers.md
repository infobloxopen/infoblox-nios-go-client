# NetworktemplateMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Struct** | Pointer to **string** | The struct type of the object. The value must be one of &#39;dhcpmember&#39; and &#39;msdhcpserver&#39;. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address or FQDN of the Microsoft server. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Grid Member. | [optional] 
**Name** | Pointer to **string** | The Grid member name | [optional] 

## Methods

### NewNetworktemplateMembers

`func NewNetworktemplateMembers() *NetworktemplateMembers`

NewNetworktemplateMembers instantiates a new NetworktemplateMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworktemplateMembersWithDefaults

`func NewNetworktemplateMembersWithDefaults() *NetworktemplateMembers`

NewNetworktemplateMembersWithDefaults instantiates a new NetworktemplateMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStruct

`func (o *NetworktemplateMembers) GetStruct() string`

GetStruct returns the Struct field if non-nil, zero value otherwise.

### GetStructOk

`func (o *NetworktemplateMembers) GetStructOk() (*string, bool)`

GetStructOk returns a tuple with the Struct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStruct

`func (o *NetworktemplateMembers) SetStruct(v string)`

SetStruct sets Struct field to given value.

### HasStruct

`func (o *NetworktemplateMembers) HasStruct() bool`

HasStruct returns a boolean if a field has been set.

### GetIpv4addr

`func (o *NetworktemplateMembers) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *NetworktemplateMembers) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *NetworktemplateMembers) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *NetworktemplateMembers) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6addr

`func (o *NetworktemplateMembers) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *NetworktemplateMembers) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *NetworktemplateMembers) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *NetworktemplateMembers) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *NetworktemplateMembers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworktemplateMembers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworktemplateMembers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworktemplateMembers) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


