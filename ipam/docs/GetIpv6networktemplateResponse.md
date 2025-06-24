# GetIpv6networktemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowAnyNetmask** | Pointer to **bool** | This flag controls whether the template allows any netmask. You must specify a netmask when creating a network using this template. If you set this parameter to False, you must specify the \&quot;cidr\&quot; field for the network template object. | [optional] 
**AutoCreateReversezone** | Pointer to **bool** | This flag controls whether reverse zones are automatically created when the network is added. | [optional] 
**Cidr** | Pointer to **int64** | The CIDR of the network in CIDR format. | [optional] 
**CloudApiCompatible** | Pointer to **bool** | This flag controls whether this template can be used to create network objects in a cloud-computing deployment. | [optional] 
**Comment** | Pointer to **string** | Comment for the network; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network. | [optional] 
**DdnsEnableOptionFqdn** | Pointer to **bool** | Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DelegatedMember** | Pointer to [**Ipv6networktemplateDelegatedMember**](Ipv6networktemplateDelegatedMember.md) |  | [optional] 
**DomainName** | Pointer to **string** | Use this method to set or retrieve the domain_name value of a DHCP IPv6 Network object. | [optional] 
**DomainNameServers** | Pointer to **[]string** | Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP IPv6 network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FixedAddressTemplates** | Pointer to **[]string** | The list of IPv6 fixed address templates assigned to this IPv6 network template object. When you create an IPv6 network based on an IPv6 network template object that contains IPv6 fixed address templates, the IPv6 fixed addresses are created based on the associated IPv6 fixed address templates. | [optional] 
**Ipv6prefix** | Pointer to **string** | The IPv6 Address prefix of the DHCP IPv6 network. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6networktemplateLogicFilterRules**](Ipv6networktemplateLogicFilterRules.md) | This field contains the logic filters to be applied on this IPv6 network template. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Members** | Pointer to [**[]Ipv6networktemplateMembers**](Ipv6networktemplateMembers.md) | A list of members that serve DHCP for the network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \&quot;_struct\&quot; member to the struct type. | [optional] 
**Name** | Pointer to **string** | The name of this IPv6 network template. | [optional] 
**Options** | Pointer to [**[]Ipv6networktemplateOptions**](Ipv6networktemplateOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PreferredLifetime** | Pointer to **int64** | Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network object. | [optional] 
**RangeTemplates** | Pointer to **[]string** | The list of IPv6 address range templates assigned to this IPv6 network template object. When you create an IPv6 network based on an IPv6 network template object that contains IPv6 range templates, the IPv6 address ranges are created based on the associated IPv6 address range templates. | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**Rir** | Pointer to **string** | The registry (RIR) that allocated the IPv6 network address space. | [optional] [readonly] 
**RirOrganization** | Pointer to **string** | The RIR organization associated with the IPv6 network. | [optional] 
**RirRegistrationAction** | Pointer to **string** | The action for the RIR registration. | [optional] 
**RirRegistrationStatus** | Pointer to **string** | The registration status of the IPv6 network in RIR. | [optional] 
**SendRirRequest** | Pointer to **bool** | Determines whether to send the RIR registration request. | [optional] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsEnableOptionFqdn** | Pointer to **bool** | Use flag for: ddns_enable_option_fqdn | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDomainName** | Pointer to **bool** | Use flag for: domain_name | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**ValidLifetime** | Pointer to **int64** | Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network object. | [optional] 
**Result** | Pointer to [**Ipv6networktemplate**](Ipv6networktemplate.md) |  | [optional] 

## Methods

### NewGetIpv6networktemplateResponse

`func NewGetIpv6networktemplateResponse() *GetIpv6networktemplateResponse`

NewGetIpv6networktemplateResponse instantiates a new GetIpv6networktemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6networktemplateResponseWithDefaults

`func NewGetIpv6networktemplateResponseWithDefaults() *GetIpv6networktemplateResponse`

NewGetIpv6networktemplateResponseWithDefaults instantiates a new GetIpv6networktemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6networktemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6networktemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6networktemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6networktemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowAnyNetmask

`func (o *GetIpv6networktemplateResponse) GetAllowAnyNetmask() bool`

GetAllowAnyNetmask returns the AllowAnyNetmask field if non-nil, zero value otherwise.

### GetAllowAnyNetmaskOk

`func (o *GetIpv6networktemplateResponse) GetAllowAnyNetmaskOk() (*bool, bool)`

GetAllowAnyNetmaskOk returns a tuple with the AllowAnyNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyNetmask

`func (o *GetIpv6networktemplateResponse) SetAllowAnyNetmask(v bool)`

SetAllowAnyNetmask sets AllowAnyNetmask field to given value.

### HasAllowAnyNetmask

`func (o *GetIpv6networktemplateResponse) HasAllowAnyNetmask() bool`

HasAllowAnyNetmask returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *GetIpv6networktemplateResponse) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *GetIpv6networktemplateResponse) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *GetIpv6networktemplateResponse) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *GetIpv6networktemplateResponse) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetCidr

`func (o *GetIpv6networktemplateResponse) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetIpv6networktemplateResponse) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetIpv6networktemplateResponse) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetIpv6networktemplateResponse) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCloudApiCompatible

`func (o *GetIpv6networktemplateResponse) GetCloudApiCompatible() bool`

GetCloudApiCompatible returns the CloudApiCompatible field if non-nil, zero value otherwise.

### GetCloudApiCompatibleOk

`func (o *GetIpv6networktemplateResponse) GetCloudApiCompatibleOk() (*bool, bool)`

GetCloudApiCompatibleOk returns a tuple with the CloudApiCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudApiCompatible

`func (o *GetIpv6networktemplateResponse) SetCloudApiCompatible(v bool)`

SetCloudApiCompatible sets CloudApiCompatible field to given value.

### HasCloudApiCompatible

`func (o *GetIpv6networktemplateResponse) HasCloudApiCompatible() bool`

HasCloudApiCompatible returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6networktemplateResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6networktemplateResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6networktemplateResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6networktemplateResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetIpv6networktemplateResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetIpv6networktemplateResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetIpv6networktemplateResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetIpv6networktemplateResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsEnableOptionFqdn

`func (o *GetIpv6networktemplateResponse) GetDdnsEnableOptionFqdn() bool`

GetDdnsEnableOptionFqdn returns the DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetDdnsEnableOptionFqdnOk

`func (o *GetIpv6networktemplateResponse) GetDdnsEnableOptionFqdnOk() (*bool, bool)`

GetDdnsEnableOptionFqdnOk returns a tuple with the DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnableOptionFqdn

`func (o *GetIpv6networktemplateResponse) SetDdnsEnableOptionFqdn(v bool)`

SetDdnsEnableOptionFqdn sets DdnsEnableOptionFqdn field to given value.

### HasDdnsEnableOptionFqdn

`func (o *GetIpv6networktemplateResponse) HasDdnsEnableOptionFqdn() bool`

HasDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GetIpv6networktemplateResponse) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GetIpv6networktemplateResponse) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GetIpv6networktemplateResponse) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GetIpv6networktemplateResponse) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *GetIpv6networktemplateResponse) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *GetIpv6networktemplateResponse) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *GetIpv6networktemplateResponse) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *GetIpv6networktemplateResponse) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *GetIpv6networktemplateResponse) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *GetIpv6networktemplateResponse) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *GetIpv6networktemplateResponse) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *GetIpv6networktemplateResponse) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDelegatedMember

`func (o *GetIpv6networktemplateResponse) GetDelegatedMember() Ipv6networktemplateDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *GetIpv6networktemplateResponse) GetDelegatedMemberOk() (*Ipv6networktemplateDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *GetIpv6networktemplateResponse) SetDelegatedMember(v Ipv6networktemplateDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *GetIpv6networktemplateResponse) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetDomainName

`func (o *GetIpv6networktemplateResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetIpv6networktemplateResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetIpv6networktemplateResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetIpv6networktemplateResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *GetIpv6networktemplateResponse) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *GetIpv6networktemplateResponse) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *GetIpv6networktemplateResponse) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *GetIpv6networktemplateResponse) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetIpv6networktemplateResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetIpv6networktemplateResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetIpv6networktemplateResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetIpv6networktemplateResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetIpv6networktemplateResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetIpv6networktemplateResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetIpv6networktemplateResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetIpv6networktemplateResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFixedAddressTemplates

`func (o *GetIpv6networktemplateResponse) GetFixedAddressTemplates() []string`

GetFixedAddressTemplates returns the FixedAddressTemplates field if non-nil, zero value otherwise.

### GetFixedAddressTemplatesOk

`func (o *GetIpv6networktemplateResponse) GetFixedAddressTemplatesOk() (*[]string, bool)`

GetFixedAddressTemplatesOk returns a tuple with the FixedAddressTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAddressTemplates

`func (o *GetIpv6networktemplateResponse) SetFixedAddressTemplates(v []string)`

SetFixedAddressTemplates sets FixedAddressTemplates field to given value.

### HasFixedAddressTemplates

`func (o *GetIpv6networktemplateResponse) HasFixedAddressTemplates() bool`

HasFixedAddressTemplates returns a boolean if a field has been set.

### GetIpv6prefix

`func (o *GetIpv6networktemplateResponse) GetIpv6prefix() string`

GetIpv6prefix returns the Ipv6prefix field if non-nil, zero value otherwise.

### GetIpv6prefixOk

`func (o *GetIpv6networktemplateResponse) GetIpv6prefixOk() (*string, bool)`

GetIpv6prefixOk returns a tuple with the Ipv6prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefix

`func (o *GetIpv6networktemplateResponse) SetIpv6prefix(v string)`

SetIpv6prefix sets Ipv6prefix field to given value.

### HasIpv6prefix

`func (o *GetIpv6networktemplateResponse) HasIpv6prefix() bool`

HasIpv6prefix returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetIpv6networktemplateResponse) GetLogicFilterRules() []Ipv6networktemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetIpv6networktemplateResponse) GetLogicFilterRulesOk() (*[]Ipv6networktemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetIpv6networktemplateResponse) SetLogicFilterRules(v []Ipv6networktemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetIpv6networktemplateResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMembers

`func (o *GetIpv6networktemplateResponse) GetMembers() []Ipv6networktemplateMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetIpv6networktemplateResponse) GetMembersOk() (*[]Ipv6networktemplateMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetIpv6networktemplateResponse) SetMembers(v []Ipv6networktemplateMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetIpv6networktemplateResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6networktemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6networktemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6networktemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6networktemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *GetIpv6networktemplateResponse) GetOptions() []Ipv6networktemplateOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetIpv6networktemplateResponse) GetOptionsOk() (*[]Ipv6networktemplateOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetIpv6networktemplateResponse) SetOptions(v []Ipv6networktemplateOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetIpv6networktemplateResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *GetIpv6networktemplateResponse) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *GetIpv6networktemplateResponse) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *GetIpv6networktemplateResponse) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *GetIpv6networktemplateResponse) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetRangeTemplates

`func (o *GetIpv6networktemplateResponse) GetRangeTemplates() []string`

GetRangeTemplates returns the RangeTemplates field if non-nil, zero value otherwise.

### GetRangeTemplatesOk

`func (o *GetIpv6networktemplateResponse) GetRangeTemplatesOk() (*[]string, bool)`

GetRangeTemplatesOk returns a tuple with the RangeTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeTemplates

`func (o *GetIpv6networktemplateResponse) SetRangeTemplates(v []string)`

SetRangeTemplates sets RangeTemplates field to given value.

### HasRangeTemplates

`func (o *GetIpv6networktemplateResponse) HasRangeTemplates() bool`

HasRangeTemplates returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetIpv6networktemplateResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetIpv6networktemplateResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetIpv6networktemplateResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetIpv6networktemplateResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRir

`func (o *GetIpv6networktemplateResponse) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *GetIpv6networktemplateResponse) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *GetIpv6networktemplateResponse) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *GetIpv6networktemplateResponse) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *GetIpv6networktemplateResponse) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *GetIpv6networktemplateResponse) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *GetIpv6networktemplateResponse) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *GetIpv6networktemplateResponse) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *GetIpv6networktemplateResponse) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *GetIpv6networktemplateResponse) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *GetIpv6networktemplateResponse) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *GetIpv6networktemplateResponse) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *GetIpv6networktemplateResponse) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *GetIpv6networktemplateResponse) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *GetIpv6networktemplateResponse) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *GetIpv6networktemplateResponse) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *GetIpv6networktemplateResponse) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *GetIpv6networktemplateResponse) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *GetIpv6networktemplateResponse) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *GetIpv6networktemplateResponse) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GetIpv6networktemplateResponse) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GetIpv6networktemplateResponse) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GetIpv6networktemplateResponse) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GetIpv6networktemplateResponse) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetIpv6networktemplateResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetIpv6networktemplateResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetIpv6networktemplateResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetIpv6networktemplateResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsEnableOptionFqdn

`func (o *GetIpv6networktemplateResponse) GetUseDdnsEnableOptionFqdn() bool`

GetUseDdnsEnableOptionFqdn returns the UseDdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetUseDdnsEnableOptionFqdnOk

`func (o *GetIpv6networktemplateResponse) GetUseDdnsEnableOptionFqdnOk() (*bool, bool)`

GetUseDdnsEnableOptionFqdnOk returns a tuple with the UseDdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsEnableOptionFqdn

`func (o *GetIpv6networktemplateResponse) SetUseDdnsEnableOptionFqdn(v bool)`

SetUseDdnsEnableOptionFqdn sets UseDdnsEnableOptionFqdn field to given value.

### HasUseDdnsEnableOptionFqdn

`func (o *GetIpv6networktemplateResponse) HasUseDdnsEnableOptionFqdn() bool`

HasUseDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *GetIpv6networktemplateResponse) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *GetIpv6networktemplateResponse) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *GetIpv6networktemplateResponse) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *GetIpv6networktemplateResponse) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *GetIpv6networktemplateResponse) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *GetIpv6networktemplateResponse) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *GetIpv6networktemplateResponse) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *GetIpv6networktemplateResponse) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDomainName

`func (o *GetIpv6networktemplateResponse) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *GetIpv6networktemplateResponse) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *GetIpv6networktemplateResponse) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *GetIpv6networktemplateResponse) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *GetIpv6networktemplateResponse) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *GetIpv6networktemplateResponse) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *GetIpv6networktemplateResponse) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *GetIpv6networktemplateResponse) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetIpv6networktemplateResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetIpv6networktemplateResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetIpv6networktemplateResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetIpv6networktemplateResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetIpv6networktemplateResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetIpv6networktemplateResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetIpv6networktemplateResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetIpv6networktemplateResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetIpv6networktemplateResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetIpv6networktemplateResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetIpv6networktemplateResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetIpv6networktemplateResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *GetIpv6networktemplateResponse) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *GetIpv6networktemplateResponse) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *GetIpv6networktemplateResponse) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *GetIpv6networktemplateResponse) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetIpv6networktemplateResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetIpv6networktemplateResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetIpv6networktemplateResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetIpv6networktemplateResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *GetIpv6networktemplateResponse) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *GetIpv6networktemplateResponse) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *GetIpv6networktemplateResponse) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *GetIpv6networktemplateResponse) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *GetIpv6networktemplateResponse) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *GetIpv6networktemplateResponse) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *GetIpv6networktemplateResponse) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *GetIpv6networktemplateResponse) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *GetIpv6networktemplateResponse) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *GetIpv6networktemplateResponse) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *GetIpv6networktemplateResponse) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *GetIpv6networktemplateResponse) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6networktemplateResponse) GetResult() Ipv6networktemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6networktemplateResponse) GetResultOk() (*Ipv6networktemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6networktemplateResponse) SetResult(v Ipv6networktemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6networktemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


