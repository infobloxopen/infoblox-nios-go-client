# GetNetworkcontainerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Authority** | Pointer to **bool** | Authority for the DHCP network container. | [optional] 
**AutoCreateReversezone** | Pointer to **bool** | This flag controls whether reverse zones are automatically created when the network is added. | [optional] 
**Bootfile** | Pointer to **string** | The boot server IPv4 Address or name in FQDN format for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot. | [optional] 
**Bootserver** | Pointer to **string** | The bootserver address for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**CloudInfo** | Pointer to [**NetworkcontainerCloudInfo**](NetworkcontainerCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the network container; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_use_option81 is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DdnsUpdateFixedAddresses** | Pointer to **bool** | By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records. | [optional] 
**DdnsUseOption81** | Pointer to **bool** | The support for DHCP Option 81 at the network container level. | [optional] 
**DeleteReason** | Pointer to **string** | The reason for deleting the RIR registration request. | [optional] 
**DenyBootp** | Pointer to **bool** | If set to True, BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this network container. | [optional] [readonly] 
**DiscoveryBasicPollSettings** | Pointer to [**NetworkcontainerDiscoveryBasicPollSettings**](NetworkcontainerDiscoveryBasicPollSettings.md) |  | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**NetworkcontainerDiscoveryBlackoutSetting**](NetworkcontainerDiscoveryBlackoutSetting.md) |  | [optional] 
**DiscoveryEngineType** | Pointer to **string** | The network discovery engine type. | [optional] [readonly] 
**DiscoveryMember** | Pointer to **string** | The member that will run discovery for this network container. | [optional] 
**EmailList** | Pointer to **[]string** | The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableDhcpThresholds** | Pointer to **bool** | Determines if DHCP thresholds are enabled for the network container. | [optional] 
**EnableDiscovery** | Pointer to **bool** | Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled. | [optional] 
**EnableEmailWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are sent through email. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the network container should be immediately enabled. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are send through SNMP. | [optional] 
**EndpointSources** | Pointer to **[]string** | The endpoints that provides data for the DHCP Network Container object. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FederatedRealms** | Pointer to [**[]NetworkcontainerFederatedRealms**](NetworkcontainerFederatedRealms.md) | This field contains the federated realms associated to this network container. | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage of DHCP network container usage threshold above which network container usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage of DHCP network container usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**IpamEmailAddresses** | Pointer to **[]string** | The e-mail lists to which the appliance sends IPAM threshold alarm e-mail messages. | [optional] 
**IpamThresholdSettings** | Pointer to [**NetworkcontainerIpamThresholdSettings**](NetworkcontainerIpamThresholdSettings.md) |  | [optional] 
**IpamTrapSettings** | Pointer to [**NetworkcontainerIpamTrapSettings**](NetworkcontainerIpamTrapSettings.md) |  | [optional] 
**LastRirRegistrationUpdateSent** | Pointer to **int64** | The timestamp when the last RIR registration update was sent. | [optional] [readonly] 
**LastRirRegistrationUpdateStatus** | Pointer to **string** | Last RIR registration update status. | [optional] [readonly] 
**LeaseScavengeTime** | Pointer to **int64** | An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogicFilterRules** | Pointer to [**[]NetworkcontainerLogicFilterRules**](NetworkcontainerLogicFilterRules.md) | This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**LowWaterMark** | Pointer to **int64** | The percentage of DHCP network container usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**LowWaterMarkReset** | Pointer to **int64** | The percentage of DHCP network container usage threshold below which network container usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value. | [optional] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MgmPrivateOverridable** | Pointer to **bool** | This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**NetworkcontainerMsAdUserData**](NetworkcontainerMsAdUserData.md) |  | [optional] 
**Network** | Pointer to [**NetworkcontainerNetwork**](NetworkcontainerNetwork.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**NetworkContainer** | Pointer to **string** | The network container to which this network belongs, if any. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this network resides. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address of the next server that the host needs to boot. | [optional] 
**Options** | Pointer to [**[]NetworkcontainerOptions**](NetworkcontainerOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PortControlBlackoutSetting** | Pointer to [**NetworkcontainerPortControlBlackoutSetting**](NetworkcontainerPortControlBlackoutSetting.md) |  | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value of a DHCP Network container object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**RemoveSubnets** | Pointer to **bool** | Remove subnets delete option. Determines whether all child objects should be removed alongside with the network container or child objects should be assigned to another parental container. By default child objects are deleted with the network container. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**Rir** | Pointer to **string** | The registry (RIR) that allocated the network container address space. | [optional] [readonly] 
**RirOrganization** | Pointer to **string** | The RIR organization assoicated with the network container. | [optional] 
**RirRegistrationAction** | Pointer to **string** | The RIR registration action. | [optional] 
**RirRegistrationStatus** | Pointer to **string** | The registration status of the network container in RIR. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | If the field is set to True, the discovery blackout setting will be used for port control blackout setting. | [optional] 
**SendRirRequest** | Pointer to **bool** | Determines whether to send the RIR registration request. | [optional] 
**SubscribeSettings** | Pointer to [**NetworkcontainerSubscribeSettings**](NetworkcontainerSubscribeSettings.md) |  | [optional] 
**Unmanaged** | Pointer to **bool** | Determines whether the network container is unmanaged or not. | [optional] 
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
**Utilization** | Pointer to **int64** | The network container utilization in percentage. | [optional] [readonly] 
**ZoneAssociations** | Pointer to [**[]NetworkcontainerZoneAssociations**](NetworkcontainerZoneAssociations.md) | The list of zones associated with this network. | [optional] 
**Result** | Pointer to [**Networkcontainer**](Networkcontainer.md) |  | [optional] 

## Methods

### NewGetNetworkcontainerResponse

`func NewGetNetworkcontainerResponse() *GetNetworkcontainerResponse`

NewGetNetworkcontainerResponse instantiates a new GetNetworkcontainerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkcontainerResponseWithDefaults

`func NewGetNetworkcontainerResponseWithDefaults() *GetNetworkcontainerResponse`

NewGetNetworkcontainerResponseWithDefaults instantiates a new GetNetworkcontainerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNetworkcontainerResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNetworkcontainerResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNetworkcontainerResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNetworkcontainerResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthority

`func (o *GetNetworkcontainerResponse) GetAuthority() bool`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GetNetworkcontainerResponse) GetAuthorityOk() (*bool, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GetNetworkcontainerResponse) SetAuthority(v bool)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GetNetworkcontainerResponse) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *GetNetworkcontainerResponse) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *GetNetworkcontainerResponse) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *GetNetworkcontainerResponse) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *GetNetworkcontainerResponse) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetBootfile

`func (o *GetNetworkcontainerResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetNetworkcontainerResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetNetworkcontainerResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetNetworkcontainerResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetNetworkcontainerResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetNetworkcontainerResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetNetworkcontainerResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetNetworkcontainerResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetNetworkcontainerResponse) GetCloudInfo() NetworkcontainerCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetNetworkcontainerResponse) GetCloudInfoOk() (*NetworkcontainerCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetNetworkcontainerResponse) SetCloudInfo(v NetworkcontainerCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetNetworkcontainerResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetNetworkcontainerResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNetworkcontainerResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNetworkcontainerResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNetworkcontainerResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetNetworkcontainerResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetNetworkcontainerResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetNetworkcontainerResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetNetworkcontainerResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GetNetworkcontainerResponse) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GetNetworkcontainerResponse) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GetNetworkcontainerResponse) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GetNetworkcontainerResponse) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *GetNetworkcontainerResponse) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *GetNetworkcontainerResponse) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *GetNetworkcontainerResponse) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *GetNetworkcontainerResponse) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *GetNetworkcontainerResponse) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *GetNetworkcontainerResponse) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *GetNetworkcontainerResponse) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *GetNetworkcontainerResponse) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUpdateFixedAddresses

`func (o *GetNetworkcontainerResponse) GetDdnsUpdateFixedAddresses() bool`

GetDdnsUpdateFixedAddresses returns the DdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetDdnsUpdateFixedAddressesOk

`func (o *GetNetworkcontainerResponse) GetDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetDdnsUpdateFixedAddressesOk returns a tuple with the DdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateFixedAddresses

`func (o *GetNetworkcontainerResponse) SetDdnsUpdateFixedAddresses(v bool)`

SetDdnsUpdateFixedAddresses sets DdnsUpdateFixedAddresses field to given value.

### HasDdnsUpdateFixedAddresses

`func (o *GetNetworkcontainerResponse) HasDdnsUpdateFixedAddresses() bool`

HasDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *GetNetworkcontainerResponse) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *GetNetworkcontainerResponse) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *GetNetworkcontainerResponse) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *GetNetworkcontainerResponse) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDeleteReason

`func (o *GetNetworkcontainerResponse) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *GetNetworkcontainerResponse) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *GetNetworkcontainerResponse) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *GetNetworkcontainerResponse) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetNetworkcontainerResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetNetworkcontainerResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetNetworkcontainerResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetNetworkcontainerResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetNetworkcontainerResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetNetworkcontainerResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetNetworkcontainerResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetNetworkcontainerResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *GetNetworkcontainerResponse) GetDiscoveryBasicPollSettings() NetworkcontainerDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *GetNetworkcontainerResponse) GetDiscoveryBasicPollSettingsOk() (*NetworkcontainerDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *GetNetworkcontainerResponse) SetDiscoveryBasicPollSettings(v NetworkcontainerDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *GetNetworkcontainerResponse) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *GetNetworkcontainerResponse) GetDiscoveryBlackoutSetting() NetworkcontainerDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *GetNetworkcontainerResponse) GetDiscoveryBlackoutSettingOk() (*NetworkcontainerDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *GetNetworkcontainerResponse) SetDiscoveryBlackoutSetting(v NetworkcontainerDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *GetNetworkcontainerResponse) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryEngineType

`func (o *GetNetworkcontainerResponse) GetDiscoveryEngineType() string`

GetDiscoveryEngineType returns the DiscoveryEngineType field if non-nil, zero value otherwise.

### GetDiscoveryEngineTypeOk

`func (o *GetNetworkcontainerResponse) GetDiscoveryEngineTypeOk() (*string, bool)`

GetDiscoveryEngineTypeOk returns a tuple with the DiscoveryEngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngineType

`func (o *GetNetworkcontainerResponse) SetDiscoveryEngineType(v string)`

SetDiscoveryEngineType sets DiscoveryEngineType field to given value.

### HasDiscoveryEngineType

`func (o *GetNetworkcontainerResponse) HasDiscoveryEngineType() bool`

HasDiscoveryEngineType returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *GetNetworkcontainerResponse) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *GetNetworkcontainerResponse) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *GetNetworkcontainerResponse) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *GetNetworkcontainerResponse) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetEmailList

`func (o *GetNetworkcontainerResponse) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *GetNetworkcontainerResponse) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *GetNetworkcontainerResponse) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *GetNetworkcontainerResponse) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetNetworkcontainerResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetNetworkcontainerResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetNetworkcontainerResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetNetworkcontainerResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *GetNetworkcontainerResponse) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *GetNetworkcontainerResponse) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *GetNetworkcontainerResponse) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *GetNetworkcontainerResponse) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *GetNetworkcontainerResponse) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *GetNetworkcontainerResponse) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *GetNetworkcontainerResponse) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *GetNetworkcontainerResponse) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *GetNetworkcontainerResponse) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *GetNetworkcontainerResponse) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *GetNetworkcontainerResponse) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *GetNetworkcontainerResponse) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *GetNetworkcontainerResponse) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *GetNetworkcontainerResponse) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *GetNetworkcontainerResponse) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *GetNetworkcontainerResponse) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetNetworkcontainerResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetNetworkcontainerResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetNetworkcontainerResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetNetworkcontainerResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *GetNetworkcontainerResponse) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *GetNetworkcontainerResponse) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *GetNetworkcontainerResponse) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *GetNetworkcontainerResponse) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetEndpointSources

`func (o *GetNetworkcontainerResponse) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *GetNetworkcontainerResponse) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *GetNetworkcontainerResponse) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *GetNetworkcontainerResponse) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetNetworkcontainerResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetNetworkcontainerResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetNetworkcontainerResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetNetworkcontainerResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetNetworkcontainerResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetNetworkcontainerResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetNetworkcontainerResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetNetworkcontainerResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetNetworkcontainerResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetNetworkcontainerResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetNetworkcontainerResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetNetworkcontainerResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *GetNetworkcontainerResponse) GetFederatedRealms() []NetworkcontainerFederatedRealms`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *GetNetworkcontainerResponse) GetFederatedRealmsOk() (*[]NetworkcontainerFederatedRealms, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *GetNetworkcontainerResponse) SetFederatedRealms(v []NetworkcontainerFederatedRealms)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *GetNetworkcontainerResponse) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetNetworkcontainerResponse) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetNetworkcontainerResponse) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetNetworkcontainerResponse) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetNetworkcontainerResponse) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *GetNetworkcontainerResponse) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *GetNetworkcontainerResponse) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *GetNetworkcontainerResponse) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *GetNetworkcontainerResponse) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GetNetworkcontainerResponse) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GetNetworkcontainerResponse) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GetNetworkcontainerResponse) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GetNetworkcontainerResponse) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *GetNetworkcontainerResponse) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *GetNetworkcontainerResponse) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *GetNetworkcontainerResponse) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *GetNetworkcontainerResponse) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *GetNetworkcontainerResponse) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *GetNetworkcontainerResponse) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *GetNetworkcontainerResponse) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *GetNetworkcontainerResponse) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetIpamEmailAddresses

`func (o *GetNetworkcontainerResponse) GetIpamEmailAddresses() []string`

GetIpamEmailAddresses returns the IpamEmailAddresses field if non-nil, zero value otherwise.

### GetIpamEmailAddressesOk

`func (o *GetNetworkcontainerResponse) GetIpamEmailAddressesOk() (*[]string, bool)`

GetIpamEmailAddressesOk returns a tuple with the IpamEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamEmailAddresses

`func (o *GetNetworkcontainerResponse) SetIpamEmailAddresses(v []string)`

SetIpamEmailAddresses sets IpamEmailAddresses field to given value.

### HasIpamEmailAddresses

`func (o *GetNetworkcontainerResponse) HasIpamEmailAddresses() bool`

HasIpamEmailAddresses returns a boolean if a field has been set.

### GetIpamThresholdSettings

`func (o *GetNetworkcontainerResponse) GetIpamThresholdSettings() NetworkcontainerIpamThresholdSettings`

GetIpamThresholdSettings returns the IpamThresholdSettings field if non-nil, zero value otherwise.

### GetIpamThresholdSettingsOk

`func (o *GetNetworkcontainerResponse) GetIpamThresholdSettingsOk() (*NetworkcontainerIpamThresholdSettings, bool)`

GetIpamThresholdSettingsOk returns a tuple with the IpamThresholdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamThresholdSettings

`func (o *GetNetworkcontainerResponse) SetIpamThresholdSettings(v NetworkcontainerIpamThresholdSettings)`

SetIpamThresholdSettings sets IpamThresholdSettings field to given value.

### HasIpamThresholdSettings

`func (o *GetNetworkcontainerResponse) HasIpamThresholdSettings() bool`

HasIpamThresholdSettings returns a boolean if a field has been set.

### GetIpamTrapSettings

`func (o *GetNetworkcontainerResponse) GetIpamTrapSettings() NetworkcontainerIpamTrapSettings`

GetIpamTrapSettings returns the IpamTrapSettings field if non-nil, zero value otherwise.

### GetIpamTrapSettingsOk

`func (o *GetNetworkcontainerResponse) GetIpamTrapSettingsOk() (*NetworkcontainerIpamTrapSettings, bool)`

GetIpamTrapSettingsOk returns a tuple with the IpamTrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamTrapSettings

`func (o *GetNetworkcontainerResponse) SetIpamTrapSettings(v NetworkcontainerIpamTrapSettings)`

SetIpamTrapSettings sets IpamTrapSettings field to given value.

### HasIpamTrapSettings

`func (o *GetNetworkcontainerResponse) HasIpamTrapSettings() bool`

HasIpamTrapSettings returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateSent

`func (o *GetNetworkcontainerResponse) GetLastRirRegistrationUpdateSent() int64`

GetLastRirRegistrationUpdateSent returns the LastRirRegistrationUpdateSent field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateSentOk

`func (o *GetNetworkcontainerResponse) GetLastRirRegistrationUpdateSentOk() (*int64, bool)`

GetLastRirRegistrationUpdateSentOk returns a tuple with the LastRirRegistrationUpdateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateSent

`func (o *GetNetworkcontainerResponse) SetLastRirRegistrationUpdateSent(v int64)`

SetLastRirRegistrationUpdateSent sets LastRirRegistrationUpdateSent field to given value.

### HasLastRirRegistrationUpdateSent

`func (o *GetNetworkcontainerResponse) HasLastRirRegistrationUpdateSent() bool`

HasLastRirRegistrationUpdateSent returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateStatus

`func (o *GetNetworkcontainerResponse) GetLastRirRegistrationUpdateStatus() string`

GetLastRirRegistrationUpdateStatus returns the LastRirRegistrationUpdateStatus field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateStatusOk

`func (o *GetNetworkcontainerResponse) GetLastRirRegistrationUpdateStatusOk() (*string, bool)`

GetLastRirRegistrationUpdateStatusOk returns a tuple with the LastRirRegistrationUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateStatus

`func (o *GetNetworkcontainerResponse) SetLastRirRegistrationUpdateStatus(v string)`

SetLastRirRegistrationUpdateStatus sets LastRirRegistrationUpdateStatus field to given value.

### HasLastRirRegistrationUpdateStatus

`func (o *GetNetworkcontainerResponse) HasLastRirRegistrationUpdateStatus() bool`

HasLastRirRegistrationUpdateStatus returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *GetNetworkcontainerResponse) GetLeaseScavengeTime() int64`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *GetNetworkcontainerResponse) GetLeaseScavengeTimeOk() (*int64, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *GetNetworkcontainerResponse) SetLeaseScavengeTime(v int64)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *GetNetworkcontainerResponse) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetNetworkcontainerResponse) GetLogicFilterRules() []NetworkcontainerLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetNetworkcontainerResponse) GetLogicFilterRulesOk() (*[]NetworkcontainerLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetNetworkcontainerResponse) SetLogicFilterRules(v []NetworkcontainerLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetNetworkcontainerResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetNetworkcontainerResponse) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetNetworkcontainerResponse) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetNetworkcontainerResponse) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetNetworkcontainerResponse) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *GetNetworkcontainerResponse) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *GetNetworkcontainerResponse) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *GetNetworkcontainerResponse) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *GetNetworkcontainerResponse) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *GetNetworkcontainerResponse) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *GetNetworkcontainerResponse) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *GetNetworkcontainerResponse) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *GetNetworkcontainerResponse) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *GetNetworkcontainerResponse) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *GetNetworkcontainerResponse) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *GetNetworkcontainerResponse) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *GetNetworkcontainerResponse) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetNetworkcontainerResponse) GetMsAdUserData() NetworkcontainerMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetNetworkcontainerResponse) GetMsAdUserDataOk() (*NetworkcontainerMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetNetworkcontainerResponse) SetMsAdUserData(v NetworkcontainerMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetNetworkcontainerResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *GetNetworkcontainerResponse) GetNetwork() NetworkcontainerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetNetworkcontainerResponse) GetNetworkOk() (*NetworkcontainerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetNetworkcontainerResponse) SetNetwork(v NetworkcontainerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetNetworkcontainerResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFuncCall

`func (o *GetNetworkcontainerResponse) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *GetNetworkcontainerResponse) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *GetNetworkcontainerResponse) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *GetNetworkcontainerResponse) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetNetworkContainer

`func (o *GetNetworkcontainerResponse) GetNetworkContainer() string`

GetNetworkContainer returns the NetworkContainer field if non-nil, zero value otherwise.

### GetNetworkContainerOk

`func (o *GetNetworkcontainerResponse) GetNetworkContainerOk() (*string, bool)`

GetNetworkContainerOk returns a tuple with the NetworkContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkContainer

`func (o *GetNetworkcontainerResponse) SetNetworkContainer(v string)`

SetNetworkContainer sets NetworkContainer field to given value.

### HasNetworkContainer

`func (o *GetNetworkcontainerResponse) HasNetworkContainer() bool`

HasNetworkContainer returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetNetworkcontainerResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetNetworkcontainerResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetNetworkcontainerResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetNetworkcontainerResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *GetNetworkcontainerResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetNetworkcontainerResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetNetworkcontainerResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetNetworkcontainerResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *GetNetworkcontainerResponse) GetOptions() []NetworkcontainerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetNetworkcontainerResponse) GetOptionsOk() (*[]NetworkcontainerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetNetworkcontainerResponse) SetOptions(v []NetworkcontainerOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetNetworkcontainerResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *GetNetworkcontainerResponse) GetPortControlBlackoutSetting() NetworkcontainerPortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *GetNetworkcontainerResponse) GetPortControlBlackoutSettingOk() (*NetworkcontainerPortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *GetNetworkcontainerResponse) SetPortControlBlackoutSetting(v NetworkcontainerPortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *GetNetworkcontainerResponse) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetNetworkcontainerResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetNetworkcontainerResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetNetworkcontainerResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetNetworkcontainerResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetNetworkcontainerResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetNetworkcontainerResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetNetworkcontainerResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetNetworkcontainerResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRemoveSubnets

`func (o *GetNetworkcontainerResponse) GetRemoveSubnets() bool`

GetRemoveSubnets returns the RemoveSubnets field if non-nil, zero value otherwise.

### GetRemoveSubnetsOk

`func (o *GetNetworkcontainerResponse) GetRemoveSubnetsOk() (*bool, bool)`

GetRemoveSubnetsOk returns a tuple with the RemoveSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSubnets

`func (o *GetNetworkcontainerResponse) SetRemoveSubnets(v bool)`

SetRemoveSubnets sets RemoveSubnets field to given value.

### HasRemoveSubnets

`func (o *GetNetworkcontainerResponse) HasRemoveSubnets() bool`

HasRemoveSubnets returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetNetworkcontainerResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetNetworkcontainerResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetNetworkcontainerResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetNetworkcontainerResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRir

`func (o *GetNetworkcontainerResponse) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *GetNetworkcontainerResponse) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *GetNetworkcontainerResponse) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *GetNetworkcontainerResponse) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *GetNetworkcontainerResponse) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *GetNetworkcontainerResponse) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *GetNetworkcontainerResponse) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *GetNetworkcontainerResponse) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *GetNetworkcontainerResponse) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *GetNetworkcontainerResponse) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *GetNetworkcontainerResponse) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *GetNetworkcontainerResponse) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *GetNetworkcontainerResponse) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *GetNetworkcontainerResponse) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *GetNetworkcontainerResponse) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *GetNetworkcontainerResponse) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *GetNetworkcontainerResponse) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *GetNetworkcontainerResponse) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *GetNetworkcontainerResponse) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *GetNetworkcontainerResponse) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *GetNetworkcontainerResponse) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *GetNetworkcontainerResponse) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *GetNetworkcontainerResponse) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *GetNetworkcontainerResponse) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *GetNetworkcontainerResponse) GetSubscribeSettings() NetworkcontainerSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *GetNetworkcontainerResponse) GetSubscribeSettingsOk() (*NetworkcontainerSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *GetNetworkcontainerResponse) SetSubscribeSettings(v NetworkcontainerSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *GetNetworkcontainerResponse) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetUnmanaged

`func (o *GetNetworkcontainerResponse) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *GetNetworkcontainerResponse) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *GetNetworkcontainerResponse) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *GetNetworkcontainerResponse) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GetNetworkcontainerResponse) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GetNetworkcontainerResponse) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GetNetworkcontainerResponse) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GetNetworkcontainerResponse) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseAuthority

`func (o *GetNetworkcontainerResponse) GetUseAuthority() bool`

GetUseAuthority returns the UseAuthority field if non-nil, zero value otherwise.

### GetUseAuthorityOk

`func (o *GetNetworkcontainerResponse) GetUseAuthorityOk() (*bool, bool)`

GetUseAuthorityOk returns a tuple with the UseAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthority

`func (o *GetNetworkcontainerResponse) SetUseAuthority(v bool)`

SetUseAuthority sets UseAuthority field to given value.

### HasUseAuthority

`func (o *GetNetworkcontainerResponse) HasUseAuthority() bool`

HasUseAuthority returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *GetNetworkcontainerResponse) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *GetNetworkcontainerResponse) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *GetNetworkcontainerResponse) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *GetNetworkcontainerResponse) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetNetworkcontainerResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetNetworkcontainerResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetNetworkcontainerResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetNetworkcontainerResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetNetworkcontainerResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetNetworkcontainerResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetNetworkcontainerResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetNetworkcontainerResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetNetworkcontainerResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetNetworkcontainerResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetNetworkcontainerResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetNetworkcontainerResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *GetNetworkcontainerResponse) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *GetNetworkcontainerResponse) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *GetNetworkcontainerResponse) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *GetNetworkcontainerResponse) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *GetNetworkcontainerResponse) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *GetNetworkcontainerResponse) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *GetNetworkcontainerResponse) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *GetNetworkcontainerResponse) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDdnsUpdateFixedAddresses

`func (o *GetNetworkcontainerResponse) GetUseDdnsUpdateFixedAddresses() bool`

GetUseDdnsUpdateFixedAddresses returns the UseDdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetUseDdnsUpdateFixedAddressesOk

`func (o *GetNetworkcontainerResponse) GetUseDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetUseDdnsUpdateFixedAddressesOk returns a tuple with the UseDdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUpdateFixedAddresses

`func (o *GetNetworkcontainerResponse) SetUseDdnsUpdateFixedAddresses(v bool)`

SetUseDdnsUpdateFixedAddresses sets UseDdnsUpdateFixedAddresses field to given value.

### HasUseDdnsUpdateFixedAddresses

`func (o *GetNetworkcontainerResponse) HasUseDdnsUpdateFixedAddresses() bool`

HasUseDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetUseDdnsUseOption81

`func (o *GetNetworkcontainerResponse) GetUseDdnsUseOption81() bool`

GetUseDdnsUseOption81 returns the UseDdnsUseOption81 field if non-nil, zero value otherwise.

### GetUseDdnsUseOption81Ok

`func (o *GetNetworkcontainerResponse) GetUseDdnsUseOption81Ok() (*bool, bool)`

GetUseDdnsUseOption81Ok returns a tuple with the UseDdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUseOption81

`func (o *GetNetworkcontainerResponse) SetUseDdnsUseOption81(v bool)`

SetUseDdnsUseOption81 sets UseDdnsUseOption81 field to given value.

### HasUseDdnsUseOption81

`func (o *GetNetworkcontainerResponse) HasUseDdnsUseOption81() bool`

HasUseDdnsUseOption81 returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetNetworkcontainerResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetNetworkcontainerResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetNetworkcontainerResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetNetworkcontainerResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *GetNetworkcontainerResponse) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *GetNetworkcontainerResponse) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *GetNetworkcontainerResponse) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *GetNetworkcontainerResponse) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEmailList

`func (o *GetNetworkcontainerResponse) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *GetNetworkcontainerResponse) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *GetNetworkcontainerResponse) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *GetNetworkcontainerResponse) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetNetworkcontainerResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetNetworkcontainerResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetNetworkcontainerResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetNetworkcontainerResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *GetNetworkcontainerResponse) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *GetNetworkcontainerResponse) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *GetNetworkcontainerResponse) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *GetNetworkcontainerResponse) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *GetNetworkcontainerResponse) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *GetNetworkcontainerResponse) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *GetNetworkcontainerResponse) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *GetNetworkcontainerResponse) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *GetNetworkcontainerResponse) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *GetNetworkcontainerResponse) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *GetNetworkcontainerResponse) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *GetNetworkcontainerResponse) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *GetNetworkcontainerResponse) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *GetNetworkcontainerResponse) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *GetNetworkcontainerResponse) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *GetNetworkcontainerResponse) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseIpamEmailAddresses

`func (o *GetNetworkcontainerResponse) GetUseIpamEmailAddresses() bool`

GetUseIpamEmailAddresses returns the UseIpamEmailAddresses field if non-nil, zero value otherwise.

### GetUseIpamEmailAddressesOk

`func (o *GetNetworkcontainerResponse) GetUseIpamEmailAddressesOk() (*bool, bool)`

GetUseIpamEmailAddressesOk returns a tuple with the UseIpamEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamEmailAddresses

`func (o *GetNetworkcontainerResponse) SetUseIpamEmailAddresses(v bool)`

SetUseIpamEmailAddresses sets UseIpamEmailAddresses field to given value.

### HasUseIpamEmailAddresses

`func (o *GetNetworkcontainerResponse) HasUseIpamEmailAddresses() bool`

HasUseIpamEmailAddresses returns a boolean if a field has been set.

### GetUseIpamThresholdSettings

`func (o *GetNetworkcontainerResponse) GetUseIpamThresholdSettings() bool`

GetUseIpamThresholdSettings returns the UseIpamThresholdSettings field if non-nil, zero value otherwise.

### GetUseIpamThresholdSettingsOk

`func (o *GetNetworkcontainerResponse) GetUseIpamThresholdSettingsOk() (*bool, bool)`

GetUseIpamThresholdSettingsOk returns a tuple with the UseIpamThresholdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamThresholdSettings

`func (o *GetNetworkcontainerResponse) SetUseIpamThresholdSettings(v bool)`

SetUseIpamThresholdSettings sets UseIpamThresholdSettings field to given value.

### HasUseIpamThresholdSettings

`func (o *GetNetworkcontainerResponse) HasUseIpamThresholdSettings() bool`

HasUseIpamThresholdSettings returns a boolean if a field has been set.

### GetUseIpamTrapSettings

`func (o *GetNetworkcontainerResponse) GetUseIpamTrapSettings() bool`

GetUseIpamTrapSettings returns the UseIpamTrapSettings field if non-nil, zero value otherwise.

### GetUseIpamTrapSettingsOk

`func (o *GetNetworkcontainerResponse) GetUseIpamTrapSettingsOk() (*bool, bool)`

GetUseIpamTrapSettingsOk returns a tuple with the UseIpamTrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpamTrapSettings

`func (o *GetNetworkcontainerResponse) SetUseIpamTrapSettings(v bool)`

SetUseIpamTrapSettings sets UseIpamTrapSettings field to given value.

### HasUseIpamTrapSettings

`func (o *GetNetworkcontainerResponse) HasUseIpamTrapSettings() bool`

HasUseIpamTrapSettings returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *GetNetworkcontainerResponse) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *GetNetworkcontainerResponse) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *GetNetworkcontainerResponse) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *GetNetworkcontainerResponse) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetNetworkcontainerResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetNetworkcontainerResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetNetworkcontainerResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetNetworkcontainerResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMgmPrivate

`func (o *GetNetworkcontainerResponse) GetUseMgmPrivate() bool`

GetUseMgmPrivate returns the UseMgmPrivate field if non-nil, zero value otherwise.

### GetUseMgmPrivateOk

`func (o *GetNetworkcontainerResponse) GetUseMgmPrivateOk() (*bool, bool)`

GetUseMgmPrivateOk returns a tuple with the UseMgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmPrivate

`func (o *GetNetworkcontainerResponse) SetUseMgmPrivate(v bool)`

SetUseMgmPrivate sets UseMgmPrivate field to given value.

### HasUseMgmPrivate

`func (o *GetNetworkcontainerResponse) HasUseMgmPrivate() bool`

HasUseMgmPrivate returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetNetworkcontainerResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetNetworkcontainerResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetNetworkcontainerResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetNetworkcontainerResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetNetworkcontainerResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetNetworkcontainerResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetNetworkcontainerResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetNetworkcontainerResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetNetworkcontainerResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetNetworkcontainerResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetNetworkcontainerResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetNetworkcontainerResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetNetworkcontainerResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetNetworkcontainerResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetNetworkcontainerResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetNetworkcontainerResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *GetNetworkcontainerResponse) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *GetNetworkcontainerResponse) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *GetNetworkcontainerResponse) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *GetNetworkcontainerResponse) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *GetNetworkcontainerResponse) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *GetNetworkcontainerResponse) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *GetNetworkcontainerResponse) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *GetNetworkcontainerResponse) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseZoneAssociations

`func (o *GetNetworkcontainerResponse) GetUseZoneAssociations() bool`

GetUseZoneAssociations returns the UseZoneAssociations field if non-nil, zero value otherwise.

### GetUseZoneAssociationsOk

`func (o *GetNetworkcontainerResponse) GetUseZoneAssociationsOk() (*bool, bool)`

GetUseZoneAssociationsOk returns a tuple with the UseZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseZoneAssociations

`func (o *GetNetworkcontainerResponse) SetUseZoneAssociations(v bool)`

SetUseZoneAssociations sets UseZoneAssociations field to given value.

### HasUseZoneAssociations

`func (o *GetNetworkcontainerResponse) HasUseZoneAssociations() bool`

HasUseZoneAssociations returns a boolean if a field has been set.

### GetUtilization

`func (o *GetNetworkcontainerResponse) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *GetNetworkcontainerResponse) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *GetNetworkcontainerResponse) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *GetNetworkcontainerResponse) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *GetNetworkcontainerResponse) GetZoneAssociations() []NetworkcontainerZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *GetNetworkcontainerResponse) GetZoneAssociationsOk() (*[]NetworkcontainerZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *GetNetworkcontainerResponse) SetZoneAssociations(v []NetworkcontainerZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *GetNetworkcontainerResponse) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.

### GetResult

`func (o *GetNetworkcontainerResponse) GetResult() Networkcontainer`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNetworkcontainerResponse) GetResultOk() (*Networkcontainer, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNetworkcontainerResponse) SetResult(v Networkcontainer)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNetworkcontainerResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


