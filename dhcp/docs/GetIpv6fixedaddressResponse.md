# GetIpv6fixedaddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddressType** | Pointer to **string** | The address type value for this IPv6 fixed address. When the address type is \&quot;ADDRESS\&quot;, a value for the &#39;ipv6addr&#39; member is required. When the address type is \&quot;PREFIX\&quot;, values for &#39;ipv6prefix&#39; and &#39;ipv6prefix_bits&#39; are required. When the address type is \&quot;BOTH\&quot;, values for &#39;ipv6addr&#39;, &#39;ipv6prefix&#39;, and &#39;ipv6prefix_bits&#39; are all required. | [optional] 
**AllowTelnet** | Pointer to **bool** | This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH. | [optional] 
**CliCredentials** | Pointer to [**[]Ipv6fixedaddressCliCredentials**](Ipv6fixedaddressCliCredentials.md) | The CLI credentials for the IPv6 fixed address. | [optional] 
**CloudInfo** | Pointer to [**Ipv6fixedaddressCloudInfo**](Ipv6fixedaddressCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the fixed address; maximum 256 characters. | [optional] 
**DeviceDescription** | Pointer to **string** | The description of the device. | [optional] 
**DeviceLocation** | Pointer to **string** | The location of the device. | [optional] 
**DeviceType** | Pointer to **string** | The type of the device. | [optional] 
**DeviceVendor** | Pointer to **string** | The vendor of the device. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a fixed address is disabled or not. When this is set to False, the IPv6 fixed address is enabled. | [optional] 
**DisableDiscovery** | Pointer to **bool** | Determines if the discovery for this IPv6 fixed address is disabled or not. False means that the discovery is enabled. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | The discovery status of this IPv6 fixed address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**Ipv6fixedaddressDiscoveredData**](Ipv6fixedaddressDiscoveredData.md) |  | [optional] 
**DomainName** | Pointer to **string** | The domain name for this IPv6 fixed address. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages. | [optional] 
**Duid** | Pointer to **string** | The DUID value for this IPv6 fixed address. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the IPv6 fixed address should be immediately enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv6addr** | Pointer to [**Ipv6fixedaddressIpv6addr**](Ipv6fixedaddressIpv6addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**Ipv6prefix** | Pointer to **string** | The IPv6 Address prefix of the DHCP IPv6 fixed address. | [optional] 
**Ipv6prefixBits** | Pointer to **int64** | Prefix bits of the DHCP IPv6 fixed address. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6fixedaddressLogicFilterRules**](Ipv6fixedaddressLogicFilterRules.md) | This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address for this IPv6 fixed address. | [optional] 
**MatchClient** | Pointer to **string** | The match_client value for this fixed address. Valid values are: \&quot;DUID\&quot;: The fixed IP address is leased to the matching DUID. \&quot;MAC_ADDRESS\&quot;: The fixed IP address is leased to the matching MAC address. | [optional] 
**MsAdUserData** | Pointer to [**Ipv6fixedaddressMsAdUserData**](Ipv6fixedaddressMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | This field contains the name of this IPv6 fixed address. | [optional] 
**Network** | Pointer to **string** | The network to which this IPv6 fixed address belongs, in IPv6 Address/CIDR format. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this IPv6 fixed address resides. | [optional] 
**Options** | Pointer to [**[]Ipv6fixedaddressOptions**](Ipv6fixedaddressOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PreferredLifetime** | Pointer to **int64** | The preferred lifetime value for this DHCP IPv6 fixed address object. | [optional] 
**ReservedInterface** | Pointer to **string** | The reference to the reserved interface to which the device belongs. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. The restart_if_needed flag can trigger a restart on DHCP services only when it is enabled on CP member. | [optional] 
**Snmp3Credential** | Pointer to [**Ipv6fixedaddressSnmp3Credential**](Ipv6fixedaddressSnmp3Credential.md) |  | [optional] 
**SnmpCredential** | Pointer to [**Ipv6fixedaddressSnmpCredential**](Ipv6fixedaddressSnmpCredential.md) |  | [optional] 
**Template** | Pointer to **string** | If set on creation, the IPv6 fixed address will be created according to the values specified in the named template. | [optional] 
**UseCliCredentials** | Pointer to **bool** | If set to true, the CLI credential will override member-level settings. | [optional] 
**UseDomainName** | Pointer to **bool** | Use flag for: domain_name | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseSnmp3Credential** | Pointer to **bool** | Determines if the SNMPv3 credential should be used for the IPv6 fixed address. | [optional] 
**UseSnmpCredential** | Pointer to **bool** | If set to true, SNMP credential will override member level settings. | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**ValidLifetime** | Pointer to **int64** | The valid lifetime value for this DHCP IPv6 Fixed Address object. | [optional] 
**Result** | Pointer to [**Ipv6fixedaddress**](Ipv6fixedaddress.md) |  | [optional] 

## Methods

### NewGetIpv6fixedaddressResponse

`func NewGetIpv6fixedaddressResponse() *GetIpv6fixedaddressResponse`

NewGetIpv6fixedaddressResponse instantiates a new GetIpv6fixedaddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6fixedaddressResponseWithDefaults

`func NewGetIpv6fixedaddressResponseWithDefaults() *GetIpv6fixedaddressResponse`

NewGetIpv6fixedaddressResponseWithDefaults instantiates a new GetIpv6fixedaddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6fixedaddressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6fixedaddressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6fixedaddressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6fixedaddressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddressType

`func (o *GetIpv6fixedaddressResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *GetIpv6fixedaddressResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *GetIpv6fixedaddressResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *GetIpv6fixedaddressResponse) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAllowTelnet

`func (o *GetIpv6fixedaddressResponse) GetAllowTelnet() bool`

GetAllowTelnet returns the AllowTelnet field if non-nil, zero value otherwise.

### GetAllowTelnetOk

`func (o *GetIpv6fixedaddressResponse) GetAllowTelnetOk() (*bool, bool)`

GetAllowTelnetOk returns a tuple with the AllowTelnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTelnet

`func (o *GetIpv6fixedaddressResponse) SetAllowTelnet(v bool)`

SetAllowTelnet sets AllowTelnet field to given value.

### HasAllowTelnet

`func (o *GetIpv6fixedaddressResponse) HasAllowTelnet() bool`

HasAllowTelnet returns a boolean if a field has been set.

### GetCliCredentials

`func (o *GetIpv6fixedaddressResponse) GetCliCredentials() []Ipv6fixedaddressCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *GetIpv6fixedaddressResponse) GetCliCredentialsOk() (*[]Ipv6fixedaddressCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *GetIpv6fixedaddressResponse) SetCliCredentials(v []Ipv6fixedaddressCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *GetIpv6fixedaddressResponse) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetIpv6fixedaddressResponse) GetCloudInfo() Ipv6fixedaddressCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetIpv6fixedaddressResponse) GetCloudInfoOk() (*Ipv6fixedaddressCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetIpv6fixedaddressResponse) SetCloudInfo(v Ipv6fixedaddressCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetIpv6fixedaddressResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6fixedaddressResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6fixedaddressResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6fixedaddressResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6fixedaddressResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeviceDescription

`func (o *GetIpv6fixedaddressResponse) GetDeviceDescription() string`

GetDeviceDescription returns the DeviceDescription field if non-nil, zero value otherwise.

### GetDeviceDescriptionOk

`func (o *GetIpv6fixedaddressResponse) GetDeviceDescriptionOk() (*string, bool)`

GetDeviceDescriptionOk returns a tuple with the DeviceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDescription

`func (o *GetIpv6fixedaddressResponse) SetDeviceDescription(v string)`

SetDeviceDescription sets DeviceDescription field to given value.

### HasDeviceDescription

`func (o *GetIpv6fixedaddressResponse) HasDeviceDescription() bool`

HasDeviceDescription returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *GetIpv6fixedaddressResponse) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *GetIpv6fixedaddressResponse) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *GetIpv6fixedaddressResponse) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *GetIpv6fixedaddressResponse) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceType

`func (o *GetIpv6fixedaddressResponse) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GetIpv6fixedaddressResponse) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GetIpv6fixedaddressResponse) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *GetIpv6fixedaddressResponse) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *GetIpv6fixedaddressResponse) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *GetIpv6fixedaddressResponse) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *GetIpv6fixedaddressResponse) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *GetIpv6fixedaddressResponse) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDisable

`func (o *GetIpv6fixedaddressResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetIpv6fixedaddressResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetIpv6fixedaddressResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetIpv6fixedaddressResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableDiscovery

`func (o *GetIpv6fixedaddressResponse) GetDisableDiscovery() bool`

GetDisableDiscovery returns the DisableDiscovery field if non-nil, zero value otherwise.

### GetDisableDiscoveryOk

`func (o *GetIpv6fixedaddressResponse) GetDisableDiscoveryOk() (*bool, bool)`

GetDisableDiscoveryOk returns a tuple with the DisableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDiscovery

`func (o *GetIpv6fixedaddressResponse) SetDisableDiscovery(v bool)`

SetDisableDiscovery sets DisableDiscovery field to given value.

### HasDisableDiscovery

`func (o *GetIpv6fixedaddressResponse) HasDisableDiscovery() bool`

HasDisableDiscovery returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetIpv6fixedaddressResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetIpv6fixedaddressResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetIpv6fixedaddressResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetIpv6fixedaddressResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *GetIpv6fixedaddressResponse) GetDiscoveredData() Ipv6fixedaddressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *GetIpv6fixedaddressResponse) GetDiscoveredDataOk() (*Ipv6fixedaddressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *GetIpv6fixedaddressResponse) SetDiscoveredData(v Ipv6fixedaddressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *GetIpv6fixedaddressResponse) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDomainName

`func (o *GetIpv6fixedaddressResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetIpv6fixedaddressResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetIpv6fixedaddressResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetIpv6fixedaddressResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *GetIpv6fixedaddressResponse) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *GetIpv6fixedaddressResponse) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *GetIpv6fixedaddressResponse) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *GetIpv6fixedaddressResponse) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetDuid

`func (o *GetIpv6fixedaddressResponse) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *GetIpv6fixedaddressResponse) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *GetIpv6fixedaddressResponse) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *GetIpv6fixedaddressResponse) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *GetIpv6fixedaddressResponse) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *GetIpv6fixedaddressResponse) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *GetIpv6fixedaddressResponse) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *GetIpv6fixedaddressResponse) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetIpv6fixedaddressResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetIpv6fixedaddressResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetIpv6fixedaddressResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetIpv6fixedaddressResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetIpv6fixedaddressResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetIpv6fixedaddressResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetIpv6fixedaddressResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetIpv6fixedaddressResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetIpv6fixedaddressResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetIpv6fixedaddressResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetIpv6fixedaddressResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetIpv6fixedaddressResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv6addr

`func (o *GetIpv6fixedaddressResponse) GetIpv6addr() Ipv6fixedaddressIpv6addr`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *GetIpv6fixedaddressResponse) GetIpv6addrOk() (*Ipv6fixedaddressIpv6addr, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *GetIpv6fixedaddressResponse) SetIpv6addr(v Ipv6fixedaddressIpv6addr)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *GetIpv6fixedaddressResponse) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *GetIpv6fixedaddressResponse) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *GetIpv6fixedaddressResponse) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *GetIpv6fixedaddressResponse) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *GetIpv6fixedaddressResponse) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetIpv6prefix

`func (o *GetIpv6fixedaddressResponse) GetIpv6prefix() string`

GetIpv6prefix returns the Ipv6prefix field if non-nil, zero value otherwise.

### GetIpv6prefixOk

`func (o *GetIpv6fixedaddressResponse) GetIpv6prefixOk() (*string, bool)`

GetIpv6prefixOk returns a tuple with the Ipv6prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefix

`func (o *GetIpv6fixedaddressResponse) SetIpv6prefix(v string)`

SetIpv6prefix sets Ipv6prefix field to given value.

### HasIpv6prefix

`func (o *GetIpv6fixedaddressResponse) HasIpv6prefix() bool`

HasIpv6prefix returns a boolean if a field has been set.

### GetIpv6prefixBits

`func (o *GetIpv6fixedaddressResponse) GetIpv6prefixBits() int64`

GetIpv6prefixBits returns the Ipv6prefixBits field if non-nil, zero value otherwise.

### GetIpv6prefixBitsOk

`func (o *GetIpv6fixedaddressResponse) GetIpv6prefixBitsOk() (*int64, bool)`

GetIpv6prefixBitsOk returns a tuple with the Ipv6prefixBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefixBits

`func (o *GetIpv6fixedaddressResponse) SetIpv6prefixBits(v int64)`

SetIpv6prefixBits sets Ipv6prefixBits field to given value.

### HasIpv6prefixBits

`func (o *GetIpv6fixedaddressResponse) HasIpv6prefixBits() bool`

HasIpv6prefixBits returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetIpv6fixedaddressResponse) GetLogicFilterRules() []Ipv6fixedaddressLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetIpv6fixedaddressResponse) GetLogicFilterRulesOk() (*[]Ipv6fixedaddressLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetIpv6fixedaddressResponse) SetLogicFilterRules(v []Ipv6fixedaddressLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetIpv6fixedaddressResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetIpv6fixedaddressResponse) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetIpv6fixedaddressResponse) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetIpv6fixedaddressResponse) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetIpv6fixedaddressResponse) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMatchClient

`func (o *GetIpv6fixedaddressResponse) GetMatchClient() string`

GetMatchClient returns the MatchClient field if non-nil, zero value otherwise.

### GetMatchClientOk

`func (o *GetIpv6fixedaddressResponse) GetMatchClientOk() (*string, bool)`

GetMatchClientOk returns a tuple with the MatchClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClient

`func (o *GetIpv6fixedaddressResponse) SetMatchClient(v string)`

SetMatchClient sets MatchClient field to given value.

### HasMatchClient

`func (o *GetIpv6fixedaddressResponse) HasMatchClient() bool`

HasMatchClient returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetIpv6fixedaddressResponse) GetMsAdUserData() Ipv6fixedaddressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetIpv6fixedaddressResponse) GetMsAdUserDataOk() (*Ipv6fixedaddressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetIpv6fixedaddressResponse) SetMsAdUserData(v Ipv6fixedaddressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetIpv6fixedaddressResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6fixedaddressResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6fixedaddressResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6fixedaddressResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6fixedaddressResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetIpv6fixedaddressResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetIpv6fixedaddressResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetIpv6fixedaddressResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetIpv6fixedaddressResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetIpv6fixedaddressResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetIpv6fixedaddressResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetIpv6fixedaddressResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetIpv6fixedaddressResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOptions

`func (o *GetIpv6fixedaddressResponse) GetOptions() []Ipv6fixedaddressOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetIpv6fixedaddressResponse) GetOptionsOk() (*[]Ipv6fixedaddressOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetIpv6fixedaddressResponse) SetOptions(v []Ipv6fixedaddressOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetIpv6fixedaddressResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *GetIpv6fixedaddressResponse) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *GetIpv6fixedaddressResponse) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *GetIpv6fixedaddressResponse) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *GetIpv6fixedaddressResponse) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetReservedInterface

`func (o *GetIpv6fixedaddressResponse) GetReservedInterface() string`

GetReservedInterface returns the ReservedInterface field if non-nil, zero value otherwise.

### GetReservedInterfaceOk

`func (o *GetIpv6fixedaddressResponse) GetReservedInterfaceOk() (*string, bool)`

GetReservedInterfaceOk returns a tuple with the ReservedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedInterface

`func (o *GetIpv6fixedaddressResponse) SetReservedInterface(v string)`

SetReservedInterface sets ReservedInterface field to given value.

### HasReservedInterface

`func (o *GetIpv6fixedaddressResponse) HasReservedInterface() bool`

HasReservedInterface returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetIpv6fixedaddressResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetIpv6fixedaddressResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetIpv6fixedaddressResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetIpv6fixedaddressResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSnmp3Credential

`func (o *GetIpv6fixedaddressResponse) GetSnmp3Credential() Ipv6fixedaddressSnmp3Credential`

GetSnmp3Credential returns the Snmp3Credential field if non-nil, zero value otherwise.

### GetSnmp3CredentialOk

`func (o *GetIpv6fixedaddressResponse) GetSnmp3CredentialOk() (*Ipv6fixedaddressSnmp3Credential, bool)`

GetSnmp3CredentialOk returns a tuple with the Snmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp3Credential

`func (o *GetIpv6fixedaddressResponse) SetSnmp3Credential(v Ipv6fixedaddressSnmp3Credential)`

SetSnmp3Credential sets Snmp3Credential field to given value.

### HasSnmp3Credential

`func (o *GetIpv6fixedaddressResponse) HasSnmp3Credential() bool`

HasSnmp3Credential returns a boolean if a field has been set.

### GetSnmpCredential

`func (o *GetIpv6fixedaddressResponse) GetSnmpCredential() Ipv6fixedaddressSnmpCredential`

GetSnmpCredential returns the SnmpCredential field if non-nil, zero value otherwise.

### GetSnmpCredentialOk

`func (o *GetIpv6fixedaddressResponse) GetSnmpCredentialOk() (*Ipv6fixedaddressSnmpCredential, bool)`

GetSnmpCredentialOk returns a tuple with the SnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredential

`func (o *GetIpv6fixedaddressResponse) SetSnmpCredential(v Ipv6fixedaddressSnmpCredential)`

SetSnmpCredential sets SnmpCredential field to given value.

### HasSnmpCredential

`func (o *GetIpv6fixedaddressResponse) HasSnmpCredential() bool`

HasSnmpCredential returns a boolean if a field has been set.

### GetTemplate

`func (o *GetIpv6fixedaddressResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetIpv6fixedaddressResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetIpv6fixedaddressResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetIpv6fixedaddressResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUseCliCredentials

`func (o *GetIpv6fixedaddressResponse) GetUseCliCredentials() bool`

GetUseCliCredentials returns the UseCliCredentials field if non-nil, zero value otherwise.

### GetUseCliCredentialsOk

`func (o *GetIpv6fixedaddressResponse) GetUseCliCredentialsOk() (*bool, bool)`

GetUseCliCredentialsOk returns a tuple with the UseCliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCliCredentials

`func (o *GetIpv6fixedaddressResponse) SetUseCliCredentials(v bool)`

SetUseCliCredentials sets UseCliCredentials field to given value.

### HasUseCliCredentials

`func (o *GetIpv6fixedaddressResponse) HasUseCliCredentials() bool`

HasUseCliCredentials returns a boolean if a field has been set.

### GetUseDomainName

`func (o *GetIpv6fixedaddressResponse) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *GetIpv6fixedaddressResponse) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *GetIpv6fixedaddressResponse) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *GetIpv6fixedaddressResponse) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *GetIpv6fixedaddressResponse) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *GetIpv6fixedaddressResponse) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *GetIpv6fixedaddressResponse) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *GetIpv6fixedaddressResponse) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetIpv6fixedaddressResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetIpv6fixedaddressResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetIpv6fixedaddressResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetIpv6fixedaddressResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetIpv6fixedaddressResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetIpv6fixedaddressResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetIpv6fixedaddressResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetIpv6fixedaddressResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *GetIpv6fixedaddressResponse) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *GetIpv6fixedaddressResponse) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *GetIpv6fixedaddressResponse) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *GetIpv6fixedaddressResponse) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseSnmp3Credential

`func (o *GetIpv6fixedaddressResponse) GetUseSnmp3Credential() bool`

GetUseSnmp3Credential returns the UseSnmp3Credential field if non-nil, zero value otherwise.

### GetUseSnmp3CredentialOk

`func (o *GetIpv6fixedaddressResponse) GetUseSnmp3CredentialOk() (*bool, bool)`

GetUseSnmp3CredentialOk returns a tuple with the UseSnmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmp3Credential

`func (o *GetIpv6fixedaddressResponse) SetUseSnmp3Credential(v bool)`

SetUseSnmp3Credential sets UseSnmp3Credential field to given value.

### HasUseSnmp3Credential

`func (o *GetIpv6fixedaddressResponse) HasUseSnmp3Credential() bool`

HasUseSnmp3Credential returns a boolean if a field has been set.

### GetUseSnmpCredential

`func (o *GetIpv6fixedaddressResponse) GetUseSnmpCredential() bool`

GetUseSnmpCredential returns the UseSnmpCredential field if non-nil, zero value otherwise.

### GetUseSnmpCredentialOk

`func (o *GetIpv6fixedaddressResponse) GetUseSnmpCredentialOk() (*bool, bool)`

GetUseSnmpCredentialOk returns a tuple with the UseSnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpCredential

`func (o *GetIpv6fixedaddressResponse) SetUseSnmpCredential(v bool)`

SetUseSnmpCredential sets UseSnmpCredential field to given value.

### HasUseSnmpCredential

`func (o *GetIpv6fixedaddressResponse) HasUseSnmpCredential() bool`

HasUseSnmpCredential returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *GetIpv6fixedaddressResponse) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *GetIpv6fixedaddressResponse) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *GetIpv6fixedaddressResponse) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *GetIpv6fixedaddressResponse) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *GetIpv6fixedaddressResponse) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *GetIpv6fixedaddressResponse) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *GetIpv6fixedaddressResponse) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *GetIpv6fixedaddressResponse) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6fixedaddressResponse) GetResult() Ipv6fixedaddress`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6fixedaddressResponse) GetResultOk() (*Ipv6fixedaddress, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6fixedaddressResponse) SetResult(v Ipv6fixedaddress)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6fixedaddressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


