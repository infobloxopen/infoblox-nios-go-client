# Ipv6rangeDiscoveryBasicPollSettings

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
**SwitchPortDataCollectionPollingSchedule** | Pointer to [**Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule**](Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule.md) |  | [optional] 
**SwitchPortDataCollectionPollingInterval** | Pointer to **int64** | Indicates the interval for switch port data collection polling. | [optional] 
**CredentialGroup** | Pointer to **string** | Credential group. | [optional] 
**PollingFrequencyModifier** | Pointer to **string** | Polling Frequency Modifier. | [optional] 
**UseGlobalPollingFrequencyModifier** | Pointer to **bool** | Use Global Polling Frequency Modifier. | [optional] 

## Methods

### NewIpv6rangeDiscoveryBasicPollSettings

`func NewIpv6rangeDiscoveryBasicPollSettings() *Ipv6rangeDiscoveryBasicPollSettings`

NewIpv6rangeDiscoveryBasicPollSettings instantiates a new Ipv6rangeDiscoveryBasicPollSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangeDiscoveryBasicPollSettingsWithDefaults

`func NewIpv6rangeDiscoveryBasicPollSettingsWithDefaults() *Ipv6rangeDiscoveryBasicPollSettings`

NewIpv6rangeDiscoveryBasicPollSettingsWithDefaults instantiates a new Ipv6rangeDiscoveryBasicPollSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortScanning

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetPortScanning() bool`

GetPortScanning returns the PortScanning field if non-nil, zero value otherwise.

### GetPortScanningOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetPortScanningOk() (*bool, bool)`

GetPortScanningOk returns a tuple with the PortScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanning

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetPortScanning(v bool)`

SetPortScanning sets PortScanning field to given value.

### HasPortScanning

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasPortScanning() bool`

HasPortScanning returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetDeviceProfile() bool`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetDeviceProfileOk() (*bool, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetDeviceProfile(v bool)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetSnmpCollection

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSnmpCollection() bool`

GetSnmpCollection returns the SnmpCollection field if non-nil, zero value otherwise.

### GetSnmpCollectionOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSnmpCollectionOk() (*bool, bool)`

GetSnmpCollectionOk returns a tuple with the SnmpCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollection

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetSnmpCollection(v bool)`

SetSnmpCollection sets SnmpCollection field to given value.

### HasSnmpCollection

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasSnmpCollection() bool`

HasSnmpCollection returns a boolean if a field has been set.

### GetCliCollection

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetCliCollection() bool`

GetCliCollection returns the CliCollection field if non-nil, zero value otherwise.

### GetCliCollectionOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetCliCollectionOk() (*bool, bool)`

GetCliCollectionOk returns a tuple with the CliCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollection

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetCliCollection(v bool)`

SetCliCollection sets CliCollection field to given value.

### HasCliCollection

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasCliCollection() bool`

HasCliCollection returns a boolean if a field has been set.

### GetNetbiosScanning

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetNetbiosScanning() bool`

GetNetbiosScanning returns the NetbiosScanning field if non-nil, zero value otherwise.

### GetNetbiosScanningOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetNetbiosScanningOk() (*bool, bool)`

GetNetbiosScanningOk returns a tuple with the NetbiosScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosScanning

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetNetbiosScanning(v bool)`

SetNetbiosScanning sets NetbiosScanning field to given value.

### HasNetbiosScanning

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasNetbiosScanning() bool`

HasNetbiosScanning returns a boolean if a field has been set.

### GetCompletePingSweep

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetCompletePingSweep() bool`

GetCompletePingSweep returns the CompletePingSweep field if non-nil, zero value otherwise.

### GetCompletePingSweepOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetCompletePingSweepOk() (*bool, bool)`

GetCompletePingSweepOk returns a tuple with the CompletePingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePingSweep

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetCompletePingSweep(v bool)`

SetCompletePingSweep sets CompletePingSweep field to given value.

### HasCompletePingSweep

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasCompletePingSweep() bool`

HasCompletePingSweep returns a boolean if a field has been set.

### GetSmartSubnetPingSweep

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSmartSubnetPingSweep() bool`

GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field if non-nil, zero value otherwise.

### GetSmartSubnetPingSweepOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSmartSubnetPingSweepOk() (*bool, bool)`

GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSubnetPingSweep

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetSmartSubnetPingSweep(v bool)`

SetSmartSubnetPingSweep sets SmartSubnetPingSweep field to given value.

### HasSmartSubnetPingSweep

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasSmartSubnetPingSweep() bool`

HasSmartSubnetPingSweep returns a boolean if a field has been set.

### GetAutoArpRefreshBeforeSwitchPortPolling

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPolling() bool`

GetAutoArpRefreshBeforeSwitchPortPolling returns the AutoArpRefreshBeforeSwitchPortPolling field if non-nil, zero value otherwise.

### GetAutoArpRefreshBeforeSwitchPortPollingOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPollingOk() (*bool, bool)`

GetAutoArpRefreshBeforeSwitchPortPollingOk returns a tuple with the AutoArpRefreshBeforeSwitchPortPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArpRefreshBeforeSwitchPortPolling

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetAutoArpRefreshBeforeSwitchPortPolling(v bool)`

SetAutoArpRefreshBeforeSwitchPortPolling sets AutoArpRefreshBeforeSwitchPortPolling field to given value.

### HasAutoArpRefreshBeforeSwitchPortPolling

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasAutoArpRefreshBeforeSwitchPortPolling() bool`

HasAutoArpRefreshBeforeSwitchPortPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPolling

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPolling() string`

GetSwitchPortDataCollectionPolling returns the SwitchPortDataCollectionPolling field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingOk() (*string, bool)`

GetSwitchPortDataCollectionPollingOk returns a tuple with the SwitchPortDataCollectionPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPolling

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPolling(v string)`

SetSwitchPortDataCollectionPolling sets SwitchPortDataCollectionPolling field to given value.

### HasSwitchPortDataCollectionPolling

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPolling() bool`

HasSwitchPortDataCollectionPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingSchedule

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingSchedule() Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

GetSwitchPortDataCollectionPollingSchedule returns the SwitchPortDataCollectionPollingSchedule field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingScheduleOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingScheduleOk() (*Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, bool)`

GetSwitchPortDataCollectionPollingScheduleOk returns a tuple with the SwitchPortDataCollectionPollingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingSchedule

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingSchedule(v Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule)`

SetSwitchPortDataCollectionPollingSchedule sets SwitchPortDataCollectionPollingSchedule field to given value.

### HasSwitchPortDataCollectionPollingSchedule

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingSchedule() bool`

HasSwitchPortDataCollectionPollingSchedule returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingInterval

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingInterval() int64`

GetSwitchPortDataCollectionPollingInterval returns the SwitchPortDataCollectionPollingInterval field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingIntervalOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingIntervalOk() (*int64, bool)`

GetSwitchPortDataCollectionPollingIntervalOk returns a tuple with the SwitchPortDataCollectionPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingInterval

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingInterval(v int64)`

SetSwitchPortDataCollectionPollingInterval sets SwitchPortDataCollectionPollingInterval field to given value.

### HasSwitchPortDataCollectionPollingInterval

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingInterval() bool`

HasSwitchPortDataCollectionPollingInterval returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.

### GetPollingFrequencyModifier

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetPollingFrequencyModifier() string`

GetPollingFrequencyModifier returns the PollingFrequencyModifier field if non-nil, zero value otherwise.

### GetPollingFrequencyModifierOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetPollingFrequencyModifierOk() (*string, bool)`

GetPollingFrequencyModifierOk returns a tuple with the PollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequencyModifier

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetPollingFrequencyModifier(v string)`

SetPollingFrequencyModifier sets PollingFrequencyModifier field to given value.

### HasPollingFrequencyModifier

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasPollingFrequencyModifier() bool`

HasPollingFrequencyModifier returns a boolean if a field has been set.

### GetUseGlobalPollingFrequencyModifier

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifier() bool`

GetUseGlobalPollingFrequencyModifier returns the UseGlobalPollingFrequencyModifier field if non-nil, zero value otherwise.

### GetUseGlobalPollingFrequencyModifierOk

`func (o *Ipv6rangeDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifierOk() (*bool, bool)`

GetUseGlobalPollingFrequencyModifierOk returns a tuple with the UseGlobalPollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalPollingFrequencyModifier

`func (o *Ipv6rangeDiscoveryBasicPollSettings) SetUseGlobalPollingFrequencyModifier(v bool)`

SetUseGlobalPollingFrequencyModifier sets UseGlobalPollingFrequencyModifier field to given value.

### HasUseGlobalPollingFrequencyModifier

`func (o *Ipv6rangeDiscoveryBasicPollSettings) HasUseGlobalPollingFrequencyModifier() bool`

HasUseGlobalPollingFrequencyModifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


