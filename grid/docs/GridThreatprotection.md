# GridThreatprotection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CurrentRuleset** | Pointer to **string** | The current Grid ruleset. | [optional] 
**DisableMultipleDnsTcpRequest** | Pointer to **bool** | Determines if multiple BIND responses via TCP connection are disabled. | [optional] 
**EnableAccelRespBeforeThreatProtection** | Pointer to **bool** | Determines if DNS responses are sent from acceleration cache before applying Threat Protection rules. Recommended for better performance when using DNS Cache Acceleration. | [optional] 
**EnableAutoDownload** | Pointer to **bool** | Determines if auto download service is enabled. | [optional] 
**EnableNatRules** | Pointer to **bool** | Determines if NAT (Network Address Translation) mapping for threat protection is enabled or not. | [optional] 
**EnableScheduledDownload** | Pointer to **bool** | Determines if scheduled download is enabled. The default frequency is once in every 24 hours if it is disabled. | [optional] 
**EventsPerSecondPerRule** | Pointer to **int64** | The number of events logged per second per rule. | [optional] 
**GridName** | Pointer to **string** | The Grid name. | [optional] [readonly] 
**LastCheckedForUpdate** | Pointer to **int64** | The time when the Grid last checked for updates. | [optional] [readonly] 
**LastRuleUpdateTimestamp** | Pointer to **int64** | The last rule update timestamp. | [optional] [readonly] 
**LastRuleUpdateVersion** | Pointer to **string** | The version of last rule update. | [optional] [readonly] 
**NatRules** | Pointer to [**[]GridThreatprotectionNatRules**](GridThreatprotectionNatRules.md) | The list of NAT mapping rules for threat protection. | [optional] 
**OutboundSettings** | Pointer to [**GridThreatprotectionOutboundSettings**](GridThreatprotectionOutboundSettings.md) |  | [optional] 
**RuleUpdatePolicy** | Pointer to **string** | The update rule policy. | [optional] 
**ScheduledDownload** | Pointer to [**GridThreatprotectionScheduledDownload**](GridThreatprotectionScheduledDownload.md) |  | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewGridThreatprotection

`func NewGridThreatprotection() *GridThreatprotection`

NewGridThreatprotection instantiates a new GridThreatprotection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridThreatprotectionWithDefaults

`func NewGridThreatprotectionWithDefaults() *GridThreatprotection`

NewGridThreatprotectionWithDefaults instantiates a new GridThreatprotection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridThreatprotection) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridThreatprotection) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridThreatprotection) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridThreatprotection) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCurrentRuleset

`func (o *GridThreatprotection) GetCurrentRuleset() string`

GetCurrentRuleset returns the CurrentRuleset field if non-nil, zero value otherwise.

### GetCurrentRulesetOk

`func (o *GridThreatprotection) GetCurrentRulesetOk() (*string, bool)`

GetCurrentRulesetOk returns a tuple with the CurrentRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRuleset

`func (o *GridThreatprotection) SetCurrentRuleset(v string)`

SetCurrentRuleset sets CurrentRuleset field to given value.

### HasCurrentRuleset

`func (o *GridThreatprotection) HasCurrentRuleset() bool`

HasCurrentRuleset returns a boolean if a field has been set.

### GetDisableMultipleDnsTcpRequest

`func (o *GridThreatprotection) GetDisableMultipleDnsTcpRequest() bool`

GetDisableMultipleDnsTcpRequest returns the DisableMultipleDnsTcpRequest field if non-nil, zero value otherwise.

### GetDisableMultipleDnsTcpRequestOk

`func (o *GridThreatprotection) GetDisableMultipleDnsTcpRequestOk() (*bool, bool)`

GetDisableMultipleDnsTcpRequestOk returns a tuple with the DisableMultipleDnsTcpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMultipleDnsTcpRequest

`func (o *GridThreatprotection) SetDisableMultipleDnsTcpRequest(v bool)`

SetDisableMultipleDnsTcpRequest sets DisableMultipleDnsTcpRequest field to given value.

### HasDisableMultipleDnsTcpRequest

`func (o *GridThreatprotection) HasDisableMultipleDnsTcpRequest() bool`

HasDisableMultipleDnsTcpRequest returns a boolean if a field has been set.

### GetEnableAccelRespBeforeThreatProtection

`func (o *GridThreatprotection) GetEnableAccelRespBeforeThreatProtection() bool`

GetEnableAccelRespBeforeThreatProtection returns the EnableAccelRespBeforeThreatProtection field if non-nil, zero value otherwise.

### GetEnableAccelRespBeforeThreatProtectionOk

`func (o *GridThreatprotection) GetEnableAccelRespBeforeThreatProtectionOk() (*bool, bool)`

GetEnableAccelRespBeforeThreatProtectionOk returns a tuple with the EnableAccelRespBeforeThreatProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccelRespBeforeThreatProtection

`func (o *GridThreatprotection) SetEnableAccelRespBeforeThreatProtection(v bool)`

SetEnableAccelRespBeforeThreatProtection sets EnableAccelRespBeforeThreatProtection field to given value.

### HasEnableAccelRespBeforeThreatProtection

`func (o *GridThreatprotection) HasEnableAccelRespBeforeThreatProtection() bool`

HasEnableAccelRespBeforeThreatProtection returns a boolean if a field has been set.

### GetEnableAutoDownload

`func (o *GridThreatprotection) GetEnableAutoDownload() bool`

GetEnableAutoDownload returns the EnableAutoDownload field if non-nil, zero value otherwise.

### GetEnableAutoDownloadOk

`func (o *GridThreatprotection) GetEnableAutoDownloadOk() (*bool, bool)`

GetEnableAutoDownloadOk returns a tuple with the EnableAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoDownload

`func (o *GridThreatprotection) SetEnableAutoDownload(v bool)`

SetEnableAutoDownload sets EnableAutoDownload field to given value.

### HasEnableAutoDownload

`func (o *GridThreatprotection) HasEnableAutoDownload() bool`

HasEnableAutoDownload returns a boolean if a field has been set.

### GetEnableNatRules

`func (o *GridThreatprotection) GetEnableNatRules() bool`

GetEnableNatRules returns the EnableNatRules field if non-nil, zero value otherwise.

### GetEnableNatRulesOk

`func (o *GridThreatprotection) GetEnableNatRulesOk() (*bool, bool)`

GetEnableNatRulesOk returns a tuple with the EnableNatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNatRules

`func (o *GridThreatprotection) SetEnableNatRules(v bool)`

SetEnableNatRules sets EnableNatRules field to given value.

### HasEnableNatRules

`func (o *GridThreatprotection) HasEnableNatRules() bool`

HasEnableNatRules returns a boolean if a field has been set.

### GetEnableScheduledDownload

`func (o *GridThreatprotection) GetEnableScheduledDownload() bool`

GetEnableScheduledDownload returns the EnableScheduledDownload field if non-nil, zero value otherwise.

### GetEnableScheduledDownloadOk

`func (o *GridThreatprotection) GetEnableScheduledDownloadOk() (*bool, bool)`

GetEnableScheduledDownloadOk returns a tuple with the EnableScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScheduledDownload

`func (o *GridThreatprotection) SetEnableScheduledDownload(v bool)`

SetEnableScheduledDownload sets EnableScheduledDownload field to given value.

### HasEnableScheduledDownload

`func (o *GridThreatprotection) HasEnableScheduledDownload() bool`

HasEnableScheduledDownload returns a boolean if a field has been set.

### GetEventsPerSecondPerRule

`func (o *GridThreatprotection) GetEventsPerSecondPerRule() int64`

GetEventsPerSecondPerRule returns the EventsPerSecondPerRule field if non-nil, zero value otherwise.

### GetEventsPerSecondPerRuleOk

`func (o *GridThreatprotection) GetEventsPerSecondPerRuleOk() (*int64, bool)`

GetEventsPerSecondPerRuleOk returns a tuple with the EventsPerSecondPerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsPerSecondPerRule

`func (o *GridThreatprotection) SetEventsPerSecondPerRule(v int64)`

SetEventsPerSecondPerRule sets EventsPerSecondPerRule field to given value.

### HasEventsPerSecondPerRule

`func (o *GridThreatprotection) HasEventsPerSecondPerRule() bool`

HasEventsPerSecondPerRule returns a boolean if a field has been set.

### GetGridName

`func (o *GridThreatprotection) GetGridName() string`

GetGridName returns the GridName field if non-nil, zero value otherwise.

### GetGridNameOk

`func (o *GridThreatprotection) GetGridNameOk() (*string, bool)`

GetGridNameOk returns a tuple with the GridName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridName

`func (o *GridThreatprotection) SetGridName(v string)`

SetGridName sets GridName field to given value.

### HasGridName

`func (o *GridThreatprotection) HasGridName() bool`

HasGridName returns a boolean if a field has been set.

### GetLastCheckedForUpdate

`func (o *GridThreatprotection) GetLastCheckedForUpdate() int64`

GetLastCheckedForUpdate returns the LastCheckedForUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForUpdateOk

`func (o *GridThreatprotection) GetLastCheckedForUpdateOk() (*int64, bool)`

GetLastCheckedForUpdateOk returns a tuple with the LastCheckedForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForUpdate

`func (o *GridThreatprotection) SetLastCheckedForUpdate(v int64)`

SetLastCheckedForUpdate sets LastCheckedForUpdate field to given value.

### HasLastCheckedForUpdate

`func (o *GridThreatprotection) HasLastCheckedForUpdate() bool`

HasLastCheckedForUpdate returns a boolean if a field has been set.

### GetLastRuleUpdateTimestamp

`func (o *GridThreatprotection) GetLastRuleUpdateTimestamp() int64`

GetLastRuleUpdateTimestamp returns the LastRuleUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastRuleUpdateTimestampOk

`func (o *GridThreatprotection) GetLastRuleUpdateTimestampOk() (*int64, bool)`

GetLastRuleUpdateTimestampOk returns a tuple with the LastRuleUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRuleUpdateTimestamp

`func (o *GridThreatprotection) SetLastRuleUpdateTimestamp(v int64)`

SetLastRuleUpdateTimestamp sets LastRuleUpdateTimestamp field to given value.

### HasLastRuleUpdateTimestamp

`func (o *GridThreatprotection) HasLastRuleUpdateTimestamp() bool`

HasLastRuleUpdateTimestamp returns a boolean if a field has been set.

### GetLastRuleUpdateVersion

`func (o *GridThreatprotection) GetLastRuleUpdateVersion() string`

GetLastRuleUpdateVersion returns the LastRuleUpdateVersion field if non-nil, zero value otherwise.

### GetLastRuleUpdateVersionOk

`func (o *GridThreatprotection) GetLastRuleUpdateVersionOk() (*string, bool)`

GetLastRuleUpdateVersionOk returns a tuple with the LastRuleUpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRuleUpdateVersion

`func (o *GridThreatprotection) SetLastRuleUpdateVersion(v string)`

SetLastRuleUpdateVersion sets LastRuleUpdateVersion field to given value.

### HasLastRuleUpdateVersion

`func (o *GridThreatprotection) HasLastRuleUpdateVersion() bool`

HasLastRuleUpdateVersion returns a boolean if a field has been set.

### GetNatRules

`func (o *GridThreatprotection) GetNatRules() []GridThreatprotectionNatRules`

GetNatRules returns the NatRules field if non-nil, zero value otherwise.

### GetNatRulesOk

`func (o *GridThreatprotection) GetNatRulesOk() (*[]GridThreatprotectionNatRules, bool)`

GetNatRulesOk returns a tuple with the NatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatRules

`func (o *GridThreatprotection) SetNatRules(v []GridThreatprotectionNatRules)`

SetNatRules sets NatRules field to given value.

### HasNatRules

`func (o *GridThreatprotection) HasNatRules() bool`

HasNatRules returns a boolean if a field has been set.

### GetOutboundSettings

`func (o *GridThreatprotection) GetOutboundSettings() GridThreatprotectionOutboundSettings`

GetOutboundSettings returns the OutboundSettings field if non-nil, zero value otherwise.

### GetOutboundSettingsOk

`func (o *GridThreatprotection) GetOutboundSettingsOk() (*GridThreatprotectionOutboundSettings, bool)`

GetOutboundSettingsOk returns a tuple with the OutboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSettings

`func (o *GridThreatprotection) SetOutboundSettings(v GridThreatprotectionOutboundSettings)`

SetOutboundSettings sets OutboundSettings field to given value.

### HasOutboundSettings

`func (o *GridThreatprotection) HasOutboundSettings() bool`

HasOutboundSettings returns a boolean if a field has been set.

### GetRuleUpdatePolicy

`func (o *GridThreatprotection) GetRuleUpdatePolicy() string`

GetRuleUpdatePolicy returns the RuleUpdatePolicy field if non-nil, zero value otherwise.

### GetRuleUpdatePolicyOk

`func (o *GridThreatprotection) GetRuleUpdatePolicyOk() (*string, bool)`

GetRuleUpdatePolicyOk returns a tuple with the RuleUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleUpdatePolicy

`func (o *GridThreatprotection) SetRuleUpdatePolicy(v string)`

SetRuleUpdatePolicy sets RuleUpdatePolicy field to given value.

### HasRuleUpdatePolicy

`func (o *GridThreatprotection) HasRuleUpdatePolicy() bool`

HasRuleUpdatePolicy returns a boolean if a field has been set.

### GetScheduledDownload

`func (o *GridThreatprotection) GetScheduledDownload() GridThreatprotectionScheduledDownload`

GetScheduledDownload returns the ScheduledDownload field if non-nil, zero value otherwise.

### GetScheduledDownloadOk

`func (o *GridThreatprotection) GetScheduledDownloadOk() (*GridThreatprotectionScheduledDownload, bool)`

GetScheduledDownloadOk returns a tuple with the ScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDownload

`func (o *GridThreatprotection) SetScheduledDownload(v GridThreatprotectionScheduledDownload)`

SetScheduledDownload sets ScheduledDownload field to given value.

### HasScheduledDownload

`func (o *GridThreatprotection) HasScheduledDownload() bool`

HasScheduledDownload returns a boolean if a field has been set.

### GetUuid

`func (o *GridThreatprotection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GridThreatprotection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GridThreatprotection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GridThreatprotection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


