# GetRecordHostIpv4addrResponse

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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**RecordHostIpv4addr**](RecordHostIpv4addr.md) |  | [optional] 

## Methods

### NewGetRecordHostIpv4addrResponse

`func NewGetRecordHostIpv4addrResponse() *GetRecordHostIpv4addrResponse`

NewGetRecordHostIpv4addrResponse instantiates a new GetRecordHostIpv4addrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordHostIpv4addrResponseWithDefaults

`func NewGetRecordHostIpv4addrResponseWithDefaults() *GetRecordHostIpv4addrResponse`

NewGetRecordHostIpv4addrResponseWithDefaults instantiates a new GetRecordHostIpv4addrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordHostIpv4addrResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordHostIpv4addrResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordHostIpv4addrResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordHostIpv4addrResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBootfile

`func (o *GetRecordHostIpv4addrResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetRecordHostIpv4addrResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetRecordHostIpv4addrResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetRecordHostIpv4addrResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetRecordHostIpv4addrResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetRecordHostIpv4addrResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetRecordHostIpv4addrResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetRecordHostIpv4addrResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetConfigureForDhcp

`func (o *GetRecordHostIpv4addrResponse) GetConfigureForDhcp() bool`

GetConfigureForDhcp returns the ConfigureForDhcp field if non-nil, zero value otherwise.

### GetConfigureForDhcpOk

`func (o *GetRecordHostIpv4addrResponse) GetConfigureForDhcpOk() (*bool, bool)`

GetConfigureForDhcpOk returns a tuple with the ConfigureForDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureForDhcp

`func (o *GetRecordHostIpv4addrResponse) SetConfigureForDhcp(v bool)`

SetConfigureForDhcp sets ConfigureForDhcp field to given value.

### HasConfigureForDhcp

`func (o *GetRecordHostIpv4addrResponse) HasConfigureForDhcp() bool`

HasConfigureForDhcp returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetRecordHostIpv4addrResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetRecordHostIpv4addrResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetRecordHostIpv4addrResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetRecordHostIpv4addrResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetRecordHostIpv4addrResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetRecordHostIpv4addrResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetRecordHostIpv4addrResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetRecordHostIpv4addrResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *GetRecordHostIpv4addrResponse) GetDiscoveredData() RecordHostIpv4addrDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *GetRecordHostIpv4addrResponse) GetDiscoveredDataOk() (*RecordHostIpv4addrDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *GetRecordHostIpv4addrResponse) SetDiscoveredData(v RecordHostIpv4addrDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *GetRecordHostIpv4addrResponse) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetRecordHostIpv4addrResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetHost

`func (o *GetRecordHostIpv4addrResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetRecordHostIpv4addrResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetRecordHostIpv4addrResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetRecordHostIpv4addrResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIgnoreClientRequestedOptions

`func (o *GetRecordHostIpv4addrResponse) GetIgnoreClientRequestedOptions() bool`

GetIgnoreClientRequestedOptions returns the IgnoreClientRequestedOptions field if non-nil, zero value otherwise.

### GetIgnoreClientRequestedOptionsOk

`func (o *GetRecordHostIpv4addrResponse) GetIgnoreClientRequestedOptionsOk() (*bool, bool)`

GetIgnoreClientRequestedOptionsOk returns a tuple with the IgnoreClientRequestedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreClientRequestedOptions

`func (o *GetRecordHostIpv4addrResponse) SetIgnoreClientRequestedOptions(v bool)`

SetIgnoreClientRequestedOptions sets IgnoreClientRequestedOptions field to given value.

### HasIgnoreClientRequestedOptions

`func (o *GetRecordHostIpv4addrResponse) HasIgnoreClientRequestedOptions() bool`

HasIgnoreClientRequestedOptions returns a boolean if a field has been set.

### GetIpv4addr

`func (o *GetRecordHostIpv4addrResponse) GetIpv4addr() RecordHostIpv4addrIpv4addr`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *GetRecordHostIpv4addrResponse) GetIpv4addrOk() (*RecordHostIpv4addrIpv4addr, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *GetRecordHostIpv4addrResponse) SetIpv4addr(v RecordHostIpv4addrIpv4addr)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *GetRecordHostIpv4addrResponse) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *GetRecordHostIpv4addrResponse) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *GetRecordHostIpv4addrResponse) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *GetRecordHostIpv4addrResponse) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *GetRecordHostIpv4addrResponse) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetIsInvalidMac

`func (o *GetRecordHostIpv4addrResponse) GetIsInvalidMac() bool`

GetIsInvalidMac returns the IsInvalidMac field if non-nil, zero value otherwise.

### GetIsInvalidMacOk

`func (o *GetRecordHostIpv4addrResponse) GetIsInvalidMacOk() (*bool, bool)`

GetIsInvalidMacOk returns a tuple with the IsInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalidMac

`func (o *GetRecordHostIpv4addrResponse) SetIsInvalidMac(v bool)`

SetIsInvalidMac sets IsInvalidMac field to given value.

### HasIsInvalidMac

`func (o *GetRecordHostIpv4addrResponse) HasIsInvalidMac() bool`

HasIsInvalidMac returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordHostIpv4addrResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordHostIpv4addrResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordHostIpv4addrResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordHostIpv4addrResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetRecordHostIpv4addrResponse) GetLogicFilterRules() []RecordHostIpv4addrLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetRecordHostIpv4addrResponse) GetLogicFilterRulesOk() (*[]RecordHostIpv4addrLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetRecordHostIpv4addrResponse) SetLogicFilterRules(v []RecordHostIpv4addrLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetRecordHostIpv4addrResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMac

`func (o *GetRecordHostIpv4addrResponse) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetRecordHostIpv4addrResponse) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetRecordHostIpv4addrResponse) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetRecordHostIpv4addrResponse) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMatchClient

`func (o *GetRecordHostIpv4addrResponse) GetMatchClient() string`

GetMatchClient returns the MatchClient field if non-nil, zero value otherwise.

### GetMatchClientOk

`func (o *GetRecordHostIpv4addrResponse) GetMatchClientOk() (*string, bool)`

GetMatchClientOk returns a tuple with the MatchClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClient

`func (o *GetRecordHostIpv4addrResponse) SetMatchClient(v string)`

SetMatchClient sets MatchClient field to given value.

### HasMatchClient

`func (o *GetRecordHostIpv4addrResponse) HasMatchClient() bool`

HasMatchClient returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetRecordHostIpv4addrResponse) GetMsAdUserData() RecordHostIpv4addrMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetRecordHostIpv4addrResponse) GetMsAdUserDataOk() (*RecordHostIpv4addrMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetRecordHostIpv4addrResponse) SetMsAdUserData(v RecordHostIpv4addrMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetRecordHostIpv4addrResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *GetRecordHostIpv4addrResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetRecordHostIpv4addrResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetRecordHostIpv4addrResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetRecordHostIpv4addrResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetRecordHostIpv4addrResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetRecordHostIpv4addrResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetRecordHostIpv4addrResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetRecordHostIpv4addrResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *GetRecordHostIpv4addrResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetRecordHostIpv4addrResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetRecordHostIpv4addrResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetRecordHostIpv4addrResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *GetRecordHostIpv4addrResponse) GetOptions() []RecordHostIpv4addrOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetRecordHostIpv4addrResponse) GetOptionsOk() (*[]RecordHostIpv4addrOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetRecordHostIpv4addrResponse) SetOptions(v []RecordHostIpv4addrOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetRecordHostIpv4addrResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetRecordHostIpv4addrResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetReservedInterface

`func (o *GetRecordHostIpv4addrResponse) GetReservedInterface() string`

GetReservedInterface returns the ReservedInterface field if non-nil, zero value otherwise.

### GetReservedInterfaceOk

`func (o *GetRecordHostIpv4addrResponse) GetReservedInterfaceOk() (*string, bool)`

GetReservedInterfaceOk returns a tuple with the ReservedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedInterface

`func (o *GetRecordHostIpv4addrResponse) SetReservedInterface(v string)`

SetReservedInterface sets ReservedInterface field to given value.

### HasReservedInterface

`func (o *GetRecordHostIpv4addrResponse) HasReservedInterface() bool`

HasReservedInterface returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetRecordHostIpv4addrResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetRecordHostIpv4addrResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetRecordHostIpv4addrResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetRecordHostIpv4addrResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetRecordHostIpv4addrResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetRecordHostIpv4addrResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetRecordHostIpv4addrResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetRecordHostIpv4addrResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetRecordHostIpv4addrResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetRecordHostIpv4addrResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetRecordHostIpv4addrResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetRecordHostIpv4addrResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseForEaInheritance

`func (o *GetRecordHostIpv4addrResponse) GetUseForEaInheritance() bool`

GetUseForEaInheritance returns the UseForEaInheritance field if non-nil, zero value otherwise.

### GetUseForEaInheritanceOk

`func (o *GetRecordHostIpv4addrResponse) GetUseForEaInheritanceOk() (*bool, bool)`

GetUseForEaInheritanceOk returns a tuple with the UseForEaInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForEaInheritance

`func (o *GetRecordHostIpv4addrResponse) SetUseForEaInheritance(v bool)`

SetUseForEaInheritance sets UseForEaInheritance field to given value.

### HasUseForEaInheritance

`func (o *GetRecordHostIpv4addrResponse) HasUseForEaInheritance() bool`

HasUseForEaInheritance returns a boolean if a field has been set.

### GetUseIgnoreClientRequestedOptions

`func (o *GetRecordHostIpv4addrResponse) GetUseIgnoreClientRequestedOptions() bool`

GetUseIgnoreClientRequestedOptions returns the UseIgnoreClientRequestedOptions field if non-nil, zero value otherwise.

### GetUseIgnoreClientRequestedOptionsOk

`func (o *GetRecordHostIpv4addrResponse) GetUseIgnoreClientRequestedOptionsOk() (*bool, bool)`

GetUseIgnoreClientRequestedOptionsOk returns a tuple with the UseIgnoreClientRequestedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreClientRequestedOptions

`func (o *GetRecordHostIpv4addrResponse) SetUseIgnoreClientRequestedOptions(v bool)`

SetUseIgnoreClientRequestedOptions sets UseIgnoreClientRequestedOptions field to given value.

### HasUseIgnoreClientRequestedOptions

`func (o *GetRecordHostIpv4addrResponse) HasUseIgnoreClientRequestedOptions() bool`

HasUseIgnoreClientRequestedOptions returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetRecordHostIpv4addrResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetRecordHostIpv4addrResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetRecordHostIpv4addrResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetRecordHostIpv4addrResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetRecordHostIpv4addrResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetRecordHostIpv4addrResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetRecordHostIpv4addrResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetRecordHostIpv4addrResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetRecordHostIpv4addrResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetRecordHostIpv4addrResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetRecordHostIpv4addrResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetRecordHostIpv4addrResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetRecordHostIpv4addrResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetRecordHostIpv4addrResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUuid

`func (o *GetRecordHostIpv4addrResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetRecordHostIpv4addrResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetRecordHostIpv4addrResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetRecordHostIpv4addrResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordHostIpv4addrResponse) GetResult() RecordHostIpv4addr`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordHostIpv4addrResponse) GetResultOk() (*RecordHostIpv4addr, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordHostIpv4addrResponse) SetResult(v RecordHostIpv4addr)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordHostIpv4addrResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


