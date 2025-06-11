# GetZoneStubResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The IP address of the server that is serving this zone. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the zone; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a zone is disabled or not. When this is set to False, the zone is enabled. | [optional] 
**DisableForwarding** | Pointer to **bool** | Determines if the name servers that host the zone should not forward queries that end with the domain name of the zone to any configured forwarders. | [optional] 
**DisplayDomain** | Pointer to **string** | The displayed name of the DNS zone. | [optional] [readonly] 
**DnsFqdn** | Pointer to **string** | The name of this DNS zone in punycode format. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format in punycode format. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalNsGroup** | Pointer to **string** | A forward stub server name server group. | [optional] 
**Fqdn** | Pointer to **string** | The name of this DNS zone. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set. | [optional] 
**Locked** | Pointer to **bool** | If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked. | [optional] 
**LockedBy** | Pointer to **string** | The name of a superuser or the administrator who locked this zone. | [optional] [readonly] 
**MaskPrefix** | Pointer to **string** | IPv4 Netmask or IPv6 prefix for this zone. | [optional] [readonly] 
**MsAdIntegrated** | Pointer to **bool** | The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \&quot;STUB\&quot;, \&quot;AUTH_PRIMARY\&quot;, or \&quot;AUTH_BOTH\&quot;. | [optional] 
**MsDdnsMode** | Pointer to **string** | Determines whether an Active Directory-integrated zone with a Microsoft DNS server as primary allows dynamic updates. Valid values are: \&quot;SECURE\&quot; if the zone allows secure updates only. \&quot;NONE\&quot; if the zone forbids dynamic updates. \&quot;ANY\&quot; if the zone accepts both secure and nonsecure updates. This field is valid only if ms_managed is either \&quot;AUTH_PRIMARY\&quot; or \&quot;AUTH_BOTH\&quot;. If the flag ms_ad_integrated is false, the value \&quot;SECURE\&quot; is not allowed. | [optional] 
**MsManaged** | Pointer to **string** | The flag that indicates whether the zone is assigned to a Microsoft DNS server. This flag returns the authoritative name server type of the Microsoft DNS server. Valid values are: \&quot;NONE\&quot; if the zone is not assigned to any Microsoft DNS server. \&quot;STUB\&quot; if the zone is assigned to a Microsoft DNS server as a stub zone. \&quot;AUTH_PRIMARY\&quot; if only the primary server of the zone is a Microsoft DNS server. \&quot;AUTH_SECONDARY\&quot; if only the secondary server of the zone is a Microsoft DNS server. \&quot;AUTH_BOTH\&quot; if both the primary and secondary servers of the zone are Microsoft DNS servers. | [optional] [readonly] 
**MsReadOnly** | Pointer to **bool** | Determines if a Grid member manages the zone served by a Microsoft DNS server in read-only mode. This flag is true when a Grid member manages the zone in read-only mode, false otherwise. When the zone has the ms_read_only flag set to True, no changes can be made to this zone. | [optional] [readonly] 
**MsSyncMasterName** | Pointer to **string** | The name of MS synchronization master for this zone. | [optional] [readonly] 
**NsGroup** | Pointer to **string** | A stub member name server group. | [optional] 
**Parent** | Pointer to **string** | The parent zone of this zone. Note that when searching for reverse zones, the \&quot;in-addr.arpa\&quot; notation should be used. | [optional] [readonly] 
**Prefix** | Pointer to **string** | The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B. | [optional] 
**SoaEmail** | Pointer to **string** | The SOA email for the zone. This value can be in unicode format. | [optional] [readonly] 
**SoaExpire** | Pointer to **int64** | This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. | [optional] [readonly] 
**SoaMname** | Pointer to **string** | The SOA mname value for this zone. The Infoblox appliance allows you to change the name of the primary server on the SOA record that is automatically created when you initially configure a zone. Use this method to change the name of the primary server on the SOA record. For example, you may want to hide the primary server for a zone. If your device is named dns1.zone.tld, and for security reasons, you want to show a secondary server called dns2.zone.tld as the primary server. To do so, you would go to dns1.zone.tld zone (being the true primary) and change the primary server on the SOA to dns2.zone.tld to hide the true identity of the real primary server. This value can be in unicode format. | [optional] [readonly] 
**SoaNegativeTtl** | Pointer to **int64** | The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \&quot;Does Not Respond\&quot; responses. | [optional] [readonly] 
**SoaRefresh** | Pointer to **int64** | This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not. | [optional] [readonly] 
**SoaRetry** | Pointer to **int64** | This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs. | [optional] [readonly] 
**SoaSerial** | Pointer to **int64** | The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server. | [optional] [readonly] 
**StubFrom** | Pointer to [**[]ZoneStubStubFrom**](ZoneStubStubFrom.md) | The primary servers (masters) of this stub zone. | [optional] 
**StubMembers** | Pointer to [**[]ZoneStubStubMembers**](ZoneStubStubMembers.md) | The Grid member servers of this stub zone. Note that the lead/stealth/grid_replicate/ preferred_primaries/override_preferred_primaries fields of the struct will be ignored when set in this field. | [optional] 
**StubMsservers** | Pointer to [**[]ZoneStubStubMsservers**](ZoneStubStubMsservers.md) | The Microsoft DNS servers of this stub zone. Note that the stealth field of the struct will be ignored when set in this field. | [optional] 
**UsingSrgAssociations** | Pointer to **bool** | This is true if the zone is associated with a shared record group. | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the zone resides. Example \&quot;external\&quot;. | [optional] 
**ZoneFormat** | Pointer to **string** | Determines the format of this zone. | [optional] 
**Result** | Pointer to [**ZoneStub**](ZoneStub.md) |  | [optional] 

## Methods

### NewGetZoneStubResponse

`func NewGetZoneStubResponse() *GetZoneStubResponse`

NewGetZoneStubResponse instantiates a new GetZoneStubResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneStubResponseWithDefaults

`func NewGetZoneStubResponseWithDefaults() *GetZoneStubResponse`

NewGetZoneStubResponseWithDefaults instantiates a new GetZoneStubResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetZoneStubResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetZoneStubResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetZoneStubResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetZoneStubResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetZoneStubResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetZoneStubResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetZoneStubResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetZoneStubResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *GetZoneStubResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetZoneStubResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetZoneStubResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetZoneStubResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetZoneStubResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetZoneStubResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetZoneStubResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetZoneStubResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableForwarding

`func (o *GetZoneStubResponse) GetDisableForwarding() bool`

GetDisableForwarding returns the DisableForwarding field if non-nil, zero value otherwise.

### GetDisableForwardingOk

`func (o *GetZoneStubResponse) GetDisableForwardingOk() (*bool, bool)`

GetDisableForwardingOk returns a tuple with the DisableForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableForwarding

`func (o *GetZoneStubResponse) SetDisableForwarding(v bool)`

SetDisableForwarding sets DisableForwarding field to given value.

### HasDisableForwarding

`func (o *GetZoneStubResponse) HasDisableForwarding() bool`

HasDisableForwarding returns a boolean if a field has been set.

### GetDisplayDomain

`func (o *GetZoneStubResponse) GetDisplayDomain() string`

GetDisplayDomain returns the DisplayDomain field if non-nil, zero value otherwise.

### GetDisplayDomainOk

`func (o *GetZoneStubResponse) GetDisplayDomainOk() (*string, bool)`

GetDisplayDomainOk returns a tuple with the DisplayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDomain

`func (o *GetZoneStubResponse) SetDisplayDomain(v string)`

SetDisplayDomain sets DisplayDomain field to given value.

### HasDisplayDomain

`func (o *GetZoneStubResponse) HasDisplayDomain() bool`

HasDisplayDomain returns a boolean if a field has been set.

### GetDnsFqdn

`func (o *GetZoneStubResponse) GetDnsFqdn() string`

GetDnsFqdn returns the DnsFqdn field if non-nil, zero value otherwise.

### GetDnsFqdnOk

`func (o *GetZoneStubResponse) GetDnsFqdnOk() (*string, bool)`

GetDnsFqdnOk returns a tuple with the DnsFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFqdn

`func (o *GetZoneStubResponse) SetDnsFqdn(v string)`

SetDnsFqdn sets DnsFqdn field to given value.

### HasDnsFqdn

`func (o *GetZoneStubResponse) HasDnsFqdn() bool`

HasDnsFqdn returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetZoneStubResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetZoneStubResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetZoneStubResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetZoneStubResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetExternalNsGroup

`func (o *GetZoneStubResponse) GetExternalNsGroup() string`

GetExternalNsGroup returns the ExternalNsGroup field if non-nil, zero value otherwise.

### GetExternalNsGroupOk

`func (o *GetZoneStubResponse) GetExternalNsGroupOk() (*string, bool)`

GetExternalNsGroupOk returns a tuple with the ExternalNsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNsGroup

`func (o *GetZoneStubResponse) SetExternalNsGroup(v string)`

SetExternalNsGroup sets ExternalNsGroup field to given value.

### HasExternalNsGroup

`func (o *GetZoneStubResponse) HasExternalNsGroup() bool`

HasExternalNsGroup returns a boolean if a field has been set.

### GetFqdn

`func (o *GetZoneStubResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GetZoneStubResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GetZoneStubResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GetZoneStubResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetLocked

`func (o *GetZoneStubResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetZoneStubResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetZoneStubResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetZoneStubResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *GetZoneStubResponse) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *GetZoneStubResponse) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *GetZoneStubResponse) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *GetZoneStubResponse) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetMaskPrefix

`func (o *GetZoneStubResponse) GetMaskPrefix() string`

GetMaskPrefix returns the MaskPrefix field if non-nil, zero value otherwise.

### GetMaskPrefixOk

`func (o *GetZoneStubResponse) GetMaskPrefixOk() (*string, bool)`

GetMaskPrefixOk returns a tuple with the MaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPrefix

`func (o *GetZoneStubResponse) SetMaskPrefix(v string)`

SetMaskPrefix sets MaskPrefix field to given value.

### HasMaskPrefix

`func (o *GetZoneStubResponse) HasMaskPrefix() bool`

HasMaskPrefix returns a boolean if a field has been set.

### GetMsAdIntegrated

`func (o *GetZoneStubResponse) GetMsAdIntegrated() bool`

GetMsAdIntegrated returns the MsAdIntegrated field if non-nil, zero value otherwise.

### GetMsAdIntegratedOk

`func (o *GetZoneStubResponse) GetMsAdIntegratedOk() (*bool, bool)`

GetMsAdIntegratedOk returns a tuple with the MsAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdIntegrated

`func (o *GetZoneStubResponse) SetMsAdIntegrated(v bool)`

SetMsAdIntegrated sets MsAdIntegrated field to given value.

### HasMsAdIntegrated

`func (o *GetZoneStubResponse) HasMsAdIntegrated() bool`

HasMsAdIntegrated returns a boolean if a field has been set.

### GetMsDdnsMode

`func (o *GetZoneStubResponse) GetMsDdnsMode() string`

GetMsDdnsMode returns the MsDdnsMode field if non-nil, zero value otherwise.

### GetMsDdnsModeOk

`func (o *GetZoneStubResponse) GetMsDdnsModeOk() (*string, bool)`

GetMsDdnsModeOk returns a tuple with the MsDdnsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDdnsMode

`func (o *GetZoneStubResponse) SetMsDdnsMode(v string)`

SetMsDdnsMode sets MsDdnsMode field to given value.

### HasMsDdnsMode

`func (o *GetZoneStubResponse) HasMsDdnsMode() bool`

HasMsDdnsMode returns a boolean if a field has been set.

### GetMsManaged

`func (o *GetZoneStubResponse) GetMsManaged() string`

GetMsManaged returns the MsManaged field if non-nil, zero value otherwise.

### GetMsManagedOk

`func (o *GetZoneStubResponse) GetMsManagedOk() (*string, bool)`

GetMsManagedOk returns a tuple with the MsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsManaged

`func (o *GetZoneStubResponse) SetMsManaged(v string)`

SetMsManaged sets MsManaged field to given value.

### HasMsManaged

`func (o *GetZoneStubResponse) HasMsManaged() bool`

HasMsManaged returns a boolean if a field has been set.

### GetMsReadOnly

`func (o *GetZoneStubResponse) GetMsReadOnly() bool`

GetMsReadOnly returns the MsReadOnly field if non-nil, zero value otherwise.

### GetMsReadOnlyOk

`func (o *GetZoneStubResponse) GetMsReadOnlyOk() (*bool, bool)`

GetMsReadOnlyOk returns a tuple with the MsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsReadOnly

`func (o *GetZoneStubResponse) SetMsReadOnly(v bool)`

SetMsReadOnly sets MsReadOnly field to given value.

### HasMsReadOnly

`func (o *GetZoneStubResponse) HasMsReadOnly() bool`

HasMsReadOnly returns a boolean if a field has been set.

### GetMsSyncMasterName

`func (o *GetZoneStubResponse) GetMsSyncMasterName() string`

GetMsSyncMasterName returns the MsSyncMasterName field if non-nil, zero value otherwise.

### GetMsSyncMasterNameOk

`func (o *GetZoneStubResponse) GetMsSyncMasterNameOk() (*string, bool)`

GetMsSyncMasterNameOk returns a tuple with the MsSyncMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncMasterName

`func (o *GetZoneStubResponse) SetMsSyncMasterName(v string)`

SetMsSyncMasterName sets MsSyncMasterName field to given value.

### HasMsSyncMasterName

`func (o *GetZoneStubResponse) HasMsSyncMasterName() bool`

HasMsSyncMasterName returns a boolean if a field has been set.

### GetNsGroup

`func (o *GetZoneStubResponse) GetNsGroup() string`

GetNsGroup returns the NsGroup field if non-nil, zero value otherwise.

### GetNsGroupOk

`func (o *GetZoneStubResponse) GetNsGroupOk() (*string, bool)`

GetNsGroupOk returns a tuple with the NsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsGroup

`func (o *GetZoneStubResponse) SetNsGroup(v string)`

SetNsGroup sets NsGroup field to given value.

### HasNsGroup

`func (o *GetZoneStubResponse) HasNsGroup() bool`

HasNsGroup returns a boolean if a field has been set.

### GetParent

`func (o *GetZoneStubResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetZoneStubResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetZoneStubResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetZoneStubResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrefix

`func (o *GetZoneStubResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetZoneStubResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetZoneStubResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetZoneStubResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSoaEmail

`func (o *GetZoneStubResponse) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *GetZoneStubResponse) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *GetZoneStubResponse) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *GetZoneStubResponse) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetSoaExpire

`func (o *GetZoneStubResponse) GetSoaExpire() int64`

GetSoaExpire returns the SoaExpire field if non-nil, zero value otherwise.

### GetSoaExpireOk

`func (o *GetZoneStubResponse) GetSoaExpireOk() (*int64, bool)`

GetSoaExpireOk returns a tuple with the SoaExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaExpire

`func (o *GetZoneStubResponse) SetSoaExpire(v int64)`

SetSoaExpire sets SoaExpire field to given value.

### HasSoaExpire

`func (o *GetZoneStubResponse) HasSoaExpire() bool`

HasSoaExpire returns a boolean if a field has been set.

### GetSoaMname

`func (o *GetZoneStubResponse) GetSoaMname() string`

GetSoaMname returns the SoaMname field if non-nil, zero value otherwise.

### GetSoaMnameOk

`func (o *GetZoneStubResponse) GetSoaMnameOk() (*string, bool)`

GetSoaMnameOk returns a tuple with the SoaMname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaMname

`func (o *GetZoneStubResponse) SetSoaMname(v string)`

SetSoaMname sets SoaMname field to given value.

### HasSoaMname

`func (o *GetZoneStubResponse) HasSoaMname() bool`

HasSoaMname returns a boolean if a field has been set.

### GetSoaNegativeTtl

`func (o *GetZoneStubResponse) GetSoaNegativeTtl() int64`

GetSoaNegativeTtl returns the SoaNegativeTtl field if non-nil, zero value otherwise.

### GetSoaNegativeTtlOk

`func (o *GetZoneStubResponse) GetSoaNegativeTtlOk() (*int64, bool)`

GetSoaNegativeTtlOk returns a tuple with the SoaNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaNegativeTtl

`func (o *GetZoneStubResponse) SetSoaNegativeTtl(v int64)`

SetSoaNegativeTtl sets SoaNegativeTtl field to given value.

### HasSoaNegativeTtl

`func (o *GetZoneStubResponse) HasSoaNegativeTtl() bool`

HasSoaNegativeTtl returns a boolean if a field has been set.

### GetSoaRefresh

`func (o *GetZoneStubResponse) GetSoaRefresh() int64`

GetSoaRefresh returns the SoaRefresh field if non-nil, zero value otherwise.

### GetSoaRefreshOk

`func (o *GetZoneStubResponse) GetSoaRefreshOk() (*int64, bool)`

GetSoaRefreshOk returns a tuple with the SoaRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRefresh

`func (o *GetZoneStubResponse) SetSoaRefresh(v int64)`

SetSoaRefresh sets SoaRefresh field to given value.

### HasSoaRefresh

`func (o *GetZoneStubResponse) HasSoaRefresh() bool`

HasSoaRefresh returns a boolean if a field has been set.

### GetSoaRetry

`func (o *GetZoneStubResponse) GetSoaRetry() int64`

GetSoaRetry returns the SoaRetry field if non-nil, zero value otherwise.

### GetSoaRetryOk

`func (o *GetZoneStubResponse) GetSoaRetryOk() (*int64, bool)`

GetSoaRetryOk returns a tuple with the SoaRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRetry

`func (o *GetZoneStubResponse) SetSoaRetry(v int64)`

SetSoaRetry sets SoaRetry field to given value.

### HasSoaRetry

`func (o *GetZoneStubResponse) HasSoaRetry() bool`

HasSoaRetry returns a boolean if a field has been set.

### GetSoaSerial

`func (o *GetZoneStubResponse) GetSoaSerial() int64`

GetSoaSerial returns the SoaSerial field if non-nil, zero value otherwise.

### GetSoaSerialOk

`func (o *GetZoneStubResponse) GetSoaSerialOk() (*int64, bool)`

GetSoaSerialOk returns a tuple with the SoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaSerial

`func (o *GetZoneStubResponse) SetSoaSerial(v int64)`

SetSoaSerial sets SoaSerial field to given value.

### HasSoaSerial

`func (o *GetZoneStubResponse) HasSoaSerial() bool`

HasSoaSerial returns a boolean if a field has been set.

### GetStubFrom

`func (o *GetZoneStubResponse) GetStubFrom() []ZoneStubStubFrom`

GetStubFrom returns the StubFrom field if non-nil, zero value otherwise.

### GetStubFromOk

`func (o *GetZoneStubResponse) GetStubFromOk() (*[]ZoneStubStubFrom, bool)`

GetStubFromOk returns a tuple with the StubFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubFrom

`func (o *GetZoneStubResponse) SetStubFrom(v []ZoneStubStubFrom)`

SetStubFrom sets StubFrom field to given value.

### HasStubFrom

`func (o *GetZoneStubResponse) HasStubFrom() bool`

HasStubFrom returns a boolean if a field has been set.

### GetStubMembers

`func (o *GetZoneStubResponse) GetStubMembers() []ZoneStubStubMembers`

GetStubMembers returns the StubMembers field if non-nil, zero value otherwise.

### GetStubMembersOk

`func (o *GetZoneStubResponse) GetStubMembersOk() (*[]ZoneStubStubMembers, bool)`

GetStubMembersOk returns a tuple with the StubMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubMembers

`func (o *GetZoneStubResponse) SetStubMembers(v []ZoneStubStubMembers)`

SetStubMembers sets StubMembers field to given value.

### HasStubMembers

`func (o *GetZoneStubResponse) HasStubMembers() bool`

HasStubMembers returns a boolean if a field has been set.

### GetStubMsservers

`func (o *GetZoneStubResponse) GetStubMsservers() []ZoneStubStubMsservers`

GetStubMsservers returns the StubMsservers field if non-nil, zero value otherwise.

### GetStubMsserversOk

`func (o *GetZoneStubResponse) GetStubMsserversOk() (*[]ZoneStubStubMsservers, bool)`

GetStubMsserversOk returns a tuple with the StubMsservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubMsservers

`func (o *GetZoneStubResponse) SetStubMsservers(v []ZoneStubStubMsservers)`

SetStubMsservers sets StubMsservers field to given value.

### HasStubMsservers

`func (o *GetZoneStubResponse) HasStubMsservers() bool`

HasStubMsservers returns a boolean if a field has been set.

### GetUsingSrgAssociations

`func (o *GetZoneStubResponse) GetUsingSrgAssociations() bool`

GetUsingSrgAssociations returns the UsingSrgAssociations field if non-nil, zero value otherwise.

### GetUsingSrgAssociationsOk

`func (o *GetZoneStubResponse) GetUsingSrgAssociationsOk() (*bool, bool)`

GetUsingSrgAssociationsOk returns a tuple with the UsingSrgAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingSrgAssociations

`func (o *GetZoneStubResponse) SetUsingSrgAssociations(v bool)`

SetUsingSrgAssociations sets UsingSrgAssociations field to given value.

### HasUsingSrgAssociations

`func (o *GetZoneStubResponse) HasUsingSrgAssociations() bool`

HasUsingSrgAssociations returns a boolean if a field has been set.

### GetView

`func (o *GetZoneStubResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetZoneStubResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetZoneStubResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetZoneStubResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZoneFormat

`func (o *GetZoneStubResponse) GetZoneFormat() string`

GetZoneFormat returns the ZoneFormat field if non-nil, zero value otherwise.

### GetZoneFormatOk

`func (o *GetZoneStubResponse) GetZoneFormatOk() (*string, bool)`

GetZoneFormatOk returns a tuple with the ZoneFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFormat

`func (o *GetZoneStubResponse) SetZoneFormat(v string)`

SetZoneFormat sets ZoneFormat field to given value.

### HasZoneFormat

`func (o *GetZoneStubResponse) HasZoneFormat() bool`

HasZoneFormat returns a boolean if a field has been set.

### GetResult

`func (o *GetZoneStubResponse) GetResult() ZoneStub`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneStubResponse) GetResultOk() (*ZoneStub, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneStubResponse) SetResult(v ZoneStub)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetZoneStubResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


