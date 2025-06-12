# GetNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Authority** | Pointer to **bool** | Authority for the DHCP network. | [optional] 
**AutoCreateReversezone** | Pointer to **bool** | This flag controls whether reverse zones are automatically created when the network is added. | [optional] 
**Bootfile** | Pointer to **string** | The bootfile name for the network. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The bootserver address for the network. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**CloudInfo** | Pointer to [**NetworkCloudInfo**](NetworkCloudInfo.md) |  | [optional] 
**CloudShared** | Pointer to **bool** | Boolean flag to indicate if the network is shared with cloud. | [optional] 
**Comment** | Pointer to **string** | Comment for the network, maximum 256 characters. | [optional] 
**ConflictCount** | Pointer to **int64** | The number of conflicts discovered via network discovery. | [optional] [readonly] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_use_option81 is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DdnsUpdateFixedAddresses** | Pointer to **bool** | By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records. | [optional] 
**DdnsUseOption81** | Pointer to **bool** | The support for DHCP Option 81 at the network level. | [optional] 
**DeleteReason** | Pointer to **string** | The reason for deleting the RIR registration request. | [optional] 
**DenyBootp** | Pointer to **bool** | If set to true, BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of the network multiplied by 1000. This is the percentage of the total number of available IP addresses belonging to the network versus the total number of all IP addresses in network. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of the network. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether a network is disabled or not. When this is set to False, the network is enabled. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this network. | [optional] [readonly] 
**DiscoveredBgpAs** | Pointer to **string** | Number of the discovered BGP AS. When multiple BGP autonomous systems are discovered in the network, this field displays \&quot;Multiple\&quot;. | [optional] [readonly] 
**DiscoveredBridgeDomain** | Pointer to **string** | Discovered bridge domain. | [optional] 
**DiscoveredTenant** | Pointer to **string** | Discovered tenant. | [optional] 
**DiscoveredVlanId** | Pointer to **string** | The identifier of the discovered VLAN. When multiple VLANs are discovered in the network, this field displays \&quot;Multiple\&quot;. | [optional] [readonly] 
**DiscoveredVlanName** | Pointer to **string** | The name of the discovered VLAN. When multiple VLANs are discovered in the network, this field displays \&quot;Multiple\&quot;. | [optional] [readonly] 
**DiscoveredVrfDescription** | Pointer to **string** | Description of the discovered VRF. When multiple VRFs are discovered in the network, this field displays \&quot;Multiple\&quot;. | [optional] [readonly] 
**DiscoveredVrfName** | Pointer to **string** | The name of the discovered VRF. When multiple VRFs are discovered in the network, this field displays \&quot;Multiple\&quot;. | [optional] [readonly] 
**DiscoveredVrfRd** | Pointer to **string** | Route distinguisher of the discovered VRF. When multiple VRFs are discovered in the network, this field displays \&quot;Multiple\&quot;. | [optional] [readonly] 
**DiscoveryBasicPollSettings** | Pointer to [**NetworkDiscoveryBasicPollSettings**](NetworkDiscoveryBasicPollSettings.md) |  | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**NetworkDiscoveryBlackoutSetting**](NetworkDiscoveryBlackoutSetting.md) |  | [optional] 
**DiscoveryEngineType** | Pointer to **string** | The network discovery engine type. | [optional] [readonly] 
**DiscoveryMember** | Pointer to **string** | The member that will run discovery for this network. | [optional] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the network. | [optional] [readonly] 
**EmailList** | Pointer to **[]string** | The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableDhcpThresholds** | Pointer to **bool** | Determines if DHCP thresholds are enabled for the network. | [optional] 
**EnableDiscovery** | Pointer to **bool** | Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery is disabled. | [optional] 
**EnableEmailWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are sent through email. | [optional] 
**EnableIfmapPublishing** | Pointer to **bool** | Determines if IFMAP publishing is enabled for the network. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the network should be immediately enabled. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are send through SNMP. | [optional] 
**EndpointSources** | Pointer to **[]string** | The endpoints that provides data for the DHCP Network object. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FederatedRealms** | Pointer to [**[]NetworkFederatedRealms**](NetworkFederatedRealms.md) | This field contains the federated realms associated to this network | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage of DHCP network usage threshold above which network usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage of DHCP network usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \&quot;NONE\&quot;, \&quot;CLIENT\&quot;, or \&quot;MACADDR\&quot;. The default is \&quot;NONE\&quot;. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**IpamEmailAddresses** | Pointer to **[]string** | The e-mail lists to which the appliance sends IPAM threshold alarm e-mail messages. | [optional] 
**IpamThresholdSettings** | Pointer to [**NetworkIpamThresholdSettings**](NetworkIpamThresholdSettings.md) |  | [optional] 
**IpamTrapSettings** | Pointer to [**NetworkIpamTrapSettings**](NetworkIpamTrapSettings.md) |  | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the network. | [optional] 
**LastRirRegistrationUpdateSent** | Pointer to **int64** | The timestamp when the last RIR registration update was sent. | [optional] [readonly] 
**LastRirRegistrationUpdateStatus** | Pointer to **string** | Last RIR registration update status. | [optional] [readonly] 
**LeaseScavengeTime** | Pointer to **int64** | An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogicFilterRules** | Pointer to [**[]NetworkLogicFilterRules**](NetworkLogicFilterRules.md) | This field contains the logic filters to be applied on the this network. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**LowWaterMark** | Pointer to **int64** | The percentage of DHCP network usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**LowWaterMarkReset** | Pointer to **int64** | The percentage of DHCP network usage threshold below which network usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value. | [optional] 
**Members** | Pointer to [**[]NetworkMembers**](NetworkMembers.md) | A list of members or Microsoft (r) servers that serve DHCP for this network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \&quot;_struct\&quot; member to the struct type. | [optional] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MgmPrivateOverridable** | Pointer to **bool** | This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**NetworkMsAdUserData**](NetworkMsAdUserData.md) |  | [optional] 
**Netmask** | Pointer to **int64** | The netmask of the network in CIDR format. | [optional] 
**Network** | Pointer to [**NetworkNetwork**](NetworkNetwork.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**NetworkContainer** | Pointer to **string** | The network container to which this network belongs (if any). | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this network resides. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address of the next server that the host needs to boot. | [optional] 
**Options** | Pointer to [**[]NetworkOptions**](NetworkOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PortControlBlackoutSetting** | Pointer to [**NetworkPortControlBlackoutSetting**](NetworkPortControlBlackoutSetting.md) |  | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value of a DHCP Network object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**Rir** | Pointer to **string** | The registry (RIR) that allocated the network address space. | [optional] [readonly] 
**RirOrganization** | Pointer to **string** | The RIR organization assoicated with the network. | [optional] 
**RirRegistrationAction** | Pointer to **string** | The RIR registration action. | [optional] 
**RirRegistrationStatus** | Pointer to **string** | The registration status of the network in RIR. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | If the field is set to True, the discovery blackout setting will be used for port control blackout setting. | [optional] 
**SendRirRequest** | Pointer to **bool** | Determines whether to send the RIR registration request. | [optional] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in the network. | [optional] [readonly] 
**SubscribeSettings** | Pointer to [**NetworkSubscribeSettings**](NetworkSubscribeSettings.md) |  | [optional] 
**Template** | Pointer to **string** | If set on creation, the network is created according to the values specified in the selected template. | [optional] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in the network. | [optional] [readonly] 
**Unmanaged** | Pointer to **bool** | Determines whether the DHCP IPv4 Network is unmanaged or not. | [optional] 
**UnmanagedCount** | Pointer to **int64** | The number of unmanaged IP addresses as discovered by network discovery. | [optional] [readonly] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseAuthority** | Pointer to **bool** | Use flag for: authority | [optional] 
**UseBlackoutSetting** | Pointer to **bool** | Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDdnsUpdateFixedAddresses** | Pointer to **bool** | Use flag for: ddns_update_fixed_addresses | [optional] 
**UseDdnsUseOption81** | Pointer to **bool** | Use flag for: ddns_use_option81 | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseDiscoveryBasicPollingSettings** | Pointer to **bool** | Use flag for: discovery_basic_poll_settings | [optional] 
**UseEmailList** | Pointer to **bool** | Use flag for: email_list | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseEnableDhcpThresholds** | Pointer to **bool** | Use flag for: enable_dhcp_thresholds | [optional] 
**UseEnableDiscovery** | Pointer to **bool** | Use flag for: discovery_member , enable_discovery | [optional] 
**UseEnableIfmapPublishing** | Pointer to **bool** | Use flag for: enable_ifmap_publishing | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseIgnoreId** | Pointer to **bool** | Use flag for: ignore_id | [optional] 
**UseIpamEmailAddresses** | Pointer to **bool** | Use flag for: ipam_email_addresses | [optional] 
**UseIpamThresholdSettings** | Pointer to **bool** | Use flag for: ipam_threshold_settings | [optional] 
**UseIpamTrapSettings** | Pointer to **bool** | Use flag for: ipam_trap_settings | [optional] 
**UseLeaseScavengeTime** | Pointer to **bool** | Use flag for: lease_scavenge_time | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMgmPrivate** | Pointer to **bool** | Use flag for: mgm_private | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseSubscribeSettings** | Pointer to **bool** | Use flag for: subscribe_settings | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**UseZoneAssociations** | Pointer to **bool** | Use flag for: zone_associations | [optional] 
**Utilization** | Pointer to **int64** | The network utilization in percentage. | [optional] [readonly] 
**UtilizationUpdate** | Pointer to **int64** | The timestamp when the utilization statistics were last updated. | [optional] [readonly] 
**Vlans** | Pointer to [**[]NetworkVlans**](NetworkVlans.md) | List of VLANs assigned to Network. | [optional] 
**ZoneAssociations** | Pointer to [**[]NetworkZoneAssociations**](NetworkZoneAssociations.md) | The list of zones associated with this network. | [optional] 
**Result** | Pointer to [**Network**](Network.md) |  | [optional] 

## Methods

### NewGetNetworkResponse

`func NewGetNetworkResponse() *GetNetworkResponse`

NewGetNetworkResponse instantiates a new GetNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkResponseWithDefaults

`func NewGetNetworkResponseWithDefaults() *GetNetworkResponse`

NewGetNetworkResponseWithDefaults instantiates a new GetNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNetworkResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNetworkResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNetworkResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNetworkResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthority

`func (o *GetNetworkResponse) GetAuthority() bool`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GetNetworkResponse) GetAuthorityOk() (*bool, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GetNetworkResponse) SetAuthority(v bool)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GetNetworkResponse) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *GetNetworkResponse) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *GetNetworkResponse) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *GetNetworkResponse) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *GetNetworkResponse) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetBootfile

`func (o *GetNetworkResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetNetworkResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetNetworkResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetNetworkResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetNetworkResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetNetworkResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetNetworkResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetNetworkResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetNetworkResponse) GetCloudInfo() NetworkCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetNetworkResponse) GetCloudInfoOk() (*NetworkCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetNetworkResponse) SetCloudInfo(v NetworkCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetNetworkResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCloudShared

`func (o *GetNetworkResponse) GetCloudShared() bool`

GetCloudShared returns the CloudShared field if non-nil, zero value otherwise.

### GetCloudSharedOk

`func (o *GetNetworkResponse) GetCloudSharedOk() (*bool, bool)`

GetCloudSharedOk returns a tuple with the CloudShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudShared

`func (o *GetNetworkResponse) SetCloudShared(v bool)`

SetCloudShared sets CloudShared field to given value.

### HasCloudShared

`func (o *GetNetworkResponse) HasCloudShared() bool`

HasCloudShared returns a boolean if a field has been set.

### GetComment

`func (o *GetNetworkResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNetworkResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNetworkResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNetworkResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConflictCount

`func (o *GetNetworkResponse) GetConflictCount() int64`

GetConflictCount returns the ConflictCount field if non-nil, zero value otherwise.

### GetConflictCountOk

`func (o *GetNetworkResponse) GetConflictCountOk() (*int64, bool)`

GetConflictCountOk returns a tuple with the ConflictCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCount

`func (o *GetNetworkResponse) SetConflictCount(v int64)`

SetConflictCount sets ConflictCount field to given value.

### HasConflictCount

`func (o *GetNetworkResponse) HasConflictCount() bool`

HasConflictCount returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetNetworkResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetNetworkResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetNetworkResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetNetworkResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GetNetworkResponse) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GetNetworkResponse) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GetNetworkResponse) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GetNetworkResponse) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *GetNetworkResponse) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *GetNetworkResponse) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *GetNetworkResponse) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *GetNetworkResponse) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *GetNetworkResponse) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *GetNetworkResponse) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *GetNetworkResponse) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *GetNetworkResponse) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUpdateFixedAddresses

`func (o *GetNetworkResponse) GetDdnsUpdateFixedAddresses() bool`

GetDdnsUpdateFixedAddresses returns the DdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetDdnsUpdateFixedAddressesOk

`func (o *GetNetworkResponse) GetDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetDdnsUpdateFixedAddressesOk returns a tuple with the DdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateFixedAddresses

`func (o *GetNetworkResponse) SetDdnsUpdateFixedAddresses(v bool)`

SetDdnsUpdateFixedAddresses sets DdnsUpdateFixedAddresses field to given value.

### HasDdnsUpdateFixedAddresses

`func (o *GetNetworkResponse) HasDdnsUpdateFixedAddresses() bool`

HasDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *GetNetworkResponse) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *GetNetworkResponse) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *GetNetworkResponse) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *GetNetworkResponse) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDeleteReason

`func (o *GetNetworkResponse) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *GetNetworkResponse) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *GetNetworkResponse) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *GetNetworkResponse) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetNetworkResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetNetworkResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetNetworkResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetNetworkResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *GetNetworkResponse) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *GetNetworkResponse) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *GetNetworkResponse) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *GetNetworkResponse) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *GetNetworkResponse) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *GetNetworkResponse) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *GetNetworkResponse) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *GetNetworkResponse) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *GetNetworkResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetNetworkResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetNetworkResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetNetworkResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetNetworkResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetNetworkResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetNetworkResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetNetworkResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredBgpAs

`func (o *GetNetworkResponse) GetDiscoveredBgpAs() string`

GetDiscoveredBgpAs returns the DiscoveredBgpAs field if non-nil, zero value otherwise.

### GetDiscoveredBgpAsOk

`func (o *GetNetworkResponse) GetDiscoveredBgpAsOk() (*string, bool)`

GetDiscoveredBgpAsOk returns a tuple with the DiscoveredBgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBgpAs

`func (o *GetNetworkResponse) SetDiscoveredBgpAs(v string)`

SetDiscoveredBgpAs sets DiscoveredBgpAs field to given value.

### HasDiscoveredBgpAs

`func (o *GetNetworkResponse) HasDiscoveredBgpAs() bool`

HasDiscoveredBgpAs returns a boolean if a field has been set.

### GetDiscoveredBridgeDomain

`func (o *GetNetworkResponse) GetDiscoveredBridgeDomain() string`

GetDiscoveredBridgeDomain returns the DiscoveredBridgeDomain field if non-nil, zero value otherwise.

### GetDiscoveredBridgeDomainOk

`func (o *GetNetworkResponse) GetDiscoveredBridgeDomainOk() (*string, bool)`

GetDiscoveredBridgeDomainOk returns a tuple with the DiscoveredBridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBridgeDomain

`func (o *GetNetworkResponse) SetDiscoveredBridgeDomain(v string)`

SetDiscoveredBridgeDomain sets DiscoveredBridgeDomain field to given value.

### HasDiscoveredBridgeDomain

`func (o *GetNetworkResponse) HasDiscoveredBridgeDomain() bool`

HasDiscoveredBridgeDomain returns a boolean if a field has been set.

### GetDiscoveredTenant

`func (o *GetNetworkResponse) GetDiscoveredTenant() string`

GetDiscoveredTenant returns the DiscoveredTenant field if non-nil, zero value otherwise.

### GetDiscoveredTenantOk

`func (o *GetNetworkResponse) GetDiscoveredTenantOk() (*string, bool)`

GetDiscoveredTenantOk returns a tuple with the DiscoveredTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTenant

`func (o *GetNetworkResponse) SetDiscoveredTenant(v string)`

SetDiscoveredTenant sets DiscoveredTenant field to given value.

### HasDiscoveredTenant

`func (o *GetNetworkResponse) HasDiscoveredTenant() bool`

HasDiscoveredTenant returns a boolean if a field has been set.

### GetDiscoveredVlanId

`func (o *GetNetworkResponse) GetDiscoveredVlanId() string`

GetDiscoveredVlanId returns the DiscoveredVlanId field if non-nil, zero value otherwise.

### GetDiscoveredVlanIdOk

`func (o *GetNetworkResponse) GetDiscoveredVlanIdOk() (*string, bool)`

GetDiscoveredVlanIdOk returns a tuple with the DiscoveredVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVlanId

`func (o *GetNetworkResponse) SetDiscoveredVlanId(v string)`

SetDiscoveredVlanId sets DiscoveredVlanId field to given value.

### HasDiscoveredVlanId

`func (o *GetNetworkResponse) HasDiscoveredVlanId() bool`

HasDiscoveredVlanId returns a boolean if a field has been set.

### GetDiscoveredVlanName

`func (o *GetNetworkResponse) GetDiscoveredVlanName() string`

GetDiscoveredVlanName returns the DiscoveredVlanName field if non-nil, zero value otherwise.

### GetDiscoveredVlanNameOk

`func (o *GetNetworkResponse) GetDiscoveredVlanNameOk() (*string, bool)`

GetDiscoveredVlanNameOk returns a tuple with the DiscoveredVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVlanName

`func (o *GetNetworkResponse) SetDiscoveredVlanName(v string)`

SetDiscoveredVlanName sets DiscoveredVlanName field to given value.

### HasDiscoveredVlanName

`func (o *GetNetworkResponse) HasDiscoveredVlanName() bool`

HasDiscoveredVlanName returns a boolean if a field has been set.

### GetDiscoveredVrfDescription

`func (o *GetNetworkResponse) GetDiscoveredVrfDescription() string`

GetDiscoveredVrfDescription returns the DiscoveredVrfDescription field if non-nil, zero value otherwise.

### GetDiscoveredVrfDescriptionOk

`func (o *GetNetworkResponse) GetDiscoveredVrfDescriptionOk() (*string, bool)`

GetDiscoveredVrfDescriptionOk returns a tuple with the DiscoveredVrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfDescription

`func (o *GetNetworkResponse) SetDiscoveredVrfDescription(v string)`

SetDiscoveredVrfDescription sets DiscoveredVrfDescription field to given value.

### HasDiscoveredVrfDescription

`func (o *GetNetworkResponse) HasDiscoveredVrfDescription() bool`

HasDiscoveredVrfDescription returns a boolean if a field has been set.

### GetDiscoveredVrfName

`func (o *GetNetworkResponse) GetDiscoveredVrfName() string`

GetDiscoveredVrfName returns the DiscoveredVrfName field if non-nil, zero value otherwise.

### GetDiscoveredVrfNameOk

`func (o *GetNetworkResponse) GetDiscoveredVrfNameOk() (*string, bool)`

GetDiscoveredVrfNameOk returns a tuple with the DiscoveredVrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfName

`func (o *GetNetworkResponse) SetDiscoveredVrfName(v string)`

SetDiscoveredVrfName sets DiscoveredVrfName field to given value.

### HasDiscoveredVrfName

`func (o *GetNetworkResponse) HasDiscoveredVrfName() bool`

HasDiscoveredVrfName returns a boolean if a field has been set.

### GetDiscoveredVrfRd

`func (o *GetNetworkResponse) GetDiscoveredVrfRd() string`

GetDiscoveredVrfRd returns the DiscoveredVrfRd field if non-nil, zero value otherwise.

### GetDiscoveredVrfRdOk

`func (o *GetNetworkResponse) GetDiscoveredVrfRdOk() (*string, bool)`

GetDiscoveredVrfRdOk returns a tuple with the DiscoveredVrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfRd

`func (o *GetNetworkResponse) SetDiscoveredVrfRd(v string)`

SetDiscoveredVrfRd sets DiscoveredVrfRd field to given value.

### HasDiscoveredVrfRd

`func (o *GetNetworkResponse) HasDiscoveredVrfRd() bool`

HasDiscoveredVrfRd returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *GetNetworkResponse) GetDiscoveryBasicPollSettings() NetworkDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *GetNetworkResponse) GetDiscoveryBasicPollSettingsOk() (*NetworkDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *GetNetworkResponse) SetDiscoveryBasicPollSettings(v NetworkDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *GetNetworkResponse) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *GetNetworkResponse) GetDiscoveryBlackoutSetting() NetworkDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *GetNetworkResponse) GetDiscoveryBlackoutSettingOk() (*NetworkDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *GetNetworkResponse) SetDiscoveryBlackoutSetting(v NetworkDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *GetNetworkResponse) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryEngineType

`func (o *GetNetworkResponse) GetDiscoveryEngineType() string`

GetDiscoveryEngineType returns the DiscoveryEngineType field if non-nil, zero value otherwise.

### GetDiscoveryEngineTypeOk

`func (o *GetNetworkResponse) GetDiscoveryEngineTypeOk() (*string, bool)`

GetDiscoveryEngineTypeOk returns a tuple with the DiscoveryEngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngineType

`func (o *GetNetworkResponse) SetDiscoveryEngineType(v string)`

SetDiscoveryEngineType sets DiscoveryEngineType field to given value.

### HasDiscoveryEngineType

`func (o *GetNetworkResponse) HasDiscoveryEngineType() bool`

HasDiscoveryEngineType returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *GetNetworkResponse) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *GetNetworkResponse) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *GetNetworkResponse) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *GetNetworkResponse) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *GetNetworkResponse) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *GetNetworkResponse) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *GetNetworkResponse) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *GetNetworkResponse) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetEmailList

`func (o *GetNetworkResponse) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *GetNetworkResponse) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *GetNetworkResponse) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *GetNetworkResponse) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetNetworkResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetNetworkResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetNetworkResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetNetworkResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *GetNetworkResponse) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *GetNetworkResponse) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *GetNetworkResponse) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *GetNetworkResponse) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *GetNetworkResponse) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *GetNetworkResponse) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *GetNetworkResponse) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *GetNetworkResponse) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *GetNetworkResponse) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *GetNetworkResponse) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *GetNetworkResponse) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *GetNetworkResponse) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableIfmapPublishing

`func (o *GetNetworkResponse) GetEnableIfmapPublishing() bool`

GetEnableIfmapPublishing returns the EnableIfmapPublishing field if non-nil, zero value otherwise.

### GetEnableIfmapPublishingOk

`func (o *GetNetworkResponse) GetEnableIfmapPublishingOk() (*bool, bool)`

GetEnableIfmapPublishingOk returns a tuple with the EnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIfmapPublishing

`func (o *GetNetworkResponse) SetEnableIfmapPublishing(v bool)`

SetEnableIfmapPublishing sets EnableIfmapPublishing field to given value.

### HasEnableIfmapPublishing

`func (o *GetNetworkResponse) HasEnableIfmapPublishing() bool`

HasEnableIfmapPublishing returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *GetNetworkResponse) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *GetNetworkResponse) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *GetNetworkResponse) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *GetNetworkResponse) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetNetworkResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetNetworkResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetNetworkResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetNetworkResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *GetNetworkResponse) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *GetNetworkResponse) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *GetNetworkResponse) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *GetNetworkResponse) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetEndpointSources

`func (o *GetNetworkResponse) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *GetNetworkResponse) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *GetNetworkResponse) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *GetNetworkResponse) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetNetworkResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetNetworkResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetNetworkResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetNetworkResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *GetNetworkResponse) GetFederatedRealms() []NetworkFederatedRealms`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *GetNetworkResponse) GetFederatedRealmsOk() (*[]NetworkFederatedRealms, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *GetNetworkResponse) SetFederatedRealms(v []NetworkFederatedRealms)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *GetNetworkResponse) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetNetworkResponse) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetNetworkResponse) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetNetworkResponse) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetNetworkResponse) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *GetNetworkResponse) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *GetNetworkResponse) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *GetNetworkResponse) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *GetNetworkResponse) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GetNetworkResponse) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GetNetworkResponse) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GetNetworkResponse) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GetNetworkResponse) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *GetNetworkResponse) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *GetNetworkResponse) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *GetNetworkResponse) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *GetNetworkResponse) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *GetNetworkResponse) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *GetNetworkResponse) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *GetNetworkResponse) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *GetNetworkResponse) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetIpamEmailAddresses

`func (o *GetNetworkResponse) GetIpamEmailAddresses() []string`

GetIpamEmailAddresses returns the IpamEmailAddresses field if non-nil, zero value otherwise.

### GetIpamEmailAddressesOk

`func (o *GetNetworkResponse) GetIpamEmailAddressesOk() (*[]string, bool)`

GetIpamEmailAddressesOk returns a tuple with the IpamEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamEmailAddresses

`func (o *GetNetworkResponse) SetIpamEmailAddresses(v []string)`

SetIpamEmailAddresses sets IpamEmailAddresses field to given value.

### HasIpamEmailAddresses

`func (o *GetNetworkResponse) HasIpamEmailAddresses() bool`

HasIpamEmailAddresses returns a boolean if a field has been set.

### GetIpamThresholdSettings

`func (o *GetNetworkResponse) GetIpamThresholdSettings() NetworkIpamThresholdSettings`

GetIpamThresholdSettings returns the IpamThresholdSettings field if non-nil, zero value otherwise.

### GetIpamThresholdSettingsOk

`func (o *GetNetworkResponse) GetIpamThresholdSettingsOk() (*NetworkIpamThresholdSettings, bool)`

GetIpamThresholdSettingsOk returns a tuple with the IpamThresholdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamThresholdSettings

`func (o *GetNetworkResponse) SetIpamThresholdSettings(v NetworkIpamThresholdSettings)`

SetIpamThresholdSettings sets IpamThresholdSettings field to given value.

### HasIpamThresholdSettings

`func (o *GetNetworkResponse) HasIpamThresholdSettings() bool`

HasIpamThresholdSettings returns a boolean if a field has been set.

### GetIpamTrapSettings

`func (o *GetNetworkResponse) GetIpamTrapSettings() NetworkIpamTrapSettings`

GetIpamTrapSettings returns the IpamTrapSettings field if non-nil, zero value otherwise.

### GetIpamTrapSettingsOk

`func (o *GetNetworkResponse) GetIpamTrapSettingsOk() (*NetworkIpamTrapSettings, bool)`

GetIpamTrapSettingsOk returns a tuple with the IpamTrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamTrapSettings

`func (o *GetNetworkResponse) SetIpamTrapSettings(v NetworkIpamTrapSettings)`

SetIpamTrapSettings sets IpamTrapSettings field to given value.

### HasIpamTrapSettings

`func (o *GetNetworkResponse) HasIpamTrapSettings() bool`

HasIpamTrapSettings returns a boolean if a field has been set.

### GetIpv4addr

`func (o *GetNetworkResponse) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *GetNetworkResponse) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *GetNetworkResponse) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *GetNetworkResponse) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateSent

`func (o *GetNetworkResponse) GetLastRirRegistrationUpdateSent() int64`

GetLastRirRegistrationUpdateSent returns the LastRirRegistrationUpdateSent field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateSentOk

`func (o *GetNetworkResponse) GetLastRirRegistrationUpdateSentOk() (*int64, bool)`

GetLastRirRegistrationUpdateSentOk returns a tuple with the LastRirRegistrationUpdateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateSent

`func (o *GetNetworkResponse) SetLastRirRegistrationUpdateSent(v int64)`

SetLastRirRegistrationUpdateSent sets LastRirRegistrationUpdateSent field to given value.

### HasLastRirRegistrationUpdateSent

`func (o *GetNetworkResponse) HasLastRirRegistrationUpdateSent() bool`

HasLastRirRegistrationUpdateSent returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateStatus

`func (o *GetNetworkResponse) GetLastRirRegistrationUpdateStatus() string`

GetLastRirRegistrationUpdateStatus returns the LastRirRegistrationUpdateStatus field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateStatusOk

`func (o *GetNetworkResponse) GetLastRirRegistrationUpdateStatusOk() (*string, bool)`

GetLastRirRegistrationUpdateStatusOk returns a tuple with the LastRirRegistrationUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateStatus

`func (o *GetNetworkResponse) SetLastRirRegistrationUpdateStatus(v string)`

SetLastRirRegistrationUpdateStatus sets LastRirRegistrationUpdateStatus field to given value.

### HasLastRirRegistrationUpdateStatus

`func (o *GetNetworkResponse) HasLastRirRegistrationUpdateStatus() bool`

HasLastRirRegistrationUpdateStatus returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *GetNetworkResponse) GetLeaseScavengeTime() int64`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *GetNetworkResponse) GetLeaseScavengeTimeOk() (*int64, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *GetNetworkResponse) SetLeaseScavengeTime(v int64)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *GetNetworkResponse) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetNetworkResponse) GetLogicFilterRules() []NetworkLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetNetworkResponse) GetLogicFilterRulesOk() (*[]NetworkLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetNetworkResponse) SetLogicFilterRules(v []NetworkLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetNetworkResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetNetworkResponse) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetNetworkResponse) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetNetworkResponse) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetNetworkResponse) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *GetNetworkResponse) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *GetNetworkResponse) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *GetNetworkResponse) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *GetNetworkResponse) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMembers

`func (o *GetNetworkResponse) GetMembers() []NetworkMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetNetworkResponse) GetMembersOk() (*[]NetworkMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetNetworkResponse) SetMembers(v []NetworkMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetNetworkResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *GetNetworkResponse) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *GetNetworkResponse) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *GetNetworkResponse) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *GetNetworkResponse) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *GetNetworkResponse) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *GetNetworkResponse) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *GetNetworkResponse) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *GetNetworkResponse) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetNetworkResponse) GetMsAdUserData() NetworkMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetNetworkResponse) GetMsAdUserDataOk() (*NetworkMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetNetworkResponse) SetMsAdUserData(v NetworkMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetNetworkResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetmask

`func (o *GetNetworkResponse) GetNetmask() int64`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GetNetworkResponse) GetNetmaskOk() (*int64, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GetNetworkResponse) SetNetmask(v int64)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GetNetworkResponse) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *GetNetworkResponse) GetNetwork() NetworkNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetNetworkResponse) GetNetworkOk() (*NetworkNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetNetworkResponse) SetNetwork(v NetworkNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetNetworkResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFuncCall

`func (o *GetNetworkResponse) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *GetNetworkResponse) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *GetNetworkResponse) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *GetNetworkResponse) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetNetworkContainer

`func (o *GetNetworkResponse) GetNetworkContainer() string`

GetNetworkContainer returns the NetworkContainer field if non-nil, zero value otherwise.

### GetNetworkContainerOk

`func (o *GetNetworkResponse) GetNetworkContainerOk() (*string, bool)`

GetNetworkContainerOk returns a tuple with the NetworkContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkContainer

`func (o *GetNetworkResponse) SetNetworkContainer(v string)`

SetNetworkContainer sets NetworkContainer field to given value.

### HasNetworkContainer

`func (o *GetNetworkResponse) HasNetworkContainer() bool`

HasNetworkContainer returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetNetworkResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetNetworkResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetNetworkResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetNetworkResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *GetNetworkResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetNetworkResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetNetworkResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetNetworkResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *GetNetworkResponse) GetOptions() []NetworkOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetNetworkResponse) GetOptionsOk() (*[]NetworkOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetNetworkResponse) SetOptions(v []NetworkOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetNetworkResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *GetNetworkResponse) GetPortControlBlackoutSetting() NetworkPortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *GetNetworkResponse) GetPortControlBlackoutSettingOk() (*NetworkPortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *GetNetworkResponse) SetPortControlBlackoutSetting(v NetworkPortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *GetNetworkResponse) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetNetworkResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetNetworkResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetNetworkResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetNetworkResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetNetworkResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetNetworkResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetNetworkResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetNetworkResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetNetworkResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetNetworkResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetNetworkResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetNetworkResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRir

`func (o *GetNetworkResponse) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *GetNetworkResponse) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *GetNetworkResponse) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *GetNetworkResponse) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *GetNetworkResponse) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *GetNetworkResponse) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *GetNetworkResponse) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *GetNetworkResponse) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *GetNetworkResponse) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *GetNetworkResponse) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *GetNetworkResponse) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *GetNetworkResponse) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *GetNetworkResponse) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *GetNetworkResponse) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *GetNetworkResponse) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *GetNetworkResponse) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *GetNetworkResponse) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *GetNetworkResponse) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *GetNetworkResponse) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *GetNetworkResponse) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *GetNetworkResponse) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *GetNetworkResponse) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *GetNetworkResponse) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *GetNetworkResponse) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetStaticHosts

`func (o *GetNetworkResponse) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *GetNetworkResponse) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *GetNetworkResponse) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *GetNetworkResponse) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *GetNetworkResponse) GetSubscribeSettings() NetworkSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *GetNetworkResponse) GetSubscribeSettingsOk() (*NetworkSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *GetNetworkResponse) SetSubscribeSettings(v NetworkSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *GetNetworkResponse) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *GetNetworkResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetNetworkResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetNetworkResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetNetworkResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTotalHosts

`func (o *GetNetworkResponse) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *GetNetworkResponse) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *GetNetworkResponse) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *GetNetworkResponse) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUnmanaged

`func (o *GetNetworkResponse) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *GetNetworkResponse) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *GetNetworkResponse) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *GetNetworkResponse) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetUnmanagedCount

`func (o *GetNetworkResponse) GetUnmanagedCount() int64`

GetUnmanagedCount returns the UnmanagedCount field if non-nil, zero value otherwise.

### GetUnmanagedCountOk

`func (o *GetNetworkResponse) GetUnmanagedCountOk() (*int64, bool)`

GetUnmanagedCountOk returns a tuple with the UnmanagedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedCount

`func (o *GetNetworkResponse) SetUnmanagedCount(v int64)`

SetUnmanagedCount sets UnmanagedCount field to given value.

### HasUnmanagedCount

`func (o *GetNetworkResponse) HasUnmanagedCount() bool`

HasUnmanagedCount returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GetNetworkResponse) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GetNetworkResponse) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GetNetworkResponse) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GetNetworkResponse) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseAuthority

`func (o *GetNetworkResponse) GetUseAuthority() bool`

GetUseAuthority returns the UseAuthority field if non-nil, zero value otherwise.

### GetUseAuthorityOk

`func (o *GetNetworkResponse) GetUseAuthorityOk() (*bool, bool)`

GetUseAuthorityOk returns a tuple with the UseAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthority

`func (o *GetNetworkResponse) SetUseAuthority(v bool)`

SetUseAuthority sets UseAuthority field to given value.

### HasUseAuthority

`func (o *GetNetworkResponse) HasUseAuthority() bool`

HasUseAuthority returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *GetNetworkResponse) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *GetNetworkResponse) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *GetNetworkResponse) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *GetNetworkResponse) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetNetworkResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetNetworkResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetNetworkResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetNetworkResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetNetworkResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetNetworkResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetNetworkResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetNetworkResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetNetworkResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetNetworkResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetNetworkResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetNetworkResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *GetNetworkResponse) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *GetNetworkResponse) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *GetNetworkResponse) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *GetNetworkResponse) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *GetNetworkResponse) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *GetNetworkResponse) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *GetNetworkResponse) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *GetNetworkResponse) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDdnsUpdateFixedAddresses

`func (o *GetNetworkResponse) GetUseDdnsUpdateFixedAddresses() bool`

GetUseDdnsUpdateFixedAddresses returns the UseDdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetUseDdnsUpdateFixedAddressesOk

`func (o *GetNetworkResponse) GetUseDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetUseDdnsUpdateFixedAddressesOk returns a tuple with the UseDdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUpdateFixedAddresses

`func (o *GetNetworkResponse) SetUseDdnsUpdateFixedAddresses(v bool)`

SetUseDdnsUpdateFixedAddresses sets UseDdnsUpdateFixedAddresses field to given value.

### HasUseDdnsUpdateFixedAddresses

`func (o *GetNetworkResponse) HasUseDdnsUpdateFixedAddresses() bool`

HasUseDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetUseDdnsUseOption81

`func (o *GetNetworkResponse) GetUseDdnsUseOption81() bool`

GetUseDdnsUseOption81 returns the UseDdnsUseOption81 field if non-nil, zero value otherwise.

### GetUseDdnsUseOption81Ok

`func (o *GetNetworkResponse) GetUseDdnsUseOption81Ok() (*bool, bool)`

GetUseDdnsUseOption81Ok returns a tuple with the UseDdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUseOption81

`func (o *GetNetworkResponse) SetUseDdnsUseOption81(v bool)`

SetUseDdnsUseOption81 sets UseDdnsUseOption81 field to given value.

### HasUseDdnsUseOption81

`func (o *GetNetworkResponse) HasUseDdnsUseOption81() bool`

HasUseDdnsUseOption81 returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetNetworkResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetNetworkResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetNetworkResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetNetworkResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *GetNetworkResponse) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *GetNetworkResponse) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *GetNetworkResponse) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *GetNetworkResponse) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEmailList

`func (o *GetNetworkResponse) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *GetNetworkResponse) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *GetNetworkResponse) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *GetNetworkResponse) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetNetworkResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetNetworkResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetNetworkResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetNetworkResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *GetNetworkResponse) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *GetNetworkResponse) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *GetNetworkResponse) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *GetNetworkResponse) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *GetNetworkResponse) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *GetNetworkResponse) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *GetNetworkResponse) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *GetNetworkResponse) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseEnableIfmapPublishing

`func (o *GetNetworkResponse) GetUseEnableIfmapPublishing() bool`

GetUseEnableIfmapPublishing returns the UseEnableIfmapPublishing field if non-nil, zero value otherwise.

### GetUseEnableIfmapPublishingOk

`func (o *GetNetworkResponse) GetUseEnableIfmapPublishingOk() (*bool, bool)`

GetUseEnableIfmapPublishingOk returns a tuple with the UseEnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableIfmapPublishing

`func (o *GetNetworkResponse) SetUseEnableIfmapPublishing(v bool)`

SetUseEnableIfmapPublishing sets UseEnableIfmapPublishing field to given value.

### HasUseEnableIfmapPublishing

`func (o *GetNetworkResponse) HasUseEnableIfmapPublishing() bool`

HasUseEnableIfmapPublishing returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *GetNetworkResponse) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *GetNetworkResponse) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *GetNetworkResponse) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *GetNetworkResponse) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *GetNetworkResponse) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *GetNetworkResponse) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *GetNetworkResponse) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *GetNetworkResponse) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseIpamEmailAddresses

`func (o *GetNetworkResponse) GetUseIpamEmailAddresses() bool`

GetUseIpamEmailAddresses returns the UseIpamEmailAddresses field if non-nil, zero value otherwise.

### GetUseIpamEmailAddressesOk

`func (o *GetNetworkResponse) GetUseIpamEmailAddressesOk() (*bool, bool)`

GetUseIpamEmailAddressesOk returns a tuple with the UseIpamEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamEmailAddresses

`func (o *GetNetworkResponse) SetUseIpamEmailAddresses(v bool)`

SetUseIpamEmailAddresses sets UseIpamEmailAddresses field to given value.

### HasUseIpamEmailAddresses

`func (o *GetNetworkResponse) HasUseIpamEmailAddresses() bool`

HasUseIpamEmailAddresses returns a boolean if a field has been set.

### GetUseIpamThresholdSettings

`func (o *GetNetworkResponse) GetUseIpamThresholdSettings() bool`

GetUseIpamThresholdSettings returns the UseIpamThresholdSettings field if non-nil, zero value otherwise.

### GetUseIpamThresholdSettingsOk

`func (o *GetNetworkResponse) GetUseIpamThresholdSettingsOk() (*bool, bool)`

GetUseIpamThresholdSettingsOk returns a tuple with the UseIpamThresholdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamThresholdSettings

`func (o *GetNetworkResponse) SetUseIpamThresholdSettings(v bool)`

SetUseIpamThresholdSettings sets UseIpamThresholdSettings field to given value.

### HasUseIpamThresholdSettings

`func (o *GetNetworkResponse) HasUseIpamThresholdSettings() bool`

HasUseIpamThresholdSettings returns a boolean if a field has been set.

### GetUseIpamTrapSettings

`func (o *GetNetworkResponse) GetUseIpamTrapSettings() bool`

GetUseIpamTrapSettings returns the UseIpamTrapSettings field if non-nil, zero value otherwise.

### GetUseIpamTrapSettingsOk

`func (o *GetNetworkResponse) GetUseIpamTrapSettingsOk() (*bool, bool)`

GetUseIpamTrapSettingsOk returns a tuple with the UseIpamTrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamTrapSettings

`func (o *GetNetworkResponse) SetUseIpamTrapSettings(v bool)`

SetUseIpamTrapSettings sets UseIpamTrapSettings field to given value.

### HasUseIpamTrapSettings

`func (o *GetNetworkResponse) HasUseIpamTrapSettings() bool`

HasUseIpamTrapSettings returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *GetNetworkResponse) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *GetNetworkResponse) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *GetNetworkResponse) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *GetNetworkResponse) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetNetworkResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetNetworkResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetNetworkResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetNetworkResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMgmPrivate

`func (o *GetNetworkResponse) GetUseMgmPrivate() bool`

GetUseMgmPrivate returns the UseMgmPrivate field if non-nil, zero value otherwise.

### GetUseMgmPrivateOk

`func (o *GetNetworkResponse) GetUseMgmPrivateOk() (*bool, bool)`

GetUseMgmPrivateOk returns a tuple with the UseMgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmPrivate

`func (o *GetNetworkResponse) SetUseMgmPrivate(v bool)`

SetUseMgmPrivate sets UseMgmPrivate field to given value.

### HasUseMgmPrivate

`func (o *GetNetworkResponse) HasUseMgmPrivate() bool`

HasUseMgmPrivate returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetNetworkResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetNetworkResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetNetworkResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetNetworkResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetNetworkResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetNetworkResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetNetworkResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetNetworkResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetNetworkResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetNetworkResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetNetworkResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetNetworkResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetNetworkResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetNetworkResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetNetworkResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetNetworkResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *GetNetworkResponse) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *GetNetworkResponse) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *GetNetworkResponse) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *GetNetworkResponse) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *GetNetworkResponse) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *GetNetworkResponse) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *GetNetworkResponse) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *GetNetworkResponse) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseZoneAssociations

`func (o *GetNetworkResponse) GetUseZoneAssociations() bool`

GetUseZoneAssociations returns the UseZoneAssociations field if non-nil, zero value otherwise.

### GetUseZoneAssociationsOk

`func (o *GetNetworkResponse) GetUseZoneAssociationsOk() (*bool, bool)`

GetUseZoneAssociationsOk returns a tuple with the UseZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseZoneAssociations

`func (o *GetNetworkResponse) SetUseZoneAssociations(v bool)`

SetUseZoneAssociations sets UseZoneAssociations field to given value.

### HasUseZoneAssociations

`func (o *GetNetworkResponse) HasUseZoneAssociations() bool`

HasUseZoneAssociations returns a boolean if a field has been set.

### GetUtilization

`func (o *GetNetworkResponse) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *GetNetworkResponse) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *GetNetworkResponse) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *GetNetworkResponse) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationUpdate

`func (o *GetNetworkResponse) GetUtilizationUpdate() int64`

GetUtilizationUpdate returns the UtilizationUpdate field if non-nil, zero value otherwise.

### GetUtilizationUpdateOk

`func (o *GetNetworkResponse) GetUtilizationUpdateOk() (*int64, bool)`

GetUtilizationUpdateOk returns a tuple with the UtilizationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationUpdate

`func (o *GetNetworkResponse) SetUtilizationUpdate(v int64)`

SetUtilizationUpdate sets UtilizationUpdate field to given value.

### HasUtilizationUpdate

`func (o *GetNetworkResponse) HasUtilizationUpdate() bool`

HasUtilizationUpdate returns a boolean if a field has been set.

### GetVlans

`func (o *GetNetworkResponse) GetVlans() []NetworkVlans`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *GetNetworkResponse) GetVlansOk() (*[]NetworkVlans, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *GetNetworkResponse) SetVlans(v []NetworkVlans)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *GetNetworkResponse) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *GetNetworkResponse) GetZoneAssociations() []NetworkZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *GetNetworkResponse) GetZoneAssociationsOk() (*[]NetworkZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *GetNetworkResponse) SetZoneAssociations(v []NetworkZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *GetNetworkResponse) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.

### GetResult

`func (o *GetNetworkResponse) GetResult() Network`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNetworkResponse) GetResultOk() (*Network, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNetworkResponse) SetResult(v Network)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNetworkResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


