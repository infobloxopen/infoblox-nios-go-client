# ZoneDelegated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IP address of the server that is serving this zone. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the zone; maximum 256 characters. | [optional] 
**DelegateTo** | Pointer to [**[]ZoneDelegatedDelegateTo**](ZoneDelegatedDelegateTo.md) | This provides information for the remote name server that maintains data for the delegated zone. The Infoblox appliance redirects queries for data for the delegated zone to this remote name server. | [optional] 
**DelegatedTtl** | Pointer to **int64** | You can specify the Time to Live (TTL) values of auto-generated NS and glue records for a delegated zone. This value is the number of seconds that data is cached. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a zone is disabled or not. When this is set to False, the zone is enabled. | [optional] 
**DisplayDomain** | Pointer to **string** | The displayed name of the DNS zone. | [optional] [readonly] 
**DnsFqdn** | Pointer to **string** | The name of this DNS zone in punycode format. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format in punycode format. | [optional] [readonly] 
**EnableRfc2317Exclusion** | Pointer to **bool** | This flag controls whether automatic generation of RFC 2317 CNAMEs for delegated reverse zones overwrite existing PTR records. The default behavior is to overwrite all the existing records in the range; this corresponds to \&quot;allow_ptr_creation_in_parent\&quot; set to False. However, when this flag is set to True the existing PTR records are not overwritten. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Fqdn** | Pointer to **string** | The name of this DNS zone. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set. | [optional] 
**Locked** | Pointer to **bool** | If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked. | [optional] 
**LockedBy** | Pointer to **string** | The name of a superuser or the administrator who locked this zone. | [optional] [readonly] 
**MaskPrefix** | Pointer to **string** | IPv4 Netmask or IPv6 prefix for this zone. | [optional] [readonly] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MgmPrivateOverridable** | Pointer to **bool** | This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object. | [optional] [readonly] 
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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the zone resides. Example \&quot;external\&quot;. | [optional] 
**ZoneFormat** | Pointer to **string** | Determines the format of this zone. | [optional] 

## Methods

### NewZoneDelegated

`func NewZoneDelegated() *ZoneDelegated`

NewZoneDelegated instantiates a new ZoneDelegated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneDelegatedWithDefaults

`func NewZoneDelegatedWithDefaults() *ZoneDelegated`

NewZoneDelegatedWithDefaults instantiates a new ZoneDelegated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ZoneDelegated) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ZoneDelegated) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ZoneDelegated) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ZoneDelegated) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *ZoneDelegated) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ZoneDelegated) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ZoneDelegated) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ZoneDelegated) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *ZoneDelegated) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ZoneDelegated) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ZoneDelegated) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ZoneDelegated) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegateTo

`func (o *ZoneDelegated) GetDelegateTo() []ZoneDelegatedDelegateTo`

GetDelegateTo returns the DelegateTo field if non-nil, zero value otherwise.

### GetDelegateToOk

`func (o *ZoneDelegated) GetDelegateToOk() (*[]ZoneDelegatedDelegateTo, bool)`

GetDelegateToOk returns a tuple with the DelegateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateTo

`func (o *ZoneDelegated) SetDelegateTo(v []ZoneDelegatedDelegateTo)`

SetDelegateTo sets DelegateTo field to given value.

### HasDelegateTo

`func (o *ZoneDelegated) HasDelegateTo() bool`

HasDelegateTo returns a boolean if a field has been set.

### GetDelegatedTtl

`func (o *ZoneDelegated) GetDelegatedTtl() int64`

GetDelegatedTtl returns the DelegatedTtl field if non-nil, zero value otherwise.

### GetDelegatedTtlOk

`func (o *ZoneDelegated) GetDelegatedTtlOk() (*int64, bool)`

GetDelegatedTtlOk returns a tuple with the DelegatedTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTtl

`func (o *ZoneDelegated) SetDelegatedTtl(v int64)`

SetDelegatedTtl sets DelegatedTtl field to given value.

### HasDelegatedTtl

`func (o *ZoneDelegated) HasDelegatedTtl() bool`

HasDelegatedTtl returns a boolean if a field has been set.

### GetDisable

`func (o *ZoneDelegated) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ZoneDelegated) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ZoneDelegated) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ZoneDelegated) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisplayDomain

`func (o *ZoneDelegated) GetDisplayDomain() string`

GetDisplayDomain returns the DisplayDomain field if non-nil, zero value otherwise.

### GetDisplayDomainOk

`func (o *ZoneDelegated) GetDisplayDomainOk() (*string, bool)`

GetDisplayDomainOk returns a tuple with the DisplayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDomain

`func (o *ZoneDelegated) SetDisplayDomain(v string)`

SetDisplayDomain sets DisplayDomain field to given value.

### HasDisplayDomain

`func (o *ZoneDelegated) HasDisplayDomain() bool`

HasDisplayDomain returns a boolean if a field has been set.

### GetDnsFqdn

`func (o *ZoneDelegated) GetDnsFqdn() string`

GetDnsFqdn returns the DnsFqdn field if non-nil, zero value otherwise.

### GetDnsFqdnOk

`func (o *ZoneDelegated) GetDnsFqdnOk() (*string, bool)`

GetDnsFqdnOk returns a tuple with the DnsFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFqdn

`func (o *ZoneDelegated) SetDnsFqdn(v string)`

SetDnsFqdn sets DnsFqdn field to given value.

### HasDnsFqdn

`func (o *ZoneDelegated) HasDnsFqdn() bool`

HasDnsFqdn returns a boolean if a field has been set.

### GetEnableRfc2317Exclusion

`func (o *ZoneDelegated) GetEnableRfc2317Exclusion() bool`

GetEnableRfc2317Exclusion returns the EnableRfc2317Exclusion field if non-nil, zero value otherwise.

### GetEnableRfc2317ExclusionOk

`func (o *ZoneDelegated) GetEnableRfc2317ExclusionOk() (*bool, bool)`

GetEnableRfc2317ExclusionOk returns a tuple with the EnableRfc2317Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRfc2317Exclusion

`func (o *ZoneDelegated) SetEnableRfc2317Exclusion(v bool)`

SetEnableRfc2317Exclusion sets EnableRfc2317Exclusion field to given value.

### HasEnableRfc2317Exclusion

`func (o *ZoneDelegated) HasEnableRfc2317Exclusion() bool`

HasEnableRfc2317Exclusion returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *ZoneDelegated) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *ZoneDelegated) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *ZoneDelegated) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *ZoneDelegated) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *ZoneDelegated) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *ZoneDelegated) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *ZoneDelegated) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *ZoneDelegated) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *ZoneDelegated) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *ZoneDelegated) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *ZoneDelegated) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *ZoneDelegated) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFqdn

`func (o *ZoneDelegated) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ZoneDelegated) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ZoneDelegated) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ZoneDelegated) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetLocked

`func (o *ZoneDelegated) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ZoneDelegated) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ZoneDelegated) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ZoneDelegated) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *ZoneDelegated) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *ZoneDelegated) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *ZoneDelegated) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *ZoneDelegated) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetMaskPrefix

`func (o *ZoneDelegated) GetMaskPrefix() string`

GetMaskPrefix returns the MaskPrefix field if non-nil, zero value otherwise.

### GetMaskPrefixOk

`func (o *ZoneDelegated) GetMaskPrefixOk() (*string, bool)`

GetMaskPrefixOk returns a tuple with the MaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPrefix

`func (o *ZoneDelegated) SetMaskPrefix(v string)`

SetMaskPrefix sets MaskPrefix field to given value.

### HasMaskPrefix

`func (o *ZoneDelegated) HasMaskPrefix() bool`

HasMaskPrefix returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *ZoneDelegated) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *ZoneDelegated) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *ZoneDelegated) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *ZoneDelegated) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *ZoneDelegated) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *ZoneDelegated) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *ZoneDelegated) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *ZoneDelegated) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdIntegrated

`func (o *ZoneDelegated) GetMsAdIntegrated() bool`

GetMsAdIntegrated returns the MsAdIntegrated field if non-nil, zero value otherwise.

### GetMsAdIntegratedOk

`func (o *ZoneDelegated) GetMsAdIntegratedOk() (*bool, bool)`

GetMsAdIntegratedOk returns a tuple with the MsAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdIntegrated

`func (o *ZoneDelegated) SetMsAdIntegrated(v bool)`

SetMsAdIntegrated sets MsAdIntegrated field to given value.

### HasMsAdIntegrated

`func (o *ZoneDelegated) HasMsAdIntegrated() bool`

HasMsAdIntegrated returns a boolean if a field has been set.

### GetMsDdnsMode

`func (o *ZoneDelegated) GetMsDdnsMode() string`

GetMsDdnsMode returns the MsDdnsMode field if non-nil, zero value otherwise.

### GetMsDdnsModeOk

`func (o *ZoneDelegated) GetMsDdnsModeOk() (*string, bool)`

GetMsDdnsModeOk returns a tuple with the MsDdnsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDdnsMode

`func (o *ZoneDelegated) SetMsDdnsMode(v string)`

SetMsDdnsMode sets MsDdnsMode field to given value.

### HasMsDdnsMode

`func (o *ZoneDelegated) HasMsDdnsMode() bool`

HasMsDdnsMode returns a boolean if a field has been set.

### GetMsManaged

`func (o *ZoneDelegated) GetMsManaged() string`

GetMsManaged returns the MsManaged field if non-nil, zero value otherwise.

### GetMsManagedOk

`func (o *ZoneDelegated) GetMsManagedOk() (*string, bool)`

GetMsManagedOk returns a tuple with the MsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsManaged

`func (o *ZoneDelegated) SetMsManaged(v string)`

SetMsManaged sets MsManaged field to given value.

### HasMsManaged

`func (o *ZoneDelegated) HasMsManaged() bool`

HasMsManaged returns a boolean if a field has been set.

### GetMsReadOnly

`func (o *ZoneDelegated) GetMsReadOnly() bool`

GetMsReadOnly returns the MsReadOnly field if non-nil, zero value otherwise.

### GetMsReadOnlyOk

`func (o *ZoneDelegated) GetMsReadOnlyOk() (*bool, bool)`

GetMsReadOnlyOk returns a tuple with the MsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsReadOnly

`func (o *ZoneDelegated) SetMsReadOnly(v bool)`

SetMsReadOnly sets MsReadOnly field to given value.

### HasMsReadOnly

`func (o *ZoneDelegated) HasMsReadOnly() bool`

HasMsReadOnly returns a boolean if a field has been set.

### GetMsSyncMasterName

`func (o *ZoneDelegated) GetMsSyncMasterName() string`

GetMsSyncMasterName returns the MsSyncMasterName field if non-nil, zero value otherwise.

### GetMsSyncMasterNameOk

`func (o *ZoneDelegated) GetMsSyncMasterNameOk() (*string, bool)`

GetMsSyncMasterNameOk returns a tuple with the MsSyncMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncMasterName

`func (o *ZoneDelegated) SetMsSyncMasterName(v string)`

SetMsSyncMasterName sets MsSyncMasterName field to given value.

### HasMsSyncMasterName

`func (o *ZoneDelegated) HasMsSyncMasterName() bool`

HasMsSyncMasterName returns a boolean if a field has been set.

### GetNsGroup

`func (o *ZoneDelegated) GetNsGroup() string`

GetNsGroup returns the NsGroup field if non-nil, zero value otherwise.

### GetNsGroupOk

`func (o *ZoneDelegated) GetNsGroupOk() (*string, bool)`

GetNsGroupOk returns a tuple with the NsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsGroup

`func (o *ZoneDelegated) SetNsGroup(v string)`

SetNsGroup sets NsGroup field to given value.

### HasNsGroup

`func (o *ZoneDelegated) HasNsGroup() bool`

HasNsGroup returns a boolean if a field has been set.

### GetParent

`func (o *ZoneDelegated) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ZoneDelegated) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ZoneDelegated) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ZoneDelegated) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrefix

`func (o *ZoneDelegated) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ZoneDelegated) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ZoneDelegated) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ZoneDelegated) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetUseDelegatedTtl

`func (o *ZoneDelegated) GetUseDelegatedTtl() bool`

GetUseDelegatedTtl returns the UseDelegatedTtl field if non-nil, zero value otherwise.

### GetUseDelegatedTtlOk

`func (o *ZoneDelegated) GetUseDelegatedTtlOk() (*bool, bool)`

GetUseDelegatedTtlOk returns a tuple with the UseDelegatedTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDelegatedTtl

`func (o *ZoneDelegated) SetUseDelegatedTtl(v bool)`

SetUseDelegatedTtl sets UseDelegatedTtl field to given value.

### HasUseDelegatedTtl

`func (o *ZoneDelegated) HasUseDelegatedTtl() bool`

HasUseDelegatedTtl returns a boolean if a field has been set.

### GetUsingSrgAssociations

`func (o *ZoneDelegated) GetUsingSrgAssociations() bool`

GetUsingSrgAssociations returns the UsingSrgAssociations field if non-nil, zero value otherwise.

### GetUsingSrgAssociationsOk

`func (o *ZoneDelegated) GetUsingSrgAssociationsOk() (*bool, bool)`

GetUsingSrgAssociationsOk returns a tuple with the UsingSrgAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingSrgAssociations

`func (o *ZoneDelegated) SetUsingSrgAssociations(v bool)`

SetUsingSrgAssociations sets UsingSrgAssociations field to given value.

### HasUsingSrgAssociations

`func (o *ZoneDelegated) HasUsingSrgAssociations() bool`

HasUsingSrgAssociations returns a boolean if a field has been set.

### GetUuid

`func (o *ZoneDelegated) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ZoneDelegated) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ZoneDelegated) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ZoneDelegated) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetView

`func (o *ZoneDelegated) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ZoneDelegated) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ZoneDelegated) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ZoneDelegated) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZoneFormat

`func (o *ZoneDelegated) GetZoneFormat() string`

GetZoneFormat returns the ZoneFormat field if non-nil, zero value otherwise.

### GetZoneFormatOk

`func (o *ZoneDelegated) GetZoneFormatOk() (*string, bool)`

GetZoneFormatOk returns a tuple with the ZoneFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFormat

`func (o *ZoneDelegated) SetZoneFormat(v string)`

SetZoneFormat sets ZoneFormat field to given value.

### HasZoneFormat

`func (o *ZoneDelegated) HasZoneFormat() bool`

HasZoneFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


