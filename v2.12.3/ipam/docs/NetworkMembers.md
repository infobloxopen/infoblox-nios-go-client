# NetworkMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4addr** | Pointer to **string** | The IPv4 Address or FQDN of the Microsoft server. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Grid Member. | [optional] 
**Name** | Pointer to **string** | The Grid member name | [optional] 

## Methods

### NewNetworkMembers

`func NewNetworkMembers() *NetworkMembers`

NewNetworkMembers instantiates a new NetworkMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMembersWithDefaults

`func NewNetworkMembersWithDefaults() *NetworkMembers`

NewNetworkMembersWithDefaults instantiates a new NetworkMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4addr

`func (o *NetworkMembers) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *NetworkMembers) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *NetworkMembers) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *NetworkMembers) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6addr

`func (o *NetworkMembers) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *NetworkMembers) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *NetworkMembers) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *NetworkMembers) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *NetworkMembers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkMembers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkMembers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkMembers) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


