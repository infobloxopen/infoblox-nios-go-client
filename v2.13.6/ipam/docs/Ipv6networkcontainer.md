# Ipv6networkcontainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AutoCreateReversezone** | Pointer to **bool** | This flag controls whether reverse zones are automatically created when the network is added. | [optional] 
**CloudInfo** | Pointer to [**Ipv6networkcontainerCloudInfo**](Ipv6networkcontainerCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the network; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container. | [optional] 
**DdnsEnableOptionFqdn** | Pointer to **bool** | Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network Container object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DeleteReason** | Pointer to **string** | The reason for deleting the RIR registration request. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this network container. | [optional] [readonly] 
**DiscoveryBasicPollSettings** | Pointer to [**Ipv6networkcontainerDiscoveryBasicPollSettings**](Ipv6networkcontainerDiscoveryBasicPollSettings.md) |  | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**Ipv6networkcontainerDiscoveryBlackoutSetting**](Ipv6networkcontainerDiscoveryBlackoutSetting.md) |  | [optional] 
**DiscoveryEngineType** | Pointer to **string** | The network discovery engine type. | [optional] [readonly] 
**DiscoveryMember** | Pointer to **string** | The member that will run discovery for this network container. | [optional] 
**DomainNameServers** | Pointer to **[]string** | Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network Container object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings. | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP IPv6 network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableDiscovery** | Pointer to **bool** | Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the network container should be immediately enabled. | [optional] 
**EndpointSources** | Pointer to **[]string** | The endpoints that provides data for the DHCP IPv6 Network Container. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FederatedRealms** | Pointer to [**[]Ipv6networkcontainerFederatedRealms**](Ipv6networkcontainerFederatedRealms.md) | This field contains the federated realms associated to this network container. | [optional] 
**LastRirRegistrationUpdateSent** | Pointer to **int64** | The timestamp when the last RIR registration update was sent. | [optional] [readonly] 
**LastRirRegistrationUpdateStatus** | Pointer to **string** | Last RIR registration update status. | [optional] [readonly] 
**LogicFilterRules** | Pointer to [**[]Ipv6networkcontainerLogicFilterRules**](Ipv6networkcontainerLogicFilterRules.md) | This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MgmPrivateOverridable** | Pointer to **bool** | This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**Ipv6networkcontainerMsAdUserData**](Ipv6networkcontainerMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network address in IPv6 Address/CIDR format. For regular expression searches, only the IPv6 Address portion is supported. Searches for the CIDR portion is always an exact match. For example, both network containers 16::0/28 and 26::0/24 are matched by expression &#39;.6&#39; and only 26::0/24 is matched by &#39;.6/24&#39;. | [optional] 
**NetworkContainer** | Pointer to **string** | The network container to which this network belongs, if any. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this network resides. | [optional] 
**Options** | Pointer to [**[]Ipv6networkcontainerOptions**](Ipv6networkcontainerOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PortControlBlackoutSetting** | Pointer to [**Ipv6networkcontainerPortControlBlackoutSetting**](Ipv6networkcontainerPortControlBlackoutSetting.md) |  | [optional] 
**PreferredLifetime** | Pointer to **int64** | Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network Container object. | [optional] 
**RemoveSubnets** | Pointer to **bool** | Remove subnets delete option. Determines whether all child objects should be removed alongside with the IPv6 network container or child objects should be assigned to another parental container. By default child objects are deleted with this network container. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**Rir** | Pointer to **string** | The registry (RIR) that allocated the IPv6 network container address space. | [optional] [readonly] 
**RirOrganization** | Pointer to **string** | The RIR organization associated with the IPv6 network container. | [optional] 
**RirRegistrationAction** | Pointer to **string** | The RIR registration action. | [optional] 
**RirRegistrationStatus** | Pointer to **string** | The registration status of the IPv6 network container in RIR. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | If the field is set to True, the discovery blackout setting will be used for port control blackout setting. | [optional] 
**SendRirRequest** | Pointer to **bool** | Determines whether to send the RIR registration request. | [optional] 
**SubscribeSettings** | Pointer to [**Ipv6networkcontainerSubscribeSettings**](Ipv6networkcontainerSubscribeSettings.md) |  | [optional] 
**Unmanaged** | Pointer to **bool** | Determines whether the network container is unmanaged or not. | [optional] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseBlackoutSetting** | Pointer to **bool** | Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsEnableOptionFqdn** | Pointer to **bool** | Use flag for: ddns_enable_option_fqdn | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDiscoveryBasicPollingSettings** | Pointer to **bool** | Use flag for: discovery_basic_poll_settings | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseEnableDiscovery** | Pointer to **bool** | Use flag for: discovery_member , enable_discovery | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMgmPrivate** | Pointer to **bool** | Use flag for: mgm_private | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseSubscribeSettings** | Pointer to **bool** | Use flag for: subscribe_settings | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**UseZoneAssociations** | Pointer to **bool** | Use flag for: zone_associations | [optional] 
**Utilization** | Pointer to **int64** | The network container utilization in percentage. | [optional] [readonly] 
**ValidLifetime** | Pointer to **int64** | Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network Container object. | [optional] 
**ZoneAssociations** | Pointer to [**[]Ipv6networkcontainerZoneAssociations**](Ipv6networkcontainerZoneAssociations.md) | The list of zones associated with this network container. | [optional] 

## Methods

### NewIpv6networkcontainer

`func NewIpv6networkcontainer() *Ipv6networkcontainer`

NewIpv6networkcontainer instantiates a new Ipv6networkcontainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6networkcontainerWithDefaults

`func NewIpv6networkcontainerWithDefaults() *Ipv6networkcontainer`

NewIpv6networkcontainerWithDefaults instantiates a new Ipv6networkcontainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6networkcontainer) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6networkcontainer) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6networkcontainer) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6networkcontainer) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreateReversezone

`func (o *Ipv6networkcontainer) GetAutoCreateReversezone() bool`

GetAutoCreateReversezone returns the AutoCreateReversezone field if non-nil, zero value otherwise.

### GetAutoCreateReversezoneOk

`func (o *Ipv6networkcontainer) GetAutoCreateReversezoneOk() (*bool, bool)`

GetAutoCreateReversezoneOk returns a tuple with the AutoCreateReversezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateReversezone

`func (o *Ipv6networkcontainer) SetAutoCreateReversezone(v bool)`

SetAutoCreateReversezone sets AutoCreateReversezone field to given value.

### HasAutoCreateReversezone

`func (o *Ipv6networkcontainer) HasAutoCreateReversezone() bool`

HasAutoCreateReversezone returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Ipv6networkcontainer) GetCloudInfo() Ipv6networkcontainerCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Ipv6networkcontainer) GetCloudInfoOk() (*Ipv6networkcontainerCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Ipv6networkcontainer) SetCloudInfo(v Ipv6networkcontainerCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Ipv6networkcontainer) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6networkcontainer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6networkcontainer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6networkcontainer) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6networkcontainer) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Ipv6networkcontainer) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Ipv6networkcontainer) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Ipv6networkcontainer) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Ipv6networkcontainer) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsEnableOptionFqdn

`func (o *Ipv6networkcontainer) GetDdnsEnableOptionFqdn() bool`

GetDdnsEnableOptionFqdn returns the DdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetDdnsEnableOptionFqdnOk

`func (o *Ipv6networkcontainer) GetDdnsEnableOptionFqdnOk() (*bool, bool)`

GetDdnsEnableOptionFqdnOk returns a tuple with the DdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnableOptionFqdn

`func (o *Ipv6networkcontainer) SetDdnsEnableOptionFqdn(v bool)`

SetDdnsEnableOptionFqdn sets DdnsEnableOptionFqdn field to given value.

### HasDdnsEnableOptionFqdn

`func (o *Ipv6networkcontainer) HasDdnsEnableOptionFqdn() bool`

HasDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *Ipv6networkcontainer) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *Ipv6networkcontainer) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *Ipv6networkcontainer) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *Ipv6networkcontainer) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *Ipv6networkcontainer) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *Ipv6networkcontainer) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *Ipv6networkcontainer) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *Ipv6networkcontainer) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *Ipv6networkcontainer) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *Ipv6networkcontainer) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *Ipv6networkcontainer) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *Ipv6networkcontainer) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDeleteReason

`func (o *Ipv6networkcontainer) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *Ipv6networkcontainer) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *Ipv6networkcontainer) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *Ipv6networkcontainer) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Ipv6networkcontainer) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Ipv6networkcontainer) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Ipv6networkcontainer) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Ipv6networkcontainer) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *Ipv6networkcontainer) GetDiscoveryBasicPollSettings() Ipv6networkcontainerDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *Ipv6networkcontainer) GetDiscoveryBasicPollSettingsOk() (*Ipv6networkcontainerDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *Ipv6networkcontainer) SetDiscoveryBasicPollSettings(v Ipv6networkcontainerDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *Ipv6networkcontainer) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *Ipv6networkcontainer) GetDiscoveryBlackoutSetting() Ipv6networkcontainerDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *Ipv6networkcontainer) GetDiscoveryBlackoutSettingOk() (*Ipv6networkcontainerDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *Ipv6networkcontainer) SetDiscoveryBlackoutSetting(v Ipv6networkcontainerDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *Ipv6networkcontainer) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryEngineType

`func (o *Ipv6networkcontainer) GetDiscoveryEngineType() string`

GetDiscoveryEngineType returns the DiscoveryEngineType field if non-nil, zero value otherwise.

### GetDiscoveryEngineTypeOk

`func (o *Ipv6networkcontainer) GetDiscoveryEngineTypeOk() (*string, bool)`

GetDiscoveryEngineTypeOk returns a tuple with the DiscoveryEngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngineType

`func (o *Ipv6networkcontainer) SetDiscoveryEngineType(v string)`

SetDiscoveryEngineType sets DiscoveryEngineType field to given value.

### HasDiscoveryEngineType

`func (o *Ipv6networkcontainer) HasDiscoveryEngineType() bool`

HasDiscoveryEngineType returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *Ipv6networkcontainer) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *Ipv6networkcontainer) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *Ipv6networkcontainer) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *Ipv6networkcontainer) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *Ipv6networkcontainer) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *Ipv6networkcontainer) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *Ipv6networkcontainer) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *Ipv6networkcontainer) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Ipv6networkcontainer) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Ipv6networkcontainer) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Ipv6networkcontainer) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Ipv6networkcontainer) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *Ipv6networkcontainer) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *Ipv6networkcontainer) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *Ipv6networkcontainer) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *Ipv6networkcontainer) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Ipv6networkcontainer) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Ipv6networkcontainer) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Ipv6networkcontainer) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Ipv6networkcontainer) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEndpointSources

`func (o *Ipv6networkcontainer) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *Ipv6networkcontainer) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *Ipv6networkcontainer) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *Ipv6networkcontainer) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Ipv6networkcontainer) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Ipv6networkcontainer) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Ipv6networkcontainer) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Ipv6networkcontainer) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *Ipv6networkcontainer) GetFederatedRealms() []Ipv6networkcontainerFederatedRealms`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *Ipv6networkcontainer) GetFederatedRealmsOk() (*[]Ipv6networkcontainerFederatedRealms, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *Ipv6networkcontainer) SetFederatedRealms(v []Ipv6networkcontainerFederatedRealms)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *Ipv6networkcontainer) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateSent

`func (o *Ipv6networkcontainer) GetLastRirRegistrationUpdateSent() int64`

GetLastRirRegistrationUpdateSent returns the LastRirRegistrationUpdateSent field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateSentOk

`func (o *Ipv6networkcontainer) GetLastRirRegistrationUpdateSentOk() (*int64, bool)`

GetLastRirRegistrationUpdateSentOk returns a tuple with the LastRirRegistrationUpdateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateSent

`func (o *Ipv6networkcontainer) SetLastRirRegistrationUpdateSent(v int64)`

SetLastRirRegistrationUpdateSent sets LastRirRegistrationUpdateSent field to given value.

### HasLastRirRegistrationUpdateSent

`func (o *Ipv6networkcontainer) HasLastRirRegistrationUpdateSent() bool`

HasLastRirRegistrationUpdateSent returns a boolean if a field has been set.

### GetLastRirRegistrationUpdateStatus

`func (o *Ipv6networkcontainer) GetLastRirRegistrationUpdateStatus() string`

GetLastRirRegistrationUpdateStatus returns the LastRirRegistrationUpdateStatus field if non-nil, zero value otherwise.

### GetLastRirRegistrationUpdateStatusOk

`func (o *Ipv6networkcontainer) GetLastRirRegistrationUpdateStatusOk() (*string, bool)`

GetLastRirRegistrationUpdateStatusOk returns a tuple with the LastRirRegistrationUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRirRegistrationUpdateStatus

`func (o *Ipv6networkcontainer) SetLastRirRegistrationUpdateStatus(v string)`

SetLastRirRegistrationUpdateStatus sets LastRirRegistrationUpdateStatus field to given value.

### HasLastRirRegistrationUpdateStatus

`func (o *Ipv6networkcontainer) HasLastRirRegistrationUpdateStatus() bool`

HasLastRirRegistrationUpdateStatus returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6networkcontainer) GetLogicFilterRules() []Ipv6networkcontainerLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6networkcontainer) GetLogicFilterRulesOk() (*[]Ipv6networkcontainerLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6networkcontainer) SetLogicFilterRules(v []Ipv6networkcontainerLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6networkcontainer) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *Ipv6networkcontainer) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *Ipv6networkcontainer) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *Ipv6networkcontainer) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *Ipv6networkcontainer) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *Ipv6networkcontainer) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *Ipv6networkcontainer) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *Ipv6networkcontainer) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *Ipv6networkcontainer) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Ipv6networkcontainer) GetMsAdUserData() Ipv6networkcontainerMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Ipv6networkcontainer) GetMsAdUserDataOk() (*Ipv6networkcontainerMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Ipv6networkcontainer) SetMsAdUserData(v Ipv6networkcontainerMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Ipv6networkcontainer) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *Ipv6networkcontainer) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ipv6networkcontainer) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ipv6networkcontainer) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ipv6networkcontainer) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkContainer

`func (o *Ipv6networkcontainer) GetNetworkContainer() string`

GetNetworkContainer returns the NetworkContainer field if non-nil, zero value otherwise.

### GetNetworkContainerOk

`func (o *Ipv6networkcontainer) GetNetworkContainerOk() (*string, bool)`

GetNetworkContainerOk returns a tuple with the NetworkContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkContainer

`func (o *Ipv6networkcontainer) SetNetworkContainer(v string)`

SetNetworkContainer sets NetworkContainer field to given value.

### HasNetworkContainer

`func (o *Ipv6networkcontainer) HasNetworkContainer() bool`

HasNetworkContainer returns a boolean if a field has been set.

### GetNetworkView

`func (o *Ipv6networkcontainer) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Ipv6networkcontainer) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Ipv6networkcontainer) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Ipv6networkcontainer) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOptions

`func (o *Ipv6networkcontainer) GetOptions() []Ipv6networkcontainerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Ipv6networkcontainer) GetOptionsOk() (*[]Ipv6networkcontainerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Ipv6networkcontainer) SetOptions(v []Ipv6networkcontainerOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Ipv6networkcontainer) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *Ipv6networkcontainer) GetPortControlBlackoutSetting() Ipv6networkcontainerPortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *Ipv6networkcontainer) GetPortControlBlackoutSettingOk() (*Ipv6networkcontainerPortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *Ipv6networkcontainer) SetPortControlBlackoutSetting(v Ipv6networkcontainerPortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *Ipv6networkcontainer) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *Ipv6networkcontainer) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *Ipv6networkcontainer) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *Ipv6networkcontainer) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *Ipv6networkcontainer) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetRemoveSubnets

`func (o *Ipv6networkcontainer) GetRemoveSubnets() bool`

GetRemoveSubnets returns the RemoveSubnets field if non-nil, zero value otherwise.

### GetRemoveSubnetsOk

`func (o *Ipv6networkcontainer) GetRemoveSubnetsOk() (*bool, bool)`

GetRemoveSubnetsOk returns a tuple with the RemoveSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSubnets

`func (o *Ipv6networkcontainer) SetRemoveSubnets(v bool)`

SetRemoveSubnets sets RemoveSubnets field to given value.

### HasRemoveSubnets

`func (o *Ipv6networkcontainer) HasRemoveSubnets() bool`

HasRemoveSubnets returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Ipv6networkcontainer) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Ipv6networkcontainer) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Ipv6networkcontainer) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Ipv6networkcontainer) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRir

`func (o *Ipv6networkcontainer) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *Ipv6networkcontainer) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *Ipv6networkcontainer) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *Ipv6networkcontainer) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRirOrganization

`func (o *Ipv6networkcontainer) GetRirOrganization() string`

GetRirOrganization returns the RirOrganization field if non-nil, zero value otherwise.

### GetRirOrganizationOk

`func (o *Ipv6networkcontainer) GetRirOrganizationOk() (*string, bool)`

GetRirOrganizationOk returns a tuple with the RirOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirOrganization

`func (o *Ipv6networkcontainer) SetRirOrganization(v string)`

SetRirOrganization sets RirOrganization field to given value.

### HasRirOrganization

`func (o *Ipv6networkcontainer) HasRirOrganization() bool`

HasRirOrganization returns a boolean if a field has been set.

### GetRirRegistrationAction

`func (o *Ipv6networkcontainer) GetRirRegistrationAction() string`

GetRirRegistrationAction returns the RirRegistrationAction field if non-nil, zero value otherwise.

### GetRirRegistrationActionOk

`func (o *Ipv6networkcontainer) GetRirRegistrationActionOk() (*string, bool)`

GetRirRegistrationActionOk returns a tuple with the RirRegistrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationAction

`func (o *Ipv6networkcontainer) SetRirRegistrationAction(v string)`

SetRirRegistrationAction sets RirRegistrationAction field to given value.

### HasRirRegistrationAction

`func (o *Ipv6networkcontainer) HasRirRegistrationAction() bool`

HasRirRegistrationAction returns a boolean if a field has been set.

### GetRirRegistrationStatus

`func (o *Ipv6networkcontainer) GetRirRegistrationStatus() string`

GetRirRegistrationStatus returns the RirRegistrationStatus field if non-nil, zero value otherwise.

### GetRirRegistrationStatusOk

`func (o *Ipv6networkcontainer) GetRirRegistrationStatusOk() (*string, bool)`

GetRirRegistrationStatusOk returns a tuple with the RirRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegistrationStatus

`func (o *Ipv6networkcontainer) SetRirRegistrationStatus(v string)`

SetRirRegistrationStatus sets RirRegistrationStatus field to given value.

### HasRirRegistrationStatus

`func (o *Ipv6networkcontainer) HasRirRegistrationStatus() bool`

HasRirRegistrationStatus returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *Ipv6networkcontainer) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *Ipv6networkcontainer) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *Ipv6networkcontainer) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *Ipv6networkcontainer) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetSendRirRequest

`func (o *Ipv6networkcontainer) GetSendRirRequest() bool`

GetSendRirRequest returns the SendRirRequest field if non-nil, zero value otherwise.

### GetSendRirRequestOk

`func (o *Ipv6networkcontainer) GetSendRirRequestOk() (*bool, bool)`

GetSendRirRequestOk returns a tuple with the SendRirRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRirRequest

`func (o *Ipv6networkcontainer) SetSendRirRequest(v bool)`

SetSendRirRequest sets SendRirRequest field to given value.

### HasSendRirRequest

`func (o *Ipv6networkcontainer) HasSendRirRequest() bool`

HasSendRirRequest returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *Ipv6networkcontainer) GetSubscribeSettings() Ipv6networkcontainerSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *Ipv6networkcontainer) GetSubscribeSettingsOk() (*Ipv6networkcontainerSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *Ipv6networkcontainer) SetSubscribeSettings(v Ipv6networkcontainerSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *Ipv6networkcontainer) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetUnmanaged

`func (o *Ipv6networkcontainer) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *Ipv6networkcontainer) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *Ipv6networkcontainer) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *Ipv6networkcontainer) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *Ipv6networkcontainer) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *Ipv6networkcontainer) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *Ipv6networkcontainer) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *Ipv6networkcontainer) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *Ipv6networkcontainer) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *Ipv6networkcontainer) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *Ipv6networkcontainer) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *Ipv6networkcontainer) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Ipv6networkcontainer) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Ipv6networkcontainer) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Ipv6networkcontainer) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Ipv6networkcontainer) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsEnableOptionFqdn

`func (o *Ipv6networkcontainer) GetUseDdnsEnableOptionFqdn() bool`

GetUseDdnsEnableOptionFqdn returns the UseDdnsEnableOptionFqdn field if non-nil, zero value otherwise.

### GetUseDdnsEnableOptionFqdnOk

`func (o *Ipv6networkcontainer) GetUseDdnsEnableOptionFqdnOk() (*bool, bool)`

GetUseDdnsEnableOptionFqdnOk returns a tuple with the UseDdnsEnableOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsEnableOptionFqdn

`func (o *Ipv6networkcontainer) SetUseDdnsEnableOptionFqdn(v bool)`

SetUseDdnsEnableOptionFqdn sets UseDdnsEnableOptionFqdn field to given value.

### HasUseDdnsEnableOptionFqdn

`func (o *Ipv6networkcontainer) HasUseDdnsEnableOptionFqdn() bool`

HasUseDdnsEnableOptionFqdn returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *Ipv6networkcontainer) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *Ipv6networkcontainer) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *Ipv6networkcontainer) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *Ipv6networkcontainer) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *Ipv6networkcontainer) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *Ipv6networkcontainer) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *Ipv6networkcontainer) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *Ipv6networkcontainer) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *Ipv6networkcontainer) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *Ipv6networkcontainer) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *Ipv6networkcontainer) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *Ipv6networkcontainer) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *Ipv6networkcontainer) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *Ipv6networkcontainer) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *Ipv6networkcontainer) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *Ipv6networkcontainer) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Ipv6networkcontainer) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Ipv6networkcontainer) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Ipv6networkcontainer) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Ipv6networkcontainer) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *Ipv6networkcontainer) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *Ipv6networkcontainer) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *Ipv6networkcontainer) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *Ipv6networkcontainer) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6networkcontainer) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6networkcontainer) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6networkcontainer) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6networkcontainer) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMgmPrivate

`func (o *Ipv6networkcontainer) GetUseMgmPrivate() bool`

GetUseMgmPrivate returns the UseMgmPrivate field if non-nil, zero value otherwise.

### GetUseMgmPrivateOk

`func (o *Ipv6networkcontainer) GetUseMgmPrivateOk() (*bool, bool)`

GetUseMgmPrivateOk returns a tuple with the UseMgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmPrivate

`func (o *Ipv6networkcontainer) SetUseMgmPrivate(v bool)`

SetUseMgmPrivate sets UseMgmPrivate field to given value.

### HasUseMgmPrivate

`func (o *Ipv6networkcontainer) HasUseMgmPrivate() bool`

HasUseMgmPrivate returns a boolean if a field has been set.

### GetUseOptions

`func (o *Ipv6networkcontainer) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Ipv6networkcontainer) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Ipv6networkcontainer) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Ipv6networkcontainer) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *Ipv6networkcontainer) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *Ipv6networkcontainer) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *Ipv6networkcontainer) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *Ipv6networkcontainer) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *Ipv6networkcontainer) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *Ipv6networkcontainer) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *Ipv6networkcontainer) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *Ipv6networkcontainer) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6networkcontainer) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *Ipv6networkcontainer) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6networkcontainer) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *Ipv6networkcontainer) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *Ipv6networkcontainer) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *Ipv6networkcontainer) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *Ipv6networkcontainer) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *Ipv6networkcontainer) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetUseZoneAssociations

`func (o *Ipv6networkcontainer) GetUseZoneAssociations() bool`

GetUseZoneAssociations returns the UseZoneAssociations field if non-nil, zero value otherwise.

### GetUseZoneAssociationsOk

`func (o *Ipv6networkcontainer) GetUseZoneAssociationsOk() (*bool, bool)`

GetUseZoneAssociationsOk returns a tuple with the UseZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseZoneAssociations

`func (o *Ipv6networkcontainer) SetUseZoneAssociations(v bool)`

SetUseZoneAssociations sets UseZoneAssociations field to given value.

### HasUseZoneAssociations

`func (o *Ipv6networkcontainer) HasUseZoneAssociations() bool`

HasUseZoneAssociations returns a boolean if a field has been set.

### GetUtilization

`func (o *Ipv6networkcontainer) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *Ipv6networkcontainer) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *Ipv6networkcontainer) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *Ipv6networkcontainer) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetValidLifetime

`func (o *Ipv6networkcontainer) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *Ipv6networkcontainer) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *Ipv6networkcontainer) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *Ipv6networkcontainer) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *Ipv6networkcontainer) GetZoneAssociations() []Ipv6networkcontainerZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *Ipv6networkcontainer) GetZoneAssociationsOk() (*[]Ipv6networkcontainerZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *Ipv6networkcontainer) SetZoneAssociations(v []Ipv6networkcontainerZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *Ipv6networkcontainer) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


