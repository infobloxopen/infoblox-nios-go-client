# MemberLomNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 Address of the Grid member. | [optional] 
**Gateway** | Pointer to **string** | The default gateway for the Grid member. | [optional] 
**SubnetMask** | Pointer to **string** | The subnet mask for the Grid member. | [optional] 
**IsLomCapable** | Pointer to **bool** | Determines if the physical node supports LOM or not. | [optional] [readonly] 

## Methods

### NewMemberLomNetworkConfig

`func NewMemberLomNetworkConfig() *MemberLomNetworkConfig`

NewMemberLomNetworkConfig instantiates a new MemberLomNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberLomNetworkConfigWithDefaults

`func NewMemberLomNetworkConfigWithDefaults() *MemberLomNetworkConfig`

NewMemberLomNetworkConfigWithDefaults instantiates a new MemberLomNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberLomNetworkConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberLomNetworkConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberLomNetworkConfig) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberLomNetworkConfig) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *MemberLomNetworkConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MemberLomNetworkConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MemberLomNetworkConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MemberLomNetworkConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetSubnetMask

`func (o *MemberLomNetworkConfig) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *MemberLomNetworkConfig) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *MemberLomNetworkConfig) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *MemberLomNetworkConfig) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetIsLomCapable

`func (o *MemberLomNetworkConfig) GetIsLomCapable() bool`

GetIsLomCapable returns the IsLomCapable field if non-nil, zero value otherwise.

### GetIsLomCapableOk

`func (o *MemberLomNetworkConfig) GetIsLomCapableOk() (*bool, bool)`

GetIsLomCapableOk returns a tuple with the IsLomCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLomCapable

`func (o *MemberLomNetworkConfig) SetIsLomCapable(v bool)`

SetIsLomCapable sets IsLomCapable field to given value.

### HasIsLomCapable

`func (o *MemberLomNetworkConfig) HasIsLomCapable() bool`

HasIsLomCapable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


