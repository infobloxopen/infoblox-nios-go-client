# MemberDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddClientIpMacOptions** | Pointer to **bool** | Add custom IP, MAC and DNS View name ENDS0 options to outgoing recursive queries. | [optional] 
**AdditionalIpList** | Pointer to **[]string** | The list of additional IP addresses on which DNS is enabled for a Grid member. Only one of \&quot;additional_ip_list\&quot; or \&quot;additional_ip_list_struct\&quot; should be set when modifying the object. | [optional] 
**AdditionalIpListStruct** | Pointer to [**[]MemberDnsAdditionalIpListStruct**](MemberDnsAdditionalIpListStruct.md) | The list of additional IP addresses and IP Space Discriminator short names on which DNS is enabled for a Grid member. Only one of \&quot;additional_ip_list\&quot; or \&quot;additional_ip_list_struct\&quot; should be set when modifying the object. | [optional] 
**AllowGssTsigZoneUpdates** | Pointer to **bool** | Determines whether the GSS-TSIG zone updates is enabled for the Grid member. | [optional] 
**AllowQuery** | Pointer to [**[]MemberDnsAllowQuery**](MemberDnsAllowQuery.md) | Determines if queries from specified IPv4 or IPv6 addresses and networks are enabled or not. The appliance can also use Transaction Signature (TSIG) keys to authenticate the queries. This setting overrides the Grid query settings. | [optional] 
**AllowRecursiveQuery** | Pointer to **bool** | Determines if the responses to recursive queries is enabled or not. This setting overrides Grid recursive query settings. | [optional] 
**AllowTransfer** | Pointer to [**[]MemberDnsAllowTransfer**](MemberDnsAllowTransfer.md) | Allows or rejects zone transfers from specified IPv4 or IPv6 addresses and networks or allows transfers from hosts authenticated by Transaction signature (TSIG) key. This setting overrides the Grid zone transfer settings. | [optional] 
**AllowUpdate** | Pointer to [**[]MemberDnsAllowUpdate**](MemberDnsAllowUpdate.md) | Allows or rejects dynamic updates from specified IPv4 or IPv6 addresses, networks or from host authenticated by TSIG key. This setting overrides Grid update settings. | [optional] 
**AnonymizeResponseLogging** | Pointer to **bool** | The flag that indicates whether the anonymization of captured DNS responses is enabled or disabled. | [optional] 
**AtcFwdEnable** | Pointer to **bool** | Enable DNS recursive query forwarding to Active Trust Cloud. | [optional] 
**AttackMitigation** | Pointer to [**MemberDnsAttackMitigation**](MemberDnsAttackMitigation.md) |  | [optional] 
**AutoBlackhole** | Pointer to [**MemberDnsAutoBlackhole**](MemberDnsAutoBlackhole.md) |  | [optional] 
**AutoCreateAAndPtrForLan2** | Pointer to **bool** | Determines if the auto-generation of A and PTR records for the LAN2 IP address is enabled or not, if DNS service is enabled on LAN2. | [optional] 
**AutoCreateAaaaAndIpv6ptrForLan2** | Pointer to **bool** | Determines if auto-generation of AAAA and IPv6 PTR records for LAN2 IPv6 address is enabled or not. | [optional] 
**AutoSortViews** | Pointer to **bool** | Determines if a Grid member to automatically sort DNS views is enabled or not. The order of the DNS views determines the order in which the appliance checks the match lists. | [optional] 
**BindCheckNamesPolicy** | Pointer to **string** | The BIND check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This method applies only if the host name restriction policy is set to &#39;Strict Hostname Checking&#39;. | [optional] 
**BindHostnameDirective** | Pointer to **string** | The value of the hostname directive for BIND. | [optional] 
**BindHostnameDirectiveFqdn** | Pointer to **string** | The value of the user-defined hostname directive for BIND. To enable user-defined hostname directive, you must set the bind_hostname_directive to \&quot;USER_DEFINED\&quot;. | [optional] 
**BlackholeList** | Pointer to [**[]MemberDnsBlackholeList**](MemberDnsBlackholeList.md) | The list of IPv4 or IPv6 addresses and networks from which DNS queries are blocked. This setting overrides the Grid blackhole_list. | [optional] 
**BlacklistAction** | Pointer to **string** | The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist_ruleset method. | [optional] 
**BlacklistLogQuery** | Pointer to **bool** | Determines if blacklist redirection queries are logged or not. | [optional] 
**BlacklistRedirectAddresses** | Pointer to **[]string** | The IP addresses the appliance includes in the response it sends in place of a blacklisted IP address. | [optional] 
**BlacklistRedirectTtl** | Pointer to **int64** | The TTL value of the synthetic DNS responses that result from blacklist redirection. | [optional] 
**BlacklistRulesets** | Pointer to **[]string** | The DNS Ruleset object names assigned at the Grid level for blacklist redirection. | [optional] 
**CaptureDnsQueriesOnAllDomains** | Pointer to **bool** | The flag that indicates whether the capture of DNS queries for all domains is enabled or disabled. | [optional] 
**CheckNamesForDdnsAndZoneTransfer** | Pointer to **bool** | Determines whether the application of BIND check-names for zone transfers and DDNS updates are enabled. | [optional] 
**CopyClientIpMacOptions** | Pointer to **bool** | Copy custom IP, MAC and DNS View name ENDS0 options from incoming to outgoing recursive queries. | [optional] 
**CopyXferToNotify** | Pointer to **bool** | Copies the allowed IPs from the zone transfer list into the also-notify statement in the named.conf file. | [optional] 
**CustomRootNameServers** | Pointer to [**[]MemberDnsCustomRootNameServers**](MemberDnsCustomRootNameServers.md) | The list of custom root name servers. You can either select and use Internet root name servers or specify custom root name servers by providing a host name and IP address to which the Infoblox appliance can send queries. | [optional] 
**DisableEdns** | Pointer to **bool** | The EDNS0 support for queries that require recursive resolution on Grid members. | [optional] 
**Dns64Groups** | Pointer to **[]string** | The list of DNS64 synthesis groups associated with this member. | [optional] 
**DnsCacheAccelerationStatus** | Pointer to **string** | The DNS cache acceleration status. | [optional] [readonly] 
**DnsCacheAccelerationTtl** | Pointer to **int64** | The minimum TTL value, in seconds, that a DNS record must have in order for it to be cached by the DNS Cache Acceleration service. An integer from 1 to 65000 that represents the TTL in seconds. | [optional] 
**DnsHealthCheckAnycastControl** | Pointer to **bool** | The flag that indicates whether the anycast failure (BFD session down) is enabled on member failure or not. | [optional] 
**DnsHealthCheckDomainList** | Pointer to **[]string** | The list of domain names for the DNS health check. | [optional] 
**DnsHealthCheckInterval** | Pointer to **int64** | The time interval (in seconds) for DNS health check. | [optional] 
**DnsHealthCheckRecursionFlag** | Pointer to **bool** | The flag that indicates whether the recursive DNS health check is enabled or not. | [optional] 
**DnsHealthCheckRetries** | Pointer to **int64** | The number of DNS health check retries. | [optional] 
**DnsHealthCheckTimeout** | Pointer to **int64** | The DNS health check timeout interval (in seconds). | [optional] 
**DnsNotifyTransferSource** | Pointer to **string** | Determines which IP address is used as the source for DDNS notify and transfer operations. | [optional] 
**DnsNotifyTransferSourceAddress** | Pointer to **string** | The source address used if dns_notify_transfer_source type is \&quot;IP\&quot;. | [optional] 
**DnsOverTlsService** | Pointer to **bool** | Enables DNS over TLS service. | [optional] 
**DnsQueryCaptureFileTimeLimit** | Pointer to **int64** | The time limit (in minutes) for the DNS query capture file. | [optional] 
**DnsQuerySourceAddress** | Pointer to **string** | The source address used if dns_query_source_interface type is \&quot;IP\&quot;. | [optional] 
**DnsQuerySourceInterface** | Pointer to **string** | Determines which IP address is used as the source for DDNS query operations. | [optional] 
**DnsViewAddressSettings** | Pointer to [**[]MemberDnsDnsViewAddressSettings**](MemberDnsDnsViewAddressSettings.md) | Array of notify/query source settings for views. | [optional] 
**DnssecBlacklistEnabled** | Pointer to **bool** | Determines if the blacklist rules for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecDns64Enabled** | Pointer to **bool** | Determines if the DNS64 groups for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Determines if the DNS security extension is enabled or not. | [optional] 
**DnssecExpiredSignaturesEnabled** | Pointer to **bool** | Determines when the DNS member accepts expired signatures. | [optional] 
**DnssecNegativeTrustAnchors** | Pointer to **[]string** | A list of zones for which the server does not perform DNSSEC validation. | [optional] 
**DnssecNxdomainEnabled** | Pointer to **bool** | Determines if the NXDOMAIN rules for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecRpzEnabled** | Pointer to **bool** | Determines if the RPZ policies for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecTrustedKeys** | Pointer to [**[]MemberDnsDnssecTrustedKeys**](MemberDnsDnssecTrustedKeys.md) | The list of trusted keys for the DNSSEC feature. | [optional] 
**DnssecValidationEnabled** | Pointer to **bool** | Determines if the DNS security validation is enabled or not. | [optional] 
**DnstapSetting** | Pointer to [**MemberDnsDnstapSetting**](MemberDnsDnstapSetting.md) |  | [optional] 
**DohHttpsSessionDuration** | Pointer to **int64** | DNS over HTTPS sessions duration. | [optional] 
**DohService** | Pointer to **bool** | Enables DNS over HTTPS service. | [optional] 
**DomainsToCaptureDnsQueries** | Pointer to **[]string** | The list of domains for DNS query capture. | [optional] 
**DtcDnsQueriesSpecificBehavior** | Pointer to **string** | Setting to control specific behavior for DTC DNS responses for incoming lbdn matched queries. | [optional] 
**DtcEdnsPreferClientSubnet** | Pointer to **bool** | Determines whether to prefer the client address from the edns-client-subnet option for DTC or not. | [optional] 
**DtcHealthSource** | Pointer to **string** | The health check source type. | [optional] 
**DtcHealthSourceAddress** | Pointer to **string** | The source address used if dtc_health_source type is \&quot;IP\&quot;. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes. | [optional] 
**EnableBlackhole** | Pointer to **bool** | Determines if the blocking of DNS queries is enabled or not. This setting overrides the Grid enable_blackhole settings. | [optional] 
**EnableBlacklist** | Pointer to **bool** | Determines if a blacklist is enabled or not on the Grid member. | [optional] 
**EnableCaptureDnsQueries** | Pointer to **bool** | The flag that indicates whether the capture of DNS queries is enabled or disabled. | [optional] 
**EnableCaptureDnsResponses** | Pointer to **bool** | The flag that indicates whether the capture of DNS responses is enabled or disabled. | [optional] 
**EnableDns** | Pointer to **bool** | Determines if the DNS service of a member is enabled or not. | [optional] 
**EnableDns64** | Pointer to **bool** | Determines if the DNS64 support is enabled or not for this member. | [optional] 
**EnableDnsCacheAcceleration** | Pointer to **bool** | Determines if the DNS Cache Acceleration service is enabled or not for a member. | [optional] 
**EnableDnsHealthCheck** | Pointer to **bool** | The flag that indicates whether the DNS health check is enabled or not. | [optional] 
**EnableDnstapQueries** | Pointer to **bool** | Determines whether the query messages need to be forwarded to DNSTAP or not. | [optional] 
**EnableDnstapResponses** | Pointer to **bool** | Determines whether the response messages need to be forwarded to DNSTAP or not. | [optional] 
**EnableDnstapViolationsTls** | Pointer to **bool** | Determines whether the violations messages need to be forwarded to DNSTAP or not. | [optional] 
**EnableExcludedDomainNames** | Pointer to **bool** | The flag that indicates whether excluding domain names from captured DNS queries and responses is enabled or disabled. | [optional] 
**EnableFixedRrsetOrderFqdns** | Pointer to **bool** | Determines if the fixed RRset order FQDN is enabled or not. | [optional] 
**EnableFtc** | Pointer to **bool** | Determines whether Fault Tolerant Caching (FTC) is enabled. | [optional] 
**EnableGssTsig** | Pointer to **bool** | Determines whether the appliance is enabled to receive GSS-TSIG authenticated updates from DHCP clients. | [optional] 
**EnableNotifySourcePort** | Pointer to **bool** | Determines if the notify source port for a member is enabled or not. | [optional] 
**EnableQueryRewrite** | Pointer to **bool** | Determines if the DNS query rewrite is enabled or not for this member. | [optional] 
**EnableQuerySourcePort** | Pointer to **bool** | Determines if the query source port for a memer is enabled or not. | [optional] 
**ExcludedDomainNames** | Pointer to **[]string** | The list of domains that are excluded from DNS query and response capture. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FileTransferSetting** | Pointer to [**MemberDnsFileTransferSetting**](MemberDnsFileTransferSetting.md) |  | [optional] 
**FilterAaaa** | Pointer to **string** | The type of AAAA filtering for this member DNS object. | [optional] 
**FilterAaaaList** | Pointer to [**[]MemberDnsFilterAaaaList**](MemberDnsFilterAaaaList.md) | The list of IPv4 addresses and networks from which queries are received. AAAA filtering is applied to these addresses. | [optional] 
**FixedRrsetOrderFqdns** | Pointer to [**[]MemberDnsFixedRrsetOrderFqdns**](MemberDnsFixedRrsetOrderFqdns.md) | The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to &#39;true&#39;, unless the same request sets the enable field to &#39;false&#39;. | [optional] 
**ForwardOnly** | Pointer to **bool** | Permits this member to send queries to forwarders only. When the value is \&quot;true\&quot;, the member sends queries to forwarders only, and not to other internal or Internet root servers. | [optional] 
**ForwardUpdates** | Pointer to **bool** | Allows secondary servers to forward updates to the DNS server. This setting overrides grid update settings. | [optional] 
**Forwarders** | Pointer to **[]string** | The forwarders for the member. A forwarder is essentially a name server to which other name servers first send all of their off-site queries. The forwarder builds up a cache of information, avoiding the need for the other name servers to send queries off-site. This setting overrides the Grid level setting. | [optional] 
**FtcExpiredRecordTimeout** | Pointer to **int64** | The timeout interval (in seconds) after which the expired Fault Tolerant Caching (FTC)record is stale and no longer valid. | [optional] 
**FtcExpiredRecordTtl** | Pointer to **int64** | The TTL value (in seconds) of the expired Fault Tolerant Caching (FTC) record in DNS responses. | [optional] 
**GlueRecordAddresses** | Pointer to [**[]MemberDnsGlueRecordAddresses**](MemberDnsGlueRecordAddresses.md) | The list of glue record addresses. | [optional] 
**GssTsigKeys** | Pointer to **[]string** | The list of GSS-TSIG keys for a member DNS object. | [optional] 
**HostName** | Pointer to **string** | The host name of the Grid member. | [optional] [readonly] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the Grid member. | [optional] [readonly] 
**Ipv6GlueRecordAddresses** | Pointer to [**[]MemberDnsIpv6GlueRecordAddresses**](MemberDnsIpv6GlueRecordAddresses.md) | The list of IPv6 glue record addresses. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Grid member. | [optional] [readonly] 
**LoggingCategories** | Pointer to [**MemberDnsLoggingCategories**](MemberDnsLoggingCategories.md) |  | [optional] 
**MaxCacheTtl** | Pointer to **int64** | The maximum time (in seconds) for which the server will cache positive answers. | [optional] 
**MaxCachedLifetime** | Pointer to **int64** | The maximum time in seconds a DNS response can be stored in the hardware acceleration cache. Valid values are unsigned integer between 60 and 86400, inclusive. | [optional] 
**MaxNcacheTtl** | Pointer to **int64** | The maximum time (in seconds) for which the server will cache negative (NXDOMAIN) responses. The maximum allowed value is 604800. | [optional] 
**MaxUdpSize** | Pointer to **int64** | The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes. | [optional] 
**MinimalResp** | Pointer to **bool** | Enables the ability to return a minimal amount of data in response to a query. This capability speeds up the DNS services provided by the appliance. | [optional] 
**NotifyDelay** | Pointer to **int64** | Specifies the number of seconds of delay the notify messages are sent to secondaries. | [optional] 
**NotifySourcePort** | Pointer to **int64** | The source port for notify messages. When requesting zone transfers from the primary server, some secondary DNS servers use the source port number (the primary server used to send the notify message) as the destination port number in the zone transfer request. This setting overrides Grid static source port settings. Valid values are between 1 and 63999. The default is selected by BIND. | [optional] 
**NxdomainLogQuery** | Pointer to **bool** | Determines if NXDOMAIN redirection queries are logged or not. | [optional] 
**NxdomainRedirect** | Pointer to **bool** | Enables NXDOMAIN redirection. | [optional] 
**NxdomainRedirectAddresses** | Pointer to **[]string** | The IPv4 NXDOMAIN redirection addresses. | [optional] 
**NxdomainRedirectAddressesV6** | Pointer to **[]string** | The IPv6 NXDOMAIN redirection addresses. | [optional] 
**NxdomainRedirectTtl** | Pointer to **int64** | The TTL value of synthetic DNS responses that result from NXDOMAIN redirection. | [optional] 
**NxdomainRulesets** | Pointer to **[]string** | The names of the Ruleset objects assigned at the Grid level for NXDOMAIN redirection. | [optional] 
**QuerySourcePort** | Pointer to **int64** | The source port for queries. Specifying a source port number for recursive queries ensures that a firewall will allow the response. Valid values are between 1 and 63999. The default is selected by BIND. | [optional] 
**RecordNamePolicy** | Pointer to **string** | The record name restriction policy. | [optional] 
**RecursiveClientLimit** | Pointer to **int64** | A limit on the number of concurrent recursive clients. | [optional] 
**RecursiveQueryList** | Pointer to [**[]MemberDnsRecursiveQueryList**](MemberDnsRecursiveQueryList.md) | The list of IPv4 or IPv6 addresses, networks or hosts authenticated by Transaction signature (TSIG) key from which recursive queries are allowed or denied. | [optional] 
**RecursiveResolver** | Pointer to **string** | The recursive resolver for member DNS. UNBOUND support has been deprecated from NIOS 9.0 onwards. | [optional] 
**ResolverQueryTimeout** | Pointer to **int64** | The recursive query timeout for the member. The value must be 0 or between 10 and 30. | [optional] 
**ResponseRateLimiting** | Pointer to [**MemberDnsResponseRateLimiting**](MemberDnsResponseRateLimiting.md) |  | [optional] 
**RootNameServerType** | Pointer to **string** | Determines the type of root name servers. | [optional] 
**RpzDisableNsdnameNsip** | Pointer to **bool** | Enables NSDNAME and NSIP resource records from RPZ feeds at member level. | [optional] 
**RpzDropIpRuleEnabled** | Pointer to **bool** | Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv4** | Pointer to **int64** | The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv6** | Pointer to **int64** | The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length. | [optional] 
**RpzQnameWaitRecurse** | Pointer to **bool** | The flag that indicates whether recursive RPZ lookups are enabled. | [optional] 
**SerialQueryRate** | Pointer to **int64** | The number of maximum concurrent SOA queries per second for the member. | [optional] 
**ServerIdDirective** | Pointer to **string** | The value of the server-id directive for BIND and Unbound DNS. | [optional] 
**ServerIdDirectiveString** | Pointer to **string** | The value of the user-defined hostname directive for BIND DNS. To enable user-defined hostname directive, you must set the bind_hostname_directive to \&quot;USER_DEFINED\&quot;. | [optional] 
**SkipInGridRpzQueries** | Pointer to **bool** | Determines if RPZ rules are applied to queries originated from this member and received by other Grid members. | [optional] 
**Sortlist** | Pointer to [**[]MemberDnsSortlist**](MemberDnsSortlist.md) | A sort list determines the order of addresses in responses made to DNS queries. This setting overrides Grid sort list settings. | [optional] 
**StoreLocally** | Pointer to **bool** | The flag that indicates whether the storage of query capture reports on the appliance is enabled or disabled. | [optional] 
**SyslogFacility** | Pointer to **string** | The syslog facility. This is the location on the syslog server to which you want to sort the DNS logging messages. This setting overrides the Grid logging facility settings. | [optional] 
**TcpIdleTimeout** | Pointer to **int64** | TCP Idle timeout for DNS over TLS connections. | [optional] 
**TlsSessionDuration** | Pointer to **int64** | DNS over TLS sessions duration. | [optional] 
**TransferExcludedServers** | Pointer to **[]string** | Excludes specified DNS servers during zone transfers. | [optional] 
**TransferFormat** | Pointer to **string** | The BIND format for a zone transfer. This provides tracking capabilities for single or multiple transfers and their associated servers. | [optional] 
**TransfersIn** | Pointer to **int64** | The number of maximum concurrent transfers for the member. | [optional] 
**TransfersOut** | Pointer to **int64** | The number of maximum outbound concurrent zone transfers for the member. | [optional] 
**TransfersPerNs** | Pointer to **int64** | The number of maximum concurrent transfers per member for the member. | [optional] 
**UpstreamAddressFamilyPreference** | Pointer to **string** | Upstream address family preference when dual mode is configured. | [optional] 
**UseAddClientIpMacOptions** | Pointer to **bool** | Use flag for: add_client_ip_mac_options | [optional] 
**UseAllowQuery** | Pointer to **bool** | Use flag for: allow_query | [optional] 
**UseAllowTransfer** | Pointer to **bool** | Use flag for: allow_transfer | [optional] 
**UseAttackMitigation** | Pointer to **bool** | Use flag for: attack_mitigation | [optional] 
**UseAutoBlackhole** | Pointer to **bool** | Use flag for: auto_blackhole | [optional] 
**UseBindHostnameDirective** | Pointer to **bool** | Use flag for: bind_hostname_directive | [optional] 
**UseBlackhole** | Pointer to **bool** | Use flag for: enable_blackhole | [optional] 
**UseBlacklist** | Pointer to **bool** | Use flag for: blackhole_list , blacklist_action, blacklist_log_query, blacklist_redirect_addresses, blacklist_redirect_ttl, blacklist_rulesets, enable_blacklist | [optional] 
**UseCaptureDnsQueriesOnAllDomains** | Pointer to **bool** | Use flag for: capture_dns_queries_on_all_domains | [optional] 
**UseCopyClientIpMacOptions** | Pointer to **bool** | Use flag for: copy_client_ip_mac_options | [optional] 
**UseCopyXferToNotify** | Pointer to **bool** | Use flag for: copy_xfer_to_notify | [optional] 
**UseDisableEdns** | Pointer to **bool** | Use flag for: disable_edns | [optional] 
**UseDns64** | Pointer to **bool** | Use flag for: enable_dns64 , dns64_groups | [optional] 
**UseDnsCacheAccelerationTtl** | Pointer to **bool** | Use flag for: dns_cache_acceleration_ttl | [optional] 
**UseDnsHealthCheck** | Pointer to **bool** | Use flag for: dns_health_check_domain_list , dns_health_check_recursion_flag, dns_health_check_anycast_control, enable_dns_health_check, dns_health_check_interval, dns_health_check_timeout, dns_health_check_retries | [optional] 
**UseDnssec** | Pointer to **bool** | Use flag for: dnssec_enabled , dnssec_expired_signatures_enabled, dnssec_validation_enabled, dnssec_trusted_keys | [optional] 
**UseDnstapSetting** | Pointer to **bool** | Use flag for: enable_dnstap_queries , enable_dnstap_responses, enable_dnstap_violations_tls, dnstap_setting | [optional] 
**UseDtcDnsQueriesSpecificBehavior** | Pointer to **bool** | Use flag for: dtc_dns_queries_specific_behavior | [optional] 
**UseDtcEdnsPreferClientSubnet** | Pointer to **bool** | Use flag for: dtc_edns_prefer_client_subnet | [optional] 
**UseEdnsUdpSize** | Pointer to **bool** | Use flag for: edns_udp_size | [optional] 
**UseEnableCaptureDns** | Pointer to **bool** | Use flag for: enable_capture_dns_queries , enable_capture_dns_responses | [optional] 
**UseEnableExcludedDomainNames** | Pointer to **bool** | Use flag for: enable_excluded_domain_names | [optional] 
**UseEnableGssTsig** | Pointer to **bool** | Use flag for: enable_gss_tsig | [optional] 
**UseEnableQueryRewrite** | Pointer to **bool** | Use flag for: enable_query_rewrite | [optional] 
**UseFilterAaaa** | Pointer to **bool** | Use flag for: filter_aaaa , filter_aaaa_list | [optional] 
**UseFixedRrsetOrderFqdns** | Pointer to **bool** | Use flag for: fixed_rrset_order_fqdns , enable_fixed_rrset_order_fqdns | [optional] 
**UseForwardUpdates** | Pointer to **bool** | Use flag for: forward_updates | [optional] 
**UseForwarders** | Pointer to **bool** | Use flag for: forwarders , forward_only | [optional] 
**UseFtc** | Pointer to **bool** | Use flag for: enable_ftc , ftc_expired_record_ttl, ftc_expired_record_timeout | [optional] 
**UseGssTsigKeys** | Pointer to **bool** | Use flag for: gss_tsig_keys | [optional] 
**UseLan2Ipv6Port** | Pointer to **bool** | Determines if the DNS service on the IPv6 LAN2 port is enabled or not. | [optional] 
**UseLan2Port** | Pointer to **bool** | Determines if the DNS service on the LAN2 port is enabled or not. | [optional] 
**UseLanIpv6Port** | Pointer to **bool** | Determines if the DNS service on the IPv6 LAN port is enabled or not. | [optional] 
**UseLanPort** | Pointer to **bool** | Determines the status of the use of DNS services on the IPv4 LAN1 port. | [optional] 
**UseLoggingCategories** | Pointer to **bool** | Use flag for: logging_categories | [optional] 
**UseMaxCacheTtl** | Pointer to **bool** | Use flag for: max_cache_ttl | [optional] 
**UseMaxCachedLifetime** | Pointer to **bool** | Use flag for: max_cached_lifetime | [optional] 
**UseMaxNcacheTtl** | Pointer to **bool** | Use flag for: max_ncache_ttl | [optional] 
**UseMaxUdpSize** | Pointer to **bool** | Use flag for: max_udp_size | [optional] 
**UseMgmtIpv6Port** | Pointer to **bool** | Determines if the DNS services on the IPv6 MGMT port is enabled or not. | [optional] 
**UseMgmtPort** | Pointer to **bool** | Determines if the DNS services on the MGMT port is enabled or not. | [optional] 
**UseNotifyDelay** | Pointer to **bool** | Use flag for: notify_delay | [optional] 
**UseNxdomainRedirect** | Pointer to **bool** | Use flag for: nxdomain_redirect , nxdomain_redirect_addresses, nxdomain_redirect_addresses_v6, nxdomain_redirect_ttl, nxdomain_log_query, nxdomain_rulesets | [optional] 
**UseRecordNamePolicy** | Pointer to **bool** | Use flag for: record_name_policy | [optional] 
**UseRecursiveClientLimit** | Pointer to **bool** | Use flag for: recursive_client_limit | [optional] 
**UseRecursiveQuerySetting** | Pointer to **bool** | Use flag for: allow_recursive_query , recursive_query_list | [optional] 
**UseResolverQueryTimeout** | Pointer to **bool** | Use flag for: resolver_query_timeout | [optional] 
**UseResponseRateLimiting** | Pointer to **bool** | Use flag for: response_rate_limiting | [optional] 
**UseRootNameServer** | Pointer to **bool** | Use flag for: root_name_server_type , custom_root_name_servers, use_root_server_for_all_views | [optional] 
**UseRootServerForAllViews** | Pointer to **bool** | Determines if root name servers should be applied to all views or only to Default view. | [optional] 
**UseRpzDisableNsdnameNsip** | Pointer to **bool** | Use flag for: rpz_disable_nsdname_nsip | [optional] 
**UseRpzDropIpRule** | Pointer to **bool** | Use flag for: rpz_drop_ip_rule_enabled , rpz_drop_ip_rule_min_prefix_length_ipv4, rpz_drop_ip_rule_min_prefix_length_ipv6 | [optional] 
**UseRpzQnameWaitRecurse** | Pointer to **bool** | Use flag for: rpz_qname_wait_recurse | [optional] 
**UseSerialQueryRate** | Pointer to **bool** | Use flag for: serial_query_rate | [optional] 
**UseServerIdDirective** | Pointer to **bool** | Use flag for: server_id_directive | [optional] 
**UseSortlist** | Pointer to **bool** | Use flag for: sortlist | [optional] 
**UseSourcePorts** | Pointer to **bool** | Use flag for: enable_notify_source_port , notify_source_port, enable_query_source_port, query_source_port | [optional] 
**UseSyslogFacility** | Pointer to **bool** | Use flag for: syslog_facility | [optional] 
**UseTransfersIn** | Pointer to **bool** | Use flag for: transfers_in | [optional] 
**UseTransfersOut** | Pointer to **bool** | Use flag for: transfers_out | [optional] 
**UseTransfersPerNs** | Pointer to **bool** | Use flag for: transfers_per_ns | [optional] 
**UseUpdateSetting** | Pointer to **bool** | Use flag for: allow_update , allow_gss_tsig_zone_updates | [optional] 
**UseZoneTransferFormat** | Pointer to **bool** | Use flag for: transfer_excluded_servers , transfer_format | [optional] 
**Views** | Pointer to **[]string** | The list of views associated with this member. | [optional] 

## Methods

### NewMemberDns

`func NewMemberDns() *MemberDns`

NewMemberDns instantiates a new MemberDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsWithDefaults

`func NewMemberDnsWithDefaults() *MemberDns`

NewMemberDnsWithDefaults instantiates a new MemberDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberDns) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberDns) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberDns) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberDns) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddClientIpMacOptions

`func (o *MemberDns) GetAddClientIpMacOptions() bool`

GetAddClientIpMacOptions returns the AddClientIpMacOptions field if non-nil, zero value otherwise.

### GetAddClientIpMacOptionsOk

`func (o *MemberDns) GetAddClientIpMacOptionsOk() (*bool, bool)`

GetAddClientIpMacOptionsOk returns a tuple with the AddClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddClientIpMacOptions

`func (o *MemberDns) SetAddClientIpMacOptions(v bool)`

SetAddClientIpMacOptions sets AddClientIpMacOptions field to given value.

### HasAddClientIpMacOptions

`func (o *MemberDns) HasAddClientIpMacOptions() bool`

HasAddClientIpMacOptions returns a boolean if a field has been set.

### GetAdditionalIpList

`func (o *MemberDns) GetAdditionalIpList() []string`

GetAdditionalIpList returns the AdditionalIpList field if non-nil, zero value otherwise.

### GetAdditionalIpListOk

`func (o *MemberDns) GetAdditionalIpListOk() (*[]string, bool)`

GetAdditionalIpListOk returns a tuple with the AdditionalIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpList

`func (o *MemberDns) SetAdditionalIpList(v []string)`

SetAdditionalIpList sets AdditionalIpList field to given value.

### HasAdditionalIpList

`func (o *MemberDns) HasAdditionalIpList() bool`

HasAdditionalIpList returns a boolean if a field has been set.

### GetAdditionalIpListStruct

`func (o *MemberDns) GetAdditionalIpListStruct() []MemberDnsAdditionalIpListStruct`

GetAdditionalIpListStruct returns the AdditionalIpListStruct field if non-nil, zero value otherwise.

### GetAdditionalIpListStructOk

`func (o *MemberDns) GetAdditionalIpListStructOk() (*[]MemberDnsAdditionalIpListStruct, bool)`

GetAdditionalIpListStructOk returns a tuple with the AdditionalIpListStruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpListStruct

`func (o *MemberDns) SetAdditionalIpListStruct(v []MemberDnsAdditionalIpListStruct)`

SetAdditionalIpListStruct sets AdditionalIpListStruct field to given value.

### HasAdditionalIpListStruct

`func (o *MemberDns) HasAdditionalIpListStruct() bool`

HasAdditionalIpListStruct returns a boolean if a field has been set.

### GetAllowGssTsigZoneUpdates

`func (o *MemberDns) GetAllowGssTsigZoneUpdates() bool`

GetAllowGssTsigZoneUpdates returns the AllowGssTsigZoneUpdates field if non-nil, zero value otherwise.

### GetAllowGssTsigZoneUpdatesOk

`func (o *MemberDns) GetAllowGssTsigZoneUpdatesOk() (*bool, bool)`

GetAllowGssTsigZoneUpdatesOk returns a tuple with the AllowGssTsigZoneUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGssTsigZoneUpdates

`func (o *MemberDns) SetAllowGssTsigZoneUpdates(v bool)`

SetAllowGssTsigZoneUpdates sets AllowGssTsigZoneUpdates field to given value.

### HasAllowGssTsigZoneUpdates

`func (o *MemberDns) HasAllowGssTsigZoneUpdates() bool`

HasAllowGssTsigZoneUpdates returns a boolean if a field has been set.

### GetAllowQuery

`func (o *MemberDns) GetAllowQuery() []MemberDnsAllowQuery`

GetAllowQuery returns the AllowQuery field if non-nil, zero value otherwise.

### GetAllowQueryOk

`func (o *MemberDns) GetAllowQueryOk() (*[]MemberDnsAllowQuery, bool)`

GetAllowQueryOk returns a tuple with the AllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuery

`func (o *MemberDns) SetAllowQuery(v []MemberDnsAllowQuery)`

SetAllowQuery sets AllowQuery field to given value.

### HasAllowQuery

`func (o *MemberDns) HasAllowQuery() bool`

HasAllowQuery returns a boolean if a field has been set.

### GetAllowRecursiveQuery

`func (o *MemberDns) GetAllowRecursiveQuery() bool`

GetAllowRecursiveQuery returns the AllowRecursiveQuery field if non-nil, zero value otherwise.

### GetAllowRecursiveQueryOk

`func (o *MemberDns) GetAllowRecursiveQueryOk() (*bool, bool)`

GetAllowRecursiveQueryOk returns a tuple with the AllowRecursiveQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRecursiveQuery

`func (o *MemberDns) SetAllowRecursiveQuery(v bool)`

SetAllowRecursiveQuery sets AllowRecursiveQuery field to given value.

### HasAllowRecursiveQuery

`func (o *MemberDns) HasAllowRecursiveQuery() bool`

HasAllowRecursiveQuery returns a boolean if a field has been set.

### GetAllowTransfer

`func (o *MemberDns) GetAllowTransfer() []MemberDnsAllowTransfer`

GetAllowTransfer returns the AllowTransfer field if non-nil, zero value otherwise.

### GetAllowTransferOk

`func (o *MemberDns) GetAllowTransferOk() (*[]MemberDnsAllowTransfer, bool)`

GetAllowTransferOk returns a tuple with the AllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransfer

`func (o *MemberDns) SetAllowTransfer(v []MemberDnsAllowTransfer)`

SetAllowTransfer sets AllowTransfer field to given value.

### HasAllowTransfer

`func (o *MemberDns) HasAllowTransfer() bool`

HasAllowTransfer returns a boolean if a field has been set.

### GetAllowUpdate

`func (o *MemberDns) GetAllowUpdate() []MemberDnsAllowUpdate`

GetAllowUpdate returns the AllowUpdate field if non-nil, zero value otherwise.

### GetAllowUpdateOk

`func (o *MemberDns) GetAllowUpdateOk() (*[]MemberDnsAllowUpdate, bool)`

GetAllowUpdateOk returns a tuple with the AllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdate

`func (o *MemberDns) SetAllowUpdate(v []MemberDnsAllowUpdate)`

SetAllowUpdate sets AllowUpdate field to given value.

### HasAllowUpdate

`func (o *MemberDns) HasAllowUpdate() bool`

HasAllowUpdate returns a boolean if a field has been set.

### GetAnonymizeResponseLogging

`func (o *MemberDns) GetAnonymizeResponseLogging() bool`

GetAnonymizeResponseLogging returns the AnonymizeResponseLogging field if non-nil, zero value otherwise.

### GetAnonymizeResponseLoggingOk

`func (o *MemberDns) GetAnonymizeResponseLoggingOk() (*bool, bool)`

GetAnonymizeResponseLoggingOk returns a tuple with the AnonymizeResponseLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeResponseLogging

`func (o *MemberDns) SetAnonymizeResponseLogging(v bool)`

SetAnonymizeResponseLogging sets AnonymizeResponseLogging field to given value.

### HasAnonymizeResponseLogging

`func (o *MemberDns) HasAnonymizeResponseLogging() bool`

HasAnonymizeResponseLogging returns a boolean if a field has been set.

### GetAtcFwdEnable

`func (o *MemberDns) GetAtcFwdEnable() bool`

GetAtcFwdEnable returns the AtcFwdEnable field if non-nil, zero value otherwise.

### GetAtcFwdEnableOk

`func (o *MemberDns) GetAtcFwdEnableOk() (*bool, bool)`

GetAtcFwdEnableOk returns a tuple with the AtcFwdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcFwdEnable

`func (o *MemberDns) SetAtcFwdEnable(v bool)`

SetAtcFwdEnable sets AtcFwdEnable field to given value.

### HasAtcFwdEnable

`func (o *MemberDns) HasAtcFwdEnable() bool`

HasAtcFwdEnable returns a boolean if a field has been set.

### GetAttackMitigation

`func (o *MemberDns) GetAttackMitigation() MemberDnsAttackMitigation`

GetAttackMitigation returns the AttackMitigation field if non-nil, zero value otherwise.

### GetAttackMitigationOk

`func (o *MemberDns) GetAttackMitigationOk() (*MemberDnsAttackMitigation, bool)`

GetAttackMitigationOk returns a tuple with the AttackMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackMitigation

`func (o *MemberDns) SetAttackMitigation(v MemberDnsAttackMitigation)`

SetAttackMitigation sets AttackMitigation field to given value.

### HasAttackMitigation

`func (o *MemberDns) HasAttackMitigation() bool`

HasAttackMitigation returns a boolean if a field has been set.

### GetAutoBlackhole

`func (o *MemberDns) GetAutoBlackhole() MemberDnsAutoBlackhole`

GetAutoBlackhole returns the AutoBlackhole field if non-nil, zero value otherwise.

### GetAutoBlackholeOk

`func (o *MemberDns) GetAutoBlackholeOk() (*MemberDnsAutoBlackhole, bool)`

GetAutoBlackholeOk returns a tuple with the AutoBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlackhole

`func (o *MemberDns) SetAutoBlackhole(v MemberDnsAutoBlackhole)`

SetAutoBlackhole sets AutoBlackhole field to given value.

### HasAutoBlackhole

`func (o *MemberDns) HasAutoBlackhole() bool`

HasAutoBlackhole returns a boolean if a field has been set.

### GetAutoCreateAAndPtrForLan2

`func (o *MemberDns) GetAutoCreateAAndPtrForLan2() bool`

GetAutoCreateAAndPtrForLan2 returns the AutoCreateAAndPtrForLan2 field if non-nil, zero value otherwise.

### GetAutoCreateAAndPtrForLan2Ok

`func (o *MemberDns) GetAutoCreateAAndPtrForLan2Ok() (*bool, bool)`

GetAutoCreateAAndPtrForLan2Ok returns a tuple with the AutoCreateAAndPtrForLan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateAAndPtrForLan2

`func (o *MemberDns) SetAutoCreateAAndPtrForLan2(v bool)`

SetAutoCreateAAndPtrForLan2 sets AutoCreateAAndPtrForLan2 field to given value.

### HasAutoCreateAAndPtrForLan2

`func (o *MemberDns) HasAutoCreateAAndPtrForLan2() bool`

HasAutoCreateAAndPtrForLan2 returns a boolean if a field has been set.

### GetAutoCreateAaaaAndIpv6ptrForLan2

`func (o *MemberDns) GetAutoCreateAaaaAndIpv6ptrForLan2() bool`

GetAutoCreateAaaaAndIpv6ptrForLan2 returns the AutoCreateAaaaAndIpv6ptrForLan2 field if non-nil, zero value otherwise.

### GetAutoCreateAaaaAndIpv6ptrForLan2Ok

`func (o *MemberDns) GetAutoCreateAaaaAndIpv6ptrForLan2Ok() (*bool, bool)`

GetAutoCreateAaaaAndIpv6ptrForLan2Ok returns a tuple with the AutoCreateAaaaAndIpv6ptrForLan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateAaaaAndIpv6ptrForLan2

`func (o *MemberDns) SetAutoCreateAaaaAndIpv6ptrForLan2(v bool)`

SetAutoCreateAaaaAndIpv6ptrForLan2 sets AutoCreateAaaaAndIpv6ptrForLan2 field to given value.

### HasAutoCreateAaaaAndIpv6ptrForLan2

`func (o *MemberDns) HasAutoCreateAaaaAndIpv6ptrForLan2() bool`

HasAutoCreateAaaaAndIpv6ptrForLan2 returns a boolean if a field has been set.

### GetAutoSortViews

`func (o *MemberDns) GetAutoSortViews() bool`

GetAutoSortViews returns the AutoSortViews field if non-nil, zero value otherwise.

### GetAutoSortViewsOk

`func (o *MemberDns) GetAutoSortViewsOk() (*bool, bool)`

GetAutoSortViewsOk returns a tuple with the AutoSortViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSortViews

`func (o *MemberDns) SetAutoSortViews(v bool)`

SetAutoSortViews sets AutoSortViews field to given value.

### HasAutoSortViews

`func (o *MemberDns) HasAutoSortViews() bool`

HasAutoSortViews returns a boolean if a field has been set.

### GetBindCheckNamesPolicy

`func (o *MemberDns) GetBindCheckNamesPolicy() string`

GetBindCheckNamesPolicy returns the BindCheckNamesPolicy field if non-nil, zero value otherwise.

### GetBindCheckNamesPolicyOk

`func (o *MemberDns) GetBindCheckNamesPolicyOk() (*string, bool)`

GetBindCheckNamesPolicyOk returns a tuple with the BindCheckNamesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCheckNamesPolicy

`func (o *MemberDns) SetBindCheckNamesPolicy(v string)`

SetBindCheckNamesPolicy sets BindCheckNamesPolicy field to given value.

### HasBindCheckNamesPolicy

`func (o *MemberDns) HasBindCheckNamesPolicy() bool`

HasBindCheckNamesPolicy returns a boolean if a field has been set.

### GetBindHostnameDirective

`func (o *MemberDns) GetBindHostnameDirective() string`

GetBindHostnameDirective returns the BindHostnameDirective field if non-nil, zero value otherwise.

### GetBindHostnameDirectiveOk

`func (o *MemberDns) GetBindHostnameDirectiveOk() (*string, bool)`

GetBindHostnameDirectiveOk returns a tuple with the BindHostnameDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindHostnameDirective

`func (o *MemberDns) SetBindHostnameDirective(v string)`

SetBindHostnameDirective sets BindHostnameDirective field to given value.

### HasBindHostnameDirective

`func (o *MemberDns) HasBindHostnameDirective() bool`

HasBindHostnameDirective returns a boolean if a field has been set.

### GetBindHostnameDirectiveFqdn

`func (o *MemberDns) GetBindHostnameDirectiveFqdn() string`

GetBindHostnameDirectiveFqdn returns the BindHostnameDirectiveFqdn field if non-nil, zero value otherwise.

### GetBindHostnameDirectiveFqdnOk

`func (o *MemberDns) GetBindHostnameDirectiveFqdnOk() (*string, bool)`

GetBindHostnameDirectiveFqdnOk returns a tuple with the BindHostnameDirectiveFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindHostnameDirectiveFqdn

`func (o *MemberDns) SetBindHostnameDirectiveFqdn(v string)`

SetBindHostnameDirectiveFqdn sets BindHostnameDirectiveFqdn field to given value.

### HasBindHostnameDirectiveFqdn

`func (o *MemberDns) HasBindHostnameDirectiveFqdn() bool`

HasBindHostnameDirectiveFqdn returns a boolean if a field has been set.

### GetBlackholeList

`func (o *MemberDns) GetBlackholeList() []MemberDnsBlackholeList`

GetBlackholeList returns the BlackholeList field if non-nil, zero value otherwise.

### GetBlackholeListOk

`func (o *MemberDns) GetBlackholeListOk() (*[]MemberDnsBlackholeList, bool)`

GetBlackholeListOk returns a tuple with the BlackholeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackholeList

`func (o *MemberDns) SetBlackholeList(v []MemberDnsBlackholeList)`

SetBlackholeList sets BlackholeList field to given value.

### HasBlackholeList

`func (o *MemberDns) HasBlackholeList() bool`

HasBlackholeList returns a boolean if a field has been set.

### GetBlacklistAction

`func (o *MemberDns) GetBlacklistAction() string`

GetBlacklistAction returns the BlacklistAction field if non-nil, zero value otherwise.

### GetBlacklistActionOk

`func (o *MemberDns) GetBlacklistActionOk() (*string, bool)`

GetBlacklistActionOk returns a tuple with the BlacklistAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistAction

`func (o *MemberDns) SetBlacklistAction(v string)`

SetBlacklistAction sets BlacklistAction field to given value.

### HasBlacklistAction

`func (o *MemberDns) HasBlacklistAction() bool`

HasBlacklistAction returns a boolean if a field has been set.

### GetBlacklistLogQuery

`func (o *MemberDns) GetBlacklistLogQuery() bool`

GetBlacklistLogQuery returns the BlacklistLogQuery field if non-nil, zero value otherwise.

### GetBlacklistLogQueryOk

`func (o *MemberDns) GetBlacklistLogQueryOk() (*bool, bool)`

GetBlacklistLogQueryOk returns a tuple with the BlacklistLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistLogQuery

`func (o *MemberDns) SetBlacklistLogQuery(v bool)`

SetBlacklistLogQuery sets BlacklistLogQuery field to given value.

### HasBlacklistLogQuery

`func (o *MemberDns) HasBlacklistLogQuery() bool`

HasBlacklistLogQuery returns a boolean if a field has been set.

### GetBlacklistRedirectAddresses

`func (o *MemberDns) GetBlacklistRedirectAddresses() []string`

GetBlacklistRedirectAddresses returns the BlacklistRedirectAddresses field if non-nil, zero value otherwise.

### GetBlacklistRedirectAddressesOk

`func (o *MemberDns) GetBlacklistRedirectAddressesOk() (*[]string, bool)`

GetBlacklistRedirectAddressesOk returns a tuple with the BlacklistRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectAddresses

`func (o *MemberDns) SetBlacklistRedirectAddresses(v []string)`

SetBlacklistRedirectAddresses sets BlacklistRedirectAddresses field to given value.

### HasBlacklistRedirectAddresses

`func (o *MemberDns) HasBlacklistRedirectAddresses() bool`

HasBlacklistRedirectAddresses returns a boolean if a field has been set.

### GetBlacklistRedirectTtl

`func (o *MemberDns) GetBlacklistRedirectTtl() int64`

GetBlacklistRedirectTtl returns the BlacklistRedirectTtl field if non-nil, zero value otherwise.

### GetBlacklistRedirectTtlOk

`func (o *MemberDns) GetBlacklistRedirectTtlOk() (*int64, bool)`

GetBlacklistRedirectTtlOk returns a tuple with the BlacklistRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectTtl

`func (o *MemberDns) SetBlacklistRedirectTtl(v int64)`

SetBlacklistRedirectTtl sets BlacklistRedirectTtl field to given value.

### HasBlacklistRedirectTtl

`func (o *MemberDns) HasBlacklistRedirectTtl() bool`

HasBlacklistRedirectTtl returns a boolean if a field has been set.

### GetBlacklistRulesets

`func (o *MemberDns) GetBlacklistRulesets() []string`

GetBlacklistRulesets returns the BlacklistRulesets field if non-nil, zero value otherwise.

### GetBlacklistRulesetsOk

`func (o *MemberDns) GetBlacklistRulesetsOk() (*[]string, bool)`

GetBlacklistRulesetsOk returns a tuple with the BlacklistRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRulesets

`func (o *MemberDns) SetBlacklistRulesets(v []string)`

SetBlacklistRulesets sets BlacklistRulesets field to given value.

### HasBlacklistRulesets

`func (o *MemberDns) HasBlacklistRulesets() bool`

HasBlacklistRulesets returns a boolean if a field has been set.

### GetCaptureDnsQueriesOnAllDomains

`func (o *MemberDns) GetCaptureDnsQueriesOnAllDomains() bool`

GetCaptureDnsQueriesOnAllDomains returns the CaptureDnsQueriesOnAllDomains field if non-nil, zero value otherwise.

### GetCaptureDnsQueriesOnAllDomainsOk

`func (o *MemberDns) GetCaptureDnsQueriesOnAllDomainsOk() (*bool, bool)`

GetCaptureDnsQueriesOnAllDomainsOk returns a tuple with the CaptureDnsQueriesOnAllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDnsQueriesOnAllDomains

`func (o *MemberDns) SetCaptureDnsQueriesOnAllDomains(v bool)`

SetCaptureDnsQueriesOnAllDomains sets CaptureDnsQueriesOnAllDomains field to given value.

### HasCaptureDnsQueriesOnAllDomains

`func (o *MemberDns) HasCaptureDnsQueriesOnAllDomains() bool`

HasCaptureDnsQueriesOnAllDomains returns a boolean if a field has been set.

### GetCheckNamesForDdnsAndZoneTransfer

`func (o *MemberDns) GetCheckNamesForDdnsAndZoneTransfer() bool`

GetCheckNamesForDdnsAndZoneTransfer returns the CheckNamesForDdnsAndZoneTransfer field if non-nil, zero value otherwise.

### GetCheckNamesForDdnsAndZoneTransferOk

`func (o *MemberDns) GetCheckNamesForDdnsAndZoneTransferOk() (*bool, bool)`

GetCheckNamesForDdnsAndZoneTransferOk returns a tuple with the CheckNamesForDdnsAndZoneTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesForDdnsAndZoneTransfer

`func (o *MemberDns) SetCheckNamesForDdnsAndZoneTransfer(v bool)`

SetCheckNamesForDdnsAndZoneTransfer sets CheckNamesForDdnsAndZoneTransfer field to given value.

### HasCheckNamesForDdnsAndZoneTransfer

`func (o *MemberDns) HasCheckNamesForDdnsAndZoneTransfer() bool`

HasCheckNamesForDdnsAndZoneTransfer returns a boolean if a field has been set.

### GetCopyClientIpMacOptions

`func (o *MemberDns) GetCopyClientIpMacOptions() bool`

GetCopyClientIpMacOptions returns the CopyClientIpMacOptions field if non-nil, zero value otherwise.

### GetCopyClientIpMacOptionsOk

`func (o *MemberDns) GetCopyClientIpMacOptionsOk() (*bool, bool)`

GetCopyClientIpMacOptionsOk returns a tuple with the CopyClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyClientIpMacOptions

`func (o *MemberDns) SetCopyClientIpMacOptions(v bool)`

SetCopyClientIpMacOptions sets CopyClientIpMacOptions field to given value.

### HasCopyClientIpMacOptions

`func (o *MemberDns) HasCopyClientIpMacOptions() bool`

HasCopyClientIpMacOptions returns a boolean if a field has been set.

### GetCopyXferToNotify

`func (o *MemberDns) GetCopyXferToNotify() bool`

GetCopyXferToNotify returns the CopyXferToNotify field if non-nil, zero value otherwise.

### GetCopyXferToNotifyOk

`func (o *MemberDns) GetCopyXferToNotifyOk() (*bool, bool)`

GetCopyXferToNotifyOk returns a tuple with the CopyXferToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyXferToNotify

`func (o *MemberDns) SetCopyXferToNotify(v bool)`

SetCopyXferToNotify sets CopyXferToNotify field to given value.

### HasCopyXferToNotify

`func (o *MemberDns) HasCopyXferToNotify() bool`

HasCopyXferToNotify returns a boolean if a field has been set.

### GetCustomRootNameServers

`func (o *MemberDns) GetCustomRootNameServers() []MemberDnsCustomRootNameServers`

GetCustomRootNameServers returns the CustomRootNameServers field if non-nil, zero value otherwise.

### GetCustomRootNameServersOk

`func (o *MemberDns) GetCustomRootNameServersOk() (*[]MemberDnsCustomRootNameServers, bool)`

GetCustomRootNameServersOk returns a tuple with the CustomRootNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNameServers

`func (o *MemberDns) SetCustomRootNameServers(v []MemberDnsCustomRootNameServers)`

SetCustomRootNameServers sets CustomRootNameServers field to given value.

### HasCustomRootNameServers

`func (o *MemberDns) HasCustomRootNameServers() bool`

HasCustomRootNameServers returns a boolean if a field has been set.

### GetDisableEdns

`func (o *MemberDns) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *MemberDns) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *MemberDns) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.

### HasDisableEdns

`func (o *MemberDns) HasDisableEdns() bool`

HasDisableEdns returns a boolean if a field has been set.

### GetDns64Groups

`func (o *MemberDns) GetDns64Groups() []string`

GetDns64Groups returns the Dns64Groups field if non-nil, zero value otherwise.

### GetDns64GroupsOk

`func (o *MemberDns) GetDns64GroupsOk() (*[]string, bool)`

GetDns64GroupsOk returns a tuple with the Dns64Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Groups

`func (o *MemberDns) SetDns64Groups(v []string)`

SetDns64Groups sets Dns64Groups field to given value.

### HasDns64Groups

`func (o *MemberDns) HasDns64Groups() bool`

HasDns64Groups returns a boolean if a field has been set.

### GetDnsCacheAccelerationStatus

`func (o *MemberDns) GetDnsCacheAccelerationStatus() string`

GetDnsCacheAccelerationStatus returns the DnsCacheAccelerationStatus field if non-nil, zero value otherwise.

### GetDnsCacheAccelerationStatusOk

`func (o *MemberDns) GetDnsCacheAccelerationStatusOk() (*string, bool)`

GetDnsCacheAccelerationStatusOk returns a tuple with the DnsCacheAccelerationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCacheAccelerationStatus

`func (o *MemberDns) SetDnsCacheAccelerationStatus(v string)`

SetDnsCacheAccelerationStatus sets DnsCacheAccelerationStatus field to given value.

### HasDnsCacheAccelerationStatus

`func (o *MemberDns) HasDnsCacheAccelerationStatus() bool`

HasDnsCacheAccelerationStatus returns a boolean if a field has been set.

### GetDnsCacheAccelerationTtl

`func (o *MemberDns) GetDnsCacheAccelerationTtl() int64`

GetDnsCacheAccelerationTtl returns the DnsCacheAccelerationTtl field if non-nil, zero value otherwise.

### GetDnsCacheAccelerationTtlOk

`func (o *MemberDns) GetDnsCacheAccelerationTtlOk() (*int64, bool)`

GetDnsCacheAccelerationTtlOk returns a tuple with the DnsCacheAccelerationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCacheAccelerationTtl

`func (o *MemberDns) SetDnsCacheAccelerationTtl(v int64)`

SetDnsCacheAccelerationTtl sets DnsCacheAccelerationTtl field to given value.

### HasDnsCacheAccelerationTtl

`func (o *MemberDns) HasDnsCacheAccelerationTtl() bool`

HasDnsCacheAccelerationTtl returns a boolean if a field has been set.

### GetDnsHealthCheckAnycastControl

`func (o *MemberDns) GetDnsHealthCheckAnycastControl() bool`

GetDnsHealthCheckAnycastControl returns the DnsHealthCheckAnycastControl field if non-nil, zero value otherwise.

### GetDnsHealthCheckAnycastControlOk

`func (o *MemberDns) GetDnsHealthCheckAnycastControlOk() (*bool, bool)`

GetDnsHealthCheckAnycastControlOk returns a tuple with the DnsHealthCheckAnycastControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckAnycastControl

`func (o *MemberDns) SetDnsHealthCheckAnycastControl(v bool)`

SetDnsHealthCheckAnycastControl sets DnsHealthCheckAnycastControl field to given value.

### HasDnsHealthCheckAnycastControl

`func (o *MemberDns) HasDnsHealthCheckAnycastControl() bool`

HasDnsHealthCheckAnycastControl returns a boolean if a field has been set.

### GetDnsHealthCheckDomainList

`func (o *MemberDns) GetDnsHealthCheckDomainList() []string`

GetDnsHealthCheckDomainList returns the DnsHealthCheckDomainList field if non-nil, zero value otherwise.

### GetDnsHealthCheckDomainListOk

`func (o *MemberDns) GetDnsHealthCheckDomainListOk() (*[]string, bool)`

GetDnsHealthCheckDomainListOk returns a tuple with the DnsHealthCheckDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckDomainList

`func (o *MemberDns) SetDnsHealthCheckDomainList(v []string)`

SetDnsHealthCheckDomainList sets DnsHealthCheckDomainList field to given value.

### HasDnsHealthCheckDomainList

`func (o *MemberDns) HasDnsHealthCheckDomainList() bool`

HasDnsHealthCheckDomainList returns a boolean if a field has been set.

### GetDnsHealthCheckInterval

`func (o *MemberDns) GetDnsHealthCheckInterval() int64`

GetDnsHealthCheckInterval returns the DnsHealthCheckInterval field if non-nil, zero value otherwise.

### GetDnsHealthCheckIntervalOk

`func (o *MemberDns) GetDnsHealthCheckIntervalOk() (*int64, bool)`

GetDnsHealthCheckIntervalOk returns a tuple with the DnsHealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckInterval

`func (o *MemberDns) SetDnsHealthCheckInterval(v int64)`

SetDnsHealthCheckInterval sets DnsHealthCheckInterval field to given value.

### HasDnsHealthCheckInterval

`func (o *MemberDns) HasDnsHealthCheckInterval() bool`

HasDnsHealthCheckInterval returns a boolean if a field has been set.

### GetDnsHealthCheckRecursionFlag

`func (o *MemberDns) GetDnsHealthCheckRecursionFlag() bool`

GetDnsHealthCheckRecursionFlag returns the DnsHealthCheckRecursionFlag field if non-nil, zero value otherwise.

### GetDnsHealthCheckRecursionFlagOk

`func (o *MemberDns) GetDnsHealthCheckRecursionFlagOk() (*bool, bool)`

GetDnsHealthCheckRecursionFlagOk returns a tuple with the DnsHealthCheckRecursionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckRecursionFlag

`func (o *MemberDns) SetDnsHealthCheckRecursionFlag(v bool)`

SetDnsHealthCheckRecursionFlag sets DnsHealthCheckRecursionFlag field to given value.

### HasDnsHealthCheckRecursionFlag

`func (o *MemberDns) HasDnsHealthCheckRecursionFlag() bool`

HasDnsHealthCheckRecursionFlag returns a boolean if a field has been set.

### GetDnsHealthCheckRetries

`func (o *MemberDns) GetDnsHealthCheckRetries() int64`

GetDnsHealthCheckRetries returns the DnsHealthCheckRetries field if non-nil, zero value otherwise.

### GetDnsHealthCheckRetriesOk

`func (o *MemberDns) GetDnsHealthCheckRetriesOk() (*int64, bool)`

GetDnsHealthCheckRetriesOk returns a tuple with the DnsHealthCheckRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckRetries

`func (o *MemberDns) SetDnsHealthCheckRetries(v int64)`

SetDnsHealthCheckRetries sets DnsHealthCheckRetries field to given value.

### HasDnsHealthCheckRetries

`func (o *MemberDns) HasDnsHealthCheckRetries() bool`

HasDnsHealthCheckRetries returns a boolean if a field has been set.

### GetDnsHealthCheckTimeout

`func (o *MemberDns) GetDnsHealthCheckTimeout() int64`

GetDnsHealthCheckTimeout returns the DnsHealthCheckTimeout field if non-nil, zero value otherwise.

### GetDnsHealthCheckTimeoutOk

`func (o *MemberDns) GetDnsHealthCheckTimeoutOk() (*int64, bool)`

GetDnsHealthCheckTimeoutOk returns a tuple with the DnsHealthCheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckTimeout

`func (o *MemberDns) SetDnsHealthCheckTimeout(v int64)`

SetDnsHealthCheckTimeout sets DnsHealthCheckTimeout field to given value.

### HasDnsHealthCheckTimeout

`func (o *MemberDns) HasDnsHealthCheckTimeout() bool`

HasDnsHealthCheckTimeout returns a boolean if a field has been set.

### GetDnsNotifyTransferSource

`func (o *MemberDns) GetDnsNotifyTransferSource() string`

GetDnsNotifyTransferSource returns the DnsNotifyTransferSource field if non-nil, zero value otherwise.

### GetDnsNotifyTransferSourceOk

`func (o *MemberDns) GetDnsNotifyTransferSourceOk() (*string, bool)`

GetDnsNotifyTransferSourceOk returns a tuple with the DnsNotifyTransferSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNotifyTransferSource

`func (o *MemberDns) SetDnsNotifyTransferSource(v string)`

SetDnsNotifyTransferSource sets DnsNotifyTransferSource field to given value.

### HasDnsNotifyTransferSource

`func (o *MemberDns) HasDnsNotifyTransferSource() bool`

HasDnsNotifyTransferSource returns a boolean if a field has been set.

### GetDnsNotifyTransferSourceAddress

`func (o *MemberDns) GetDnsNotifyTransferSourceAddress() string`

GetDnsNotifyTransferSourceAddress returns the DnsNotifyTransferSourceAddress field if non-nil, zero value otherwise.

### GetDnsNotifyTransferSourceAddressOk

`func (o *MemberDns) GetDnsNotifyTransferSourceAddressOk() (*string, bool)`

GetDnsNotifyTransferSourceAddressOk returns a tuple with the DnsNotifyTransferSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNotifyTransferSourceAddress

`func (o *MemberDns) SetDnsNotifyTransferSourceAddress(v string)`

SetDnsNotifyTransferSourceAddress sets DnsNotifyTransferSourceAddress field to given value.

### HasDnsNotifyTransferSourceAddress

`func (o *MemberDns) HasDnsNotifyTransferSourceAddress() bool`

HasDnsNotifyTransferSourceAddress returns a boolean if a field has been set.

### GetDnsOverTlsService

`func (o *MemberDns) GetDnsOverTlsService() bool`

GetDnsOverTlsService returns the DnsOverTlsService field if non-nil, zero value otherwise.

### GetDnsOverTlsServiceOk

`func (o *MemberDns) GetDnsOverTlsServiceOk() (*bool, bool)`

GetDnsOverTlsServiceOk returns a tuple with the DnsOverTlsService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOverTlsService

`func (o *MemberDns) SetDnsOverTlsService(v bool)`

SetDnsOverTlsService sets DnsOverTlsService field to given value.

### HasDnsOverTlsService

`func (o *MemberDns) HasDnsOverTlsService() bool`

HasDnsOverTlsService returns a boolean if a field has been set.

### GetDnsQueryCaptureFileTimeLimit

`func (o *MemberDns) GetDnsQueryCaptureFileTimeLimit() int64`

GetDnsQueryCaptureFileTimeLimit returns the DnsQueryCaptureFileTimeLimit field if non-nil, zero value otherwise.

### GetDnsQueryCaptureFileTimeLimitOk

`func (o *MemberDns) GetDnsQueryCaptureFileTimeLimitOk() (*int64, bool)`

GetDnsQueryCaptureFileTimeLimitOk returns a tuple with the DnsQueryCaptureFileTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueryCaptureFileTimeLimit

`func (o *MemberDns) SetDnsQueryCaptureFileTimeLimit(v int64)`

SetDnsQueryCaptureFileTimeLimit sets DnsQueryCaptureFileTimeLimit field to given value.

### HasDnsQueryCaptureFileTimeLimit

`func (o *MemberDns) HasDnsQueryCaptureFileTimeLimit() bool`

HasDnsQueryCaptureFileTimeLimit returns a boolean if a field has been set.

### GetDnsQuerySourceAddress

`func (o *MemberDns) GetDnsQuerySourceAddress() string`

GetDnsQuerySourceAddress returns the DnsQuerySourceAddress field if non-nil, zero value otherwise.

### GetDnsQuerySourceAddressOk

`func (o *MemberDns) GetDnsQuerySourceAddressOk() (*string, bool)`

GetDnsQuerySourceAddressOk returns a tuple with the DnsQuerySourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQuerySourceAddress

`func (o *MemberDns) SetDnsQuerySourceAddress(v string)`

SetDnsQuerySourceAddress sets DnsQuerySourceAddress field to given value.

### HasDnsQuerySourceAddress

`func (o *MemberDns) HasDnsQuerySourceAddress() bool`

HasDnsQuerySourceAddress returns a boolean if a field has been set.

### GetDnsQuerySourceInterface

`func (o *MemberDns) GetDnsQuerySourceInterface() string`

GetDnsQuerySourceInterface returns the DnsQuerySourceInterface field if non-nil, zero value otherwise.

### GetDnsQuerySourceInterfaceOk

`func (o *MemberDns) GetDnsQuerySourceInterfaceOk() (*string, bool)`

GetDnsQuerySourceInterfaceOk returns a tuple with the DnsQuerySourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQuerySourceInterface

`func (o *MemberDns) SetDnsQuerySourceInterface(v string)`

SetDnsQuerySourceInterface sets DnsQuerySourceInterface field to given value.

### HasDnsQuerySourceInterface

`func (o *MemberDns) HasDnsQuerySourceInterface() bool`

HasDnsQuerySourceInterface returns a boolean if a field has been set.

### GetDnsViewAddressSettings

`func (o *MemberDns) GetDnsViewAddressSettings() []MemberDnsDnsViewAddressSettings`

GetDnsViewAddressSettings returns the DnsViewAddressSettings field if non-nil, zero value otherwise.

### GetDnsViewAddressSettingsOk

`func (o *MemberDns) GetDnsViewAddressSettingsOk() (*[]MemberDnsDnsViewAddressSettings, bool)`

GetDnsViewAddressSettingsOk returns a tuple with the DnsViewAddressSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsViewAddressSettings

`func (o *MemberDns) SetDnsViewAddressSettings(v []MemberDnsDnsViewAddressSettings)`

SetDnsViewAddressSettings sets DnsViewAddressSettings field to given value.

### HasDnsViewAddressSettings

`func (o *MemberDns) HasDnsViewAddressSettings() bool`

HasDnsViewAddressSettings returns a boolean if a field has been set.

### GetDnssecBlacklistEnabled

`func (o *MemberDns) GetDnssecBlacklistEnabled() bool`

GetDnssecBlacklistEnabled returns the DnssecBlacklistEnabled field if non-nil, zero value otherwise.

### GetDnssecBlacklistEnabledOk

`func (o *MemberDns) GetDnssecBlacklistEnabledOk() (*bool, bool)`

GetDnssecBlacklistEnabledOk returns a tuple with the DnssecBlacklistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecBlacklistEnabled

`func (o *MemberDns) SetDnssecBlacklistEnabled(v bool)`

SetDnssecBlacklistEnabled sets DnssecBlacklistEnabled field to given value.

### HasDnssecBlacklistEnabled

`func (o *MemberDns) HasDnssecBlacklistEnabled() bool`

HasDnssecBlacklistEnabled returns a boolean if a field has been set.

### GetDnssecDns64Enabled

`func (o *MemberDns) GetDnssecDns64Enabled() bool`

GetDnssecDns64Enabled returns the DnssecDns64Enabled field if non-nil, zero value otherwise.

### GetDnssecDns64EnabledOk

`func (o *MemberDns) GetDnssecDns64EnabledOk() (*bool, bool)`

GetDnssecDns64EnabledOk returns a tuple with the DnssecDns64Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecDns64Enabled

`func (o *MemberDns) SetDnssecDns64Enabled(v bool)`

SetDnssecDns64Enabled sets DnssecDns64Enabled field to given value.

### HasDnssecDns64Enabled

`func (o *MemberDns) HasDnssecDns64Enabled() bool`

HasDnssecDns64Enabled returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *MemberDns) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *MemberDns) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *MemberDns) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *MemberDns) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecExpiredSignaturesEnabled

`func (o *MemberDns) GetDnssecExpiredSignaturesEnabled() bool`

GetDnssecExpiredSignaturesEnabled returns the DnssecExpiredSignaturesEnabled field if non-nil, zero value otherwise.

### GetDnssecExpiredSignaturesEnabledOk

`func (o *MemberDns) GetDnssecExpiredSignaturesEnabledOk() (*bool, bool)`

GetDnssecExpiredSignaturesEnabledOk returns a tuple with the DnssecExpiredSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecExpiredSignaturesEnabled

`func (o *MemberDns) SetDnssecExpiredSignaturesEnabled(v bool)`

SetDnssecExpiredSignaturesEnabled sets DnssecExpiredSignaturesEnabled field to given value.

### HasDnssecExpiredSignaturesEnabled

`func (o *MemberDns) HasDnssecExpiredSignaturesEnabled() bool`

HasDnssecExpiredSignaturesEnabled returns a boolean if a field has been set.

### GetDnssecNegativeTrustAnchors

`func (o *MemberDns) GetDnssecNegativeTrustAnchors() []string`

GetDnssecNegativeTrustAnchors returns the DnssecNegativeTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecNegativeTrustAnchorsOk

`func (o *MemberDns) GetDnssecNegativeTrustAnchorsOk() (*[]string, bool)`

GetDnssecNegativeTrustAnchorsOk returns a tuple with the DnssecNegativeTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNegativeTrustAnchors

`func (o *MemberDns) SetDnssecNegativeTrustAnchors(v []string)`

SetDnssecNegativeTrustAnchors sets DnssecNegativeTrustAnchors field to given value.

### HasDnssecNegativeTrustAnchors

`func (o *MemberDns) HasDnssecNegativeTrustAnchors() bool`

HasDnssecNegativeTrustAnchors returns a boolean if a field has been set.

### GetDnssecNxdomainEnabled

`func (o *MemberDns) GetDnssecNxdomainEnabled() bool`

GetDnssecNxdomainEnabled returns the DnssecNxdomainEnabled field if non-nil, zero value otherwise.

### GetDnssecNxdomainEnabledOk

`func (o *MemberDns) GetDnssecNxdomainEnabledOk() (*bool, bool)`

GetDnssecNxdomainEnabledOk returns a tuple with the DnssecNxdomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNxdomainEnabled

`func (o *MemberDns) SetDnssecNxdomainEnabled(v bool)`

SetDnssecNxdomainEnabled sets DnssecNxdomainEnabled field to given value.

### HasDnssecNxdomainEnabled

`func (o *MemberDns) HasDnssecNxdomainEnabled() bool`

HasDnssecNxdomainEnabled returns a boolean if a field has been set.

### GetDnssecRpzEnabled

`func (o *MemberDns) GetDnssecRpzEnabled() bool`

GetDnssecRpzEnabled returns the DnssecRpzEnabled field if non-nil, zero value otherwise.

### GetDnssecRpzEnabledOk

`func (o *MemberDns) GetDnssecRpzEnabledOk() (*bool, bool)`

GetDnssecRpzEnabledOk returns a tuple with the DnssecRpzEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRpzEnabled

`func (o *MemberDns) SetDnssecRpzEnabled(v bool)`

SetDnssecRpzEnabled sets DnssecRpzEnabled field to given value.

### HasDnssecRpzEnabled

`func (o *MemberDns) HasDnssecRpzEnabled() bool`

HasDnssecRpzEnabled returns a boolean if a field has been set.

### GetDnssecTrustedKeys

`func (o *MemberDns) GetDnssecTrustedKeys() []MemberDnsDnssecTrustedKeys`

GetDnssecTrustedKeys returns the DnssecTrustedKeys field if non-nil, zero value otherwise.

### GetDnssecTrustedKeysOk

`func (o *MemberDns) GetDnssecTrustedKeysOk() (*[]MemberDnsDnssecTrustedKeys, bool)`

GetDnssecTrustedKeysOk returns a tuple with the DnssecTrustedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustedKeys

`func (o *MemberDns) SetDnssecTrustedKeys(v []MemberDnsDnssecTrustedKeys)`

SetDnssecTrustedKeys sets DnssecTrustedKeys field to given value.

### HasDnssecTrustedKeys

`func (o *MemberDns) HasDnssecTrustedKeys() bool`

HasDnssecTrustedKeys returns a boolean if a field has been set.

### GetDnssecValidationEnabled

`func (o *MemberDns) GetDnssecValidationEnabled() bool`

GetDnssecValidationEnabled returns the DnssecValidationEnabled field if non-nil, zero value otherwise.

### GetDnssecValidationEnabledOk

`func (o *MemberDns) GetDnssecValidationEnabledOk() (*bool, bool)`

GetDnssecValidationEnabledOk returns a tuple with the DnssecValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationEnabled

`func (o *MemberDns) SetDnssecValidationEnabled(v bool)`

SetDnssecValidationEnabled sets DnssecValidationEnabled field to given value.

### HasDnssecValidationEnabled

`func (o *MemberDns) HasDnssecValidationEnabled() bool`

HasDnssecValidationEnabled returns a boolean if a field has been set.

### GetDnstapSetting

`func (o *MemberDns) GetDnstapSetting() MemberDnsDnstapSetting`

GetDnstapSetting returns the DnstapSetting field if non-nil, zero value otherwise.

### GetDnstapSettingOk

`func (o *MemberDns) GetDnstapSettingOk() (*MemberDnsDnstapSetting, bool)`

GetDnstapSettingOk returns a tuple with the DnstapSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapSetting

`func (o *MemberDns) SetDnstapSetting(v MemberDnsDnstapSetting)`

SetDnstapSetting sets DnstapSetting field to given value.

### HasDnstapSetting

`func (o *MemberDns) HasDnstapSetting() bool`

HasDnstapSetting returns a boolean if a field has been set.

### GetDohHttpsSessionDuration

`func (o *MemberDns) GetDohHttpsSessionDuration() int64`

GetDohHttpsSessionDuration returns the DohHttpsSessionDuration field if non-nil, zero value otherwise.

### GetDohHttpsSessionDurationOk

`func (o *MemberDns) GetDohHttpsSessionDurationOk() (*int64, bool)`

GetDohHttpsSessionDurationOk returns a tuple with the DohHttpsSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDohHttpsSessionDuration

`func (o *MemberDns) SetDohHttpsSessionDuration(v int64)`

SetDohHttpsSessionDuration sets DohHttpsSessionDuration field to given value.

### HasDohHttpsSessionDuration

`func (o *MemberDns) HasDohHttpsSessionDuration() bool`

HasDohHttpsSessionDuration returns a boolean if a field has been set.

### GetDohService

`func (o *MemberDns) GetDohService() bool`

GetDohService returns the DohService field if non-nil, zero value otherwise.

### GetDohServiceOk

`func (o *MemberDns) GetDohServiceOk() (*bool, bool)`

GetDohServiceOk returns a tuple with the DohService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDohService

`func (o *MemberDns) SetDohService(v bool)`

SetDohService sets DohService field to given value.

### HasDohService

`func (o *MemberDns) HasDohService() bool`

HasDohService returns a boolean if a field has been set.

### GetDomainsToCaptureDnsQueries

`func (o *MemberDns) GetDomainsToCaptureDnsQueries() []string`

GetDomainsToCaptureDnsQueries returns the DomainsToCaptureDnsQueries field if non-nil, zero value otherwise.

### GetDomainsToCaptureDnsQueriesOk

`func (o *MemberDns) GetDomainsToCaptureDnsQueriesOk() (*[]string, bool)`

GetDomainsToCaptureDnsQueriesOk returns a tuple with the DomainsToCaptureDnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsToCaptureDnsQueries

`func (o *MemberDns) SetDomainsToCaptureDnsQueries(v []string)`

SetDomainsToCaptureDnsQueries sets DomainsToCaptureDnsQueries field to given value.

### HasDomainsToCaptureDnsQueries

`func (o *MemberDns) HasDomainsToCaptureDnsQueries() bool`

HasDomainsToCaptureDnsQueries returns a boolean if a field has been set.

### GetDtcDnsQueriesSpecificBehavior

`func (o *MemberDns) GetDtcDnsQueriesSpecificBehavior() string`

GetDtcDnsQueriesSpecificBehavior returns the DtcDnsQueriesSpecificBehavior field if non-nil, zero value otherwise.

### GetDtcDnsQueriesSpecificBehaviorOk

`func (o *MemberDns) GetDtcDnsQueriesSpecificBehaviorOk() (*string, bool)`

GetDtcDnsQueriesSpecificBehaviorOk returns a tuple with the DtcDnsQueriesSpecificBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcDnsQueriesSpecificBehavior

`func (o *MemberDns) SetDtcDnsQueriesSpecificBehavior(v string)`

SetDtcDnsQueriesSpecificBehavior sets DtcDnsQueriesSpecificBehavior field to given value.

### HasDtcDnsQueriesSpecificBehavior

`func (o *MemberDns) HasDtcDnsQueriesSpecificBehavior() bool`

HasDtcDnsQueriesSpecificBehavior returns a boolean if a field has been set.

### GetDtcEdnsPreferClientSubnet

`func (o *MemberDns) GetDtcEdnsPreferClientSubnet() bool`

GetDtcEdnsPreferClientSubnet returns the DtcEdnsPreferClientSubnet field if non-nil, zero value otherwise.

### GetDtcEdnsPreferClientSubnetOk

`func (o *MemberDns) GetDtcEdnsPreferClientSubnetOk() (*bool, bool)`

GetDtcEdnsPreferClientSubnetOk returns a tuple with the DtcEdnsPreferClientSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcEdnsPreferClientSubnet

`func (o *MemberDns) SetDtcEdnsPreferClientSubnet(v bool)`

SetDtcEdnsPreferClientSubnet sets DtcEdnsPreferClientSubnet field to given value.

### HasDtcEdnsPreferClientSubnet

`func (o *MemberDns) HasDtcEdnsPreferClientSubnet() bool`

HasDtcEdnsPreferClientSubnet returns a boolean if a field has been set.

### GetDtcHealthSource

`func (o *MemberDns) GetDtcHealthSource() string`

GetDtcHealthSource returns the DtcHealthSource field if non-nil, zero value otherwise.

### GetDtcHealthSourceOk

`func (o *MemberDns) GetDtcHealthSourceOk() (*string, bool)`

GetDtcHealthSourceOk returns a tuple with the DtcHealthSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcHealthSource

`func (o *MemberDns) SetDtcHealthSource(v string)`

SetDtcHealthSource sets DtcHealthSource field to given value.

### HasDtcHealthSource

`func (o *MemberDns) HasDtcHealthSource() bool`

HasDtcHealthSource returns a boolean if a field has been set.

### GetDtcHealthSourceAddress

`func (o *MemberDns) GetDtcHealthSourceAddress() string`

GetDtcHealthSourceAddress returns the DtcHealthSourceAddress field if non-nil, zero value otherwise.

### GetDtcHealthSourceAddressOk

`func (o *MemberDns) GetDtcHealthSourceAddressOk() (*string, bool)`

GetDtcHealthSourceAddressOk returns a tuple with the DtcHealthSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcHealthSourceAddress

`func (o *MemberDns) SetDtcHealthSourceAddress(v string)`

SetDtcHealthSourceAddress sets DtcHealthSourceAddress field to given value.

### HasDtcHealthSourceAddress

`func (o *MemberDns) HasDtcHealthSourceAddress() bool`

HasDtcHealthSourceAddress returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *MemberDns) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *MemberDns) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *MemberDns) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *MemberDns) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetEnableBlackhole

`func (o *MemberDns) GetEnableBlackhole() bool`

GetEnableBlackhole returns the EnableBlackhole field if non-nil, zero value otherwise.

### GetEnableBlackholeOk

`func (o *MemberDns) GetEnableBlackholeOk() (*bool, bool)`

GetEnableBlackholeOk returns a tuple with the EnableBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackhole

`func (o *MemberDns) SetEnableBlackhole(v bool)`

SetEnableBlackhole sets EnableBlackhole field to given value.

### HasEnableBlackhole

`func (o *MemberDns) HasEnableBlackhole() bool`

HasEnableBlackhole returns a boolean if a field has been set.

### GetEnableBlacklist

`func (o *MemberDns) GetEnableBlacklist() bool`

GetEnableBlacklist returns the EnableBlacklist field if non-nil, zero value otherwise.

### GetEnableBlacklistOk

`func (o *MemberDns) GetEnableBlacklistOk() (*bool, bool)`

GetEnableBlacklistOk returns a tuple with the EnableBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlacklist

`func (o *MemberDns) SetEnableBlacklist(v bool)`

SetEnableBlacklist sets EnableBlacklist field to given value.

### HasEnableBlacklist

`func (o *MemberDns) HasEnableBlacklist() bool`

HasEnableBlacklist returns a boolean if a field has been set.

### GetEnableCaptureDnsQueries

`func (o *MemberDns) GetEnableCaptureDnsQueries() bool`

GetEnableCaptureDnsQueries returns the EnableCaptureDnsQueries field if non-nil, zero value otherwise.

### GetEnableCaptureDnsQueriesOk

`func (o *MemberDns) GetEnableCaptureDnsQueriesOk() (*bool, bool)`

GetEnableCaptureDnsQueriesOk returns a tuple with the EnableCaptureDnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaptureDnsQueries

`func (o *MemberDns) SetEnableCaptureDnsQueries(v bool)`

SetEnableCaptureDnsQueries sets EnableCaptureDnsQueries field to given value.

### HasEnableCaptureDnsQueries

`func (o *MemberDns) HasEnableCaptureDnsQueries() bool`

HasEnableCaptureDnsQueries returns a boolean if a field has been set.

### GetEnableCaptureDnsResponses

`func (o *MemberDns) GetEnableCaptureDnsResponses() bool`

GetEnableCaptureDnsResponses returns the EnableCaptureDnsResponses field if non-nil, zero value otherwise.

### GetEnableCaptureDnsResponsesOk

`func (o *MemberDns) GetEnableCaptureDnsResponsesOk() (*bool, bool)`

GetEnableCaptureDnsResponsesOk returns a tuple with the EnableCaptureDnsResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaptureDnsResponses

`func (o *MemberDns) SetEnableCaptureDnsResponses(v bool)`

SetEnableCaptureDnsResponses sets EnableCaptureDnsResponses field to given value.

### HasEnableCaptureDnsResponses

`func (o *MemberDns) HasEnableCaptureDnsResponses() bool`

HasEnableCaptureDnsResponses returns a boolean if a field has been set.

### GetEnableDns

`func (o *MemberDns) GetEnableDns() bool`

GetEnableDns returns the EnableDns field if non-nil, zero value otherwise.

### GetEnableDnsOk

`func (o *MemberDns) GetEnableDnsOk() (*bool, bool)`

GetEnableDnsOk returns a tuple with the EnableDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns

`func (o *MemberDns) SetEnableDns(v bool)`

SetEnableDns sets EnableDns field to given value.

### HasEnableDns

`func (o *MemberDns) HasEnableDns() bool`

HasEnableDns returns a boolean if a field has been set.

### GetEnableDns64

`func (o *MemberDns) GetEnableDns64() bool`

GetEnableDns64 returns the EnableDns64 field if non-nil, zero value otherwise.

### GetEnableDns64Ok

`func (o *MemberDns) GetEnableDns64Ok() (*bool, bool)`

GetEnableDns64Ok returns a tuple with the EnableDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns64

`func (o *MemberDns) SetEnableDns64(v bool)`

SetEnableDns64 sets EnableDns64 field to given value.

### HasEnableDns64

`func (o *MemberDns) HasEnableDns64() bool`

HasEnableDns64 returns a boolean if a field has been set.

### GetEnableDnsCacheAcceleration

`func (o *MemberDns) GetEnableDnsCacheAcceleration() bool`

GetEnableDnsCacheAcceleration returns the EnableDnsCacheAcceleration field if non-nil, zero value otherwise.

### GetEnableDnsCacheAccelerationOk

`func (o *MemberDns) GetEnableDnsCacheAccelerationOk() (*bool, bool)`

GetEnableDnsCacheAccelerationOk returns a tuple with the EnableDnsCacheAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsCacheAcceleration

`func (o *MemberDns) SetEnableDnsCacheAcceleration(v bool)`

SetEnableDnsCacheAcceleration sets EnableDnsCacheAcceleration field to given value.

### HasEnableDnsCacheAcceleration

`func (o *MemberDns) HasEnableDnsCacheAcceleration() bool`

HasEnableDnsCacheAcceleration returns a boolean if a field has been set.

### GetEnableDnsHealthCheck

`func (o *MemberDns) GetEnableDnsHealthCheck() bool`

GetEnableDnsHealthCheck returns the EnableDnsHealthCheck field if non-nil, zero value otherwise.

### GetEnableDnsHealthCheckOk

`func (o *MemberDns) GetEnableDnsHealthCheckOk() (*bool, bool)`

GetEnableDnsHealthCheckOk returns a tuple with the EnableDnsHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsHealthCheck

`func (o *MemberDns) SetEnableDnsHealthCheck(v bool)`

SetEnableDnsHealthCheck sets EnableDnsHealthCheck field to given value.

### HasEnableDnsHealthCheck

`func (o *MemberDns) HasEnableDnsHealthCheck() bool`

HasEnableDnsHealthCheck returns a boolean if a field has been set.

### GetEnableDnstapQueries

`func (o *MemberDns) GetEnableDnstapQueries() bool`

GetEnableDnstapQueries returns the EnableDnstapQueries field if non-nil, zero value otherwise.

### GetEnableDnstapQueriesOk

`func (o *MemberDns) GetEnableDnstapQueriesOk() (*bool, bool)`

GetEnableDnstapQueriesOk returns a tuple with the EnableDnstapQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapQueries

`func (o *MemberDns) SetEnableDnstapQueries(v bool)`

SetEnableDnstapQueries sets EnableDnstapQueries field to given value.

### HasEnableDnstapQueries

`func (o *MemberDns) HasEnableDnstapQueries() bool`

HasEnableDnstapQueries returns a boolean if a field has been set.

### GetEnableDnstapResponses

`func (o *MemberDns) GetEnableDnstapResponses() bool`

GetEnableDnstapResponses returns the EnableDnstapResponses field if non-nil, zero value otherwise.

### GetEnableDnstapResponsesOk

`func (o *MemberDns) GetEnableDnstapResponsesOk() (*bool, bool)`

GetEnableDnstapResponsesOk returns a tuple with the EnableDnstapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapResponses

`func (o *MemberDns) SetEnableDnstapResponses(v bool)`

SetEnableDnstapResponses sets EnableDnstapResponses field to given value.

### HasEnableDnstapResponses

`func (o *MemberDns) HasEnableDnstapResponses() bool`

HasEnableDnstapResponses returns a boolean if a field has been set.

### GetEnableDnstapViolationsTls

`func (o *MemberDns) GetEnableDnstapViolationsTls() bool`

GetEnableDnstapViolationsTls returns the EnableDnstapViolationsTls field if non-nil, zero value otherwise.

### GetEnableDnstapViolationsTlsOk

`func (o *MemberDns) GetEnableDnstapViolationsTlsOk() (*bool, bool)`

GetEnableDnstapViolationsTlsOk returns a tuple with the EnableDnstapViolationsTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapViolationsTls

`func (o *MemberDns) SetEnableDnstapViolationsTls(v bool)`

SetEnableDnstapViolationsTls sets EnableDnstapViolationsTls field to given value.

### HasEnableDnstapViolationsTls

`func (o *MemberDns) HasEnableDnstapViolationsTls() bool`

HasEnableDnstapViolationsTls returns a boolean if a field has been set.

### GetEnableExcludedDomainNames

`func (o *MemberDns) GetEnableExcludedDomainNames() bool`

GetEnableExcludedDomainNames returns the EnableExcludedDomainNames field if non-nil, zero value otherwise.

### GetEnableExcludedDomainNamesOk

`func (o *MemberDns) GetEnableExcludedDomainNamesOk() (*bool, bool)`

GetEnableExcludedDomainNamesOk returns a tuple with the EnableExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExcludedDomainNames

`func (o *MemberDns) SetEnableExcludedDomainNames(v bool)`

SetEnableExcludedDomainNames sets EnableExcludedDomainNames field to given value.

### HasEnableExcludedDomainNames

`func (o *MemberDns) HasEnableExcludedDomainNames() bool`

HasEnableExcludedDomainNames returns a boolean if a field has been set.

### GetEnableFixedRrsetOrderFqdns

`func (o *MemberDns) GetEnableFixedRrsetOrderFqdns() bool`

GetEnableFixedRrsetOrderFqdns returns the EnableFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetEnableFixedRrsetOrderFqdnsOk

`func (o *MemberDns) GetEnableFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetEnableFixedRrsetOrderFqdnsOk returns a tuple with the EnableFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixedRrsetOrderFqdns

`func (o *MemberDns) SetEnableFixedRrsetOrderFqdns(v bool)`

SetEnableFixedRrsetOrderFqdns sets EnableFixedRrsetOrderFqdns field to given value.

### HasEnableFixedRrsetOrderFqdns

`func (o *MemberDns) HasEnableFixedRrsetOrderFqdns() bool`

HasEnableFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetEnableFtc

`func (o *MemberDns) GetEnableFtc() bool`

GetEnableFtc returns the EnableFtc field if non-nil, zero value otherwise.

### GetEnableFtcOk

`func (o *MemberDns) GetEnableFtcOk() (*bool, bool)`

GetEnableFtcOk returns a tuple with the EnableFtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtc

`func (o *MemberDns) SetEnableFtc(v bool)`

SetEnableFtc sets EnableFtc field to given value.

### HasEnableFtc

`func (o *MemberDns) HasEnableFtc() bool`

HasEnableFtc returns a boolean if a field has been set.

### GetEnableGssTsig

`func (o *MemberDns) GetEnableGssTsig() bool`

GetEnableGssTsig returns the EnableGssTsig field if non-nil, zero value otherwise.

### GetEnableGssTsigOk

`func (o *MemberDns) GetEnableGssTsigOk() (*bool, bool)`

GetEnableGssTsigOk returns a tuple with the EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGssTsig

`func (o *MemberDns) SetEnableGssTsig(v bool)`

SetEnableGssTsig sets EnableGssTsig field to given value.

### HasEnableGssTsig

`func (o *MemberDns) HasEnableGssTsig() bool`

HasEnableGssTsig returns a boolean if a field has been set.

### GetEnableNotifySourcePort

`func (o *MemberDns) GetEnableNotifySourcePort() bool`

GetEnableNotifySourcePort returns the EnableNotifySourcePort field if non-nil, zero value otherwise.

### GetEnableNotifySourcePortOk

`func (o *MemberDns) GetEnableNotifySourcePortOk() (*bool, bool)`

GetEnableNotifySourcePortOk returns a tuple with the EnableNotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySourcePort

`func (o *MemberDns) SetEnableNotifySourcePort(v bool)`

SetEnableNotifySourcePort sets EnableNotifySourcePort field to given value.

### HasEnableNotifySourcePort

`func (o *MemberDns) HasEnableNotifySourcePort() bool`

HasEnableNotifySourcePort returns a boolean if a field has been set.

### GetEnableQueryRewrite

`func (o *MemberDns) GetEnableQueryRewrite() bool`

GetEnableQueryRewrite returns the EnableQueryRewrite field if non-nil, zero value otherwise.

### GetEnableQueryRewriteOk

`func (o *MemberDns) GetEnableQueryRewriteOk() (*bool, bool)`

GetEnableQueryRewriteOk returns a tuple with the EnableQueryRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryRewrite

`func (o *MemberDns) SetEnableQueryRewrite(v bool)`

SetEnableQueryRewrite sets EnableQueryRewrite field to given value.

### HasEnableQueryRewrite

`func (o *MemberDns) HasEnableQueryRewrite() bool`

HasEnableQueryRewrite returns a boolean if a field has been set.

### GetEnableQuerySourcePort

`func (o *MemberDns) GetEnableQuerySourcePort() bool`

GetEnableQuerySourcePort returns the EnableQuerySourcePort field if non-nil, zero value otherwise.

### GetEnableQuerySourcePortOk

`func (o *MemberDns) GetEnableQuerySourcePortOk() (*bool, bool)`

GetEnableQuerySourcePortOk returns a tuple with the EnableQuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuerySourcePort

`func (o *MemberDns) SetEnableQuerySourcePort(v bool)`

SetEnableQuerySourcePort sets EnableQuerySourcePort field to given value.

### HasEnableQuerySourcePort

`func (o *MemberDns) HasEnableQuerySourcePort() bool`

HasEnableQuerySourcePort returns a boolean if a field has been set.

### GetExcludedDomainNames

`func (o *MemberDns) GetExcludedDomainNames() []string`

GetExcludedDomainNames returns the ExcludedDomainNames field if non-nil, zero value otherwise.

### GetExcludedDomainNamesOk

`func (o *MemberDns) GetExcludedDomainNamesOk() (*[]string, bool)`

GetExcludedDomainNamesOk returns a tuple with the ExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDomainNames

`func (o *MemberDns) SetExcludedDomainNames(v []string)`

SetExcludedDomainNames sets ExcludedDomainNames field to given value.

### HasExcludedDomainNames

`func (o *MemberDns) HasExcludedDomainNames() bool`

HasExcludedDomainNames returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *MemberDns) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *MemberDns) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *MemberDns) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *MemberDns) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *MemberDns) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *MemberDns) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *MemberDns) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *MemberDns) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *MemberDns) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *MemberDns) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *MemberDns) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *MemberDns) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFileTransferSetting

`func (o *MemberDns) GetFileTransferSetting() MemberDnsFileTransferSetting`

GetFileTransferSetting returns the FileTransferSetting field if non-nil, zero value otherwise.

### GetFileTransferSettingOk

`func (o *MemberDns) GetFileTransferSettingOk() (*MemberDnsFileTransferSetting, bool)`

GetFileTransferSettingOk returns a tuple with the FileTransferSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferSetting

`func (o *MemberDns) SetFileTransferSetting(v MemberDnsFileTransferSetting)`

SetFileTransferSetting sets FileTransferSetting field to given value.

### HasFileTransferSetting

`func (o *MemberDns) HasFileTransferSetting() bool`

HasFileTransferSetting returns a boolean if a field has been set.

### GetFilterAaaa

`func (o *MemberDns) GetFilterAaaa() string`

GetFilterAaaa returns the FilterAaaa field if non-nil, zero value otherwise.

### GetFilterAaaaOk

`func (o *MemberDns) GetFilterAaaaOk() (*string, bool)`

GetFilterAaaaOk returns a tuple with the FilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaa

`func (o *MemberDns) SetFilterAaaa(v string)`

SetFilterAaaa sets FilterAaaa field to given value.

### HasFilterAaaa

`func (o *MemberDns) HasFilterAaaa() bool`

HasFilterAaaa returns a boolean if a field has been set.

### GetFilterAaaaList

`func (o *MemberDns) GetFilterAaaaList() []MemberDnsFilterAaaaList`

GetFilterAaaaList returns the FilterAaaaList field if non-nil, zero value otherwise.

### GetFilterAaaaListOk

`func (o *MemberDns) GetFilterAaaaListOk() (*[]MemberDnsFilterAaaaList, bool)`

GetFilterAaaaListOk returns a tuple with the FilterAaaaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaList

`func (o *MemberDns) SetFilterAaaaList(v []MemberDnsFilterAaaaList)`

SetFilterAaaaList sets FilterAaaaList field to given value.

### HasFilterAaaaList

`func (o *MemberDns) HasFilterAaaaList() bool`

HasFilterAaaaList returns a boolean if a field has been set.

### GetFixedRrsetOrderFqdns

`func (o *MemberDns) GetFixedRrsetOrderFqdns() []MemberDnsFixedRrsetOrderFqdns`

GetFixedRrsetOrderFqdns returns the FixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetFixedRrsetOrderFqdnsOk

`func (o *MemberDns) GetFixedRrsetOrderFqdnsOk() (*[]MemberDnsFixedRrsetOrderFqdns, bool)`

GetFixedRrsetOrderFqdnsOk returns a tuple with the FixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedRrsetOrderFqdns

`func (o *MemberDns) SetFixedRrsetOrderFqdns(v []MemberDnsFixedRrsetOrderFqdns)`

SetFixedRrsetOrderFqdns sets FixedRrsetOrderFqdns field to given value.

### HasFixedRrsetOrderFqdns

`func (o *MemberDns) HasFixedRrsetOrderFqdns() bool`

HasFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetForwardOnly

`func (o *MemberDns) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *MemberDns) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *MemberDns) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *MemberDns) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetForwardUpdates

`func (o *MemberDns) GetForwardUpdates() bool`

GetForwardUpdates returns the ForwardUpdates field if non-nil, zero value otherwise.

### GetForwardUpdatesOk

`func (o *MemberDns) GetForwardUpdatesOk() (*bool, bool)`

GetForwardUpdatesOk returns a tuple with the ForwardUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUpdates

`func (o *MemberDns) SetForwardUpdates(v bool)`

SetForwardUpdates sets ForwardUpdates field to given value.

### HasForwardUpdates

`func (o *MemberDns) HasForwardUpdates() bool`

HasForwardUpdates returns a boolean if a field has been set.

### GetForwarders

`func (o *MemberDns) GetForwarders() []string`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *MemberDns) GetForwardersOk() (*[]string, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *MemberDns) SetForwarders(v []string)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *MemberDns) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetFtcExpiredRecordTimeout

`func (o *MemberDns) GetFtcExpiredRecordTimeout() int64`

GetFtcExpiredRecordTimeout returns the FtcExpiredRecordTimeout field if non-nil, zero value otherwise.

### GetFtcExpiredRecordTimeoutOk

`func (o *MemberDns) GetFtcExpiredRecordTimeoutOk() (*int64, bool)`

GetFtcExpiredRecordTimeoutOk returns a tuple with the FtcExpiredRecordTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtcExpiredRecordTimeout

`func (o *MemberDns) SetFtcExpiredRecordTimeout(v int64)`

SetFtcExpiredRecordTimeout sets FtcExpiredRecordTimeout field to given value.

### HasFtcExpiredRecordTimeout

`func (o *MemberDns) HasFtcExpiredRecordTimeout() bool`

HasFtcExpiredRecordTimeout returns a boolean if a field has been set.

### GetFtcExpiredRecordTtl

`func (o *MemberDns) GetFtcExpiredRecordTtl() int64`

GetFtcExpiredRecordTtl returns the FtcExpiredRecordTtl field if non-nil, zero value otherwise.

### GetFtcExpiredRecordTtlOk

`func (o *MemberDns) GetFtcExpiredRecordTtlOk() (*int64, bool)`

GetFtcExpiredRecordTtlOk returns a tuple with the FtcExpiredRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtcExpiredRecordTtl

`func (o *MemberDns) SetFtcExpiredRecordTtl(v int64)`

SetFtcExpiredRecordTtl sets FtcExpiredRecordTtl field to given value.

### HasFtcExpiredRecordTtl

`func (o *MemberDns) HasFtcExpiredRecordTtl() bool`

HasFtcExpiredRecordTtl returns a boolean if a field has been set.

### GetGlueRecordAddresses

`func (o *MemberDns) GetGlueRecordAddresses() []MemberDnsGlueRecordAddresses`

GetGlueRecordAddresses returns the GlueRecordAddresses field if non-nil, zero value otherwise.

### GetGlueRecordAddressesOk

`func (o *MemberDns) GetGlueRecordAddressesOk() (*[]MemberDnsGlueRecordAddresses, bool)`

GetGlueRecordAddressesOk returns a tuple with the GlueRecordAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlueRecordAddresses

`func (o *MemberDns) SetGlueRecordAddresses(v []MemberDnsGlueRecordAddresses)`

SetGlueRecordAddresses sets GlueRecordAddresses field to given value.

### HasGlueRecordAddresses

`func (o *MemberDns) HasGlueRecordAddresses() bool`

HasGlueRecordAddresses returns a boolean if a field has been set.

### GetGssTsigKeys

`func (o *MemberDns) GetGssTsigKeys() []string`

GetGssTsigKeys returns the GssTsigKeys field if non-nil, zero value otherwise.

### GetGssTsigKeysOk

`func (o *MemberDns) GetGssTsigKeysOk() (*[]string, bool)`

GetGssTsigKeysOk returns a tuple with the GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigKeys

`func (o *MemberDns) SetGssTsigKeys(v []string)`

SetGssTsigKeys sets GssTsigKeys field to given value.

### HasGssTsigKeys

`func (o *MemberDns) HasGssTsigKeys() bool`

HasGssTsigKeys returns a boolean if a field has been set.

### GetHostName

`func (o *MemberDns) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *MemberDns) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *MemberDns) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *MemberDns) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpv4addr

`func (o *MemberDns) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *MemberDns) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *MemberDns) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *MemberDns) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6GlueRecordAddresses

`func (o *MemberDns) GetIpv6GlueRecordAddresses() []MemberDnsIpv6GlueRecordAddresses`

GetIpv6GlueRecordAddresses returns the Ipv6GlueRecordAddresses field if non-nil, zero value otherwise.

### GetIpv6GlueRecordAddressesOk

`func (o *MemberDns) GetIpv6GlueRecordAddressesOk() (*[]MemberDnsIpv6GlueRecordAddresses, bool)`

GetIpv6GlueRecordAddressesOk returns a tuple with the Ipv6GlueRecordAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6GlueRecordAddresses

`func (o *MemberDns) SetIpv6GlueRecordAddresses(v []MemberDnsIpv6GlueRecordAddresses)`

SetIpv6GlueRecordAddresses sets Ipv6GlueRecordAddresses field to given value.

### HasIpv6GlueRecordAddresses

`func (o *MemberDns) HasIpv6GlueRecordAddresses() bool`

HasIpv6GlueRecordAddresses returns a boolean if a field has been set.

### GetIpv6addr

`func (o *MemberDns) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *MemberDns) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *MemberDns) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *MemberDns) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetLoggingCategories

`func (o *MemberDns) GetLoggingCategories() MemberDnsLoggingCategories`

GetLoggingCategories returns the LoggingCategories field if non-nil, zero value otherwise.

### GetLoggingCategoriesOk

`func (o *MemberDns) GetLoggingCategoriesOk() (*MemberDnsLoggingCategories, bool)`

GetLoggingCategoriesOk returns a tuple with the LoggingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCategories

`func (o *MemberDns) SetLoggingCategories(v MemberDnsLoggingCategories)`

SetLoggingCategories sets LoggingCategories field to given value.

### HasLoggingCategories

`func (o *MemberDns) HasLoggingCategories() bool`

HasLoggingCategories returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *MemberDns) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *MemberDns) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *MemberDns) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *MemberDns) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxCachedLifetime

`func (o *MemberDns) GetMaxCachedLifetime() int64`

GetMaxCachedLifetime returns the MaxCachedLifetime field if non-nil, zero value otherwise.

### GetMaxCachedLifetimeOk

`func (o *MemberDns) GetMaxCachedLifetimeOk() (*int64, bool)`

GetMaxCachedLifetimeOk returns a tuple with the MaxCachedLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCachedLifetime

`func (o *MemberDns) SetMaxCachedLifetime(v int64)`

SetMaxCachedLifetime sets MaxCachedLifetime field to given value.

### HasMaxCachedLifetime

`func (o *MemberDns) HasMaxCachedLifetime() bool`

HasMaxCachedLifetime returns a boolean if a field has been set.

### GetMaxNcacheTtl

`func (o *MemberDns) GetMaxNcacheTtl() int64`

GetMaxNcacheTtl returns the MaxNcacheTtl field if non-nil, zero value otherwise.

### GetMaxNcacheTtlOk

`func (o *MemberDns) GetMaxNcacheTtlOk() (*int64, bool)`

GetMaxNcacheTtlOk returns a tuple with the MaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNcacheTtl

`func (o *MemberDns) SetMaxNcacheTtl(v int64)`

SetMaxNcacheTtl sets MaxNcacheTtl field to given value.

### HasMaxNcacheTtl

`func (o *MemberDns) HasMaxNcacheTtl() bool`

HasMaxNcacheTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *MemberDns) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *MemberDns) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *MemberDns) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *MemberDns) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResp

`func (o *MemberDns) GetMinimalResp() bool`

GetMinimalResp returns the MinimalResp field if non-nil, zero value otherwise.

### GetMinimalRespOk

`func (o *MemberDns) GetMinimalRespOk() (*bool, bool)`

GetMinimalRespOk returns a tuple with the MinimalResp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResp

`func (o *MemberDns) SetMinimalResp(v bool)`

SetMinimalResp sets MinimalResp field to given value.

### HasMinimalResp

`func (o *MemberDns) HasMinimalResp() bool`

HasMinimalResp returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *MemberDns) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *MemberDns) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *MemberDns) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *MemberDns) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetNotifySourcePort

`func (o *MemberDns) GetNotifySourcePort() int64`

GetNotifySourcePort returns the NotifySourcePort field if non-nil, zero value otherwise.

### GetNotifySourcePortOk

`func (o *MemberDns) GetNotifySourcePortOk() (*int64, bool)`

GetNotifySourcePortOk returns a tuple with the NotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySourcePort

`func (o *MemberDns) SetNotifySourcePort(v int64)`

SetNotifySourcePort sets NotifySourcePort field to given value.

### HasNotifySourcePort

`func (o *MemberDns) HasNotifySourcePort() bool`

HasNotifySourcePort returns a boolean if a field has been set.

### GetNxdomainLogQuery

`func (o *MemberDns) GetNxdomainLogQuery() bool`

GetNxdomainLogQuery returns the NxdomainLogQuery field if non-nil, zero value otherwise.

### GetNxdomainLogQueryOk

`func (o *MemberDns) GetNxdomainLogQueryOk() (*bool, bool)`

GetNxdomainLogQueryOk returns a tuple with the NxdomainLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainLogQuery

`func (o *MemberDns) SetNxdomainLogQuery(v bool)`

SetNxdomainLogQuery sets NxdomainLogQuery field to given value.

### HasNxdomainLogQuery

`func (o *MemberDns) HasNxdomainLogQuery() bool`

HasNxdomainLogQuery returns a boolean if a field has been set.

### GetNxdomainRedirect

`func (o *MemberDns) GetNxdomainRedirect() bool`

GetNxdomainRedirect returns the NxdomainRedirect field if non-nil, zero value otherwise.

### GetNxdomainRedirectOk

`func (o *MemberDns) GetNxdomainRedirectOk() (*bool, bool)`

GetNxdomainRedirectOk returns a tuple with the NxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirect

`func (o *MemberDns) SetNxdomainRedirect(v bool)`

SetNxdomainRedirect sets NxdomainRedirect field to given value.

### HasNxdomainRedirect

`func (o *MemberDns) HasNxdomainRedirect() bool`

HasNxdomainRedirect returns a boolean if a field has been set.

### GetNxdomainRedirectAddresses

`func (o *MemberDns) GetNxdomainRedirectAddresses() []string`

GetNxdomainRedirectAddresses returns the NxdomainRedirectAddresses field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesOk

`func (o *MemberDns) GetNxdomainRedirectAddressesOk() (*[]string, bool)`

GetNxdomainRedirectAddressesOk returns a tuple with the NxdomainRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddresses

`func (o *MemberDns) SetNxdomainRedirectAddresses(v []string)`

SetNxdomainRedirectAddresses sets NxdomainRedirectAddresses field to given value.

### HasNxdomainRedirectAddresses

`func (o *MemberDns) HasNxdomainRedirectAddresses() bool`

HasNxdomainRedirectAddresses returns a boolean if a field has been set.

### GetNxdomainRedirectAddressesV6

`func (o *MemberDns) GetNxdomainRedirectAddressesV6() []string`

GetNxdomainRedirectAddressesV6 returns the NxdomainRedirectAddressesV6 field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesV6Ok

`func (o *MemberDns) GetNxdomainRedirectAddressesV6Ok() (*[]string, bool)`

GetNxdomainRedirectAddressesV6Ok returns a tuple with the NxdomainRedirectAddressesV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddressesV6

`func (o *MemberDns) SetNxdomainRedirectAddressesV6(v []string)`

SetNxdomainRedirectAddressesV6 sets NxdomainRedirectAddressesV6 field to given value.

### HasNxdomainRedirectAddressesV6

`func (o *MemberDns) HasNxdomainRedirectAddressesV6() bool`

HasNxdomainRedirectAddressesV6 returns a boolean if a field has been set.

### GetNxdomainRedirectTtl

`func (o *MemberDns) GetNxdomainRedirectTtl() int64`

GetNxdomainRedirectTtl returns the NxdomainRedirectTtl field if non-nil, zero value otherwise.

### GetNxdomainRedirectTtlOk

`func (o *MemberDns) GetNxdomainRedirectTtlOk() (*int64, bool)`

GetNxdomainRedirectTtlOk returns a tuple with the NxdomainRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectTtl

`func (o *MemberDns) SetNxdomainRedirectTtl(v int64)`

SetNxdomainRedirectTtl sets NxdomainRedirectTtl field to given value.

### HasNxdomainRedirectTtl

`func (o *MemberDns) HasNxdomainRedirectTtl() bool`

HasNxdomainRedirectTtl returns a boolean if a field has been set.

### GetNxdomainRulesets

`func (o *MemberDns) GetNxdomainRulesets() []string`

GetNxdomainRulesets returns the NxdomainRulesets field if non-nil, zero value otherwise.

### GetNxdomainRulesetsOk

`func (o *MemberDns) GetNxdomainRulesetsOk() (*[]string, bool)`

GetNxdomainRulesetsOk returns a tuple with the NxdomainRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRulesets

`func (o *MemberDns) SetNxdomainRulesets(v []string)`

SetNxdomainRulesets sets NxdomainRulesets field to given value.

### HasNxdomainRulesets

`func (o *MemberDns) HasNxdomainRulesets() bool`

HasNxdomainRulesets returns a boolean if a field has been set.

### GetQuerySourcePort

`func (o *MemberDns) GetQuerySourcePort() int64`

GetQuerySourcePort returns the QuerySourcePort field if non-nil, zero value otherwise.

### GetQuerySourcePortOk

`func (o *MemberDns) GetQuerySourcePortOk() (*int64, bool)`

GetQuerySourcePortOk returns a tuple with the QuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySourcePort

`func (o *MemberDns) SetQuerySourcePort(v int64)`

SetQuerySourcePort sets QuerySourcePort field to given value.

### HasQuerySourcePort

`func (o *MemberDns) HasQuerySourcePort() bool`

HasQuerySourcePort returns a boolean if a field has been set.

### GetRecordNamePolicy

`func (o *MemberDns) GetRecordNamePolicy() string`

GetRecordNamePolicy returns the RecordNamePolicy field if non-nil, zero value otherwise.

### GetRecordNamePolicyOk

`func (o *MemberDns) GetRecordNamePolicyOk() (*string, bool)`

GetRecordNamePolicyOk returns a tuple with the RecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNamePolicy

`func (o *MemberDns) SetRecordNamePolicy(v string)`

SetRecordNamePolicy sets RecordNamePolicy field to given value.

### HasRecordNamePolicy

`func (o *MemberDns) HasRecordNamePolicy() bool`

HasRecordNamePolicy returns a boolean if a field has been set.

### GetRecursiveClientLimit

`func (o *MemberDns) GetRecursiveClientLimit() int64`

GetRecursiveClientLimit returns the RecursiveClientLimit field if non-nil, zero value otherwise.

### GetRecursiveClientLimitOk

`func (o *MemberDns) GetRecursiveClientLimitOk() (*int64, bool)`

GetRecursiveClientLimitOk returns a tuple with the RecursiveClientLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClientLimit

`func (o *MemberDns) SetRecursiveClientLimit(v int64)`

SetRecursiveClientLimit sets RecursiveClientLimit field to given value.

### HasRecursiveClientLimit

`func (o *MemberDns) HasRecursiveClientLimit() bool`

HasRecursiveClientLimit returns a boolean if a field has been set.

### GetRecursiveQueryList

`func (o *MemberDns) GetRecursiveQueryList() []MemberDnsRecursiveQueryList`

GetRecursiveQueryList returns the RecursiveQueryList field if non-nil, zero value otherwise.

### GetRecursiveQueryListOk

`func (o *MemberDns) GetRecursiveQueryListOk() (*[]MemberDnsRecursiveQueryList, bool)`

GetRecursiveQueryListOk returns a tuple with the RecursiveQueryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveQueryList

`func (o *MemberDns) SetRecursiveQueryList(v []MemberDnsRecursiveQueryList)`

SetRecursiveQueryList sets RecursiveQueryList field to given value.

### HasRecursiveQueryList

`func (o *MemberDns) HasRecursiveQueryList() bool`

HasRecursiveQueryList returns a boolean if a field has been set.

### GetRecursiveResolver

`func (o *MemberDns) GetRecursiveResolver() string`

GetRecursiveResolver returns the RecursiveResolver field if non-nil, zero value otherwise.

### GetRecursiveResolverOk

`func (o *MemberDns) GetRecursiveResolverOk() (*string, bool)`

GetRecursiveResolverOk returns a tuple with the RecursiveResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveResolver

`func (o *MemberDns) SetRecursiveResolver(v string)`

SetRecursiveResolver sets RecursiveResolver field to given value.

### HasRecursiveResolver

`func (o *MemberDns) HasRecursiveResolver() bool`

HasRecursiveResolver returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *MemberDns) GetResolverQueryTimeout() int64`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *MemberDns) GetResolverQueryTimeoutOk() (*int64, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *MemberDns) SetResolverQueryTimeout(v int64)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *MemberDns) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetResponseRateLimiting

`func (o *MemberDns) GetResponseRateLimiting() MemberDnsResponseRateLimiting`

GetResponseRateLimiting returns the ResponseRateLimiting field if non-nil, zero value otherwise.

### GetResponseRateLimitingOk

`func (o *MemberDns) GetResponseRateLimitingOk() (*MemberDnsResponseRateLimiting, bool)`

GetResponseRateLimitingOk returns a tuple with the ResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRateLimiting

`func (o *MemberDns) SetResponseRateLimiting(v MemberDnsResponseRateLimiting)`

SetResponseRateLimiting sets ResponseRateLimiting field to given value.

### HasResponseRateLimiting

`func (o *MemberDns) HasResponseRateLimiting() bool`

HasResponseRateLimiting returns a boolean if a field has been set.

### GetRootNameServerType

`func (o *MemberDns) GetRootNameServerType() string`

GetRootNameServerType returns the RootNameServerType field if non-nil, zero value otherwise.

### GetRootNameServerTypeOk

`func (o *MemberDns) GetRootNameServerTypeOk() (*string, bool)`

GetRootNameServerTypeOk returns a tuple with the RootNameServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNameServerType

`func (o *MemberDns) SetRootNameServerType(v string)`

SetRootNameServerType sets RootNameServerType field to given value.

### HasRootNameServerType

`func (o *MemberDns) HasRootNameServerType() bool`

HasRootNameServerType returns a boolean if a field has been set.

### GetRpzDisableNsdnameNsip

`func (o *MemberDns) GetRpzDisableNsdnameNsip() bool`

GetRpzDisableNsdnameNsip returns the RpzDisableNsdnameNsip field if non-nil, zero value otherwise.

### GetRpzDisableNsdnameNsipOk

`func (o *MemberDns) GetRpzDisableNsdnameNsipOk() (*bool, bool)`

GetRpzDisableNsdnameNsipOk returns a tuple with the RpzDisableNsdnameNsip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDisableNsdnameNsip

`func (o *MemberDns) SetRpzDisableNsdnameNsip(v bool)`

SetRpzDisableNsdnameNsip sets RpzDisableNsdnameNsip field to given value.

### HasRpzDisableNsdnameNsip

`func (o *MemberDns) HasRpzDisableNsdnameNsip() bool`

HasRpzDisableNsdnameNsip returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *MemberDns) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *MemberDns) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *MemberDns) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *MemberDns) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *MemberDns) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *MemberDns) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *MemberDns) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *MemberDns) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *MemberDns) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *MemberDns) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *MemberDns) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *MemberDns) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzQnameWaitRecurse

`func (o *MemberDns) GetRpzQnameWaitRecurse() bool`

GetRpzQnameWaitRecurse returns the RpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetRpzQnameWaitRecurseOk

`func (o *MemberDns) GetRpzQnameWaitRecurseOk() (*bool, bool)`

GetRpzQnameWaitRecurseOk returns a tuple with the RpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzQnameWaitRecurse

`func (o *MemberDns) SetRpzQnameWaitRecurse(v bool)`

SetRpzQnameWaitRecurse sets RpzQnameWaitRecurse field to given value.

### HasRpzQnameWaitRecurse

`func (o *MemberDns) HasRpzQnameWaitRecurse() bool`

HasRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetSerialQueryRate

`func (o *MemberDns) GetSerialQueryRate() int64`

GetSerialQueryRate returns the SerialQueryRate field if non-nil, zero value otherwise.

### GetSerialQueryRateOk

`func (o *MemberDns) GetSerialQueryRateOk() (*int64, bool)`

GetSerialQueryRateOk returns a tuple with the SerialQueryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialQueryRate

`func (o *MemberDns) SetSerialQueryRate(v int64)`

SetSerialQueryRate sets SerialQueryRate field to given value.

### HasSerialQueryRate

`func (o *MemberDns) HasSerialQueryRate() bool`

HasSerialQueryRate returns a boolean if a field has been set.

### GetServerIdDirective

`func (o *MemberDns) GetServerIdDirective() string`

GetServerIdDirective returns the ServerIdDirective field if non-nil, zero value otherwise.

### GetServerIdDirectiveOk

`func (o *MemberDns) GetServerIdDirectiveOk() (*string, bool)`

GetServerIdDirectiveOk returns a tuple with the ServerIdDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIdDirective

`func (o *MemberDns) SetServerIdDirective(v string)`

SetServerIdDirective sets ServerIdDirective field to given value.

### HasServerIdDirective

`func (o *MemberDns) HasServerIdDirective() bool`

HasServerIdDirective returns a boolean if a field has been set.

### GetServerIdDirectiveString

`func (o *MemberDns) GetServerIdDirectiveString() string`

GetServerIdDirectiveString returns the ServerIdDirectiveString field if non-nil, zero value otherwise.

### GetServerIdDirectiveStringOk

`func (o *MemberDns) GetServerIdDirectiveStringOk() (*string, bool)`

GetServerIdDirectiveStringOk returns a tuple with the ServerIdDirectiveString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIdDirectiveString

`func (o *MemberDns) SetServerIdDirectiveString(v string)`

SetServerIdDirectiveString sets ServerIdDirectiveString field to given value.

### HasServerIdDirectiveString

`func (o *MemberDns) HasServerIdDirectiveString() bool`

HasServerIdDirectiveString returns a boolean if a field has been set.

### GetSkipInGridRpzQueries

`func (o *MemberDns) GetSkipInGridRpzQueries() bool`

GetSkipInGridRpzQueries returns the SkipInGridRpzQueries field if non-nil, zero value otherwise.

### GetSkipInGridRpzQueriesOk

`func (o *MemberDns) GetSkipInGridRpzQueriesOk() (*bool, bool)`

GetSkipInGridRpzQueriesOk returns a tuple with the SkipInGridRpzQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipInGridRpzQueries

`func (o *MemberDns) SetSkipInGridRpzQueries(v bool)`

SetSkipInGridRpzQueries sets SkipInGridRpzQueries field to given value.

### HasSkipInGridRpzQueries

`func (o *MemberDns) HasSkipInGridRpzQueries() bool`

HasSkipInGridRpzQueries returns a boolean if a field has been set.

### GetSortlist

`func (o *MemberDns) GetSortlist() []MemberDnsSortlist`

GetSortlist returns the Sortlist field if non-nil, zero value otherwise.

### GetSortlistOk

`func (o *MemberDns) GetSortlistOk() (*[]MemberDnsSortlist, bool)`

GetSortlistOk returns a tuple with the Sortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortlist

`func (o *MemberDns) SetSortlist(v []MemberDnsSortlist)`

SetSortlist sets Sortlist field to given value.

### HasSortlist

`func (o *MemberDns) HasSortlist() bool`

HasSortlist returns a boolean if a field has been set.

### GetStoreLocally

`func (o *MemberDns) GetStoreLocally() bool`

GetStoreLocally returns the StoreLocally field if non-nil, zero value otherwise.

### GetStoreLocallyOk

`func (o *MemberDns) GetStoreLocallyOk() (*bool, bool)`

GetStoreLocallyOk returns a tuple with the StoreLocally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreLocally

`func (o *MemberDns) SetStoreLocally(v bool)`

SetStoreLocally sets StoreLocally field to given value.

### HasStoreLocally

`func (o *MemberDns) HasStoreLocally() bool`

HasStoreLocally returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *MemberDns) GetSyslogFacility() string`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *MemberDns) GetSyslogFacilityOk() (*string, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *MemberDns) SetSyslogFacility(v string)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *MemberDns) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetTcpIdleTimeout

`func (o *MemberDns) GetTcpIdleTimeout() int64`

GetTcpIdleTimeout returns the TcpIdleTimeout field if non-nil, zero value otherwise.

### GetTcpIdleTimeoutOk

`func (o *MemberDns) GetTcpIdleTimeoutOk() (*int64, bool)`

GetTcpIdleTimeoutOk returns a tuple with the TcpIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIdleTimeout

`func (o *MemberDns) SetTcpIdleTimeout(v int64)`

SetTcpIdleTimeout sets TcpIdleTimeout field to given value.

### HasTcpIdleTimeout

`func (o *MemberDns) HasTcpIdleTimeout() bool`

HasTcpIdleTimeout returns a boolean if a field has been set.

### GetTlsSessionDuration

`func (o *MemberDns) GetTlsSessionDuration() int64`

GetTlsSessionDuration returns the TlsSessionDuration field if non-nil, zero value otherwise.

### GetTlsSessionDurationOk

`func (o *MemberDns) GetTlsSessionDurationOk() (*int64, bool)`

GetTlsSessionDurationOk returns a tuple with the TlsSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSessionDuration

`func (o *MemberDns) SetTlsSessionDuration(v int64)`

SetTlsSessionDuration sets TlsSessionDuration field to given value.

### HasTlsSessionDuration

`func (o *MemberDns) HasTlsSessionDuration() bool`

HasTlsSessionDuration returns a boolean if a field has been set.

### GetTransferExcludedServers

`func (o *MemberDns) GetTransferExcludedServers() []string`

GetTransferExcludedServers returns the TransferExcludedServers field if non-nil, zero value otherwise.

### GetTransferExcludedServersOk

`func (o *MemberDns) GetTransferExcludedServersOk() (*[]string, bool)`

GetTransferExcludedServersOk returns a tuple with the TransferExcludedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferExcludedServers

`func (o *MemberDns) SetTransferExcludedServers(v []string)`

SetTransferExcludedServers sets TransferExcludedServers field to given value.

### HasTransferExcludedServers

`func (o *MemberDns) HasTransferExcludedServers() bool`

HasTransferExcludedServers returns a boolean if a field has been set.

### GetTransferFormat

`func (o *MemberDns) GetTransferFormat() string`

GetTransferFormat returns the TransferFormat field if non-nil, zero value otherwise.

### GetTransferFormatOk

`func (o *MemberDns) GetTransferFormatOk() (*string, bool)`

GetTransferFormatOk returns a tuple with the TransferFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFormat

`func (o *MemberDns) SetTransferFormat(v string)`

SetTransferFormat sets TransferFormat field to given value.

### HasTransferFormat

`func (o *MemberDns) HasTransferFormat() bool`

HasTransferFormat returns a boolean if a field has been set.

### GetTransfersIn

`func (o *MemberDns) GetTransfersIn() int64`

GetTransfersIn returns the TransfersIn field if non-nil, zero value otherwise.

### GetTransfersInOk

`func (o *MemberDns) GetTransfersInOk() (*int64, bool)`

GetTransfersInOk returns a tuple with the TransfersIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersIn

`func (o *MemberDns) SetTransfersIn(v int64)`

SetTransfersIn sets TransfersIn field to given value.

### HasTransfersIn

`func (o *MemberDns) HasTransfersIn() bool`

HasTransfersIn returns a boolean if a field has been set.

### GetTransfersOut

`func (o *MemberDns) GetTransfersOut() int64`

GetTransfersOut returns the TransfersOut field if non-nil, zero value otherwise.

### GetTransfersOutOk

`func (o *MemberDns) GetTransfersOutOk() (*int64, bool)`

GetTransfersOutOk returns a tuple with the TransfersOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersOut

`func (o *MemberDns) SetTransfersOut(v int64)`

SetTransfersOut sets TransfersOut field to given value.

### HasTransfersOut

`func (o *MemberDns) HasTransfersOut() bool`

HasTransfersOut returns a boolean if a field has been set.

### GetTransfersPerNs

`func (o *MemberDns) GetTransfersPerNs() int64`

GetTransfersPerNs returns the TransfersPerNs field if non-nil, zero value otherwise.

### GetTransfersPerNsOk

`func (o *MemberDns) GetTransfersPerNsOk() (*int64, bool)`

GetTransfersPerNsOk returns a tuple with the TransfersPerNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersPerNs

`func (o *MemberDns) SetTransfersPerNs(v int64)`

SetTransfersPerNs sets TransfersPerNs field to given value.

### HasTransfersPerNs

`func (o *MemberDns) HasTransfersPerNs() bool`

HasTransfersPerNs returns a boolean if a field has been set.

### GetUpstreamAddressFamilyPreference

`func (o *MemberDns) GetUpstreamAddressFamilyPreference() string`

GetUpstreamAddressFamilyPreference returns the UpstreamAddressFamilyPreference field if non-nil, zero value otherwise.

### GetUpstreamAddressFamilyPreferenceOk

`func (o *MemberDns) GetUpstreamAddressFamilyPreferenceOk() (*string, bool)`

GetUpstreamAddressFamilyPreferenceOk returns a tuple with the UpstreamAddressFamilyPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamAddressFamilyPreference

`func (o *MemberDns) SetUpstreamAddressFamilyPreference(v string)`

SetUpstreamAddressFamilyPreference sets UpstreamAddressFamilyPreference field to given value.

### HasUpstreamAddressFamilyPreference

`func (o *MemberDns) HasUpstreamAddressFamilyPreference() bool`

HasUpstreamAddressFamilyPreference returns a boolean if a field has been set.

### GetUseAddClientIpMacOptions

`func (o *MemberDns) GetUseAddClientIpMacOptions() bool`

GetUseAddClientIpMacOptions returns the UseAddClientIpMacOptions field if non-nil, zero value otherwise.

### GetUseAddClientIpMacOptionsOk

`func (o *MemberDns) GetUseAddClientIpMacOptionsOk() (*bool, bool)`

GetUseAddClientIpMacOptionsOk returns a tuple with the UseAddClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAddClientIpMacOptions

`func (o *MemberDns) SetUseAddClientIpMacOptions(v bool)`

SetUseAddClientIpMacOptions sets UseAddClientIpMacOptions field to given value.

### HasUseAddClientIpMacOptions

`func (o *MemberDns) HasUseAddClientIpMacOptions() bool`

HasUseAddClientIpMacOptions returns a boolean if a field has been set.

### GetUseAllowQuery

`func (o *MemberDns) GetUseAllowQuery() bool`

GetUseAllowQuery returns the UseAllowQuery field if non-nil, zero value otherwise.

### GetUseAllowQueryOk

`func (o *MemberDns) GetUseAllowQueryOk() (*bool, bool)`

GetUseAllowQueryOk returns a tuple with the UseAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowQuery

`func (o *MemberDns) SetUseAllowQuery(v bool)`

SetUseAllowQuery sets UseAllowQuery field to given value.

### HasUseAllowQuery

`func (o *MemberDns) HasUseAllowQuery() bool`

HasUseAllowQuery returns a boolean if a field has been set.

### GetUseAllowTransfer

`func (o *MemberDns) GetUseAllowTransfer() bool`

GetUseAllowTransfer returns the UseAllowTransfer field if non-nil, zero value otherwise.

### GetUseAllowTransferOk

`func (o *MemberDns) GetUseAllowTransferOk() (*bool, bool)`

GetUseAllowTransferOk returns a tuple with the UseAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowTransfer

`func (o *MemberDns) SetUseAllowTransfer(v bool)`

SetUseAllowTransfer sets UseAllowTransfer field to given value.

### HasUseAllowTransfer

`func (o *MemberDns) HasUseAllowTransfer() bool`

HasUseAllowTransfer returns a boolean if a field has been set.

### GetUseAttackMitigation

`func (o *MemberDns) GetUseAttackMitigation() bool`

GetUseAttackMitigation returns the UseAttackMitigation field if non-nil, zero value otherwise.

### GetUseAttackMitigationOk

`func (o *MemberDns) GetUseAttackMitigationOk() (*bool, bool)`

GetUseAttackMitigationOk returns a tuple with the UseAttackMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAttackMitigation

`func (o *MemberDns) SetUseAttackMitigation(v bool)`

SetUseAttackMitigation sets UseAttackMitigation field to given value.

### HasUseAttackMitigation

`func (o *MemberDns) HasUseAttackMitigation() bool`

HasUseAttackMitigation returns a boolean if a field has been set.

### GetUseAutoBlackhole

`func (o *MemberDns) GetUseAutoBlackhole() bool`

GetUseAutoBlackhole returns the UseAutoBlackhole field if non-nil, zero value otherwise.

### GetUseAutoBlackholeOk

`func (o *MemberDns) GetUseAutoBlackholeOk() (*bool, bool)`

GetUseAutoBlackholeOk returns a tuple with the UseAutoBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAutoBlackhole

`func (o *MemberDns) SetUseAutoBlackhole(v bool)`

SetUseAutoBlackhole sets UseAutoBlackhole field to given value.

### HasUseAutoBlackhole

`func (o *MemberDns) HasUseAutoBlackhole() bool`

HasUseAutoBlackhole returns a boolean if a field has been set.

### GetUseBindHostnameDirective

`func (o *MemberDns) GetUseBindHostnameDirective() bool`

GetUseBindHostnameDirective returns the UseBindHostnameDirective field if non-nil, zero value otherwise.

### GetUseBindHostnameDirectiveOk

`func (o *MemberDns) GetUseBindHostnameDirectiveOk() (*bool, bool)`

GetUseBindHostnameDirectiveOk returns a tuple with the UseBindHostnameDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBindHostnameDirective

`func (o *MemberDns) SetUseBindHostnameDirective(v bool)`

SetUseBindHostnameDirective sets UseBindHostnameDirective field to given value.

### HasUseBindHostnameDirective

`func (o *MemberDns) HasUseBindHostnameDirective() bool`

HasUseBindHostnameDirective returns a boolean if a field has been set.

### GetUseBlackhole

`func (o *MemberDns) GetUseBlackhole() bool`

GetUseBlackhole returns the UseBlackhole field if non-nil, zero value otherwise.

### GetUseBlackholeOk

`func (o *MemberDns) GetUseBlackholeOk() (*bool, bool)`

GetUseBlackholeOk returns a tuple with the UseBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackhole

`func (o *MemberDns) SetUseBlackhole(v bool)`

SetUseBlackhole sets UseBlackhole field to given value.

### HasUseBlackhole

`func (o *MemberDns) HasUseBlackhole() bool`

HasUseBlackhole returns a boolean if a field has been set.

### GetUseBlacklist

`func (o *MemberDns) GetUseBlacklist() bool`

GetUseBlacklist returns the UseBlacklist field if non-nil, zero value otherwise.

### GetUseBlacklistOk

`func (o *MemberDns) GetUseBlacklistOk() (*bool, bool)`

GetUseBlacklistOk returns a tuple with the UseBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlacklist

`func (o *MemberDns) SetUseBlacklist(v bool)`

SetUseBlacklist sets UseBlacklist field to given value.

### HasUseBlacklist

`func (o *MemberDns) HasUseBlacklist() bool`

HasUseBlacklist returns a boolean if a field has been set.

### GetUseCaptureDnsQueriesOnAllDomains

`func (o *MemberDns) GetUseCaptureDnsQueriesOnAllDomains() bool`

GetUseCaptureDnsQueriesOnAllDomains returns the UseCaptureDnsQueriesOnAllDomains field if non-nil, zero value otherwise.

### GetUseCaptureDnsQueriesOnAllDomainsOk

`func (o *MemberDns) GetUseCaptureDnsQueriesOnAllDomainsOk() (*bool, bool)`

GetUseCaptureDnsQueriesOnAllDomainsOk returns a tuple with the UseCaptureDnsQueriesOnAllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCaptureDnsQueriesOnAllDomains

`func (o *MemberDns) SetUseCaptureDnsQueriesOnAllDomains(v bool)`

SetUseCaptureDnsQueriesOnAllDomains sets UseCaptureDnsQueriesOnAllDomains field to given value.

### HasUseCaptureDnsQueriesOnAllDomains

`func (o *MemberDns) HasUseCaptureDnsQueriesOnAllDomains() bool`

HasUseCaptureDnsQueriesOnAllDomains returns a boolean if a field has been set.

### GetUseCopyClientIpMacOptions

`func (o *MemberDns) GetUseCopyClientIpMacOptions() bool`

GetUseCopyClientIpMacOptions returns the UseCopyClientIpMacOptions field if non-nil, zero value otherwise.

### GetUseCopyClientIpMacOptionsOk

`func (o *MemberDns) GetUseCopyClientIpMacOptionsOk() (*bool, bool)`

GetUseCopyClientIpMacOptionsOk returns a tuple with the UseCopyClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCopyClientIpMacOptions

`func (o *MemberDns) SetUseCopyClientIpMacOptions(v bool)`

SetUseCopyClientIpMacOptions sets UseCopyClientIpMacOptions field to given value.

### HasUseCopyClientIpMacOptions

`func (o *MemberDns) HasUseCopyClientIpMacOptions() bool`

HasUseCopyClientIpMacOptions returns a boolean if a field has been set.

### GetUseCopyXferToNotify

`func (o *MemberDns) GetUseCopyXferToNotify() bool`

GetUseCopyXferToNotify returns the UseCopyXferToNotify field if non-nil, zero value otherwise.

### GetUseCopyXferToNotifyOk

`func (o *MemberDns) GetUseCopyXferToNotifyOk() (*bool, bool)`

GetUseCopyXferToNotifyOk returns a tuple with the UseCopyXferToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCopyXferToNotify

`func (o *MemberDns) SetUseCopyXferToNotify(v bool)`

SetUseCopyXferToNotify sets UseCopyXferToNotify field to given value.

### HasUseCopyXferToNotify

`func (o *MemberDns) HasUseCopyXferToNotify() bool`

HasUseCopyXferToNotify returns a boolean if a field has been set.

### GetUseDisableEdns

`func (o *MemberDns) GetUseDisableEdns() bool`

GetUseDisableEdns returns the UseDisableEdns field if non-nil, zero value otherwise.

### GetUseDisableEdnsOk

`func (o *MemberDns) GetUseDisableEdnsOk() (*bool, bool)`

GetUseDisableEdnsOk returns a tuple with the UseDisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisableEdns

`func (o *MemberDns) SetUseDisableEdns(v bool)`

SetUseDisableEdns sets UseDisableEdns field to given value.

### HasUseDisableEdns

`func (o *MemberDns) HasUseDisableEdns() bool`

HasUseDisableEdns returns a boolean if a field has been set.

### GetUseDns64

`func (o *MemberDns) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *MemberDns) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *MemberDns) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.

### HasUseDns64

`func (o *MemberDns) HasUseDns64() bool`

HasUseDns64 returns a boolean if a field has been set.

### GetUseDnsCacheAccelerationTtl

`func (o *MemberDns) GetUseDnsCacheAccelerationTtl() bool`

GetUseDnsCacheAccelerationTtl returns the UseDnsCacheAccelerationTtl field if non-nil, zero value otherwise.

### GetUseDnsCacheAccelerationTtlOk

`func (o *MemberDns) GetUseDnsCacheAccelerationTtlOk() (*bool, bool)`

GetUseDnsCacheAccelerationTtlOk returns a tuple with the UseDnsCacheAccelerationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsCacheAccelerationTtl

`func (o *MemberDns) SetUseDnsCacheAccelerationTtl(v bool)`

SetUseDnsCacheAccelerationTtl sets UseDnsCacheAccelerationTtl field to given value.

### HasUseDnsCacheAccelerationTtl

`func (o *MemberDns) HasUseDnsCacheAccelerationTtl() bool`

HasUseDnsCacheAccelerationTtl returns a boolean if a field has been set.

### GetUseDnsHealthCheck

`func (o *MemberDns) GetUseDnsHealthCheck() bool`

GetUseDnsHealthCheck returns the UseDnsHealthCheck field if non-nil, zero value otherwise.

### GetUseDnsHealthCheckOk

`func (o *MemberDns) GetUseDnsHealthCheckOk() (*bool, bool)`

GetUseDnsHealthCheckOk returns a tuple with the UseDnsHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsHealthCheck

`func (o *MemberDns) SetUseDnsHealthCheck(v bool)`

SetUseDnsHealthCheck sets UseDnsHealthCheck field to given value.

### HasUseDnsHealthCheck

`func (o *MemberDns) HasUseDnsHealthCheck() bool`

HasUseDnsHealthCheck returns a boolean if a field has been set.

### GetUseDnssec

`func (o *MemberDns) GetUseDnssec() bool`

GetUseDnssec returns the UseDnssec field if non-nil, zero value otherwise.

### GetUseDnssecOk

`func (o *MemberDns) GetUseDnssecOk() (*bool, bool)`

GetUseDnssecOk returns a tuple with the UseDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnssec

`func (o *MemberDns) SetUseDnssec(v bool)`

SetUseDnssec sets UseDnssec field to given value.

### HasUseDnssec

`func (o *MemberDns) HasUseDnssec() bool`

HasUseDnssec returns a boolean if a field has been set.

### GetUseDnstapSetting

`func (o *MemberDns) GetUseDnstapSetting() bool`

GetUseDnstapSetting returns the UseDnstapSetting field if non-nil, zero value otherwise.

### GetUseDnstapSettingOk

`func (o *MemberDns) GetUseDnstapSettingOk() (*bool, bool)`

GetUseDnstapSettingOk returns a tuple with the UseDnstapSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnstapSetting

`func (o *MemberDns) SetUseDnstapSetting(v bool)`

SetUseDnstapSetting sets UseDnstapSetting field to given value.

### HasUseDnstapSetting

`func (o *MemberDns) HasUseDnstapSetting() bool`

HasUseDnstapSetting returns a boolean if a field has been set.

### GetUseDtcDnsQueriesSpecificBehavior

`func (o *MemberDns) GetUseDtcDnsQueriesSpecificBehavior() bool`

GetUseDtcDnsQueriesSpecificBehavior returns the UseDtcDnsQueriesSpecificBehavior field if non-nil, zero value otherwise.

### GetUseDtcDnsQueriesSpecificBehaviorOk

`func (o *MemberDns) GetUseDtcDnsQueriesSpecificBehaviorOk() (*bool, bool)`

GetUseDtcDnsQueriesSpecificBehaviorOk returns a tuple with the UseDtcDnsQueriesSpecificBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDtcDnsQueriesSpecificBehavior

`func (o *MemberDns) SetUseDtcDnsQueriesSpecificBehavior(v bool)`

SetUseDtcDnsQueriesSpecificBehavior sets UseDtcDnsQueriesSpecificBehavior field to given value.

### HasUseDtcDnsQueriesSpecificBehavior

`func (o *MemberDns) HasUseDtcDnsQueriesSpecificBehavior() bool`

HasUseDtcDnsQueriesSpecificBehavior returns a boolean if a field has been set.

### GetUseDtcEdnsPreferClientSubnet

`func (o *MemberDns) GetUseDtcEdnsPreferClientSubnet() bool`

GetUseDtcEdnsPreferClientSubnet returns the UseDtcEdnsPreferClientSubnet field if non-nil, zero value otherwise.

### GetUseDtcEdnsPreferClientSubnetOk

`func (o *MemberDns) GetUseDtcEdnsPreferClientSubnetOk() (*bool, bool)`

GetUseDtcEdnsPreferClientSubnetOk returns a tuple with the UseDtcEdnsPreferClientSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDtcEdnsPreferClientSubnet

`func (o *MemberDns) SetUseDtcEdnsPreferClientSubnet(v bool)`

SetUseDtcEdnsPreferClientSubnet sets UseDtcEdnsPreferClientSubnet field to given value.

### HasUseDtcEdnsPreferClientSubnet

`func (o *MemberDns) HasUseDtcEdnsPreferClientSubnet() bool`

HasUseDtcEdnsPreferClientSubnet returns a boolean if a field has been set.

### GetUseEdnsUdpSize

`func (o *MemberDns) GetUseEdnsUdpSize() bool`

GetUseEdnsUdpSize returns the UseEdnsUdpSize field if non-nil, zero value otherwise.

### GetUseEdnsUdpSizeOk

`func (o *MemberDns) GetUseEdnsUdpSizeOk() (*bool, bool)`

GetUseEdnsUdpSizeOk returns a tuple with the UseEdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEdnsUdpSize

`func (o *MemberDns) SetUseEdnsUdpSize(v bool)`

SetUseEdnsUdpSize sets UseEdnsUdpSize field to given value.

### HasUseEdnsUdpSize

`func (o *MemberDns) HasUseEdnsUdpSize() bool`

HasUseEdnsUdpSize returns a boolean if a field has been set.

### GetUseEnableCaptureDns

`func (o *MemberDns) GetUseEnableCaptureDns() bool`

GetUseEnableCaptureDns returns the UseEnableCaptureDns field if non-nil, zero value otherwise.

### GetUseEnableCaptureDnsOk

`func (o *MemberDns) GetUseEnableCaptureDnsOk() (*bool, bool)`

GetUseEnableCaptureDnsOk returns a tuple with the UseEnableCaptureDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableCaptureDns

`func (o *MemberDns) SetUseEnableCaptureDns(v bool)`

SetUseEnableCaptureDns sets UseEnableCaptureDns field to given value.

### HasUseEnableCaptureDns

`func (o *MemberDns) HasUseEnableCaptureDns() bool`

HasUseEnableCaptureDns returns a boolean if a field has been set.

### GetUseEnableExcludedDomainNames

`func (o *MemberDns) GetUseEnableExcludedDomainNames() bool`

GetUseEnableExcludedDomainNames returns the UseEnableExcludedDomainNames field if non-nil, zero value otherwise.

### GetUseEnableExcludedDomainNamesOk

`func (o *MemberDns) GetUseEnableExcludedDomainNamesOk() (*bool, bool)`

GetUseEnableExcludedDomainNamesOk returns a tuple with the UseEnableExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableExcludedDomainNames

`func (o *MemberDns) SetUseEnableExcludedDomainNames(v bool)`

SetUseEnableExcludedDomainNames sets UseEnableExcludedDomainNames field to given value.

### HasUseEnableExcludedDomainNames

`func (o *MemberDns) HasUseEnableExcludedDomainNames() bool`

HasUseEnableExcludedDomainNames returns a boolean if a field has been set.

### GetUseEnableGssTsig

`func (o *MemberDns) GetUseEnableGssTsig() bool`

GetUseEnableGssTsig returns the UseEnableGssTsig field if non-nil, zero value otherwise.

### GetUseEnableGssTsigOk

`func (o *MemberDns) GetUseEnableGssTsigOk() (*bool, bool)`

GetUseEnableGssTsigOk returns a tuple with the UseEnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableGssTsig

`func (o *MemberDns) SetUseEnableGssTsig(v bool)`

SetUseEnableGssTsig sets UseEnableGssTsig field to given value.

### HasUseEnableGssTsig

`func (o *MemberDns) HasUseEnableGssTsig() bool`

HasUseEnableGssTsig returns a boolean if a field has been set.

### GetUseEnableQueryRewrite

`func (o *MemberDns) GetUseEnableQueryRewrite() bool`

GetUseEnableQueryRewrite returns the UseEnableQueryRewrite field if non-nil, zero value otherwise.

### GetUseEnableQueryRewriteOk

`func (o *MemberDns) GetUseEnableQueryRewriteOk() (*bool, bool)`

GetUseEnableQueryRewriteOk returns a tuple with the UseEnableQueryRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableQueryRewrite

`func (o *MemberDns) SetUseEnableQueryRewrite(v bool)`

SetUseEnableQueryRewrite sets UseEnableQueryRewrite field to given value.

### HasUseEnableQueryRewrite

`func (o *MemberDns) HasUseEnableQueryRewrite() bool`

HasUseEnableQueryRewrite returns a boolean if a field has been set.

### GetUseFilterAaaa

`func (o *MemberDns) GetUseFilterAaaa() bool`

GetUseFilterAaaa returns the UseFilterAaaa field if non-nil, zero value otherwise.

### GetUseFilterAaaaOk

`func (o *MemberDns) GetUseFilterAaaaOk() (*bool, bool)`

GetUseFilterAaaaOk returns a tuple with the UseFilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFilterAaaa

`func (o *MemberDns) SetUseFilterAaaa(v bool)`

SetUseFilterAaaa sets UseFilterAaaa field to given value.

### HasUseFilterAaaa

`func (o *MemberDns) HasUseFilterAaaa() bool`

HasUseFilterAaaa returns a boolean if a field has been set.

### GetUseFixedRrsetOrderFqdns

`func (o *MemberDns) GetUseFixedRrsetOrderFqdns() bool`

GetUseFixedRrsetOrderFqdns returns the UseFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetUseFixedRrsetOrderFqdnsOk

`func (o *MemberDns) GetUseFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetUseFixedRrsetOrderFqdnsOk returns a tuple with the UseFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedRrsetOrderFqdns

`func (o *MemberDns) SetUseFixedRrsetOrderFqdns(v bool)`

SetUseFixedRrsetOrderFqdns sets UseFixedRrsetOrderFqdns field to given value.

### HasUseFixedRrsetOrderFqdns

`func (o *MemberDns) HasUseFixedRrsetOrderFqdns() bool`

HasUseFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetUseForwardUpdates

`func (o *MemberDns) GetUseForwardUpdates() bool`

GetUseForwardUpdates returns the UseForwardUpdates field if non-nil, zero value otherwise.

### GetUseForwardUpdatesOk

`func (o *MemberDns) GetUseForwardUpdatesOk() (*bool, bool)`

GetUseForwardUpdatesOk returns a tuple with the UseForwardUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardUpdates

`func (o *MemberDns) SetUseForwardUpdates(v bool)`

SetUseForwardUpdates sets UseForwardUpdates field to given value.

### HasUseForwardUpdates

`func (o *MemberDns) HasUseForwardUpdates() bool`

HasUseForwardUpdates returns a boolean if a field has been set.

### GetUseForwarders

`func (o *MemberDns) GetUseForwarders() bool`

GetUseForwarders returns the UseForwarders field if non-nil, zero value otherwise.

### GetUseForwardersOk

`func (o *MemberDns) GetUseForwardersOk() (*bool, bool)`

GetUseForwardersOk returns a tuple with the UseForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwarders

`func (o *MemberDns) SetUseForwarders(v bool)`

SetUseForwarders sets UseForwarders field to given value.

### HasUseForwarders

`func (o *MemberDns) HasUseForwarders() bool`

HasUseForwarders returns a boolean if a field has been set.

### GetUseFtc

`func (o *MemberDns) GetUseFtc() bool`

GetUseFtc returns the UseFtc field if non-nil, zero value otherwise.

### GetUseFtcOk

`func (o *MemberDns) GetUseFtcOk() (*bool, bool)`

GetUseFtcOk returns a tuple with the UseFtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFtc

`func (o *MemberDns) SetUseFtc(v bool)`

SetUseFtc sets UseFtc field to given value.

### HasUseFtc

`func (o *MemberDns) HasUseFtc() bool`

HasUseFtc returns a boolean if a field has been set.

### GetUseGssTsigKeys

`func (o *MemberDns) GetUseGssTsigKeys() bool`

GetUseGssTsigKeys returns the UseGssTsigKeys field if non-nil, zero value otherwise.

### GetUseGssTsigKeysOk

`func (o *MemberDns) GetUseGssTsigKeysOk() (*bool, bool)`

GetUseGssTsigKeysOk returns a tuple with the UseGssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGssTsigKeys

`func (o *MemberDns) SetUseGssTsigKeys(v bool)`

SetUseGssTsigKeys sets UseGssTsigKeys field to given value.

### HasUseGssTsigKeys

`func (o *MemberDns) HasUseGssTsigKeys() bool`

HasUseGssTsigKeys returns a boolean if a field has been set.

### GetUseLan2Ipv6Port

`func (o *MemberDns) GetUseLan2Ipv6Port() bool`

GetUseLan2Ipv6Port returns the UseLan2Ipv6Port field if non-nil, zero value otherwise.

### GetUseLan2Ipv6PortOk

`func (o *MemberDns) GetUseLan2Ipv6PortOk() (*bool, bool)`

GetUseLan2Ipv6PortOk returns a tuple with the UseLan2Ipv6Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLan2Ipv6Port

`func (o *MemberDns) SetUseLan2Ipv6Port(v bool)`

SetUseLan2Ipv6Port sets UseLan2Ipv6Port field to given value.

### HasUseLan2Ipv6Port

`func (o *MemberDns) HasUseLan2Ipv6Port() bool`

HasUseLan2Ipv6Port returns a boolean if a field has been set.

### GetUseLan2Port

`func (o *MemberDns) GetUseLan2Port() bool`

GetUseLan2Port returns the UseLan2Port field if non-nil, zero value otherwise.

### GetUseLan2PortOk

`func (o *MemberDns) GetUseLan2PortOk() (*bool, bool)`

GetUseLan2PortOk returns a tuple with the UseLan2Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLan2Port

`func (o *MemberDns) SetUseLan2Port(v bool)`

SetUseLan2Port sets UseLan2Port field to given value.

### HasUseLan2Port

`func (o *MemberDns) HasUseLan2Port() bool`

HasUseLan2Port returns a boolean if a field has been set.

### GetUseLanIpv6Port

`func (o *MemberDns) GetUseLanIpv6Port() bool`

GetUseLanIpv6Port returns the UseLanIpv6Port field if non-nil, zero value otherwise.

### GetUseLanIpv6PortOk

`func (o *MemberDns) GetUseLanIpv6PortOk() (*bool, bool)`

GetUseLanIpv6PortOk returns a tuple with the UseLanIpv6Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLanIpv6Port

`func (o *MemberDns) SetUseLanIpv6Port(v bool)`

SetUseLanIpv6Port sets UseLanIpv6Port field to given value.

### HasUseLanIpv6Port

`func (o *MemberDns) HasUseLanIpv6Port() bool`

HasUseLanIpv6Port returns a boolean if a field has been set.

### GetUseLanPort

`func (o *MemberDns) GetUseLanPort() bool`

GetUseLanPort returns the UseLanPort field if non-nil, zero value otherwise.

### GetUseLanPortOk

`func (o *MemberDns) GetUseLanPortOk() (*bool, bool)`

GetUseLanPortOk returns a tuple with the UseLanPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLanPort

`func (o *MemberDns) SetUseLanPort(v bool)`

SetUseLanPort sets UseLanPort field to given value.

### HasUseLanPort

`func (o *MemberDns) HasUseLanPort() bool`

HasUseLanPort returns a boolean if a field has been set.

### GetUseLoggingCategories

`func (o *MemberDns) GetUseLoggingCategories() bool`

GetUseLoggingCategories returns the UseLoggingCategories field if non-nil, zero value otherwise.

### GetUseLoggingCategoriesOk

`func (o *MemberDns) GetUseLoggingCategoriesOk() (*bool, bool)`

GetUseLoggingCategoriesOk returns a tuple with the UseLoggingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLoggingCategories

`func (o *MemberDns) SetUseLoggingCategories(v bool)`

SetUseLoggingCategories sets UseLoggingCategories field to given value.

### HasUseLoggingCategories

`func (o *MemberDns) HasUseLoggingCategories() bool`

HasUseLoggingCategories returns a boolean if a field has been set.

### GetUseMaxCacheTtl

`func (o *MemberDns) GetUseMaxCacheTtl() bool`

GetUseMaxCacheTtl returns the UseMaxCacheTtl field if non-nil, zero value otherwise.

### GetUseMaxCacheTtlOk

`func (o *MemberDns) GetUseMaxCacheTtlOk() (*bool, bool)`

GetUseMaxCacheTtlOk returns a tuple with the UseMaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxCacheTtl

`func (o *MemberDns) SetUseMaxCacheTtl(v bool)`

SetUseMaxCacheTtl sets UseMaxCacheTtl field to given value.

### HasUseMaxCacheTtl

`func (o *MemberDns) HasUseMaxCacheTtl() bool`

HasUseMaxCacheTtl returns a boolean if a field has been set.

### GetUseMaxCachedLifetime

`func (o *MemberDns) GetUseMaxCachedLifetime() bool`

GetUseMaxCachedLifetime returns the UseMaxCachedLifetime field if non-nil, zero value otherwise.

### GetUseMaxCachedLifetimeOk

`func (o *MemberDns) GetUseMaxCachedLifetimeOk() (*bool, bool)`

GetUseMaxCachedLifetimeOk returns a tuple with the UseMaxCachedLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxCachedLifetime

`func (o *MemberDns) SetUseMaxCachedLifetime(v bool)`

SetUseMaxCachedLifetime sets UseMaxCachedLifetime field to given value.

### HasUseMaxCachedLifetime

`func (o *MemberDns) HasUseMaxCachedLifetime() bool`

HasUseMaxCachedLifetime returns a boolean if a field has been set.

### GetUseMaxNcacheTtl

`func (o *MemberDns) GetUseMaxNcacheTtl() bool`

GetUseMaxNcacheTtl returns the UseMaxNcacheTtl field if non-nil, zero value otherwise.

### GetUseMaxNcacheTtlOk

`func (o *MemberDns) GetUseMaxNcacheTtlOk() (*bool, bool)`

GetUseMaxNcacheTtlOk returns a tuple with the UseMaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxNcacheTtl

`func (o *MemberDns) SetUseMaxNcacheTtl(v bool)`

SetUseMaxNcacheTtl sets UseMaxNcacheTtl field to given value.

### HasUseMaxNcacheTtl

`func (o *MemberDns) HasUseMaxNcacheTtl() bool`

HasUseMaxNcacheTtl returns a boolean if a field has been set.

### GetUseMaxUdpSize

`func (o *MemberDns) GetUseMaxUdpSize() bool`

GetUseMaxUdpSize returns the UseMaxUdpSize field if non-nil, zero value otherwise.

### GetUseMaxUdpSizeOk

`func (o *MemberDns) GetUseMaxUdpSizeOk() (*bool, bool)`

GetUseMaxUdpSizeOk returns a tuple with the UseMaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaxUdpSize

`func (o *MemberDns) SetUseMaxUdpSize(v bool)`

SetUseMaxUdpSize sets UseMaxUdpSize field to given value.

### HasUseMaxUdpSize

`func (o *MemberDns) HasUseMaxUdpSize() bool`

HasUseMaxUdpSize returns a boolean if a field has been set.

### GetUseMgmtIpv6Port

`func (o *MemberDns) GetUseMgmtIpv6Port() bool`

GetUseMgmtIpv6Port returns the UseMgmtIpv6Port field if non-nil, zero value otherwise.

### GetUseMgmtIpv6PortOk

`func (o *MemberDns) GetUseMgmtIpv6PortOk() (*bool, bool)`

GetUseMgmtIpv6PortOk returns a tuple with the UseMgmtIpv6Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtIpv6Port

`func (o *MemberDns) SetUseMgmtIpv6Port(v bool)`

SetUseMgmtIpv6Port sets UseMgmtIpv6Port field to given value.

### HasUseMgmtIpv6Port

`func (o *MemberDns) HasUseMgmtIpv6Port() bool`

HasUseMgmtIpv6Port returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *MemberDns) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *MemberDns) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *MemberDns) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *MemberDns) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.

### GetUseNotifyDelay

`func (o *MemberDns) GetUseNotifyDelay() bool`

GetUseNotifyDelay returns the UseNotifyDelay field if non-nil, zero value otherwise.

### GetUseNotifyDelayOk

`func (o *MemberDns) GetUseNotifyDelayOk() (*bool, bool)`

GetUseNotifyDelayOk returns a tuple with the UseNotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNotifyDelay

`func (o *MemberDns) SetUseNotifyDelay(v bool)`

SetUseNotifyDelay sets UseNotifyDelay field to given value.

### HasUseNotifyDelay

`func (o *MemberDns) HasUseNotifyDelay() bool`

HasUseNotifyDelay returns a boolean if a field has been set.

### GetUseNxdomainRedirect

`func (o *MemberDns) GetUseNxdomainRedirect() bool`

GetUseNxdomainRedirect returns the UseNxdomainRedirect field if non-nil, zero value otherwise.

### GetUseNxdomainRedirectOk

`func (o *MemberDns) GetUseNxdomainRedirectOk() (*bool, bool)`

GetUseNxdomainRedirectOk returns a tuple with the UseNxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNxdomainRedirect

`func (o *MemberDns) SetUseNxdomainRedirect(v bool)`

SetUseNxdomainRedirect sets UseNxdomainRedirect field to given value.

### HasUseNxdomainRedirect

`func (o *MemberDns) HasUseNxdomainRedirect() bool`

HasUseNxdomainRedirect returns a boolean if a field has been set.

### GetUseRecordNamePolicy

`func (o *MemberDns) GetUseRecordNamePolicy() bool`

GetUseRecordNamePolicy returns the UseRecordNamePolicy field if non-nil, zero value otherwise.

### GetUseRecordNamePolicyOk

`func (o *MemberDns) GetUseRecordNamePolicyOk() (*bool, bool)`

GetUseRecordNamePolicyOk returns a tuple with the UseRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecordNamePolicy

`func (o *MemberDns) SetUseRecordNamePolicy(v bool)`

SetUseRecordNamePolicy sets UseRecordNamePolicy field to given value.

### HasUseRecordNamePolicy

`func (o *MemberDns) HasUseRecordNamePolicy() bool`

HasUseRecordNamePolicy returns a boolean if a field has been set.

### GetUseRecursiveClientLimit

`func (o *MemberDns) GetUseRecursiveClientLimit() bool`

GetUseRecursiveClientLimit returns the UseRecursiveClientLimit field if non-nil, zero value otherwise.

### GetUseRecursiveClientLimitOk

`func (o *MemberDns) GetUseRecursiveClientLimitOk() (*bool, bool)`

GetUseRecursiveClientLimitOk returns a tuple with the UseRecursiveClientLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecursiveClientLimit

`func (o *MemberDns) SetUseRecursiveClientLimit(v bool)`

SetUseRecursiveClientLimit sets UseRecursiveClientLimit field to given value.

### HasUseRecursiveClientLimit

`func (o *MemberDns) HasUseRecursiveClientLimit() bool`

HasUseRecursiveClientLimit returns a boolean if a field has been set.

### GetUseRecursiveQuerySetting

`func (o *MemberDns) GetUseRecursiveQuerySetting() bool`

GetUseRecursiveQuerySetting returns the UseRecursiveQuerySetting field if non-nil, zero value otherwise.

### GetUseRecursiveQuerySettingOk

`func (o *MemberDns) GetUseRecursiveQuerySettingOk() (*bool, bool)`

GetUseRecursiveQuerySettingOk returns a tuple with the UseRecursiveQuerySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecursiveQuerySetting

`func (o *MemberDns) SetUseRecursiveQuerySetting(v bool)`

SetUseRecursiveQuerySetting sets UseRecursiveQuerySetting field to given value.

### HasUseRecursiveQuerySetting

`func (o *MemberDns) HasUseRecursiveQuerySetting() bool`

HasUseRecursiveQuerySetting returns a boolean if a field has been set.

### GetUseResolverQueryTimeout

`func (o *MemberDns) GetUseResolverQueryTimeout() bool`

GetUseResolverQueryTimeout returns the UseResolverQueryTimeout field if non-nil, zero value otherwise.

### GetUseResolverQueryTimeoutOk

`func (o *MemberDns) GetUseResolverQueryTimeoutOk() (*bool, bool)`

GetUseResolverQueryTimeoutOk returns a tuple with the UseResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResolverQueryTimeout

`func (o *MemberDns) SetUseResolverQueryTimeout(v bool)`

SetUseResolverQueryTimeout sets UseResolverQueryTimeout field to given value.

### HasUseResolverQueryTimeout

`func (o *MemberDns) HasUseResolverQueryTimeout() bool`

HasUseResolverQueryTimeout returns a boolean if a field has been set.

### GetUseResponseRateLimiting

`func (o *MemberDns) GetUseResponseRateLimiting() bool`

GetUseResponseRateLimiting returns the UseResponseRateLimiting field if non-nil, zero value otherwise.

### GetUseResponseRateLimitingOk

`func (o *MemberDns) GetUseResponseRateLimitingOk() (*bool, bool)`

GetUseResponseRateLimitingOk returns a tuple with the UseResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponseRateLimiting

`func (o *MemberDns) SetUseResponseRateLimiting(v bool)`

SetUseResponseRateLimiting sets UseResponseRateLimiting field to given value.

### HasUseResponseRateLimiting

`func (o *MemberDns) HasUseResponseRateLimiting() bool`

HasUseResponseRateLimiting returns a boolean if a field has been set.

### GetUseRootNameServer

`func (o *MemberDns) GetUseRootNameServer() bool`

GetUseRootNameServer returns the UseRootNameServer field if non-nil, zero value otherwise.

### GetUseRootNameServerOk

`func (o *MemberDns) GetUseRootNameServerOk() (*bool, bool)`

GetUseRootNameServerOk returns a tuple with the UseRootNameServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootNameServer

`func (o *MemberDns) SetUseRootNameServer(v bool)`

SetUseRootNameServer sets UseRootNameServer field to given value.

### HasUseRootNameServer

`func (o *MemberDns) HasUseRootNameServer() bool`

HasUseRootNameServer returns a boolean if a field has been set.

### GetUseRootServerForAllViews

`func (o *MemberDns) GetUseRootServerForAllViews() bool`

GetUseRootServerForAllViews returns the UseRootServerForAllViews field if non-nil, zero value otherwise.

### GetUseRootServerForAllViewsOk

`func (o *MemberDns) GetUseRootServerForAllViewsOk() (*bool, bool)`

GetUseRootServerForAllViewsOk returns a tuple with the UseRootServerForAllViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootServerForAllViews

`func (o *MemberDns) SetUseRootServerForAllViews(v bool)`

SetUseRootServerForAllViews sets UseRootServerForAllViews field to given value.

### HasUseRootServerForAllViews

`func (o *MemberDns) HasUseRootServerForAllViews() bool`

HasUseRootServerForAllViews returns a boolean if a field has been set.

### GetUseRpzDisableNsdnameNsip

`func (o *MemberDns) GetUseRpzDisableNsdnameNsip() bool`

GetUseRpzDisableNsdnameNsip returns the UseRpzDisableNsdnameNsip field if non-nil, zero value otherwise.

### GetUseRpzDisableNsdnameNsipOk

`func (o *MemberDns) GetUseRpzDisableNsdnameNsipOk() (*bool, bool)`

GetUseRpzDisableNsdnameNsipOk returns a tuple with the UseRpzDisableNsdnameNsip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzDisableNsdnameNsip

`func (o *MemberDns) SetUseRpzDisableNsdnameNsip(v bool)`

SetUseRpzDisableNsdnameNsip sets UseRpzDisableNsdnameNsip field to given value.

### HasUseRpzDisableNsdnameNsip

`func (o *MemberDns) HasUseRpzDisableNsdnameNsip() bool`

HasUseRpzDisableNsdnameNsip returns a boolean if a field has been set.

### GetUseRpzDropIpRule

`func (o *MemberDns) GetUseRpzDropIpRule() bool`

GetUseRpzDropIpRule returns the UseRpzDropIpRule field if non-nil, zero value otherwise.

### GetUseRpzDropIpRuleOk

`func (o *MemberDns) GetUseRpzDropIpRuleOk() (*bool, bool)`

GetUseRpzDropIpRuleOk returns a tuple with the UseRpzDropIpRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzDropIpRule

`func (o *MemberDns) SetUseRpzDropIpRule(v bool)`

SetUseRpzDropIpRule sets UseRpzDropIpRule field to given value.

### HasUseRpzDropIpRule

`func (o *MemberDns) HasUseRpzDropIpRule() bool`

HasUseRpzDropIpRule returns a boolean if a field has been set.

### GetUseRpzQnameWaitRecurse

`func (o *MemberDns) GetUseRpzQnameWaitRecurse() bool`

GetUseRpzQnameWaitRecurse returns the UseRpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetUseRpzQnameWaitRecurseOk

`func (o *MemberDns) GetUseRpzQnameWaitRecurseOk() (*bool, bool)`

GetUseRpzQnameWaitRecurseOk returns a tuple with the UseRpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzQnameWaitRecurse

`func (o *MemberDns) SetUseRpzQnameWaitRecurse(v bool)`

SetUseRpzQnameWaitRecurse sets UseRpzQnameWaitRecurse field to given value.

### HasUseRpzQnameWaitRecurse

`func (o *MemberDns) HasUseRpzQnameWaitRecurse() bool`

HasUseRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetUseSerialQueryRate

`func (o *MemberDns) GetUseSerialQueryRate() bool`

GetUseSerialQueryRate returns the UseSerialQueryRate field if non-nil, zero value otherwise.

### GetUseSerialQueryRateOk

`func (o *MemberDns) GetUseSerialQueryRateOk() (*bool, bool)`

GetUseSerialQueryRateOk returns a tuple with the UseSerialQueryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSerialQueryRate

`func (o *MemberDns) SetUseSerialQueryRate(v bool)`

SetUseSerialQueryRate sets UseSerialQueryRate field to given value.

### HasUseSerialQueryRate

`func (o *MemberDns) HasUseSerialQueryRate() bool`

HasUseSerialQueryRate returns a boolean if a field has been set.

### GetUseServerIdDirective

`func (o *MemberDns) GetUseServerIdDirective() bool`

GetUseServerIdDirective returns the UseServerIdDirective field if non-nil, zero value otherwise.

### GetUseServerIdDirectiveOk

`func (o *MemberDns) GetUseServerIdDirectiveOk() (*bool, bool)`

GetUseServerIdDirectiveOk returns a tuple with the UseServerIdDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseServerIdDirective

`func (o *MemberDns) SetUseServerIdDirective(v bool)`

SetUseServerIdDirective sets UseServerIdDirective field to given value.

### HasUseServerIdDirective

`func (o *MemberDns) HasUseServerIdDirective() bool`

HasUseServerIdDirective returns a boolean if a field has been set.

### GetUseSortlist

`func (o *MemberDns) GetUseSortlist() bool`

GetUseSortlist returns the UseSortlist field if non-nil, zero value otherwise.

### GetUseSortlistOk

`func (o *MemberDns) GetUseSortlistOk() (*bool, bool)`

GetUseSortlistOk returns a tuple with the UseSortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSortlist

`func (o *MemberDns) SetUseSortlist(v bool)`

SetUseSortlist sets UseSortlist field to given value.

### HasUseSortlist

`func (o *MemberDns) HasUseSortlist() bool`

HasUseSortlist returns a boolean if a field has been set.

### GetUseSourcePorts

`func (o *MemberDns) GetUseSourcePorts() bool`

GetUseSourcePorts returns the UseSourcePorts field if non-nil, zero value otherwise.

### GetUseSourcePortsOk

`func (o *MemberDns) GetUseSourcePortsOk() (*bool, bool)`

GetUseSourcePortsOk returns a tuple with the UseSourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSourcePorts

`func (o *MemberDns) SetUseSourcePorts(v bool)`

SetUseSourcePorts sets UseSourcePorts field to given value.

### HasUseSourcePorts

`func (o *MemberDns) HasUseSourcePorts() bool`

HasUseSourcePorts returns a boolean if a field has been set.

### GetUseSyslogFacility

`func (o *MemberDns) GetUseSyslogFacility() bool`

GetUseSyslogFacility returns the UseSyslogFacility field if non-nil, zero value otherwise.

### GetUseSyslogFacilityOk

`func (o *MemberDns) GetUseSyslogFacilityOk() (*bool, bool)`

GetUseSyslogFacilityOk returns a tuple with the UseSyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSyslogFacility

`func (o *MemberDns) SetUseSyslogFacility(v bool)`

SetUseSyslogFacility sets UseSyslogFacility field to given value.

### HasUseSyslogFacility

`func (o *MemberDns) HasUseSyslogFacility() bool`

HasUseSyslogFacility returns a boolean if a field has been set.

### GetUseTransfersIn

`func (o *MemberDns) GetUseTransfersIn() bool`

GetUseTransfersIn returns the UseTransfersIn field if non-nil, zero value otherwise.

### GetUseTransfersInOk

`func (o *MemberDns) GetUseTransfersInOk() (*bool, bool)`

GetUseTransfersInOk returns a tuple with the UseTransfersIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTransfersIn

`func (o *MemberDns) SetUseTransfersIn(v bool)`

SetUseTransfersIn sets UseTransfersIn field to given value.

### HasUseTransfersIn

`func (o *MemberDns) HasUseTransfersIn() bool`

HasUseTransfersIn returns a boolean if a field has been set.

### GetUseTransfersOut

`func (o *MemberDns) GetUseTransfersOut() bool`

GetUseTransfersOut returns the UseTransfersOut field if non-nil, zero value otherwise.

### GetUseTransfersOutOk

`func (o *MemberDns) GetUseTransfersOutOk() (*bool, bool)`

GetUseTransfersOutOk returns a tuple with the UseTransfersOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTransfersOut

`func (o *MemberDns) SetUseTransfersOut(v bool)`

SetUseTransfersOut sets UseTransfersOut field to given value.

### HasUseTransfersOut

`func (o *MemberDns) HasUseTransfersOut() bool`

HasUseTransfersOut returns a boolean if a field has been set.

### GetUseTransfersPerNs

`func (o *MemberDns) GetUseTransfersPerNs() bool`

GetUseTransfersPerNs returns the UseTransfersPerNs field if non-nil, zero value otherwise.

### GetUseTransfersPerNsOk

`func (o *MemberDns) GetUseTransfersPerNsOk() (*bool, bool)`

GetUseTransfersPerNsOk returns a tuple with the UseTransfersPerNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTransfersPerNs

`func (o *MemberDns) SetUseTransfersPerNs(v bool)`

SetUseTransfersPerNs sets UseTransfersPerNs field to given value.

### HasUseTransfersPerNs

`func (o *MemberDns) HasUseTransfersPerNs() bool`

HasUseTransfersPerNs returns a boolean if a field has been set.

### GetUseUpdateSetting

`func (o *MemberDns) GetUseUpdateSetting() bool`

GetUseUpdateSetting returns the UseUpdateSetting field if non-nil, zero value otherwise.

### GetUseUpdateSettingOk

`func (o *MemberDns) GetUseUpdateSettingOk() (*bool, bool)`

GetUseUpdateSettingOk returns a tuple with the UseUpdateSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateSetting

`func (o *MemberDns) SetUseUpdateSetting(v bool)`

SetUseUpdateSetting sets UseUpdateSetting field to given value.

### HasUseUpdateSetting

`func (o *MemberDns) HasUseUpdateSetting() bool`

HasUseUpdateSetting returns a boolean if a field has been set.

### GetUseZoneTransferFormat

`func (o *MemberDns) GetUseZoneTransferFormat() bool`

GetUseZoneTransferFormat returns the UseZoneTransferFormat field if non-nil, zero value otherwise.

### GetUseZoneTransferFormatOk

`func (o *MemberDns) GetUseZoneTransferFormatOk() (*bool, bool)`

GetUseZoneTransferFormatOk returns a tuple with the UseZoneTransferFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseZoneTransferFormat

`func (o *MemberDns) SetUseZoneTransferFormat(v bool)`

SetUseZoneTransferFormat sets UseZoneTransferFormat field to given value.

### HasUseZoneTransferFormat

`func (o *MemberDns) HasUseZoneTransferFormat() bool`

HasUseZoneTransferFormat returns a boolean if a field has been set.

### GetViews

`func (o *MemberDns) GetViews() []string`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *MemberDns) GetViewsOk() (*[]string, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *MemberDns) SetViews(v []string)`

SetViews sets Views field to given value.

### HasViews

`func (o *MemberDns) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


