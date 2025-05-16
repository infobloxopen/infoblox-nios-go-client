# MemberadditionaliplistIpv4NetworkSetting

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

### NewMemberadditionaliplistIpv4NetworkSetting

`func NewMemberadditionaliplistIpv4NetworkSetting() *MemberadditionaliplistIpv4NetworkSetting`

NewMemberadditionaliplistIpv4NetworkSetting instantiates a new MemberadditionaliplistIpv4NetworkSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberadditionaliplistIpv4NetworkSettingWithDefaults

`func NewMemberadditionaliplistIpv4NetworkSettingWithDefaults() *MemberadditionaliplistIpv4NetworkSetting`

NewMemberadditionaliplistIpv4NetworkSettingWithDefaults instantiates a new MemberadditionaliplistIpv4NetworkSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetSubnetMask

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetVlanId

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPrimary

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetDscp

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetLanSubnetMask

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanSubnetMask() string`

GetLanSubnetMask returns the LanSubnetMask field if non-nil, zero value otherwise.

### GetLanSubnetMaskOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanSubnetMaskOk() (*string, bool)`

GetLanSubnetMaskOk returns a tuple with the LanSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSubnetMask

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetLanSubnetMask(v string)`

SetLanSubnetMask sets LanSubnetMask field to given value.

### HasLanSubnetMask

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasLanSubnetMask() bool`

HasLanSubnetMask returns a boolean if a field has been set.

### GetLanGateway

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanGateway() string`

GetLanGateway returns the LanGateway field if non-nil, zero value otherwise.

### GetLanGatewayOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanGatewayOk() (*string, bool)`

GetLanGatewayOk returns a tuple with the LanGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanGateway

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetLanGateway(v string)`

SetLanGateway sets LanGateway field to given value.

### HasLanGateway

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasLanGateway() bool`

HasLanGateway returns a boolean if a field has been set.

### GetUseDscp

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *MemberadditionaliplistIpv4NetworkSetting) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *MemberadditionaliplistIpv4NetworkSetting) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *MemberadditionaliplistIpv4NetworkSetting) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


