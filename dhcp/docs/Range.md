# Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FailoverAssociation** | Pointer to **string** | The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {range:range} must be set to &#39;FAILOVER&#39; or &#39;FAILOVER_MS&#39; if you want the failover association specified here to serve the range. | [optional] 
**FingerprintFilterRules** | Pointer to [**[]RangeFingerprintFilterRules**](RangeFingerprintFilterRules.md) | This field contains the fingerprint filters for this DHCP range. The appliance uses matching rules in these filters to select the address range from which it assigns a lease. | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage of DHCP range usage threshold above which range usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage of DHCP range usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \&quot;NONE\&quot;, \&quot;CLIENT\&quot;, or \&quot;MACADDR\&quot;. The default is \&quot;NONE\&quot;. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**IsSplitScope** | Pointer to **bool** | This field will be &#39;true&#39; if this particular range is part of a split scope. | [optional] [readonly] 
**KnownClients** | Pointer to **string** | Permission for known clients. This can be &#39;Allow&#39; or &#39;Deny&#39;. If set to &#39;Deny&#39; known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries. | [optional] 
**LeaseScavengeTime** | Pointer to **string** | An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
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
**NextAvailableIp** | Pointer to **map[string]interface{}** |  | [optional] 
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

## Methods

### NewRange

`func NewRange() *Range`

NewRange instantiates a new Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeWithDefaults

`func NewRangeWithDefaults() *Range`

NewRangeWithDefaults instantiates a new Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Range) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Range) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Range) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Range) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlwaysUpdateDns

`func (o *Range) GetAlwaysUpdateDns() bool`

GetAlwaysUpdateDns returns the AlwaysUpdateDns field if non-nil, zero value otherwise.

### GetAlwaysUpdateDnsOk

`func (o *Range) GetAlwaysUpdateDnsOk() (*bool, bool)`

GetAlwaysUpdateDnsOk returns a tuple with the AlwaysUpdateDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUpdateDns

`func (o *Range) SetAlwaysUpdateDns(v bool)`

SetAlwaysUpdateDns sets AlwaysUpdateDns field to given value.

### HasAlwaysUpdateDns

`func (o *Range) HasAlwaysUpdateDns() bool`

HasAlwaysUpdateDns returns a boolean if a field has been set.

### GetBootfile

`func (o *Range) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *Range) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *Range) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *Range) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *Range) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *Range) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *Range) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *Range) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Range) GetCloudInfo() RangeCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Range) GetCloudInfoOk() (*RangeCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Range) SetCloudInfo(v RangeCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Range) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Range) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Range) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Range) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Range) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Range) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Range) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Range) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Range) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *Range) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *Range) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *Range) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *Range) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDenyAllClients

`func (o *Range) GetDenyAllClients() bool`

GetDenyAllClients returns the DenyAllClients field if non-nil, zero value otherwise.

### GetDenyAllClientsOk

`func (o *Range) GetDenyAllClientsOk() (*bool, bool)`

GetDenyAllClientsOk returns a tuple with the DenyAllClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyAllClients

`func (o *Range) SetDenyAllClients(v bool)`

SetDenyAllClients sets DenyAllClients field to given value.

### HasDenyAllClients

`func (o *Range) HasDenyAllClients() bool`

HasDenyAllClients returns a boolean if a field has been set.

### GetDenyBootp

`func (o *Range) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *Range) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *Range) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *Range) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *Range) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *Range) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *Range) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *Range) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *Range) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *Range) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *Range) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *Range) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *Range) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Range) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Range) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Range) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Range) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Range) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Range) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Range) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *Range) GetDiscoveryBasicPollSettings() RangeDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *Range) GetDiscoveryBasicPollSettingsOk() (*RangeDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *Range) SetDiscoveryBasicPollSettings(v RangeDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *Range) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *Range) GetDiscoveryBlackoutSetting() RangeDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *Range) GetDiscoveryBlackoutSettingOk() (*RangeDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *Range) SetDiscoveryBlackoutSetting(v RangeDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *Range) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *Range) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *Range) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *Range) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *Range) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *Range) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *Range) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *Range) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *Range) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetEmailList

`func (o *Range) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *Range) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *Range) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *Range) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Range) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Range) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Range) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Range) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *Range) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *Range) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *Range) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *Range) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *Range) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *Range) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *Range) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *Range) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *Range) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *Range) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *Range) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *Range) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableIfmapPublishing

`func (o *Range) GetEnableIfmapPublishing() bool`

GetEnableIfmapPublishing returns the EnableIfmapPublishing field if non-nil, zero value otherwise.

### GetEnableIfmapPublishingOk

`func (o *Range) GetEnableIfmapPublishingOk() (*bool, bool)`

GetEnableIfmapPublishingOk returns a tuple with the EnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIfmapPublishing

`func (o *Range) SetEnableIfmapPublishing(v bool)`

SetEnableIfmapPublishing sets EnableIfmapPublishing field to given value.

### HasEnableIfmapPublishing

`func (o *Range) HasEnableIfmapPublishing() bool`

HasEnableIfmapPublishing returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Range) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Range) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Range) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Range) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *Range) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *Range) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *Range) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *Range) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *Range) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *Range) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *Range) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *Range) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetEndAddr

`func (o *Range) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *Range) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *Range) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *Range) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetEndpointSources

`func (o *Range) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *Range) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *Range) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *Range) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExclude

`func (o *Range) GetExclude() []RangeExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Range) GetExcludeOk() (*[]RangeExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Range) SetExclude(v []RangeExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *Range) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExtattrs

`func (o *Range) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Range) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Range) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Range) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFailoverAssociation

`func (o *Range) GetFailoverAssociation() string`

GetFailoverAssociation returns the FailoverAssociation field if non-nil, zero value otherwise.

### GetFailoverAssociationOk

`func (o *Range) GetFailoverAssociationOk() (*string, bool)`

GetFailoverAssociationOk returns a tuple with the FailoverAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAssociation

`func (o *Range) SetFailoverAssociation(v string)`

SetFailoverAssociation sets FailoverAssociation field to given value.

### HasFailoverAssociation

`func (o *Range) HasFailoverAssociation() bool`

HasFailoverAssociation returns a boolean if a field has been set.

### GetFingerprintFilterRules

`func (o *Range) GetFingerprintFilterRules() []RangeFingerprintFilterRules`

GetFingerprintFilterRules returns the FingerprintFilterRules field if non-nil, zero value otherwise.

### GetFingerprintFilterRulesOk

`func (o *Range) GetFingerprintFilterRulesOk() (*[]RangeFingerprintFilterRules, bool)`

GetFingerprintFilterRulesOk returns a tuple with the FingerprintFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintFilterRules

`func (o *Range) SetFingerprintFilterRules(v []RangeFingerprintFilterRules)`

SetFingerprintFilterRules sets FingerprintFilterRules field to given value.

### HasFingerprintFilterRules

`func (o *Range) HasFingerprintFilterRules() bool`

HasFingerprintFilterRules returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *Range) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *Range) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *Range) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *Range) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *Range) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *Range) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *Range) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *Range) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *Range) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *Range) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *Range) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *Range) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *Range) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *Range) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *Range) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *Range) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *Range) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *Range) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *Range) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *Range) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetIsSplitScope

`func (o *Range) GetIsSplitScope() bool`

GetIsSplitScope returns the IsSplitScope field if non-nil, zero value otherwise.

### GetIsSplitScopeOk

`func (o *Range) GetIsSplitScopeOk() (*bool, bool)`

GetIsSplitScopeOk returns a tuple with the IsSplitScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSplitScope

`func (o *Range) SetIsSplitScope(v bool)`

SetIsSplitScope sets IsSplitScope field to given value.

### HasIsSplitScope

`func (o *Range) HasIsSplitScope() bool`

HasIsSplitScope returns a boolean if a field has been set.

### GetKnownClients

`func (o *Range) GetKnownClients() string`

GetKnownClients returns the KnownClients field if non-nil, zero value otherwise.

### GetKnownClientsOk

`func (o *Range) GetKnownClientsOk() (*string, bool)`

GetKnownClientsOk returns a tuple with the KnownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClients

`func (o *Range) SetKnownClients(v string)`

SetKnownClients sets KnownClients field to given value.

### HasKnownClients

`func (o *Range) HasKnownClients() bool`

HasKnownClients returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *Range) GetLeaseScavengeTime() string`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *Range) GetLeaseScavengeTimeOk() (*string, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *Range) SetLeaseScavengeTime(v string)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *Range) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Range) GetLogicFilterRules() []RangeLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Range) GetLogicFilterRulesOk() (*[]RangeLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Range) SetLogicFilterRules(v []RangeLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Range) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *Range) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *Range) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *Range) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *Range) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *Range) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *Range) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *Range) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *Range) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMacFilterRules

`func (o *Range) GetMacFilterRules() []RangeMacFilterRules`

GetMacFilterRules returns the MacFilterRules field if non-nil, zero value otherwise.

### GetMacFilterRulesOk

`func (o *Range) GetMacFilterRulesOk() (*[]RangeMacFilterRules, bool)`

GetMacFilterRulesOk returns a tuple with the MacFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilterRules

`func (o *Range) SetMacFilterRules(v []RangeMacFilterRules)`

SetMacFilterRules sets MacFilterRules field to given value.

### HasMacFilterRules

`func (o *Range) HasMacFilterRules() bool`

HasMacFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *Range) GetMember() RangeMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Range) GetMemberOk() (*RangeMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Range) SetMember(v RangeMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *Range) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Range) GetMsAdUserData() RangeMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Range) GetMsAdUserDataOk() (*RangeMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Range) SetMsAdUserData(v RangeMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Range) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetMsOptions

`func (o *Range) GetMsOptions() []RangeMsOptions`

GetMsOptions returns the MsOptions field if non-nil, zero value otherwise.

### GetMsOptionsOk

`func (o *Range) GetMsOptionsOk() (*[]RangeMsOptions, bool)`

GetMsOptionsOk returns a tuple with the MsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsOptions

`func (o *Range) SetMsOptions(v []RangeMsOptions)`

SetMsOptions sets MsOptions field to given value.

### HasMsOptions

`func (o *Range) HasMsOptions() bool`

HasMsOptions returns a boolean if a field has been set.

### GetMsServer

`func (o *Range) GetMsServer() RangeMsServer`

GetMsServer returns the MsServer field if non-nil, zero value otherwise.

### GetMsServerOk

`func (o *Range) GetMsServerOk() (*RangeMsServer, bool)`

GetMsServerOk returns a tuple with the MsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServer

`func (o *Range) SetMsServer(v RangeMsServer)`

SetMsServer sets MsServer field to given value.

### HasMsServer

`func (o *Range) HasMsServer() bool`

HasMsServer returns a boolean if a field has been set.

### GetNacFilterRules

`func (o *Range) GetNacFilterRules() []RangeNacFilterRules`

GetNacFilterRules returns the NacFilterRules field if non-nil, zero value otherwise.

### GetNacFilterRulesOk

`func (o *Range) GetNacFilterRulesOk() (*[]RangeNacFilterRules, bool)`

GetNacFilterRulesOk returns a tuple with the NacFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacFilterRules

`func (o *Range) SetNacFilterRules(v []RangeNacFilterRules)`

SetNacFilterRules sets NacFilterRules field to given value.

### HasNacFilterRules

`func (o *Range) HasNacFilterRules() bool`

HasNacFilterRules returns a boolean if a field has been set.

### GetName

`func (o *Range) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Range) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Range) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Range) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Range) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Range) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Range) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Range) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Range) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Range) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Range) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Range) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextAvailableIp

`func (o *Range) GetNextAvailableIp() map[string]interface{}`

GetNextAvailableIp returns the NextAvailableIp field if non-nil, zero value otherwise.

### GetNextAvailableIpOk

`func (o *Range) GetNextAvailableIpOk() (*map[string]interface{}, bool)`

GetNextAvailableIpOk returns a tuple with the NextAvailableIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableIp

`func (o *Range) SetNextAvailableIp(v map[string]interface{})`

SetNextAvailableIp sets NextAvailableIp field to given value.

### HasNextAvailableIp

`func (o *Range) HasNextAvailableIp() bool`

HasNextAvailableIp returns a boolean if a field has been set.

### GetNextserver

`func (o *Range) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *Range) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *Range) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *Range) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *Range) GetOptionFilterRules() []RangeOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *Range) GetOptionFilterRulesOk() (*[]RangeOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *Range) SetOptionFilterRules(v []RangeOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *Range) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetOptions

`func (o *Range) GetOptions() []RangeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Range) GetOptionsOk() (*[]RangeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Range) SetOptions(v []RangeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Range) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *Range) GetPortControlBlackoutSetting() RangePortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *Range) GetPortControlBlackoutSettingOk() (*RangePortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *Range) SetPortControlBlackoutSetting(v RangePortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *Range) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *Range) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *Range) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *Range) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *Range) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Range) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Range) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Range) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Range) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRelayAgentFilterRules

`func (o *Range) GetRelayAgentFilterRules() []RangeRelayAgentFilterRules`

GetRelayAgentFilterRules returns the RelayAgentFilterRules field if non-nil, zero value otherwise.

### GetRelayAgentFilterRulesOk

`func (o *Range) GetRelayAgentFilterRulesOk() (*[]RangeRelayAgentFilterRules, bool)`

GetRelayAgentFilterRulesOk returns a tuple with the RelayAgentFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAgentFilterRules

`func (o *Range) SetRelayAgentFilterRules(v []RangeRelayAgentFilterRules)`

SetRelayAgentFilterRules sets RelayAgentFilterRules field to given value.

### HasRelayAgentFilterRules

`func (o *Range) HasRelayAgentFilterRules() bool`

HasRelayAgentFilterRules returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Range) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Range) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Range) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Range) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *Range) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *Range) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *Range) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *Range) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *Range) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *Range) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *Range) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *Range) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetSplitMember

`func (o *Range) GetSplitMember() RangeSplitMember`

GetSplitMember returns the SplitMember field if non-nil, zero value otherwise.

### GetSplitMemberOk

`func (o *Range) GetSplitMemberOk() (*RangeSplitMember, bool)`

GetSplitMemberOk returns a tuple with the SplitMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitMember

`func (o *Range) SetSplitMember(v RangeSplitMember)`

SetSplitMember sets SplitMember field to given value.

### HasSplitMember

`func (o *Range) HasSplitMember() bool`

HasSplitMember returns a boolean if a field has been set.

### GetSplitScopeExclusionPercent

`func (o *Range) GetSplitScopeExclusionPercent() int64`

GetSplitScopeExclusionPercent returns the SplitScopeExclusionPercent field if non-nil, zero value otherwise.

### GetSplitScopeExclusionPercentOk

`func (o *Range) GetSplitScopeExclusionPercentOk() (*int64, bool)`

GetSplitScopeExclusionPercentOk returns a tuple with the SplitScopeExclusionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitScopeExclusionPercent

`func (o *Range) SetSplitScopeExclusionPercent(v int64)`

SetSplitScopeExclusionPercent sets SplitScopeExclusionPercent field to given value.

### HasSplitScopeExclusionPercent

`func (o *Range) HasSplitScopeExclusionPercent() bool`

HasSplitScopeExclusionPercent returns a boolean if a field has been set.

### GetStartAddr

`func (o *Range) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *Range) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *Range) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *Range) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.

### GetStaticHosts

`func (o *Range) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *Range) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *Range) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *Range) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *Range) GetSubscribeSettings() RangeSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *Range) GetSubscribeSettingsOk() (*RangeSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *Range) SetSubscribeSettings(v RangeSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *Range) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *Range) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Range) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Range) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Range) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTotalHosts

`func (o *Range) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *Range) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *Range) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *Range) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUnknownClients

`func (o *Range) GetUnknownClients() string`

GetUnknownClients returns the UnknownClients field if non-nil, zero value otherwise.

### GetUnknownClientsOk

`func (o *Range) GetUnknownClientsOk() (*string, bool)`

GetUnknownClientsOk returns a tuple with the UnknownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownClients

`func (o *Range) SetUnknownClients(v string)`

SetUnknownClients sets UnknownClients field to given value.

### HasUnknownClients

`func (o *Range) HasUnknownClients() bool`

HasUnknownClients returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *Range) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *Range) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *Range) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *Range) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *Range) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *Range) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *Range) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *Range) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseBootfile

`func (o *Range) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *Range) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *Range) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *Range) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *Range) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *Range) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *Range) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *Range) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Range) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Range) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Range) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Range) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *Range) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *Range) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *Range) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *Range) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *Range) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *Range) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *Range) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *Range) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *Range) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *Range) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *Range) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *Range) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEmailList

`func (o *Range) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *Range) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *Range) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *Range) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Range) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Range) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Range) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Range) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *Range) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *Range) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *Range) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *Range) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *Range) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *Range) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *Range) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *Range) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseEnableIfmapPublishing

`func (o *Range) GetUseEnableIfmapPublishing() bool`

GetUseEnableIfmapPublishing returns the UseEnableIfmapPublishing field if non-nil, zero value otherwise.

### GetUseEnableIfmapPublishingOk

`func (o *Range) GetUseEnableIfmapPublishingOk() (*bool, bool)`

GetUseEnableIfmapPublishingOk returns a tuple with the UseEnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableIfmapPublishing

`func (o *Range) SetUseEnableIfmapPublishing(v bool)`

SetUseEnableIfmapPublishing sets UseEnableIfmapPublishing field to given value.

### HasUseEnableIfmapPublishing

`func (o *Range) HasUseEnableIfmapPublishing() bool`

HasUseEnableIfmapPublishing returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *Range) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *Range) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *Range) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *Range) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *Range) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *Range) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *Range) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *Range) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseKnownClients

`func (o *Range) GetUseKnownClients() bool`

GetUseKnownClients returns the UseKnownClients field if non-nil, zero value otherwise.

### GetUseKnownClientsOk

`func (o *Range) GetUseKnownClientsOk() (*bool, bool)`

GetUseKnownClientsOk returns a tuple with the UseKnownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKnownClients

`func (o *Range) SetUseKnownClients(v bool)`

SetUseKnownClients sets UseKnownClients field to given value.

### HasUseKnownClients

`func (o *Range) HasUseKnownClients() bool`

HasUseKnownClients returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *Range) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *Range) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *Range) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *Range) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Range) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Range) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Range) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Range) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMsOptions

`func (o *Range) GetUseMsOptions() bool`

GetUseMsOptions returns the UseMsOptions field if non-nil, zero value otherwise.

### GetUseMsOptionsOk

`func (o *Range) GetUseMsOptionsOk() (*bool, bool)`

GetUseMsOptionsOk returns a tuple with the UseMsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsOptions

`func (o *Range) SetUseMsOptions(v bool)`

SetUseMsOptions sets UseMsOptions field to given value.

### HasUseMsOptions

`func (o *Range) HasUseMsOptions() bool`

HasUseMsOptions returns a boolean if a field has been set.

### GetUseNextserver

`func (o *Range) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *Range) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *Range) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *Range) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *Range) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Range) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Range) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Range) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *Range) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *Range) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *Range) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *Range) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Range) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Range) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Range) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Range) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *Range) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *Range) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *Range) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *Range) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUnknownClients

`func (o *Range) GetUseUnknownClients() bool`

GetUseUnknownClients returns the UseUnknownClients field if non-nil, zero value otherwise.

### GetUseUnknownClientsOk

`func (o *Range) GetUseUnknownClientsOk() (*bool, bool)`

GetUseUnknownClientsOk returns a tuple with the UseUnknownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUnknownClients

`func (o *Range) SetUseUnknownClients(v bool)`

SetUseUnknownClients sets UseUnknownClients field to given value.

### HasUseUnknownClients

`func (o *Range) HasUseUnknownClients() bool`

HasUseUnknownClients returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *Range) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *Range) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *Range) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *Range) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


