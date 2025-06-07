# MemberStaticRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 Address of the Grid Member. | [optional] 
**Gateway** | Pointer to **string** | The default gateway for the Grid Member. | [optional] 
**SubnetMask** | Pointer to **string** | The subnet mask for the Grid Member. | [optional] 
**VlanId** | Pointer to **int64** | The identifier for the VLAN. Valid values are from 1 to 4096. | [optional] 
**Primary** | Pointer to **bool** | Determines if the current address is the primary VLAN address or not. | [optional] 
**Dscp** | Pointer to **int64** | The DSCP (Differentiated Services Code Point) value determines relative priorities for the type of services on your network. The appliance implements QoS (Quality of Service) rules based on this configuration. Valid values are from 0 to 63. | [optional] 
**LanSubnetMask** | Pointer to **string** | LAN netmask only for GCP HA. | [optional] 
**LanGateway** | Pointer to **string** | LAN gateway only for GCP HA. | [optional] 
**UseDscp** | Pointer to **bool** | Use flag for: dscp | [optional] 

## Methods

### NewMemberStaticRoutes

`func NewMemberStaticRoutes() *MemberStaticRoutes`

NewMemberStaticRoutes instantiates a new MemberStaticRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberStaticRoutesWithDefaults

`func NewMemberStaticRoutesWithDefaults() *MemberStaticRoutes`

NewMemberStaticRoutesWithDefaults instantiates a new MemberStaticRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberStaticRoutes) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberStaticRoutes) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberStaticRoutes) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberStaticRoutes) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *MemberStaticRoutes) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MemberStaticRoutes) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MemberStaticRoutes) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MemberStaticRoutes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetSubnetMask

`func (o *MemberStaticRoutes) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *MemberStaticRoutes) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *MemberStaticRoutes) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *MemberStaticRoutes) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetVlanId

`func (o *MemberStaticRoutes) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *MemberStaticRoutes) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *MemberStaticRoutes) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *MemberStaticRoutes) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPrimary

`func (o *MemberStaticRoutes) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MemberStaticRoutes) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MemberStaticRoutes) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MemberStaticRoutes) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetDscp

`func (o *MemberStaticRoutes) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *MemberStaticRoutes) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *MemberStaticRoutes) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *MemberStaticRoutes) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetLanSubnetMask

`func (o *MemberStaticRoutes) GetLanSubnetMask() string`

GetLanSubnetMask returns the LanSubnetMask field if non-nil, zero value otherwise.

### GetLanSubnetMaskOk

`func (o *MemberStaticRoutes) GetLanSubnetMaskOk() (*string, bool)`

GetLanSubnetMaskOk returns a tuple with the LanSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSubnetMask

`func (o *MemberStaticRoutes) SetLanSubnetMask(v string)`

SetLanSubnetMask sets LanSubnetMask field to given value.

### HasLanSubnetMask

`func (o *MemberStaticRoutes) HasLanSubnetMask() bool`

HasLanSubnetMask returns a boolean if a field has been set.

### GetLanGateway

`func (o *MemberStaticRoutes) GetLanGateway() string`

GetLanGateway returns the LanGateway field if non-nil, zero value otherwise.

### GetLanGatewayOk

`func (o *MemberStaticRoutes) GetLanGatewayOk() (*string, bool)`

GetLanGatewayOk returns a tuple with the LanGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanGateway

`func (o *MemberStaticRoutes) SetLanGateway(v string)`

SetLanGateway sets LanGateway field to given value.

### HasLanGateway

`func (o *MemberStaticRoutes) HasLanGateway() bool`

HasLanGateway returns a boolean if a field has been set.

### GetUseDscp

`func (o *MemberStaticRoutes) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *MemberStaticRoutes) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *MemberStaticRoutes) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *MemberStaticRoutes) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


