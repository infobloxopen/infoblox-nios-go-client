# GetIpv6sharednetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the IPv6 shared network, maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_use_option81 is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of an IPv6 shared network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DdnsUseOption81** | Pointer to **bool** | The support for DHCP Option 81 at the IPv6 shared network level. | [optional] 
**Disable** | Pointer to **bool** | Determines whether an IPv6 shared network is disabled or not. When this is set to False, the IPv6 shared network is enabled. | [optional] 
**DomainName** | Pointer to **string** | Use this method to set or retrieve the domain_name value of a DHCP IPv6 Shared Network object. | [optional] 
**DomainNameServers** | Pointer to **[]string** | Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Shared Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of an IPv6 shared network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6sharednetworkLogicFilterRules**](Ipv6sharednetworkLogicFilterRules.md) | This field contains the logic filters to be applied on the this IPv6 shared network. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Name** | Pointer to **string** | The name of the IPv6 Shared Network. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this IPv6 shared network resides. | [optional] 
**Networks** | Pointer to [**[]Ipv6sharednetworkNetworks**](Ipv6sharednetworkNetworks.md) | A list of IPv6 networks belonging to the shared network Each individual list item must be specified as an object containing a &#39;_ref&#39; parameter to a network reference, for example:: [{ \&quot;_ref\&quot;: \&quot;ipv6network/ZG5zdHdvcmskMTAuAvMTYvMA\&quot;, }] if the reference of the wanted network is not known, it is possible to specify search parameters for the network instead in the following way:: [{ \&quot;_ref\&quot;: { &#39;network&#39;: &#39;aabb::/64&#39;, } }] note that in this case the search must match exactly one network for the assignment to be successful. | [optional] 
**Options** | Pointer to [**[]Ipv6sharednetworkOptions**](Ipv6sharednetworkOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PreferredLifetime** | Pointer to **int64** | Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Shared Network object. | [optional] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDdnsUseOption81** | Pointer to **bool** | Use flag for: ddns_use_option81 | [optional] 
**UseDomainName** | Pointer to **bool** | Use flag for: domain_name | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**ValidLifetime** | Pointer to **int64** | Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Shared Network object. | [optional] 
**Result** | Pointer to [**Ipv6sharednetwork**](Ipv6sharednetwork.md) |  | [optional] 

## Methods

### NewGetIpv6sharednetworkResponse

`func NewGetIpv6sharednetworkResponse() *GetIpv6sharednetworkResponse`

NewGetIpv6sharednetworkResponse instantiates a new GetIpv6sharednetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6sharednetworkResponseWithDefaults

`func NewGetIpv6sharednetworkResponseWithDefaults() *GetIpv6sharednetworkResponse`

NewGetIpv6sharednetworkResponseWithDefaults instantiates a new GetIpv6sharednetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6sharednetworkResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6sharednetworkResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6sharednetworkResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6sharednetworkResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6sharednetworkResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6sharednetworkResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6sharednetworkResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6sharednetworkResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetIpv6sharednetworkResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetIpv6sharednetworkResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetIpv6sharednetworkResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetIpv6sharednetworkResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GetIpv6sharednetworkResponse) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GetIpv6sharednetworkResponse) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GetIpv6sharednetworkResponse) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GetIpv6sharednetworkResponse) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *GetIpv6sharednetworkResponse) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *GetIpv6sharednetworkResponse) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *GetIpv6sharednetworkResponse) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *GetIpv6sharednetworkResponse) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *GetIpv6sharednetworkResponse) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *GetIpv6sharednetworkResponse) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *GetIpv6sharednetworkResponse) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *GetIpv6sharednetworkResponse) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *GetIpv6sharednetworkResponse) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *GetIpv6sharednetworkResponse) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *GetIpv6sharednetworkResponse) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *GetIpv6sharednetworkResponse) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDisable

`func (o *GetIpv6sharednetworkResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetIpv6sharednetworkResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetIpv6sharednetworkResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetIpv6sharednetworkResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDomainName

`func (o *GetIpv6sharednetworkResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetIpv6sharednetworkResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetIpv6sharednetworkResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetIpv6sharednetworkResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *GetIpv6sharednetworkResponse) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *GetIpv6sharednetworkResponse) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *GetIpv6sharednetworkResponse) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *GetIpv6sharednetworkResponse) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetIpv6sharednetworkResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetIpv6sharednetworkResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetIpv6sharednetworkResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetIpv6sharednetworkResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetIpv6sharednetworkResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetIpv6sharednetworkResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetIpv6sharednetworkResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetIpv6sharednetworkResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetIpv6sharednetworkResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetIpv6sharednetworkResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetIpv6sharednetworkResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetIpv6sharednetworkResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetIpv6sharednetworkResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetIpv6sharednetworkResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetIpv6sharednetworkResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetIpv6sharednetworkResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetIpv6sharednetworkResponse) GetLogicFilterRules() []Ipv6sharednetworkLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetIpv6sharednetworkResponse) GetLogicFilterRulesOk() (*[]Ipv6sharednetworkLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetIpv6sharednetworkResponse) SetLogicFilterRules(v []Ipv6sharednetworkLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetIpv6sharednetworkResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6sharednetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6sharednetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6sharednetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6sharednetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetIpv6sharednetworkResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetIpv6sharednetworkResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetIpv6sharednetworkResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetIpv6sharednetworkResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworks

`func (o *GetIpv6sharednetworkResponse) GetNetworks() []Ipv6sharednetworkNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetIpv6sharednetworkResponse) GetNetworksOk() (*[]Ipv6sharednetworkNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetIpv6sharednetworkResponse) SetNetworks(v []Ipv6sharednetworkNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetIpv6sharednetworkResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOptions

`func (o *GetIpv6sharednetworkResponse) GetOptions() []Ipv6sharednetworkOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetIpv6sharednetworkResponse) GetOptionsOk() (*[]Ipv6sharednetworkOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetIpv6sharednetworkResponse) SetOptions(v []Ipv6sharednetworkOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetIpv6sharednetworkResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *GetIpv6sharednetworkResponse) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *GetIpv6sharednetworkResponse) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *GetIpv6sharednetworkResponse) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *GetIpv6sharednetworkResponse) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GetIpv6sharednetworkResponse) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GetIpv6sharednetworkResponse) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GetIpv6sharednetworkResponse) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GetIpv6sharednetworkResponse) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetIpv6sharednetworkResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetIpv6sharednetworkResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *GetIpv6sharednetworkResponse) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *GetIpv6sharednetworkResponse) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *GetIpv6sharednetworkResponse) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *GetIpv6sharednetworkResponse) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDdnsUseOption81

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsUseOption81() bool`

GetUseDdnsUseOption81 returns the UseDdnsUseOption81 field if non-nil, zero value otherwise.

### GetUseDdnsUseOption81Ok

`func (o *GetIpv6sharednetworkResponse) GetUseDdnsUseOption81Ok() (*bool, bool)`

GetUseDdnsUseOption81Ok returns a tuple with the UseDdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUseOption81

`func (o *GetIpv6sharednetworkResponse) SetUseDdnsUseOption81(v bool)`

SetUseDdnsUseOption81 sets UseDdnsUseOption81 field to given value.

### HasUseDdnsUseOption81

`func (o *GetIpv6sharednetworkResponse) HasUseDdnsUseOption81() bool`

HasUseDdnsUseOption81 returns a boolean if a field has been set.

### GetUseDomainName

`func (o *GetIpv6sharednetworkResponse) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *GetIpv6sharednetworkResponse) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *GetIpv6sharednetworkResponse) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *GetIpv6sharednetworkResponse) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *GetIpv6sharednetworkResponse) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *GetIpv6sharednetworkResponse) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *GetIpv6sharednetworkResponse) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *GetIpv6sharednetworkResponse) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetIpv6sharednetworkResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetIpv6sharednetworkResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetIpv6sharednetworkResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetIpv6sharednetworkResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetIpv6sharednetworkResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetIpv6sharednetworkResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetIpv6sharednetworkResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetIpv6sharednetworkResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetIpv6sharednetworkResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetIpv6sharednetworkResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetIpv6sharednetworkResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetIpv6sharednetworkResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *GetIpv6sharednetworkResponse) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *GetIpv6sharednetworkResponse) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *GetIpv6sharednetworkResponse) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *GetIpv6sharednetworkResponse) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *GetIpv6sharednetworkResponse) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *GetIpv6sharednetworkResponse) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *GetIpv6sharednetworkResponse) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *GetIpv6sharednetworkResponse) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *GetIpv6sharednetworkResponse) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *GetIpv6sharednetworkResponse) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *GetIpv6sharednetworkResponse) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *GetIpv6sharednetworkResponse) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *GetIpv6sharednetworkResponse) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *GetIpv6sharednetworkResponse) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *GetIpv6sharednetworkResponse) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *GetIpv6sharednetworkResponse) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6sharednetworkResponse) GetResult() Ipv6sharednetwork`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6sharednetworkResponse) GetResultOk() (*Ipv6sharednetwork, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6sharednetworkResponse) SetResult(v Ipv6sharednetwork)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6sharednetworkResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


