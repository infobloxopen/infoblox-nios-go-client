# GetViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**View**](View.md) |  | [optional] 

## Methods

### NewGetViewResponse

`func NewGetViewResponse() *GetViewResponse`

NewGetViewResponse instantiates a new GetViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewResponseWithDefaults

`func NewGetViewResponseWithDefaults() *GetViewResponse`

NewGetViewResponseWithDefaults instantiates a new GetViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetViewResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetViewResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetViewResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetViewResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBlacklistAction

`func (o *GetViewResponse) GetBlacklistAction() string`

GetBlacklistAction returns the BlacklistAction field if non-nil, zero value otherwise.

### GetBlacklistActionOk

`func (o *GetViewResponse) GetBlacklistActionOk() (*string, bool)`

GetBlacklistActionOk returns a tuple with the BlacklistAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistAction

`func (o *GetViewResponse) SetBlacklistAction(v string)`

SetBlacklistAction sets BlacklistAction field to given value.

### HasBlacklistAction

`func (o *GetViewResponse) HasBlacklistAction() bool`

HasBlacklistAction returns a boolean if a field has been set.

### GetBlacklistLogQuery

`func (o *GetViewResponse) GetBlacklistLogQuery() bool`

GetBlacklistLogQuery returns the BlacklistLogQuery field if non-nil, zero value otherwise.

### GetBlacklistLogQueryOk

`func (o *GetViewResponse) GetBlacklistLogQueryOk() (*bool, bool)`

GetBlacklistLogQueryOk returns a tuple with the BlacklistLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistLogQuery

`func (o *GetViewResponse) SetBlacklistLogQuery(v bool)`

SetBlacklistLogQuery sets BlacklistLogQuery field to given value.

### HasBlacklistLogQuery

`func (o *GetViewResponse) HasBlacklistLogQuery() bool`

HasBlacklistLogQuery returns a boolean if a field has been set.

### GetBlacklistRedirectAddresses

`func (o *GetViewResponse) GetBlacklistRedirectAddresses() []string`

GetBlacklistRedirectAddresses returns the BlacklistRedirectAddresses field if non-nil, zero value otherwise.

### GetBlacklistRedirectAddressesOk

`func (o *GetViewResponse) GetBlacklistRedirectAddressesOk() (*[]string, bool)`

GetBlacklistRedirectAddressesOk returns a tuple with the BlacklistRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectAddresses

`func (o *GetViewResponse) SetBlacklistRedirectAddresses(v []string)`

SetBlacklistRedirectAddresses sets BlacklistRedirectAddresses field to given value.

### HasBlacklistRedirectAddresses

`func (o *GetViewResponse) HasBlacklistRedirectAddresses() bool`

HasBlacklistRedirectAddresses returns a boolean if a field has been set.

### GetBlacklistRedirectTtl

`func (o *GetViewResponse) GetBlacklistRedirectTtl() int64`

GetBlacklistRedirectTtl returns the BlacklistRedirectTtl field if non-nil, zero value otherwise.

### GetBlacklistRedirectTtlOk

`func (o *GetViewResponse) GetBlacklistRedirectTtlOk() (*int64, bool)`

GetBlacklistRedirectTtlOk returns a tuple with the BlacklistRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectTtl

`func (o *GetViewResponse) SetBlacklistRedirectTtl(v int64)`

SetBlacklistRedirectTtl sets BlacklistRedirectTtl field to given value.

### HasBlacklistRedirectTtl

`func (o *GetViewResponse) HasBlacklistRedirectTtl() bool`

HasBlacklistRedirectTtl returns a boolean if a field has been set.

### GetBlacklistRulesets

`func (o *GetViewResponse) GetBlacklistRulesets() []string`

GetBlacklistRulesets returns the BlacklistRulesets field if non-nil, zero value otherwise.

### GetBlacklistRulesetsOk

`func (o *GetViewResponse) GetBlacklistRulesetsOk() (*[]string, bool)`

GetBlacklistRulesetsOk returns a tuple with the BlacklistRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRulesets

`func (o *GetViewResponse) SetBlacklistRulesets(v []string)`

SetBlacklistRulesets sets BlacklistRulesets field to given value.

### HasBlacklistRulesets

`func (o *GetViewResponse) HasBlacklistRulesets() bool`

HasBlacklistRulesets returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetViewResponse) GetCloudInfo() ViewCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetViewResponse) GetCloudInfoOk() (*ViewCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetViewResponse) SetCloudInfo(v ViewCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetViewResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetViewResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetViewResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetViewResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetViewResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCustomRootNameServers

`func (o *GetViewResponse) GetCustomRootNameServers() []ViewCustomRootNameServers`

GetCustomRootNameServers returns the CustomRootNameServers field if non-nil, zero value otherwise.

### GetCustomRootNameServersOk

`func (o *GetViewResponse) GetCustomRootNameServersOk() (*[]ViewCustomRootNameServers, bool)`

GetCustomRootNameServersOk returns a tuple with the CustomRootNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNameServers

`func (o *GetViewResponse) SetCustomRootNameServers(v []ViewCustomRootNameServers)`

SetCustomRootNameServers sets CustomRootNameServers field to given value.

### HasCustomRootNameServers

`func (o *GetViewResponse) HasCustomRootNameServers() bool`

HasCustomRootNameServers returns a boolean if a field has been set.

### GetDdnsForceCreationTimestampUpdate

`func (o *GetViewResponse) GetDdnsForceCreationTimestampUpdate() bool`

GetDdnsForceCreationTimestampUpdate returns the DdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetDdnsForceCreationTimestampUpdateOk

`func (o *GetViewResponse) GetDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetDdnsForceCreationTimestampUpdateOk returns a tuple with the DdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsForceCreationTimestampUpdate

`func (o *GetViewResponse) SetDdnsForceCreationTimestampUpdate(v bool)`

SetDdnsForceCreationTimestampUpdate sets DdnsForceCreationTimestampUpdate field to given value.

### HasDdnsForceCreationTimestampUpdate

`func (o *GetViewResponse) HasDdnsForceCreationTimestampUpdate() bool`

HasDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetDdnsPrincipalGroup

`func (o *GetViewResponse) GetDdnsPrincipalGroup() string`

GetDdnsPrincipalGroup returns the DdnsPrincipalGroup field if non-nil, zero value otherwise.

### GetDdnsPrincipalGroupOk

`func (o *GetViewResponse) GetDdnsPrincipalGroupOk() (*string, bool)`

GetDdnsPrincipalGroupOk returns a tuple with the DdnsPrincipalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalGroup

`func (o *GetViewResponse) SetDdnsPrincipalGroup(v string)`

SetDdnsPrincipalGroup sets DdnsPrincipalGroup field to given value.

### HasDdnsPrincipalGroup

`func (o *GetViewResponse) HasDdnsPrincipalGroup() bool`

HasDdnsPrincipalGroup returns a boolean if a field has been set.

### GetDdnsPrincipalTracking

`func (o *GetViewResponse) GetDdnsPrincipalTracking() bool`

GetDdnsPrincipalTracking returns the DdnsPrincipalTracking field if non-nil, zero value otherwise.

### GetDdnsPrincipalTrackingOk

`func (o *GetViewResponse) GetDdnsPrincipalTrackingOk() (*bool, bool)`

GetDdnsPrincipalTrackingOk returns a tuple with the DdnsPrincipalTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalTracking

`func (o *GetViewResponse) SetDdnsPrincipalTracking(v bool)`

SetDdnsPrincipalTracking sets DdnsPrincipalTracking field to given value.

### HasDdnsPrincipalTracking

`func (o *GetViewResponse) HasDdnsPrincipalTracking() bool`

HasDdnsPrincipalTracking returns a boolean if a field has been set.

### GetDdnsRestrictPatterns

`func (o *GetViewResponse) GetDdnsRestrictPatterns() bool`

GetDdnsRestrictPatterns returns the DdnsRestrictPatterns field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsOk

`func (o *GetViewResponse) GetDdnsRestrictPatternsOk() (*bool, bool)`

GetDdnsRestrictPatternsOk returns a tuple with the DdnsRestrictPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatterns

`func (o *GetViewResponse) SetDdnsRestrictPatterns(v bool)`

SetDdnsRestrictPatterns sets DdnsRestrictPatterns field to given value.

### HasDdnsRestrictPatterns

`func (o *GetViewResponse) HasDdnsRestrictPatterns() bool`

HasDdnsRestrictPatterns returns a boolean if a field has been set.

### GetDdnsRestrictPatternsList

`func (o *GetViewResponse) GetDdnsRestrictPatternsList() []string`

GetDdnsRestrictPatternsList returns the DdnsRestrictPatternsList field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsListOk

`func (o *GetViewResponse) GetDdnsRestrictPatternsListOk() (*[]string, bool)`

GetDdnsRestrictPatternsListOk returns a tuple with the DdnsRestrictPatternsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatternsList

`func (o *GetViewResponse) SetDdnsRestrictPatternsList(v []string)`

SetDdnsRestrictPatternsList sets DdnsRestrictPatternsList field to given value.

### HasDdnsRestrictPatternsList

`func (o *GetViewResponse) HasDdnsRestrictPatternsList() bool`

HasDdnsRestrictPatternsList returns a boolean if a field has been set.

### GetDdnsRestrictProtected

`func (o *GetViewResponse) GetDdnsRestrictProtected() bool`

GetDdnsRestrictProtected returns the DdnsRestrictProtected field if non-nil, zero value otherwise.

### GetDdnsRestrictProtectedOk

`func (o *GetViewResponse) GetDdnsRestrictProtectedOk() (*bool, bool)`

GetDdnsRestrictProtectedOk returns a tuple with the DdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictProtected

`func (o *GetViewResponse) SetDdnsRestrictProtected(v bool)`

SetDdnsRestrictProtected sets DdnsRestrictProtected field to given value.

### HasDdnsRestrictProtected

`func (o *GetViewResponse) HasDdnsRestrictProtected() bool`

HasDdnsRestrictProtected returns a boolean if a field has been set.

### GetDdnsRestrictSecure

`func (o *GetViewResponse) GetDdnsRestrictSecure() bool`

GetDdnsRestrictSecure returns the DdnsRestrictSecure field if non-nil, zero value otherwise.

### GetDdnsRestrictSecureOk

`func (o *GetViewResponse) GetDdnsRestrictSecureOk() (*bool, bool)`

GetDdnsRestrictSecureOk returns a tuple with the DdnsRestrictSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictSecure

`func (o *GetViewResponse) SetDdnsRestrictSecure(v bool)`

SetDdnsRestrictSecure sets DdnsRestrictSecure field to given value.

### HasDdnsRestrictSecure

`func (o *GetViewResponse) HasDdnsRestrictSecure() bool`

HasDdnsRestrictSecure returns a boolean if a field has been set.

### GetDdnsRestrictStatic

`func (o *GetViewResponse) GetDdnsRestrictStatic() bool`

GetDdnsRestrictStatic returns the DdnsRestrictStatic field if non-nil, zero value otherwise.

### GetDdnsRestrictStaticOk

`func (o *GetViewResponse) GetDdnsRestrictStaticOk() (*bool, bool)`

GetDdnsRestrictStaticOk returns a tuple with the DdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictStatic

`func (o *GetViewResponse) SetDdnsRestrictStatic(v bool)`

SetDdnsRestrictStatic sets DdnsRestrictStatic field to given value.

### HasDdnsRestrictStatic

`func (o *GetViewResponse) HasDdnsRestrictStatic() bool`

HasDdnsRestrictStatic returns a boolean if a field has been set.

### GetDisable

`func (o *GetViewResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetViewResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetViewResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetViewResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDns64Enabled

`func (o *GetViewResponse) GetDns64Enabled() bool`

GetDns64Enabled returns the Dns64Enabled field if non-nil, zero value otherwise.

### GetDns64EnabledOk

`func (o *GetViewResponse) GetDns64EnabledOk() (*bool, bool)`

GetDns64EnabledOk returns a tuple with the Dns64Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Enabled

`func (o *GetViewResponse) SetDns64Enabled(v bool)`

SetDns64Enabled sets Dns64Enabled field to given value.

### HasDns64Enabled

`func (o *GetViewResponse) HasDns64Enabled() bool`

HasDns64Enabled returns a boolean if a field has been set.

### GetDns64Groups

`func (o *GetViewResponse) GetDns64Groups() []string`

GetDns64Groups returns the Dns64Groups field if non-nil, zero value otherwise.

### GetDns64GroupsOk

`func (o *GetViewResponse) GetDns64GroupsOk() (*[]string, bool)`

GetDns64GroupsOk returns a tuple with the Dns64Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Groups

`func (o *GetViewResponse) SetDns64Groups(v []string)`

SetDns64Groups sets Dns64Groups field to given value.

### HasDns64Groups

`func (o *GetViewResponse) HasDns64Groups() bool`

HasDns64Groups returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *GetViewResponse) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *GetViewResponse) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *GetViewResponse) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *GetViewResponse) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecExpiredSignaturesEnabled

`func (o *GetViewResponse) GetDnssecExpiredSignaturesEnabled() bool`

GetDnssecExpiredSignaturesEnabled returns the DnssecExpiredSignaturesEnabled field if non-nil, zero value otherwise.

### GetDnssecExpiredSignaturesEnabledOk

`func (o *GetViewResponse) GetDnssecExpiredSignaturesEnabledOk() (*bool, bool)`

GetDnssecExpiredSignaturesEnabledOk returns a tuple with the DnssecExpiredSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecExpiredSignaturesEnabled

`func (o *GetViewResponse) SetDnssecExpiredSignaturesEnabled(v bool)`

SetDnssecExpiredSignaturesEnabled sets DnssecExpiredSignaturesEnabled field to given value.

### HasDnssecExpiredSignaturesEnabled

`func (o *GetViewResponse) HasDnssecExpiredSignaturesEnabled() bool`

HasDnssecExpiredSignaturesEnabled returns a boolean if a field has been set.

### GetDnssecNegativeTrustAnchors

`func (o *GetViewResponse) GetDnssecNegativeTrustAnchors() []string`

GetDnssecNegativeTrustAnchors returns the DnssecNegativeTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecNegativeTrustAnchorsOk

`func (o *GetViewResponse) GetDnssecNegativeTrustAnchorsOk() (*[]string, bool)`

GetDnssecNegativeTrustAnchorsOk returns a tuple with the DnssecNegativeTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNegativeTrustAnchors

`func (o *GetViewResponse) SetDnssecNegativeTrustAnchors(v []string)`

SetDnssecNegativeTrustAnchors sets DnssecNegativeTrustAnchors field to given value.

### HasDnssecNegativeTrustAnchors

`func (o *GetViewResponse) HasDnssecNegativeTrustAnchors() bool`

HasDnssecNegativeTrustAnchors returns a boolean if a field has been set.

### GetDnssecTrustedKeys

`func (o *GetViewResponse) GetDnssecTrustedKeys() []ViewDnssecTrustedKeys`

GetDnssecTrustedKeys returns the DnssecTrustedKeys field if non-nil, zero value otherwise.

### GetDnssecTrustedKeysOk

`func (o *GetViewResponse) GetDnssecTrustedKeysOk() (*[]ViewDnssecTrustedKeys, bool)`

GetDnssecTrustedKeysOk returns a tuple with the DnssecTrustedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustedKeys

`func (o *GetViewResponse) SetDnssecTrustedKeys(v []ViewDnssecTrustedKeys)`

SetDnssecTrustedKeys sets DnssecTrustedKeys field to given value.

### HasDnssecTrustedKeys

`func (o *GetViewResponse) HasDnssecTrustedKeys() bool`

HasDnssecTrustedKeys returns a boolean if a field has been set.

### GetDnssecValidationEnabled

`func (o *GetViewResponse) GetDnssecValidationEnabled() bool`

GetDnssecValidationEnabled returns the DnssecValidationEnabled field if non-nil, zero value otherwise.

### GetDnssecValidationEnabledOk

`func (o *GetViewResponse) GetDnssecValidationEnabledOk() (*bool, bool)`

GetDnssecValidationEnabledOk returns a tuple with the DnssecValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationEnabled

`func (o *GetViewResponse) SetDnssecValidationEnabled(v bool)`

SetDnssecValidationEnabled sets DnssecValidationEnabled field to given value.

### HasDnssecValidationEnabled

`func (o *GetViewResponse) HasDnssecValidationEnabled() bool`

HasDnssecValidationEnabled returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *GetViewResponse) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *GetViewResponse) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *GetViewResponse) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *GetViewResponse) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetEnableBlacklist

`func (o *GetViewResponse) GetEnableBlacklist() bool`

GetEnableBlacklist returns the EnableBlacklist field if non-nil, zero value otherwise.

### GetEnableBlacklistOk

`func (o *GetViewResponse) GetEnableBlacklistOk() (*bool, bool)`

GetEnableBlacklistOk returns a tuple with the EnableBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlacklist

`func (o *GetViewResponse) SetEnableBlacklist(v bool)`

SetEnableBlacklist sets EnableBlacklist field to given value.

### HasEnableBlacklist

`func (o *GetViewResponse) HasEnableBlacklist() bool`

HasEnableBlacklist returns a boolean if a field has been set.

### GetEnableFixedRrsetOrderFqdns

`func (o *GetViewResponse) GetEnableFixedRrsetOrderFqdns() bool`

GetEnableFixedRrsetOrderFqdns returns the EnableFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetEnableFixedRrsetOrderFqdnsOk

`func (o *GetViewResponse) GetEnableFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetEnableFixedRrsetOrderFqdnsOk returns a tuple with the EnableFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixedRrsetOrderFqdns

`func (o *GetViewResponse) SetEnableFixedRrsetOrderFqdns(v bool)`

SetEnableFixedRrsetOrderFqdns sets EnableFixedRrsetOrderFqdns field to given value.

### HasEnableFixedRrsetOrderFqdns

`func (o *GetViewResponse) HasEnableFixedRrsetOrderFqdns() bool`

HasEnableFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetEnableMatchRecursiveOnly

`func (o *GetViewResponse) GetEnableMatchRecursiveOnly() bool`

GetEnableMatchRecursiveOnly returns the EnableMatchRecursiveOnly field if non-nil, zero value otherwise.

### GetEnableMatchRecursiveOnlyOk

`func (o *GetViewResponse) GetEnableMatchRecursiveOnlyOk() (*bool, bool)`

GetEnableMatchRecursiveOnlyOk returns a tuple with the EnableMatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMatchRecursiveOnly

`func (o *GetViewResponse) SetEnableMatchRecursiveOnly(v bool)`

SetEnableMatchRecursiveOnly sets EnableMatchRecursiveOnly field to given value.

### HasEnableMatchRecursiveOnly

`func (o *GetViewResponse) HasEnableMatchRecursiveOnly() bool`

HasEnableMatchRecursiveOnly returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetViewResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetViewResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetViewResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetViewResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFilterAaaa

`func (o *GetViewResponse) GetFilterAaaa() string`

GetFilterAaaa returns the FilterAaaa field if non-nil, zero value otherwise.

### GetFilterAaaaOk

`func (o *GetViewResponse) GetFilterAaaaOk() (*string, bool)`

GetFilterAaaaOk returns a tuple with the FilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaa

`func (o *GetViewResponse) SetFilterAaaa(v string)`

SetFilterAaaa sets FilterAaaa field to given value.

### HasFilterAaaa

`func (o *GetViewResponse) HasFilterAaaa() bool`

HasFilterAaaa returns a boolean if a field has been set.

### GetFilterAaaaList

`func (o *GetViewResponse) GetFilterAaaaList() []ViewFilterAaaaList`

GetFilterAaaaList returns the FilterAaaaList field if non-nil, zero value otherwise.

### GetFilterAaaaListOk

`func (o *GetViewResponse) GetFilterAaaaListOk() (*[]ViewFilterAaaaList, bool)`

GetFilterAaaaListOk returns a tuple with the FilterAaaaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaList

`func (o *GetViewResponse) SetFilterAaaaList(v []ViewFilterAaaaList)`

SetFilterAaaaList sets FilterAaaaList field to given value.

### HasFilterAaaaList

`func (o *GetViewResponse) HasFilterAaaaList() bool`

HasFilterAaaaList returns a boolean if a field has been set.

### GetFixedRrsetOrderFqdns

`func (o *GetViewResponse) GetFixedRrsetOrderFqdns() []ViewFixedRrsetOrderFqdns`

GetFixedRrsetOrderFqdns returns the FixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetFixedRrsetOrderFqdnsOk

`func (o *GetViewResponse) GetFixedRrsetOrderFqdnsOk() (*[]ViewFixedRrsetOrderFqdns, bool)`

GetFixedRrsetOrderFqdnsOk returns a tuple with the FixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedRrsetOrderFqdns

`func (o *GetViewResponse) SetFixedRrsetOrderFqdns(v []ViewFixedRrsetOrderFqdns)`

SetFixedRrsetOrderFqdns sets FixedRrsetOrderFqdns field to given value.

### HasFixedRrsetOrderFqdns

`func (o *GetViewResponse) HasFixedRrsetOrderFqdns() bool`

HasFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetForwardOnly

`func (o *GetViewResponse) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *GetViewResponse) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *GetViewResponse) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *GetViewResponse) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetForwarders

`func (o *GetViewResponse) GetForwarders() []string`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *GetViewResponse) GetForwardersOk() (*[]string, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *GetViewResponse) SetForwarders(v []string)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *GetViewResponse) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetViewResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetViewResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetViewResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetViewResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastQueriedAcl

`func (o *GetViewResponse) GetLastQueriedAcl() []ViewLastQueriedAcl`

GetLastQueriedAcl returns the LastQueriedAcl field if non-nil, zero value otherwise.

### GetLastQueriedAclOk

`func (o *GetViewResponse) GetLastQueriedAclOk() (*[]ViewLastQueriedAcl, bool)`

GetLastQueriedAclOk returns a tuple with the LastQueriedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueriedAcl

`func (o *GetViewResponse) SetLastQueriedAcl(v []ViewLastQueriedAcl)`

SetLastQueriedAcl sets LastQueriedAcl field to given value.

### HasLastQueriedAcl

`func (o *GetViewResponse) HasLastQueriedAcl() bool`

HasLastQueriedAcl returns a boolean if a field has been set.

### GetMatchClients

`func (o *GetViewResponse) GetMatchClients() []ViewMatchClients`

GetMatchClients returns the MatchClients field if non-nil, zero value otherwise.

### GetMatchClientsOk

`func (o *GetViewResponse) GetMatchClientsOk() (*[]ViewMatchClients, bool)`

GetMatchClientsOk returns a tuple with the MatchClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClients

`func (o *GetViewResponse) SetMatchClients(v []ViewMatchClients)`

SetMatchClients sets MatchClients field to given value.

### HasMatchClients

`func (o *GetViewResponse) HasMatchClients() bool`

HasMatchClients returns a boolean if a field has been set.

### GetMatchDestinations

`func (o *GetViewResponse) GetMatchDestinations() []ViewMatchDestinations`

GetMatchDestinations returns the MatchDestinations field if non-nil, zero value otherwise.

### GetMatchDestinationsOk

`func (o *GetViewResponse) GetMatchDestinationsOk() (*[]ViewMatchDestinations, bool)`

GetMatchDestinationsOk returns a tuple with the MatchDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDestinations

`func (o *GetViewResponse) SetMatchDestinations(v []ViewMatchDestinations)`

SetMatchDestinations sets MatchDestinations field to given value.

### HasMatchDestinations

`func (o *GetViewResponse) HasMatchDestinations() bool`

HasMatchDestinations returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *GetViewResponse) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *GetViewResponse) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *GetViewResponse) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *GetViewResponse) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNcacheTtl

`func (o *GetViewResponse) GetMaxNcacheTtl() int64`

GetMaxNcacheTtl returns the MaxNcacheTtl field if non-nil, zero value otherwise.

### GetMaxNcacheTtlOk

`func (o *GetViewResponse) GetMaxNcacheTtlOk() (*int64, bool)`

GetMaxNcacheTtlOk returns a tuple with the MaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNcacheTtl

`func (o *GetViewResponse) SetMaxNcacheTtl(v int64)`

SetMaxNcacheTtl sets MaxNcacheTtl field to given value.

### HasMaxNcacheTtl

`func (o *GetViewResponse) HasMaxNcacheTtl() bool`

HasMaxNcacheTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *GetViewResponse) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *GetViewResponse) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *GetViewResponse) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *GetViewResponse) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetName

`func (o *GetViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetViewResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetViewResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetViewResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetViewResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetViewResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetViewResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *GetViewResponse) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *GetViewResponse) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *GetViewResponse) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *GetViewResponse) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetNxdomainLogQuery

`func (o *GetViewResponse) GetNxdomainLogQuery() bool`

GetNxdomainLogQuery returns the NxdomainLogQuery field if non-nil, zero value otherwise.

### GetNxdomainLogQueryOk

`func (o *GetViewResponse) GetNxdomainLogQueryOk() (*bool, bool)`

GetNxdomainLogQueryOk returns a tuple with the NxdomainLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainLogQuery

`func (o *GetViewResponse) SetNxdomainLogQuery(v bool)`

SetNxdomainLogQuery sets NxdomainLogQuery field to given value.

### HasNxdomainLogQuery

`func (o *GetViewResponse) HasNxdomainLogQuery() bool`

HasNxdomainLogQuery returns a boolean if a field has been set.

### GetNxdomainRedirect

`func (o *GetViewResponse) GetNxdomainRedirect() bool`

GetNxdomainRedirect returns the NxdomainRedirect field if non-nil, zero value otherwise.

### GetNxdomainRedirectOk

`func (o *GetViewResponse) GetNxdomainRedirectOk() (*bool, bool)`

GetNxdomainRedirectOk returns a tuple with the NxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirect

`func (o *GetViewResponse) SetNxdomainRedirect(v bool)`

SetNxdomainRedirect sets NxdomainRedirect field to given value.

### HasNxdomainRedirect

`func (o *GetViewResponse) HasNxdomainRedirect() bool`

HasNxdomainRedirect returns a boolean if a field has been set.

### GetNxdomainRedirectAddresses

`func (o *GetViewResponse) GetNxdomainRedirectAddresses() []string`

GetNxdomainRedirectAddresses returns the NxdomainRedirectAddresses field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesOk

`func (o *GetViewResponse) GetNxdomainRedirectAddressesOk() (*[]string, bool)`

GetNxdomainRedirectAddressesOk returns a tuple with the NxdomainRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddresses

`func (o *GetViewResponse) SetNxdomainRedirectAddresses(v []string)`

SetNxdomainRedirectAddresses sets NxdomainRedirectAddresses field to given value.

### HasNxdomainRedirectAddresses

`func (o *GetViewResponse) HasNxdomainRedirectAddresses() bool`

HasNxdomainRedirectAddresses returns a boolean if a field has been set.

### GetNxdomainRedirectAddressesV6

`func (o *GetViewResponse) GetNxdomainRedirectAddressesV6() []string`

GetNxdomainRedirectAddressesV6 returns the NxdomainRedirectAddressesV6 field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesV6Ok

`func (o *GetViewResponse) GetNxdomainRedirectAddressesV6Ok() (*[]string, bool)`

GetNxdomainRedirectAddressesV6Ok returns a tuple with the NxdomainRedirectAddressesV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddressesV6

`func (o *GetViewResponse) SetNxdomainRedirectAddressesV6(v []string)`

SetNxdomainRedirectAddressesV6 sets NxdomainRedirectAddressesV6 field to given value.

### HasNxdomainRedirectAddressesV6

`func (o *GetViewResponse) HasNxdomainRedirectAddressesV6() bool`

HasNxdomainRedirectAddressesV6 returns a boolean if a field has been set.

### GetNxdomainRedirectTtl

`func (o *GetViewResponse) GetNxdomainRedirectTtl() int64`

GetNxdomainRedirectTtl returns the NxdomainRedirectTtl field if non-nil, zero value otherwise.

### GetNxdomainRedirectTtlOk

`func (o *GetViewResponse) GetNxdomainRedirectTtlOk() (*int64, bool)`

GetNxdomainRedirectTtlOk returns a tuple with the NxdomainRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectTtl

`func (o *GetViewResponse) SetNxdomainRedirectTtl(v int64)`

SetNxdomainRedirectTtl sets NxdomainRedirectTtl field to given value.

### HasNxdomainRedirectTtl

`func (o *GetViewResponse) HasNxdomainRedirectTtl() bool`

HasNxdomainRedirectTtl returns a boolean if a field has been set.

### GetNxdomainRulesets

`func (o *GetViewResponse) GetNxdomainRulesets() []string`

GetNxdomainRulesets returns the NxdomainRulesets field if non-nil, zero value otherwise.

### GetNxdomainRulesetsOk

`func (o *GetViewResponse) GetNxdomainRulesetsOk() (*[]string, bool)`

GetNxdomainRulesetsOk returns a tuple with the NxdomainRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRulesets

`func (o *GetViewResponse) SetNxdomainRulesets(v []string)`

SetNxdomainRulesets sets NxdomainRulesets field to given value.

### HasNxdomainRulesets

`func (o *GetViewResponse) HasNxdomainRulesets() bool`

HasNxdomainRulesets returns a boolean if a field has been set.

### GetRecursion

`func (o *GetViewResponse) GetRecursion() bool`

GetRecursion returns the Recursion field if non-nil, zero value otherwise.

### GetRecursionOk

`func (o *GetViewResponse) GetRecursionOk() (*bool, bool)`

GetRecursionOk returns a tuple with the Recursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursion

`func (o *GetViewResponse) SetRecursion(v bool)`

SetRecursion sets Recursion field to given value.

### HasRecursion

`func (o *GetViewResponse) HasRecursion() bool`

HasRecursion returns a boolean if a field has been set.

### GetResponseRateLimiting

`func (o *GetViewResponse) GetResponseRateLimiting() ViewResponseRateLimiting`

GetResponseRateLimiting returns the ResponseRateLimiting field if non-nil, zero value otherwise.

### GetResponseRateLimitingOk

`func (o *GetViewResponse) GetResponseRateLimitingOk() (*ViewResponseRateLimiting, bool)`

GetResponseRateLimitingOk returns a tuple with the ResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRateLimiting

`func (o *GetViewResponse) SetResponseRateLimiting(v ViewResponseRateLimiting)`

SetResponseRateLimiting sets ResponseRateLimiting field to given value.

### HasResponseRateLimiting

`func (o *GetViewResponse) HasResponseRateLimiting() bool`

HasResponseRateLimiting returns a boolean if a field has been set.

### GetRootNameServerType

`func (o *GetViewResponse) GetRootNameServerType() string`

GetRootNameServerType returns the RootNameServerType field if non-nil, zero value otherwise.

### GetRootNameServerTypeOk

`func (o *GetViewResponse) GetRootNameServerTypeOk() (*string, bool)`

GetRootNameServerTypeOk returns a tuple with the RootNameServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNameServerType

`func (o *GetViewResponse) SetRootNameServerType(v string)`

SetRootNameServerType sets RootNameServerType field to given value.

### HasRootNameServerType

`func (o *GetViewResponse) HasRootNameServerType() bool`

HasRootNameServerType returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *GetViewResponse) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *GetViewResponse) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *GetViewResponse) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *GetViewResponse) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetViewResponse) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *GetViewResponse) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetViewResponse) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetViewResponse) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetViewResponse) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *GetViewResponse) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetViewResponse) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetViewResponse) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzQnameWaitRecurse

`func (o *GetViewResponse) GetRpzQnameWaitRecurse() bool`

GetRpzQnameWaitRecurse returns the RpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetRpzQnameWaitRecurseOk

`func (o *GetViewResponse) GetRpzQnameWaitRecurseOk() (*bool, bool)`

GetRpzQnameWaitRecurseOk returns a tuple with the RpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzQnameWaitRecurse

`func (o *GetViewResponse) SetRpzQnameWaitRecurse(v bool)`

SetRpzQnameWaitRecurse sets RpzQnameWaitRecurse field to given value.

### HasRpzQnameWaitRecurse

`func (o *GetViewResponse) HasRpzQnameWaitRecurse() bool`

HasRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetScavengingSettings

`func (o *GetViewResponse) GetScavengingSettings() ViewScavengingSettings`

GetScavengingSettings returns the ScavengingSettings field if non-nil, zero value otherwise.

### GetScavengingSettingsOk

`func (o *GetViewResponse) GetScavengingSettingsOk() (*ViewScavengingSettings, bool)`

GetScavengingSettingsOk returns a tuple with the ScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSettings

`func (o *GetViewResponse) SetScavengingSettings(v ViewScavengingSettings)`

SetScavengingSettings sets ScavengingSettings field to given value.

### HasScavengingSettings

`func (o *GetViewResponse) HasScavengingSettings() bool`

HasScavengingSettings returns a boolean if a field has been set.

### GetSortlist

`func (o *GetViewResponse) GetSortlist() []ViewSortlist`

GetSortlist returns the Sortlist field if non-nil, zero value otherwise.

### GetSortlistOk

`func (o *GetViewResponse) GetSortlistOk() (*[]ViewSortlist, bool)`

GetSortlistOk returns a tuple with the Sortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortlist

`func (o *GetViewResponse) SetSortlist(v []ViewSortlist)`

SetSortlist sets Sortlist field to given value.

### HasSortlist

`func (o *GetViewResponse) HasSortlist() bool`

HasSortlist returns a boolean if a field has been set.

### GetUseBlacklist

`func (o *GetViewResponse) GetUseBlacklist() bool`

GetUseBlacklist returns the UseBlacklist field if non-nil, zero value otherwise.

### GetUseBlacklistOk

`func (o *GetViewResponse) GetUseBlacklistOk() (*bool, bool)`

GetUseBlacklistOk returns a tuple with the UseBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlacklist

`func (o *GetViewResponse) SetUseBlacklist(v bool)`

SetUseBlacklist sets UseBlacklist field to given value.

### HasUseBlacklist

`func (o *GetViewResponse) HasUseBlacklist() bool`

HasUseBlacklist returns a boolean if a field has been set.

### GetUseDdnsForceCreationTimestampUpdate

`func (o *GetViewResponse) GetUseDdnsForceCreationTimestampUpdate() bool`

GetUseDdnsForceCreationTimestampUpdate returns the UseDdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetUseDdnsForceCreationTimestampUpdateOk

`func (o *GetViewResponse) GetUseDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetUseDdnsForceCreationTimestampUpdateOk returns a tuple with the UseDdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsForceCreationTimestampUpdate

`func (o *GetViewResponse) SetUseDdnsForceCreationTimestampUpdate(v bool)`

SetUseDdnsForceCreationTimestampUpdate sets UseDdnsForceCreationTimestampUpdate field to given value.

### HasUseDdnsForceCreationTimestampUpdate

`func (o *GetViewResponse) HasUseDdnsForceCreationTimestampUpdate() bool`

HasUseDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetUseDdnsPatternsRestriction

`func (o *GetViewResponse) GetUseDdnsPatternsRestriction() bool`

GetUseDdnsPatternsRestriction returns the UseDdnsPatternsRestriction field if non-nil, zero value otherwise.

### GetUseDdnsPatternsRestrictionOk

`func (o *GetViewResponse) GetUseDdnsPatternsRestrictionOk() (*bool, bool)`

GetUseDdnsPatternsRestrictionOk returns a tuple with the UseDdnsPatternsRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsPatternsRestriction

`func (o *GetViewResponse) SetUseDdnsPatternsRestriction(v bool)`

SetUseDdnsPatternsRestriction sets UseDdnsPatternsRestriction field to given value.

### HasUseDdnsPatternsRestriction

`func (o *GetViewResponse) HasUseDdnsPatternsRestriction() bool`

HasUseDdnsPatternsRestriction returns a boolean if a field has been set.

### GetUseDdnsPrincipalSecurity

`func (o *GetViewResponse) GetUseDdnsPrincipalSecurity() bool`

GetUseDdnsPrincipalSecurity returns the UseDdnsPrincipalSecurity field if non-nil, zero value otherwise.

### GetUseDdnsPrincipalSecurityOk

`func (o *GetViewResponse) GetUseDdnsPrincipalSecurityOk() (*bool, bool)`

GetUseDdnsPrincipalSecurityOk returns a tuple with the UseDdnsPrincipalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsPrincipalSecurity

`func (o *GetViewResponse) SetUseDdnsPrincipalSecurity(v bool)`

SetUseDdnsPrincipalSecurity sets UseDdnsPrincipalSecurity field to given value.

### HasUseDdnsPrincipalSecurity

`func (o *GetViewResponse) HasUseDdnsPrincipalSecurity() bool`

HasUseDdnsPrincipalSecurity returns a boolean if a field has been set.

### GetUseDdnsRestrictProtected

`func (o *GetViewResponse) GetUseDdnsRestrictProtected() bool`

GetUseDdnsRestrictProtected returns the UseDdnsRestrictProtected field if non-nil, zero value otherwise.

### GetUseDdnsRestrictProtectedOk

`func (o *GetViewResponse) GetUseDdnsRestrictProtectedOk() (*bool, bool)`

GetUseDdnsRestrictProtectedOk returns a tuple with the UseDdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsRestrictProtected

`func (o *GetViewResponse) SetUseDdnsRestrictProtected(v bool)`

SetUseDdnsRestrictProtected sets UseDdnsRestrictProtected field to given value.

### HasUseDdnsRestrictProtected

`func (o *GetViewResponse) HasUseDdnsRestrictProtected() bool`

HasUseDdnsRestrictProtected returns a boolean if a field has been set.

### GetUseDdnsRestrictStatic

`func (o *GetViewResponse) GetUseDdnsRestrictStatic() bool`

GetUseDdnsRestrictStatic returns the UseDdnsRestrictStatic field if non-nil, zero value otherwise.

### GetUseDdnsRestrictStaticOk

`func (o *GetViewResponse) GetUseDdnsRestrictStaticOk() (*bool, bool)`

GetUseDdnsRestrictStaticOk returns a tuple with the UseDdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsRestrictStatic

`func (o *GetViewResponse) SetUseDdnsRestrictStatic(v bool)`

SetUseDdnsRestrictStatic sets UseDdnsRestrictStatic field to given value.

### HasUseDdnsRestrictStatic

`func (o *GetViewResponse) HasUseDdnsRestrictStatic() bool`

HasUseDdnsRestrictStatic returns a boolean if a field has been set.

### GetUseDns64

`func (o *GetViewResponse) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *GetViewResponse) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *GetViewResponse) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.

### HasUseDns64

`func (o *GetViewResponse) HasUseDns64() bool`

HasUseDns64 returns a boolean if a field has been set.

### GetUseDnssec

`func (o *GetViewResponse) GetUseDnssec() bool`

GetUseDnssec returns the UseDnssec field if non-nil, zero value otherwise.

### GetUseDnssecOk

`func (o *GetViewResponse) GetUseDnssecOk() (*bool, bool)`

GetUseDnssecOk returns a tuple with the UseDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnssec

`func (o *GetViewResponse) SetUseDnssec(v bool)`

SetUseDnssec sets UseDnssec field to given value.

### HasUseDnssec

`func (o *GetViewResponse) HasUseDnssec() bool`

HasUseDnssec returns a boolean if a field has been set.

### GetUseEdnsUdpSize

`func (o *GetViewResponse) GetUseEdnsUdpSize() bool`

GetUseEdnsUdpSize returns the UseEdnsUdpSize field if non-nil, zero value otherwise.

### GetUseEdnsUdpSizeOk

`func (o *GetViewResponse) GetUseEdnsUdpSizeOk() (*bool, bool)`

GetUseEdnsUdpSizeOk returns a tuple with the UseEdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEdnsUdpSize

`func (o *GetViewResponse) SetUseEdnsUdpSize(v bool)`

SetUseEdnsUdpSize sets UseEdnsUdpSize field to given value.

### HasUseEdnsUdpSize

`func (o *GetViewResponse) HasUseEdnsUdpSize() bool`

HasUseEdnsUdpSize returns a boolean if a field has been set.

### GetUseFilterAaaa

`func (o *GetViewResponse) GetUseFilterAaaa() bool`

GetUseFilterAaaa returns the UseFilterAaaa field if non-nil, zero value otherwise.

### GetUseFilterAaaaOk

`func (o *GetViewResponse) GetUseFilterAaaaOk() (*bool, bool)`

GetUseFilterAaaaOk returns a tuple with the UseFilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFilterAaaa

`func (o *GetViewResponse) SetUseFilterAaaa(v bool)`

SetUseFilterAaaa sets UseFilterAaaa field to given value.

### HasUseFilterAaaa

`func (o *GetViewResponse) HasUseFilterAaaa() bool`

HasUseFilterAaaa returns a boolean if a field has been set.

### GetUseFixedRrsetOrderFqdns

`func (o *GetViewResponse) GetUseFixedRrsetOrderFqdns() bool`

GetUseFixedRrsetOrderFqdns returns the UseFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetUseFixedRrsetOrderFqdnsOk

`func (o *GetViewResponse) GetUseFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetUseFixedRrsetOrderFqdnsOk returns a tuple with the UseFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedRrsetOrderFqdns

`func (o *GetViewResponse) SetUseFixedRrsetOrderFqdns(v bool)`

SetUseFixedRrsetOrderFqdns sets UseFixedRrsetOrderFqdns field to given value.

### HasUseFixedRrsetOrderFqdns

`func (o *GetViewResponse) HasUseFixedRrsetOrderFqdns() bool`

HasUseFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetUseForwarders

`func (o *GetViewResponse) GetUseForwarders() bool`

GetUseForwarders returns the UseForwarders field if non-nil, zero value otherwise.

### GetUseForwardersOk

`func (o *GetViewResponse) GetUseForwardersOk() (*bool, bool)`

GetUseForwardersOk returns a tuple with the UseForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwarders

`func (o *GetViewResponse) SetUseForwarders(v bool)`

SetUseForwarders sets UseForwarders field to given value.

### HasUseForwarders

`func (o *GetViewResponse) HasUseForwarders() bool`

HasUseForwarders returns a boolean if a field has been set.

### GetUseMaxCacheTtl

`func (o *GetViewResponse) GetUseMaxCacheTtl() bool`

GetUseMaxCacheTtl returns the UseMaxCacheTtl field if non-nil, zero value otherwise.

### GetUseMaxCacheTtlOk

`func (o *GetViewResponse) GetUseMaxCacheTtlOk() (*bool, bool)`

GetUseMaxCacheTtlOk returns a tuple with the UseMaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxCacheTtl

`func (o *GetViewResponse) SetUseMaxCacheTtl(v bool)`

SetUseMaxCacheTtl sets UseMaxCacheTtl field to given value.

### HasUseMaxCacheTtl

`func (o *GetViewResponse) HasUseMaxCacheTtl() bool`

HasUseMaxCacheTtl returns a boolean if a field has been set.

### GetUseMaxNcacheTtl

`func (o *GetViewResponse) GetUseMaxNcacheTtl() bool`

GetUseMaxNcacheTtl returns the UseMaxNcacheTtl field if non-nil, zero value otherwise.

### GetUseMaxNcacheTtlOk

`func (o *GetViewResponse) GetUseMaxNcacheTtlOk() (*bool, bool)`

GetUseMaxNcacheTtlOk returns a tuple with the UseMaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxNcacheTtl

`func (o *GetViewResponse) SetUseMaxNcacheTtl(v bool)`

SetUseMaxNcacheTtl sets UseMaxNcacheTtl field to given value.

### HasUseMaxNcacheTtl

`func (o *GetViewResponse) HasUseMaxNcacheTtl() bool`

HasUseMaxNcacheTtl returns a boolean if a field has been set.

### GetUseMaxUdpSize

`func (o *GetViewResponse) GetUseMaxUdpSize() bool`

GetUseMaxUdpSize returns the UseMaxUdpSize field if non-nil, zero value otherwise.

### GetUseMaxUdpSizeOk

`func (o *GetViewResponse) GetUseMaxUdpSizeOk() (*bool, bool)`

GetUseMaxUdpSizeOk returns a tuple with the UseMaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxUdpSize

`func (o *GetViewResponse) SetUseMaxUdpSize(v bool)`

SetUseMaxUdpSize sets UseMaxUdpSize field to given value.

### HasUseMaxUdpSize

`func (o *GetViewResponse) HasUseMaxUdpSize() bool`

HasUseMaxUdpSize returns a boolean if a field has been set.

### GetUseNxdomainRedirect

`func (o *GetViewResponse) GetUseNxdomainRedirect() bool`

GetUseNxdomainRedirect returns the UseNxdomainRedirect field if non-nil, zero value otherwise.

### GetUseNxdomainRedirectOk

`func (o *GetViewResponse) GetUseNxdomainRedirectOk() (*bool, bool)`

GetUseNxdomainRedirectOk returns a tuple with the UseNxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNxdomainRedirect

`func (o *GetViewResponse) SetUseNxdomainRedirect(v bool)`

SetUseNxdomainRedirect sets UseNxdomainRedirect field to given value.

### HasUseNxdomainRedirect

`func (o *GetViewResponse) HasUseNxdomainRedirect() bool`

HasUseNxdomainRedirect returns a boolean if a field has been set.

### GetUseRecursion

`func (o *GetViewResponse) GetUseRecursion() bool`

GetUseRecursion returns the UseRecursion field if non-nil, zero value otherwise.

### GetUseRecursionOk

`func (o *GetViewResponse) GetUseRecursionOk() (*bool, bool)`

GetUseRecursionOk returns a tuple with the UseRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecursion

`func (o *GetViewResponse) SetUseRecursion(v bool)`

SetUseRecursion sets UseRecursion field to given value.

### HasUseRecursion

`func (o *GetViewResponse) HasUseRecursion() bool`

HasUseRecursion returns a boolean if a field has been set.

### GetUseResponseRateLimiting

`func (o *GetViewResponse) GetUseResponseRateLimiting() bool`

GetUseResponseRateLimiting returns the UseResponseRateLimiting field if non-nil, zero value otherwise.

### GetUseResponseRateLimitingOk

`func (o *GetViewResponse) GetUseResponseRateLimitingOk() (*bool, bool)`

GetUseResponseRateLimitingOk returns a tuple with the UseResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponseRateLimiting

`func (o *GetViewResponse) SetUseResponseRateLimiting(v bool)`

SetUseResponseRateLimiting sets UseResponseRateLimiting field to given value.

### HasUseResponseRateLimiting

`func (o *GetViewResponse) HasUseResponseRateLimiting() bool`

HasUseResponseRateLimiting returns a boolean if a field has been set.

### GetUseRootNameServer

`func (o *GetViewResponse) GetUseRootNameServer() bool`

GetUseRootNameServer returns the UseRootNameServer field if non-nil, zero value otherwise.

### GetUseRootNameServerOk

`func (o *GetViewResponse) GetUseRootNameServerOk() (*bool, bool)`

GetUseRootNameServerOk returns a tuple with the UseRootNameServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootNameServer

`func (o *GetViewResponse) SetUseRootNameServer(v bool)`

SetUseRootNameServer sets UseRootNameServer field to given value.

### HasUseRootNameServer

`func (o *GetViewResponse) HasUseRootNameServer() bool`

HasUseRootNameServer returns a boolean if a field has been set.

### GetUseRpzDropIpRule

`func (o *GetViewResponse) GetUseRpzDropIpRule() bool`

GetUseRpzDropIpRule returns the UseRpzDropIpRule field if non-nil, zero value otherwise.

### GetUseRpzDropIpRuleOk

`func (o *GetViewResponse) GetUseRpzDropIpRuleOk() (*bool, bool)`

GetUseRpzDropIpRuleOk returns a tuple with the UseRpzDropIpRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzDropIpRule

`func (o *GetViewResponse) SetUseRpzDropIpRule(v bool)`

SetUseRpzDropIpRule sets UseRpzDropIpRule field to given value.

### HasUseRpzDropIpRule

`func (o *GetViewResponse) HasUseRpzDropIpRule() bool`

HasUseRpzDropIpRule returns a boolean if a field has been set.

### GetUseRpzQnameWaitRecurse

`func (o *GetViewResponse) GetUseRpzQnameWaitRecurse() bool`

GetUseRpzQnameWaitRecurse returns the UseRpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetUseRpzQnameWaitRecurseOk

`func (o *GetViewResponse) GetUseRpzQnameWaitRecurseOk() (*bool, bool)`

GetUseRpzQnameWaitRecurseOk returns a tuple with the UseRpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzQnameWaitRecurse

`func (o *GetViewResponse) SetUseRpzQnameWaitRecurse(v bool)`

SetUseRpzQnameWaitRecurse sets UseRpzQnameWaitRecurse field to given value.

### HasUseRpzQnameWaitRecurse

`func (o *GetViewResponse) HasUseRpzQnameWaitRecurse() bool`

HasUseRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetUseScavengingSettings

`func (o *GetViewResponse) GetUseScavengingSettings() bool`

GetUseScavengingSettings returns the UseScavengingSettings field if non-nil, zero value otherwise.

### GetUseScavengingSettingsOk

`func (o *GetViewResponse) GetUseScavengingSettingsOk() (*bool, bool)`

GetUseScavengingSettingsOk returns a tuple with the UseScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseScavengingSettings

`func (o *GetViewResponse) SetUseScavengingSettings(v bool)`

SetUseScavengingSettings sets UseScavengingSettings field to given value.

### HasUseScavengingSettings

`func (o *GetViewResponse) HasUseScavengingSettings() bool`

HasUseScavengingSettings returns a boolean if a field has been set.

### GetUseSortlist

`func (o *GetViewResponse) GetUseSortlist() bool`

GetUseSortlist returns the UseSortlist field if non-nil, zero value otherwise.

### GetUseSortlistOk

`func (o *GetViewResponse) GetUseSortlistOk() (*bool, bool)`

GetUseSortlistOk returns a tuple with the UseSortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSortlist

`func (o *GetViewResponse) SetUseSortlist(v bool)`

SetUseSortlist sets UseSortlist field to given value.

### HasUseSortlist

`func (o *GetViewResponse) HasUseSortlist() bool`

HasUseSortlist returns a boolean if a field has been set.

### GetResult

`func (o *GetViewResponse) GetResult() View`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetViewResponse) GetResultOk() (*View, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetViewResponse) SetResult(v View)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetViewResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


