# Fixedaddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AgentCircuitId** | Pointer to **string** | The agent circuit ID for the fixed address. | [optional] 
**AgentRemoteId** | Pointer to **string** | The agent remote ID for the fixed address. | [optional] 
**AllowTelnet** | Pointer to **bool** | This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH. | [optional] 
**AlwaysUpdateDns** | Pointer to **bool** | This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP client requests. | [optional] 
**Bootfile** | Pointer to **string** | The bootfile name for the fixed address. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The bootserver address for the fixed address. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**CliCredentials** | Pointer to [**[]FixedaddressCliCredentials**](FixedaddressCliCredentials.md) | The CLI credentials for the fixed address. | [optional] 
**ClientIdentifierPrependZero** | Pointer to **bool** | This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address. | [optional] 
**CloudInfo** | Pointer to [**FixedaddressCloudInfo**](FixedaddressCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the fixed address; maximum 256 characters. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this fixed address. | [optional] 
**DdnsHostname** | Pointer to **string** | The DDNS host name for this fixed address. | [optional] 
**DenyBootp** | Pointer to **bool** | If set to true, BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**DeviceDescription** | Pointer to **string** | The description of the device. | [optional] 
**DeviceLocation** | Pointer to **string** | The location of the device. | [optional] 
**DeviceType** | Pointer to **string** | The type of the device. | [optional] 
**DeviceVendor** | Pointer to **string** | The vendor of the device. | [optional] 
**DhcpClientIdentifier** | Pointer to **string** | The DHCP client ID for the fixed address. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled. | [optional] 
**DisableDiscovery** | Pointer to **bool** | Determines if the discovery for this fixed address is disabled or not. False means that the discovery is enabled. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | The discovery status of this fixed address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**FixedaddressDiscoveredData**](FixedaddressDiscoveredData.md) |  | [optional] 
**EnableDdns** | Pointer to **bool** | The dynamic DNS updates flag of a DHCP Fixed Address object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the fixed address should be immediately enabled. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**Ipv4addr** | Pointer to [**FixedaddressIpv4addr**](FixedaddressIpv4addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**IsInvalidMac** | Pointer to **bool** | This flag reflects whether the MAC address for this fixed address is invalid. | [optional] [readonly] 
**LogicFilterRules** | Pointer to [**[]FixedaddressLogicFilterRules**](FixedaddressLogicFilterRules.md) | This field contains the logic filters to be applied on the this fixed address. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**Mac** | Pointer to **string** | The MAC address value for this fixed address. | [optional] 
**MatchClient** | Pointer to **string** | The match_client value for this fixed address. Valid values are: \&quot;MAC_ADDRESS\&quot;: The fixed IP address is leased to the matching MAC address. \&quot;CLIENT_ID\&quot;: The fixed IP address is leased to the matching DHCP client identifier. \&quot;RESERVED\&quot;: The fixed IP address is reserved for later use with a MAC address that only has zeros. \&quot;CIRCUIT_ID\&quot;: The fixed IP address is leased to the DHCP client with a matching circuit ID. Note that the \&quot;agent_circuit_id\&quot; field must be set in this case. \&quot;REMOTE_ID\&quot;: The fixed IP address is leased to the DHCP client with a matching remote ID. Note that the \&quot;agent_remote_id\&quot; field must be set in this case. | [optional] 
**MsAdUserData** | Pointer to [**FixedaddressMsAdUserData**](FixedaddressMsAdUserData.md) |  | [optional] 
**MsOptions** | Pointer to [**[]FixedaddressMsOptions**](FixedaddressMsOptions.md) | This field contains the Microsoft DHCP options for this fixed address. | [optional] 
**MsServer** | Pointer to [**FixedaddressMsServer**](FixedaddressMsServer.md) |  | [optional] 
**Name** | Pointer to **string** | This field contains the name of this fixed address. | [optional] 
**Network** | Pointer to **string** | The network to which this fixed address belongs, in IPv4 Address/CIDR format. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which this fixed address resides. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot. | [optional] 
**Options** | Pointer to [**[]FixedaddressOptions**](FixedaddressOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value for a DHCP Fixed Address object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**ReservedInterface** | Pointer to **string** | The ref to the reserved interface to which the device belongs. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. The restart_if_needed flag can trigger a restart on DHCP services only when it is enabled on CP member. | [optional] 
**Snmp3Credential** | Pointer to [**FixedaddressSnmp3Credential**](FixedaddressSnmp3Credential.md) |  | [optional] 
**SnmpCredential** | Pointer to [**FixedaddressSnmpCredential**](FixedaddressSnmpCredential.md) |  | [optional] 
**Template** | Pointer to **string** | If set on creation, the fixed address will be created according to the values specified in the named template. | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseCliCredentials** | Pointer to **bool** | If set to true, the CLI credential will override member-level settings. | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMsOptions** | Pointer to **bool** | Use flag for: ms_options | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseSnmp3Credential** | Pointer to **bool** | Determines if the SNMPv3 credential should be used for the fixed address. | [optional] 
**UseSnmpCredential** | Pointer to **bool** | If set to true, the SNMP credential will override member-level settings. | [optional] 

## Methods

### NewFixedaddress

`func NewFixedaddress() *Fixedaddress`

NewFixedaddress instantiates a new Fixedaddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedaddressWithDefaults

`func NewFixedaddressWithDefaults() *Fixedaddress`

NewFixedaddressWithDefaults instantiates a new Fixedaddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Fixedaddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Fixedaddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Fixedaddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Fixedaddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAgentCircuitId

`func (o *Fixedaddress) GetAgentCircuitId() string`

GetAgentCircuitId returns the AgentCircuitId field if non-nil, zero value otherwise.

### GetAgentCircuitIdOk

`func (o *Fixedaddress) GetAgentCircuitIdOk() (*string, bool)`

GetAgentCircuitIdOk returns a tuple with the AgentCircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCircuitId

`func (o *Fixedaddress) SetAgentCircuitId(v string)`

SetAgentCircuitId sets AgentCircuitId field to given value.

### HasAgentCircuitId

`func (o *Fixedaddress) HasAgentCircuitId() bool`

HasAgentCircuitId returns a boolean if a field has been set.

### GetAgentRemoteId

`func (o *Fixedaddress) GetAgentRemoteId() string`

GetAgentRemoteId returns the AgentRemoteId field if non-nil, zero value otherwise.

### GetAgentRemoteIdOk

`func (o *Fixedaddress) GetAgentRemoteIdOk() (*string, bool)`

GetAgentRemoteIdOk returns a tuple with the AgentRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRemoteId

`func (o *Fixedaddress) SetAgentRemoteId(v string)`

SetAgentRemoteId sets AgentRemoteId field to given value.

### HasAgentRemoteId

`func (o *Fixedaddress) HasAgentRemoteId() bool`

HasAgentRemoteId returns a boolean if a field has been set.

### GetAllowTelnet

`func (o *Fixedaddress) GetAllowTelnet() bool`

GetAllowTelnet returns the AllowTelnet field if non-nil, zero value otherwise.

### GetAllowTelnetOk

`func (o *Fixedaddress) GetAllowTelnetOk() (*bool, bool)`

GetAllowTelnetOk returns a tuple with the AllowTelnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTelnet

`func (o *Fixedaddress) SetAllowTelnet(v bool)`

SetAllowTelnet sets AllowTelnet field to given value.

### HasAllowTelnet

`func (o *Fixedaddress) HasAllowTelnet() bool`

HasAllowTelnet returns a boolean if a field has been set.

### GetAlwaysUpdateDns

`func (o *Fixedaddress) GetAlwaysUpdateDns() bool`

GetAlwaysUpdateDns returns the AlwaysUpdateDns field if non-nil, zero value otherwise.

### GetAlwaysUpdateDnsOk

`func (o *Fixedaddress) GetAlwaysUpdateDnsOk() (*bool, bool)`

GetAlwaysUpdateDnsOk returns a tuple with the AlwaysUpdateDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUpdateDns

`func (o *Fixedaddress) SetAlwaysUpdateDns(v bool)`

SetAlwaysUpdateDns sets AlwaysUpdateDns field to given value.

### HasAlwaysUpdateDns

`func (o *Fixedaddress) HasAlwaysUpdateDns() bool`

HasAlwaysUpdateDns returns a boolean if a field has been set.

### GetBootfile

`func (o *Fixedaddress) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *Fixedaddress) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *Fixedaddress) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *Fixedaddress) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *Fixedaddress) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *Fixedaddress) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *Fixedaddress) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *Fixedaddress) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCliCredentials

`func (o *Fixedaddress) GetCliCredentials() []FixedaddressCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *Fixedaddress) GetCliCredentialsOk() (*[]FixedaddressCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *Fixedaddress) SetCliCredentials(v []FixedaddressCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *Fixedaddress) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetClientIdentifierPrependZero

`func (o *Fixedaddress) GetClientIdentifierPrependZero() bool`

GetClientIdentifierPrependZero returns the ClientIdentifierPrependZero field if non-nil, zero value otherwise.

### GetClientIdentifierPrependZeroOk

`func (o *Fixedaddress) GetClientIdentifierPrependZeroOk() (*bool, bool)`

GetClientIdentifierPrependZeroOk returns a tuple with the ClientIdentifierPrependZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifierPrependZero

`func (o *Fixedaddress) SetClientIdentifierPrependZero(v bool)`

SetClientIdentifierPrependZero sets ClientIdentifierPrependZero field to given value.

### HasClientIdentifierPrependZero

`func (o *Fixedaddress) HasClientIdentifierPrependZero() bool`

HasClientIdentifierPrependZero returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Fixedaddress) GetCloudInfo() FixedaddressCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Fixedaddress) GetCloudInfoOk() (*FixedaddressCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Fixedaddress) SetCloudInfo(v FixedaddressCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Fixedaddress) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Fixedaddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Fixedaddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Fixedaddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Fixedaddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Fixedaddress) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Fixedaddress) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Fixedaddress) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Fixedaddress) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsHostname

`func (o *Fixedaddress) GetDdnsHostname() string`

GetDdnsHostname returns the DdnsHostname field if non-nil, zero value otherwise.

### GetDdnsHostnameOk

`func (o *Fixedaddress) GetDdnsHostnameOk() (*string, bool)`

GetDdnsHostnameOk returns a tuple with the DdnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostname

`func (o *Fixedaddress) SetDdnsHostname(v string)`

SetDdnsHostname sets DdnsHostname field to given value.

### HasDdnsHostname

`func (o *Fixedaddress) HasDdnsHostname() bool`

HasDdnsHostname returns a boolean if a field has been set.

### GetDenyBootp

`func (o *Fixedaddress) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *Fixedaddress) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *Fixedaddress) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *Fixedaddress) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetDeviceDescription

`func (o *Fixedaddress) GetDeviceDescription() string`

GetDeviceDescription returns the DeviceDescription field if non-nil, zero value otherwise.

### GetDeviceDescriptionOk

`func (o *Fixedaddress) GetDeviceDescriptionOk() (*string, bool)`

GetDeviceDescriptionOk returns a tuple with the DeviceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDescription

`func (o *Fixedaddress) SetDeviceDescription(v string)`

SetDeviceDescription sets DeviceDescription field to given value.

### HasDeviceDescription

`func (o *Fixedaddress) HasDeviceDescription() bool`

HasDeviceDescription returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *Fixedaddress) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *Fixedaddress) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *Fixedaddress) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *Fixedaddress) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceType

`func (o *Fixedaddress) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Fixedaddress) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Fixedaddress) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Fixedaddress) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *Fixedaddress) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *Fixedaddress) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *Fixedaddress) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *Fixedaddress) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDhcpClientIdentifier

`func (o *Fixedaddress) GetDhcpClientIdentifier() string`

GetDhcpClientIdentifier returns the DhcpClientIdentifier field if non-nil, zero value otherwise.

### GetDhcpClientIdentifierOk

`func (o *Fixedaddress) GetDhcpClientIdentifierOk() (*string, bool)`

GetDhcpClientIdentifierOk returns a tuple with the DhcpClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpClientIdentifier

`func (o *Fixedaddress) SetDhcpClientIdentifier(v string)`

SetDhcpClientIdentifier sets DhcpClientIdentifier field to given value.

### HasDhcpClientIdentifier

`func (o *Fixedaddress) HasDhcpClientIdentifier() bool`

HasDhcpClientIdentifier returns a boolean if a field has been set.

### GetDisable

`func (o *Fixedaddress) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Fixedaddress) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Fixedaddress) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Fixedaddress) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableDiscovery

`func (o *Fixedaddress) GetDisableDiscovery() bool`

GetDisableDiscovery returns the DisableDiscovery field if non-nil, zero value otherwise.

### GetDisableDiscoveryOk

`func (o *Fixedaddress) GetDisableDiscoveryOk() (*bool, bool)`

GetDisableDiscoveryOk returns a tuple with the DisableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDiscovery

`func (o *Fixedaddress) SetDisableDiscovery(v bool)`

SetDisableDiscovery sets DisableDiscovery field to given value.

### HasDisableDiscovery

`func (o *Fixedaddress) HasDisableDiscovery() bool`

HasDisableDiscovery returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Fixedaddress) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Fixedaddress) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Fixedaddress) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Fixedaddress) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *Fixedaddress) GetDiscoveredData() FixedaddressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *Fixedaddress) GetDiscoveredDataOk() (*FixedaddressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *Fixedaddress) SetDiscoveredData(v FixedaddressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *Fixedaddress) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Fixedaddress) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Fixedaddress) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Fixedaddress) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Fixedaddress) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *Fixedaddress) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *Fixedaddress) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *Fixedaddress) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *Fixedaddress) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *Fixedaddress) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *Fixedaddress) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *Fixedaddress) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *Fixedaddress) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetExtattrs

`func (o *Fixedaddress) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Fixedaddress) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Fixedaddress) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Fixedaddress) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *Fixedaddress) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *Fixedaddress) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *Fixedaddress) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *Fixedaddress) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetIpv4addr

`func (o *Fixedaddress) GetIpv4addr() FixedaddressIpv4addr`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *Fixedaddress) GetIpv4addrOk() (*FixedaddressIpv4addr, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *Fixedaddress) SetIpv4addr(v FixedaddressIpv4addr)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *Fixedaddress) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *Fixedaddress) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *Fixedaddress) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *Fixedaddress) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *Fixedaddress) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetIsInvalidMac

`func (o *Fixedaddress) GetIsInvalidMac() bool`

GetIsInvalidMac returns the IsInvalidMac field if non-nil, zero value otherwise.

### GetIsInvalidMacOk

`func (o *Fixedaddress) GetIsInvalidMacOk() (*bool, bool)`

GetIsInvalidMacOk returns a tuple with the IsInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalidMac

`func (o *Fixedaddress) SetIsInvalidMac(v bool)`

SetIsInvalidMac sets IsInvalidMac field to given value.

### HasIsInvalidMac

`func (o *Fixedaddress) HasIsInvalidMac() bool`

HasIsInvalidMac returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Fixedaddress) GetLogicFilterRules() []FixedaddressLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Fixedaddress) GetLogicFilterRulesOk() (*[]FixedaddressLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Fixedaddress) SetLogicFilterRules(v []FixedaddressLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Fixedaddress) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMac

`func (o *Fixedaddress) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Fixedaddress) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Fixedaddress) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Fixedaddress) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMatchClient

`func (o *Fixedaddress) GetMatchClient() string`

GetMatchClient returns the MatchClient field if non-nil, zero value otherwise.

### GetMatchClientOk

`func (o *Fixedaddress) GetMatchClientOk() (*string, bool)`

GetMatchClientOk returns a tuple with the MatchClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClient

`func (o *Fixedaddress) SetMatchClient(v string)`

SetMatchClient sets MatchClient field to given value.

### HasMatchClient

`func (o *Fixedaddress) HasMatchClient() bool`

HasMatchClient returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Fixedaddress) GetMsAdUserData() FixedaddressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Fixedaddress) GetMsAdUserDataOk() (*FixedaddressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Fixedaddress) SetMsAdUserData(v FixedaddressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Fixedaddress) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetMsOptions

`func (o *Fixedaddress) GetMsOptions() []FixedaddressMsOptions`

GetMsOptions returns the MsOptions field if non-nil, zero value otherwise.

### GetMsOptionsOk

`func (o *Fixedaddress) GetMsOptionsOk() (*[]FixedaddressMsOptions, bool)`

GetMsOptionsOk returns a tuple with the MsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsOptions

`func (o *Fixedaddress) SetMsOptions(v []FixedaddressMsOptions)`

SetMsOptions sets MsOptions field to given value.

### HasMsOptions

`func (o *Fixedaddress) HasMsOptions() bool`

HasMsOptions returns a boolean if a field has been set.

### GetMsServer

`func (o *Fixedaddress) GetMsServer() FixedaddressMsServer`

GetMsServer returns the MsServer field if non-nil, zero value otherwise.

### GetMsServerOk

`func (o *Fixedaddress) GetMsServerOk() (*FixedaddressMsServer, bool)`

GetMsServerOk returns a tuple with the MsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServer

`func (o *Fixedaddress) SetMsServer(v FixedaddressMsServer)`

SetMsServer sets MsServer field to given value.

### HasMsServer

`func (o *Fixedaddress) HasMsServer() bool`

HasMsServer returns a boolean if a field has been set.

### GetName

`func (o *Fixedaddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Fixedaddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Fixedaddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Fixedaddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Fixedaddress) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Fixedaddress) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Fixedaddress) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Fixedaddress) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Fixedaddress) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Fixedaddress) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Fixedaddress) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Fixedaddress) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNextserver

`func (o *Fixedaddress) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *Fixedaddress) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *Fixedaddress) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *Fixedaddress) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetOptions

`func (o *Fixedaddress) GetOptions() []FixedaddressOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Fixedaddress) GetOptionsOk() (*[]FixedaddressOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Fixedaddress) SetOptions(v []FixedaddressOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Fixedaddress) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *Fixedaddress) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *Fixedaddress) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *Fixedaddress) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *Fixedaddress) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetReservedInterface

`func (o *Fixedaddress) GetReservedInterface() string`

GetReservedInterface returns the ReservedInterface field if non-nil, zero value otherwise.

### GetReservedInterfaceOk

`func (o *Fixedaddress) GetReservedInterfaceOk() (*string, bool)`

GetReservedInterfaceOk returns a tuple with the ReservedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedInterface

`func (o *Fixedaddress) SetReservedInterface(v string)`

SetReservedInterface sets ReservedInterface field to given value.

### HasReservedInterface

`func (o *Fixedaddress) HasReservedInterface() bool`

HasReservedInterface returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *Fixedaddress) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *Fixedaddress) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *Fixedaddress) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *Fixedaddress) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetSnmp3Credential

`func (o *Fixedaddress) GetSnmp3Credential() FixedaddressSnmp3Credential`

GetSnmp3Credential returns the Snmp3Credential field if non-nil, zero value otherwise.

### GetSnmp3CredentialOk

`func (o *Fixedaddress) GetSnmp3CredentialOk() (*FixedaddressSnmp3Credential, bool)`

GetSnmp3CredentialOk returns a tuple with the Snmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp3Credential

`func (o *Fixedaddress) SetSnmp3Credential(v FixedaddressSnmp3Credential)`

SetSnmp3Credential sets Snmp3Credential field to given value.

### HasSnmp3Credential

`func (o *Fixedaddress) HasSnmp3Credential() bool`

HasSnmp3Credential returns a boolean if a field has been set.

### GetSnmpCredential

`func (o *Fixedaddress) GetSnmpCredential() FixedaddressSnmpCredential`

GetSnmpCredential returns the SnmpCredential field if non-nil, zero value otherwise.

### GetSnmpCredentialOk

`func (o *Fixedaddress) GetSnmpCredentialOk() (*FixedaddressSnmpCredential, bool)`

GetSnmpCredentialOk returns a tuple with the SnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredential

`func (o *Fixedaddress) SetSnmpCredential(v FixedaddressSnmpCredential)`

SetSnmpCredential sets SnmpCredential field to given value.

### HasSnmpCredential

`func (o *Fixedaddress) HasSnmpCredential() bool`

HasSnmpCredential returns a boolean if a field has been set.

### GetTemplate

`func (o *Fixedaddress) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Fixedaddress) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Fixedaddress) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Fixedaddress) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUseBootfile

`func (o *Fixedaddress) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *Fixedaddress) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *Fixedaddress) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *Fixedaddress) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *Fixedaddress) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *Fixedaddress) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *Fixedaddress) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *Fixedaddress) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseCliCredentials

`func (o *Fixedaddress) GetUseCliCredentials() bool`

GetUseCliCredentials returns the UseCliCredentials field if non-nil, zero value otherwise.

### GetUseCliCredentialsOk

`func (o *Fixedaddress) GetUseCliCredentialsOk() (*bool, bool)`

GetUseCliCredentialsOk returns a tuple with the UseCliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCliCredentials

`func (o *Fixedaddress) SetUseCliCredentials(v bool)`

SetUseCliCredentials sets UseCliCredentials field to given value.

### HasUseCliCredentials

`func (o *Fixedaddress) HasUseCliCredentials() bool`

HasUseCliCredentials returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Fixedaddress) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Fixedaddress) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Fixedaddress) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Fixedaddress) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *Fixedaddress) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *Fixedaddress) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *Fixedaddress) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *Fixedaddress) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Fixedaddress) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Fixedaddress) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Fixedaddress) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Fixedaddress) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *Fixedaddress) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *Fixedaddress) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *Fixedaddress) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *Fixedaddress) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Fixedaddress) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Fixedaddress) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Fixedaddress) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Fixedaddress) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMsOptions

`func (o *Fixedaddress) GetUseMsOptions() bool`

GetUseMsOptions returns the UseMsOptions field if non-nil, zero value otherwise.

### GetUseMsOptionsOk

`func (o *Fixedaddress) GetUseMsOptionsOk() (*bool, bool)`

GetUseMsOptionsOk returns a tuple with the UseMsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsOptions

`func (o *Fixedaddress) SetUseMsOptions(v bool)`

SetUseMsOptions sets UseMsOptions field to given value.

### HasUseMsOptions

`func (o *Fixedaddress) HasUseMsOptions() bool`

HasUseMsOptions returns a boolean if a field has been set.

### GetUseNextserver

`func (o *Fixedaddress) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *Fixedaddress) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *Fixedaddress) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *Fixedaddress) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *Fixedaddress) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Fixedaddress) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Fixedaddress) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Fixedaddress) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *Fixedaddress) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *Fixedaddress) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *Fixedaddress) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *Fixedaddress) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseSnmp3Credential

`func (o *Fixedaddress) GetUseSnmp3Credential() bool`

GetUseSnmp3Credential returns the UseSnmp3Credential field if non-nil, zero value otherwise.

### GetUseSnmp3CredentialOk

`func (o *Fixedaddress) GetUseSnmp3CredentialOk() (*bool, bool)`

GetUseSnmp3CredentialOk returns a tuple with the UseSnmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmp3Credential

`func (o *Fixedaddress) SetUseSnmp3Credential(v bool)`

SetUseSnmp3Credential sets UseSnmp3Credential field to given value.

### HasUseSnmp3Credential

`func (o *Fixedaddress) HasUseSnmp3Credential() bool`

HasUseSnmp3Credential returns a boolean if a field has been set.

### GetUseSnmpCredential

`func (o *Fixedaddress) GetUseSnmpCredential() bool`

GetUseSnmpCredential returns the UseSnmpCredential field if non-nil, zero value otherwise.

### GetUseSnmpCredentialOk

`func (o *Fixedaddress) GetUseSnmpCredentialOk() (*bool, bool)`

GetUseSnmpCredentialOk returns a tuple with the UseSnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpCredential

`func (o *Fixedaddress) SetUseSnmpCredential(v bool)`

SetUseSnmpCredential sets UseSnmpCredential field to given value.

### HasUseSnmpCredential

`func (o *Fixedaddress) HasUseSnmpCredential() bool`

HasUseSnmpCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


