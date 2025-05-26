# Discoverytask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CsvFileName** | Pointer to **string** | The network discovery CSV file name. | [optional] [readonly] 
**DisableIpScanning** | Pointer to **bool** | Determines whether IP scanning is disabled. | [optional] 
**DisableVmwareScanning** | Pointer to **bool** | Determines whether VMWare scanning is disabled. | [optional] 
**DiscoveryTaskOid** | Pointer to **string** | The discovery task identifier. | [optional] [readonly] 
**MemberName** | Pointer to **string** | The Grid member that runs the discovery. | [optional] 
**MergeData** | Pointer to **bool** | Determines whether to replace or merge new data with existing data. | [optional] 
**Mode** | Pointer to **string** | Network discovery scanning mode. | [optional] 
**NetworkDiscoveryControl** | Pointer to **map[string]interface{}** |  | [optional] 
**NetworkView** | Pointer to **string** | Name of the network view in which target networks for network discovery reside. | [optional] 
**Networks** | Pointer to **[]map[string]interface{}** | The list of the networks on which the network discovery will be invoked. | [optional] 
**PingRetries** | Pointer to **int64** | The number of times to perfrom ping for ICMP and FULL modes. | [optional] 
**PingTimeout** | Pointer to **int64** | The ping timeout for ICMP and FULL modes. | [optional] 
**ScheduledRun** | Pointer to [**DiscoverytaskScheduledRun**](DiscoverytaskScheduledRun.md) |  | [optional] 
**State** | Pointer to **string** | The network discovery process state. | [optional] [readonly] 
**StateTime** | Pointer to **int64** | Time when the network discovery process state was last updated. | [optional] [readonly] 
**Status** | Pointer to **string** | The network discovery process descriptive status. | [optional] [readonly] 
**StatusTime** | Pointer to **int64** | The time when the network discovery process status was last updated. | [optional] [readonly] 
**TcpPorts** | Pointer to [**[]DiscoverytaskTcpPorts**](DiscoverytaskTcpPorts.md) | The ports to scan for FULL and TCP modes. | [optional] 
**TcpScanTechnique** | Pointer to **string** | The TCP scan techinque for FULL and TCP modes. | [optional] 
**VNetworkView** | Pointer to **string** | Name of the network view in which target networks for VMWare scanning reside. | [optional] 
**Vservers** | Pointer to [**[]DiscoverytaskVservers**](DiscoverytaskVservers.md) | The list of VMware vSphere servers for VM discovery. | [optional] 
**Warning** | Pointer to **string** | The network discovery process warning. | [optional] [readonly] 

## Methods

### NewDiscoverytask

`func NewDiscoverytask() *Discoverytask`

NewDiscoverytask instantiates a new Discoverytask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoverytaskWithDefaults

`func NewDiscoverytaskWithDefaults() *Discoverytask`

NewDiscoverytaskWithDefaults instantiates a new Discoverytask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Discoverytask) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Discoverytask) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Discoverytask) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Discoverytask) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCsvFileName

`func (o *Discoverytask) GetCsvFileName() string`

GetCsvFileName returns the CsvFileName field if non-nil, zero value otherwise.

### GetCsvFileNameOk

`func (o *Discoverytask) GetCsvFileNameOk() (*string, bool)`

GetCsvFileNameOk returns a tuple with the CsvFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFileName

`func (o *Discoverytask) SetCsvFileName(v string)`

SetCsvFileName sets CsvFileName field to given value.

### HasCsvFileName

`func (o *Discoverytask) HasCsvFileName() bool`

HasCsvFileName returns a boolean if a field has been set.

### GetDisableIpScanning

`func (o *Discoverytask) GetDisableIpScanning() bool`

GetDisableIpScanning returns the DisableIpScanning field if non-nil, zero value otherwise.

### GetDisableIpScanningOk

`func (o *Discoverytask) GetDisableIpScanningOk() (*bool, bool)`

GetDisableIpScanningOk returns a tuple with the DisableIpScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIpScanning

`func (o *Discoverytask) SetDisableIpScanning(v bool)`

SetDisableIpScanning sets DisableIpScanning field to given value.

### HasDisableIpScanning

`func (o *Discoverytask) HasDisableIpScanning() bool`

HasDisableIpScanning returns a boolean if a field has been set.

### GetDisableVmwareScanning

`func (o *Discoverytask) GetDisableVmwareScanning() bool`

GetDisableVmwareScanning returns the DisableVmwareScanning field if non-nil, zero value otherwise.

### GetDisableVmwareScanningOk

`func (o *Discoverytask) GetDisableVmwareScanningOk() (*bool, bool)`

GetDisableVmwareScanningOk returns a tuple with the DisableVmwareScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVmwareScanning

`func (o *Discoverytask) SetDisableVmwareScanning(v bool)`

SetDisableVmwareScanning sets DisableVmwareScanning field to given value.

### HasDisableVmwareScanning

`func (o *Discoverytask) HasDisableVmwareScanning() bool`

HasDisableVmwareScanning returns a boolean if a field has been set.

### GetDiscoveryTaskOid

`func (o *Discoverytask) GetDiscoveryTaskOid() string`

GetDiscoveryTaskOid returns the DiscoveryTaskOid field if non-nil, zero value otherwise.

### GetDiscoveryTaskOidOk

`func (o *Discoverytask) GetDiscoveryTaskOidOk() (*string, bool)`

GetDiscoveryTaskOidOk returns a tuple with the DiscoveryTaskOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryTaskOid

`func (o *Discoverytask) SetDiscoveryTaskOid(v string)`

SetDiscoveryTaskOid sets DiscoveryTaskOid field to given value.

### HasDiscoveryTaskOid

`func (o *Discoverytask) HasDiscoveryTaskOid() bool`

HasDiscoveryTaskOid returns a boolean if a field has been set.

### GetMemberName

`func (o *Discoverytask) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *Discoverytask) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *Discoverytask) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *Discoverytask) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetMergeData

`func (o *Discoverytask) GetMergeData() bool`

GetMergeData returns the MergeData field if non-nil, zero value otherwise.

### GetMergeDataOk

`func (o *Discoverytask) GetMergeDataOk() (*bool, bool)`

GetMergeDataOk returns a tuple with the MergeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeData

`func (o *Discoverytask) SetMergeData(v bool)`

SetMergeData sets MergeData field to given value.

### HasMergeData

`func (o *Discoverytask) HasMergeData() bool`

HasMergeData returns a boolean if a field has been set.

### GetMode

`func (o *Discoverytask) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Discoverytask) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Discoverytask) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Discoverytask) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNetworkDiscoveryControl

`func (o *Discoverytask) GetNetworkDiscoveryControl() map[string]interface{}`

GetNetworkDiscoveryControl returns the NetworkDiscoveryControl field if non-nil, zero value otherwise.

### GetNetworkDiscoveryControlOk

`func (o *Discoverytask) GetNetworkDiscoveryControlOk() (*map[string]interface{}, bool)`

GetNetworkDiscoveryControlOk returns a tuple with the NetworkDiscoveryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDiscoveryControl

`func (o *Discoverytask) SetNetworkDiscoveryControl(v map[string]interface{})`

SetNetworkDiscoveryControl sets NetworkDiscoveryControl field to given value.

### HasNetworkDiscoveryControl

`func (o *Discoverytask) HasNetworkDiscoveryControl() bool`

HasNetworkDiscoveryControl returns a boolean if a field has been set.

### GetNetworkView

`func (o *Discoverytask) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Discoverytask) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Discoverytask) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Discoverytask) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworks

`func (o *Discoverytask) GetNetworks() []map[string]interface{}`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Discoverytask) GetNetworksOk() (*[]map[string]interface{}, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Discoverytask) SetNetworks(v []map[string]interface{})`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *Discoverytask) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPingRetries

`func (o *Discoverytask) GetPingRetries() int64`

GetPingRetries returns the PingRetries field if non-nil, zero value otherwise.

### GetPingRetriesOk

`func (o *Discoverytask) GetPingRetriesOk() (*int64, bool)`

GetPingRetriesOk returns a tuple with the PingRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRetries

`func (o *Discoverytask) SetPingRetries(v int64)`

SetPingRetries sets PingRetries field to given value.

### HasPingRetries

`func (o *Discoverytask) HasPingRetries() bool`

HasPingRetries returns a boolean if a field has been set.

### GetPingTimeout

`func (o *Discoverytask) GetPingTimeout() int64`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *Discoverytask) GetPingTimeoutOk() (*int64, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *Discoverytask) SetPingTimeout(v int64)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *Discoverytask) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetScheduledRun

`func (o *Discoverytask) GetScheduledRun() DiscoverytaskScheduledRun`

GetScheduledRun returns the ScheduledRun field if non-nil, zero value otherwise.

### GetScheduledRunOk

`func (o *Discoverytask) GetScheduledRunOk() (*DiscoverytaskScheduledRun, bool)`

GetScheduledRunOk returns a tuple with the ScheduledRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRun

`func (o *Discoverytask) SetScheduledRun(v DiscoverytaskScheduledRun)`

SetScheduledRun sets ScheduledRun field to given value.

### HasScheduledRun

`func (o *Discoverytask) HasScheduledRun() bool`

HasScheduledRun returns a boolean if a field has been set.

### GetState

`func (o *Discoverytask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Discoverytask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Discoverytask) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Discoverytask) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateTime

`func (o *Discoverytask) GetStateTime() int64`

GetStateTime returns the StateTime field if non-nil, zero value otherwise.

### GetStateTimeOk

`func (o *Discoverytask) GetStateTimeOk() (*int64, bool)`

GetStateTimeOk returns a tuple with the StateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTime

`func (o *Discoverytask) SetStateTime(v int64)`

SetStateTime sets StateTime field to given value.

### HasStateTime

`func (o *Discoverytask) HasStateTime() bool`

HasStateTime returns a boolean if a field has been set.

### GetStatus

`func (o *Discoverytask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Discoverytask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Discoverytask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Discoverytask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *Discoverytask) GetStatusTime() int64`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *Discoverytask) GetStatusTimeOk() (*int64, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *Discoverytask) SetStatusTime(v int64)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *Discoverytask) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetTcpPorts

`func (o *Discoverytask) GetTcpPorts() []DiscoverytaskTcpPorts`

GetTcpPorts returns the TcpPorts field if non-nil, zero value otherwise.

### GetTcpPortsOk

`func (o *Discoverytask) GetTcpPortsOk() (*[]DiscoverytaskTcpPorts, bool)`

GetTcpPortsOk returns a tuple with the TcpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPorts

`func (o *Discoverytask) SetTcpPorts(v []DiscoverytaskTcpPorts)`

SetTcpPorts sets TcpPorts field to given value.

### HasTcpPorts

`func (o *Discoverytask) HasTcpPorts() bool`

HasTcpPorts returns a boolean if a field has been set.

### GetTcpScanTechnique

`func (o *Discoverytask) GetTcpScanTechnique() string`

GetTcpScanTechnique returns the TcpScanTechnique field if non-nil, zero value otherwise.

### GetTcpScanTechniqueOk

`func (o *Discoverytask) GetTcpScanTechniqueOk() (*string, bool)`

GetTcpScanTechniqueOk returns a tuple with the TcpScanTechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanTechnique

`func (o *Discoverytask) SetTcpScanTechnique(v string)`

SetTcpScanTechnique sets TcpScanTechnique field to given value.

### HasTcpScanTechnique

`func (o *Discoverytask) HasTcpScanTechnique() bool`

HasTcpScanTechnique returns a boolean if a field has been set.

### GetVNetworkView

`func (o *Discoverytask) GetVNetworkView() string`

GetVNetworkView returns the VNetworkView field if non-nil, zero value otherwise.

### GetVNetworkViewOk

`func (o *Discoverytask) GetVNetworkViewOk() (*string, bool)`

GetVNetworkViewOk returns a tuple with the VNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetworkView

`func (o *Discoverytask) SetVNetworkView(v string)`

SetVNetworkView sets VNetworkView field to given value.

### HasVNetworkView

`func (o *Discoverytask) HasVNetworkView() bool`

HasVNetworkView returns a boolean if a field has been set.

### GetVservers

`func (o *Discoverytask) GetVservers() []DiscoverytaskVservers`

GetVservers returns the Vservers field if non-nil, zero value otherwise.

### GetVserversOk

`func (o *Discoverytask) GetVserversOk() (*[]DiscoverytaskVservers, bool)`

GetVserversOk returns a tuple with the Vservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVservers

`func (o *Discoverytask) SetVservers(v []DiscoverytaskVservers)`

SetVservers sets Vservers field to given value.

### HasVservers

`func (o *Discoverytask) HasVservers() bool`

HasVservers returns a boolean if a field has been set.

### GetWarning

`func (o *Discoverytask) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *Discoverytask) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *Discoverytask) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *Discoverytask) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


