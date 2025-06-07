# RangeDiscoveryBasicPollSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortScanning** | Pointer to **bool** | Determines whether port scanning is enabled or not. | [optional] 
**DeviceProfile** | Pointer to **bool** | Determines whether device profile is enabled or not. | [optional] 
**SnmpCollection** | Pointer to **bool** | Determines whether SNMP collection is enabled or not. | [optional] 
**CliCollection** | Pointer to **bool** | Determines whether CLI collection is enabled or not. | [optional] 
**NetbiosScanning** | Pointer to **bool** | Determines whether netbios scanning is enabled or not. | [optional] 
**CompletePingSweep** | Pointer to **bool** | Determines whether complete ping sweep is enabled or not. | [optional] 
**SmartSubnetPingSweep** | Pointer to **bool** | Determines whether smart subnet ping sweep is enabled or not. | [optional] 
**AutoArpRefreshBeforeSwitchPortPolling** | Pointer to **bool** | Determines whether auto ARP refresh before switch port polling is enabled or not. | [optional] 
**SwitchPortDataCollectionPolling** | Pointer to **string** | A switch port data collection polling mode. | [optional] 
**SwitchPortDataCollectionPollingSchedule** | Pointer to [**RangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule**](RangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule.md) |  | [optional] 
**SwitchPortDataCollectionPollingInterval** | Pointer to **int64** | Indicates the interval for switch port data collection polling. | [optional] 
**CredentialGroup** | Pointer to **string** | Credential group. | [optional] 
**PollingFrequencyModifier** | Pointer to **string** | Polling Frequency Modifier. | [optional] 
**UseGlobalPollingFrequencyModifier** | Pointer to **bool** | Use Global Polling Frequency Modifier. | [optional] 

## Methods

### NewRangeDiscoveryBasicPollSettings

`func NewRangeDiscoveryBasicPollSettings() *RangeDiscoveryBasicPollSettings`

NewRangeDiscoveryBasicPollSettings instantiates a new RangeDiscoveryBasicPollSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeDiscoveryBasicPollSettingsWithDefaults

`func NewRangeDiscoveryBasicPollSettingsWithDefaults() *RangeDiscoveryBasicPollSettings`

NewRangeDiscoveryBasicPollSettingsWithDefaults instantiates a new RangeDiscoveryBasicPollSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortScanning

`func (o *RangeDiscoveryBasicPollSettings) GetPortScanning() bool`

GetPortScanning returns the PortScanning field if non-nil, zero value otherwise.

### GetPortScanningOk

`func (o *RangeDiscoveryBasicPollSettings) GetPortScanningOk() (*bool, bool)`

GetPortScanningOk returns a tuple with the PortScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanning

`func (o *RangeDiscoveryBasicPollSettings) SetPortScanning(v bool)`

SetPortScanning sets PortScanning field to given value.

### HasPortScanning

`func (o *RangeDiscoveryBasicPollSettings) HasPortScanning() bool`

HasPortScanning returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *RangeDiscoveryBasicPollSettings) GetDeviceProfile() bool`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *RangeDiscoveryBasicPollSettings) GetDeviceProfileOk() (*bool, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *RangeDiscoveryBasicPollSettings) SetDeviceProfile(v bool)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *RangeDiscoveryBasicPollSettings) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetSnmpCollection

`func (o *RangeDiscoveryBasicPollSettings) GetSnmpCollection() bool`

GetSnmpCollection returns the SnmpCollection field if non-nil, zero value otherwise.

### GetSnmpCollectionOk

`func (o *RangeDiscoveryBasicPollSettings) GetSnmpCollectionOk() (*bool, bool)`

GetSnmpCollectionOk returns a tuple with the SnmpCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollection

`func (o *RangeDiscoveryBasicPollSettings) SetSnmpCollection(v bool)`

SetSnmpCollection sets SnmpCollection field to given value.

### HasSnmpCollection

`func (o *RangeDiscoveryBasicPollSettings) HasSnmpCollection() bool`

HasSnmpCollection returns a boolean if a field has been set.

### GetCliCollection

`func (o *RangeDiscoveryBasicPollSettings) GetCliCollection() bool`

GetCliCollection returns the CliCollection field if non-nil, zero value otherwise.

### GetCliCollectionOk

`func (o *RangeDiscoveryBasicPollSettings) GetCliCollectionOk() (*bool, bool)`

GetCliCollectionOk returns a tuple with the CliCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollection

`func (o *RangeDiscoveryBasicPollSettings) SetCliCollection(v bool)`

SetCliCollection sets CliCollection field to given value.

### HasCliCollection

`func (o *RangeDiscoveryBasicPollSettings) HasCliCollection() bool`

HasCliCollection returns a boolean if a field has been set.

### GetNetbiosScanning

`func (o *RangeDiscoveryBasicPollSettings) GetNetbiosScanning() bool`

GetNetbiosScanning returns the NetbiosScanning field if non-nil, zero value otherwise.

### GetNetbiosScanningOk

`func (o *RangeDiscoveryBasicPollSettings) GetNetbiosScanningOk() (*bool, bool)`

GetNetbiosScanningOk returns a tuple with the NetbiosScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosScanning

`func (o *RangeDiscoveryBasicPollSettings) SetNetbiosScanning(v bool)`

SetNetbiosScanning sets NetbiosScanning field to given value.

### HasNetbiosScanning

`func (o *RangeDiscoveryBasicPollSettings) HasNetbiosScanning() bool`

HasNetbiosScanning returns a boolean if a field has been set.

### GetCompletePingSweep

`func (o *RangeDiscoveryBasicPollSettings) GetCompletePingSweep() bool`

GetCompletePingSweep returns the CompletePingSweep field if non-nil, zero value otherwise.

### GetCompletePingSweepOk

`func (o *RangeDiscoveryBasicPollSettings) GetCompletePingSweepOk() (*bool, bool)`

GetCompletePingSweepOk returns a tuple with the CompletePingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePingSweep

`func (o *RangeDiscoveryBasicPollSettings) SetCompletePingSweep(v bool)`

SetCompletePingSweep sets CompletePingSweep field to given value.

### HasCompletePingSweep

`func (o *RangeDiscoveryBasicPollSettings) HasCompletePingSweep() bool`

HasCompletePingSweep returns a boolean if a field has been set.

### GetSmartSubnetPingSweep

`func (o *RangeDiscoveryBasicPollSettings) GetSmartSubnetPingSweep() bool`

GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field if non-nil, zero value otherwise.

### GetSmartSubnetPingSweepOk

`func (o *RangeDiscoveryBasicPollSettings) GetSmartSubnetPingSweepOk() (*bool, bool)`

GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSubnetPingSweep

`func (o *RangeDiscoveryBasicPollSettings) SetSmartSubnetPingSweep(v bool)`

SetSmartSubnetPingSweep sets SmartSubnetPingSweep field to given value.

### HasSmartSubnetPingSweep

`func (o *RangeDiscoveryBasicPollSettings) HasSmartSubnetPingSweep() bool`

HasSmartSubnetPingSweep returns a boolean if a field has been set.

### GetAutoArpRefreshBeforeSwitchPortPolling

`func (o *RangeDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPolling() bool`

GetAutoArpRefreshBeforeSwitchPortPolling returns the AutoArpRefreshBeforeSwitchPortPolling field if non-nil, zero value otherwise.

### GetAutoArpRefreshBeforeSwitchPortPollingOk

`func (o *RangeDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPollingOk() (*bool, bool)`

GetAutoArpRefreshBeforeSwitchPortPollingOk returns a tuple with the AutoArpRefreshBeforeSwitchPortPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArpRefreshBeforeSwitchPortPolling

`func (o *RangeDiscoveryBasicPollSettings) SetAutoArpRefreshBeforeSwitchPortPolling(v bool)`

SetAutoArpRefreshBeforeSwitchPortPolling sets AutoArpRefreshBeforeSwitchPortPolling field to given value.

### HasAutoArpRefreshBeforeSwitchPortPolling

`func (o *RangeDiscoveryBasicPollSettings) HasAutoArpRefreshBeforeSwitchPortPolling() bool`

HasAutoArpRefreshBeforeSwitchPortPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPolling

`func (o *RangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPolling() string`

GetSwitchPortDataCollectionPolling returns the SwitchPortDataCollectionPolling field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingOk

`func (o *RangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingOk() (*string, bool)`

GetSwitchPortDataCollectionPollingOk returns a tuple with the SwitchPortDataCollectionPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPolling

`func (o *RangeDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPolling(v string)`

SetSwitchPortDataCollectionPolling sets SwitchPortDataCollectionPolling field to given value.

### HasSwitchPortDataCollectionPolling

`func (o *RangeDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPolling() bool`

HasSwitchPortDataCollectionPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingSchedule

`func (o *RangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingSchedule() RangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

GetSwitchPortDataCollectionPollingSchedule returns the SwitchPortDataCollectionPollingSchedule field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingScheduleOk

`func (o *RangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingScheduleOk() (*RangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, bool)`

GetSwitchPortDataCollectionPollingScheduleOk returns a tuple with the SwitchPortDataCollectionPollingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingSchedule

`func (o *RangeDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingSchedule(v RangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule)`

SetSwitchPortDataCollectionPollingSchedule sets SwitchPortDataCollectionPollingSchedule field to given value.

### HasSwitchPortDataCollectionPollingSchedule

`func (o *RangeDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingSchedule() bool`

HasSwitchPortDataCollectionPollingSchedule returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingInterval

`func (o *RangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingInterval() int64`

GetSwitchPortDataCollectionPollingInterval returns the SwitchPortDataCollectionPollingInterval field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingIntervalOk

`func (o *RangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingIntervalOk() (*int64, bool)`

GetSwitchPortDataCollectionPollingIntervalOk returns a tuple with the SwitchPortDataCollectionPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingInterval

`func (o *RangeDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingInterval(v int64)`

SetSwitchPortDataCollectionPollingInterval sets SwitchPortDataCollectionPollingInterval field to given value.

### HasSwitchPortDataCollectionPollingInterval

`func (o *RangeDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingInterval() bool`

HasSwitchPortDataCollectionPollingInterval returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *RangeDiscoveryBasicPollSettings) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *RangeDiscoveryBasicPollSettings) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *RangeDiscoveryBasicPollSettings) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *RangeDiscoveryBasicPollSettings) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.

### GetPollingFrequencyModifier

`func (o *RangeDiscoveryBasicPollSettings) GetPollingFrequencyModifier() string`

GetPollingFrequencyModifier returns the PollingFrequencyModifier field if non-nil, zero value otherwise.

### GetPollingFrequencyModifierOk

`func (o *RangeDiscoveryBasicPollSettings) GetPollingFrequencyModifierOk() (*string, bool)`

GetPollingFrequencyModifierOk returns a tuple with the PollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequencyModifier

`func (o *RangeDiscoveryBasicPollSettings) SetPollingFrequencyModifier(v string)`

SetPollingFrequencyModifier sets PollingFrequencyModifier field to given value.

### HasPollingFrequencyModifier

`func (o *RangeDiscoveryBasicPollSettings) HasPollingFrequencyModifier() bool`

HasPollingFrequencyModifier returns a boolean if a field has been set.

### GetUseGlobalPollingFrequencyModifier

`func (o *RangeDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifier() bool`

GetUseGlobalPollingFrequencyModifier returns the UseGlobalPollingFrequencyModifier field if non-nil, zero value otherwise.

### GetUseGlobalPollingFrequencyModifierOk

`func (o *RangeDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifierOk() (*bool, bool)`

GetUseGlobalPollingFrequencyModifierOk returns a tuple with the UseGlobalPollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalPollingFrequencyModifier

`func (o *RangeDiscoveryBasicPollSettings) SetUseGlobalPollingFrequencyModifier(v bool)`

SetUseGlobalPollingFrequencyModifier sets UseGlobalPollingFrequencyModifier field to given value.

### HasUseGlobalPollingFrequencyModifier

`func (o *RangeDiscoveryBasicPollSettings) HasUseGlobalPollingFrequencyModifier() bool`

HasUseGlobalPollingFrequencyModifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


