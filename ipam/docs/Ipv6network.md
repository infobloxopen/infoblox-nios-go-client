# Ipv6network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AutoCreateReversezone** | Pointer to **bool** | This flag controls whether reverse zones are automatically created when the network is added. | [optional] 
**CloudInfo** | Pointer to [**Ipv6networkCloudInfo**](Ipv6networkCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the network; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network. | [optional] 
**DdnsEnableOptionFqdn** | Pointer to **bool** | Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DeleteReason** | Pointer to **string** | The reason for deleting the RIR registration request. | [optional] 
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
**DiscoveryBasicPollSettings** | Pointer to [**Ipv6networkDiscoveryBasicPollSettings**](Ipv6networkDiscoveryBasicPollSettings.md) |  | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**Ipv6networkDiscoveryBlackoutSetting**](Ipv6networkDiscoveryBlackoutSetting.md) |  | [optional] 
**DiscoveryEngineType** | Pointer to **string** | The network discovery engine type. | [optional] [readonly] 
**DiscoveryMember** | Pointer to **string** | The member that will run discovery for this network. | [optional] 
**DomainName** | Pointer to **string** | Use this method to set or retrieve the domain_name value of a DHCP IPv6 Network object. | [optional] 
**DomainNameServers** | Pointer to **[]string** | Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP IPv6 network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableDiscovery** | Pointer to **bool** | Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery is disabled. | [optional] 
**EnableIfmapPublishing** | Pointer to **bool** | Determines if IFMAP publishing is enabled for the network. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the network should be immediately enabled. | [optional] 
**EndpointSources** | Pointer to **[]string** | The endpoints that provides data for the DHCP IPv6 Network object. | [optional] [readonly] 
**ExpandNetwork** | Pointer to **map[string]interface{}** |  | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LastRirRegistrationUpdateSent** | Pointer to **int64** | The timestamp when the last RIR registration update was sent. | [optional] [readonly] 
**LastRirRegistrationUpdateStatus** | Pointer to **string** | Last RIR registration update status. | [optional] [readonly] 
**LogicFilterRules** | Pointer to [**[]Ipv6networkLogicFilterRules**](Ipv6networkLogicFilterRules.md) | This field contains the logic filters to be applied on this IPv6 network. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Members** | Pointer to [**[]Ipv6networkMembers**](Ipv6networkMembers.md) | A list of members servers that serve DHCP for the network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \&quot;_struct\&quot; member to the struct type. | [optional] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MgmPrivateOverridable** | Pointer to **bool** | This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**Ipv6networkMsAdUserData**](Ipv6networkMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network address in IPv6 Address/CIDR format. For regular expression searches, only the IPv6 Address portion is supported. Searches for the CIDR portion is always an exact match. For example, both network containers 16::0/28 and 26::0/24 are matched by expression &#39;.6&#39; and only 26::0/24 is matched by &#39;.6/24&#39;. | [optional] 
**NetworkContainer** | Pointer to **string** | The network container to which this network belongs, if any. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this network resides. | [optional] 
**NextAvailableIp** | Pointer to **map[string]interface{}** |  | [optional] 
**NextAvailableNetwork** | Pointer to **map[string]interface{}** |  | [optional] 
**NextAvailableVlan** | Pointer to **map[string]interface{}** |  | [optional] 
**Options** | Pointer to [**[]Ipv6networkOptions**](Ipv6networkOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PortControlBlackoutSetting** | Pointer to [**Ipv6networkPortControlBlackoutSetting**](Ipv6networkPortControlBlackoutSetting.md) |  | [optional] 
**PreferredLifetime** | Pointer to **int64** | Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network object. | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**Rir** | Pointer to **string** | The registry (RIR) that allocated the IPv6 network address space. | [optional] [readonly] 
**RirOrganization** | Pointer to **string** | The RIR organization associated with the IPv6 network. | [optional] 
**RirRegistrationAction** | Pointer to **string** | The RIR registration action. | [optional] 
**RirRegistrationStatus** | Pointer to **string** | The registration status of the IPv6 network in RIR. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | If the field is set to True, the discovery blackout setting will be used for port control blackout setting. | [optional] 
**SendRirRequest** | Pointer to **bool** | Determines whether to send the RIR registration request. | [optional] 
**SplitNetwork** | Pointer to **map[string]interface{}** |  | [optional] 
**SubscribeSettings** | Pointer to [**Ipv6networkSubscribeSettings**](Ipv6networkSubscribeSettings.md) |  | [optional] 
**Template** | Pointer to **string** | If set on creation, the network is created according to the values specified in the selected template. | [optional] 
**Unmanaged** | Pointer to **bool** | Determines whether the DHCP IPv6 Network is unmanaged or not. | [optional] 
**UnmanagedCount** | Pointer to **int64** | The number of unmanaged IP addresses as discovered by network discovery. | [optional] [readonly] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseBlackoutSetting** | Pointer to **bool** | Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsEnableOptionFqdn** | Pointer to **bool** | Use flag for: ddns_enable_option_fqdn | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDiscoveryBasicPollingSettings** | Pointer to **bool** | Use flag for: discovery_basic_poll_settings | [optional] 
**UseDomainName** | Pointer to **bool** | Use flag for: domain_name | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseEnableDiscovery** | Pointer to **bool** | Use flag for: discovery_member , enable_discovery | [optional] 
**UseEnableIfmapPublishing** | Pointer to **bool** | Use flag for: enable_ifmap_publishing | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMgmPrivate** | Pointer to **bool** | Use flag for: mgm_private | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseSubscribeSettings** | Pointer to **bool** | Use flag for: subscribe_settings | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**UseZoneAssociations** | Pointer to **bool** | Use flag for: zone_associations | [optional] 
**ValidLifetime** | Pointer to **int64** | Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network object. | [optional] 
**Vlans** | Pointer to [**[]Ipv6networkVlans**](Ipv6networkVlans.md) | List of VLANs assigned to Network. | [optional] 
**ZoneAssociations** | Pointer to [**[]Ipv6networkZoneAssociations**](Ipv6networkZoneAssociations.md) | The list of zones associated with this network. | [optional] 

## Methods

### NewIpv6network

`func NewIpv6network() *Ipv6network`

NewIpv6network instantiates a new Ipv6network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6networkWithDefaults

`func NewIpv6networkWithDefaults() *Ipv6network`

NewIpv6networkWithDefaults instantiates a new Ipv6network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6network) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6network) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6network) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6network) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *Ipv6network) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *Ipv6network) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *Ipv6network) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *Ipv6network) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Ipv6network) GetCloudInfo() Ipv6networkCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Ipv6network) GetCloudInfoOk() (*Ipv6networkCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Ipv6network) SetCloudInfo(v Ipv6networkCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Ipv6network) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6network) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6network) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6network) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6network) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Ipv6network) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Ipv6network) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Ipv6network) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Ipv6network) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsEnableOptionFqdn

`func (o *Ipv6network) GetDdnsEnableOptionFqdn() bool`

GetDdnsEnableOptionFqdn returns the DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetDdnsEnableOptionFqdnOk

`func (o *Ipv6network) GetDdnsEnableOptionFqdnOk() (*bool, bool)`

GetDdnsEnableOptionFqdnOk returns a tuple with the DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnableOptionFqdn

`func (o *Ipv6network) SetDdnsEnableOptionFqdn(v bool)`

SetDdnsEnableOptionFqdn sets DdnsEnableOptionFqdn field to given value.

### HasDdnsEnableOptionFqdn

`func (o *Ipv6network) HasDdnsEnableOptionFqdn() bool`

HasDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *Ipv6network) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *Ipv6network) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *Ipv6network) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *Ipv6network) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *Ipv6network) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *Ipv6network) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *Ipv6network) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *Ipv6network) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *Ipv6network) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *Ipv6network) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *Ipv6network) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *Ipv6network) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDeleteReason

`func (o *Ipv6network) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *Ipv6network) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *Ipv6network) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *Ipv6network) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### GetDisable

`func (o *Ipv6network) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Ipv6network) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Ipv6network) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Ipv6network) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Ipv6network) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Ipv6network) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Ipv6network) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Ipv6network) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredBgpAs

`func (o *Ipv6network) GetDiscoveredBgpAs() string`

GetDiscoveredBgpAs returns the DiscoveredBgpAs field if non-nil, zero value otherwise.

### GetDiscoveredBgpAsOk

`func (o *Ipv6network) GetDiscoveredBgpAsOk() (*string, bool)`

GetDiscoveredBgpAsOk returns a tuple with the DiscoveredBgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBgpAs

`func (o *Ipv6network) SetDiscoveredBgpAs(v string)`

SetDiscoveredBgpAs sets DiscoveredBgpAs field to given value.

### HasDiscoveredBgpAs

`func (o *Ipv6network) HasDiscoveredBgpAs() bool`

HasDiscoveredBgpAs returns a boolean if a field has been set.

### GetDiscoveredBridgeDomain

`func (o *Ipv6network) GetDiscoveredBridgeDomain() string`

GetDiscoveredBridgeDomain returns the DiscoveredBridgeDomain field if non-nil, zero value otherwise.

### GetDiscoveredBridgeDomainOk

`func (o *Ipv6network) GetDiscoveredBridgeDomainOk() (*string, bool)`

GetDiscoveredBridgeDomainOk returns a tuple with the DiscoveredBridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBridgeDomain

`func (o *Ipv6network) SetDiscoveredBridgeDomain(v string)`

SetDiscoveredBridgeDomain sets DiscoveredBridgeDomain field to given value.

### HasDiscoveredBridgeDomain

`func (o *Ipv6network) HasDiscoveredBridgeDomain() bool`

HasDiscoveredBridgeDomain returns a boolean if a field has been set.

### GetDiscoveredTenant

`func (o *Ipv6network) GetDiscoveredTenant() string`

GetDiscoveredTenant returns the DiscoveredTenant field if non-nil, zero value otherwise.

### GetDiscoveredTenantOk

`func (o *Ipv6network) GetDiscoveredTenantOk() (*string, bool)`

GetDiscoveredTenantOk returns a tuple with the DiscoveredTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTenant

`func (o *Ipv6network) SetDiscoveredTenant(v string)`

SetDiscoveredTenant sets DiscoveredTenant field to given value.

### HasDiscoveredTenant

`func (o *Ipv6network) HasDiscoveredTenant() bool`

HasDiscoveredTenant returns a boolean if a field has been set.

### GetDiscoveredVlanId

`func (o *Ipv6network) GetDiscoveredVlanId() string`

GetDiscoveredVlanId returns the DiscoveredVlanId field if non-nil, zero value otherwise.

### GetDiscoveredVlanIdOk

`func (o *Ipv6network) GetDiscoveredVlanIdOk() (*string, bool)`

GetDiscoveredVlanIdOk returns a tuple with the DiscoveredVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVlanId

`func (o *Ipv6network) SetDiscoveredVlanId(v string)`

SetDiscoveredVlanId sets DiscoveredVlanId field to given value.

### HasDiscoveredVlanId

`func (o *Ipv6network) HasDiscoveredVlanId() bool`

HasDiscoveredVlanId returns a boolean if a field has been set.

### GetDiscoveredVlanName

`func (o *Ipv6network) GetDiscoveredVlanName() string`

GetDiscoveredVlanName returns the DiscoveredVlanName field if non-nil, zero value otherwise.

### GetDiscoveredVlanNameOk

`func (o *Ipv6network) GetDiscoveredVlanNameOk() (*string, bool)`

GetDiscoveredVlanNameOk returns a tuple with the DiscoveredVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVlanName

`func (o *Ipv6network) SetDiscoveredVlanName(v string)`

SetDiscoveredVlanName sets DiscoveredVlanName field to given value.

### HasDiscoveredVlanName

`func (o *Ipv6network) HasDiscoveredVlanName() bool`

HasDiscoveredVlanName returns a boolean if a field has been set.

### GetDiscoveredVrfDescription

`func (o *Ipv6network) GetDiscoveredVrfDescription() string`

GetDiscoveredVrfDescription returns the DiscoveredVrfDescription field if non-nil, zero value otherwise.

### GetDiscoveredVrfDescriptionOk

`func (o *Ipv6network) GetDiscoveredVrfDescriptionOk() (*string, bool)`

GetDiscoveredVrfDescriptionOk returns a tuple with the DiscoveredVrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfDescription

`func (o *Ipv6network) SetDiscoveredVrfDescription(v string)`

SetDiscoveredVrfDescription sets DiscoveredVrfDescription field to given value.

### HasDiscoveredVrfDescription

`func (o *Ipv6network) HasDiscoveredVrfDescription() bool`

HasDiscoveredVrfDescription returns a boolean if a field has been set.

### GetDiscoveredVrfName

`func (o *Ipv6network) GetDiscoveredVrfName() string`

GetDiscoveredVrfName returns the DiscoveredVrfName field if non-nil, zero value otherwise.

### GetDiscoveredVrfNameOk

`func (o *Ipv6network) GetDiscoveredVrfNameOk() (*string, bool)`

GetDiscoveredVrfNameOk returns a tuple with the DiscoveredVrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfName

`func (o *Ipv6network) SetDiscoveredVrfName(v string)`

SetDiscoveredVrfName sets DiscoveredVrfName field to given value.

### HasDiscoveredVrfName

`func (o *Ipv6network) HasDiscoveredVrfName() bool`

HasDiscoveredVrfName returns a boolean if a field has been set.

### GetDiscoveredVrfRd

`func (o *Ipv6network) GetDiscoveredVrfRd() string`

GetDiscoveredVrfRd returns the DiscoveredVrfRd field if non-nil, zero value otherwise.

### GetDiscoveredVrfRdOk

`func (o *Ipv6network) GetDiscoveredVrfRdOk() (*string, bool)`

GetDiscoveredVrfRdOk returns a tuple with the DiscoveredVrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVrfRd

`func (o *Ipv6network) SetDiscoveredVrfRd(v string)`

SetDiscoveredVrfRd sets DiscoveredVrfRd field to given value.

### HasDiscoveredVrfRd

`func (o *Ipv6network) HasDiscoveredVrfRd() bool`

HasDiscoveredVrfRd returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *Ipv6network) GetDiscoveryBasicPollSettings() Ipv6networkDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *Ipv6network) GetDiscoveryBasicPollSettingsOk() (*Ipv6networkDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *Ipv6network) SetDiscoveryBasicPollSettings(v Ipv6networkDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *Ipv6network) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *Ipv6network) GetDiscoveryBlackoutSetting() Ipv6networkDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *Ipv6network) GetDiscoveryBlackoutSettingOk() (*Ipv6networkDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *Ipv6network) SetDiscoveryBlackoutSetting(v Ipv6networkDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *Ipv6network) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryEngineType

`func (o *Ipv6network) GetDiscoveryEngineType() string`

GetDiscoveryEngineType returns the DiscoveryEngineType field if non-nil, zero value otherwise.

### GetDiscoveryEngineTypeOk

`func (o *Ipv6network) GetDiscoveryEngineTypeOk() (*string, bool)`

GetDiscoveryEngineTypeOk returns a tuple with the DiscoveryEngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngineType

`func (o *Ipv6network) SetDiscoveryEngineType(v string)`

SetDiscoveryEngineType sets DiscoveryEngineType field to given value.

### HasDiscoveryEngineType

`func (o *Ipv6network) HasDiscoveryEngineType() bool`

HasDiscoveryEngineType returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *Ipv6network) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *Ipv6network) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *Ipv6network) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *Ipv6network) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetDomainName

`func (o *Ipv6network) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Ipv6network) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Ipv6network) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Ipv6network) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *Ipv6network) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *Ipv6network) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *Ipv6network) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *Ipv6network) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Ipv6network) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Ipv6network) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Ipv6network) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Ipv6network) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *Ipv6network) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *Ipv6network) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *Ipv6network) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *Ipv6network) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableIfmapPublishing

`func (o *Ipv6network) GetEnableIfmapPublishing() bool`

GetEnableIfmapPublishing returns the EnableIfmapPublishing field if non-nil, zero value otherwise.

### GetEnableIfmapPublishingOk

`func (o *Ipv6network) GetEnableIfmapPublishingOk() (*bool, bool)`

GetEnableIfmapPublishingOk returns a tuple with the EnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIfmapPublishing

`func (o *Ipv6network) SetEnableIfmapPublishing(v bool)`

SetEnableIfmapPublishing sets EnableIfmapPublishing field to given value.

### HasEnableIfmapPublishing

`func (o *Ipv6network) HasEnableIfmapPublishing() bool`

HasEnableIfmapPublishing returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Ipv6network) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Ipv6network) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Ipv6network) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Ipv6network) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEndpointSources

`func (o *Ipv6network) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *Ipv6network) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *Ipv6network) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *Ipv6network) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExpandNetwork

`func (o *Ipv6network) GetExpandNetwork() map[string]interface{}`

GetExpandNetwork returns the ExpandNetwork field if non-nil, zero value otherwise.

### GetExpandNetworkOk

`func (o *Ipv6network) GetExpandNetworkOk() (*map[string]interface{}, bool)`

GetExpandNetworkOk returns a tuple with the ExpandNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandNetwork

`func (o *Ipv6network) SetExpandNetwork(v map[string]interface{})`

SetExpandNetwork sets ExpandNetwork field to given value.

### HasExpandNetwork

`func (o *Ipv6network) HasExpandNetwork() bool`

HasExpandNetwork returns a boolean if a field has been set.

### GetExtattrs

`func (o *Ipv6network) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Ipv6network) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Ipv6network) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Ipv6network) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateSent

`func (o *Ipv6network) GetLastRirRegistrationUpdateSent() int64`

GetLastRirRegistrationUpdateSent returns the LastRirRegistrationUpdateSent field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateSentOk

`func (o *Ipv6network) GetLastRirRegistrationUpdateSentOk() (*int64, bool)`

GetLastRirRegistrationUpdateSentOk returns a tuple with the LastRirRegistrationUpdateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateSent

`func (o *Ipv6network) SetLastRirRegistrationUpdateSent(v int64)`

SetLastRirRegistrationUpdateSent sets LastRirRegistrationUpdateSent field to given value.

### HasLastRirRegistrationUpdateSent

`func (o *Ipv6network) HasLastRirRegistrationUpdateSent() bool`

HasLastRirRegistrationUpdateSent returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateStatus

`func (o *Ipv6network) GetLastRirRegistrationUpdateStatus() string`

GetLastRirRegistrationUpdateStatus returns the LastRirRegistrationUpdateStatus field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateStatusOk

`func (o *Ipv6network) GetLastRirRegistrationUpdateStatusOk() (*string, bool)`

GetLastRirRegistrationUpdateStatusOk returns a tuple with the LastRirRegistrationUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateStatus

`func (o *Ipv6network) SetLastRirRegistrationUpdateStatus(v string)`

SetLastRirRegistrationUpdateStatus sets LastRirRegistrationUpdateStatus field to given value.

### HasLastRirRegistrationUpdateStatus

`func (o *Ipv6network) HasLastRirRegistrationUpdateStatus() bool`

HasLastRirRegistrationUpdateStatus returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6network) GetLogicFilterRules() []Ipv6networkLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6network) GetLogicFilterRulesOk() (*[]Ipv6networkLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6network) SetLogicFilterRules(v []Ipv6networkLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6network) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMembers

`func (o *Ipv6network) GetMembers() []Ipv6networkMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Ipv6network) GetMembersOk() (*[]Ipv6networkMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Ipv6network) SetMembers(v []Ipv6networkMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Ipv6network) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *Ipv6network) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *Ipv6network) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *Ipv6network) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *Ipv6network) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *Ipv6network) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *Ipv6network) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *Ipv6network) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *Ipv6network) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Ipv6network) GetMsAdUserData() Ipv6networkMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Ipv6network) GetMsAdUserDataOk() (*Ipv6networkMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Ipv6network) SetMsAdUserData(v Ipv6networkMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Ipv6network) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *Ipv6network) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ipv6network) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ipv6network) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ipv6network) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkContainer

`func (o *Ipv6network) GetNetworkContainer() string`

GetNetworkContainer returns the NetworkContainer field if non-nil, zero value otherwise.

### GetNetworkContainerOk

`func (o *Ipv6network) GetNetworkContainerOk() (*string, bool)`

GetNetworkContainerOk returns a tuple with the NetworkContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkContainer

`func (o *Ipv6network) SetNetworkContainer(v string)`

SetNetworkContainer sets NetworkContainer field to given value.

### HasNetworkContainer

`func (o *Ipv6network) HasNetworkContainer() bool`

HasNetworkContainer returns a boolean if a field has been set.

### GetNetworkView

`func (o *Ipv6network) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Ipv6network) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Ipv6network) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Ipv6network) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextAvailableIp

`func (o *Ipv6network) GetNextAvailableIp() map[string]interface{}`

GetNextAvailableIp returns the NextAvailableIp field if non-nil, zero value otherwise.

### GetNextAvailableIpOk

`func (o *Ipv6network) GetNextAvailableIpOk() (*map[string]interface{}, bool)`

GetNextAvailableIpOk returns a tuple with the NextAvailableIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableIp

`func (o *Ipv6network) SetNextAvailableIp(v map[string]interface{})`

SetNextAvailableIp sets NextAvailableIp field to given value.

### HasNextAvailableIp

`func (o *Ipv6network) HasNextAvailableIp() bool`

HasNextAvailableIp returns a boolean if a field has been set.

### GetNextAvailableNetwork

`func (o *Ipv6network) GetNextAvailableNetwork() map[string]interface{}`

GetNextAvailableNetwork returns the NextAvailableNetwork field if non-nil, zero value otherwise.

### GetNextAvailableNetworkOk

`func (o *Ipv6network) GetNextAvailableNetworkOk() (*map[string]interface{}, bool)`

GetNextAvailableNetworkOk returns a tuple with the NextAvailableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableNetwork

`func (o *Ipv6network) SetNextAvailableNetwork(v map[string]interface{})`

SetNextAvailableNetwork sets NextAvailableNetwork field to given value.

### HasNextAvailableNetwork

`func (o *Ipv6network) HasNextAvailableNetwork() bool`

HasNextAvailableNetwork returns a boolean if a field has been set.

### GetNextAvailableVlan

`func (o *Ipv6network) GetNextAvailableVlan() map[string]interface{}`

GetNextAvailableVlan returns the NextAvailableVlan field if non-nil, zero value otherwise.

### GetNextAvailableVlanOk

`func (o *Ipv6network) GetNextAvailableVlanOk() (*map[string]interface{}, bool)`

GetNextAvailableVlanOk returns a tuple with the NextAvailableVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableVlan

`func (o *Ipv6network) SetNextAvailableVlan(v map[string]interface{})`

SetNextAvailableVlan sets NextAvailableVlan field to given value.

### HasNextAvailableVlan

`func (o *Ipv6network) HasNextAvailableVlan() bool`

HasNextAvailableVlan returns a boolean if a field has been set.

### GetOptions

`func (o *Ipv6network) GetOptions() []Ipv6networkOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Ipv6network) GetOptionsOk() (*[]Ipv6networkOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Ipv6network) SetOptions(v []Ipv6networkOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Ipv6network) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *Ipv6network) GetPortControlBlackoutSetting() Ipv6networkPortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *Ipv6network) GetPortControlBlackoutSettingOk() (*Ipv6networkPortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *Ipv6network) SetPortControlBlackoutSetting(v Ipv6networkPortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *Ipv6network) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *Ipv6network) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *Ipv6network) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *Ipv6network) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *Ipv6network) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Ipv6network) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Ipv6network) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Ipv6network) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Ipv6network) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Ipv6network) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Ipv6network) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Ipv6network) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Ipv6network) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRir

`func (o *Ipv6network) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *Ipv6network) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *Ipv6network) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *Ipv6network) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *Ipv6network) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *Ipv6network) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *Ipv6network) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *Ipv6network) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *Ipv6network) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *Ipv6network) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *Ipv6network) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *Ipv6network) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *Ipv6network) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *Ipv6network) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *Ipv6network) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *Ipv6network) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *Ipv6network) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *Ipv6network) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *Ipv6network) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *Ipv6network) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *Ipv6network) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *Ipv6network) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *Ipv6network) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *Ipv6network) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetSplitNetwork

`func (o *Ipv6network) GetSplitNetwork() map[string]interface{}`

GetSplitNetwork returns the SplitNetwork field if non-nil, zero value otherwise.

### GetSplitNetworkOk

`func (o *Ipv6network) GetSplitNetworkOk() (*map[string]interface{}, bool)`

GetSplitNetworkOk returns a tuple with the SplitNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitNetwork

`func (o *Ipv6network) SetSplitNetwork(v map[string]interface{})`

SetSplitNetwork sets SplitNetwork field to given value.

### HasSplitNetwork

`func (o *Ipv6network) HasSplitNetwork() bool`

HasSplitNetwork returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *Ipv6network) GetSubscribeSettings() Ipv6networkSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *Ipv6network) GetSubscribeSettingsOk() (*Ipv6networkSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *Ipv6network) SetSubscribeSettings(v Ipv6networkSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *Ipv6network) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *Ipv6network) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Ipv6network) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Ipv6network) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Ipv6network) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUnmanaged

`func (o *Ipv6network) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *Ipv6network) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *Ipv6network) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *Ipv6network) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetUnmanagedCount

`func (o *Ipv6network) GetUnmanagedCount() int64`

GetUnmanagedCount returns the UnmanagedCount field if non-nil, zero value otherwise.

### GetUnmanagedCountOk

`func (o *Ipv6network) GetUnmanagedCountOk() (*int64, bool)`

GetUnmanagedCountOk returns a tuple with the UnmanagedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedCount

`func (o *Ipv6network) SetUnmanagedCount(v int64)`

SetUnmanagedCount sets UnmanagedCount field to given value.

### HasUnmanagedCount

`func (o *Ipv6network) HasUnmanagedCount() bool`

HasUnmanagedCount returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *Ipv6network) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *Ipv6network) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *Ipv6network) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *Ipv6network) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *Ipv6network) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *Ipv6network) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *Ipv6network) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *Ipv6network) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Ipv6network) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Ipv6network) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Ipv6network) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Ipv6network) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsEnableOptionFqdn

`func (o *Ipv6network) GetUseDdnsEnableOptionFqdn() bool`

GetUseDdnsEnableOptionFqdn returns the UseDdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetUseDdnsEnableOptionFqdnOk

`func (o *Ipv6network) GetUseDdnsEnableOptionFqdnOk() (*bool, bool)`

GetUseDdnsEnableOptionFqdnOk returns a tuple with the UseDdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsEnableOptionFqdn

`func (o *Ipv6network) SetUseDdnsEnableOptionFqdn(v bool)`

SetUseDdnsEnableOptionFqdn sets UseDdnsEnableOptionFqdn field to given value.

### HasUseDdnsEnableOptionFqdn

`func (o *Ipv6network) HasUseDdnsEnableOptionFqdn() bool`

HasUseDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *Ipv6network) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *Ipv6network) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *Ipv6network) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *Ipv6network) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *Ipv6network) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *Ipv6network) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *Ipv6network) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *Ipv6network) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *Ipv6network) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *Ipv6network) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *Ipv6network) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *Ipv6network) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseDomainName

`func (o *Ipv6network) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *Ipv6network) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *Ipv6network) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *Ipv6network) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *Ipv6network) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *Ipv6network) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *Ipv6network) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *Ipv6network) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Ipv6network) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Ipv6network) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Ipv6network) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Ipv6network) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *Ipv6network) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *Ipv6network) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *Ipv6network) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *Ipv6network) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseEnableIfmapPublishing

`func (o *Ipv6network) GetUseEnableIfmapPublishing() bool`

GetUseEnableIfmapPublishing returns the UseEnableIfmapPublishing field if non-nil, zero value otherwise.

### GetUseEnableIfmapPublishingOk

`func (o *Ipv6network) GetUseEnableIfmapPublishingOk() (*bool, bool)`

GetUseEnableIfmapPublishingOk returns a tuple with the UseEnableIfmapPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableIfmapPublishing

`func (o *Ipv6network) SetUseEnableIfmapPublishing(v bool)`

SetUseEnableIfmapPublishing sets UseEnableIfmapPublishing field to given value.

### HasUseEnableIfmapPublishing

`func (o *Ipv6network) HasUseEnableIfmapPublishing() bool`

HasUseEnableIfmapPublishing returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6network) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6network) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6network) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6network) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMgmPrivate

`func (o *Ipv6network) GetUseMgmPrivate() bool`

GetUseMgmPrivate returns the UseMgmPrivate field if non-nil, zero value otherwise.

### GetUseMgmPrivateOk

`func (o *Ipv6network) GetUseMgmPrivateOk() (*bool, bool)`

GetUseMgmPrivateOk returns a tuple with the UseMgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmPrivate

`func (o *Ipv6network) SetUseMgmPrivate(v bool)`

SetUseMgmPrivate sets UseMgmPrivate field to given value.

### HasUseMgmPrivate

`func (o *Ipv6network) HasUseMgmPrivate() bool`

HasUseMgmPrivate returns a boolean if a field has been set.

### GetUseOptions

`func (o *Ipv6network) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Ipv6network) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Ipv6network) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Ipv6network) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *Ipv6network) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *Ipv6network) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *Ipv6network) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *Ipv6network) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Ipv6network) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Ipv6network) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Ipv6network) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Ipv6network) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *Ipv6network) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *Ipv6network) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *Ipv6network) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *Ipv6network) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6network) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *Ipv6network) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6network) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6network) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *Ipv6network) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *Ipv6network) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *Ipv6network) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *Ipv6network) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetUseZoneAssociations

`func (o *Ipv6network) GetUseZoneAssociations() bool`

GetUseZoneAssociations returns the UseZoneAssociations field if non-nil, zero value otherwise.

### GetUseZoneAssociationsOk

`func (o *Ipv6network) GetUseZoneAssociationsOk() (*bool, bool)`

GetUseZoneAssociationsOk returns a tuple with the UseZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseZoneAssociations

`func (o *Ipv6network) SetUseZoneAssociations(v bool)`

SetUseZoneAssociations sets UseZoneAssociations field to given value.

### HasUseZoneAssociations

`func (o *Ipv6network) HasUseZoneAssociations() bool`

HasUseZoneAssociations returns a boolean if a field has been set.

### GetValidLifetime

`func (o *Ipv6network) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *Ipv6network) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *Ipv6network) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *Ipv6network) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetVlans

`func (o *Ipv6network) GetVlans() []Ipv6networkVlans`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *Ipv6network) GetVlansOk() (*[]Ipv6networkVlans, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *Ipv6network) SetVlans(v []Ipv6networkVlans)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *Ipv6network) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *Ipv6network) GetZoneAssociations() []Ipv6networkZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *Ipv6network) GetZoneAssociationsOk() (*[]Ipv6networkZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *Ipv6network) SetZoneAssociations(v []Ipv6networkZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *Ipv6network) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


