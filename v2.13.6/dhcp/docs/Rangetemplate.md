# Rangetemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Bootfile** | Pointer to **string** | The bootfile name for the range. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The bootserver address for the range. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**CloudApiCompatible** | Pointer to **bool** | This flag controls whether this template can be used to create network objects in a cloud-computing deployment. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of a range template object. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this range. | [optional] 
**DdnsGenerateHostname** | Pointer to **bool** | If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname. | [optional] 
**DelegatedMember** | Pointer to [**RangetemplateDelegatedMember**](RangetemplateDelegatedMember.md) |  | [optional] 
**DenyAllClients** | Pointer to **bool** | If True, send NAK forcing the client to take the new address. | [optional] 
**DenyBootp** | Pointer to **bool** | Determines if BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**EmailList** | Pointer to **[]string** | The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages. | [optional] 
**EnableDdns** | Pointer to **bool** | Determines if the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnableDhcpThresholds** | Pointer to **bool** | Determines if DHCP thresholds are enabled for the range. | [optional] 
**EnableEmailWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are sent through email. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determines if DHCP threshold warnings are sent through SNMP. | [optional] 
**Exclude** | Pointer to [**[]RangetemplateExclude**](RangetemplateExclude.md) | These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally, information about this exclusion range. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FailoverAssociation** | Pointer to **string** | The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {rangetemplate:rangetemplate} must be set to &#39;FAILOVER&#39; or &#39;FAILOVER_MS&#39; if you want the failover association specified here to serve the range. | [optional] 
**FingerprintFilterRules** | Pointer to [**[]RangetemplateFingerprintFilterRules**](RangetemplateFingerprintFilterRules.md) | This field contains the fingerprint filters for this DHCP range. The appliance uses matching rules in these filters to select the address range from which it assigns a lease. | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage of DHCP range usage threshold above which range usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage of DHCP range usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**KnownClients** | Pointer to **string** | Permission for known clients. If set to &#39;Deny&#39; known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries. | [optional] 
**LeaseScavengeTime** | Pointer to **int64** | An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day). | [optional] 
**LogicFilterRules** | Pointer to [**[]RangetemplateLogicFilterRules**](RangetemplateLogicFilterRules.md) | This field contains the logic filters to be applied on this range. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**LowWaterMark** | Pointer to **int64** | The percentage of DHCP range usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. | [optional] 
**LowWaterMarkReset** | Pointer to **int64** | The percentage of DHCP range usage threshold below which range usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value. | [optional] 
**MacFilterRules** | Pointer to [**[]RangetemplateMacFilterRules**](RangetemplateMacFilterRules.md) | This field contains the MAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**Member** | Pointer to [**RangetemplateMember**](RangetemplateMember.md) |  | [optional] 
**MsOptions** | Pointer to [**[]RangetemplateMsOptions**](RangetemplateMsOptions.md) | The Microsoft DHCP options for this range. | [optional] 
**MsServer** | Pointer to [**RangetemplateMsServer**](RangetemplateMsServer.md) |  | [optional] 
**NacFilterRules** | Pointer to [**[]RangetemplateNacFilterRules**](RangetemplateNacFilterRules.md) | This field contains the NAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**Name** | Pointer to **string** | The name of a range template object. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot. | [optional] 
**NumberOfAddresses** | Pointer to **int64** | The number of addresses for this range. | [optional] 
**Offset** | Pointer to **int64** | The start address offset for this range. | [optional] 
**OptionFilterRules** | Pointer to [**[]RangetemplateOptionFilterRules**](RangetemplateOptionFilterRules.md) | This field contains the Option filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**Options** | Pointer to [**[]RangetemplateOptions**](RangetemplateOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value for a range object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**RecycleLeases** | Pointer to **bool** | If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted. | [optional] 
**RelayAgentFilterRules** | Pointer to [**[]RangetemplateRelayAgentFilterRules**](RangetemplateRelayAgentFilterRules.md) | This field contains the Relay Agent filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**ServerAssociationType** | Pointer to **string** | The type of server that is going to serve the range. | [optional] 
**UnknownClients** | Pointer to **string** | Permission for unknown clients. If set to &#39;Deny&#39; unknown clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries. | [optional] 
**UpdateDnsOnLeaseRenewal** | Pointer to **bool** | This field controls whether the DHCP server updates DNS when a DHCP lease is renewed. | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDdnsGenerateHostname** | Pointer to **bool** | Use flag for: ddns_generate_hostname | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseEmailList** | Pointer to **bool** | Use flag for: email_list | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseEnableDhcpThresholds** | Pointer to **bool** | Use flag for: enable_dhcp_thresholds | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseKnownClients** | Pointer to **bool** | Use flag for: known_clients | [optional] 
**UseLeaseScavengeTime** | Pointer to **bool** | Use flag for: lease_scavenge_time | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseMsOptions** | Pointer to **bool** | Use flag for: ms_options | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**UseUnknownClients** | Pointer to **bool** | Use flag for: unknown_clients | [optional] 
**UseUpdateDnsOnLeaseRenewal** | Pointer to **bool** | Use flag for: update_dns_on_lease_renewal | [optional] 

## Methods

### NewRangetemplate

`func NewRangetemplate() *Rangetemplate`

NewRangetemplate instantiates a new Rangetemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangetemplateWithDefaults

`func NewRangetemplateWithDefaults() *Rangetemplate`

NewRangetemplateWithDefaults instantiates a new Rangetemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Rangetemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Rangetemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Rangetemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Rangetemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBootfile

`func (o *Rangetemplate) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *Rangetemplate) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *Rangetemplate) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *Rangetemplate) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *Rangetemplate) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *Rangetemplate) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *Rangetemplate) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *Rangetemplate) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetCloudApiCompatible

`func (o *Rangetemplate) GetCloudApiCompatible() bool`

GetCloudApiCompatible returns the CloudApiCompatible field if non-nil, zero value otherwise.

### GetCloudApiCompatibleOk

`func (o *Rangetemplate) GetCloudApiCompatibleOk() (*bool, bool)`

GetCloudApiCompatibleOk returns a tuple with the CloudApiCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudApiCompatible

`func (o *Rangetemplate) SetCloudApiCompatible(v bool)`

SetCloudApiCompatible sets CloudApiCompatible field to given value.

### HasCloudApiCompatible

`func (o *Rangetemplate) HasCloudApiCompatible() bool`

HasCloudApiCompatible returns a boolean if a field has been set.

### GetComment

`func (o *Rangetemplate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Rangetemplate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Rangetemplate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Rangetemplate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *Rangetemplate) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *Rangetemplate) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *Rangetemplate) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *Rangetemplate) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsGenerateHostname

`func (o *Rangetemplate) GetDdnsGenerateHostname() bool`

GetDdnsGenerateHostname returns the DdnsGenerateHostname field if non-nil, zero value otherwise.

### GetDdnsGenerateHostnameOk

`func (o *Rangetemplate) GetDdnsGenerateHostnameOk() (*bool, bool)`

GetDdnsGenerateHostnameOk returns a tuple with the DdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateHostname

`func (o *Rangetemplate) SetDdnsGenerateHostname(v bool)`

SetDdnsGenerateHostname sets DdnsGenerateHostname field to given value.

### HasDdnsGenerateHostname

`func (o *Rangetemplate) HasDdnsGenerateHostname() bool`

HasDdnsGenerateHostname returns a boolean if a field has been set.

### GetDelegatedMember

`func (o *Rangetemplate) GetDelegatedMember() RangetemplateDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *Rangetemplate) GetDelegatedMemberOk() (*RangetemplateDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *Rangetemplate) SetDelegatedMember(v RangetemplateDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *Rangetemplate) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetDenyAllClients

`func (o *Rangetemplate) GetDenyAllClients() bool`

GetDenyAllClients returns the DenyAllClients field if non-nil, zero value otherwise.

### GetDenyAllClientsOk

`func (o *Rangetemplate) GetDenyAllClientsOk() (*bool, bool)`

GetDenyAllClientsOk returns a tuple with the DenyAllClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyAllClients

`func (o *Rangetemplate) SetDenyAllClients(v bool)`

SetDenyAllClients sets DenyAllClients field to given value.

### HasDenyAllClients

`func (o *Rangetemplate) HasDenyAllClients() bool`

HasDenyAllClients returns a boolean if a field has been set.

### GetDenyBootp

`func (o *Rangetemplate) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *Rangetemplate) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *Rangetemplate) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *Rangetemplate) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetEmailList

`func (o *Rangetemplate) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *Rangetemplate) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *Rangetemplate) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *Rangetemplate) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnableDdns

`func (o *Rangetemplate) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *Rangetemplate) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *Rangetemplate) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *Rangetemplate) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnableDhcpThresholds

`func (o *Rangetemplate) GetEnableDhcpThresholds() bool`

GetEnableDhcpThresholds returns the EnableDhcpThresholds field if non-nil, zero value otherwise.

### GetEnableDhcpThresholdsOk

`func (o *Rangetemplate) GetEnableDhcpThresholdsOk() (*bool, bool)`

GetEnableDhcpThresholdsOk returns a tuple with the EnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcpThresholds

`func (o *Rangetemplate) SetEnableDhcpThresholds(v bool)`

SetEnableDhcpThresholds sets EnableDhcpThresholds field to given value.

### HasEnableDhcpThresholds

`func (o *Rangetemplate) HasEnableDhcpThresholds() bool`

HasEnableDhcpThresholds returns a boolean if a field has been set.

### GetEnableEmailWarnings

`func (o *Rangetemplate) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *Rangetemplate) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *Rangetemplate) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *Rangetemplate) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *Rangetemplate) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *Rangetemplate) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *Rangetemplate) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *Rangetemplate) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *Rangetemplate) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *Rangetemplate) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *Rangetemplate) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *Rangetemplate) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.

### GetExclude

`func (o *Rangetemplate) GetExclude() []RangetemplateExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Rangetemplate) GetExcludeOk() (*[]RangetemplateExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Rangetemplate) SetExclude(v []RangetemplateExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *Rangetemplate) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Rangetemplate) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Rangetemplate) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Rangetemplate) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Rangetemplate) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFailoverAssociation

`func (o *Rangetemplate) GetFailoverAssociation() string`

GetFailoverAssociation returns the FailoverAssociation field if non-nil, zero value otherwise.

### GetFailoverAssociationOk

`func (o *Rangetemplate) GetFailoverAssociationOk() (*string, bool)`

GetFailoverAssociationOk returns a tuple with the FailoverAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAssociation

`func (o *Rangetemplate) SetFailoverAssociation(v string)`

SetFailoverAssociation sets FailoverAssociation field to given value.

### HasFailoverAssociation

`func (o *Rangetemplate) HasFailoverAssociation() bool`

HasFailoverAssociation returns a boolean if a field has been set.

### GetFingerprintFilterRules

`func (o *Rangetemplate) GetFingerprintFilterRules() []RangetemplateFingerprintFilterRules`

GetFingerprintFilterRules returns the FingerprintFilterRules field if non-nil, zero value otherwise.

### GetFingerprintFilterRulesOk

`func (o *Rangetemplate) GetFingerprintFilterRulesOk() (*[]RangetemplateFingerprintFilterRules, bool)`

GetFingerprintFilterRulesOk returns a tuple with the FingerprintFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintFilterRules

`func (o *Rangetemplate) SetFingerprintFilterRules(v []RangetemplateFingerprintFilterRules)`

SetFingerprintFilterRules sets FingerprintFilterRules field to given value.

### HasFingerprintFilterRules

`func (o *Rangetemplate) HasFingerprintFilterRules() bool`

HasFingerprintFilterRules returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *Rangetemplate) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *Rangetemplate) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *Rangetemplate) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *Rangetemplate) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *Rangetemplate) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *Rangetemplate) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *Rangetemplate) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *Rangetemplate) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *Rangetemplate) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *Rangetemplate) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *Rangetemplate) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *Rangetemplate) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetKnownClients

`func (o *Rangetemplate) GetKnownClients() string`

GetKnownClients returns the KnownClients field if non-nil, zero value otherwise.

### GetKnownClientsOk

`func (o *Rangetemplate) GetKnownClientsOk() (*string, bool)`

GetKnownClientsOk returns a tuple with the KnownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClients

`func (o *Rangetemplate) SetKnownClients(v string)`

SetKnownClients sets KnownClients field to given value.

### HasKnownClients

`func (o *Rangetemplate) HasKnownClients() bool`

HasKnownClients returns a boolean if a field has been set.

### GetLeaseScavengeTime

`func (o *Rangetemplate) GetLeaseScavengeTime() int64`

GetLeaseScavengeTime returns the LeaseScavengeTime field if non-nil, zero value otherwise.

### GetLeaseScavengeTimeOk

`func (o *Rangetemplate) GetLeaseScavengeTimeOk() (*int64, bool)`

GetLeaseScavengeTimeOk returns a tuple with the LeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseScavengeTime

`func (o *Rangetemplate) SetLeaseScavengeTime(v int64)`

SetLeaseScavengeTime sets LeaseScavengeTime field to given value.

### HasLeaseScavengeTime

`func (o *Rangetemplate) HasLeaseScavengeTime() bool`

HasLeaseScavengeTime returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Rangetemplate) GetLogicFilterRules() []RangetemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Rangetemplate) GetLogicFilterRulesOk() (*[]RangetemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Rangetemplate) SetLogicFilterRules(v []RangetemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Rangetemplate) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *Rangetemplate) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *Rangetemplate) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *Rangetemplate) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *Rangetemplate) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *Rangetemplate) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *Rangetemplate) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *Rangetemplate) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *Rangetemplate) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetMacFilterRules

`func (o *Rangetemplate) GetMacFilterRules() []RangetemplateMacFilterRules`

GetMacFilterRules returns the MacFilterRules field if non-nil, zero value otherwise.

### GetMacFilterRulesOk

`func (o *Rangetemplate) GetMacFilterRulesOk() (*[]RangetemplateMacFilterRules, bool)`

GetMacFilterRulesOk returns a tuple with the MacFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilterRules

`func (o *Rangetemplate) SetMacFilterRules(v []RangetemplateMacFilterRules)`

SetMacFilterRules sets MacFilterRules field to given value.

### HasMacFilterRules

`func (o *Rangetemplate) HasMacFilterRules() bool`

HasMacFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *Rangetemplate) GetMember() RangetemplateMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Rangetemplate) GetMemberOk() (*RangetemplateMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Rangetemplate) SetMember(v RangetemplateMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *Rangetemplate) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMsOptions

`func (o *Rangetemplate) GetMsOptions() []RangetemplateMsOptions`

GetMsOptions returns the MsOptions field if non-nil, zero value otherwise.

### GetMsOptionsOk

`func (o *Rangetemplate) GetMsOptionsOk() (*[]RangetemplateMsOptions, bool)`

GetMsOptionsOk returns a tuple with the MsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsOptions

`func (o *Rangetemplate) SetMsOptions(v []RangetemplateMsOptions)`

SetMsOptions sets MsOptions field to given value.

### HasMsOptions

`func (o *Rangetemplate) HasMsOptions() bool`

HasMsOptions returns a boolean if a field has been set.

### GetMsServer

`func (o *Rangetemplate) GetMsServer() RangetemplateMsServer`

GetMsServer returns the MsServer field if non-nil, zero value otherwise.

### GetMsServerOk

`func (o *Rangetemplate) GetMsServerOk() (*RangetemplateMsServer, bool)`

GetMsServerOk returns a tuple with the MsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServer

`func (o *Rangetemplate) SetMsServer(v RangetemplateMsServer)`

SetMsServer sets MsServer field to given value.

### HasMsServer

`func (o *Rangetemplate) HasMsServer() bool`

HasMsServer returns a boolean if a field has been set.

### GetNacFilterRules

`func (o *Rangetemplate) GetNacFilterRules() []RangetemplateNacFilterRules`

GetNacFilterRules returns the NacFilterRules field if non-nil, zero value otherwise.

### GetNacFilterRulesOk

`func (o *Rangetemplate) GetNacFilterRulesOk() (*[]RangetemplateNacFilterRules, bool)`

GetNacFilterRulesOk returns a tuple with the NacFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacFilterRules

`func (o *Rangetemplate) SetNacFilterRules(v []RangetemplateNacFilterRules)`

SetNacFilterRules sets NacFilterRules field to given value.

### HasNacFilterRules

`func (o *Rangetemplate) HasNacFilterRules() bool`

HasNacFilterRules returns a boolean if a field has been set.

### GetName

`func (o *Rangetemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rangetemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rangetemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Rangetemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextserver

`func (o *Rangetemplate) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *Rangetemplate) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *Rangetemplate) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *Rangetemplate) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetNumberOfAddresses

`func (o *Rangetemplate) GetNumberOfAddresses() int64`

GetNumberOfAddresses returns the NumberOfAddresses field if non-nil, zero value otherwise.

### GetNumberOfAddressesOk

`func (o *Rangetemplate) GetNumberOfAddressesOk() (*int64, bool)`

GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddresses

`func (o *Rangetemplate) SetNumberOfAddresses(v int64)`

SetNumberOfAddresses sets NumberOfAddresses field to given value.

### HasNumberOfAddresses

`func (o *Rangetemplate) HasNumberOfAddresses() bool`

HasNumberOfAddresses returns a boolean if a field has been set.

### GetOffset

`func (o *Rangetemplate) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Rangetemplate) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Rangetemplate) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Rangetemplate) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *Rangetemplate) GetOptionFilterRules() []RangetemplateOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *Rangetemplate) GetOptionFilterRulesOk() (*[]RangetemplateOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *Rangetemplate) SetOptionFilterRules(v []RangetemplateOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *Rangetemplate) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetOptions

`func (o *Rangetemplate) GetOptions() []RangetemplateOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Rangetemplate) GetOptionsOk() (*[]RangetemplateOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Rangetemplate) SetOptions(v []RangetemplateOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Rangetemplate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *Rangetemplate) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *Rangetemplate) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *Rangetemplate) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *Rangetemplate) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Rangetemplate) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Rangetemplate) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Rangetemplate) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Rangetemplate) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetRelayAgentFilterRules

`func (o *Rangetemplate) GetRelayAgentFilterRules() []RangetemplateRelayAgentFilterRules`

GetRelayAgentFilterRules returns the RelayAgentFilterRules field if non-nil, zero value otherwise.

### GetRelayAgentFilterRulesOk

`func (o *Rangetemplate) GetRelayAgentFilterRulesOk() (*[]RangetemplateRelayAgentFilterRules, bool)`

GetRelayAgentFilterRulesOk returns a tuple with the RelayAgentFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAgentFilterRules

`func (o *Rangetemplate) SetRelayAgentFilterRules(v []RangetemplateRelayAgentFilterRules)`

SetRelayAgentFilterRules sets RelayAgentFilterRules field to given value.

### HasRelayAgentFilterRules

`func (o *Rangetemplate) HasRelayAgentFilterRules() bool`

HasRelayAgentFilterRules returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *Rangetemplate) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *Rangetemplate) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *Rangetemplate) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *Rangetemplate) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetUnknownClients

`func (o *Rangetemplate) GetUnknownClients() string`

GetUnknownClients returns the UnknownClients field if non-nil, zero value otherwise.

### GetUnknownClientsOk

`func (o *Rangetemplate) GetUnknownClientsOk() (*string, bool)`

GetUnknownClientsOk returns a tuple with the UnknownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownClients

`func (o *Rangetemplate) SetUnknownClients(v string)`

SetUnknownClients sets UnknownClients field to given value.

### HasUnknownClients

`func (o *Rangetemplate) HasUnknownClients() bool`

HasUnknownClients returns a boolean if a field has been set.

### GetUpdateDnsOnLeaseRenewal

`func (o *Rangetemplate) GetUpdateDnsOnLeaseRenewal() bool`

GetUpdateDnsOnLeaseRenewal returns the UpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUpdateDnsOnLeaseRenewalOk

`func (o *Rangetemplate) GetUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUpdateDnsOnLeaseRenewalOk returns a tuple with the UpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsOnLeaseRenewal

`func (o *Rangetemplate) SetUpdateDnsOnLeaseRenewal(v bool)`

SetUpdateDnsOnLeaseRenewal sets UpdateDnsOnLeaseRenewal field to given value.

### HasUpdateDnsOnLeaseRenewal

`func (o *Rangetemplate) HasUpdateDnsOnLeaseRenewal() bool`

HasUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.

### GetUseBootfile

`func (o *Rangetemplate) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *Rangetemplate) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *Rangetemplate) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *Rangetemplate) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *Rangetemplate) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *Rangetemplate) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *Rangetemplate) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *Rangetemplate) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *Rangetemplate) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *Rangetemplate) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *Rangetemplate) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *Rangetemplate) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDdnsGenerateHostname

`func (o *Rangetemplate) GetUseDdnsGenerateHostname() bool`

GetUseDdnsGenerateHostname returns the UseDdnsGenerateHostname field if non-nil, zero value otherwise.

### GetUseDdnsGenerateHostnameOk

`func (o *Rangetemplate) GetUseDdnsGenerateHostnameOk() (*bool, bool)`

GetUseDdnsGenerateHostnameOk returns a tuple with the UseDdnsGenerateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsGenerateHostname

`func (o *Rangetemplate) SetUseDdnsGenerateHostname(v bool)`

SetUseDdnsGenerateHostname sets UseDdnsGenerateHostname field to given value.

### HasUseDdnsGenerateHostname

`func (o *Rangetemplate) HasUseDdnsGenerateHostname() bool`

HasUseDdnsGenerateHostname returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *Rangetemplate) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *Rangetemplate) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *Rangetemplate) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *Rangetemplate) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseEmailList

`func (o *Rangetemplate) GetUseEmailList() bool`

GetUseEmailList returns the UseEmailList field if non-nil, zero value otherwise.

### GetUseEmailListOk

`func (o *Rangetemplate) GetUseEmailListOk() (*bool, bool)`

GetUseEmailListOk returns a tuple with the UseEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailList

`func (o *Rangetemplate) SetUseEmailList(v bool)`

SetUseEmailList sets UseEmailList field to given value.

### HasUseEmailList

`func (o *Rangetemplate) HasUseEmailList() bool`

HasUseEmailList returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *Rangetemplate) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *Rangetemplate) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *Rangetemplate) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *Rangetemplate) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseEnableDhcpThresholds

`func (o *Rangetemplate) GetUseEnableDhcpThresholds() bool`

GetUseEnableDhcpThresholds returns the UseEnableDhcpThresholds field if non-nil, zero value otherwise.

### GetUseEnableDhcpThresholdsOk

`func (o *Rangetemplate) GetUseEnableDhcpThresholdsOk() (*bool, bool)`

GetUseEnableDhcpThresholdsOk returns a tuple with the UseEnableDhcpThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDhcpThresholds

`func (o *Rangetemplate) SetUseEnableDhcpThresholds(v bool)`

SetUseEnableDhcpThresholds sets UseEnableDhcpThresholds field to given value.

### HasUseEnableDhcpThresholds

`func (o *Rangetemplate) HasUseEnableDhcpThresholds() bool`

HasUseEnableDhcpThresholds returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *Rangetemplate) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *Rangetemplate) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *Rangetemplate) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *Rangetemplate) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseKnownClients

`func (o *Rangetemplate) GetUseKnownClients() bool`

GetUseKnownClients returns the UseKnownClients field if non-nil, zero value otherwise.

### GetUseKnownClientsOk

`func (o *Rangetemplate) GetUseKnownClientsOk() (*bool, bool)`

GetUseKnownClientsOk returns a tuple with the UseKnownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKnownClients

`func (o *Rangetemplate) SetUseKnownClients(v bool)`

SetUseKnownClients sets UseKnownClients field to given value.

### HasUseKnownClients

`func (o *Rangetemplate) HasUseKnownClients() bool`

HasUseKnownClients returns a boolean if a field has been set.

### GetUseLeaseScavengeTime

`func (o *Rangetemplate) GetUseLeaseScavengeTime() bool`

GetUseLeaseScavengeTime returns the UseLeaseScavengeTime field if non-nil, zero value otherwise.

### GetUseLeaseScavengeTimeOk

`func (o *Rangetemplate) GetUseLeaseScavengeTimeOk() (*bool, bool)`

GetUseLeaseScavengeTimeOk returns a tuple with the UseLeaseScavengeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLeaseScavengeTime

`func (o *Rangetemplate) SetUseLeaseScavengeTime(v bool)`

SetUseLeaseScavengeTime sets UseLeaseScavengeTime field to given value.

### HasUseLeaseScavengeTime

`func (o *Rangetemplate) HasUseLeaseScavengeTime() bool`

HasUseLeaseScavengeTime returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Rangetemplate) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Rangetemplate) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Rangetemplate) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Rangetemplate) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseMsOptions

`func (o *Rangetemplate) GetUseMsOptions() bool`

GetUseMsOptions returns the UseMsOptions field if non-nil, zero value otherwise.

### GetUseMsOptionsOk

`func (o *Rangetemplate) GetUseMsOptionsOk() (*bool, bool)`

GetUseMsOptionsOk returns a tuple with the UseMsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsOptions

`func (o *Rangetemplate) SetUseMsOptions(v bool)`

SetUseMsOptions sets UseMsOptions field to given value.

### HasUseMsOptions

`func (o *Rangetemplate) HasUseMsOptions() bool`

HasUseMsOptions returns a boolean if a field has been set.

### GetUseNextserver

`func (o *Rangetemplate) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *Rangetemplate) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *Rangetemplate) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *Rangetemplate) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *Rangetemplate) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Rangetemplate) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Rangetemplate) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Rangetemplate) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *Rangetemplate) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *Rangetemplate) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *Rangetemplate) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *Rangetemplate) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Rangetemplate) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Rangetemplate) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Rangetemplate) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Rangetemplate) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetUseUnknownClients

`func (o *Rangetemplate) GetUseUnknownClients() bool`

GetUseUnknownClients returns the UseUnknownClients field if non-nil, zero value otherwise.

### GetUseUnknownClientsOk

`func (o *Rangetemplate) GetUseUnknownClientsOk() (*bool, bool)`

GetUseUnknownClientsOk returns a tuple with the UseUnknownClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUnknownClients

`func (o *Rangetemplate) SetUseUnknownClients(v bool)`

SetUseUnknownClients sets UseUnknownClients field to given value.

### HasUseUnknownClients

`func (o *Rangetemplate) HasUseUnknownClients() bool`

HasUseUnknownClients returns a boolean if a field has been set.

### GetUseUpdateDnsOnLeaseRenewal

`func (o *Rangetemplate) GetUseUpdateDnsOnLeaseRenewal() bool`

GetUseUpdateDnsOnLeaseRenewal returns the UseUpdateDnsOnLeaseRenewal field if non-nil, zero value otherwise.

### GetUseUpdateDnsOnLeaseRenewalOk

`func (o *Rangetemplate) GetUseUpdateDnsOnLeaseRenewalOk() (*bool, bool)`

GetUseUpdateDnsOnLeaseRenewalOk returns a tuple with the UseUpdateDnsOnLeaseRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUpdateDnsOnLeaseRenewal

`func (o *Rangetemplate) SetUseUpdateDnsOnLeaseRenewal(v bool)`

SetUseUpdateDnsOnLeaseRenewal sets UseUpdateDnsOnLeaseRenewal field to given value.

### HasUseUpdateDnsOnLeaseRenewal

`func (o *Rangetemplate) HasUseUpdateDnsOnLeaseRenewal() bool`

HasUseUpdateDnsOnLeaseRenewal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


