# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**BlacklistAction** | Pointer to **string** | The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist_ruleset method. Valid values are \&quot;REDIRECT\&quot; or \&quot;REFUSE\&quot;. The default value is \&quot;REFUSE\&quot;. | [optional] 
**BlacklistLogQuery** | Pointer to **bool** | The flag that indicates whether blacklist redirection queries are logged. Specify \&quot;true\&quot; to enable logging, or \&quot;false\&quot; to disable it. The default value is \&quot;false\&quot;. | [optional] 
**BlacklistRedirectAddresses** | Pointer to **[]string** | The array of IP addresses the appliance includes in the response it sends in place of a blacklisted IP address. | [optional] 
**BlacklistRedirectTtl** | Pointer to **int64** | The Time To Live (TTL) value of the synthetic DNS responses resulted from blacklist redirection. The TTL value is a 32-bit unsigned integer that represents the TTL in seconds. | [optional] 
**BlacklistRulesets** | Pointer to **[]string** | The name of the Ruleset object assigned at the Grid level for blacklist redirection. | [optional] 
**CloudInfo** | Pointer to [**ViewCloudInfo**](ViewCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the DNS view; maximum 64 characters. | [optional] 
**CustomRootNameServers** | Pointer to [**[]ViewCustomRootNameServers**](ViewCustomRootNameServers.md) | The list of customized root name servers. You can either select and use Internet root name servers or specify custom root name servers by providing a host name and IP address to which the Infoblox appliance can send queries. Include the specified parameter to set the attribute value. Omit the parameter to retrieve the attribute value. | [optional] 
**DdnsForceCreationTimestampUpdate** | Pointer to **bool** | Defines whether creation timestamp of RR should be updated &#39; when DDNS update happens even if there is no change to &#39; the RR. | [optional] 
**DdnsPrincipalGroup** | Pointer to **string** | The DDNS Principal cluster group name. | [optional] 
**DdnsPrincipalTracking** | Pointer to **bool** | The flag that indicates whether the DDNS principal track is enabled or disabled. | [optional] 
**DdnsRestrictPatterns** | Pointer to **bool** | The flag that indicates whether an option to restrict DDNS update request based on FQDN patterns is enabled or disabled. | [optional] 
**DdnsRestrictPatternsList** | Pointer to **[]string** | The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns. | [optional] 
**DdnsRestrictProtected** | Pointer to **bool** | The flag that indicates whether an option to restrict DDNS update request to protected resource records is enabled or disabled. | [optional] 
**DdnsRestrictSecure** | Pointer to **bool** | The flag that indicates whether DDNS update request for principal other than target resource record&#39;s principal is restricted. | [optional] 
**DdnsRestrictStatic** | Pointer to **bool** | The flag that indicates whether an option to restrict DDNS update request to resource records which are marked as &#39;STATIC&#39; is enabled or disabled. | [optional] 
**Disable** | Pointer to **bool** | Determines if the DNS view is disabled or not. When this is set to False, the DNS view is enabled. | [optional] 
**Dns64Enabled** | Pointer to **bool** | Determines if the DNS64 s enabled or not. | [optional] 
**Dns64Groups** | Pointer to **[]string** | The list of DNS64 synthesis groups associated with this DNS view. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Determines if the DNS security extension is enabled or not. | [optional] 
**DnssecExpiredSignaturesEnabled** | Pointer to **bool** | Determines if the DNS security extension accepts expired signatures or not. | [optional] 
**DnssecNegativeTrustAnchors** | Pointer to **[]string** | A list of zones for which the server does not perform DNSSEC validation. | [optional] 
**DnssecTrustedKeys** | Pointer to [**[]ViewDnssecTrustedKeys**](ViewDnssecTrustedKeys.md) | The list of trusted keys for the DNS security extension. | [optional] 
**DnssecValidationEnabled** | Pointer to **bool** | Determines if the DNS security validation is enabled or not. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes. | [optional] 
**EnableBlacklist** | Pointer to **bool** | Determines if the blacklist in a DNS view is enabled or not. | [optional] 
**EnableFixedRrsetOrderFqdns** | Pointer to **bool** | Determines if the fixed RRset order FQDN is enabled or not. | [optional] 
**EnableMatchRecursiveOnly** | Pointer to **bool** | Determines if the &#39;match-recursive-only&#39; option in a DNS view is enabled or not. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FilterAaaa** | Pointer to **string** | The type of AAAA filtering for this DNS view object. | [optional] 
**FilterAaaaList** | Pointer to [**[]ViewFilterAaaaList**](ViewFilterAaaaList.md) | Applies AAAA filtering to a named ACL, or to a list of IPv4/IPv6 addresses and networks from which queries are received. This field does not allow TSIG keys. | [optional] 
**FixedRrsetOrderFqdns** | Pointer to [**[]ViewFixedRrsetOrderFqdns**](ViewFixedRrsetOrderFqdns.md) | The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to &#39;true&#39;, unless the same request sets the enable field to &#39;false&#39;. | [optional] 
**ForwardOnly** | Pointer to **bool** | Determines if this DNS view sends queries to forwarders only or not. When the value is True, queries are sent to forwarders only, and not to other internal or Internet root servers. | [optional] 
**Forwarders** | Pointer to **[]string** | The list of forwarders for the DNS view. A forwarder is a name server to which other name servers first send their off-site queries. The forwarder builds up a cache of information, avoiding the need for other name servers to send queries off-site. | [optional] 
**IsDefault** | Pointer to **bool** | The NIOS appliance provides one default DNS view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one DNS view in the appliance. | [optional] [readonly] 
**LastQueriedAcl** | Pointer to [**[]ViewLastQueriedAcl**](ViewLastQueriedAcl.md) | Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings. | [optional] 
**MatchClients** | Pointer to [**[]ViewMatchClients**](ViewMatchClients.md) | A list of forwarders for the match clients. This list specifies a named ACL, or a list of IPv4/IPv6 addresses, networks, TSIG keys of clients that are allowed or denied access to the DNS view. | [optional] 
**MatchDestinations** | Pointer to [**[]ViewMatchDestinations**](ViewMatchDestinations.md) | A list of forwarders for the match destinations. This list specifies a name ACL, or a list of IPv4/IPv6 addresses, networks, TSIG keys of clients that are allowed or denied access to the DNS view. | [optional] 
**MaxCacheTtl** | Pointer to **int64** | The maximum number of seconds to cache ordinary (positive) answers. | [optional] 
**MaxNcacheTtl** | Pointer to **int64** | The maximum number of seconds to cache negative (NXDOMAIN) answers. | [optional] 
**MaxUdpSize** | Pointer to **int64** | The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes. | [optional] 
**Name** | Pointer to **string** | Name of the DNS view. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view object associated with this DNS view. | [optional] 
**NotifyDelay** | Pointer to **int64** | The number of seconds of delay the notify messages are sent to secondaries. | [optional] 
**NxdomainLogQuery** | Pointer to **bool** | The flag that indicates whether NXDOMAIN redirection queries are logged. Specify \&quot;true\&quot; to enable logging, or \&quot;false\&quot; to disable it. The default value is \&quot;false\&quot;. | [optional] 
**NxdomainRedirect** | Pointer to **bool** | Determines if NXDOMAIN redirection in a DNS view is enabled or not. | [optional] 
**NxdomainRedirectAddresses** | Pointer to **[]string** | The array with IPv4 addresses the appliance includes in the response it sends in place of an NXDOMAIN response. | [optional] 
**NxdomainRedirectAddressesV6** | Pointer to **[]string** | The array with IPv6 addresses the appliance includes in the response it sends in place of an NXDOMAIN response. | [optional] 
**NxdomainRedirectTtl** | Pointer to **int64** | The Time To Live (TTL) value of the synthetic DNS responses resulted from NXDOMAIN redirection. The TTL value is a 32-bit unsigned integer that represents the TTL in seconds. | [optional] 
**NxdomainRulesets** | Pointer to **[]string** | The names of the Ruleset objects assigned at the grid level for NXDOMAIN redirection. | [optional] 
**Recursion** | Pointer to **bool** | Determines if recursion is enabled or not. | [optional] 
**ResponseRateLimiting** | Pointer to [**ViewResponseRateLimiting**](ViewResponseRateLimiting.md) |  | [optional] 
**RootNameServerType** | Pointer to **string** | Determines the type of root name servers. | [optional] 
**RpzDropIpRuleEnabled** | Pointer to **bool** | Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv4** | Pointer to **int64** | The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv6** | Pointer to **int64** | The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length. | [optional] 
**RpzQnameWaitRecurse** | Pointer to **bool** | The flag that indicates whether recursive RPZ lookups are enabled. | [optional] 
**ScavengingSettings** | Pointer to [**ViewScavengingSettings**](ViewScavengingSettings.md) |  | [optional] 
**Sortlist** | Pointer to [**[]ViewSortlist**](ViewSortlist.md) | A sort list that determines the order of IP addresses in responses sent to DNS queries. | [optional] 
**UseBlacklist** | Pointer to **bool** | Use flag for: blacklist_action , blacklist_log_query, blacklist_redirect_addresses, blacklist_redirect_ttl, blacklist_rulesets, enable_blacklist | [optional] 
**UseDdnsForceCreationTimestampUpdate** | Pointer to **bool** | Use flag for: ddns_force_creation_timestamp_update | [optional] 
**UseDdnsPatternsRestriction** | Pointer to **bool** | Use flag for: ddns_restrict_patterns_list , ddns_restrict_patterns | [optional] 
**UseDdnsPrincipalSecurity** | Pointer to **bool** | Use flag for: ddns_restrict_secure , ddns_principal_tracking, ddns_principal_group | [optional] 
**UseDdnsRestrictProtected** | Pointer to **bool** | Use flag for: ddns_restrict_protected | [optional] 
**UseDdnsRestrictStatic** | Pointer to **bool** | Use flag for: ddns_restrict_static | [optional] 
**UseDns64** | Pointer to **bool** | Use flag for: dns64_enabled , dns64_groups | [optional] 
**UseDnssec** | Pointer to **bool** | Use flag for: dnssec_enabled , dnssec_expired_signatures_enabled, dnssec_validation_enabled, dnssec_trusted_keys | [optional] 
**UseEdnsUdpSize** | Pointer to **bool** | Use flag for: edns_udp_size | [optional] 
**UseFilterAaaa** | Pointer to **bool** | Use flag for: filter_aaaa , filter_aaaa_list | [optional] 
**UseFixedRrsetOrderFqdns** | Pointer to **bool** | Use flag for: fixed_rrset_order_fqdns , enable_fixed_rrset_order_fqdns | [optional] 
**UseForwarders** | Pointer to **bool** | Use flag for: forwarders , forward_only | [optional] 
**UseMaxCacheTtl** | Pointer to **bool** | Use flag for: max_cache_ttl | [optional] 
**UseMaxNcacheTtl** | Pointer to **bool** | Use flag for: max_ncache_ttl | [optional] 
**UseMaxUdpSize** | Pointer to **bool** | Use flag for: max_udp_size | [optional] 
**UseNxdomainRedirect** | Pointer to **bool** | Use flag for: nxdomain_redirect , nxdomain_redirect_addresses, nxdomain_redirect_addresses_v6, nxdomain_redirect_ttl, nxdomain_log_query, nxdomain_rulesets | [optional] 
**UseRecursion** | Pointer to **bool** | Use flag for: recursion | [optional] 
**UseResponseRateLimiting** | Pointer to **bool** | Use flag for: response_rate_limiting | [optional] 
**UseRootNameServer** | Pointer to **bool** | Use flag for: custom_root_name_servers , root_name_server_type | [optional] 
**UseRpzDropIpRule** | Pointer to **bool** | Use flag for: rpz_drop_ip_rule_enabled , rpz_drop_ip_rule_min_prefix_length_ipv4, rpz_drop_ip_rule_min_prefix_length_ipv6 | [optional] 
**UseRpzQnameWaitRecurse** | Pointer to **bool** | Use flag for: rpz_qname_wait_recurse | [optional] 
**UseScavengingSettings** | Pointer to **bool** | Use flag for: scavenging_settings , last_queried_acl | [optional] 
**UseSortlist** | Pointer to **bool** | Use flag for: sortlist | [optional] 

## Methods

### NewView

`func NewView() *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *View) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *View) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *View) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *View) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBlacklistAction

`func (o *View) GetBlacklistAction() string`

GetBlacklistAction returns the BlacklistAction field if non-nil, zero value otherwise.

### GetBlacklistActionOk

`func (o *View) GetBlacklistActionOk() (*string, bool)`

GetBlacklistActionOk returns a tuple with the BlacklistAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistAction

`func (o *View) SetBlacklistAction(v string)`

SetBlacklistAction sets BlacklistAction field to given value.

### HasBlacklistAction

`func (o *View) HasBlacklistAction() bool`

HasBlacklistAction returns a boolean if a field has been set.

### GetBlacklistLogQuery

`func (o *View) GetBlacklistLogQuery() bool`

GetBlacklistLogQuery returns the BlacklistLogQuery field if non-nil, zero value otherwise.

### GetBlacklistLogQueryOk

`func (o *View) GetBlacklistLogQueryOk() (*bool, bool)`

GetBlacklistLogQueryOk returns a tuple with the BlacklistLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistLogQuery

`func (o *View) SetBlacklistLogQuery(v bool)`

SetBlacklistLogQuery sets BlacklistLogQuery field to given value.

### HasBlacklistLogQuery

`func (o *View) HasBlacklistLogQuery() bool`

HasBlacklistLogQuery returns a boolean if a field has been set.

### GetBlacklistRedirectAddresses

`func (o *View) GetBlacklistRedirectAddresses() []string`

GetBlacklistRedirectAddresses returns the BlacklistRedirectAddresses field if non-nil, zero value otherwise.

### GetBlacklistRedirectAddressesOk

`func (o *View) GetBlacklistRedirectAddressesOk() (*[]string, bool)`

GetBlacklistRedirectAddressesOk returns a tuple with the BlacklistRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectAddresses

`func (o *View) SetBlacklistRedirectAddresses(v []string)`

SetBlacklistRedirectAddresses sets BlacklistRedirectAddresses field to given value.

### HasBlacklistRedirectAddresses

`func (o *View) HasBlacklistRedirectAddresses() bool`

HasBlacklistRedirectAddresses returns a boolean if a field has been set.

### GetBlacklistRedirectTtl

`func (o *View) GetBlacklistRedirectTtl() int64`

GetBlacklistRedirectTtl returns the BlacklistRedirectTtl field if non-nil, zero value otherwise.

### GetBlacklistRedirectTtlOk

`func (o *View) GetBlacklistRedirectTtlOk() (*int64, bool)`

GetBlacklistRedirectTtlOk returns a tuple with the BlacklistRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectTtl

`func (o *View) SetBlacklistRedirectTtl(v int64)`

SetBlacklistRedirectTtl sets BlacklistRedirectTtl field to given value.

### HasBlacklistRedirectTtl

`func (o *View) HasBlacklistRedirectTtl() bool`

HasBlacklistRedirectTtl returns a boolean if a field has been set.

### GetBlacklistRulesets

`func (o *View) GetBlacklistRulesets() []string`

GetBlacklistRulesets returns the BlacklistRulesets field if non-nil, zero value otherwise.

### GetBlacklistRulesetsOk

`func (o *View) GetBlacklistRulesetsOk() (*[]string, bool)`

GetBlacklistRulesetsOk returns a tuple with the BlacklistRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRulesets

`func (o *View) SetBlacklistRulesets(v []string)`

SetBlacklistRulesets sets BlacklistRulesets field to given value.

### HasBlacklistRulesets

`func (o *View) HasBlacklistRulesets() bool`

HasBlacklistRulesets returns a boolean if a field has been set.

### GetCloudInfo

`func (o *View) GetCloudInfo() ViewCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *View) GetCloudInfoOk() (*ViewCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *View) SetCloudInfo(v ViewCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *View) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *View) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *View) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *View) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *View) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCustomRootNameServers

`func (o *View) GetCustomRootNameServers() []ViewCustomRootNameServers`

GetCustomRootNameServers returns the CustomRootNameServers field if non-nil, zero value otherwise.

### GetCustomRootNameServersOk

`func (o *View) GetCustomRootNameServersOk() (*[]ViewCustomRootNameServers, bool)`

GetCustomRootNameServersOk returns a tuple with the CustomRootNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNameServers

`func (o *View) SetCustomRootNameServers(v []ViewCustomRootNameServers)`

SetCustomRootNameServers sets CustomRootNameServers field to given value.

### HasCustomRootNameServers

`func (o *View) HasCustomRootNameServers() bool`

HasCustomRootNameServers returns a boolean if a field has been set.

### GetDdnsForceCreationTimestampUpdate

`func (o *View) GetDdnsForceCreationTimestampUpdate() bool`

GetDdnsForceCreationTimestampUpdate returns the DdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetDdnsForceCreationTimestampUpdateOk

`func (o *View) GetDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetDdnsForceCreationTimestampUpdateOk returns a tuple with the DdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsForceCreationTimestampUpdate

`func (o *View) SetDdnsForceCreationTimestampUpdate(v bool)`

SetDdnsForceCreationTimestampUpdate sets DdnsForceCreationTimestampUpdate field to given value.

### HasDdnsForceCreationTimestampUpdate

`func (o *View) HasDdnsForceCreationTimestampUpdate() bool`

HasDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetDdnsPrincipalGroup

`func (o *View) GetDdnsPrincipalGroup() string`

GetDdnsPrincipalGroup returns the DdnsPrincipalGroup field if non-nil, zero value otherwise.

### GetDdnsPrincipalGroupOk

`func (o *View) GetDdnsPrincipalGroupOk() (*string, bool)`

GetDdnsPrincipalGroupOk returns a tuple with the DdnsPrincipalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalGroup

`func (o *View) SetDdnsPrincipalGroup(v string)`

SetDdnsPrincipalGroup sets DdnsPrincipalGroup field to given value.

### HasDdnsPrincipalGroup

`func (o *View) HasDdnsPrincipalGroup() bool`

HasDdnsPrincipalGroup returns a boolean if a field has been set.

### GetDdnsPrincipalTracking

`func (o *View) GetDdnsPrincipalTracking() bool`

GetDdnsPrincipalTracking returns the DdnsPrincipalTracking field if non-nil, zero value otherwise.

### GetDdnsPrincipalTrackingOk

`func (o *View) GetDdnsPrincipalTrackingOk() (*bool, bool)`

GetDdnsPrincipalTrackingOk returns a tuple with the DdnsPrincipalTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalTracking

`func (o *View) SetDdnsPrincipalTracking(v bool)`

SetDdnsPrincipalTracking sets DdnsPrincipalTracking field to given value.

### HasDdnsPrincipalTracking

`func (o *View) HasDdnsPrincipalTracking() bool`

HasDdnsPrincipalTracking returns a boolean if a field has been set.

### GetDdnsRestrictPatterns

`func (o *View) GetDdnsRestrictPatterns() bool`

GetDdnsRestrictPatterns returns the DdnsRestrictPatterns field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsOk

`func (o *View) GetDdnsRestrictPatternsOk() (*bool, bool)`

GetDdnsRestrictPatternsOk returns a tuple with the DdnsRestrictPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatterns

`func (o *View) SetDdnsRestrictPatterns(v bool)`

SetDdnsRestrictPatterns sets DdnsRestrictPatterns field to given value.

### HasDdnsRestrictPatterns

`func (o *View) HasDdnsRestrictPatterns() bool`

HasDdnsRestrictPatterns returns a boolean if a field has been set.

### GetDdnsRestrictPatternsList

`func (o *View) GetDdnsRestrictPatternsList() []string`

GetDdnsRestrictPatternsList returns the DdnsRestrictPatternsList field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsListOk

`func (o *View) GetDdnsRestrictPatternsListOk() (*[]string, bool)`

GetDdnsRestrictPatternsListOk returns a tuple with the DdnsRestrictPatternsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatternsList

`func (o *View) SetDdnsRestrictPatternsList(v []string)`

SetDdnsRestrictPatternsList sets DdnsRestrictPatternsList field to given value.

### HasDdnsRestrictPatternsList

`func (o *View) HasDdnsRestrictPatternsList() bool`

HasDdnsRestrictPatternsList returns a boolean if a field has been set.

### GetDdnsRestrictProtected

`func (o *View) GetDdnsRestrictProtected() bool`

GetDdnsRestrictProtected returns the DdnsRestrictProtected field if non-nil, zero value otherwise.

### GetDdnsRestrictProtectedOk

`func (o *View) GetDdnsRestrictProtectedOk() (*bool, bool)`

GetDdnsRestrictProtectedOk returns a tuple with the DdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictProtected

`func (o *View) SetDdnsRestrictProtected(v bool)`

SetDdnsRestrictProtected sets DdnsRestrictProtected field to given value.

### HasDdnsRestrictProtected

`func (o *View) HasDdnsRestrictProtected() bool`

HasDdnsRestrictProtected returns a boolean if a field has been set.

### GetDdnsRestrictSecure

`func (o *View) GetDdnsRestrictSecure() bool`

GetDdnsRestrictSecure returns the DdnsRestrictSecure field if non-nil, zero value otherwise.

### GetDdnsRestrictSecureOk

`func (o *View) GetDdnsRestrictSecureOk() (*bool, bool)`

GetDdnsRestrictSecureOk returns a tuple with the DdnsRestrictSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictSecure

`func (o *View) SetDdnsRestrictSecure(v bool)`

SetDdnsRestrictSecure sets DdnsRestrictSecure field to given value.

### HasDdnsRestrictSecure

`func (o *View) HasDdnsRestrictSecure() bool`

HasDdnsRestrictSecure returns a boolean if a field has been set.

### GetDdnsRestrictStatic

`func (o *View) GetDdnsRestrictStatic() bool`

GetDdnsRestrictStatic returns the DdnsRestrictStatic field if non-nil, zero value otherwise.

### GetDdnsRestrictStaticOk

`func (o *View) GetDdnsRestrictStaticOk() (*bool, bool)`

GetDdnsRestrictStaticOk returns a tuple with the DdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictStatic

`func (o *View) SetDdnsRestrictStatic(v bool)`

SetDdnsRestrictStatic sets DdnsRestrictStatic field to given value.

### HasDdnsRestrictStatic

`func (o *View) HasDdnsRestrictStatic() bool`

HasDdnsRestrictStatic returns a boolean if a field has been set.

### GetDisable

`func (o *View) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *View) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *View) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *View) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDns64Enabled

`func (o *View) GetDns64Enabled() bool`

GetDns64Enabled returns the Dns64Enabled field if non-nil, zero value otherwise.

### GetDns64EnabledOk

`func (o *View) GetDns64EnabledOk() (*bool, bool)`

GetDns64EnabledOk returns a tuple with the Dns64Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Enabled

`func (o *View) SetDns64Enabled(v bool)`

SetDns64Enabled sets Dns64Enabled field to given value.

### HasDns64Enabled

`func (o *View) HasDns64Enabled() bool`

HasDns64Enabled returns a boolean if a field has been set.

### GetDns64Groups

`func (o *View) GetDns64Groups() []string`

GetDns64Groups returns the Dns64Groups field if non-nil, zero value otherwise.

### GetDns64GroupsOk

`func (o *View) GetDns64GroupsOk() (*[]string, bool)`

GetDns64GroupsOk returns a tuple with the Dns64Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Groups

`func (o *View) SetDns64Groups(v []string)`

SetDns64Groups sets Dns64Groups field to given value.

### HasDns64Groups

`func (o *View) HasDns64Groups() bool`

HasDns64Groups returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *View) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *View) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *View) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *View) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecExpiredSignaturesEnabled

`func (o *View) GetDnssecExpiredSignaturesEnabled() bool`

GetDnssecExpiredSignaturesEnabled returns the DnssecExpiredSignaturesEnabled field if non-nil, zero value otherwise.

### GetDnssecExpiredSignaturesEnabledOk

`func (o *View) GetDnssecExpiredSignaturesEnabledOk() (*bool, bool)`

GetDnssecExpiredSignaturesEnabledOk returns a tuple with the DnssecExpiredSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecExpiredSignaturesEnabled

`func (o *View) SetDnssecExpiredSignaturesEnabled(v bool)`

SetDnssecExpiredSignaturesEnabled sets DnssecExpiredSignaturesEnabled field to given value.

### HasDnssecExpiredSignaturesEnabled

`func (o *View) HasDnssecExpiredSignaturesEnabled() bool`

HasDnssecExpiredSignaturesEnabled returns a boolean if a field has been set.

### GetDnssecNegativeTrustAnchors

`func (o *View) GetDnssecNegativeTrustAnchors() []string`

GetDnssecNegativeTrustAnchors returns the DnssecNegativeTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecNegativeTrustAnchorsOk

`func (o *View) GetDnssecNegativeTrustAnchorsOk() (*[]string, bool)`

GetDnssecNegativeTrustAnchorsOk returns a tuple with the DnssecNegativeTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNegativeTrustAnchors

`func (o *View) SetDnssecNegativeTrustAnchors(v []string)`

SetDnssecNegativeTrustAnchors sets DnssecNegativeTrustAnchors field to given value.

### HasDnssecNegativeTrustAnchors

`func (o *View) HasDnssecNegativeTrustAnchors() bool`

HasDnssecNegativeTrustAnchors returns a boolean if a field has been set.

### GetDnssecTrustedKeys

`func (o *View) GetDnssecTrustedKeys() []ViewDnssecTrustedKeys`

GetDnssecTrustedKeys returns the DnssecTrustedKeys field if non-nil, zero value otherwise.

### GetDnssecTrustedKeysOk

`func (o *View) GetDnssecTrustedKeysOk() (*[]ViewDnssecTrustedKeys, bool)`

GetDnssecTrustedKeysOk returns a tuple with the DnssecTrustedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustedKeys

`func (o *View) SetDnssecTrustedKeys(v []ViewDnssecTrustedKeys)`

SetDnssecTrustedKeys sets DnssecTrustedKeys field to given value.

### HasDnssecTrustedKeys

`func (o *View) HasDnssecTrustedKeys() bool`

HasDnssecTrustedKeys returns a boolean if a field has been set.

### GetDnssecValidationEnabled

`func (o *View) GetDnssecValidationEnabled() bool`

GetDnssecValidationEnabled returns the DnssecValidationEnabled field if non-nil, zero value otherwise.

### GetDnssecValidationEnabledOk

`func (o *View) GetDnssecValidationEnabledOk() (*bool, bool)`

GetDnssecValidationEnabledOk returns a tuple with the DnssecValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationEnabled

`func (o *View) SetDnssecValidationEnabled(v bool)`

SetDnssecValidationEnabled sets DnssecValidationEnabled field to given value.

### HasDnssecValidationEnabled

`func (o *View) HasDnssecValidationEnabled() bool`

HasDnssecValidationEnabled returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *View) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *View) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *View) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *View) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetEnableBlacklist

`func (o *View) GetEnableBlacklist() bool`

GetEnableBlacklist returns the EnableBlacklist field if non-nil, zero value otherwise.

### GetEnableBlacklistOk

`func (o *View) GetEnableBlacklistOk() (*bool, bool)`

GetEnableBlacklistOk returns a tuple with the EnableBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlacklist

`func (o *View) SetEnableBlacklist(v bool)`

SetEnableBlacklist sets EnableBlacklist field to given value.

### HasEnableBlacklist

`func (o *View) HasEnableBlacklist() bool`

HasEnableBlacklist returns a boolean if a field has been set.

### GetEnableFixedRrsetOrderFqdns

`func (o *View) GetEnableFixedRrsetOrderFqdns() bool`

GetEnableFixedRrsetOrderFqdns returns the EnableFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetEnableFixedRrsetOrderFqdnsOk

`func (o *View) GetEnableFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetEnableFixedRrsetOrderFqdnsOk returns a tuple with the EnableFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixedRrsetOrderFqdns

`func (o *View) SetEnableFixedRrsetOrderFqdns(v bool)`

SetEnableFixedRrsetOrderFqdns sets EnableFixedRrsetOrderFqdns field to given value.

### HasEnableFixedRrsetOrderFqdns

`func (o *View) HasEnableFixedRrsetOrderFqdns() bool`

HasEnableFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetEnableMatchRecursiveOnly

`func (o *View) GetEnableMatchRecursiveOnly() bool`

GetEnableMatchRecursiveOnly returns the EnableMatchRecursiveOnly field if non-nil, zero value otherwise.

### GetEnableMatchRecursiveOnlyOk

`func (o *View) GetEnableMatchRecursiveOnlyOk() (*bool, bool)`

GetEnableMatchRecursiveOnlyOk returns a tuple with the EnableMatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMatchRecursiveOnly

`func (o *View) SetEnableMatchRecursiveOnly(v bool)`

SetEnableMatchRecursiveOnly sets EnableMatchRecursiveOnly field to given value.

### HasEnableMatchRecursiveOnly

`func (o *View) HasEnableMatchRecursiveOnly() bool`

HasEnableMatchRecursiveOnly returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *View) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *View) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *View) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *View) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *View) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *View) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *View) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *View) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *View) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *View) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *View) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *View) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFilterAaaa

`func (o *View) GetFilterAaaa() string`

GetFilterAaaa returns the FilterAaaa field if non-nil, zero value otherwise.

### GetFilterAaaaOk

`func (o *View) GetFilterAaaaOk() (*string, bool)`

GetFilterAaaaOk returns a tuple with the FilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaa

`func (o *View) SetFilterAaaa(v string)`

SetFilterAaaa sets FilterAaaa field to given value.

### HasFilterAaaa

`func (o *View) HasFilterAaaa() bool`

HasFilterAaaa returns a boolean if a field has been set.

### GetFilterAaaaList

`func (o *View) GetFilterAaaaList() []ViewFilterAaaaList`

GetFilterAaaaList returns the FilterAaaaList field if non-nil, zero value otherwise.

### GetFilterAaaaListOk

`func (o *View) GetFilterAaaaListOk() (*[]ViewFilterAaaaList, bool)`

GetFilterAaaaListOk returns a tuple with the FilterAaaaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaList

`func (o *View) SetFilterAaaaList(v []ViewFilterAaaaList)`

SetFilterAaaaList sets FilterAaaaList field to given value.

### HasFilterAaaaList

`func (o *View) HasFilterAaaaList() bool`

HasFilterAaaaList returns a boolean if a field has been set.

### GetFixedRrsetOrderFqdns

`func (o *View) GetFixedRrsetOrderFqdns() []ViewFixedRrsetOrderFqdns`

GetFixedRrsetOrderFqdns returns the FixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetFixedRrsetOrderFqdnsOk

`func (o *View) GetFixedRrsetOrderFqdnsOk() (*[]ViewFixedRrsetOrderFqdns, bool)`

GetFixedRrsetOrderFqdnsOk returns a tuple with the FixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedRrsetOrderFqdns

`func (o *View) SetFixedRrsetOrderFqdns(v []ViewFixedRrsetOrderFqdns)`

SetFixedRrsetOrderFqdns sets FixedRrsetOrderFqdns field to given value.

### HasFixedRrsetOrderFqdns

`func (o *View) HasFixedRrsetOrderFqdns() bool`

HasFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetForwardOnly

`func (o *View) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *View) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *View) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *View) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetForwarders

`func (o *View) GetForwarders() []string`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *View) GetForwardersOk() (*[]string, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *View) SetForwarders(v []string)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *View) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetIsDefault

`func (o *View) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *View) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *View) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *View) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastQueriedAcl

`func (o *View) GetLastQueriedAcl() []ViewLastQueriedAcl`

GetLastQueriedAcl returns the LastQueriedAcl field if non-nil, zero value otherwise.

### GetLastQueriedAclOk

`func (o *View) GetLastQueriedAclOk() (*[]ViewLastQueriedAcl, bool)`

GetLastQueriedAclOk returns a tuple with the LastQueriedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueriedAcl

`func (o *View) SetLastQueriedAcl(v []ViewLastQueriedAcl)`

SetLastQueriedAcl sets LastQueriedAcl field to given value.

### HasLastQueriedAcl

`func (o *View) HasLastQueriedAcl() bool`

HasLastQueriedAcl returns a boolean if a field has been set.

### GetMatchClients

`func (o *View) GetMatchClients() []ViewMatchClients`

GetMatchClients returns the MatchClients field if non-nil, zero value otherwise.

### GetMatchClientsOk

`func (o *View) GetMatchClientsOk() (*[]ViewMatchClients, bool)`

GetMatchClientsOk returns a tuple with the MatchClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClients

`func (o *View) SetMatchClients(v []ViewMatchClients)`

SetMatchClients sets MatchClients field to given value.

### HasMatchClients

`func (o *View) HasMatchClients() bool`

HasMatchClients returns a boolean if a field has been set.

### GetMatchDestinations

`func (o *View) GetMatchDestinations() []ViewMatchDestinations`

GetMatchDestinations returns the MatchDestinations field if non-nil, zero value otherwise.

### GetMatchDestinationsOk

`func (o *View) GetMatchDestinationsOk() (*[]ViewMatchDestinations, bool)`

GetMatchDestinationsOk returns a tuple with the MatchDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDestinations

`func (o *View) SetMatchDestinations(v []ViewMatchDestinations)`

SetMatchDestinations sets MatchDestinations field to given value.

### HasMatchDestinations

`func (o *View) HasMatchDestinations() bool`

HasMatchDestinations returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *View) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *View) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *View) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *View) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNcacheTtl

`func (o *View) GetMaxNcacheTtl() int64`

GetMaxNcacheTtl returns the MaxNcacheTtl field if non-nil, zero value otherwise.

### GetMaxNcacheTtlOk

`func (o *View) GetMaxNcacheTtlOk() (*int64, bool)`

GetMaxNcacheTtlOk returns a tuple with the MaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNcacheTtl

`func (o *View) SetMaxNcacheTtl(v int64)`

SetMaxNcacheTtl sets MaxNcacheTtl field to given value.

### HasMaxNcacheTtl

`func (o *View) HasMaxNcacheTtl() bool`

HasMaxNcacheTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *View) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *View) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *View) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *View) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *View) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *View) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *View) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *View) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *View) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *View) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *View) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *View) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *View) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetNxdomainLogQuery

`func (o *View) GetNxdomainLogQuery() bool`

GetNxdomainLogQuery returns the NxdomainLogQuery field if non-nil, zero value otherwise.

### GetNxdomainLogQueryOk

`func (o *View) GetNxdomainLogQueryOk() (*bool, bool)`

GetNxdomainLogQueryOk returns a tuple with the NxdomainLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainLogQuery

`func (o *View) SetNxdomainLogQuery(v bool)`

SetNxdomainLogQuery sets NxdomainLogQuery field to given value.

### HasNxdomainLogQuery

`func (o *View) HasNxdomainLogQuery() bool`

HasNxdomainLogQuery returns a boolean if a field has been set.

### GetNxdomainRedirect

`func (o *View) GetNxdomainRedirect() bool`

GetNxdomainRedirect returns the NxdomainRedirect field if non-nil, zero value otherwise.

### GetNxdomainRedirectOk

`func (o *View) GetNxdomainRedirectOk() (*bool, bool)`

GetNxdomainRedirectOk returns a tuple with the NxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirect

`func (o *View) SetNxdomainRedirect(v bool)`

SetNxdomainRedirect sets NxdomainRedirect field to given value.

### HasNxdomainRedirect

`func (o *View) HasNxdomainRedirect() bool`

HasNxdomainRedirect returns a boolean if a field has been set.

### GetNxdomainRedirectAddresses

`func (o *View) GetNxdomainRedirectAddresses() []string`

GetNxdomainRedirectAddresses returns the NxdomainRedirectAddresses field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesOk

`func (o *View) GetNxdomainRedirectAddressesOk() (*[]string, bool)`

GetNxdomainRedirectAddressesOk returns a tuple with the NxdomainRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddresses

`func (o *View) SetNxdomainRedirectAddresses(v []string)`

SetNxdomainRedirectAddresses sets NxdomainRedirectAddresses field to given value.

### HasNxdomainRedirectAddresses

`func (o *View) HasNxdomainRedirectAddresses() bool`

HasNxdomainRedirectAddresses returns a boolean if a field has been set.

### GetNxdomainRedirectAddressesV6

`func (o *View) GetNxdomainRedirectAddressesV6() []string`

GetNxdomainRedirectAddressesV6 returns the NxdomainRedirectAddressesV6 field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesV6Ok

`func (o *View) GetNxdomainRedirectAddressesV6Ok() (*[]string, bool)`

GetNxdomainRedirectAddressesV6Ok returns a tuple with the NxdomainRedirectAddressesV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddressesV6

`func (o *View) SetNxdomainRedirectAddressesV6(v []string)`

SetNxdomainRedirectAddressesV6 sets NxdomainRedirectAddressesV6 field to given value.

### HasNxdomainRedirectAddressesV6

`func (o *View) HasNxdomainRedirectAddressesV6() bool`

HasNxdomainRedirectAddressesV6 returns a boolean if a field has been set.

### GetNxdomainRedirectTtl

`func (o *View) GetNxdomainRedirectTtl() int64`

GetNxdomainRedirectTtl returns the NxdomainRedirectTtl field if non-nil, zero value otherwise.

### GetNxdomainRedirectTtlOk

`func (o *View) GetNxdomainRedirectTtlOk() (*int64, bool)`

GetNxdomainRedirectTtlOk returns a tuple with the NxdomainRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectTtl

`func (o *View) SetNxdomainRedirectTtl(v int64)`

SetNxdomainRedirectTtl sets NxdomainRedirectTtl field to given value.

### HasNxdomainRedirectTtl

`func (o *View) HasNxdomainRedirectTtl() bool`

HasNxdomainRedirectTtl returns a boolean if a field has been set.

### GetNxdomainRulesets

`func (o *View) GetNxdomainRulesets() []string`

GetNxdomainRulesets returns the NxdomainRulesets field if non-nil, zero value otherwise.

### GetNxdomainRulesetsOk

`func (o *View) GetNxdomainRulesetsOk() (*[]string, bool)`

GetNxdomainRulesetsOk returns a tuple with the NxdomainRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRulesets

`func (o *View) SetNxdomainRulesets(v []string)`

SetNxdomainRulesets sets NxdomainRulesets field to given value.

### HasNxdomainRulesets

`func (o *View) HasNxdomainRulesets() bool`

HasNxdomainRulesets returns a boolean if a field has been set.

### GetRecursion

`func (o *View) GetRecursion() bool`

GetRecursion returns the Recursion field if non-nil, zero value otherwise.

### GetRecursionOk

`func (o *View) GetRecursionOk() (*bool, bool)`

GetRecursionOk returns a tuple with the Recursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursion

`func (o *View) SetRecursion(v bool)`

SetRecursion sets Recursion field to given value.

### HasRecursion

`func (o *View) HasRecursion() bool`

HasRecursion returns a boolean if a field has been set.

### GetResponseRateLimiting

`func (o *View) GetResponseRateLimiting() ViewResponseRateLimiting`

GetResponseRateLimiting returns the ResponseRateLimiting field if non-nil, zero value otherwise.

### GetResponseRateLimitingOk

`func (o *View) GetResponseRateLimitingOk() (*ViewResponseRateLimiting, bool)`

GetResponseRateLimitingOk returns a tuple with the ResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRateLimiting

`func (o *View) SetResponseRateLimiting(v ViewResponseRateLimiting)`

SetResponseRateLimiting sets ResponseRateLimiting field to given value.

### HasResponseRateLimiting

`func (o *View) HasResponseRateLimiting() bool`

HasResponseRateLimiting returns a boolean if a field has been set.

### GetRootNameServerType

`func (o *View) GetRootNameServerType() string`

GetRootNameServerType returns the RootNameServerType field if non-nil, zero value otherwise.

### GetRootNameServerTypeOk

`func (o *View) GetRootNameServerTypeOk() (*string, bool)`

GetRootNameServerTypeOk returns a tuple with the RootNameServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNameServerType

`func (o *View) SetRootNameServerType(v string)`

SetRootNameServerType sets RootNameServerType field to given value.

### HasRootNameServerType

`func (o *View) HasRootNameServerType() bool`

HasRootNameServerType returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *View) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *View) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *View) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *View) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *View) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *View) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *View) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *View) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *View) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *View) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *View) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *View) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzQnameWaitRecurse

`func (o *View) GetRpzQnameWaitRecurse() bool`

GetRpzQnameWaitRecurse returns the RpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetRpzQnameWaitRecurseOk

`func (o *View) GetRpzQnameWaitRecurseOk() (*bool, bool)`

GetRpzQnameWaitRecurseOk returns a tuple with the RpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzQnameWaitRecurse

`func (o *View) SetRpzQnameWaitRecurse(v bool)`

SetRpzQnameWaitRecurse sets RpzQnameWaitRecurse field to given value.

### HasRpzQnameWaitRecurse

`func (o *View) HasRpzQnameWaitRecurse() bool`

HasRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetScavengingSettings

`func (o *View) GetScavengingSettings() ViewScavengingSettings`

GetScavengingSettings returns the ScavengingSettings field if non-nil, zero value otherwise.

### GetScavengingSettingsOk

`func (o *View) GetScavengingSettingsOk() (*ViewScavengingSettings, bool)`

GetScavengingSettingsOk returns a tuple with the ScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSettings

`func (o *View) SetScavengingSettings(v ViewScavengingSettings)`

SetScavengingSettings sets ScavengingSettings field to given value.

### HasScavengingSettings

`func (o *View) HasScavengingSettings() bool`

HasScavengingSettings returns a boolean if a field has been set.

### GetSortlist

`func (o *View) GetSortlist() []ViewSortlist`

GetSortlist returns the Sortlist field if non-nil, zero value otherwise.

### GetSortlistOk

`func (o *View) GetSortlistOk() (*[]ViewSortlist, bool)`

GetSortlistOk returns a tuple with the Sortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortlist

`func (o *View) SetSortlist(v []ViewSortlist)`

SetSortlist sets Sortlist field to given value.

### HasSortlist

`func (o *View) HasSortlist() bool`

HasSortlist returns a boolean if a field has been set.

### GetUseBlacklist

`func (o *View) GetUseBlacklist() bool`

GetUseBlacklist returns the UseBlacklist field if non-nil, zero value otherwise.

### GetUseBlacklistOk

`func (o *View) GetUseBlacklistOk() (*bool, bool)`

GetUseBlacklistOk returns a tuple with the UseBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlacklist

`func (o *View) SetUseBlacklist(v bool)`

SetUseBlacklist sets UseBlacklist field to given value.

### HasUseBlacklist

`func (o *View) HasUseBlacklist() bool`

HasUseBlacklist returns a boolean if a field has been set.

### GetUseDdnsForceCreationTimestampUpdate

`func (o *View) GetUseDdnsForceCreationTimestampUpdate() bool`

GetUseDdnsForceCreationTimestampUpdate returns the UseDdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetUseDdnsForceCreationTimestampUpdateOk

`func (o *View) GetUseDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetUseDdnsForceCreationTimestampUpdateOk returns a tuple with the UseDdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsForceCreationTimestampUpdate

`func (o *View) SetUseDdnsForceCreationTimestampUpdate(v bool)`

SetUseDdnsForceCreationTimestampUpdate sets UseDdnsForceCreationTimestampUpdate field to given value.

### HasUseDdnsForceCreationTimestampUpdate

`func (o *View) HasUseDdnsForceCreationTimestampUpdate() bool`

HasUseDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetUseDdnsPatternsRestriction

`func (o *View) GetUseDdnsPatternsRestriction() bool`

GetUseDdnsPatternsRestriction returns the UseDdnsPatternsRestriction field if non-nil, zero value otherwise.

### GetUseDdnsPatternsRestrictionOk

`func (o *View) GetUseDdnsPatternsRestrictionOk() (*bool, bool)`

GetUseDdnsPatternsRestrictionOk returns a tuple with the UseDdnsPatternsRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsPatternsRestriction

`func (o *View) SetUseDdnsPatternsRestriction(v bool)`

SetUseDdnsPatternsRestriction sets UseDdnsPatternsRestriction field to given value.

### HasUseDdnsPatternsRestriction

`func (o *View) HasUseDdnsPatternsRestriction() bool`

HasUseDdnsPatternsRestriction returns a boolean if a field has been set.

### GetUseDdnsPrincipalSecurity

`func (o *View) GetUseDdnsPrincipalSecurity() bool`

GetUseDdnsPrincipalSecurity returns the UseDdnsPrincipalSecurity field if non-nil, zero value otherwise.

### GetUseDdnsPrincipalSecurityOk

`func (o *View) GetUseDdnsPrincipalSecurityOk() (*bool, bool)`

GetUseDdnsPrincipalSecurityOk returns a tuple with the UseDdnsPrincipalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsPrincipalSecurity

`func (o *View) SetUseDdnsPrincipalSecurity(v bool)`

SetUseDdnsPrincipalSecurity sets UseDdnsPrincipalSecurity field to given value.

### HasUseDdnsPrincipalSecurity

`func (o *View) HasUseDdnsPrincipalSecurity() bool`

HasUseDdnsPrincipalSecurity returns a boolean if a field has been set.

### GetUseDdnsRestrictProtected

`func (o *View) GetUseDdnsRestrictProtected() bool`

GetUseDdnsRestrictProtected returns the UseDdnsRestrictProtected field if non-nil, zero value otherwise.

### GetUseDdnsRestrictProtectedOk

`func (o *View) GetUseDdnsRestrictProtectedOk() (*bool, bool)`

GetUseDdnsRestrictProtectedOk returns a tuple with the UseDdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsRestrictProtected

`func (o *View) SetUseDdnsRestrictProtected(v bool)`

SetUseDdnsRestrictProtected sets UseDdnsRestrictProtected field to given value.

### HasUseDdnsRestrictProtected

`func (o *View) HasUseDdnsRestrictProtected() bool`

HasUseDdnsRestrictProtected returns a boolean if a field has been set.

### GetUseDdnsRestrictStatic

`func (o *View) GetUseDdnsRestrictStatic() bool`

GetUseDdnsRestrictStatic returns the UseDdnsRestrictStatic field if non-nil, zero value otherwise.

### GetUseDdnsRestrictStaticOk

`func (o *View) GetUseDdnsRestrictStaticOk() (*bool, bool)`

GetUseDdnsRestrictStaticOk returns a tuple with the UseDdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsRestrictStatic

`func (o *View) SetUseDdnsRestrictStatic(v bool)`

SetUseDdnsRestrictStatic sets UseDdnsRestrictStatic field to given value.

### HasUseDdnsRestrictStatic

`func (o *View) HasUseDdnsRestrictStatic() bool`

HasUseDdnsRestrictStatic returns a boolean if a field has been set.

### GetUseDns64

`func (o *View) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *View) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *View) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.

### HasUseDns64

`func (o *View) HasUseDns64() bool`

HasUseDns64 returns a boolean if a field has been set.

### GetUseDnssec

`func (o *View) GetUseDnssec() bool`

GetUseDnssec returns the UseDnssec field if non-nil, zero value otherwise.

### GetUseDnssecOk

`func (o *View) GetUseDnssecOk() (*bool, bool)`

GetUseDnssecOk returns a tuple with the UseDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnssec

`func (o *View) SetUseDnssec(v bool)`

SetUseDnssec sets UseDnssec field to given value.

### HasUseDnssec

`func (o *View) HasUseDnssec() bool`

HasUseDnssec returns a boolean if a field has been set.

### GetUseEdnsUdpSize

`func (o *View) GetUseEdnsUdpSize() bool`

GetUseEdnsUdpSize returns the UseEdnsUdpSize field if non-nil, zero value otherwise.

### GetUseEdnsUdpSizeOk

`func (o *View) GetUseEdnsUdpSizeOk() (*bool, bool)`

GetUseEdnsUdpSizeOk returns a tuple with the UseEdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEdnsUdpSize

`func (o *View) SetUseEdnsUdpSize(v bool)`

SetUseEdnsUdpSize sets UseEdnsUdpSize field to given value.

### HasUseEdnsUdpSize

`func (o *View) HasUseEdnsUdpSize() bool`

HasUseEdnsUdpSize returns a boolean if a field has been set.

### GetUseFilterAaaa

`func (o *View) GetUseFilterAaaa() bool`

GetUseFilterAaaa returns the UseFilterAaaa field if non-nil, zero value otherwise.

### GetUseFilterAaaaOk

`func (o *View) GetUseFilterAaaaOk() (*bool, bool)`

GetUseFilterAaaaOk returns a tuple with the UseFilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFilterAaaa

`func (o *View) SetUseFilterAaaa(v bool)`

SetUseFilterAaaa sets UseFilterAaaa field to given value.

### HasUseFilterAaaa

`func (o *View) HasUseFilterAaaa() bool`

HasUseFilterAaaa returns a boolean if a field has been set.

### GetUseFixedRrsetOrderFqdns

`func (o *View) GetUseFixedRrsetOrderFqdns() bool`

GetUseFixedRrsetOrderFqdns returns the UseFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetUseFixedRrsetOrderFqdnsOk

`func (o *View) GetUseFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetUseFixedRrsetOrderFqdnsOk returns a tuple with the UseFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedRrsetOrderFqdns

`func (o *View) SetUseFixedRrsetOrderFqdns(v bool)`

SetUseFixedRrsetOrderFqdns sets UseFixedRrsetOrderFqdns field to given value.

### HasUseFixedRrsetOrderFqdns

`func (o *View) HasUseFixedRrsetOrderFqdns() bool`

HasUseFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetUseForwarders

`func (o *View) GetUseForwarders() bool`

GetUseForwarders returns the UseForwarders field if non-nil, zero value otherwise.

### GetUseForwardersOk

`func (o *View) GetUseForwardersOk() (*bool, bool)`

GetUseForwardersOk returns a tuple with the UseForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwarders

`func (o *View) SetUseForwarders(v bool)`

SetUseForwarders sets UseForwarders field to given value.

### HasUseForwarders

`func (o *View) HasUseForwarders() bool`

HasUseForwarders returns a boolean if a field has been set.

### GetUseMaxCacheTtl

`func (o *View) GetUseMaxCacheTtl() bool`

GetUseMaxCacheTtl returns the UseMaxCacheTtl field if non-nil, zero value otherwise.

### GetUseMaxCacheTtlOk

`func (o *View) GetUseMaxCacheTtlOk() (*bool, bool)`

GetUseMaxCacheTtlOk returns a tuple with the UseMaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxCacheTtl

`func (o *View) SetUseMaxCacheTtl(v bool)`

SetUseMaxCacheTtl sets UseMaxCacheTtl field to given value.

### HasUseMaxCacheTtl

`func (o *View) HasUseMaxCacheTtl() bool`

HasUseMaxCacheTtl returns a boolean if a field has been set.

### GetUseMaxNcacheTtl

`func (o *View) GetUseMaxNcacheTtl() bool`

GetUseMaxNcacheTtl returns the UseMaxNcacheTtl field if non-nil, zero value otherwise.

### GetUseMaxNcacheTtlOk

`func (o *View) GetUseMaxNcacheTtlOk() (*bool, bool)`

GetUseMaxNcacheTtlOk returns a tuple with the UseMaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxNcacheTtl

`func (o *View) SetUseMaxNcacheTtl(v bool)`

SetUseMaxNcacheTtl sets UseMaxNcacheTtl field to given value.

### HasUseMaxNcacheTtl

`func (o *View) HasUseMaxNcacheTtl() bool`

HasUseMaxNcacheTtl returns a boolean if a field has been set.

### GetUseMaxUdpSize

`func (o *View) GetUseMaxUdpSize() bool`

GetUseMaxUdpSize returns the UseMaxUdpSize field if non-nil, zero value otherwise.

### GetUseMaxUdpSizeOk

`func (o *View) GetUseMaxUdpSizeOk() (*bool, bool)`

GetUseMaxUdpSizeOk returns a tuple with the UseMaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxUdpSize

`func (o *View) SetUseMaxUdpSize(v bool)`

SetUseMaxUdpSize sets UseMaxUdpSize field to given value.

### HasUseMaxUdpSize

`func (o *View) HasUseMaxUdpSize() bool`

HasUseMaxUdpSize returns a boolean if a field has been set.

### GetUseNxdomainRedirect

`func (o *View) GetUseNxdomainRedirect() bool`

GetUseNxdomainRedirect returns the UseNxdomainRedirect field if non-nil, zero value otherwise.

### GetUseNxdomainRedirectOk

`func (o *View) GetUseNxdomainRedirectOk() (*bool, bool)`

GetUseNxdomainRedirectOk returns a tuple with the UseNxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNxdomainRedirect

`func (o *View) SetUseNxdomainRedirect(v bool)`

SetUseNxdomainRedirect sets UseNxdomainRedirect field to given value.

### HasUseNxdomainRedirect

`func (o *View) HasUseNxdomainRedirect() bool`

HasUseNxdomainRedirect returns a boolean if a field has been set.

### GetUseRecursion

`func (o *View) GetUseRecursion() bool`

GetUseRecursion returns the UseRecursion field if non-nil, zero value otherwise.

### GetUseRecursionOk

`func (o *View) GetUseRecursionOk() (*bool, bool)`

GetUseRecursionOk returns a tuple with the UseRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecursion

`func (o *View) SetUseRecursion(v bool)`

SetUseRecursion sets UseRecursion field to given value.

### HasUseRecursion

`func (o *View) HasUseRecursion() bool`

HasUseRecursion returns a boolean if a field has been set.

### GetUseResponseRateLimiting

`func (o *View) GetUseResponseRateLimiting() bool`

GetUseResponseRateLimiting returns the UseResponseRateLimiting field if non-nil, zero value otherwise.

### GetUseResponseRateLimitingOk

`func (o *View) GetUseResponseRateLimitingOk() (*bool, bool)`

GetUseResponseRateLimitingOk returns a tuple with the UseResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponseRateLimiting

`func (o *View) SetUseResponseRateLimiting(v bool)`

SetUseResponseRateLimiting sets UseResponseRateLimiting field to given value.

### HasUseResponseRateLimiting

`func (o *View) HasUseResponseRateLimiting() bool`

HasUseResponseRateLimiting returns a boolean if a field has been set.

### GetUseRootNameServer

`func (o *View) GetUseRootNameServer() bool`

GetUseRootNameServer returns the UseRootNameServer field if non-nil, zero value otherwise.

### GetUseRootNameServerOk

`func (o *View) GetUseRootNameServerOk() (*bool, bool)`

GetUseRootNameServerOk returns a tuple with the UseRootNameServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootNameServer

`func (o *View) SetUseRootNameServer(v bool)`

SetUseRootNameServer sets UseRootNameServer field to given value.

### HasUseRootNameServer

`func (o *View) HasUseRootNameServer() bool`

HasUseRootNameServer returns a boolean if a field has been set.

### GetUseRpzDropIpRule

`func (o *View) GetUseRpzDropIpRule() bool`

GetUseRpzDropIpRule returns the UseRpzDropIpRule field if non-nil, zero value otherwise.

### GetUseRpzDropIpRuleOk

`func (o *View) GetUseRpzDropIpRuleOk() (*bool, bool)`

GetUseRpzDropIpRuleOk returns a tuple with the UseRpzDropIpRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzDropIpRule

`func (o *View) SetUseRpzDropIpRule(v bool)`

SetUseRpzDropIpRule sets UseRpzDropIpRule field to given value.

### HasUseRpzDropIpRule

`func (o *View) HasUseRpzDropIpRule() bool`

HasUseRpzDropIpRule returns a boolean if a field has been set.

### GetUseRpzQnameWaitRecurse

`func (o *View) GetUseRpzQnameWaitRecurse() bool`

GetUseRpzQnameWaitRecurse returns the UseRpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetUseRpzQnameWaitRecurseOk

`func (o *View) GetUseRpzQnameWaitRecurseOk() (*bool, bool)`

GetUseRpzQnameWaitRecurseOk returns a tuple with the UseRpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzQnameWaitRecurse

`func (o *View) SetUseRpzQnameWaitRecurse(v bool)`

SetUseRpzQnameWaitRecurse sets UseRpzQnameWaitRecurse field to given value.

### HasUseRpzQnameWaitRecurse

`func (o *View) HasUseRpzQnameWaitRecurse() bool`

HasUseRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetUseScavengingSettings

`func (o *View) GetUseScavengingSettings() bool`

GetUseScavengingSettings returns the UseScavengingSettings field if non-nil, zero value otherwise.

### GetUseScavengingSettingsOk

`func (o *View) GetUseScavengingSettingsOk() (*bool, bool)`

GetUseScavengingSettingsOk returns a tuple with the UseScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseScavengingSettings

`func (o *View) SetUseScavengingSettings(v bool)`

SetUseScavengingSettings sets UseScavengingSettings field to given value.

### HasUseScavengingSettings

`func (o *View) HasUseScavengingSettings() bool`

HasUseScavengingSettings returns a boolean if a field has been set.

### GetUseSortlist

`func (o *View) GetUseSortlist() bool`

GetUseSortlist returns the UseSortlist field if non-nil, zero value otherwise.

### GetUseSortlistOk

`func (o *View) GetUseSortlistOk() (*bool, bool)`

GetUseSortlistOk returns a tuple with the UseSortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSortlist

`func (o *View) SetUseSortlist(v bool)`

SetUseSortlist sets UseSortlist field to given value.

### HasUseSortlist

`func (o *View) HasUseSortlist() bool`

HasUseSortlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


