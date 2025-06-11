# MemberDhcpproperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AuthServerGroup** | Pointer to **string** | The Authentication Server Group object associated with this member. | [optional] 
**AuthnCaptivePortal** | Pointer to **string** | The captive portal responsible for authenticating this DHCP member. | [optional] 
**AuthnCaptivePortalAuthenticatedFilter** | Pointer to **string** | The MAC filter representing the authenticated range. | [optional] 
**AuthnCaptivePortalEnabled** | Pointer to **bool** | The flag that controls if this DHCP member is enabled for captive portal authentication. | [optional] 
**AuthnCaptivePortalGuestFilter** | Pointer to **string** | The MAC filter representing the guest range. | [optional] 
**AuthnServerGroupEnabled** | Pointer to **bool** | The flag that controls if this DHCP member can send authentication requests to an authentication server group. | [optional] 
**Authority** | Pointer to **bool** | The authority flag of a Grid member. This flag specifies if a DHCP server is authoritative for a domain. | [optional] 
**Bootfile** | Pointer to **string** | The name of a file that DHCP clients need to boot. This setting overrides the Grid level setting. | [optional] 
**Bootserver** | Pointer to **string** | The name of the server on which a boot file is stored. This setting overrides the Grid level setting. | [optional] 
**DdnsDomainname** | Pointer to **string** | The member DDNS domain name value. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | Determines the ability of a member DHCP server to generate a host name and update DNS with this host name when it receives a DHCP REQUEST message that does not include a host name. | [optional] 
**DdnsRetryInterval** | Pointer to **int64** | Determines the retry interval when the member DHCP server makes repeated attempts to send DDNS updates to a DNS server. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | Determines that only the DHCP server is allowed to update DNS, regardless of the requests from the DHCP clients. This setting overrides the Grid level setting. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DDNS TTL (Dynamic DNS Time To Live) value specifies the number of seconds an IP address for the name is cached. | [optional] 
**DdnsUpdateFixedAddresses** | Pointer to **bool** | Determines if the member DHCP server&#39;s ability to update the A and PTR records with a fixed address is enabled or not. | [optional] 
**DdnsUseOption81** | Pointer to **bool** | Determines if support for option 81 is enabled or not. | [optional] 
**DdnsZonePrimaries** | Pointer to [**[]MemberDhcppropertiesDdnsZonePrimaries**](MemberDhcppropertiesDdnsZonePrimaries.md) | An ordered list of zone primaries that will receive DDNS updates. | [optional] 
**DenyBootp** | Pointer to **bool** | Determines if a BOOTP server denies BOOTP request or not. This setting overrides the Grid level setting. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of DHCP objects belonging to the Grid Member multiplied by 1000. This is the percentage of the total number of available IP addresses from all the DHCP objects belonging to the Grid Member versus the total number of all IP addresses in all of the DHCP objects on the Grid Member. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of DHCP objects that belong to the Grid Member. | [optional] [readonly] 
**DnsUpdateStyle** | Pointer to **string** | The update style for dynamic DNS updates. | [optional] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the DHCP objects on the Grid Member. | [optional] [readonly] 
**EmailList** | Pointer to **[]string** | The email_list value of a member DHCP server. | [optional] 
**EnableDdns** | Pointer to **bool** | Determines if the member DHCP server&#39;s ability to send DDNS updates is enabled or not. | [optional] 
**EnableDhcp** | Pointer to **bool** | Determines if the DHCP service of a member is enabled or not. | [optional] 
**EnableDhcpOnIpv6Lan2** | Pointer to **bool** | Determines if the DHCP service on the IPv6 LAN2 interface is enabled or not. | [optional] 
**EnableDhcpOnLan2** | Pointer to **bool** | Determines if the DHCP service on the LAN2 interface is enabled or not. | [optional] 
**EnableDhcpThresholds** | Pointer to **bool** | Represents the watermarks above or below which address usage in a network is unexpected and might warrant your attention. This setting overrides the Grid level setting. | [optional] 
**EnableDhcpv6Service** | Pointer to **bool** | Determines if DHCPv6 service for the member is enabled or not. | [optional] 
**EnableEmailWarnings** | Pointer to **bool** | Determines if e-mail warnings are enabled or disabled. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an e-mail notification to an administrator. | [optional] 
**EnableFingerprint** | Pointer to **bool** | Determines if fingerprint feature is enabled on this member. If you enable this feature, the server will match a fingerprint for incoming lease requests. | [optional] 
**EnableGssTsig** | Pointer to **bool** | Determines whether the appliance is enabled to receive GSS-TSIG authenticated updates from DHCP clients. | [optional] 
**EnableHostnameRewrite** | Pointer to **bool** | Determines if the Grid member&#39;s host name rewrite feature is enabled or not. | [optional] 
**EnableLeasequery** | Pointer to **bool** | Determines if lease query is allowed or not. This setting overrides the Grid-level setting. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determines if SNMP warnings are enabled or disabled on this DHCP member. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an SNMP trap to the trap receiver that was defined for the Grid member level. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**GssTsigKeys** | Pointer to **[]string** | The list of GSS-TSIG keys for a member DHCP object. | [optional] 
**HighWaterMark** | Pointer to **int64** | Determines the high watermark value of a member DHCP server. If the percentage of allocated addresses exceeds this watermark, the appliance makes a syslog entry and sends an e-mail notification (if enabled). Specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | Determines the high watermark reset value of a member DHCP server. If the percentage of allocated addresses drops below this value, a corresponding SNMP trap is reset. Specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**HostName** | Pointer to **string** | Host name of the Grid member. | [optional] [readonly] 
**HostnameRewritePolicy** | Pointer to **string** | The hostname rewrite policy that is in the protocol hostname rewrite policies array of the Grid DHCP object. This attribute is mandatory if enable_hostname_rewrite is \&quot;true\&quot;. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | Determines if the ignore DHCP option list request flag of a Grid member DHCP is enabled or not. If this flag is set to true all available DHCP options will be returned to the client. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \&quot;NONE\&quot;, \&quot;CLIENT\&quot;, or \&quot;MACADDR\&quot;. The default is \&quot;NONE\&quot;. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**ImmediateFaConfiguration** | Pointer to **bool** | Determines if the Immediate Fixed address configuration apply feature for the DHCP member is enabled or not. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the Grid member. | [optional] [readonly] 
**Ipv6DdnsDomainname** | Pointer to **string** | The member DDNS IPv6 domain name value. | [optional] 
**Ipv6DdnsEnableOptionFqdn** | Pointer to **bool** | Controls whether the FQDN option sent by the DHCPv6 client is to be used, or if the server can automatically generate the FQDN. | [optional] 
**Ipv6DdnsHostname** | Pointer to **string** | The member IPv6 DDNS hostname value. | [optional] 
**Ipv6DdnsServerAlwaysUpdates** | Pointer to **bool** | Determines if the server always updates DNS or updates only if requested by the client. | [optional] 
**Ipv6DdnsTtl** | Pointer to **int64** | The member IPv6 DDNS TTL value. | [optional] 
**Ipv6DnsUpdateStyle** | Pointer to **string** | The update style for dynamic DHCPv6 DNS updates. | [optional] 
**Ipv6DomainName** | Pointer to **string** | The IPv6 domain name. | [optional] 
**Ipv6DomainNameServers** | Pointer to **[]string** | The comma separated list of domain name server addresses in IPv6 address format. | [optional] 
**Ipv6EnableDdns** | Pointer to **bool** | Determines if sending DDNS updates by the member DHCPv6 server is enabled or not. | [optional] 
**Ipv6EnableGssTsig** | Pointer to **bool** | Determines whether the appliance is enabled to receive GSS-TSIG authenticated updates from DHCPv6 clients. | [optional] 
**Ipv6EnableLeaseScavenging** | Pointer to **bool** | Indicates whether DHCPv6 lease scavenging is enabled or disabled. | [optional] 
**Ipv6EnableRetryUpdates** | Pointer to **bool** | Determines if the DHCPv6 server retries failed dynamic DNS updates or not. | [optional] 
**Ipv6GenerateHostname** | Pointer to **bool** | Determines if the server generates the hostname if it is not sent by the client. | [optional] 
**Ipv6GssTsigKeys** | Pointer to **[]string** | The list of GSS-TSIG keys for a member DHCPv6 object. | [optional] 
**Ipv6KdcServer** | Pointer to **string** | Determines the IPv6 address or FQDN of the Kerberos server for DHCPv6 GSS-TSIG authentication. This setting overrides the Grid level setting. | [optional] 
**Ipv6LeaseScavengingTime** | Pointer to **int64** | The member-level grace period (in seconds) to keep an expired lease before it is deleted by the scavenging process. | [optional] 
**Ipv6MicrosoftCodePage** | Pointer to **string** | The Microsoft client DHCP IPv6 code page value of a Grid member. This value is the hostname translation code page for Microsoft DHCP IPv6 clients and overrides the Grid level Microsoft DHCP IPv6 client code page. | [optional] 
**Ipv6Options** | Pointer to [**[]MemberDhcppropertiesIpv6Options**](MemberDhcppropertiesIpv6Options.md) | An array of DHCP option dhcpoption structs that lists the DHCPv6 options associated with the object. | [optional] 
**Ipv6RecycleLeases** | Pointer to **bool** | Determines if the IPv6 recycle leases feature is enabled or not. If the feature is enabled, leases are kept in the Recycle Bin until one week after lease expiration. When the feature is disabled, the leases are irrecoverably deleted. | [optional] 
**Ipv6RememberExpiredClientAssociation** | Pointer to **bool** | Enable binding for expired DHCPv6 leases. | [optional] 
**Ipv6RetryUpdatesInterval** | Pointer to **int64** | Determines the retry interval when the member DHCPv6 server makes repeated attempts to send DDNS updates to a DNS server. | [optional] 
**Ipv6ServerDuid** | Pointer to **string** | The server DHCPv6 unique identifier (DUID) for the Grid member. | [optional] 
**Ipv6UpdateDnsOnLeaseRenewal** | Pointer to **bool** | Controls whether the DHCPv6 server updates DNS when an IPv6 DHCP lease is renewed. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Grid member. | [optional] [readonly] 
**KdcServer** | Pointer to **string** | The IPv4 address or FQDN of the Kerberos server for DHCPv4 GSS-TSIG authentication. This setting overrides the Grid level setting. | [optional] 
**LeasePerClientSettings** | Pointer to **string** | Defines how the appliance releases DHCP leases. Valid values are \&quot;RELEASE_MACHING_ID\&quot;, \&quot;NEVER_RELEASE\&quot;, or \&quot;ONE_LEASE_PER_CLIENT\&quot;. The default is \&quot;RELEASE_MATCHING_ID\&quot;. | [optional] 
**LeaseScavengeTime** | Pointer to **int32** | Determines the lease scavenging time value. When this field is set, the appliance permanently deletes the free and backup leases that remain in the database beyond a specified period of time. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogLeaseEvents** | Pointer to **bool** | This value specifies whether the grid member logs lease events. This setting overrides the Grid level setting. | [optional] 
**LogicFilterRules** | Pointer to [**[]MemberDhcppropertiesLogicFilterRules**](MemberDhcppropertiesLogicFilterRules.md) | This field contains the logic filters to be applied on the Grid member. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**LowWaterMark** | Pointer to **int64** | Determines the low watermark value. If the percent of allocated addresses drops below this watermark, the appliance makes a syslog entry and sends an e-mail notification (if enabled). | [optional] 
**LowWaterMarkReset** | Pointer to **int64** | Determines the low watermark reset value. If the percentage of allocated addresses exceeds this value, a corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value. | [optional] 
**MicrosoftCodePage** | Pointer to **string** | The Microsoft client DHCP IPv4 code page value of a grid member. This value is the hostname translation code page for Microsoft DHCP IPv4 clients and overrides the Grid level Microsoft DHCP IPv4 client code page. | [optional] 
**Nextserver** | Pointer to **string** | The next server value of a member DHCP server. This value is the IP address or name of the boot file server on which the boot file is stored. | [optional] 
**Option60MatchRules** | Pointer to [**[]MemberDhcppropertiesOption60MatchRules**](MemberDhcppropertiesOption60MatchRules.md) | The list of option 60 match rules. | [optional] 
**Options** | Pointer to [**[]MemberDhcppropertiesOptions**](MemberDhcppropertiesOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PingCount** | Pointer to **int64** | Specifies the number of pings that the Infoblox appliance sends to an IP address to verify that it is not in use. Values are from 0 to 10, where 0 disables pings. | [optional] 
**PingTimeout** | Pointer to **int64** | Indicates the number of milliseconds the appliance waits for a response to its ping. Valid values are 100, 500, 1000, 2000, 3000, 4000 and 5000 milliseconds. | [optional] 
**PreferredLifetime** | Pointer to **int64** | The preferred lifetime value. | [optional] 
**PrefixLengthMode** | Pointer to **string** | The Prefix length mode for DHCPv6. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | Specifies the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**RecycleLeases** | Pointer to **bool** | Determines if the recycle leases feature is enabled or not. If you enabled this feature and then delete a DHCP range, the appliance stores active leases from this range up to one week after the leases expires. | [optional] 
**RetryDdnsUpdates** | Pointer to **bool** | Indicates whether the DHCP server makes repeated attempts to send DDNS updates to a DNS server. | [optional] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in DHCP objects that belong to the Grid Member. | [optional] [readonly] 
**SyslogFacility** | Pointer to **string** | The syslog facility is the location on the syslog server to which you want to sort the syslog messages. | [optional] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in DHCP objects that belong to the Grid Member. | [optional] [readonly] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | Controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseAuthority** | Pointer to **bool** | Use flag for: authority | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDdnsUpdateFixedAddresses** | Pointer to **bool** | Use flag for: ddns_update_fixed_addresses | [optional] 
**UseDdnsUseOption81** | Pointer to **bool** | Use flag for: ddns_use_option81 | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseDnsUpdateStyle** | Pointer to **bool** | Use flag for: dns_update_style | [optional] 
**UseEmailList** | Pointer to **bool** | Use flag for: email_list | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseEnableDhcpThresholds** | Pointer to **bool** | Use flag for: enable_dhcp_thresholds , high_water_mark, high_water_mark_reset, low_water_mark, low_water_mark_reset | [optional] 
**UseEnableFingerprint** | Pointer to **bool** | Use flag for: enable_fingerprint | [optional] 
**UseEnableGssTsig** | Pointer to **bool** | Use flag for: kdc_server , enable_gss_tsig | [optional] 
**UseEnableHostnameRewrite** | Pointer to **bool** | Use flag for: enable_hostname_rewrite , hostname_rewrite_policy | [optional] 
**UseEnableLeasequery** | Pointer to **bool** | Use flag for: enable_leasequery | [optional] 
**UseEnableOneLeasePerClient** | Pointer to **bool** | Use flag for: enable_one_lease_per_client | [optional] 
**UseGssTsigKeys** | Pointer to **bool** | Use flag for: gss_tsig_keys | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseIgnoreId** | Pointer to **bool** | Use flag for: ignore_id | [optional] 
**UseImmediateFaConfiguration** | Pointer to **bool** | Use flag for: immediate_fa_configuration | [optional] 
**UseIpv6DdnsDomainname** | Pointer to **bool** | Use flag for: ipv6_ddns_domainname | [optional] 
**UseIpv6DdnsEnableOptionFqdn** | Pointer to **bool** | Use flag for: ipv6_ddns_enable_option_fqdn | [optional] 
**UseIpv6DdnsHostname** | Pointer to **bool** | Use flag for: ipv6_ddns_hostname | [optional] 
**UseIpv6DdnsTtl** | Pointer to **bool** | Use flag for: ipv6_ddns_ttl | [optional] 
**UseIpv6DnsUpdateStyle** | Pointer to **bool** | Use flag for: ipv6_dns_update_style | [optional] 
**UseIpv6DomainName** | Pointer to **bool** | Use flag for: ipv6_domain_name | [optional] 
**UseIpv6DomainNameServers** | Pointer to **bool** | Use flag for: ipv6_domain_name_servers | [optional] 
**UseIpv6EnableDdns** | Pointer to **bool** | Use flag for: ipv6_enable_ddns | [optional] 
**UseIpv6EnableGssTsig** | Pointer to **bool** | Use flag for: ipv6_kdc_server , ipv6_enable_gss_tsig | [optional] 
**UseIpv6EnableRetryUpdates** | Pointer to **bool** | Use flag for: ipv6_enable_retry_updates , ipv6_retry_updates_interval | [optional] 
**UseIpv6GenerateHostname** | Pointer to **bool** | Use flag for: ipv6_generate_hostname | [optional] 
**UseIpv6GssTsigKeys** | Pointer to **bool** | Use flag for: ipv6_gss_tsig_keys | [optional] 
**UseIpv6LeaseScavenging** | Pointer to **bool** | Use flag for: ipv6_enable_lease_scavenging , ipv6_lease_scavenging_time, ipv6_remember_expired_client_association | [optional] 
**UseIpv6MicrosoftCodePage** | Pointer to **bool** | Use flag for: ipv6_microsoft_code_page | [optional] 
**UseIpv6Options** | Pointer to **bool** | Use flag for: ipv6_options | [optional] 
**UseIpv6RecycleLeases** | Pointer to **bool** | Use flag for: ipv6_recycle_leases | [optional] 
**UseIpv6UpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: ipv6_update_dns_on_lease_renewal | [optional] 
**UseLeasePerClientSettings** | Pointer to **bool** | Use flag for: lease_per_client_settings | [optional] 
**UseLeaseScavengeTime** | Pointer to **bool** | Use flag for: lease_scavenge_time | [optional] 
**UseLogLeaseEvents** | Pointer to **bool** | Use flag for: log_lease_events | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMicrosoftCodePage** | Pointer to **bool** | Use flag for: microsoft_code_page | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePingCount** | Pointer to **bool** | Use flag for: ping_count | [optional] 
**UsePingTimeout** | Pointer to **bool** | Use flag for: ping_timeout | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UsePrefixLengthMode** | Pointer to **bool** | Use flag for: prefix_length_mode | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseRetryDdnsUpdates** | Pointer to **bool** | Use flag for: ddns_retry_interval , retry_ddns_updates | [optional] 
**UseSyslogFacility** | Pointer to **bool** | Use flag for: syslog_facility | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**ValidLifetime** | Pointer to **int64** | The valid lifetime for Grid Member DHCP. Specifies the length of time addresses that are assigned to DHCPv6 clients remain in the valid state. | [optional] 

## Methods

### NewMemberDhcpproperties

`func NewMemberDhcpproperties() *MemberDhcpproperties`

NewMemberDhcpproperties instantiates a new MemberDhcpproperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDhcppropertiesWithDefaults

`func NewMemberDhcppropertiesWithDefaults() *MemberDhcpproperties`

NewMemberDhcppropertiesWithDefaults instantiates a new MemberDhcpproperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberDhcpproperties) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberDhcpproperties) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberDhcpproperties) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberDhcpproperties) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthServerGroup

`func (o *MemberDhcpproperties) GetAuthServerGroup() string`

GetAuthServerGroup returns the AuthServerGroup field if non-nil, zero value otherwise.

### GetAuthServerGroupOk

`func (o *MemberDhcpproperties) GetAuthServerGroupOk() (*string, bool)`

GetAuthServerGroupOk returns a tuple with the AuthServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerGroup

`func (o *MemberDhcpproperties) SetAuthServerGroup(v string)`

SetAuthServerGroup sets AuthServerGroup field to given value.

### HasAuthServerGroup

`func (o *MemberDhcpproperties) HasAuthServerGroup() bool`

HasAuthServerGroup returns a boolean if a field has been set.

### GetAuthnCaptivePortal

`func (o *MemberDhcpproperties) GetAuthnCaptivePortal() string`

GetAuthnCaptivePortal returns the AuthnCaptivePortal field if non-nil, zero value otherwise.

### GetAuthnCaptivePortalOk

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalOk() (*string, bool)`

GetAuthnCaptivePortalOk returns a tuple with the AuthnCaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCaptivePortal

`func (o *MemberDhcpproperties) SetAuthnCaptivePortal(v string)`

SetAuthnCaptivePortal sets AuthnCaptivePortal field to given value.

### HasAuthnCaptivePortal

`func (o *MemberDhcpproperties) HasAuthnCaptivePortal() bool`

HasAuthnCaptivePortal returns a boolean if a field has been set.

### GetAuthnCaptivePortalAuthenticatedFilter

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalAuthenticatedFilter() string`

GetAuthnCaptivePortalAuthenticatedFilter returns the AuthnCaptivePortalAuthenticatedFilter field if non-nil, zero value otherwise.

### GetAuthnCaptivePortalAuthenticatedFilterOk

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalAuthenticatedFilterOk() (*string, bool)`

GetAuthnCaptivePortalAuthenticatedFilterOk returns a tuple with the AuthnCaptivePortalAuthenticatedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCaptivePortalAuthenticatedFilter

`func (o *MemberDhcpproperties) SetAuthnCaptivePortalAuthenticatedFilter(v string)`

SetAuthnCaptivePortalAuthenticatedFilter sets AuthnCaptivePortalAuthenticatedFilter field to given value.

### HasAuthnCaptivePortalAuthenticatedFilter

`func (o *MemberDhcpproperties) HasAuthnCaptivePortalAuthenticatedFilter() bool`

HasAuthnCaptivePortalAuthenticatedFilter returns a boolean if a field has been set.

### GetAuthnCaptivePortalEnabled

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalEnabled() bool`

GetAuthnCaptivePortalEnabled returns the AuthnCaptivePortalEnabled field if non-nil, zero value otherwise.

### GetAuthnCaptivePortalEnabledOk

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalEnabledOk() (*bool, bool)`

GetAuthnCaptivePortalEnabledOk returns a tuple with the AuthnCaptivePortalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCaptivePortalEnabled

`func (o *MemberDhcpproperties) SetAuthnCaptivePortalEnabled(v bool)`

SetAuthnCaptivePortalEnabled sets AuthnCaptivePortalEnabled field to given value.

### HasAuthnCaptivePortalEnabled

`func (o *MemberDhcpproperties) HasAuthnCaptivePortalEnabled() bool`

HasAuthnCaptivePortalEnabled returns a boolean if a field has been set.

### GetAuthnCaptivePortalGuestFilter

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalGuestFilter() string`

GetAuthnCaptivePortalGuestFilter returns the AuthnCaptivePortalGuestFilter field if non-nil, zero value otherwise.

### GetAuthnCaptivePortalGuestFilterOk

`func (o *MemberDhcpproperties) GetAuthnCaptivePortalGuestFilterOk() (*string, bool)`

GetAuthnCaptivePortalGuestFilterOk returns a tuple with the AuthnCaptivePortalGuestFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCaptivePortalGuestFilter

`func (o *MemberDhcpproperties) SetAuthnCaptivePortalGuestFilter(v string)`

SetAuthnCaptivePortalGuestFilter sets AuthnCaptivePortalGuestFilter field to given value.

### HasAuthnCaptivePortalGuestFilter

`func (o *MemberDhcpproperties) HasAuthnCaptivePortalGuestFilter() bool`

HasAuthnCaptivePortalGuestFilter returns a boolean if a field has been set.

### GetAuthnServerGroupEnabled

`func (o *MemberDhcpproperties) GetAuthnServerGroupEnabled() bool`

GetAuthnServerGroupEnabled returns the AuthnServerGroupEnabled field if non-nil, zero value otherwise.

### GetAuthnServerGroupEnabledOk

`func (o *MemberDhcpproperties) GetAuthnServerGroupEnabledOk() (*bool, bool)`

GetAuthnServerGroupEnabledOk returns a tuple with the AuthnServerGroupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnServerGroupEnabled

`func (o *MemberDhcpproperties) SetAuthnServerGroupEnabled(v bool)`

SetAuthnServerGroupEnabled sets AuthnServerGroupEnabled field to given value.

### HasAuthnServerGroupEnabled

`func (o *MemberDhcpproperties) HasAuthnServerGroupEnabled() bool`

HasAuthnServerGroupEnabled returns a boolean if a field has been set.

### GetAuthority

`func (o *MemberDhcpproperties) GetAuthority() bool`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *MemberDhcpproperties) GetAuthorityOk() (*bool, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *MemberDhcpproperties) SetAuthority(v bool)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *MemberDhcpproperties) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetBootfile

`func (o *MemberDhcpproperties) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *MemberDhcpproperties) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *MemberDhcpproperties) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *MemberDhcpproperties) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *MemberDhcpproperties) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *MemberDhcpproperties) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *MemberDhcpproperties) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *MemberDhcpproperties) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *MemberDhcpproperties) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *MemberDhcpproperties) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *MemberDhcpproperties) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *MemberDhcpproperties) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *MemberDhcpproperties) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *MemberDhcpproperties) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *MemberDhcpproperties) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *MemberDhcpproperties) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsRetryInterval

`func (o *MemberDhcpproperties) GetDdnsRetryInterval() int64`

GetDdnsRetryInterval returns the DdnsRetryInterval field if non-nil, zero value otherwise.

### GetDdnsRetryIntervalOk

`func (o *MemberDhcpproperties) GetDdnsRetryIntervalOk() (*int64, bool)`

GetDdnsRetryIntervalOk returns a tuple with the DdnsRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRetryInterval

`func (o *MemberDhcpproperties) SetDdnsRetryInterval(v int64)`

SetDdnsRetryInterval sets DdnsRetryInterval field to given value.

### HasDdnsRetryInterval

`func (o *MemberDhcpproperties) HasDdnsRetryInterval() bool`

HasDdnsRetryInterval returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *MemberDhcpproperties) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *MemberDhcpproperties) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *MemberDhcpproperties) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *MemberDhcpproperties) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *MemberDhcpproperties) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *MemberDhcpproperties) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *MemberDhcpproperties) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *MemberDhcpproperties) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUpdateFixedAddresses

`func (o *MemberDhcpproperties) GetDdnsUpdateFixedAddresses() bool`

GetDdnsUpdateFixedAddresses returns the DdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetDdnsUpdateFixedAddressesOk

`func (o *MemberDhcpproperties) GetDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetDdnsUpdateFixedAddressesOk returns a tuple with the DdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateFixedAddresses

`func (o *MemberDhcpproperties) SetDdnsUpdateFixedAddresses(v bool)`

SetDdnsUpdateFixedAddresses sets DdnsUpdateFixedAddresses field to given value.

### HasDdnsUpdateFixedAddresses

`func (o *MemberDhcpproperties) HasDdnsUpdateFixedAddresses() bool`

HasDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *MemberDhcpproperties) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *MemberDhcpproperties) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *MemberDhcpproperties) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *MemberDhcpproperties) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDdnsZonePrimaries

`func (o *MemberDhcpproperties) GetDdnsZonePrimaries() []MemberDhcppropertiesDdnsZonePrimaries`

GetDdnsZonePrimaries returns the DdnsZonePrimaries field if non-nil, zero value otherwise.

### GetDdnsZonePrimariesOk

`func (o *MemberDhcpproperties) GetDdnsZonePrimariesOk() (*[]MemberDhcppropertiesDdnsZonePrimaries, bool)`

GetDdnsZonePrimariesOk returns a tuple with the DdnsZonePrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsZonePrimaries

`func (o *MemberDhcpproperties) SetDdnsZonePrimaries(v []MemberDhcppropertiesDdnsZonePrimaries)`

SetDdnsZonePrimaries sets DdnsZonePrimaries field to given value.

### HasDdnsZonePrimaries

`func (o *MemberDhcpproperties) HasDdnsZonePrimaries() bool`

HasDdnsZonePrimaries returns a boolean if a field has been set.

### GetDenyBootp

`func (o *MemberDhcpproperties) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *MemberDhcpproperties) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *MemberDhcpproperties) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *MemberDhcpproperties) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *MemberDhcpproperties) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *MemberDhcpproperties) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *MemberDhcpproperties) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *MemberDhcpproperties) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *MemberDhcpproperties) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *MemberDhcpproperties) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *MemberDhcpproperties) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *MemberDhcpproperties) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDnsUpdateStyle

`func (o *MemberDhcpproperties) GetDnsUpdateStyle() string`

GetDnsUpdateStyle returns the DnsUpdateStyle field if non-nil, zero value otherwise.

### GetDnsUpdateStyleOk

`func (o *MemberDhcpproperties) GetDnsUpdateStyleOk() (*string, bool)`

GetDnsUpdateStyleOk returns a tuple with the DnsUpdateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsUpdateStyle

`func (o *MemberDhcpproperties) SetDnsUpdateStyle(v string)`

SetDnsUpdateStyle sets DnsUpdateStyle field to given value.

### HasDnsUpdateStyle

`func (o *MemberDhcpproperties) HasDnsUpdateStyle() bool`

HasDnsUpdateStyle returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *MemberDhcpproperties) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *MemberDhcpproperties) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *MemberDhcpproperties) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *MemberDhcpproperties) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetEmailList

`func (o *MemberDhcpproperties) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *MemberDhcpproperties) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *MemberDhcpproperties) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *MemberDhcpproperties) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *MemberDhcpproperties) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *MemberDhcpproperties) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *MemberDhcpproperties) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *MemberDhcpproperties) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcp

`func (o *MemberDhcpproperties) GetEnableDhcp() bool`

GetEnableDhcp returns the EnableDhcp field if non-nil, zero value otherwise.

### GetEnableDhcpOk

`func (o *MemberDhcpproperties) GetEnableDhcpOk() (*bool, bool)`

GetEnableDhcpOk returns a tuple with the EnableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcp

`func (o *MemberDhcpproperties) SetEnableDhcp(v bool)`

SetEnableDhcp sets EnableDhcp field to given value.

### HasEnableDhcp

`func (o *MemberDhcpproperties) HasEnableDhcp() bool`

HasEnableDhcp returns a boolean if a field has been set.

### GetEnableDhcpOnIpv6Lan2

`func (o *MemberDhcpproperties) GetEnableDhcpOnIpv6Lan2() bool`

GetEnableDhcpOnIpv6Lan2 returns the EnableDhcpOnIpv6Lan2 field if non-nil, zero value otherwise.

### GetEnableDhcpOnIpv6Lan2Ok

`func (o *MemberDhcpproperties) GetEnableDhcpOnIpv6Lan2Ok() (*bool, bool)`

GetEnableDhcpOnIpv6Lan2Ok returns a tuple with the EnableDhcpOnIpv6Lan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpOnIpv6Lan2

`func (o *MemberDhcpproperties) SetEnableDhcpOnIpv6Lan2(v bool)`

SetEnableDhcpOnIpv6Lan2 sets EnableDhcpOnIpv6Lan2 field to given value.

### HasEnableDhcpOnIpv6Lan2

`func (o *MemberDhcpproperties) HasEnableDhcpOnIpv6Lan2() bool`

HasEnableDhcpOnIpv6Lan2 returns a boolean if a field has been set.

### GetEnableDhcpOnLan2

`func (o *MemberDhcpproperties) GetEnableDhcpOnLan2() bool`

GetEnableDhcpOnLan2 returns the EnableDhcpOnLan2 field if non-nil, zero value otherwise.

### GetEnableDhcpOnLan2Ok

`func (o *MemberDhcpproperties) GetEnableDhcpOnLan2Ok() (*bool, bool)`

GetEnableDhcpOnLan2Ok returns a tuple with the EnableDhcpOnLan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpOnLan2

`func (o *MemberDhcpproperties) SetEnableDhcpOnLan2(v bool)`

SetEnableDhcpOnLan2 sets EnableDhcpOnLan2 field to given value.

### HasEnableDhcpOnLan2

`func (o *MemberDhcpproperties) HasEnableDhcpOnLan2() bool`

HasEnableDhcpOnLan2 returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *MemberDhcpproperties) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *MemberDhcpproperties) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *MemberDhcpproperties) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *MemberDhcpproperties) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableDhcpv6Service

`func (o *MemberDhcpproperties) GetEnableDhcpv6Service() bool`

GetEnableDhcpv6Service returns the EnableDhcpv6Service field if non-nil, zero value otherwise.

### GetEnableDhcpv6ServiceOk

`func (o *MemberDhcpproperties) GetEnableDhcpv6ServiceOk() (*bool, bool)`

GetEnableDhcpv6ServiceOk returns a tuple with the EnableDhcpv6Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpv6Service

`func (o *MemberDhcpproperties) SetEnableDhcpv6Service(v bool)`

SetEnableDhcpv6Service sets EnableDhcpv6Service field to given value.

### HasEnableDhcpv6Service

`func (o *MemberDhcpproperties) HasEnableDhcpv6Service() bool`

HasEnableDhcpv6Service returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *MemberDhcpproperties) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *MemberDhcpproperties) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *MemberDhcpproperties) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *MemberDhcpproperties) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableFingerprint

`func (o *MemberDhcpproperties) GetEnableFingerprint() bool`

GetEnableFingerprint returns the EnableFingerprint field if non-nil, zero value otherwise.

### GetEnableFingerprintOk

`func (o *MemberDhcpproperties) GetEnableFingerprintOk() (*bool, bool)`

GetEnableFingerprintOk returns a tuple with the EnableFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFingerprint

`func (o *MemberDhcpproperties) SetEnableFingerprint(v bool)`

SetEnableFingerprint sets EnableFingerprint field to given value.

### HasEnableFingerprint

`func (o *MemberDhcpproperties) HasEnableFingerprint() bool`

HasEnableFingerprint returns a boolean if a field has been set.

### GetEnableGssTsig

`func (o *MemberDhcpproperties) GetEnableGssTsig() bool`

GetEnableGssTsig returns the EnableGssTsig field if non-nil, zero value otherwise.

### GetEnableGssTsigOk

`func (o *MemberDhcpproperties) GetEnableGssTsigOk() (*bool, bool)`

GetEnableGssTsigOk returns a tuple with the EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGssTsig

`func (o *MemberDhcpproperties) SetEnableGssTsig(v bool)`

SetEnableGssTsig sets EnableGssTsig field to given value.

### HasEnableGssTsig

`func (o *MemberDhcpproperties) HasEnableGssTsig() bool`

HasEnableGssTsig returns a boolean if a field has been set.

### GetEnableHostnameRewrite

`func (o *MemberDhcpproperties) GetEnableHostnameRewrite() bool`

GetEnableHostnameRewrite returns the EnableHostnameRewrite field if non-nil, zero value otherwise.

### GetEnableHostnameRewriteOk

`func (o *MemberDhcpproperties) GetEnableHostnameRewriteOk() (*bool, bool)`

GetEnableHostnameRewriteOk returns a tuple with the EnableHostnameRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHostnameRewrite

`func (o *MemberDhcpproperties) SetEnableHostnameRewrite(v bool)`

SetEnableHostnameRewrite sets EnableHostnameRewrite field to given value.

### HasEnableHostnameRewrite

`func (o *MemberDhcpproperties) HasEnableHostnameRewrite() bool`

HasEnableHostnameRewrite returns a boolean if a field has been set.

### GetEnableLeasequery

`func (o *MemberDhcpproperties) GetEnableLeasequery() bool`

GetEnableLeasequery returns the EnableLeasequery field if non-nil, zero value otherwise.

### GetEnableLeasequeryOk

`func (o *MemberDhcpproperties) GetEnableLeasequeryOk() (*bool, bool)`

GetEnableLeasequeryOk returns a tuple with the EnableLeasequery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLeasequery

`func (o *MemberDhcpproperties) SetEnableLeasequery(v bool)`

SetEnableLeasequery sets EnableLeasequery field to given value.

### HasEnableLeasequery

`func (o *MemberDhcpproperties) HasEnableLeasequery() bool`

HasEnableLeasequery returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *MemberDhcpproperties) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *MemberDhcpproperties) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *MemberDhcpproperties) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *MemberDhcpproperties) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetExtattrs

`func (o *MemberDhcpproperties) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *MemberDhcpproperties) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *MemberDhcpproperties) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *MemberDhcpproperties) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetGssTsigKeys

`func (o *MemberDhcpproperties) GetGssTsigKeys() []string`

GetGssTsigKeys returns the GssTsigKeys field if non-nil, zero value otherwise.

### GetGssTsigKeysOk

`func (o *MemberDhcpproperties) GetGssTsigKeysOk() (*[]string, bool)`

GetGssTsigKeysOk returns a tuple with the GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigKeys

`func (o *MemberDhcpproperties) SetGssTsigKeys(v []string)`

SetGssTsigKeys sets GssTsigKeys field to given value.

### HasGssTsigKeys

`func (o *MemberDhcpproperties) HasGssTsigKeys() bool`

HasGssTsigKeys returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *MemberDhcpproperties) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *MemberDhcpproperties) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *MemberDhcpproperties) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *MemberDhcpproperties) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *MemberDhcpproperties) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *MemberDhcpproperties) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *MemberDhcpproperties) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *MemberDhcpproperties) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetHostName

`func (o *MemberDhcpproperties) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *MemberDhcpproperties) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *MemberDhcpproperties) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *MemberDhcpproperties) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostnameRewritePolicy

`func (o *MemberDhcpproperties) GetHostnameRewritePolicy() string`

GetHostnameRewritePolicy returns the HostnameRewritePolicy field if non-nil, zero value otherwise.

### GetHostnameRewritePolicyOk

`func (o *MemberDhcpproperties) GetHostnameRewritePolicyOk() (*string, bool)`

GetHostnameRewritePolicyOk returns a tuple with the HostnameRewritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewritePolicy

`func (o *MemberDhcpproperties) SetHostnameRewritePolicy(v string)`

SetHostnameRewritePolicy sets HostnameRewritePolicy field to given value.

### HasHostnameRewritePolicy

`func (o *MemberDhcpproperties) HasHostnameRewritePolicy() bool`

HasHostnameRewritePolicy returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *MemberDhcpproperties) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *MemberDhcpproperties) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *MemberDhcpproperties) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *MemberDhcpproperties) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *MemberDhcpproperties) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *MemberDhcpproperties) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *MemberDhcpproperties) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *MemberDhcpproperties) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *MemberDhcpproperties) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *MemberDhcpproperties) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *MemberDhcpproperties) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *MemberDhcpproperties) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetImmediateFaConfiguration

`func (o *MemberDhcpproperties) GetImmediateFaConfiguration() bool`

GetImmediateFaConfiguration returns the ImmediateFaConfiguration field if non-nil, zero value otherwise.

### GetImmediateFaConfigurationOk

`func (o *MemberDhcpproperties) GetImmediateFaConfigurationOk() (*bool, bool)`

GetImmediateFaConfigurationOk returns a tuple with the ImmediateFaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFaConfiguration

`func (o *MemberDhcpproperties) SetImmediateFaConfiguration(v bool)`

SetImmediateFaConfiguration sets ImmediateFaConfiguration field to given value.

### HasImmediateFaConfiguration

`func (o *MemberDhcpproperties) HasImmediateFaConfiguration() bool`

HasImmediateFaConfiguration returns a boolean if a field has been set.

### GetIpv4addr

`func (o *MemberDhcpproperties) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *MemberDhcpproperties) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *MemberDhcpproperties) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *MemberDhcpproperties) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6DdnsDomainname

`func (o *MemberDhcpproperties) GetIpv6DdnsDomainname() string`

GetIpv6DdnsDomainname returns the Ipv6DdnsDomainname field if non-nil, zero value otherwise.

### GetIpv6DdnsDomainnameOk

`func (o *MemberDhcpproperties) GetIpv6DdnsDomainnameOk() (*string, bool)`

GetIpv6DdnsDomainnameOk returns a tuple with the Ipv6DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsDomainname

`func (o *MemberDhcpproperties) SetIpv6DdnsDomainname(v string)`

SetIpv6DdnsDomainname sets Ipv6DdnsDomainname field to given value.

### HasIpv6DdnsDomainname

`func (o *MemberDhcpproperties) HasIpv6DdnsDomainname() bool`

HasIpv6DdnsDomainname returns a boolean if a field has been set.

### GetIpv6DdnsEnableOptionFqdn

`func (o *MemberDhcpproperties) GetIpv6DdnsEnableOptionFqdn() bool`

GetIpv6DdnsEnableOptionFqdn returns the Ipv6DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetIpv6DdnsEnableOptionFqdnOk

`func (o *MemberDhcpproperties) GetIpv6DdnsEnableOptionFqdnOk() (*bool, bool)`

GetIpv6DdnsEnableOptionFqdnOk returns a tuple with the Ipv6DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsEnableOptionFqdn

`func (o *MemberDhcpproperties) SetIpv6DdnsEnableOptionFqdn(v bool)`

SetIpv6DdnsEnableOptionFqdn sets Ipv6DdnsEnableOptionFqdn field to given value.

### HasIpv6DdnsEnableOptionFqdn

`func (o *MemberDhcpproperties) HasIpv6DdnsEnableOptionFqdn() bool`

HasIpv6DdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetIpv6DdnsHostname

`func (o *MemberDhcpproperties) GetIpv6DdnsHostname() string`

GetIpv6DdnsHostname returns the Ipv6DdnsHostname field if non-nil, zero value otherwise.

### GetIpv6DdnsHostnameOk

`func (o *MemberDhcpproperties) GetIpv6DdnsHostnameOk() (*string, bool)`

GetIpv6DdnsHostnameOk returns a tuple with the Ipv6DdnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsHostname

`func (o *MemberDhcpproperties) SetIpv6DdnsHostname(v string)`

SetIpv6DdnsHostname sets Ipv6DdnsHostname field to given value.

### HasIpv6DdnsHostname

`func (o *MemberDhcpproperties) HasIpv6DdnsHostname() bool`

HasIpv6DdnsHostname returns a boolean if a field has been set.

### GetIpv6DdnsServerAlwaysUpdates

`func (o *MemberDhcpproperties) GetIpv6DdnsServerAlwaysUpdates() bool`

GetIpv6DdnsServerAlwaysUpdates returns the Ipv6DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetIpv6DdnsServerAlwaysUpdatesOk

`func (o *MemberDhcpproperties) GetIpv6DdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetIpv6DdnsServerAlwaysUpdatesOk returns a tuple with the Ipv6DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsServerAlwaysUpdates

`func (o *MemberDhcpproperties) SetIpv6DdnsServerAlwaysUpdates(v bool)`

SetIpv6DdnsServerAlwaysUpdates sets Ipv6DdnsServerAlwaysUpdates field to given value.

### HasIpv6DdnsServerAlwaysUpdates

`func (o *MemberDhcpproperties) HasIpv6DdnsServerAlwaysUpdates() bool`

HasIpv6DdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetIpv6DdnsTtl

`func (o *MemberDhcpproperties) GetIpv6DdnsTtl() int64`

GetIpv6DdnsTtl returns the Ipv6DdnsTtl field if non-nil, zero value otherwise.

### GetIpv6DdnsTtlOk

`func (o *MemberDhcpproperties) GetIpv6DdnsTtlOk() (*int64, bool)`

GetIpv6DdnsTtlOk returns a tuple with the Ipv6DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsTtl

`func (o *MemberDhcpproperties) SetIpv6DdnsTtl(v int64)`

SetIpv6DdnsTtl sets Ipv6DdnsTtl field to given value.

### HasIpv6DdnsTtl

`func (o *MemberDhcpproperties) HasIpv6DdnsTtl() bool`

HasIpv6DdnsTtl returns a boolean if a field has been set.

### GetIpv6DnsUpdateStyle

`func (o *MemberDhcpproperties) GetIpv6DnsUpdateStyle() string`

GetIpv6DnsUpdateStyle returns the Ipv6DnsUpdateStyle field if non-nil, zero value otherwise.

### GetIpv6DnsUpdateStyleOk

`func (o *MemberDhcpproperties) GetIpv6DnsUpdateStyleOk() (*string, bool)`

GetIpv6DnsUpdateStyleOk returns a tuple with the Ipv6DnsUpdateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DnsUpdateStyle

`func (o *MemberDhcpproperties) SetIpv6DnsUpdateStyle(v string)`

SetIpv6DnsUpdateStyle sets Ipv6DnsUpdateStyle field to given value.

### HasIpv6DnsUpdateStyle

`func (o *MemberDhcpproperties) HasIpv6DnsUpdateStyle() bool`

HasIpv6DnsUpdateStyle returns a boolean if a field has been set.

### GetIpv6DomainName

`func (o *MemberDhcpproperties) GetIpv6DomainName() string`

GetIpv6DomainName returns the Ipv6DomainName field if non-nil, zero value otherwise.

### GetIpv6DomainNameOk

`func (o *MemberDhcpproperties) GetIpv6DomainNameOk() (*string, bool)`

GetIpv6DomainNameOk returns a tuple with the Ipv6DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DomainName

`func (o *MemberDhcpproperties) SetIpv6DomainName(v string)`

SetIpv6DomainName sets Ipv6DomainName field to given value.

### HasIpv6DomainName

`func (o *MemberDhcpproperties) HasIpv6DomainName() bool`

HasIpv6DomainName returns a boolean if a field has been set.

### GetIpv6DomainNameServers

`func (o *MemberDhcpproperties) GetIpv6DomainNameServers() []string`

GetIpv6DomainNameServers returns the Ipv6DomainNameServers field if non-nil, zero value otherwise.

### GetIpv6DomainNameServersOk

`func (o *MemberDhcpproperties) GetIpv6DomainNameServersOk() (*[]string, bool)`

GetIpv6DomainNameServersOk returns a tuple with the Ipv6DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DomainNameServers

`func (o *MemberDhcpproperties) SetIpv6DomainNameServers(v []string)`

SetIpv6DomainNameServers sets Ipv6DomainNameServers field to given value.

### HasIpv6DomainNameServers

`func (o *MemberDhcpproperties) HasIpv6DomainNameServers() bool`

HasIpv6DomainNameServers returns a boolean if a field has been set.

### GetIpv6EnableDdns

`func (o *MemberDhcpproperties) GetIpv6EnableDdns() bool`

GetIpv6EnableDdns returns the Ipv6EnableDdns field if non-nil, zero value otherwise.

### GetIpv6EnableDdnsOk

`func (o *MemberDhcpproperties) GetIpv6EnableDdnsOk() (*bool, bool)`

GetIpv6EnableDdnsOk returns a tuple with the Ipv6EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableDdns

`func (o *MemberDhcpproperties) SetIpv6EnableDdns(v bool)`

SetIpv6EnableDdns sets Ipv6EnableDdns field to given value.

### HasIpv6EnableDdns

`func (o *MemberDhcpproperties) HasIpv6EnableDdns() bool`

HasIpv6EnableDdns returns a boolean if a field has been set.

### GetIpv6EnableGssTsig

`func (o *MemberDhcpproperties) GetIpv6EnableGssTsig() bool`

GetIpv6EnableGssTsig returns the Ipv6EnableGssTsig field if non-nil, zero value otherwise.

### GetIpv6EnableGssTsigOk

`func (o *MemberDhcpproperties) GetIpv6EnableGssTsigOk() (*bool, bool)`

GetIpv6EnableGssTsigOk returns a tuple with the Ipv6EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableGssTsig

`func (o *MemberDhcpproperties) SetIpv6EnableGssTsig(v bool)`

SetIpv6EnableGssTsig sets Ipv6EnableGssTsig field to given value.

### HasIpv6EnableGssTsig

`func (o *MemberDhcpproperties) HasIpv6EnableGssTsig() bool`

HasIpv6EnableGssTsig returns a boolean if a field has been set.

### GetIpv6EnableLeaseScavenging

`func (o *MemberDhcpproperties) GetIpv6EnableLeaseScavenging() bool`

GetIpv6EnableLeaseScavenging returns the Ipv6EnableLeaseScavenging field if non-nil, zero value otherwise.

### GetIpv6EnableLeaseScavengingOk

`func (o *MemberDhcpproperties) GetIpv6EnableLeaseScavengingOk() (*bool, bool)`

GetIpv6EnableLeaseScavengingOk returns a tuple with the Ipv6EnableLeaseScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableLeaseScavenging

`func (o *MemberDhcpproperties) SetIpv6EnableLeaseScavenging(v bool)`

SetIpv6EnableLeaseScavenging sets Ipv6EnableLeaseScavenging field to given value.

### HasIpv6EnableLeaseScavenging

`func (o *MemberDhcpproperties) HasIpv6EnableLeaseScavenging() bool`

HasIpv6EnableLeaseScavenging returns a boolean if a field has been set.

### GetIpv6EnableRetryUpdates

`func (o *MemberDhcpproperties) GetIpv6EnableRetryUpdates() bool`

GetIpv6EnableRetryUpdates returns the Ipv6EnableRetryUpdates field if non-nil, zero value otherwise.

### GetIpv6EnableRetryUpdatesOk

`func (o *MemberDhcpproperties) GetIpv6EnableRetryUpdatesOk() (*bool, bool)`

GetIpv6EnableRetryUpdatesOk returns a tuple with the Ipv6EnableRetryUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableRetryUpdates

`func (o *MemberDhcpproperties) SetIpv6EnableRetryUpdates(v bool)`

SetIpv6EnableRetryUpdates sets Ipv6EnableRetryUpdates field to given value.

### HasIpv6EnableRetryUpdates

`func (o *MemberDhcpproperties) HasIpv6EnableRetryUpdates() bool`

HasIpv6EnableRetryUpdates returns a boolean if a field has been set.

### GetIpv6GenerateHostname

`func (o *MemberDhcpproperties) GetIpv6GenerateHostname() bool`

GetIpv6GenerateHostname returns the Ipv6GenerateHostname field if non-nil, zero value otherwise.

### GetIpv6GenerateHostnameOk

`func (o *MemberDhcpproperties) GetIpv6GenerateHostnameOk() (*bool, bool)`

GetIpv6GenerateHostnameOk returns a tuple with the Ipv6GenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6GenerateHostname

`func (o *MemberDhcpproperties) SetIpv6GenerateHostname(v bool)`

SetIpv6GenerateHostname sets Ipv6GenerateHostname field to given value.

### HasIpv6GenerateHostname

`func (o *MemberDhcpproperties) HasIpv6GenerateHostname() bool`

HasIpv6GenerateHostname returns a boolean if a field has been set.

### GetIpv6GssTsigKeys

`func (o *MemberDhcpproperties) GetIpv6GssTsigKeys() []string`

GetIpv6GssTsigKeys returns the Ipv6GssTsigKeys field if non-nil, zero value otherwise.

### GetIpv6GssTsigKeysOk

`func (o *MemberDhcpproperties) GetIpv6GssTsigKeysOk() (*[]string, bool)`

GetIpv6GssTsigKeysOk returns a tuple with the Ipv6GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6GssTsigKeys

`func (o *MemberDhcpproperties) SetIpv6GssTsigKeys(v []string)`

SetIpv6GssTsigKeys sets Ipv6GssTsigKeys field to given value.

### HasIpv6GssTsigKeys

`func (o *MemberDhcpproperties) HasIpv6GssTsigKeys() bool`

HasIpv6GssTsigKeys returns a boolean if a field has been set.

### GetIpv6KdcServer

`func (o *MemberDhcpproperties) GetIpv6KdcServer() string`

GetIpv6KdcServer returns the Ipv6KdcServer field if non-nil, zero value otherwise.

### GetIpv6KdcServerOk

`func (o *MemberDhcpproperties) GetIpv6KdcServerOk() (*string, bool)`

GetIpv6KdcServerOk returns a tuple with the Ipv6KdcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6KdcServer

`func (o *MemberDhcpproperties) SetIpv6KdcServer(v string)`

SetIpv6KdcServer sets Ipv6KdcServer field to given value.

### HasIpv6KdcServer

`func (o *MemberDhcpproperties) HasIpv6KdcServer() bool`

HasIpv6KdcServer returns a boolean if a field has been set.

### GetIpv6LeaseScavengingTime

`func (o *MemberDhcpproperties) GetIpv6LeaseScavengingTime() int64`

GetIpv6LeaseScavengingTime returns the Ipv6LeaseScavengingTime field if non-nil, zero value otherwise.

### GetIpv6LeaseScavengingTimeOk

`func (o *MemberDhcpproperties) GetIpv6LeaseScavengingTimeOk() (*int64, bool)`

GetIpv6LeaseScavengingTimeOk returns a tuple with the Ipv6LeaseScavengingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6LeaseScavengingTime

`func (o *MemberDhcpproperties) SetIpv6LeaseScavengingTime(v int64)`

SetIpv6LeaseScavengingTime sets Ipv6LeaseScavengingTime field to given value.

### HasIpv6LeaseScavengingTime

`func (o *MemberDhcpproperties) HasIpv6LeaseScavengingTime() bool`

HasIpv6LeaseScavengingTime returns a boolean if a field has been set.

### GetIpv6MicrosoftCodePage

`func (o *MemberDhcpproperties) GetIpv6MicrosoftCodePage() string`

GetIpv6MicrosoftCodePage returns the Ipv6MicrosoftCodePage field if non-nil, zero value otherwise.

### GetIpv6MicrosoftCodePageOk

`func (o *MemberDhcpproperties) GetIpv6MicrosoftCodePageOk() (*string, bool)`

GetIpv6MicrosoftCodePageOk returns a tuple with the Ipv6MicrosoftCodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MicrosoftCodePage

`func (o *MemberDhcpproperties) SetIpv6MicrosoftCodePage(v string)`

SetIpv6MicrosoftCodePage sets Ipv6MicrosoftCodePage field to given value.

### HasIpv6MicrosoftCodePage

`func (o *MemberDhcpproperties) HasIpv6MicrosoftCodePage() bool`

HasIpv6MicrosoftCodePage returns a boolean if a field has been set.

### GetIpv6Options

`func (o *MemberDhcpproperties) GetIpv6Options() []MemberDhcppropertiesIpv6Options`

GetIpv6Options returns the Ipv6Options field if non-nil, zero value otherwise.

### GetIpv6OptionsOk

`func (o *MemberDhcpproperties) GetIpv6OptionsOk() (*[]MemberDhcppropertiesIpv6Options, bool)`

GetIpv6OptionsOk returns a tuple with the Ipv6Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Options

`func (o *MemberDhcpproperties) SetIpv6Options(v []MemberDhcppropertiesIpv6Options)`

SetIpv6Options sets Ipv6Options field to given value.

### HasIpv6Options

`func (o *MemberDhcpproperties) HasIpv6Options() bool`

HasIpv6Options returns a boolean if a field has been set.

### GetIpv6RecycleLeases

`func (o *MemberDhcpproperties) GetIpv6RecycleLeases() bool`

GetIpv6RecycleLeases returns the Ipv6RecycleLeases field if non-nil, zero value otherwise.

### GetIpv6RecycleLeasesOk

`func (o *MemberDhcpproperties) GetIpv6RecycleLeasesOk() (*bool, bool)`

GetIpv6RecycleLeasesOk returns a tuple with the Ipv6RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RecycleLeases

`func (o *MemberDhcpproperties) SetIpv6RecycleLeases(v bool)`

SetIpv6RecycleLeases sets Ipv6RecycleLeases field to given value.

### HasIpv6RecycleLeases

`func (o *MemberDhcpproperties) HasIpv6RecycleLeases() bool`

HasIpv6RecycleLeases returns a boolean if a field has been set.

### GetIpv6RememberExpiredClientAssociation

`func (o *MemberDhcpproperties) GetIpv6RememberExpiredClientAssociation() bool`

GetIpv6RememberExpiredClientAssociation returns the Ipv6RememberExpiredClientAssociation field if non-nil, zero value otherwise.

### GetIpv6RememberExpiredClientAssociationOk

`func (o *MemberDhcpproperties) GetIpv6RememberExpiredClientAssociationOk() (*bool, bool)`

GetIpv6RememberExpiredClientAssociationOk returns a tuple with the Ipv6RememberExpiredClientAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RememberExpiredClientAssociation

`func (o *MemberDhcpproperties) SetIpv6RememberExpiredClientAssociation(v bool)`

SetIpv6RememberExpiredClientAssociation sets Ipv6RememberExpiredClientAssociation field to given value.

### HasIpv6RememberExpiredClientAssociation

`func (o *MemberDhcpproperties) HasIpv6RememberExpiredClientAssociation() bool`

HasIpv6RememberExpiredClientAssociation returns a boolean if a field has been set.

### GetIpv6RetryUpdatesInterval

`func (o *MemberDhcpproperties) GetIpv6RetryUpdatesInterval() int64`

GetIpv6RetryUpdatesInterval returns the Ipv6RetryUpdatesInterval field if non-nil, zero value otherwise.

### GetIpv6RetryUpdatesIntervalOk

`func (o *MemberDhcpproperties) GetIpv6RetryUpdatesIntervalOk() (*int64, bool)`

GetIpv6RetryUpdatesIntervalOk returns a tuple with the Ipv6RetryUpdatesInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RetryUpdatesInterval

`func (o *MemberDhcpproperties) SetIpv6RetryUpdatesInterval(v int64)`

SetIpv6RetryUpdatesInterval sets Ipv6RetryUpdatesInterval field to given value.

### HasIpv6RetryUpdatesInterval

`func (o *MemberDhcpproperties) HasIpv6RetryUpdatesInterval() bool`

HasIpv6RetryUpdatesInterval returns a boolean if a field has been set.

### GetIpv6ServerDuid

`func (o *MemberDhcpproperties) GetIpv6ServerDuid() string`

GetIpv6ServerDuid returns the Ipv6ServerDuid field if non-nil, zero value otherwise.

### GetIpv6ServerDuidOk

`func (o *MemberDhcpproperties) GetIpv6ServerDuidOk() (*string, bool)`

GetIpv6ServerDuidOk returns a tuple with the Ipv6ServerDuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6ServerDuid

`func (o *MemberDhcpproperties) SetIpv6ServerDuid(v string)`

SetIpv6ServerDuid sets Ipv6ServerDuid field to given value.

### HasIpv6ServerDuid

`func (o *MemberDhcpproperties) HasIpv6ServerDuid() bool`

HasIpv6ServerDuid returns a boolean if a field has been set.

### GetIpv6UpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) GetIpv6UpdateDnsOnLeaseRenewal() bool`

GetIpv6UpdateDnsOnLeaseRenewal returns the Ipv6UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetIpv6UpdateDnsOnLeaseRenewalOk

`func (o *MemberDhcpproperties) GetIpv6UpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetIpv6UpdateDnsOnLeaseRenewalOk returns a tuple with the Ipv6UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6UpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) SetIpv6UpdateDnsOnLeaseRenewal(v bool)`

SetIpv6UpdateDnsOnLeaseRenewal sets Ipv6UpdateDnsOnLeaseRenewal field to given value.

### HasIpv6UpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) HasIpv6UpdateDnsOnLeaseRenewal() bool`

HasIpv6UpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetIpv6addr

`func (o *MemberDhcpproperties) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *MemberDhcpproperties) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *MemberDhcpproperties) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *MemberDhcpproperties) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetKdcServer

`func (o *MemberDhcpproperties) GetKdcServer() string`

GetKdcServer returns the KdcServer field if non-nil, zero value otherwise.

### GetKdcServerOk

`func (o *MemberDhcpproperties) GetKdcServerOk() (*string, bool)`

GetKdcServerOk returns a tuple with the KdcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcServer

`func (o *MemberDhcpproperties) SetKdcServer(v string)`

SetKdcServer sets KdcServer field to given value.

### HasKdcServer

`func (o *MemberDhcpproperties) HasKdcServer() bool`

HasKdcServer returns a boolean if a field has been set.

### GetLeasePerClientSettings

`func (o *MemberDhcpproperties) GetLeasePerClientSettings() string`

GetLeasePerClientSettings returns the LeasePerClientSettings field if non-nil, zero value otherwise.

### GetLeasePerClientSettingsOk

`func (o *MemberDhcpproperties) GetLeasePerClientSettingsOk() (*string, bool)`

GetLeasePerClientSettingsOk returns a tuple with the LeasePerClientSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasePerClientSettings

`func (o *MemberDhcpproperties) SetLeasePerClientSettings(v string)`

SetLeasePerClientSettings sets LeasePerClientSettings field to given value.

### HasLeasePerClientSettings

`func (o *MemberDhcpproperties) HasLeasePerClientSettings() bool`

HasLeasePerClientSettings returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *MemberDhcpproperties) GetLeaseScavengeTime() int32`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *MemberDhcpproperties) GetLeaseScavengeTimeOk() (*int32, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *MemberDhcpproperties) SetLeaseScavengeTime(v int32)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *MemberDhcpproperties) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogLeaseEvents

`func (o *MemberDhcpproperties) GetLogLeaseEvents() bool`

GetLogLeaseEvents returns the LogLeaseEvents field if non-nil, zero value otherwise.

### GetLogLeaseEventsOk

`func (o *MemberDhcpproperties) GetLogLeaseEventsOk() (*bool, bool)`

GetLogLeaseEventsOk returns a tuple with the LogLeaseEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLeaseEvents

`func (o *MemberDhcpproperties) SetLogLeaseEvents(v bool)`

SetLogLeaseEvents sets LogLeaseEvents field to given value.

### HasLogLeaseEvents

`func (o *MemberDhcpproperties) HasLogLeaseEvents() bool`

HasLogLeaseEvents returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *MemberDhcpproperties) GetLogicFilterRules() []MemberDhcppropertiesLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *MemberDhcpproperties) GetLogicFilterRulesOk() (*[]MemberDhcppropertiesLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *MemberDhcpproperties) SetLogicFilterRules(v []MemberDhcppropertiesLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *MemberDhcpproperties) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *MemberDhcpproperties) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *MemberDhcpproperties) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *MemberDhcpproperties) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *MemberDhcpproperties) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *MemberDhcpproperties) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *MemberDhcpproperties) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *MemberDhcpproperties) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *MemberDhcpproperties) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMicrosoftCodePage

`func (o *MemberDhcpproperties) GetMicrosoftCodePage() string`

GetMicrosoftCodePage returns the MicrosoftCodePage field if non-nil, zero value otherwise.

### GetMicrosoftCodePageOk

`func (o *MemberDhcpproperties) GetMicrosoftCodePageOk() (*string, bool)`

GetMicrosoftCodePageOk returns a tuple with the MicrosoftCodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftCodePage

`func (o *MemberDhcpproperties) SetMicrosoftCodePage(v string)`

SetMicrosoftCodePage sets MicrosoftCodePage field to given value.

### HasMicrosoftCodePage

`func (o *MemberDhcpproperties) HasMicrosoftCodePage() bool`

HasMicrosoftCodePage returns a boolean if a field has been set.

### GetNextserver

`func (o *MemberDhcpproperties) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *MemberDhcpproperties) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *MemberDhcpproperties) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *MemberDhcpproperties) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOption60MatchRules

`func (o *MemberDhcpproperties) GetOption60MatchRules() []MemberDhcppropertiesOption60MatchRules`

GetOption60MatchRules returns the Option60MatchRules field if non-nil, zero value otherwise.

### GetOption60MatchRulesOk

`func (o *MemberDhcpproperties) GetOption60MatchRulesOk() (*[]MemberDhcppropertiesOption60MatchRules, bool)`

GetOption60MatchRulesOk returns a tuple with the Option60MatchRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption60MatchRules

`func (o *MemberDhcpproperties) SetOption60MatchRules(v []MemberDhcppropertiesOption60MatchRules)`

SetOption60MatchRules sets Option60MatchRules field to given value.

### HasOption60MatchRules

`func (o *MemberDhcpproperties) HasOption60MatchRules() bool`

HasOption60MatchRules returns a boolean if a field has been set.

### GetOptions

`func (o *MemberDhcpproperties) GetOptions() []MemberDhcppropertiesOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MemberDhcpproperties) GetOptionsOk() (*[]MemberDhcppropertiesOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MemberDhcpproperties) SetOptions(v []MemberDhcppropertiesOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MemberDhcpproperties) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPingCount

`func (o *MemberDhcpproperties) GetPingCount() int64`

GetPingCount returns the PingCount field if non-nil, zero value otherwise.

### GetPingCountOk

`func (o *MemberDhcpproperties) GetPingCountOk() (*int64, bool)`

GetPingCountOk returns a tuple with the PingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingCount

`func (o *MemberDhcpproperties) SetPingCount(v int64)`

SetPingCount sets PingCount field to given value.

### HasPingCount

`func (o *MemberDhcpproperties) HasPingCount() bool`

HasPingCount returns a boolean if a field has been set.

### GetPingTimeout

`func (o *MemberDhcpproperties) GetPingTimeout() int64`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *MemberDhcpproperties) GetPingTimeoutOk() (*int64, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *MemberDhcpproperties) SetPingTimeout(v int64)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *MemberDhcpproperties) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *MemberDhcpproperties) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *MemberDhcpproperties) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *MemberDhcpproperties) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *MemberDhcpproperties) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetPrefixLengthMode

`func (o *MemberDhcpproperties) GetPrefixLengthMode() string`

GetPrefixLengthMode returns the PrefixLengthMode field if non-nil, zero value otherwise.

### GetPrefixLengthModeOk

`func (o *MemberDhcpproperties) GetPrefixLengthModeOk() (*string, bool)`

GetPrefixLengthModeOk returns a tuple with the PrefixLengthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLengthMode

`func (o *MemberDhcpproperties) SetPrefixLengthMode(v string)`

SetPrefixLengthMode sets PrefixLengthMode field to given value.

### HasPrefixLengthMode

`func (o *MemberDhcpproperties) HasPrefixLengthMode() bool`

HasPrefixLengthMode returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *MemberDhcpproperties) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *MemberDhcpproperties) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *MemberDhcpproperties) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *MemberDhcpproperties) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *MemberDhcpproperties) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *MemberDhcpproperties) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *MemberDhcpproperties) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *MemberDhcpproperties) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRetryDdnsUpdates

`func (o *MemberDhcpproperties) GetRetryDdnsUpdates() bool`

GetRetryDdnsUpdates returns the RetryDdnsUpdates field if non-nil, zero value otherwise.

### GetRetryDdnsUpdatesOk

`func (o *MemberDhcpproperties) GetRetryDdnsUpdatesOk() (*bool, bool)`

GetRetryDdnsUpdatesOk returns a tuple with the RetryDdnsUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDdnsUpdates

`func (o *MemberDhcpproperties) SetRetryDdnsUpdates(v bool)`

SetRetryDdnsUpdates sets RetryDdnsUpdates field to given value.

### HasRetryDdnsUpdates

`func (o *MemberDhcpproperties) HasRetryDdnsUpdates() bool`

HasRetryDdnsUpdates returns a boolean if a field has been set.

### GetStaticHosts

`func (o *MemberDhcpproperties) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *MemberDhcpproperties) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *MemberDhcpproperties) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *MemberDhcpproperties) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *MemberDhcpproperties) GetSyslogFacility() string`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *MemberDhcpproperties) GetSyslogFacilityOk() (*string, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *MemberDhcpproperties) SetSyslogFacility(v string)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *MemberDhcpproperties) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetTotalHosts

`func (o *MemberDhcpproperties) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *MemberDhcpproperties) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *MemberDhcpproperties) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *MemberDhcpproperties) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *MemberDhcpproperties) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseAuthority

`func (o *MemberDhcpproperties) GetUseAuthority() bool`

GetUseAuthority returns the UseAuthority field if non-nil, zero value otherwise.

### GetUseAuthorityOk

`func (o *MemberDhcpproperties) GetUseAuthorityOk() (*bool, bool)`

GetUseAuthorityOk returns a tuple with the UseAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthority

`func (o *MemberDhcpproperties) SetUseAuthority(v bool)`

SetUseAuthority sets UseAuthority field to given value.

### HasUseAuthority

`func (o *MemberDhcpproperties) HasUseAuthority() bool`

HasUseAuthority returns a boolean if a field has been set.

### GetUseBootfile

`func (o *MemberDhcpproperties) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *MemberDhcpproperties) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *MemberDhcpproperties) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *MemberDhcpproperties) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *MemberDhcpproperties) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *MemberDhcpproperties) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *MemberDhcpproperties) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *MemberDhcpproperties) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *MemberDhcpproperties) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *MemberDhcpproperties) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *MemberDhcpproperties) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *MemberDhcpproperties) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *MemberDhcpproperties) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *MemberDhcpproperties) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *MemberDhcpproperties) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *MemberDhcpproperties) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *MemberDhcpproperties) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *MemberDhcpproperties) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *MemberDhcpproperties) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *MemberDhcpproperties) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDdnsUpdateFixedAddresses

`func (o *MemberDhcpproperties) GetUseDdnsUpdateFixedAddresses() bool`

GetUseDdnsUpdateFixedAddresses returns the UseDdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetUseDdnsUpdateFixedAddressesOk

`func (o *MemberDhcpproperties) GetUseDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetUseDdnsUpdateFixedAddressesOk returns a tuple with the UseDdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUpdateFixedAddresses

`func (o *MemberDhcpproperties) SetUseDdnsUpdateFixedAddresses(v bool)`

SetUseDdnsUpdateFixedAddresses sets UseDdnsUpdateFixedAddresses field to given value.

### HasUseDdnsUpdateFixedAddresses

`func (o *MemberDhcpproperties) HasUseDdnsUpdateFixedAddresses() bool`

HasUseDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetUseDdnsUseOption81

`func (o *MemberDhcpproperties) GetUseDdnsUseOption81() bool`

GetUseDdnsUseOption81 returns the UseDdnsUseOption81 field if non-nil, zero value otherwise.

### GetUseDdnsUseOption81Ok

`func (o *MemberDhcpproperties) GetUseDdnsUseOption81Ok() (*bool, bool)`

GetUseDdnsUseOption81Ok returns a tuple with the UseDdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUseOption81

`func (o *MemberDhcpproperties) SetUseDdnsUseOption81(v bool)`

SetUseDdnsUseOption81 sets UseDdnsUseOption81 field to given value.

### HasUseDdnsUseOption81

`func (o *MemberDhcpproperties) HasUseDdnsUseOption81() bool`

HasUseDdnsUseOption81 returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *MemberDhcpproperties) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *MemberDhcpproperties) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *MemberDhcpproperties) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *MemberDhcpproperties) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseDnsUpdateStyle

`func (o *MemberDhcpproperties) GetUseDnsUpdateStyle() bool`

GetUseDnsUpdateStyle returns the UseDnsUpdateStyle field if non-nil, zero value otherwise.

### GetUseDnsUpdateStyleOk

`func (o *MemberDhcpproperties) GetUseDnsUpdateStyleOk() (*bool, bool)`

GetUseDnsUpdateStyleOk returns a tuple with the UseDnsUpdateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsUpdateStyle

`func (o *MemberDhcpproperties) SetUseDnsUpdateStyle(v bool)`

SetUseDnsUpdateStyle sets UseDnsUpdateStyle field to given value.

### HasUseDnsUpdateStyle

`func (o *MemberDhcpproperties) HasUseDnsUpdateStyle() bool`

HasUseDnsUpdateStyle returns a boolean if a field has been set.

### GetUseEmailList

`func (o *MemberDhcpproperties) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *MemberDhcpproperties) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *MemberDhcpproperties) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *MemberDhcpproperties) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *MemberDhcpproperties) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *MemberDhcpproperties) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *MemberDhcpproperties) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *MemberDhcpproperties) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *MemberDhcpproperties) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *MemberDhcpproperties) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *MemberDhcpproperties) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *MemberDhcpproperties) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseEnableFingerprint

`func (o *MemberDhcpproperties) GetUseEnableFingerprint() bool`

GetUseEnableFingerprint returns the UseEnableFingerprint field if non-nil, zero value otherwise.

### GetUseEnableFingerprintOk

`func (o *MemberDhcpproperties) GetUseEnableFingerprintOk() (*bool, bool)`

GetUseEnableFingerprintOk returns a tuple with the UseEnableFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableFingerprint

`func (o *MemberDhcpproperties) SetUseEnableFingerprint(v bool)`

SetUseEnableFingerprint sets UseEnableFingerprint field to given value.

### HasUseEnableFingerprint

`func (o *MemberDhcpproperties) HasUseEnableFingerprint() bool`

HasUseEnableFingerprint returns a boolean if a field has been set.

### GetUseEnableGssTsig

`func (o *MemberDhcpproperties) GetUseEnableGssTsig() bool`

GetUseEnableGssTsig returns the UseEnableGssTsig field if non-nil, zero value otherwise.

### GetUseEnableGssTsigOk

`func (o *MemberDhcpproperties) GetUseEnableGssTsigOk() (*bool, bool)`

GetUseEnableGssTsigOk returns a tuple with the UseEnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableGssTsig

`func (o *MemberDhcpproperties) SetUseEnableGssTsig(v bool)`

SetUseEnableGssTsig sets UseEnableGssTsig field to given value.

### HasUseEnableGssTsig

`func (o *MemberDhcpproperties) HasUseEnableGssTsig() bool`

HasUseEnableGssTsig returns a boolean if a field has been set.

### GetUseEnableHostnameRewrite

`func (o *MemberDhcpproperties) GetUseEnableHostnameRewrite() bool`

GetUseEnableHostnameRewrite returns the UseEnableHostnameRewrite field if non-nil, zero value otherwise.

### GetUseEnableHostnameRewriteOk

`func (o *MemberDhcpproperties) GetUseEnableHostnameRewriteOk() (*bool, bool)`

GetUseEnableHostnameRewriteOk returns a tuple with the UseEnableHostnameRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableHostnameRewrite

`func (o *MemberDhcpproperties) SetUseEnableHostnameRewrite(v bool)`

SetUseEnableHostnameRewrite sets UseEnableHostnameRewrite field to given value.

### HasUseEnableHostnameRewrite

`func (o *MemberDhcpproperties) HasUseEnableHostnameRewrite() bool`

HasUseEnableHostnameRewrite returns a boolean if a field has been set.

### GetUseEnableLeasequery

`func (o *MemberDhcpproperties) GetUseEnableLeasequery() bool`

GetUseEnableLeasequery returns the UseEnableLeasequery field if non-nil, zero value otherwise.

### GetUseEnableLeasequeryOk

`func (o *MemberDhcpproperties) GetUseEnableLeasequeryOk() (*bool, bool)`

GetUseEnableLeasequeryOk returns a tuple with the UseEnableLeasequery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableLeasequery

`func (o *MemberDhcpproperties) SetUseEnableLeasequery(v bool)`

SetUseEnableLeasequery sets UseEnableLeasequery field to given value.

### HasUseEnableLeasequery

`func (o *MemberDhcpproperties) HasUseEnableLeasequery() bool`

HasUseEnableLeasequery returns a boolean if a field has been set.

### GetUseEnableOneLeasePerClient

`func (o *MemberDhcpproperties) GetUseEnableOneLeasePerClient() bool`

GetUseEnableOneLeasePerClient returns the UseEnableOneLeasePerClient field if non-nil, zero value otherwise.

### GetUseEnableOneLeasePerClientOk

`func (o *MemberDhcpproperties) GetUseEnableOneLeasePerClientOk() (*bool, bool)`

GetUseEnableOneLeasePerClientOk returns a tuple with the UseEnableOneLeasePerClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableOneLeasePerClient

`func (o *MemberDhcpproperties) SetUseEnableOneLeasePerClient(v bool)`

SetUseEnableOneLeasePerClient sets UseEnableOneLeasePerClient field to given value.

### HasUseEnableOneLeasePerClient

`func (o *MemberDhcpproperties) HasUseEnableOneLeasePerClient() bool`

HasUseEnableOneLeasePerClient returns a boolean if a field has been set.

### GetUseGssTsigKeys

`func (o *MemberDhcpproperties) GetUseGssTsigKeys() bool`

GetUseGssTsigKeys returns the UseGssTsigKeys field if non-nil, zero value otherwise.

### GetUseGssTsigKeysOk

`func (o *MemberDhcpproperties) GetUseGssTsigKeysOk() (*bool, bool)`

GetUseGssTsigKeysOk returns a tuple with the UseGssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGssTsigKeys

`func (o *MemberDhcpproperties) SetUseGssTsigKeys(v bool)`

SetUseGssTsigKeys sets UseGssTsigKeys field to given value.

### HasUseGssTsigKeys

`func (o *MemberDhcpproperties) HasUseGssTsigKeys() bool`

HasUseGssTsigKeys returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *MemberDhcpproperties) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *MemberDhcpproperties) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *MemberDhcpproperties) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *MemberDhcpproperties) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *MemberDhcpproperties) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *MemberDhcpproperties) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *MemberDhcpproperties) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *MemberDhcpproperties) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseImmediateFaConfiguration

`func (o *MemberDhcpproperties) GetUseImmediateFaConfiguration() bool`

GetUseImmediateFaConfiguration returns the UseImmediateFaConfiguration field if non-nil, zero value otherwise.

### GetUseImmediateFaConfigurationOk

`func (o *MemberDhcpproperties) GetUseImmediateFaConfigurationOk() (*bool, bool)`

GetUseImmediateFaConfigurationOk returns a tuple with the UseImmediateFaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseImmediateFaConfiguration

`func (o *MemberDhcpproperties) SetUseImmediateFaConfiguration(v bool)`

SetUseImmediateFaConfiguration sets UseImmediateFaConfiguration field to given value.

### HasUseImmediateFaConfiguration

`func (o *MemberDhcpproperties) HasUseImmediateFaConfiguration() bool`

HasUseImmediateFaConfiguration returns a boolean if a field has been set.

### GetUseIpv6DdnsDomainname

`func (o *MemberDhcpproperties) GetUseIpv6DdnsDomainname() bool`

GetUseIpv6DdnsDomainname returns the UseIpv6DdnsDomainname field if non-nil, zero value otherwise.

### GetUseIpv6DdnsDomainnameOk

`func (o *MemberDhcpproperties) GetUseIpv6DdnsDomainnameOk() (*bool, bool)`

GetUseIpv6DdnsDomainnameOk returns a tuple with the UseIpv6DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DdnsDomainname

`func (o *MemberDhcpproperties) SetUseIpv6DdnsDomainname(v bool)`

SetUseIpv6DdnsDomainname sets UseIpv6DdnsDomainname field to given value.

### HasUseIpv6DdnsDomainname

`func (o *MemberDhcpproperties) HasUseIpv6DdnsDomainname() bool`

HasUseIpv6DdnsDomainname returns a boolean if a field has been set.

### GetUseIpv6DdnsEnableOptionFqdn

`func (o *MemberDhcpproperties) GetUseIpv6DdnsEnableOptionFqdn() bool`

GetUseIpv6DdnsEnableOptionFqdn returns the UseIpv6DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetUseIpv6DdnsEnableOptionFqdnOk

`func (o *MemberDhcpproperties) GetUseIpv6DdnsEnableOptionFqdnOk() (*bool, bool)`

GetUseIpv6DdnsEnableOptionFqdnOk returns a tuple with the UseIpv6DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DdnsEnableOptionFqdn

`func (o *MemberDhcpproperties) SetUseIpv6DdnsEnableOptionFqdn(v bool)`

SetUseIpv6DdnsEnableOptionFqdn sets UseIpv6DdnsEnableOptionFqdn field to given value.

### HasUseIpv6DdnsEnableOptionFqdn

`func (o *MemberDhcpproperties) HasUseIpv6DdnsEnableOptionFqdn() bool`

HasUseIpv6DdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetUseIpv6DdnsHostname

`func (o *MemberDhcpproperties) GetUseIpv6DdnsHostname() bool`

GetUseIpv6DdnsHostname returns the UseIpv6DdnsHostname field if non-nil, zero value otherwise.

### GetUseIpv6DdnsHostnameOk

`func (o *MemberDhcpproperties) GetUseIpv6DdnsHostnameOk() (*bool, bool)`

GetUseIpv6DdnsHostnameOk returns a tuple with the UseIpv6DdnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DdnsHostname

`func (o *MemberDhcpproperties) SetUseIpv6DdnsHostname(v bool)`

SetUseIpv6DdnsHostname sets UseIpv6DdnsHostname field to given value.

### HasUseIpv6DdnsHostname

`func (o *MemberDhcpproperties) HasUseIpv6DdnsHostname() bool`

HasUseIpv6DdnsHostname returns a boolean if a field has been set.

### GetUseIpv6DdnsTtl

`func (o *MemberDhcpproperties) GetUseIpv6DdnsTtl() bool`

GetUseIpv6DdnsTtl returns the UseIpv6DdnsTtl field if non-nil, zero value otherwise.

### GetUseIpv6DdnsTtlOk

`func (o *MemberDhcpproperties) GetUseIpv6DdnsTtlOk() (*bool, bool)`

GetUseIpv6DdnsTtlOk returns a tuple with the UseIpv6DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DdnsTtl

`func (o *MemberDhcpproperties) SetUseIpv6DdnsTtl(v bool)`

SetUseIpv6DdnsTtl sets UseIpv6DdnsTtl field to given value.

### HasUseIpv6DdnsTtl

`func (o *MemberDhcpproperties) HasUseIpv6DdnsTtl() bool`

HasUseIpv6DdnsTtl returns a boolean if a field has been set.

### GetUseIpv6DnsUpdateStyle

`func (o *MemberDhcpproperties) GetUseIpv6DnsUpdateStyle() bool`

GetUseIpv6DnsUpdateStyle returns the UseIpv6DnsUpdateStyle field if non-nil, zero value otherwise.

### GetUseIpv6DnsUpdateStyleOk

`func (o *MemberDhcpproperties) GetUseIpv6DnsUpdateStyleOk() (*bool, bool)`

GetUseIpv6DnsUpdateStyleOk returns a tuple with the UseIpv6DnsUpdateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DnsUpdateStyle

`func (o *MemberDhcpproperties) SetUseIpv6DnsUpdateStyle(v bool)`

SetUseIpv6DnsUpdateStyle sets UseIpv6DnsUpdateStyle field to given value.

### HasUseIpv6DnsUpdateStyle

`func (o *MemberDhcpproperties) HasUseIpv6DnsUpdateStyle() bool`

HasUseIpv6DnsUpdateStyle returns a boolean if a field has been set.

### GetUseIpv6DomainName

`func (o *MemberDhcpproperties) GetUseIpv6DomainName() bool`

GetUseIpv6DomainName returns the UseIpv6DomainName field if non-nil, zero value otherwise.

### GetUseIpv6DomainNameOk

`func (o *MemberDhcpproperties) GetUseIpv6DomainNameOk() (*bool, bool)`

GetUseIpv6DomainNameOk returns a tuple with the UseIpv6DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DomainName

`func (o *MemberDhcpproperties) SetUseIpv6DomainName(v bool)`

SetUseIpv6DomainName sets UseIpv6DomainName field to given value.

### HasUseIpv6DomainName

`func (o *MemberDhcpproperties) HasUseIpv6DomainName() bool`

HasUseIpv6DomainName returns a boolean if a field has been set.

### GetUseIpv6DomainNameServers

`func (o *MemberDhcpproperties) GetUseIpv6DomainNameServers() bool`

GetUseIpv6DomainNameServers returns the UseIpv6DomainNameServers field if non-nil, zero value otherwise.

### GetUseIpv6DomainNameServersOk

`func (o *MemberDhcpproperties) GetUseIpv6DomainNameServersOk() (*bool, bool)`

GetUseIpv6DomainNameServersOk returns a tuple with the UseIpv6DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6DomainNameServers

`func (o *MemberDhcpproperties) SetUseIpv6DomainNameServers(v bool)`

SetUseIpv6DomainNameServers sets UseIpv6DomainNameServers field to given value.

### HasUseIpv6DomainNameServers

`func (o *MemberDhcpproperties) HasUseIpv6DomainNameServers() bool`

HasUseIpv6DomainNameServers returns a boolean if a field has been set.

### GetUseIpv6EnableDdns

`func (o *MemberDhcpproperties) GetUseIpv6EnableDdns() bool`

GetUseIpv6EnableDdns returns the UseIpv6EnableDdns field if non-nil, zero value otherwise.

### GetUseIpv6EnableDdnsOk

`func (o *MemberDhcpproperties) GetUseIpv6EnableDdnsOk() (*bool, bool)`

GetUseIpv6EnableDdnsOk returns a tuple with the UseIpv6EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6EnableDdns

`func (o *MemberDhcpproperties) SetUseIpv6EnableDdns(v bool)`

SetUseIpv6EnableDdns sets UseIpv6EnableDdns field to given value.

### HasUseIpv6EnableDdns

`func (o *MemberDhcpproperties) HasUseIpv6EnableDdns() bool`

HasUseIpv6EnableDdns returns a boolean if a field has been set.

### GetUseIpv6EnableGssTsig

`func (o *MemberDhcpproperties) GetUseIpv6EnableGssTsig() bool`

GetUseIpv6EnableGssTsig returns the UseIpv6EnableGssTsig field if non-nil, zero value otherwise.

### GetUseIpv6EnableGssTsigOk

`func (o *MemberDhcpproperties) GetUseIpv6EnableGssTsigOk() (*bool, bool)`

GetUseIpv6EnableGssTsigOk returns a tuple with the UseIpv6EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6EnableGssTsig

`func (o *MemberDhcpproperties) SetUseIpv6EnableGssTsig(v bool)`

SetUseIpv6EnableGssTsig sets UseIpv6EnableGssTsig field to given value.

### HasUseIpv6EnableGssTsig

`func (o *MemberDhcpproperties) HasUseIpv6EnableGssTsig() bool`

HasUseIpv6EnableGssTsig returns a boolean if a field has been set.

### GetUseIpv6EnableRetryUpdates

`func (o *MemberDhcpproperties) GetUseIpv6EnableRetryUpdates() bool`

GetUseIpv6EnableRetryUpdates returns the UseIpv6EnableRetryUpdates field if non-nil, zero value otherwise.

### GetUseIpv6EnableRetryUpdatesOk

`func (o *MemberDhcpproperties) GetUseIpv6EnableRetryUpdatesOk() (*bool, bool)`

GetUseIpv6EnableRetryUpdatesOk returns a tuple with the UseIpv6EnableRetryUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6EnableRetryUpdates

`func (o *MemberDhcpproperties) SetUseIpv6EnableRetryUpdates(v bool)`

SetUseIpv6EnableRetryUpdates sets UseIpv6EnableRetryUpdates field to given value.

### HasUseIpv6EnableRetryUpdates

`func (o *MemberDhcpproperties) HasUseIpv6EnableRetryUpdates() bool`

HasUseIpv6EnableRetryUpdates returns a boolean if a field has been set.

### GetUseIpv6GenerateHostname

`func (o *MemberDhcpproperties) GetUseIpv6GenerateHostname() bool`

GetUseIpv6GenerateHostname returns the UseIpv6GenerateHostname field if non-nil, zero value otherwise.

### GetUseIpv6GenerateHostnameOk

`func (o *MemberDhcpproperties) GetUseIpv6GenerateHostnameOk() (*bool, bool)`

GetUseIpv6GenerateHostnameOk returns a tuple with the UseIpv6GenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6GenerateHostname

`func (o *MemberDhcpproperties) SetUseIpv6GenerateHostname(v bool)`

SetUseIpv6GenerateHostname sets UseIpv6GenerateHostname field to given value.

### HasUseIpv6GenerateHostname

`func (o *MemberDhcpproperties) HasUseIpv6GenerateHostname() bool`

HasUseIpv6GenerateHostname returns a boolean if a field has been set.

### GetUseIpv6GssTsigKeys

`func (o *MemberDhcpproperties) GetUseIpv6GssTsigKeys() bool`

GetUseIpv6GssTsigKeys returns the UseIpv6GssTsigKeys field if non-nil, zero value otherwise.

### GetUseIpv6GssTsigKeysOk

`func (o *MemberDhcpproperties) GetUseIpv6GssTsigKeysOk() (*bool, bool)`

GetUseIpv6GssTsigKeysOk returns a tuple with the UseIpv6GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6GssTsigKeys

`func (o *MemberDhcpproperties) SetUseIpv6GssTsigKeys(v bool)`

SetUseIpv6GssTsigKeys sets UseIpv6GssTsigKeys field to given value.

### HasUseIpv6GssTsigKeys

`func (o *MemberDhcpproperties) HasUseIpv6GssTsigKeys() bool`

HasUseIpv6GssTsigKeys returns a boolean if a field has been set.

### GetUseIpv6LeaseScavenging

`func (o *MemberDhcpproperties) GetUseIpv6LeaseScavenging() bool`

GetUseIpv6LeaseScavenging returns the UseIpv6LeaseScavenging field if non-nil, zero value otherwise.

### GetUseIpv6LeaseScavengingOk

`func (o *MemberDhcpproperties) GetUseIpv6LeaseScavengingOk() (*bool, bool)`

GetUseIpv6LeaseScavengingOk returns a tuple with the UseIpv6LeaseScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6LeaseScavenging

`func (o *MemberDhcpproperties) SetUseIpv6LeaseScavenging(v bool)`

SetUseIpv6LeaseScavenging sets UseIpv6LeaseScavenging field to given value.

### HasUseIpv6LeaseScavenging

`func (o *MemberDhcpproperties) HasUseIpv6LeaseScavenging() bool`

HasUseIpv6LeaseScavenging returns a boolean if a field has been set.

### GetUseIpv6MicrosoftCodePage

`func (o *MemberDhcpproperties) GetUseIpv6MicrosoftCodePage() bool`

GetUseIpv6MicrosoftCodePage returns the UseIpv6MicrosoftCodePage field if non-nil, zero value otherwise.

### GetUseIpv6MicrosoftCodePageOk

`func (o *MemberDhcpproperties) GetUseIpv6MicrosoftCodePageOk() (*bool, bool)`

GetUseIpv6MicrosoftCodePageOk returns a tuple with the UseIpv6MicrosoftCodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6MicrosoftCodePage

`func (o *MemberDhcpproperties) SetUseIpv6MicrosoftCodePage(v bool)`

SetUseIpv6MicrosoftCodePage sets UseIpv6MicrosoftCodePage field to given value.

### HasUseIpv6MicrosoftCodePage

`func (o *MemberDhcpproperties) HasUseIpv6MicrosoftCodePage() bool`

HasUseIpv6MicrosoftCodePage returns a boolean if a field has been set.

### GetUseIpv6Options

`func (o *MemberDhcpproperties) GetUseIpv6Options() bool`

GetUseIpv6Options returns the UseIpv6Options field if non-nil, zero value otherwise.

### GetUseIpv6OptionsOk

`func (o *MemberDhcpproperties) GetUseIpv6OptionsOk() (*bool, bool)`

GetUseIpv6OptionsOk returns a tuple with the UseIpv6Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6Options

`func (o *MemberDhcpproperties) SetUseIpv6Options(v bool)`

SetUseIpv6Options sets UseIpv6Options field to given value.

### HasUseIpv6Options

`func (o *MemberDhcpproperties) HasUseIpv6Options() bool`

HasUseIpv6Options returns a boolean if a field has been set.

### GetUseIpv6RecycleLeases

`func (o *MemberDhcpproperties) GetUseIpv6RecycleLeases() bool`

GetUseIpv6RecycleLeases returns the UseIpv6RecycleLeases field if non-nil, zero value otherwise.

### GetUseIpv6RecycleLeasesOk

`func (o *MemberDhcpproperties) GetUseIpv6RecycleLeasesOk() (*bool, bool)`

GetUseIpv6RecycleLeasesOk returns a tuple with the UseIpv6RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6RecycleLeases

`func (o *MemberDhcpproperties) SetUseIpv6RecycleLeases(v bool)`

SetUseIpv6RecycleLeases sets UseIpv6RecycleLeases field to given value.

### HasUseIpv6RecycleLeases

`func (o *MemberDhcpproperties) HasUseIpv6RecycleLeases() bool`

HasUseIpv6RecycleLeases returns a boolean if a field has been set.

### GetUseIpv6UpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) GetUseIpv6UpdateDnsOnLeaseRenewal() bool`

GetUseIpv6UpdateDnsOnLeaseRenewal returns the UseIpv6UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseIpv6UpdateDnsOnLeaseRenewalOk

`func (o *MemberDhcpproperties) GetUseIpv6UpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseIpv6UpdateDnsOnLeaseRenewalOk returns a tuple with the UseIpv6UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6UpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) SetUseIpv6UpdateDnsOnLeaseRenewal(v bool)`

SetUseIpv6UpdateDnsOnLeaseRenewal sets UseIpv6UpdateDnsOnLeaseRenewal field to given value.

### HasUseIpv6UpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) HasUseIpv6UpdateDnsOnLeaseRenewal() bool`

HasUseIpv6UpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseLeasePerClientSettings

`func (o *MemberDhcpproperties) GetUseLeasePerClientSettings() bool`

GetUseLeasePerClientSettings returns the UseLeasePerClientSettings field if non-nil, zero value otherwise.

### GetUseLeasePerClientSettingsOk

`func (o *MemberDhcpproperties) GetUseLeasePerClientSettingsOk() (*bool, bool)`

GetUseLeasePerClientSettingsOk returns a tuple with the UseLeasePerClientSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeasePerClientSettings

`func (o *MemberDhcpproperties) SetUseLeasePerClientSettings(v bool)`

SetUseLeasePerClientSettings sets UseLeasePerClientSettings field to given value.

### HasUseLeasePerClientSettings

`func (o *MemberDhcpproperties) HasUseLeasePerClientSettings() bool`

HasUseLeasePerClientSettings returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *MemberDhcpproperties) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *MemberDhcpproperties) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *MemberDhcpproperties) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *MemberDhcpproperties) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogLeaseEvents

`func (o *MemberDhcpproperties) GetUseLogLeaseEvents() bool`

GetUseLogLeaseEvents returns the UseLogLeaseEvents field if non-nil, zero value otherwise.

### GetUseLogLeaseEventsOk

`func (o *MemberDhcpproperties) GetUseLogLeaseEventsOk() (*bool, bool)`

GetUseLogLeaseEventsOk returns a tuple with the UseLogLeaseEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogLeaseEvents

`func (o *MemberDhcpproperties) SetUseLogLeaseEvents(v bool)`

SetUseLogLeaseEvents sets UseLogLeaseEvents field to given value.

### HasUseLogLeaseEvents

`func (o *MemberDhcpproperties) HasUseLogLeaseEvents() bool`

HasUseLogLeaseEvents returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *MemberDhcpproperties) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *MemberDhcpproperties) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *MemberDhcpproperties) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *MemberDhcpproperties) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMicrosoftCodePage

`func (o *MemberDhcpproperties) GetUseMicrosoftCodePage() bool`

GetUseMicrosoftCodePage returns the UseMicrosoftCodePage field if non-nil, zero value otherwise.

### GetUseMicrosoftCodePageOk

`func (o *MemberDhcpproperties) GetUseMicrosoftCodePageOk() (*bool, bool)`

GetUseMicrosoftCodePageOk returns a tuple with the UseMicrosoftCodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMicrosoftCodePage

`func (o *MemberDhcpproperties) SetUseMicrosoftCodePage(v bool)`

SetUseMicrosoftCodePage sets UseMicrosoftCodePage field to given value.

### HasUseMicrosoftCodePage

`func (o *MemberDhcpproperties) HasUseMicrosoftCodePage() bool`

HasUseMicrosoftCodePage returns a boolean if a field has been set.

### GetUseNextserver

`func (o *MemberDhcpproperties) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *MemberDhcpproperties) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *MemberDhcpproperties) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *MemberDhcpproperties) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *MemberDhcpproperties) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *MemberDhcpproperties) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *MemberDhcpproperties) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *MemberDhcpproperties) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePingCount

`func (o *MemberDhcpproperties) GetUsePingCount() bool`

GetUsePingCount returns the UsePingCount field if non-nil, zero value otherwise.

### GetUsePingCountOk

`func (o *MemberDhcpproperties) GetUsePingCountOk() (*bool, bool)`

GetUsePingCountOk returns a tuple with the UsePingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePingCount

`func (o *MemberDhcpproperties) SetUsePingCount(v bool)`

SetUsePingCount sets UsePingCount field to given value.

### HasUsePingCount

`func (o *MemberDhcpproperties) HasUsePingCount() bool`

HasUsePingCount returns a boolean if a field has been set.

### GetUsePingTimeout

`func (o *MemberDhcpproperties) GetUsePingTimeout() bool`

GetUsePingTimeout returns the UsePingTimeout field if non-nil, zero value otherwise.

### GetUsePingTimeoutOk

`func (o *MemberDhcpproperties) GetUsePingTimeoutOk() (*bool, bool)`

GetUsePingTimeoutOk returns a tuple with the UsePingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePingTimeout

`func (o *MemberDhcpproperties) SetUsePingTimeout(v bool)`

SetUsePingTimeout sets UsePingTimeout field to given value.

### HasUsePingTimeout

`func (o *MemberDhcpproperties) HasUsePingTimeout() bool`

HasUsePingTimeout returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *MemberDhcpproperties) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *MemberDhcpproperties) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *MemberDhcpproperties) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *MemberDhcpproperties) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUsePrefixLengthMode

`func (o *MemberDhcpproperties) GetUsePrefixLengthMode() bool`

GetUsePrefixLengthMode returns the UsePrefixLengthMode field if non-nil, zero value otherwise.

### GetUsePrefixLengthModeOk

`func (o *MemberDhcpproperties) GetUsePrefixLengthModeOk() (*bool, bool)`

GetUsePrefixLengthModeOk returns a tuple with the UsePrefixLengthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePrefixLengthMode

`func (o *MemberDhcpproperties) SetUsePrefixLengthMode(v bool)`

SetUsePrefixLengthMode sets UsePrefixLengthMode field to given value.

### HasUsePrefixLengthMode

`func (o *MemberDhcpproperties) HasUsePrefixLengthMode() bool`

HasUsePrefixLengthMode returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *MemberDhcpproperties) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *MemberDhcpproperties) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *MemberDhcpproperties) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *MemberDhcpproperties) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *MemberDhcpproperties) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *MemberDhcpproperties) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *MemberDhcpproperties) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *MemberDhcpproperties) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseRetryDdnsUpdates

`func (o *MemberDhcpproperties) GetUseRetryDdnsUpdates() bool`

GetUseRetryDdnsUpdates returns the UseRetryDdnsUpdates field if non-nil, zero value otherwise.

### GetUseRetryDdnsUpdatesOk

`func (o *MemberDhcpproperties) GetUseRetryDdnsUpdatesOk() (*bool, bool)`

GetUseRetryDdnsUpdatesOk returns a tuple with the UseRetryDdnsUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRetryDdnsUpdates

`func (o *MemberDhcpproperties) SetUseRetryDdnsUpdates(v bool)`

SetUseRetryDdnsUpdates sets UseRetryDdnsUpdates field to given value.

### HasUseRetryDdnsUpdates

`func (o *MemberDhcpproperties) HasUseRetryDdnsUpdates() bool`

HasUseRetryDdnsUpdates returns a boolean if a field has been set.

### GetUseSyslogFacility

`func (o *MemberDhcpproperties) GetUseSyslogFacility() bool`

GetUseSyslogFacility returns the UseSyslogFacility field if non-nil, zero value otherwise.

### GetUseSyslogFacilityOk

`func (o *MemberDhcpproperties) GetUseSyslogFacilityOk() (*bool, bool)`

GetUseSyslogFacilityOk returns a tuple with the UseSyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSyslogFacility

`func (o *MemberDhcpproperties) SetUseSyslogFacility(v bool)`

SetUseSyslogFacility sets UseSyslogFacility field to given value.

### HasUseSyslogFacility

`func (o *MemberDhcpproperties) HasUseSyslogFacility() bool`

HasUseSyslogFacility returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *MemberDhcpproperties) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *MemberDhcpproperties) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *MemberDhcpproperties) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *MemberDhcpproperties) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *MemberDhcpproperties) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *MemberDhcpproperties) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *MemberDhcpproperties) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *MemberDhcpproperties) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *MemberDhcpproperties) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *MemberDhcpproperties) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


