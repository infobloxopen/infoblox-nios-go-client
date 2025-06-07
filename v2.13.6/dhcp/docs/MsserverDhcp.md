# MsserverDhcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The address or FQDN of the DHCP Microsoft Server. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment from Microsoft Server | [optional] [readonly] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of DHCP objects belonging to the DHCP Microsoft Server multiplied by 1000. This is the percentage of the total number of available IP addresses from all the DHCP objects belonging to the DHCP Microsoft Server versus the total number of all IP addresses in all of the DHCP objects on the DHCP Microsoft Server. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of DHCP objects that belong to the DHCP Microsoft Server. | [optional] [readonly] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the DHCP objects on the DHCP Microsoft Server. | [optional] [readonly] 
**LastSyncTs** | Pointer to **int64** | Timestamp of the last synchronization attempt | [optional] [readonly] 
**LoginName** | Pointer to **string** | The login name of the DHCP Microsoft Server. | [optional] 
**LoginPassword** | Pointer to **string** | The login password of the DHCP Microsoft Server. | [optional] 
**NetworkView** | Pointer to **string** | Network view to update | [optional] [readonly] 
**NextSyncControl** | Pointer to **string** | Defines what control to apply on the DHCP server | [optional] 
**ReadOnly** | Pointer to **bool** | Whether Microsoft server is read only | [optional] [readonly] 
**ServerName** | Pointer to **string** | Microsoft server address | [optional] [readonly] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in DHCP objects that belong to the DHCP Microsoft Server. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the Microsoft DHCP Service | [optional] 
**StatusDetail** | Pointer to **string** | Detailed status of the DHCP status | [optional] [readonly] 
**StatusLastUpdated** | Pointer to **int64** | Timestamp of the last update | [optional] [readonly] 
**SupportsFailover** | Pointer to **bool** | Flag indicating if the DHCP supports Failover | [optional] [readonly] 
**SynchronizationInterval** | Pointer to **int64** | The minimum number of minutes between two synchronizations. | [optional] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in DHCP objects that belong to the DHCP Microsoft Server. | [optional] [readonly] 
**UseLogin** | Pointer to **bool** | Use flag for: login_name , login_password | [optional] 
**UseSynchronizationInterval** | Pointer to **bool** | Use flag for: synchronization_interval | [optional] 

## Methods

### NewMsserverDhcp

`func NewMsserverDhcp() *MsserverDhcp`

NewMsserverDhcp instantiates a new MsserverDhcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverDhcpWithDefaults

`func NewMsserverDhcpWithDefaults() *MsserverDhcp`

NewMsserverDhcpWithDefaults instantiates a new MsserverDhcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MsserverDhcp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MsserverDhcp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MsserverDhcp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MsserverDhcp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *MsserverDhcp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MsserverDhcp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MsserverDhcp) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MsserverDhcp) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *MsserverDhcp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MsserverDhcp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MsserverDhcp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MsserverDhcp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *MsserverDhcp) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *MsserverDhcp) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *MsserverDhcp) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *MsserverDhcp) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *MsserverDhcp) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *MsserverDhcp) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *MsserverDhcp) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *MsserverDhcp) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *MsserverDhcp) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *MsserverDhcp) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *MsserverDhcp) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *MsserverDhcp) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetLastSyncTs

`func (o *MsserverDhcp) GetLastSyncTs() int64`

GetLastSyncTs returns the LastSyncTs field if non-nil, zero value otherwise.

### GetLastSyncTsOk

`func (o *MsserverDhcp) GetLastSyncTsOk() (*int64, bool)`

GetLastSyncTsOk returns a tuple with the LastSyncTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTs

`func (o *MsserverDhcp) SetLastSyncTs(v int64)`

SetLastSyncTs sets LastSyncTs field to given value.

### HasLastSyncTs

`func (o *MsserverDhcp) HasLastSyncTs() bool`

HasLastSyncTs returns a boolean if a field has been set.

### GetLoginName

`func (o *MsserverDhcp) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MsserverDhcp) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MsserverDhcp) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MsserverDhcp) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *MsserverDhcp) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *MsserverDhcp) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *MsserverDhcp) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *MsserverDhcp) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetNetworkView

`func (o *MsserverDhcp) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *MsserverDhcp) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *MsserverDhcp) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *MsserverDhcp) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextSyncControl

`func (o *MsserverDhcp) GetNextSyncControl() string`

GetNextSyncControl returns the NextSyncControl field if non-nil, zero value otherwise.

### GetNextSyncControlOk

`func (o *MsserverDhcp) GetNextSyncControlOk() (*string, bool)`

GetNextSyncControlOk returns a tuple with the NextSyncControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSyncControl

`func (o *MsserverDhcp) SetNextSyncControl(v string)`

SetNextSyncControl sets NextSyncControl field to given value.

### HasNextSyncControl

`func (o *MsserverDhcp) HasNextSyncControl() bool`

HasNextSyncControl returns a boolean if a field has been set.

### GetReadOnly

`func (o *MsserverDhcp) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MsserverDhcp) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MsserverDhcp) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MsserverDhcp) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetServerName

`func (o *MsserverDhcp) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *MsserverDhcp) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *MsserverDhcp) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *MsserverDhcp) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStaticHosts

`func (o *MsserverDhcp) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *MsserverDhcp) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *MsserverDhcp) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *MsserverDhcp) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetStatus

`func (o *MsserverDhcp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MsserverDhcp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MsserverDhcp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MsserverDhcp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *MsserverDhcp) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *MsserverDhcp) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *MsserverDhcp) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *MsserverDhcp) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetStatusLastUpdated

`func (o *MsserverDhcp) GetStatusLastUpdated() int64`

GetStatusLastUpdated returns the StatusLastUpdated field if non-nil, zero value otherwise.

### GetStatusLastUpdatedOk

`func (o *MsserverDhcp) GetStatusLastUpdatedOk() (*int64, bool)`

GetStatusLastUpdatedOk returns a tuple with the StatusLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastUpdated

`func (o *MsserverDhcp) SetStatusLastUpdated(v int64)`

SetStatusLastUpdated sets StatusLastUpdated field to given value.

### HasStatusLastUpdated

`func (o *MsserverDhcp) HasStatusLastUpdated() bool`

HasStatusLastUpdated returns a boolean if a field has been set.

### GetSupportsFailover

`func (o *MsserverDhcp) GetSupportsFailover() bool`

GetSupportsFailover returns the SupportsFailover field if non-nil, zero value otherwise.

### GetSupportsFailoverOk

`func (o *MsserverDhcp) GetSupportsFailoverOk() (*bool, bool)`

GetSupportsFailoverOk returns a tuple with the SupportsFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFailover

`func (o *MsserverDhcp) SetSupportsFailover(v bool)`

SetSupportsFailover sets SupportsFailover field to given value.

### HasSupportsFailover

`func (o *MsserverDhcp) HasSupportsFailover() bool`

HasSupportsFailover returns a boolean if a field has been set.

### GetSynchronizationInterval

`func (o *MsserverDhcp) GetSynchronizationInterval() int64`

GetSynchronizationInterval returns the SynchronizationInterval field if non-nil, zero value otherwise.

### GetSynchronizationIntervalOk

`func (o *MsserverDhcp) GetSynchronizationIntervalOk() (*int64, bool)`

GetSynchronizationIntervalOk returns a tuple with the SynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationInterval

`func (o *MsserverDhcp) SetSynchronizationInterval(v int64)`

SetSynchronizationInterval sets SynchronizationInterval field to given value.

### HasSynchronizationInterval

`func (o *MsserverDhcp) HasSynchronizationInterval() bool`

HasSynchronizationInterval returns a boolean if a field has been set.

### GetTotalHosts

`func (o *MsserverDhcp) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *MsserverDhcp) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *MsserverDhcp) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *MsserverDhcp) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUseLogin

`func (o *MsserverDhcp) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *MsserverDhcp) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *MsserverDhcp) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *MsserverDhcp) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetUseSynchronizationInterval

`func (o *MsserverDhcp) GetUseSynchronizationInterval() bool`

GetUseSynchronizationInterval returns the UseSynchronizationInterval field if non-nil, zero value otherwise.

### GetUseSynchronizationIntervalOk

`func (o *MsserverDhcp) GetUseSynchronizationIntervalOk() (*bool, bool)`

GetUseSynchronizationIntervalOk returns a tuple with the UseSynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationInterval

`func (o *MsserverDhcp) SetUseSynchronizationInterval(v bool)`

SetUseSynchronizationInterval sets UseSynchronizationInterval field to given value.

### HasUseSynchronizationInterval

`func (o *MsserverDhcp) HasUseSynchronizationInterval() bool`

HasUseSynchronizationInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


