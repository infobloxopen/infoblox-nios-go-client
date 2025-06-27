# GetIpv6rangeResponse

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
**EndpointSources** | Pointer to **[]string** | The endpoints that provides data for the DHCP IPv6 Range object. | [optional] [readonly] 
**Exclude** | Pointer to [**[]Ipv6rangeExclude**](Ipv6rangeExclude.md) | These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally,information about this exclusion range. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv6EndPrefix** | Pointer to **string** | The IPv6 Address end prefix of the DHCP IPv6 range. | [optional] 
**Ipv6PrefixBits** | Pointer to **int64** | Prefix bits of the DHCP IPv6 range. | [optional] 
**Ipv6StartPrefix** | Pointer to **string** | The IPv6 Address starting prefix of the DHCP IPv6 range. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6rangeLogicFilterRules**](Ipv6rangeLogicFilterRules.md) | This field contains the logic filters to be applied to this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Member** | Pointer to [**Ipv6rangeMember**](Ipv6rangeMember.md) |  | [optional] 
**Name** | Pointer to **string** | This field contains the name of the Microsoft scope. | [optional] 
**Network** | Pointer to **string** | The network this range belongs to, in IPv6 Address/CIDR format. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this range resides. | [optional] 
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
**Result** | Pointer to [**Ipv6range**](Ipv6range.md) |  | [optional] 

## Methods

### NewGetIpv6rangeResponse

`func NewGetIpv6rangeResponse() *GetIpv6rangeResponse`

NewGetIpv6rangeResponse instantiates a new GetIpv6rangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6rangeResponseWithDefaults

`func NewGetIpv6rangeResponseWithDefaults() *GetIpv6rangeResponse`

NewGetIpv6rangeResponseWithDefaults instantiates a new GetIpv6rangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6rangeResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6rangeResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6rangeResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6rangeResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddressType

`func (o *GetIpv6rangeResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *GetIpv6rangeResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *GetIpv6rangeResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *GetIpv6rangeResponse) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetIpv6rangeResponse) GetCloudInfo() Ipv6rangeCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetIpv6rangeResponse) GetCloudInfoOk() (*Ipv6rangeCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetIpv6rangeResponse) SetCloudInfo(v Ipv6rangeCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetIpv6rangeResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6rangeResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6rangeResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6rangeResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6rangeResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetIpv6rangeResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetIpv6rangeResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetIpv6rangeResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetIpv6rangeResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetIpv6rangeResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetIpv6rangeResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetIpv6rangeResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetIpv6rangeResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveryBasicPollSettings

`func (o *GetIpv6rangeResponse) GetDiscoveryBasicPollSettings() Ipv6rangeDiscoveryBasicPollSettings`

GetDiscoveryBasicPollSettings returns the DiscoveryBasicPollSettings field if non-nil, zero value otherwise.

### GetDiscoveryBasicPollSettingsOk

`func (o *GetIpv6rangeResponse) GetDiscoveryBasicPollSettingsOk() (*Ipv6rangeDiscoveryBasicPollSettings, bool)`

GetDiscoveryBasicPollSettingsOk returns a tuple with the DiscoveryBasicPollSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBasicPollSettings

`func (o *GetIpv6rangeResponse) SetDiscoveryBasicPollSettings(v Ipv6rangeDiscoveryBasicPollSettings)`

SetDiscoveryBasicPollSettings sets DiscoveryBasicPollSettings field to given value.

### HasDiscoveryBasicPollSettings

`func (o *GetIpv6rangeResponse) HasDiscoveryBasicPollSettings() bool`

HasDiscoveryBasicPollSettings returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *GetIpv6rangeResponse) GetDiscoveryBlackoutSetting() Ipv6rangeDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *GetIpv6rangeResponse) GetDiscoveryBlackoutSettingOk() (*Ipv6rangeDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *GetIpv6rangeResponse) SetDiscoveryBlackoutSetting(v Ipv6rangeDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *GetIpv6rangeResponse) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *GetIpv6rangeResponse) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *GetIpv6rangeResponse) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *GetIpv6rangeResponse) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *GetIpv6rangeResponse) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetEnableDiscovery

`func (o *GetIpv6rangeResponse) GetEnableDiscovery() bool`

GetEnableDiscovery returns the EnableDiscovery field if non-nil, zero value otherwise.

### GetEnableDiscoveryOk

`func (o *GetIpv6rangeResponse) GetEnableDiscoveryOk() (*bool, bool)`

GetEnableDiscoveryOk returns a tuple with the EnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscovery

`func (o *GetIpv6rangeResponse) SetEnableDiscovery(v bool)`

SetEnableDiscovery sets EnableDiscovery field to given value.

### HasEnableDiscovery

`func (o *GetIpv6rangeResponse) HasEnableDiscovery() bool`

HasEnableDiscovery returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *GetIpv6rangeResponse) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *GetIpv6rangeResponse) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *GetIpv6rangeResponse) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *GetIpv6rangeResponse) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEndAddr

`func (o *GetIpv6rangeResponse) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *GetIpv6rangeResponse) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *GetIpv6rangeResponse) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *GetIpv6rangeResponse) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetEndpointSources

`func (o *GetIpv6rangeResponse) GetEndpointSources() []string`

GetEndpointSources returns the EndpointSources field if non-nil, zero value otherwise.

### GetEndpointSourcesOk

`func (o *GetIpv6rangeResponse) GetEndpointSourcesOk() (*[]string, bool)`

GetEndpointSourcesOk returns a tuple with the EndpointSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSources

`func (o *GetIpv6rangeResponse) SetEndpointSources(v []string)`

SetEndpointSources sets EndpointSources field to given value.

### HasEndpointSources

`func (o *GetIpv6rangeResponse) HasEndpointSources() bool`

HasEndpointSources returns a boolean if a field has been set.

### GetExclude

`func (o *GetIpv6rangeResponse) GetExclude() []Ipv6rangeExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *GetIpv6rangeResponse) GetExcludeOk() (*[]Ipv6rangeExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *GetIpv6rangeResponse) SetExclude(v []Ipv6rangeExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *GetIpv6rangeResponse) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetIpv6rangeResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetIpv6rangeResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetIpv6rangeResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetIpv6rangeResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv6EndPrefix

`func (o *GetIpv6rangeResponse) GetIpv6EndPrefix() string`

GetIpv6EndPrefix returns the Ipv6EndPrefix field if non-nil, zero value otherwise.

### GetIpv6EndPrefixOk

`func (o *GetIpv6rangeResponse) GetIpv6EndPrefixOk() (*string, bool)`

GetIpv6EndPrefixOk returns a tuple with the Ipv6EndPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndPrefix

`func (o *GetIpv6rangeResponse) SetIpv6EndPrefix(v string)`

SetIpv6EndPrefix sets Ipv6EndPrefix field to given value.

### HasIpv6EndPrefix

`func (o *GetIpv6rangeResponse) HasIpv6EndPrefix() bool`

HasIpv6EndPrefix returns a boolean if a field has been set.

### GetIpv6PrefixBits

`func (o *GetIpv6rangeResponse) GetIpv6PrefixBits() int64`

GetIpv6PrefixBits returns the Ipv6PrefixBits field if non-nil, zero value otherwise.

### GetIpv6PrefixBitsOk

`func (o *GetIpv6rangeResponse) GetIpv6PrefixBitsOk() (*int64, bool)`

GetIpv6PrefixBitsOk returns a tuple with the Ipv6PrefixBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixBits

`func (o *GetIpv6rangeResponse) SetIpv6PrefixBits(v int64)`

SetIpv6PrefixBits sets Ipv6PrefixBits field to given value.

### HasIpv6PrefixBits

`func (o *GetIpv6rangeResponse) HasIpv6PrefixBits() bool`

HasIpv6PrefixBits returns a boolean if a field has been set.

### GetIpv6StartPrefix

`func (o *GetIpv6rangeResponse) GetIpv6StartPrefix() string`

GetIpv6StartPrefix returns the Ipv6StartPrefix field if non-nil, zero value otherwise.

### GetIpv6StartPrefixOk

`func (o *GetIpv6rangeResponse) GetIpv6StartPrefixOk() (*string, bool)`

GetIpv6StartPrefixOk returns a tuple with the Ipv6StartPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6StartPrefix

`func (o *GetIpv6rangeResponse) SetIpv6StartPrefix(v string)`

SetIpv6StartPrefix sets Ipv6StartPrefix field to given value.

### HasIpv6StartPrefix

`func (o *GetIpv6rangeResponse) HasIpv6StartPrefix() bool`

HasIpv6StartPrefix returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetIpv6rangeResponse) GetLogicFilterRules() []Ipv6rangeLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetIpv6rangeResponse) GetLogicFilterRulesOk() (*[]Ipv6rangeLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetIpv6rangeResponse) SetLogicFilterRules(v []Ipv6rangeLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetIpv6rangeResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *GetIpv6rangeResponse) GetMember() Ipv6rangeMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetIpv6rangeResponse) GetMemberOk() (*Ipv6rangeMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetIpv6rangeResponse) SetMember(v Ipv6rangeMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetIpv6rangeResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6rangeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6rangeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6rangeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6rangeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetIpv6rangeResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetIpv6rangeResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetIpv6rangeResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetIpv6rangeResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetIpv6rangeResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetIpv6rangeResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetIpv6rangeResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetIpv6rangeResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *GetIpv6rangeResponse) GetOptionFilterRules() []Ipv6rangeOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *GetIpv6rangeResponse) GetOptionFilterRulesOk() (*[]Ipv6rangeOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *GetIpv6rangeResponse) SetOptionFilterRules(v []Ipv6rangeOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *GetIpv6rangeResponse) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *GetIpv6rangeResponse) GetPortControlBlackoutSetting() Ipv6rangePortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *GetIpv6rangeResponse) GetPortControlBlackoutSettingOk() (*Ipv6rangePortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *GetIpv6rangeResponse) SetPortControlBlackoutSetting(v Ipv6rangePortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *GetIpv6rangeResponse) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetIpv6rangeResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetIpv6rangeResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetIpv6rangeResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetIpv6rangeResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetIpv6rangeResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetIpv6rangeResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetIpv6rangeResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetIpv6rangeResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *GetIpv6rangeResponse) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *GetIpv6rangeResponse) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *GetIpv6rangeResponse) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *GetIpv6rangeResponse) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *GetIpv6rangeResponse) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *GetIpv6rangeResponse) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *GetIpv6rangeResponse) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *GetIpv6rangeResponse) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetStartAddr

`func (o *GetIpv6rangeResponse) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *GetIpv6rangeResponse) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *GetIpv6rangeResponse) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *GetIpv6rangeResponse) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *GetIpv6rangeResponse) GetSubscribeSettings() Ipv6rangeSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *GetIpv6rangeResponse) GetSubscribeSettingsOk() (*Ipv6rangeSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *GetIpv6rangeResponse) SetSubscribeSettings(v Ipv6rangeSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *GetIpv6rangeResponse) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplate

`func (o *GetIpv6rangeResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetIpv6rangeResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetIpv6rangeResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetIpv6rangeResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUseBlackoutSetting

`func (o *GetIpv6rangeResponse) GetUseBlackoutSetting() bool`

GetUseBlackoutSetting returns the UseBlackoutSetting field if non-nil, zero value otherwise.

### GetUseBlackoutSettingOk

`func (o *GetIpv6rangeResponse) GetUseBlackoutSettingOk() (*bool, bool)`

GetUseBlackoutSettingOk returns a tuple with the UseBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBlackoutSetting

`func (o *GetIpv6rangeResponse) SetUseBlackoutSetting(v bool)`

SetUseBlackoutSetting sets UseBlackoutSetting field to given value.

### HasUseBlackoutSetting

`func (o *GetIpv6rangeResponse) HasUseBlackoutSetting() bool`

HasUseBlackoutSetting returns a boolean if a field has been set.

### GetUseDiscoveryBasicPollingSettings

`func (o *GetIpv6rangeResponse) GetUseDiscoveryBasicPollingSettings() bool`

GetUseDiscoveryBasicPollingSettings returns the UseDiscoveryBasicPollingSettings field if non-nil, zero value otherwise.

### GetUseDiscoveryBasicPollingSettingsOk

`func (o *GetIpv6rangeResponse) GetUseDiscoveryBasicPollingSettingsOk() (*bool, bool)`

GetUseDiscoveryBasicPollingSettingsOk returns a tuple with the UseDiscoveryBasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryBasicPollingSettings

`func (o *GetIpv6rangeResponse) SetUseDiscoveryBasicPollingSettings(v bool)`

SetUseDiscoveryBasicPollingSettings sets UseDiscoveryBasicPollingSettings field to given value.

### HasUseDiscoveryBasicPollingSettings

`func (o *GetIpv6rangeResponse) HasUseDiscoveryBasicPollingSettings() bool`

HasUseDiscoveryBasicPollingSettings returns a boolean if a field has been set.

### GetUseEnableDiscovery

`func (o *GetIpv6rangeResponse) GetUseEnableDiscovery() bool`

GetUseEnableDiscovery returns the UseEnableDiscovery field if non-nil, zero value otherwise.

### GetUseEnableDiscoveryOk

`func (o *GetIpv6rangeResponse) GetUseEnableDiscoveryOk() (*bool, bool)`

GetUseEnableDiscoveryOk returns a tuple with the UseEnableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDiscovery

`func (o *GetIpv6rangeResponse) SetUseEnableDiscovery(v bool)`

SetUseEnableDiscovery sets UseEnableDiscovery field to given value.

### HasUseEnableDiscovery

`func (o *GetIpv6rangeResponse) HasUseEnableDiscovery() bool`

HasUseEnableDiscovery returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetIpv6rangeResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetIpv6rangeResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetIpv6rangeResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetIpv6rangeResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetIpv6rangeResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetIpv6rangeResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetIpv6rangeResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetIpv6rangeResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseSubscribeSettings

`func (o *GetIpv6rangeResponse) GetUseSubscribeSettings() bool`

GetUseSubscribeSettings returns the UseSubscribeSettings field if non-nil, zero value otherwise.

### GetUseSubscribeSettingsOk

`func (o *GetIpv6rangeResponse) GetUseSubscribeSettingsOk() (*bool, bool)`

GetUseSubscribeSettingsOk returns a tuple with the UseSubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubscribeSettings

`func (o *GetIpv6rangeResponse) SetUseSubscribeSettings(v bool)`

SetUseSubscribeSettings sets UseSubscribeSettings field to given value.

### HasUseSubscribeSettings

`func (o *GetIpv6rangeResponse) HasUseSubscribeSettings() bool`

HasUseSubscribeSettings returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6rangeResponse) GetResult() Ipv6range`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6rangeResponse) GetResultOk() (*Ipv6range, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6rangeResponse) SetResult(v Ipv6range)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6rangeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


