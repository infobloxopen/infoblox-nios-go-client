# GetMemberThreatprotectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The human readable comment for member threat protection properties. | [optional] [readonly] 
**CurrentRuleset** | Pointer to **string** | The ruleset used for threat protection. | [optional] 
**DisableMultipleDnsTcpRequest** | Pointer to **bool** | Determines if multiple BIND responses via TCP connection is enabled or not. | [optional] 
**EnableAccelRespBeforeThreatProtection** | Pointer to **bool** | Determines if DNS responses are sent from acceleration cache before applying Threat Protection rules. Recommended for better performance when using DNS Cache Acceleration. | [optional] 
**EnableNatRules** | Pointer to **bool** | Determines if NAT (Network Address Translation) mapping for threat protection is enabled or not. | [optional] 
**EnableService** | Pointer to **bool** | Determines if the Threat protection service is enabled or not. | [optional] 
**EventsPerSecondPerRule** | Pointer to **int64** | The number of events logged per second per rule. | [optional] 
**HardwareModel** | Pointer to **string** | The hardware model of the member. | [optional] [readonly] 
**HardwareType** | Pointer to **string** | The hardware type of the member. | [optional] [readonly] 
**HostName** | Pointer to **string** | A Grid member name. | [optional] [readonly] 
**Ipv4address** | Pointer to **string** | The IPv4 address of member threat protection service. | [optional] [readonly] 
**Ipv6address** | Pointer to **string** | The IPv6 address of member threat protection service. | [optional] [readonly] 
**NatRules** | Pointer to [**[]MemberThreatprotectionNatRules**](MemberThreatprotectionNatRules.md) | The list of NAT rules. | [optional] 
**OutboundSettings** | Pointer to [**MemberThreatprotectionOutboundSettings**](MemberThreatprotectionOutboundSettings.md) |  | [optional] 
**Profile** | Pointer to **string** | The Threat Protection profile associated with the member. | [optional] 
**UseCurrentRuleset** | Pointer to **bool** | Use flag for: current_ruleset | [optional] 
**UseDisableMultipleDnsTcpRequest** | Pointer to **bool** | Use flag for: disable_multiple_dns_tcp_request | [optional] 
**UseEnableAccelRespBeforeThreatProtection** | Pointer to **bool** | Use flag for: enable_accel_resp_before_threat_protection | [optional] 
**UseEnableNatRules** | Pointer to **bool** | Use flag for: enable_nat_rules | [optional] 
**UseEventsPerSecondPerRule** | Pointer to **bool** | Use flag for: events_per_second_per_rule | [optional] 
**UseOutboundSettings** | Pointer to **bool** | Use flag for: outbound_settings | [optional] 
**Result** | Pointer to [**MemberThreatprotection**](MemberThreatprotection.md) |  | [optional] 

## Methods

### NewGetMemberThreatprotectionResponse

`func NewGetMemberThreatprotectionResponse() *GetMemberThreatprotectionResponse`

NewGetMemberThreatprotectionResponse instantiates a new GetMemberThreatprotectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMemberThreatprotectionResponseWithDefaults

`func NewGetMemberThreatprotectionResponseWithDefaults() *GetMemberThreatprotectionResponse`

NewGetMemberThreatprotectionResponseWithDefaults instantiates a new GetMemberThreatprotectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMemberThreatprotectionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMemberThreatprotectionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMemberThreatprotectionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMemberThreatprotectionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetMemberThreatprotectionResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetMemberThreatprotectionResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetMemberThreatprotectionResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetMemberThreatprotectionResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentRuleset

`func (o *GetMemberThreatprotectionResponse) GetCurrentRuleset() string`

GetCurrentRuleset returns the CurrentRuleset field if non-nil, zero value otherwise.

### GetCurrentRulesetOk

`func (o *GetMemberThreatprotectionResponse) GetCurrentRulesetOk() (*string, bool)`

GetCurrentRulesetOk returns a tuple with the CurrentRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRuleset

`func (o *GetMemberThreatprotectionResponse) SetCurrentRuleset(v string)`

SetCurrentRuleset sets CurrentRuleset field to given value.

### HasCurrentRuleset

`func (o *GetMemberThreatprotectionResponse) HasCurrentRuleset() bool`

HasCurrentRuleset returns a boolean if a field has been set.

### GetDisableMultipleDnsTcpRequest

`func (o *GetMemberThreatprotectionResponse) GetDisableMultipleDnsTcpRequest() bool`

GetDisableMultipleDnsTcpRequest returns the DisableMultipleDnsTcpRequest field if non-nil, zero value otherwise.

### GetDisableMultipleDnsTcpRequestOk

`func (o *GetMemberThreatprotectionResponse) GetDisableMultipleDnsTcpRequestOk() (*bool, bool)`

GetDisableMultipleDnsTcpRequestOk returns a tuple with the DisableMultipleDnsTcpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMultipleDnsTcpRequest

`func (o *GetMemberThreatprotectionResponse) SetDisableMultipleDnsTcpRequest(v bool)`

SetDisableMultipleDnsTcpRequest sets DisableMultipleDnsTcpRequest field to given value.

### HasDisableMultipleDnsTcpRequest

`func (o *GetMemberThreatprotectionResponse) HasDisableMultipleDnsTcpRequest() bool`

HasDisableMultipleDnsTcpRequest returns a boolean if a field has been set.

### GetEnableAccelRespBeforeThreatProtection

`func (o *GetMemberThreatprotectionResponse) GetEnableAccelRespBeforeThreatProtection() bool`

GetEnableAccelRespBeforeThreatProtection returns the EnableAccelRespBeforeThreatProtection field if non-nil, zero value otherwise.

### GetEnableAccelRespBeforeThreatProtectionOk

`func (o *GetMemberThreatprotectionResponse) GetEnableAccelRespBeforeThreatProtectionOk() (*bool, bool)`

GetEnableAccelRespBeforeThreatProtectionOk returns a tuple with the EnableAccelRespBeforeThreatProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccelRespBeforeThreatProtection

`func (o *GetMemberThreatprotectionResponse) SetEnableAccelRespBeforeThreatProtection(v bool)`

SetEnableAccelRespBeforeThreatProtection sets EnableAccelRespBeforeThreatProtection field to given value.

### HasEnableAccelRespBeforeThreatProtection

`func (o *GetMemberThreatprotectionResponse) HasEnableAccelRespBeforeThreatProtection() bool`

HasEnableAccelRespBeforeThreatProtection returns a boolean if a field has been set.

### GetEnableNatRules

`func (o *GetMemberThreatprotectionResponse) GetEnableNatRules() bool`

GetEnableNatRules returns the EnableNatRules field if non-nil, zero value otherwise.

### GetEnableNatRulesOk

`func (o *GetMemberThreatprotectionResponse) GetEnableNatRulesOk() (*bool, bool)`

GetEnableNatRulesOk returns a tuple with the EnableNatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNatRules

`func (o *GetMemberThreatprotectionResponse) SetEnableNatRules(v bool)`

SetEnableNatRules sets EnableNatRules field to given value.

### HasEnableNatRules

`func (o *GetMemberThreatprotectionResponse) HasEnableNatRules() bool`

HasEnableNatRules returns a boolean if a field has been set.

### GetEnableService

`func (o *GetMemberThreatprotectionResponse) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *GetMemberThreatprotectionResponse) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *GetMemberThreatprotectionResponse) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *GetMemberThreatprotectionResponse) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetEventsPerSecondPerRule

`func (o *GetMemberThreatprotectionResponse) GetEventsPerSecondPerRule() int64`

GetEventsPerSecondPerRule returns the EventsPerSecondPerRule field if non-nil, zero value otherwise.

### GetEventsPerSecondPerRuleOk

`func (o *GetMemberThreatprotectionResponse) GetEventsPerSecondPerRuleOk() (*int64, bool)`

GetEventsPerSecondPerRuleOk returns a tuple with the EventsPerSecondPerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsPerSecondPerRule

`func (o *GetMemberThreatprotectionResponse) SetEventsPerSecondPerRule(v int64)`

SetEventsPerSecondPerRule sets EventsPerSecondPerRule field to given value.

### HasEventsPerSecondPerRule

`func (o *GetMemberThreatprotectionResponse) HasEventsPerSecondPerRule() bool`

HasEventsPerSecondPerRule returns a boolean if a field has been set.

### GetHardwareModel

`func (o *GetMemberThreatprotectionResponse) GetHardwareModel() string`

GetHardwareModel returns the HardwareModel field if non-nil, zero value otherwise.

### GetHardwareModelOk

`func (o *GetMemberThreatprotectionResponse) GetHardwareModelOk() (*string, bool)`

GetHardwareModelOk returns a tuple with the HardwareModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareModel

`func (o *GetMemberThreatprotectionResponse) SetHardwareModel(v string)`

SetHardwareModel sets HardwareModel field to given value.

### HasHardwareModel

`func (o *GetMemberThreatprotectionResponse) HasHardwareModel() bool`

HasHardwareModel returns a boolean if a field has been set.

### GetHardwareType

`func (o *GetMemberThreatprotectionResponse) GetHardwareType() string`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *GetMemberThreatprotectionResponse) GetHardwareTypeOk() (*string, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *GetMemberThreatprotectionResponse) SetHardwareType(v string)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *GetMemberThreatprotectionResponse) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### GetHostName

`func (o *GetMemberThreatprotectionResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetMemberThreatprotectionResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetMemberThreatprotectionResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetMemberThreatprotectionResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpv4address

`func (o *GetMemberThreatprotectionResponse) GetIpv4address() string`

GetIpv4address returns the Ipv4address field if non-nil, zero value otherwise.

### GetIpv4addressOk

`func (o *GetMemberThreatprotectionResponse) GetIpv4addressOk() (*string, bool)`

GetIpv4addressOk returns a tuple with the Ipv4address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4address

`func (o *GetMemberThreatprotectionResponse) SetIpv4address(v string)`

SetIpv4address sets Ipv4address field to given value.

### HasIpv4address

`func (o *GetMemberThreatprotectionResponse) HasIpv4address() bool`

HasIpv4address returns a boolean if a field has been set.

### GetIpv6address

`func (o *GetMemberThreatprotectionResponse) GetIpv6address() string`

GetIpv6address returns the Ipv6address field if non-nil, zero value otherwise.

### GetIpv6addressOk

`func (o *GetMemberThreatprotectionResponse) GetIpv6addressOk() (*string, bool)`

GetIpv6addressOk returns a tuple with the Ipv6address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6address

`func (o *GetMemberThreatprotectionResponse) SetIpv6address(v string)`

SetIpv6address sets Ipv6address field to given value.

### HasIpv6address

`func (o *GetMemberThreatprotectionResponse) HasIpv6address() bool`

HasIpv6address returns a boolean if a field has been set.

### GetNatRules

`func (o *GetMemberThreatprotectionResponse) GetNatRules() []MemberThreatprotectionNatRules`

GetNatRules returns the NatRules field if non-nil, zero value otherwise.

### GetNatRulesOk

`func (o *GetMemberThreatprotectionResponse) GetNatRulesOk() (*[]MemberThreatprotectionNatRules, bool)`

GetNatRulesOk returns a tuple with the NatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatRules

`func (o *GetMemberThreatprotectionResponse) SetNatRules(v []MemberThreatprotectionNatRules)`

SetNatRules sets NatRules field to given value.

### HasNatRules

`func (o *GetMemberThreatprotectionResponse) HasNatRules() bool`

HasNatRules returns a boolean if a field has been set.

### GetOutboundSettings

`func (o *GetMemberThreatprotectionResponse) GetOutboundSettings() MemberThreatprotectionOutboundSettings`

GetOutboundSettings returns the OutboundSettings field if non-nil, zero value otherwise.

### GetOutboundSettingsOk

`func (o *GetMemberThreatprotectionResponse) GetOutboundSettingsOk() (*MemberThreatprotectionOutboundSettings, bool)`

GetOutboundSettingsOk returns a tuple with the OutboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSettings

`func (o *GetMemberThreatprotectionResponse) SetOutboundSettings(v MemberThreatprotectionOutboundSettings)`

SetOutboundSettings sets OutboundSettings field to given value.

### HasOutboundSettings

`func (o *GetMemberThreatprotectionResponse) HasOutboundSettings() bool`

HasOutboundSettings returns a boolean if a field has been set.

### GetProfile

`func (o *GetMemberThreatprotectionResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetMemberThreatprotectionResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetMemberThreatprotectionResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetMemberThreatprotectionResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetUseCurrentRuleset

`func (o *GetMemberThreatprotectionResponse) GetUseCurrentRuleset() bool`

GetUseCurrentRuleset returns the UseCurrentRuleset field if non-nil, zero value otherwise.

### GetUseCurrentRulesetOk

`func (o *GetMemberThreatprotectionResponse) GetUseCurrentRulesetOk() (*bool, bool)`

GetUseCurrentRulesetOk returns a tuple with the UseCurrentRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCurrentRuleset

`func (o *GetMemberThreatprotectionResponse) SetUseCurrentRuleset(v bool)`

SetUseCurrentRuleset sets UseCurrentRuleset field to given value.

### HasUseCurrentRuleset

`func (o *GetMemberThreatprotectionResponse) HasUseCurrentRuleset() bool`

HasUseCurrentRuleset returns a boolean if a field has been set.

### GetUseDisableMultipleDnsTcpRequest

`func (o *GetMemberThreatprotectionResponse) GetUseDisableMultipleDnsTcpRequest() bool`

GetUseDisableMultipleDnsTcpRequest returns the UseDisableMultipleDnsTcpRequest field if non-nil, zero value otherwise.

### GetUseDisableMultipleDnsTcpRequestOk

`func (o *GetMemberThreatprotectionResponse) GetUseDisableMultipleDnsTcpRequestOk() (*bool, bool)`

GetUseDisableMultipleDnsTcpRequestOk returns a tuple with the UseDisableMultipleDnsTcpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisableMultipleDnsTcpRequest

`func (o *GetMemberThreatprotectionResponse) SetUseDisableMultipleDnsTcpRequest(v bool)`

SetUseDisableMultipleDnsTcpRequest sets UseDisableMultipleDnsTcpRequest field to given value.

### HasUseDisableMultipleDnsTcpRequest

`func (o *GetMemberThreatprotectionResponse) HasUseDisableMultipleDnsTcpRequest() bool`

HasUseDisableMultipleDnsTcpRequest returns a boolean if a field has been set.

### GetUseEnableAccelRespBeforeThreatProtection

`func (o *GetMemberThreatprotectionResponse) GetUseEnableAccelRespBeforeThreatProtection() bool`

GetUseEnableAccelRespBeforeThreatProtection returns the UseEnableAccelRespBeforeThreatProtection field if non-nil, zero value otherwise.

### GetUseEnableAccelRespBeforeThreatProtectionOk

`func (o *GetMemberThreatprotectionResponse) GetUseEnableAccelRespBeforeThreatProtectionOk() (*bool, bool)`

GetUseEnableAccelRespBeforeThreatProtectionOk returns a tuple with the UseEnableAccelRespBeforeThreatProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableAccelRespBeforeThreatProtection

`func (o *GetMemberThreatprotectionResponse) SetUseEnableAccelRespBeforeThreatProtection(v bool)`

SetUseEnableAccelRespBeforeThreatProtection sets UseEnableAccelRespBeforeThreatProtection field to given value.

### HasUseEnableAccelRespBeforeThreatProtection

`func (o *GetMemberThreatprotectionResponse) HasUseEnableAccelRespBeforeThreatProtection() bool`

HasUseEnableAccelRespBeforeThreatProtection returns a boolean if a field has been set.

### GetUseEnableNatRules

`func (o *GetMemberThreatprotectionResponse) GetUseEnableNatRules() bool`

GetUseEnableNatRules returns the UseEnableNatRules field if non-nil, zero value otherwise.

### GetUseEnableNatRulesOk

`func (o *GetMemberThreatprotectionResponse) GetUseEnableNatRulesOk() (*bool, bool)`

GetUseEnableNatRulesOk returns a tuple with the UseEnableNatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableNatRules

`func (o *GetMemberThreatprotectionResponse) SetUseEnableNatRules(v bool)`

SetUseEnableNatRules sets UseEnableNatRules field to given value.

### HasUseEnableNatRules

`func (o *GetMemberThreatprotectionResponse) HasUseEnableNatRules() bool`

HasUseEnableNatRules returns a boolean if a field has been set.

### GetUseEventsPerSecondPerRule

`func (o *GetMemberThreatprotectionResponse) GetUseEventsPerSecondPerRule() bool`

GetUseEventsPerSecondPerRule returns the UseEventsPerSecondPerRule field if non-nil, zero value otherwise.

### GetUseEventsPerSecondPerRuleOk

`func (o *GetMemberThreatprotectionResponse) GetUseEventsPerSecondPerRuleOk() (*bool, bool)`

GetUseEventsPerSecondPerRuleOk returns a tuple with the UseEventsPerSecondPerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEventsPerSecondPerRule

`func (o *GetMemberThreatprotectionResponse) SetUseEventsPerSecondPerRule(v bool)`

SetUseEventsPerSecondPerRule sets UseEventsPerSecondPerRule field to given value.

### HasUseEventsPerSecondPerRule

`func (o *GetMemberThreatprotectionResponse) HasUseEventsPerSecondPerRule() bool`

HasUseEventsPerSecondPerRule returns a boolean if a field has been set.

### GetUseOutboundSettings

`func (o *GetMemberThreatprotectionResponse) GetUseOutboundSettings() bool`

GetUseOutboundSettings returns the UseOutboundSettings field if non-nil, zero value otherwise.

### GetUseOutboundSettingsOk

`func (o *GetMemberThreatprotectionResponse) GetUseOutboundSettingsOk() (*bool, bool)`

GetUseOutboundSettingsOk returns a tuple with the UseOutboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOutboundSettings

`func (o *GetMemberThreatprotectionResponse) SetUseOutboundSettings(v bool)`

SetUseOutboundSettings sets UseOutboundSettings field to given value.

### HasUseOutboundSettings

`func (o *GetMemberThreatprotectionResponse) HasUseOutboundSettings() bool`

HasUseOutboundSettings returns a boolean if a field has been set.

### GetResult

`func (o *GetMemberThreatprotectionResponse) GetResult() MemberThreatprotection`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMemberThreatprotectionResponse) GetResultOk() (*MemberThreatprotection, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMemberThreatprotectionResponse) SetResult(v MemberThreatprotection)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMemberThreatprotectionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


