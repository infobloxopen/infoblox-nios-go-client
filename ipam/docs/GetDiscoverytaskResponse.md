# GetDiscoverytaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CsvFileName** | Pointer to **string** | The network discovery CSV file name. | [optional] [readonly] 
**DisableIpScanning** | Pointer to **bool** | Determines whether IP scanning is disabled. | [optional] 
**DisableVmwareScanning** | Pointer to **bool** | Determines whether VMWare scanning is disabled. | [optional] 
**DiscoveryTaskOid** | Pointer to **string** | The discovery task identifier. | [optional] [readonly] 
**MemberName** | Pointer to **string** | The Grid member that runs the discovery. | [optional] 
**MergeData** | Pointer to **bool** | Determines whether to replace or merge new data with existing data. | [optional] 
**Mode** | Pointer to **string** | Network discovery scanning mode. | [optional] 
**NetworkView** | Pointer to **string** | Name of the network view in which target networks for network discovery reside. | [optional] 
**Networks** | Pointer to **[]string** | The list of the networks on which the network discovery will be invoked. | [optional] 
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
**Result** | Pointer to [**Discoverytask**](Discoverytask.md) |  | [optional] 

## Methods

### NewGetDiscoverytaskResponse

`func NewGetDiscoverytaskResponse() *GetDiscoverytaskResponse`

NewGetDiscoverytaskResponse instantiates a new GetDiscoverytaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoverytaskResponseWithDefaults

`func NewGetDiscoverytaskResponseWithDefaults() *GetDiscoverytaskResponse`

NewGetDiscoverytaskResponseWithDefaults instantiates a new GetDiscoverytaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoverytaskResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoverytaskResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoverytaskResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoverytaskResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCsvFileName

`func (o *GetDiscoverytaskResponse) GetCsvFileName() string`

GetCsvFileName returns the CsvFileName field if non-nil, zero value otherwise.

### GetCsvFileNameOk

`func (o *GetDiscoverytaskResponse) GetCsvFileNameOk() (*string, bool)`

GetCsvFileNameOk returns a tuple with the CsvFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFileName

`func (o *GetDiscoverytaskResponse) SetCsvFileName(v string)`

SetCsvFileName sets CsvFileName field to given value.

### HasCsvFileName

`func (o *GetDiscoverytaskResponse) HasCsvFileName() bool`

HasCsvFileName returns a boolean if a field has been set.

### GetDisableIpScanning

`func (o *GetDiscoverytaskResponse) GetDisableIpScanning() bool`

GetDisableIpScanning returns the DisableIpScanning field if non-nil, zero value otherwise.

### GetDisableIpScanningOk

`func (o *GetDiscoverytaskResponse) GetDisableIpScanningOk() (*bool, bool)`

GetDisableIpScanningOk returns a tuple with the DisableIpScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIpScanning

`func (o *GetDiscoverytaskResponse) SetDisableIpScanning(v bool)`

SetDisableIpScanning sets DisableIpScanning field to given value.

### HasDisableIpScanning

`func (o *GetDiscoverytaskResponse) HasDisableIpScanning() bool`

HasDisableIpScanning returns a boolean if a field has been set.

### GetDisableVmwareScanning

`func (o *GetDiscoverytaskResponse) GetDisableVmwareScanning() bool`

GetDisableVmwareScanning returns the DisableVmwareScanning field if non-nil, zero value otherwise.

### GetDisableVmwareScanningOk

`func (o *GetDiscoverytaskResponse) GetDisableVmwareScanningOk() (*bool, bool)`

GetDisableVmwareScanningOk returns a tuple with the DisableVmwareScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVmwareScanning

`func (o *GetDiscoverytaskResponse) SetDisableVmwareScanning(v bool)`

SetDisableVmwareScanning sets DisableVmwareScanning field to given value.

### HasDisableVmwareScanning

`func (o *GetDiscoverytaskResponse) HasDisableVmwareScanning() bool`

HasDisableVmwareScanning returns a boolean if a field has been set.

### GetDiscoveryTaskOid

`func (o *GetDiscoverytaskResponse) GetDiscoveryTaskOid() string`

GetDiscoveryTaskOid returns the DiscoveryTaskOid field if non-nil, zero value otherwise.

### GetDiscoveryTaskOidOk

`func (o *GetDiscoverytaskResponse) GetDiscoveryTaskOidOk() (*string, bool)`

GetDiscoveryTaskOidOk returns a tuple with the DiscoveryTaskOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryTaskOid

`func (o *GetDiscoverytaskResponse) SetDiscoveryTaskOid(v string)`

SetDiscoveryTaskOid sets DiscoveryTaskOid field to given value.

### HasDiscoveryTaskOid

`func (o *GetDiscoverytaskResponse) HasDiscoveryTaskOid() bool`

HasDiscoveryTaskOid returns a boolean if a field has been set.

### GetMemberName

`func (o *GetDiscoverytaskResponse) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *GetDiscoverytaskResponse) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *GetDiscoverytaskResponse) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *GetDiscoverytaskResponse) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetMergeData

`func (o *GetDiscoverytaskResponse) GetMergeData() bool`

GetMergeData returns the MergeData field if non-nil, zero value otherwise.

### GetMergeDataOk

`func (o *GetDiscoverytaskResponse) GetMergeDataOk() (*bool, bool)`

GetMergeDataOk returns a tuple with the MergeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeData

`func (o *GetDiscoverytaskResponse) SetMergeData(v bool)`

SetMergeData sets MergeData field to given value.

### HasMergeData

`func (o *GetDiscoverytaskResponse) HasMergeData() bool`

HasMergeData returns a boolean if a field has been set.

### GetMode

`func (o *GetDiscoverytaskResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetDiscoverytaskResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetDiscoverytaskResponse) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetDiscoverytaskResponse) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetDiscoverytaskResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetDiscoverytaskResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetDiscoverytaskResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetDiscoverytaskResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworks

`func (o *GetDiscoverytaskResponse) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetDiscoverytaskResponse) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetDiscoverytaskResponse) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetDiscoverytaskResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPingRetries

`func (o *GetDiscoverytaskResponse) GetPingRetries() int64`

GetPingRetries returns the PingRetries field if non-nil, zero value otherwise.

### GetPingRetriesOk

`func (o *GetDiscoverytaskResponse) GetPingRetriesOk() (*int64, bool)`

GetPingRetriesOk returns a tuple with the PingRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRetries

`func (o *GetDiscoverytaskResponse) SetPingRetries(v int64)`

SetPingRetries sets PingRetries field to given value.

### HasPingRetries

`func (o *GetDiscoverytaskResponse) HasPingRetries() bool`

HasPingRetries returns a boolean if a field has been set.

### GetPingTimeout

`func (o *GetDiscoverytaskResponse) GetPingTimeout() int64`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *GetDiscoverytaskResponse) GetPingTimeoutOk() (*int64, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *GetDiscoverytaskResponse) SetPingTimeout(v int64)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *GetDiscoverytaskResponse) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetScheduledRun

`func (o *GetDiscoverytaskResponse) GetScheduledRun() DiscoverytaskScheduledRun`

GetScheduledRun returns the ScheduledRun field if non-nil, zero value otherwise.

### GetScheduledRunOk

`func (o *GetDiscoverytaskResponse) GetScheduledRunOk() (*DiscoverytaskScheduledRun, bool)`

GetScheduledRunOk returns a tuple with the ScheduledRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRun

`func (o *GetDiscoverytaskResponse) SetScheduledRun(v DiscoverytaskScheduledRun)`

SetScheduledRun sets ScheduledRun field to given value.

### HasScheduledRun

`func (o *GetDiscoverytaskResponse) HasScheduledRun() bool`

HasScheduledRun returns a boolean if a field has been set.

### GetState

`func (o *GetDiscoverytaskResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetDiscoverytaskResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetDiscoverytaskResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetDiscoverytaskResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateTime

`func (o *GetDiscoverytaskResponse) GetStateTime() int64`

GetStateTime returns the StateTime field if non-nil, zero value otherwise.

### GetStateTimeOk

`func (o *GetDiscoverytaskResponse) GetStateTimeOk() (*int64, bool)`

GetStateTimeOk returns a tuple with the StateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTime

`func (o *GetDiscoverytaskResponse) SetStateTime(v int64)`

SetStateTime sets StateTime field to given value.

### HasStateTime

`func (o *GetDiscoverytaskResponse) HasStateTime() bool`

HasStateTime returns a boolean if a field has been set.

### GetStatus

`func (o *GetDiscoverytaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDiscoverytaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDiscoverytaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDiscoverytaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *GetDiscoverytaskResponse) GetStatusTime() int64`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *GetDiscoverytaskResponse) GetStatusTimeOk() (*int64, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *GetDiscoverytaskResponse) SetStatusTime(v int64)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *GetDiscoverytaskResponse) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetTcpPorts

`func (o *GetDiscoverytaskResponse) GetTcpPorts() []DiscoverytaskTcpPorts`

GetTcpPorts returns the TcpPorts field if non-nil, zero value otherwise.

### GetTcpPortsOk

`func (o *GetDiscoverytaskResponse) GetTcpPortsOk() (*[]DiscoverytaskTcpPorts, bool)`

GetTcpPortsOk returns a tuple with the TcpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPorts

`func (o *GetDiscoverytaskResponse) SetTcpPorts(v []DiscoverytaskTcpPorts)`

SetTcpPorts sets TcpPorts field to given value.

### HasTcpPorts

`func (o *GetDiscoverytaskResponse) HasTcpPorts() bool`

HasTcpPorts returns a boolean if a field has been set.

### GetTcpScanTechnique

`func (o *GetDiscoverytaskResponse) GetTcpScanTechnique() string`

GetTcpScanTechnique returns the TcpScanTechnique field if non-nil, zero value otherwise.

### GetTcpScanTechniqueOk

`func (o *GetDiscoverytaskResponse) GetTcpScanTechniqueOk() (*string, bool)`

GetTcpScanTechniqueOk returns a tuple with the TcpScanTechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanTechnique

`func (o *GetDiscoverytaskResponse) SetTcpScanTechnique(v string)`

SetTcpScanTechnique sets TcpScanTechnique field to given value.

### HasTcpScanTechnique

`func (o *GetDiscoverytaskResponse) HasTcpScanTechnique() bool`

HasTcpScanTechnique returns a boolean if a field has been set.

### GetVNetworkView

`func (o *GetDiscoverytaskResponse) GetVNetworkView() string`

GetVNetworkView returns the VNetworkView field if non-nil, zero value otherwise.

### GetVNetworkViewOk

`func (o *GetDiscoverytaskResponse) GetVNetworkViewOk() (*string, bool)`

GetVNetworkViewOk returns a tuple with the VNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetworkView

`func (o *GetDiscoverytaskResponse) SetVNetworkView(v string)`

SetVNetworkView sets VNetworkView field to given value.

### HasVNetworkView

`func (o *GetDiscoverytaskResponse) HasVNetworkView() bool`

HasVNetworkView returns a boolean if a field has been set.

### GetVservers

`func (o *GetDiscoverytaskResponse) GetVservers() []DiscoverytaskVservers`

GetVservers returns the Vservers field if non-nil, zero value otherwise.

### GetVserversOk

`func (o *GetDiscoverytaskResponse) GetVserversOk() (*[]DiscoverytaskVservers, bool)`

GetVserversOk returns a tuple with the Vservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVservers

`func (o *GetDiscoverytaskResponse) SetVservers(v []DiscoverytaskVservers)`

SetVservers sets Vservers field to given value.

### HasVservers

`func (o *GetDiscoverytaskResponse) HasVservers() bool`

HasVservers returns a boolean if a field has been set.

### GetWarning

`func (o *GetDiscoverytaskResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GetDiscoverytaskResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GetDiscoverytaskResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GetDiscoverytaskResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoverytaskResponse) GetResult() Discoverytask`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoverytaskResponse) GetResultOk() (*Discoverytask, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoverytaskResponse) SetResult(v Discoverytask)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoverytaskResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


