# GetRoaminghostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AddressType** | Pointer to **string** | The address type for this roaming host. | [optional] 
**Bootfile** | Pointer to **string** | The bootfile name for the roaming host. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The boot server address for the roaming host. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**ClientIdentifierPrependZero** | Pointer to **bool** | This field controls whether there is a prepend for the dhcp-client-identifier of a roaming host. | [optional] 
**Comment** | Pointer to **string** | Comment for the roaming host; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The DDNS domain name for this roaming host. | [optional] 
**DdnsHostname** | Pointer to **string** | The DDNS host name for this roaming host. | [optional] 
**DenyBootp** | Pointer to **bool** | If set to true, BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**DhcpClientIdentifier** | Pointer to **string** | The DHCP client ID for the roaming host. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a roaming host is disabled or not. When this is set to False, the roaming host is enabled. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of the roaming host object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForceRoamingHostname** | Pointer to **bool** | Set this to True to use the roaming host name as its ddns_hostname. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all the DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**Ipv6ClientHostname** | Pointer to **string** | The client hostname of a DHCP roaming host object. This field specifies the host name that the DHCP client sends to the Infoblox appliance using DHCP option 12. | [optional] [readonly] 
**Ipv6DdnsDomainname** | Pointer to **string** | The IPv6 DDNS domain name for this roaming host. | [optional] 
**Ipv6DdnsHostname** | Pointer to **string** | The IPv6 DDNS host name for this roaming host. | [optional] 
**Ipv6DomainName** | Pointer to **string** | The IPv6 domain name for this roaming host. | [optional] 
**Ipv6DomainNameServers** | Pointer to **[]string** | The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages. | [optional] 
**Ipv6Duid** | Pointer to **string** | The DUID value for this roaming host. | [optional] 
**Ipv6EnableDdns** | Pointer to **bool** | Set this to True to enable IPv6 DDNS. | [optional] 
**Ipv6ForceRoamingHostname** | Pointer to **bool** | Set this to True to use the roaming host name as its ddns_hostname. | [optional] 
**Ipv6MacAddress** | Pointer to **string** | The MAC address for this roaming host. | [optional] 
**Ipv6MatchOption** | Pointer to **string** | The identification method for an IPv6 or mixed IPv4/IPv6 roaming host. The supported values for this field are \&quot;DUID\&quot; or \&quot;V6_MAC_ADDRESS\&quot;, which specify what option should be used to identify the specific DHCPv6 client. | [optional] 
**Ipv6Options** | Pointer to [**[]RoaminghostIpv6Options**](RoaminghostIpv6Options.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**Ipv6Template** | Pointer to **string** | If set on creation, the roaming host will be created according to the values specified in the named IPv6 roaming host template. | [optional] 
**Mac** | Pointer to **string** | The MAC address for this roaming host. | [optional] 
**MatchClient** | Pointer to **string** | The match-client value for this roaming host. Valid values are: \&quot;MAC_ADDRESS\&quot;: The fixed IP address is leased to the matching MAC address. \&quot;CLIENT_ID\&quot;: The fixed IP address is leased to the matching DHCP client identifier. | [optional] 
**Name** | Pointer to **string** | The name of this roaming host. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this roaming host resides. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot. | [optional] 
**Options** | Pointer to [**[]RoaminghostOptions**](RoaminghostOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PreferredLifetime** | Pointer to **int64** | The preferred lifetime value for this roaming host object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value for this roaming host object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**Template** | Pointer to **string** | If set on creation, the roaming host will be created according to the values specified in the named template. | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseIpv6DdnsDomainname** | Pointer to **bool** | Use flag for: ipv6_ddns_domainname | [optional] 
**UseIpv6DomainName** | Pointer to **bool** | Use flag for: ipv6_domain_name | [optional] 
**UseIpv6DomainNameServers** | Pointer to **bool** | Use flag for: ipv6_domain_name_servers | [optional] 
**UseIpv6EnableDdns** | Pointer to **bool** | Use flag for: ipv6_enable_ddns | [optional] 
**UseIpv6Options** | Pointer to **bool** | Use flag for: ipv6_options | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**ValidLifetime** | Pointer to **int64** | The valid lifetime value for this roaming host object. | [optional] 
**Result** | Pointer to [**Roaminghost**](Roaminghost.md) |  | [optional] 

## Methods

### NewGetRoaminghostResponse

`func NewGetRoaminghostResponse() *GetRoaminghostResponse`

NewGetRoaminghostResponse instantiates a new GetRoaminghostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoaminghostResponseWithDefaults

`func NewGetRoaminghostResponseWithDefaults() *GetRoaminghostResponse`

NewGetRoaminghostResponseWithDefaults instantiates a new GetRoaminghostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRoaminghostResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRoaminghostResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRoaminghostResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRoaminghostResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddressType

`func (o *GetRoaminghostResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *GetRoaminghostResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *GetRoaminghostResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *GetRoaminghostResponse) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetBootfile

`func (o *GetRoaminghostResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetRoaminghostResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetRoaminghostResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetRoaminghostResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetRoaminghostResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetRoaminghostResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetRoaminghostResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetRoaminghostResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetClientIdentifierPrependZero

`func (o *GetRoaminghostResponse) GetClientIdentifierPrependZero() bool`

GetClientIdentifierPrependZero returns the ClientIdentifierPrependZero field if non-nil, zero value otherwise.

### GetClientIdentifierPrependZeroOk

`func (o *GetRoaminghostResponse) GetClientIdentifierPrependZeroOk() (*bool, bool)`

GetClientIdentifierPrependZeroOk returns a tuple with the ClientIdentifierPrependZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifierPrependZero

`func (o *GetRoaminghostResponse) SetClientIdentifierPrependZero(v bool)`

SetClientIdentifierPrependZero sets ClientIdentifierPrependZero field to given value.

### HasClientIdentifierPrependZero

`func (o *GetRoaminghostResponse) HasClientIdentifierPrependZero() bool`

HasClientIdentifierPrependZero returns a boolean if a field has been set.

### GetComment

`func (o *GetRoaminghostResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRoaminghostResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRoaminghostResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRoaminghostResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetRoaminghostResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetRoaminghostResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetRoaminghostResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetRoaminghostResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsHostname

`func (o *GetRoaminghostResponse) GetDdnsHostname() string`

GetDdnsHostname returns the DdnsHostname field if non-nil, zero value otherwise.

### GetDdnsHostnameOk

`func (o *GetRoaminghostResponse) GetDdnsHostnameOk() (*string, bool)`

GetDdnsHostnameOk returns a tuple with the DdnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostname

`func (o *GetRoaminghostResponse) SetDdnsHostname(v string)`

SetDdnsHostname sets DdnsHostname field to given value.

### HasDdnsHostname

`func (o *GetRoaminghostResponse) HasDdnsHostname() bool`

HasDdnsHostname returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetRoaminghostResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetRoaminghostResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetRoaminghostResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetRoaminghostResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpClientIdentifier

`func (o *GetRoaminghostResponse) GetDhcpClientIdentifier() string`

GetDhcpClientIdentifier returns the DhcpClientIdentifier field if non-nil, zero value otherwise.

### GetDhcpClientIdentifierOk

`func (o *GetRoaminghostResponse) GetDhcpClientIdentifierOk() (*string, bool)`

GetDhcpClientIdentifierOk returns a tuple with the DhcpClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpClientIdentifier

`func (o *GetRoaminghostResponse) SetDhcpClientIdentifier(v string)`

SetDhcpClientIdentifier sets DhcpClientIdentifier field to given value.

### HasDhcpClientIdentifier

`func (o *GetRoaminghostResponse) HasDhcpClientIdentifier() bool`

HasDhcpClientIdentifier returns a boolean if a field has been set.

### GetDisable

`func (o *GetRoaminghostResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRoaminghostResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRoaminghostResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRoaminghostResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetRoaminghostResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetRoaminghostResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetRoaminghostResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetRoaminghostResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetRoaminghostResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetRoaminghostResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetRoaminghostResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetRoaminghostResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetRoaminghostResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetRoaminghostResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetRoaminghostResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetRoaminghostResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetRoaminghostResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetRoaminghostResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetRoaminghostResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetRoaminghostResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRoaminghostResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRoaminghostResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRoaminghostResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRoaminghostResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForceRoamingHostname

`func (o *GetRoaminghostResponse) GetForceRoamingHostname() bool`

GetForceRoamingHostname returns the ForceRoamingHostname field if non-nil, zero value otherwise.

### GetForceRoamingHostnameOk

`func (o *GetRoaminghostResponse) GetForceRoamingHostnameOk() (*bool, bool)`

GetForceRoamingHostnameOk returns a tuple with the ForceRoamingHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRoamingHostname

`func (o *GetRoaminghostResponse) SetForceRoamingHostname(v bool)`

SetForceRoamingHostname sets ForceRoamingHostname field to given value.

### HasForceRoamingHostname

`func (o *GetRoaminghostResponse) HasForceRoamingHostname() bool`

HasForceRoamingHostname returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GetRoaminghostResponse) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GetRoaminghostResponse) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GetRoaminghostResponse) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GetRoaminghostResponse) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIpv6ClientHostname

`func (o *GetRoaminghostResponse) GetIpv6ClientHostname() string`

GetIpv6ClientHostname returns the Ipv6ClientHostname field if non-nil, zero value otherwise.

### GetIpv6ClientHostnameOk

`func (o *GetRoaminghostResponse) GetIpv6ClientHostnameOk() (*string, bool)`

GetIpv6ClientHostnameOk returns a tuple with the Ipv6ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6ClientHostname

`func (o *GetRoaminghostResponse) SetIpv6ClientHostname(v string)`

SetIpv6ClientHostname sets Ipv6ClientHostname field to given value.

### HasIpv6ClientHostname

`func (o *GetRoaminghostResponse) HasIpv6ClientHostname() bool`

HasIpv6ClientHostname returns a boolean if a field has been set.

### GetIpv6DdnsDomainname

`func (o *GetRoaminghostResponse) GetIpv6DdnsDomainname() string`

GetIpv6DdnsDomainname returns the Ipv6DdnsDomainname field if non-nil, zero value otherwise.

### GetIpv6DdnsDomainnameOk

`func (o *GetRoaminghostResponse) GetIpv6DdnsDomainnameOk() (*string, bool)`

GetIpv6DdnsDomainnameOk returns a tuple with the Ipv6DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsDomainname

`func (o *GetRoaminghostResponse) SetIpv6DdnsDomainname(v string)`

SetIpv6DdnsDomainname sets Ipv6DdnsDomainname field to given value.

### HasIpv6DdnsDomainname

`func (o *GetRoaminghostResponse) HasIpv6DdnsDomainname() bool`

HasIpv6DdnsDomainname returns a boolean if a field has been set.

### GetIpv6DdnsHostname

`func (o *GetRoaminghostResponse) GetIpv6DdnsHostname() string`

GetIpv6DdnsHostname returns the Ipv6DdnsHostname field if non-nil, zero value otherwise.

### GetIpv6DdnsHostnameOk

`func (o *GetRoaminghostResponse) GetIpv6DdnsHostnameOk() (*string, bool)`

GetIpv6DdnsHostnameOk returns a tuple with the Ipv6DdnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsHostname

`func (o *GetRoaminghostResponse) SetIpv6DdnsHostname(v string)`

SetIpv6DdnsHostname sets Ipv6DdnsHostname field to given value.

### HasIpv6DdnsHostname

`func (o *GetRoaminghostResponse) HasIpv6DdnsHostname() bool`

HasIpv6DdnsHostname returns a boolean if a field has been set.

### GetIpv6DomainName

`func (o *GetRoaminghostResponse) GetIpv6DomainName() string`

GetIpv6DomainName returns the Ipv6DomainName field if non-nil, zero value otherwise.

### GetIpv6DomainNameOk

`func (o *GetRoaminghostResponse) GetIpv6DomainNameOk() (*string, bool)`

GetIpv6DomainNameOk returns a tuple with the Ipv6DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DomainName

`func (o *GetRoaminghostResponse) SetIpv6DomainName(v string)`

SetIpv6DomainName sets Ipv6DomainName field to given value.

### HasIpv6DomainName

`func (o *GetRoaminghostResponse) HasIpv6DomainName() bool`

HasIpv6DomainName returns a boolean if a field has been set.

### GetIpv6DomainNameServers

`func (o *GetRoaminghostResponse) GetIpv6DomainNameServers() []string`

GetIpv6DomainNameServers returns the Ipv6DomainNameServers field if non-nil, zero value otherwise.

### GetIpv6DomainNameServersOk

`func (o *GetRoaminghostResponse) GetIpv6DomainNameServersOk() (*[]string, bool)`

GetIpv6DomainNameServersOk returns a tuple with the Ipv6DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DomainNameServers

`func (o *GetRoaminghostResponse) SetIpv6DomainNameServers(v []string)`

SetIpv6DomainNameServers sets Ipv6DomainNameServers field to given value.

### HasIpv6DomainNameServers

`func (o *GetRoaminghostResponse) HasIpv6DomainNameServers() bool`

HasIpv6DomainNameServers returns a boolean if a field has been set.

### GetIpv6Duid

`func (o *GetRoaminghostResponse) GetIpv6Duid() string`

GetIpv6Duid returns the Ipv6Duid field if non-nil, zero value otherwise.

### GetIpv6DuidOk

`func (o *GetRoaminghostResponse) GetIpv6DuidOk() (*string, bool)`

GetIpv6DuidOk returns a tuple with the Ipv6Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Duid

`func (o *GetRoaminghostResponse) SetIpv6Duid(v string)`

SetIpv6Duid sets Ipv6Duid field to given value.

### HasIpv6Duid

`func (o *GetRoaminghostResponse) HasIpv6Duid() bool`

HasIpv6Duid returns a boolean if a field has been set.

### GetIpv6EnableDdns

`func (o *GetRoaminghostResponse) GetIpv6EnableDdns() bool`

GetIpv6EnableDdns returns the Ipv6EnableDdns field if non-nil, zero value otherwise.

### GetIpv6EnableDdnsOk

`func (o *GetRoaminghostResponse) GetIpv6EnableDdnsOk() (*bool, bool)`

GetIpv6EnableDdnsOk returns a tuple with the Ipv6EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableDdns

`func (o *GetRoaminghostResponse) SetIpv6EnableDdns(v bool)`

SetIpv6EnableDdns sets Ipv6EnableDdns field to given value.

### HasIpv6EnableDdns

`func (o *GetRoaminghostResponse) HasIpv6EnableDdns() bool`

HasIpv6EnableDdns returns a boolean if a field has been set.

### GetIpv6ForceRoamingHostname

`func (o *GetRoaminghostResponse) GetIpv6ForceRoamingHostname() bool`

GetIpv6ForceRoamingHostname returns the Ipv6ForceRoamingHostname field if non-nil, zero value otherwise.

### GetIpv6ForceRoamingHostnameOk

`func (o *GetRoaminghostResponse) GetIpv6ForceRoamingHostnameOk() (*bool, bool)`

GetIpv6ForceRoamingHostnameOk returns a tuple with the Ipv6ForceRoamingHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6ForceRoamingHostname

`func (o *GetRoaminghostResponse) SetIpv6ForceRoamingHostname(v bool)`

SetIpv6ForceRoamingHostname sets Ipv6ForceRoamingHostname field to given value.

### HasIpv6ForceRoamingHostname

`func (o *GetRoaminghostResponse) HasIpv6ForceRoamingHostname() bool`

HasIpv6ForceRoamingHostname returns a boolean if a field has been set.

### GetIpv6MacAddress

`func (o *GetRoaminghostResponse) GetIpv6MacAddress() string`

GetIpv6MacAddress returns the Ipv6MacAddress field if non-nil, zero value otherwise.

### GetIpv6MacAddressOk

`func (o *GetRoaminghostResponse) GetIpv6MacAddressOk() (*string, bool)`

GetIpv6MacAddressOk returns a tuple with the Ipv6MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MacAddress

`func (o *GetRoaminghostResponse) SetIpv6MacAddress(v string)`

SetIpv6MacAddress sets Ipv6MacAddress field to given value.

### HasIpv6MacAddress

`func (o *GetRoaminghostResponse) HasIpv6MacAddress() bool`

HasIpv6MacAddress returns a boolean if a field has been set.

### GetIpv6MatchOption

`func (o *GetRoaminghostResponse) GetIpv6MatchOption() string`

GetIpv6MatchOption returns the Ipv6MatchOption field if non-nil, zero value otherwise.

### GetIpv6MatchOptionOk

`func (o *GetRoaminghostResponse) GetIpv6MatchOptionOk() (*string, bool)`

GetIpv6MatchOptionOk returns a tuple with the Ipv6MatchOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MatchOption

`func (o *GetRoaminghostResponse) SetIpv6MatchOption(v string)`

SetIpv6MatchOption sets Ipv6MatchOption field to given value.

### HasIpv6MatchOption

`func (o *GetRoaminghostResponse) HasIpv6MatchOption() bool`

HasIpv6MatchOption returns a boolean if a field has been set.

### GetIpv6Options

`func (o *GetRoaminghostResponse) GetIpv6Options() []RoaminghostIpv6Options`

GetIpv6Options returns the Ipv6Options field if non-nil, zero value otherwise.

### GetIpv6OptionsOk

`func (o *GetRoaminghostResponse) GetIpv6OptionsOk() (*[]RoaminghostIpv6Options, bool)`

GetIpv6OptionsOk returns a tuple with the Ipv6Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Options

`func (o *GetRoaminghostResponse) SetIpv6Options(v []RoaminghostIpv6Options)`

SetIpv6Options sets Ipv6Options field to given value.

### HasIpv6Options

`func (o *GetRoaminghostResponse) HasIpv6Options() bool`

HasIpv6Options returns a boolean if a field has been set.

### GetIpv6Template

`func (o *GetRoaminghostResponse) GetIpv6Template() string`

GetIpv6Template returns the Ipv6Template field if non-nil, zero value otherwise.

### GetIpv6TemplateOk

`func (o *GetRoaminghostResponse) GetIpv6TemplateOk() (*string, bool)`

GetIpv6TemplateOk returns a tuple with the Ipv6Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Template

`func (o *GetRoaminghostResponse) SetIpv6Template(v string)`

SetIpv6Template sets Ipv6Template field to given value.

### HasIpv6Template

`func (o *GetRoaminghostResponse) HasIpv6Template() bool`

HasIpv6Template returns a boolean if a field has been set.

### GetMac

`func (o *GetRoaminghostResponse) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetRoaminghostResponse) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetRoaminghostResponse) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetRoaminghostResponse) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMatchClient

`func (o *GetRoaminghostResponse) GetMatchClient() string`

GetMatchClient returns the MatchClient field if non-nil, zero value otherwise.

### GetMatchClientOk

`func (o *GetRoaminghostResponse) GetMatchClientOk() (*string, bool)`

GetMatchClientOk returns a tuple with the MatchClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClient

`func (o *GetRoaminghostResponse) SetMatchClient(v string)`

SetMatchClient sets MatchClient field to given value.

### HasMatchClient

`func (o *GetRoaminghostResponse) HasMatchClient() bool`

HasMatchClient returns a boolean if a field has been set.

### GetName

`func (o *GetRoaminghostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRoaminghostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRoaminghostResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRoaminghostResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetRoaminghostResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetRoaminghostResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetRoaminghostResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetRoaminghostResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *GetRoaminghostResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetRoaminghostResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetRoaminghostResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetRoaminghostResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *GetRoaminghostResponse) GetOptions() []RoaminghostOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetRoaminghostResponse) GetOptionsOk() (*[]RoaminghostOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetRoaminghostResponse) SetOptions(v []RoaminghostOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetRoaminghostResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *GetRoaminghostResponse) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *GetRoaminghostResponse) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *GetRoaminghostResponse) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *GetRoaminghostResponse) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetRoaminghostResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetRoaminghostResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetRoaminghostResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetRoaminghostResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetTemplate

`func (o *GetRoaminghostResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetRoaminghostResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetRoaminghostResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetRoaminghostResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetRoaminghostResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetRoaminghostResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetRoaminghostResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetRoaminghostResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetRoaminghostResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetRoaminghostResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetRoaminghostResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetRoaminghostResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetRoaminghostResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetRoaminghostResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetRoaminghostResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetRoaminghostResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetRoaminghostResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetRoaminghostResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetRoaminghostResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetRoaminghostResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetRoaminghostResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetRoaminghostResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetRoaminghostResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetRoaminghostResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *GetRoaminghostResponse) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *GetRoaminghostResponse) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *GetRoaminghostResponse) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *GetRoaminghostResponse) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIpv6DdnsDomainname

`func (o *GetRoaminghostResponse) GetUseIpv6DdnsDomainname() bool`

GetUseIpv6DdnsDomainname returns the UseIpv6DdnsDomainname field if non-nil, zero value otherwise.

### GetUseIpv6DdnsDomainnameOk

`func (o *GetRoaminghostResponse) GetUseIpv6DdnsDomainnameOk() (*bool, bool)`

GetUseIpv6DdnsDomainnameOk returns a tuple with the UseIpv6DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DdnsDomainname

`func (o *GetRoaminghostResponse) SetUseIpv6DdnsDomainname(v bool)`

SetUseIpv6DdnsDomainname sets UseIpv6DdnsDomainname field to given value.

### HasUseIpv6DdnsDomainname

`func (o *GetRoaminghostResponse) HasUseIpv6DdnsDomainname() bool`

HasUseIpv6DdnsDomainname returns a boolean if a field has been set.

### GetUseIpv6DomainName

`func (o *GetRoaminghostResponse) GetUseIpv6DomainName() bool`

GetUseIpv6DomainName returns the UseIpv6DomainName field if non-nil, zero value otherwise.

### GetUseIpv6DomainNameOk

`func (o *GetRoaminghostResponse) GetUseIpv6DomainNameOk() (*bool, bool)`

GetUseIpv6DomainNameOk returns a tuple with the UseIpv6DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DomainName

`func (o *GetRoaminghostResponse) SetUseIpv6DomainName(v bool)`

SetUseIpv6DomainName sets UseIpv6DomainName field to given value.

### HasUseIpv6DomainName

`func (o *GetRoaminghostResponse) HasUseIpv6DomainName() bool`

HasUseIpv6DomainName returns a boolean if a field has been set.

### GetUseIpv6DomainNameServers

`func (o *GetRoaminghostResponse) GetUseIpv6DomainNameServers() bool`

GetUseIpv6DomainNameServers returns the UseIpv6DomainNameServers field if non-nil, zero value otherwise.

### GetUseIpv6DomainNameServersOk

`func (o *GetRoaminghostResponse) GetUseIpv6DomainNameServersOk() (*bool, bool)`

GetUseIpv6DomainNameServersOk returns a tuple with the UseIpv6DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DomainNameServers

`func (o *GetRoaminghostResponse) SetUseIpv6DomainNameServers(v bool)`

SetUseIpv6DomainNameServers sets UseIpv6DomainNameServers field to given value.

### HasUseIpv6DomainNameServers

`func (o *GetRoaminghostResponse) HasUseIpv6DomainNameServers() bool`

HasUseIpv6DomainNameServers returns a boolean if a field has been set.

### GetUseIpv6EnableDdns

`func (o *GetRoaminghostResponse) GetUseIpv6EnableDdns() bool`

GetUseIpv6EnableDdns returns the UseIpv6EnableDdns field if non-nil, zero value otherwise.

### GetUseIpv6EnableDdnsOk

`func (o *GetRoaminghostResponse) GetUseIpv6EnableDdnsOk() (*bool, bool)`

GetUseIpv6EnableDdnsOk returns a tuple with the UseIpv6EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6EnableDdns

`func (o *GetRoaminghostResponse) SetUseIpv6EnableDdns(v bool)`

SetUseIpv6EnableDdns sets UseIpv6EnableDdns field to given value.

### HasUseIpv6EnableDdns

`func (o *GetRoaminghostResponse) HasUseIpv6EnableDdns() bool`

HasUseIpv6EnableDdns returns a boolean if a field has been set.

### GetUseIpv6Options

`func (o *GetRoaminghostResponse) GetUseIpv6Options() bool`

GetUseIpv6Options returns the UseIpv6Options field if non-nil, zero value otherwise.

### GetUseIpv6OptionsOk

`func (o *GetRoaminghostResponse) GetUseIpv6OptionsOk() (*bool, bool)`

GetUseIpv6OptionsOk returns a tuple with the UseIpv6Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6Options

`func (o *GetRoaminghostResponse) SetUseIpv6Options(v bool)`

SetUseIpv6Options sets UseIpv6Options field to given value.

### HasUseIpv6Options

`func (o *GetRoaminghostResponse) HasUseIpv6Options() bool`

HasUseIpv6Options returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetRoaminghostResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetRoaminghostResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetRoaminghostResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetRoaminghostResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetRoaminghostResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetRoaminghostResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetRoaminghostResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetRoaminghostResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *GetRoaminghostResponse) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *GetRoaminghostResponse) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *GetRoaminghostResponse) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *GetRoaminghostResponse) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetRoaminghostResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetRoaminghostResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetRoaminghostResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetRoaminghostResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *GetRoaminghostResponse) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *GetRoaminghostResponse) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *GetRoaminghostResponse) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *GetRoaminghostResponse) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetUuid

`func (o *GetRoaminghostResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetRoaminghostResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetRoaminghostResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetRoaminghostResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValidLifetime

`func (o *GetRoaminghostResponse) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *GetRoaminghostResponse) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *GetRoaminghostResponse) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *GetRoaminghostResponse) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetResult

`func (o *GetRoaminghostResponse) GetResult() Roaminghost`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRoaminghostResponse) GetResultOk() (*Roaminghost, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRoaminghostResponse) SetResult(v Roaminghost)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRoaminghostResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


