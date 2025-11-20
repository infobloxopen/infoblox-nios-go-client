# RecordHostIpv4addr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Bootfile** | Pointer to **string** | The name of the boot file the client must download. | [optional] 
**Bootserver** | Pointer to **string** | The IP address or hostname of the boot file server where the boot file is stored. | [optional] 
**ConfigureForDhcp** | Pointer to **bool** | Set this to True to enable the DHCP configuration for this host address. | [optional] 
**DenyBootp** | Pointer to **bool** | Set this to True to disable the BOOTP settings and deny BOOTP boot requests. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | The discovery status of this Host Address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**RecordHostIpv4addrDiscoveredData**](RecordHostIpv4addrDiscoveredData.md) |  | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. You can specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. For example, set a longer lease time if the client downloads an OS (operating system) or configuration file, or set a shorter lease time if the client downloads only configuration changes. Enter the lease time for the preboot execution environment for hosts to boot remotely from a server. | [optional] 
**Host** | Pointer to **string** | The host to which the host address belongs, in FQDN format. It is only present when the host address object is not returned as part of a host. | [optional] [readonly] 
**IgnoreClientRequestedOptions** | Pointer to **bool** | If this field is set to false, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**Ipv4addr** | Pointer to [**RecordHostIpv4addrIpv4addr**](RecordHostIpv4addrIpv4addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**IsInvalidMac** | Pointer to **bool** | This flag reflects whether the MAC address for this host address is invalid. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**LogicFilterRules** | Pointer to [**[]RecordHostIpv4addrLogicFilterRules**](RecordHostIpv4addrLogicFilterRules.md) | This field contains the logic filters to be applied on the this host address. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**Mac** | Pointer to **string** | The MAC address for this host address. | [optional] 
**MatchClient** | Pointer to **string** | Set this to &#39;MAC_ADDRESS&#39; to assign the IP address to the selected host, provided that the MAC address of the requesting host matches the MAC address that you specify in the field. Set this to &#39;RESERVED&#39; to reserve this particular IP address for future use, or if the IP address is statically configured on a system (the Infoblox server does not assign the address from a DHCP request). | [optional] 
**MsAdUserData** | Pointer to [**RecordHostIpv4addrMsAdUserData**](RecordHostIpv4addrMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network of the host address, in FQDN/CIDR format. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which the host address resides. | [optional] [readonly] 
**Nextserver** | Pointer to **string** | The name in FQDN format and/or IPv4 Address of the next server that the host needs to boot. | [optional] 
**Options** | Pointer to [**[]RecordHostIpv4addrOptions**](RecordHostIpv4addrOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The lease time for PXE clients, see *enable_pxe_lease_time* for more information. | [optional] 
**ReservedInterface** | Pointer to **string** | The reference to the reserved interface to which the device belongs. | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseForEaInheritance** | Pointer to **bool** | Set this to True when using this host address for EA inheritance. | [optional] 
**UseIgnoreClientRequestedOptions** | Pointer to **bool** | Use flag for: ignore_client_requested_options | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 

## Methods

### NewRecordHostIpv4addr

`func NewRecordHostIpv4addr() *RecordHostIpv4addr`

NewRecordHostIpv4addr instantiates a new RecordHostIpv4addr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHostIpv4addrWithDefaults

`func NewRecordHostIpv4addrWithDefaults() *RecordHostIpv4addr`

NewRecordHostIpv4addrWithDefaults instantiates a new RecordHostIpv4addr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordHostIpv4addr) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordHostIpv4addr) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordHostIpv4addr) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordHostIpv4addr) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBootfile

`func (o *RecordHostIpv4addr) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *RecordHostIpv4addr) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *RecordHostIpv4addr) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *RecordHostIpv4addr) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *RecordHostIpv4addr) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *RecordHostIpv4addr) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *RecordHostIpv4addr) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *RecordHostIpv4addr) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetConfigureForDhcp

`func (o *RecordHostIpv4addr) GetConfigureForDhcp() bool`

GetConfigureForDhcp returns the ConfigureForDhcp field if non-nil, zero value otherwise.

### GetConfigureForDhcpOk

`func (o *RecordHostIpv4addr) GetConfigureForDhcpOk() (*bool, bool)`

GetConfigureForDhcpOk returns a tuple with the ConfigureForDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureForDhcp

`func (o *RecordHostIpv4addr) SetConfigureForDhcp(v bool)`

SetConfigureForDhcp sets ConfigureForDhcp field to given value.

### HasConfigureForDhcp

`func (o *RecordHostIpv4addr) HasConfigureForDhcp() bool`

HasConfigureForDhcp returns a boolean if a field has been set.

### GetDenyBootp

`func (o *RecordHostIpv4addr) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *RecordHostIpv4addr) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *RecordHostIpv4addr) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *RecordHostIpv4addr) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *RecordHostIpv4addr) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *RecordHostIpv4addr) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *RecordHostIpv4addr) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *RecordHostIpv4addr) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *RecordHostIpv4addr) GetDiscoveredData() RecordHostIpv4addrDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *RecordHostIpv4addr) GetDiscoveredDataOk() (*RecordHostIpv4addrDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *RecordHostIpv4addr) SetDiscoveredData(v RecordHostIpv4addrDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *RecordHostIpv4addr) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *RecordHostIpv4addr) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *RecordHostIpv4addr) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *RecordHostIpv4addr) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *RecordHostIpv4addr) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetHost

`func (o *RecordHostIpv4addr) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RecordHostIpv4addr) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RecordHostIpv4addr) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RecordHostIpv4addr) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIgnoreClientRequestedOptions

`func (o *RecordHostIpv4addr) GetIgnoreClientRequestedOptions() bool`

GetIgnoreClientRequestedOptions returns the IgnoreClientRequestedOptions field if non-nil, zero value otherwise.

### GetIgnoreClientRequestedOptionsOk

`func (o *RecordHostIpv4addr) GetIgnoreClientRequestedOptionsOk() (*bool, bool)`

GetIgnoreClientRequestedOptionsOk returns a tuple with the IgnoreClientRequestedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreClientRequestedOptions

`func (o *RecordHostIpv4addr) SetIgnoreClientRequestedOptions(v bool)`

SetIgnoreClientRequestedOptions sets IgnoreClientRequestedOptions field to given value.

### HasIgnoreClientRequestedOptions

`func (o *RecordHostIpv4addr) HasIgnoreClientRequestedOptions() bool`

HasIgnoreClientRequestedOptions returns a boolean if a field has been set.

### GetIpv4addr

`func (o *RecordHostIpv4addr) GetIpv4addr() RecordHostIpv4addrIpv4addr`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *RecordHostIpv4addr) GetIpv4addrOk() (*RecordHostIpv4addrIpv4addr, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *RecordHostIpv4addr) SetIpv4addr(v RecordHostIpv4addrIpv4addr)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *RecordHostIpv4addr) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *RecordHostIpv4addr) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *RecordHostIpv4addr) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *RecordHostIpv4addr) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *RecordHostIpv4addr) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetIsInvalidMac

`func (o *RecordHostIpv4addr) GetIsInvalidMac() bool`

GetIsInvalidMac returns the IsInvalidMac field if non-nil, zero value otherwise.

### GetIsInvalidMacOk

`func (o *RecordHostIpv4addr) GetIsInvalidMacOk() (*bool, bool)`

GetIsInvalidMacOk returns a tuple with the IsInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalidMac

`func (o *RecordHostIpv4addr) SetIsInvalidMac(v bool)`

SetIsInvalidMac sets IsInvalidMac field to given value.

### HasIsInvalidMac

`func (o *RecordHostIpv4addr) HasIsInvalidMac() bool`

HasIsInvalidMac returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordHostIpv4addr) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordHostIpv4addr) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordHostIpv4addr) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordHostIpv4addr) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *RecordHostIpv4addr) GetLogicFilterRules() []RecordHostIpv4addrLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *RecordHostIpv4addr) GetLogicFilterRulesOk() (*[]RecordHostIpv4addrLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *RecordHostIpv4addr) SetLogicFilterRules(v []RecordHostIpv4addrLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *RecordHostIpv4addr) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMac

`func (o *RecordHostIpv4addr) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RecordHostIpv4addr) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RecordHostIpv4addr) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RecordHostIpv4addr) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMatchClient

`func (o *RecordHostIpv4addr) GetMatchClient() string`

GetMatchClient returns the MatchClient field if non-nil, zero value otherwise.

### GetMatchClientOk

`func (o *RecordHostIpv4addr) GetMatchClientOk() (*string, bool)`

GetMatchClientOk returns a tuple with the MatchClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClient

`func (o *RecordHostIpv4addr) SetMatchClient(v string)`

SetMatchClient sets MatchClient field to given value.

### HasMatchClient

`func (o *RecordHostIpv4addr) HasMatchClient() bool`

HasMatchClient returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *RecordHostIpv4addr) GetMsAdUserData() RecordHostIpv4addrMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *RecordHostIpv4addr) GetMsAdUserDataOk() (*RecordHostIpv4addrMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *RecordHostIpv4addr) SetMsAdUserData(v RecordHostIpv4addrMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *RecordHostIpv4addr) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *RecordHostIpv4addr) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RecordHostIpv4addr) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RecordHostIpv4addr) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RecordHostIpv4addr) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *RecordHostIpv4addr) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *RecordHostIpv4addr) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *RecordHostIpv4addr) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *RecordHostIpv4addr) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *RecordHostIpv4addr) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *RecordHostIpv4addr) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *RecordHostIpv4addr) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *RecordHostIpv4addr) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *RecordHostIpv4addr) GetOptions() []RecordHostIpv4addrOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RecordHostIpv4addr) GetOptionsOk() (*[]RecordHostIpv4addrOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RecordHostIpv4addr) SetOptions(v []RecordHostIpv4addrOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RecordHostIpv4addr) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *RecordHostIpv4addr) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *RecordHostIpv4addr) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *RecordHostIpv4addr) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *RecordHostIpv4addr) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetReservedInterface

`func (o *RecordHostIpv4addr) GetReservedInterface() string`

GetReservedInterface returns the ReservedInterface field if non-nil, zero value otherwise.

### GetReservedInterfaceOk

`func (o *RecordHostIpv4addr) GetReservedInterfaceOk() (*string, bool)`

GetReservedInterfaceOk returns a tuple with the ReservedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedInterface

`func (o *RecordHostIpv4addr) SetReservedInterface(v string)`

SetReservedInterface sets ReservedInterface field to given value.

### HasReservedInterface

`func (o *RecordHostIpv4addr) HasReservedInterface() bool`

HasReservedInterface returns a boolean if a field has been set.

### GetUseBootfile

`func (o *RecordHostIpv4addr) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *RecordHostIpv4addr) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *RecordHostIpv4addr) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *RecordHostIpv4addr) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *RecordHostIpv4addr) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *RecordHostIpv4addr) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *RecordHostIpv4addr) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *RecordHostIpv4addr) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *RecordHostIpv4addr) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *RecordHostIpv4addr) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *RecordHostIpv4addr) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *RecordHostIpv4addr) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseForEaInheritance

`func (o *RecordHostIpv4addr) GetUseForEaInheritance() bool`

GetUseForEaInheritance returns the UseForEaInheritance field if non-nil, zero value otherwise.

### GetUseForEaInheritanceOk

`func (o *RecordHostIpv4addr) GetUseForEaInheritanceOk() (*bool, bool)`

GetUseForEaInheritanceOk returns a tuple with the UseForEaInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForEaInheritance

`func (o *RecordHostIpv4addr) SetUseForEaInheritance(v bool)`

SetUseForEaInheritance sets UseForEaInheritance field to given value.

### HasUseForEaInheritance

`func (o *RecordHostIpv4addr) HasUseForEaInheritance() bool`

HasUseForEaInheritance returns a boolean if a field has been set.

### GetUseIgnoreClientRequestedOptions

`func (o *RecordHostIpv4addr) GetUseIgnoreClientRequestedOptions() bool`

GetUseIgnoreClientRequestedOptions returns the UseIgnoreClientRequestedOptions field if non-nil, zero value otherwise.

### GetUseIgnoreClientRequestedOptionsOk

`func (o *RecordHostIpv4addr) GetUseIgnoreClientRequestedOptionsOk() (*bool, bool)`

GetUseIgnoreClientRequestedOptionsOk returns a tuple with the UseIgnoreClientRequestedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreClientRequestedOptions

`func (o *RecordHostIpv4addr) SetUseIgnoreClientRequestedOptions(v bool)`

SetUseIgnoreClientRequestedOptions sets UseIgnoreClientRequestedOptions field to given value.

### HasUseIgnoreClientRequestedOptions

`func (o *RecordHostIpv4addr) HasUseIgnoreClientRequestedOptions() bool`

HasUseIgnoreClientRequestedOptions returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *RecordHostIpv4addr) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *RecordHostIpv4addr) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *RecordHostIpv4addr) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *RecordHostIpv4addr) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseNextserver

`func (o *RecordHostIpv4addr) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *RecordHostIpv4addr) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *RecordHostIpv4addr) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *RecordHostIpv4addr) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *RecordHostIpv4addr) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *RecordHostIpv4addr) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *RecordHostIpv4addr) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *RecordHostIpv4addr) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *RecordHostIpv4addr) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *RecordHostIpv4addr) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *RecordHostIpv4addr) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *RecordHostIpv4addr) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


