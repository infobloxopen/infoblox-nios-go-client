# GetRangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AlwaysUpdateDns** | Pointer to **bool** | This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. | [optional] 
**Bootfile** | Pointer to **string** | The bootfile name for the range. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The bootserver address for the range. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**CloudInfo** | Pointer to [**RangeCloudInfo**](RangeCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the range; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this range. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DenyAllClients** | Pointer to **bool** | If True, send NAK forcing the client to take the new address. | [optional] 
**DenyBootp** | Pointer to **bool** | If set to true, BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of the range multiplied by 1000. This is the percentage of the total number of available IP addresses belonging to the range versus the total number of all IP addresses in the range. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of the range. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether a range is disabled or not. When this is set to False, the range is enabled. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this range. | [optional] [readonly] 
**DiscoveryBasicPollSettings** | Pointer to [**RangeDiscoveryBasicPollSettings**](RangeDiscoveryBasicPollSettings.md) |  | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**RangeDiscoveryBlackoutSetting**](RangeDiscoveryBlackoutSetting.md) |  | [optional] 
**DiscoveryMember** | Pointer to **string** | The member that will run discovery for this range. | [optional] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the range. | [optional] [readonly] 
**EmailList** | Pointer to **[]string** | The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP range object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableDhcpThresholds** | Pointer to **bool** | Determines if DHCP thresholds are enabled for the range. | [optional] 
**EnableDiscovery** | Pointer to **bool** | Determines whether a discovery is enabled or not for this range. When this is set to False, the discovery for this range is disabled. | [optional] 
**EnableEmailWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are sent through email. | [optional] 
**EnableIfmapPublishing** | Pointer to **bool** | Determines if IFMAP publishing is enabled for the range. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the range should be immediately enabled. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are send through SNMP. | [optional] 
**EndAddr** | Pointer to **string** | The IPv4 Address end address of the range. | [optional] 
**EndpointSources** | Pointer to **[]string** | The endpoints that provides data for the DHCP Range object. | [optional] [readonly] 
**Exclude** | Pointer to [**[]RangeExclude**](RangeExclude.md) | These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally, information about this exclusion range. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FailoverAssociation** | Pointer to **string** | The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {range:range} must be set to &#39;FAILOVER&#39; or &#39;FAILOVER_MS&#39; if you want the failover association specified here to serve the range. | [optional] 
**FingerprintFilterRules** | Pointer to [**[]RangeFingerprintFilterRules**](RangeFingerprintFilterRules.md) | This field contains the fingerprint filters for this DHCP range. The appliance uses matching rules in these filters to select the address range from which it assigns a lease. | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage of DHCP range usage threshold above which range usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage of DHCP range usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \&quot;NONE\&quot;, \&quot;CLIENT\&quot;, or \&quot;MACADDR\&quot;. The default is \&quot;NONE\&quot;. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**IsSplitScope** | Pointer to **bool** | This field will be &#39;true&#39; if this particular range is part of a split scope. | [optional] [readonly] 
**KnownClients** | Pointer to **string** | Permission for known clients. This can be &#39;Allow&#39; or &#39;Deny&#39;. If set to &#39;Deny&#39; known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries. | [optional] 
**LeaseScavengeTime** | Pointer to **int64** | An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogicFilterRules** | Pointer to [**[]RangeLogicFilterRules**](RangeLogicFilterRules.md) | This field contains the logic filters to be applied to this range. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**LowWaterMark** | Pointer to **int64** | The percentage of DHCP range usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**LowWaterMarkReset** | Pointer to **int64** | The percentage of DHCP range usage threshold below which range usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value. | [optional] 
**MacFilterRules** | Pointer to [**[]RangeMacFilterRules**](RangeMacFilterRules.md) | This field contains the MAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**Member** | Pointer to [**RangeMember**](RangeMember.md) |  | [optional] 
**MsAdUserData** | Pointer to [**RangeMsAdUserData**](RangeMsAdUserData.md) |  | [optional] 
**MsOptions** | Pointer to [**[]RangeMsOptions**](RangeMsOptions.md) | This field contains the Microsoft DHCP options for this range. | [optional] 
**MsServer** | Pointer to [**RangeMsServer**](RangeMsServer.md) |  | [optional] 
**NacFilterRules** | Pointer to [**[]RangeNacFilterRules**](RangeNacFilterRules.md) | This field contains the NAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**Name** | Pointer to **string** | This field contains the name of the Microsoft scope. | [optional] 
**Network** | Pointer to **string** | The network to which this range belongs, in IPv4 Address/CIDR format. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this range resides. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address of the next server that the host needs to boot. | [optional] 
**OptionFilterRules** | Pointer to [**[]RangeOptionFilterRules**](RangeOptionFilterRules.md) | This field contains the Option filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**Options** | Pointer to [**[]RangeOptions**](RangeOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PortControlBlackoutSetting** | Pointer to [**RangePortControlBlackoutSetting**](RangePortControlBlackoutSetting.md) |  | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value of a DHCP Range object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**RelayAgentFilterRules** | Pointer to [**[]RangeRelayAgentFilterRules**](RangeRelayAgentFilterRules.md) | This field contains the Relay Agent filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | If the field is set to True, the discovery blackout setting will be used for port control blackout setting. | [optional] 
**ServerAssociationType** | Pointer to **string** | The type of server that is going to serve the range. | [optional] 
**SplitMember** | Pointer to [**RangeSplitMember**](RangeSplitMember.md) |  | [optional] 
**SplitScopeExclusionPercent** | Pointer to **int64** | This field controls the percentage used when creating a split scope. Valid values are numbers between 1 and 99. If the value is 40, it means that the top 40% of the exclusion will be created on the DHCP range assigned to {next_available_ip:next_available_ip} and the lower 60% of the range will be assigned to DHCP range assigned to {next_available_ip:next_available_ip} | [optional] 
**StartAddr** | Pointer to **string** | The IPv4 Address starting address of the range. | [optional] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in the range. | [optional] [readonly] 
**SubscribeSettings** | Pointer to [**RangeSubscribeSettings**](RangeSubscribeSettings.md) |  | [optional] 
**Template** | Pointer to **string** | If set on creation, the range will be created according to the values specified in the named template. | [optional] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in the range. | [optional] [readonly] 
**UnknownClients** | Pointer to **string** | Permission for unknown clients. This can be &#39;Allow&#39; or &#39;Deny&#39;. If set to &#39;Deny&#39;, unknown clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries. | [optional] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseBlackoutSetting** | Pointer to **bool** | Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseDiscoveryBasicPollingSettings** | Pointer to **bool** | Use flag for: discovery_basic_poll_settings | [optional] 
**UseEmailList** | Pointer to **bool** | Use flag for: email_list | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseEnableDhcpThresholds** | Pointer to **bool** | Use flag for: enable_dhcp_thresholds | [optional] 
**UseEnableDiscovery** | Pointer to **bool** | Use flag for: discovery_member , enable_discovery | [optional] 
**UseEnableIfmapPublishing** | Pointer to **bool** | Use flag for: enable_ifmap_publishing | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseIgnoreId** | Pointer to **bool** | Use flag for: ignore_id | [optional] 
**UseKnownClients** | Pointer to **bool** | Use flag for: known_clients | [optional] 
**UseLeaseScavengeTime** | Pointer to **bool** | Use flag for: lease_scavenge_time | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMsOptions** | Pointer to **bool** | Use flag for: ms_options | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseSubscribeSettings** | Pointer to **bool** | Use flag for: subscribe_settings | [optional] 
**UseUnknownClients** | Pointer to **bool** | Use flag for: unknown_clients | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Range**](Range.md) |  | [optional] 

## Methods

### NewGetRangeResponse

`func NewGetRangeResponse() *GetRangeResponse`

NewGetRangeResponse instantiates a new GetRangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRangeResponseWithDefaults

`func NewGetRangeResponseWithDefaults() *GetRangeResponse`

NewGetRangeResponseWithDefaults instantiates a new GetRangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRangeResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRangeResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRangeResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRangeResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlwaysUpdateDns

`func (o *GetRangeResponse) GetAlwaysUpdateDns() bool`

GetAlwaysUpdateDns returns the AlwaysUpdateDns field if non-nil, zero value otherwise.

### GetAlwaysUpdateDnsOk

`func (o *GetRangeResponse) GetAlwaysUpdateDnsOk() (*bool, bool)`

GetAlwaysUpdateDnsOk returns a tuple with the AlwaysUpdateDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUpdateDns

`func (o *GetRangeResponse) SetAlwaysUpdateDns(v bool)`

SetAlwaysUpdateDns sets AlwaysUpdateDns field to given value.

### HasAlwaysUpdateDns

`func (o *GetRangeResponse) HasAlwaysUpdateDns() bool`

HasAlwaysUpdateDns returns a boolean if a field has been set.

### GetBootfile

`func (o *GetRangeResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetRangeResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetRangeResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetRangeResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetRangeResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetRangeResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetRangeResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetRangeResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRangeResponse) GetCloudInfo() RangeCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRangeResponse) GetCloudInfoOk() (*RangeCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRangeResponse) SetCloudInfo(v RangeCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRangeResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetRangeResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRangeResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRangeResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRangeResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetRangeResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetRangeResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetRangeResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetRangeResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GetRangeResponse) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GetRangeResponse) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GetRangeResponse) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GetRangeResponse) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDenyAllClients

`func (o *GetRangeResponse) GetDenyAllClients() bool`

GetDenyAllClients returns the DenyAllClients field if non-nil, zero value otherwise.

### GetDenyAllClientsOk

`func (o *GetRangeResponse) GetDenyAllClientsOk() (*bool, bool)`

GetDenyAllClientsOk returns a tuple with the DenyAllClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyAllClients

`func (o *GetRangeResponse) SetDenyAllClients(v bool)`

SetDenyAllClients sets DenyAllClients field to given value.

### HasDenyAllClients

`func (o *GetRangeResponse) HasDenyAllClients() bool`

HasDenyAllClients returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetRangeResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetRangeResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetRangeResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetRangeResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *GetRangeResponse) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *GetRangeResponse) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *GetRangeResponse) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *GetRangeResponse) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *GetRangeResponse) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *GetRangeResponse) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *GetRangeResponse) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *GetRangeResponse) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *GetRangeResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRangeResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRangeResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRangeResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetRangeResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetRangeResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetRangeResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetRangeResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *GetRangeResponse) GetDiscoveryBasicPollSettings() RangeDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *GetRangeResponse) GetDiscoveryBasicPollSettingsOk() (*RangeDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *GetRangeResponse) SetDiscoveryBasicPollSettings(v RangeDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *GetRangeResponse) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *GetRangeResponse) GetDiscoveryBlackoutSetting() RangeDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *GetRangeResponse) GetDiscoveryBlackoutSettingOk() (*RangeDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *GetRangeResponse) SetDiscoveryBlackoutSetting(v RangeDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *GetRangeResponse) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *GetRangeResponse) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *GetRangeResponse) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *GetRangeResponse) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *GetRangeResponse) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *GetRangeResponse) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *GetRangeResponse) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *GetRangeResponse) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *GetRangeResponse) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetEmailList

`func (o *GetRangeResponse) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *GetRangeResponse) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *GetRangeResponse) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *GetRangeResponse) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetRangeResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetRangeResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetRangeResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetRangeResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *GetRangeResponse) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *GetRangeResponse) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *GetRangeResponse) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *GetRangeResponse) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *GetRangeResponse) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *GetRangeResponse) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *GetRangeResponse) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *GetRangeResponse) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *GetRangeResponse) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *GetRangeResponse) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *GetRangeResponse) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *GetRangeResponse) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableIfmapPublishing

`func (o *GetRangeResponse) GetEnableIfmapPublishing() bool`

GetEnableIfmapPublishing returns the EnableIfmapPublishing field if non-nil, zero value otherwise.

### GetEnableIfmapPublishingOk

`func (o *GetRangeResponse) GetEnableIfmapPublishingOk() (*bool, bool)`

GetEnableIfmapPublishingOk returns a tuple with the EnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIfmapPublishing

`func (o *GetRangeResponse) SetEnableIfmapPublishing(v bool)`

SetEnableIfmapPublishing sets EnableIfmapPublishing field to given value.

### HasEnableIfmapPublishing

`func (o *GetRangeResponse) HasEnableIfmapPublishing() bool`

HasEnableIfmapPublishing returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *GetRangeResponse) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *GetRangeResponse) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *GetRangeResponse) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *GetRangeResponse) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetRangeResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetRangeResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetRangeResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetRangeResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *GetRangeResponse) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *GetRangeResponse) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *GetRangeResponse) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *GetRangeResponse) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetEndAddr

`func (o *GetRangeResponse) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *GetRangeResponse) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *GetRangeResponse) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *GetRangeResponse) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetEndpointSources

`func (o *GetRangeResponse) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *GetRangeResponse) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *GetRangeResponse) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *GetRangeResponse) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExclude

`func (o *GetRangeResponse) GetExclude() []RangeExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *GetRangeResponse) GetExcludeOk() (*[]RangeExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *GetRangeResponse) SetExclude(v []RangeExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *GetRangeResponse) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetRangeResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetRangeResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetRangeResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetRangeResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetRangeResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetRangeResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetRangeResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetRangeResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRangeResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRangeResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRangeResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRangeResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFailoverAssociation

`func (o *GetRangeResponse) GetFailoverAssociation() string`

GetFailoverAssociation returns the FailoverAssociation field if non-nil, zero value otherwise.

### GetFailoverAssociationOk

`func (o *GetRangeResponse) GetFailoverAssociationOk() (*string, bool)`

GetFailoverAssociationOk returns a tuple with the FailoverAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAssociation

`func (o *GetRangeResponse) SetFailoverAssociation(v string)`

SetFailoverAssociation sets FailoverAssociation field to given value.

### HasFailoverAssociation

`func (o *GetRangeResponse) HasFailoverAssociation() bool`

HasFailoverAssociation returns a boolean if a field has been set.

### GetFingerprintFilterRules

`func (o *GetRangeResponse) GetFingerprintFilterRules() []RangeFingerprintFilterRules`

GetFingerprintFilterRules returns the FingerprintFilterRules field if non-nil, zero value otherwise.

### GetFingerprintFilterRulesOk

`func (o *GetRangeResponse) GetFingerprintFilterRulesOk() (*[]RangeFingerprintFilterRules, bool)`

GetFingerprintFilterRulesOk returns a tuple with the FingerprintFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintFilterRules

`func (o *GetRangeResponse) SetFingerprintFilterRules(v []RangeFingerprintFilterRules)`

SetFingerprintFilterRules sets FingerprintFilterRules field to given value.

### HasFingerprintFilterRules

`func (o *GetRangeResponse) HasFingerprintFilterRules() bool`

HasFingerprintFilterRules returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetRangeResponse) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetRangeResponse) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetRangeResponse) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetRangeResponse) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *GetRangeResponse) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *GetRangeResponse) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *GetRangeResponse) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *GetRangeResponse) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GetRangeResponse) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GetRangeResponse) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GetRangeResponse) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GetRangeResponse) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *GetRangeResponse) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *GetRangeResponse) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *GetRangeResponse) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *GetRangeResponse) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *GetRangeResponse) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *GetRangeResponse) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *GetRangeResponse) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *GetRangeResponse) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetIsSplitScope

`func (o *GetRangeResponse) GetIsSplitScope() bool`

GetIsSplitScope returns the IsSplitScope field if non-nil, zero value otherwise.

### GetIsSplitScopeOk

`func (o *GetRangeResponse) GetIsSplitScopeOk() (*bool, bool)`

GetIsSplitScopeOk returns a tuple with the IsSplitScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSplitScope

`func (o *GetRangeResponse) SetIsSplitScope(v bool)`

SetIsSplitScope sets IsSplitScope field to given value.

### HasIsSplitScope

`func (o *GetRangeResponse) HasIsSplitScope() bool`

HasIsSplitScope returns a boolean if a field has been set.

### GetKnownClients

`func (o *GetRangeResponse) GetKnownClients() string`

GetKnownClients returns the KnownClients field if non-nil, zero value otherwise.

### GetKnownClientsOk

`func (o *GetRangeResponse) GetKnownClientsOk() (*string, bool)`

GetKnownClientsOk returns a tuple with the KnownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClients

`func (o *GetRangeResponse) SetKnownClients(v string)`

SetKnownClients sets KnownClients field to given value.

### HasKnownClients

`func (o *GetRangeResponse) HasKnownClients() bool`

HasKnownClients returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *GetRangeResponse) GetLeaseScavengeTime() int64`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *GetRangeResponse) GetLeaseScavengeTimeOk() (*int64, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *GetRangeResponse) SetLeaseScavengeTime(v int64)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *GetRangeResponse) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetRangeResponse) GetLogicFilterRules() []RangeLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetRangeResponse) GetLogicFilterRulesOk() (*[]RangeLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetRangeResponse) SetLogicFilterRules(v []RangeLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetRangeResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetRangeResponse) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetRangeResponse) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetRangeResponse) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetRangeResponse) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *GetRangeResponse) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *GetRangeResponse) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *GetRangeResponse) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *GetRangeResponse) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMacFilterRules

`func (o *GetRangeResponse) GetMacFilterRules() []RangeMacFilterRules`

GetMacFilterRules returns the MacFilterRules field if non-nil, zero value otherwise.

### GetMacFilterRulesOk

`func (o *GetRangeResponse) GetMacFilterRulesOk() (*[]RangeMacFilterRules, bool)`

GetMacFilterRulesOk returns a tuple with the MacFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilterRules

`func (o *GetRangeResponse) SetMacFilterRules(v []RangeMacFilterRules)`

SetMacFilterRules sets MacFilterRules field to given value.

### HasMacFilterRules

`func (o *GetRangeResponse) HasMacFilterRules() bool`

HasMacFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *GetRangeResponse) GetMember() RangeMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetRangeResponse) GetMemberOk() (*RangeMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetRangeResponse) SetMember(v RangeMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetRangeResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetRangeResponse) GetMsAdUserData() RangeMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetRangeResponse) GetMsAdUserDataOk() (*RangeMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetRangeResponse) SetMsAdUserData(v RangeMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetRangeResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetMsOptions

`func (o *GetRangeResponse) GetMsOptions() []RangeMsOptions`

GetMsOptions returns the MsOptions field if non-nil, zero value otherwise.

### GetMsOptionsOk

`func (o *GetRangeResponse) GetMsOptionsOk() (*[]RangeMsOptions, bool)`

GetMsOptionsOk returns a tuple with the MsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsOptions

`func (o *GetRangeResponse) SetMsOptions(v []RangeMsOptions)`

SetMsOptions sets MsOptions field to given value.

### HasMsOptions

`func (o *GetRangeResponse) HasMsOptions() bool`

HasMsOptions returns a boolean if a field has been set.

### GetMsServer

`func (o *GetRangeResponse) GetMsServer() RangeMsServer`

GetMsServer returns the MsServer field if non-nil, zero value otherwise.

### GetMsServerOk

`func (o *GetRangeResponse) GetMsServerOk() (*RangeMsServer, bool)`

GetMsServerOk returns a tuple with the MsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServer

`func (o *GetRangeResponse) SetMsServer(v RangeMsServer)`

SetMsServer sets MsServer field to given value.

### HasMsServer

`func (o *GetRangeResponse) HasMsServer() bool`

HasMsServer returns a boolean if a field has been set.

### GetNacFilterRules

`func (o *GetRangeResponse) GetNacFilterRules() []RangeNacFilterRules`

GetNacFilterRules returns the NacFilterRules field if non-nil, zero value otherwise.

### GetNacFilterRulesOk

`func (o *GetRangeResponse) GetNacFilterRulesOk() (*[]RangeNacFilterRules, bool)`

GetNacFilterRulesOk returns a tuple with the NacFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacFilterRules

`func (o *GetRangeResponse) SetNacFilterRules(v []RangeNacFilterRules)`

SetNacFilterRules sets NacFilterRules field to given value.

### HasNacFilterRules

`func (o *GetRangeResponse) HasNacFilterRules() bool`

HasNacFilterRules returns a boolean if a field has been set.

### GetName

`func (o *GetRangeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRangeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRangeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRangeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetRangeResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetRangeResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetRangeResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetRangeResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetRangeResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetRangeResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetRangeResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetRangeResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *GetRangeResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetRangeResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetRangeResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetRangeResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *GetRangeResponse) GetOptionFilterRules() []RangeOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *GetRangeResponse) GetOptionFilterRulesOk() (*[]RangeOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *GetRangeResponse) SetOptionFilterRules(v []RangeOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *GetRangeResponse) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetOptions

`func (o *GetRangeResponse) GetOptions() []RangeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetRangeResponse) GetOptionsOk() (*[]RangeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetRangeResponse) SetOptions(v []RangeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetRangeResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *GetRangeResponse) GetPortControlBlackoutSetting() RangePortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *GetRangeResponse) GetPortControlBlackoutSettingOk() (*RangePortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *GetRangeResponse) SetPortControlBlackoutSetting(v RangePortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *GetRangeResponse) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetRangeResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetRangeResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetRangeResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetRangeResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetRangeResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetRangeResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetRangeResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetRangeResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRelayAgentFilterRules

`func (o *GetRangeResponse) GetRelayAgentFilterRules() []RangeRelayAgentFilterRules`

GetRelayAgentFilterRules returns the RelayAgentFilterRules field if non-nil, zero value otherwise.

### GetRelayAgentFilterRulesOk

`func (o *GetRangeResponse) GetRelayAgentFilterRulesOk() (*[]RangeRelayAgentFilterRules, bool)`

GetRelayAgentFilterRulesOk returns a tuple with the RelayAgentFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAgentFilterRules

`func (o *GetRangeResponse) SetRelayAgentFilterRules(v []RangeRelayAgentFilterRules)`

SetRelayAgentFilterRules sets RelayAgentFilterRules field to given value.

### HasRelayAgentFilterRules

`func (o *GetRangeResponse) HasRelayAgentFilterRules() bool`

HasRelayAgentFilterRules returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetRangeResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetRangeResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetRangeResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetRangeResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *GetRangeResponse) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *GetRangeResponse) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *GetRangeResponse) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *GetRangeResponse) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *GetRangeResponse) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *GetRangeResponse) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *GetRangeResponse) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *GetRangeResponse) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetSplitMember

`func (o *GetRangeResponse) GetSplitMember() RangeSplitMember`

GetSplitMember returns the SplitMember field if non-nil, zero value otherwise.

### GetSplitMemberOk

`func (o *GetRangeResponse) GetSplitMemberOk() (*RangeSplitMember, bool)`

GetSplitMemberOk returns a tuple with the SplitMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitMember

`func (o *GetRangeResponse) SetSplitMember(v RangeSplitMember)`

SetSplitMember sets SplitMember field to given value.

### HasSplitMember

`func (o *GetRangeResponse) HasSplitMember() bool`

HasSplitMember returns a boolean if a field has been set.

### GetSplitScopeExclusionPercent

`func (o *GetRangeResponse) GetSplitScopeExclusionPercent() int64`

GetSplitScopeExclusionPercent returns the SplitScopeExclusionPercent field if non-nil, zero value otherwise.

### GetSplitScopeExclusionPercentOk

`func (o *GetRangeResponse) GetSplitScopeExclusionPercentOk() (*int64, bool)`

GetSplitScopeExclusionPercentOk returns a tuple with the SplitScopeExclusionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitScopeExclusionPercent

`func (o *GetRangeResponse) SetSplitScopeExclusionPercent(v int64)`

SetSplitScopeExclusionPercent sets SplitScopeExclusionPercent field to given value.

### HasSplitScopeExclusionPercent

`func (o *GetRangeResponse) HasSplitScopeExclusionPercent() bool`

HasSplitScopeExclusionPercent returns a boolean if a field has been set.

### GetStartAddr

`func (o *GetRangeResponse) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *GetRangeResponse) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *GetRangeResponse) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *GetRangeResponse) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.

### GetStaticHosts

`func (o *GetRangeResponse) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *GetRangeResponse) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *GetRangeResponse) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *GetRangeResponse) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *GetRangeResponse) GetSubscribeSettings() RangeSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *GetRangeResponse) GetSubscribeSettingsOk() (*RangeSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *GetRangeResponse) SetSubscribeSettings(v RangeSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *GetRangeResponse) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *GetRangeResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetRangeResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetRangeResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetRangeResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTotalHosts

`func (o *GetRangeResponse) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *GetRangeResponse) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *GetRangeResponse) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *GetRangeResponse) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUnknownClients

`func (o *GetRangeResponse) GetUnknownClients() string`

GetUnknownClients returns the UnknownClients field if non-nil, zero value otherwise.

### GetUnknownClientsOk

`func (o *GetRangeResponse) GetUnknownClientsOk() (*string, bool)`

GetUnknownClientsOk returns a tuple with the UnknownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownClients

`func (o *GetRangeResponse) SetUnknownClients(v string)`

SetUnknownClients sets UnknownClients field to given value.

### HasUnknownClients

`func (o *GetRangeResponse) HasUnknownClients() bool`

HasUnknownClients returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GetRangeResponse) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GetRangeResponse) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GetRangeResponse) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GetRangeResponse) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *GetRangeResponse) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *GetRangeResponse) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *GetRangeResponse) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *GetRangeResponse) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetRangeResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetRangeResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetRangeResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetRangeResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetRangeResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetRangeResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetRangeResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetRangeResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetRangeResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetRangeResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetRangeResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetRangeResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *GetRangeResponse) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *GetRangeResponse) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *GetRangeResponse) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *GetRangeResponse) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetRangeResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetRangeResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetRangeResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetRangeResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *GetRangeResponse) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *GetRangeResponse) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *GetRangeResponse) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *GetRangeResponse) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEmailList

`func (o *GetRangeResponse) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *GetRangeResponse) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *GetRangeResponse) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *GetRangeResponse) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetRangeResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetRangeResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetRangeResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetRangeResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *GetRangeResponse) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *GetRangeResponse) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *GetRangeResponse) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *GetRangeResponse) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *GetRangeResponse) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *GetRangeResponse) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *GetRangeResponse) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *GetRangeResponse) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseEnableIfmapPublishing

`func (o *GetRangeResponse) GetUseEnableIfmapPublishing() bool`

GetUseEnableIfmapPublishing returns the UseEnableIfmapPublishing field if non-nil, zero value otherwise.

### GetUseEnableIfmapPublishingOk

`func (o *GetRangeResponse) GetUseEnableIfmapPublishingOk() (*bool, bool)`

GetUseEnableIfmapPublishingOk returns a tuple with the UseEnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableIfmapPublishing

`func (o *GetRangeResponse) SetUseEnableIfmapPublishing(v bool)`

SetUseEnableIfmapPublishing sets UseEnableIfmapPublishing field to given value.

### HasUseEnableIfmapPublishing

`func (o *GetRangeResponse) HasUseEnableIfmapPublishing() bool`

HasUseEnableIfmapPublishing returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *GetRangeResponse) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *GetRangeResponse) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *GetRangeResponse) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *GetRangeResponse) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *GetRangeResponse) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *GetRangeResponse) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *GetRangeResponse) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *GetRangeResponse) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseKnownClients

`func (o *GetRangeResponse) GetUseKnownClients() bool`

GetUseKnownClients returns the UseKnownClients field if non-nil, zero value otherwise.

### GetUseKnownClientsOk

`func (o *GetRangeResponse) GetUseKnownClientsOk() (*bool, bool)`

GetUseKnownClientsOk returns a tuple with the UseKnownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKnownClients

`func (o *GetRangeResponse) SetUseKnownClients(v bool)`

SetUseKnownClients sets UseKnownClients field to given value.

### HasUseKnownClients

`func (o *GetRangeResponse) HasUseKnownClients() bool`

HasUseKnownClients returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *GetRangeResponse) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *GetRangeResponse) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *GetRangeResponse) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *GetRangeResponse) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetRangeResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetRangeResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetRangeResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetRangeResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMsOptions

`func (o *GetRangeResponse) GetUseMsOptions() bool`

GetUseMsOptions returns the UseMsOptions field if non-nil, zero value otherwise.

### GetUseMsOptionsOk

`func (o *GetRangeResponse) GetUseMsOptionsOk() (*bool, bool)`

GetUseMsOptionsOk returns a tuple with the UseMsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsOptions

`func (o *GetRangeResponse) SetUseMsOptions(v bool)`

SetUseMsOptions sets UseMsOptions field to given value.

### HasUseMsOptions

`func (o *GetRangeResponse) HasUseMsOptions() bool`

HasUseMsOptions returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetRangeResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetRangeResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetRangeResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetRangeResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetRangeResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetRangeResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetRangeResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetRangeResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetRangeResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetRangeResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetRangeResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetRangeResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetRangeResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetRangeResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetRangeResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetRangeResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *GetRangeResponse) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *GetRangeResponse) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *GetRangeResponse) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *GetRangeResponse) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUnknownClients

`func (o *GetRangeResponse) GetUseUnknownClients() bool`

GetUseUnknownClients returns the UseUnknownClients field if non-nil, zero value otherwise.

### GetUseUnknownClientsOk

`func (o *GetRangeResponse) GetUseUnknownClientsOk() (*bool, bool)`

GetUseUnknownClientsOk returns a tuple with the UseUnknownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUnknownClients

`func (o *GetRangeResponse) SetUseUnknownClients(v bool)`

SetUseUnknownClients sets UseUnknownClients field to given value.

### HasUseUnknownClients

`func (o *GetRangeResponse) HasUseUnknownClients() bool`

HasUseUnknownClients returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *GetRangeResponse) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *GetRangeResponse) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *GetRangeResponse) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *GetRangeResponse) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUuid

`func (o *GetRangeResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetRangeResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetRangeResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetRangeResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetRangeResponse) GetResult() Range`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRangeResponse) GetResultOk() (*Range, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRangeResponse) SetResult(v Range)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRangeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


