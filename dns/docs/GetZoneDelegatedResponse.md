# GetZoneDelegatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The IP address of the server that is serving this zone. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the zone; maximum 256 characters. | [optional] 
**DelegateTo** | Pointer to [**[]ZoneDelegatedDelegateTo**](ZoneDelegatedDelegateTo.md) | This provides information for the remote name server that maintains data for the delegated zone. The Infoblox appliance redirects queries for data for the delegated zone to this remote name server. | [optional] 
**DelegatedTtl** | Pointer to **int64** | You can specify the Time to Live (TTL) values of auto-generated NS and glue records for a delegated zone. This value is the number of seconds that data is cached. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a zone is disabled or not. When this is set to False, the zone is enabled. | [optional] 
**DisplayDomain** | Pointer to **string** | The displayed name of the DNS zone. | [optional] [readonly] 
**DnsFqdn** | Pointer to **string** | The name of this DNS zone in punycode format. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format in punycode format. | [optional] [readonly] 
**EnableRfc2317Exclusion** | Pointer to **bool** | This flag controls whether automatic generation of RFC 2317 CNAMEs for delegated reverse zones overwrite existing PTR records. The default behavior is to overwrite all the existing records in the range; this corresponds to \&quot;allow_ptr_creation_in_parent\&quot; set to False. However, when this flag is set to True the existing PTR records are not overwritten. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Fqdn** | Pointer to **string** | The name of this DNS zone. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set. | [optional] 
**Locked** | Pointer to **bool** | If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked. | [optional] 
**LockedBy** | Pointer to **string** | The name of a superuser or the administrator who locked this zone. | [optional] [readonly] 
**MaskPrefix** | Pointer to **string** | IPv4 Netmask or IPv6 prefix for this zone. | [optional] [readonly] 
**MsAdIntegrated** | Pointer to **bool** | The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \&quot;STUB\&quot;, \&quot;AUTH_PRIMARY\&quot;, or \&quot;AUTH_BOTH\&quot;. | [optional] 
**MsDdnsMode** | Pointer to **string** | Determines whether an Active Directory-integrated zone with a Microsoft DNS server as primary allows dynamic updates. Valid values are: \&quot;SECURE\&quot; if the zone allows secure updates only. \&quot;NONE\&quot; if the zone forbids dynamic updates. \&quot;ANY\&quot; if the zone accepts both secure and nonsecure updates. This field is valid only if ms_managed is either \&quot;AUTH_PRIMARY\&quot; or \&quot;AUTH_BOTH\&quot;. If the flag ms_ad_integrated is false, the value \&quot;SECURE\&quot; is not allowed. | [optional] 
**MsManaged** | Pointer to **string** | The flag that indicates whether the zone is assigned to a Microsoft DNS server. This flag returns the authoritative name server type of the Microsoft DNS server. Valid values are: \&quot;NONE\&quot; if the zone is not assigned to any Microsoft DNS server. \&quot;STUB\&quot; if the zone is assigned to a Microsoft DNS server as a stub zone. \&quot;AUTH_PRIMARY\&quot; if only the primary server of the zone is a Microsoft DNS server. \&quot;AUTH_SECONDARY\&quot; if only the secondary server of the zone is a Microsoft DNS server. \&quot;AUTH_BOTH\&quot; if both the primary and secondary servers of the zone are Microsoft DNS servers. | [optional] [readonly] 
**MsReadOnly** | Pointer to **bool** | Determines if a Grid member manages the zone served by a Microsoft DNS server in read-only mode. This flag is true when a Grid member manages the zone in read-only mode, false otherwise. When the zone has the ms_read_only flag set to True, no changes can be made to this zone. | [optional] [readonly] 
**MsSyncMasterName** | Pointer to **string** | The name of MS synchronization master for this zone. | [optional] [readonly] 
**NsGroup** | Pointer to **string** | The delegation NS group bound with delegated zone. | [optional] 
**Parent** | Pointer to **string** | The parent zone of this zone. Note that when searching for reverse zones, the \&quot;in-addr.arpa\&quot; notation should be used. | [optional] [readonly] 
**Prefix** | Pointer to **string** | The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B. | [optional] 
**UseDelegatedTtl** | Pointer to **bool** | Use flag for: delegated_ttl | [optional] 
**UsingSrgAssociations** | Pointer to **bool** | This is true if the zone is associated with a shared record group. | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the zone resides. Example \&quot;external\&quot;. | [optional] 
**ZoneFormat** | Pointer to **string** | Determines the format of this zone. | [optional] 
**Result** | Pointer to [**ZoneDelegated**](ZoneDelegated.md) |  | [optional] 

## Methods

### NewGetZoneDelegatedResponse

`func NewGetZoneDelegatedResponse() *GetZoneDelegatedResponse`

NewGetZoneDelegatedResponse instantiates a new GetZoneDelegatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneDelegatedResponseWithDefaults

`func NewGetZoneDelegatedResponseWithDefaults() *GetZoneDelegatedResponse`

NewGetZoneDelegatedResponseWithDefaults instantiates a new GetZoneDelegatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetZoneDelegatedResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetZoneDelegatedResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetZoneDelegatedResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetZoneDelegatedResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetZoneDelegatedResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetZoneDelegatedResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetZoneDelegatedResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetZoneDelegatedResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *GetZoneDelegatedResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetZoneDelegatedResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetZoneDelegatedResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetZoneDelegatedResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegateTo

`func (o *GetZoneDelegatedResponse) GetDelegateTo() []ZoneDelegatedDelegateTo`

GetDelegateTo returns the DelegateTo field if non-nil, zero value otherwise.

### GetDelegateToOk

`func (o *GetZoneDelegatedResponse) GetDelegateToOk() (*[]ZoneDelegatedDelegateTo, bool)`

GetDelegateToOk returns a tuple with the DelegateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateTo

`func (o *GetZoneDelegatedResponse) SetDelegateTo(v []ZoneDelegatedDelegateTo)`

SetDelegateTo sets DelegateTo field to given value.

### HasDelegateTo

`func (o *GetZoneDelegatedResponse) HasDelegateTo() bool`

HasDelegateTo returns a boolean if a field has been set.

### GetDelegatedTtl

`func (o *GetZoneDelegatedResponse) GetDelegatedTtl() int64`

GetDelegatedTtl returns the DelegatedTtl field if non-nil, zero value otherwise.

### GetDelegatedTtlOk

`func (o *GetZoneDelegatedResponse) GetDelegatedTtlOk() (*int64, bool)`

GetDelegatedTtlOk returns a tuple with the DelegatedTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTtl

`func (o *GetZoneDelegatedResponse) SetDelegatedTtl(v int64)`

SetDelegatedTtl sets DelegatedTtl field to given value.

### HasDelegatedTtl

`func (o *GetZoneDelegatedResponse) HasDelegatedTtl() bool`

HasDelegatedTtl returns a boolean if a field has been set.

### GetDisable

`func (o *GetZoneDelegatedResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetZoneDelegatedResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetZoneDelegatedResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetZoneDelegatedResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisplayDomain

`func (o *GetZoneDelegatedResponse) GetDisplayDomain() string`

GetDisplayDomain returns the DisplayDomain field if non-nil, zero value otherwise.

### GetDisplayDomainOk

`func (o *GetZoneDelegatedResponse) GetDisplayDomainOk() (*string, bool)`

GetDisplayDomainOk returns a tuple with the DisplayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDomain

`func (o *GetZoneDelegatedResponse) SetDisplayDomain(v string)`

SetDisplayDomain sets DisplayDomain field to given value.

### HasDisplayDomain

`func (o *GetZoneDelegatedResponse) HasDisplayDomain() bool`

HasDisplayDomain returns a boolean if a field has been set.

### GetDnsFqdn

`func (o *GetZoneDelegatedResponse) GetDnsFqdn() string`

GetDnsFqdn returns the DnsFqdn field if non-nil, zero value otherwise.

### GetDnsFqdnOk

`func (o *GetZoneDelegatedResponse) GetDnsFqdnOk() (*string, bool)`

GetDnsFqdnOk returns a tuple with the DnsFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFqdn

`func (o *GetZoneDelegatedResponse) SetDnsFqdn(v string)`

SetDnsFqdn sets DnsFqdn field to given value.

### HasDnsFqdn

`func (o *GetZoneDelegatedResponse) HasDnsFqdn() bool`

HasDnsFqdn returns a boolean if a field has been set.

### GetEnableRfc2317Exclusion

`func (o *GetZoneDelegatedResponse) GetEnableRfc2317Exclusion() bool`

GetEnableRfc2317Exclusion returns the EnableRfc2317Exclusion field if non-nil, zero value otherwise.

### GetEnableRfc2317ExclusionOk

`func (o *GetZoneDelegatedResponse) GetEnableRfc2317ExclusionOk() (*bool, bool)`

GetEnableRfc2317ExclusionOk returns a tuple with the EnableRfc2317Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRfc2317Exclusion

`func (o *GetZoneDelegatedResponse) SetEnableRfc2317Exclusion(v bool)`

SetEnableRfc2317Exclusion sets EnableRfc2317Exclusion field to given value.

### HasEnableRfc2317Exclusion

`func (o *GetZoneDelegatedResponse) HasEnableRfc2317Exclusion() bool`

HasEnableRfc2317Exclusion returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetZoneDelegatedResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetZoneDelegatedResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetZoneDelegatedResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetZoneDelegatedResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFqdn

`func (o *GetZoneDelegatedResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GetZoneDelegatedResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GetZoneDelegatedResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GetZoneDelegatedResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetLocked

`func (o *GetZoneDelegatedResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetZoneDelegatedResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetZoneDelegatedResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetZoneDelegatedResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *GetZoneDelegatedResponse) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *GetZoneDelegatedResponse) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *GetZoneDelegatedResponse) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *GetZoneDelegatedResponse) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetMaskPrefix

`func (o *GetZoneDelegatedResponse) GetMaskPrefix() string`

GetMaskPrefix returns the MaskPrefix field if non-nil, zero value otherwise.

### GetMaskPrefixOk

`func (o *GetZoneDelegatedResponse) GetMaskPrefixOk() (*string, bool)`

GetMaskPrefixOk returns a tuple with the MaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPrefix

`func (o *GetZoneDelegatedResponse) SetMaskPrefix(v string)`

SetMaskPrefix sets MaskPrefix field to given value.

### HasMaskPrefix

`func (o *GetZoneDelegatedResponse) HasMaskPrefix() bool`

HasMaskPrefix returns a boolean if a field has been set.

### GetMsAdIntegrated

`func (o *GetZoneDelegatedResponse) GetMsAdIntegrated() bool`

GetMsAdIntegrated returns the MsAdIntegrated field if non-nil, zero value otherwise.

### GetMsAdIntegratedOk

`func (o *GetZoneDelegatedResponse) GetMsAdIntegratedOk() (*bool, bool)`

GetMsAdIntegratedOk returns a tuple with the MsAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdIntegrated

`func (o *GetZoneDelegatedResponse) SetMsAdIntegrated(v bool)`

SetMsAdIntegrated sets MsAdIntegrated field to given value.

### HasMsAdIntegrated

`func (o *GetZoneDelegatedResponse) HasMsAdIntegrated() bool`

HasMsAdIntegrated returns a boolean if a field has been set.

### GetMsDdnsMode

`func (o *GetZoneDelegatedResponse) GetMsDdnsMode() string`

GetMsDdnsMode returns the MsDdnsMode field if non-nil, zero value otherwise.

### GetMsDdnsModeOk

`func (o *GetZoneDelegatedResponse) GetMsDdnsModeOk() (*string, bool)`

GetMsDdnsModeOk returns a tuple with the MsDdnsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDdnsMode

`func (o *GetZoneDelegatedResponse) SetMsDdnsMode(v string)`

SetMsDdnsMode sets MsDdnsMode field to given value.

### HasMsDdnsMode

`func (o *GetZoneDelegatedResponse) HasMsDdnsMode() bool`

HasMsDdnsMode returns a boolean if a field has been set.

### GetMsManaged

`func (o *GetZoneDelegatedResponse) GetMsManaged() string`

GetMsManaged returns the MsManaged field if non-nil, zero value otherwise.

### GetMsManagedOk

`func (o *GetZoneDelegatedResponse) GetMsManagedOk() (*string, bool)`

GetMsManagedOk returns a tuple with the MsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsManaged

`func (o *GetZoneDelegatedResponse) SetMsManaged(v string)`

SetMsManaged sets MsManaged field to given value.

### HasMsManaged

`func (o *GetZoneDelegatedResponse) HasMsManaged() bool`

HasMsManaged returns a boolean if a field has been set.

### GetMsReadOnly

`func (o *GetZoneDelegatedResponse) GetMsReadOnly() bool`

GetMsReadOnly returns the MsReadOnly field if non-nil, zero value otherwise.

### GetMsReadOnlyOk

`func (o *GetZoneDelegatedResponse) GetMsReadOnlyOk() (*bool, bool)`

GetMsReadOnlyOk returns a tuple with the MsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsReadOnly

`func (o *GetZoneDelegatedResponse) SetMsReadOnly(v bool)`

SetMsReadOnly sets MsReadOnly field to given value.

### HasMsReadOnly

`func (o *GetZoneDelegatedResponse) HasMsReadOnly() bool`

HasMsReadOnly returns a boolean if a field has been set.

### GetMsSyncMasterName

`func (o *GetZoneDelegatedResponse) GetMsSyncMasterName() string`

GetMsSyncMasterName returns the MsSyncMasterName field if non-nil, zero value otherwise.

### GetMsSyncMasterNameOk

`func (o *GetZoneDelegatedResponse) GetMsSyncMasterNameOk() (*string, bool)`

GetMsSyncMasterNameOk returns a tuple with the MsSyncMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncMasterName

`func (o *GetZoneDelegatedResponse) SetMsSyncMasterName(v string)`

SetMsSyncMasterName sets MsSyncMasterName field to given value.

### HasMsSyncMasterName

`func (o *GetZoneDelegatedResponse) HasMsSyncMasterName() bool`

HasMsSyncMasterName returns a boolean if a field has been set.

### GetNsGroup

`func (o *GetZoneDelegatedResponse) GetNsGroup() string`

GetNsGroup returns the NsGroup field if non-nil, zero value otherwise.

### GetNsGroupOk

`func (o *GetZoneDelegatedResponse) GetNsGroupOk() (*string, bool)`

GetNsGroupOk returns a tuple with the NsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsGroup

`func (o *GetZoneDelegatedResponse) SetNsGroup(v string)`

SetNsGroup sets NsGroup field to given value.

### HasNsGroup

`func (o *GetZoneDelegatedResponse) HasNsGroup() bool`

HasNsGroup returns a boolean if a field has been set.

### GetParent

`func (o *GetZoneDelegatedResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetZoneDelegatedResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetZoneDelegatedResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetZoneDelegatedResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrefix

`func (o *GetZoneDelegatedResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetZoneDelegatedResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetZoneDelegatedResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetZoneDelegatedResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetUseDelegatedTtl

`func (o *GetZoneDelegatedResponse) GetUseDelegatedTtl() bool`

GetUseDelegatedTtl returns the UseDelegatedTtl field if non-nil, zero value otherwise.

### GetUseDelegatedTtlOk

`func (o *GetZoneDelegatedResponse) GetUseDelegatedTtlOk() (*bool, bool)`

GetUseDelegatedTtlOk returns a tuple with the UseDelegatedTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDelegatedTtl

`func (o *GetZoneDelegatedResponse) SetUseDelegatedTtl(v bool)`

SetUseDelegatedTtl sets UseDelegatedTtl field to given value.

### HasUseDelegatedTtl

`func (o *GetZoneDelegatedResponse) HasUseDelegatedTtl() bool`

HasUseDelegatedTtl returns a boolean if a field has been set.

### GetUsingSrgAssociations

`func (o *GetZoneDelegatedResponse) GetUsingSrgAssociations() bool`

GetUsingSrgAssociations returns the UsingSrgAssociations field if non-nil, zero value otherwise.

### GetUsingSrgAssociationsOk

`func (o *GetZoneDelegatedResponse) GetUsingSrgAssociationsOk() (*bool, bool)`

GetUsingSrgAssociationsOk returns a tuple with the UsingSrgAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingSrgAssociations

`func (o *GetZoneDelegatedResponse) SetUsingSrgAssociations(v bool)`

SetUsingSrgAssociations sets UsingSrgAssociations field to given value.

### HasUsingSrgAssociations

`func (o *GetZoneDelegatedResponse) HasUsingSrgAssociations() bool`

HasUsingSrgAssociations returns a boolean if a field has been set.

### GetView

`func (o *GetZoneDelegatedResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetZoneDelegatedResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetZoneDelegatedResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetZoneDelegatedResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZoneFormat

`func (o *GetZoneDelegatedResponse) GetZoneFormat() string`

GetZoneFormat returns the ZoneFormat field if non-nil, zero value otherwise.

### GetZoneFormatOk

`func (o *GetZoneDelegatedResponse) GetZoneFormatOk() (*string, bool)`

GetZoneFormatOk returns a tuple with the ZoneFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFormat

`func (o *GetZoneDelegatedResponse) SetZoneFormat(v string)`

SetZoneFormat sets ZoneFormat field to given value.

### HasZoneFormat

`func (o *GetZoneDelegatedResponse) HasZoneFormat() bool`

HasZoneFormat returns a boolean if a field has been set.

### GetResult

`func (o *GetZoneDelegatedResponse) GetResult() ZoneDelegated`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneDelegatedResponse) GetResultOk() (*ZoneDelegated, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneDelegatedResponse) SetResult(v ZoneDelegated)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetZoneDelegatedResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


