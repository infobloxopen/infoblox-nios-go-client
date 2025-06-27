# TftpfiledirVtftpDirMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to **string** | The Grid member on which to create the virtual TFTP directory. | [optional] 
**IpType** | Pointer to **string** | The IP type of the virtual TFTP root directory. | [optional] 
**Address** | Pointer to **string** | The IP address of the clients which will see the virtual TFTP directory as the root directory. | [optional] 
**StartAddress** | Pointer to **string** | The start IP address of the range within which the clients will see the virtual TFTP directory as the root directory. | [optional] 
**EndAddress** | Pointer to **string** | The end IP address of the range within which the clients will see the virtual TFTP directory as the root directory. | [optional] 
**Network** | Pointer to **string** | The IP address of network the clients from which will see the virtual TFTP directory as the root directory. | [optional] 
**Cidr** | Pointer to **int64** | The CIDR of network the clients from which will see the virtual TFTP directory as the root directory. | [optional] 

## Methods

### NewTftpfiledirVtftpDirMembers

`func NewTftpfiledirVtftpDirMembers() *TftpfiledirVtftpDirMembers`

NewTftpfiledirVtftpDirMembers instantiates a new TftpfiledirVtftpDirMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTftpfiledirVtftpDirMembersWithDefaults

`func NewTftpfiledirVtftpDirMembersWithDefaults() *TftpfiledirVtftpDirMembers`

NewTftpfiledirVtftpDirMembersWithDefaults instantiates a new TftpfiledirVtftpDirMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *TftpfiledirVtftpDirMembers) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TftpfiledirVtftpDirMembers) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TftpfiledirVtftpDirMembers) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *TftpfiledirVtftpDirMembers) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetIpType

`func (o *TftpfiledirVtftpDirMembers) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *TftpfiledirVtftpDirMembers) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *TftpfiledirVtftpDirMembers) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *TftpfiledirVtftpDirMembers) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetAddress

`func (o *TftpfiledirVtftpDirMembers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TftpfiledirVtftpDirMembers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TftpfiledirVtftpDirMembers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TftpfiledirVtftpDirMembers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStartAddress

`func (o *TftpfiledirVtftpDirMembers) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *TftpfiledirVtftpDirMembers) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *TftpfiledirVtftpDirMembers) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *TftpfiledirVtftpDirMembers) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *TftpfiledirVtftpDirMembers) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *TftpfiledirVtftpDirMembers) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *TftpfiledirVtftpDirMembers) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *TftpfiledirVtftpDirMembers) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetNetwork

`func (o *TftpfiledirVtftpDirMembers) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TftpfiledirVtftpDirMembers) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TftpfiledirVtftpDirMembers) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *TftpfiledirVtftpDirMembers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCidr

`func (o *TftpfiledirVtftpDirMembers) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *TftpfiledirVtftpDirMembers) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *TftpfiledirVtftpDirMembers) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *TftpfiledirVtftpDirMembers) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


