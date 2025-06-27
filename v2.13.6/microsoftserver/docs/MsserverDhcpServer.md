# MsserverDhcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseLogin** | Pointer to **bool** | Flag to override login name and password from the MS Server | [optional] 
**LoginName** | Pointer to **string** | Microsoft Server login name | [optional] 
**LoginPassword** | Pointer to **string** | Microsoft Server login password | [optional] 
**Managed** | Pointer to **bool** | flag indicating if the DNS service is managed | [optional] 
**NextSyncControl** | Pointer to **string** | Defines what control to apply on the DNS server | [optional] 
**Status** | Pointer to **string** | Status of the Microsoft DNS Service | [optional] [readonly] 
**StatusLastUpdated** | Pointer to **int64** | Timestamp of the last update | [optional] [readonly] 
**UseEnableMonitoring** | Pointer to **bool** | Override enable monitoring inherited from grid level | [optional] 
**EnableMonitoring** | Pointer to **bool** | Flag indicating if the DNS service is monitored and controlled | [optional] 
**UseEnableInvalidMac** | Pointer to **bool** | Override setting for Enable Invalid Mac Address | [optional] 
**EnableInvalidMac** | Pointer to **bool** | Enable Invalid Mac Address | [optional] 
**SupportsFailover** | Pointer to **bool** | Flag indicating if the DHCP supports Failover | [optional] [readonly] 
**UseSynchronizationMinDelay** | Pointer to **bool** | Flag to override synchronization interval from the MS Server | [optional] 
**SynchronizationMinDelay** | Pointer to **int64** | Minimum number of minutes between two synchronizations | [optional] 

## Methods

### NewMsserverDhcpServer

`func NewMsserverDhcpServer() *MsserverDhcpServer`

NewMsserverDhcpServer instantiates a new MsserverDhcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverDhcpServerWithDefaults

`func NewMsserverDhcpServerWithDefaults() *MsserverDhcpServer`

NewMsserverDhcpServerWithDefaults instantiates a new MsserverDhcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseLogin

`func (o *MsserverDhcpServer) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *MsserverDhcpServer) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *MsserverDhcpServer) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *MsserverDhcpServer) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetLoginName

`func (o *MsserverDhcpServer) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MsserverDhcpServer) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MsserverDhcpServer) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MsserverDhcpServer) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *MsserverDhcpServer) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *MsserverDhcpServer) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *MsserverDhcpServer) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *MsserverDhcpServer) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetManaged

`func (o *MsserverDhcpServer) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MsserverDhcpServer) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MsserverDhcpServer) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MsserverDhcpServer) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNextSyncControl

`func (o *MsserverDhcpServer) GetNextSyncControl() string`

GetNextSyncControl returns the NextSyncControl field if non-nil, zero value otherwise.

### GetNextSyncControlOk

`func (o *MsserverDhcpServer) GetNextSyncControlOk() (*string, bool)`

GetNextSyncControlOk returns a tuple with the NextSyncControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSyncControl

`func (o *MsserverDhcpServer) SetNextSyncControl(v string)`

SetNextSyncControl sets NextSyncControl field to given value.

### HasNextSyncControl

`func (o *MsserverDhcpServer) HasNextSyncControl() bool`

HasNextSyncControl returns a boolean if a field has been set.

### GetStatus

`func (o *MsserverDhcpServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MsserverDhcpServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MsserverDhcpServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MsserverDhcpServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusLastUpdated

`func (o *MsserverDhcpServer) GetStatusLastUpdated() int64`

GetStatusLastUpdated returns the StatusLastUpdated field if non-nil, zero value otherwise.

### GetStatusLastUpdatedOk

`func (o *MsserverDhcpServer) GetStatusLastUpdatedOk() (*int64, bool)`

GetStatusLastUpdatedOk returns a tuple with the StatusLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastUpdated

`func (o *MsserverDhcpServer) SetStatusLastUpdated(v int64)`

SetStatusLastUpdated sets StatusLastUpdated field to given value.

### HasStatusLastUpdated

`func (o *MsserverDhcpServer) HasStatusLastUpdated() bool`

HasStatusLastUpdated returns a boolean if a field has been set.

### GetUseEnableMonitoring

`func (o *MsserverDhcpServer) GetUseEnableMonitoring() bool`

GetUseEnableMonitoring returns the UseEnableMonitoring field if non-nil, zero value otherwise.

### GetUseEnableMonitoringOk

`func (o *MsserverDhcpServer) GetUseEnableMonitoringOk() (*bool, bool)`

GetUseEnableMonitoringOk returns a tuple with the UseEnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableMonitoring

`func (o *MsserverDhcpServer) SetUseEnableMonitoring(v bool)`

SetUseEnableMonitoring sets UseEnableMonitoring field to given value.

### HasUseEnableMonitoring

`func (o *MsserverDhcpServer) HasUseEnableMonitoring() bool`

HasUseEnableMonitoring returns a boolean if a field has been set.

### GetEnableMonitoring

`func (o *MsserverDhcpServer) GetEnableMonitoring() bool`

GetEnableMonitoring returns the EnableMonitoring field if non-nil, zero value otherwise.

### GetEnableMonitoringOk

`func (o *MsserverDhcpServer) GetEnableMonitoringOk() (*bool, bool)`

GetEnableMonitoringOk returns a tuple with the EnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitoring

`func (o *MsserverDhcpServer) SetEnableMonitoring(v bool)`

SetEnableMonitoring sets EnableMonitoring field to given value.

### HasEnableMonitoring

`func (o *MsserverDhcpServer) HasEnableMonitoring() bool`

HasEnableMonitoring returns a boolean if a field has been set.

### GetUseEnableInvalidMac

`func (o *MsserverDhcpServer) GetUseEnableInvalidMac() bool`

GetUseEnableInvalidMac returns the UseEnableInvalidMac field if non-nil, zero value otherwise.

### GetUseEnableInvalidMacOk

`func (o *MsserverDhcpServer) GetUseEnableInvalidMacOk() (*bool, bool)`

GetUseEnableInvalidMacOk returns a tuple with the UseEnableInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableInvalidMac

`func (o *MsserverDhcpServer) SetUseEnableInvalidMac(v bool)`

SetUseEnableInvalidMac sets UseEnableInvalidMac field to given value.

### HasUseEnableInvalidMac

`func (o *MsserverDhcpServer) HasUseEnableInvalidMac() bool`

HasUseEnableInvalidMac returns a boolean if a field has been set.

### GetEnableInvalidMac

`func (o *MsserverDhcpServer) GetEnableInvalidMac() bool`

GetEnableInvalidMac returns the EnableInvalidMac field if non-nil, zero value otherwise.

### GetEnableInvalidMacOk

`func (o *MsserverDhcpServer) GetEnableInvalidMacOk() (*bool, bool)`

GetEnableInvalidMacOk returns a tuple with the EnableInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInvalidMac

`func (o *MsserverDhcpServer) SetEnableInvalidMac(v bool)`

SetEnableInvalidMac sets EnableInvalidMac field to given value.

### HasEnableInvalidMac

`func (o *MsserverDhcpServer) HasEnableInvalidMac() bool`

HasEnableInvalidMac returns a boolean if a field has been set.

### GetSupportsFailover

`func (o *MsserverDhcpServer) GetSupportsFailover() bool`

GetSupportsFailover returns the SupportsFailover field if non-nil, zero value otherwise.

### GetSupportsFailoverOk

`func (o *MsserverDhcpServer) GetSupportsFailoverOk() (*bool, bool)`

GetSupportsFailoverOk returns a tuple with the SupportsFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFailover

`func (o *MsserverDhcpServer) SetSupportsFailover(v bool)`

SetSupportsFailover sets SupportsFailover field to given value.

### HasSupportsFailover

`func (o *MsserverDhcpServer) HasSupportsFailover() bool`

HasSupportsFailover returns a boolean if a field has been set.

### GetUseSynchronizationMinDelay

`func (o *MsserverDhcpServer) GetUseSynchronizationMinDelay() bool`

GetUseSynchronizationMinDelay returns the UseSynchronizationMinDelay field if non-nil, zero value otherwise.

### GetUseSynchronizationMinDelayOk

`func (o *MsserverDhcpServer) GetUseSynchronizationMinDelayOk() (*bool, bool)`

GetUseSynchronizationMinDelayOk returns a tuple with the UseSynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationMinDelay

`func (o *MsserverDhcpServer) SetUseSynchronizationMinDelay(v bool)`

SetUseSynchronizationMinDelay sets UseSynchronizationMinDelay field to given value.

### HasUseSynchronizationMinDelay

`func (o *MsserverDhcpServer) HasUseSynchronizationMinDelay() bool`

HasUseSynchronizationMinDelay returns a boolean if a field has been set.

### GetSynchronizationMinDelay

`func (o *MsserverDhcpServer) GetSynchronizationMinDelay() int64`

GetSynchronizationMinDelay returns the SynchronizationMinDelay field if non-nil, zero value otherwise.

### GetSynchronizationMinDelayOk

`func (o *MsserverDhcpServer) GetSynchronizationMinDelayOk() (*int64, bool)`

GetSynchronizationMinDelayOk returns a tuple with the SynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationMinDelay

`func (o *MsserverDhcpServer) SetSynchronizationMinDelay(v int64)`

SetSynchronizationMinDelay sets SynchronizationMinDelay field to given value.

### HasSynchronizationMinDelay

`func (o *MsserverDhcpServer) HasSynchronizationMinDelay() bool`

HasSynchronizationMinDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


