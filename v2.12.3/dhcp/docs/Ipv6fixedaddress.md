# Ipv6fixedaddress

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
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the DHCP IPv6 fixed address. | [optional] 
**Ipv6prefix** | Pointer to **string** | The IPv6 Address prefix of the DHCP IPv6 fixed address. | [optional] 
**Ipv6prefixBits** | Pointer to **int64** | Prefix bits of the DHCP IPv6 fixed address. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6fixedaddressLogicFilterRules**](Ipv6fixedaddressLogicFilterRules.md) | This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
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

## Methods

### NewIpv6fixedaddress

`func NewIpv6fixedaddress() *Ipv6fixedaddress`

NewIpv6fixedaddress instantiates a new Ipv6fixedaddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6fixedaddressWithDefaults

`func NewIpv6fixedaddressWithDefaults() *Ipv6fixedaddress`

NewIpv6fixedaddressWithDefaults instantiates a new Ipv6fixedaddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6fixedaddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6fixedaddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6fixedaddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6fixedaddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddressType

`func (o *Ipv6fixedaddress) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *Ipv6fixedaddress) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *Ipv6fixedaddress) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *Ipv6fixedaddress) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAllowTelnet

`func (o *Ipv6fixedaddress) GetAllowTelnet() bool`

GetAllowTelnet returns the AllowTelnet field if non-nil, zero value otherwise.

### GetAllowTelnetOk

`func (o *Ipv6fixedaddress) GetAllowTelnetOk() (*bool, bool)`

GetAllowTelnetOk returns a tuple with the AllowTelnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTelnet

`func (o *Ipv6fixedaddress) SetAllowTelnet(v bool)`

SetAllowTelnet sets AllowTelnet field to given value.

### HasAllowTelnet

`func (o *Ipv6fixedaddress) HasAllowTelnet() bool`

HasAllowTelnet returns a boolean if a field has been set.

### GetCliCredentials

`func (o *Ipv6fixedaddress) GetCliCredentials() []Ipv6fixedaddressCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *Ipv6fixedaddress) GetCliCredentialsOk() (*[]Ipv6fixedaddressCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *Ipv6fixedaddress) SetCliCredentials(v []Ipv6fixedaddressCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *Ipv6fixedaddress) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Ipv6fixedaddress) GetCloudInfo() Ipv6fixedaddressCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Ipv6fixedaddress) GetCloudInfoOk() (*Ipv6fixedaddressCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Ipv6fixedaddress) SetCloudInfo(v Ipv6fixedaddressCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Ipv6fixedaddress) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6fixedaddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6fixedaddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6fixedaddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6fixedaddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeviceDescription

`func (o *Ipv6fixedaddress) GetDeviceDescription() string`

GetDeviceDescription returns the DeviceDescription field if non-nil, zero value otherwise.

### GetDeviceDescriptionOk

`func (o *Ipv6fixedaddress) GetDeviceDescriptionOk() (*string, bool)`

GetDeviceDescriptionOk returns a tuple with the DeviceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDescription

`func (o *Ipv6fixedaddress) SetDeviceDescription(v string)`

SetDeviceDescription sets DeviceDescription field to given value.

### HasDeviceDescription

`func (o *Ipv6fixedaddress) HasDeviceDescription() bool`

HasDeviceDescription returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *Ipv6fixedaddress) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *Ipv6fixedaddress) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *Ipv6fixedaddress) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *Ipv6fixedaddress) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceType

`func (o *Ipv6fixedaddress) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Ipv6fixedaddress) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Ipv6fixedaddress) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Ipv6fixedaddress) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *Ipv6fixedaddress) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *Ipv6fixedaddress) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *Ipv6fixedaddress) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *Ipv6fixedaddress) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDisable

`func (o *Ipv6fixedaddress) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Ipv6fixedaddress) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Ipv6fixedaddress) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Ipv6fixedaddress) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableDiscovery

`func (o *Ipv6fixedaddress) GetDisableDiscovery() bool`

GetDisableDiscovery returns the DisableDiscovery field if non-nil, zero value otherwise.

### GetDisableDiscoveryOk

`func (o *Ipv6fixedaddress) GetDisableDiscoveryOk() (*bool, bool)`

GetDisableDiscoveryOk returns a tuple with the DisableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDiscovery

`func (o *Ipv6fixedaddress) SetDisableDiscovery(v bool)`

SetDisableDiscovery sets DisableDiscovery field to given value.

### HasDisableDiscovery

`func (o *Ipv6fixedaddress) HasDisableDiscovery() bool`

HasDisableDiscovery returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Ipv6fixedaddress) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Ipv6fixedaddress) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Ipv6fixedaddress) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Ipv6fixedaddress) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *Ipv6fixedaddress) GetDiscoveredData() Ipv6fixedaddressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *Ipv6fixedaddress) GetDiscoveredDataOk() (*Ipv6fixedaddressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *Ipv6fixedaddress) SetDiscoveredData(v Ipv6fixedaddressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *Ipv6fixedaddress) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDomainName

`func (o *Ipv6fixedaddress) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Ipv6fixedaddress) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Ipv6fixedaddress) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Ipv6fixedaddress) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *Ipv6fixedaddress) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *Ipv6fixedaddress) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *Ipv6fixedaddress) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *Ipv6fixedaddress) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetDuid

`func (o *Ipv6fixedaddress) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *Ipv6fixedaddress) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *Ipv6fixedaddress) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *Ipv6fixedaddress) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Ipv6fixedaddress) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Ipv6fixedaddress) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Ipv6fixedaddress) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Ipv6fixedaddress) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetExtattrs

`func (o *Ipv6fixedaddress) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Ipv6fixedaddress) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Ipv6fixedaddress) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Ipv6fixedaddress) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIpv6addr

`func (o *Ipv6fixedaddress) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *Ipv6fixedaddress) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *Ipv6fixedaddress) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *Ipv6fixedaddress) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetIpv6prefix

`func (o *Ipv6fixedaddress) GetIpv6prefix() string`

GetIpv6prefix returns the Ipv6prefix field if non-nil, zero value otherwise.

### GetIpv6prefixOk

`func (o *Ipv6fixedaddress) GetIpv6prefixOk() (*string, bool)`

GetIpv6prefixOk returns a tuple with the Ipv6prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefix

`func (o *Ipv6fixedaddress) SetIpv6prefix(v string)`

SetIpv6prefix sets Ipv6prefix field to given value.

### HasIpv6prefix

`func (o *Ipv6fixedaddress) HasIpv6prefix() bool`

HasIpv6prefix returns a boolean if a field has been set.

### GetIpv6prefixBits

`func (o *Ipv6fixedaddress) GetIpv6prefixBits() int64`

GetIpv6prefixBits returns the Ipv6prefixBits field if non-nil, zero value otherwise.

### GetIpv6prefixBitsOk

`func (o *Ipv6fixedaddress) GetIpv6prefixBitsOk() (*int64, bool)`

GetIpv6prefixBitsOk returns a tuple with the Ipv6prefixBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefixBits

`func (o *Ipv6fixedaddress) SetIpv6prefixBits(v int64)`

SetIpv6prefixBits sets Ipv6prefixBits field to given value.

### HasIpv6prefixBits

`func (o *Ipv6fixedaddress) HasIpv6prefixBits() bool`

HasIpv6prefixBits returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6fixedaddress) GetLogicFilterRules() []Ipv6fixedaddressLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6fixedaddress) GetLogicFilterRulesOk() (*[]Ipv6fixedaddressLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6fixedaddress) SetLogicFilterRules(v []Ipv6fixedaddressLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6fixedaddress) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Ipv6fixedaddress) GetMsAdUserData() Ipv6fixedaddressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Ipv6fixedaddress) GetMsAdUserDataOk() (*Ipv6fixedaddressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Ipv6fixedaddress) SetMsAdUserData(v Ipv6fixedaddressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Ipv6fixedaddress) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *Ipv6fixedaddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6fixedaddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6fixedaddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6fixedaddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Ipv6fixedaddress) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ipv6fixedaddress) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ipv6fixedaddress) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ipv6fixedaddress) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Ipv6fixedaddress) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Ipv6fixedaddress) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Ipv6fixedaddress) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Ipv6fixedaddress) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOptions

`func (o *Ipv6fixedaddress) GetOptions() []Ipv6fixedaddressOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Ipv6fixedaddress) GetOptionsOk() (*[]Ipv6fixedaddressOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Ipv6fixedaddress) SetOptions(v []Ipv6fixedaddressOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Ipv6fixedaddress) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *Ipv6fixedaddress) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *Ipv6fixedaddress) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *Ipv6fixedaddress) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *Ipv6fixedaddress) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetReservedInterface

`func (o *Ipv6fixedaddress) GetReservedInterface() string`

GetReservedInterface returns the ReservedInterface field if non-nil, zero value otherwise.

### GetReservedInterfaceOk

`func (o *Ipv6fixedaddress) GetReservedInterfaceOk() (*string, bool)`

GetReservedInterfaceOk returns a tuple with the ReservedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedInterface

`func (o *Ipv6fixedaddress) SetReservedInterface(v string)`

SetReservedInterface sets ReservedInterface field to given value.

### HasReservedInterface

`func (o *Ipv6fixedaddress) HasReservedInterface() bool`

HasReservedInterface returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Ipv6fixedaddress) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Ipv6fixedaddress) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Ipv6fixedaddress) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Ipv6fixedaddress) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSnmp3Credential

`func (o *Ipv6fixedaddress) GetSnmp3Credential() Ipv6fixedaddressSnmp3Credential`

GetSnmp3Credential returns the Snmp3Credential field if non-nil, zero value otherwise.

### GetSnmp3CredentialOk

`func (o *Ipv6fixedaddress) GetSnmp3CredentialOk() (*Ipv6fixedaddressSnmp3Credential, bool)`

GetSnmp3CredentialOk returns a tuple with the Snmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp3Credential

`func (o *Ipv6fixedaddress) SetSnmp3Credential(v Ipv6fixedaddressSnmp3Credential)`

SetSnmp3Credential sets Snmp3Credential field to given value.

### HasSnmp3Credential

`func (o *Ipv6fixedaddress) HasSnmp3Credential() bool`

HasSnmp3Credential returns a boolean if a field has been set.

### GetSnmpCredential

`func (o *Ipv6fixedaddress) GetSnmpCredential() Ipv6fixedaddressSnmpCredential`

GetSnmpCredential returns the SnmpCredential field if non-nil, zero value otherwise.

### GetSnmpCredentialOk

`func (o *Ipv6fixedaddress) GetSnmpCredentialOk() (*Ipv6fixedaddressSnmpCredential, bool)`

GetSnmpCredentialOk returns a tuple with the SnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredential

`func (o *Ipv6fixedaddress) SetSnmpCredential(v Ipv6fixedaddressSnmpCredential)`

SetSnmpCredential sets SnmpCredential field to given value.

### HasSnmpCredential

`func (o *Ipv6fixedaddress) HasSnmpCredential() bool`

HasSnmpCredential returns a boolean if a field has been set.

### GetTemplate

`func (o *Ipv6fixedaddress) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Ipv6fixedaddress) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Ipv6fixedaddress) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Ipv6fixedaddress) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUseCliCredentials

`func (o *Ipv6fixedaddress) GetUseCliCredentials() bool`

GetUseCliCredentials returns the UseCliCredentials field if non-nil, zero value otherwise.

### GetUseCliCredentialsOk

`func (o *Ipv6fixedaddress) GetUseCliCredentialsOk() (*bool, bool)`

GetUseCliCredentialsOk returns a tuple with the UseCliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCliCredentials

`func (o *Ipv6fixedaddress) SetUseCliCredentials(v bool)`

SetUseCliCredentials sets UseCliCredentials field to given value.

### HasUseCliCredentials

`func (o *Ipv6fixedaddress) HasUseCliCredentials() bool`

HasUseCliCredentials returns a boolean if a field has been set.

### GetUseDomainName

`func (o *Ipv6fixedaddress) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *Ipv6fixedaddress) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *Ipv6fixedaddress) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *Ipv6fixedaddress) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *Ipv6fixedaddress) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *Ipv6fixedaddress) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *Ipv6fixedaddress) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *Ipv6fixedaddress) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6fixedaddress) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6fixedaddress) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6fixedaddress) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6fixedaddress) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *Ipv6fixedaddress) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Ipv6fixedaddress) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Ipv6fixedaddress) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Ipv6fixedaddress) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *Ipv6fixedaddress) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *Ipv6fixedaddress) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *Ipv6fixedaddress) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *Ipv6fixedaddress) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseSnmp3Credential

`func (o *Ipv6fixedaddress) GetUseSnmp3Credential() bool`

GetUseSnmp3Credential returns the UseSnmp3Credential field if non-nil, zero value otherwise.

### GetUseSnmp3CredentialOk

`func (o *Ipv6fixedaddress) GetUseSnmp3CredentialOk() (*bool, bool)`

GetUseSnmp3CredentialOk returns a tuple with the UseSnmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmp3Credential

`func (o *Ipv6fixedaddress) SetUseSnmp3Credential(v bool)`

SetUseSnmp3Credential sets UseSnmp3Credential field to given value.

### HasUseSnmp3Credential

`func (o *Ipv6fixedaddress) HasUseSnmp3Credential() bool`

HasUseSnmp3Credential returns a boolean if a field has been set.

### GetUseSnmpCredential

`func (o *Ipv6fixedaddress) GetUseSnmpCredential() bool`

GetUseSnmpCredential returns the UseSnmpCredential field if non-nil, zero value otherwise.

### GetUseSnmpCredentialOk

`func (o *Ipv6fixedaddress) GetUseSnmpCredentialOk() (*bool, bool)`

GetUseSnmpCredentialOk returns a tuple with the UseSnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpCredential

`func (o *Ipv6fixedaddress) SetUseSnmpCredential(v bool)`

SetUseSnmpCredential sets UseSnmpCredential field to given value.

### HasUseSnmpCredential

`func (o *Ipv6fixedaddress) HasUseSnmpCredential() bool`

HasUseSnmpCredential returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *Ipv6fixedaddress) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *Ipv6fixedaddress) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *Ipv6fixedaddress) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *Ipv6fixedaddress) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *Ipv6fixedaddress) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *Ipv6fixedaddress) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *Ipv6fixedaddress) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *Ipv6fixedaddress) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


