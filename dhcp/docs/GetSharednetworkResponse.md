# GetSharednetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Authority** | Pointer to **bool** | Authority for the shared network. | [optional] 
**Bootfile** | Pointer to **string** | The bootfile name for the shared network. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The bootserver address for the shared network. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**Comment** | Pointer to **string** | Comment for the shared network, maximum 256 characters. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DdnsServerAlwaysUpdates** | Pointer to **bool** | This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_use_option81 is True. | [optional] 
**DdnsTtl** | Pointer to **int64** | The DNS update Time to Live (TTL) value of a shared network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**DdnsUpdateFixedAddresses** | Pointer to **bool** | By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records. | [optional] 
**DdnsUseOption81** | Pointer to **bool** | The support for DHCP Option 81 at the shared network level. | [optional] 
**DenyBootp** | Pointer to **bool** | If set to true, BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of the networks belonging to the shared network multiplied by 1000. This is the percentage of the total number of available IP addresses from all the networks belonging to the shared network versus the total number of all IP addresses in all of the networks in the shared network. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of the shared network. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether a shared network is disabled or not. When this is set to False, the shared network is enabled. | [optional] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the shared network. | [optional] [readonly] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a shared network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IgnoreClientIdentifier** | Pointer to **bool** | If set to true, the client identifier will be ignored. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**IgnoreId** | Pointer to **string** | Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \&quot;NONE\&quot;, \&quot;CLIENT\&quot;, or \&quot;MACADDR\&quot;. The default is \&quot;NONE\&quot;. | [optional] 
**IgnoreMacAddresses** | Pointer to **[]string** | A list of MAC addresses the appliance will ignore. | [optional] 
**LeaseScavengeTime** | Pointer to **int32** | An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogicFilterRules** | Pointer to [**[]SharednetworkLogicFilterRules**](SharednetworkLogicFilterRules.md) | This field contains the logic filters to be applied on the this shared network. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**MsAdUserData** | Pointer to [**SharednetworkMsAdUserData**](SharednetworkMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the IPv6 Shared Network. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this shared network resides. | [optional] 
**Networks** | Pointer to **[]string** | A list of networks belonging to the shared network Each individual list item must be specified as an object containing a &#39;_ref&#39; parameter to a network reference, for example:: [{ \&quot;_ref\&quot;: \&quot;network/ZG5zLm5ldHdvcmskMTAuMwLvMTYvMA\&quot;, }] if the reference of the wanted network is not known, it is possible to specify search parameters for the network instead in the following way:: [{ \&quot;_ref\&quot;: { &#39;network&#39;: &#39;10.0.0.0/8&#39;, } }] note that in this case the search must match exactly one network for the assignment to be successful. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address of the next server that the host needs to boot. | [optional] 
**Options** | Pointer to [**[]SharednetworkOptions**](SharednetworkOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value of a shared network object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in the shared network. | [optional] [readonly] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in the shared network. | [optional] [readonly] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseAuthority** | Pointer to **bool** | Use flag for: authority | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDdnsTtl** | Pointer to **bool** | Use flag for: ddns_ttl | [optional] 
**UseDdnsUpdateFixedAddresses** | Pointer to **bool** | Use flag for: ddns_update_fixed_addresses | [optional] 
**UseDdnsUseOption81** | Pointer to **bool** | Use flag for: ddns_use_option81 | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseIgnoreClientIdentifier** | Pointer to **bool** | Use flag for: ignore_client_identifier | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseIgnoreId** | Pointer to **bool** | Use flag for: ignore_id | [optional] 
**UseLeaseScavengeTime** | Pointer to **bool** | Use flag for: lease_scavenge_time | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 
**Result** | Pointer to [**Sharednetwork**](Sharednetwork.md) |  | [optional] 

## Methods

### NewGetSharednetworkResponse

`func NewGetSharednetworkResponse() *GetSharednetworkResponse`

NewGetSharednetworkResponse instantiates a new GetSharednetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharednetworkResponseWithDefaults

`func NewGetSharednetworkResponseWithDefaults() *GetSharednetworkResponse`

NewGetSharednetworkResponseWithDefaults instantiates a new GetSharednetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSharednetworkResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSharednetworkResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSharednetworkResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSharednetworkResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthority

`func (o *GetSharednetworkResponse) GetAuthority() bool`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GetSharednetworkResponse) GetAuthorityOk() (*bool, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GetSharednetworkResponse) SetAuthority(v bool)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GetSharednetworkResponse) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetBootfile

`func (o *GetSharednetworkResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetSharednetworkResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetSharednetworkResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetSharednetworkResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetSharednetworkResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetSharednetworkResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetSharednetworkResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetSharednetworkResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetComment

`func (o *GetSharednetworkResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSharednetworkResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSharednetworkResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSharednetworkResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *GetSharednetworkResponse) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *GetSharednetworkResponse) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *GetSharednetworkResponse) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *GetSharednetworkResponse) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDdnsServerAlwaysUpdates

`func (o *GetSharednetworkResponse) GetDdnsServerAlwaysUpdates() bool`

GetDdnsServerAlwaysUpdates returns the DdnsServerAlwaysUpdates field if non-nil, zero value otherwise.

### GetDdnsServerAlwaysUpdatesOk

`func (o *GetSharednetworkResponse) GetDdnsServerAlwaysUpdatesOk() (*bool, bool)`

GetDdnsServerAlwaysUpdatesOk returns a tuple with the DdnsServerAlwaysUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsServerAlwaysUpdates

`func (o *GetSharednetworkResponse) SetDdnsServerAlwaysUpdates(v bool)`

SetDdnsServerAlwaysUpdates sets DdnsServerAlwaysUpdates field to given value.

### HasDdnsServerAlwaysUpdates

`func (o *GetSharednetworkResponse) HasDdnsServerAlwaysUpdates() bool`

HasDdnsServerAlwaysUpdates returns a boolean if a field has been set.

### GetDdnsTtl

`func (o *GetSharednetworkResponse) GetDdnsTtl() int64`

GetDdnsTtl returns the DdnsTtl field if non-nil, zero value otherwise.

### GetDdnsTtlOk

`func (o *GetSharednetworkResponse) GetDdnsTtlOk() (*int64, bool)`

GetDdnsTtlOk returns a tuple with the DdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtl

`func (o *GetSharednetworkResponse) SetDdnsTtl(v int64)`

SetDdnsTtl sets DdnsTtl field to given value.

### HasDdnsTtl

`func (o *GetSharednetworkResponse) HasDdnsTtl() bool`

HasDdnsTtl returns a boolean if a field has been set.

### GetDdnsUpdateFixedAddresses

`func (o *GetSharednetworkResponse) GetDdnsUpdateFixedAddresses() bool`

GetDdnsUpdateFixedAddresses returns the DdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetDdnsUpdateFixedAddressesOk

`func (o *GetSharednetworkResponse) GetDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetDdnsUpdateFixedAddressesOk returns a tuple with the DdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateFixedAddresses

`func (o *GetSharednetworkResponse) SetDdnsUpdateFixedAddresses(v bool)`

SetDdnsUpdateFixedAddresses sets DdnsUpdateFixedAddresses field to given value.

### HasDdnsUpdateFixedAddresses

`func (o *GetSharednetworkResponse) HasDdnsUpdateFixedAddresses() bool`

HasDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetDdnsUseOption81

`func (o *GetSharednetworkResponse) GetDdnsUseOption81() bool`

GetDdnsUseOption81 returns the DdnsUseOption81 field if non-nil, zero value otherwise.

### GetDdnsUseOption81Ok

`func (o *GetSharednetworkResponse) GetDdnsUseOption81Ok() (*bool, bool)`

GetDdnsUseOption81Ok returns a tuple with the DdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseOption81

`func (o *GetSharednetworkResponse) SetDdnsUseOption81(v bool)`

SetDdnsUseOption81 sets DdnsUseOption81 field to given value.

### HasDdnsUseOption81

`func (o *GetSharednetworkResponse) HasDdnsUseOption81() bool`

HasDdnsUseOption81 returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetSharednetworkResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetSharednetworkResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetSharednetworkResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetSharednetworkResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *GetSharednetworkResponse) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *GetSharednetworkResponse) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *GetSharednetworkResponse) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *GetSharednetworkResponse) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *GetSharednetworkResponse) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *GetSharednetworkResponse) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *GetSharednetworkResponse) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *GetSharednetworkResponse) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *GetSharednetworkResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetSharednetworkResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetSharednetworkResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetSharednetworkResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *GetSharednetworkResponse) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *GetSharednetworkResponse) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *GetSharednetworkResponse) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *GetSharednetworkResponse) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetSharednetworkResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetSharednetworkResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetSharednetworkResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetSharednetworkResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetSharednetworkResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetSharednetworkResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetSharednetworkResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetSharednetworkResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetSharednetworkResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetSharednetworkResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetSharednetworkResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetSharednetworkResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIgnoreClientIdentifier

`func (o *GetSharednetworkResponse) GetIgnoreClientIdentifier() bool`

GetIgnoreClientIdentifier returns the IgnoreClientIdentifier field if non-nil, zero value otherwise.

### GetIgnoreClientIdentifierOk

`func (o *GetSharednetworkResponse) GetIgnoreClientIdentifierOk() (*bool, bool)`

GetIgnoreClientIdentifierOk returns a tuple with the IgnoreClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreClientIdentifier

`func (o *GetSharednetworkResponse) SetIgnoreClientIdentifier(v bool)`

SetIgnoreClientIdentifier sets IgnoreClientIdentifier field to given value.

### HasIgnoreClientIdentifier

`func (o *GetSharednetworkResponse) HasIgnoreClientIdentifier() bool`

HasIgnoreClientIdentifier returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GetSharednetworkResponse) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GetSharednetworkResponse) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GetSharednetworkResponse) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GetSharednetworkResponse) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIgnoreId

`func (o *GetSharednetworkResponse) GetIgnoreId() string`

GetIgnoreId returns the IgnoreId field if non-nil, zero value otherwise.

### GetIgnoreIdOk

`func (o *GetSharednetworkResponse) GetIgnoreIdOk() (*string, bool)`

GetIgnoreIdOk returns a tuple with the IgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreId

`func (o *GetSharednetworkResponse) SetIgnoreId(v string)`

SetIgnoreId sets IgnoreId field to given value.

### HasIgnoreId

`func (o *GetSharednetworkResponse) HasIgnoreId() bool`

HasIgnoreId returns a boolean if a field has been set.

### GetIgnoreMacAddresses

`func (o *GetSharednetworkResponse) GetIgnoreMacAddresses() []string`

GetIgnoreMacAddresses returns the IgnoreMacAddresses field if non-nil, zero value otherwise.

### GetIgnoreMacAddressesOk

`func (o *GetSharednetworkResponse) GetIgnoreMacAddressesOk() (*[]string, bool)`

GetIgnoreMacAddressesOk returns a tuple with the IgnoreMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMacAddresses

`func (o *GetSharednetworkResponse) SetIgnoreMacAddresses(v []string)`

SetIgnoreMacAddresses sets IgnoreMacAddresses field to given value.

### HasIgnoreMacAddresses

`func (o *GetSharednetworkResponse) HasIgnoreMacAddresses() bool`

HasIgnoreMacAddresses returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *GetSharednetworkResponse) GetLeaseScavengeTime() int32`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *GetSharednetworkResponse) GetLeaseScavengeTimeOk() (*int32, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *GetSharednetworkResponse) SetLeaseScavengeTime(v int32)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *GetSharednetworkResponse) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetSharednetworkResponse) GetLogicFilterRules() []SharednetworkLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetSharednetworkResponse) GetLogicFilterRulesOk() (*[]SharednetworkLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetSharednetworkResponse) SetLogicFilterRules(v []SharednetworkLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetSharednetworkResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetSharednetworkResponse) GetMsAdUserData() SharednetworkMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetSharednetworkResponse) GetMsAdUserDataOk() (*SharednetworkMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetSharednetworkResponse) SetMsAdUserData(v SharednetworkMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetSharednetworkResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetSharednetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSharednetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSharednetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSharednetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetSharednetworkResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetSharednetworkResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetSharednetworkResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetSharednetworkResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworks

`func (o *GetSharednetworkResponse) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetSharednetworkResponse) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetSharednetworkResponse) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetSharednetworkResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNextserver

`func (o *GetSharednetworkResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetSharednetworkResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetSharednetworkResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetSharednetworkResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *GetSharednetworkResponse) GetOptions() []SharednetworkOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetSharednetworkResponse) GetOptionsOk() (*[]SharednetworkOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetSharednetworkResponse) SetOptions(v []SharednetworkOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetSharednetworkResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetSharednetworkResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetSharednetworkResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetSharednetworkResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetSharednetworkResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetStaticHosts

`func (o *GetSharednetworkResponse) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *GetSharednetworkResponse) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *GetSharednetworkResponse) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *GetSharednetworkResponse) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetTotalHosts

`func (o *GetSharednetworkResponse) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *GetSharednetworkResponse) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *GetSharednetworkResponse) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *GetSharednetworkResponse) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *GetSharednetworkResponse) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *GetSharednetworkResponse) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *GetSharednetworkResponse) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *GetSharednetworkResponse) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseAuthority

`func (o *GetSharednetworkResponse) GetUseAuthority() bool`

GetUseAuthority returns the UseAuthority field if non-nil, zero value otherwise.

### GetUseAuthorityOk

`func (o *GetSharednetworkResponse) GetUseAuthorityOk() (*bool, bool)`

GetUseAuthorityOk returns a tuple with the UseAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthority

`func (o *GetSharednetworkResponse) SetUseAuthority(v bool)`

SetUseAuthority sets UseAuthority field to given value.

### HasUseAuthority

`func (o *GetSharednetworkResponse) HasUseAuthority() bool`

HasUseAuthority returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetSharednetworkResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetSharednetworkResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetSharednetworkResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetSharednetworkResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetSharednetworkResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetSharednetworkResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetSharednetworkResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetSharednetworkResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *GetSharednetworkResponse) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *GetSharednetworkResponse) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *GetSharednetworkResponse) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *GetSharednetworkResponse) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDdnsTtl

`func (o *GetSharednetworkResponse) GetUseDdnsTtl() bool`

GetUseDdnsTtl returns the UseDdnsTtl field if non-nil, zero value otherwise.

### GetUseDdnsTtlOk

`func (o *GetSharednetworkResponse) GetUseDdnsTtlOk() (*bool, bool)`

GetUseDdnsTtlOk returns a tuple with the UseDdnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsTtl

`func (o *GetSharednetworkResponse) SetUseDdnsTtl(v bool)`

SetUseDdnsTtl sets UseDdnsTtl field to given value.

### HasUseDdnsTtl

`func (o *GetSharednetworkResponse) HasUseDdnsTtl() bool`

HasUseDdnsTtl returns a boolean if a field has been set.

### GetUseDdnsUpdateFixedAddresses

`func (o *GetSharednetworkResponse) GetUseDdnsUpdateFixedAddresses() bool`

GetUseDdnsUpdateFixedAddresses returns the UseDdnsUpdateFixedAddresses field if non-nil, zero value otherwise.

### GetUseDdnsUpdateFixedAddressesOk

`func (o *GetSharednetworkResponse) GetUseDdnsUpdateFixedAddressesOk() (*bool, bool)`

GetUseDdnsUpdateFixedAddressesOk returns a tuple with the UseDdnsUpdateFixedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUpdateFixedAddresses

`func (o *GetSharednetworkResponse) SetUseDdnsUpdateFixedAddresses(v bool)`

SetUseDdnsUpdateFixedAddresses sets UseDdnsUpdateFixedAddresses field to given value.

### HasUseDdnsUpdateFixedAddresses

`func (o *GetSharednetworkResponse) HasUseDdnsUpdateFixedAddresses() bool`

HasUseDdnsUpdateFixedAddresses returns a boolean if a field has been set.

### GetUseDdnsUseOption81

`func (o *GetSharednetworkResponse) GetUseDdnsUseOption81() bool`

GetUseDdnsUseOption81 returns the UseDdnsUseOption81 field if non-nil, zero value otherwise.

### GetUseDdnsUseOption81Ok

`func (o *GetSharednetworkResponse) GetUseDdnsUseOption81Ok() (*bool, bool)`

GetUseDdnsUseOption81Ok returns a tuple with the UseDdnsUseOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsUseOption81

`func (o *GetSharednetworkResponse) SetUseDdnsUseOption81(v bool)`

SetUseDdnsUseOption81 sets UseDdnsUseOption81 field to given value.

### HasUseDdnsUseOption81

`func (o *GetSharednetworkResponse) HasUseDdnsUseOption81() bool`

HasUseDdnsUseOption81 returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetSharednetworkResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetSharednetworkResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetSharednetworkResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetSharednetworkResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetSharednetworkResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetSharednetworkResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetSharednetworkResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetSharednetworkResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseIgnoreClientIdentifier

`func (o *GetSharednetworkResponse) GetUseIgnoreClientIdentifier() bool`

GetUseIgnoreClientIdentifier returns the UseIgnoreClientIdentifier field if non-nil, zero value otherwise.

### GetUseIgnoreClientIdentifierOk

`func (o *GetSharednetworkResponse) GetUseIgnoreClientIdentifierOk() (*bool, bool)`

GetUseIgnoreClientIdentifierOk returns a tuple with the UseIgnoreClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreClientIdentifier

`func (o *GetSharednetworkResponse) SetUseIgnoreClientIdentifier(v bool)`

SetUseIgnoreClientIdentifier sets UseIgnoreClientIdentifier field to given value.

### HasUseIgnoreClientIdentifier

`func (o *GetSharednetworkResponse) HasUseIgnoreClientIdentifier() bool`

HasUseIgnoreClientIdentifier returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *GetSharednetworkResponse) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *GetSharednetworkResponse) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *GetSharednetworkResponse) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *GetSharednetworkResponse) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseIgnoreId

`func (o *GetSharednetworkResponse) GetUseIgnoreId() bool`

GetUseIgnoreId returns the UseIgnoreId field if non-nil, zero value otherwise.

### GetUseIgnoreIdOk

`func (o *GetSharednetworkResponse) GetUseIgnoreIdOk() (*bool, bool)`

GetUseIgnoreIdOk returns a tuple with the UseIgnoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreId

`func (o *GetSharednetworkResponse) SetUseIgnoreId(v bool)`

SetUseIgnoreId sets UseIgnoreId field to given value.

### HasUseIgnoreId

`func (o *GetSharednetworkResponse) HasUseIgnoreId() bool`

HasUseIgnoreId returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *GetSharednetworkResponse) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *GetSharednetworkResponse) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *GetSharednetworkResponse) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *GetSharednetworkResponse) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetSharednetworkResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetSharednetworkResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetSharednetworkResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetSharednetworkResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetSharednetworkResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetSharednetworkResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetSharednetworkResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetSharednetworkResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetSharednetworkResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetSharednetworkResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetSharednetworkResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetSharednetworkResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetSharednetworkResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetSharednetworkResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetSharednetworkResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetSharednetworkResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *GetSharednetworkResponse) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *GetSharednetworkResponse) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *GetSharednetworkResponse) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *GetSharednetworkResponse) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetResult

`func (o *GetSharednetworkResponse) GetResult() Sharednetwork`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSharednetworkResponse) GetResultOk() (*Sharednetwork, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSharednetworkResponse) SetResult(v Sharednetwork)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSharednetworkResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


