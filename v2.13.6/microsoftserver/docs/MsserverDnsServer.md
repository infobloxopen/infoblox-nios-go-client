# MsserverDnsServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseLogin** | Pointer to **bool** | Flag to override login name and password from the MS Server | [optional] 
**LoginName** | Pointer to **string** | Microsoft Server login name | [optional] 
**LoginPassword** | Pointer to **string** | Microsoft Server login password | [optional] 
**Managed** | Pointer to **bool** | flag indicating if the DNS service is managed | [optional] 
**NextSyncControl** | Pointer to **string** | Defines what control to apply on the DNS server | [optional] 
**Status** | Pointer to **string** | Status of the Microsoft DNS Service | [optional] [readonly] 
**StatusDetail** | Pointer to **string** | Detailed status of the DNS status | [optional] [readonly] 
**StatusLastUpdated** | Pointer to **int64** | Timestamp of the last update | [optional] [readonly] 
**LastSyncTs** | Pointer to **int64** | Timestamp of the last synchronization attempt | [optional] [readonly] 
**LastSyncStatus** | Pointer to **string** | Status of the last synchronization attempt | [optional] [readonly] 
**LastSyncDetail** | Pointer to **string** | Detailled status of the last synchronization attempt | [optional] [readonly] 
**Forwarders** | Pointer to **string** | Ordered list of IP addresses to forward queries to | [optional] [readonly] 
**SupportsIpv6** | Pointer to **bool** | Flag indicating if the server supports IPv6 | [optional] [readonly] 
**SupportsIpv6Reverse** | Pointer to **bool** | Flag indicating if the server supports reverse IPv6 zones | [optional] [readonly] 
**SupportsRrDname** | Pointer to **bool** | Flag indicating if the server supports DNAME records | [optional] [readonly] 
**SupportsDnssec** | Pointer to **bool** | Flag indicating if the server supports | [optional] [readonly] 
**SupportsActiveDirectory** | Pointer to **bool** | Flag indicating if the server supports AD integrated zones | [optional] [readonly] 
**Address** | Pointer to **string** | MS Server ip address | [optional] [readonly] 
**SupportsRrNaptr** | Pointer to **bool** | Flag indicating if the server supports NAPTR records | [optional] [readonly] 
**UseEnableMonitoring** | Pointer to **bool** | Override enable monitoring inherited from grid level | [optional] 
**EnableMonitoring** | Pointer to **bool** | Flag indicating if the DNS service is monitored and controlled | [optional] 
**UseSynchronizationMinDelay** | Pointer to **bool** | Flag to override synchronization interval from the MS Server | [optional] 
**SynchronizationMinDelay** | Pointer to **int64** | Minimum number of minutes between two synchronizations | [optional] 
**UseEnableDnsReportsSync** | Pointer to **bool** | Override enable reports data inherited from grid level | [optional] 
**EnableDnsReportsSync** | Pointer to **bool** | Enable or Disable MS DNS data for reports from this MS Server | [optional] 

## Methods

### NewMsserverDnsServer

`func NewMsserverDnsServer() *MsserverDnsServer`

NewMsserverDnsServer instantiates a new MsserverDnsServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverDnsServerWithDefaults

`func NewMsserverDnsServerWithDefaults() *MsserverDnsServer`

NewMsserverDnsServerWithDefaults instantiates a new MsserverDnsServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseLogin

`func (o *MsserverDnsServer) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *MsserverDnsServer) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *MsserverDnsServer) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *MsserverDnsServer) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetLoginName

`func (o *MsserverDnsServer) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MsserverDnsServer) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MsserverDnsServer) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MsserverDnsServer) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *MsserverDnsServer) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *MsserverDnsServer) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *MsserverDnsServer) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *MsserverDnsServer) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetManaged

`func (o *MsserverDnsServer) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MsserverDnsServer) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MsserverDnsServer) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MsserverDnsServer) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNextSyncControl

`func (o *MsserverDnsServer) GetNextSyncControl() string`

GetNextSyncControl returns the NextSyncControl field if non-nil, zero value otherwise.

### GetNextSyncControlOk

`func (o *MsserverDnsServer) GetNextSyncControlOk() (*string, bool)`

GetNextSyncControlOk returns a tuple with the NextSyncControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSyncControl

`func (o *MsserverDnsServer) SetNextSyncControl(v string)`

SetNextSyncControl sets NextSyncControl field to given value.

### HasNextSyncControl

`func (o *MsserverDnsServer) HasNextSyncControl() bool`

HasNextSyncControl returns a boolean if a field has been set.

### GetStatus

`func (o *MsserverDnsServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MsserverDnsServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MsserverDnsServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MsserverDnsServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *MsserverDnsServer) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *MsserverDnsServer) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *MsserverDnsServer) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *MsserverDnsServer) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetStatusLastUpdated

`func (o *MsserverDnsServer) GetStatusLastUpdated() int64`

GetStatusLastUpdated returns the StatusLastUpdated field if non-nil, zero value otherwise.

### GetStatusLastUpdatedOk

`func (o *MsserverDnsServer) GetStatusLastUpdatedOk() (*int64, bool)`

GetStatusLastUpdatedOk returns a tuple with the StatusLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastUpdated

`func (o *MsserverDnsServer) SetStatusLastUpdated(v int64)`

SetStatusLastUpdated sets StatusLastUpdated field to given value.

### HasStatusLastUpdated

`func (o *MsserverDnsServer) HasStatusLastUpdated() bool`

HasStatusLastUpdated returns a boolean if a field has been set.

### GetLastSyncTs

`func (o *MsserverDnsServer) GetLastSyncTs() int64`

GetLastSyncTs returns the LastSyncTs field if non-nil, zero value otherwise.

### GetLastSyncTsOk

`func (o *MsserverDnsServer) GetLastSyncTsOk() (*int64, bool)`

GetLastSyncTsOk returns a tuple with the LastSyncTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTs

`func (o *MsserverDnsServer) SetLastSyncTs(v int64)`

SetLastSyncTs sets LastSyncTs field to given value.

### HasLastSyncTs

`func (o *MsserverDnsServer) HasLastSyncTs() bool`

HasLastSyncTs returns a boolean if a field has been set.

### GetLastSyncStatus

`func (o *MsserverDnsServer) GetLastSyncStatus() string`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *MsserverDnsServer) GetLastSyncStatusOk() (*string, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *MsserverDnsServer) SetLastSyncStatus(v string)`

SetLastSyncStatus sets LastSyncStatus field to given value.

### HasLastSyncStatus

`func (o *MsserverDnsServer) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.

### GetLastSyncDetail

`func (o *MsserverDnsServer) GetLastSyncDetail() string`

GetLastSyncDetail returns the LastSyncDetail field if non-nil, zero value otherwise.

### GetLastSyncDetailOk

`func (o *MsserverDnsServer) GetLastSyncDetailOk() (*string, bool)`

GetLastSyncDetailOk returns a tuple with the LastSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDetail

`func (o *MsserverDnsServer) SetLastSyncDetail(v string)`

SetLastSyncDetail sets LastSyncDetail field to given value.

### HasLastSyncDetail

`func (o *MsserverDnsServer) HasLastSyncDetail() bool`

HasLastSyncDetail returns a boolean if a field has been set.

### GetForwarders

`func (o *MsserverDnsServer) GetForwarders() string`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *MsserverDnsServer) GetForwardersOk() (*string, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *MsserverDnsServer) SetForwarders(v string)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *MsserverDnsServer) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetSupportsIpv6

`func (o *MsserverDnsServer) GetSupportsIpv6() bool`

GetSupportsIpv6 returns the SupportsIpv6 field if non-nil, zero value otherwise.

### GetSupportsIpv6Ok

`func (o *MsserverDnsServer) GetSupportsIpv6Ok() (*bool, bool)`

GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIpv6

`func (o *MsserverDnsServer) SetSupportsIpv6(v bool)`

SetSupportsIpv6 sets SupportsIpv6 field to given value.

### HasSupportsIpv6

`func (o *MsserverDnsServer) HasSupportsIpv6() bool`

HasSupportsIpv6 returns a boolean if a field has been set.

### GetSupportsIpv6Reverse

`func (o *MsserverDnsServer) GetSupportsIpv6Reverse() bool`

GetSupportsIpv6Reverse returns the SupportsIpv6Reverse field if non-nil, zero value otherwise.

### GetSupportsIpv6ReverseOk

`func (o *MsserverDnsServer) GetSupportsIpv6ReverseOk() (*bool, bool)`

GetSupportsIpv6ReverseOk returns a tuple with the SupportsIpv6Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIpv6Reverse

`func (o *MsserverDnsServer) SetSupportsIpv6Reverse(v bool)`

SetSupportsIpv6Reverse sets SupportsIpv6Reverse field to given value.

### HasSupportsIpv6Reverse

`func (o *MsserverDnsServer) HasSupportsIpv6Reverse() bool`

HasSupportsIpv6Reverse returns a boolean if a field has been set.

### GetSupportsRrDname

`func (o *MsserverDnsServer) GetSupportsRrDname() bool`

GetSupportsRrDname returns the SupportsRrDname field if non-nil, zero value otherwise.

### GetSupportsRrDnameOk

`func (o *MsserverDnsServer) GetSupportsRrDnameOk() (*bool, bool)`

GetSupportsRrDnameOk returns a tuple with the SupportsRrDname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRrDname

`func (o *MsserverDnsServer) SetSupportsRrDname(v bool)`

SetSupportsRrDname sets SupportsRrDname field to given value.

### HasSupportsRrDname

`func (o *MsserverDnsServer) HasSupportsRrDname() bool`

HasSupportsRrDname returns a boolean if a field has been set.

### GetSupportsDnssec

`func (o *MsserverDnsServer) GetSupportsDnssec() bool`

GetSupportsDnssec returns the SupportsDnssec field if non-nil, zero value otherwise.

### GetSupportsDnssecOk

`func (o *MsserverDnsServer) GetSupportsDnssecOk() (*bool, bool)`

GetSupportsDnssecOk returns a tuple with the SupportsDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDnssec

`func (o *MsserverDnsServer) SetSupportsDnssec(v bool)`

SetSupportsDnssec sets SupportsDnssec field to given value.

### HasSupportsDnssec

`func (o *MsserverDnsServer) HasSupportsDnssec() bool`

HasSupportsDnssec returns a boolean if a field has been set.

### GetSupportsActiveDirectory

`func (o *MsserverDnsServer) GetSupportsActiveDirectory() bool`

GetSupportsActiveDirectory returns the SupportsActiveDirectory field if non-nil, zero value otherwise.

### GetSupportsActiveDirectoryOk

`func (o *MsserverDnsServer) GetSupportsActiveDirectoryOk() (*bool, bool)`

GetSupportsActiveDirectoryOk returns a tuple with the SupportsActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsActiveDirectory

`func (o *MsserverDnsServer) SetSupportsActiveDirectory(v bool)`

SetSupportsActiveDirectory sets SupportsActiveDirectory field to given value.

### HasSupportsActiveDirectory

`func (o *MsserverDnsServer) HasSupportsActiveDirectory() bool`

HasSupportsActiveDirectory returns a boolean if a field has been set.

### GetAddress

`func (o *MsserverDnsServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MsserverDnsServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MsserverDnsServer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MsserverDnsServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSupportsRrNaptr

`func (o *MsserverDnsServer) GetSupportsRrNaptr() bool`

GetSupportsRrNaptr returns the SupportsRrNaptr field if non-nil, zero value otherwise.

### GetSupportsRrNaptrOk

`func (o *MsserverDnsServer) GetSupportsRrNaptrOk() (*bool, bool)`

GetSupportsRrNaptrOk returns a tuple with the SupportsRrNaptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRrNaptr

`func (o *MsserverDnsServer) SetSupportsRrNaptr(v bool)`

SetSupportsRrNaptr sets SupportsRrNaptr field to given value.

### HasSupportsRrNaptr

`func (o *MsserverDnsServer) HasSupportsRrNaptr() bool`

HasSupportsRrNaptr returns a boolean if a field has been set.

### GetUseEnableMonitoring

`func (o *MsserverDnsServer) GetUseEnableMonitoring() bool`

GetUseEnableMonitoring returns the UseEnableMonitoring field if non-nil, zero value otherwise.

### GetUseEnableMonitoringOk

`func (o *MsserverDnsServer) GetUseEnableMonitoringOk() (*bool, bool)`

GetUseEnableMonitoringOk returns a tuple with the UseEnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableMonitoring

`func (o *MsserverDnsServer) SetUseEnableMonitoring(v bool)`

SetUseEnableMonitoring sets UseEnableMonitoring field to given value.

### HasUseEnableMonitoring

`func (o *MsserverDnsServer) HasUseEnableMonitoring() bool`

HasUseEnableMonitoring returns a boolean if a field has been set.

### GetEnableMonitoring

`func (o *MsserverDnsServer) GetEnableMonitoring() bool`

GetEnableMonitoring returns the EnableMonitoring field if non-nil, zero value otherwise.

### GetEnableMonitoringOk

`func (o *MsserverDnsServer) GetEnableMonitoringOk() (*bool, bool)`

GetEnableMonitoringOk returns a tuple with the EnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitoring

`func (o *MsserverDnsServer) SetEnableMonitoring(v bool)`

SetEnableMonitoring sets EnableMonitoring field to given value.

### HasEnableMonitoring

`func (o *MsserverDnsServer) HasEnableMonitoring() bool`

HasEnableMonitoring returns a boolean if a field has been set.

### GetUseSynchronizationMinDelay

`func (o *MsserverDnsServer) GetUseSynchronizationMinDelay() bool`

GetUseSynchronizationMinDelay returns the UseSynchronizationMinDelay field if non-nil, zero value otherwise.

### GetUseSynchronizationMinDelayOk

`func (o *MsserverDnsServer) GetUseSynchronizationMinDelayOk() (*bool, bool)`

GetUseSynchronizationMinDelayOk returns a tuple with the UseSynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationMinDelay

`func (o *MsserverDnsServer) SetUseSynchronizationMinDelay(v bool)`

SetUseSynchronizationMinDelay sets UseSynchronizationMinDelay field to given value.

### HasUseSynchronizationMinDelay

`func (o *MsserverDnsServer) HasUseSynchronizationMinDelay() bool`

HasUseSynchronizationMinDelay returns a boolean if a field has been set.

### GetSynchronizationMinDelay

`func (o *MsserverDnsServer) GetSynchronizationMinDelay() int64`

GetSynchronizationMinDelay returns the SynchronizationMinDelay field if non-nil, zero value otherwise.

### GetSynchronizationMinDelayOk

`func (o *MsserverDnsServer) GetSynchronizationMinDelayOk() (*int64, bool)`

GetSynchronizationMinDelayOk returns a tuple with the SynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationMinDelay

`func (o *MsserverDnsServer) SetSynchronizationMinDelay(v int64)`

SetSynchronizationMinDelay sets SynchronizationMinDelay field to given value.

### HasSynchronizationMinDelay

`func (o *MsserverDnsServer) HasSynchronizationMinDelay() bool`

HasSynchronizationMinDelay returns a boolean if a field has been set.

### GetUseEnableDnsReportsSync

`func (o *MsserverDnsServer) GetUseEnableDnsReportsSync() bool`

GetUseEnableDnsReportsSync returns the UseEnableDnsReportsSync field if non-nil, zero value otherwise.

### GetUseEnableDnsReportsSyncOk

`func (o *MsserverDnsServer) GetUseEnableDnsReportsSyncOk() (*bool, bool)`

GetUseEnableDnsReportsSyncOk returns a tuple with the UseEnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDnsReportsSync

`func (o *MsserverDnsServer) SetUseEnableDnsReportsSync(v bool)`

SetUseEnableDnsReportsSync sets UseEnableDnsReportsSync field to given value.

### HasUseEnableDnsReportsSync

`func (o *MsserverDnsServer) HasUseEnableDnsReportsSync() bool`

HasUseEnableDnsReportsSync returns a boolean if a field has been set.

### GetEnableDnsReportsSync

`func (o *MsserverDnsServer) GetEnableDnsReportsSync() bool`

GetEnableDnsReportsSync returns the EnableDnsReportsSync field if non-nil, zero value otherwise.

### GetEnableDnsReportsSyncOk

`func (o *MsserverDnsServer) GetEnableDnsReportsSyncOk() (*bool, bool)`

GetEnableDnsReportsSyncOk returns a tuple with the EnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsReportsSync

`func (o *MsserverDnsServer) SetEnableDnsReportsSync(v bool)`

SetEnableDnsReportsSync sets EnableDnsReportsSync field to given value.

### HasEnableDnsReportsSync

`func (o *MsserverDnsServer) HasEnableDnsReportsSync() bool`

HasEnableDnsReportsSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


