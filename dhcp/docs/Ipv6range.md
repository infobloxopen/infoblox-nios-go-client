# Ipv6range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddressType** | Pointer to **string** | Type of a DHCP IPv6 Range object. Valid values are \&quot;ADDRESS\&quot;, \&quot;PREFIX\&quot;, or \&quot;BOTH\&quot;. When the address type is \&quot;ADDRESS\&quot;, values for the &#39;start_addr&#39; and &#39;end_addr&#39; members are required. When the address type is \&quot;PREFIX\&quot;, values for &#39;ipv6_start_prefix&#39;, &#39;ipv6_end_prefix&#39;, and &#39;ipv6_prefix_bits&#39; are required. When the address type is \&quot;BOTH\&quot;, values for &#39;start_addr&#39;, &#39;end_addr&#39;, &#39;ipv6_start_prefix&#39;, &#39;ipv6_end_prefix&#39;, and &#39;ipv6_prefix_bits&#39; are all required. | [optional] 
**CloudInfo** | Pointer to [**Ipv6rangeCloudInfo**](Ipv6rangeCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the range; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a range is disabled or not. When this is set to False, the range is enabled. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this range. | [optional] [readonly] 
**DiscoveryBasicPollSettings** | Pointer to [**Ipv6rangeDiscoveryBasicPollSettings**](Ipv6rangeDiscoveryBasicPollSettings.md) |  | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**Ipv6rangeDiscoveryBlackoutSetting**](Ipv6rangeDiscoveryBlackoutSetting.md) |  | [optional] 
**DiscoveryMember** | Pointer to **string** | The member that will run discovery for this range. | [optional] 
**EnableDiscovery** | Pointer to **bool** | Determines whether a discovery is enabled or not for this range. When this is set to False, the discovery for this range is disabled. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the range should be immediately enabled. | [optional] 
**EndAddr** | Pointer to **string** | The IPv6 Address end address of the DHCP IPv6 range. | [optional] 
**EndpointSources** | Pointer to **[]map[string]interface{}** | The endpoints that provides data for the DHCP IPv6 Range object. | [optional] [readonly] 
**Exclude** | Pointer to [**[]Ipv6rangeExclude**](Ipv6rangeExclude.md) | These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally,information about this exclusion range. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv6EndPrefix** | Pointer to **string** | The IPv6 Address end prefix of the DHCP IPv6 range. | [optional] 
**Ipv6PrefixBits** | Pointer to **int64** | Prefix bits of the DHCP IPv6 range. | [optional] 
**Ipv6StartPrefix** | Pointer to **string** | The IPv6 Address starting prefix of the DHCP IPv6 range. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6rangeLogicFilterRules**](Ipv6rangeLogicFilterRules.md) | This field contains the logic filters to be applied to this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Member** | Pointer to [**Ipv6rangeMember**](Ipv6rangeMember.md) |  | [optional] 
**Name** | Pointer to **string** | This field contains the name of the Microsoft scope. | [optional] 
**Network** | Pointer to **string** | The network this range belongs to, in IPv6 Address/CIDR format. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this range resides. | [optional] 
**NextAvailableIp** | Pointer to **map[string]interface{}** |  | [optional] 
**OptionFilterRules** | Pointer to [**[]Ipv6rangeOptionFilterRules**](Ipv6rangeOptionFilterRules.md) | This field contains the Option filters to be applied to this IPv6 range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**PortControlBlackoutSetting** | Pointer to [**Ipv6rangePortControlBlackoutSetting**](Ipv6rangePortControlBlackoutSetting.md) |  | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | If the field is set to True, the discovery blackout setting will be used for port control blackout setting. | [optional] 
**ServerAssociationType** | Pointer to **string** | The type of server that is going to serve the range. Valid values are: * MEMBER * NONE | [optional] 
**StartAddr** | Pointer to **string** | The IPv6 Address starting address of the DHCP IPv6 range. | [optional] 
**SubscribeSettings** | Pointer to [**Ipv6rangeSubscribeSettings**](Ipv6rangeSubscribeSettings.md) |  | [optional] 
**Template** | Pointer to **string** | If set on creation, the range will be created according to the values specified in the named template. | [optional] 
**UseBlackoutSetting** | Pointer to **bool** | Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout | [optional] 
**UseDiscoveryBasicPollingSettings** | Pointer to **bool** | Use flag for: discovery_basic_poll_settings | [optional] 
**UseEnableDiscovery** | Pointer to **bool** | Use flag for: discovery_member , enable_discovery | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseSubscribeSettings** | Pointer to **bool** | Use flag for: subscribe_settings | [optional] 

## Methods

### NewIpv6range

`func NewIpv6range() *Ipv6range`

NewIpv6range instantiates a new Ipv6range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangeWithDefaults

`func NewIpv6rangeWithDefaults() *Ipv6range`

NewIpv6rangeWithDefaults instantiates a new Ipv6range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6range) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6range) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6range) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6range) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddressType

`func (o *Ipv6range) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *Ipv6range) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *Ipv6range) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *Ipv6range) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Ipv6range) GetCloudInfo() Ipv6rangeCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Ipv6range) GetCloudInfoOk() (*Ipv6rangeCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Ipv6range) SetCloudInfo(v Ipv6rangeCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Ipv6range) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6range) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6range) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6range) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6range) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *Ipv6range) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Ipv6range) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Ipv6range) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Ipv6range) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Ipv6range) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Ipv6range) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Ipv6range) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Ipv6range) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *Ipv6range) GetDiscoveryBasicPollSettings() Ipv6rangeDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *Ipv6range) GetDiscoveryBasicPollSettingsOk() (*Ipv6rangeDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *Ipv6range) SetDiscoveryBasicPollSettings(v Ipv6rangeDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *Ipv6range) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *Ipv6range) GetDiscoveryBlackoutSetting() Ipv6rangeDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *Ipv6range) GetDiscoveryBlackoutSettingOk() (*Ipv6rangeDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *Ipv6range) SetDiscoveryBlackoutSetting(v Ipv6rangeDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *Ipv6range) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *Ipv6range) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *Ipv6range) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *Ipv6range) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *Ipv6range) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *Ipv6range) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *Ipv6range) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *Ipv6range) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *Ipv6range) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Ipv6range) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Ipv6range) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Ipv6range) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Ipv6range) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEndAddr

`func (o *Ipv6range) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *Ipv6range) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *Ipv6range) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *Ipv6range) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetEndpointSources

`func (o *Ipv6range) GetEndpointSources() []map[string]interface{}`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *Ipv6range) GetEndpointSourcesOk() (*[]map[string]interface{}, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *Ipv6range) SetEndpointSources(v []map[string]interface{})`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *Ipv6range) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExclude

`func (o *Ipv6range) GetExclude() []Ipv6rangeExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Ipv6range) GetExcludeOk() (*[]Ipv6rangeExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Ipv6range) SetExclude(v []Ipv6rangeExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *Ipv6range) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExtattrs

`func (o *Ipv6range) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Ipv6range) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Ipv6range) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Ipv6range) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIpv6EndPrefix

`func (o *Ipv6range) GetIpv6EndPrefix() string`

GetIpv6EndPrefix returns the Ipv6EndPrefix field if non-nil, zero value otherwise.

### GetIpv6EndPrefixOk

`func (o *Ipv6range) GetIpv6EndPrefixOk() (*string, bool)`

GetIpv6EndPrefixOk returns a tuple with the Ipv6EndPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndPrefix

`func (o *Ipv6range) SetIpv6EndPrefix(v string)`

SetIpv6EndPrefix sets Ipv6EndPrefix field to given value.

### HasIpv6EndPrefix

`func (o *Ipv6range) HasIpv6EndPrefix() bool`

HasIpv6EndPrefix returns a boolean if a field has been set.

### GetIpv6PrefixBits

`func (o *Ipv6range) GetIpv6PrefixBits() int64`

GetIpv6PrefixBits returns the Ipv6PrefixBits field if non-nil, zero value otherwise.

### GetIpv6PrefixBitsOk

`func (o *Ipv6range) GetIpv6PrefixBitsOk() (*int64, bool)`

GetIpv6PrefixBitsOk returns a tuple with the Ipv6PrefixBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixBits

`func (o *Ipv6range) SetIpv6PrefixBits(v int64)`

SetIpv6PrefixBits sets Ipv6PrefixBits field to given value.

### HasIpv6PrefixBits

`func (o *Ipv6range) HasIpv6PrefixBits() bool`

HasIpv6PrefixBits returns a boolean if a field has been set.

### GetIpv6StartPrefix

`func (o *Ipv6range) GetIpv6StartPrefix() string`

GetIpv6StartPrefix returns the Ipv6StartPrefix field if non-nil, zero value otherwise.

### GetIpv6StartPrefixOk

`func (o *Ipv6range) GetIpv6StartPrefixOk() (*string, bool)`

GetIpv6StartPrefixOk returns a tuple with the Ipv6StartPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6StartPrefix

`func (o *Ipv6range) SetIpv6StartPrefix(v string)`

SetIpv6StartPrefix sets Ipv6StartPrefix field to given value.

### HasIpv6StartPrefix

`func (o *Ipv6range) HasIpv6StartPrefix() bool`

HasIpv6StartPrefix returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6range) GetLogicFilterRules() []Ipv6rangeLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6range) GetLogicFilterRulesOk() (*[]Ipv6rangeLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6range) SetLogicFilterRules(v []Ipv6rangeLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6range) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *Ipv6range) GetMember() Ipv6rangeMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Ipv6range) GetMemberOk() (*Ipv6rangeMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Ipv6range) SetMember(v Ipv6rangeMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *Ipv6range) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *Ipv6range) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6range) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6range) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6range) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Ipv6range) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ipv6range) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ipv6range) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ipv6range) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Ipv6range) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Ipv6range) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Ipv6range) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Ipv6range) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextAvailableIp

`func (o *Ipv6range) GetNextAvailableIp() map[string]interface{}`

GetNextAvailableIp returns the NextAvailableIp field if non-nil, zero value otherwise.

### GetNextAvailableIpOk

`func (o *Ipv6range) GetNextAvailableIpOk() (*map[string]interface{}, bool)`

GetNextAvailableIpOk returns a tuple with the NextAvailableIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableIp

`func (o *Ipv6range) SetNextAvailableIp(v map[string]interface{})`

SetNextAvailableIp sets NextAvailableIp field to given value.

### HasNextAvailableIp

`func (o *Ipv6range) HasNextAvailableIp() bool`

HasNextAvailableIp returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *Ipv6range) GetOptionFilterRules() []Ipv6rangeOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *Ipv6range) GetOptionFilterRulesOk() (*[]Ipv6rangeOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *Ipv6range) SetOptionFilterRules(v []Ipv6rangeOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *Ipv6range) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *Ipv6range) GetPortControlBlackoutSetting() Ipv6rangePortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *Ipv6range) GetPortControlBlackoutSettingOk() (*Ipv6rangePortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *Ipv6range) SetPortControlBlackoutSetting(v Ipv6rangePortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *Ipv6range) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Ipv6range) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Ipv6range) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Ipv6range) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Ipv6range) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Ipv6range) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Ipv6range) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Ipv6range) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Ipv6range) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *Ipv6range) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *Ipv6range) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *Ipv6range) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *Ipv6range) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *Ipv6range) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *Ipv6range) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *Ipv6range) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *Ipv6range) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetStartAddr

`func (o *Ipv6range) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *Ipv6range) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *Ipv6range) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *Ipv6range) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *Ipv6range) GetSubscribeSettings() Ipv6rangeSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *Ipv6range) GetSubscribeSettingsOk() (*Ipv6rangeSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *Ipv6range) SetSubscribeSettings(v Ipv6rangeSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *Ipv6range) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *Ipv6range) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Ipv6range) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Ipv6range) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Ipv6range) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *Ipv6range) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *Ipv6range) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *Ipv6range) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *Ipv6range) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *Ipv6range) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *Ipv6range) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *Ipv6range) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *Ipv6range) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *Ipv6range) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *Ipv6range) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *Ipv6range) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *Ipv6range) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6range) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6range) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6range) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6range) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Ipv6range) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Ipv6range) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Ipv6range) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Ipv6range) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *Ipv6range) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *Ipv6range) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *Ipv6range) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *Ipv6range) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


