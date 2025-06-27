# DiscoveryGridproperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AdvancedPollingSettings** | Pointer to [**DiscoveryGridpropertiesAdvancedPollingSettings**](DiscoveryGridpropertiesAdvancedPollingSettings.md) |  | [optional] 
**AdvancedSdnPollingSettings** | Pointer to [**DiscoveryGridpropertiesAdvancedSdnPollingSettings**](DiscoveryGridpropertiesAdvancedSdnPollingSettings.md) |  | [optional] 
**AdvisorSettings** | Pointer to [**DiscoveryGridpropertiesAdvisorSettings**](DiscoveryGridpropertiesAdvisorSettings.md) |  | [optional] 
**AutoConversionSettings** | Pointer to [**[]DiscoveryGridpropertiesAutoConversionSettings**](DiscoveryGridpropertiesAutoConversionSettings.md) | Automatic conversion settings. | [optional] 
**BasicPollingSettings** | Pointer to [**DiscoveryGridpropertiesBasicPollingSettings**](DiscoveryGridpropertiesBasicPollingSettings.md) |  | [optional] 
**BasicSdnPollingSettings** | Pointer to [**DiscoveryGridpropertiesBasicSdnPollingSettings**](DiscoveryGridpropertiesBasicSdnPollingSettings.md) |  | [optional] 
**CliCredentials** | Pointer to [**[]DiscoveryGridpropertiesCliCredentials**](DiscoveryGridpropertiesCliCredentials.md) | Discovery CLI credentials. | [optional] 
**DeviceHints** | Pointer to [**[]DiscoveryGridpropertiesDeviceHints**](DiscoveryGridpropertiesDeviceHints.md) | Device Hints. | [optional] 
**DiscoveryBlackoutSetting** | Pointer to [**DiscoveryGridpropertiesDiscoveryBlackoutSetting**](DiscoveryGridpropertiesDiscoveryBlackoutSetting.md) |  | [optional] 
**DnsLookupOption** | Pointer to **string** | The type of the devices the DNS processor operates on. | [optional] 
**DnsLookupThrottle** | Pointer to **int64** | The percentage of available capacity the DNS processor operates at. Valid values are unsigned integer between 1 and 100, inclusive. | [optional] 
**EnableAdvisor** | Pointer to **bool** | Advisor application enabled/disabled. | [optional] 
**EnableAutoConversion** | Pointer to **bool** | The flag that enables automatic conversion of discovered data. | [optional] 
**EnableAutoUpdates** | Pointer to **bool** | The flag that enables updating discovered data for managed objects. | [optional] 
**GridName** | Pointer to **string** | The Grid name. | [optional] [readonly] 
**IgnoreConflictDuration** | Pointer to **int64** | Determines the timeout to ignore the discovery conflict duration (in seconds). | [optional] 
**PortControlBlackoutSetting** | Pointer to [**DiscoveryGridpropertiesPortControlBlackoutSetting**](DiscoveryGridpropertiesPortControlBlackoutSetting.md) |  | [optional] 
**Ports** | Pointer to [**[]DiscoveryGridpropertiesPorts**](DiscoveryGridpropertiesPorts.md) | Ports to scan. | [optional] 
**SamePortControlDiscoveryBlackout** | Pointer to **bool** | Determines if the same port control is used for discovery blackout. | [optional] 
**Snmpv1v2Credentials** | Pointer to [**[]DiscoveryGridpropertiesSnmpv1v2Credentials**](DiscoveryGridpropertiesSnmpv1v2Credentials.md) | Discovery SNMP v1 and v2 credentials. | [optional] 
**Snmpv3Credentials** | Pointer to [**[]DiscoveryGridpropertiesSnmpv3Credentials**](DiscoveryGridpropertiesSnmpv3Credentials.md) | Discovery SNMP v3 credentials. | [optional] 
**UnmanagedIpsLimit** | Pointer to **int64** | Limit of discovered unmanaged IP address which determines how frequently the user is notified about the new unmanaged IP address in a particular network. | [optional] 
**UnmanagedIpsTimeout** | Pointer to **int64** | Determines the timeout between two notifications (in seconds) about the new unmanaged IP address in a particular network. The value must be between 60 seconds and the number of seconds remaining to Jan 2038. | [optional] 
**VrfMappingPolicy** | Pointer to **string** | The policy type used to define the behavior of the VRF mapping. | [optional] 
**VrfMappingRules** | Pointer to [**[]DiscoveryGridpropertiesVrfMappingRules**](DiscoveryGridpropertiesVrfMappingRules.md) | VRF mapping rules. | [optional] 

## Methods

### NewDiscoveryGridproperties

`func NewDiscoveryGridproperties() *DiscoveryGridproperties`

NewDiscoveryGridproperties instantiates a new DiscoveryGridproperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesWithDefaults

`func NewDiscoveryGridpropertiesWithDefaults() *DiscoveryGridproperties`

NewDiscoveryGridpropertiesWithDefaults instantiates a new DiscoveryGridproperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryGridproperties) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryGridproperties) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryGridproperties) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryGridproperties) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdvancedPollingSettings

`func (o *DiscoveryGridproperties) GetAdvancedPollingSettings() DiscoveryGridpropertiesAdvancedPollingSettings`

GetAdvancedPollingSettings returns the AdvancedPollingSettings field if non-nil, zero value otherwise.

### GetAdvancedPollingSettingsOk

`func (o *DiscoveryGridproperties) GetAdvancedPollingSettingsOk() (*DiscoveryGridpropertiesAdvancedPollingSettings, bool)`

GetAdvancedPollingSettingsOk returns a tuple with the AdvancedPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPollingSettings

`func (o *DiscoveryGridproperties) SetAdvancedPollingSettings(v DiscoveryGridpropertiesAdvancedPollingSettings)`

SetAdvancedPollingSettings sets AdvancedPollingSettings field to given value.

### HasAdvancedPollingSettings

`func (o *DiscoveryGridproperties) HasAdvancedPollingSettings() bool`

HasAdvancedPollingSettings returns a boolean if a field has been set.

### GetAdvancedSdnPollingSettings

`func (o *DiscoveryGridproperties) GetAdvancedSdnPollingSettings() DiscoveryGridpropertiesAdvancedSdnPollingSettings`

GetAdvancedSdnPollingSettings returns the AdvancedSdnPollingSettings field if non-nil, zero value otherwise.

### GetAdvancedSdnPollingSettingsOk

`func (o *DiscoveryGridproperties) GetAdvancedSdnPollingSettingsOk() (*DiscoveryGridpropertiesAdvancedSdnPollingSettings, bool)`

GetAdvancedSdnPollingSettingsOk returns a tuple with the AdvancedSdnPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSdnPollingSettings

`func (o *DiscoveryGridproperties) SetAdvancedSdnPollingSettings(v DiscoveryGridpropertiesAdvancedSdnPollingSettings)`

SetAdvancedSdnPollingSettings sets AdvancedSdnPollingSettings field to given value.

### HasAdvancedSdnPollingSettings

`func (o *DiscoveryGridproperties) HasAdvancedSdnPollingSettings() bool`

HasAdvancedSdnPollingSettings returns a boolean if a field has been set.

### GetAdvisorSettings

`func (o *DiscoveryGridproperties) GetAdvisorSettings() DiscoveryGridpropertiesAdvisorSettings`

GetAdvisorSettings returns the AdvisorSettings field if non-nil, zero value otherwise.

### GetAdvisorSettingsOk

`func (o *DiscoveryGridproperties) GetAdvisorSettingsOk() (*DiscoveryGridpropertiesAdvisorSettings, bool)`

GetAdvisorSettingsOk returns a tuple with the AdvisorSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisorSettings

`func (o *DiscoveryGridproperties) SetAdvisorSettings(v DiscoveryGridpropertiesAdvisorSettings)`

SetAdvisorSettings sets AdvisorSettings field to given value.

### HasAdvisorSettings

`func (o *DiscoveryGridproperties) HasAdvisorSettings() bool`

HasAdvisorSettings returns a boolean if a field has been set.

### GetAutoConversionSettings

`func (o *DiscoveryGridproperties) GetAutoConversionSettings() []DiscoveryGridpropertiesAutoConversionSettings`

GetAutoConversionSettings returns the AutoConversionSettings field if non-nil, zero value otherwise.

### GetAutoConversionSettingsOk

`func (o *DiscoveryGridproperties) GetAutoConversionSettingsOk() (*[]DiscoveryGridpropertiesAutoConversionSettings, bool)`

GetAutoConversionSettingsOk returns a tuple with the AutoConversionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConversionSettings

`func (o *DiscoveryGridproperties) SetAutoConversionSettings(v []DiscoveryGridpropertiesAutoConversionSettings)`

SetAutoConversionSettings sets AutoConversionSettings field to given value.

### HasAutoConversionSettings

`func (o *DiscoveryGridproperties) HasAutoConversionSettings() bool`

HasAutoConversionSettings returns a boolean if a field has been set.

### GetBasicPollingSettings

`func (o *DiscoveryGridproperties) GetBasicPollingSettings() DiscoveryGridpropertiesBasicPollingSettings`

GetBasicPollingSettings returns the BasicPollingSettings field if non-nil, zero value otherwise.

### GetBasicPollingSettingsOk

`func (o *DiscoveryGridproperties) GetBasicPollingSettingsOk() (*DiscoveryGridpropertiesBasicPollingSettings, bool)`

GetBasicPollingSettingsOk returns a tuple with the BasicPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicPollingSettings

`func (o *DiscoveryGridproperties) SetBasicPollingSettings(v DiscoveryGridpropertiesBasicPollingSettings)`

SetBasicPollingSettings sets BasicPollingSettings field to given value.

### HasBasicPollingSettings

`func (o *DiscoveryGridproperties) HasBasicPollingSettings() bool`

HasBasicPollingSettings returns a boolean if a field has been set.

### GetBasicSdnPollingSettings

`func (o *DiscoveryGridproperties) GetBasicSdnPollingSettings() DiscoveryGridpropertiesBasicSdnPollingSettings`

GetBasicSdnPollingSettings returns the BasicSdnPollingSettings field if non-nil, zero value otherwise.

### GetBasicSdnPollingSettingsOk

`func (o *DiscoveryGridproperties) GetBasicSdnPollingSettingsOk() (*DiscoveryGridpropertiesBasicSdnPollingSettings, bool)`

GetBasicSdnPollingSettingsOk returns a tuple with the BasicSdnPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSdnPollingSettings

`func (o *DiscoveryGridproperties) SetBasicSdnPollingSettings(v DiscoveryGridpropertiesBasicSdnPollingSettings)`

SetBasicSdnPollingSettings sets BasicSdnPollingSettings field to given value.

### HasBasicSdnPollingSettings

`func (o *DiscoveryGridproperties) HasBasicSdnPollingSettings() bool`

HasBasicSdnPollingSettings returns a boolean if a field has been set.

### GetCliCredentials

`func (o *DiscoveryGridproperties) GetCliCredentials() []DiscoveryGridpropertiesCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *DiscoveryGridproperties) GetCliCredentialsOk() (*[]DiscoveryGridpropertiesCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *DiscoveryGridproperties) SetCliCredentials(v []DiscoveryGridpropertiesCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *DiscoveryGridproperties) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetDeviceHints

`func (o *DiscoveryGridproperties) GetDeviceHints() []DiscoveryGridpropertiesDeviceHints`

GetDeviceHints returns the DeviceHints field if non-nil, zero value otherwise.

### GetDeviceHintsOk

`func (o *DiscoveryGridproperties) GetDeviceHintsOk() (*[]DiscoveryGridpropertiesDeviceHints, bool)`

GetDeviceHintsOk returns a tuple with the DeviceHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHints

`func (o *DiscoveryGridproperties) SetDeviceHints(v []DiscoveryGridpropertiesDeviceHints)`

SetDeviceHints sets DeviceHints field to given value.

### HasDeviceHints

`func (o *DiscoveryGridproperties) HasDeviceHints() bool`

HasDeviceHints returns a boolean if a field has been set.

### GetDiscoveryBlackoutSetting

`func (o *DiscoveryGridproperties) GetDiscoveryBlackoutSetting() DiscoveryGridpropertiesDiscoveryBlackoutSetting`

GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field if non-nil, zero value otherwise.

### GetDiscoveryBlackoutSettingOk

`func (o *DiscoveryGridproperties) GetDiscoveryBlackoutSettingOk() (*DiscoveryGridpropertiesDiscoveryBlackoutSetting, bool)`

GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryBlackoutSetting

`func (o *DiscoveryGridproperties) SetDiscoveryBlackoutSetting(v DiscoveryGridpropertiesDiscoveryBlackoutSetting)`

SetDiscoveryBlackoutSetting sets DiscoveryBlackoutSetting field to given value.

### HasDiscoveryBlackoutSetting

`func (o *DiscoveryGridproperties) HasDiscoveryBlackoutSetting() bool`

HasDiscoveryBlackoutSetting returns a boolean if a field has been set.

### GetDnsLookupOption

`func (o *DiscoveryGridproperties) GetDnsLookupOption() string`

GetDnsLookupOption returns the DnsLookupOption field if non-nil, zero value otherwise.

### GetDnsLookupOptionOk

`func (o *DiscoveryGridproperties) GetDnsLookupOptionOk() (*string, bool)`

GetDnsLookupOptionOk returns a tuple with the DnsLookupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupOption

`func (o *DiscoveryGridproperties) SetDnsLookupOption(v string)`

SetDnsLookupOption sets DnsLookupOption field to given value.

### HasDnsLookupOption

`func (o *DiscoveryGridproperties) HasDnsLookupOption() bool`

HasDnsLookupOption returns a boolean if a field has been set.

### GetDnsLookupThrottle

`func (o *DiscoveryGridproperties) GetDnsLookupThrottle() int64`

GetDnsLookupThrottle returns the DnsLookupThrottle field if non-nil, zero value otherwise.

### GetDnsLookupThrottleOk

`func (o *DiscoveryGridproperties) GetDnsLookupThrottleOk() (*int64, bool)`

GetDnsLookupThrottleOk returns a tuple with the DnsLookupThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupThrottle

`func (o *DiscoveryGridproperties) SetDnsLookupThrottle(v int64)`

SetDnsLookupThrottle sets DnsLookupThrottle field to given value.

### HasDnsLookupThrottle

`func (o *DiscoveryGridproperties) HasDnsLookupThrottle() bool`

HasDnsLookupThrottle returns a boolean if a field has been set.

### GetEnableAdvisor

`func (o *DiscoveryGridproperties) GetEnableAdvisor() bool`

GetEnableAdvisor returns the EnableAdvisor field if non-nil, zero value otherwise.

### GetEnableAdvisorOk

`func (o *DiscoveryGridproperties) GetEnableAdvisorOk() (*bool, bool)`

GetEnableAdvisorOk returns a tuple with the EnableAdvisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdvisor

`func (o *DiscoveryGridproperties) SetEnableAdvisor(v bool)`

SetEnableAdvisor sets EnableAdvisor field to given value.

### HasEnableAdvisor

`func (o *DiscoveryGridproperties) HasEnableAdvisor() bool`

HasEnableAdvisor returns a boolean if a field has been set.

### GetEnableAutoConversion

`func (o *DiscoveryGridproperties) GetEnableAutoConversion() bool`

GetEnableAutoConversion returns the EnableAutoConversion field if non-nil, zero value otherwise.

### GetEnableAutoConversionOk

`func (o *DiscoveryGridproperties) GetEnableAutoConversionOk() (*bool, bool)`

GetEnableAutoConversionOk returns a tuple with the EnableAutoConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoConversion

`func (o *DiscoveryGridproperties) SetEnableAutoConversion(v bool)`

SetEnableAutoConversion sets EnableAutoConversion field to given value.

### HasEnableAutoConversion

`func (o *DiscoveryGridproperties) HasEnableAutoConversion() bool`

HasEnableAutoConversion returns a boolean if a field has been set.

### GetEnableAutoUpdates

`func (o *DiscoveryGridproperties) GetEnableAutoUpdates() bool`

GetEnableAutoUpdates returns the EnableAutoUpdates field if non-nil, zero value otherwise.

### GetEnableAutoUpdatesOk

`func (o *DiscoveryGridproperties) GetEnableAutoUpdatesOk() (*bool, bool)`

GetEnableAutoUpdatesOk returns a tuple with the EnableAutoUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoUpdates

`func (o *DiscoveryGridproperties) SetEnableAutoUpdates(v bool)`

SetEnableAutoUpdates sets EnableAutoUpdates field to given value.

### HasEnableAutoUpdates

`func (o *DiscoveryGridproperties) HasEnableAutoUpdates() bool`

HasEnableAutoUpdates returns a boolean if a field has been set.

### GetGridName

`func (o *DiscoveryGridproperties) GetGridName() string`

GetGridName returns the GridName field if non-nil, zero value otherwise.

### GetGridNameOk

`func (o *DiscoveryGridproperties) GetGridNameOk() (*string, bool)`

GetGridNameOk returns a tuple with the GridName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridName

`func (o *DiscoveryGridproperties) SetGridName(v string)`

SetGridName sets GridName field to given value.

### HasGridName

`func (o *DiscoveryGridproperties) HasGridName() bool`

HasGridName returns a boolean if a field has been set.

### GetIgnoreConflictDuration

`func (o *DiscoveryGridproperties) GetIgnoreConflictDuration() int64`

GetIgnoreConflictDuration returns the IgnoreConflictDuration field if non-nil, zero value otherwise.

### GetIgnoreConflictDurationOk

`func (o *DiscoveryGridproperties) GetIgnoreConflictDurationOk() (*int64, bool)`

GetIgnoreConflictDurationOk returns a tuple with the IgnoreConflictDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreConflictDuration

`func (o *DiscoveryGridproperties) SetIgnoreConflictDuration(v int64)`

SetIgnoreConflictDuration sets IgnoreConflictDuration field to given value.

### HasIgnoreConflictDuration

`func (o *DiscoveryGridproperties) HasIgnoreConflictDuration() bool`

HasIgnoreConflictDuration returns a boolean if a field has been set.

### GetPortControlBlackoutSetting

`func (o *DiscoveryGridproperties) GetPortControlBlackoutSetting() DiscoveryGridpropertiesPortControlBlackoutSetting`

GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field if non-nil, zero value otherwise.

### GetPortControlBlackoutSettingOk

`func (o *DiscoveryGridproperties) GetPortControlBlackoutSettingOk() (*DiscoveryGridpropertiesPortControlBlackoutSetting, bool)`

GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlBlackoutSetting

`func (o *DiscoveryGridproperties) SetPortControlBlackoutSetting(v DiscoveryGridpropertiesPortControlBlackoutSetting)`

SetPortControlBlackoutSetting sets PortControlBlackoutSetting field to given value.

### HasPortControlBlackoutSetting

`func (o *DiscoveryGridproperties) HasPortControlBlackoutSetting() bool`

HasPortControlBlackoutSetting returns a boolean if a field has been set.

### GetPorts

`func (o *DiscoveryGridproperties) GetPorts() []DiscoveryGridpropertiesPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DiscoveryGridproperties) GetPortsOk() (*[]DiscoveryGridpropertiesPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DiscoveryGridproperties) SetPorts(v []DiscoveryGridpropertiesPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DiscoveryGridproperties) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSamePortControlDiscoveryBlackout

`func (o *DiscoveryGridproperties) GetSamePortControlDiscoveryBlackout() bool`

GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field if non-nil, zero value otherwise.

### GetSamePortControlDiscoveryBlackoutOk

`func (o *DiscoveryGridproperties) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool)`

GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePortControlDiscoveryBlackout

`func (o *DiscoveryGridproperties) SetSamePortControlDiscoveryBlackout(v bool)`

SetSamePortControlDiscoveryBlackout sets SamePortControlDiscoveryBlackout field to given value.

### HasSamePortControlDiscoveryBlackout

`func (o *DiscoveryGridproperties) HasSamePortControlDiscoveryBlackout() bool`

HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.

### GetSnmpv1v2Credentials

`func (o *DiscoveryGridproperties) GetSnmpv1v2Credentials() []DiscoveryGridpropertiesSnmpv1v2Credentials`

GetSnmpv1v2Credentials returns the Snmpv1v2Credentials field if non-nil, zero value otherwise.

### GetSnmpv1v2CredentialsOk

`func (o *DiscoveryGridproperties) GetSnmpv1v2CredentialsOk() (*[]DiscoveryGridpropertiesSnmpv1v2Credentials, bool)`

GetSnmpv1v2CredentialsOk returns a tuple with the Snmpv1v2Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1v2Credentials

`func (o *DiscoveryGridproperties) SetSnmpv1v2Credentials(v []DiscoveryGridpropertiesSnmpv1v2Credentials)`

SetSnmpv1v2Credentials sets Snmpv1v2Credentials field to given value.

### HasSnmpv1v2Credentials

`func (o *DiscoveryGridproperties) HasSnmpv1v2Credentials() bool`

HasSnmpv1v2Credentials returns a boolean if a field has been set.

### GetSnmpv3Credentials

`func (o *DiscoveryGridproperties) GetSnmpv3Credentials() []DiscoveryGridpropertiesSnmpv3Credentials`

GetSnmpv3Credentials returns the Snmpv3Credentials field if non-nil, zero value otherwise.

### GetSnmpv3CredentialsOk

`func (o *DiscoveryGridproperties) GetSnmpv3CredentialsOk() (*[]DiscoveryGridpropertiesSnmpv3Credentials, bool)`

GetSnmpv3CredentialsOk returns a tuple with the Snmpv3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3Credentials

`func (o *DiscoveryGridproperties) SetSnmpv3Credentials(v []DiscoveryGridpropertiesSnmpv3Credentials)`

SetSnmpv3Credentials sets Snmpv3Credentials field to given value.

### HasSnmpv3Credentials

`func (o *DiscoveryGridproperties) HasSnmpv3Credentials() bool`

HasSnmpv3Credentials returns a boolean if a field has been set.

### GetUnmanagedIpsLimit

`func (o *DiscoveryGridproperties) GetUnmanagedIpsLimit() int64`

GetUnmanagedIpsLimit returns the UnmanagedIpsLimit field if non-nil, zero value otherwise.

### GetUnmanagedIpsLimitOk

`func (o *DiscoveryGridproperties) GetUnmanagedIpsLimitOk() (*int64, bool)`

GetUnmanagedIpsLimitOk returns a tuple with the UnmanagedIpsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedIpsLimit

`func (o *DiscoveryGridproperties) SetUnmanagedIpsLimit(v int64)`

SetUnmanagedIpsLimit sets UnmanagedIpsLimit field to given value.

### HasUnmanagedIpsLimit

`func (o *DiscoveryGridproperties) HasUnmanagedIpsLimit() bool`

HasUnmanagedIpsLimit returns a boolean if a field has been set.

### GetUnmanagedIpsTimeout

`func (o *DiscoveryGridproperties) GetUnmanagedIpsTimeout() int64`

GetUnmanagedIpsTimeout returns the UnmanagedIpsTimeout field if non-nil, zero value otherwise.

### GetUnmanagedIpsTimeoutOk

`func (o *DiscoveryGridproperties) GetUnmanagedIpsTimeoutOk() (*int64, bool)`

GetUnmanagedIpsTimeoutOk returns a tuple with the UnmanagedIpsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedIpsTimeout

`func (o *DiscoveryGridproperties) SetUnmanagedIpsTimeout(v int64)`

SetUnmanagedIpsTimeout sets UnmanagedIpsTimeout field to given value.

### HasUnmanagedIpsTimeout

`func (o *DiscoveryGridproperties) HasUnmanagedIpsTimeout() bool`

HasUnmanagedIpsTimeout returns a boolean if a field has been set.

### GetVrfMappingPolicy

`func (o *DiscoveryGridproperties) GetVrfMappingPolicy() string`

GetVrfMappingPolicy returns the VrfMappingPolicy field if non-nil, zero value otherwise.

### GetVrfMappingPolicyOk

`func (o *DiscoveryGridproperties) GetVrfMappingPolicyOk() (*string, bool)`

GetVrfMappingPolicyOk returns a tuple with the VrfMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfMappingPolicy

`func (o *DiscoveryGridproperties) SetVrfMappingPolicy(v string)`

SetVrfMappingPolicy sets VrfMappingPolicy field to given value.

### HasVrfMappingPolicy

`func (o *DiscoveryGridproperties) HasVrfMappingPolicy() bool`

HasVrfMappingPolicy returns a boolean if a field has been set.

### GetVrfMappingRules

`func (o *DiscoveryGridproperties) GetVrfMappingRules() []DiscoveryGridpropertiesVrfMappingRules`

GetVrfMappingRules returns the VrfMappingRules field if non-nil, zero value otherwise.

### GetVrfMappingRulesOk

`func (o *DiscoveryGridproperties) GetVrfMappingRulesOk() (*[]DiscoveryGridpropertiesVrfMappingRules, bool)`

GetVrfMappingRulesOk returns a tuple with the VrfMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfMappingRules

`func (o *DiscoveryGridproperties) SetVrfMappingRules(v []DiscoveryGridpropertiesVrfMappingRules)`

SetVrfMappingRules sets VrfMappingRules field to given value.

### HasVrfMappingRules

`func (o *DiscoveryGridproperties) HasVrfMappingRules() bool`

HasVrfMappingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


