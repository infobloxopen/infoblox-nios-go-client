# NetworkDiscoveryBasicPollSettings

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
**SwitchPortDataCollectionPollingSchedule** | Pointer to [**NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule**](NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule.md) |  | [optional] 
**SwitchPortDataCollectionPollingInterval** | Pointer to **int64** | Indicates the interval for switch port data collection polling. | [optional] 
**CredentialGroup** | Pointer to **string** | Credential group. | [optional] 

## Methods

### NewNetworkDiscoveryBasicPollSettings

`func NewNetworkDiscoveryBasicPollSettings() *NetworkDiscoveryBasicPollSettings`

NewNetworkDiscoveryBasicPollSettings instantiates a new NetworkDiscoveryBasicPollSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDiscoveryBasicPollSettingsWithDefaults

`func NewNetworkDiscoveryBasicPollSettingsWithDefaults() *NetworkDiscoveryBasicPollSettings`

NewNetworkDiscoveryBasicPollSettingsWithDefaults instantiates a new NetworkDiscoveryBasicPollSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortScanning

`func (o *NetworkDiscoveryBasicPollSettings) GetPortScanning() bool`

GetPortScanning returns the PortScanning field if non-nil, zero value otherwise.

### GetPortScanningOk

`func (o *NetworkDiscoveryBasicPollSettings) GetPortScanningOk() (*bool, bool)`

GetPortScanningOk returns a tuple with the PortScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanning

`func (o *NetworkDiscoveryBasicPollSettings) SetPortScanning(v bool)`

SetPortScanning sets PortScanning field to given value.

### HasPortScanning

`func (o *NetworkDiscoveryBasicPollSettings) HasPortScanning() bool`

HasPortScanning returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *NetworkDiscoveryBasicPollSettings) GetDeviceProfile() bool`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *NetworkDiscoveryBasicPollSettings) GetDeviceProfileOk() (*bool, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *NetworkDiscoveryBasicPollSettings) SetDeviceProfile(v bool)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *NetworkDiscoveryBasicPollSettings) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetSnmpCollection

`func (o *NetworkDiscoveryBasicPollSettings) GetSnmpCollection() bool`

GetSnmpCollection returns the SnmpCollection field if non-nil, zero value otherwise.

### GetSnmpCollectionOk

`func (o *NetworkDiscoveryBasicPollSettings) GetSnmpCollectionOk() (*bool, bool)`

GetSnmpCollectionOk returns a tuple with the SnmpCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollection

`func (o *NetworkDiscoveryBasicPollSettings) SetSnmpCollection(v bool)`

SetSnmpCollection sets SnmpCollection field to given value.

### HasSnmpCollection

`func (o *NetworkDiscoveryBasicPollSettings) HasSnmpCollection() bool`

HasSnmpCollection returns a boolean if a field has been set.

### GetCliCollection

`func (o *NetworkDiscoveryBasicPollSettings) GetCliCollection() bool`

GetCliCollection returns the CliCollection field if non-nil, zero value otherwise.

### GetCliCollectionOk

`func (o *NetworkDiscoveryBasicPollSettings) GetCliCollectionOk() (*bool, bool)`

GetCliCollectionOk returns a tuple with the CliCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollection

`func (o *NetworkDiscoveryBasicPollSettings) SetCliCollection(v bool)`

SetCliCollection sets CliCollection field to given value.

### HasCliCollection

`func (o *NetworkDiscoveryBasicPollSettings) HasCliCollection() bool`

HasCliCollection returns a boolean if a field has been set.

### GetNetbiosScanning

`func (o *NetworkDiscoveryBasicPollSettings) GetNetbiosScanning() bool`

GetNetbiosScanning returns the NetbiosScanning field if non-nil, zero value otherwise.

### GetNetbiosScanningOk

`func (o *NetworkDiscoveryBasicPollSettings) GetNetbiosScanningOk() (*bool, bool)`

GetNetbiosScanningOk returns a tuple with the NetbiosScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosScanning

`func (o *NetworkDiscoveryBasicPollSettings) SetNetbiosScanning(v bool)`

SetNetbiosScanning sets NetbiosScanning field to given value.

### HasNetbiosScanning

`func (o *NetworkDiscoveryBasicPollSettings) HasNetbiosScanning() bool`

HasNetbiosScanning returns a boolean if a field has been set.

### GetCompletePingSweep

`func (o *NetworkDiscoveryBasicPollSettings) GetCompletePingSweep() bool`

GetCompletePingSweep returns the CompletePingSweep field if non-nil, zero value otherwise.

### GetCompletePingSweepOk

`func (o *NetworkDiscoveryBasicPollSettings) GetCompletePingSweepOk() (*bool, bool)`

GetCompletePingSweepOk returns a tuple with the CompletePingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePingSweep

`func (o *NetworkDiscoveryBasicPollSettings) SetCompletePingSweep(v bool)`

SetCompletePingSweep sets CompletePingSweep field to given value.

### HasCompletePingSweep

`func (o *NetworkDiscoveryBasicPollSettings) HasCompletePingSweep() bool`

HasCompletePingSweep returns a boolean if a field has been set.

### GetSmartSubnetPingSweep

`func (o *NetworkDiscoveryBasicPollSettings) GetSmartSubnetPingSweep() bool`

GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field if non-nil, zero value otherwise.

### GetSmartSubnetPingSweepOk

`func (o *NetworkDiscoveryBasicPollSettings) GetSmartSubnetPingSweepOk() (*bool, bool)`

GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSubnetPingSweep

`func (o *NetworkDiscoveryBasicPollSettings) SetSmartSubnetPingSweep(v bool)`

SetSmartSubnetPingSweep sets SmartSubnetPingSweep field to given value.

### HasSmartSubnetPingSweep

`func (o *NetworkDiscoveryBasicPollSettings) HasSmartSubnetPingSweep() bool`

HasSmartSubnetPingSweep returns a boolean if a field has been set.

### GetAutoArpRefreshBeforeSwitchPortPolling

`func (o *NetworkDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPolling() bool`

GetAutoArpRefreshBeforeSwitchPortPolling returns the AutoArpRefreshBeforeSwitchPortPolling field if non-nil, zero value otherwise.

### GetAutoArpRefreshBeforeSwitchPortPollingOk

`func (o *NetworkDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPollingOk() (*bool, bool)`

GetAutoArpRefreshBeforeSwitchPortPollingOk returns a tuple with the AutoArpRefreshBeforeSwitchPortPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArpRefreshBeforeSwitchPortPolling

`func (o *NetworkDiscoveryBasicPollSettings) SetAutoArpRefreshBeforeSwitchPortPolling(v bool)`

SetAutoArpRefreshBeforeSwitchPortPolling sets AutoArpRefreshBeforeSwitchPortPolling field to given value.

### HasAutoArpRefreshBeforeSwitchPortPolling

`func (o *NetworkDiscoveryBasicPollSettings) HasAutoArpRefreshBeforeSwitchPortPolling() bool`

HasAutoArpRefreshBeforeSwitchPortPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPolling

`func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPolling() string`

GetSwitchPortDataCollectionPolling returns the SwitchPortDataCollectionPolling field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingOk

`func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingOk() (*string, bool)`

GetSwitchPortDataCollectionPollingOk returns a tuple with the SwitchPortDataCollectionPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPolling

`func (o *NetworkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPolling(v string)`

SetSwitchPortDataCollectionPolling sets SwitchPortDataCollectionPolling field to given value.

### HasSwitchPortDataCollectionPolling

`func (o *NetworkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPolling() bool`

HasSwitchPortDataCollectionPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingSchedule

`func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingSchedule() NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

GetSwitchPortDataCollectionPollingSchedule returns the SwitchPortDataCollectionPollingSchedule field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingScheduleOk

`func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingScheduleOk() (*NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, bool)`

GetSwitchPortDataCollectionPollingScheduleOk returns a tuple with the SwitchPortDataCollectionPollingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingSchedule

`func (o *NetworkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingSchedule(v NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule)`

SetSwitchPortDataCollectionPollingSchedule sets SwitchPortDataCollectionPollingSchedule field to given value.

### HasSwitchPortDataCollectionPollingSchedule

`func (o *NetworkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingSchedule() bool`

HasSwitchPortDataCollectionPollingSchedule returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingInterval

`func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingInterval() int64`

GetSwitchPortDataCollectionPollingInterval returns the SwitchPortDataCollectionPollingInterval field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingIntervalOk

`func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingIntervalOk() (*int64, bool)`

GetSwitchPortDataCollectionPollingIntervalOk returns a tuple with the SwitchPortDataCollectionPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingInterval

`func (o *NetworkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingInterval(v int64)`

SetSwitchPortDataCollectionPollingInterval sets SwitchPortDataCollectionPollingInterval field to given value.

### HasSwitchPortDataCollectionPollingInterval

`func (o *NetworkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingInterval() bool`

HasSwitchPortDataCollectionPollingInterval returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *NetworkDiscoveryBasicPollSettings) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *NetworkDiscoveryBasicPollSettings) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *NetworkDiscoveryBasicPollSettings) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *NetworkDiscoveryBasicPollSettings) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


