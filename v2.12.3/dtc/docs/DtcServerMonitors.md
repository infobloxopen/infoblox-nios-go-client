# DtcServerMonitors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to **string** | The monitor related to server. | [optional] 
**Host** | Pointer to **string** | IP address or FQDN of the server used for monitoring. | [optional] 

## Methods

### NewDtcServerMonitors

`func NewDtcServerMonitors() *DtcServerMonitors`

NewDtcServerMonitors instantiates a new DtcServerMonitors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcServerMonitorsWithDefaults

`func NewDtcServerMonitorsWithDefaults() *DtcServerMonitors`

NewDtcServerMonitorsWithDefaults instantiates a new DtcServerMonitors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *DtcServerMonitors) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *DtcServerMonitors) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *DtcServerMonitors) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *DtcServerMonitors) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetHost

`func (o *DtcServerMonitors) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DtcServerMonitors) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DtcServerMonitors) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DtcServerMonitors) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


