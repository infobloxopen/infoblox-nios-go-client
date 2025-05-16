# GetGridDnsResponse

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
**Result** | Pointer to [**GridDns**](GridDns.md) |  | [optional] 

## Methods

### NewGetGridDnsResponse

`func NewGetGridDnsResponse() *GetGridDnsResponse`

NewGetGridDnsResponse instantiates a new GetGridDnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridDnsResponseWithDefaults

`func NewGetGridDnsResponseWithDefaults() *GetGridDnsResponse`

NewGetGridDnsResponseWithDefaults instantiates a new GetGridDnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridDnsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridDnsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridDnsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridDnsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddClientIpMacOptions

`func (o *GetGridDnsResponse) GetAddClientIpMacOptions() bool`

GetAddClientIpMacOptions returns the AddClientIpMacOptions field if non-nil, zero value otherwise.

### GetAddClientIpMacOptionsOk

`func (o *GetGridDnsResponse) GetAddClientIpMacOptionsOk() (*bool, bool)`

GetAddClientIpMacOptionsOk returns a tuple with the AddClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddClientIpMacOptions

`func (o *GetGridDnsResponse) SetAddClientIpMacOptions(v bool)`

SetAddClientIpMacOptions sets AddClientIpMacOptions field to given value.

### HasAddClientIpMacOptions

`func (o *GetGridDnsResponse) HasAddClientIpMacOptions() bool`

HasAddClientIpMacOptions returns a boolean if a field has been set.

### GetAllowBulkhostDdns

`func (o *GetGridDnsResponse) GetAllowBulkhostDdns() string`

GetAllowBulkhostDdns returns the AllowBulkhostDdns field if non-nil, zero value otherwise.

### GetAllowBulkhostDdnsOk

`func (o *GetGridDnsResponse) GetAllowBulkhostDdnsOk() (*string, bool)`

GetAllowBulkhostDdnsOk returns a tuple with the AllowBulkhostDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBulkhostDdns

`func (o *GetGridDnsResponse) SetAllowBulkhostDdns(v string)`

SetAllowBulkhostDdns sets AllowBulkhostDdns field to given value.

### HasAllowBulkhostDdns

`func (o *GetGridDnsResponse) HasAllowBulkhostDdns() bool`

HasAllowBulkhostDdns returns a boolean if a field has been set.

### GetAllowGssTsigZoneUpdates

`func (o *GetGridDnsResponse) GetAllowGssTsigZoneUpdates() bool`

GetAllowGssTsigZoneUpdates returns the AllowGssTsigZoneUpdates field if non-nil, zero value otherwise.

### GetAllowGssTsigZoneUpdatesOk

`func (o *GetGridDnsResponse) GetAllowGssTsigZoneUpdatesOk() (*bool, bool)`

GetAllowGssTsigZoneUpdatesOk returns a tuple with the AllowGssTsigZoneUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGssTsigZoneUpdates

`func (o *GetGridDnsResponse) SetAllowGssTsigZoneUpdates(v bool)`

SetAllowGssTsigZoneUpdates sets AllowGssTsigZoneUpdates field to given value.

### HasAllowGssTsigZoneUpdates

`func (o *GetGridDnsResponse) HasAllowGssTsigZoneUpdates() bool`

HasAllowGssTsigZoneUpdates returns a boolean if a field has been set.

### GetAllowQuery

`func (o *GetGridDnsResponse) GetAllowQuery() []GridDnsAllowQuery`

GetAllowQuery returns the AllowQuery field if non-nil, zero value otherwise.

### GetAllowQueryOk

`func (o *GetGridDnsResponse) GetAllowQueryOk() (*[]GridDnsAllowQuery, bool)`

GetAllowQueryOk returns a tuple with the AllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuery

`func (o *GetGridDnsResponse) SetAllowQuery(v []GridDnsAllowQuery)`

SetAllowQuery sets AllowQuery field to given value.

### HasAllowQuery

`func (o *GetGridDnsResponse) HasAllowQuery() bool`

HasAllowQuery returns a boolean if a field has been set.

### GetAllowRecursiveQuery

`func (o *GetGridDnsResponse) GetAllowRecursiveQuery() bool`

GetAllowRecursiveQuery returns the AllowRecursiveQuery field if non-nil, zero value otherwise.

### GetAllowRecursiveQueryOk

`func (o *GetGridDnsResponse) GetAllowRecursiveQueryOk() (*bool, bool)`

GetAllowRecursiveQueryOk returns a tuple with the AllowRecursiveQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRecursiveQuery

`func (o *GetGridDnsResponse) SetAllowRecursiveQuery(v bool)`

SetAllowRecursiveQuery sets AllowRecursiveQuery field to given value.

### HasAllowRecursiveQuery

`func (o *GetGridDnsResponse) HasAllowRecursiveQuery() bool`

HasAllowRecursiveQuery returns a boolean if a field has been set.

### GetAllowTransfer

`func (o *GetGridDnsResponse) GetAllowTransfer() []GridDnsAllowTransfer`

GetAllowTransfer returns the AllowTransfer field if non-nil, zero value otherwise.

### GetAllowTransferOk

`func (o *GetGridDnsResponse) GetAllowTransferOk() (*[]GridDnsAllowTransfer, bool)`

GetAllowTransferOk returns a tuple with the AllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransfer

`func (o *GetGridDnsResponse) SetAllowTransfer(v []GridDnsAllowTransfer)`

SetAllowTransfer sets AllowTransfer field to given value.

### HasAllowTransfer

`func (o *GetGridDnsResponse) HasAllowTransfer() bool`

HasAllowTransfer returns a boolean if a field has been set.

### GetAllowUpdate

`func (o *GetGridDnsResponse) GetAllowUpdate() []GridDnsAllowUpdate`

GetAllowUpdate returns the AllowUpdate field if non-nil, zero value otherwise.

### GetAllowUpdateOk

`func (o *GetGridDnsResponse) GetAllowUpdateOk() (*[]GridDnsAllowUpdate, bool)`

GetAllowUpdateOk returns a tuple with the AllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdate

`func (o *GetGridDnsResponse) SetAllowUpdate(v []GridDnsAllowUpdate)`

SetAllowUpdate sets AllowUpdate field to given value.

### HasAllowUpdate

`func (o *GetGridDnsResponse) HasAllowUpdate() bool`

HasAllowUpdate returns a boolean if a field has been set.

### GetAnonymizeResponseLogging

`func (o *GetGridDnsResponse) GetAnonymizeResponseLogging() bool`

GetAnonymizeResponseLogging returns the AnonymizeResponseLogging field if non-nil, zero value otherwise.

### GetAnonymizeResponseLoggingOk

`func (o *GetGridDnsResponse) GetAnonymizeResponseLoggingOk() (*bool, bool)`

GetAnonymizeResponseLoggingOk returns a tuple with the AnonymizeResponseLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeResponseLogging

`func (o *GetGridDnsResponse) SetAnonymizeResponseLogging(v bool)`

SetAnonymizeResponseLogging sets AnonymizeResponseLogging field to given value.

### HasAnonymizeResponseLogging

`func (o *GetGridDnsResponse) HasAnonymizeResponseLogging() bool`

HasAnonymizeResponseLogging returns a boolean if a field has been set.

### GetAttackMitigation

`func (o *GetGridDnsResponse) GetAttackMitigation() GridDnsAttackMitigation`

GetAttackMitigation returns the AttackMitigation field if non-nil, zero value otherwise.

### GetAttackMitigationOk

`func (o *GetGridDnsResponse) GetAttackMitigationOk() (*GridDnsAttackMitigation, bool)`

GetAttackMitigationOk returns a tuple with the AttackMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackMitigation

`func (o *GetGridDnsResponse) SetAttackMitigation(v GridDnsAttackMitigation)`

SetAttackMitigation sets AttackMitigation field to given value.

### HasAttackMitigation

`func (o *GetGridDnsResponse) HasAttackMitigation() bool`

HasAttackMitigation returns a boolean if a field has been set.

### GetAutoBlackhole

`func (o *GetGridDnsResponse) GetAutoBlackhole() GridDnsAutoBlackhole`

GetAutoBlackhole returns the AutoBlackhole field if non-nil, zero value otherwise.

### GetAutoBlackholeOk

`func (o *GetGridDnsResponse) GetAutoBlackholeOk() (*GridDnsAutoBlackhole, bool)`

GetAutoBlackholeOk returns a tuple with the AutoBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlackhole

`func (o *GetGridDnsResponse) SetAutoBlackhole(v GridDnsAutoBlackhole)`

SetAutoBlackhole sets AutoBlackhole field to given value.

### HasAutoBlackhole

`func (o *GetGridDnsResponse) HasAutoBlackhole() bool`

HasAutoBlackhole returns a boolean if a field has been set.

### GetBindCheckNamesPolicy

`func (o *GetGridDnsResponse) GetBindCheckNamesPolicy() string`

GetBindCheckNamesPolicy returns the BindCheckNamesPolicy field if non-nil, zero value otherwise.

### GetBindCheckNamesPolicyOk

`func (o *GetGridDnsResponse) GetBindCheckNamesPolicyOk() (*string, bool)`

GetBindCheckNamesPolicyOk returns a tuple with the BindCheckNamesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCheckNamesPolicy

`func (o *GetGridDnsResponse) SetBindCheckNamesPolicy(v string)`

SetBindCheckNamesPolicy sets BindCheckNamesPolicy field to given value.

### HasBindCheckNamesPolicy

`func (o *GetGridDnsResponse) HasBindCheckNamesPolicy() bool`

HasBindCheckNamesPolicy returns a boolean if a field has been set.

### GetBindHostnameDirective

`func (o *GetGridDnsResponse) GetBindHostnameDirective() string`

GetBindHostnameDirective returns the BindHostnameDirective field if non-nil, zero value otherwise.

### GetBindHostnameDirectiveOk

`func (o *GetGridDnsResponse) GetBindHostnameDirectiveOk() (*string, bool)`

GetBindHostnameDirectiveOk returns a tuple with the BindHostnameDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindHostnameDirective

`func (o *GetGridDnsResponse) SetBindHostnameDirective(v string)`

SetBindHostnameDirective sets BindHostnameDirective field to given value.

### HasBindHostnameDirective

`func (o *GetGridDnsResponse) HasBindHostnameDirective() bool`

HasBindHostnameDirective returns a boolean if a field has been set.

### GetBlackholeList

`func (o *GetGridDnsResponse) GetBlackholeList() []GridDnsBlackholeList`

GetBlackholeList returns the BlackholeList field if non-nil, zero value otherwise.

### GetBlackholeListOk

`func (o *GetGridDnsResponse) GetBlackholeListOk() (*[]GridDnsBlackholeList, bool)`

GetBlackholeListOk returns a tuple with the BlackholeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackholeList

`func (o *GetGridDnsResponse) SetBlackholeList(v []GridDnsBlackholeList)`

SetBlackholeList sets BlackholeList field to given value.

### HasBlackholeList

`func (o *GetGridDnsResponse) HasBlackholeList() bool`

HasBlackholeList returns a boolean if a field has been set.

### GetBlacklistAction

`func (o *GetGridDnsResponse) GetBlacklistAction() string`

GetBlacklistAction returns the BlacklistAction field if non-nil, zero value otherwise.

### GetBlacklistActionOk

`func (o *GetGridDnsResponse) GetBlacklistActionOk() (*string, bool)`

GetBlacklistActionOk returns a tuple with the BlacklistAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistAction

`func (o *GetGridDnsResponse) SetBlacklistAction(v string)`

SetBlacklistAction sets BlacklistAction field to given value.

### HasBlacklistAction

`func (o *GetGridDnsResponse) HasBlacklistAction() bool`

HasBlacklistAction returns a boolean if a field has been set.

### GetBlacklistLogQuery

`func (o *GetGridDnsResponse) GetBlacklistLogQuery() bool`

GetBlacklistLogQuery returns the BlacklistLogQuery field if non-nil, zero value otherwise.

### GetBlacklistLogQueryOk

`func (o *GetGridDnsResponse) GetBlacklistLogQueryOk() (*bool, bool)`

GetBlacklistLogQueryOk returns a tuple with the BlacklistLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistLogQuery

`func (o *GetGridDnsResponse) SetBlacklistLogQuery(v bool)`

SetBlacklistLogQuery sets BlacklistLogQuery field to given value.

### HasBlacklistLogQuery

`func (o *GetGridDnsResponse) HasBlacklistLogQuery() bool`

HasBlacklistLogQuery returns a boolean if a field has been set.

### GetBlacklistRedirectAddresses

`func (o *GetGridDnsResponse) GetBlacklistRedirectAddresses() []string`

GetBlacklistRedirectAddresses returns the BlacklistRedirectAddresses field if non-nil, zero value otherwise.

### GetBlacklistRedirectAddressesOk

`func (o *GetGridDnsResponse) GetBlacklistRedirectAddressesOk() (*[]string, bool)`

GetBlacklistRedirectAddressesOk returns a tuple with the BlacklistRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectAddresses

`func (o *GetGridDnsResponse) SetBlacklistRedirectAddresses(v []string)`

SetBlacklistRedirectAddresses sets BlacklistRedirectAddresses field to given value.

### HasBlacklistRedirectAddresses

`func (o *GetGridDnsResponse) HasBlacklistRedirectAddresses() bool`

HasBlacklistRedirectAddresses returns a boolean if a field has been set.

### GetBlacklistRedirectTtl

`func (o *GetGridDnsResponse) GetBlacklistRedirectTtl() int64`

GetBlacklistRedirectTtl returns the BlacklistRedirectTtl field if non-nil, zero value otherwise.

### GetBlacklistRedirectTtlOk

`func (o *GetGridDnsResponse) GetBlacklistRedirectTtlOk() (*int64, bool)`

GetBlacklistRedirectTtlOk returns a tuple with the BlacklistRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRedirectTtl

`func (o *GetGridDnsResponse) SetBlacklistRedirectTtl(v int64)`

SetBlacklistRedirectTtl sets BlacklistRedirectTtl field to given value.

### HasBlacklistRedirectTtl

`func (o *GetGridDnsResponse) HasBlacklistRedirectTtl() bool`

HasBlacklistRedirectTtl returns a boolean if a field has been set.

### GetBlacklistRulesets

`func (o *GetGridDnsResponse) GetBlacklistRulesets() []string`

GetBlacklistRulesets returns the BlacklistRulesets field if non-nil, zero value otherwise.

### GetBlacklistRulesetsOk

`func (o *GetGridDnsResponse) GetBlacklistRulesetsOk() (*[]string, bool)`

GetBlacklistRulesetsOk returns a tuple with the BlacklistRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRulesets

`func (o *GetGridDnsResponse) SetBlacklistRulesets(v []string)`

SetBlacklistRulesets sets BlacklistRulesets field to given value.

### HasBlacklistRulesets

`func (o *GetGridDnsResponse) HasBlacklistRulesets() bool`

HasBlacklistRulesets returns a boolean if a field has been set.

### GetBulkHostNameTemplates

`func (o *GetGridDnsResponse) GetBulkHostNameTemplates() []string`

GetBulkHostNameTemplates returns the BulkHostNameTemplates field if non-nil, zero value otherwise.

### GetBulkHostNameTemplatesOk

`func (o *GetGridDnsResponse) GetBulkHostNameTemplatesOk() (*[]string, bool)`

GetBulkHostNameTemplatesOk returns a tuple with the BulkHostNameTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkHostNameTemplates

`func (o *GetGridDnsResponse) SetBulkHostNameTemplates(v []string)`

SetBulkHostNameTemplates sets BulkHostNameTemplates field to given value.

### HasBulkHostNameTemplates

`func (o *GetGridDnsResponse) HasBulkHostNameTemplates() bool`

HasBulkHostNameTemplates returns a boolean if a field has been set.

### GetCaptureDnsQueriesOnAllDomains

`func (o *GetGridDnsResponse) GetCaptureDnsQueriesOnAllDomains() bool`

GetCaptureDnsQueriesOnAllDomains returns the CaptureDnsQueriesOnAllDomains field if non-nil, zero value otherwise.

### GetCaptureDnsQueriesOnAllDomainsOk

`func (o *GetGridDnsResponse) GetCaptureDnsQueriesOnAllDomainsOk() (*bool, bool)`

GetCaptureDnsQueriesOnAllDomainsOk returns a tuple with the CaptureDnsQueriesOnAllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDnsQueriesOnAllDomains

`func (o *GetGridDnsResponse) SetCaptureDnsQueriesOnAllDomains(v bool)`

SetCaptureDnsQueriesOnAllDomains sets CaptureDnsQueriesOnAllDomains field to given value.

### HasCaptureDnsQueriesOnAllDomains

`func (o *GetGridDnsResponse) HasCaptureDnsQueriesOnAllDomains() bool`

HasCaptureDnsQueriesOnAllDomains returns a boolean if a field has been set.

### GetCheckNamesForDdnsAndZoneTransfer

`func (o *GetGridDnsResponse) GetCheckNamesForDdnsAndZoneTransfer() bool`

GetCheckNamesForDdnsAndZoneTransfer returns the CheckNamesForDdnsAndZoneTransfer field if non-nil, zero value otherwise.

### GetCheckNamesForDdnsAndZoneTransferOk

`func (o *GetGridDnsResponse) GetCheckNamesForDdnsAndZoneTransferOk() (*bool, bool)`

GetCheckNamesForDdnsAndZoneTransferOk returns a tuple with the CheckNamesForDdnsAndZoneTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesForDdnsAndZoneTransfer

`func (o *GetGridDnsResponse) SetCheckNamesForDdnsAndZoneTransfer(v bool)`

SetCheckNamesForDdnsAndZoneTransfer sets CheckNamesForDdnsAndZoneTransfer field to given value.

### HasCheckNamesForDdnsAndZoneTransfer

`func (o *GetGridDnsResponse) HasCheckNamesForDdnsAndZoneTransfer() bool`

HasCheckNamesForDdnsAndZoneTransfer returns a boolean if a field has been set.

### GetClientSubnetDomains

`func (o *GetGridDnsResponse) GetClientSubnetDomains() []GridDnsClientSubnetDomains`

GetClientSubnetDomains returns the ClientSubnetDomains field if non-nil, zero value otherwise.

### GetClientSubnetDomainsOk

`func (o *GetGridDnsResponse) GetClientSubnetDomainsOk() (*[]GridDnsClientSubnetDomains, bool)`

GetClientSubnetDomainsOk returns a tuple with the ClientSubnetDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetDomains

`func (o *GetGridDnsResponse) SetClientSubnetDomains(v []GridDnsClientSubnetDomains)`

SetClientSubnetDomains sets ClientSubnetDomains field to given value.

### HasClientSubnetDomains

`func (o *GetGridDnsResponse) HasClientSubnetDomains() bool`

HasClientSubnetDomains returns a boolean if a field has been set.

### GetClientSubnetIpv4PrefixLength

`func (o *GetGridDnsResponse) GetClientSubnetIpv4PrefixLength() int64`

GetClientSubnetIpv4PrefixLength returns the ClientSubnetIpv4PrefixLength field if non-nil, zero value otherwise.

### GetClientSubnetIpv4PrefixLengthOk

`func (o *GetGridDnsResponse) GetClientSubnetIpv4PrefixLengthOk() (*int64, bool)`

GetClientSubnetIpv4PrefixLengthOk returns a tuple with the ClientSubnetIpv4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetIpv4PrefixLength

`func (o *GetGridDnsResponse) SetClientSubnetIpv4PrefixLength(v int64)`

SetClientSubnetIpv4PrefixLength sets ClientSubnetIpv4PrefixLength field to given value.

### HasClientSubnetIpv4PrefixLength

`func (o *GetGridDnsResponse) HasClientSubnetIpv4PrefixLength() bool`

HasClientSubnetIpv4PrefixLength returns a boolean if a field has been set.

### GetClientSubnetIpv6PrefixLength

`func (o *GetGridDnsResponse) GetClientSubnetIpv6PrefixLength() int64`

GetClientSubnetIpv6PrefixLength returns the ClientSubnetIpv6PrefixLength field if non-nil, zero value otherwise.

### GetClientSubnetIpv6PrefixLengthOk

`func (o *GetGridDnsResponse) GetClientSubnetIpv6PrefixLengthOk() (*int64, bool)`

GetClientSubnetIpv6PrefixLengthOk returns a tuple with the ClientSubnetIpv6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetIpv6PrefixLength

`func (o *GetGridDnsResponse) SetClientSubnetIpv6PrefixLength(v int64)`

SetClientSubnetIpv6PrefixLength sets ClientSubnetIpv6PrefixLength field to given value.

### HasClientSubnetIpv6PrefixLength

`func (o *GetGridDnsResponse) HasClientSubnetIpv6PrefixLength() bool`

HasClientSubnetIpv6PrefixLength returns a boolean if a field has been set.

### GetCopyClientIpMacOptions

`func (o *GetGridDnsResponse) GetCopyClientIpMacOptions() bool`

GetCopyClientIpMacOptions returns the CopyClientIpMacOptions field if non-nil, zero value otherwise.

### GetCopyClientIpMacOptionsOk

`func (o *GetGridDnsResponse) GetCopyClientIpMacOptionsOk() (*bool, bool)`

GetCopyClientIpMacOptionsOk returns a tuple with the CopyClientIpMacOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyClientIpMacOptions

`func (o *GetGridDnsResponse) SetCopyClientIpMacOptions(v bool)`

SetCopyClientIpMacOptions sets CopyClientIpMacOptions field to given value.

### HasCopyClientIpMacOptions

`func (o *GetGridDnsResponse) HasCopyClientIpMacOptions() bool`

HasCopyClientIpMacOptions returns a boolean if a field has been set.

### GetCopyXferToNotify

`func (o *GetGridDnsResponse) GetCopyXferToNotify() bool`

GetCopyXferToNotify returns the CopyXferToNotify field if non-nil, zero value otherwise.

### GetCopyXferToNotifyOk

`func (o *GetGridDnsResponse) GetCopyXferToNotifyOk() (*bool, bool)`

GetCopyXferToNotifyOk returns a tuple with the CopyXferToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyXferToNotify

`func (o *GetGridDnsResponse) SetCopyXferToNotify(v bool)`

SetCopyXferToNotify sets CopyXferToNotify field to given value.

### HasCopyXferToNotify

`func (o *GetGridDnsResponse) HasCopyXferToNotify() bool`

HasCopyXferToNotify returns a boolean if a field has been set.

### GetCustomRootNameServers

`func (o *GetGridDnsResponse) GetCustomRootNameServers() []GridDnsCustomRootNameServers`

GetCustomRootNameServers returns the CustomRootNameServers field if non-nil, zero value otherwise.

### GetCustomRootNameServersOk

`func (o *GetGridDnsResponse) GetCustomRootNameServersOk() (*[]GridDnsCustomRootNameServers, bool)`

GetCustomRootNameServersOk returns a tuple with the CustomRootNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNameServers

`func (o *GetGridDnsResponse) SetCustomRootNameServers(v []GridDnsCustomRootNameServers)`

SetCustomRootNameServers sets CustomRootNameServers field to given value.

### HasCustomRootNameServers

`func (o *GetGridDnsResponse) HasCustomRootNameServers() bool`

HasCustomRootNameServers returns a boolean if a field has been set.

### GetDdnsForceCreationTimestampUpdate

`func (o *GetGridDnsResponse) GetDdnsForceCreationTimestampUpdate() bool`

GetDdnsForceCreationTimestampUpdate returns the DdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetDdnsForceCreationTimestampUpdateOk

`func (o *GetGridDnsResponse) GetDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetDdnsForceCreationTimestampUpdateOk returns a tuple with the DdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsForceCreationTimestampUpdate

`func (o *GetGridDnsResponse) SetDdnsForceCreationTimestampUpdate(v bool)`

SetDdnsForceCreationTimestampUpdate sets DdnsForceCreationTimestampUpdate field to given value.

### HasDdnsForceCreationTimestampUpdate

`func (o *GetGridDnsResponse) HasDdnsForceCreationTimestampUpdate() bool`

HasDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetDdnsPrincipalGroup

`func (o *GetGridDnsResponse) GetDdnsPrincipalGroup() string`

GetDdnsPrincipalGroup returns the DdnsPrincipalGroup field if non-nil, zero value otherwise.

### GetDdnsPrincipalGroupOk

`func (o *GetGridDnsResponse) GetDdnsPrincipalGroupOk() (*string, bool)`

GetDdnsPrincipalGroupOk returns a tuple with the DdnsPrincipalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalGroup

`func (o *GetGridDnsResponse) SetDdnsPrincipalGroup(v string)`

SetDdnsPrincipalGroup sets DdnsPrincipalGroup field to given value.

### HasDdnsPrincipalGroup

`func (o *GetGridDnsResponse) HasDdnsPrincipalGroup() bool`

HasDdnsPrincipalGroup returns a boolean if a field has been set.

### GetDdnsPrincipalTracking

`func (o *GetGridDnsResponse) GetDdnsPrincipalTracking() bool`

GetDdnsPrincipalTracking returns the DdnsPrincipalTracking field if non-nil, zero value otherwise.

### GetDdnsPrincipalTrackingOk

`func (o *GetGridDnsResponse) GetDdnsPrincipalTrackingOk() (*bool, bool)`

GetDdnsPrincipalTrackingOk returns a tuple with the DdnsPrincipalTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalTracking

`func (o *GetGridDnsResponse) SetDdnsPrincipalTracking(v bool)`

SetDdnsPrincipalTracking sets DdnsPrincipalTracking field to given value.

### HasDdnsPrincipalTracking

`func (o *GetGridDnsResponse) HasDdnsPrincipalTracking() bool`

HasDdnsPrincipalTracking returns a boolean if a field has been set.

### GetDdnsRestrictPatterns

`func (o *GetGridDnsResponse) GetDdnsRestrictPatterns() bool`

GetDdnsRestrictPatterns returns the DdnsRestrictPatterns field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsOk

`func (o *GetGridDnsResponse) GetDdnsRestrictPatternsOk() (*bool, bool)`

GetDdnsRestrictPatternsOk returns a tuple with the DdnsRestrictPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatterns

`func (o *GetGridDnsResponse) SetDdnsRestrictPatterns(v bool)`

SetDdnsRestrictPatterns sets DdnsRestrictPatterns field to given value.

### HasDdnsRestrictPatterns

`func (o *GetGridDnsResponse) HasDdnsRestrictPatterns() bool`

HasDdnsRestrictPatterns returns a boolean if a field has been set.

### GetDdnsRestrictPatternsList

`func (o *GetGridDnsResponse) GetDdnsRestrictPatternsList() []string`

GetDdnsRestrictPatternsList returns the DdnsRestrictPatternsList field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsListOk

`func (o *GetGridDnsResponse) GetDdnsRestrictPatternsListOk() (*[]string, bool)`

GetDdnsRestrictPatternsListOk returns a tuple with the DdnsRestrictPatternsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatternsList

`func (o *GetGridDnsResponse) SetDdnsRestrictPatternsList(v []string)`

SetDdnsRestrictPatternsList sets DdnsRestrictPatternsList field to given value.

### HasDdnsRestrictPatternsList

`func (o *GetGridDnsResponse) HasDdnsRestrictPatternsList() bool`

HasDdnsRestrictPatternsList returns a boolean if a field has been set.

### GetDdnsRestrictProtected

`func (o *GetGridDnsResponse) GetDdnsRestrictProtected() bool`

GetDdnsRestrictProtected returns the DdnsRestrictProtected field if non-nil, zero value otherwise.

### GetDdnsRestrictProtectedOk

`func (o *GetGridDnsResponse) GetDdnsRestrictProtectedOk() (*bool, bool)`

GetDdnsRestrictProtectedOk returns a tuple with the DdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictProtected

`func (o *GetGridDnsResponse) SetDdnsRestrictProtected(v bool)`

SetDdnsRestrictProtected sets DdnsRestrictProtected field to given value.

### HasDdnsRestrictProtected

`func (o *GetGridDnsResponse) HasDdnsRestrictProtected() bool`

HasDdnsRestrictProtected returns a boolean if a field has been set.

### GetDdnsRestrictSecure

`func (o *GetGridDnsResponse) GetDdnsRestrictSecure() bool`

GetDdnsRestrictSecure returns the DdnsRestrictSecure field if non-nil, zero value otherwise.

### GetDdnsRestrictSecureOk

`func (o *GetGridDnsResponse) GetDdnsRestrictSecureOk() (*bool, bool)`

GetDdnsRestrictSecureOk returns a tuple with the DdnsRestrictSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictSecure

`func (o *GetGridDnsResponse) SetDdnsRestrictSecure(v bool)`

SetDdnsRestrictSecure sets DdnsRestrictSecure field to given value.

### HasDdnsRestrictSecure

`func (o *GetGridDnsResponse) HasDdnsRestrictSecure() bool`

HasDdnsRestrictSecure returns a boolean if a field has been set.

### GetDdnsRestrictStatic

`func (o *GetGridDnsResponse) GetDdnsRestrictStatic() bool`

GetDdnsRestrictStatic returns the DdnsRestrictStatic field if non-nil, zero value otherwise.

### GetDdnsRestrictStaticOk

`func (o *GetGridDnsResponse) GetDdnsRestrictStaticOk() (*bool, bool)`

GetDdnsRestrictStaticOk returns a tuple with the DdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictStatic

`func (o *GetGridDnsResponse) SetDdnsRestrictStatic(v bool)`

SetDdnsRestrictStatic sets DdnsRestrictStatic field to given value.

### HasDdnsRestrictStatic

`func (o *GetGridDnsResponse) HasDdnsRestrictStatic() bool`

HasDdnsRestrictStatic returns a boolean if a field has been set.

### GetDefaultBulkHostNameTemplate

`func (o *GetGridDnsResponse) GetDefaultBulkHostNameTemplate() string`

GetDefaultBulkHostNameTemplate returns the DefaultBulkHostNameTemplate field if non-nil, zero value otherwise.

### GetDefaultBulkHostNameTemplateOk

`func (o *GetGridDnsResponse) GetDefaultBulkHostNameTemplateOk() (*string, bool)`

GetDefaultBulkHostNameTemplateOk returns a tuple with the DefaultBulkHostNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBulkHostNameTemplate

`func (o *GetGridDnsResponse) SetDefaultBulkHostNameTemplate(v string)`

SetDefaultBulkHostNameTemplate sets DefaultBulkHostNameTemplate field to given value.

### HasDefaultBulkHostNameTemplate

`func (o *GetGridDnsResponse) HasDefaultBulkHostNameTemplate() bool`

HasDefaultBulkHostNameTemplate returns a boolean if a field has been set.

### GetDefaultTtl

`func (o *GetGridDnsResponse) GetDefaultTtl() int64`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *GetGridDnsResponse) GetDefaultTtlOk() (*int64, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *GetGridDnsResponse) SetDefaultTtl(v int64)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *GetGridDnsResponse) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetDisableEdns

`func (o *GetGridDnsResponse) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *GetGridDnsResponse) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *GetGridDnsResponse) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.

### HasDisableEdns

`func (o *GetGridDnsResponse) HasDisableEdns() bool`

HasDisableEdns returns a boolean if a field has been set.

### GetDns64Groups

`func (o *GetGridDnsResponse) GetDns64Groups() []string`

GetDns64Groups returns the Dns64Groups field if non-nil, zero value otherwise.

### GetDns64GroupsOk

`func (o *GetGridDnsResponse) GetDns64GroupsOk() (*[]string, bool)`

GetDns64GroupsOk returns a tuple with the Dns64Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns64Groups

`func (o *GetGridDnsResponse) SetDns64Groups(v []string)`

SetDns64Groups sets Dns64Groups field to given value.

### HasDns64Groups

`func (o *GetGridDnsResponse) HasDns64Groups() bool`

HasDns64Groups returns a boolean if a field has been set.

### GetDnsCacheAccelerationTtl

`func (o *GetGridDnsResponse) GetDnsCacheAccelerationTtl() int64`

GetDnsCacheAccelerationTtl returns the DnsCacheAccelerationTtl field if non-nil, zero value otherwise.

### GetDnsCacheAccelerationTtlOk

`func (o *GetGridDnsResponse) GetDnsCacheAccelerationTtlOk() (*int64, bool)`

GetDnsCacheAccelerationTtlOk returns a tuple with the DnsCacheAccelerationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCacheAccelerationTtl

`func (o *GetGridDnsResponse) SetDnsCacheAccelerationTtl(v int64)`

SetDnsCacheAccelerationTtl sets DnsCacheAccelerationTtl field to given value.

### HasDnsCacheAccelerationTtl

`func (o *GetGridDnsResponse) HasDnsCacheAccelerationTtl() bool`

HasDnsCacheAccelerationTtl returns a boolean if a field has been set.

### GetDnsHealthCheckAnycastControl

`func (o *GetGridDnsResponse) GetDnsHealthCheckAnycastControl() bool`

GetDnsHealthCheckAnycastControl returns the DnsHealthCheckAnycastControl field if non-nil, zero value otherwise.

### GetDnsHealthCheckAnycastControlOk

`func (o *GetGridDnsResponse) GetDnsHealthCheckAnycastControlOk() (*bool, bool)`

GetDnsHealthCheckAnycastControlOk returns a tuple with the DnsHealthCheckAnycastControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckAnycastControl

`func (o *GetGridDnsResponse) SetDnsHealthCheckAnycastControl(v bool)`

SetDnsHealthCheckAnycastControl sets DnsHealthCheckAnycastControl field to given value.

### HasDnsHealthCheckAnycastControl

`func (o *GetGridDnsResponse) HasDnsHealthCheckAnycastControl() bool`

HasDnsHealthCheckAnycastControl returns a boolean if a field has been set.

### GetDnsHealthCheckDomainList

`func (o *GetGridDnsResponse) GetDnsHealthCheckDomainList() []string`

GetDnsHealthCheckDomainList returns the DnsHealthCheckDomainList field if non-nil, zero value otherwise.

### GetDnsHealthCheckDomainListOk

`func (o *GetGridDnsResponse) GetDnsHealthCheckDomainListOk() (*[]string, bool)`

GetDnsHealthCheckDomainListOk returns a tuple with the DnsHealthCheckDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckDomainList

`func (o *GetGridDnsResponse) SetDnsHealthCheckDomainList(v []string)`

SetDnsHealthCheckDomainList sets DnsHealthCheckDomainList field to given value.

### HasDnsHealthCheckDomainList

`func (o *GetGridDnsResponse) HasDnsHealthCheckDomainList() bool`

HasDnsHealthCheckDomainList returns a boolean if a field has been set.

### GetDnsHealthCheckInterval

`func (o *GetGridDnsResponse) GetDnsHealthCheckInterval() int64`

GetDnsHealthCheckInterval returns the DnsHealthCheckInterval field if non-nil, zero value otherwise.

### GetDnsHealthCheckIntervalOk

`func (o *GetGridDnsResponse) GetDnsHealthCheckIntervalOk() (*int64, bool)`

GetDnsHealthCheckIntervalOk returns a tuple with the DnsHealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckInterval

`func (o *GetGridDnsResponse) SetDnsHealthCheckInterval(v int64)`

SetDnsHealthCheckInterval sets DnsHealthCheckInterval field to given value.

### HasDnsHealthCheckInterval

`func (o *GetGridDnsResponse) HasDnsHealthCheckInterval() bool`

HasDnsHealthCheckInterval returns a boolean if a field has been set.

### GetDnsHealthCheckRecursionFlag

`func (o *GetGridDnsResponse) GetDnsHealthCheckRecursionFlag() bool`

GetDnsHealthCheckRecursionFlag returns the DnsHealthCheckRecursionFlag field if non-nil, zero value otherwise.

### GetDnsHealthCheckRecursionFlagOk

`func (o *GetGridDnsResponse) GetDnsHealthCheckRecursionFlagOk() (*bool, bool)`

GetDnsHealthCheckRecursionFlagOk returns a tuple with the DnsHealthCheckRecursionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckRecursionFlag

`func (o *GetGridDnsResponse) SetDnsHealthCheckRecursionFlag(v bool)`

SetDnsHealthCheckRecursionFlag sets DnsHealthCheckRecursionFlag field to given value.

### HasDnsHealthCheckRecursionFlag

`func (o *GetGridDnsResponse) HasDnsHealthCheckRecursionFlag() bool`

HasDnsHealthCheckRecursionFlag returns a boolean if a field has been set.

### GetDnsHealthCheckRetries

`func (o *GetGridDnsResponse) GetDnsHealthCheckRetries() int64`

GetDnsHealthCheckRetries returns the DnsHealthCheckRetries field if non-nil, zero value otherwise.

### GetDnsHealthCheckRetriesOk

`func (o *GetGridDnsResponse) GetDnsHealthCheckRetriesOk() (*int64, bool)`

GetDnsHealthCheckRetriesOk returns a tuple with the DnsHealthCheckRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckRetries

`func (o *GetGridDnsResponse) SetDnsHealthCheckRetries(v int64)`

SetDnsHealthCheckRetries sets DnsHealthCheckRetries field to given value.

### HasDnsHealthCheckRetries

`func (o *GetGridDnsResponse) HasDnsHealthCheckRetries() bool`

HasDnsHealthCheckRetries returns a boolean if a field has been set.

### GetDnsHealthCheckTimeout

`func (o *GetGridDnsResponse) GetDnsHealthCheckTimeout() int64`

GetDnsHealthCheckTimeout returns the DnsHealthCheckTimeout field if non-nil, zero value otherwise.

### GetDnsHealthCheckTimeoutOk

`func (o *GetGridDnsResponse) GetDnsHealthCheckTimeoutOk() (*int64, bool)`

GetDnsHealthCheckTimeoutOk returns a tuple with the DnsHealthCheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHealthCheckTimeout

`func (o *GetGridDnsResponse) SetDnsHealthCheckTimeout(v int64)`

SetDnsHealthCheckTimeout sets DnsHealthCheckTimeout field to given value.

### HasDnsHealthCheckTimeout

`func (o *GetGridDnsResponse) HasDnsHealthCheckTimeout() bool`

HasDnsHealthCheckTimeout returns a boolean if a field has been set.

### GetDnsQueryCaptureFileTimeLimit

`func (o *GetGridDnsResponse) GetDnsQueryCaptureFileTimeLimit() int64`

GetDnsQueryCaptureFileTimeLimit returns the DnsQueryCaptureFileTimeLimit field if non-nil, zero value otherwise.

### GetDnsQueryCaptureFileTimeLimitOk

`func (o *GetGridDnsResponse) GetDnsQueryCaptureFileTimeLimitOk() (*int64, bool)`

GetDnsQueryCaptureFileTimeLimitOk returns a tuple with the DnsQueryCaptureFileTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueryCaptureFileTimeLimit

`func (o *GetGridDnsResponse) SetDnsQueryCaptureFileTimeLimit(v int64)`

SetDnsQueryCaptureFileTimeLimit sets DnsQueryCaptureFileTimeLimit field to given value.

### HasDnsQueryCaptureFileTimeLimit

`func (o *GetGridDnsResponse) HasDnsQueryCaptureFileTimeLimit() bool`

HasDnsQueryCaptureFileTimeLimit returns a boolean if a field has been set.

### GetDnssecBlacklistEnabled

`func (o *GetGridDnsResponse) GetDnssecBlacklistEnabled() bool`

GetDnssecBlacklistEnabled returns the DnssecBlacklistEnabled field if non-nil, zero value otherwise.

### GetDnssecBlacklistEnabledOk

`func (o *GetGridDnsResponse) GetDnssecBlacklistEnabledOk() (*bool, bool)`

GetDnssecBlacklistEnabledOk returns a tuple with the DnssecBlacklistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecBlacklistEnabled

`func (o *GetGridDnsResponse) SetDnssecBlacklistEnabled(v bool)`

SetDnssecBlacklistEnabled sets DnssecBlacklistEnabled field to given value.

### HasDnssecBlacklistEnabled

`func (o *GetGridDnsResponse) HasDnssecBlacklistEnabled() bool`

HasDnssecBlacklistEnabled returns a boolean if a field has been set.

### GetDnssecDns64Enabled

`func (o *GetGridDnsResponse) GetDnssecDns64Enabled() bool`

GetDnssecDns64Enabled returns the DnssecDns64Enabled field if non-nil, zero value otherwise.

### GetDnssecDns64EnabledOk

`func (o *GetGridDnsResponse) GetDnssecDns64EnabledOk() (*bool, bool)`

GetDnssecDns64EnabledOk returns a tuple with the DnssecDns64Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecDns64Enabled

`func (o *GetGridDnsResponse) SetDnssecDns64Enabled(v bool)`

SetDnssecDns64Enabled sets DnssecDns64Enabled field to given value.

### HasDnssecDns64Enabled

`func (o *GetGridDnsResponse) HasDnssecDns64Enabled() bool`

HasDnssecDns64Enabled returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *GetGridDnsResponse) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *GetGridDnsResponse) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *GetGridDnsResponse) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *GetGridDnsResponse) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecExpiredSignaturesEnabled

`func (o *GetGridDnsResponse) GetDnssecExpiredSignaturesEnabled() bool`

GetDnssecExpiredSignaturesEnabled returns the DnssecExpiredSignaturesEnabled field if non-nil, zero value otherwise.

### GetDnssecExpiredSignaturesEnabledOk

`func (o *GetGridDnsResponse) GetDnssecExpiredSignaturesEnabledOk() (*bool, bool)`

GetDnssecExpiredSignaturesEnabledOk returns a tuple with the DnssecExpiredSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecExpiredSignaturesEnabled

`func (o *GetGridDnsResponse) SetDnssecExpiredSignaturesEnabled(v bool)`

SetDnssecExpiredSignaturesEnabled sets DnssecExpiredSignaturesEnabled field to given value.

### HasDnssecExpiredSignaturesEnabled

`func (o *GetGridDnsResponse) HasDnssecExpiredSignaturesEnabled() bool`

HasDnssecExpiredSignaturesEnabled returns a boolean if a field has been set.

### GetDnssecKeyParams

`func (o *GetGridDnsResponse) GetDnssecKeyParams() GridDnsDnssecKeyParams`

GetDnssecKeyParams returns the DnssecKeyParams field if non-nil, zero value otherwise.

### GetDnssecKeyParamsOk

`func (o *GetGridDnsResponse) GetDnssecKeyParamsOk() (*GridDnsDnssecKeyParams, bool)`

GetDnssecKeyParamsOk returns a tuple with the DnssecKeyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKeyParams

`func (o *GetGridDnsResponse) SetDnssecKeyParams(v GridDnsDnssecKeyParams)`

SetDnssecKeyParams sets DnssecKeyParams field to given value.

### HasDnssecKeyParams

`func (o *GetGridDnsResponse) HasDnssecKeyParams() bool`

HasDnssecKeyParams returns a boolean if a field has been set.

### GetDnssecNegativeTrustAnchors

`func (o *GetGridDnsResponse) GetDnssecNegativeTrustAnchors() []string`

GetDnssecNegativeTrustAnchors returns the DnssecNegativeTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecNegativeTrustAnchorsOk

`func (o *GetGridDnsResponse) GetDnssecNegativeTrustAnchorsOk() (*[]string, bool)`

GetDnssecNegativeTrustAnchorsOk returns a tuple with the DnssecNegativeTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNegativeTrustAnchors

`func (o *GetGridDnsResponse) SetDnssecNegativeTrustAnchors(v []string)`

SetDnssecNegativeTrustAnchors sets DnssecNegativeTrustAnchors field to given value.

### HasDnssecNegativeTrustAnchors

`func (o *GetGridDnsResponse) HasDnssecNegativeTrustAnchors() bool`

HasDnssecNegativeTrustAnchors returns a boolean if a field has been set.

### GetDnssecNxdomainEnabled

`func (o *GetGridDnsResponse) GetDnssecNxdomainEnabled() bool`

GetDnssecNxdomainEnabled returns the DnssecNxdomainEnabled field if non-nil, zero value otherwise.

### GetDnssecNxdomainEnabledOk

`func (o *GetGridDnsResponse) GetDnssecNxdomainEnabledOk() (*bool, bool)`

GetDnssecNxdomainEnabledOk returns a tuple with the DnssecNxdomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecNxdomainEnabled

`func (o *GetGridDnsResponse) SetDnssecNxdomainEnabled(v bool)`

SetDnssecNxdomainEnabled sets DnssecNxdomainEnabled field to given value.

### HasDnssecNxdomainEnabled

`func (o *GetGridDnsResponse) HasDnssecNxdomainEnabled() bool`

HasDnssecNxdomainEnabled returns a boolean if a field has been set.

### GetDnssecRpzEnabled

`func (o *GetGridDnsResponse) GetDnssecRpzEnabled() bool`

GetDnssecRpzEnabled returns the DnssecRpzEnabled field if non-nil, zero value otherwise.

### GetDnssecRpzEnabledOk

`func (o *GetGridDnsResponse) GetDnssecRpzEnabledOk() (*bool, bool)`

GetDnssecRpzEnabledOk returns a tuple with the DnssecRpzEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRpzEnabled

`func (o *GetGridDnsResponse) SetDnssecRpzEnabled(v bool)`

SetDnssecRpzEnabled sets DnssecRpzEnabled field to given value.

### HasDnssecRpzEnabled

`func (o *GetGridDnsResponse) HasDnssecRpzEnabled() bool`

HasDnssecRpzEnabled returns a boolean if a field has been set.

### GetDnssecTrustedKeys

`func (o *GetGridDnsResponse) GetDnssecTrustedKeys() []GridDnsDnssecTrustedKeys`

GetDnssecTrustedKeys returns the DnssecTrustedKeys field if non-nil, zero value otherwise.

### GetDnssecTrustedKeysOk

`func (o *GetGridDnsResponse) GetDnssecTrustedKeysOk() (*[]GridDnsDnssecTrustedKeys, bool)`

GetDnssecTrustedKeysOk returns a tuple with the DnssecTrustedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustedKeys

`func (o *GetGridDnsResponse) SetDnssecTrustedKeys(v []GridDnsDnssecTrustedKeys)`

SetDnssecTrustedKeys sets DnssecTrustedKeys field to given value.

### HasDnssecTrustedKeys

`func (o *GetGridDnsResponse) HasDnssecTrustedKeys() bool`

HasDnssecTrustedKeys returns a boolean if a field has been set.

### GetDnssecValidationEnabled

`func (o *GetGridDnsResponse) GetDnssecValidationEnabled() bool`

GetDnssecValidationEnabled returns the DnssecValidationEnabled field if non-nil, zero value otherwise.

### GetDnssecValidationEnabledOk

`func (o *GetGridDnsResponse) GetDnssecValidationEnabledOk() (*bool, bool)`

GetDnssecValidationEnabledOk returns a tuple with the DnssecValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationEnabled

`func (o *GetGridDnsResponse) SetDnssecValidationEnabled(v bool)`

SetDnssecValidationEnabled sets DnssecValidationEnabled field to given value.

### HasDnssecValidationEnabled

`func (o *GetGridDnsResponse) HasDnssecValidationEnabled() bool`

HasDnssecValidationEnabled returns a boolean if a field has been set.

### GetDnstapSetting

`func (o *GetGridDnsResponse) GetDnstapSetting() GridDnsDnstapSetting`

GetDnstapSetting returns the DnstapSetting field if non-nil, zero value otherwise.

### GetDnstapSettingOk

`func (o *GetGridDnsResponse) GetDnstapSettingOk() (*GridDnsDnstapSetting, bool)`

GetDnstapSettingOk returns a tuple with the DnstapSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapSetting

`func (o *GetGridDnsResponse) SetDnstapSetting(v GridDnsDnstapSetting)`

SetDnstapSetting sets DnstapSetting field to given value.

### HasDnstapSetting

`func (o *GetGridDnsResponse) HasDnstapSetting() bool`

HasDnstapSetting returns a boolean if a field has been set.

### GetDomainsToCaptureDnsQueries

`func (o *GetGridDnsResponse) GetDomainsToCaptureDnsQueries() []string`

GetDomainsToCaptureDnsQueries returns the DomainsToCaptureDnsQueries field if non-nil, zero value otherwise.

### GetDomainsToCaptureDnsQueriesOk

`func (o *GetGridDnsResponse) GetDomainsToCaptureDnsQueriesOk() (*[]string, bool)`

GetDomainsToCaptureDnsQueriesOk returns a tuple with the DomainsToCaptureDnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsToCaptureDnsQueries

`func (o *GetGridDnsResponse) SetDomainsToCaptureDnsQueries(v []string)`

SetDomainsToCaptureDnsQueries sets DomainsToCaptureDnsQueries field to given value.

### HasDomainsToCaptureDnsQueries

`func (o *GetGridDnsResponse) HasDomainsToCaptureDnsQueries() bool`

HasDomainsToCaptureDnsQueries returns a boolean if a field has been set.

### GetDtcDnsQueriesSpecificBehavior

`func (o *GetGridDnsResponse) GetDtcDnsQueriesSpecificBehavior() string`

GetDtcDnsQueriesSpecificBehavior returns the DtcDnsQueriesSpecificBehavior field if non-nil, zero value otherwise.

### GetDtcDnsQueriesSpecificBehaviorOk

`func (o *GetGridDnsResponse) GetDtcDnsQueriesSpecificBehaviorOk() (*string, bool)`

GetDtcDnsQueriesSpecificBehaviorOk returns a tuple with the DtcDnsQueriesSpecificBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcDnsQueriesSpecificBehavior

`func (o *GetGridDnsResponse) SetDtcDnsQueriesSpecificBehavior(v string)`

SetDtcDnsQueriesSpecificBehavior sets DtcDnsQueriesSpecificBehavior field to given value.

### HasDtcDnsQueriesSpecificBehavior

`func (o *GetGridDnsResponse) HasDtcDnsQueriesSpecificBehavior() bool`

HasDtcDnsQueriesSpecificBehavior returns a boolean if a field has been set.

### GetDtcDnssecMode

`func (o *GetGridDnsResponse) GetDtcDnssecMode() string`

GetDtcDnssecMode returns the DtcDnssecMode field if non-nil, zero value otherwise.

### GetDtcDnssecModeOk

`func (o *GetGridDnsResponse) GetDtcDnssecModeOk() (*string, bool)`

GetDtcDnssecModeOk returns a tuple with the DtcDnssecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcDnssecMode

`func (o *GetGridDnsResponse) SetDtcDnssecMode(v string)`

SetDtcDnssecMode sets DtcDnssecMode field to given value.

### HasDtcDnssecMode

`func (o *GetGridDnsResponse) HasDtcDnssecMode() bool`

HasDtcDnssecMode returns a boolean if a field has been set.

### GetDtcEdnsPreferClientSubnet

`func (o *GetGridDnsResponse) GetDtcEdnsPreferClientSubnet() bool`

GetDtcEdnsPreferClientSubnet returns the DtcEdnsPreferClientSubnet field if non-nil, zero value otherwise.

### GetDtcEdnsPreferClientSubnetOk

`func (o *GetGridDnsResponse) GetDtcEdnsPreferClientSubnetOk() (*bool, bool)`

GetDtcEdnsPreferClientSubnetOk returns a tuple with the DtcEdnsPreferClientSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcEdnsPreferClientSubnet

`func (o *GetGridDnsResponse) SetDtcEdnsPreferClientSubnet(v bool)`

SetDtcEdnsPreferClientSubnet sets DtcEdnsPreferClientSubnet field to given value.

### HasDtcEdnsPreferClientSubnet

`func (o *GetGridDnsResponse) HasDtcEdnsPreferClientSubnet() bool`

HasDtcEdnsPreferClientSubnet returns a boolean if a field has been set.

### GetDtcScheduledBackup

`func (o *GetGridDnsResponse) GetDtcScheduledBackup() GridDnsDtcScheduledBackup`

GetDtcScheduledBackup returns the DtcScheduledBackup field if non-nil, zero value otherwise.

### GetDtcScheduledBackupOk

`func (o *GetGridDnsResponse) GetDtcScheduledBackupOk() (*GridDnsDtcScheduledBackup, bool)`

GetDtcScheduledBackupOk returns a tuple with the DtcScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcScheduledBackup

`func (o *GetGridDnsResponse) SetDtcScheduledBackup(v GridDnsDtcScheduledBackup)`

SetDtcScheduledBackup sets DtcScheduledBackup field to given value.

### HasDtcScheduledBackup

`func (o *GetGridDnsResponse) HasDtcScheduledBackup() bool`

HasDtcScheduledBackup returns a boolean if a field has been set.

### GetDtcTopologyEaList

`func (o *GetGridDnsResponse) GetDtcTopologyEaList() []string`

GetDtcTopologyEaList returns the DtcTopologyEaList field if non-nil, zero value otherwise.

### GetDtcTopologyEaListOk

`func (o *GetGridDnsResponse) GetDtcTopologyEaListOk() (*[]string, bool)`

GetDtcTopologyEaListOk returns a tuple with the DtcTopologyEaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcTopologyEaList

`func (o *GetGridDnsResponse) SetDtcTopologyEaList(v []string)`

SetDtcTopologyEaList sets DtcTopologyEaList field to given value.

### HasDtcTopologyEaList

`func (o *GetGridDnsResponse) HasDtcTopologyEaList() bool`

HasDtcTopologyEaList returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *GetGridDnsResponse) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *GetGridDnsResponse) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *GetGridDnsResponse) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *GetGridDnsResponse) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetEmail

`func (o *GetGridDnsResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetGridDnsResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetGridDnsResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetGridDnsResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableBlackhole

`func (o *GetGridDnsResponse) GetEnableBlackhole() bool`

GetEnableBlackhole returns the EnableBlackhole field if non-nil, zero value otherwise.

### GetEnableBlackholeOk

`func (o *GetGridDnsResponse) GetEnableBlackholeOk() (*bool, bool)`

GetEnableBlackholeOk returns a tuple with the EnableBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackhole

`func (o *GetGridDnsResponse) SetEnableBlackhole(v bool)`

SetEnableBlackhole sets EnableBlackhole field to given value.

### HasEnableBlackhole

`func (o *GetGridDnsResponse) HasEnableBlackhole() bool`

HasEnableBlackhole returns a boolean if a field has been set.

### GetEnableBlacklist

`func (o *GetGridDnsResponse) GetEnableBlacklist() bool`

GetEnableBlacklist returns the EnableBlacklist field if non-nil, zero value otherwise.

### GetEnableBlacklistOk

`func (o *GetGridDnsResponse) GetEnableBlacklistOk() (*bool, bool)`

GetEnableBlacklistOk returns a tuple with the EnableBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlacklist

`func (o *GetGridDnsResponse) SetEnableBlacklist(v bool)`

SetEnableBlacklist sets EnableBlacklist field to given value.

### HasEnableBlacklist

`func (o *GetGridDnsResponse) HasEnableBlacklist() bool`

HasEnableBlacklist returns a boolean if a field has been set.

### GetEnableCaptureDnsQueries

`func (o *GetGridDnsResponse) GetEnableCaptureDnsQueries() bool`

GetEnableCaptureDnsQueries returns the EnableCaptureDnsQueries field if non-nil, zero value otherwise.

### GetEnableCaptureDnsQueriesOk

`func (o *GetGridDnsResponse) GetEnableCaptureDnsQueriesOk() (*bool, bool)`

GetEnableCaptureDnsQueriesOk returns a tuple with the EnableCaptureDnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaptureDnsQueries

`func (o *GetGridDnsResponse) SetEnableCaptureDnsQueries(v bool)`

SetEnableCaptureDnsQueries sets EnableCaptureDnsQueries field to given value.

### HasEnableCaptureDnsQueries

`func (o *GetGridDnsResponse) HasEnableCaptureDnsQueries() bool`

HasEnableCaptureDnsQueries returns a boolean if a field has been set.

### GetEnableCaptureDnsResponses

`func (o *GetGridDnsResponse) GetEnableCaptureDnsResponses() bool`

GetEnableCaptureDnsResponses returns the EnableCaptureDnsResponses field if non-nil, zero value otherwise.

### GetEnableCaptureDnsResponsesOk

`func (o *GetGridDnsResponse) GetEnableCaptureDnsResponsesOk() (*bool, bool)`

GetEnableCaptureDnsResponsesOk returns a tuple with the EnableCaptureDnsResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaptureDnsResponses

`func (o *GetGridDnsResponse) SetEnableCaptureDnsResponses(v bool)`

SetEnableCaptureDnsResponses sets EnableCaptureDnsResponses field to given value.

### HasEnableCaptureDnsResponses

`func (o *GetGridDnsResponse) HasEnableCaptureDnsResponses() bool`

HasEnableCaptureDnsResponses returns a boolean if a field has been set.

### GetEnableClientSubnetForwarding

`func (o *GetGridDnsResponse) GetEnableClientSubnetForwarding() bool`

GetEnableClientSubnetForwarding returns the EnableClientSubnetForwarding field if non-nil, zero value otherwise.

### GetEnableClientSubnetForwardingOk

`func (o *GetGridDnsResponse) GetEnableClientSubnetForwardingOk() (*bool, bool)`

GetEnableClientSubnetForwardingOk returns a tuple with the EnableClientSubnetForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientSubnetForwarding

`func (o *GetGridDnsResponse) SetEnableClientSubnetForwarding(v bool)`

SetEnableClientSubnetForwarding sets EnableClientSubnetForwarding field to given value.

### HasEnableClientSubnetForwarding

`func (o *GetGridDnsResponse) HasEnableClientSubnetForwarding() bool`

HasEnableClientSubnetForwarding returns a boolean if a field has been set.

### GetEnableClientSubnetRecursive

`func (o *GetGridDnsResponse) GetEnableClientSubnetRecursive() bool`

GetEnableClientSubnetRecursive returns the EnableClientSubnetRecursive field if non-nil, zero value otherwise.

### GetEnableClientSubnetRecursiveOk

`func (o *GetGridDnsResponse) GetEnableClientSubnetRecursiveOk() (*bool, bool)`

GetEnableClientSubnetRecursiveOk returns a tuple with the EnableClientSubnetRecursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientSubnetRecursive

`func (o *GetGridDnsResponse) SetEnableClientSubnetRecursive(v bool)`

SetEnableClientSubnetRecursive sets EnableClientSubnetRecursive field to given value.

### HasEnableClientSubnetRecursive

`func (o *GetGridDnsResponse) HasEnableClientSubnetRecursive() bool`

HasEnableClientSubnetRecursive returns a boolean if a field has been set.

### GetEnableDeleteAssociatedPtr

`func (o *GetGridDnsResponse) GetEnableDeleteAssociatedPtr() bool`

GetEnableDeleteAssociatedPtr returns the EnableDeleteAssociatedPtr field if non-nil, zero value otherwise.

### GetEnableDeleteAssociatedPtrOk

`func (o *GetGridDnsResponse) GetEnableDeleteAssociatedPtrOk() (*bool, bool)`

GetEnableDeleteAssociatedPtrOk returns a tuple with the EnableDeleteAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeleteAssociatedPtr

`func (o *GetGridDnsResponse) SetEnableDeleteAssociatedPtr(v bool)`

SetEnableDeleteAssociatedPtr sets EnableDeleteAssociatedPtr field to given value.

### HasEnableDeleteAssociatedPtr

`func (o *GetGridDnsResponse) HasEnableDeleteAssociatedPtr() bool`

HasEnableDeleteAssociatedPtr returns a boolean if a field has been set.

### GetEnableDns64

`func (o *GetGridDnsResponse) GetEnableDns64() bool`

GetEnableDns64 returns the EnableDns64 field if non-nil, zero value otherwise.

### GetEnableDns64Ok

`func (o *GetGridDnsResponse) GetEnableDns64Ok() (*bool, bool)`

GetEnableDns64Ok returns a tuple with the EnableDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns64

`func (o *GetGridDnsResponse) SetEnableDns64(v bool)`

SetEnableDns64 sets EnableDns64 field to given value.

### HasEnableDns64

`func (o *GetGridDnsResponse) HasEnableDns64() bool`

HasEnableDns64 returns a boolean if a field has been set.

### GetEnableDnsHealthCheck

`func (o *GetGridDnsResponse) GetEnableDnsHealthCheck() bool`

GetEnableDnsHealthCheck returns the EnableDnsHealthCheck field if non-nil, zero value otherwise.

### GetEnableDnsHealthCheckOk

`func (o *GetGridDnsResponse) GetEnableDnsHealthCheckOk() (*bool, bool)`

GetEnableDnsHealthCheckOk returns a tuple with the EnableDnsHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsHealthCheck

`func (o *GetGridDnsResponse) SetEnableDnsHealthCheck(v bool)`

SetEnableDnsHealthCheck sets EnableDnsHealthCheck field to given value.

### HasEnableDnsHealthCheck

`func (o *GetGridDnsResponse) HasEnableDnsHealthCheck() bool`

HasEnableDnsHealthCheck returns a boolean if a field has been set.

### GetEnableDnstapQueries

`func (o *GetGridDnsResponse) GetEnableDnstapQueries() bool`

GetEnableDnstapQueries returns the EnableDnstapQueries field if non-nil, zero value otherwise.

### GetEnableDnstapQueriesOk

`func (o *GetGridDnsResponse) GetEnableDnstapQueriesOk() (*bool, bool)`

GetEnableDnstapQueriesOk returns a tuple with the EnableDnstapQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapQueries

`func (o *GetGridDnsResponse) SetEnableDnstapQueries(v bool)`

SetEnableDnstapQueries sets EnableDnstapQueries field to given value.

### HasEnableDnstapQueries

`func (o *GetGridDnsResponse) HasEnableDnstapQueries() bool`

HasEnableDnstapQueries returns a boolean if a field has been set.

### GetEnableDnstapResponses

`func (o *GetGridDnsResponse) GetEnableDnstapResponses() bool`

GetEnableDnstapResponses returns the EnableDnstapResponses field if non-nil, zero value otherwise.

### GetEnableDnstapResponsesOk

`func (o *GetGridDnsResponse) GetEnableDnstapResponsesOk() (*bool, bool)`

GetEnableDnstapResponsesOk returns a tuple with the EnableDnstapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnstapResponses

`func (o *GetGridDnsResponse) SetEnableDnstapResponses(v bool)`

SetEnableDnstapResponses sets EnableDnstapResponses field to given value.

### HasEnableDnstapResponses

`func (o *GetGridDnsResponse) HasEnableDnstapResponses() bool`

HasEnableDnstapResponses returns a boolean if a field has been set.

### GetEnableExcludedDomainNames

`func (o *GetGridDnsResponse) GetEnableExcludedDomainNames() bool`

GetEnableExcludedDomainNames returns the EnableExcludedDomainNames field if non-nil, zero value otherwise.

### GetEnableExcludedDomainNamesOk

`func (o *GetGridDnsResponse) GetEnableExcludedDomainNamesOk() (*bool, bool)`

GetEnableExcludedDomainNamesOk returns a tuple with the EnableExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExcludedDomainNames

`func (o *GetGridDnsResponse) SetEnableExcludedDomainNames(v bool)`

SetEnableExcludedDomainNames sets EnableExcludedDomainNames field to given value.

### HasEnableExcludedDomainNames

`func (o *GetGridDnsResponse) HasEnableExcludedDomainNames() bool`

HasEnableExcludedDomainNames returns a boolean if a field has been set.

### GetEnableFixedRrsetOrderFqdns

`func (o *GetGridDnsResponse) GetEnableFixedRrsetOrderFqdns() bool`

GetEnableFixedRrsetOrderFqdns returns the EnableFixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetEnableFixedRrsetOrderFqdnsOk

`func (o *GetGridDnsResponse) GetEnableFixedRrsetOrderFqdnsOk() (*bool, bool)`

GetEnableFixedRrsetOrderFqdnsOk returns a tuple with the EnableFixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixedRrsetOrderFqdns

`func (o *GetGridDnsResponse) SetEnableFixedRrsetOrderFqdns(v bool)`

SetEnableFixedRrsetOrderFqdns sets EnableFixedRrsetOrderFqdns field to given value.

### HasEnableFixedRrsetOrderFqdns

`func (o *GetGridDnsResponse) HasEnableFixedRrsetOrderFqdns() bool`

HasEnableFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetEnableFtc

`func (o *GetGridDnsResponse) GetEnableFtc() bool`

GetEnableFtc returns the EnableFtc field if non-nil, zero value otherwise.

### GetEnableFtcOk

`func (o *GetGridDnsResponse) GetEnableFtcOk() (*bool, bool)`

GetEnableFtcOk returns a tuple with the EnableFtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtc

`func (o *GetGridDnsResponse) SetEnableFtc(v bool)`

SetEnableFtc sets EnableFtc field to given value.

### HasEnableFtc

`func (o *GetGridDnsResponse) HasEnableFtc() bool`

HasEnableFtc returns a boolean if a field has been set.

### GetEnableGssTsig

`func (o *GetGridDnsResponse) GetEnableGssTsig() bool`

GetEnableGssTsig returns the EnableGssTsig field if non-nil, zero value otherwise.

### GetEnableGssTsigOk

`func (o *GetGridDnsResponse) GetEnableGssTsigOk() (*bool, bool)`

GetEnableGssTsigOk returns a tuple with the EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGssTsig

`func (o *GetGridDnsResponse) SetEnableGssTsig(v bool)`

SetEnableGssTsig sets EnableGssTsig field to given value.

### HasEnableGssTsig

`func (o *GetGridDnsResponse) HasEnableGssTsig() bool`

HasEnableGssTsig returns a boolean if a field has been set.

### GetEnableHostRrsetOrder

`func (o *GetGridDnsResponse) GetEnableHostRrsetOrder() bool`

GetEnableHostRrsetOrder returns the EnableHostRrsetOrder field if non-nil, zero value otherwise.

### GetEnableHostRrsetOrderOk

`func (o *GetGridDnsResponse) GetEnableHostRrsetOrderOk() (*bool, bool)`

GetEnableHostRrsetOrderOk returns a tuple with the EnableHostRrsetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHostRrsetOrder

`func (o *GetGridDnsResponse) SetEnableHostRrsetOrder(v bool)`

SetEnableHostRrsetOrder sets EnableHostRrsetOrder field to given value.

### HasEnableHostRrsetOrder

`func (o *GetGridDnsResponse) HasEnableHostRrsetOrder() bool`

HasEnableHostRrsetOrder returns a boolean if a field has been set.

### GetEnableHsmSigning

`func (o *GetGridDnsResponse) GetEnableHsmSigning() bool`

GetEnableHsmSigning returns the EnableHsmSigning field if non-nil, zero value otherwise.

### GetEnableHsmSigningOk

`func (o *GetGridDnsResponse) GetEnableHsmSigningOk() (*bool, bool)`

GetEnableHsmSigningOk returns a tuple with the EnableHsmSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHsmSigning

`func (o *GetGridDnsResponse) SetEnableHsmSigning(v bool)`

SetEnableHsmSigning sets EnableHsmSigning field to given value.

### HasEnableHsmSigning

`func (o *GetGridDnsResponse) HasEnableHsmSigning() bool`

HasEnableHsmSigning returns a boolean if a field has been set.

### GetEnableNotifySourcePort

`func (o *GetGridDnsResponse) GetEnableNotifySourcePort() bool`

GetEnableNotifySourcePort returns the EnableNotifySourcePort field if non-nil, zero value otherwise.

### GetEnableNotifySourcePortOk

`func (o *GetGridDnsResponse) GetEnableNotifySourcePortOk() (*bool, bool)`

GetEnableNotifySourcePortOk returns a tuple with the EnableNotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySourcePort

`func (o *GetGridDnsResponse) SetEnableNotifySourcePort(v bool)`

SetEnableNotifySourcePort sets EnableNotifySourcePort field to given value.

### HasEnableNotifySourcePort

`func (o *GetGridDnsResponse) HasEnableNotifySourcePort() bool`

HasEnableNotifySourcePort returns a boolean if a field has been set.

### GetEnableQueryRewrite

`func (o *GetGridDnsResponse) GetEnableQueryRewrite() bool`

GetEnableQueryRewrite returns the EnableQueryRewrite field if non-nil, zero value otherwise.

### GetEnableQueryRewriteOk

`func (o *GetGridDnsResponse) GetEnableQueryRewriteOk() (*bool, bool)`

GetEnableQueryRewriteOk returns a tuple with the EnableQueryRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryRewrite

`func (o *GetGridDnsResponse) SetEnableQueryRewrite(v bool)`

SetEnableQueryRewrite sets EnableQueryRewrite field to given value.

### HasEnableQueryRewrite

`func (o *GetGridDnsResponse) HasEnableQueryRewrite() bool`

HasEnableQueryRewrite returns a boolean if a field has been set.

### GetEnableQuerySourcePort

`func (o *GetGridDnsResponse) GetEnableQuerySourcePort() bool`

GetEnableQuerySourcePort returns the EnableQuerySourcePort field if non-nil, zero value otherwise.

### GetEnableQuerySourcePortOk

`func (o *GetGridDnsResponse) GetEnableQuerySourcePortOk() (*bool, bool)`

GetEnableQuerySourcePortOk returns a tuple with the EnableQuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuerySourcePort

`func (o *GetGridDnsResponse) SetEnableQuerySourcePort(v bool)`

SetEnableQuerySourcePort sets EnableQuerySourcePort field to given value.

### HasEnableQuerySourcePort

`func (o *GetGridDnsResponse) HasEnableQuerySourcePort() bool`

HasEnableQuerySourcePort returns a boolean if a field has been set.

### GetExcludedDomainNames

`func (o *GetGridDnsResponse) GetExcludedDomainNames() []string`

GetExcludedDomainNames returns the ExcludedDomainNames field if non-nil, zero value otherwise.

### GetExcludedDomainNamesOk

`func (o *GetGridDnsResponse) GetExcludedDomainNamesOk() (*[]string, bool)`

GetExcludedDomainNamesOk returns a tuple with the ExcludedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDomainNames

`func (o *GetGridDnsResponse) SetExcludedDomainNames(v []string)`

SetExcludedDomainNames sets ExcludedDomainNames field to given value.

### HasExcludedDomainNames

`func (o *GetGridDnsResponse) HasExcludedDomainNames() bool`

HasExcludedDomainNames returns a boolean if a field has been set.

### GetExpireAfter

`func (o *GetGridDnsResponse) GetExpireAfter() int64`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *GetGridDnsResponse) GetExpireAfterOk() (*int64, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *GetGridDnsResponse) SetExpireAfter(v int64)`

SetExpireAfter sets ExpireAfter field to given value.

### HasExpireAfter

`func (o *GetGridDnsResponse) HasExpireAfter() bool`

HasExpireAfter returns a boolean if a field has been set.

### GetFileTransferSetting

`func (o *GetGridDnsResponse) GetFileTransferSetting() GridDnsFileTransferSetting`

GetFileTransferSetting returns the FileTransferSetting field if non-nil, zero value otherwise.

### GetFileTransferSettingOk

`func (o *GetGridDnsResponse) GetFileTransferSettingOk() (*GridDnsFileTransferSetting, bool)`

GetFileTransferSettingOk returns a tuple with the FileTransferSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferSetting

`func (o *GetGridDnsResponse) SetFileTransferSetting(v GridDnsFileTransferSetting)`

SetFileTransferSetting sets FileTransferSetting field to given value.

### HasFileTransferSetting

`func (o *GetGridDnsResponse) HasFileTransferSetting() bool`

HasFileTransferSetting returns a boolean if a field has been set.

### GetFilterAaaa

`func (o *GetGridDnsResponse) GetFilterAaaa() string`

GetFilterAaaa returns the FilterAaaa field if non-nil, zero value otherwise.

### GetFilterAaaaOk

`func (o *GetGridDnsResponse) GetFilterAaaaOk() (*string, bool)`

GetFilterAaaaOk returns a tuple with the FilterAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaa

`func (o *GetGridDnsResponse) SetFilterAaaa(v string)`

SetFilterAaaa sets FilterAaaa field to given value.

### HasFilterAaaa

`func (o *GetGridDnsResponse) HasFilterAaaa() bool`

HasFilterAaaa returns a boolean if a field has been set.

### GetFilterAaaaList

`func (o *GetGridDnsResponse) GetFilterAaaaList() []GridDnsFilterAaaaList`

GetFilterAaaaList returns the FilterAaaaList field if non-nil, zero value otherwise.

### GetFilterAaaaListOk

`func (o *GetGridDnsResponse) GetFilterAaaaListOk() (*[]GridDnsFilterAaaaList, bool)`

GetFilterAaaaListOk returns a tuple with the FilterAaaaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaList

`func (o *GetGridDnsResponse) SetFilterAaaaList(v []GridDnsFilterAaaaList)`

SetFilterAaaaList sets FilterAaaaList field to given value.

### HasFilterAaaaList

`func (o *GetGridDnsResponse) HasFilterAaaaList() bool`

HasFilterAaaaList returns a boolean if a field has been set.

### GetFixedRrsetOrderFqdns

`func (o *GetGridDnsResponse) GetFixedRrsetOrderFqdns() []GridDnsFixedRrsetOrderFqdns`

GetFixedRrsetOrderFqdns returns the FixedRrsetOrderFqdns field if non-nil, zero value otherwise.

### GetFixedRrsetOrderFqdnsOk

`func (o *GetGridDnsResponse) GetFixedRrsetOrderFqdnsOk() (*[]GridDnsFixedRrsetOrderFqdns, bool)`

GetFixedRrsetOrderFqdnsOk returns a tuple with the FixedRrsetOrderFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedRrsetOrderFqdns

`func (o *GetGridDnsResponse) SetFixedRrsetOrderFqdns(v []GridDnsFixedRrsetOrderFqdns)`

SetFixedRrsetOrderFqdns sets FixedRrsetOrderFqdns field to given value.

### HasFixedRrsetOrderFqdns

`func (o *GetGridDnsResponse) HasFixedRrsetOrderFqdns() bool`

HasFixedRrsetOrderFqdns returns a boolean if a field has been set.

### GetForwardOnly

`func (o *GetGridDnsResponse) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *GetGridDnsResponse) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *GetGridDnsResponse) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *GetGridDnsResponse) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetForwardUpdates

`func (o *GetGridDnsResponse) GetForwardUpdates() bool`

GetForwardUpdates returns the ForwardUpdates field if non-nil, zero value otherwise.

### GetForwardUpdatesOk

`func (o *GetGridDnsResponse) GetForwardUpdatesOk() (*bool, bool)`

GetForwardUpdatesOk returns a tuple with the ForwardUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUpdates

`func (o *GetGridDnsResponse) SetForwardUpdates(v bool)`

SetForwardUpdates sets ForwardUpdates field to given value.

### HasForwardUpdates

`func (o *GetGridDnsResponse) HasForwardUpdates() bool`

HasForwardUpdates returns a boolean if a field has been set.

### GetForwarders

`func (o *GetGridDnsResponse) GetForwarders() []string`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *GetGridDnsResponse) GetForwardersOk() (*[]string, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *GetGridDnsResponse) SetForwarders(v []string)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *GetGridDnsResponse) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetFtcExpiredRecordTimeout

`func (o *GetGridDnsResponse) GetFtcExpiredRecordTimeout() int64`

GetFtcExpiredRecordTimeout returns the FtcExpiredRecordTimeout field if non-nil, zero value otherwise.

### GetFtcExpiredRecordTimeoutOk

`func (o *GetGridDnsResponse) GetFtcExpiredRecordTimeoutOk() (*int64, bool)`

GetFtcExpiredRecordTimeoutOk returns a tuple with the FtcExpiredRecordTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtcExpiredRecordTimeout

`func (o *GetGridDnsResponse) SetFtcExpiredRecordTimeout(v int64)`

SetFtcExpiredRecordTimeout sets FtcExpiredRecordTimeout field to given value.

### HasFtcExpiredRecordTimeout

`func (o *GetGridDnsResponse) HasFtcExpiredRecordTimeout() bool`

HasFtcExpiredRecordTimeout returns a boolean if a field has been set.

### GetFtcExpiredRecordTtl

`func (o *GetGridDnsResponse) GetFtcExpiredRecordTtl() int64`

GetFtcExpiredRecordTtl returns the FtcExpiredRecordTtl field if non-nil, zero value otherwise.

### GetFtcExpiredRecordTtlOk

`func (o *GetGridDnsResponse) GetFtcExpiredRecordTtlOk() (*int64, bool)`

GetFtcExpiredRecordTtlOk returns a tuple with the FtcExpiredRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtcExpiredRecordTtl

`func (o *GetGridDnsResponse) SetFtcExpiredRecordTtl(v int64)`

SetFtcExpiredRecordTtl sets FtcExpiredRecordTtl field to given value.

### HasFtcExpiredRecordTtl

`func (o *GetGridDnsResponse) HasFtcExpiredRecordTtl() bool`

HasFtcExpiredRecordTtl returns a boolean if a field has been set.

### GetGenEadbFromHosts

`func (o *GetGridDnsResponse) GetGenEadbFromHosts() bool`

GetGenEadbFromHosts returns the GenEadbFromHosts field if non-nil, zero value otherwise.

### GetGenEadbFromHostsOk

`func (o *GetGridDnsResponse) GetGenEadbFromHostsOk() (*bool, bool)`

GetGenEadbFromHostsOk returns a tuple with the GenEadbFromHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromHosts

`func (o *GetGridDnsResponse) SetGenEadbFromHosts(v bool)`

SetGenEadbFromHosts sets GenEadbFromHosts field to given value.

### HasGenEadbFromHosts

`func (o *GetGridDnsResponse) HasGenEadbFromHosts() bool`

HasGenEadbFromHosts returns a boolean if a field has been set.

### GetGenEadbFromNetworkContainers

`func (o *GetGridDnsResponse) GetGenEadbFromNetworkContainers() bool`

GetGenEadbFromNetworkContainers returns the GenEadbFromNetworkContainers field if non-nil, zero value otherwise.

### GetGenEadbFromNetworkContainersOk

`func (o *GetGridDnsResponse) GetGenEadbFromNetworkContainersOk() (*bool, bool)`

GetGenEadbFromNetworkContainersOk returns a tuple with the GenEadbFromNetworkContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromNetworkContainers

`func (o *GetGridDnsResponse) SetGenEadbFromNetworkContainers(v bool)`

SetGenEadbFromNetworkContainers sets GenEadbFromNetworkContainers field to given value.

### HasGenEadbFromNetworkContainers

`func (o *GetGridDnsResponse) HasGenEadbFromNetworkContainers() bool`

HasGenEadbFromNetworkContainers returns a boolean if a field has been set.

### GetGenEadbFromNetworks

`func (o *GetGridDnsResponse) GetGenEadbFromNetworks() bool`

GetGenEadbFromNetworks returns the GenEadbFromNetworks field if non-nil, zero value otherwise.

### GetGenEadbFromNetworksOk

`func (o *GetGridDnsResponse) GetGenEadbFromNetworksOk() (*bool, bool)`

GetGenEadbFromNetworksOk returns a tuple with the GenEadbFromNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromNetworks

`func (o *GetGridDnsResponse) SetGenEadbFromNetworks(v bool)`

SetGenEadbFromNetworks sets GenEadbFromNetworks field to given value.

### HasGenEadbFromNetworks

`func (o *GetGridDnsResponse) HasGenEadbFromNetworks() bool`

HasGenEadbFromNetworks returns a boolean if a field has been set.

### GetGenEadbFromRanges

`func (o *GetGridDnsResponse) GetGenEadbFromRanges() bool`

GetGenEadbFromRanges returns the GenEadbFromRanges field if non-nil, zero value otherwise.

### GetGenEadbFromRangesOk

`func (o *GetGridDnsResponse) GetGenEadbFromRangesOk() (*bool, bool)`

GetGenEadbFromRangesOk returns a tuple with the GenEadbFromRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenEadbFromRanges

`func (o *GetGridDnsResponse) SetGenEadbFromRanges(v bool)`

SetGenEadbFromRanges sets GenEadbFromRanges field to given value.

### HasGenEadbFromRanges

`func (o *GetGridDnsResponse) HasGenEadbFromRanges() bool`

HasGenEadbFromRanges returns a boolean if a field has been set.

### GetGssTsigKeys

`func (o *GetGridDnsResponse) GetGssTsigKeys() []string`

GetGssTsigKeys returns the GssTsigKeys field if non-nil, zero value otherwise.

### GetGssTsigKeysOk

`func (o *GetGridDnsResponse) GetGssTsigKeysOk() (*[]string, bool)`

GetGssTsigKeysOk returns a tuple with the GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigKeys

`func (o *GetGridDnsResponse) SetGssTsigKeys(v []string)`

SetGssTsigKeys sets GssTsigKeys field to given value.

### HasGssTsigKeys

`func (o *GetGridDnsResponse) HasGssTsigKeys() bool`

HasGssTsigKeys returns a boolean if a field has been set.

### GetLastQueriedAcl

`func (o *GetGridDnsResponse) GetLastQueriedAcl() []GridDnsLastQueriedAcl`

GetLastQueriedAcl returns the LastQueriedAcl field if non-nil, zero value otherwise.

### GetLastQueriedAclOk

`func (o *GetGridDnsResponse) GetLastQueriedAclOk() (*[]GridDnsLastQueriedAcl, bool)`

GetLastQueriedAclOk returns a tuple with the LastQueriedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueriedAcl

`func (o *GetGridDnsResponse) SetLastQueriedAcl(v []GridDnsLastQueriedAcl)`

SetLastQueriedAcl sets LastQueriedAcl field to given value.

### HasLastQueriedAcl

`func (o *GetGridDnsResponse) HasLastQueriedAcl() bool`

HasLastQueriedAcl returns a boolean if a field has been set.

### GetLoggingCategories

`func (o *GetGridDnsResponse) GetLoggingCategories() GridDnsLoggingCategories`

GetLoggingCategories returns the LoggingCategories field if non-nil, zero value otherwise.

### GetLoggingCategoriesOk

`func (o *GetGridDnsResponse) GetLoggingCategoriesOk() (*GridDnsLoggingCategories, bool)`

GetLoggingCategoriesOk returns a tuple with the LoggingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCategories

`func (o *GetGridDnsResponse) SetLoggingCategories(v GridDnsLoggingCategories)`

SetLoggingCategories sets LoggingCategories field to given value.

### HasLoggingCategories

`func (o *GetGridDnsResponse) HasLoggingCategories() bool`

HasLoggingCategories returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *GetGridDnsResponse) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *GetGridDnsResponse) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *GetGridDnsResponse) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *GetGridDnsResponse) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxCachedLifetime

`func (o *GetGridDnsResponse) GetMaxCachedLifetime() int64`

GetMaxCachedLifetime returns the MaxCachedLifetime field if non-nil, zero value otherwise.

### GetMaxCachedLifetimeOk

`func (o *GetGridDnsResponse) GetMaxCachedLifetimeOk() (*int64, bool)`

GetMaxCachedLifetimeOk returns a tuple with the MaxCachedLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCachedLifetime

`func (o *GetGridDnsResponse) SetMaxCachedLifetime(v int64)`

SetMaxCachedLifetime sets MaxCachedLifetime field to given value.

### HasMaxCachedLifetime

`func (o *GetGridDnsResponse) HasMaxCachedLifetime() bool`

HasMaxCachedLifetime returns a boolean if a field has been set.

### GetMaxNcacheTtl

`func (o *GetGridDnsResponse) GetMaxNcacheTtl() int64`

GetMaxNcacheTtl returns the MaxNcacheTtl field if non-nil, zero value otherwise.

### GetMaxNcacheTtlOk

`func (o *GetGridDnsResponse) GetMaxNcacheTtlOk() (*int64, bool)`

GetMaxNcacheTtlOk returns a tuple with the MaxNcacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNcacheTtl

`func (o *GetGridDnsResponse) SetMaxNcacheTtl(v int64)`

SetMaxNcacheTtl sets MaxNcacheTtl field to given value.

### HasMaxNcacheTtl

`func (o *GetGridDnsResponse) HasMaxNcacheTtl() bool`

HasMaxNcacheTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *GetGridDnsResponse) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *GetGridDnsResponse) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *GetGridDnsResponse) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *GetGridDnsResponse) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMemberSecondaryNotify

`func (o *GetGridDnsResponse) GetMemberSecondaryNotify() bool`

GetMemberSecondaryNotify returns the MemberSecondaryNotify field if non-nil, zero value otherwise.

### GetMemberSecondaryNotifyOk

`func (o *GetGridDnsResponse) GetMemberSecondaryNotifyOk() (*bool, bool)`

GetMemberSecondaryNotifyOk returns a tuple with the MemberSecondaryNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSecondaryNotify

`func (o *GetGridDnsResponse) SetMemberSecondaryNotify(v bool)`

SetMemberSecondaryNotify sets MemberSecondaryNotify field to given value.

### HasMemberSecondaryNotify

`func (o *GetGridDnsResponse) HasMemberSecondaryNotify() bool`

HasMemberSecondaryNotify returns a boolean if a field has been set.

### GetNegativeTtl

`func (o *GetGridDnsResponse) GetNegativeTtl() int64`

GetNegativeTtl returns the NegativeTtl field if non-nil, zero value otherwise.

### GetNegativeTtlOk

`func (o *GetGridDnsResponse) GetNegativeTtlOk() (*int64, bool)`

GetNegativeTtlOk returns a tuple with the NegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeTtl

`func (o *GetGridDnsResponse) SetNegativeTtl(v int64)`

SetNegativeTtl sets NegativeTtl field to given value.

### HasNegativeTtl

`func (o *GetGridDnsResponse) HasNegativeTtl() bool`

HasNegativeTtl returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *GetGridDnsResponse) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *GetGridDnsResponse) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *GetGridDnsResponse) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *GetGridDnsResponse) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetNotifySourcePort

`func (o *GetGridDnsResponse) GetNotifySourcePort() int64`

GetNotifySourcePort returns the NotifySourcePort field if non-nil, zero value otherwise.

### GetNotifySourcePortOk

`func (o *GetGridDnsResponse) GetNotifySourcePortOk() (*int64, bool)`

GetNotifySourcePortOk returns a tuple with the NotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySourcePort

`func (o *GetGridDnsResponse) SetNotifySourcePort(v int64)`

SetNotifySourcePort sets NotifySourcePort field to given value.

### HasNotifySourcePort

`func (o *GetGridDnsResponse) HasNotifySourcePort() bool`

HasNotifySourcePort returns a boolean if a field has been set.

### GetNsgroupDefault

`func (o *GetGridDnsResponse) GetNsgroupDefault() string`

GetNsgroupDefault returns the NsgroupDefault field if non-nil, zero value otherwise.

### GetNsgroupDefaultOk

`func (o *GetGridDnsResponse) GetNsgroupDefaultOk() (*string, bool)`

GetNsgroupDefaultOk returns a tuple with the NsgroupDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgroupDefault

`func (o *GetGridDnsResponse) SetNsgroupDefault(v string)`

SetNsgroupDefault sets NsgroupDefault field to given value.

### HasNsgroupDefault

`func (o *GetGridDnsResponse) HasNsgroupDefault() bool`

HasNsgroupDefault returns a boolean if a field has been set.

### GetNsgroups

`func (o *GetGridDnsResponse) GetNsgroups() []string`

GetNsgroups returns the Nsgroups field if non-nil, zero value otherwise.

### GetNsgroupsOk

`func (o *GetGridDnsResponse) GetNsgroupsOk() (*[]string, bool)`

GetNsgroupsOk returns a tuple with the Nsgroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgroups

`func (o *GetGridDnsResponse) SetNsgroups(v []string)`

SetNsgroups sets Nsgroups field to given value.

### HasNsgroups

`func (o *GetGridDnsResponse) HasNsgroups() bool`

HasNsgroups returns a boolean if a field has been set.

### GetNxdomainLogQuery

`func (o *GetGridDnsResponse) GetNxdomainLogQuery() bool`

GetNxdomainLogQuery returns the NxdomainLogQuery field if non-nil, zero value otherwise.

### GetNxdomainLogQueryOk

`func (o *GetGridDnsResponse) GetNxdomainLogQueryOk() (*bool, bool)`

GetNxdomainLogQueryOk returns a tuple with the NxdomainLogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainLogQuery

`func (o *GetGridDnsResponse) SetNxdomainLogQuery(v bool)`

SetNxdomainLogQuery sets NxdomainLogQuery field to given value.

### HasNxdomainLogQuery

`func (o *GetGridDnsResponse) HasNxdomainLogQuery() bool`

HasNxdomainLogQuery returns a boolean if a field has been set.

### GetNxdomainRedirect

`func (o *GetGridDnsResponse) GetNxdomainRedirect() bool`

GetNxdomainRedirect returns the NxdomainRedirect field if non-nil, zero value otherwise.

### GetNxdomainRedirectOk

`func (o *GetGridDnsResponse) GetNxdomainRedirectOk() (*bool, bool)`

GetNxdomainRedirectOk returns a tuple with the NxdomainRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirect

`func (o *GetGridDnsResponse) SetNxdomainRedirect(v bool)`

SetNxdomainRedirect sets NxdomainRedirect field to given value.

### HasNxdomainRedirect

`func (o *GetGridDnsResponse) HasNxdomainRedirect() bool`

HasNxdomainRedirect returns a boolean if a field has been set.

### GetNxdomainRedirectAddresses

`func (o *GetGridDnsResponse) GetNxdomainRedirectAddresses() []string`

GetNxdomainRedirectAddresses returns the NxdomainRedirectAddresses field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesOk

`func (o *GetGridDnsResponse) GetNxdomainRedirectAddressesOk() (*[]string, bool)`

GetNxdomainRedirectAddressesOk returns a tuple with the NxdomainRedirectAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddresses

`func (o *GetGridDnsResponse) SetNxdomainRedirectAddresses(v []string)`

SetNxdomainRedirectAddresses sets NxdomainRedirectAddresses field to given value.

### HasNxdomainRedirectAddresses

`func (o *GetGridDnsResponse) HasNxdomainRedirectAddresses() bool`

HasNxdomainRedirectAddresses returns a boolean if a field has been set.

### GetNxdomainRedirectAddressesV6

`func (o *GetGridDnsResponse) GetNxdomainRedirectAddressesV6() []string`

GetNxdomainRedirectAddressesV6 returns the NxdomainRedirectAddressesV6 field if non-nil, zero value otherwise.

### GetNxdomainRedirectAddressesV6Ok

`func (o *GetGridDnsResponse) GetNxdomainRedirectAddressesV6Ok() (*[]string, bool)`

GetNxdomainRedirectAddressesV6Ok returns a tuple with the NxdomainRedirectAddressesV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectAddressesV6

`func (o *GetGridDnsResponse) SetNxdomainRedirectAddressesV6(v []string)`

SetNxdomainRedirectAddressesV6 sets NxdomainRedirectAddressesV6 field to given value.

### HasNxdomainRedirectAddressesV6

`func (o *GetGridDnsResponse) HasNxdomainRedirectAddressesV6() bool`

HasNxdomainRedirectAddressesV6 returns a boolean if a field has been set.

### GetNxdomainRedirectTtl

`func (o *GetGridDnsResponse) GetNxdomainRedirectTtl() int64`

GetNxdomainRedirectTtl returns the NxdomainRedirectTtl field if non-nil, zero value otherwise.

### GetNxdomainRedirectTtlOk

`func (o *GetGridDnsResponse) GetNxdomainRedirectTtlOk() (*int64, bool)`

GetNxdomainRedirectTtlOk returns a tuple with the NxdomainRedirectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRedirectTtl

`func (o *GetGridDnsResponse) SetNxdomainRedirectTtl(v int64)`

SetNxdomainRedirectTtl sets NxdomainRedirectTtl field to given value.

### HasNxdomainRedirectTtl

`func (o *GetGridDnsResponse) HasNxdomainRedirectTtl() bool`

HasNxdomainRedirectTtl returns a boolean if a field has been set.

### GetNxdomainRulesets

`func (o *GetGridDnsResponse) GetNxdomainRulesets() []string`

GetNxdomainRulesets returns the NxdomainRulesets field if non-nil, zero value otherwise.

### GetNxdomainRulesetsOk

`func (o *GetGridDnsResponse) GetNxdomainRulesetsOk() (*[]string, bool)`

GetNxdomainRulesetsOk returns a tuple with the NxdomainRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRulesets

`func (o *GetGridDnsResponse) SetNxdomainRulesets(v []string)`

SetNxdomainRulesets sets NxdomainRulesets field to given value.

### HasNxdomainRulesets

`func (o *GetGridDnsResponse) HasNxdomainRulesets() bool`

HasNxdomainRulesets returns a boolean if a field has been set.

### GetPreserveHostRrsetOrderOnSecondaries

`func (o *GetGridDnsResponse) GetPreserveHostRrsetOrderOnSecondaries() bool`

GetPreserveHostRrsetOrderOnSecondaries returns the PreserveHostRrsetOrderOnSecondaries field if non-nil, zero value otherwise.

### GetPreserveHostRrsetOrderOnSecondariesOk

`func (o *GetGridDnsResponse) GetPreserveHostRrsetOrderOnSecondariesOk() (*bool, bool)`

GetPreserveHostRrsetOrderOnSecondariesOk returns a tuple with the PreserveHostRrsetOrderOnSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveHostRrsetOrderOnSecondaries

`func (o *GetGridDnsResponse) SetPreserveHostRrsetOrderOnSecondaries(v bool)`

SetPreserveHostRrsetOrderOnSecondaries sets PreserveHostRrsetOrderOnSecondaries field to given value.

### HasPreserveHostRrsetOrderOnSecondaries

`func (o *GetGridDnsResponse) HasPreserveHostRrsetOrderOnSecondaries() bool`

HasPreserveHostRrsetOrderOnSecondaries returns a boolean if a field has been set.

### GetProtocolRecordNamePolicies

`func (o *GetGridDnsResponse) GetProtocolRecordNamePolicies() []string`

GetProtocolRecordNamePolicies returns the ProtocolRecordNamePolicies field if non-nil, zero value otherwise.

### GetProtocolRecordNamePoliciesOk

`func (o *GetGridDnsResponse) GetProtocolRecordNamePoliciesOk() (*[]string, bool)`

GetProtocolRecordNamePoliciesOk returns a tuple with the ProtocolRecordNamePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRecordNamePolicies

`func (o *GetGridDnsResponse) SetProtocolRecordNamePolicies(v []string)`

SetProtocolRecordNamePolicies sets ProtocolRecordNamePolicies field to given value.

### HasProtocolRecordNamePolicies

`func (o *GetGridDnsResponse) HasProtocolRecordNamePolicies() bool`

HasProtocolRecordNamePolicies returns a boolean if a field has been set.

### GetQueryRewriteDomainNames

`func (o *GetGridDnsResponse) GetQueryRewriteDomainNames() []string`

GetQueryRewriteDomainNames returns the QueryRewriteDomainNames field if non-nil, zero value otherwise.

### GetQueryRewriteDomainNamesOk

`func (o *GetGridDnsResponse) GetQueryRewriteDomainNamesOk() (*[]string, bool)`

GetQueryRewriteDomainNamesOk returns a tuple with the QueryRewriteDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRewriteDomainNames

`func (o *GetGridDnsResponse) SetQueryRewriteDomainNames(v []string)`

SetQueryRewriteDomainNames sets QueryRewriteDomainNames field to given value.

### HasQueryRewriteDomainNames

`func (o *GetGridDnsResponse) HasQueryRewriteDomainNames() bool`

HasQueryRewriteDomainNames returns a boolean if a field has been set.

### GetQueryRewritePrefix

`func (o *GetGridDnsResponse) GetQueryRewritePrefix() string`

GetQueryRewritePrefix returns the QueryRewritePrefix field if non-nil, zero value otherwise.

### GetQueryRewritePrefixOk

`func (o *GetGridDnsResponse) GetQueryRewritePrefixOk() (*string, bool)`

GetQueryRewritePrefixOk returns a tuple with the QueryRewritePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRewritePrefix

`func (o *GetGridDnsResponse) SetQueryRewritePrefix(v string)`

SetQueryRewritePrefix sets QueryRewritePrefix field to given value.

### HasQueryRewritePrefix

`func (o *GetGridDnsResponse) HasQueryRewritePrefix() bool`

HasQueryRewritePrefix returns a boolean if a field has been set.

### GetQuerySourcePort

`func (o *GetGridDnsResponse) GetQuerySourcePort() int64`

GetQuerySourcePort returns the QuerySourcePort field if non-nil, zero value otherwise.

### GetQuerySourcePortOk

`func (o *GetGridDnsResponse) GetQuerySourcePortOk() (*int64, bool)`

GetQuerySourcePortOk returns a tuple with the QuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySourcePort

`func (o *GetGridDnsResponse) SetQuerySourcePort(v int64)`

SetQuerySourcePort sets QuerySourcePort field to given value.

### HasQuerySourcePort

`func (o *GetGridDnsResponse) HasQuerySourcePort() bool`

HasQuerySourcePort returns a boolean if a field has been set.

### GetRecursiveQueryList

`func (o *GetGridDnsResponse) GetRecursiveQueryList() []GridDnsRecursiveQueryList`

GetRecursiveQueryList returns the RecursiveQueryList field if non-nil, zero value otherwise.

### GetRecursiveQueryListOk

`func (o *GetGridDnsResponse) GetRecursiveQueryListOk() (*[]GridDnsRecursiveQueryList, bool)`

GetRecursiveQueryListOk returns a tuple with the RecursiveQueryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveQueryList

`func (o *GetGridDnsResponse) SetRecursiveQueryList(v []GridDnsRecursiveQueryList)`

SetRecursiveQueryList sets RecursiveQueryList field to given value.

### HasRecursiveQueryList

`func (o *GetGridDnsResponse) HasRecursiveQueryList() bool`

HasRecursiveQueryList returns a boolean if a field has been set.

### GetRefreshTimer

`func (o *GetGridDnsResponse) GetRefreshTimer() int64`

GetRefreshTimer returns the RefreshTimer field if non-nil, zero value otherwise.

### GetRefreshTimerOk

`func (o *GetGridDnsResponse) GetRefreshTimerOk() (*int64, bool)`

GetRefreshTimerOk returns a tuple with the RefreshTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTimer

`func (o *GetGridDnsResponse) SetRefreshTimer(v int64)`

SetRefreshTimer sets RefreshTimer field to given value.

### HasRefreshTimer

`func (o *GetGridDnsResponse) HasRefreshTimer() bool`

HasRefreshTimer returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *GetGridDnsResponse) GetResolverQueryTimeout() int64`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *GetGridDnsResponse) GetResolverQueryTimeoutOk() (*int64, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *GetGridDnsResponse) SetResolverQueryTimeout(v int64)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *GetGridDnsResponse) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetResponseRateLimiting

`func (o *GetGridDnsResponse) GetResponseRateLimiting() GridDnsResponseRateLimiting`

GetResponseRateLimiting returns the ResponseRateLimiting field if non-nil, zero value otherwise.

### GetResponseRateLimitingOk

`func (o *GetGridDnsResponse) GetResponseRateLimitingOk() (*GridDnsResponseRateLimiting, bool)`

GetResponseRateLimitingOk returns a tuple with the ResponseRateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRateLimiting

`func (o *GetGridDnsResponse) SetResponseRateLimiting(v GridDnsResponseRateLimiting)`

SetResponseRateLimiting sets ResponseRateLimiting field to given value.

### HasResponseRateLimiting

`func (o *GetGridDnsResponse) HasResponseRateLimiting() bool`

HasResponseRateLimiting returns a boolean if a field has been set.

### GetRestartSetting

`func (o *GetGridDnsResponse) GetRestartSetting() GridDnsRestartSetting`

GetRestartSetting returns the RestartSetting field if non-nil, zero value otherwise.

### GetRestartSettingOk

`func (o *GetGridDnsResponse) GetRestartSettingOk() (*GridDnsRestartSetting, bool)`

GetRestartSettingOk returns a tuple with the RestartSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSetting

`func (o *GetGridDnsResponse) SetRestartSetting(v GridDnsRestartSetting)`

SetRestartSetting sets RestartSetting field to given value.

### HasRestartSetting

`func (o *GetGridDnsResponse) HasRestartSetting() bool`

HasRestartSetting returns a boolean if a field has been set.

### GetRetryTimer

`func (o *GetGridDnsResponse) GetRetryTimer() int64`

GetRetryTimer returns the RetryTimer field if non-nil, zero value otherwise.

### GetRetryTimerOk

`func (o *GetGridDnsResponse) GetRetryTimerOk() (*int64, bool)`

GetRetryTimerOk returns a tuple with the RetryTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimer

`func (o *GetGridDnsResponse) SetRetryTimer(v int64)`

SetRetryTimer sets RetryTimer field to given value.

### HasRetryTimer

`func (o *GetGridDnsResponse) HasRetryTimer() bool`

HasRetryTimer returns a boolean if a field has been set.

### GetRootNameServerType

`func (o *GetGridDnsResponse) GetRootNameServerType() string`

GetRootNameServerType returns the RootNameServerType field if non-nil, zero value otherwise.

### GetRootNameServerTypeOk

`func (o *GetGridDnsResponse) GetRootNameServerTypeOk() (*string, bool)`

GetRootNameServerTypeOk returns a tuple with the RootNameServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNameServerType

`func (o *GetGridDnsResponse) SetRootNameServerType(v string)`

SetRootNameServerType sets RootNameServerType field to given value.

### HasRootNameServerType

`func (o *GetGridDnsResponse) HasRootNameServerType() bool`

HasRootNameServerType returns a boolean if a field has been set.

### GetRpzDisableNsdnameNsip

`func (o *GetGridDnsResponse) GetRpzDisableNsdnameNsip() bool`

GetRpzDisableNsdnameNsip returns the RpzDisableNsdnameNsip field if non-nil, zero value otherwise.

### GetRpzDisableNsdnameNsipOk

`func (o *GetGridDnsResponse) GetRpzDisableNsdnameNsipOk() (*bool, bool)`

GetRpzDisableNsdnameNsipOk returns a tuple with the RpzDisableNsdnameNsip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDisableNsdnameNsip

`func (o *GetGridDnsResponse) SetRpzDisableNsdnameNsip(v bool)`

SetRpzDisableNsdnameNsip sets RpzDisableNsdnameNsip field to given value.

### HasRpzDisableNsdnameNsip

`func (o *GetGridDnsResponse) HasRpzDisableNsdnameNsip() bool`

HasRpzDisableNsdnameNsip returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *GetGridDnsResponse) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *GetGridDnsResponse) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *GetGridDnsResponse) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *GetGridDnsResponse) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetGridDnsResponse) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *GetGridDnsResponse) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetGridDnsResponse) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetGridDnsResponse) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetGridDnsResponse) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *GetGridDnsResponse) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetGridDnsResponse) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetGridDnsResponse) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzQnameWaitRecurse

`func (o *GetGridDnsResponse) GetRpzQnameWaitRecurse() bool`

GetRpzQnameWaitRecurse returns the RpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetRpzQnameWaitRecurseOk

`func (o *GetGridDnsResponse) GetRpzQnameWaitRecurseOk() (*bool, bool)`

GetRpzQnameWaitRecurseOk returns a tuple with the RpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzQnameWaitRecurse

`func (o *GetGridDnsResponse) SetRpzQnameWaitRecurse(v bool)`

SetRpzQnameWaitRecurse sets RpzQnameWaitRecurse field to given value.

### HasRpzQnameWaitRecurse

`func (o *GetGridDnsResponse) HasRpzQnameWaitRecurse() bool`

HasRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetRunScavenging

`func (o *GetGridDnsResponse) GetRunScavenging() map[string]interface{}`

GetRunScavenging returns the RunScavenging field if non-nil, zero value otherwise.

### GetRunScavengingOk

`func (o *GetGridDnsResponse) GetRunScavengingOk() (*map[string]interface{}, bool)`

GetRunScavengingOk returns a tuple with the RunScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunScavenging

`func (o *GetGridDnsResponse) SetRunScavenging(v map[string]interface{})`

SetRunScavenging sets RunScavenging field to given value.

### HasRunScavenging

`func (o *GetGridDnsResponse) HasRunScavenging() bool`

HasRunScavenging returns a boolean if a field has been set.

### GetScavengingSettings

`func (o *GetGridDnsResponse) GetScavengingSettings() GridDnsScavengingSettings`

GetScavengingSettings returns the ScavengingSettings field if non-nil, zero value otherwise.

### GetScavengingSettingsOk

`func (o *GetGridDnsResponse) GetScavengingSettingsOk() (*GridDnsScavengingSettings, bool)`

GetScavengingSettingsOk returns a tuple with the ScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSettings

`func (o *GetGridDnsResponse) SetScavengingSettings(v GridDnsScavengingSettings)`

SetScavengingSettings sets ScavengingSettings field to given value.

### HasScavengingSettings

`func (o *GetGridDnsResponse) HasScavengingSettings() bool`

HasScavengingSettings returns a boolean if a field has been set.

### GetSerialQueryRate

`func (o *GetGridDnsResponse) GetSerialQueryRate() int64`

GetSerialQueryRate returns the SerialQueryRate field if non-nil, zero value otherwise.

### GetSerialQueryRateOk

`func (o *GetGridDnsResponse) GetSerialQueryRateOk() (*int64, bool)`

GetSerialQueryRateOk returns a tuple with the SerialQueryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialQueryRate

`func (o *GetGridDnsResponse) SetSerialQueryRate(v int64)`

SetSerialQueryRate sets SerialQueryRate field to given value.

### HasSerialQueryRate

`func (o *GetGridDnsResponse) HasSerialQueryRate() bool`

HasSerialQueryRate returns a boolean if a field has been set.

### GetServerIdDirective

`func (o *GetGridDnsResponse) GetServerIdDirective() string`

GetServerIdDirective returns the ServerIdDirective field if non-nil, zero value otherwise.

### GetServerIdDirectiveOk

`func (o *GetGridDnsResponse) GetServerIdDirectiveOk() (*string, bool)`

GetServerIdDirectiveOk returns a tuple with the ServerIdDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIdDirective

`func (o *GetGridDnsResponse) SetServerIdDirective(v string)`

SetServerIdDirective sets ServerIdDirective field to given value.

### HasServerIdDirective

`func (o *GetGridDnsResponse) HasServerIdDirective() bool`

HasServerIdDirective returns a boolean if a field has been set.

### GetSortlist

`func (o *GetGridDnsResponse) GetSortlist() []GridDnsSortlist`

GetSortlist returns the Sortlist field if non-nil, zero value otherwise.

### GetSortlistOk

`func (o *GetGridDnsResponse) GetSortlistOk() (*[]GridDnsSortlist, bool)`

GetSortlistOk returns a tuple with the Sortlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortlist

`func (o *GetGridDnsResponse) SetSortlist(v []GridDnsSortlist)`

SetSortlist sets Sortlist field to given value.

### HasSortlist

`func (o *GetGridDnsResponse) HasSortlist() bool`

HasSortlist returns a boolean if a field has been set.

### GetStoreLocally

`func (o *GetGridDnsResponse) GetStoreLocally() bool`

GetStoreLocally returns the StoreLocally field if non-nil, zero value otherwise.

### GetStoreLocallyOk

`func (o *GetGridDnsResponse) GetStoreLocallyOk() (*bool, bool)`

GetStoreLocallyOk returns a tuple with the StoreLocally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreLocally

`func (o *GetGridDnsResponse) SetStoreLocally(v bool)`

SetStoreLocally sets StoreLocally field to given value.

### HasStoreLocally

`func (o *GetGridDnsResponse) HasStoreLocally() bool`

HasStoreLocally returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *GetGridDnsResponse) GetSyslogFacility() string`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *GetGridDnsResponse) GetSyslogFacilityOk() (*string, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *GetGridDnsResponse) SetSyslogFacility(v string)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *GetGridDnsResponse) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetTransferExcludedServers

`func (o *GetGridDnsResponse) GetTransferExcludedServers() []string`

GetTransferExcludedServers returns the TransferExcludedServers field if non-nil, zero value otherwise.

### GetTransferExcludedServersOk

`func (o *GetGridDnsResponse) GetTransferExcludedServersOk() (*[]string, bool)`

GetTransferExcludedServersOk returns a tuple with the TransferExcludedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferExcludedServers

`func (o *GetGridDnsResponse) SetTransferExcludedServers(v []string)`

SetTransferExcludedServers sets TransferExcludedServers field to given value.

### HasTransferExcludedServers

`func (o *GetGridDnsResponse) HasTransferExcludedServers() bool`

HasTransferExcludedServers returns a boolean if a field has been set.

### GetTransferFormat

`func (o *GetGridDnsResponse) GetTransferFormat() string`

GetTransferFormat returns the TransferFormat field if non-nil, zero value otherwise.

### GetTransferFormatOk

`func (o *GetGridDnsResponse) GetTransferFormatOk() (*string, bool)`

GetTransferFormatOk returns a tuple with the TransferFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFormat

`func (o *GetGridDnsResponse) SetTransferFormat(v string)`

SetTransferFormat sets TransferFormat field to given value.

### HasTransferFormat

`func (o *GetGridDnsResponse) HasTransferFormat() bool`

HasTransferFormat returns a boolean if a field has been set.

### GetTransfersIn

`func (o *GetGridDnsResponse) GetTransfersIn() int64`

GetTransfersIn returns the TransfersIn field if non-nil, zero value otherwise.

### GetTransfersInOk

`func (o *GetGridDnsResponse) GetTransfersInOk() (*int64, bool)`

GetTransfersInOk returns a tuple with the TransfersIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersIn

`func (o *GetGridDnsResponse) SetTransfersIn(v int64)`

SetTransfersIn sets TransfersIn field to given value.

### HasTransfersIn

`func (o *GetGridDnsResponse) HasTransfersIn() bool`

HasTransfersIn returns a boolean if a field has been set.

### GetTransfersOut

`func (o *GetGridDnsResponse) GetTransfersOut() int64`

GetTransfersOut returns the TransfersOut field if non-nil, zero value otherwise.

### GetTransfersOutOk

`func (o *GetGridDnsResponse) GetTransfersOutOk() (*int64, bool)`

GetTransfersOutOk returns a tuple with the TransfersOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersOut

`func (o *GetGridDnsResponse) SetTransfersOut(v int64)`

SetTransfersOut sets TransfersOut field to given value.

### HasTransfersOut

`func (o *GetGridDnsResponse) HasTransfersOut() bool`

HasTransfersOut returns a boolean if a field has been set.

### GetTransfersPerNs

`func (o *GetGridDnsResponse) GetTransfersPerNs() int64`

GetTransfersPerNs returns the TransfersPerNs field if non-nil, zero value otherwise.

### GetTransfersPerNsOk

`func (o *GetGridDnsResponse) GetTransfersPerNsOk() (*int64, bool)`

GetTransfersPerNsOk returns a tuple with the TransfersPerNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersPerNs

`func (o *GetGridDnsResponse) SetTransfersPerNs(v int64)`

SetTransfersPerNs sets TransfersPerNs field to given value.

### HasTransfersPerNs

`func (o *GetGridDnsResponse) HasTransfersPerNs() bool`

HasTransfersPerNs returns a boolean if a field has been set.

### GetZoneDeletionDoubleConfirm

`func (o *GetGridDnsResponse) GetZoneDeletionDoubleConfirm() bool`

GetZoneDeletionDoubleConfirm returns the ZoneDeletionDoubleConfirm field if non-nil, zero value otherwise.

### GetZoneDeletionDoubleConfirmOk

`func (o *GetGridDnsResponse) GetZoneDeletionDoubleConfirmOk() (*bool, bool)`

GetZoneDeletionDoubleConfirmOk returns a tuple with the ZoneDeletionDoubleConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDeletionDoubleConfirm

`func (o *GetGridDnsResponse) SetZoneDeletionDoubleConfirm(v bool)`

SetZoneDeletionDoubleConfirm sets ZoneDeletionDoubleConfirm field to given value.

### HasZoneDeletionDoubleConfirm

`func (o *GetGridDnsResponse) HasZoneDeletionDoubleConfirm() bool`

HasZoneDeletionDoubleConfirm returns a boolean if a field has been set.

### GetResult

`func (o *GetGridDnsResponse) GetResult() GridDns`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridDnsResponse) GetResultOk() (*GridDns, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridDnsResponse) SetResult(v GridDns)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridDnsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


