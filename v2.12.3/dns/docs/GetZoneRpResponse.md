# GetZoneRpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The IP address of the server that is serving this zone. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the zone; maximum 256 characters. | [optional] 
**CopyRpzRecords** | Pointer to **map[string]interface{}** |  | [optional] 
**Disable** | Pointer to **bool** | Determines whether a zone is disabled or not. When this is set to False, the zone is enabled. | [optional] 
**DisplayDomain** | Pointer to **string** | The displayed name of the DNS zone. | [optional] [readonly] 
**DnsSoaEmail** | Pointer to **string** | The SOA email for the zone in punycode format. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ZoneRpExternalPrimaries**](ZoneRpExternalPrimaries.md) | The list of external primary servers. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ZoneRpExternalSecondaries**](ZoneRpExternalSecondaries.md) | The list of external secondary servers. | [optional] 
**FireeyeRuleMapping** | Pointer to [**ZoneRpFireeyeRuleMapping**](ZoneRpFireeyeRuleMapping.md) |  | [optional] 
**Fqdn** | Pointer to **string** | The name of this DNS zone in FQDN format. | [optional] 
**GridPrimary** | Pointer to [**[]ZoneRpGridPrimary**](ZoneRpGridPrimary.md) | The grid primary servers for this zone. | [optional] 
**GridSecondaries** | Pointer to [**[]ZoneRpGridSecondaries**](ZoneRpGridSecondaries.md) | The list with Grid members that are secondary servers for this zone. | [optional] 
**LockUnlockZone** | Pointer to **map[string]interface{}** |  | [optional] 
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
**View** | Pointer to **string** | The name of the DNS view in which the zone resides. Example \&quot;external\&quot;. | [optional] 
**Result** | Pointer to [**ZoneRp**](ZoneRp.md) |  | [optional] 

## Methods

### NewGetZoneRpResponse

`func NewGetZoneRpResponse() *GetZoneRpResponse`

NewGetZoneRpResponse instantiates a new GetZoneRpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneRpResponseWithDefaults

`func NewGetZoneRpResponseWithDefaults() *GetZoneRpResponse`

NewGetZoneRpResponseWithDefaults instantiates a new GetZoneRpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetZoneRpResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetZoneRpResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetZoneRpResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetZoneRpResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetZoneRpResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetZoneRpResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetZoneRpResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetZoneRpResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *GetZoneRpResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetZoneRpResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetZoneRpResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetZoneRpResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyRpzRecords

`func (o *GetZoneRpResponse) GetCopyRpzRecords() map[string]interface{}`

GetCopyRpzRecords returns the CopyRpzRecords field if non-nil, zero value otherwise.

### GetCopyRpzRecordsOk

`func (o *GetZoneRpResponse) GetCopyRpzRecordsOk() (*map[string]interface{}, bool)`

GetCopyRpzRecordsOk returns a tuple with the CopyRpzRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRpzRecords

`func (o *GetZoneRpResponse) SetCopyRpzRecords(v map[string]interface{})`

SetCopyRpzRecords sets CopyRpzRecords field to given value.

### HasCopyRpzRecords

`func (o *GetZoneRpResponse) HasCopyRpzRecords() bool`

HasCopyRpzRecords returns a boolean if a field has been set.

### GetDisable

`func (o *GetZoneRpResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetZoneRpResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetZoneRpResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetZoneRpResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisplayDomain

`func (o *GetZoneRpResponse) GetDisplayDomain() string`

GetDisplayDomain returns the DisplayDomain field if non-nil, zero value otherwise.

### GetDisplayDomainOk

`func (o *GetZoneRpResponse) GetDisplayDomainOk() (*string, bool)`

GetDisplayDomainOk returns a tuple with the DisplayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDomain

`func (o *GetZoneRpResponse) SetDisplayDomain(v string)`

SetDisplayDomain sets DisplayDomain field to given value.

### HasDisplayDomain

`func (o *GetZoneRpResponse) HasDisplayDomain() bool`

HasDisplayDomain returns a boolean if a field has been set.

### GetDnsSoaEmail

`func (o *GetZoneRpResponse) GetDnsSoaEmail() string`

GetDnsSoaEmail returns the DnsSoaEmail field if non-nil, zero value otherwise.

### GetDnsSoaEmailOk

`func (o *GetZoneRpResponse) GetDnsSoaEmailOk() (*string, bool)`

GetDnsSoaEmailOk returns a tuple with the DnsSoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSoaEmail

`func (o *GetZoneRpResponse) SetDnsSoaEmail(v string)`

SetDnsSoaEmail sets DnsSoaEmail field to given value.

### HasDnsSoaEmail

`func (o *GetZoneRpResponse) HasDnsSoaEmail() bool`

HasDnsSoaEmail returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetZoneRpResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetZoneRpResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetZoneRpResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetZoneRpResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *GetZoneRpResponse) GetExternalPrimaries() []ZoneRpExternalPrimaries`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *GetZoneRpResponse) GetExternalPrimariesOk() (*[]ZoneRpExternalPrimaries, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *GetZoneRpResponse) SetExternalPrimaries(v []ZoneRpExternalPrimaries)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *GetZoneRpResponse) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *GetZoneRpResponse) GetExternalSecondaries() []ZoneRpExternalSecondaries`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *GetZoneRpResponse) GetExternalSecondariesOk() (*[]ZoneRpExternalSecondaries, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *GetZoneRpResponse) SetExternalSecondaries(v []ZoneRpExternalSecondaries)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *GetZoneRpResponse) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetFireeyeRuleMapping

`func (o *GetZoneRpResponse) GetFireeyeRuleMapping() ZoneRpFireeyeRuleMapping`

GetFireeyeRuleMapping returns the FireeyeRuleMapping field if non-nil, zero value otherwise.

### GetFireeyeRuleMappingOk

`func (o *GetZoneRpResponse) GetFireeyeRuleMappingOk() (*ZoneRpFireeyeRuleMapping, bool)`

GetFireeyeRuleMappingOk returns a tuple with the FireeyeRuleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireeyeRuleMapping

`func (o *GetZoneRpResponse) SetFireeyeRuleMapping(v ZoneRpFireeyeRuleMapping)`

SetFireeyeRuleMapping sets FireeyeRuleMapping field to given value.

### HasFireeyeRuleMapping

`func (o *GetZoneRpResponse) HasFireeyeRuleMapping() bool`

HasFireeyeRuleMapping returns a boolean if a field has been set.

### GetFqdn

`func (o *GetZoneRpResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GetZoneRpResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GetZoneRpResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GetZoneRpResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGridPrimary

`func (o *GetZoneRpResponse) GetGridPrimary() []ZoneRpGridPrimary`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *GetZoneRpResponse) GetGridPrimaryOk() (*[]ZoneRpGridPrimary, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *GetZoneRpResponse) SetGridPrimary(v []ZoneRpGridPrimary)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *GetZoneRpResponse) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *GetZoneRpResponse) GetGridSecondaries() []ZoneRpGridSecondaries`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *GetZoneRpResponse) GetGridSecondariesOk() (*[]ZoneRpGridSecondaries, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *GetZoneRpResponse) SetGridSecondaries(v []ZoneRpGridSecondaries)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *GetZoneRpResponse) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetLockUnlockZone

`func (o *GetZoneRpResponse) GetLockUnlockZone() map[string]interface{}`

GetLockUnlockZone returns the LockUnlockZone field if non-nil, zero value otherwise.

### GetLockUnlockZoneOk

`func (o *GetZoneRpResponse) GetLockUnlockZoneOk() (*map[string]interface{}, bool)`

GetLockUnlockZoneOk returns a tuple with the LockUnlockZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUnlockZone

`func (o *GetZoneRpResponse) SetLockUnlockZone(v map[string]interface{})`

SetLockUnlockZone sets LockUnlockZone field to given value.

### HasLockUnlockZone

`func (o *GetZoneRpResponse) HasLockUnlockZone() bool`

HasLockUnlockZone returns a boolean if a field has been set.

### GetLocked

`func (o *GetZoneRpResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetZoneRpResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetZoneRpResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetZoneRpResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *GetZoneRpResponse) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *GetZoneRpResponse) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *GetZoneRpResponse) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *GetZoneRpResponse) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetLogRpz

`func (o *GetZoneRpResponse) GetLogRpz() bool`

GetLogRpz returns the LogRpz field if non-nil, zero value otherwise.

### GetLogRpzOk

`func (o *GetZoneRpResponse) GetLogRpzOk() (*bool, bool)`

GetLogRpzOk returns a tuple with the LogRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRpz

`func (o *GetZoneRpResponse) SetLogRpz(v bool)`

SetLogRpz sets LogRpz field to given value.

### HasLogRpz

`func (o *GetZoneRpResponse) HasLogRpz() bool`

HasLogRpz returns a boolean if a field has been set.

### GetMaskPrefix

`func (o *GetZoneRpResponse) GetMaskPrefix() string`

GetMaskPrefix returns the MaskPrefix field if non-nil, zero value otherwise.

### GetMaskPrefixOk

`func (o *GetZoneRpResponse) GetMaskPrefixOk() (*string, bool)`

GetMaskPrefixOk returns a tuple with the MaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPrefix

`func (o *GetZoneRpResponse) SetMaskPrefix(v string)`

SetMaskPrefix sets MaskPrefix field to given value.

### HasMaskPrefix

`func (o *GetZoneRpResponse) HasMaskPrefix() bool`

HasMaskPrefix returns a boolean if a field has been set.

### GetMemberSoaMnames

`func (o *GetZoneRpResponse) GetMemberSoaMnames() []ZoneRpMemberSoaMnames`

GetMemberSoaMnames returns the MemberSoaMnames field if non-nil, zero value otherwise.

### GetMemberSoaMnamesOk

`func (o *GetZoneRpResponse) GetMemberSoaMnamesOk() (*[]ZoneRpMemberSoaMnames, bool)`

GetMemberSoaMnamesOk returns a tuple with the MemberSoaMnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSoaMnames

`func (o *GetZoneRpResponse) SetMemberSoaMnames(v []ZoneRpMemberSoaMnames)`

SetMemberSoaMnames sets MemberSoaMnames field to given value.

### HasMemberSoaMnames

`func (o *GetZoneRpResponse) HasMemberSoaMnames() bool`

HasMemberSoaMnames returns a boolean if a field has been set.

### GetMemberSoaSerials

`func (o *GetZoneRpResponse) GetMemberSoaSerials() []ZoneRpMemberSoaSerials`

GetMemberSoaSerials returns the MemberSoaSerials field if non-nil, zero value otherwise.

### GetMemberSoaSerialsOk

`func (o *GetZoneRpResponse) GetMemberSoaSerialsOk() (*[]ZoneRpMemberSoaSerials, bool)`

GetMemberSoaSerialsOk returns a tuple with the MemberSoaSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSoaSerials

`func (o *GetZoneRpResponse) SetMemberSoaSerials(v []ZoneRpMemberSoaSerials)`

SetMemberSoaSerials sets MemberSoaSerials field to given value.

### HasMemberSoaSerials

`func (o *GetZoneRpResponse) HasMemberSoaSerials() bool`

HasMemberSoaSerials returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetZoneRpResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetZoneRpResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetZoneRpResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetZoneRpResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNsGroup

`func (o *GetZoneRpResponse) GetNsGroup() string`

GetNsGroup returns the NsGroup field if non-nil, zero value otherwise.

### GetNsGroupOk

`func (o *GetZoneRpResponse) GetNsGroupOk() (*string, bool)`

GetNsGroupOk returns a tuple with the NsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsGroup

`func (o *GetZoneRpResponse) SetNsGroup(v string)`

SetNsGroup sets NsGroup field to given value.

### HasNsGroup

`func (o *GetZoneRpResponse) HasNsGroup() bool`

HasNsGroup returns a boolean if a field has been set.

### GetParent

`func (o *GetZoneRpResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetZoneRpResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetZoneRpResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetZoneRpResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrefix

`func (o *GetZoneRpResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetZoneRpResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetZoneRpResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetZoneRpResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrimaryType

`func (o *GetZoneRpResponse) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *GetZoneRpResponse) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *GetZoneRpResponse) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *GetZoneRpResponse) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetRecordNamePolicy

`func (o *GetZoneRpResponse) GetRecordNamePolicy() string`

GetRecordNamePolicy returns the RecordNamePolicy field if non-nil, zero value otherwise.

### GetRecordNamePolicyOk

`func (o *GetZoneRpResponse) GetRecordNamePolicyOk() (*string, bool)`

GetRecordNamePolicyOk returns a tuple with the RecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNamePolicy

`func (o *GetZoneRpResponse) SetRecordNamePolicy(v string)`

SetRecordNamePolicy sets RecordNamePolicy field to given value.

### HasRecordNamePolicy

`func (o *GetZoneRpResponse) HasRecordNamePolicy() bool`

HasRecordNamePolicy returns a boolean if a field has been set.

### GetRpzDropIpRuleEnabled

`func (o *GetZoneRpResponse) GetRpzDropIpRuleEnabled() bool`

GetRpzDropIpRuleEnabled returns the RpzDropIpRuleEnabled field if non-nil, zero value otherwise.

### GetRpzDropIpRuleEnabledOk

`func (o *GetZoneRpResponse) GetRpzDropIpRuleEnabledOk() (*bool, bool)`

GetRpzDropIpRuleEnabledOk returns a tuple with the RpzDropIpRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleEnabled

`func (o *GetZoneRpResponse) SetRpzDropIpRuleEnabled(v bool)`

SetRpzDropIpRuleEnabled sets RpzDropIpRuleEnabled field to given value.

### HasRpzDropIpRuleEnabled

`func (o *GetZoneRpResponse) HasRpzDropIpRuleEnabled() bool`

HasRpzDropIpRuleEnabled returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetZoneRpResponse) GetRpzDropIpRuleMinPrefixLengthIpv4() int64`

GetRpzDropIpRuleMinPrefixLengthIpv4 returns the RpzDropIpRuleMinPrefixLengthIpv4 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv4Ok

`func (o *GetZoneRpResponse) GetRpzDropIpRuleMinPrefixLengthIpv4Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv4Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetZoneRpResponse) SetRpzDropIpRuleMinPrefixLengthIpv4(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv4 sets RpzDropIpRuleMinPrefixLengthIpv4 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv4

`func (o *GetZoneRpResponse) HasRpzDropIpRuleMinPrefixLengthIpv4() bool`

HasRpzDropIpRuleMinPrefixLengthIpv4 returns a boolean if a field has been set.

### GetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetZoneRpResponse) GetRpzDropIpRuleMinPrefixLengthIpv6() int64`

GetRpzDropIpRuleMinPrefixLengthIpv6 returns the RpzDropIpRuleMinPrefixLengthIpv6 field if non-nil, zero value otherwise.

### GetRpzDropIpRuleMinPrefixLengthIpv6Ok

`func (o *GetZoneRpResponse) GetRpzDropIpRuleMinPrefixLengthIpv6Ok() (*int64, bool)`

GetRpzDropIpRuleMinPrefixLengthIpv6Ok returns a tuple with the RpzDropIpRuleMinPrefixLengthIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetZoneRpResponse) SetRpzDropIpRuleMinPrefixLengthIpv6(v int64)`

SetRpzDropIpRuleMinPrefixLengthIpv6 sets RpzDropIpRuleMinPrefixLengthIpv6 field to given value.

### HasRpzDropIpRuleMinPrefixLengthIpv6

`func (o *GetZoneRpResponse) HasRpzDropIpRuleMinPrefixLengthIpv6() bool`

HasRpzDropIpRuleMinPrefixLengthIpv6 returns a boolean if a field has been set.

### GetRpzLastUpdatedTime

`func (o *GetZoneRpResponse) GetRpzLastUpdatedTime() int64`

GetRpzLastUpdatedTime returns the RpzLastUpdatedTime field if non-nil, zero value otherwise.

### GetRpzLastUpdatedTimeOk

`func (o *GetZoneRpResponse) GetRpzLastUpdatedTimeOk() (*int64, bool)`

GetRpzLastUpdatedTimeOk returns a tuple with the RpzLastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzLastUpdatedTime

`func (o *GetZoneRpResponse) SetRpzLastUpdatedTime(v int64)`

SetRpzLastUpdatedTime sets RpzLastUpdatedTime field to given value.

### HasRpzLastUpdatedTime

`func (o *GetZoneRpResponse) HasRpzLastUpdatedTime() bool`

HasRpzLastUpdatedTime returns a boolean if a field has been set.

### GetRpzPolicy

`func (o *GetZoneRpResponse) GetRpzPolicy() string`

GetRpzPolicy returns the RpzPolicy field if non-nil, zero value otherwise.

### GetRpzPolicyOk

`func (o *GetZoneRpResponse) GetRpzPolicyOk() (*string, bool)`

GetRpzPolicyOk returns a tuple with the RpzPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPolicy

`func (o *GetZoneRpResponse) SetRpzPolicy(v string)`

SetRpzPolicy sets RpzPolicy field to given value.

### HasRpzPolicy

`func (o *GetZoneRpResponse) HasRpzPolicy() bool`

HasRpzPolicy returns a boolean if a field has been set.

### GetRpzPriority

`func (o *GetZoneRpResponse) GetRpzPriority() int64`

GetRpzPriority returns the RpzPriority field if non-nil, zero value otherwise.

### GetRpzPriorityOk

`func (o *GetZoneRpResponse) GetRpzPriorityOk() (*int64, bool)`

GetRpzPriorityOk returns a tuple with the RpzPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPriority

`func (o *GetZoneRpResponse) SetRpzPriority(v int64)`

SetRpzPriority sets RpzPriority field to given value.

### HasRpzPriority

`func (o *GetZoneRpResponse) HasRpzPriority() bool`

HasRpzPriority returns a boolean if a field has been set.

### GetRpzPriorityEnd

`func (o *GetZoneRpResponse) GetRpzPriorityEnd() int64`

GetRpzPriorityEnd returns the RpzPriorityEnd field if non-nil, zero value otherwise.

### GetRpzPriorityEndOk

`func (o *GetZoneRpResponse) GetRpzPriorityEndOk() (*int64, bool)`

GetRpzPriorityEndOk returns a tuple with the RpzPriorityEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPriorityEnd

`func (o *GetZoneRpResponse) SetRpzPriorityEnd(v int64)`

SetRpzPriorityEnd sets RpzPriorityEnd field to given value.

### HasRpzPriorityEnd

`func (o *GetZoneRpResponse) HasRpzPriorityEnd() bool`

HasRpzPriorityEnd returns a boolean if a field has been set.

### GetRpzSeverity

`func (o *GetZoneRpResponse) GetRpzSeverity() string`

GetRpzSeverity returns the RpzSeverity field if non-nil, zero value otherwise.

### GetRpzSeverityOk

`func (o *GetZoneRpResponse) GetRpzSeverityOk() (*string, bool)`

GetRpzSeverityOk returns a tuple with the RpzSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzSeverity

`func (o *GetZoneRpResponse) SetRpzSeverity(v string)`

SetRpzSeverity sets RpzSeverity field to given value.

### HasRpzSeverity

`func (o *GetZoneRpResponse) HasRpzSeverity() bool`

HasRpzSeverity returns a boolean if a field has been set.

### GetRpzType

`func (o *GetZoneRpResponse) GetRpzType() string`

GetRpzType returns the RpzType field if non-nil, zero value otherwise.

### GetRpzTypeOk

`func (o *GetZoneRpResponse) GetRpzTypeOk() (*string, bool)`

GetRpzTypeOk returns a tuple with the RpzType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzType

`func (o *GetZoneRpResponse) SetRpzType(v string)`

SetRpzType sets RpzType field to given value.

### HasRpzType

`func (o *GetZoneRpResponse) HasRpzType() bool`

HasRpzType returns a boolean if a field has been set.

### GetSetSoaSerialNumber

`func (o *GetZoneRpResponse) GetSetSoaSerialNumber() bool`

GetSetSoaSerialNumber returns the SetSoaSerialNumber field if non-nil, zero value otherwise.

### GetSetSoaSerialNumberOk

`func (o *GetZoneRpResponse) GetSetSoaSerialNumberOk() (*bool, bool)`

GetSetSoaSerialNumberOk returns a tuple with the SetSoaSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSoaSerialNumber

`func (o *GetZoneRpResponse) SetSetSoaSerialNumber(v bool)`

SetSetSoaSerialNumber sets SetSoaSerialNumber field to given value.

### HasSetSoaSerialNumber

`func (o *GetZoneRpResponse) HasSetSoaSerialNumber() bool`

HasSetSoaSerialNumber returns a boolean if a field has been set.

### GetSoaDefaultTtl

`func (o *GetZoneRpResponse) GetSoaDefaultTtl() int64`

GetSoaDefaultTtl returns the SoaDefaultTtl field if non-nil, zero value otherwise.

### GetSoaDefaultTtlOk

`func (o *GetZoneRpResponse) GetSoaDefaultTtlOk() (*int64, bool)`

GetSoaDefaultTtlOk returns a tuple with the SoaDefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaDefaultTtl

`func (o *GetZoneRpResponse) SetSoaDefaultTtl(v int64)`

SetSoaDefaultTtl sets SoaDefaultTtl field to given value.

### HasSoaDefaultTtl

`func (o *GetZoneRpResponse) HasSoaDefaultTtl() bool`

HasSoaDefaultTtl returns a boolean if a field has been set.

### GetSoaEmail

`func (o *GetZoneRpResponse) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *GetZoneRpResponse) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *GetZoneRpResponse) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *GetZoneRpResponse) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetSoaExpire

`func (o *GetZoneRpResponse) GetSoaExpire() int64`

GetSoaExpire returns the SoaExpire field if non-nil, zero value otherwise.

### GetSoaExpireOk

`func (o *GetZoneRpResponse) GetSoaExpireOk() (*int64, bool)`

GetSoaExpireOk returns a tuple with the SoaExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaExpire

`func (o *GetZoneRpResponse) SetSoaExpire(v int64)`

SetSoaExpire sets SoaExpire field to given value.

### HasSoaExpire

`func (o *GetZoneRpResponse) HasSoaExpire() bool`

HasSoaExpire returns a boolean if a field has been set.

### GetSoaNegativeTtl

`func (o *GetZoneRpResponse) GetSoaNegativeTtl() int64`

GetSoaNegativeTtl returns the SoaNegativeTtl field if non-nil, zero value otherwise.

### GetSoaNegativeTtlOk

`func (o *GetZoneRpResponse) GetSoaNegativeTtlOk() (*int64, bool)`

GetSoaNegativeTtlOk returns a tuple with the SoaNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaNegativeTtl

`func (o *GetZoneRpResponse) SetSoaNegativeTtl(v int64)`

SetSoaNegativeTtl sets SoaNegativeTtl field to given value.

### HasSoaNegativeTtl

`func (o *GetZoneRpResponse) HasSoaNegativeTtl() bool`

HasSoaNegativeTtl returns a boolean if a field has been set.

### GetSoaRefresh

`func (o *GetZoneRpResponse) GetSoaRefresh() int64`

GetSoaRefresh returns the SoaRefresh field if non-nil, zero value otherwise.

### GetSoaRefreshOk

`func (o *GetZoneRpResponse) GetSoaRefreshOk() (*int64, bool)`

GetSoaRefreshOk returns a tuple with the SoaRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRefresh

`func (o *GetZoneRpResponse) SetSoaRefresh(v int64)`

SetSoaRefresh sets SoaRefresh field to given value.

### HasSoaRefresh

`func (o *GetZoneRpResponse) HasSoaRefresh() bool`

HasSoaRefresh returns a boolean if a field has been set.

### GetSoaRetry

`func (o *GetZoneRpResponse) GetSoaRetry() int64`

GetSoaRetry returns the SoaRetry field if non-nil, zero value otherwise.

### GetSoaRetryOk

`func (o *GetZoneRpResponse) GetSoaRetryOk() (*int64, bool)`

GetSoaRetryOk returns a tuple with the SoaRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRetry

`func (o *GetZoneRpResponse) SetSoaRetry(v int64)`

SetSoaRetry sets SoaRetry field to given value.

### HasSoaRetry

`func (o *GetZoneRpResponse) HasSoaRetry() bool`

HasSoaRetry returns a boolean if a field has been set.

### GetSoaSerial

`func (o *GetZoneRpResponse) GetSoaSerial() int64`

GetSoaSerial returns the SoaSerial field if non-nil, zero value otherwise.

### GetSoaSerialOk

`func (o *GetZoneRpResponse) GetSoaSerialOk() (*int64, bool)`

GetSoaSerialOk returns a tuple with the SoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaSerial

`func (o *GetZoneRpResponse) SetSoaSerial(v int64)`

SetSoaSerial sets SoaSerial field to given value.

### HasSoaSerial

`func (o *GetZoneRpResponse) HasSoaSerial() bool`

HasSoaSerial returns a boolean if a field has been set.

### GetSubstituteName

`func (o *GetZoneRpResponse) GetSubstituteName() string`

GetSubstituteName returns the SubstituteName field if non-nil, zero value otherwise.

### GetSubstituteNameOk

`func (o *GetZoneRpResponse) GetSubstituteNameOk() (*string, bool)`

GetSubstituteNameOk returns a tuple with the SubstituteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstituteName

`func (o *GetZoneRpResponse) SetSubstituteName(v string)`

SetSubstituteName sets SubstituteName field to given value.

### HasSubstituteName

`func (o *GetZoneRpResponse) HasSubstituteName() bool`

HasSubstituteName returns a boolean if a field has been set.

### GetUseExternalPrimary

`func (o *GetZoneRpResponse) GetUseExternalPrimary() bool`

GetUseExternalPrimary returns the UseExternalPrimary field if non-nil, zero value otherwise.

### GetUseExternalPrimaryOk

`func (o *GetZoneRpResponse) GetUseExternalPrimaryOk() (*bool, bool)`

GetUseExternalPrimaryOk returns a tuple with the UseExternalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalPrimary

`func (o *GetZoneRpResponse) SetUseExternalPrimary(v bool)`

SetUseExternalPrimary sets UseExternalPrimary field to given value.

### HasUseExternalPrimary

`func (o *GetZoneRpResponse) HasUseExternalPrimary() bool`

HasUseExternalPrimary returns a boolean if a field has been set.

### GetUseGridZoneTimer

`func (o *GetZoneRpResponse) GetUseGridZoneTimer() bool`

GetUseGridZoneTimer returns the UseGridZoneTimer field if non-nil, zero value otherwise.

### GetUseGridZoneTimerOk

`func (o *GetZoneRpResponse) GetUseGridZoneTimerOk() (*bool, bool)`

GetUseGridZoneTimerOk returns a tuple with the UseGridZoneTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGridZoneTimer

`func (o *GetZoneRpResponse) SetUseGridZoneTimer(v bool)`

SetUseGridZoneTimer sets UseGridZoneTimer field to given value.

### HasUseGridZoneTimer

`func (o *GetZoneRpResponse) HasUseGridZoneTimer() bool`

HasUseGridZoneTimer returns a boolean if a field has been set.

### GetUseLogRpz

`func (o *GetZoneRpResponse) GetUseLogRpz() bool`

GetUseLogRpz returns the UseLogRpz field if non-nil, zero value otherwise.

### GetUseLogRpzOk

`func (o *GetZoneRpResponse) GetUseLogRpzOk() (*bool, bool)`

GetUseLogRpzOk returns a tuple with the UseLogRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogRpz

`func (o *GetZoneRpResponse) SetUseLogRpz(v bool)`

SetUseLogRpz sets UseLogRpz field to given value.

### HasUseLogRpz

`func (o *GetZoneRpResponse) HasUseLogRpz() bool`

HasUseLogRpz returns a boolean if a field has been set.

### GetUseRecordNamePolicy

`func (o *GetZoneRpResponse) GetUseRecordNamePolicy() bool`

GetUseRecordNamePolicy returns the UseRecordNamePolicy field if non-nil, zero value otherwise.

### GetUseRecordNamePolicyOk

`func (o *GetZoneRpResponse) GetUseRecordNamePolicyOk() (*bool, bool)`

GetUseRecordNamePolicyOk returns a tuple with the UseRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecordNamePolicy

`func (o *GetZoneRpResponse) SetUseRecordNamePolicy(v bool)`

SetUseRecordNamePolicy sets UseRecordNamePolicy field to given value.

### HasUseRecordNamePolicy

`func (o *GetZoneRpResponse) HasUseRecordNamePolicy() bool`

HasUseRecordNamePolicy returns a boolean if a field has been set.

### GetUseRpzDropIpRule

`func (o *GetZoneRpResponse) GetUseRpzDropIpRule() bool`

GetUseRpzDropIpRule returns the UseRpzDropIpRule field if non-nil, zero value otherwise.

### GetUseRpzDropIpRuleOk

`func (o *GetZoneRpResponse) GetUseRpzDropIpRuleOk() (*bool, bool)`

GetUseRpzDropIpRuleOk returns a tuple with the UseRpzDropIpRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRpzDropIpRule

`func (o *GetZoneRpResponse) SetUseRpzDropIpRule(v bool)`

SetUseRpzDropIpRule sets UseRpzDropIpRule field to given value.

### HasUseRpzDropIpRule

`func (o *GetZoneRpResponse) HasUseRpzDropIpRule() bool`

HasUseRpzDropIpRule returns a boolean if a field has been set.

### GetUseSoaEmail

`func (o *GetZoneRpResponse) GetUseSoaEmail() bool`

GetUseSoaEmail returns the UseSoaEmail field if non-nil, zero value otherwise.

### GetUseSoaEmailOk

`func (o *GetZoneRpResponse) GetUseSoaEmailOk() (*bool, bool)`

GetUseSoaEmailOk returns a tuple with the UseSoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSoaEmail

`func (o *GetZoneRpResponse) SetUseSoaEmail(v bool)`

SetUseSoaEmail sets UseSoaEmail field to given value.

### HasUseSoaEmail

`func (o *GetZoneRpResponse) HasUseSoaEmail() bool`

HasUseSoaEmail returns a boolean if a field has been set.

### GetView

`func (o *GetZoneRpResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetZoneRpResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetZoneRpResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetZoneRpResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetResult

`func (o *GetZoneRpResponse) GetResult() ZoneRp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneRpResponse) GetResultOk() (*ZoneRp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneRpResponse) SetResult(v ZoneRp)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetZoneRpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


