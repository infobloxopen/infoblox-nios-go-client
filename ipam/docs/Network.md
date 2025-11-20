# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
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
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Network) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Network) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Network) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Network) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthority

`func (o *Network) GetAuthority() bool`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *Network) GetAuthorityOk() (*bool, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *Network) SetAuthority(v bool)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *Network) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *Network) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *Network) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *Network) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *Network) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetBootfile

`func (o *Network) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *Network) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *Network) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *Network) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *Network) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *Network) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *Network) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *Network) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Network) GetCloudInfo() NetworkCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Network) GetCloudInfoOk() (*NetworkCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Network) SetCloudInfo(v NetworkCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Network) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCloudShared

`func (o *Network) GetCloudShared() bool`

GetCloudShared returns the CloudShared field if non-nil, zero value otherwise.

### GetCloudSharedOk

`func (o *Network) GetCloudSharedOk() (*bool, bool)`

GetCloudSharedOk returns a tuple with the CloudShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudShared

`func (o *Network) SetCloudShared(v bool)`

SetCloudShared sets CloudShared field to given value.

### HasCloudShared

`func (o *Network) HasCloudShared() bool`

HasCloudShared returns a boolean if a field has been set.

### GetComment

`func (o *Network) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Network) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Network) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Network) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConflictCount

`func (o *Network) GetConflictCount() int64`

GetConflictCount returns the ConflictCount field if non-nil, zero value otherwise.

### GetConflictCountOk

`func (o *Network) GetConflictCountOk() (*int64, bool)`

GetConflictCountOk returns a tuple with the ConflictCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCount

`func (o *Network) SetConflictCount(v int64)`

SetConflictCount sets ConflictCount field to given value.

### HasConflictCount

`func (o *Network) HasConflictCount() bool`

HasConflictCount returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Network) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Network) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Network) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Network) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *Network) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *Network) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *Network) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *Network) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *Network) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *Network) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *Network) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *Network) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *Network) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *Network) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *Network) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *Network) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUpdateFixedAddresses

`func (o *Network) GetDdnsUpdateFixedAddresses() bool`

GetDdnsUpdateFixedAddresses returns the DdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetDdnsUpdateFixedAddressesOk

`func (o *Network) GetDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetDdnsUpdateFixedAddressesOk returns a tuple with the DdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateFixedAddresses

`func (o *Network) SetDdnsUpdateFixedAddresses(v bool)`

SetDdnsUpdateFixedAddresses sets DdnsUpdateFixedAddresses field to given value.

### HasDdnsUpdateFixedAddresses

`func (o *Network) HasDdnsUpdateFixedAddresses() bool`

HasDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *Network) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *Network) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *Network) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *Network) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDeleteReason

`func (o *Network) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *Network) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *Network) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *Network) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### GetDenyBootp

`func (o *Network) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *Network) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *Network) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *Network) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *Network) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *Network) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *Network) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *Network) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *Network) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *Network) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *Network) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *Network) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *Network) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Network) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Network) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Network) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Network) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Network) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Network) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Network) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredBgpAs

`func (o *Network) GetDiscoveredBgpAs() string`

GetDiscoveredBgpAs returns the DiscoveredBgpAs field if non-nil, zero value otherwise.

### GetDiscoveredBgpAsOk

`func (o *Network) GetDiscoveredBgpAsOk() (*string, bool)`

GetDiscoveredBgpAsOk returns a tuple with the DiscoveredBgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBgpAs

`func (o *Network) SetDiscoveredBgpAs(v string)`

SetDiscoveredBgpAs sets DiscoveredBgpAs field to given value.

### HasDiscoveredBgpAs

`func (o *Network) HasDiscoveredBgpAs() bool`

HasDiscoveredBgpAs returns a boolean if a field has been set.

### GetDiscoveredBridgeDomain

`func (o *Network) GetDiscoveredBridgeDomain() string`

GetDiscoveredBridgeDomain returns the DiscoveredBridgeDomain field if non-nil, zero value otherwise.

### GetDiscoveredBridgeDomainOk

`func (o *Network) GetDiscoveredBridgeDomainOk() (*string, bool)`

GetDiscoveredBridgeDomainOk returns a tuple with the DiscoveredBridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBridgeDomain

`func (o *Network) SetDiscoveredBridgeDomain(v string)`

SetDiscoveredBridgeDomain sets DiscoveredBridgeDomain field to given value.

### HasDiscoveredBridgeDomain

`func (o *Network) HasDiscoveredBridgeDomain() bool`

HasDiscoveredBridgeDomain returns a boolean if a field has been set.

### GetDiscoveredTenant

`func (o *Network) GetDiscoveredTenant() string`

GetDiscoveredTenant returns the DiscoveredTenant field if non-nil, zero value otherwise.

### GetDiscoveredTenantOk

`func (o *Network) GetDiscoveredTenantOk() (*string, bool)`

GetDiscoveredTenantOk returns a tuple with the DiscoveredTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTenant

`func (o *Network) SetDiscoveredTenant(v string)`

SetDiscoveredTenant sets DiscoveredTenant field to given value.

### HasDiscoveredTenant

`func (o *Network) HasDiscoveredTenant() bool`

HasDiscoveredTenant returns a boolean if a field has been set.

### GetDiscoveredVlanId

`func (o *Network) GetDiscoveredVlanId() string`

GetDiscoveredVlanId returns the DiscoveredVlanId field if non-nil, zero value otherwise.

### GetDiscoveredVlanIdOk

`func (o *Network) GetDiscoveredVlanIdOk() (*string, bool)`

GetDiscoveredVlanIdOk returns a tuple with the DiscoveredVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVlanId

`func (o *Network) SetDiscoveredVlanId(v string)`

SetDiscoveredVlanId sets DiscoveredVlanId field to given value.

### HasDiscoveredVlanId

`func (o *Network) HasDiscoveredVlanId() bool`

HasDiscoveredVlanId returns a boolean if a field has been set.

### GetDiscoveredVlanName

`func (o *Network) GetDiscoveredVlanName() string`

GetDiscoveredVlanName returns the DiscoveredVlanName field if non-nil, zero value otherwise.

### GetDiscoveredVlanNameOk

`func (o *Network) GetDiscoveredVlanNameOk() (*string, bool)`

GetDiscoveredVlanNameOk returns a tuple with the DiscoveredVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVlanName

`func (o *Network) SetDiscoveredVlanName(v string)`

SetDiscoveredVlanName sets DiscoveredVlanName field to given value.

### HasDiscoveredVlanName

`func (o *Network) HasDiscoveredVlanName() bool`

HasDiscoveredVlanName returns a boolean if a field has been set.

### GetDiscoveredVrfDescription

`func (o *Network) GetDiscoveredVrfDescription() string`

GetDiscoveredVrfDescription returns the DiscoveredVrfDescription field if non-nil, zero value otherwise.

### GetDiscoveredVrfDescriptionOk

`func (o *Network) GetDiscoveredVrfDescriptionOk() (*string, bool)`

GetDiscoveredVrfDescriptionOk returns a tuple with the DiscoveredVrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfDescription

`func (o *Network) SetDiscoveredVrfDescription(v string)`

SetDiscoveredVrfDescription sets DiscoveredVrfDescription field to given value.

### HasDiscoveredVrfDescription

`func (o *Network) HasDiscoveredVrfDescription() bool`

HasDiscoveredVrfDescription returns a boolean if a field has been set.

### GetDiscoveredVrfName

`func (o *Network) GetDiscoveredVrfName() string`

GetDiscoveredVrfName returns the DiscoveredVrfName field if non-nil, zero value otherwise.

### GetDiscoveredVrfNameOk

`func (o *Network) GetDiscoveredVrfNameOk() (*string, bool)`

GetDiscoveredVrfNameOk returns a tuple with the DiscoveredVrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfName

`func (o *Network) SetDiscoveredVrfName(v string)`

SetDiscoveredVrfName sets DiscoveredVrfName field to given value.

### HasDiscoveredVrfName

`func (o *Network) HasDiscoveredVrfName() bool`

HasDiscoveredVrfName returns a boolean if a field has been set.

### GetDiscoveredVrfRd

`func (o *Network) GetDiscoveredVrfRd() string`

GetDiscoveredVrfRd returns the DiscoveredVrfRd field if non-nil, zero value otherwise.

### GetDiscoveredVrfRdOk

`func (o *Network) GetDiscoveredVrfRdOk() (*string, bool)`

GetDiscoveredVrfRdOk returns a tuple with the DiscoveredVrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfRd

`func (o *Network) SetDiscoveredVrfRd(v string)`

SetDiscoveredVrfRd sets DiscoveredVrfRd field to given value.

### HasDiscoveredVrfRd

`func (o *Network) HasDiscoveredVrfRd() bool`

HasDiscoveredVrfRd returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *Network) GetDiscoveryBasicPollSettings() NetworkDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *Network) GetDiscoveryBasicPollSettingsOk() (*NetworkDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *Network) SetDiscoveryBasicPollSettings(v NetworkDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *Network) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *Network) GetDiscoveryBlackoutSetting() NetworkDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *Network) GetDiscoveryBlackoutSettingOk() (*NetworkDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *Network) SetDiscoveryBlackoutSetting(v NetworkDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *Network) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryEngineType

`func (o *Network) GetDiscoveryEngineType() string`

GetDiscoveryEngineType returns the DiscoveryEngineType field if non-nil, zero value otherwise.

### GetDiscoveryEngineTypeOk

`func (o *Network) GetDiscoveryEngineTypeOk() (*string, bool)`

GetDiscoveryEngineTypeOk returns a tuple with the DiscoveryEngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngineType

`func (o *Network) SetDiscoveryEngineType(v string)`

SetDiscoveryEngineType sets DiscoveryEngineType field to given value.

### HasDiscoveryEngineType

`func (o *Network) HasDiscoveryEngineType() bool`

HasDiscoveryEngineType returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *Network) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *Network) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *Network) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *Network) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *Network) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *Network) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *Network) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *Network) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetEmailList

`func (o *Network) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *Network) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *Network) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *Network) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Network) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Network) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Network) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Network) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *Network) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *Network) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *Network) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *Network) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *Network) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *Network) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *Network) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *Network) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *Network) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *Network) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *Network) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *Network) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableIfmapPublishing

`func (o *Network) GetEnableIfmapPublishing() bool`

GetEnableIfmapPublishing returns the EnableIfmapPublishing field if non-nil, zero value otherwise.

### GetEnableIfmapPublishingOk

`func (o *Network) GetEnableIfmapPublishingOk() (*bool, bool)`

GetEnableIfmapPublishingOk returns a tuple with the EnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIfmapPublishing

`func (o *Network) SetEnableIfmapPublishing(v bool)`

SetEnableIfmapPublishing sets EnableIfmapPublishing field to given value.

### HasEnableIfmapPublishing

`func (o *Network) HasEnableIfmapPublishing() bool`

HasEnableIfmapPublishing returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Network) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Network) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Network) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Network) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *Network) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *Network) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *Network) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *Network) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *Network) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *Network) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *Network) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *Network) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetEndpointSources

`func (o *Network) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *Network) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *Network) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *Network) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Network) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Network) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Network) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Network) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Network) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Network) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Network) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Network) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Network) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Network) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Network) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Network) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *Network) GetFederatedRealms() []NetworkFederatedRealms`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *Network) GetFederatedRealmsOk() (*[]NetworkFederatedRealms, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *Network) SetFederatedRealms(v []NetworkFederatedRealms)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *Network) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *Network) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *Network) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *Network) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *Network) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *Network) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *Network) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *Network) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *Network) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *Network) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *Network) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *Network) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *Network) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *Network) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *Network) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *Network) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *Network) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *Network) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *Network) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *Network) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *Network) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetIpamEmailAddresses

`func (o *Network) GetIpamEmailAddresses() []string`

GetIpamEmailAddresses returns the IpamEmailAddresses field if non-nil, zero value otherwise.

### GetIpamEmailAddressesOk

`func (o *Network) GetIpamEmailAddressesOk() (*[]string, bool)`

GetIpamEmailAddressesOk returns a tuple with the IpamEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamEmailAddresses

`func (o *Network) SetIpamEmailAddresses(v []string)`

SetIpamEmailAddresses sets IpamEmailAddresses field to given value.

### HasIpamEmailAddresses

`func (o *Network) HasIpamEmailAddresses() bool`

HasIpamEmailAddresses returns a boolean if a field has been set.

### GetIpamThresholdSettings

`func (o *Network) GetIpamThresholdSettings() NetworkIpamThresholdSettings`

GetIpamThresholdSettings returns the IpamThresholdSettings field if non-nil, zero value otherwise.

### GetIpamThresholdSettingsOk

`func (o *Network) GetIpamThresholdSettingsOk() (*NetworkIpamThresholdSettings, bool)`

GetIpamThresholdSettingsOk returns a tuple with the IpamThresholdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamThresholdSettings

`func (o *Network) SetIpamThresholdSettings(v NetworkIpamThresholdSettings)`

SetIpamThresholdSettings sets IpamThresholdSettings field to given value.

### HasIpamThresholdSettings

`func (o *Network) HasIpamThresholdSettings() bool`

HasIpamThresholdSettings returns a boolean if a field has been set.

### GetIpamTrapSettings

`func (o *Network) GetIpamTrapSettings() NetworkIpamTrapSettings`

GetIpamTrapSettings returns the IpamTrapSettings field if non-nil, zero value otherwise.

### GetIpamTrapSettingsOk

`func (o *Network) GetIpamTrapSettingsOk() (*NetworkIpamTrapSettings, bool)`

GetIpamTrapSettingsOk returns a tuple with the IpamTrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamTrapSettings

`func (o *Network) SetIpamTrapSettings(v NetworkIpamTrapSettings)`

SetIpamTrapSettings sets IpamTrapSettings field to given value.

### HasIpamTrapSettings

`func (o *Network) HasIpamTrapSettings() bool`

HasIpamTrapSettings returns a boolean if a field has been set.

### GetIpv4addr

`func (o *Network) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *Network) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *Network) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *Network) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateSent

`func (o *Network) GetLastRirRegistrationUpdateSent() int64`

GetLastRirRegistrationUpdateSent returns the LastRirRegistrationUpdateSent field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateSentOk

`func (o *Network) GetLastRirRegistrationUpdateSentOk() (*int64, bool)`

GetLastRirRegistrationUpdateSentOk returns a tuple with the LastRirRegistrationUpdateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateSent

`func (o *Network) SetLastRirRegistrationUpdateSent(v int64)`

SetLastRirRegistrationUpdateSent sets LastRirRegistrationUpdateSent field to given value.

### HasLastRirRegistrationUpdateSent

`func (o *Network) HasLastRirRegistrationUpdateSent() bool`

HasLastRirRegistrationUpdateSent returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateStatus

`func (o *Network) GetLastRirRegistrationUpdateStatus() string`

GetLastRirRegistrationUpdateStatus returns the LastRirRegistrationUpdateStatus field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateStatusOk

`func (o *Network) GetLastRirRegistrationUpdateStatusOk() (*string, bool)`

GetLastRirRegistrationUpdateStatusOk returns a tuple with the LastRirRegistrationUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateStatus

`func (o *Network) SetLastRirRegistrationUpdateStatus(v string)`

SetLastRirRegistrationUpdateStatus sets LastRirRegistrationUpdateStatus field to given value.

### HasLastRirRegistrationUpdateStatus

`func (o *Network) HasLastRirRegistrationUpdateStatus() bool`

HasLastRirRegistrationUpdateStatus returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *Network) GetLeaseScavengeTime() int64`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *Network) GetLeaseScavengeTimeOk() (*int64, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *Network) SetLeaseScavengeTime(v int64)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *Network) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Network) GetLogicFilterRules() []NetworkLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Network) GetLogicFilterRulesOk() (*[]NetworkLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Network) SetLogicFilterRules(v []NetworkLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Network) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *Network) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *Network) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *Network) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *Network) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *Network) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *Network) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *Network) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *Network) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMembers

`func (o *Network) GetMembers() []NetworkMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Network) GetMembersOk() (*[]NetworkMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Network) SetMembers(v []NetworkMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Network) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *Network) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *Network) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *Network) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *Network) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *Network) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *Network) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *Network) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *Network) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Network) GetMsAdUserData() NetworkMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Network) GetMsAdUserDataOk() (*NetworkMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Network) SetMsAdUserData(v NetworkMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Network) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetmask

`func (o *Network) GetNetmask() int64`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Network) GetNetmaskOk() (*int64, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Network) SetNetmask(v int64)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *Network) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *Network) GetNetwork() NetworkNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Network) GetNetworkOk() (*NetworkNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Network) SetNetwork(v NetworkNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Network) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFuncCall

`func (o *Network) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *Network) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *Network) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *Network) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetNetworkContainer

`func (o *Network) GetNetworkContainer() string`

GetNetworkContainer returns the NetworkContainer field if non-nil, zero value otherwise.

### GetNetworkContainerOk

`func (o *Network) GetNetworkContainerOk() (*string, bool)`

GetNetworkContainerOk returns a tuple with the NetworkContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkContainer

`func (o *Network) SetNetworkContainer(v string)`

SetNetworkContainer sets NetworkContainer field to given value.

### HasNetworkContainer

`func (o *Network) HasNetworkContainer() bool`

HasNetworkContainer returns a boolean if a field has been set.

### GetNetworkView

`func (o *Network) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Network) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Network) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Network) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *Network) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *Network) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *Network) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *Network) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *Network) GetOptions() []NetworkOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Network) GetOptionsOk() (*[]NetworkOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Network) SetOptions(v []NetworkOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Network) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *Network) GetPortControlBlackoutSetting() NetworkPortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *Network) GetPortControlBlackoutSettingOk() (*NetworkPortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *Network) SetPortControlBlackoutSetting(v NetworkPortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *Network) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *Network) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *Network) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *Network) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *Network) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Network) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Network) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Network) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Network) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Network) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Network) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Network) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Network) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRir

`func (o *Network) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *Network) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *Network) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *Network) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *Network) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *Network) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *Network) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *Network) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *Network) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *Network) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *Network) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *Network) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *Network) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *Network) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *Network) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *Network) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *Network) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *Network) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *Network) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *Network) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *Network) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *Network) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *Network) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *Network) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetStaticHosts

`func (o *Network) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *Network) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *Network) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *Network) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *Network) GetSubscribeSettings() NetworkSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *Network) GetSubscribeSettingsOk() (*NetworkSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *Network) SetSubscribeSettings(v NetworkSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *Network) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *Network) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Network) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Network) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Network) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTotalHosts

`func (o *Network) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *Network) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *Network) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *Network) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUnmanaged

`func (o *Network) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *Network) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *Network) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *Network) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetUnmanagedCount

`func (o *Network) GetUnmanagedCount() int64`

GetUnmanagedCount returns the UnmanagedCount field if non-nil, zero value otherwise.

### GetUnmanagedCountOk

`func (o *Network) GetUnmanagedCountOk() (*int64, bool)`

GetUnmanagedCountOk returns a tuple with the UnmanagedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedCount

`func (o *Network) SetUnmanagedCount(v int64)`

SetUnmanagedCount sets UnmanagedCount field to given value.

### HasUnmanagedCount

`func (o *Network) HasUnmanagedCount() bool`

HasUnmanagedCount returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *Network) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *Network) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *Network) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *Network) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseAuthority

`func (o *Network) GetUseAuthority() bool`

GetUseAuthority returns the UseAuthority field if non-nil, zero value otherwise.

### GetUseAuthorityOk

`func (o *Network) GetUseAuthorityOk() (*bool, bool)`

GetUseAuthorityOk returns a tuple with the UseAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthority

`func (o *Network) SetUseAuthority(v bool)`

SetUseAuthority sets UseAuthority field to given value.

### HasUseAuthority

`func (o *Network) HasUseAuthority() bool`

HasUseAuthority returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *Network) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *Network) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *Network) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *Network) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseBootfile

`func (o *Network) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *Network) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *Network) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *Network) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *Network) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *Network) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *Network) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *Network) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Network) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Network) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Network) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Network) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *Network) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *Network) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *Network) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *Network) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *Network) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *Network) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *Network) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *Network) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDdnsUpdateFixedAddresses

`func (o *Network) GetUseDdnsUpdateFixedAddresses() bool`

GetUseDdnsUpdateFixedAddresses returns the UseDdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetUseDdnsUpdateFixedAddressesOk

`func (o *Network) GetUseDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetUseDdnsUpdateFixedAddressesOk returns a tuple with the UseDdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUpdateFixedAddresses

`func (o *Network) SetUseDdnsUpdateFixedAddresses(v bool)`

SetUseDdnsUpdateFixedAddresses sets UseDdnsUpdateFixedAddresses field to given value.

### HasUseDdnsUpdateFixedAddresses

`func (o *Network) HasUseDdnsUpdateFixedAddresses() bool`

HasUseDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetUseDdnsUseOption81

`func (o *Network) GetUseDdnsUseOption81() bool`

GetUseDdnsUseOption81 returns the UseDdnsUseOption81 field if non-nil, zero value otherwise.

### GetUseDdnsUseOption81Ok

`func (o *Network) GetUseDdnsUseOption81Ok() (*bool, bool)`

GetUseDdnsUseOption81Ok returns a tuple with the UseDdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUseOption81

`func (o *Network) SetUseDdnsUseOption81(v bool)`

SetUseDdnsUseOption81 sets UseDdnsUseOption81 field to given value.

### HasUseDdnsUseOption81

`func (o *Network) HasUseDdnsUseOption81() bool`

HasUseDdnsUseOption81 returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *Network) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *Network) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *Network) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *Network) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *Network) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *Network) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *Network) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *Network) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEmailList

`func (o *Network) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *Network) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *Network) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *Network) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Network) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Network) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Network) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Network) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *Network) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *Network) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *Network) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *Network) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *Network) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *Network) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *Network) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *Network) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseEnableIfmapPublishing

`func (o *Network) GetUseEnableIfmapPublishing() bool`

GetUseEnableIfmapPublishing returns the UseEnableIfmapPublishing field if non-nil, zero value otherwise.

### GetUseEnableIfmapPublishingOk

`func (o *Network) GetUseEnableIfmapPublishingOk() (*bool, bool)`

GetUseEnableIfmapPublishingOk returns a tuple with the UseEnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableIfmapPublishing

`func (o *Network) SetUseEnableIfmapPublishing(v bool)`

SetUseEnableIfmapPublishing sets UseEnableIfmapPublishing field to given value.

### HasUseEnableIfmapPublishing

`func (o *Network) HasUseEnableIfmapPublishing() bool`

HasUseEnableIfmapPublishing returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *Network) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *Network) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *Network) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *Network) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *Network) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *Network) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *Network) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *Network) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseIpamEmailAddresses

`func (o *Network) GetUseIpamEmailAddresses() bool`

GetUseIpamEmailAddresses returns the UseIpamEmailAddresses field if non-nil, zero value otherwise.

### GetUseIpamEmailAddressesOk

`func (o *Network) GetUseIpamEmailAddressesOk() (*bool, bool)`

GetUseIpamEmailAddressesOk returns a tuple with the UseIpamEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamEmailAddresses

`func (o *Network) SetUseIpamEmailAddresses(v bool)`

SetUseIpamEmailAddresses sets UseIpamEmailAddresses field to given value.

### HasUseIpamEmailAddresses

`func (o *Network) HasUseIpamEmailAddresses() bool`

HasUseIpamEmailAddresses returns a boolean if a field has been set.

### GetUseIpamThresholdSettings

`func (o *Network) GetUseIpamThresholdSettings() bool`

GetUseIpamThresholdSettings returns the UseIpamThresholdSettings field if non-nil, zero value otherwise.

### GetUseIpamThresholdSettingsOk

`func (o *Network) GetUseIpamThresholdSettingsOk() (*bool, bool)`

GetUseIpamThresholdSettingsOk returns a tuple with the UseIpamThresholdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamThresholdSettings

`func (o *Network) SetUseIpamThresholdSettings(v bool)`

SetUseIpamThresholdSettings sets UseIpamThresholdSettings field to given value.

### HasUseIpamThresholdSettings

`func (o *Network) HasUseIpamThresholdSettings() bool`

HasUseIpamThresholdSettings returns a boolean if a field has been set.

### GetUseIpamTrapSettings

`func (o *Network) GetUseIpamTrapSettings() bool`

GetUseIpamTrapSettings returns the UseIpamTrapSettings field if non-nil, zero value otherwise.

### GetUseIpamTrapSettingsOk

`func (o *Network) GetUseIpamTrapSettingsOk() (*bool, bool)`

GetUseIpamTrapSettingsOk returns a tuple with the UseIpamTrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamTrapSettings

`func (o *Network) SetUseIpamTrapSettings(v bool)`

SetUseIpamTrapSettings sets UseIpamTrapSettings field to given value.

### HasUseIpamTrapSettings

`func (o *Network) HasUseIpamTrapSettings() bool`

HasUseIpamTrapSettings returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *Network) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *Network) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *Network) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *Network) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Network) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Network) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Network) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Network) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMgmPrivate

`func (o *Network) GetUseMgmPrivate() bool`

GetUseMgmPrivate returns the UseMgmPrivate field if non-nil, zero value otherwise.

### GetUseMgmPrivateOk

`func (o *Network) GetUseMgmPrivateOk() (*bool, bool)`

GetUseMgmPrivateOk returns a tuple with the UseMgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmPrivate

`func (o *Network) SetUseMgmPrivate(v bool)`

SetUseMgmPrivate sets UseMgmPrivate field to given value.

### HasUseMgmPrivate

`func (o *Network) HasUseMgmPrivate() bool`

HasUseMgmPrivate returns a boolean if a field has been set.

### GetUseNextserver

`func (o *Network) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *Network) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *Network) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *Network) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *Network) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Network) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Network) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Network) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *Network) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *Network) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *Network) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *Network) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Network) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Network) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Network) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Network) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *Network) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *Network) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *Network) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *Network) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *Network) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *Network) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *Network) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *Network) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseZoneAssociations

`func (o *Network) GetUseZoneAssociations() bool`

GetUseZoneAssociations returns the UseZoneAssociations field if non-nil, zero value otherwise.

### GetUseZoneAssociationsOk

`func (o *Network) GetUseZoneAssociationsOk() (*bool, bool)`

GetUseZoneAssociationsOk returns a tuple with the UseZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseZoneAssociations

`func (o *Network) SetUseZoneAssociations(v bool)`

SetUseZoneAssociations sets UseZoneAssociations field to given value.

### HasUseZoneAssociations

`func (o *Network) HasUseZoneAssociations() bool`

HasUseZoneAssociations returns a boolean if a field has been set.

### GetUtilization

`func (o *Network) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *Network) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *Network) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *Network) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationUpdate

`func (o *Network) GetUtilizationUpdate() int64`

GetUtilizationUpdate returns the UtilizationUpdate field if non-nil, zero value otherwise.

### GetUtilizationUpdateOk

`func (o *Network) GetUtilizationUpdateOk() (*int64, bool)`

GetUtilizationUpdateOk returns a tuple with the UtilizationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationUpdate

`func (o *Network) SetUtilizationUpdate(v int64)`

SetUtilizationUpdate sets UtilizationUpdate field to given value.

### HasUtilizationUpdate

`func (o *Network) HasUtilizationUpdate() bool`

HasUtilizationUpdate returns a boolean if a field has been set.

### GetVlans

`func (o *Network) GetVlans() []NetworkVlans`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *Network) GetVlansOk() (*[]NetworkVlans, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *Network) SetVlans(v []NetworkVlans)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *Network) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *Network) GetZoneAssociations() []NetworkZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *Network) GetZoneAssociationsOk() (*[]NetworkZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *Network) SetZoneAssociations(v []NetworkZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *Network) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


