# Ipv6networktemplate

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
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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

## Methods

### NewIpv6networktemplate

`func NewIpv6networktemplate() *Ipv6networktemplate`

NewIpv6networktemplate instantiates a new Ipv6networktemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6networktemplateWithDefaults

`func NewIpv6networktemplateWithDefaults() *Ipv6networktemplate`

NewIpv6networktemplateWithDefaults instantiates a new Ipv6networktemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6networktemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6networktemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6networktemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6networktemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowAnyNetmask

`func (o *Ipv6networktemplate) GetAllowAnyNetmask() bool`

GetAllowAnyNetmask returns the AllowAnyNetmask field if non-nil, zero value otherwise.

### GetAllowAnyNetmaskOk

`func (o *Ipv6networktemplate) GetAllowAnyNetmaskOk() (*bool, bool)`

GetAllowAnyNetmaskOk returns a tuple with the AllowAnyNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyNetmask

`func (o *Ipv6networktemplate) SetAllowAnyNetmask(v bool)`

SetAllowAnyNetmask sets AllowAnyNetmask field to given value.

### HasAllowAnyNetmask

`func (o *Ipv6networktemplate) HasAllowAnyNetmask() bool`

HasAllowAnyNetmask returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *Ipv6networktemplate) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *Ipv6networktemplate) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *Ipv6networktemplate) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *Ipv6networktemplate) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetCidr

`func (o *Ipv6networktemplate) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Ipv6networktemplate) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Ipv6networktemplate) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *Ipv6networktemplate) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCloudApiCompatible

`func (o *Ipv6networktemplate) GetCloudApiCompatible() bool`

GetCloudApiCompatible returns the CloudApiCompatible field if non-nil, zero value otherwise.

### GetCloudApiCompatibleOk

`func (o *Ipv6networktemplate) GetCloudApiCompatibleOk() (*bool, bool)`

GetCloudApiCompatibleOk returns a tuple with the CloudApiCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudApiCompatible

`func (o *Ipv6networktemplate) SetCloudApiCompatible(v bool)`

SetCloudApiCompatible sets CloudApiCompatible field to given value.

### HasCloudApiCompatible

`func (o *Ipv6networktemplate) HasCloudApiCompatible() bool`

HasCloudApiCompatible returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6networktemplate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6networktemplate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6networktemplate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6networktemplate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Ipv6networktemplate) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Ipv6networktemplate) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Ipv6networktemplate) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Ipv6networktemplate) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsEnableOptionFqdn

`func (o *Ipv6networktemplate) GetDdnsEnableOptionFqdn() bool`

GetDdnsEnableOptionFqdn returns the DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetDdnsEnableOptionFqdnOk

`func (o *Ipv6networktemplate) GetDdnsEnableOptionFqdnOk() (*bool, bool)`

GetDdnsEnableOptionFqdnOk returns a tuple with the DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnableOptionFqdn

`func (o *Ipv6networktemplate) SetDdnsEnableOptionFqdn(v bool)`

SetDdnsEnableOptionFqdn sets DdnsEnableOptionFqdn field to given value.

### HasDdnsEnableOptionFqdn

`func (o *Ipv6networktemplate) HasDdnsEnableOptionFqdn() bool`

HasDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *Ipv6networktemplate) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *Ipv6networktemplate) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *Ipv6networktemplate) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *Ipv6networktemplate) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *Ipv6networktemplate) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *Ipv6networktemplate) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *Ipv6networktemplate) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *Ipv6networktemplate) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *Ipv6networktemplate) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *Ipv6networktemplate) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *Ipv6networktemplate) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *Ipv6networktemplate) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDelegatedMember

`func (o *Ipv6networktemplate) GetDelegatedMember() Ipv6networktemplateDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *Ipv6networktemplate) GetDelegatedMemberOk() (*Ipv6networktemplateDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *Ipv6networktemplate) SetDelegatedMember(v Ipv6networktemplateDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *Ipv6networktemplate) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetDomainName

`func (o *Ipv6networktemplate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Ipv6networktemplate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Ipv6networktemplate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Ipv6networktemplate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *Ipv6networktemplate) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *Ipv6networktemplate) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *Ipv6networktemplate) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *Ipv6networktemplate) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Ipv6networktemplate) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Ipv6networktemplate) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Ipv6networktemplate) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Ipv6networktemplate) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetExtattrs

`func (o *Ipv6networktemplate) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Ipv6networktemplate) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Ipv6networktemplate) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Ipv6networktemplate) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFixedAddressTemplates

`func (o *Ipv6networktemplate) GetFixedAddressTemplates() []string`

GetFixedAddressTemplates returns the FixedAddressTemplates field if non-nil, zero value otherwise.

### GetFixedAddressTemplatesOk

`func (o *Ipv6networktemplate) GetFixedAddressTemplatesOk() (*[]string, bool)`

GetFixedAddressTemplatesOk returns a tuple with the FixedAddressTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAddressTemplates

`func (o *Ipv6networktemplate) SetFixedAddressTemplates(v []string)`

SetFixedAddressTemplates sets FixedAddressTemplates field to given value.

### HasFixedAddressTemplates

`func (o *Ipv6networktemplate) HasFixedAddressTemplates() bool`

HasFixedAddressTemplates returns a boolean if a field has been set.

### GetIpv6prefix

`func (o *Ipv6networktemplate) GetIpv6prefix() string`

GetIpv6prefix returns the Ipv6prefix field if non-nil, zero value otherwise.

### GetIpv6prefixOk

`func (o *Ipv6networktemplate) GetIpv6prefixOk() (*string, bool)`

GetIpv6prefixOk returns a tuple with the Ipv6prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefix

`func (o *Ipv6networktemplate) SetIpv6prefix(v string)`

SetIpv6prefix sets Ipv6prefix field to given value.

### HasIpv6prefix

`func (o *Ipv6networktemplate) HasIpv6prefix() bool`

HasIpv6prefix returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6networktemplate) GetLogicFilterRules() []Ipv6networktemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6networktemplate) GetLogicFilterRulesOk() (*[]Ipv6networktemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6networktemplate) SetLogicFilterRules(v []Ipv6networktemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6networktemplate) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMembers

`func (o *Ipv6networktemplate) GetMembers() []Ipv6networktemplateMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Ipv6networktemplate) GetMembersOk() (*[]Ipv6networktemplateMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Ipv6networktemplate) SetMembers(v []Ipv6networktemplateMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Ipv6networktemplate) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *Ipv6networktemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6networktemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6networktemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6networktemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *Ipv6networktemplate) GetOptions() []Ipv6networktemplateOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Ipv6networktemplate) GetOptionsOk() (*[]Ipv6networktemplateOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Ipv6networktemplate) SetOptions(v []Ipv6networktemplateOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Ipv6networktemplate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *Ipv6networktemplate) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *Ipv6networktemplate) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *Ipv6networktemplate) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *Ipv6networktemplate) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetRangeTemplates

`func (o *Ipv6networktemplate) GetRangeTemplates() []string`

GetRangeTemplates returns the RangeTemplates field if non-nil, zero value otherwise.

### GetRangeTemplatesOk

`func (o *Ipv6networktemplate) GetRangeTemplatesOk() (*[]string, bool)`

GetRangeTemplatesOk returns a tuple with the RangeTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeTemplates

`func (o *Ipv6networktemplate) SetRangeTemplates(v []string)`

SetRangeTemplates sets RangeTemplates field to given value.

### HasRangeTemplates

`func (o *Ipv6networktemplate) HasRangeTemplates() bool`

HasRangeTemplates returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Ipv6networktemplate) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Ipv6networktemplate) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Ipv6networktemplate) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Ipv6networktemplate) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRir

`func (o *Ipv6networktemplate) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *Ipv6networktemplate) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *Ipv6networktemplate) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *Ipv6networktemplate) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *Ipv6networktemplate) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *Ipv6networktemplate) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *Ipv6networktemplate) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *Ipv6networktemplate) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *Ipv6networktemplate) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *Ipv6networktemplate) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *Ipv6networktemplate) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *Ipv6networktemplate) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *Ipv6networktemplate) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *Ipv6networktemplate) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *Ipv6networktemplate) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *Ipv6networktemplate) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *Ipv6networktemplate) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *Ipv6networktemplate) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *Ipv6networktemplate) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *Ipv6networktemplate) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *Ipv6networktemplate) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *Ipv6networktemplate) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *Ipv6networktemplate) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *Ipv6networktemplate) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Ipv6networktemplate) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Ipv6networktemplate) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Ipv6networktemplate) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Ipv6networktemplate) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsEnableOptionFqdn

`func (o *Ipv6networktemplate) GetUseDdnsEnableOptionFqdn() bool`

GetUseDdnsEnableOptionFqdn returns the UseDdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetUseDdnsEnableOptionFqdnOk

`func (o *Ipv6networktemplate) GetUseDdnsEnableOptionFqdnOk() (*bool, bool)`

GetUseDdnsEnableOptionFqdnOk returns a tuple with the UseDdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsEnableOptionFqdn

`func (o *Ipv6networktemplate) SetUseDdnsEnableOptionFqdn(v bool)`

SetUseDdnsEnableOptionFqdn sets UseDdnsEnableOptionFqdn field to given value.

### HasUseDdnsEnableOptionFqdn

`func (o *Ipv6networktemplate) HasUseDdnsEnableOptionFqdn() bool`

HasUseDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *Ipv6networktemplate) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *Ipv6networktemplate) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *Ipv6networktemplate) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *Ipv6networktemplate) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *Ipv6networktemplate) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *Ipv6networktemplate) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *Ipv6networktemplate) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *Ipv6networktemplate) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDomainName

`func (o *Ipv6networktemplate) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *Ipv6networktemplate) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *Ipv6networktemplate) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *Ipv6networktemplate) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *Ipv6networktemplate) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *Ipv6networktemplate) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *Ipv6networktemplate) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *Ipv6networktemplate) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Ipv6networktemplate) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Ipv6networktemplate) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Ipv6networktemplate) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Ipv6networktemplate) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6networktemplate) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6networktemplate) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6networktemplate) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6networktemplate) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *Ipv6networktemplate) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Ipv6networktemplate) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Ipv6networktemplate) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Ipv6networktemplate) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *Ipv6networktemplate) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *Ipv6networktemplate) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *Ipv6networktemplate) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *Ipv6networktemplate) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Ipv6networktemplate) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Ipv6networktemplate) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Ipv6networktemplate) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Ipv6networktemplate) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6networktemplate) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *Ipv6networktemplate) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6networktemplate) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6networktemplate) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *Ipv6networktemplate) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *Ipv6networktemplate) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *Ipv6networktemplate) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *Ipv6networktemplate) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *Ipv6networktemplate) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *Ipv6networktemplate) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *Ipv6networktemplate) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *Ipv6networktemplate) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


