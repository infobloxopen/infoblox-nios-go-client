# Msserver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AdDomain** | Pointer to **string** | The Active Directory domain to which this server belongs (if applicable). | [optional] [readonly] 
**AdSites** | Pointer to [**MsserverAdSites**](MsserverAdSites.md) |  | [optional] 
**AdUser** | Pointer to [**MsserverAdUser**](MsserverAdUser.md) |  | [optional] 
**Address** | Pointer to **string** | The address or FQDN of the server. | [optional] 
**Comment** | Pointer to **string** | User comments for this Microsoft Server | [optional] 
**ConnectionStatus** | Pointer to **string** | Result of the last RPC connection attempt made | [optional] [readonly] 
**ConnectionStatusDetail** | Pointer to **string** | Detail of the last connection attempt made | [optional] [readonly] 
**DhcpServer** | Pointer to [**MsserverDhcpServer**](MsserverDhcpServer.md) |  | [optional] 
**Disabled** | Pointer to **bool** | Allow/forbids usage of this Microsoft Server | [optional] 
**DnsServer** | Pointer to [**MsserverDnsServer**](MsserverDnsServer.md) |  | [optional] 
**DnsView** | Pointer to **string** | Reference to the DNS view | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**GridMember** | Pointer to **string** | eference to the assigned grid member | [optional] 
**LastSeen** | Pointer to **int64** | Timestamp of the last message received | [optional] [readonly] 
**LogDestination** | Pointer to **string** | Directs logging of sync messages either to syslog or mslog | [optional] 
**LogLevel** | Pointer to **string** | Log level for this Microsoft Server | [optional] 
**LoginName** | Pointer to **string** | Microsoft Server login name, with optional domainname | [optional] 
**LoginPassword** | Pointer to **string** | Microsoft Server login password | [optional] 
**ManagingMember** | Pointer to **string** | Hostname of grid member managing this Microsoft Server | [optional] [readonly] 
**MsMaxConnection** | Pointer to **int64** | Maximum number of connections to MS server | [optional] 
**MsRpcTimeoutInSeconds** | Pointer to **int64** | Timeout in seconds of RPC connections for this MS Server | [optional] 
**NetworkView** | Pointer to **string** | Reference to the network view | [optional] 
**ReadOnly** | Pointer to **bool** | Enable read-only management for this Microsoft Server | [optional] 
**RootAdDomain** | Pointer to **string** | The root Active Directory domain to which this server belongs (if applicable). | [optional] [readonly] 
**ServerName** | Pointer to **string** | Gives the server name as reported by itself | [optional] [readonly] 
**SynchronizationMinDelay** | Pointer to **int64** | Minimum number of minutes between two synchronizations | [optional] 
**SynchronizationStatus** | Pointer to **string** | Synchronization status summary | [optional] [readonly] 
**SynchronizationStatusDetail** | Pointer to **string** | Detail status if synchronization_status is ERROR | [optional] [readonly] 
**UseLogDestination** | Pointer to **bool** | Override log_destination inherited from grid level | [optional] 
**UseMsMaxConnection** | Pointer to **bool** | Override grid ms_max_connection setting | [optional] 
**UseMsRpcTimeoutInSeconds** | Pointer to **bool** | Flag to override cluster RPC timeout | [optional] 
**Version** | Pointer to **string** | Version of the Microsoft Server | [optional] [readonly] 

## Methods

### NewMsserver

`func NewMsserver() *Msserver`

NewMsserver instantiates a new Msserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverWithDefaults

`func NewMsserverWithDefaults() *Msserver`

NewMsserverWithDefaults instantiates a new Msserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Msserver) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Msserver) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Msserver) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Msserver) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdDomain

`func (o *Msserver) GetAdDomain() string`

GetAdDomain returns the AdDomain field if non-nil, zero value otherwise.

### GetAdDomainOk

`func (o *Msserver) GetAdDomainOk() (*string, bool)`

GetAdDomainOk returns a tuple with the AdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomain

`func (o *Msserver) SetAdDomain(v string)`

SetAdDomain sets AdDomain field to given value.

### HasAdDomain

`func (o *Msserver) HasAdDomain() bool`

HasAdDomain returns a boolean if a field has been set.

### GetAdSites

`func (o *Msserver) GetAdSites() MsserverAdSites`

GetAdSites returns the AdSites field if non-nil, zero value otherwise.

### GetAdSitesOk

`func (o *Msserver) GetAdSitesOk() (*MsserverAdSites, bool)`

GetAdSitesOk returns a tuple with the AdSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSites

`func (o *Msserver) SetAdSites(v MsserverAdSites)`

SetAdSites sets AdSites field to given value.

### HasAdSites

`func (o *Msserver) HasAdSites() bool`

HasAdSites returns a boolean if a field has been set.

### GetAdUser

`func (o *Msserver) GetAdUser() MsserverAdUser`

GetAdUser returns the AdUser field if non-nil, zero value otherwise.

### GetAdUserOk

`func (o *Msserver) GetAdUserOk() (*MsserverAdUser, bool)`

GetAdUserOk returns a tuple with the AdUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUser

`func (o *Msserver) SetAdUser(v MsserverAdUser)`

SetAdUser sets AdUser field to given value.

### HasAdUser

`func (o *Msserver) HasAdUser() bool`

HasAdUser returns a boolean if a field has been set.

### GetAddress

`func (o *Msserver) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Msserver) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Msserver) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Msserver) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *Msserver) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Msserver) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Msserver) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Msserver) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *Msserver) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *Msserver) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *Msserver) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *Msserver) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetConnectionStatusDetail

`func (o *Msserver) GetConnectionStatusDetail() string`

GetConnectionStatusDetail returns the ConnectionStatusDetail field if non-nil, zero value otherwise.

### GetConnectionStatusDetailOk

`func (o *Msserver) GetConnectionStatusDetailOk() (*string, bool)`

GetConnectionStatusDetailOk returns a tuple with the ConnectionStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusDetail

`func (o *Msserver) SetConnectionStatusDetail(v string)`

SetConnectionStatusDetail sets ConnectionStatusDetail field to given value.

### HasConnectionStatusDetail

`func (o *Msserver) HasConnectionStatusDetail() bool`

HasConnectionStatusDetail returns a boolean if a field has been set.

### GetDhcpServer

`func (o *Msserver) GetDhcpServer() MsserverDhcpServer`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *Msserver) GetDhcpServerOk() (*MsserverDhcpServer, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *Msserver) SetDhcpServer(v MsserverDhcpServer)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *Msserver) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDisabled

`func (o *Msserver) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Msserver) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Msserver) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Msserver) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnsServer

`func (o *Msserver) GetDnsServer() MsserverDnsServer`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *Msserver) GetDnsServerOk() (*MsserverDnsServer, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *Msserver) SetDnsServer(v MsserverDnsServer)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *Msserver) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetDnsView

`func (o *Msserver) GetDnsView() string`

GetDnsView returns the DnsView field if non-nil, zero value otherwise.

### GetDnsViewOk

`func (o *Msserver) GetDnsViewOk() (*string, bool)`

GetDnsViewOk returns a tuple with the DnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsView

`func (o *Msserver) SetDnsView(v string)`

SetDnsView sets DnsView field to given value.

### HasDnsView

`func (o *Msserver) HasDnsView() bool`

HasDnsView returns a boolean if a field has been set.

### GetExtattrs

`func (o *Msserver) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Msserver) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Msserver) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Msserver) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetGridMember

`func (o *Msserver) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *Msserver) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *Msserver) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *Msserver) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetLastSeen

`func (o *Msserver) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Msserver) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Msserver) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Msserver) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLogDestination

`func (o *Msserver) GetLogDestination() string`

GetLogDestination returns the LogDestination field if non-nil, zero value otherwise.

### GetLogDestinationOk

`func (o *Msserver) GetLogDestinationOk() (*string, bool)`

GetLogDestinationOk returns a tuple with the LogDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDestination

`func (o *Msserver) SetLogDestination(v string)`

SetLogDestination sets LogDestination field to given value.

### HasLogDestination

`func (o *Msserver) HasLogDestination() bool`

HasLogDestination returns a boolean if a field has been set.

### GetLogLevel

`func (o *Msserver) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *Msserver) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *Msserver) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *Msserver) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLoginName

`func (o *Msserver) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *Msserver) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *Msserver) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *Msserver) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *Msserver) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *Msserver) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *Msserver) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *Msserver) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetManagingMember

`func (o *Msserver) GetManagingMember() string`

GetManagingMember returns the ManagingMember field if non-nil, zero value otherwise.

### GetManagingMemberOk

`func (o *Msserver) GetManagingMemberOk() (*string, bool)`

GetManagingMemberOk returns a tuple with the ManagingMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingMember

`func (o *Msserver) SetManagingMember(v string)`

SetManagingMember sets ManagingMember field to given value.

### HasManagingMember

`func (o *Msserver) HasManagingMember() bool`

HasManagingMember returns a boolean if a field has been set.

### GetMsMaxConnection

`func (o *Msserver) GetMsMaxConnection() int64`

GetMsMaxConnection returns the MsMaxConnection field if non-nil, zero value otherwise.

### GetMsMaxConnectionOk

`func (o *Msserver) GetMsMaxConnectionOk() (*int64, bool)`

GetMsMaxConnectionOk returns a tuple with the MsMaxConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsMaxConnection

`func (o *Msserver) SetMsMaxConnection(v int64)`

SetMsMaxConnection sets MsMaxConnection field to given value.

### HasMsMaxConnection

`func (o *Msserver) HasMsMaxConnection() bool`

HasMsMaxConnection returns a boolean if a field has been set.

### GetMsRpcTimeoutInSeconds

`func (o *Msserver) GetMsRpcTimeoutInSeconds() int64`

GetMsRpcTimeoutInSeconds returns the MsRpcTimeoutInSeconds field if non-nil, zero value otherwise.

### GetMsRpcTimeoutInSecondsOk

`func (o *Msserver) GetMsRpcTimeoutInSecondsOk() (*int64, bool)`

GetMsRpcTimeoutInSecondsOk returns a tuple with the MsRpcTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsRpcTimeoutInSeconds

`func (o *Msserver) SetMsRpcTimeoutInSeconds(v int64)`

SetMsRpcTimeoutInSeconds sets MsRpcTimeoutInSeconds field to given value.

### HasMsRpcTimeoutInSeconds

`func (o *Msserver) HasMsRpcTimeoutInSeconds() bool`

HasMsRpcTimeoutInSeconds returns a boolean if a field has been set.

### GetNetworkView

`func (o *Msserver) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Msserver) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Msserver) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Msserver) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetReadOnly

`func (o *Msserver) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Msserver) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Msserver) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *Msserver) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRootAdDomain

`func (o *Msserver) GetRootAdDomain() string`

GetRootAdDomain returns the RootAdDomain field if non-nil, zero value otherwise.

### GetRootAdDomainOk

`func (o *Msserver) GetRootAdDomainOk() (*string, bool)`

GetRootAdDomainOk returns a tuple with the RootAdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAdDomain

`func (o *Msserver) SetRootAdDomain(v string)`

SetRootAdDomain sets RootAdDomain field to given value.

### HasRootAdDomain

`func (o *Msserver) HasRootAdDomain() bool`

HasRootAdDomain returns a boolean if a field has been set.

### GetServerName

`func (o *Msserver) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *Msserver) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *Msserver) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *Msserver) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSynchronizationMinDelay

`func (o *Msserver) GetSynchronizationMinDelay() int64`

GetSynchronizationMinDelay returns the SynchronizationMinDelay field if non-nil, zero value otherwise.

### GetSynchronizationMinDelayOk

`func (o *Msserver) GetSynchronizationMinDelayOk() (*int64, bool)`

GetSynchronizationMinDelayOk returns a tuple with the SynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationMinDelay

`func (o *Msserver) SetSynchronizationMinDelay(v int64)`

SetSynchronizationMinDelay sets SynchronizationMinDelay field to given value.

### HasSynchronizationMinDelay

`func (o *Msserver) HasSynchronizationMinDelay() bool`

HasSynchronizationMinDelay returns a boolean if a field has been set.

### GetSynchronizationStatus

`func (o *Msserver) GetSynchronizationStatus() string`

GetSynchronizationStatus returns the SynchronizationStatus field if non-nil, zero value otherwise.

### GetSynchronizationStatusOk

`func (o *Msserver) GetSynchronizationStatusOk() (*string, bool)`

GetSynchronizationStatusOk returns a tuple with the SynchronizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationStatus

`func (o *Msserver) SetSynchronizationStatus(v string)`

SetSynchronizationStatus sets SynchronizationStatus field to given value.

### HasSynchronizationStatus

`func (o *Msserver) HasSynchronizationStatus() bool`

HasSynchronizationStatus returns a boolean if a field has been set.

### GetSynchronizationStatusDetail

`func (o *Msserver) GetSynchronizationStatusDetail() string`

GetSynchronizationStatusDetail returns the SynchronizationStatusDetail field if non-nil, zero value otherwise.

### GetSynchronizationStatusDetailOk

`func (o *Msserver) GetSynchronizationStatusDetailOk() (*string, bool)`

GetSynchronizationStatusDetailOk returns a tuple with the SynchronizationStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationStatusDetail

`func (o *Msserver) SetSynchronizationStatusDetail(v string)`

SetSynchronizationStatusDetail sets SynchronizationStatusDetail field to given value.

### HasSynchronizationStatusDetail

`func (o *Msserver) HasSynchronizationStatusDetail() bool`

HasSynchronizationStatusDetail returns a boolean if a field has been set.

### GetUseLogDestination

`func (o *Msserver) GetUseLogDestination() bool`

GetUseLogDestination returns the UseLogDestination field if non-nil, zero value otherwise.

### GetUseLogDestinationOk

`func (o *Msserver) GetUseLogDestinationOk() (*bool, bool)`

GetUseLogDestinationOk returns a tuple with the UseLogDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogDestination

`func (o *Msserver) SetUseLogDestination(v bool)`

SetUseLogDestination sets UseLogDestination field to given value.

### HasUseLogDestination

`func (o *Msserver) HasUseLogDestination() bool`

HasUseLogDestination returns a boolean if a field has been set.

### GetUseMsMaxConnection

`func (o *Msserver) GetUseMsMaxConnection() bool`

GetUseMsMaxConnection returns the UseMsMaxConnection field if non-nil, zero value otherwise.

### GetUseMsMaxConnectionOk

`func (o *Msserver) GetUseMsMaxConnectionOk() (*bool, bool)`

GetUseMsMaxConnectionOk returns a tuple with the UseMsMaxConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsMaxConnection

`func (o *Msserver) SetUseMsMaxConnection(v bool)`

SetUseMsMaxConnection sets UseMsMaxConnection field to given value.

### HasUseMsMaxConnection

`func (o *Msserver) HasUseMsMaxConnection() bool`

HasUseMsMaxConnection returns a boolean if a field has been set.

### GetUseMsRpcTimeoutInSeconds

`func (o *Msserver) GetUseMsRpcTimeoutInSeconds() bool`

GetUseMsRpcTimeoutInSeconds returns the UseMsRpcTimeoutInSeconds field if non-nil, zero value otherwise.

### GetUseMsRpcTimeoutInSecondsOk

`func (o *Msserver) GetUseMsRpcTimeoutInSecondsOk() (*bool, bool)`

GetUseMsRpcTimeoutInSecondsOk returns a tuple with the UseMsRpcTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsRpcTimeoutInSeconds

`func (o *Msserver) SetUseMsRpcTimeoutInSeconds(v bool)`

SetUseMsRpcTimeoutInSeconds sets UseMsRpcTimeoutInSeconds field to given value.

### HasUseMsRpcTimeoutInSeconds

`func (o *Msserver) HasUseMsRpcTimeoutInSeconds() bool`

HasUseMsRpcTimeoutInSeconds returns a boolean if a field has been set.

### GetVersion

`func (o *Msserver) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Msserver) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Msserver) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Msserver) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


