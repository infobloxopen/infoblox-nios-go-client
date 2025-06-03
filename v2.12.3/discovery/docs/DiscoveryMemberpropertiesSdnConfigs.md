# DiscoveryMemberpropertiesSdnConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdnType** | Pointer to **string** | Type of SDN controller. | [optional] 
**Addresses** | Pointer to **[]string** | SDN controller IP addresses or FQDNs. | [optional] 
**NetworkView** | Pointer to **string** | The network view associated with SDN controller. | [optional] 
**Protocol** | Pointer to **string** | The connection protocol. Valid values are &#39;HTTP&#39; and &#39;HTTPS&#39;. | [optional] 
**Handle** | Pointer to **string** | Unique configuration handle. | [optional] 
**Password** | Pointer to **string** | SDN controller login password. | [optional] 
**Username** | Pointer to **string** | SDN controller login name. | [optional] 
**ApiKey** | Pointer to **string** | SDN controller API key. | [optional] 
**OnPrem** | Pointer to **bool** | True if controller is on-premises, otherwise we consider it is in cloud. | [optional] 
**UseGlobalProxy** | Pointer to **bool** | Use global proxy settings to access SDN controller. | [optional] 
**Comment** | Pointer to **string** | Additional information about the SDN configuration. | [optional] 
**NetworkInterfaceType** | Pointer to **string** | The type of the network interface on discovery member used for SDN controller discovery | [optional] 
**NetworkInterfaceVirtualIp** | Pointer to **string** | Virtual IP of VLAN network interface on discovery member | [optional] 
**Uuid** | Pointer to **string** | Unique key of SDN controller structure. Must be specified for existing sdn_config structures, otherwise they will be removed. If not specified, new structure will be created. | [optional] 

## Methods

### NewDiscoveryMemberpropertiesSdnConfigs

`func NewDiscoveryMemberpropertiesSdnConfigs() *DiscoveryMemberpropertiesSdnConfigs`

NewDiscoveryMemberpropertiesSdnConfigs instantiates a new DiscoveryMemberpropertiesSdnConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryMemberpropertiesSdnConfigsWithDefaults

`func NewDiscoveryMemberpropertiesSdnConfigsWithDefaults() *DiscoveryMemberpropertiesSdnConfigs`

NewDiscoveryMemberpropertiesSdnConfigsWithDefaults instantiates a new DiscoveryMemberpropertiesSdnConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdnType

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetSdnType() string`

GetSdnType returns the SdnType field if non-nil, zero value otherwise.

### GetSdnTypeOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetSdnTypeOk() (*string, bool)`

GetSdnTypeOk returns a tuple with the SdnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnType

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetSdnType(v string)`

SetSdnType sets SdnType field to given value.

### HasSdnType

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasSdnType() bool`

HasSdnType returns a boolean if a field has been set.

### GetAddresses

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetNetworkView

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetProtocol

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetHandle

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetPassword

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetApiKey

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetOnPrem

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetOnPrem() bool`

GetOnPrem returns the OnPrem field if non-nil, zero value otherwise.

### GetOnPremOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetOnPremOk() (*bool, bool)`

GetOnPremOk returns a tuple with the OnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPrem

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetOnPrem(v bool)`

SetOnPrem sets OnPrem field to given value.

### HasOnPrem

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasOnPrem() bool`

HasOnPrem returns a boolean if a field has been set.

### GetUseGlobalProxy

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetUseGlobalProxy() bool`

GetUseGlobalProxy returns the UseGlobalProxy field if non-nil, zero value otherwise.

### GetUseGlobalProxyOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetUseGlobalProxyOk() (*bool, bool)`

GetUseGlobalProxyOk returns a tuple with the UseGlobalProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalProxy

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetUseGlobalProxy(v bool)`

SetUseGlobalProxy sets UseGlobalProxy field to given value.

### HasUseGlobalProxy

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasUseGlobalProxy() bool`

HasUseGlobalProxy returns a boolean if a field has been set.

### GetComment

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetNetworkInterfaceType

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetNetworkInterfaceType() string`

GetNetworkInterfaceType returns the NetworkInterfaceType field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetNetworkInterfaceTypeOk() (*string, bool)`

GetNetworkInterfaceTypeOk returns a tuple with the NetworkInterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceType

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetNetworkInterfaceType(v string)`

SetNetworkInterfaceType sets NetworkInterfaceType field to given value.

### HasNetworkInterfaceType

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasNetworkInterfaceType() bool`

HasNetworkInterfaceType returns a boolean if a field has been set.

### GetNetworkInterfaceVirtualIp

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetNetworkInterfaceVirtualIp() string`

GetNetworkInterfaceVirtualIp returns the NetworkInterfaceVirtualIp field if non-nil, zero value otherwise.

### GetNetworkInterfaceVirtualIpOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetNetworkInterfaceVirtualIpOk() (*string, bool)`

GetNetworkInterfaceVirtualIpOk returns a tuple with the NetworkInterfaceVirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceVirtualIp

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetNetworkInterfaceVirtualIp(v string)`

SetNetworkInterfaceVirtualIp sets NetworkInterfaceVirtualIp field to given value.

### HasNetworkInterfaceVirtualIp

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasNetworkInterfaceVirtualIp() bool`

HasNetworkInterfaceVirtualIp returns a boolean if a field has been set.

### GetUuid

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DiscoveryMemberpropertiesSdnConfigs) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DiscoveryMemberpropertiesSdnConfigs) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DiscoveryMemberpropertiesSdnConfigs) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


