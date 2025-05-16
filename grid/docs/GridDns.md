# GridDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddClientIpMacOptions** | Pointer to **bool** | Add custom IP, MAC and DNS View name ENDS0 options to outgoing recursive queries. | [optional] 
**AllowBulkhostDdns** | Pointer to **string** | Determines if DDNS bulk host is allowed or not. | [optional] 
**AllowGssTsigZoneUpdates** | Pointer to **bool** | Determines whether GSS-TSIG zone update is enabled for all Grid members. | [optional] 
**AllowQuery** | Pointer to [**[]GridDnsAllowQuery**](GridDnsAllowQuery.md) | Determines if queries from the specified IPv4 or IPv6 addresses and networks are allowed or not. The appliance can also use Transaction Signature (TSIG) keys to authenticate the queries. | [optional] 
**AllowRecursiveQuery** | Pointer to **bool** | Determines if the responses to recursive queries are enabled or not. | [optional] 
**AllowTransfer** | Pointer to [**[]GridDnsAllowTransfer**](GridDnsAllowTransfer.md) | Determines if zone transfers from specified IPv4 or IPv6 addresses and networks or transfers from hosts authenticated by Transaction signature (TSIG) key are allowed or not. | [optional] 
**AllowUpdate** | Pointer to [**[]GridDnsAllowUpdate**](GridDnsAllowUpdate.md) | Determines if dynamic updates from specified IPv4 or IPv6 addresses, networks or from host authenticated by TSIG key are allowed or not. | [optional] 
**AnonymizeResponseLogging** | Pointer to **bool** | Determines if the anonymization of captured DNS responses is enabled or disabled. | [optional] 
**AttackMitigation** | Pointer to [**GridDnsAttackMitigation**](GridDnsAttackMitigation.md) |  | [optional] 
**AutoBlackhole** | Pointer to [**GridDnsAutoBlackhole**](GridDnsAutoBlackhole.md) |  | [optional] 
**BindCheckNamesPolicy** | Pointer to **string** | The BIND check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This method applies only if the host name restriction policy is set to \&quot;Strict Hostname Checking\&quot;. | [optional] 
**BindHostnameDirective** | Pointer to **string** | The value of the hostname directive for BIND. | [optional] 
**BlackholeList** | Pointer to [**[]GridDnsBlackholeList**](GridDnsBlackholeList.md) | The list of IPv4 or IPv6 addresses and networks from which DNS queries are blocked. | [optional] 
**BlacklistAction** | Pointer to **string** | The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist ruleset. | [optional] 
**BlacklistLogQuery** | Pointer to **bool** | Determines if blacklist redirection queries are logged or not. | [optional] 
**BlacklistRedirectAddresses** | Pointer to **[]string** | The IP addresses the appliance includes in the response it sends in place of a blacklisted IP address. | [optional] 
**BlacklistRedirectTtl** | Pointer to **int64** | The TTL value (in seconds) of the synthetic DNS responses that result from blacklist redirection. | [optional] 
**BlacklistRulesets** | Pointer to **[]string** | The DNS Ruleset object names assigned at the Grid level for blacklist redirection. | [optional] 
**BulkHostNameTemplates** | Pointer to **[]string** | The list of bulk host name templates. There are four Infoblox predefined bulk host name templates. Template Name Template Format \&quot;Four Octets\&quot; -$1-$2-$3-$4 \&quot;Three Octets\&quot; -$2-$3-$4 \&quot;Two Octets\&quot; -$3-$4 \&quot;One Octet\&quot; -$4 | [optional] 
**CaptureDnsQueriesOnAllDomains** | Pointer to **bool** | Determines if the capture of DNS queries for all domains is enabled or disabled. | [optional] 
**CheckNamesForDdnsAndZoneTransfer** | Pointer to **bool** | Determines whether the application of BIND check-names for zone transfers and DDNS updates are enabled. | [optional] 
**ClientSubnetDomains** | Pointer to [**[]GridDnsClientSubnetDomains**](GridDnsClientSubnetDomains.md) | The list of zone domain names that are allowed or forbidden for EDNS client subnet (ECS) recursion. | [optional] 
**ClientSubnetIpv4PrefixLength** | Pointer to **int64** | Default IPv4 Source Prefix-Length used when sending queries with EDNS client subnet option. | [optional] 
**ClientSubnetIpv6PrefixLength** | Pointer to **int64** | Default IPv6 Source Prefix-Length used when sending queries with EDNS client subnet option. | [optional] 
**CopyClientIpMacOptions** | Pointer to **bool** | Copy custom IP, MAC and DNS View name ENDS0 options from incoming to outgoing recursive queries. | [optional] 
**CopyXferToNotify** | Pointer to **bool** | The allowed IPs, from the zone transfer list, added to the also-notify statement in the named.conf file. | [optional] 
**CustomRootNameServers** | Pointer to [**[]GridDnsCustomRootNameServers**](GridDnsCustomRootNameServers.md) | The list of customized root nameserver(s). You can use Internet root name servers or specify host names and IP addresses of custom root name servers. | [optional] 
**DdnsForceCreationTimestampUpdate** | Pointer to **bool** | Defines whether creation timestamp of RR should be updated &#39; when DDNS update happens even if there is no change to &#39; the RR. | [optional] 
**DdnsPrincipalGroup** | Pointer to **string** | The DDNS Principal cluster group name. | [optional] 
**DdnsPrincipalTracking** | Pointer to **bool** | Determines if the DDNS principal track is enabled or disabled. | [optional] 
**DdnsRestrictPatterns** | Pointer to **bool** | Determines if an option to restrict DDNS update request based on FQDN patterns is enabled or disabled. | [optional] 
**DdnsRestrictPatternsList** | Pointer to **[]string** | The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns. | [optional] 
**DdnsRestrictProtected** | Pointer to **bool** | Determines if an option to restrict DDNS update request to protected resource records is enabled or disabled. | [optional] 
**DdnsRestrictSecure** | Pointer to **bool** | Determines if DDNS update request for principal other than target resource record&#39;s principal is restricted. | [optional] 
**DdnsRestrictStatic** | Pointer to **bool** | Determines if an option to restrict DDNS update request to resource records which are marked as &#39;STATIC&#39; is enabled or disabled. | [optional] 
**DefaultBulkHostNameTemplate** | Pointer to **string** | Default bulk host name of a Grid DNS. | [optional] 
**DefaultTtl** | Pointer to **int64** | The default TTL value of a Grid DNS object. This interval tells the secondary how long the data can be cached. | [optional] 
**DisableEdns** | Pointer to **bool** | Determines if the EDNS0 support for queries that require recursive resolution on Grid members is enabled or not. | [optional] 
**Dns64Groups** | Pointer to **[]string** | The list of DNS64 synthesis groups associated with this Grid DNS object. | [optional] 
**DnsCacheAccelerationTtl** | Pointer to **int64** | The minimum TTL value, in seconds, that a DNS record must have in order for it to be cached by the DNS Cache Acceleration service. An integer from 1 to 65000 that represents the TTL in seconds. | [optional] 
**DnsHealthCheckAnycastControl** | Pointer to **bool** | Determines if the anycast failure (BFD session down) is enabled on member failure or not. | [optional] 
**DnsHealthCheckDomainList** | Pointer to **[]string** | The list of domain names for the DNS health check. | [optional] 
**DnsHealthCheckInterval** | Pointer to **int64** | The time interval (in seconds) for DNS health check. | [optional] 
**DnsHealthCheckRecursionFlag** | Pointer to **bool** | Determines if the recursive DNS health check is enabled or not. | [optional] 
**DnsHealthCheckRetries** | Pointer to **int64** | The number of DNS health check retries. | [optional] 
**DnsHealthCheckTimeout** | Pointer to **int64** | The DNS health check timeout interval (in seconds). | [optional] 
**DnsQueryCaptureFileTimeLimit** | Pointer to **int64** | The time limit (in minutes) for the DNS query capture file. | [optional] 
**DnssecBlacklistEnabled** | Pointer to **bool** | Determines if the blacklist rules for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecDns64Enabled** | Pointer to **bool** | Determines if the DNS64 groups for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Determines if the DNS security extension is enabled or not. | [optional] 
**DnssecExpiredSignaturesEnabled** | Pointer to **bool** | Determines when the DNS member accepts expired signatures. | [optional] 
**DnssecKeyParams** | Pointer to [**GridDnsDnssecKeyParams**](GridDnsDnssecKeyParams.md) |  | [optional] 
**DnssecNegativeTrustAnchors** | Pointer to **[]string** | A list of zones for which the server does not perform DNSSEC validation. | [optional] 
**DnssecNxdomainEnabled** | Pointer to **bool** | Determines if the NXDOMAIN rules for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecRpzEnabled** | Pointer to **bool** | Determines if the RPZ policies for DNSSEC-enabled clients are enabled or not. | [optional] 
**DnssecTrustedKeys** | Pointer to [**[]GridDnsDnssecTrustedKeys**](GridDnsDnssecTrustedKeys.md) | The list of trusted keys for the DNSSEC feature. | [optional] 
**DnssecValidationEnabled** | Pointer to **bool** | Determines if the DNS security validation is enabled or not. | [optional] 
**DnstapSetting** | Pointer to [**GridDnsDnstapSetting**](GridDnsDnstapSetting.md) |  | [optional] 
**DomainsToCaptureDnsQueries** | Pointer to **[]string** | The list of domains for DNS query capture. | [optional] 
**DtcDnsQueriesSpecificBehavior** | Pointer to **string** | Setting to control specific behavior for DTC DNS responses for incoming lbdn matched queries. | [optional] 
**DtcDnssecMode** | Pointer to **string** | DTC DNSSEC operation mode. | [optional] 
**DtcEdnsPreferClientSubnet** | Pointer to **bool** | Determines whether to prefer the client address from the edns-client-subnet option for DTC or not. | [optional] 
**DtcScheduledBackup** | Pointer to [**GridDnsDtcScheduledBackup**](GridDnsDtcScheduledBackup.md) |  | [optional] 
**DtcTopologyEaList** | Pointer to **[]string** | The DTC topology extensible attribute definition list. When configuring a DTC topology, users may configure classification as either \&quot;Geographic\&quot; or \&quot;Extensible Attributes\&quot;. Selecting extensible attributes will replace supported Topology database labels (Continent, Country, Subdivision, City) with the names of the selection EA types and provide values extracted from DHCP Network Container, Network and Range objects with those extensible attributes. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes. | [optional] 
**Email** | Pointer to **string** | The email address of a Grid DNS object. | [optional] 
**EnableBlackhole** | Pointer to **bool** | Determines if the blocking of DNS queries is enabled or not. | [optional] 
**EnableBlacklist** | Pointer to **bool** | Determines if a blacklist is enabled or not. | [optional] 
**EnableCaptureDnsQueries** | Pointer to **bool** | Determines if the capture of DNS queries is enabled or disabled. | [optional] 
**EnableCaptureDnsResponses** | Pointer to **bool** | Determines if the capture of DNS responses is enabled or disabled. | [optional] 
**EnableClientSubnetForwarding** | Pointer to **bool** | Determines whether to enable forwarding EDNS client subnet options to upstream servers. | [optional] 
**EnableClientSubnetRecursive** | Pointer to **bool** | Determines whether to enable adding EDNS client subnet options in recursive resolution. | [optional] 
**EnableDeleteAssociatedPtr** | Pointer to **bool** | Determines if the ability to automatically remove associated PTR records while deleting A or AAAA records is enabled or not. | [optional] 
**EnableDns64** | Pointer to **bool** | Determines if the DNS64 support is enabled or not. | [optional] 
**EnableDnsHealthCheck** | Pointer to **bool** | Determines if the DNS health check is enabled or not. | [optional] 
**EnableDnstapQueries** | Pointer to **bool** | Determines whether the query messages need to be forwarded to DNSTAP or not. | [optional] 
**EnableDnstapResponses** | Pointer to **bool** | Determines whether the response messages need to be forwarded to DNSTAP or not. | [optional] 
**EnableExcludedDomainNames** | Pointer to **bool** | Determines if excluding domain names from captured DNS queries and responses is enabled or disabled. | [optional] 
**EnableFixedRrsetOrderFqdns** | Pointer to **bool** | Determines if the fixed RRset order FQDN is enabled or not. | [optional] 
**EnableFtc** | Pointer to **bool** | Determines whether Fault Tolerant Caching (FTC) is enabled. | [optional] 
**EnableGssTsig** | Pointer to **bool** | Determines whether all appliances in the Grid are enabled to receive GSS-TSIG authenticated updates from DNS clients. | [optional] 
**EnableHostRrsetOrder** | Pointer to **bool** | Determines if the host RRset order is enabled or not. | [optional] 
**EnableHsmSigning** | Pointer to **bool** | Determines whether Hardware Security Modules (HSMs) are enabled for key generation and signing. Note, that you must configure the HSM group with at least one enabled HSM. | [optional] 
**EnableNotifySourcePort** | Pointer to **bool** | Determines if the notify source port at the Grid Level is enabled or not. | [optional] 
**EnableQueryRewrite** | Pointer to **bool** | Determines if the DNS query rewrite is enabled or not. | [optional] 
**EnableQuerySourcePort** | Pointer to **bool** | Determines if the query source port at the Grid Level is enabled or not. | [optional] 
**ExcludedDomainNames** | Pointer to **[]string** | The list of domains that are excluded from DNS query and response capture. | [optional] 
**ExpireAfter** | Pointer to **int64** | The expiration time of a Grid DNS object. If the secondary DNS server fails to contact the primary server for the specified interval, the secondary server stops giving out answers about the zone because the zone data is too old to be useful. | [optional] 
**FileTransferSetting** | Pointer to [**GridDnsFileTransferSetting**](GridDnsFileTransferSetting.md) |  | [optional] 
**FilterAaaa** | Pointer to **string** | The type of AAAA filtering for this member DNS object. | [optional] 
**FilterAaaaList** | Pointer to [**[]GridDnsFilterAaaaList**](GridDnsFilterAaaaList.md) | The list of IPv4 addresses and networks from which queries are received. AAAA filtering is applied to these addresses. | [optional] 
**FixedRrsetOrderFqdns** | Pointer to [**[]GridDnsFixedRrsetOrderFqdns**](GridDnsFixedRrsetOrderFqdns.md) | The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to &#39;true&#39;, unless the same request sets the enable field to &#39;false&#39;. | [optional] 
**ForwardOnly** | Pointer to **bool** | Determines if member sends queries to forwarders only. When the value is \&quot;true\&quot;, the member sends queries to forwarders only, and not to other internal or Internet root servers. | [optional] 
**ForwardUpdates** | Pointer to **bool** | Determines if secondary servers is allowed to forward updates to the DNS server or not. | [optional] 
**Forwarders** | Pointer to **[]string** | The forwarders for the member. A forwarder is essentially a name server to which other name servers first send all of their off-site queries. The forwarder builds up a cache of information, avoiding the need for the other name servers to send queries off-site. | [optional] 
**FtcExpiredRecordTimeout** | Pointer to **int64** | The timeout interval (in seconds) after which the expired Fault Tolerant Caching (FTC)record is stale and no longer valid. | [optional] 
**FtcExpiredRecordTtl** | Pointer to **int64** | The TTL value (in seconds) of the expired Fault Tolerant Caching (FTC) record in DNS responses. | [optional] 
**GenEadbFromHosts** | Pointer to **bool** | Flag for taking EA values from IPAM Hosts into consideration for the DTC topology EA database. | [optional] 
**GenEadbFromNetworkContainers** | Pointer to **bool** | Flag for taking EA values from IPAM Network Containers into consideration for the DTC topology EA database. | [optional] 
**GenEadbFromNetworks** | Pointer to **bool** | Flag for taking EA values from IPAM Network into consideration for the DTC topology EA database. | [optional] 
**GenEadbFromRanges** | Pointer to **bool** | Flag for taking EA values from IPAM Ranges into consideration for the DTC topology EA database. | [optional] 
**GssTsigKeys** | Pointer to **[]string** | The list of GSS-TSIG keys for a Grid DNS object. | [optional] 
**LastQueriedAcl** | Pointer to [**[]GridDnsLastQueriedAcl**](GridDnsLastQueriedAcl.md) | Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings. | [optional] 
**LoggingCategories** | Pointer to [**GridDnsLoggingCategories**](GridDnsLoggingCategories.md) |  | [optional] 
**MaxCacheTtl** | Pointer to **int64** | The maximum time (in seconds) for which the server will cache positive answers. | [optional] 
**MaxCachedLifetime** | Pointer to **int64** | The maximum time (in seconds) a DNS response can be stored in the hardware acceleration cache. Valid values are unsigned integer between 60 and 86400, inclusive. | [optional] 
**MaxNcacheTtl** | Pointer to **int64** | The maximum time (in seconds) for which the server will cache negative (NXDOMAIN) responses. The maximum allowed value is 604800. | [optional] 
**MaxUdpSize** | Pointer to **int64** | The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes. | [optional] 
**MemberSecondaryNotify** | Pointer to **bool** | Determines if Grid members that are authoritative secondary servers are allowed to send notification messages to external name servers, if the Grid member that is primary for a zone fails or loses connectivity. | [optional] 
**NegativeTtl** | Pointer to **int64** | The negative TTL value of a Grid DNS object. This interval tells the secondary how long data can be cached for \&quot;Does Not Respond\&quot; responses. | [optional] 
**NotifyDelay** | Pointer to **int64** | Specifies with how many seconds of delay the notify messages are sent to secondaries. | [optional] 
**NotifySourcePort** | Pointer to **int64** | The source port for notify messages. When requesting zone transfers from the primary server, some secondary DNS servers use the source port number (the primary server used to send the notify message) as the destination port number in the zone transfer request. Valid values are between 1 and 63999. The default is picked by BIND. | [optional] 
**NsgroupDefault** | Pointer to **string** | The default nameserver group. | [optional] 
**Nsgroups** | Pointer to **[]string** | A name server group is a collection of one primary DNS server and one or more secondary DNS servers. | [optional] 
**NxdomainLogQuery** | Pointer to **bool** | Determines if NXDOMAIN redirection queries are logged or not. | [optional] 
**NxdomainRedirect** | Pointer to **bool** | Determines if NXDOMAIN redirection is enabled or not. | [optional] 
**NxdomainRedirectAddresses** | Pointer to **[]string** | The list of IPv4 NXDOMAIN redirection addresses. | [optional] 
**NxdomainRedirectAddressesV6** | Pointer to **[]string** | The list of IPv6 NXDOMAIN redirection addresses. | [optional] 
**NxdomainRedirectTtl** | Pointer to **int64** | The TTL value (in seconds) of synthetic DNS responses that result from NXDOMAIN redirection. | [optional] 
**NxdomainRulesets** | Pointer to **[]string** | The Ruleset object names assigned at the Grid level for NXDOMAIN redirection. | [optional] 
**PreserveHostRrsetOrderOnSecondaries** | Pointer to **bool** | Determines if the host RRset order on secondaries is preserved or not. | [optional] 
**ProtocolRecordNamePolicies** | Pointer to **[]string** | The list of record name policies. | [optional] 
**QueryRewriteDomainNames** | Pointer to **[]string** | The list of domain names that trigger DNS query rewrite. | [optional] 
**QueryRewritePrefix** | Pointer to **string** | The domain name prefix for DNS query rewrite. | [optional] 
**QuerySourcePort** | Pointer to **int64** | The source port for queries. Specifying a source port number for recursive queries ensures that a firewall will allow the response. Valid values are between 1 and 63999. The default is picked by BIND. | [optional] 
**RecursiveQueryList** | Pointer to [**[]GridDnsRecursiveQueryList**](GridDnsRecursiveQueryList.md) | The list of IPv4 or IPv6 addresses, networks or hosts authenticated by Transaction signature (TSIG) key from which recursive queries are allowed or denied. | [optional] 
**RefreshTimer** | Pointer to **int64** | The refresh time. This interval tells the secondary how often to send a message to the primary for a zone to check that its data is current, and retrieve fresh data if it is not. | [optional] 
**ResolverQueryTimeout** | Pointer to **int64** | The recursive query timeout for the member. | [optional] 
**ResponseRateLimiting** | Pointer to [**GridDnsResponseRateLimiting**](GridDnsResponseRateLimiting.md) |  | [optional] 
**RestartSetting** | Pointer to [**GridDnsRestartSetting**](GridDnsRestartSetting.md) |  | [optional] 
**RetryTimer** | Pointer to **int64** | The retry time. This interval tells the secondary how long to wait before attempting to recontact the primary after a connection failure occurs between the two servers. | [optional] 
**RootNameServerType** | Pointer to **string** | Determines the type of root name servers. | [optional] 
**RpzDisableNsdnameNsip** | Pointer to **bool** | Determines if NSDNAME and NSIP resource records from RPZ feeds are enabled or not. | [optional] 
**RpzDropIpRuleEnabled** | Pointer to **bool** | Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv4** | Pointer to **int64** | The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv6** | Pointer to **int64** | The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length. | [optional] 
**RpzQnameWaitRecurse** | Pointer to **bool** | Determines if recursive RPZ lookups are enabled. | [optional] 
**RunScavenging** | Pointer to **map[string]interface{}** |  | [optional] 
**ScavengingSettings** | Pointer to [**GridDnsScavengingSettings**](GridDnsScavengingSettings.md) |  | [optional] 
**SerialQueryRate** | Pointer to **int64** | The number of maximum concurrent SOA queries per second. Valid values are unsigned integer between 20 and 1000, inclusive. | [optional] 
**ServerIdDirective** | Pointer to **string** | The value of the server-id directive for BIND DNS. | [optional] 
**Sortlist** | Pointer to [**[]GridDnsSortlist**](GridDnsSortlist.md) | A sort list determines the order of addresses in responses made to DNS queries. | [optional] 
**StoreLocally** | Pointer to **bool** | Determines if the storage of query capture reports on the appliance is enabled or disabled. | [optional] 
**SyslogFacility** | Pointer to **string** | The syslog facility. This is the location on the syslog server to which you want to sort the DNS logging messages. | [optional] 
**TransferExcludedServers** | Pointer to **[]string** | The list of excluded DNS servers during zone transfers. | [optional] 
**TransferFormat** | Pointer to **string** | The BIND format for a zone transfer. This provides tracking capabilities for single or multiple transfers and their associated servers. | [optional] 
**TransfersIn** | Pointer to **int64** | The number of maximum concurrent transfers for the Grid. Valid values are unsigned integer between 10 and 10000, inclusive. | [optional] 
**TransfersOut** | Pointer to **int64** | The number of maximum outbound concurrent zone transfers. Valid values are unsigned integer between 10 and 10000, inclusive. | [optional] 
**TransfersPerNs** | Pointer to **int64** | The number of maximum concurrent transfers per member. Valid values are unsigned integer between 2 and 10000, inclusive. | [optional] 
**ZoneDeletionDoubleConfirm** | Pointer to **bool** | Determines if the double confirmation during zone deletion is enabled or not. | [optional] 

## Methods

### NewGridDns

`func NewGridDns() *GridDns`

NewGridDns instantiates a new GridDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsWithDefaults

`func NewGridDnsWithDefaults() *GridDns`

NewGridDnsWithDefaults instantiates a new GridDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridDns) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridDns) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridDns) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridDns) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddClientIpMacOptions

`func (o *GridDns) GetAddClientIpMacOptions() bool`

GetAddClientIpMacOptions returns the AddClientIpMacOptions field if non-nil, zero value otherwise.

### GetAddClientIpMacOptionsOk

`func (o *GridDns) GetAddClientIpMacOptionsOk() (*bool, bool)`

GetAddClientIpMacOptionsOk returns a tuple with the AddClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddClientIpMacOptions

`func (o *GridDns) SetAddClientIpMacOptions(v bool)`

SetAddClientIpMacOptions sets AddClientIpMacOptions field to given value.

### HasAddClientIpMacOptions

`func (o *GridDns) HasAddClientIpMacOptions() bool`

HasAddClientIpMacOptions returns a boolean if a field has been set.

### GetAllowBulkhostDdns

`func (o *GridDns) GetAllowBulkhostDdns() string`

GetAllowBulkhostDdns returns the AllowBulkhostDdns field if non-nil, zero value otherwise.

### GetAllowBulkhostDdnsOk

`func (o *GridDns) GetAllowBulkhostDdnsOk() (*string, bool)`

GetAllowBulkhostDdnsOk returns a tuple with the AllowBulkhostDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBulkhostDdns

`func (o *GridDns) SetAllowBulkhostDdns(v string)`

SetAllowBulkhostDdns sets AllowBulkhostDdns field to given value.

### HasAllowBulkhostDdns

`func (o *GridDns) HasAllowBulkhostDdns() bool`

HasAllowBulkhostDdns returns a boolean if a field has been set.

### GetAllowGssTsigZoneUpdates

`func (o *GridDns) GetAllowGssTsigZoneUpdates() bool`

GetAllowGssTsigZoneUpdates returns the AllowGssTsigZoneUpdates field if non-nil, zero value otherwise.

### GetAllowGssTsigZoneUpdatesOk

`func (o *GridDns) GetAllowGssTsigZoneUpdatesOk() (*bool, bool)`

GetAllowGssTsigZoneUpdatesOk returns a tuple with the AllowGssTsigZoneUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGssTsigZoneUpdates

`func (o *GridDns) SetAllowGssTsigZoneUpdates(v bool)`

SetAllowGssTsigZoneUpdates sets AllowGssTsigZoneUpdates field to given value.

### HasAllowGssTsigZoneUpdates

`func (o *GridDns) HasAllowGssTsigZoneUpdates() bool`

HasAllowGssTsigZoneUpdates returns a boolean if a field has been set.

### GetAllowQuery

`func (o *GridDns) GetAllowQuery() []GridDnsAllowQuery`

GetAllowQuery returns the AllowQuery field if non-nil, zero value otherwise.

### GetAllowQueryOk

`func (o *GridDns) GetAllowQueryOk() (*[]GridDnsAllowQuery, bool)`

GetAllowQueryOk returns a tuple with the AllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuery

`func (o *GridDns) SetAllowQuery(v []GridDnsAllowQuery)`

SetAllowQuery sets AllowQuery field to given value.

### HasAllowQuery

`func (o *GridDns) HasAllowQuery() bool`

HasAllowQuery returns a boolean if a field has been set.

### GetAllowRecursiveQuery

`func (o *GridDns) GetAllowRecursiveQuery() bool`

GetAllowRecursiveQuery returns the AllowRecursiveQuery field if non-nil, zero value otherwise.

### GetAllowRecursiveQueryOk

`func (o *GridDns) GetAllowRecursiveQueryOk() (*bool, bool)`

GetAllowRecursiveQueryOk returns a tuple with the AllowRecursiveQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRecursiveQuery

`func (o *GridDns) SetAllowRecursiveQuery(v bool)`

SetAllowRecursiveQuery sets AllowRecursiveQuery field to given value.

### HasAllowRecursiveQuery

`func (o *GridDns) HasAllowRecursiveQuery() bool`

HasAllowRecursiveQuery returns a boolean if a field has been set.

### GetAllowTransfer

`func (o *GridDns) GetAllowTransfer() []GridDnsAllowTransfer`

GetAllowTransfer returns the AllowTransfer field if non-nil, zero value otherwise.

### GetAllowTransferOk

`func (o *GridDns) GetAllowTransferOk() (*[]GridDnsAllowTransfer, bool)`

GetAllowTransferOk returns a tuple with the AllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransfer

`func (o *GridDns) SetAllowTransfer(v []GridDnsAllowTransfer)`

SetAllowTransfer sets AllowTransfer field to given value.

### HasAllowTransfer

`func (o *GridDns) HasAllowTransfer() bool`

HasAllowTransfer returns a boolean if a field has been set.

### GetAllowUpdate

`func (o *GridDns) GetAllowUpdate() []GridDnsAllowUpdate`

GetAllowUpdate returns the AllowUpdate field if non-nil, zero value otherwise.

### GetAllowUpdateOk

`func (o *GridDns) GetAllowUpdateOk() (*[]GridDnsAllowUpdate, bool)`

GetAllowUpdateOk returns a tuple with the AllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdate

`func (o *GridDns) SetAllowUpdate(v []GridDnsAllowUpdate)`

SetAllowUpdate sets AllowUpdate field to given value.

### HasAllowUpdate

`func (o *GridDns) HasAllowUpdate() bool`

HasAllowUpdate returns a boolean if a field has been set.

### GetAnonymizeResponseLogging

`func (o *GridDns) GetAnonymizeResponseLogging() bool`

GetAnonymizeResponseLogging returns the AnonymizeResponseLogging field if non-nil, zero value otherwise.

### GetAnonymizeResponseLoggingOk

`func (o *GridDns) GetAnonymizeResponseLoggingOk() (*bool, bool)`

GetAnonymizeResponseLoggingOk returns a tuple with the AnonymizeResponseLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeResponseLogging

`func (o *GridDns) SetAnonymizeResponseLogging(v bool)`

SetAnonymizeResponseLogging sets AnonymizeResponseLogging field to given value.

### HasAnonymizeResponseLogging

`func (o *GridDns) HasAnonymizeResponseLogging() bool`

HasAnonymizeResponseLogging returns a boolean if a field has been set.

### GetAttackMitigation

`func (o *GridDns) GetAttackMitigation() GridDnsAttackMitigation`

GetAttackMitigation returns the AttackMitigation field if non-nil, zero value otherwise.

### GetAttackMitigationOk

`func (o *GridDns) GetAttackMitigationOk() (*GridDnsAttackMitigation, bool)`

GetAttackMitigationOk returns a tuple with the AttackMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackMitigation

`func (o *GridDns) SetAttackMitigation(v GridDnsAttackMitigation)`

SetAttackMitigation sets AttackMitigation field to given value.

### HasAttackMitigation

`func (o *GridDns) HasAttackMitigation() bool`

HasAttackMitigation returns a boolean if a field has been set.

### GetAutoBlackhole

`func (o *GridDns) GetAutoBlackhole() GridDnsAutoBlackhole`

GetAutoBlackhole returns the AutoBlackhole field if non-nil, zero value otherwise.

### GetAutoBlackholeOk

`func (o *GridDns) GetAutoBlackholeOk() (*GridDnsAutoBlackhole, bool)`

GetAutoBlackholeOk returns a tuple with the AutoBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlackhole

`func (o *GridDns) SetAutoBlackhole(v GridDnsAutoBlackhole)`

SetAutoBlackhole sets AutoBlackhole field to given value.

### HasAutoBlackhole

`func (o *GridDns) HasAutoBlackhole() bool`

HasAutoBlackhole returns a boolean if a field has been set.

### GetBindCheckNamesPolicy

`func (o *GridDns) GetBindCheckNamesPolicy() string`

GetBindCheckNamesPolicy returns the BindCheckNamesPolicy field if non-nil, zero value otherwise.

### GetBindCheckNamesPolicyOk

`func (o *GridDns) GetBindCheckNamesPolicyOk() (*string, bool)`

GetBindCheckNamesPolicyOk returns a tuple with the BindCheckNamesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCheckNamesPolicy

`func (o *GridDns) SetBindCheckNamesPolicy(v string)`

SetBindCheckNamesPolicy sets BindCheckNamesPolicy field to given value.

### HasBindCheckNamesPolicy

`func (o *GridDns) HasBindCheckNamesPolicy() bool`

HasBindCheckNamesPolicy returns a boolean if a field has been set.

### GetBindHostnameDirective

`func (o *GridDns) GetBindHostnameDirective() string`

GetBindHostnameDirective returns the BindHostnameDirective field if non-nil, zero value otherwise.

### GetBindHostnameDirectiveOk

`func (o *GridDns) GetBindHostnameDirectiveOk() (*string, bool)`

GetBindHostnameDirectiveOk returns a tuple with the BindHostnameDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindHostnameDirective

`func (o *GridDns) SetBindHostnameDirective(v string)`

SetBindHostnameDirective sets BindHostnameDirective field to given value.

### HasBindHostnameDirective

`func (o *GridDns) HasBindHostnameDirective() bool`

HasBindHostnameDirective returns a boolean if a field has been set.

### GetBlackholeList

`func (o *GridDns) GetBlackholeList() []GridDnsBlackholeList`

GetBlackholeList returns the BlackholeList field if non-nil, zero value otherwise.

### GetBlackholeListOk

`func (o *GridDns) GetBlackholeListOk() (*[]GridDnsBlackholeList, bool)`

GetBlackholeListOk returns a tuple with the BlackholeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackholeList

`func (o *GridDns) SetBlackholeList(v []GridDnsBlackholeList)`

SetBlackholeList sets BlackholeList field to given value.

### HasBlackholeList

`func (o *GridDns) HasBlackholeList() bool`

HasBlackholeList returns a boolean if a field has been set.

### GetBlacklistAction

`func (o *GridDns) GetBlacklistAction() string`

GetBlacklistAction returns the BlacklistAction field if non-nil, zero value otherwise.

### GetBlacklistActionOk

`func (o *GridDns) GetBlacklistActionOk() (*string, bool)`

GetBlacklistActionOk returns a tuple with the BlacklistAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistAction

`func (o *GridDns) SetBlacklistAction(v string)`

SetBlacklistAction sets BlacklistAction field to given value.

### HasBlacklistAction

`func (o *GridDns) HasBlacklistAction() bool`

HasBlacklistAction returns a boolean if a field has been set.

### GetBlacklistLogQuery

`func (o *GridDns) GetBlacklistLogQuery() bool`

GetBlacklistLogQuery returns the BlacklistLogQuery field if non-nil, zero value otherwise.

### GetBlacklistLogQueryOk

`func (o *GridDns) GetBlacklistLogQueryOk() (*bool, bool)`

GetBlacklistLogQueryOk returns a tuple with the BlacklistLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistLogQuery

`func (o *GridDns) SetBlacklistLogQuery(v bool)`

SetBlacklistLogQuery sets BlacklistLogQuery field to given value.

### HasBlacklistLogQuery

`func (o *GridDns) HasBlacklistLogQuery() bool`

HasBlacklistLogQuery returns a boolean if a field has been set.

### GetBlacklistRedirectAddresses

`func (o *GridDns) GetBlacklistRedirectAddresses() []string`

GetBlacklistRedirectAddresses returns the BlacklistRedirectAddresses field if non-nil, zero value otherwise.

### GetBlacklistRedirectAddressesOk

`func (o *GridDns) GetBlacklistRedirectAddressesOk() (*[]string, bool)`

GetBlacklistRedirectAddressesOk returns a tuple with the BlacklistRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectAddresses

`func (o *GridDns) SetBlacklistRedirectAddresses(v []string)`

SetBlacklistRedirectAddresses sets BlacklistRedirectAddresses field to given value.

### HasBlacklistRedirectAddresses

`func (o *GridDns) HasBlacklistRedirectAddresses() bool`

HasBlacklistRedirectAddresses returns a boolean if a field has been set.

### GetBlacklistRedirectTtl

`func (o *GridDns) GetBlacklistRedirectTtl() int64`

GetBlacklistRedirectTtl returns the BlacklistRedirectTtl field if non-nil, zero value otherwise.

### GetBlacklistRedirectTtlOk

`func (o *GridDns) GetBlacklistRedirectTtlOk() (*int64, bool)`

GetBlacklistRedirectTtlOk returns a tuple with the BlacklistRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectTtl

`func (o *GridDns) SetBlacklistRedirectTtl(v int64)`

SetBlacklistRedirectTtl sets BlacklistRedirectTtl field to given value.

### HasBlacklistRedirectTtl

`func (o *GridDns) HasBlacklistRedirectTtl() bool`

HasBlacklistRedirectTtl returns a boolean if a field has been set.

### GetBlacklistRulesets

`func (o *GridDns) GetBlacklistRulesets() []string`

GetBlacklistRulesets returns the BlacklistRulesets field if non-nil, zero value otherwise.

### GetBlacklistRulesetsOk

`func (o *GridDns) GetBlacklistRulesetsOk() (*[]string, bool)`

GetBlacklistRulesetsOk returns a tuple with the BlacklistRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRulesets

`func (o *GridDns) SetBlacklistRulesets(v []string)`

SetBlacklistRulesets sets BlacklistRulesets field to given value.

### HasBlacklistRulesets

`func (o *GridDns) HasBlacklistRulesets() bool`

HasBlacklistRulesets returns a boolean if a field has been set.

### GetBulkHostNameTemplates

`func (o *GridDns) GetBulkHostNameTemplates() []string`

GetBulkHostNameTemplates returns the BulkHostNameTemplates field if non-nil, zero value otherwise.

### GetBulkHostNameTemplatesOk

`func (o *GridDns) GetBulkHostNameTemplatesOk() (*[]string, bool)`

GetBulkHostNameTemplatesOk returns a tuple with the BulkHostNameTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkHostNameTemplates

`func (o *GridDns) SetBulkHostNameTemplates(v []string)`

SetBulkHostNameTemplates sets BulkHostNameTemplates field to given value.

### HasBulkHostNameTemplates

`func (o *GridDns) HasBulkHostNameTemplates() bool`

HasBulkHostNameTemplates returns a boolean if a field has been set.

### GetCaptureDnsQueriesOnAllDomains

`func (o *GridDns) GetCaptureDnsQueriesOnAllDomains() bool`

GetCaptureDnsQueriesOnAllDomains returns the CaptureDnsQueriesOnAllDomains field if non-nil, zero value otherwise.

### GetCaptureDnsQueriesOnAllDomainsOk

`func (o *GridDns) GetCaptureDnsQueriesOnAllDomainsOk() (*bool, bool)`

GetCaptureDnsQueriesOnAllDomainsOk returns a tuple with the CaptureDnsQueriesOnAllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDnsQueriesOnAllDomains

`func (o *GridDns) SetCaptureDnsQueriesOnAllDomains(v bool)`

SetCaptureDnsQueriesOnAllDomains sets CaptureDnsQueriesOnAllDomains field to given value.

### HasCaptureDnsQueriesOnAllDomains

`func (o *GridDns) HasCaptureDnsQueriesOnAllDomains() bool`

HasCaptureDnsQueriesOnAllDomains returns a boolean if a field has been set.

### GetCheckNamesForDdnsAndZoneTransfer

`func (o *GridDns) GetCheckNamesForDdnsAndZoneTransfer() bool`

GetCheckNamesForDdnsAndZoneTransfer returns the CheckNamesForDdnsAndZoneTransfer field if non-nil, zero value otherwise.

### GetCheckNamesForDdnsAndZoneTransferOk

`func (o *GridDns) GetCheckNamesForDdnsAndZoneTransferOk() (*bool, bool)`

GetCheckNamesForDdnsAndZoneTransferOk returns a tuple with the CheckNamesForDdnsAndZoneTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesForDdnsAndZoneTransfer

`func (o *GridDns) SetCheckNamesForDdnsAndZoneTransfer(v bool)`

SetCheckNamesForDdnsAndZoneTransfer sets CheckNamesForDdnsAndZoneTransfer field to given value.

### HasCheckNamesForDdnsAndZoneTransfer

`func (o *GridDns) HasCheckNamesForDdnsAndZoneTransfer() bool`

HasCheckNamesForDdnsAndZoneTransfer returns a boolean if a field has been set.

### GetClientSubnetDomains

`func (o *GridDns) GetClientSubnetDomains() []GridDnsClientSubnetDomains`

GetClientSubnetDomains returns the ClientSubnetDomains field if non-nil, zero value otherwise.

### GetClientSubnetDomainsOk

`func (o *GridDns) GetClientSubnetDomainsOk() (*[]GridDnsClientSubnetDomains, bool)`

GetClientSubnetDomainsOk returns a tuple with the ClientSubnetDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetDomains

`func (o *GridDns) SetClientSubnetDomains(v []GridDnsClientSubnetDomains)`

SetClientSubnetDomains sets ClientSubnetDomains field to given value.

### HasClientSubnetDomains

`func (o *GridDns) HasClientSubnetDomains() bool`

HasClientSubnetDomains returns a boolean if a field has been set.

### GetClientSubnetIpv4PrefixLength

`func (o *GridDns) GetClientSubnetIpv4PrefixLength() int64`

GetClientSubnetIpv4PrefixLength returns the ClientSubnetIpv4PrefixLength field if non-nil, zero value otherwise.

### GetClientSubnetIpv4PrefixLengthOk

`func (o *GridDns) GetClientSubnetIpv4PrefixLengthOk() (*int64, bool)`

GetClientSubnetIpv4PrefixLengthOk returns a tuple with the ClientSubnetIpv4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetIpv4PrefixLength

`func (o *GridDns) SetClientSubnetIpv4PrefixLength(v int64)`

SetClientSubnetIpv4PrefixLength sets ClientSubnetIpv4PrefixLength field to given value.

### HasClientSubnetIpv4PrefixLength

`func (o *GridDns) HasClientSubnetIpv4PrefixLength() bool`

HasClientSubnetIpv4PrefixLength returns a boolean if a field has been set.

### GetClientSubnetIpv6PrefixLength

`func (o *GridDns) GetClientSubnetIpv6PrefixLength() int64`

GetClientSubnetIpv6PrefixLength returns the ClientSubnetIpv6PrefixLength field if non-nil, zero value otherwise.

### GetClientSubnetIpv6PrefixLengthOk

`func (o *GridDns) GetClientSubnetIpv6PrefixLengthOk() (*int64, bool)`

GetClientSubnetIpv6PrefixLengthOk returns a tuple with the ClientSubnetIpv6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetIpv6PrefixLength

`func (o *GridDns) SetClientSubnetIpv6PrefixLength(v int64)`

SetClientSubnetIpv6PrefixLength sets ClientSubnetIpv6PrefixLength field to given value.

### HasClientSubnetIpv6PrefixLength

`func (o *GridDns) HasClientSubnetIpv6PrefixLength() bool`

HasClientSubnetIpv6PrefixLength returns a boolean if a field has been set.

### GetCopyClientIpMacOptions

`func (o *GridDns) GetCopyClientIpMacOptions() bool`

GetCopyClientIpMacOptions returns the CopyClientIpMacOptions field if non-nil, zero value otherwise.

### GetCopyClientIpMacOptionsOk

`func (o *GridDns) GetCopyClientIpMacOptionsOk() (*bool, bool)`

GetCopyClientIpMacOptionsOk returns a tuple with the CopyClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyClientIpMacOptions

`func (o *GridDns) SetCopyClientIpMacOptions(v bool)`

SetCopyClientIpMacOptions sets CopyClientIpMacOptions field to given value.

### HasCopyClientIpMacOptions

`func (o *GridDns) HasCopyClientIpMacOptions() bool`

HasCopyClientIpMacOptions returns a boolean if a field has been set.

### GetCopyXferToNotify

`func (o *GridDns) GetCopyXferToNotify() bool`

GetCopyXferToNotify returns the CopyXferToNotify field if non-nil, zero value otherwise.

### GetCopyXferToNotifyOk

`func (o *GridDns) GetCopyXferToNotifyOk() (*bool, bool)`

GetCopyXferToNotifyOk returns a tuple with the CopyXferToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyXferToNotify

`func (o *GridDns) SetCopyXferToNotify(v bool)`

SetCopyXferToNotify sets CopyXferToNotify field to given value.

### HasCopyXferToNotify

`func (o *GridDns) HasCopyXferToNotify() bool`

HasCopyXferToNotify returns a boolean if a field has been set.

### GetCustomRootNameServers

`func (o *GridDns) GetCustomRootNameServers() []GridDnsCustomRootNameServers`

GetCustomRootNameServers returns the CustomRootNameServers field if non-nil, zero value otherwise.

### GetCustomRootNameServersOk

`func (o *GridDns) GetCustomRootNameServersOk() (*[]GridDnsCustomRootNameServers, bool)`

GetCustomRootNameServersOk returns a tuple with the CustomRootNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNameServers

`func (o *GridDns) SetCustomRootNameServers(v []GridDnsCustomRootNameServers)`

SetCustomRootNameServers sets CustomRootNameServers field to given value.

### HasCustomRootNameServers

`func (o *GridDns) HasCustomRootNameServers() bool`

HasCustomRootNameServers returns a boolean if a field has been set.

### GetDdnsForceCreationTimestampUpdate

`func (o *GridDns) GetDdnsForceCreationTimestampUpdate() bool`

GetDdnsForceCreationTimestampUpdate returns the DdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetDdnsForceCreationTimestampUpdateOk

`func (o *GridDns) GetDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetDdnsForceCreationTimestampUpdateOk returns a tuple with the DdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsForceCreationTimestampUpdate

`func (o *GridDns) SetDdnsForceCreationTimestampUpdate(v bool)`

SetDdnsForceCreationTimestampUpdate sets DdnsForceCreationTimestampUpdate field to given value.

### HasDdnsForceCreationTimestampUpdate

`func (o *GridDns) HasDdnsForceCreationTimestampUpdate() bool`

HasDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetDdnsPrincipalGroup

`func (o *GridDns) GetDdnsPrincipalGroup() string`

GetDdnsPrincipalGroup returns the DdnsPrincipalGroup field if non-nil, zero value otherwise.

### GetDdnsPrincipalGroupOk

`func (o *GridDns) GetDdnsPrincipalGroupOk() (*string, bool)`

GetDdnsPrincipalGroupOk returns a tuple with the DdnsPrincipalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalGroup

`func (o *GridDns) SetDdnsPrincipalGroup(v string)`

SetDdnsPrincipalGroup sets DdnsPrincipalGroup field to given value.

### HasDdnsPrincipalGroup

`func (o *GridDns) HasDdnsPrincipalGroup() bool`

HasDdnsPrincipalGroup returns a boolean if a field has been set.

### GetDdnsPrincipalTracking

`func (o *GridDns) GetDdnsPrincipalTracking() bool`

GetDdnsPrincipalTracking returns the DdnsPrincipalTracking field if non-nil, zero value otherwise.

### GetDdnsPrincipalTrackingOk

`func (o *GridDns) GetDdnsPrincipalTrackingOk() (*bool, bool)`

GetDdnsPrincipalTrackingOk returns a tuple with the DdnsPrincipalTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalTracking

`func (o *GridDns) SetDdnsPrincipalTracking(v bool)`

SetDdnsPrincipalTracking sets DdnsPrincipalTracking field to given value.

### HasDdnsPrincipalTracking

`func (o *GridDns) HasDdnsPrincipalTracking() bool`

HasDdnsPrincipalTracking returns a boolean if a field has been set.

### GetDdnsRestrictPatterns

`func (o *GridDns) GetDdnsRestrictPatterns() bool`

GetDdnsRestrictPatterns returns the DdnsRestrictPatterns field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsOk

`func (o *GridDns) GetDdnsRestrictPatternsOk() (*bool, bool)`

GetDdnsRestrictPatternsOk returns a tuple with the DdnsRestrictPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatterns

`func (o *GridDns) SetDdnsRestrictPatterns(v bool)`

SetDdnsRestrictPatterns sets DdnsRestrictPatterns field to given value.

### HasDdnsRestrictPatterns

`func (o *GridDns) HasDdnsRestrictPatterns() bool`

HasDdnsRestrictPatterns returns a boolean if a field has been set.

### GetDdnsRestrictPatternsList

`func (o *GridDns) GetDdnsRestrictPatternsList() []string`

GetDdnsRestrictPatternsList returns the DdnsRestrictPatternsList field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsListOk

`func (o *GridDns) GetDdnsRestrictPatternsListOk() (*[]string, bool)`

GetDdnsRestrictPatternsListOk returns a tuple with the DdnsRestrictPatternsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatternsList

`func (o *GridDns) SetDdnsRestrictPatternsList(v []string)`

SetDdnsRestrictPatternsList sets DdnsRestrictPatternsList field to given value.

### HasDdnsRestrictPatternsList

`func (o *GridDns) HasDdnsRestrictPatternsList() bool`

HasDdnsRestrictPatternsList returns a boolean if a field has been set.

### GetDdnsRestrictProtected

`func (o *GridDns) GetDdnsRestrictProtected() bool`

GetDdnsRestrictProtected returns the DdnsRestrictProtected field if non-nil, zero value otherwise.

### GetDdnsRestrictProtectedOk

`func (o *GridDns) GetDdnsRestrictProtectedOk() (*bool, bool)`

GetDdnsRestrictProtectedOk returns a tuple with the DdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictProtected

`func (o *GridDns) SetDdnsRestrictProtected(v bool)`

SetDdnsRestrictProtected sets DdnsRestrictProtected field to given value.

### HasDdnsRestrictProtected

`func (o *GridDns) HasDdnsRestrictProtected() bool`

HasDdnsRestrictProtected returns a boolean if a field has been set.

### GetDdnsRestrictSecure

`func (o *GridDns) GetDdnsRestrictSecure() bool`

GetDdnsRestrictSecure returns the DdnsRestrictSecure field if non-nil, zero value otherwise.

### GetDdnsRestrictSecureOk

`func (o *GridDns) GetDdnsRestrictSecureOk() (*bool, bool)`

GetDdnsRestrictSecureOk returns a tuple with the DdnsRestrictSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictSecure

`func (o *GridDns) SetDdnsRestrictSecure(v bool)`

SetDdnsRestrictSecure sets DdnsRestrictSecure field to given value.

### HasDdnsRestrictSecure

`func (o *GridDns) HasDdnsRestrictSecure() bool`

HasDdnsRestrictSecure returns a boolean if a field has been set.

### GetDdnsRestrictStatic

`func (o *GridDns) GetDdnsRestrictStatic() bool`

GetDdnsRestrictStatic returns the DdnsRestrictStatic field if non-nil, zero value otherwise.

### GetDdnsRestrictStaticOk

`func (o *GridDns) GetDdnsRestrictStaticOk() (*bool, bool)`

GetDdnsRestrictStaticOk returns a tuple with the DdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictStatic

`func (o *GridDns) SetDdnsRestrictStatic(v bool)`

SetDdnsRestrictStatic sets DdnsRestrictStatic field to given value.

### HasDdnsRestrictStatic

`func (o *GridDns) HasDdnsRestrictStatic() bool`

HasDdnsRestrictStatic returns a boolean if a field has been set.

### GetDefaultBulkHostNameTemplate

`func (o *GridDns) GetDefaultBulkHostNameTemplate() string`

GetDefaultBulkHostNameTemplate returns the DefaultBulkHostNameTemplate field if non-nil, zero value otherwise.

### GetDefaultBulkHostNameTemplateOk

`func (o *GridDns) GetDefaultBulkHostNameTemplateOk() (*string, bool)`

GetDefaultBulkHostNameTemplateOk returns a tuple with the DefaultBulkHostNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBulkHostNameTemplate

`func (o *GridDns) SetDefaultBulkHostNameTemplate(v string)`

SetDefaultBulkHostNameTemplate sets DefaultBulkHostNameTemplate field to given value.

### HasDefaultBulkHostNameTemplate

`func (o *GridDns) HasDefaultBulkHostNameTemplate() bool`

HasDefaultBulkHostNameTemplate returns a boolean if a field has been set.

### GetDefaultTtl

`func (o *GridDns) GetDefaultTtl() int64`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *GridDns) GetDefaultTtlOk() (*int64, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *GridDns) SetDefaultTtl(v int64)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *GridDns) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetDisableEdns

`func (o *GridDns) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *GridDns) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *GridDns) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.

### HasDisableEdns

`func (o *GridDns) HasDisableEdns() bool`

HasDisableEdns returns a boolean if a field has been set.

### GetDns64Groups

`func (o *GridDns) GetDns64Groups() []string`

GetDns64Groups returns the Dns64Groups field if non-nil, zero value otherwise.

### GetDns64GroupsOk

`func (o *GridDns) GetDns64GroupsOk() (*[]string, bool)`

GetDns64GroupsOk returns a tuple with the Dns64Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Groups

`func (o *GridDns) SetDns64Groups(v []string)`

SetDns64Groups sets Dns64Groups field to given value.

### HasDns64Groups

`func (o *GridDns) HasDns64Groups() bool`

HasDns64Groups returns a boolean if a field has been set.

### GetDnsCacheAccelerationTtl

`func (o *GridDns) GetDnsCacheAccelerationTtl() int64`

GetDnsCacheAccelerationTtl returns the DnsCacheAccelerationTtl field if non-nil, zero value otherwise.

### GetDnsCacheAccelerationTtlOk

`func (o *GridDns) GetDnsCacheAccelerationTtlOk() (*int64, bool)`

GetDnsCacheAccelerationTtlOk returns a tuple with the DnsCacheAccelerationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCacheAccelerationTtl

`func (o *GridDns) SetDnsCacheAccelerationTtl(v int64)`

SetDnsCacheAccelerationTtl sets DnsCacheAccelerationTtl field to given value.

### HasDnsCacheAccelerationTtl

`func (o *GridDns) HasDnsCacheAccelerationTtl() bool`

HasDnsCacheAccelerationTtl returns a boolean if a field has been set.

### GetDnsHealthCheckAnycastControl

`func (o *GridDns) GetDnsHealthCheckAnycastControl() bool`

GetDnsHealthCheckAnycastControl returns the DnsHealthCheckAnycastControl field if non-nil, zero value otherwise.

### GetDnsHealthCheckAnycastControlOk

`func (o *GridDns) GetDnsHealthCheckAnycastControlOk() (*bool, bool)`

GetDnsHealthCheckAnycastControlOk returns a tuple with the DnsHealthCheckAnycastControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckAnycastControl

`func (o *GridDns) SetDnsHealthCheckAnycastControl(v bool)`

SetDnsHealthCheckAnycastControl sets DnsHealthCheckAnycastControl field to given value.

### HasDnsHealthCheckAnycastControl

`func (o *GridDns) HasDnsHealthCheckAnycastControl() bool`

HasDnsHealthCheckAnycastControl returns a boolean if a field has been set.

### GetDnsHealthCheckDomainList

`func (o *GridDns) GetDnsHealthCheckDomainList() []string`

GetDnsHealthCheckDomainList returns the DnsHealthCheckDomainList field if non-nil, zero value otherwise.

### GetDnsHealthCheckDomainListOk

`func (o *GridDns) GetDnsHealthCheckDomainListOk() (*[]string, bool)`

GetDnsHealthCheckDomainListOk returns a tuple with the DnsHealthCheckDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckDomainList

`func (o *GridDns) SetDnsHealthCheckDomainList(v []string)`

SetDnsHealthCheckDomainList sets DnsHealthCheckDomainList field to given value.

### HasDnsHealthCheckDomainList

`func (o *GridDns) HasDnsHealthCheckDomainList() bool`

HasDnsHealthCheckDomainList returns a boolean if a field has been set.

### GetDnsHealthCheckInterval

`func (o *GridDns) GetDnsHealthCheckInterval() int64`

GetDnsHealthCheckInterval returns the DnsHealthCheckInterval field if non-nil, zero value otherwise.

### GetDnsHealthCheckIntervalOk

`func (o *GridDns) GetDnsHealthCheckIntervalOk() (*int64, bool)`

GetDnsHealthCheckIntervalOk returns a tuple with the DnsHealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckInterval

`func (o *GridDns) SetDnsHealthCheckInterval(v int64)`

SetDnsHealthCheckInterval sets DnsHealthCheckInterval field to given value.

### HasDnsHealthCheckInterval

`func (o *GridDns) HasDnsHealthCheckInterval() bool`

HasDnsHealthCheckInterval returns a boolean if a field has been set.

### GetDnsHealthCheckRecursionFlag

`func (o *GridDns) GetDnsHealthCheckRecursionFlag() bool`

GetDnsHealthCheckRecursionFlag returns the DnsHealthCheckRecursionFlag field if non-nil, zero value otherwise.

### GetDnsHealthCheckRecursionFlagOk

`func (o *GridDns) GetDnsHealthCheckRecursionFlagOk() (*bool, bool)`

GetDnsHealthCheckRecursionFlagOk returns a tuple with the DnsHealthCheckRecursionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckRecursionFlag

`func (o *GridDns) SetDnsHealthCheckRecursionFlag(v bool)`

SetDnsHealthCheckRecursionFlag sets DnsHealthCheckRecursionFlag field to given value.

### HasDnsHealthCheckRecursionFlag

`func (o *GridDns) HasDnsHealthCheckRecursionFlag() bool`

HasDnsHealthCheckRecursionFlag returns a boolean if a field has been set.

### GetDnsHealthCheckRetries

`func (o *GridDns) GetDnsHealthCheckRetries() int64`

GetDnsHealthCheckRetries returns the DnsHealthCheckRetries field if non-nil, zero value otherwise.

### GetDnsHealthCheckRetriesOk

`func (o *GridDns) GetDnsHealthCheckRetriesOk() (*int64, bool)`

GetDnsHealthCheckRetriesOk returns a tuple with the DnsHealthCheckRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckRetries

`func (o *GridDns) SetDnsHealthCheckRetries(v int64)`

SetDnsHealthCheckRetries sets DnsHealthCheckRetries field to given value.

### HasDnsHealthCheckRetries

`func (o *GridDns) HasDnsHealthCheckRetries() bool`

HasDnsHealthCheckRetries returns a boolean if a field has been set.

### GetDnsHealthCheckTimeout

`func (o *GridDns) GetDnsHealthCheckTimeout() int64`

GetDnsHealthCheckTimeout returns the DnsHealthCheckTimeout field if non-nil, zero value otherwise.

### GetDnsHealthCheckTimeoutOk

`func (o *GridDns) GetDnsHealthCheckTimeoutOk() (*int64, bool)`

GetDnsHealthCheckTimeoutOk returns a tuple with the DnsHealthCheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckTimeout

`func (o *GridDns) SetDnsHealthCheckTimeout(v int64)`

SetDnsHealthCheckTimeout sets DnsHealthCheckTimeout field to given value.

### HasDnsHealthCheckTimeout

`func (o *GridDns) HasDnsHealthCheckTimeout() bool`

HasDnsHealthCheckTimeout returns a boolean if a field has been set.

### GetDnsQueryCaptureFileTimeLimit

`func (o *GridDns) GetDnsQueryCaptureFileTimeLimit() int64`

GetDnsQueryCaptureFileTimeLimit returns the DnsQueryCaptureFileTimeLimit field if non-nil, zero value otherwise.

### GetDnsQueryCaptureFileTimeLimitOk

`func (o *GridDns) GetDnsQueryCaptureFileTimeLimitOk() (*int64, bool)`

GetDnsQueryCaptureFileTimeLimitOk returns a tuple with the DnsQueryCaptureFileTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueryCaptureFileTimeLimit

`func (o *GridDns) SetDnsQueryCaptureFileTimeLimit(v int64)`

SetDnsQueryCaptureFileTimeLimit sets DnsQueryCaptureFileTimeLimit field to given value.

### HasDnsQueryCaptureFileTimeLimit

`func (o *GridDns) HasDnsQueryCaptureFileTimeLimit() bool`

HasDnsQueryCaptureFileTimeLimit returns a boolean if a field has been set.

### GetDnssecBlacklistEnabled

`func (o *GridDns) GetDnssecBlacklistEnabled() bool`

GetDnssecBlacklistEnabled returns the DnssecBlacklistEnabled field if non-nil, zero value otherwise.

### GetDnssecBlacklistEnabledOk

`func (o *GridDns) GetDnssecBlacklistEnabledOk() (*bool, bool)`

GetDnssecBlacklistEnabledOk returns a tuple with the DnssecBlacklistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecBlacklistEnabled

`func (o *GridDns) SetDnssecBlacklistEnabled(v bool)`

SetDnssecBlacklistEnabled sets DnssecBlacklistEnabled field to given value.

### HasDnssecBlacklistEnabled

`func (o *GridDns) HasDnssecBlacklistEnabled() bool`

HasDnssecBlacklistEnabled returns a boolean if a field has been set.

### GetDnssecDns64Enabled

`func (o *GridDns) GetDnssecDns64Enabled() bool`

GetDnssecDns64Enabled returns the DnssecDns64Enabled field if non-nil, zero value otherwise.

### GetDnssecDns64EnabledOk

`func (o *GridDns) GetDnssecDns64EnabledOk() (*bool, bool)`

GetDnssecDns64EnabledOk returns a tuple with the DnssecDns64Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecDns64Enabled

`func (o *GridDns) SetDnssecDns64Enabled(v bool)`

SetDnssecDns64Enabled sets DnssecDns64Enabled field to given value.

### HasDnssecDns64Enabled

`func (o *GridDns) HasDnssecDns64Enabled() bool`

HasDnssecDns64Enabled returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *GridDns) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *GridDns) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *GridDns) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *GridDns) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecExpiredSignaturesEnabled

`func (o *GridDns) GetDnssecExpiredSignaturesEnabled() bool`

GetDnssecExpiredSignaturesEnabled returns the DnssecExpiredSignaturesEnabled field if non-nil, zero value otherwise.

### GetDnssecExpiredSignaturesEnabledOk

`func (o *GridDns) GetDnssecExpiredSignaturesEnabledOk() (*bool, bool)`

GetDnssecExpiredSignaturesEnabledOk returns a tuple with the DnssecExpiredSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecExpiredSignaturesEnabled

`func (o *GridDns) SetDnssecExpiredSignaturesEnabled(v bool)`

SetDnssecExpiredSignaturesEnabled sets DnssecExpiredSignaturesEnabled field to given value.

### HasDnssecExpiredSignaturesEnabled

`func (o *GridDns) HasDnssecExpiredSignaturesEnabled() bool`

HasDnssecExpiredSignaturesEnabled returns a boolean if a field has been set.

### GetDnssecKeyParams

`func (o *GridDns) GetDnssecKeyParams() GridDnsDnssecKeyParams`

GetDnssecKeyParams returns the DnssecKeyParams field if non-nil, zero value otherwise.

### GetDnssecKeyParamsOk

`func (o *GridDns) GetDnssecKeyParamsOk() (*GridDnsDnssecKeyParams, bool)`

GetDnssecKeyParamsOk returns a tuple with the DnssecKeyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKeyParams

`func (o *GridDns) SetDnssecKeyParams(v GridDnsDnssecKeyParams)`

SetDnssecKeyParams sets DnssecKeyParams field to given value.

### HasDnssecKeyParams

`func (o *GridDns) HasDnssecKeyParams() bool`

HasDnssecKeyParams returns a boolean if a field has been set.

### GetDnssecNegativeTrustAnchors

`func (o *GridDns) GetDnssecNegativeTrustAnchors() []string`

GetDnssecNegativeTrustAnchors returns the DnssecNegativeTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecNegativeTrustAnchorsOk

`func (o *GridDns) GetDnssecNegativeTrustAnchorsOk() (*[]string, bool)`

GetDnssecNegativeTrustAnchorsOk returns a tuple with the DnssecNegativeTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNegativeTrustAnchors

`func (o *GridDns) SetDnssecNegativeTrustAnchors(v []string)`

SetDnssecNegativeTrustAnchors sets DnssecNegativeTrustAnchors field to given value.

### HasDnssecNegativeTrustAnchors

`func (o *GridDns) HasDnssecNegativeTrustAnchors() bool`

HasDnssecNegativeTrustAnchors returns a boolean if a field has been set.

### GetDnssecNxdomainEnabled

`func (o *GridDns) GetDnssecNxdomainEnabled() bool`

GetDnssecNxdomainEnabled returns the DnssecNxdomainEnabled field if non-nil, zero value otherwise.

### GetDnssecNxdomainEnabledOk

`func (o *GridDns) GetDnssecNxdomainEnabledOk() (*bool, bool)`

GetDnssecNxdomainEnabledOk returns a tuple with the DnssecNxdomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNxdomainEnabled

`func (o *GridDns) SetDnssecNxdomainEnabled(v bool)`

SetDnssecNxdomainEnabled sets DnssecNxdomainEnabled field to given value.

### HasDnssecNxdomainEnabled

`func (o *GridDns) HasDnssecNxdomainEnabled() bool`

HasDnssecNxdomainEnabled returns a boolean if a field has been set.

### GetDnssecRpzEnabled

`func (o *GridDns) GetDnssecRpzEnabled() bool`

GetDnssecRpzEnabled returns the DnssecRpzEnabled field if non-nil, zero value otherwise.

### GetDnssecRpzEnabledOk

`func (o *GridDns) GetDnssecRpzEnabledOk() (*bool, bool)`

GetDnssecRpzEnabledOk returns a tuple with the DnssecRpzEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRpzEnabled

`func (o *GridDns) SetDnssecRpzEnabled(v bool)`

SetDnssecRpzEnabled sets DnssecRpzEnabled field to given value.

### HasDnssecRpzEnabled

`func (o *GridDns) HasDnssecRpzEnabled() bool`

HasDnssecRpzEnabled returns a boolean if a field has been set.

### GetDnssecTrustedKeys

`func (o *GridDns) GetDnssecTrustedKeys() []GridDnsDnssecTrustedKeys`

GetDnssecTrustedKeys returns the DnssecTrustedKeys field if non-nil, zero value otherwise.

### GetDnssecTrustedKeysOk

`func (o *GridDns) GetDnssecTrustedKeysOk() (*[]GridDnsDnssecTrustedKeys, bool)`

GetDnssecTrustedKeysOk returns a tuple with the DnssecTrustedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustedKeys

`func (o *GridDns) SetDnssecTrustedKeys(v []GridDnsDnssecTrustedKeys)`

SetDnssecTrustedKeys sets DnssecTrustedKeys field to given value.

### HasDnssecTrustedKeys

`func (o *GridDns) HasDnssecTrustedKeys() bool`

HasDnssecTrustedKeys returns a boolean if a field has been set.

### GetDnssecValidationEnabled

`func (o *GridDns) GetDnssecValidationEnabled() bool`

GetDnssecValidationEnabled returns the DnssecValidationEnabled field if non-nil, zero value otherwise.

### GetDnssecValidationEnabledOk

`func (o *GridDns) GetDnssecValidationEnabledOk() (*bool, bool)`

GetDnssecValidationEnabledOk returns a tuple with the DnssecValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationEnabled

`func (o *GridDns) SetDnssecValidationEnabled(v bool)`

SetDnssecValidationEnabled sets DnssecValidationEnabled field to given value.

### HasDnssecValidationEnabled

`func (o *GridDns) HasDnssecValidationEnabled() bool`

HasDnssecValidationEnabled returns a boolean if a field has been set.

### GetDnstapSetting

`func (o *GridDns) GetDnstapSetting() GridDnsDnstapSetting`

GetDnstapSetting returns the DnstapSetting field if non-nil, zero value otherwise.

### GetDnstapSettingOk

`func (o *GridDns) GetDnstapSettingOk() (*GridDnsDnstapSetting, bool)`

GetDnstapSettingOk returns a tuple with the DnstapSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapSetting

`func (o *GridDns) SetDnstapSetting(v GridDnsDnstapSetting)`

SetDnstapSetting sets DnstapSetting field to given value.

### HasDnstapSetting

`func (o *GridDns) HasDnstapSetting() bool`

HasDnstapSetting returns a boolean if a field has been set.

### GetDomainsToCaptureDnsQueries

`func (o *GridDns) GetDomainsToCaptureDnsQueries() []string`

GetDomainsToCaptureDnsQueries returns the DomainsToCaptureDnsQueries field if non-nil, zero value otherwise.

### GetDomainsToCaptureDnsQueriesOk

`func (o *GridDns) GetDomainsToCaptureDnsQueriesOk() (*[]string, bool)`

GetDomainsToCaptureDnsQueriesOk returns a tuple with the DomainsToCaptureDnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsToCaptureDnsQueries

`func (o *GridDns) SetDomainsToCaptureDnsQueries(v []string)`

SetDomainsToCaptureDnsQueries sets DomainsToCaptureDnsQueries field to given value.

### HasDomainsToCaptureDnsQueries

`func (o *GridDns) HasDomainsToCaptureDnsQueries() bool`

HasDomainsToCaptureDnsQueries returns a boolean if a field has been set.

### GetDtcDnsQueriesSpecificBehavior

`func (o *GridDns) GetDtcDnsQueriesSpecificBehavior() string`

GetDtcDnsQueriesSpecificBehavior returns the DtcDnsQueriesSpecificBehavior field if non-nil, zero value otherwise.

### GetDtcDnsQueriesSpecificBehaviorOk

`func (o *GridDns) GetDtcDnsQueriesSpecificBehaviorOk() (*string, bool)`

GetDtcDnsQueriesSpecificBehaviorOk returns a tuple with the DtcDnsQueriesSpecificBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcDnsQueriesSpecificBehavior

`func (o *GridDns) SetDtcDnsQueriesSpecificBehavior(v string)`

SetDtcDnsQueriesSpecificBehavior sets DtcDnsQueriesSpecificBehavior field to given value.

### HasDtcDnsQueriesSpecificBehavior

`func (o *GridDns) HasDtcDnsQueriesSpecificBehavior() bool`

HasDtcDnsQueriesSpecificBehavior returns a boolean if a field has been set.

### GetDtcDnssecMode

`func (o *GridDns) GetDtcDnssecMode() string`

GetDtcDnssecMode returns the DtcDnssecMode field if non-nil, zero value otherwise.

### GetDtcDnssecModeOk

`func (o *GridDns) GetDtcDnssecModeOk() (*string, bool)`

GetDtcDnssecModeOk returns a tuple with the DtcDnssecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcDnssecMode

`func (o *GridDns) SetDtcDnssecMode(v string)`

SetDtcDnssecMode sets DtcDnssecMode field to given value.

### HasDtcDnssecMode

`func (o *GridDns) HasDtcDnssecMode() bool`

HasDtcDnssecMode returns a boolean if a field has been set.

### GetDtcEdnsPreferClientSubnet

`func (o *GridDns) GetDtcEdnsPreferClientSubnet() bool`

GetDtcEdnsPreferClientSubnet returns the DtcEdnsPreferClientSubnet field if non-nil, zero value otherwise.

### GetDtcEdnsPreferClientSubnetOk

`func (o *GridDns) GetDtcEdnsPreferClientSubnetOk() (*bool, bool)`

GetDtcEdnsPreferClientSubnetOk returns a tuple with the DtcEdnsPreferClientSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcEdnsPreferClientSubnet

`func (o *GridDns) SetDtcEdnsPreferClientSubnet(v bool)`

SetDtcEdnsPreferClientSubnet sets DtcEdnsPreferClientSubnet field to given value.

### HasDtcEdnsPreferClientSubnet

`func (o *GridDns) HasDtcEdnsPreferClientSubnet() bool`

HasDtcEdnsPreferClientSubnet returns a boolean if a field has been set.

### GetDtcScheduledBackup

`func (o *GridDns) GetDtcScheduledBackup() GridDnsDtcScheduledBackup`

GetDtcScheduledBackup returns the DtcScheduledBackup field if non-nil, zero value otherwise.

### GetDtcScheduledBackupOk

`func (o *GridDns) GetDtcScheduledBackupOk() (*GridDnsDtcScheduledBackup, bool)`

GetDtcScheduledBackupOk returns a tuple with the DtcScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcScheduledBackup

`func (o *GridDns) SetDtcScheduledBackup(v GridDnsDtcScheduledBackup)`

SetDtcScheduledBackup sets DtcScheduledBackup field to given value.

### HasDtcScheduledBackup

`func (o *GridDns) HasDtcScheduledBackup() bool`

HasDtcScheduledBackup returns a boolean if a field has been set.

### GetDtcTopologyEaList

`func (o *GridDns) GetDtcTopologyEaList() []string`

GetDtcTopologyEaList returns the DtcTopologyEaList field if non-nil, zero value otherwise.

### GetDtcTopologyEaListOk

`func (o *GridDns) GetDtcTopologyEaListOk() (*[]string, bool)`

GetDtcTopologyEaListOk returns a tuple with the DtcTopologyEaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcTopologyEaList

`func (o *GridDns) SetDtcTopologyEaList(v []string)`

SetDtcTopologyEaList sets DtcTopologyEaList field to given value.

### HasDtcTopologyEaList

`func (o *GridDns) HasDtcTopologyEaList() bool`

HasDtcTopologyEaList returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *GridDns) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *GridDns) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *GridDns) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *GridDns) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetEmail

`func (o *GridDns) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GridDns) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GridDns) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GridDns) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableBlackhole

`func (o *GridDns) GetEnableBlackhole() bool`

GetEnableBlackhole returns the EnableBlackhole field if non-nil, zero value otherwise.

### GetEnableBlackholeOk

`func (o *GridDns) GetEnableBlackholeOk() (*bool, bool)`

GetEnableBlackholeOk returns a tuple with the EnableBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackhole

`func (o *GridDns) SetEnableBlackhole(v bool)`

SetEnableBlackhole sets EnableBlackhole field to given value.

### HasEnableBlackhole

`func (o *GridDns) HasEnableBlackhole() bool`

HasEnableBlackhole returns a boolean if a field has been set.

### GetEnableBlacklist

`func (o *GridDns) GetEnableBlacklist() bool`

GetEnableBlacklist returns the EnableBlacklist field if non-nil, zero value otherwise.

### GetEnableBlacklistOk

`func (o *GridDns) GetEnableBlacklistOk() (*bool, bool)`

GetEnableBlacklistOk returns a tuple with the EnableBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlacklist

`func (o *GridDns) SetEnableBlacklist(v bool)`

SetEnableBlacklist sets EnableBlacklist field to given value.

### HasEnableBlacklist

`func (o *GridDns) HasEnableBlacklist() bool`

HasEnableBlacklist returns a boolean if a field has been set.

### GetEnableCaptureDnsQueries

`func (o *GridDns) GetEnableCaptureDnsQueries() bool`

GetEnableCaptureDnsQueries returns the EnableCaptureDnsQueries field if non-nil, zero value otherwise.

### GetEnableCaptureDnsQueriesOk

`func (o *GridDns) GetEnableCaptureDnsQueriesOk() (*bool, bool)`

GetEnableCaptureDnsQueriesOk returns a tuple with the EnableCaptureDnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaptureDnsQueries

`func (o *GridDns) SetEnableCaptureDnsQueries(v bool)`

SetEnableCaptureDnsQueries sets EnableCaptureDnsQueries field to given value.

### HasEnableCaptureDnsQueries

`func (o *GridDns) HasEnableCaptureDnsQueries() bool`

HasEnableCaptureDnsQueries returns a boolean if a field has been set.

### GetEnableCaptureDnsResponses

`func (o *GridDns) GetEnableCaptureDnsResponses() bool`

GetEnableCaptureDnsResponses returns the EnableCaptureDnsResponses field if non-nil, zero value otherwise.

### GetEnableCaptureDnsResponsesOk

`func (o *GridDns) GetEnableCaptureDnsResponsesOk() (*bool, bool)`

GetEnableCaptureDnsResponsesOk returns a tuple with the EnableCaptureDnsResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaptureDnsResponses

`func (o *GridDns) SetEnableCaptureDnsResponses(v bool)`

SetEnableCaptureDnsResponses sets EnableCaptureDnsResponses field to given value.

### HasEnableCaptureDnsResponses

`func (o *GridDns) HasEnableCaptureDnsResponses() bool`

HasEnableCaptureDnsResponses returns a boolean if a field has been set.

### GetEnableClientSubnetForwarding

`func (o *GridDns) GetEnableClientSubnetForwarding() bool`

GetEnableClientSubnetForwarding returns the EnableClientSubnetForwarding field if non-nil, zero value otherwise.

### GetEnableClientSubnetForwardingOk

`func (o *GridDns) GetEnableClientSubnetForwardingOk() (*bool, bool)`

GetEnableClientSubnetForwardingOk returns a tuple with the EnableClientSubnetForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientSubnetForwarding

`func (o *GridDns) SetEnableClientSubnetForwarding(v bool)`

SetEnableClientSubnetForwarding sets EnableClientSubnetForwarding field to given value.

### HasEnableClientSubnetForwarding

`func (o *GridDns) HasEnableClientSubnetForwarding() bool`

HasEnableClientSubnetForwarding returns a boolean if a field has been set.

### GetEnableClientSubnetRecursive

`func (o *GridDns) GetEnableClientSubnetRecursive() bool`

GetEnableClientSubnetRecursive returns the EnableClientSubnetRecursive field if non-nil, zero value otherwise.

### GetEnableClientSubnetRecursiveOk

`func (o *GridDns) GetEnableClientSubnetRecursiveOk() (*bool, bool)`

GetEnableClientSubnetRecursiveOk returns a tuple with the EnableClientSubnetRecursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientSubnetRecursive

`func (o *GridDns) SetEnableClientSubnetRecursive(v bool)`

SetEnableClientSubnetRecursive sets EnableClientSubnetRecursive field to given value.

### HasEnableClientSubnetRecursive

`func (o *GridDns) HasEnableClientSubnetRecursive() bool`

HasEnableClientSubnetRecursive returns a boolean if a field has been set.

### GetEnableDeleteAssociatedPtr

`func (o *GridDns) GetEnableDeleteAssociatedPtr() bool`

GetEnableDeleteAssociatedPtr returns the EnableDeleteAssociatedPtr field if non-nil, zero value otherwise.

### GetEnableDeleteAssociatedPtrOk

`func (o *GridDns) GetEnableDeleteAssociatedPtrOk() (*bool, bool)`

GetEnableDeleteAssociatedPtrOk returns a tuple with the EnableDeleteAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeleteAssociatedPtr

`func (o *GridDns) SetEnableDeleteAssociatedPtr(v bool)`

SetEnableDeleteAssociatedPtr sets EnableDeleteAssociatedPtr field to given value.

### HasEnableDeleteAssociatedPtr

`func (o *GridDns) HasEnableDeleteAssociatedPtr() bool`

HasEnableDeleteAssociatedPtr returns a boolean if a field has been set.

### GetEnableDns64

`func (o *GridDns) GetEnableDns64() bool`

GetEnableDns64 returns the EnableDns64 field if non-nil, zero value otherwise.

### GetEnableDns64Ok

`func (o *GridDns) GetEnableDns64Ok() (*bool, bool)`

GetEnableDns64Ok returns a tuple with the EnableDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns64

`func (o *GridDns) SetEnableDns64(v bool)`

SetEnableDns64 sets EnableDns64 field to given value.

### HasEnableDns64

`func (o *GridDns) HasEnableDns64() bool`

HasEnableDns64 returns a boolean if a field has been set.

### GetEnableDnsHealthCheck

`func (o *GridDns) GetEnableDnsHealthCheck() bool`

GetEnableDnsHealthCheck returns the EnableDnsHealthCheck field if non-nil, zero value otherwise.

### GetEnableDnsHealthCheckOk

`func (o *GridDns) GetEnableDnsHealthCheckOk() (*bool, bool)`

GetEnableDnsHealthCheckOk returns a tuple with the EnableDnsHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsHealthCheck

`func (o *GridDns) SetEnableDnsHealthCheck(v bool)`

SetEnableDnsHealthCheck sets EnableDnsHealthCheck field to given value.

### HasEnableDnsHealthCheck

`func (o *GridDns) HasEnableDnsHealthCheck() bool`

HasEnableDnsHealthCheck returns a boolean if a field has been set.

### GetEnableDnstapQueries

`func (o *GridDns) GetEnableDnstapQueries() bool`

GetEnableDnstapQueries returns the EnableDnstapQueries field if non-nil, zero value otherwise.

### GetEnableDnstapQueriesOk

`func (o *GridDns) GetEnableDnstapQueriesOk() (*bool, bool)`

GetEnableDnstapQueriesOk returns a tuple with the EnableDnstapQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapQueries

`func (o *GridDns) SetEnableDnstapQueries(v bool)`

SetEnableDnstapQueries sets EnableDnstapQueries field to given value.

### HasEnableDnstapQueries

`func (o *GridDns) HasEnableDnstapQueries() bool`

HasEnableDnstapQueries returns a boolean if a field has been set.

### GetEnableDnstapResponses

`func (o *GridDns) GetEnableDnstapResponses() bool`

GetEnableDnstapResponses returns the EnableDnstapResponses field if non-nil, zero value otherwise.

### GetEnableDnstapResponsesOk

`func (o *GridDns) GetEnableDnstapResponsesOk() (*bool, bool)`

GetEnableDnstapResponsesOk returns a tuple with the EnableDnstapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapResponses

`func (o *GridDns) SetEnableDnstapResponses(v bool)`

SetEnableDnstapResponses sets EnableDnstapResponses field to given value.

### HasEnableDnstapResponses

`func (o *GridDns) HasEnableDnstapResponses() bool`

HasEnableDnstapResponses returns a boolean if a field has been set.

### GetEnableExcludedDomainNames

`func (o *GridDns) GetEnableExcludedDomainNames() bool`

GetEnableExcludedDomainNames returns the EnableExcludedDomainNames field if non-nil, zero value otherwise.

### GetEnableExcludedDomainNamesOk

`func (o *GridDns) GetEnableExcludedDomainNamesOk() (*bool, bool)`

GetEnableExcludedDomainNamesOk returns a tuple with the EnableExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExcludedDomainNames

`func (o *GridDns) SetEnableExcludedDomainNames(v bool)`

SetEnableExcludedDomainNames sets EnableExcludedDomainNames field to given value.

### HasEnableExcludedDomainNames

`func (o *GridDns) HasEnableExcludedDomainNames() bool`

HasEnableExcludedDomainNames returns a boolean if a field has been set.

### GetEnableFixedRrsetOrderFqdns

`func (o *GridDns) GetEnableFixedRrsetOrderFqdns() bool`

GetEnableFixedRrsetOrderFqdns returns the EnableFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetEnableFixedRrsetOrderFqdnsOk

`func (o *GridDns) GetEnableFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetEnableFixedRrsetOrderFqdnsOk returns a tuple with the EnableFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixedRrsetOrderFqdns

`func (o *GridDns) SetEnableFixedRrsetOrderFqdns(v bool)`

SetEnableFixedRrsetOrderFqdns sets EnableFixedRrsetOrderFqdns field to given value.

### HasEnableFixedRrsetOrderFqdns

`func (o *GridDns) HasEnableFixedRrsetOrderFqdns() bool`

HasEnableFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetEnableFtc

`func (o *GridDns) GetEnableFtc() bool`

GetEnableFtc returns the EnableFtc field if non-nil, zero value otherwise.

### GetEnableFtcOk

`func (o *GridDns) GetEnableFtcOk() (*bool, bool)`

GetEnableFtcOk returns a tuple with the EnableFtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtc

`func (o *GridDns) SetEnableFtc(v bool)`

SetEnableFtc sets EnableFtc field to given value.

### HasEnableFtc

`func (o *GridDns) HasEnableFtc() bool`

HasEnableFtc returns a boolean if a field has been set.

### GetEnableGssTsig

`func (o *GridDns) GetEnableGssTsig() bool`

GetEnableGssTsig returns the EnableGssTsig field if non-nil, zero value otherwise.

### GetEnableGssTsigOk

`func (o *GridDns) GetEnableGssTsigOk() (*bool, bool)`

GetEnableGssTsigOk returns a tuple with the EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGssTsig

`func (o *GridDns) SetEnableGssTsig(v bool)`

SetEnableGssTsig sets EnableGssTsig field to given value.

### HasEnableGssTsig

`func (o *GridDns) HasEnableGssTsig() bool`

HasEnableGssTsig returns a boolean if a field has been set.

### GetEnableHostRrsetOrder

`func (o *GridDns) GetEnableHostRrsetOrder() bool`

GetEnableHostRrsetOrder returns the EnableHostRrsetOrder field if non-nil, zero value otherwise.

### GetEnableHostRrsetOrderOk

`func (o *GridDns) GetEnableHostRrsetOrderOk() (*bool, bool)`

GetEnableHostRrsetOrderOk returns a tuple with the EnableHostRrsetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHostRrsetOrder

`func (o *GridDns) SetEnableHostRrsetOrder(v bool)`

SetEnableHostRrsetOrder sets EnableHostRrsetOrder field to given value.

### HasEnableHostRrsetOrder

`func (o *GridDns) HasEnableHostRrsetOrder() bool`

HasEnableHostRrsetOrder returns a boolean if a field has been set.

### GetEnableHsmSigning

`func (o *GridDns) GetEnableHsmSigning() bool`

GetEnableHsmSigning returns the EnableHsmSigning field if non-nil, zero value otherwise.

### GetEnableHsmSigningOk

`func (o *GridDns) GetEnableHsmSigningOk() (*bool, bool)`

GetEnableHsmSigningOk returns a tuple with the EnableHsmSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHsmSigning

`func (o *GridDns) SetEnableHsmSigning(v bool)`

SetEnableHsmSigning sets EnableHsmSigning field to given value.

### HasEnableHsmSigning

`func (o *GridDns) HasEnableHsmSigning() bool`

HasEnableHsmSigning returns a boolean if a field has been set.

### GetEnableNotifySourcePort

`func (o *GridDns) GetEnableNotifySourcePort() bool`

GetEnableNotifySourcePort returns the EnableNotifySourcePort field if non-nil, zero value otherwise.

### GetEnableNotifySourcePortOk

`func (o *GridDns) GetEnableNotifySourcePortOk() (*bool, bool)`

GetEnableNotifySourcePortOk returns a tuple with the EnableNotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySourcePort

`func (o *GridDns) SetEnableNotifySourcePort(v bool)`

SetEnableNotifySourcePort sets EnableNotifySourcePort field to given value.

### HasEnableNotifySourcePort

`func (o *GridDns) HasEnableNotifySourcePort() bool`

HasEnableNotifySourcePort returns a boolean if a field has been set.

### GetEnableQueryRewrite

`func (o *GridDns) GetEnableQueryRewrite() bool`

GetEnableQueryRewrite returns the EnableQueryRewrite field if non-nil, zero value otherwise.

### GetEnableQueryRewriteOk

`func (o *GridDns) GetEnableQueryRewriteOk() (*bool, bool)`

GetEnableQueryRewriteOk returns a tuple with the EnableQueryRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryRewrite

`func (o *GridDns) SetEnableQueryRewrite(v bool)`

SetEnableQueryRewrite sets EnableQueryRewrite field to given value.

### HasEnableQueryRewrite

`func (o *GridDns) HasEnableQueryRewrite() bool`

HasEnableQueryRewrite returns a boolean if a field has been set.

### GetEnableQuerySourcePort

`func (o *GridDns) GetEnableQuerySourcePort() bool`

GetEnableQuerySourcePort returns the EnableQuerySourcePort field if non-nil, zero value otherwise.

### GetEnableQuerySourcePortOk

`func (o *GridDns) GetEnableQuerySourcePortOk() (*bool, bool)`

GetEnableQuerySourcePortOk returns a tuple with the EnableQuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuerySourcePort

`func (o *GridDns) SetEnableQuerySourcePort(v bool)`

SetEnableQuerySourcePort sets EnableQuerySourcePort field to given value.

### HasEnableQuerySourcePort

`func (o *GridDns) HasEnableQuerySourcePort() bool`

HasEnableQuerySourcePort returns a boolean if a field has been set.

### GetExcludedDomainNames

`func (o *GridDns) GetExcludedDomainNames() []string`

GetExcludedDomainNames returns the ExcludedDomainNames field if non-nil, zero value otherwise.

### GetExcludedDomainNamesOk

`func (o *GridDns) GetExcludedDomainNamesOk() (*[]string, bool)`

GetExcludedDomainNamesOk returns a tuple with the ExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDomainNames

`func (o *GridDns) SetExcludedDomainNames(v []string)`

SetExcludedDomainNames sets ExcludedDomainNames field to given value.

### HasExcludedDomainNames

`func (o *GridDns) HasExcludedDomainNames() bool`

HasExcludedDomainNames returns a boolean if a field has been set.

### GetExpireAfter

`func (o *GridDns) GetExpireAfter() int64`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *GridDns) GetExpireAfterOk() (*int64, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *GridDns) SetExpireAfter(v int64)`

SetExpireAfter sets ExpireAfter field to given value.

### HasExpireAfter

`func (o *GridDns) HasExpireAfter() bool`

HasExpireAfter returns a boolean if a field has been set.

### GetFileTransferSetting

`func (o *GridDns) GetFileTransferSetting() GridDnsFileTransferSetting`

GetFileTransferSetting returns the FileTransferSetting field if non-nil, zero value otherwise.

### GetFileTransferSettingOk

`func (o *GridDns) GetFileTransferSettingOk() (*GridDnsFileTransferSetting, bool)`

GetFileTransferSettingOk returns a tuple with the FileTransferSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferSetting

`func (o *GridDns) SetFileTransferSetting(v GridDnsFileTransferSetting)`

SetFileTransferSetting sets FileTransferSetting field to given value.

### HasFileTransferSetting

`func (o *GridDns) HasFileTransferSetting() bool`

HasFileTransferSetting returns a boolean if a field has been set.

### GetFilterAaaa

`func (o *GridDns) GetFilterAaaa() string`

GetFilterAaaa returns the FilterAaaa field if non-nil, zero value otherwise.

### GetFilterAaaaOk

`func (o *GridDns) GetFilterAaaaOk() (*string, bool)`

GetFilterAaaaOk returns a tuple with the FilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaa

`func (o *GridDns) SetFilterAaaa(v string)`

SetFilterAaaa sets FilterAaaa field to given value.

### HasFilterAaaa

`func (o *GridDns) HasFilterAaaa() bool`

HasFilterAaaa returns a boolean if a field has been set.

### GetFilterAaaaList

`func (o *GridDns) GetFilterAaaaList() []GridDnsFilterAaaaList`

GetFilterAaaaList returns the FilterAaaaList field if non-nil, zero value otherwise.

### GetFilterAaaaListOk

`func (o *GridDns) GetFilterAaaaListOk() (*[]GridDnsFilterAaaaList, bool)`

GetFilterAaaaListOk returns a tuple with the FilterAaaaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaList

`func (o *GridDns) SetFilterAaaaList(v []GridDnsFilterAaaaList)`

SetFilterAaaaList sets FilterAaaaList field to given value.

### HasFilterAaaaList

`func (o *GridDns) HasFilterAaaaList() bool`

HasFilterAaaaList returns a boolean if a field has been set.

### GetFixedRrsetOrderFqdns

`func (o *GridDns) GetFixedRrsetOrderFqdns() []GridDnsFixedRrsetOrderFqdns`

GetFixedRrsetOrderFqdns returns the FixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetFixedRrsetOrderFqdnsOk

`func (o *GridDns) GetFixedRrsetOrderFqdnsOk() (*[]GridDnsFixedRrsetOrderFqdns, bool)`

GetFixedRrsetOrderFqdnsOk returns a tuple with the FixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedRrsetOrderFqdns

`func (o *GridDns) SetFixedRrsetOrderFqdns(v []GridDnsFixedRrsetOrderFqdns)`

SetFixedRrsetOrderFqdns sets FixedRrsetOrderFqdns field to given value.

### HasFixedRrsetOrderFqdns

`func (o *GridDns) HasFixedRrsetOrderFqdns() bool`

HasFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetForwardOnly

`func (o *GridDns) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *GridDns) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *GridDns) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *GridDns) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetForwardUpdates

`func (o *GridDns) GetForwardUpdates() bool`

GetForwardUpdates returns the ForwardUpdates field if non-nil, zero value otherwise.

### GetForwardUpdatesOk

`func (o *GridDns) GetForwardUpdatesOk() (*bool, bool)`

GetForwardUpdatesOk returns a tuple with the ForwardUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUpdates

`func (o *GridDns) SetForwardUpdates(v bool)`

SetForwardUpdates sets ForwardUpdates field to given value.

### HasForwardUpdates

`func (o *GridDns) HasForwardUpdates() bool`

HasForwardUpdates returns a boolean if a field has been set.

### GetForwarders

`func (o *GridDns) GetForwarders() []string`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *GridDns) GetForwardersOk() (*[]string, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *GridDns) SetForwarders(v []string)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *GridDns) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetFtcExpiredRecordTimeout

`func (o *GridDns) GetFtcExpiredRecordTimeout() int64`

GetFtcExpiredRecordTimeout returns the FtcExpiredRecordTimeout field if non-nil, zero value otherwise.

### GetFtcExpiredRecordTimeoutOk

`func (o *GridDns) GetFtcExpiredRecordTimeoutOk() (*int64, bool)`

GetFtcExpiredRecordTimeoutOk returns a tuple with the FtcExpiredRecordTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtcExpiredRecordTimeout

`func (o *GridDns) SetFtcExpiredRecordTimeout(v int64)`

SetFtcExpiredRecordTimeout sets FtcExpiredRecordTimeout field to given value.

### HasFtcExpiredRecordTimeout

`func (o *GridDns) HasFtcExpiredRecordTimeout() bool`

HasFtcExpiredRecordTimeout returns a boolean if a field has been set.

### GetFtcExpiredRecordTtl

`func (o *GridDns) GetFtcExpiredRecordTtl() int64`

GetFtcExpiredRecordTtl returns the FtcExpiredRecordTtl field if non-nil, zero value otherwise.

### GetFtcExpiredRecordTtlOk

`func (o *GridDns) GetFtcExpiredRecordTtlOk() (*int64, bool)`

GetFtcExpiredRecordTtlOk returns a tuple with the FtcExpiredRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtcExpiredRecordTtl

`func (o *GridDns) SetFtcExpiredRecordTtl(v int64)`

SetFtcExpiredRecordTtl sets FtcExpiredRecordTtl field to given value.

### HasFtcExpiredRecordTtl

`func (o *GridDns) HasFtcExpiredRecordTtl() bool`

HasFtcExpiredRecordTtl returns a boolean if a field has been set.

### GetGenEadbFromHosts

`func (o *GridDns) GetGenEadbFromHosts() bool`

GetGenEadbFromHosts returns the GenEadbFromHosts field if non-nil, zero value otherwise.

### GetGenEadbFromHostsOk

`func (o *GridDns) GetGenEadbFromHostsOk() (*bool, bool)`

GetGenEadbFromHostsOk returns a tuple with the GenEadbFromHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromHosts

`func (o *GridDns) SetGenEadbFromHosts(v bool)`

SetGenEadbFromHosts sets GenEadbFromHosts field to given value.

### HasGenEadbFromHosts

`func (o *GridDns) HasGenEadbFromHosts() bool`

HasGenEadbFromHosts returns a boolean if a field has been set.

### GetGenEadbFromNetworkContainers

`func (o *GridDns) GetGenEadbFromNetworkContainers() bool`

GetGenEadbFromNetworkContainers returns the GenEadbFromNetworkContainers field if non-nil, zero value otherwise.

### GetGenEadbFromNetworkContainersOk

`func (o *GridDns) GetGenEadbFromNetworkContainersOk() (*bool, bool)`

GetGenEadbFromNetworkContainersOk returns a tuple with the GenEadbFromNetworkContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromNetworkContainers

`func (o *GridDns) SetGenEadbFromNetworkContainers(v bool)`

SetGenEadbFromNetworkContainers sets GenEadbFromNetworkContainers field to given value.

### HasGenEadbFromNetworkContainers

`func (o *GridDns) HasGenEadbFromNetworkContainers() bool`

HasGenEadbFromNetworkContainers returns a boolean if a field has been set.

### GetGenEadbFromNetworks

`func (o *GridDns) GetGenEadbFromNetworks() bool`

GetGenEadbFromNetworks returns the GenEadbFromNetworks field if non-nil, zero value otherwise.

### GetGenEadbFromNetworksOk

`func (o *GridDns) GetGenEadbFromNetworksOk() (*bool, bool)`

GetGenEadbFromNetworksOk returns a tuple with the GenEadbFromNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromNetworks

`func (o *GridDns) SetGenEadbFromNetworks(v bool)`

SetGenEadbFromNetworks sets GenEadbFromNetworks field to given value.

### HasGenEadbFromNetworks

`func (o *GridDns) HasGenEadbFromNetworks() bool`

HasGenEadbFromNetworks returns a boolean if a field has been set.

### GetGenEadbFromRanges

`func (o *GridDns) GetGenEadbFromRanges() bool`

GetGenEadbFromRanges returns the GenEadbFromRanges field if non-nil, zero value otherwise.

### GetGenEadbFromRangesOk

`func (o *GridDns) GetGenEadbFromRangesOk() (*bool, bool)`

GetGenEadbFromRangesOk returns a tuple with the GenEadbFromRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromRanges

`func (o *GridDns) SetGenEadbFromRanges(v bool)`

SetGenEadbFromRanges sets GenEadbFromRanges field to given value.

### HasGenEadbFromRanges

`func (o *GridDns) HasGenEadbFromRanges() bool`

HasGenEadbFromRanges returns a boolean if a field has been set.

### GetGssTsigKeys

`func (o *GridDns) GetGssTsigKeys() []string`

GetGssTsigKeys returns the GssTsigKeys field if non-nil, zero value otherwise.

### GetGssTsigKeysOk

`func (o *GridDns) GetGssTsigKeysOk() (*[]string, bool)`

GetGssTsigKeysOk returns a tuple with the GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigKeys

`func (o *GridDns) SetGssTsigKeys(v []string)`

SetGssTsigKeys sets GssTsigKeys field to given value.

### HasGssTsigKeys

`func (o *GridDns) HasGssTsigKeys() bool`

HasGssTsigKeys returns a boolean if a field has been set.

### GetLastQueriedAcl

`func (o *GridDns) GetLastQueriedAcl() []GridDnsLastQueriedAcl`

GetLastQueriedAcl returns the LastQueriedAcl field if non-nil, zero value otherwise.

### GetLastQueriedAclOk

`func (o *GridDns) GetLastQueriedAclOk() (*[]GridDnsLastQueriedAcl, bool)`

GetLastQueriedAclOk returns a tuple with the LastQueriedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueriedAcl

`func (o *GridDns) SetLastQueriedAcl(v []GridDnsLastQueriedAcl)`

SetLastQueriedAcl sets LastQueriedAcl field to given value.

### HasLastQueriedAcl

`func (o *GridDns) HasLastQueriedAcl() bool`

HasLastQueriedAcl returns a boolean if a field has been set.

### GetLoggingCategories

`func (o *GridDns) GetLoggingCategories() GridDnsLoggingCategories`

GetLoggingCategories returns the LoggingCategories field if non-nil, zero value otherwise.

### GetLoggingCategoriesOk

`func (o *GridDns) GetLoggingCategoriesOk() (*GridDnsLoggingCategories, bool)`

GetLoggingCategoriesOk returns a tuple with the LoggingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCategories

`func (o *GridDns) SetLoggingCategories(v GridDnsLoggingCategories)`

SetLoggingCategories sets LoggingCategories field to given value.

### HasLoggingCategories

`func (o *GridDns) HasLoggingCategories() bool`

HasLoggingCategories returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *GridDns) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *GridDns) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *GridDns) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *GridDns) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxCachedLifetime

`func (o *GridDns) GetMaxCachedLifetime() int64`

GetMaxCachedLifetime returns the MaxCachedLifetime field if non-nil, zero value otherwise.

### GetMaxCachedLifetimeOk

`func (o *GridDns) GetMaxCachedLifetimeOk() (*int64, bool)`

GetMaxCachedLifetimeOk returns a tuple with the MaxCachedLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCachedLifetime

`func (o *GridDns) SetMaxCachedLifetime(v int64)`

SetMaxCachedLifetime sets MaxCachedLifetime field to given value.

### HasMaxCachedLifetime

`func (o *GridDns) HasMaxCachedLifetime() bool`

HasMaxCachedLifetime returns a boolean if a field has been set.

### GetMaxNcacheTtl

`func (o *GridDns) GetMaxNcacheTtl() int64`

GetMaxNcacheTtl returns the MaxNcacheTtl field if non-nil, zero value otherwise.

### GetMaxNcacheTtlOk

`func (o *GridDns) GetMaxNcacheTtlOk() (*int64, bool)`

GetMaxNcacheTtlOk returns a tuple with the MaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNcacheTtl

`func (o *GridDns) SetMaxNcacheTtl(v int64)`

SetMaxNcacheTtl sets MaxNcacheTtl field to given value.

### HasMaxNcacheTtl

`func (o *GridDns) HasMaxNcacheTtl() bool`

HasMaxNcacheTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *GridDns) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *GridDns) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *GridDns) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *GridDns) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMemberSecondaryNotify

`func (o *GridDns) GetMemberSecondaryNotify() bool`

GetMemberSecondaryNotify returns the MemberSecondaryNotify field if non-nil, zero value otherwise.

### GetMemberSecondaryNotifyOk

`func (o *GridDns) GetMemberSecondaryNotifyOk() (*bool, bool)`

GetMemberSecondaryNotifyOk returns a tuple with the MemberSecondaryNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSecondaryNotify

`func (o *GridDns) SetMemberSecondaryNotify(v bool)`

SetMemberSecondaryNotify sets MemberSecondaryNotify field to given value.

### HasMemberSecondaryNotify

`func (o *GridDns) HasMemberSecondaryNotify() bool`

HasMemberSecondaryNotify returns a boolean if a field has been set.

### GetNegativeTtl

`func (o *GridDns) GetNegativeTtl() int64`

GetNegativeTtl returns the NegativeTtl field if non-nil, zero value otherwise.

### GetNegativeTtlOk

`func (o *GridDns) GetNegativeTtlOk() (*int64, bool)`

GetNegativeTtlOk returns a tuple with the NegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeTtl

`func (o *GridDns) SetNegativeTtl(v int64)`

SetNegativeTtl sets NegativeTtl field to given value.

### HasNegativeTtl

`func (o *GridDns) HasNegativeTtl() bool`

HasNegativeTtl returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *GridDns) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *GridDns) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *GridDns) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *GridDns) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetNotifySourcePort

`func (o *GridDns) GetNotifySourcePort() int64`

GetNotifySourcePort returns the NotifySourcePort field if non-nil, zero value otherwise.

### GetNotifySourcePortOk

`func (o *GridDns) GetNotifySourcePortOk() (*int64, bool)`

GetNotifySourcePortOk returns a tuple with the NotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySourcePort

`func (o *GridDns) SetNotifySourcePort(v int64)`

SetNotifySourcePort sets NotifySourcePort field to given value.

### HasNotifySourcePort

`func (o *GridDns) HasNotifySourcePort() bool`

HasNotifySourcePort returns a boolean if a field has been set.

### GetNsgroupDefault

`func (o *GridDns) GetNsgroupDefault() string`

GetNsgroupDefault returns the NsgroupDefault field if non-nil, zero value otherwise.

### GetNsgroupDefaultOk

`func (o *GridDns) GetNsgroupDefaultOk() (*string, bool)`

GetNsgroupDefaultOk returns a tuple with the NsgroupDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgroupDefault

`func (o *GridDns) SetNsgroupDefault(v string)`

SetNsgroupDefault sets NsgroupDefault field to given value.

### HasNsgroupDefault

`func (o *GridDns) HasNsgroupDefault() bool`

HasNsgroupDefault returns a boolean if a field has been set.

### GetNsgroups

`func (o *GridDns) GetNsgroups() []string`

GetNsgroups returns the Nsgroups field if non-nil, zero value otherwise.

### GetNsgroupsOk

`func (o *GridDns) GetNsgroupsOk() (*[]string, bool)`

GetNsgroupsOk returns a tuple with the Nsgroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgroups

`func (o *GridDns) SetNsgroups(v []string)`

SetNsgroups sets Nsgroups field to given value.

### HasNsgroups

`func (o *GridDns) HasNsgroups() bool`

HasNsgroups returns a boolean if a field has been set.

### GetNxdomainLogQuery

`func (o *GridDns) GetNxdomainLogQuery() bool`

GetNxdomainLogQuery returns the NxdomainLogQuery field if non-nil, zero value otherwise.

### GetNxdomainLogQueryOk

`func (o *GridDns) GetNxdomainLogQueryOk() (*bool, bool)`

GetNxdomainLogQueryOk returns a tuple with the NxdomainLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainLogQuery

`func (o *GridDns) SetNxdomainLogQuery(v bool)`

SetNxdomainLogQuery sets NxdomainLogQuery field to given value.

### HasNxdomainLogQuery

`func (o *GridDns) HasNxdomainLogQuery() bool`

HasNxdomainLogQuery returns a boolean if a field has been set.

### GetNxdomainRedirect

`func (o *GridDns) GetNxdomainRedirect() bool`

GetNxdomainRedirect returns the NxdomainRedirect field if non-nil, zero value otherwise.

### GetNxdomainRedirectOk

`func (o *GridDns) GetNxdomainRedirectOk() (*bool, bool)`

GetNxdomainRedirectOk returns a tuple with the NxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirect

`func (o *GridDns) SetNxdomainRedirect(v bool)`

SetNxdomainRedirect sets NxdomainRedirect field to given value.

### HasNxdomainRedirect

`func (o *GridDns) HasNxdomainRedirect() bool`

HasNxdomainRedirect returns a boolean if a field has been set.

### GetNxdomainRedirectAddresses

`func (o *GridDns) GetNxdomainRedirectAddresses() []string`

GetNxdomainRedirectAddresses returns the NxdomainRedirectAddresses field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesOk

`func (o *GridDns) GetNxdomainRedirectAddressesOk() (*[]string, bool)`

GetNxdomainRedirectAddressesOk returns a tuple with the NxdomainRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddresses

`func (o *GridDns) SetNxdomainRedirectAddresses(v []string)`

SetNxdomainRedirectAddresses sets NxdomainRedirectAddresses field to given value.

### HasNxdomainRedirectAddresses

`func (o *GridDns) HasNxdomainRedirectAddresses() bool`

HasNxdomainRedirectAddresses returns a boolean if a field has been set.

### GetNxdomainRedirectAddressesV6

`func (o *GridDns) GetNxdomainRedirectAddressesV6() []string`

GetNxdomainRedirectAddressesV6 returns the NxdomainRedirectAddressesV6 field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesV6Ok

`func (o *GridDns) GetNxdomainRedirectAddressesV6Ok() (*[]string, bool)`

GetNxdomainRedirectAddressesV6Ok returns a tuple with the NxdomainRedirectAddressesV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddressesV6

`func (o *GridDns) SetNxdomainRedirectAddressesV6(v []string)`

SetNxdomainRedirectAddressesV6 sets NxdomainRedirectAddressesV6 field to given value.

### HasNxdomainRedirectAddressesV6

`func (o *GridDns) HasNxdomainRedirectAddressesV6() bool`

HasNxdomainRedirectAddressesV6 returns a boolean if a field has been set.

### GetNxdomainRedirectTtl

`func (o *GridDns) GetNxdomainRedirectTtl() int64`

GetNxdomainRedirectTtl returns the NxdomainRedirectTtl field if non-nil, zero value otherwise.

### GetNxdomainRedirectTtlOk

`func (o *GridDns) GetNxdomainRedirectTtlOk() (*int64, bool)`

GetNxdomainRedirectTtlOk returns a tuple with the NxdomainRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectTtl

`func (o *GridDns) SetNxdomainRedirectTtl(v int64)`

SetNxdomainRedirectTtl sets NxdomainRedirectTtl field to given value.

### HasNxdomainRedirectTtl

`func (o *GridDns) HasNxdomainRedirectTtl() bool`

HasNxdomainRedirectTtl returns a boolean if a field has been set.

### GetNxdomainRulesets

`func (o *GridDns) GetNxdomainRulesets() []string`

GetNxdomainRulesets returns the NxdomainRulesets field if non-nil, zero value otherwise.

### GetNxdomainRulesetsOk

`func (o *GridDns) GetNxdomainRulesetsOk() (*[]string, bool)`

GetNxdomainRulesetsOk returns a tuple with the NxdomainRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRulesets

`func (o *GridDns) SetNxdomainRulesets(v []string)`

SetNxdomainRulesets sets NxdomainRulesets field to given value.

### HasNxdomainRulesets

`func (o *GridDns) HasNxdomainRulesets() bool`

HasNxdomainRulesets returns a boolean if a field has been set.

### GetPreserveHostRrsetOrderOnSecondaries

`func (o *GridDns) GetPreserveHostRrsetOrderOnSecondaries() bool`

GetPreserveHostRrsetOrderOnSecondaries returns the PreserveHostRrsetOrderOnSecondaries field if non-nil, zero value otherwise.

### GetPreserveHostRrsetOrderOnSecondariesOk

`func (o *GridDns) GetPreserveHostRrsetOrderOnSecondariesOk() (*bool, bool)`

GetPreserveHostRrsetOrderOnSecondariesOk returns a tuple with the PreserveHostRrsetOrderOnSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveHostRrsetOrderOnSecondaries

`func (o *GridDns) SetPreserveHostRrsetOrderOnSecondaries(v bool)`

SetPreserveHostRrsetOrderOnSecondaries sets PreserveHostRrsetOrderOnSecondaries field to given value.

### HasPreserveHostRrsetOrderOnSecondaries

`func (o *GridDns) HasPreserveHostRrsetOrderOnSecondaries() bool`

HasPreserveHostRrsetOrderOnSecondaries returns a boolean if a field has been set.

### GetProtocolRecordNamePolicies

`func (o *GridDns) GetProtocolRecordNamePolicies() []string`

GetProtocolRecordNamePolicies returns the ProtocolRecordNamePolicies field if non-nil, zero value otherwise.

### GetProtocolRecordNamePoliciesOk

`func (o *GridDns) GetProtocolRecordNamePoliciesOk() (*[]string, bool)`

GetProtocolRecordNamePoliciesOk returns a tuple with the ProtocolRecordNamePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRecordNamePolicies

`func (o *GridDns) SetProtocolRecordNamePolicies(v []string)`

SetProtocolRecordNamePolicies sets ProtocolRecordNamePolicies field to given value.

### HasProtocolRecordNamePolicies

`func (o *GridDns) HasProtocolRecordNamePolicies() bool`

HasProtocolRecordNamePolicies returns a boolean if a field has been set.

### GetQueryRewriteDomainNames

`func (o *GridDns) GetQueryRewriteDomainNames() []string`

GetQueryRewriteDomainNames returns the QueryRewriteDomainNames field if non-nil, zero value otherwise.

### GetQueryRewriteDomainNamesOk

`func (o *GridDns) GetQueryRewriteDomainNamesOk() (*[]string, bool)`

GetQueryRewriteDomainNamesOk returns a tuple with the QueryRewriteDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRewriteDomainNames

`func (o *GridDns) SetQueryRewriteDomainNames(v []string)`

SetQueryRewriteDomainNames sets QueryRewriteDomainNames field to given value.

### HasQueryRewriteDomainNames

`func (o *GridDns) HasQueryRewriteDomainNames() bool`

HasQueryRewriteDomainNames returns a boolean if a field has been set.

### GetQueryRewritePrefix

`func (o *GridDns) GetQueryRewritePrefix() string`

GetQueryRewritePrefix returns the QueryRewritePrefix field if non-nil, zero value otherwise.

### GetQueryRewritePrefixOk

`func (o *GridDns) GetQueryRewritePrefixOk() (*string, bool)`

GetQueryRewritePrefixOk returns a tuple with the QueryRewritePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRewritePrefix

`func (o *GridDns) SetQueryRewritePrefix(v string)`

SetQueryRewritePrefix sets QueryRewritePrefix field to given value.

### HasQueryRewritePrefix

`func (o *GridDns) HasQueryRewritePrefix() bool`

HasQueryRewritePrefix returns a boolean if a field has been set.

### GetQuerySourcePort

`func (o *GridDns) GetQuerySourcePort() int64`

GetQuerySourcePort returns the QuerySourcePort field if non-nil, zero value otherwise.

### GetQuerySourcePortOk

`func (o *GridDns) GetQuerySourcePortOk() (*int64, bool)`

GetQuerySourcePortOk returns a tuple with the QuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySourcePort

`func (o *GridDns) SetQuerySourcePort(v int64)`

SetQuerySourcePort sets QuerySourcePort field to given value.

### HasQuerySourcePort

`func (o *GridDns) HasQuerySourcePort() bool`

HasQuerySourcePort returns a boolean if a field has been set.

### GetRecursiveQueryList

`func (o *GridDns) GetRecursiveQueryList() []GridDnsRecursiveQueryList`

GetRecursiveQueryList returns the RecursiveQueryList field if non-nil, zero value otherwise.

### GetRecursiveQueryListOk

`func (o *GridDns) GetRecursiveQueryListOk() (*[]GridDnsRecursiveQueryList, bool)`

GetRecursiveQueryListOk returns a tuple with the RecursiveQueryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveQueryList

`func (o *GridDns) SetRecursiveQueryList(v []GridDnsRecursiveQueryList)`

SetRecursiveQueryList sets RecursiveQueryList field to given value.

### HasRecursiveQueryList

`func (o *GridDns) HasRecursiveQueryList() bool`

HasRecursiveQueryList returns a boolean if a field has been set.

### GetRefreshTimer

`func (o *GridDns) GetRefreshTimer() int64`

GetRefreshTimer returns the RefreshTimer field if non-nil, zero value otherwise.

### GetRefreshTimerOk

`func (o *GridDns) GetRefreshTimerOk() (*int64, bool)`

GetRefreshTimerOk returns a tuple with the RefreshTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTimer

`func (o *GridDns) SetRefreshTimer(v int64)`

SetRefreshTimer sets RefreshTimer field to given value.

### HasRefreshTimer

`func (o *GridDns) HasRefreshTimer() bool`

HasRefreshTimer returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *GridDns) GetResolverQueryTimeout() int64`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *GridDns) GetResolverQueryTimeoutOk() (*int64, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *GridDns) SetResolverQueryTimeout(v int64)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *GridDns) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetResponseRateLimiting

`func (o *GridDns) GetResponseRateLimiting() GridDnsResponseRateLimiting`

GetResponseRateLimiting returns the ResponseRateLimiting field if non-nil, zero value otherwise.

### GetResponseRateLimitingOk

`func (o *GridDns) GetResponseRateLimitingOk() (*GridDnsResponseRateLimiting, bool)`

GetResponseRateLimitingOk returns a tuple with the ResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRateLimiting

`func (o *GridDns) SetResponseRateLimiting(v GridDnsResponseRateLimiting)`

SetResponseRateLimiting sets ResponseRateLimiting field to given value.

### HasResponseRateLimiting

`func (o *GridDns) HasResponseRateLimiting() bool`

HasResponseRateLimiting returns a boolean if a field has been set.

### GetRestartSetting

`func (o *GridDns) GetRestartSetting() GridDnsRestartSetting`

GetRestartSetting returns the RestartSetting field if non-nil, zero value otherwise.

### GetRestartSettingOk

`func (o *GridDns) GetRestartSettingOk() (*GridDnsRestartSetting, bool)`

GetRestartSettingOk returns a tuple with the RestartSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSetting

`func (o *GridDns) SetRestartSetting(v GridDnsRestartSetting)`

SetRestartSetting sets RestartSetting field to given value.

### HasRestartSetting

`func (o *GridDns) HasRestartSetting() bool`

HasRestartSetting returns a boolean if a field has been set.

### GetRetryTimer

`func (o *GridDns) GetRetryTimer() int64`

GetRetryTimer returns the RetryTimer field if non-nil, zero value otherwise.

### GetRetryTimerOk

`func (o *GridDns) GetRetryTimerOk() (*int64, bool)`

GetRetryTimerOk returns a tuple with the RetryTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimer

`func (o *GridDns) SetRetryTimer(v int64)`

SetRetryTimer sets RetryTimer field to given value.

### HasRetryTimer

`func (o *GridDns) HasRetryTimer() bool`

HasRetryTimer returns a boolean if a field has been set.

### GetRootNameServerType

`func (o *GridDns) GetRootNameServerType() string`

GetRootNameServerType returns the RootNameServerType field if non-nil, zero value otherwise.

### GetRootNameServerTypeOk

`func (o *GridDns) GetRootNameServerTypeOk() (*string, bool)`

GetRootNameServerTypeOk returns a tuple with the RootNameServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNameServerType

`func (o *GridDns) SetRootNameServerType(v string)`

SetRootNameServerType sets RootNameServerType field to given value.

### HasRootNameServerType

`func (o *GridDns) HasRootNameServerType() bool`

HasRootNameServerType returns a boolean if a field has been set.

### GetRpzDisableNsdnameNsip

`func (o *GridDns) GetRpzDisableNsdnameNsip() bool`

GetRpzDisableNsdnameNsip returns the RpzDisableNsdnameNsip field if non-nil, zero value otherwise.

### GetRpzDisableNsdnameNsipOk

`func (o *GridDns) GetRpzDisableNsdnameNsipOk() (*bool, bool)`

GetRpzDisableNsdnameNsipOk returns a tuple with the RpzDisableNsdnameNsip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDisableNsdnameNsip

`func (o *GridDns) SetRpzDisableNsdnameNsip(v bool)`

SetRpzDisableNsdnameNsip sets RpzDisableNsdnameNsip field to given value.

### HasRpzDisableNsdnameNsip

`func (o *GridDns) HasRpzDisableNsdnameNsip() bool`

HasRpzDisableNsdnameNsip returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *GridDns) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *GridDns) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *GridDns) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *GridDns) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GridDns) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *GridDns) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GridDns) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GridDns) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GridDns) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *GridDns) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GridDns) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GridDns) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzQnameWaitRecurse

`func (o *GridDns) GetRpzQnameWaitRecurse() bool`

GetRpzQnameWaitRecurse returns the RpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetRpzQnameWaitRecurseOk

`func (o *GridDns) GetRpzQnameWaitRecurseOk() (*bool, bool)`

GetRpzQnameWaitRecurseOk returns a tuple with the RpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzQnameWaitRecurse

`func (o *GridDns) SetRpzQnameWaitRecurse(v bool)`

SetRpzQnameWaitRecurse sets RpzQnameWaitRecurse field to given value.

### HasRpzQnameWaitRecurse

`func (o *GridDns) HasRpzQnameWaitRecurse() bool`

HasRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetRunScavenging

`func (o *GridDns) GetRunScavenging() map[string]interface{}`

GetRunScavenging returns the RunScavenging field if non-nil, zero value otherwise.

### GetRunScavengingOk

`func (o *GridDns) GetRunScavengingOk() (*map[string]interface{}, bool)`

GetRunScavengingOk returns a tuple with the RunScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunScavenging

`func (o *GridDns) SetRunScavenging(v map[string]interface{})`

SetRunScavenging sets RunScavenging field to given value.

### HasRunScavenging

`func (o *GridDns) HasRunScavenging() bool`

HasRunScavenging returns a boolean if a field has been set.

### GetScavengingSettings

`func (o *GridDns) GetScavengingSettings() GridDnsScavengingSettings`

GetScavengingSettings returns the ScavengingSettings field if non-nil, zero value otherwise.

### GetScavengingSettingsOk

`func (o *GridDns) GetScavengingSettingsOk() (*GridDnsScavengingSettings, bool)`

GetScavengingSettingsOk returns a tuple with the ScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSettings

`func (o *GridDns) SetScavengingSettings(v GridDnsScavengingSettings)`

SetScavengingSettings sets ScavengingSettings field to given value.

### HasScavengingSettings

`func (o *GridDns) HasScavengingSettings() bool`

HasScavengingSettings returns a boolean if a field has been set.

### GetSerialQueryRate

`func (o *GridDns) GetSerialQueryRate() int64`

GetSerialQueryRate returns the SerialQueryRate field if non-nil, zero value otherwise.

### GetSerialQueryRateOk

`func (o *GridDns) GetSerialQueryRateOk() (*int64, bool)`

GetSerialQueryRateOk returns a tuple with the SerialQueryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialQueryRate

`func (o *GridDns) SetSerialQueryRate(v int64)`

SetSerialQueryRate sets SerialQueryRate field to given value.

### HasSerialQueryRate

`func (o *GridDns) HasSerialQueryRate() bool`

HasSerialQueryRate returns a boolean if a field has been set.

### GetServerIdDirective

`func (o *GridDns) GetServerIdDirective() string`

GetServerIdDirective returns the ServerIdDirective field if non-nil, zero value otherwise.

### GetServerIdDirectiveOk

`func (o *GridDns) GetServerIdDirectiveOk() (*string, bool)`

GetServerIdDirectiveOk returns a tuple with the ServerIdDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIdDirective

`func (o *GridDns) SetServerIdDirective(v string)`

SetServerIdDirective sets ServerIdDirective field to given value.

### HasServerIdDirective

`func (o *GridDns) HasServerIdDirective() bool`

HasServerIdDirective returns a boolean if a field has been set.

### GetSortlist

`func (o *GridDns) GetSortlist() []GridDnsSortlist`

GetSortlist returns the Sortlist field if non-nil, zero value otherwise.

### GetSortlistOk

`func (o *GridDns) GetSortlistOk() (*[]GridDnsSortlist, bool)`

GetSortlistOk returns a tuple with the Sortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortlist

`func (o *GridDns) SetSortlist(v []GridDnsSortlist)`

SetSortlist sets Sortlist field to given value.

### HasSortlist

`func (o *GridDns) HasSortlist() bool`

HasSortlist returns a boolean if a field has been set.

### GetStoreLocally

`func (o *GridDns) GetStoreLocally() bool`

GetStoreLocally returns the StoreLocally field if non-nil, zero value otherwise.

### GetStoreLocallyOk

`func (o *GridDns) GetStoreLocallyOk() (*bool, bool)`

GetStoreLocallyOk returns a tuple with the StoreLocally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreLocally

`func (o *GridDns) SetStoreLocally(v bool)`

SetStoreLocally sets StoreLocally field to given value.

### HasStoreLocally

`func (o *GridDns) HasStoreLocally() bool`

HasStoreLocally returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *GridDns) GetSyslogFacility() string`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *GridDns) GetSyslogFacilityOk() (*string, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *GridDns) SetSyslogFacility(v string)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *GridDns) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetTransferExcludedServers

`func (o *GridDns) GetTransferExcludedServers() []string`

GetTransferExcludedServers returns the TransferExcludedServers field if non-nil, zero value otherwise.

### GetTransferExcludedServersOk

`func (o *GridDns) GetTransferExcludedServersOk() (*[]string, bool)`

GetTransferExcludedServersOk returns a tuple with the TransferExcludedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferExcludedServers

`func (o *GridDns) SetTransferExcludedServers(v []string)`

SetTransferExcludedServers sets TransferExcludedServers field to given value.

### HasTransferExcludedServers

`func (o *GridDns) HasTransferExcludedServers() bool`

HasTransferExcludedServers returns a boolean if a field has been set.

### GetTransferFormat

`func (o *GridDns) GetTransferFormat() string`

GetTransferFormat returns the TransferFormat field if non-nil, zero value otherwise.

### GetTransferFormatOk

`func (o *GridDns) GetTransferFormatOk() (*string, bool)`

GetTransferFormatOk returns a tuple with the TransferFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFormat

`func (o *GridDns) SetTransferFormat(v string)`

SetTransferFormat sets TransferFormat field to given value.

### HasTransferFormat

`func (o *GridDns) HasTransferFormat() bool`

HasTransferFormat returns a boolean if a field has been set.

### GetTransfersIn

`func (o *GridDns) GetTransfersIn() int64`

GetTransfersIn returns the TransfersIn field if non-nil, zero value otherwise.

### GetTransfersInOk

`func (o *GridDns) GetTransfersInOk() (*int64, bool)`

GetTransfersInOk returns a tuple with the TransfersIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersIn

`func (o *GridDns) SetTransfersIn(v int64)`

SetTransfersIn sets TransfersIn field to given value.

### HasTransfersIn

`func (o *GridDns) HasTransfersIn() bool`

HasTransfersIn returns a boolean if a field has been set.

### GetTransfersOut

`func (o *GridDns) GetTransfersOut() int64`

GetTransfersOut returns the TransfersOut field if non-nil, zero value otherwise.

### GetTransfersOutOk

`func (o *GridDns) GetTransfersOutOk() (*int64, bool)`

GetTransfersOutOk returns a tuple with the TransfersOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersOut

`func (o *GridDns) SetTransfersOut(v int64)`

SetTransfersOut sets TransfersOut field to given value.

### HasTransfersOut

`func (o *GridDns) HasTransfersOut() bool`

HasTransfersOut returns a boolean if a field has been set.

### GetTransfersPerNs

`func (o *GridDns) GetTransfersPerNs() int64`

GetTransfersPerNs returns the TransfersPerNs field if non-nil, zero value otherwise.

### GetTransfersPerNsOk

`func (o *GridDns) GetTransfersPerNsOk() (*int64, bool)`

GetTransfersPerNsOk returns a tuple with the TransfersPerNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersPerNs

`func (o *GridDns) SetTransfersPerNs(v int64)`

SetTransfersPerNs sets TransfersPerNs field to given value.

### HasTransfersPerNs

`func (o *GridDns) HasTransfersPerNs() bool`

HasTransfersPerNs returns a boolean if a field has been set.

### GetZoneDeletionDoubleConfirm

`func (o *GridDns) GetZoneDeletionDoubleConfirm() bool`

GetZoneDeletionDoubleConfirm returns the ZoneDeletionDoubleConfirm field if non-nil, zero value otherwise.

### GetZoneDeletionDoubleConfirmOk

`func (o *GridDns) GetZoneDeletionDoubleConfirmOk() (*bool, bool)`

GetZoneDeletionDoubleConfirmOk returns a tuple with the ZoneDeletionDoubleConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDeletionDoubleConfirm

`func (o *GridDns) SetZoneDeletionDoubleConfirm(v bool)`

SetZoneDeletionDoubleConfirm sets ZoneDeletionDoubleConfirm field to given value.

### HasZoneDeletionDoubleConfirm

`func (o *GridDns) HasZoneDeletionDoubleConfirm() bool`

HasZoneDeletionDoubleConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


