# NetworkcontainerDiscoveryBasicPollSettings

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
**SwitchPortDataCollectionPollingSchedule** | Pointer to [**NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule**](NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule.md) |  | [optional] 
**SwitchPortDataCollectionPollingInterval** | Pointer to **int64** | Indicates the interval for switch port data collection polling. | [optional] 
**CredentialGroup** | Pointer to **string** | Credential group. | [optional] 

## Methods

### NewNetworkcontainerDiscoveryBasicPollSettings

`func NewNetworkcontainerDiscoveryBasicPollSettings() *NetworkcontainerDiscoveryBasicPollSettings`

NewNetworkcontainerDiscoveryBasicPollSettings instantiates a new NetworkcontainerDiscoveryBasicPollSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkcontainerDiscoveryBasicPollSettingsWithDefaults

`func NewNetworkcontainerDiscoveryBasicPollSettingsWithDefaults() *NetworkcontainerDiscoveryBasicPollSettings`

NewNetworkcontainerDiscoveryBasicPollSettingsWithDefaults instantiates a new NetworkcontainerDiscoveryBasicPollSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortScanning

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetPortScanning() bool`

GetPortScanning returns the PortScanning field if non-nil, zero value otherwise.

### GetPortScanningOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetPortScanningOk() (*bool, bool)`

GetPortScanningOk returns a tuple with the PortScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanning

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetPortScanning(v bool)`

SetPortScanning sets PortScanning field to given value.

### HasPortScanning

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasPortScanning() bool`

HasPortScanning returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetDeviceProfile() bool`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetDeviceProfileOk() (*bool, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetDeviceProfile(v bool)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetSnmpCollection

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSnmpCollection() bool`

GetSnmpCollection returns the SnmpCollection field if non-nil, zero value otherwise.

### GetSnmpCollectionOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSnmpCollectionOk() (*bool, bool)`

GetSnmpCollectionOk returns a tuple with the SnmpCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollection

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetSnmpCollection(v bool)`

SetSnmpCollection sets SnmpCollection field to given value.

### HasSnmpCollection

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasSnmpCollection() bool`

HasSnmpCollection returns a boolean if a field has been set.

### GetCliCollection

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetCliCollection() bool`

GetCliCollection returns the CliCollection field if non-nil, zero value otherwise.

### GetCliCollectionOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetCliCollectionOk() (*bool, bool)`

GetCliCollectionOk returns a tuple with the CliCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollection

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetCliCollection(v bool)`

SetCliCollection sets CliCollection field to given value.

### HasCliCollection

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasCliCollection() bool`

HasCliCollection returns a boolean if a field has been set.

### GetNetbiosScanning

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetNetbiosScanning() bool`

GetNetbiosScanning returns the NetbiosScanning field if non-nil, zero value otherwise.

### GetNetbiosScanningOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetNetbiosScanningOk() (*bool, bool)`

GetNetbiosScanningOk returns a tuple with the NetbiosScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosScanning

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetNetbiosScanning(v bool)`

SetNetbiosScanning sets NetbiosScanning field to given value.

### HasNetbiosScanning

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasNetbiosScanning() bool`

HasNetbiosScanning returns a boolean if a field has been set.

### GetCompletePingSweep

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetCompletePingSweep() bool`

GetCompletePingSweep returns the CompletePingSweep field if non-nil, zero value otherwise.

### GetCompletePingSweepOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetCompletePingSweepOk() (*bool, bool)`

GetCompletePingSweepOk returns a tuple with the CompletePingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePingSweep

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetCompletePingSweep(v bool)`

SetCompletePingSweep sets CompletePingSweep field to given value.

### HasCompletePingSweep

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasCompletePingSweep() bool`

HasCompletePingSweep returns a boolean if a field has been set.

### GetSmartSubnetPingSweep

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSmartSubnetPingSweep() bool`

GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field if non-nil, zero value otherwise.

### GetSmartSubnetPingSweepOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSmartSubnetPingSweepOk() (*bool, bool)`

GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSubnetPingSweep

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetSmartSubnetPingSweep(v bool)`

SetSmartSubnetPingSweep sets SmartSubnetPingSweep field to given value.

### HasSmartSubnetPingSweep

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasSmartSubnetPingSweep() bool`

HasSmartSubnetPingSweep returns a boolean if a field has been set.

### GetAutoArpRefreshBeforeSwitchPortPolling

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPolling() bool`

GetAutoArpRefreshBeforeSwitchPortPolling returns the AutoArpRefreshBeforeSwitchPortPolling field if non-nil, zero value otherwise.

### GetAutoArpRefreshBeforeSwitchPortPollingOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPollingOk() (*bool, bool)`

GetAutoArpRefreshBeforeSwitchPortPollingOk returns a tuple with the AutoArpRefreshBeforeSwitchPortPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArpRefreshBeforeSwitchPortPolling

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetAutoArpRefreshBeforeSwitchPortPolling(v bool)`

SetAutoArpRefreshBeforeSwitchPortPolling sets AutoArpRefreshBeforeSwitchPortPolling field to given value.

### HasAutoArpRefreshBeforeSwitchPortPolling

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasAutoArpRefreshBeforeSwitchPortPolling() bool`

HasAutoArpRefreshBeforeSwitchPortPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPolling

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPolling() string`

GetSwitchPortDataCollectionPolling returns the SwitchPortDataCollectionPolling field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingOk() (*string, bool)`

GetSwitchPortDataCollectionPollingOk returns a tuple with the SwitchPortDataCollectionPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPolling

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPolling(v string)`

SetSwitchPortDataCollectionPolling sets SwitchPortDataCollectionPolling field to given value.

### HasSwitchPortDataCollectionPolling

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPolling() bool`

HasSwitchPortDataCollectionPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingSchedule

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingSchedule() NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

GetSwitchPortDataCollectionPollingSchedule returns the SwitchPortDataCollectionPollingSchedule field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingScheduleOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingScheduleOk() (*NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, bool)`

GetSwitchPortDataCollectionPollingScheduleOk returns a tuple with the SwitchPortDataCollectionPollingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingSchedule

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingSchedule(v NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule)`

SetSwitchPortDataCollectionPollingSchedule sets SwitchPortDataCollectionPollingSchedule field to given value.

### HasSwitchPortDataCollectionPollingSchedule

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingSchedule() bool`

HasSwitchPortDataCollectionPollingSchedule returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingInterval

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingInterval() int64`

GetSwitchPortDataCollectionPollingInterval returns the SwitchPortDataCollectionPollingInterval field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingIntervalOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingIntervalOk() (*int64, bool)`

GetSwitchPortDataCollectionPollingIntervalOk returns a tuple with the SwitchPortDataCollectionPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingInterval

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingInterval(v int64)`

SetSwitchPortDataCollectionPollingInterval sets SwitchPortDataCollectionPollingInterval field to given value.

### HasSwitchPortDataCollectionPollingInterval

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingInterval() bool`

HasSwitchPortDataCollectionPollingInterval returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *NetworkcontainerDiscoveryBasicPollSettings) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *NetworkcontainerDiscoveryBasicPollSettings) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *NetworkcontainerDiscoveryBasicPollSettings) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


