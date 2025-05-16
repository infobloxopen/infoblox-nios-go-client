# GridMsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDestination** | Pointer to **string** | The logging of synchronization messages to the syslog or mslog. | [optional] 
**EnableInvalidMac** | Pointer to **bool** | Determines if the invalid MAC address synchronization for DHCP leases and fixed addresses is enabled or not. | [optional] 
**MaxConnection** | Pointer to **int64** | Determines the maximum number of connections to Microsoft servers. | [optional] 
**RpcTimeout** | Pointer to **int64** | Determines the timeout value (in seconds) for RPC connections to all Microsoft servers. | [optional] 
**EnableDhcpMonitoring** | Pointer to **bool** | Determines if the monitoring and control of DHCP service on all Microsoft servers in the Grid is enabled or not. | [optional] 
**EnableDnsMonitoring** | Pointer to **bool** | Determines if the monitoring and control of DNS service on all Microsoft servers in the Grid is enabled or not. | [optional] 
**LdapTimeout** | Pointer to **int64** | Determines an LDAP connection timeout interval (in seconds) for all Microsoft servers. | [optional] 
**DefaultIpSiteLink** | Pointer to **string** | The default IP site link for sites created on NIOS for all Microsoft servers. | [optional] 
**EnableNetworkUsers** | Pointer to **bool** | Determines if the Network Users creation is enabled or not. | [optional] 
**EnableAdUserSync** | Pointer to **bool** | Determines if Active Directory user synchronization for all Microsoft servers in the Grid is enabled or not. | [optional] 
**AdUserDefaultTimeout** | Pointer to **int64** | Determines the default timeout value (in seconds) for Active Directory user synchronization for all Microsoft servers. | [optional] 
**EnableDnsReportsSync** | Pointer to **bool** | Determines if synchronization of DNS reporting data from all Microsoft servers in the Grid is enabled or not. | [optional] 

## Methods

### NewGridMsSetting

`func NewGridMsSetting() *GridMsSetting`

NewGridMsSetting instantiates a new GridMsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridMsSettingWithDefaults

`func NewGridMsSettingWithDefaults() *GridMsSetting`

NewGridMsSettingWithDefaults instantiates a new GridMsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDestination

`func (o *GridMsSetting) GetLogDestination() string`

GetLogDestination returns the LogDestination field if non-nil, zero value otherwise.

### GetLogDestinationOk

`func (o *GridMsSetting) GetLogDestinationOk() (*string, bool)`

GetLogDestinationOk returns a tuple with the LogDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDestination

`func (o *GridMsSetting) SetLogDestination(v string)`

SetLogDestination sets LogDestination field to given value.

### HasLogDestination

`func (o *GridMsSetting) HasLogDestination() bool`

HasLogDestination returns a boolean if a field has been set.

### GetEnableInvalidMac

`func (o *GridMsSetting) GetEnableInvalidMac() bool`

GetEnableInvalidMac returns the EnableInvalidMac field if non-nil, zero value otherwise.

### GetEnableInvalidMacOk

`func (o *GridMsSetting) GetEnableInvalidMacOk() (*bool, bool)`

GetEnableInvalidMacOk returns a tuple with the EnableInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInvalidMac

`func (o *GridMsSetting) SetEnableInvalidMac(v bool)`

SetEnableInvalidMac sets EnableInvalidMac field to given value.

### HasEnableInvalidMac

`func (o *GridMsSetting) HasEnableInvalidMac() bool`

HasEnableInvalidMac returns a boolean if a field has been set.

### GetMaxConnection

`func (o *GridMsSetting) GetMaxConnection() int64`

GetMaxConnection returns the MaxConnection field if non-nil, zero value otherwise.

### GetMaxConnectionOk

`func (o *GridMsSetting) GetMaxConnectionOk() (*int64, bool)`

GetMaxConnectionOk returns a tuple with the MaxConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnection

`func (o *GridMsSetting) SetMaxConnection(v int64)`

SetMaxConnection sets MaxConnection field to given value.

### HasMaxConnection

`func (o *GridMsSetting) HasMaxConnection() bool`

HasMaxConnection returns a boolean if a field has been set.

### GetRpcTimeout

`func (o *GridMsSetting) GetRpcTimeout() int64`

GetRpcTimeout returns the RpcTimeout field if non-nil, zero value otherwise.

### GetRpcTimeoutOk

`func (o *GridMsSetting) GetRpcTimeoutOk() (*int64, bool)`

GetRpcTimeoutOk returns a tuple with the RpcTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcTimeout

`func (o *GridMsSetting) SetRpcTimeout(v int64)`

SetRpcTimeout sets RpcTimeout field to given value.

### HasRpcTimeout

`func (o *GridMsSetting) HasRpcTimeout() bool`

HasRpcTimeout returns a boolean if a field has been set.

### GetEnableDhcpMonitoring

`func (o *GridMsSetting) GetEnableDhcpMonitoring() bool`

GetEnableDhcpMonitoring returns the EnableDhcpMonitoring field if non-nil, zero value otherwise.

### GetEnableDhcpMonitoringOk

`func (o *GridMsSetting) GetEnableDhcpMonitoringOk() (*bool, bool)`

GetEnableDhcpMonitoringOk returns a tuple with the EnableDhcpMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpMonitoring

`func (o *GridMsSetting) SetEnableDhcpMonitoring(v bool)`

SetEnableDhcpMonitoring sets EnableDhcpMonitoring field to given value.

### HasEnableDhcpMonitoring

`func (o *GridMsSetting) HasEnableDhcpMonitoring() bool`

HasEnableDhcpMonitoring returns a boolean if a field has been set.

### GetEnableDnsMonitoring

`func (o *GridMsSetting) GetEnableDnsMonitoring() bool`

GetEnableDnsMonitoring returns the EnableDnsMonitoring field if non-nil, zero value otherwise.

### GetEnableDnsMonitoringOk

`func (o *GridMsSetting) GetEnableDnsMonitoringOk() (*bool, bool)`

GetEnableDnsMonitoringOk returns a tuple with the EnableDnsMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsMonitoring

`func (o *GridMsSetting) SetEnableDnsMonitoring(v bool)`

SetEnableDnsMonitoring sets EnableDnsMonitoring field to given value.

### HasEnableDnsMonitoring

`func (o *GridMsSetting) HasEnableDnsMonitoring() bool`

HasEnableDnsMonitoring returns a boolean if a field has been set.

### GetLdapTimeout

`func (o *GridMsSetting) GetLdapTimeout() int64`

GetLdapTimeout returns the LdapTimeout field if non-nil, zero value otherwise.

### GetLdapTimeoutOk

`func (o *GridMsSetting) GetLdapTimeoutOk() (*int64, bool)`

GetLdapTimeoutOk returns a tuple with the LdapTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTimeout

`func (o *GridMsSetting) SetLdapTimeout(v int64)`

SetLdapTimeout sets LdapTimeout field to given value.

### HasLdapTimeout

`func (o *GridMsSetting) HasLdapTimeout() bool`

HasLdapTimeout returns a boolean if a field has been set.

### GetDefaultIpSiteLink

`func (o *GridMsSetting) GetDefaultIpSiteLink() string`

GetDefaultIpSiteLink returns the DefaultIpSiteLink field if non-nil, zero value otherwise.

### GetDefaultIpSiteLinkOk

`func (o *GridMsSetting) GetDefaultIpSiteLinkOk() (*string, bool)`

GetDefaultIpSiteLinkOk returns a tuple with the DefaultIpSiteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIpSiteLink

`func (o *GridMsSetting) SetDefaultIpSiteLink(v string)`

SetDefaultIpSiteLink sets DefaultIpSiteLink field to given value.

### HasDefaultIpSiteLink

`func (o *GridMsSetting) HasDefaultIpSiteLink() bool`

HasDefaultIpSiteLink returns a boolean if a field has been set.

### GetEnableNetworkUsers

`func (o *GridMsSetting) GetEnableNetworkUsers() bool`

GetEnableNetworkUsers returns the EnableNetworkUsers field if non-nil, zero value otherwise.

### GetEnableNetworkUsersOk

`func (o *GridMsSetting) GetEnableNetworkUsersOk() (*bool, bool)`

GetEnableNetworkUsersOk returns a tuple with the EnableNetworkUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkUsers

`func (o *GridMsSetting) SetEnableNetworkUsers(v bool)`

SetEnableNetworkUsers sets EnableNetworkUsers field to given value.

### HasEnableNetworkUsers

`func (o *GridMsSetting) HasEnableNetworkUsers() bool`

HasEnableNetworkUsers returns a boolean if a field has been set.

### GetEnableAdUserSync

`func (o *GridMsSetting) GetEnableAdUserSync() bool`

GetEnableAdUserSync returns the EnableAdUserSync field if non-nil, zero value otherwise.

### GetEnableAdUserSyncOk

`func (o *GridMsSetting) GetEnableAdUserSyncOk() (*bool, bool)`

GetEnableAdUserSyncOk returns a tuple with the EnableAdUserSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdUserSync

`func (o *GridMsSetting) SetEnableAdUserSync(v bool)`

SetEnableAdUserSync sets EnableAdUserSync field to given value.

### HasEnableAdUserSync

`func (o *GridMsSetting) HasEnableAdUserSync() bool`

HasEnableAdUserSync returns a boolean if a field has been set.

### GetAdUserDefaultTimeout

`func (o *GridMsSetting) GetAdUserDefaultTimeout() int64`

GetAdUserDefaultTimeout returns the AdUserDefaultTimeout field if non-nil, zero value otherwise.

### GetAdUserDefaultTimeoutOk

`func (o *GridMsSetting) GetAdUserDefaultTimeoutOk() (*int64, bool)`

GetAdUserDefaultTimeoutOk returns a tuple with the AdUserDefaultTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserDefaultTimeout

`func (o *GridMsSetting) SetAdUserDefaultTimeout(v int64)`

SetAdUserDefaultTimeout sets AdUserDefaultTimeout field to given value.

### HasAdUserDefaultTimeout

`func (o *GridMsSetting) HasAdUserDefaultTimeout() bool`

HasAdUserDefaultTimeout returns a boolean if a field has been set.

### GetEnableDnsReportsSync

`func (o *GridMsSetting) GetEnableDnsReportsSync() bool`

GetEnableDnsReportsSync returns the EnableDnsReportsSync field if non-nil, zero value otherwise.

### GetEnableDnsReportsSyncOk

`func (o *GridMsSetting) GetEnableDnsReportsSyncOk() (*bool, bool)`

GetEnableDnsReportsSyncOk returns a tuple with the EnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsReportsSync

`func (o *GridMsSetting) SetEnableDnsReportsSync(v bool)`

SetEnableDnsReportsSync sets EnableDnsReportsSync field to given value.

### HasEnableDnsReportsSync

`func (o *GridMsSetting) HasEnableDnsReportsSync() bool`

HasEnableDnsReportsSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


