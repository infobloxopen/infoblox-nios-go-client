# MemberIpv6StaticRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | IPv6 address. | [optional] 
**Cidr** | Pointer to **int64** | IPv6 CIDR | [optional] 
**Gateway** | Pointer to **string** | Gateway address. | [optional] 

## Methods

### NewMemberIpv6StaticRoutes

`func NewMemberIpv6StaticRoutes() *MemberIpv6StaticRoutes`

NewMemberIpv6StaticRoutes instantiates a new MemberIpv6StaticRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberIpv6StaticRoutesWithDefaults

`func NewMemberIpv6StaticRoutesWithDefaults() *MemberIpv6StaticRoutes`

NewMemberIpv6StaticRoutesWithDefaults instantiates a new MemberIpv6StaticRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberIpv6StaticRoutes) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberIpv6StaticRoutes) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberIpv6StaticRoutes) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberIpv6StaticRoutes) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCidr

`func (o *MemberIpv6StaticRoutes) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *MemberIpv6StaticRoutes) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *MemberIpv6StaticRoutes) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *MemberIpv6StaticRoutes) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *MemberIpv6StaticRoutes) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MemberIpv6StaticRoutes) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MemberIpv6StaticRoutes) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MemberIpv6StaticRoutes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


