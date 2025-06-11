# GridDhcpproperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Authority** | Pointer to **bool** | The Grid-level authority flag. This flag specifies whether a DHCP server is authoritative for a domain. | [optional] 
**Bootfile** | Pointer to **string** | The name of a file that DHCP clients need to boot. Some DHCP clients use BOOTP (bootstrap protocol) or include the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The name of the server on which a boot file is stored. | [optional] 
**CaptureHostname** | Pointer to **bool** | The Grid-level capture hostname flag. Set this flag to capture the hostname and lease time when assigning a fixed address. | [optional] 
**DdnsDomainname** | Pointer to **string** | The member DDNS domain name value. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | Determines if the ability of a DHCP server to generate a host name and update DNS with this host name when it receives a DHCP REQUEST message that does not include a host name is enabled or not. | [optional] 
**DdnsRetryInterval** | Pointer to **int64** | Determines the retry interval when the DHCP server makes repeated attempts to send DDNS updates to a DNS server. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | Determines that only the DHCP server is allowed to update DNS, regardless of the requests from the DHCP clients. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DDNS TTL (Dynamic DNS Time To Live) value specifies the number of seconds an IP address for the name is cached. | [optional] 
**DdnsUpdateFixedAddresses** | Pointer to **bool** | Determines if the Grid DHCP server&#39;s ability to update the A and PTR records with a fixed address is enabled or not. | [optional] 
**DdnsUseOption81** | Pointer to **bool** | Determines if support for option 81 is enabled or not. | [optional] 
**DenyBootp** | Pointer to **bool** | Determines if deny BOOTP is enabled or not. | [optional] 
**DisableAllNacFilters** | Pointer to **bool** | If set to True, NAC filters will be disabled on the Infoblox Grid. | [optional] 
**DnsUpdateStyle** | Pointer to **string** | The update style for dynamic DNS updates. | [optional] 
**EmailList** | Pointer to **[]string** | The Grid-level email_list value. Specify an e-mail address to which you want the Infoblox appliance to send e-mail notifications when the DHCP address usage for the grid crosses a threshold. You can create a list of several e-mail addresses. | [optional] 
**EnableDdns** | Pointer to **bool** | Determines if the member DHCP server&#39;s ability to send DDNS updates is enabled or not. | [optional] 
**EnableDhcpThresholds** | Pointer to **bool** | Represents the watermarks above or below which address usage in a network is unexpected and might warrant your attention. | [optional] 
**EnableEmailWarnings** | Pointer to **bool** | Determines if e-mail warnings are enabled or disabled. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an e-mail notification to an administrator. | [optional] 
**EnableFingerprint** | Pointer to **bool** | Determines if the fingerprint feature is enabled or not. If you enable this feature, the server will match a fingerprint for incoming lease requests. | [optional] 
**EnableGssTsig** | Pointer to **bool** | Determines whether all appliances are enabled to receive GSS-TSIG authenticated updates from DHCP clients. | [optional] 
**EnableHostnameRewrite** | Pointer to **bool** | Determines if the Grid-level host name rewrite feature is enabled or not. | [optional] 
**EnableLeasequery** | Pointer to **bool** | Determines if lease query is allowed or not. | [optional] 
**EnableRoamingHosts** | Pointer to **bool** | Determines if DHCP servers in a Grid support roaming hosts or not. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determined if the SNMP warnings on Grid-level are enabled or not. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an SNMP trap to the trap receiver that you defined you defined at the Grid member level. | [optional] 
**FormatLogOption82** | Pointer to **string** | The format option for Option 82 logging. | [optional] 
**Grid** | Pointer to **string** | Determines the Grid that serves DHCP. This specifies a group of Infoblox appliances that are connected together to provide a single point of device administration and service configuration in a secure, highly available environment. | [optional] [readonly] 
**GssTsigKeys** | Pointer to **[]string** | The list of GSS-TSIG keys for a Grid DHCP object. | [optional] 
**HighWaterMark** | Pointer to **int64** | Determines the high watermark value of a Grid DHCP server. If the percentage of allocated addresses exceeds this watermark, the appliance makes a syslog entry and sends an e-mail notification (if enabled). Specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | Determines the high watermark reset value of a member DHCP server. If the percentage of allocated addresses drops below this value, a corresponding SNMP trap is reset. Specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**HostnameRewritePolicy** | Pointer to **string** | The name of the default hostname rewrite policy, which is also in the protocol_hostname_rewrite_policies array. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | Determines if the ignore DHCP option list request flag of a Grid DHCP is enabled or not. If this flag is set to true all available DHCP options will be returned to the client. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \&quot;NONE\&quot;, \&quot;CLIENT\&quot;, or \&quot;MACADDR\&quot;. The default is \&quot;NONE\&quot;. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**ImmediateFaConfiguration** | Pointer to **bool** | Determines if the fixed address configuration takes effect immediately without DHCP service restart or not. | [optional] 
**Ipv6CaptureHostname** | Pointer to **bool** | Determines if the IPv6 host name and lease time is captured or not while assigning a fixed address. | [optional] 
**Ipv6DdnsDomainname** | Pointer to **string** | The Grid-level DDNS domain name value. | [optional] 
**Ipv6DdnsEnableOptionFqdn** | Pointer to **bool** | Controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. | [optional] 
**Ipv6DdnsServerAlwaysUpdates** | Pointer to **bool** | Determines if the server always updates DNS or updates only if requested by the client. | [optional] 
**Ipv6DdnsTtl** | Pointer to **int64** | The Grid-level IPv6 DDNS TTL value. | [optional] 
**Ipv6DefaultPrefix** | Pointer to **string** | The Grid-level IPv6 default prefix. | [optional] 
**Ipv6DnsUpdateStyle** | Pointer to **string** | The update style for dynamic DHCPv6 DNS updates. | [optional] 
**Ipv6DomainName** | Pointer to **string** | The IPv6 domain name. | [optional] 
**Ipv6DomainNameServers** | Pointer to **[]string** | The comma separated list of domain name server addresses in IPv6 address format. | [optional] 
**Ipv6EnableDdns** | Pointer to **bool** | Determines if sending DDNS updates by the DHCPv6 server is enabled or not. | [optional] 
**Ipv6EnableGssTsig** | Pointer to **bool** | Determines whether the all appliances are enabled to receive GSS-TSIG authenticated updates from DHCPv6 clients. | [optional] 
**Ipv6EnableLeaseScavenging** | Pointer to **bool** | Indicates whether DHCPv6 lease scavenging is enabled or disabled. | [optional] 
**Ipv6EnableRetryUpdates** | Pointer to **bool** | Determines if the DHCPv6 server retries failed dynamic DNS updates or not. | [optional] 
**Ipv6GenerateHostname** | Pointer to **bool** | Determines if the server generates the hostname if it is not sent by the client. | [optional] 
**Ipv6GssTsigKeys** | Pointer to **[]string** | The list of GSS-TSIG keys for a Grid DHCPv6 object. | [optional] 
**Ipv6KdcServer** | Pointer to **string** | The IPv6 address or FQDN of the Kerberos server for DHCPv6 GSS-TSIG authentication. | [optional] 
**Ipv6LeaseScavengingTime** | Pointer to **int64** | The Grid-level grace period (in seconds) to keep an expired lease before it is deleted by the scavenging process. | [optional] 
**Ipv6MicrosoftCodePage** | Pointer to **string** | The Grid-level Microsoft client DHCP IPv6 code page value. This value is the hostname translation code page for Microsoft DHCP IPv6 clients. | [optional] 
**Ipv6Options** | Pointer to [**[]GridDhcppropertiesIpv6Options**](GridDhcppropertiesIpv6Options.md) | An array of DHCP option dhcpoption structs that lists the DHCPv6 options associated with the object. | [optional] 
**Ipv6Prefixes** | Pointer to **[]string** | The Grid-level list of IPv6 prefixes. | [optional] 
**Ipv6RecycleLeases** | Pointer to **bool** | Determines if the IPv6 recycle leases feature is enabled or not. If the feature is enabled, leases are kept in the Recycle Bin until one week after expiration. When the feature is disabled, the leases are irrecoverably deleted. | [optional] 
**Ipv6RememberExpiredClientAssociation** | Pointer to **bool** | Enable binding for expired DHCPv6 leases. | [optional] 
**Ipv6RetryUpdatesInterval** | Pointer to **int64** | Determines the retry interval when the member DHCPv6 server makes repeated attempts to send DDNS updates to a DNS server. | [optional] 
**Ipv6TxtRecordHandling** | Pointer to **string** | The Grid-level TXT record handling value. This value specifies how DHCPv6 should treat the TXT records when performing DNS updates. | [optional] 
**Ipv6UpdateDnsOnLeaseRenewal** | Pointer to **bool** | Controls whether the DHCPv6 server updates DNS when an IPv6 DHCP lease is renewed. | [optional] 
**KdcServer** | Pointer to **string** | The IPv4 address or FQDN of the Kerberos server for DHCPv4 GSS-TSIG authentication. | [optional] 
**LeaseLoggingMember** | Pointer to **string** | The Grid member on which you want to store the DHCP lease history log. Infoblox recommends that you dedicate a member other than the master as a logging member. If possible, use this member solely for storing the DHCP lease history log. If you do not select a member, no logging can occur. | [optional] 
**LeasePerClientSettings** | Pointer to **string** | Defines how the appliance releases DHCP leases. Valid values are \&quot;RELEASE_MACHING_ID\&quot;, \&quot;NEVER_RELEASE\&quot;, or \&quot;ONE_LEASE_PER_CLIENT\&quot;. The default is \&quot;RELEASE_MATCHING_ID\&quot;. | [optional] 
**LeaseScavengeTime** | Pointer to **int32** | Determines the lease scavenging time value. When this field is set, the appliance permanently deletes the free and backup leases, that remain in the database beyond a specified period of time. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogLeaseEvents** | Pointer to **bool** | This value specifies whether the Grid DHCP members log lease events is enabled or not. | [optional] 
**LogicFilterRules** | Pointer to [**[]GridDhcppropertiesLogicFilterRules**](GridDhcppropertiesLogicFilterRules.md) | This field contains the logic filters to be applied on the Infoblox Grid. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**LowWaterMark** | Pointer to **int64** | Determines the low watermark value. If the percent of allocated addresses drops below this watermark, the appliance makes a syslog entry and if enabled, sends an e-mail notification. | [optional] 
**LowWaterMarkReset** | Pointer to **int64** | Determines the low watermark reset value.If the percentage of allocated addresses exceeds this value, a corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value. | [optional] 
**MicrosoftCodePage** | Pointer to **string** | The Microsoft client DHCP IPv4 code page value of a Grid. This value is the hostname translation code page for Microsoft DHCP IPv4 clients. | [optional] 
**Nextserver** | Pointer to **string** | The next server value of a DHCP server. This value is the IP address or name of the boot file server on which the boot file is stored. | [optional] 
**Option60MatchRules** | Pointer to [**[]GridDhcppropertiesOption60MatchRules**](GridDhcppropertiesOption60MatchRules.md) | The list of option 60 match rules. | [optional] 
**Options** | Pointer to [**[]GridDhcppropertiesOptions**](GridDhcppropertiesOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. Note that WAPI does not return special options &#39;routers&#39;, &#39;domain-name-servers&#39;, &#39;domain-name&#39; and &#39;broadcast-address&#39; with empty values for this object. | [optional] 
**PingCount** | Pointer to **int64** | Specifies the number of pings that the Infoblox appliance sends to an IP address to verify that it is not in use. Values are range is from 0 to 10, where 0 disables pings. | [optional] 
**PingTimeout** | Pointer to **int64** | Indicates the number of milliseconds the appliance waits for a response to its ping. Valid values are 100, 500, 1000, 2000, 3000, 4000 and 5000 milliseconds. | [optional] 
**PreferredLifetime** | Pointer to **int64** | The preferred lifetime value. | [optional] 
**PrefixLengthMode** | Pointer to **string** | The Prefix length mode for DHCPv6. | [optional] 
**ProtocolHostnameRewritePolicies** | Pointer to **[]string** | The list of hostname rewrite policies. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | Specifies the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**RecycleLeases** | Pointer to **bool** | Determines if the recycle leases feature is enabled or not. If you enabled this feature, and then delete a DHCP range, the appliance stores active leases from this range up to one week after the leases expires. | [optional] 
**RestartSetting** | Pointer to [**GridDhcppropertiesRestartSetting**](GridDhcppropertiesRestartSetting.md) |  | [optional] 
**RetryDdnsUpdates** | Pointer to **bool** | Indicates whether the DHCP server makes repeated attempts to send DDNS updates to a DNS server. | [optional] 
**SyslogFacility** | Pointer to **string** | The syslog facility is the location on the syslog server to which you want to sort the syslog messages. | [optional] 
**TxtRecordHandling** | Pointer to **string** | The Grid-level TXT record handling value. This value specifies how DHCP should treat the TXT records when performing DNS updates. | [optional] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | Controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**ValidLifetime** | Pointer to **int64** | The valid lifetime for the Grid members. | [optional] 

## Methods

### NewGridDhcpproperties

`func NewGridDhcpproperties() *GridDhcpproperties`

NewGridDhcpproperties instantiates a new GridDhcpproperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDhcppropertiesWithDefaults

`func NewGridDhcppropertiesWithDefaults() *GridDhcpproperties`

NewGridDhcppropertiesWithDefaults instantiates a new GridDhcpproperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridDhcpproperties) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridDhcpproperties) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridDhcpproperties) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridDhcpproperties) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthority

`func (o *GridDhcpproperties) GetAuthority() bool`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GridDhcpproperties) GetAuthorityOk() (*bool, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GridDhcpproperties) SetAuthority(v bool)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GridDhcpproperties) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetBootfile

`func (o *GridDhcpproperties) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GridDhcpproperties) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GridDhcpproperties) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GridDhcpproperties) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GridDhcpproperties) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GridDhcpproperties) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GridDhcpproperties) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GridDhcpproperties) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCaptureHostname

`func (o *GridDhcpproperties) GetCaptureHostname() bool`

GetCaptureHostname returns the CaptureHostname field if non-nil, zero value otherwise.

### GetCaptureHostnameOk

`func (o *GridDhcpproperties) GetCaptureHostnameOk() (*bool, bool)`

GetCaptureHostnameOk returns a tuple with the CaptureHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureHostname

`func (o *GridDhcpproperties) SetCaptureHostname(v bool)`

SetCaptureHostname sets CaptureHostname field to given value.

### HasCaptureHostname

`func (o *GridDhcpproperties) HasCaptureHostname() bool`

HasCaptureHostname returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GridDhcpproperties) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GridDhcpproperties) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GridDhcpproperties) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GridDhcpproperties) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GridDhcpproperties) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GridDhcpproperties) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GridDhcpproperties) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GridDhcpproperties) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsRetryInterval

`func (o *GridDhcpproperties) GetDdnsRetryInterval() int64`

GetDdnsRetryInterval returns the DdnsRetryInterval field if non-nil, zero value otherwise.

### GetDdnsRetryIntervalOk

`func (o *GridDhcpproperties) GetDdnsRetryIntervalOk() (*int64, bool)`

GetDdnsRetryIntervalOk returns a tuple with the DdnsRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRetryInterval

`func (o *GridDhcpproperties) SetDdnsRetryInterval(v int64)`

SetDdnsRetryInterval sets DdnsRetryInterval field to given value.

### HasDdnsRetryInterval

`func (o *GridDhcpproperties) HasDdnsRetryInterval() bool`

HasDdnsRetryInterval returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *GridDhcpproperties) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *GridDhcpproperties) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *GridDhcpproperties) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *GridDhcpproperties) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *GridDhcpproperties) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *GridDhcpproperties) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *GridDhcpproperties) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *GridDhcpproperties) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUpdateFixedAddresses

`func (o *GridDhcpproperties) GetDdnsUpdateFixedAddresses() bool`

GetDdnsUpdateFixedAddresses returns the DdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetDdnsUpdateFixedAddressesOk

`func (o *GridDhcpproperties) GetDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetDdnsUpdateFixedAddressesOk returns a tuple with the DdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateFixedAddresses

`func (o *GridDhcpproperties) SetDdnsUpdateFixedAddresses(v bool)`

SetDdnsUpdateFixedAddresses sets DdnsUpdateFixedAddresses field to given value.

### HasDdnsUpdateFixedAddresses

`func (o *GridDhcpproperties) HasDdnsUpdateFixedAddresses() bool`

HasDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *GridDhcpproperties) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *GridDhcpproperties) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *GridDhcpproperties) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *GridDhcpproperties) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GridDhcpproperties) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GridDhcpproperties) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GridDhcpproperties) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GridDhcpproperties) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDisableAllNacFilters

`func (o *GridDhcpproperties) GetDisableAllNacFilters() bool`

GetDisableAllNacFilters returns the DisableAllNacFilters field if non-nil, zero value otherwise.

### GetDisableAllNacFiltersOk

`func (o *GridDhcpproperties) GetDisableAllNacFiltersOk() (*bool, bool)`

GetDisableAllNacFiltersOk returns a tuple with the DisableAllNacFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAllNacFilters

`func (o *GridDhcpproperties) SetDisableAllNacFilters(v bool)`

SetDisableAllNacFilters sets DisableAllNacFilters field to given value.

### HasDisableAllNacFilters

`func (o *GridDhcpproperties) HasDisableAllNacFilters() bool`

HasDisableAllNacFilters returns a boolean if a field has been set.

### GetDnsUpdateStyle

`func (o *GridDhcpproperties) GetDnsUpdateStyle() string`

GetDnsUpdateStyle returns the DnsUpdateStyle field if non-nil, zero value otherwise.

### GetDnsUpdateStyleOk

`func (o *GridDhcpproperties) GetDnsUpdateStyleOk() (*string, bool)`

GetDnsUpdateStyleOk returns a tuple with the DnsUpdateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsUpdateStyle

`func (o *GridDhcpproperties) SetDnsUpdateStyle(v string)`

SetDnsUpdateStyle sets DnsUpdateStyle field to given value.

### HasDnsUpdateStyle

`func (o *GridDhcpproperties) HasDnsUpdateStyle() bool`

HasDnsUpdateStyle returns a boolean if a field has been set.

### GetEmailList

`func (o *GridDhcpproperties) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *GridDhcpproperties) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *GridDhcpproperties) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *GridDhcpproperties) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GridDhcpproperties) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GridDhcpproperties) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GridDhcpproperties) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GridDhcpproperties) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *GridDhcpproperties) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *GridDhcpproperties) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *GridDhcpproperties) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *GridDhcpproperties) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *GridDhcpproperties) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *GridDhcpproperties) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *GridDhcpproperties) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *GridDhcpproperties) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableFingerprint

`func (o *GridDhcpproperties) GetEnableFingerprint() bool`

GetEnableFingerprint returns the EnableFingerprint field if non-nil, zero value otherwise.

### GetEnableFingerprintOk

`func (o *GridDhcpproperties) GetEnableFingerprintOk() (*bool, bool)`

GetEnableFingerprintOk returns a tuple with the EnableFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFingerprint

`func (o *GridDhcpproperties) SetEnableFingerprint(v bool)`

SetEnableFingerprint sets EnableFingerprint field to given value.

### HasEnableFingerprint

`func (o *GridDhcpproperties) HasEnableFingerprint() bool`

HasEnableFingerprint returns a boolean if a field has been set.

### GetEnableGssTsig

`func (o *GridDhcpproperties) GetEnableGssTsig() bool`

GetEnableGssTsig returns the EnableGssTsig field if non-nil, zero value otherwise.

### GetEnableGssTsigOk

`func (o *GridDhcpproperties) GetEnableGssTsigOk() (*bool, bool)`

GetEnableGssTsigOk returns a tuple with the EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGssTsig

`func (o *GridDhcpproperties) SetEnableGssTsig(v bool)`

SetEnableGssTsig sets EnableGssTsig field to given value.

### HasEnableGssTsig

`func (o *GridDhcpproperties) HasEnableGssTsig() bool`

HasEnableGssTsig returns a boolean if a field has been set.

### GetEnableHostnameRewrite

`func (o *GridDhcpproperties) GetEnableHostnameRewrite() bool`

GetEnableHostnameRewrite returns the EnableHostnameRewrite field if non-nil, zero value otherwise.

### GetEnableHostnameRewriteOk

`func (o *GridDhcpproperties) GetEnableHostnameRewriteOk() (*bool, bool)`

GetEnableHostnameRewriteOk returns a tuple with the EnableHostnameRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHostnameRewrite

`func (o *GridDhcpproperties) SetEnableHostnameRewrite(v bool)`

SetEnableHostnameRewrite sets EnableHostnameRewrite field to given value.

### HasEnableHostnameRewrite

`func (o *GridDhcpproperties) HasEnableHostnameRewrite() bool`

HasEnableHostnameRewrite returns a boolean if a field has been set.

### GetEnableLeasequery

`func (o *GridDhcpproperties) GetEnableLeasequery() bool`

GetEnableLeasequery returns the EnableLeasequery field if non-nil, zero value otherwise.

### GetEnableLeasequeryOk

`func (o *GridDhcpproperties) GetEnableLeasequeryOk() (*bool, bool)`

GetEnableLeasequeryOk returns a tuple with the EnableLeasequery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLeasequery

`func (o *GridDhcpproperties) SetEnableLeasequery(v bool)`

SetEnableLeasequery sets EnableLeasequery field to given value.

### HasEnableLeasequery

`func (o *GridDhcpproperties) HasEnableLeasequery() bool`

HasEnableLeasequery returns a boolean if a field has been set.

### GetEnableRoamingHosts

`func (o *GridDhcpproperties) GetEnableRoamingHosts() bool`

GetEnableRoamingHosts returns the EnableRoamingHosts field if non-nil, zero value otherwise.

### GetEnableRoamingHostsOk

`func (o *GridDhcpproperties) GetEnableRoamingHostsOk() (*bool, bool)`

GetEnableRoamingHostsOk returns a tuple with the EnableRoamingHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRoamingHosts

`func (o *GridDhcpproperties) SetEnableRoamingHosts(v bool)`

SetEnableRoamingHosts sets EnableRoamingHosts field to given value.

### HasEnableRoamingHosts

`func (o *GridDhcpproperties) HasEnableRoamingHosts() bool`

HasEnableRoamingHosts returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *GridDhcpproperties) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *GridDhcpproperties) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *GridDhcpproperties) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *GridDhcpproperties) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetFormatLogOption82

`func (o *GridDhcpproperties) GetFormatLogOption82() string`

GetFormatLogOption82 returns the FormatLogOption82 field if non-nil, zero value otherwise.

### GetFormatLogOption82Ok

`func (o *GridDhcpproperties) GetFormatLogOption82Ok() (*string, bool)`

GetFormatLogOption82Ok returns a tuple with the FormatLogOption82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatLogOption82

`func (o *GridDhcpproperties) SetFormatLogOption82(v string)`

SetFormatLogOption82 sets FormatLogOption82 field to given value.

### HasFormatLogOption82

`func (o *GridDhcpproperties) HasFormatLogOption82() bool`

HasFormatLogOption82 returns a boolean if a field has been set.

### GetGrid

`func (o *GridDhcpproperties) GetGrid() string`

GetGrid returns the Grid field if non-nil, zero value otherwise.

### GetGridOk

`func (o *GridDhcpproperties) GetGridOk() (*string, bool)`

GetGridOk returns a tuple with the Grid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrid

`func (o *GridDhcpproperties) SetGrid(v string)`

SetGrid sets Grid field to given value.

### HasGrid

`func (o *GridDhcpproperties) HasGrid() bool`

HasGrid returns a boolean if a field has been set.

### GetGssTsigKeys

`func (o *GridDhcpproperties) GetGssTsigKeys() []string`

GetGssTsigKeys returns the GssTsigKeys field if non-nil, zero value otherwise.

### GetGssTsigKeysOk

`func (o *GridDhcpproperties) GetGssTsigKeysOk() (*[]string, bool)`

GetGssTsigKeysOk returns a tuple with the GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigKeys

`func (o *GridDhcpproperties) SetGssTsigKeys(v []string)`

SetGssTsigKeys sets GssTsigKeys field to given value.

### HasGssTsigKeys

`func (o *GridDhcpproperties) HasGssTsigKeys() bool`

HasGssTsigKeys returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GridDhcpproperties) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GridDhcpproperties) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GridDhcpproperties) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GridDhcpproperties) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *GridDhcpproperties) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *GridDhcpproperties) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *GridDhcpproperties) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *GridDhcpproperties) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetHostnameRewritePolicy

`func (o *GridDhcpproperties) GetHostnameRewritePolicy() string`

GetHostnameRewritePolicy returns the HostnameRewritePolicy field if non-nil, zero value otherwise.

### GetHostnameRewritePolicyOk

`func (o *GridDhcpproperties) GetHostnameRewritePolicyOk() (*string, bool)`

GetHostnameRewritePolicyOk returns a tuple with the HostnameRewritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewritePolicy

`func (o *GridDhcpproperties) SetHostnameRewritePolicy(v string)`

SetHostnameRewritePolicy sets HostnameRewritePolicy field to given value.

### HasHostnameRewritePolicy

`func (o *GridDhcpproperties) HasHostnameRewritePolicy() bool`

HasHostnameRewritePolicy returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GridDhcpproperties) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GridDhcpproperties) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GridDhcpproperties) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GridDhcpproperties) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *GridDhcpproperties) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *GridDhcpproperties) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *GridDhcpproperties) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *GridDhcpproperties) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *GridDhcpproperties) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *GridDhcpproperties) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *GridDhcpproperties) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *GridDhcpproperties) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetImmediateFaConfiguration

`func (o *GridDhcpproperties) GetImmediateFaConfiguration() bool`

GetImmediateFaConfiguration returns the ImmediateFaConfiguration field if non-nil, zero value otherwise.

### GetImmediateFaConfigurationOk

`func (o *GridDhcpproperties) GetImmediateFaConfigurationOk() (*bool, bool)`

GetImmediateFaConfigurationOk returns a tuple with the ImmediateFaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFaConfiguration

`func (o *GridDhcpproperties) SetImmediateFaConfiguration(v bool)`

SetImmediateFaConfiguration sets ImmediateFaConfiguration field to given value.

### HasImmediateFaConfiguration

`func (o *GridDhcpproperties) HasImmediateFaConfiguration() bool`

HasImmediateFaConfiguration returns a boolean if a field has been set.

### GetIpv6CaptureHostname

`func (o *GridDhcpproperties) GetIpv6CaptureHostname() bool`

GetIpv6CaptureHostname returns the Ipv6CaptureHostname field if non-nil, zero value otherwise.

### GetIpv6CaptureHostnameOk

`func (o *GridDhcpproperties) GetIpv6CaptureHostnameOk() (*bool, bool)`

GetIpv6CaptureHostnameOk returns a tuple with the Ipv6CaptureHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CaptureHostname

`func (o *GridDhcpproperties) SetIpv6CaptureHostname(v bool)`

SetIpv6CaptureHostname sets Ipv6CaptureHostname field to given value.

### HasIpv6CaptureHostname

`func (o *GridDhcpproperties) HasIpv6CaptureHostname() bool`

HasIpv6CaptureHostname returns a boolean if a field has been set.

### GetIpv6DdnsDomainname

`func (o *GridDhcpproperties) GetIpv6DdnsDomainname() string`

GetIpv6DdnsDomainname returns the Ipv6DdnsDomainname field if non-nil, zero value otherwise.

### GetIpv6DdnsDomainnameOk

`func (o *GridDhcpproperties) GetIpv6DdnsDomainnameOk() (*string, bool)`

GetIpv6DdnsDomainnameOk returns a tuple with the Ipv6DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsDomainname

`func (o *GridDhcpproperties) SetIpv6DdnsDomainname(v string)`

SetIpv6DdnsDomainname sets Ipv6DdnsDomainname field to given value.

### HasIpv6DdnsDomainname

`func (o *GridDhcpproperties) HasIpv6DdnsDomainname() bool`

HasIpv6DdnsDomainname returns a boolean if a field has been set.

### GetIpv6DdnsEnableOptionFqdn

`func (o *GridDhcpproperties) GetIpv6DdnsEnableOptionFqdn() bool`

GetIpv6DdnsEnableOptionFqdn returns the Ipv6DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetIpv6DdnsEnableOptionFqdnOk

`func (o *GridDhcpproperties) GetIpv6DdnsEnableOptionFqdnOk() (*bool, bool)`

GetIpv6DdnsEnableOptionFqdnOk returns a tuple with the Ipv6DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsEnableOptionFqdn

`func (o *GridDhcpproperties) SetIpv6DdnsEnableOptionFqdn(v bool)`

SetIpv6DdnsEnableOptionFqdn sets Ipv6DdnsEnableOptionFqdn field to given value.

### HasIpv6DdnsEnableOptionFqdn

`func (o *GridDhcpproperties) HasIpv6DdnsEnableOptionFqdn() bool`

HasIpv6DdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetIpv6DdnsServerAlwaysUpdates

`func (o *GridDhcpproperties) GetIpv6DdnsServerAlwaysUpdates() bool`

GetIpv6DdnsServerAlwaysUpdates returns the Ipv6DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetIpv6DdnsServerAlwaysUpdatesOk

`func (o *GridDhcpproperties) GetIpv6DdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetIpv6DdnsServerAlwaysUpdatesOk returns a tuple with the Ipv6DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsServerAlwaysUpdates

`func (o *GridDhcpproperties) SetIpv6DdnsServerAlwaysUpdates(v bool)`

SetIpv6DdnsServerAlwaysUpdates sets Ipv6DdnsServerAlwaysUpdates field to given value.

### HasIpv6DdnsServerAlwaysUpdates

`func (o *GridDhcpproperties) HasIpv6DdnsServerAlwaysUpdates() bool`

HasIpv6DdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetIpv6DdnsTtl

`func (o *GridDhcpproperties) GetIpv6DdnsTtl() int64`

GetIpv6DdnsTtl returns the Ipv6DdnsTtl field if non-nil, zero value otherwise.

### GetIpv6DdnsTtlOk

`func (o *GridDhcpproperties) GetIpv6DdnsTtlOk() (*int64, bool)`

GetIpv6DdnsTtlOk returns a tuple with the Ipv6DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DdnsTtl

`func (o *GridDhcpproperties) SetIpv6DdnsTtl(v int64)`

SetIpv6DdnsTtl sets Ipv6DdnsTtl field to given value.

### HasIpv6DdnsTtl

`func (o *GridDhcpproperties) HasIpv6DdnsTtl() bool`

HasIpv6DdnsTtl returns a boolean if a field has been set.

### GetIpv6DefaultPrefix

`func (o *GridDhcpproperties) GetIpv6DefaultPrefix() string`

GetIpv6DefaultPrefix returns the Ipv6DefaultPrefix field if non-nil, zero value otherwise.

### GetIpv6DefaultPrefixOk

`func (o *GridDhcpproperties) GetIpv6DefaultPrefixOk() (*string, bool)`

GetIpv6DefaultPrefixOk returns a tuple with the Ipv6DefaultPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DefaultPrefix

`func (o *GridDhcpproperties) SetIpv6DefaultPrefix(v string)`

SetIpv6DefaultPrefix sets Ipv6DefaultPrefix field to given value.

### HasIpv6DefaultPrefix

`func (o *GridDhcpproperties) HasIpv6DefaultPrefix() bool`

HasIpv6DefaultPrefix returns a boolean if a field has been set.

### GetIpv6DnsUpdateStyle

`func (o *GridDhcpproperties) GetIpv6DnsUpdateStyle() string`

GetIpv6DnsUpdateStyle returns the Ipv6DnsUpdateStyle field if non-nil, zero value otherwise.

### GetIpv6DnsUpdateStyleOk

`func (o *GridDhcpproperties) GetIpv6DnsUpdateStyleOk() (*string, bool)`

GetIpv6DnsUpdateStyleOk returns a tuple with the Ipv6DnsUpdateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DnsUpdateStyle

`func (o *GridDhcpproperties) SetIpv6DnsUpdateStyle(v string)`

SetIpv6DnsUpdateStyle sets Ipv6DnsUpdateStyle field to given value.

### HasIpv6DnsUpdateStyle

`func (o *GridDhcpproperties) HasIpv6DnsUpdateStyle() bool`

HasIpv6DnsUpdateStyle returns a boolean if a field has been set.

### GetIpv6DomainName

`func (o *GridDhcpproperties) GetIpv6DomainName() string`

GetIpv6DomainName returns the Ipv6DomainName field if non-nil, zero value otherwise.

### GetIpv6DomainNameOk

`func (o *GridDhcpproperties) GetIpv6DomainNameOk() (*string, bool)`

GetIpv6DomainNameOk returns a tuple with the Ipv6DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DomainName

`func (o *GridDhcpproperties) SetIpv6DomainName(v string)`

SetIpv6DomainName sets Ipv6DomainName field to given value.

### HasIpv6DomainName

`func (o *GridDhcpproperties) HasIpv6DomainName() bool`

HasIpv6DomainName returns a boolean if a field has been set.

### GetIpv6DomainNameServers

`func (o *GridDhcpproperties) GetIpv6DomainNameServers() []string`

GetIpv6DomainNameServers returns the Ipv6DomainNameServers field if non-nil, zero value otherwise.

### GetIpv6DomainNameServersOk

`func (o *GridDhcpproperties) GetIpv6DomainNameServersOk() (*[]string, bool)`

GetIpv6DomainNameServersOk returns a tuple with the Ipv6DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6DomainNameServers

`func (o *GridDhcpproperties) SetIpv6DomainNameServers(v []string)`

SetIpv6DomainNameServers sets Ipv6DomainNameServers field to given value.

### HasIpv6DomainNameServers

`func (o *GridDhcpproperties) HasIpv6DomainNameServers() bool`

HasIpv6DomainNameServers returns a boolean if a field has been set.

### GetIpv6EnableDdns

`func (o *GridDhcpproperties) GetIpv6EnableDdns() bool`

GetIpv6EnableDdns returns the Ipv6EnableDdns field if non-nil, zero value otherwise.

### GetIpv6EnableDdnsOk

`func (o *GridDhcpproperties) GetIpv6EnableDdnsOk() (*bool, bool)`

GetIpv6EnableDdnsOk returns a tuple with the Ipv6EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableDdns

`func (o *GridDhcpproperties) SetIpv6EnableDdns(v bool)`

SetIpv6EnableDdns sets Ipv6EnableDdns field to given value.

### HasIpv6EnableDdns

`func (o *GridDhcpproperties) HasIpv6EnableDdns() bool`

HasIpv6EnableDdns returns a boolean if a field has been set.

### GetIpv6EnableGssTsig

`func (o *GridDhcpproperties) GetIpv6EnableGssTsig() bool`

GetIpv6EnableGssTsig returns the Ipv6EnableGssTsig field if non-nil, zero value otherwise.

### GetIpv6EnableGssTsigOk

`func (o *GridDhcpproperties) GetIpv6EnableGssTsigOk() (*bool, bool)`

GetIpv6EnableGssTsigOk returns a tuple with the Ipv6EnableGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableGssTsig

`func (o *GridDhcpproperties) SetIpv6EnableGssTsig(v bool)`

SetIpv6EnableGssTsig sets Ipv6EnableGssTsig field to given value.

### HasIpv6EnableGssTsig

`func (o *GridDhcpproperties) HasIpv6EnableGssTsig() bool`

HasIpv6EnableGssTsig returns a boolean if a field has been set.

### GetIpv6EnableLeaseScavenging

`func (o *GridDhcpproperties) GetIpv6EnableLeaseScavenging() bool`

GetIpv6EnableLeaseScavenging returns the Ipv6EnableLeaseScavenging field if non-nil, zero value otherwise.

### GetIpv6EnableLeaseScavengingOk

`func (o *GridDhcpproperties) GetIpv6EnableLeaseScavengingOk() (*bool, bool)`

GetIpv6EnableLeaseScavengingOk returns a tuple with the Ipv6EnableLeaseScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableLeaseScavenging

`func (o *GridDhcpproperties) SetIpv6EnableLeaseScavenging(v bool)`

SetIpv6EnableLeaseScavenging sets Ipv6EnableLeaseScavenging field to given value.

### HasIpv6EnableLeaseScavenging

`func (o *GridDhcpproperties) HasIpv6EnableLeaseScavenging() bool`

HasIpv6EnableLeaseScavenging returns a boolean if a field has been set.

### GetIpv6EnableRetryUpdates

`func (o *GridDhcpproperties) GetIpv6EnableRetryUpdates() bool`

GetIpv6EnableRetryUpdates returns the Ipv6EnableRetryUpdates field if non-nil, zero value otherwise.

### GetIpv6EnableRetryUpdatesOk

`func (o *GridDhcpproperties) GetIpv6EnableRetryUpdatesOk() (*bool, bool)`

GetIpv6EnableRetryUpdatesOk returns a tuple with the Ipv6EnableRetryUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EnableRetryUpdates

`func (o *GridDhcpproperties) SetIpv6EnableRetryUpdates(v bool)`

SetIpv6EnableRetryUpdates sets Ipv6EnableRetryUpdates field to given value.

### HasIpv6EnableRetryUpdates

`func (o *GridDhcpproperties) HasIpv6EnableRetryUpdates() bool`

HasIpv6EnableRetryUpdates returns a boolean if a field has been set.

### GetIpv6GenerateHostname

`func (o *GridDhcpproperties) GetIpv6GenerateHostname() bool`

GetIpv6GenerateHostname returns the Ipv6GenerateHostname field if non-nil, zero value otherwise.

### GetIpv6GenerateHostnameOk

`func (o *GridDhcpproperties) GetIpv6GenerateHostnameOk() (*bool, bool)`

GetIpv6GenerateHostnameOk returns a tuple with the Ipv6GenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6GenerateHostname

`func (o *GridDhcpproperties) SetIpv6GenerateHostname(v bool)`

SetIpv6GenerateHostname sets Ipv6GenerateHostname field to given value.

### HasIpv6GenerateHostname

`func (o *GridDhcpproperties) HasIpv6GenerateHostname() bool`

HasIpv6GenerateHostname returns a boolean if a field has been set.

### GetIpv6GssTsigKeys

`func (o *GridDhcpproperties) GetIpv6GssTsigKeys() []string`

GetIpv6GssTsigKeys returns the Ipv6GssTsigKeys field if non-nil, zero value otherwise.

### GetIpv6GssTsigKeysOk

`func (o *GridDhcpproperties) GetIpv6GssTsigKeysOk() (*[]string, bool)`

GetIpv6GssTsigKeysOk returns a tuple with the Ipv6GssTsigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6GssTsigKeys

`func (o *GridDhcpproperties) SetIpv6GssTsigKeys(v []string)`

SetIpv6GssTsigKeys sets Ipv6GssTsigKeys field to given value.

### HasIpv6GssTsigKeys

`func (o *GridDhcpproperties) HasIpv6GssTsigKeys() bool`

HasIpv6GssTsigKeys returns a boolean if a field has been set.

### GetIpv6KdcServer

`func (o *GridDhcpproperties) GetIpv6KdcServer() string`

GetIpv6KdcServer returns the Ipv6KdcServer field if non-nil, zero value otherwise.

### GetIpv6KdcServerOk

`func (o *GridDhcpproperties) GetIpv6KdcServerOk() (*string, bool)`

GetIpv6KdcServerOk returns a tuple with the Ipv6KdcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6KdcServer

`func (o *GridDhcpproperties) SetIpv6KdcServer(v string)`

SetIpv6KdcServer sets Ipv6KdcServer field to given value.

### HasIpv6KdcServer

`func (o *GridDhcpproperties) HasIpv6KdcServer() bool`

HasIpv6KdcServer returns a boolean if a field has been set.

### GetIpv6LeaseScavengingTime

`func (o *GridDhcpproperties) GetIpv6LeaseScavengingTime() int64`

GetIpv6LeaseScavengingTime returns the Ipv6LeaseScavengingTime field if non-nil, zero value otherwise.

### GetIpv6LeaseScavengingTimeOk

`func (o *GridDhcpproperties) GetIpv6LeaseScavengingTimeOk() (*int64, bool)`

GetIpv6LeaseScavengingTimeOk returns a tuple with the Ipv6LeaseScavengingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6LeaseScavengingTime

`func (o *GridDhcpproperties) SetIpv6LeaseScavengingTime(v int64)`

SetIpv6LeaseScavengingTime sets Ipv6LeaseScavengingTime field to given value.

### HasIpv6LeaseScavengingTime

`func (o *GridDhcpproperties) HasIpv6LeaseScavengingTime() bool`

HasIpv6LeaseScavengingTime returns a boolean if a field has been set.

### GetIpv6MicrosoftCodePage

`func (o *GridDhcpproperties) GetIpv6MicrosoftCodePage() string`

GetIpv6MicrosoftCodePage returns the Ipv6MicrosoftCodePage field if non-nil, zero value otherwise.

### GetIpv6MicrosoftCodePageOk

`func (o *GridDhcpproperties) GetIpv6MicrosoftCodePageOk() (*string, bool)`

GetIpv6MicrosoftCodePageOk returns a tuple with the Ipv6MicrosoftCodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MicrosoftCodePage

`func (o *GridDhcpproperties) SetIpv6MicrosoftCodePage(v string)`

SetIpv6MicrosoftCodePage sets Ipv6MicrosoftCodePage field to given value.

### HasIpv6MicrosoftCodePage

`func (o *GridDhcpproperties) HasIpv6MicrosoftCodePage() bool`

HasIpv6MicrosoftCodePage returns a boolean if a field has been set.

### GetIpv6Options

`func (o *GridDhcpproperties) GetIpv6Options() []GridDhcppropertiesIpv6Options`

GetIpv6Options returns the Ipv6Options field if non-nil, zero value otherwise.

### GetIpv6OptionsOk

`func (o *GridDhcpproperties) GetIpv6OptionsOk() (*[]GridDhcppropertiesIpv6Options, bool)`

GetIpv6OptionsOk returns a tuple with the Ipv6Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Options

`func (o *GridDhcpproperties) SetIpv6Options(v []GridDhcppropertiesIpv6Options)`

SetIpv6Options sets Ipv6Options field to given value.

### HasIpv6Options

`func (o *GridDhcpproperties) HasIpv6Options() bool`

HasIpv6Options returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *GridDhcpproperties) GetIpv6Prefixes() []string`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *GridDhcpproperties) GetIpv6PrefixesOk() (*[]string, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *GridDhcpproperties) SetIpv6Prefixes(v []string)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *GridDhcpproperties) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetIpv6RecycleLeases

`func (o *GridDhcpproperties) GetIpv6RecycleLeases() bool`

GetIpv6RecycleLeases returns the Ipv6RecycleLeases field if non-nil, zero value otherwise.

### GetIpv6RecycleLeasesOk

`func (o *GridDhcpproperties) GetIpv6RecycleLeasesOk() (*bool, bool)`

GetIpv6RecycleLeasesOk returns a tuple with the Ipv6RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RecycleLeases

`func (o *GridDhcpproperties) SetIpv6RecycleLeases(v bool)`

SetIpv6RecycleLeases sets Ipv6RecycleLeases field to given value.

### HasIpv6RecycleLeases

`func (o *GridDhcpproperties) HasIpv6RecycleLeases() bool`

HasIpv6RecycleLeases returns a boolean if a field has been set.

### GetIpv6RememberExpiredClientAssociation

`func (o *GridDhcpproperties) GetIpv6RememberExpiredClientAssociation() bool`

GetIpv6RememberExpiredClientAssociation returns the Ipv6RememberExpiredClientAssociation field if non-nil, zero value otherwise.

### GetIpv6RememberExpiredClientAssociationOk

`func (o *GridDhcpproperties) GetIpv6RememberExpiredClientAssociationOk() (*bool, bool)`

GetIpv6RememberExpiredClientAssociationOk returns a tuple with the Ipv6RememberExpiredClientAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RememberExpiredClientAssociation

`func (o *GridDhcpproperties) SetIpv6RememberExpiredClientAssociation(v bool)`

SetIpv6RememberExpiredClientAssociation sets Ipv6RememberExpiredClientAssociation field to given value.

### HasIpv6RememberExpiredClientAssociation

`func (o *GridDhcpproperties) HasIpv6RememberExpiredClientAssociation() bool`

HasIpv6RememberExpiredClientAssociation returns a boolean if a field has been set.

### GetIpv6RetryUpdatesInterval

`func (o *GridDhcpproperties) GetIpv6RetryUpdatesInterval() int64`

GetIpv6RetryUpdatesInterval returns the Ipv6RetryUpdatesInterval field if non-nil, zero value otherwise.

### GetIpv6RetryUpdatesIntervalOk

`func (o *GridDhcpproperties) GetIpv6RetryUpdatesIntervalOk() (*int64, bool)`

GetIpv6RetryUpdatesIntervalOk returns a tuple with the Ipv6RetryUpdatesInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RetryUpdatesInterval

`func (o *GridDhcpproperties) SetIpv6RetryUpdatesInterval(v int64)`

SetIpv6RetryUpdatesInterval sets Ipv6RetryUpdatesInterval field to given value.

### HasIpv6RetryUpdatesInterval

`func (o *GridDhcpproperties) HasIpv6RetryUpdatesInterval() bool`

HasIpv6RetryUpdatesInterval returns a boolean if a field has been set.

### GetIpv6TxtRecordHandling

`func (o *GridDhcpproperties) GetIpv6TxtRecordHandling() string`

GetIpv6TxtRecordHandling returns the Ipv6TxtRecordHandling field if non-nil, zero value otherwise.

### GetIpv6TxtRecordHandlingOk

`func (o *GridDhcpproperties) GetIpv6TxtRecordHandlingOk() (*string, bool)`

GetIpv6TxtRecordHandlingOk returns a tuple with the Ipv6TxtRecordHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6TxtRecordHandling

`func (o *GridDhcpproperties) SetIpv6TxtRecordHandling(v string)`

SetIpv6TxtRecordHandling sets Ipv6TxtRecordHandling field to given value.

### HasIpv6TxtRecordHandling

`func (o *GridDhcpproperties) HasIpv6TxtRecordHandling() bool`

HasIpv6TxtRecordHandling returns a boolean if a field has been set.

### GetIpv6UpdateDnsOnLeaseRenewal

`func (o *GridDhcpproperties) GetIpv6UpdateDnsOnLeaseRenewal() bool`

GetIpv6UpdateDnsOnLeaseRenewal returns the Ipv6UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetIpv6UpdateDnsOnLeaseRenewalOk

`func (o *GridDhcpproperties) GetIpv6UpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetIpv6UpdateDnsOnLeaseRenewalOk returns a tuple with the Ipv6UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6UpdateDnsOnLeaseRenewal

`func (o *GridDhcpproperties) SetIpv6UpdateDnsOnLeaseRenewal(v bool)`

SetIpv6UpdateDnsOnLeaseRenewal sets Ipv6UpdateDnsOnLeaseRenewal field to given value.

### HasIpv6UpdateDnsOnLeaseRenewal

`func (o *GridDhcpproperties) HasIpv6UpdateDnsOnLeaseRenewal() bool`

HasIpv6UpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetKdcServer

`func (o *GridDhcpproperties) GetKdcServer() string`

GetKdcServer returns the KdcServer field if non-nil, zero value otherwise.

### GetKdcServerOk

`func (o *GridDhcpproperties) GetKdcServerOk() (*string, bool)`

GetKdcServerOk returns a tuple with the KdcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcServer

`func (o *GridDhcpproperties) SetKdcServer(v string)`

SetKdcServer sets KdcServer field to given value.

### HasKdcServer

`func (o *GridDhcpproperties) HasKdcServer() bool`

HasKdcServer returns a boolean if a field has been set.

### GetLeaseLoggingMember

`func (o *GridDhcpproperties) GetLeaseLoggingMember() string`

GetLeaseLoggingMember returns the LeaseLoggingMember field if non-nil, zero value otherwise.

### GetLeaseLoggingMemberOk

`func (o *GridDhcpproperties) GetLeaseLoggingMemberOk() (*string, bool)`

GetLeaseLoggingMemberOk returns a tuple with the LeaseLoggingMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseLoggingMember

`func (o *GridDhcpproperties) SetLeaseLoggingMember(v string)`

SetLeaseLoggingMember sets LeaseLoggingMember field to given value.

### HasLeaseLoggingMember

`func (o *GridDhcpproperties) HasLeaseLoggingMember() bool`

HasLeaseLoggingMember returns a boolean if a field has been set.

### GetLeasePerClientSettings

`func (o *GridDhcpproperties) GetLeasePerClientSettings() string`

GetLeasePerClientSettings returns the LeasePerClientSettings field if non-nil, zero value otherwise.

### GetLeasePerClientSettingsOk

`func (o *GridDhcpproperties) GetLeasePerClientSettingsOk() (*string, bool)`

GetLeasePerClientSettingsOk returns a tuple with the LeasePerClientSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasePerClientSettings

`func (o *GridDhcpproperties) SetLeasePerClientSettings(v string)`

SetLeasePerClientSettings sets LeasePerClientSettings field to given value.

### HasLeasePerClientSettings

`func (o *GridDhcpproperties) HasLeasePerClientSettings() bool`

HasLeasePerClientSettings returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *GridDhcpproperties) GetLeaseScavengeTime() int32`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *GridDhcpproperties) GetLeaseScavengeTimeOk() (*int32, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *GridDhcpproperties) SetLeaseScavengeTime(v int32)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *GridDhcpproperties) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogLeaseEvents

`func (o *GridDhcpproperties) GetLogLeaseEvents() bool`

GetLogLeaseEvents returns the LogLeaseEvents field if non-nil, zero value otherwise.

### GetLogLeaseEventsOk

`func (o *GridDhcpproperties) GetLogLeaseEventsOk() (*bool, bool)`

GetLogLeaseEventsOk returns a tuple with the LogLeaseEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLeaseEvents

`func (o *GridDhcpproperties) SetLogLeaseEvents(v bool)`

SetLogLeaseEvents sets LogLeaseEvents field to given value.

### HasLogLeaseEvents

`func (o *GridDhcpproperties) HasLogLeaseEvents() bool`

HasLogLeaseEvents returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GridDhcpproperties) GetLogicFilterRules() []GridDhcppropertiesLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GridDhcpproperties) GetLogicFilterRulesOk() (*[]GridDhcppropertiesLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GridDhcpproperties) SetLogicFilterRules(v []GridDhcppropertiesLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GridDhcpproperties) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GridDhcpproperties) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GridDhcpproperties) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GridDhcpproperties) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GridDhcpproperties) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *GridDhcpproperties) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *GridDhcpproperties) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *GridDhcpproperties) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *GridDhcpproperties) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMicrosoftCodePage

`func (o *GridDhcpproperties) GetMicrosoftCodePage() string`

GetMicrosoftCodePage returns the MicrosoftCodePage field if non-nil, zero value otherwise.

### GetMicrosoftCodePageOk

`func (o *GridDhcpproperties) GetMicrosoftCodePageOk() (*string, bool)`

GetMicrosoftCodePageOk returns a tuple with the MicrosoftCodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftCodePage

`func (o *GridDhcpproperties) SetMicrosoftCodePage(v string)`

SetMicrosoftCodePage sets MicrosoftCodePage field to given value.

### HasMicrosoftCodePage

`func (o *GridDhcpproperties) HasMicrosoftCodePage() bool`

HasMicrosoftCodePage returns a boolean if a field has been set.

### GetNextserver

`func (o *GridDhcpproperties) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GridDhcpproperties) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GridDhcpproperties) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GridDhcpproperties) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOption60MatchRules

`func (o *GridDhcpproperties) GetOption60MatchRules() []GridDhcppropertiesOption60MatchRules`

GetOption60MatchRules returns the Option60MatchRules field if non-nil, zero value otherwise.

### GetOption60MatchRulesOk

`func (o *GridDhcpproperties) GetOption60MatchRulesOk() (*[]GridDhcppropertiesOption60MatchRules, bool)`

GetOption60MatchRulesOk returns a tuple with the Option60MatchRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption60MatchRules

`func (o *GridDhcpproperties) SetOption60MatchRules(v []GridDhcppropertiesOption60MatchRules)`

SetOption60MatchRules sets Option60MatchRules field to given value.

### HasOption60MatchRules

`func (o *GridDhcpproperties) HasOption60MatchRules() bool`

HasOption60MatchRules returns a boolean if a field has been set.

### GetOptions

`func (o *GridDhcpproperties) GetOptions() []GridDhcppropertiesOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GridDhcpproperties) GetOptionsOk() (*[]GridDhcppropertiesOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GridDhcpproperties) SetOptions(v []GridDhcppropertiesOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GridDhcpproperties) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPingCount

`func (o *GridDhcpproperties) GetPingCount() int64`

GetPingCount returns the PingCount field if non-nil, zero value otherwise.

### GetPingCountOk

`func (o *GridDhcpproperties) GetPingCountOk() (*int64, bool)`

GetPingCountOk returns a tuple with the PingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingCount

`func (o *GridDhcpproperties) SetPingCount(v int64)`

SetPingCount sets PingCount field to given value.

### HasPingCount

`func (o *GridDhcpproperties) HasPingCount() bool`

HasPingCount returns a boolean if a field has been set.

### GetPingTimeout

`func (o *GridDhcpproperties) GetPingTimeout() int64`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *GridDhcpproperties) GetPingTimeoutOk() (*int64, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *GridDhcpproperties) SetPingTimeout(v int64)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *GridDhcpproperties) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *GridDhcpproperties) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *GridDhcpproperties) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *GridDhcpproperties) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *GridDhcpproperties) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetPrefixLengthMode

`func (o *GridDhcpproperties) GetPrefixLengthMode() string`

GetPrefixLengthMode returns the PrefixLengthMode field if non-nil, zero value otherwise.

### GetPrefixLengthModeOk

`func (o *GridDhcpproperties) GetPrefixLengthModeOk() (*string, bool)`

GetPrefixLengthModeOk returns a tuple with the PrefixLengthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLengthMode

`func (o *GridDhcpproperties) SetPrefixLengthMode(v string)`

SetPrefixLengthMode sets PrefixLengthMode field to given value.

### HasPrefixLengthMode

`func (o *GridDhcpproperties) HasPrefixLengthMode() bool`

HasPrefixLengthMode returns a boolean if a field has been set.

### GetProtocolHostnameRewritePolicies

`func (o *GridDhcpproperties) GetProtocolHostnameRewritePolicies() []string`

GetProtocolHostnameRewritePolicies returns the ProtocolHostnameRewritePolicies field if non-nil, zero value otherwise.

### GetProtocolHostnameRewritePoliciesOk

`func (o *GridDhcpproperties) GetProtocolHostnameRewritePoliciesOk() (*[]string, bool)`

GetProtocolHostnameRewritePoliciesOk returns a tuple with the ProtocolHostnameRewritePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolHostnameRewritePolicies

`func (o *GridDhcpproperties) SetProtocolHostnameRewritePolicies(v []string)`

SetProtocolHostnameRewritePolicies sets ProtocolHostnameRewritePolicies field to given value.

### HasProtocolHostnameRewritePolicies

`func (o *GridDhcpproperties) HasProtocolHostnameRewritePolicies() bool`

HasProtocolHostnameRewritePolicies returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GridDhcpproperties) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GridDhcpproperties) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GridDhcpproperties) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GridDhcpproperties) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GridDhcpproperties) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GridDhcpproperties) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GridDhcpproperties) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GridDhcpproperties) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRestartSetting

`func (o *GridDhcpproperties) GetRestartSetting() GridDhcppropertiesRestartSetting`

GetRestartSetting returns the RestartSetting field if non-nil, zero value otherwise.

### GetRestartSettingOk

`func (o *GridDhcpproperties) GetRestartSettingOk() (*GridDhcppropertiesRestartSetting, bool)`

GetRestartSettingOk returns a tuple with the RestartSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSetting

`func (o *GridDhcpproperties) SetRestartSetting(v GridDhcppropertiesRestartSetting)`

SetRestartSetting sets RestartSetting field to given value.

### HasRestartSetting

`func (o *GridDhcpproperties) HasRestartSetting() bool`

HasRestartSetting returns a boolean if a field has been set.

### GetRetryDdnsUpdates

`func (o *GridDhcpproperties) GetRetryDdnsUpdates() bool`

GetRetryDdnsUpdates returns the RetryDdnsUpdates field if non-nil, zero value otherwise.

### GetRetryDdnsUpdatesOk

`func (o *GridDhcpproperties) GetRetryDdnsUpdatesOk() (*bool, bool)`

GetRetryDdnsUpdatesOk returns a tuple with the RetryDdnsUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDdnsUpdates

`func (o *GridDhcpproperties) SetRetryDdnsUpdates(v bool)`

SetRetryDdnsUpdates sets RetryDdnsUpdates field to given value.

### HasRetryDdnsUpdates

`func (o *GridDhcpproperties) HasRetryDdnsUpdates() bool`

HasRetryDdnsUpdates returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *GridDhcpproperties) GetSyslogFacility() string`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *GridDhcpproperties) GetSyslogFacilityOk() (*string, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *GridDhcpproperties) SetSyslogFacility(v string)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *GridDhcpproperties) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetTxtRecordHandling

`func (o *GridDhcpproperties) GetTxtRecordHandling() string`

GetTxtRecordHandling returns the TxtRecordHandling field if non-nil, zero value otherwise.

### GetTxtRecordHandlingOk

`func (o *GridDhcpproperties) GetTxtRecordHandlingOk() (*string, bool)`

GetTxtRecordHandlingOk returns a tuple with the TxtRecordHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtRecordHandling

`func (o *GridDhcpproperties) SetTxtRecordHandling(v string)`

SetTxtRecordHandling sets TxtRecordHandling field to given value.

### HasTxtRecordHandling

`func (o *GridDhcpproperties) HasTxtRecordHandling() bool`

HasTxtRecordHandling returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GridDhcpproperties) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GridDhcpproperties) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GridDhcpproperties) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GridDhcpproperties) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetValidLifetime

`func (o *GridDhcpproperties) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *GridDhcpproperties) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *GridDhcpproperties) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *GridDhcpproperties) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


