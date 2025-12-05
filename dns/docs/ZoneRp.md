# ZoneRp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IP address of the server that is serving this zone. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the zone; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a zone is disabled or not. When this is set to False, the zone is enabled. | [optional] 
**DisplayDomain** | Pointer to **string** | The displayed name of the DNS zone. | [optional] [readonly] 
**DnsSoaEmail** | Pointer to **string** | The SOA email for the zone in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ZoneRpExternalPrimaries**](ZoneRpExternalPrimaries.md) | The list of external primary servers. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ZoneRpExternalSecondaries**](ZoneRpExternalSecondaries.md) | The list of external secondary servers. | [optional] 
**FireeyeRuleMapping** | Pointer to [**ZoneRpFireeyeRuleMapping**](ZoneRpFireeyeRuleMapping.md) |  | [optional] 
**Fqdn** | Pointer to **string** | The name of this DNS zone in FQDN format. | [optional] 
**GridPrimary** | Pointer to [**[]ZoneRpGridPrimary**](ZoneRpGridPrimary.md) | The grid primary servers for this zone. | [optional] 
**GridSecondaries** | Pointer to [**[]ZoneRpGridSecondaries**](ZoneRpGridSecondaries.md) | The list with Grid members that are secondary servers for this zone. | [optional] 
**Locked** | Pointer to **bool** | If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked. | [optional] 
**LockedBy** | Pointer to **string** | The name of a superuser or the administrator who locked this zone. | [optional] [readonly] 
**LogRpz** | Pointer to **bool** | Determines whether RPZ logging enabled or not at zone level. When this is set to False, the logging is disabled. | [optional] 
**MaskPrefix** | Pointer to **string** | IPv4 Netmask or IPv6 prefix for this zone. | [optional] [readonly] 
**MemberSoaMnames** | Pointer to [**[]ZoneRpMemberSoaMnames**](ZoneRpMemberSoaMnames.md) | The list of per-member SOA MNAME information. | [optional] 
**MemberSoaSerials** | Pointer to [**[]ZoneRpMemberSoaSerials**](ZoneRpMemberSoaSerials.md) | The list of per-member SOA serial information. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this zone resides. | [optional] [readonly] 
**NsGroup** | Pointer to **string** | The name server group that serves DNS for this zone. | [optional] 
**Parent** | Pointer to **string** | The parent zone of this zone. Note that when searching for reverse zones, the \&quot;in-addr.arpa\&quot; notation should be used. | [optional] [readonly] 
**Prefix** | Pointer to **string** | The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B. | [optional] 
**PrimaryType** | Pointer to **string** | The type of the primary server. | [optional] [readonly] 
**RecordNamePolicy** | Pointer to **string** | The hostname policy for records under this zone. | [optional] 
**RpzDropIpRuleEnabled** | Pointer to **bool** | Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv4** | Pointer to **int64** | The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length. | [optional] 
**RpzDropIpRuleMinPrefixLengthIpv6** | Pointer to **int64** | The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length. | [optional] 
**RpzLastUpdatedTime** | Pointer to **int64** | The timestamp of the last update for zone data. | [optional] [readonly] 
**RpzPolicy** | Pointer to **string** | The response policy zone override policy. | [optional] 
**RpzPriority** | Pointer to **int64** | The priority of this response policy zone. | [optional] [readonly] 
**RpzPriorityEnd** | Pointer to **int64** | This number is for UI to identify the end of qualified zone list. | [optional] [readonly] 
**RpzSeverity** | Pointer to **string** | The severity of this response policy zone. | [optional] 
**RpzType** | Pointer to **string** | The type of rpz zone. | [optional] 
**SetSoaSerialNumber** | Pointer to **bool** | The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \&quot;soa_serial_number\&quot; and pass \&quot;set_soa_serial_number\&quot; as True. | [optional] 
**SoaDefaultTtl** | Pointer to **int64** | The Time to Live (TTL) value of the SOA record of this zone. This value is the number of seconds that data is cached. | [optional] 
**SoaEmail** | Pointer to **string** | The SOA email value for this zone. This value can be in unicode format. | [optional] 
**SoaExpire** | Pointer to **int64** | This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week. | [optional] 
**SoaNegativeTtl** | Pointer to **int64** | The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \&quot;Does Not Respond\&quot; responses. | [optional] 
**SoaRefresh** | Pointer to **int64** | This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not. | [optional] 
**SoaRetry** | Pointer to **int64** | This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs. | [optional] 
**SoaSerial** | Pointer to **int64** | The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \&quot;soa_serial_number\&quot; and pass \&quot;set_soa_serial_number\&quot; as True. | [optional] 
**SubstituteName** | Pointer to **string** | The canonical name of redirect target in substitute policy of response policy zone. | [optional] 
**UseExternalPrimary** | Pointer to **bool** | This flag controls whether the zone is using an external primary. | [optional] 
**UseGridZoneTimer** | Pointer to **bool** | Use flag for: soa_default_ttl , soa_expire, soa_negative_ttl, soa_refresh, soa_retry | [optional] 
**UseLogRpz** | Pointer to **bool** | Use flag for: log_rpz | [optional] 
**UseRecordNamePolicy** | Pointer to **bool** | Use flag for: record_name_policy | [optional] 
**UseRpzDropIpRule** | Pointer to **bool** | Use flag for: rpz_drop_ip_rule_enabled , rpz_drop_ip_rule_min_prefix_length_ipv4, rpz_drop_ip_rule_min_prefix_length_ipv6 | [optional] 
**UseSoaEmail** | Pointer to **bool** | Use flag for: soa_email | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the zone resides. Example \&quot;external\&quot;. | [optional] 

## Methods

### NewZoneRp

`func NewZoneRp() *ZoneRp`

NewZoneRp instantiates a new ZoneRp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneRpWithDefaults

`func NewZoneRpWithDefaults() *ZoneRp`

NewZoneRpWithDefaults instantiates a new ZoneRp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ZoneRp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ZoneRp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ZoneRp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ZoneRp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *ZoneRp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ZoneRp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ZoneRp) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ZoneRp) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *ZoneRp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ZoneRp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ZoneRp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ZoneRp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *ZoneRp) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ZoneRp) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ZoneRp) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ZoneRp) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisplayDomain

`func (o *ZoneRp) GetDisplayDomain() string`

GetDisplayDomain returns the DisplayDomain field if non-nil, zero value otherwise.

### GetDisplayDomainOk

`func (o *ZoneRp) GetDisplayDomainOk() (*string, bool)`

GetDisplayDomainOk returns a tuple with the DisplayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDomain

`func (o *ZoneRp) SetDisplayDomain(v string)`

SetDisplayDomain sets DisplayDomain field to given value.

### HasDisplayDomain

`func (o *ZoneRp) HasDisplayDomain() bool`

HasDisplayDomain returns a boolean if a field has been set.

### GetDnsSoaEmail

`func (o *ZoneRp) GetDnsSoaEmail() string`

GetDnsSoaEmail returns the DnsSoaEmail field if non-nil, zero value otherwise.

### GetDnsSoaEmailOk

`func (o *ZoneRp) GetDnsSoaEmailOk() (*string, bool)`

GetDnsSoaEmailOk returns a tuple with the DnsSoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSoaEmail

`func (o *ZoneRp) SetDnsSoaEmail(v string)`

SetDnsSoaEmail sets DnsSoaEmail field to given value.

### HasDnsSoaEmail

`func (o *ZoneRp) HasDnsSoaEmail() bool`

HasDnsSoaEmail returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *ZoneRp) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *ZoneRp) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *ZoneRp) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *ZoneRp) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *ZoneRp) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *ZoneRp) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *ZoneRp) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *ZoneRp) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *ZoneRp) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *ZoneRp) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *ZoneRp) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *ZoneRp) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *ZoneRp) GetExternalPrimaries() []ZoneRpExternalPrimaries`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *ZoneRp) GetExternalPrimariesOk() (*[]ZoneRpExternalPrimaries, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *ZoneRp) SetExternalPrimaries(v []ZoneRpExternalPrimaries)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *ZoneRp) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *ZoneRp) GetExternalSecondaries() []ZoneRpExternalSecondaries`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *ZoneRp) GetExternalSecondariesOk() (*[]ZoneRpExternalSecondaries, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *ZoneRp) SetExternalSecondaries(v []ZoneRpExternalSecondaries)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *ZoneRp) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetFireeyeRuleMapping

`func (o *ZoneRp) GetFireeyeRuleMapping() ZoneRpFireeyeRuleMapping`

GetFireeyeRuleMapping returns the FireeyeRuleMapping field if non-nil, zero value otherwise.

### GetFireeyeRuleMappingOk

`func (o *ZoneRp) GetFireeyeRuleMappingOk() (*ZoneRpFireeyeRuleMapping, bool)`

GetFireeyeRuleMappingOk returns a tuple with the FireeyeRuleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireeyeRuleMapping

`func (o *ZoneRp) SetFireeyeRuleMapping(v ZoneRpFireeyeRuleMapping)`

SetFireeyeRuleMapping sets FireeyeRuleMapping field to given value.

### HasFireeyeRuleMapping

`func (o *ZoneRp) HasFireeyeRuleMapping() bool`

HasFireeyeRuleMapping returns a boolean if a field has been set.

### GetFqdn

`func (o *ZoneRp) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ZoneRp) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ZoneRp) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ZoneRp) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGridPrimary

`func (o *ZoneRp) GetGridPrimary() []ZoneRpGridPrimary`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *ZoneRp) GetGridPrimaryOk() (*[]ZoneRpGridPrimary, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *ZoneRp) SetGridPrimary(v []ZoneRpGridPrimary)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *ZoneRp) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *ZoneRp) GetGridSecondaries() []ZoneRpGridSecondaries`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *ZoneRp) GetGridSecondariesOk() (*[]ZoneRpGridSecondaries, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *ZoneRp) SetGridSecondaries(v []ZoneRpGridSecondaries)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *ZoneRp) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetLocked

`func (o *ZoneRp) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ZoneRp) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ZoneRp) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ZoneRp) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *ZoneRp) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *ZoneRp) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *ZoneRp) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *ZoneRp) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetLogRpz

`func (o *ZoneRp) GetLogRpz() bool`

GetLogRpz returns the LogRpz field if non-nil, zero value otherwise.

### GetLogRpzOk

`func (o *ZoneRp) GetLogRpzOk() (*bool, bool)`

GetLogRpzOk returns a tuple with the LogRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRpz

`func (o *ZoneRp) SetLogRpz(v bool)`

SetLogRpz sets LogRpz field to given value.

### HasLogRpz

`func (o *ZoneRp) HasLogRpz() bool`

HasLogRpz returns a boolean if a field has been set.

### GetMaskPrefix

`func (o *ZoneRp) GetMaskPrefix() string`

GetMaskPrefix returns the MaskPrefix field if non-nil, zero value otherwise.

### GetMaskPrefixOk

`func (o *ZoneRp) GetMaskPrefixOk() (*string, bool)`

GetMaskPrefixOk returns a tuple with the MaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPrefix

`func (o *ZoneRp) SetMaskPrefix(v string)`

SetMaskPrefix sets MaskPrefix field to given value.

### HasMaskPrefix

`func (o *ZoneRp) HasMaskPrefix() bool`

HasMaskPrefix returns a boolean if a field has been set.

### GetMemberSoaMnames

`func (o *ZoneRp) GetMemberSoaMnames() []ZoneRpMemberSoaMnames`

GetMemberSoaMnames returns the MemberSoaMnames field if non-nil, zero value otherwise.

### GetMemberSoaMnamesOk

`func (o *ZoneRp) GetMemberSoaMnamesOk() (*[]ZoneRpMemberSoaMnames, bool)`

GetMemberSoaMnamesOk returns a tuple with the MemberSoaMnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSoaMnames

`func (o *ZoneRp) SetMemberSoaMnames(v []ZoneRpMemberSoaMnames)`

SetMemberSoaMnames sets MemberSoaMnames field to given value.

### HasMemberSoaMnames

`func (o *ZoneRp) HasMemberSoaMnames() bool`

HasMemberSoaMnames returns a boolean if a field has been set.

### GetMemberSoaSerials

`func (o *ZoneRp) GetMemberSoaSerials() []ZoneRpMemberSoaSerials`

GetMemberSoaSerials returns the MemberSoaSerials field if non-nil, zero value otherwise.

### GetMemberSoaSerialsOk

`func (o *ZoneRp) GetMemberSoaSerialsOk() (*[]ZoneRpMemberSoaSerials, bool)`

GetMemberSoaSerialsOk returns a tuple with the MemberSoaSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSoaSerials

`func (o *ZoneRp) SetMemberSoaSerials(v []ZoneRpMemberSoaSerials)`

SetMemberSoaSerials sets MemberSoaSerials field to given value.

### HasMemberSoaSerials

`func (o *ZoneRp) HasMemberSoaSerials() bool`

HasMemberSoaSerials returns a boolean if a field has been set.

### GetNetworkView

`func (o *ZoneRp) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *ZoneRp) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *ZoneRp) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *ZoneRp) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNsGroup

`func (o *ZoneRp) GetNsGroup() string`

GetNsGroup returns the NsGroup field if non-nil, zero value otherwise.

### GetNsGroupOk

`func (o *ZoneRp) GetNsGroupOk() (*string, bool)`

GetNsGroupOk returns a tuple with the NsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsGroup

`func (o *ZoneRp) SetNsGroup(v string)`

SetNsGroup sets NsGroup field to given value.

### HasNsGroup

`func (o *ZoneRp) HasNsGroup() bool`

HasNsGroup returns a boolean if a field has been set.

### GetParent

`func (o *ZoneRp) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ZoneRp) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ZoneRp) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ZoneRp) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrefix

`func (o *ZoneRp) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ZoneRp) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ZoneRp) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ZoneRp) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrimaryType

`func (o *ZoneRp) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *ZoneRp) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *ZoneRp) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *ZoneRp) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetRecordNamePolicy

`func (o *ZoneRp) GetRecordNamePolicy() string`

GetRecordNamePolicy returns the RecordNamePolicy field if non-nil, zero value otherwise.

### GetRecordNamePolicyOk

`func (o *ZoneRp) GetRecordNamePolicyOk() (*string, bool)`

GetRecordNamePolicyOk returns a tuple with the RecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNamePolicy

`func (o *ZoneRp) SetRecordNamePolicy(v string)`

SetRecordNamePolicy sets RecordNamePolicy field to given value.

### HasRecordNamePolicy

`func (o *ZoneRp) HasRecordNamePolicy() bool`

HasRecordNamePolicy returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *ZoneRp) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *ZoneRp) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *ZoneRp) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *ZoneRp) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *ZoneRp) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *ZoneRp) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *ZoneRp) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *ZoneRp) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *ZoneRp) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *ZoneRp) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *ZoneRp) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *ZoneRp) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzLastUpdatedTime

`func (o *ZoneRp) GetRpzLastUpdatedTime() int64`

GetRpzLastUpdatedTime returns the RpzLastUpdatedTime field if non-nil, zero value otherwise.

### GetRpzLastUpdatedTimeOk

`func (o *ZoneRp) GetRpzLastUpdatedTimeOk() (*int64, bool)`

GetRpzLastUpdatedTimeOk returns a tuple with the RpzLastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzLastUpdatedTime

`func (o *ZoneRp) SetRpzLastUpdatedTime(v int64)`

SetRpzLastUpdatedTime sets RpzLastUpdatedTime field to given value.

### HasRpzLastUpdatedTime

`func (o *ZoneRp) HasRpzLastUpdatedTime() bool`

HasRpzLastUpdatedTime returns a boolean if a field has been set.

### GetRpzPolicy

`func (o *ZoneRp) GetRpzPolicy() string`

GetRpzPolicy returns the RpzPolicy field if non-nil, zero value otherwise.

### GetRpzPolicyOk

`func (o *ZoneRp) GetRpzPolicyOk() (*string, bool)`

GetRpzPolicyOk returns a tuple with the RpzPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPolicy

`func (o *ZoneRp) SetRpzPolicy(v string)`

SetRpzPolicy sets RpzPolicy field to given value.

### HasRpzPolicy

`func (o *ZoneRp) HasRpzPolicy() bool`

HasRpzPolicy returns a boolean if a field has been set.

### GetRpzPriority

`func (o *ZoneRp) GetRpzPriority() int64`

GetRpzPriority returns the RpzPriority field if non-nil, zero value otherwise.

### GetRpzPriorityOk

`func (o *ZoneRp) GetRpzPriorityOk() (*int64, bool)`

GetRpzPriorityOk returns a tuple with the RpzPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPriority

`func (o *ZoneRp) SetRpzPriority(v int64)`

SetRpzPriority sets RpzPriority field to given value.

### HasRpzPriority

`func (o *ZoneRp) HasRpzPriority() bool`

HasRpzPriority returns a boolean if a field has been set.

### GetRpzPriorityEnd

`func (o *ZoneRp) GetRpzPriorityEnd() int64`

GetRpzPriorityEnd returns the RpzPriorityEnd field if non-nil, zero value otherwise.

### GetRpzPriorityEndOk

`func (o *ZoneRp) GetRpzPriorityEndOk() (*int64, bool)`

GetRpzPriorityEndOk returns a tuple with the RpzPriorityEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPriorityEnd

`func (o *ZoneRp) SetRpzPriorityEnd(v int64)`

SetRpzPriorityEnd sets RpzPriorityEnd field to given value.

### HasRpzPriorityEnd

`func (o *ZoneRp) HasRpzPriorityEnd() bool`

HasRpzPriorityEnd returns a boolean if a field has been set.

### GetRpzSeverity

`func (o *ZoneRp) GetRpzSeverity() string`

GetRpzSeverity returns the RpzSeverity field if non-nil, zero value otherwise.

### GetRpzSeverityOk

`func (o *ZoneRp) GetRpzSeverityOk() (*string, bool)`

GetRpzSeverityOk returns a tuple with the RpzSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzSeverity

`func (o *ZoneRp) SetRpzSeverity(v string)`

SetRpzSeverity sets RpzSeverity field to given value.

### HasRpzSeverity

`func (o *ZoneRp) HasRpzSeverity() bool`

HasRpzSeverity returns a boolean if a field has been set.

### GetRpzType

`func (o *ZoneRp) GetRpzType() string`

GetRpzType returns the RpzType field if non-nil, zero value otherwise.

### GetRpzTypeOk

`func (o *ZoneRp) GetRpzTypeOk() (*string, bool)`

GetRpzTypeOk returns a tuple with the RpzType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzType

`func (o *ZoneRp) SetRpzType(v string)`

SetRpzType sets RpzType field to given value.

### HasRpzType

`func (o *ZoneRp) HasRpzType() bool`

HasRpzType returns a boolean if a field has been set.

### GetSetSoaSerialNumber

`func (o *ZoneRp) GetSetSoaSerialNumber() bool`

GetSetSoaSerialNumber returns the SetSoaSerialNumber field if non-nil, zero value otherwise.

### GetSetSoaSerialNumberOk

`func (o *ZoneRp) GetSetSoaSerialNumberOk() (*bool, bool)`

GetSetSoaSerialNumberOk returns a tuple with the SetSoaSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSoaSerialNumber

`func (o *ZoneRp) SetSetSoaSerialNumber(v bool)`

SetSetSoaSerialNumber sets SetSoaSerialNumber field to given value.

### HasSetSoaSerialNumber

`func (o *ZoneRp) HasSetSoaSerialNumber() bool`

HasSetSoaSerialNumber returns a boolean if a field has been set.

### GetSoaDefaultTtl

`func (o *ZoneRp) GetSoaDefaultTtl() int64`

GetSoaDefaultTtl returns the SoaDefaultTtl field if non-nil, zero value otherwise.

### GetSoaDefaultTtlOk

`func (o *ZoneRp) GetSoaDefaultTtlOk() (*int64, bool)`

GetSoaDefaultTtlOk returns a tuple with the SoaDefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaDefaultTtl

`func (o *ZoneRp) SetSoaDefaultTtl(v int64)`

SetSoaDefaultTtl sets SoaDefaultTtl field to given value.

### HasSoaDefaultTtl

`func (o *ZoneRp) HasSoaDefaultTtl() bool`

HasSoaDefaultTtl returns a boolean if a field has been set.

### GetSoaEmail

`func (o *ZoneRp) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *ZoneRp) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *ZoneRp) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *ZoneRp) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetSoaExpire

`func (o *ZoneRp) GetSoaExpire() int64`

GetSoaExpire returns the SoaExpire field if non-nil, zero value otherwise.

### GetSoaExpireOk

`func (o *ZoneRp) GetSoaExpireOk() (*int64, bool)`

GetSoaExpireOk returns a tuple with the SoaExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaExpire

`func (o *ZoneRp) SetSoaExpire(v int64)`

SetSoaExpire sets SoaExpire field to given value.

### HasSoaExpire

`func (o *ZoneRp) HasSoaExpire() bool`

HasSoaExpire returns a boolean if a field has been set.

### GetSoaNegativeTtl

`func (o *ZoneRp) GetSoaNegativeTtl() int64`

GetSoaNegativeTtl returns the SoaNegativeTtl field if non-nil, zero value otherwise.

### GetSoaNegativeTtlOk

`func (o *ZoneRp) GetSoaNegativeTtlOk() (*int64, bool)`

GetSoaNegativeTtlOk returns a tuple with the SoaNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaNegativeTtl

`func (o *ZoneRp) SetSoaNegativeTtl(v int64)`

SetSoaNegativeTtl sets SoaNegativeTtl field to given value.

### HasSoaNegativeTtl

`func (o *ZoneRp) HasSoaNegativeTtl() bool`

HasSoaNegativeTtl returns a boolean if a field has been set.

### GetSoaRefresh

`func (o *ZoneRp) GetSoaRefresh() int64`

GetSoaRefresh returns the SoaRefresh field if non-nil, zero value otherwise.

### GetSoaRefreshOk

`func (o *ZoneRp) GetSoaRefreshOk() (*int64, bool)`

GetSoaRefreshOk returns a tuple with the SoaRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRefresh

`func (o *ZoneRp) SetSoaRefresh(v int64)`

SetSoaRefresh sets SoaRefresh field to given value.

### HasSoaRefresh

`func (o *ZoneRp) HasSoaRefresh() bool`

HasSoaRefresh returns a boolean if a field has been set.

### GetSoaRetry

`func (o *ZoneRp) GetSoaRetry() int64`

GetSoaRetry returns the SoaRetry field if non-nil, zero value otherwise.

### GetSoaRetryOk

`func (o *ZoneRp) GetSoaRetryOk() (*int64, bool)`

GetSoaRetryOk returns a tuple with the SoaRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRetry

`func (o *ZoneRp) SetSoaRetry(v int64)`

SetSoaRetry sets SoaRetry field to given value.

### HasSoaRetry

`func (o *ZoneRp) HasSoaRetry() bool`

HasSoaRetry returns a boolean if a field has been set.

### GetSoaSerial

`func (o *ZoneRp) GetSoaSerial() int64`

GetSoaSerial returns the SoaSerial field if non-nil, zero value otherwise.

### GetSoaSerialOk

`func (o *ZoneRp) GetSoaSerialOk() (*int64, bool)`

GetSoaSerialOk returns a tuple with the SoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaSerial

`func (o *ZoneRp) SetSoaSerial(v int64)`

SetSoaSerial sets SoaSerial field to given value.

### HasSoaSerial

`func (o *ZoneRp) HasSoaSerial() bool`

HasSoaSerial returns a boolean if a field has been set.

### GetSubstituteName

`func (o *ZoneRp) GetSubstituteName() string`

GetSubstituteName returns the SubstituteName field if non-nil, zero value otherwise.

### GetSubstituteNameOk

`func (o *ZoneRp) GetSubstituteNameOk() (*string, bool)`

GetSubstituteNameOk returns a tuple with the SubstituteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstituteName

`func (o *ZoneRp) SetSubstituteName(v string)`

SetSubstituteName sets SubstituteName field to given value.

### HasSubstituteName

`func (o *ZoneRp) HasSubstituteName() bool`

HasSubstituteName returns a boolean if a field has been set.

### GetUseExternalPrimary

`func (o *ZoneRp) GetUseExternalPrimary() bool`

GetUseExternalPrimary returns the UseExternalPrimary field if non-nil, zero value otherwise.

### GetUseExternalPrimaryOk

`func (o *ZoneRp) GetUseExternalPrimaryOk() (*bool, bool)`

GetUseExternalPrimaryOk returns a tuple with the UseExternalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalPrimary

`func (o *ZoneRp) SetUseExternalPrimary(v bool)`

SetUseExternalPrimary sets UseExternalPrimary field to given value.

### HasUseExternalPrimary

`func (o *ZoneRp) HasUseExternalPrimary() bool`

HasUseExternalPrimary returns a boolean if a field has been set.

### GetUseGridZoneTimer

`func (o *ZoneRp) GetUseGridZoneTimer() bool`

GetUseGridZoneTimer returns the UseGridZoneTimer field if non-nil, zero value otherwise.

### GetUseGridZoneTimerOk

`func (o *ZoneRp) GetUseGridZoneTimerOk() (*bool, bool)`

GetUseGridZoneTimerOk returns a tuple with the UseGridZoneTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGridZoneTimer

`func (o *ZoneRp) SetUseGridZoneTimer(v bool)`

SetUseGridZoneTimer sets UseGridZoneTimer field to given value.

### HasUseGridZoneTimer

`func (o *ZoneRp) HasUseGridZoneTimer() bool`

HasUseGridZoneTimer returns a boolean if a field has been set.

### GetUseLogRpz

`func (o *ZoneRp) GetUseLogRpz() bool`

GetUseLogRpz returns the UseLogRpz field if non-nil, zero value otherwise.

### GetUseLogRpzOk

`func (o *ZoneRp) GetUseLogRpzOk() (*bool, bool)`

GetUseLogRpzOk returns a tuple with the UseLogRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogRpz

`func (o *ZoneRp) SetUseLogRpz(v bool)`

SetUseLogRpz sets UseLogRpz field to given value.

### HasUseLogRpz

`func (o *ZoneRp) HasUseLogRpz() bool`

HasUseLogRpz returns a boolean if a field has been set.

### GetUseRecordNamePolicy

`func (o *ZoneRp) GetUseRecordNamePolicy() bool`

GetUseRecordNamePolicy returns the UseRecordNamePolicy field if non-nil, zero value otherwise.

### GetUseRecordNamePolicyOk

`func (o *ZoneRp) GetUseRecordNamePolicyOk() (*bool, bool)`

GetUseRecordNamePolicyOk returns a tuple with the UseRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecordNamePolicy

`func (o *ZoneRp) SetUseRecordNamePolicy(v bool)`

SetUseRecordNamePolicy sets UseRecordNamePolicy field to given value.

### HasUseRecordNamePolicy

`func (o *ZoneRp) HasUseRecordNamePolicy() bool`

HasUseRecordNamePolicy returns a boolean if a field has been set.

### GetUseRpzDropIpRule

`func (o *ZoneRp) GetUseRpzDropIpRule() bool`

GetUseRpzDropIpRule returns the UseRpzDropIpRule field if non-nil, zero value otherwise.

### GetUseRpzDropIpRuleOk

`func (o *ZoneRp) GetUseRpzDropIpRuleOk() (*bool, bool)`

GetUseRpzDropIpRuleOk returns a tuple with the UseRpzDropIpRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzDropIpRule

`func (o *ZoneRp) SetUseRpzDropIpRule(v bool)`

SetUseRpzDropIpRule sets UseRpzDropIpRule field to given value.

### HasUseRpzDropIpRule

`func (o *ZoneRp) HasUseRpzDropIpRule() bool`

HasUseRpzDropIpRule returns a boolean if a field has been set.

### GetUseSoaEmail

`func (o *ZoneRp) GetUseSoaEmail() bool`

GetUseSoaEmail returns the UseSoaEmail field if non-nil, zero value otherwise.

### GetUseSoaEmailOk

`func (o *ZoneRp) GetUseSoaEmailOk() (*bool, bool)`

GetUseSoaEmailOk returns a tuple with the UseSoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSoaEmail

`func (o *ZoneRp) SetUseSoaEmail(v bool)`

SetUseSoaEmail sets UseSoaEmail field to given value.

### HasUseSoaEmail

`func (o *ZoneRp) HasUseSoaEmail() bool`

HasUseSoaEmail returns a boolean if a field has been set.

### GetUuid

`func (o *ZoneRp) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ZoneRp) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ZoneRp) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ZoneRp) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetView

`func (o *ZoneRp) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ZoneRp) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ZoneRp) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ZoneRp) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


