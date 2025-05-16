# AdmingroupNetworkingShowCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowConnectionLimit** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowConnections** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowInterface** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIpRateLimit** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIpv6Bgp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIpv6DisableOnDad** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIpv6Neighbor** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIpv6Ospf** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowLom** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMldVersion** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowNamedRecvSockBufSize** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowNamedTcpClientsLimit** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowNetwork** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowOspf** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowRemoteConsole** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowRoutes** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowStaticRoutes** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowTcpTimestamps** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowTrafficCaptureStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowWinsForwarding** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDefaultRoute** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIproute** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIprule** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIptables** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMtuSize** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowNetworkConnectivity** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowTrafficfiles** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowInterfaceStats** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupNetworkingShowCommands

`func NewAdmingroupNetworkingShowCommands() *AdmingroupNetworkingShowCommands`

NewAdmingroupNetworkingShowCommands instantiates a new AdmingroupNetworkingShowCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupNetworkingShowCommandsWithDefaults

`func NewAdmingroupNetworkingShowCommandsWithDefaults() *AdmingroupNetworkingShowCommands`

NewAdmingroupNetworkingShowCommandsWithDefaults instantiates a new AdmingroupNetworkingShowCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowConnectionLimit

`func (o *AdmingroupNetworkingShowCommands) GetShowConnectionLimit() bool`

GetShowConnectionLimit returns the ShowConnectionLimit field if non-nil, zero value otherwise.

### GetShowConnectionLimitOk

`func (o *AdmingroupNetworkingShowCommands) GetShowConnectionLimitOk() (*bool, bool)`

GetShowConnectionLimitOk returns a tuple with the ShowConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConnectionLimit

`func (o *AdmingroupNetworkingShowCommands) SetShowConnectionLimit(v bool)`

SetShowConnectionLimit sets ShowConnectionLimit field to given value.

### HasShowConnectionLimit

`func (o *AdmingroupNetworkingShowCommands) HasShowConnectionLimit() bool`

HasShowConnectionLimit returns a boolean if a field has been set.

### GetShowConnections

`func (o *AdmingroupNetworkingShowCommands) GetShowConnections() bool`

GetShowConnections returns the ShowConnections field if non-nil, zero value otherwise.

### GetShowConnectionsOk

`func (o *AdmingroupNetworkingShowCommands) GetShowConnectionsOk() (*bool, bool)`

GetShowConnectionsOk returns a tuple with the ShowConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConnections

`func (o *AdmingroupNetworkingShowCommands) SetShowConnections(v bool)`

SetShowConnections sets ShowConnections field to given value.

### HasShowConnections

`func (o *AdmingroupNetworkingShowCommands) HasShowConnections() bool`

HasShowConnections returns a boolean if a field has been set.

### GetShowInterface

`func (o *AdmingroupNetworkingShowCommands) GetShowInterface() bool`

GetShowInterface returns the ShowInterface field if non-nil, zero value otherwise.

### GetShowInterfaceOk

`func (o *AdmingroupNetworkingShowCommands) GetShowInterfaceOk() (*bool, bool)`

GetShowInterfaceOk returns a tuple with the ShowInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInterface

`func (o *AdmingroupNetworkingShowCommands) SetShowInterface(v bool)`

SetShowInterface sets ShowInterface field to given value.

### HasShowInterface

`func (o *AdmingroupNetworkingShowCommands) HasShowInterface() bool`

HasShowInterface returns a boolean if a field has been set.

### GetShowIpRateLimit

`func (o *AdmingroupNetworkingShowCommands) GetShowIpRateLimit() bool`

GetShowIpRateLimit returns the ShowIpRateLimit field if non-nil, zero value otherwise.

### GetShowIpRateLimitOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIpRateLimitOk() (*bool, bool)`

GetShowIpRateLimitOk returns a tuple with the ShowIpRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIpRateLimit

`func (o *AdmingroupNetworkingShowCommands) SetShowIpRateLimit(v bool)`

SetShowIpRateLimit sets ShowIpRateLimit field to given value.

### HasShowIpRateLimit

`func (o *AdmingroupNetworkingShowCommands) HasShowIpRateLimit() bool`

HasShowIpRateLimit returns a boolean if a field has been set.

### GetShowIpv6Bgp

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6Bgp() bool`

GetShowIpv6Bgp returns the ShowIpv6Bgp field if non-nil, zero value otherwise.

### GetShowIpv6BgpOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6BgpOk() (*bool, bool)`

GetShowIpv6BgpOk returns a tuple with the ShowIpv6Bgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIpv6Bgp

`func (o *AdmingroupNetworkingShowCommands) SetShowIpv6Bgp(v bool)`

SetShowIpv6Bgp sets ShowIpv6Bgp field to given value.

### HasShowIpv6Bgp

`func (o *AdmingroupNetworkingShowCommands) HasShowIpv6Bgp() bool`

HasShowIpv6Bgp returns a boolean if a field has been set.

### GetShowIpv6DisableOnDad

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6DisableOnDad() bool`

GetShowIpv6DisableOnDad returns the ShowIpv6DisableOnDad field if non-nil, zero value otherwise.

### GetShowIpv6DisableOnDadOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6DisableOnDadOk() (*bool, bool)`

GetShowIpv6DisableOnDadOk returns a tuple with the ShowIpv6DisableOnDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIpv6DisableOnDad

`func (o *AdmingroupNetworkingShowCommands) SetShowIpv6DisableOnDad(v bool)`

SetShowIpv6DisableOnDad sets ShowIpv6DisableOnDad field to given value.

### HasShowIpv6DisableOnDad

`func (o *AdmingroupNetworkingShowCommands) HasShowIpv6DisableOnDad() bool`

HasShowIpv6DisableOnDad returns a boolean if a field has been set.

### GetShowIpv6Neighbor

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6Neighbor() bool`

GetShowIpv6Neighbor returns the ShowIpv6Neighbor field if non-nil, zero value otherwise.

### GetShowIpv6NeighborOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6NeighborOk() (*bool, bool)`

GetShowIpv6NeighborOk returns a tuple with the ShowIpv6Neighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIpv6Neighbor

`func (o *AdmingroupNetworkingShowCommands) SetShowIpv6Neighbor(v bool)`

SetShowIpv6Neighbor sets ShowIpv6Neighbor field to given value.

### HasShowIpv6Neighbor

`func (o *AdmingroupNetworkingShowCommands) HasShowIpv6Neighbor() bool`

HasShowIpv6Neighbor returns a boolean if a field has been set.

### GetShowIpv6Ospf

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6Ospf() bool`

GetShowIpv6Ospf returns the ShowIpv6Ospf field if non-nil, zero value otherwise.

### GetShowIpv6OspfOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIpv6OspfOk() (*bool, bool)`

GetShowIpv6OspfOk returns a tuple with the ShowIpv6Ospf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIpv6Ospf

`func (o *AdmingroupNetworkingShowCommands) SetShowIpv6Ospf(v bool)`

SetShowIpv6Ospf sets ShowIpv6Ospf field to given value.

### HasShowIpv6Ospf

`func (o *AdmingroupNetworkingShowCommands) HasShowIpv6Ospf() bool`

HasShowIpv6Ospf returns a boolean if a field has been set.

### GetShowLom

`func (o *AdmingroupNetworkingShowCommands) GetShowLom() bool`

GetShowLom returns the ShowLom field if non-nil, zero value otherwise.

### GetShowLomOk

`func (o *AdmingroupNetworkingShowCommands) GetShowLomOk() (*bool, bool)`

GetShowLomOk returns a tuple with the ShowLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLom

`func (o *AdmingroupNetworkingShowCommands) SetShowLom(v bool)`

SetShowLom sets ShowLom field to given value.

### HasShowLom

`func (o *AdmingroupNetworkingShowCommands) HasShowLom() bool`

HasShowLom returns a boolean if a field has been set.

### GetShowMldVersion

`func (o *AdmingroupNetworkingShowCommands) GetShowMldVersion() bool`

GetShowMldVersion returns the ShowMldVersion field if non-nil, zero value otherwise.

### GetShowMldVersionOk

`func (o *AdmingroupNetworkingShowCommands) GetShowMldVersionOk() (*bool, bool)`

GetShowMldVersionOk returns a tuple with the ShowMldVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMldVersion

`func (o *AdmingroupNetworkingShowCommands) SetShowMldVersion(v bool)`

SetShowMldVersion sets ShowMldVersion field to given value.

### HasShowMldVersion

`func (o *AdmingroupNetworkingShowCommands) HasShowMldVersion() bool`

HasShowMldVersion returns a boolean if a field has been set.

### GetShowNamedRecvSockBufSize

`func (o *AdmingroupNetworkingShowCommands) GetShowNamedRecvSockBufSize() bool`

GetShowNamedRecvSockBufSize returns the ShowNamedRecvSockBufSize field if non-nil, zero value otherwise.

### GetShowNamedRecvSockBufSizeOk

`func (o *AdmingroupNetworkingShowCommands) GetShowNamedRecvSockBufSizeOk() (*bool, bool)`

GetShowNamedRecvSockBufSizeOk returns a tuple with the ShowNamedRecvSockBufSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNamedRecvSockBufSize

`func (o *AdmingroupNetworkingShowCommands) SetShowNamedRecvSockBufSize(v bool)`

SetShowNamedRecvSockBufSize sets ShowNamedRecvSockBufSize field to given value.

### HasShowNamedRecvSockBufSize

`func (o *AdmingroupNetworkingShowCommands) HasShowNamedRecvSockBufSize() bool`

HasShowNamedRecvSockBufSize returns a boolean if a field has been set.

### GetShowNamedTcpClientsLimit

`func (o *AdmingroupNetworkingShowCommands) GetShowNamedTcpClientsLimit() bool`

GetShowNamedTcpClientsLimit returns the ShowNamedTcpClientsLimit field if non-nil, zero value otherwise.

### GetShowNamedTcpClientsLimitOk

`func (o *AdmingroupNetworkingShowCommands) GetShowNamedTcpClientsLimitOk() (*bool, bool)`

GetShowNamedTcpClientsLimitOk returns a tuple with the ShowNamedTcpClientsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNamedTcpClientsLimit

`func (o *AdmingroupNetworkingShowCommands) SetShowNamedTcpClientsLimit(v bool)`

SetShowNamedTcpClientsLimit sets ShowNamedTcpClientsLimit field to given value.

### HasShowNamedTcpClientsLimit

`func (o *AdmingroupNetworkingShowCommands) HasShowNamedTcpClientsLimit() bool`

HasShowNamedTcpClientsLimit returns a boolean if a field has been set.

### GetShowNetwork

`func (o *AdmingroupNetworkingShowCommands) GetShowNetwork() bool`

GetShowNetwork returns the ShowNetwork field if non-nil, zero value otherwise.

### GetShowNetworkOk

`func (o *AdmingroupNetworkingShowCommands) GetShowNetworkOk() (*bool, bool)`

GetShowNetworkOk returns a tuple with the ShowNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetwork

`func (o *AdmingroupNetworkingShowCommands) SetShowNetwork(v bool)`

SetShowNetwork sets ShowNetwork field to given value.

### HasShowNetwork

`func (o *AdmingroupNetworkingShowCommands) HasShowNetwork() bool`

HasShowNetwork returns a boolean if a field has been set.

### GetShowOspf

`func (o *AdmingroupNetworkingShowCommands) GetShowOspf() bool`

GetShowOspf returns the ShowOspf field if non-nil, zero value otherwise.

### GetShowOspfOk

`func (o *AdmingroupNetworkingShowCommands) GetShowOspfOk() (*bool, bool)`

GetShowOspfOk returns a tuple with the ShowOspf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOspf

`func (o *AdmingroupNetworkingShowCommands) SetShowOspf(v bool)`

SetShowOspf sets ShowOspf field to given value.

### HasShowOspf

`func (o *AdmingroupNetworkingShowCommands) HasShowOspf() bool`

HasShowOspf returns a boolean if a field has been set.

### GetShowRemoteConsole

`func (o *AdmingroupNetworkingShowCommands) GetShowRemoteConsole() bool`

GetShowRemoteConsole returns the ShowRemoteConsole field if non-nil, zero value otherwise.

### GetShowRemoteConsoleOk

`func (o *AdmingroupNetworkingShowCommands) GetShowRemoteConsoleOk() (*bool, bool)`

GetShowRemoteConsoleOk returns a tuple with the ShowRemoteConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRemoteConsole

`func (o *AdmingroupNetworkingShowCommands) SetShowRemoteConsole(v bool)`

SetShowRemoteConsole sets ShowRemoteConsole field to given value.

### HasShowRemoteConsole

`func (o *AdmingroupNetworkingShowCommands) HasShowRemoteConsole() bool`

HasShowRemoteConsole returns a boolean if a field has been set.

### GetShowRoutes

`func (o *AdmingroupNetworkingShowCommands) GetShowRoutes() bool`

GetShowRoutes returns the ShowRoutes field if non-nil, zero value otherwise.

### GetShowRoutesOk

`func (o *AdmingroupNetworkingShowCommands) GetShowRoutesOk() (*bool, bool)`

GetShowRoutesOk returns a tuple with the ShowRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRoutes

`func (o *AdmingroupNetworkingShowCommands) SetShowRoutes(v bool)`

SetShowRoutes sets ShowRoutes field to given value.

### HasShowRoutes

`func (o *AdmingroupNetworkingShowCommands) HasShowRoutes() bool`

HasShowRoutes returns a boolean if a field has been set.

### GetShowStaticRoutes

`func (o *AdmingroupNetworkingShowCommands) GetShowStaticRoutes() bool`

GetShowStaticRoutes returns the ShowStaticRoutes field if non-nil, zero value otherwise.

### GetShowStaticRoutesOk

`func (o *AdmingroupNetworkingShowCommands) GetShowStaticRoutesOk() (*bool, bool)`

GetShowStaticRoutesOk returns a tuple with the ShowStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowStaticRoutes

`func (o *AdmingroupNetworkingShowCommands) SetShowStaticRoutes(v bool)`

SetShowStaticRoutes sets ShowStaticRoutes field to given value.

### HasShowStaticRoutes

`func (o *AdmingroupNetworkingShowCommands) HasShowStaticRoutes() bool`

HasShowStaticRoutes returns a boolean if a field has been set.

### GetShowTcpTimestamps

`func (o *AdmingroupNetworkingShowCommands) GetShowTcpTimestamps() bool`

GetShowTcpTimestamps returns the ShowTcpTimestamps field if non-nil, zero value otherwise.

### GetShowTcpTimestampsOk

`func (o *AdmingroupNetworkingShowCommands) GetShowTcpTimestampsOk() (*bool, bool)`

GetShowTcpTimestampsOk returns a tuple with the ShowTcpTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTcpTimestamps

`func (o *AdmingroupNetworkingShowCommands) SetShowTcpTimestamps(v bool)`

SetShowTcpTimestamps sets ShowTcpTimestamps field to given value.

### HasShowTcpTimestamps

`func (o *AdmingroupNetworkingShowCommands) HasShowTcpTimestamps() bool`

HasShowTcpTimestamps returns a boolean if a field has been set.

### GetShowTrafficCaptureStatus

`func (o *AdmingroupNetworkingShowCommands) GetShowTrafficCaptureStatus() bool`

GetShowTrafficCaptureStatus returns the ShowTrafficCaptureStatus field if non-nil, zero value otherwise.

### GetShowTrafficCaptureStatusOk

`func (o *AdmingroupNetworkingShowCommands) GetShowTrafficCaptureStatusOk() (*bool, bool)`

GetShowTrafficCaptureStatusOk returns a tuple with the ShowTrafficCaptureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTrafficCaptureStatus

`func (o *AdmingroupNetworkingShowCommands) SetShowTrafficCaptureStatus(v bool)`

SetShowTrafficCaptureStatus sets ShowTrafficCaptureStatus field to given value.

### HasShowTrafficCaptureStatus

`func (o *AdmingroupNetworkingShowCommands) HasShowTrafficCaptureStatus() bool`

HasShowTrafficCaptureStatus returns a boolean if a field has been set.

### GetShowWinsForwarding

`func (o *AdmingroupNetworkingShowCommands) GetShowWinsForwarding() bool`

GetShowWinsForwarding returns the ShowWinsForwarding field if non-nil, zero value otherwise.

### GetShowWinsForwardingOk

`func (o *AdmingroupNetworkingShowCommands) GetShowWinsForwardingOk() (*bool, bool)`

GetShowWinsForwardingOk returns a tuple with the ShowWinsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWinsForwarding

`func (o *AdmingroupNetworkingShowCommands) SetShowWinsForwarding(v bool)`

SetShowWinsForwarding sets ShowWinsForwarding field to given value.

### HasShowWinsForwarding

`func (o *AdmingroupNetworkingShowCommands) HasShowWinsForwarding() bool`

HasShowWinsForwarding returns a boolean if a field has been set.

### GetShowDefaultRoute

`func (o *AdmingroupNetworkingShowCommands) GetShowDefaultRoute() bool`

GetShowDefaultRoute returns the ShowDefaultRoute field if non-nil, zero value otherwise.

### GetShowDefaultRouteOk

`func (o *AdmingroupNetworkingShowCommands) GetShowDefaultRouteOk() (*bool, bool)`

GetShowDefaultRouteOk returns a tuple with the ShowDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDefaultRoute

`func (o *AdmingroupNetworkingShowCommands) SetShowDefaultRoute(v bool)`

SetShowDefaultRoute sets ShowDefaultRoute field to given value.

### HasShowDefaultRoute

`func (o *AdmingroupNetworkingShowCommands) HasShowDefaultRoute() bool`

HasShowDefaultRoute returns a boolean if a field has been set.

### GetShowIproute

`func (o *AdmingroupNetworkingShowCommands) GetShowIproute() bool`

GetShowIproute returns the ShowIproute field if non-nil, zero value otherwise.

### GetShowIprouteOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIprouteOk() (*bool, bool)`

GetShowIprouteOk returns a tuple with the ShowIproute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIproute

`func (o *AdmingroupNetworkingShowCommands) SetShowIproute(v bool)`

SetShowIproute sets ShowIproute field to given value.

### HasShowIproute

`func (o *AdmingroupNetworkingShowCommands) HasShowIproute() bool`

HasShowIproute returns a boolean if a field has been set.

### GetShowIprule

`func (o *AdmingroupNetworkingShowCommands) GetShowIprule() bool`

GetShowIprule returns the ShowIprule field if non-nil, zero value otherwise.

### GetShowIpruleOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIpruleOk() (*bool, bool)`

GetShowIpruleOk returns a tuple with the ShowIprule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIprule

`func (o *AdmingroupNetworkingShowCommands) SetShowIprule(v bool)`

SetShowIprule sets ShowIprule field to given value.

### HasShowIprule

`func (o *AdmingroupNetworkingShowCommands) HasShowIprule() bool`

HasShowIprule returns a boolean if a field has been set.

### GetShowIptables

`func (o *AdmingroupNetworkingShowCommands) GetShowIptables() bool`

GetShowIptables returns the ShowIptables field if non-nil, zero value otherwise.

### GetShowIptablesOk

`func (o *AdmingroupNetworkingShowCommands) GetShowIptablesOk() (*bool, bool)`

GetShowIptablesOk returns a tuple with the ShowIptables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIptables

`func (o *AdmingroupNetworkingShowCommands) SetShowIptables(v bool)`

SetShowIptables sets ShowIptables field to given value.

### HasShowIptables

`func (o *AdmingroupNetworkingShowCommands) HasShowIptables() bool`

HasShowIptables returns a boolean if a field has been set.

### GetShowMtuSize

`func (o *AdmingroupNetworkingShowCommands) GetShowMtuSize() bool`

GetShowMtuSize returns the ShowMtuSize field if non-nil, zero value otherwise.

### GetShowMtuSizeOk

`func (o *AdmingroupNetworkingShowCommands) GetShowMtuSizeOk() (*bool, bool)`

GetShowMtuSizeOk returns a tuple with the ShowMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMtuSize

`func (o *AdmingroupNetworkingShowCommands) SetShowMtuSize(v bool)`

SetShowMtuSize sets ShowMtuSize field to given value.

### HasShowMtuSize

`func (o *AdmingroupNetworkingShowCommands) HasShowMtuSize() bool`

HasShowMtuSize returns a boolean if a field has been set.

### GetShowNetworkConnectivity

`func (o *AdmingroupNetworkingShowCommands) GetShowNetworkConnectivity() bool`

GetShowNetworkConnectivity returns the ShowNetworkConnectivity field if non-nil, zero value otherwise.

### GetShowNetworkConnectivityOk

`func (o *AdmingroupNetworkingShowCommands) GetShowNetworkConnectivityOk() (*bool, bool)`

GetShowNetworkConnectivityOk returns a tuple with the ShowNetworkConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkConnectivity

`func (o *AdmingroupNetworkingShowCommands) SetShowNetworkConnectivity(v bool)`

SetShowNetworkConnectivity sets ShowNetworkConnectivity field to given value.

### HasShowNetworkConnectivity

`func (o *AdmingroupNetworkingShowCommands) HasShowNetworkConnectivity() bool`

HasShowNetworkConnectivity returns a boolean if a field has been set.

### GetShowTrafficfiles

`func (o *AdmingroupNetworkingShowCommands) GetShowTrafficfiles() bool`

GetShowTrafficfiles returns the ShowTrafficfiles field if non-nil, zero value otherwise.

### GetShowTrafficfilesOk

`func (o *AdmingroupNetworkingShowCommands) GetShowTrafficfilesOk() (*bool, bool)`

GetShowTrafficfilesOk returns a tuple with the ShowTrafficfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTrafficfiles

`func (o *AdmingroupNetworkingShowCommands) SetShowTrafficfiles(v bool)`

SetShowTrafficfiles sets ShowTrafficfiles field to given value.

### HasShowTrafficfiles

`func (o *AdmingroupNetworkingShowCommands) HasShowTrafficfiles() bool`

HasShowTrafficfiles returns a boolean if a field has been set.

### GetShowInterfaceStats

`func (o *AdmingroupNetworkingShowCommands) GetShowInterfaceStats() bool`

GetShowInterfaceStats returns the ShowInterfaceStats field if non-nil, zero value otherwise.

### GetShowInterfaceStatsOk

`func (o *AdmingroupNetworkingShowCommands) GetShowInterfaceStatsOk() (*bool, bool)`

GetShowInterfaceStatsOk returns a tuple with the ShowInterfaceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInterfaceStats

`func (o *AdmingroupNetworkingShowCommands) SetShowInterfaceStats(v bool)`

SetShowInterfaceStats sets ShowInterfaceStats field to given value.

### HasShowInterfaceStats

`func (o *AdmingroupNetworkingShowCommands) HasShowInterfaceStats() bool`

HasShowInterfaceStats returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupNetworkingShowCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupNetworkingShowCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupNetworkingShowCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupNetworkingShowCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupNetworkingShowCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupNetworkingShowCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupNetworkingShowCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupNetworkingShowCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


