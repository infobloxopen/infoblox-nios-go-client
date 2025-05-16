# GetFixedaddresstemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Bootfile** | Pointer to **string** | The boot file name for the fixed address. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages. | [optional] 
**Bootserver** | Pointer to **string** | The boot server address for the fixed address. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of a fixed address template object. | [optional] 
**DdnsDomainname** | Pointer to **string** | The dynamic DNS domain name the appliance uses specifically for DDNS updates for this fixed address. | [optional] 
**DdnsHostname** | Pointer to **string** | The DDNS host name for this fixed address. | [optional] 
**DenyBootp** | Pointer to **bool** | Determines if BOOTP settings are disabled and BOOTP requests will be denied. | [optional] 
**EnableDdns** | Pointer to **bool** | Determines if the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers. | [optional] 
**EnablePxeLeaseTime** | Pointer to **bool** | Set this to True if you want the DHCP server to use a different lease time for PXE clients. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IgnoreDhcpOptionListRequest** | Pointer to **bool** | If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested. | [optional] 
**LogicFilterRules** | Pointer to [**[]FixedaddresstemplateLogicFilterRules**](FixedaddresstemplateLogicFilterRules.md) | This field contains the logic filters to be applied on this fixed address. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**Name** | Pointer to **string** | The name of a fixed address template object. | [optional] 
**Nextserver** | Pointer to **string** | The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot. | [optional] 
**NumberOfAddresses** | Pointer to **int64** | The number of addresses for this fixed address. | [optional] 
**Offset** | Pointer to **int64** | The start address offset for this fixed address. | [optional] 
**Options** | Pointer to [**[]FixedaddresstemplateOptions**](FixedaddresstemplateOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | The PXE lease time value for a DHCP Fixed Address object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached. | [optional] 
**UseBootfile** | Pointer to **bool** | Use flag for: bootfile | [optional] 
**UseBootserver** | Pointer to **bool** | Use flag for: bootserver | [optional] 
**UseDdnsDomainname** | Pointer to **bool** | Use flag for: ddns_domainname | [optional] 
**UseDenyBootp** | Pointer to **bool** | Use flag for: deny_bootp | [optional] 
**UseEnableDdns** | Pointer to **bool** | Use flag for: enable_ddns | [optional] 
**UseIgnoreDhcpOptionListRequest** | Pointer to **bool** | Use flag for: ignore_dhcp_option_list_request | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseNextserver** | Pointer to **bool** | Use flag for: nextserver | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePxeLeaseTime** | Pointer to **bool** | Use flag for: pxe_lease_time | [optional] 
**Result** | Pointer to [**Fixedaddresstemplate**](Fixedaddresstemplate.md) |  | [optional] 

## Methods

### NewGetFixedaddresstemplateResponse

`func NewGetFixedaddresstemplateResponse() *GetFixedaddresstemplateResponse`

NewGetFixedaddresstemplateResponse instantiates a new GetFixedaddresstemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFixedaddresstemplateResponseWithDefaults

`func NewGetFixedaddresstemplateResponseWithDefaults() *GetFixedaddresstemplateResponse`

NewGetFixedaddresstemplateResponseWithDefaults instantiates a new GetFixedaddresstemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFixedaddresstemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFixedaddresstemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFixedaddresstemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFixedaddresstemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBootfile

`func (o *GetFixedaddresstemplateResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetFixedaddresstemplateResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetFixedaddresstemplateResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetFixedaddresstemplateResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetFixedaddresstemplateResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetFixedaddresstemplateResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetFixedaddresstemplateResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetFixedaddresstemplateResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetComment

`func (o *GetFixedaddresstemplateResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFixedaddresstemplateResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFixedaddresstemplateResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFixedaddresstemplateResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDomainname

`func (o *GetFixedaddresstemplateResponse) GetDdnsDomainname() string`

GetDdnsDomainname returns the DdnsDomainname field if non-nil, zero value otherwise.

### GetDdnsDomainnameOk

`func (o *GetFixedaddresstemplateResponse) GetDdnsDomainnameOk() (*string, bool)`

GetDdnsDomainnameOk returns a tuple with the DdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomainname

`func (o *GetFixedaddresstemplateResponse) SetDdnsDomainname(v string)`

SetDdnsDomainname sets DdnsDomainname field to given value.

### HasDdnsDomainname

`func (o *GetFixedaddresstemplateResponse) HasDdnsDomainname() bool`

HasDdnsDomainname returns a boolean if a field has been set.

### GetDdnsHostname

`func (o *GetFixedaddresstemplateResponse) GetDdnsHostname() string`

GetDdnsHostname returns the DdnsHostname field if non-nil, zero value otherwise.

### GetDdnsHostnameOk

`func (o *GetFixedaddresstemplateResponse) GetDdnsHostnameOk() (*string, bool)`

GetDdnsHostnameOk returns a tuple with the DdnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostname

`func (o *GetFixedaddresstemplateResponse) SetDdnsHostname(v string)`

SetDdnsHostname sets DdnsHostname field to given value.

### HasDdnsHostname

`func (o *GetFixedaddresstemplateResponse) HasDdnsHostname() bool`

HasDdnsHostname returns a boolean if a field has been set.

### GetDenyBootp

`func (o *GetFixedaddresstemplateResponse) GetDenyBootp() bool`

GetDenyBootp returns the DenyBootp field if non-nil, zero value otherwise.

### GetDenyBootpOk

`func (o *GetFixedaddresstemplateResponse) GetDenyBootpOk() (*bool, bool)`

GetDenyBootpOk returns a tuple with the DenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyBootp

`func (o *GetFixedaddresstemplateResponse) SetDenyBootp(v bool)`

SetDenyBootp sets DenyBootp field to given value.

### HasDenyBootp

`func (o *GetFixedaddresstemplateResponse) HasDenyBootp() bool`

HasDenyBootp returns a boolean if a field has been set.

### GetEnableDdns

`func (o *GetFixedaddresstemplateResponse) GetEnableDdns() bool`

GetEnableDdns returns the EnableDdns field if non-nil, zero value otherwise.

### GetEnableDdnsOk

`func (o *GetFixedaddresstemplateResponse) GetEnableDdnsOk() (*bool, bool)`

GetEnableDdnsOk returns a tuple with the EnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDdns

`func (o *GetFixedaddresstemplateResponse) SetEnableDdns(v bool)`

SetEnableDdns sets EnableDdns field to given value.

### HasEnableDdns

`func (o *GetFixedaddresstemplateResponse) HasEnableDdns() bool`

HasEnableDdns returns a boolean if a field has been set.

### GetEnablePxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) GetEnablePxeLeaseTime() bool`

GetEnablePxeLeaseTime returns the EnablePxeLeaseTime field if non-nil, zero value otherwise.

### GetEnablePxeLeaseTimeOk

`func (o *GetFixedaddresstemplateResponse) GetEnablePxeLeaseTimeOk() (*bool, bool)`

GetEnablePxeLeaseTimeOk returns a tuple with the EnablePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) SetEnablePxeLeaseTime(v bool)`

SetEnablePxeLeaseTime sets EnablePxeLeaseTime field to given value.

### HasEnablePxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) HasEnablePxeLeaseTime() bool`

HasEnablePxeLeaseTime returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetFixedaddresstemplateResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetFixedaddresstemplateResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetFixedaddresstemplateResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetFixedaddresstemplateResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIgnoreDhcpOptionListRequest

`func (o *GetFixedaddresstemplateResponse) GetIgnoreDhcpOptionListRequest() bool`

GetIgnoreDhcpOptionListRequest returns the IgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetIgnoreDhcpOptionListRequestOk

`func (o *GetFixedaddresstemplateResponse) GetIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetIgnoreDhcpOptionListRequestOk returns a tuple with the IgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDhcpOptionListRequest

`func (o *GetFixedaddresstemplateResponse) SetIgnoreDhcpOptionListRequest(v bool)`

SetIgnoreDhcpOptionListRequest sets IgnoreDhcpOptionListRequest field to given value.

### HasIgnoreDhcpOptionListRequest

`func (o *GetFixedaddresstemplateResponse) HasIgnoreDhcpOptionListRequest() bool`

HasIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetFixedaddresstemplateResponse) GetLogicFilterRules() []FixedaddresstemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetFixedaddresstemplateResponse) GetLogicFilterRulesOk() (*[]FixedaddresstemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetFixedaddresstemplateResponse) SetLogicFilterRules(v []FixedaddresstemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetFixedaddresstemplateResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetName

`func (o *GetFixedaddresstemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFixedaddresstemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFixedaddresstemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFixedaddresstemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextserver

`func (o *GetFixedaddresstemplateResponse) GetNextserver() string`

GetNextserver returns the Nextserver field if non-nil, zero value otherwise.

### GetNextserverOk

`func (o *GetFixedaddresstemplateResponse) GetNextserverOk() (*string, bool)`

GetNextserverOk returns a tuple with the Nextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextserver

`func (o *GetFixedaddresstemplateResponse) SetNextserver(v string)`

SetNextserver sets Nextserver field to given value.

### HasNextserver

`func (o *GetFixedaddresstemplateResponse) HasNextserver() bool`

HasNextserver returns a boolean if a field has been set.

### GetNumberOfAddresses

`func (o *GetFixedaddresstemplateResponse) GetNumberOfAddresses() int64`

GetNumberOfAddresses returns the NumberOfAddresses field if non-nil, zero value otherwise.

### GetNumberOfAddressesOk

`func (o *GetFixedaddresstemplateResponse) GetNumberOfAddressesOk() (*int64, bool)`

GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddresses

`func (o *GetFixedaddresstemplateResponse) SetNumberOfAddresses(v int64)`

SetNumberOfAddresses sets NumberOfAddresses field to given value.

### HasNumberOfAddresses

`func (o *GetFixedaddresstemplateResponse) HasNumberOfAddresses() bool`

HasNumberOfAddresses returns a boolean if a field has been set.

### GetOffset

`func (o *GetFixedaddresstemplateResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetFixedaddresstemplateResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetFixedaddresstemplateResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetFixedaddresstemplateResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptions

`func (o *GetFixedaddresstemplateResponse) GetOptions() []FixedaddresstemplateOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetFixedaddresstemplateResponse) GetOptionsOk() (*[]FixedaddresstemplateOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetFixedaddresstemplateResponse) SetOptions(v []FixedaddresstemplateOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetFixedaddresstemplateResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetFixedaddresstemplateResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetUseBootfile

`func (o *GetFixedaddresstemplateResponse) GetUseBootfile() bool`

GetUseBootfile returns the UseBootfile field if non-nil, zero value otherwise.

### GetUseBootfileOk

`func (o *GetFixedaddresstemplateResponse) GetUseBootfileOk() (*bool, bool)`

GetUseBootfileOk returns a tuple with the UseBootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootfile

`func (o *GetFixedaddresstemplateResponse) SetUseBootfile(v bool)`

SetUseBootfile sets UseBootfile field to given value.

### HasUseBootfile

`func (o *GetFixedaddresstemplateResponse) HasUseBootfile() bool`

HasUseBootfile returns a boolean if a field has been set.

### GetUseBootserver

`func (o *GetFixedaddresstemplateResponse) GetUseBootserver() bool`

GetUseBootserver returns the UseBootserver field if non-nil, zero value otherwise.

### GetUseBootserverOk

`func (o *GetFixedaddresstemplateResponse) GetUseBootserverOk() (*bool, bool)`

GetUseBootserverOk returns a tuple with the UseBootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBootserver

`func (o *GetFixedaddresstemplateResponse) SetUseBootserver(v bool)`

SetUseBootserver sets UseBootserver field to given value.

### HasUseBootserver

`func (o *GetFixedaddresstemplateResponse) HasUseBootserver() bool`

HasUseBootserver returns a boolean if a field has been set.

### GetUseDdnsDomainname

`func (o *GetFixedaddresstemplateResponse) GetUseDdnsDomainname() bool`

GetUseDdnsDomainname returns the UseDdnsDomainname field if non-nil, zero value otherwise.

### GetUseDdnsDomainnameOk

`func (o *GetFixedaddresstemplateResponse) GetUseDdnsDomainnameOk() (*bool, bool)`

GetUseDdnsDomainnameOk returns a tuple with the UseDdnsDomainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsDomainname

`func (o *GetFixedaddresstemplateResponse) SetUseDdnsDomainname(v bool)`

SetUseDdnsDomainname sets UseDdnsDomainname field to given value.

### HasUseDdnsDomainname

`func (o *GetFixedaddresstemplateResponse) HasUseDdnsDomainname() bool`

HasUseDdnsDomainname returns a boolean if a field has been set.

### GetUseDenyBootp

`func (o *GetFixedaddresstemplateResponse) GetUseDenyBootp() bool`

GetUseDenyBootp returns the UseDenyBootp field if non-nil, zero value otherwise.

### GetUseDenyBootpOk

`func (o *GetFixedaddresstemplateResponse) GetUseDenyBootpOk() (*bool, bool)`

GetUseDenyBootpOk returns a tuple with the UseDenyBootp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDenyBootp

`func (o *GetFixedaddresstemplateResponse) SetUseDenyBootp(v bool)`

SetUseDenyBootp sets UseDenyBootp field to given value.

### HasUseDenyBootp

`func (o *GetFixedaddresstemplateResponse) HasUseDenyBootp() bool`

HasUseDenyBootp returns a boolean if a field has been set.

### GetUseEnableDdns

`func (o *GetFixedaddresstemplateResponse) GetUseEnableDdns() bool`

GetUseEnableDdns returns the UseEnableDdns field if non-nil, zero value otherwise.

### GetUseEnableDdnsOk

`func (o *GetFixedaddresstemplateResponse) GetUseEnableDdnsOk() (*bool, bool)`

GetUseEnableDdnsOk returns a tuple with the UseEnableDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDdns

`func (o *GetFixedaddresstemplateResponse) SetUseEnableDdns(v bool)`

SetUseEnableDdns sets UseEnableDdns field to given value.

### HasUseEnableDdns

`func (o *GetFixedaddresstemplateResponse) HasUseEnableDdns() bool`

HasUseEnableDdns returns a boolean if a field has been set.

### GetUseIgnoreDhcpOptionListRequest

`func (o *GetFixedaddresstemplateResponse) GetUseIgnoreDhcpOptionListRequest() bool`

GetUseIgnoreDhcpOptionListRequest returns the UseIgnoreDhcpOptionListRequest field if non-nil, zero value otherwise.

### GetUseIgnoreDhcpOptionListRequestOk

`func (o *GetFixedaddresstemplateResponse) GetUseIgnoreDhcpOptionListRequestOk() (*bool, bool)`

GetUseIgnoreDhcpOptionListRequestOk returns a tuple with the UseIgnoreDhcpOptionListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgnoreDhcpOptionListRequest

`func (o *GetFixedaddresstemplateResponse) SetUseIgnoreDhcpOptionListRequest(v bool)`

SetUseIgnoreDhcpOptionListRequest sets UseIgnoreDhcpOptionListRequest field to given value.

### HasUseIgnoreDhcpOptionListRequest

`func (o *GetFixedaddresstemplateResponse) HasUseIgnoreDhcpOptionListRequest() bool`

HasUseIgnoreDhcpOptionListRequest returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetFixedaddresstemplateResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetFixedaddresstemplateResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetFixedaddresstemplateResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetFixedaddresstemplateResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseNextserver

`func (o *GetFixedaddresstemplateResponse) GetUseNextserver() bool`

GetUseNextserver returns the UseNextserver field if non-nil, zero value otherwise.

### GetUseNextserverOk

`func (o *GetFixedaddresstemplateResponse) GetUseNextserverOk() (*bool, bool)`

GetUseNextserverOk returns a tuple with the UseNextserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNextserver

`func (o *GetFixedaddresstemplateResponse) SetUseNextserver(v bool)`

SetUseNextserver sets UseNextserver field to given value.

### HasUseNextserver

`func (o *GetFixedaddresstemplateResponse) HasUseNextserver() bool`

HasUseNextserver returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetFixedaddresstemplateResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetFixedaddresstemplateResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetFixedaddresstemplateResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetFixedaddresstemplateResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) GetUsePxeLeaseTime() bool`

GetUsePxeLeaseTime returns the UsePxeLeaseTime field if non-nil, zero value otherwise.

### GetUsePxeLeaseTimeOk

`func (o *GetFixedaddresstemplateResponse) GetUsePxeLeaseTimeOk() (*bool, bool)`

GetUsePxeLeaseTimeOk returns a tuple with the UsePxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) SetUsePxeLeaseTime(v bool)`

SetUsePxeLeaseTime sets UsePxeLeaseTime field to given value.

### HasUsePxeLeaseTime

`func (o *GetFixedaddresstemplateResponse) HasUsePxeLeaseTime() bool`

HasUsePxeLeaseTime returns a boolean if a field has been set.

### GetResult

`func (o *GetFixedaddresstemplateResponse) GetResult() Fixedaddresstemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFixedaddresstemplateResponse) GetResultOk() (*Fixedaddresstemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFixedaddresstemplateResponse) SetResult(v Fixedaddresstemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFixedaddresstemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


