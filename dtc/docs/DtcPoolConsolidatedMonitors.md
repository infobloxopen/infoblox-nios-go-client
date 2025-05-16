# DtcPoolConsolidatedMonitors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** | Members whose monitor statuses are shared across other members in a pool. | [optional] 
**Monitor** | Pointer to **string** | Monitor whose statuses are shared across other members in a pool. | [optional] 
**Availability** | Pointer to **string** | Servers assigned to a pool with monitor defined are healthy if ANY or ALL members report healthy status. | [optional] 
**FullHealthCommunication** | Pointer to **bool** | Flag for switching health performing and sharing behavior to perform health checks on each DTC grid member that serves related LBDN(s) and send them across all DTC grid members from both selected and non-selected lists. | [optional] 

## Methods

### NewDtcPoolConsolidatedMonitors

`func NewDtcPoolConsolidatedMonitors() *DtcPoolConsolidatedMonitors`

NewDtcPoolConsolidatedMonitors instantiates a new DtcPoolConsolidatedMonitors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcPoolConsolidatedMonitorsWithDefaults

`func NewDtcPoolConsolidatedMonitorsWithDefaults() *DtcPoolConsolidatedMonitors`

NewDtcPoolConsolidatedMonitorsWithDefaults instantiates a new DtcPoolConsolidatedMonitors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *DtcPoolConsolidatedMonitors) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DtcPoolConsolidatedMonitors) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DtcPoolConsolidatedMonitors) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DtcPoolConsolidatedMonitors) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMonitor

`func (o *DtcPoolConsolidatedMonitors) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *DtcPoolConsolidatedMonitors) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *DtcPoolConsolidatedMonitors) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *DtcPoolConsolidatedMonitors) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetAvailability

`func (o *DtcPoolConsolidatedMonitors) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *DtcPoolConsolidatedMonitors) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *DtcPoolConsolidatedMonitors) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *DtcPoolConsolidatedMonitors) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetFullHealthCommunication

`func (o *DtcPoolConsolidatedMonitors) GetFullHealthCommunication() bool`

GetFullHealthCommunication returns the FullHealthCommunication field if non-nil, zero value otherwise.

### GetFullHealthCommunicationOk

`func (o *DtcPoolConsolidatedMonitors) GetFullHealthCommunicationOk() (*bool, bool)`

GetFullHealthCommunicationOk returns a tuple with the FullHealthCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullHealthCommunication

`func (o *DtcPoolConsolidatedMonitors) SetFullHealthCommunication(v bool)`

SetFullHealthCommunication sets FullHealthCommunication field to given value.

### HasFullHealthCommunication

`func (o *DtcPoolConsolidatedMonitors) HasFullHealthCommunication() bool`

HasFullHealthCommunication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


