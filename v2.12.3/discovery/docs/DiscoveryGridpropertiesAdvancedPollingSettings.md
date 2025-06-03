# DiscoveryGridpropertiesAdvancedPollingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TcpScanTechnique** | Pointer to **string** | The TCP scan method. | [optional] 
**PingTimeout** | Pointer to **int64** | The ping timeout in seconds. | [optional] 
**PingRetries** | Pointer to **int64** | The number of ping retries. | [optional] 
**PurgeExpiredDeviceData** | Pointer to **int64** | The number of days a device remains in database after it is no longer found in network. | [optional] 
**EnablePurgeExpiredEndhostData** | Pointer to **bool** | Determines if purge of expired end host data is enabled. | [optional] 
**PurgeExpiredEndhostData** | Pointer to **int64** | The number of days a end host remains in database after it is no longer found in network. | [optional] 
**ArpAggregateLimit** | Pointer to **int64** | The ARP aggregate limit. | [optional] 
**RouteLimit** | Pointer to **int64** | Route limit. | [optional] 
**PingSweepInterval** | Pointer to **int64** | The hourly wait interval between ping sweeps for individual discovery ranges. | [optional] 
**ArpCacheRefreshInterval** | Pointer to **int64** | The refresh interval in seconds for ARP cache. | [optional] 
**PollingAuthenticateSnmpv2cOrLaterOnly** | Pointer to **bool** | Determines if polling and authenticating using SNMPv2c or later is enabled. | [optional] 
**DisableDiscoveryOutsideIpam** | Pointer to **bool** | Determines if discovery of networks that are not in IPAM is disabled. | [optional] 
**DhcpRouterAsSeed** | Pointer to **bool** | Determines if DHCP router is used as seed for discovery. | [optional] 
**SyslogIpamEvents** | Pointer to **bool** | Determines if syslogging of IPAM sync events is enabled. | [optional] 
**SyslogNetworkEvents** | Pointer to **bool** | Determines if syslogging of Network sync events is enabled. | [optional] 

## Methods

### NewDiscoveryGridpropertiesAdvancedPollingSettings

`func NewDiscoveryGridpropertiesAdvancedPollingSettings() *DiscoveryGridpropertiesAdvancedPollingSettings`

NewDiscoveryGridpropertiesAdvancedPollingSettings instantiates a new DiscoveryGridpropertiesAdvancedPollingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesAdvancedPollingSettingsWithDefaults

`func NewDiscoveryGridpropertiesAdvancedPollingSettingsWithDefaults() *DiscoveryGridpropertiesAdvancedPollingSettings`

NewDiscoveryGridpropertiesAdvancedPollingSettingsWithDefaults instantiates a new DiscoveryGridpropertiesAdvancedPollingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcpScanTechnique

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetTcpScanTechnique() string`

GetTcpScanTechnique returns the TcpScanTechnique field if non-nil, zero value otherwise.

### GetTcpScanTechniqueOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetTcpScanTechniqueOk() (*string, bool)`

GetTcpScanTechniqueOk returns a tuple with the TcpScanTechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanTechnique

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetTcpScanTechnique(v string)`

SetTcpScanTechnique sets TcpScanTechnique field to given value.

### HasTcpScanTechnique

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasTcpScanTechnique() bool`

HasTcpScanTechnique returns a boolean if a field has been set.

### GetPingTimeout

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPingTimeout() int64`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPingTimeoutOk() (*int64, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetPingTimeout(v int64)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetPingRetries

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPingRetries() int64`

GetPingRetries returns the PingRetries field if non-nil, zero value otherwise.

### GetPingRetriesOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPingRetriesOk() (*int64, bool)`

GetPingRetriesOk returns a tuple with the PingRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRetries

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetPingRetries(v int64)`

SetPingRetries sets PingRetries field to given value.

### HasPingRetries

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasPingRetries() bool`

HasPingRetries returns a boolean if a field has been set.

### GetPurgeExpiredDeviceData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPurgeExpiredDeviceData() int64`

GetPurgeExpiredDeviceData returns the PurgeExpiredDeviceData field if non-nil, zero value otherwise.

### GetPurgeExpiredDeviceDataOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPurgeExpiredDeviceDataOk() (*int64, bool)`

GetPurgeExpiredDeviceDataOk returns a tuple with the PurgeExpiredDeviceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeExpiredDeviceData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetPurgeExpiredDeviceData(v int64)`

SetPurgeExpiredDeviceData sets PurgeExpiredDeviceData field to given value.

### HasPurgeExpiredDeviceData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasPurgeExpiredDeviceData() bool`

HasPurgeExpiredDeviceData returns a boolean if a field has been set.

### GetEnablePurgeExpiredEndhostData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetEnablePurgeExpiredEndhostData() bool`

GetEnablePurgeExpiredEndhostData returns the EnablePurgeExpiredEndhostData field if non-nil, zero value otherwise.

### GetEnablePurgeExpiredEndhostDataOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetEnablePurgeExpiredEndhostDataOk() (*bool, bool)`

GetEnablePurgeExpiredEndhostDataOk returns a tuple with the EnablePurgeExpiredEndhostData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePurgeExpiredEndhostData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetEnablePurgeExpiredEndhostData(v bool)`

SetEnablePurgeExpiredEndhostData sets EnablePurgeExpiredEndhostData field to given value.

### HasEnablePurgeExpiredEndhostData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasEnablePurgeExpiredEndhostData() bool`

HasEnablePurgeExpiredEndhostData returns a boolean if a field has been set.

### GetPurgeExpiredEndhostData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPurgeExpiredEndhostData() int64`

GetPurgeExpiredEndhostData returns the PurgeExpiredEndhostData field if non-nil, zero value otherwise.

### GetPurgeExpiredEndhostDataOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPurgeExpiredEndhostDataOk() (*int64, bool)`

GetPurgeExpiredEndhostDataOk returns a tuple with the PurgeExpiredEndhostData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeExpiredEndhostData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetPurgeExpiredEndhostData(v int64)`

SetPurgeExpiredEndhostData sets PurgeExpiredEndhostData field to given value.

### HasPurgeExpiredEndhostData

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasPurgeExpiredEndhostData() bool`

HasPurgeExpiredEndhostData returns a boolean if a field has been set.

### GetArpAggregateLimit

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetArpAggregateLimit() int64`

GetArpAggregateLimit returns the ArpAggregateLimit field if non-nil, zero value otherwise.

### GetArpAggregateLimitOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetArpAggregateLimitOk() (*int64, bool)`

GetArpAggregateLimitOk returns a tuple with the ArpAggregateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpAggregateLimit

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetArpAggregateLimit(v int64)`

SetArpAggregateLimit sets ArpAggregateLimit field to given value.

### HasArpAggregateLimit

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasArpAggregateLimit() bool`

HasArpAggregateLimit returns a boolean if a field has been set.

### GetRouteLimit

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetRouteLimit() int64`

GetRouteLimit returns the RouteLimit field if non-nil, zero value otherwise.

### GetRouteLimitOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetRouteLimitOk() (*int64, bool)`

GetRouteLimitOk returns a tuple with the RouteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteLimit

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetRouteLimit(v int64)`

SetRouteLimit sets RouteLimit field to given value.

### HasRouteLimit

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasRouteLimit() bool`

HasRouteLimit returns a boolean if a field has been set.

### GetPingSweepInterval

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPingSweepInterval() int64`

GetPingSweepInterval returns the PingSweepInterval field if non-nil, zero value otherwise.

### GetPingSweepIntervalOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPingSweepIntervalOk() (*int64, bool)`

GetPingSweepIntervalOk returns a tuple with the PingSweepInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweepInterval

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetPingSweepInterval(v int64)`

SetPingSweepInterval sets PingSweepInterval field to given value.

### HasPingSweepInterval

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasPingSweepInterval() bool`

HasPingSweepInterval returns a boolean if a field has been set.

### GetArpCacheRefreshInterval

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetArpCacheRefreshInterval() int64`

GetArpCacheRefreshInterval returns the ArpCacheRefreshInterval field if non-nil, zero value otherwise.

### GetArpCacheRefreshIntervalOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetArpCacheRefreshIntervalOk() (*int64, bool)`

GetArpCacheRefreshIntervalOk returns a tuple with the ArpCacheRefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpCacheRefreshInterval

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetArpCacheRefreshInterval(v int64)`

SetArpCacheRefreshInterval sets ArpCacheRefreshInterval field to given value.

### HasArpCacheRefreshInterval

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasArpCacheRefreshInterval() bool`

HasArpCacheRefreshInterval returns a boolean if a field has been set.

### GetPollingAuthenticateSnmpv2cOrLaterOnly

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPollingAuthenticateSnmpv2cOrLaterOnly() bool`

GetPollingAuthenticateSnmpv2cOrLaterOnly returns the PollingAuthenticateSnmpv2cOrLaterOnly field if non-nil, zero value otherwise.

### GetPollingAuthenticateSnmpv2cOrLaterOnlyOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetPollingAuthenticateSnmpv2cOrLaterOnlyOk() (*bool, bool)`

GetPollingAuthenticateSnmpv2cOrLaterOnlyOk returns a tuple with the PollingAuthenticateSnmpv2cOrLaterOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingAuthenticateSnmpv2cOrLaterOnly

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetPollingAuthenticateSnmpv2cOrLaterOnly(v bool)`

SetPollingAuthenticateSnmpv2cOrLaterOnly sets PollingAuthenticateSnmpv2cOrLaterOnly field to given value.

### HasPollingAuthenticateSnmpv2cOrLaterOnly

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasPollingAuthenticateSnmpv2cOrLaterOnly() bool`

HasPollingAuthenticateSnmpv2cOrLaterOnly returns a boolean if a field has been set.

### GetDisableDiscoveryOutsideIpam

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetDisableDiscoveryOutsideIpam() bool`

GetDisableDiscoveryOutsideIpam returns the DisableDiscoveryOutsideIpam field if non-nil, zero value otherwise.

### GetDisableDiscoveryOutsideIpamOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetDisableDiscoveryOutsideIpamOk() (*bool, bool)`

GetDisableDiscoveryOutsideIpamOk returns a tuple with the DisableDiscoveryOutsideIpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDiscoveryOutsideIpam

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetDisableDiscoveryOutsideIpam(v bool)`

SetDisableDiscoveryOutsideIpam sets DisableDiscoveryOutsideIpam field to given value.

### HasDisableDiscoveryOutsideIpam

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasDisableDiscoveryOutsideIpam() bool`

HasDisableDiscoveryOutsideIpam returns a boolean if a field has been set.

### GetDhcpRouterAsSeed

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetDhcpRouterAsSeed() bool`

GetDhcpRouterAsSeed returns the DhcpRouterAsSeed field if non-nil, zero value otherwise.

### GetDhcpRouterAsSeedOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetDhcpRouterAsSeedOk() (*bool, bool)`

GetDhcpRouterAsSeedOk returns a tuple with the DhcpRouterAsSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRouterAsSeed

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetDhcpRouterAsSeed(v bool)`

SetDhcpRouterAsSeed sets DhcpRouterAsSeed field to given value.

### HasDhcpRouterAsSeed

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasDhcpRouterAsSeed() bool`

HasDhcpRouterAsSeed returns a boolean if a field has been set.

### GetSyslogIpamEvents

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetSyslogIpamEvents() bool`

GetSyslogIpamEvents returns the SyslogIpamEvents field if non-nil, zero value otherwise.

### GetSyslogIpamEventsOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetSyslogIpamEventsOk() (*bool, bool)`

GetSyslogIpamEventsOk returns a tuple with the SyslogIpamEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogIpamEvents

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetSyslogIpamEvents(v bool)`

SetSyslogIpamEvents sets SyslogIpamEvents field to given value.

### HasSyslogIpamEvents

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasSyslogIpamEvents() bool`

HasSyslogIpamEvents returns a boolean if a field has been set.

### GetSyslogNetworkEvents

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetSyslogNetworkEvents() bool`

GetSyslogNetworkEvents returns the SyslogNetworkEvents field if non-nil, zero value otherwise.

### GetSyslogNetworkEventsOk

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) GetSyslogNetworkEventsOk() (*bool, bool)`

GetSyslogNetworkEventsOk returns a tuple with the SyslogNetworkEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNetworkEvents

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) SetSyslogNetworkEvents(v bool)`

SetSyslogNetworkEvents sets SyslogNetworkEvents field to given value.

### HasSyslogNetworkEvents

`func (o *DiscoveryGridpropertiesAdvancedPollingSettings) HasSyslogNetworkEvents() bool`

HasSyslogNetworkEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


