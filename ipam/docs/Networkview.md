# Networkview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AssociatedDnsViews** | Pointer to **[]string** | The list of DNS views associated with this network view. | [optional] [readonly] 
**AssociatedMembers** | Pointer to [**[]NetworkviewAssociatedMembers**](NetworkviewAssociatedMembers.md) | The list of members associated with a network view. | [optional] [readonly] 
**CloudInfo** | Pointer to [**NetworkviewCloudInfo**](NetworkviewCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the network view; maximum 256 characters. | [optional] 
**DdnsDnsView** | Pointer to **string** | DNS views that will receive the updates if you enable the appliance to send updates to Grid members. | [optional] 
**DdnsZonePrimaries** | Pointer to [**[]NetworkviewDdnsZonePrimaries**](NetworkviewDdnsZonePrimaries.md) | An array of Ddns Zone Primary dhcpddns structs that lists the information of primary zone to wich DDNS updates should be sent. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FederatedRealms** | Pointer to [**[]NetworkviewFederatedRealms**](NetworkviewFederatedRealms.md) | This field contains the federated realms associated to this network view | [optional] 
**InternalForwardZones** | Pointer to **[]string** | The list of linked authoritative DNS zones. | [optional] 
**IsDefault** | Pointer to **bool** | The NIOS appliance provides one default network view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one network view in the appliance. | [optional] [readonly] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MsAdUserData** | Pointer to [**NetworkviewMsAdUserData**](NetworkviewMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the network view. | [optional] 
**RemoteForwardZones** | Pointer to [**[]NetworkviewRemoteForwardZones**](NetworkviewRemoteForwardZones.md) | The list of forward-mapping zones to which the DHCP server sends the updates. | [optional] 
**RemoteReverseZones** | Pointer to [**[]NetworkviewRemoteReverseZones**](NetworkviewRemoteReverseZones.md) | The list of reverse-mapping zones to which the DHCP server sends the updates. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewNetworkview

`func NewNetworkview() *Networkview`

NewNetworkview instantiates a new Networkview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkviewWithDefaults

`func NewNetworkviewWithDefaults() *Networkview`

NewNetworkviewWithDefaults instantiates a new Networkview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Networkview) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Networkview) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Networkview) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Networkview) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssociatedDnsViews

`func (o *Networkview) GetAssociatedDnsViews() []string`

GetAssociatedDnsViews returns the AssociatedDnsViews field if non-nil, zero value otherwise.

### GetAssociatedDnsViewsOk

`func (o *Networkview) GetAssociatedDnsViewsOk() (*[]string, bool)`

GetAssociatedDnsViewsOk returns a tuple with the AssociatedDnsViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDnsViews

`func (o *Networkview) SetAssociatedDnsViews(v []string)`

SetAssociatedDnsViews sets AssociatedDnsViews field to given value.

### HasAssociatedDnsViews

`func (o *Networkview) HasAssociatedDnsViews() bool`

HasAssociatedDnsViews returns a boolean if a field has been set.

### GetAssociatedMembers

`func (o *Networkview) GetAssociatedMembers() []NetworkviewAssociatedMembers`

GetAssociatedMembers returns the AssociatedMembers field if non-nil, zero value otherwise.

### GetAssociatedMembersOk

`func (o *Networkview) GetAssociatedMembersOk() (*[]NetworkviewAssociatedMembers, bool)`

GetAssociatedMembersOk returns a tuple with the AssociatedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMembers

`func (o *Networkview) SetAssociatedMembers(v []NetworkviewAssociatedMembers)`

SetAssociatedMembers sets AssociatedMembers field to given value.

### HasAssociatedMembers

`func (o *Networkview) HasAssociatedMembers() bool`

HasAssociatedMembers returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Networkview) GetCloudInfo() NetworkviewCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Networkview) GetCloudInfoOk() (*NetworkviewCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Networkview) SetCloudInfo(v NetworkviewCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Networkview) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Networkview) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Networkview) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Networkview) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Networkview) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDnsView

`func (o *Networkview) GetDdnsDnsView() string`

GetDdnsDnsView returns the DdnsDnsView field if non-nil, zero value otherwise.

### GetDdnsDnsViewOk

`func (o *Networkview) GetDdnsDnsViewOk() (*string, bool)`

GetDdnsDnsViewOk returns a tuple with the DdnsDnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDnsView

`func (o *Networkview) SetDdnsDnsView(v string)`

SetDdnsDnsView sets DdnsDnsView field to given value.

### HasDdnsDnsView

`func (o *Networkview) HasDdnsDnsView() bool`

HasDdnsDnsView returns a boolean if a field has been set.

### GetDdnsZonePrimaries

`func (o *Networkview) GetDdnsZonePrimaries() []NetworkviewDdnsZonePrimaries`

GetDdnsZonePrimaries returns the DdnsZonePrimaries field if non-nil, zero value otherwise.

### GetDdnsZonePrimariesOk

`func (o *Networkview) GetDdnsZonePrimariesOk() (*[]NetworkviewDdnsZonePrimaries, bool)`

GetDdnsZonePrimariesOk returns a tuple with the DdnsZonePrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsZonePrimaries

`func (o *Networkview) SetDdnsZonePrimaries(v []NetworkviewDdnsZonePrimaries)`

SetDdnsZonePrimaries sets DdnsZonePrimaries field to given value.

### HasDdnsZonePrimaries

`func (o *Networkview) HasDdnsZonePrimaries() bool`

HasDdnsZonePrimaries returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Networkview) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Networkview) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Networkview) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Networkview) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Networkview) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Networkview) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Networkview) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Networkview) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Networkview) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Networkview) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Networkview) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Networkview) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *Networkview) GetFederatedRealms() []NetworkviewFederatedRealms`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *Networkview) GetFederatedRealmsOk() (*[]NetworkviewFederatedRealms, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *Networkview) SetFederatedRealms(v []NetworkviewFederatedRealms)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *Networkview) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetInternalForwardZones

`func (o *Networkview) GetInternalForwardZones() []string`

GetInternalForwardZones returns the InternalForwardZones field if non-nil, zero value otherwise.

### GetInternalForwardZonesOk

`func (o *Networkview) GetInternalForwardZonesOk() (*[]string, bool)`

GetInternalForwardZonesOk returns a tuple with the InternalForwardZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwardZones

`func (o *Networkview) SetInternalForwardZones(v []string)`

SetInternalForwardZones sets InternalForwardZones field to given value.

### HasInternalForwardZones

`func (o *Networkview) HasInternalForwardZones() bool`

HasInternalForwardZones returns a boolean if a field has been set.

### GetIsDefault

`func (o *Networkview) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Networkview) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Networkview) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Networkview) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *Networkview) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *Networkview) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *Networkview) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *Networkview) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Networkview) GetMsAdUserData() NetworkviewMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Networkview) GetMsAdUserDataOk() (*NetworkviewMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Networkview) SetMsAdUserData(v NetworkviewMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Networkview) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *Networkview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Networkview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Networkview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Networkview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteForwardZones

`func (o *Networkview) GetRemoteForwardZones() []NetworkviewRemoteForwardZones`

GetRemoteForwardZones returns the RemoteForwardZones field if non-nil, zero value otherwise.

### GetRemoteForwardZonesOk

`func (o *Networkview) GetRemoteForwardZonesOk() (*[]NetworkviewRemoteForwardZones, bool)`

GetRemoteForwardZonesOk returns a tuple with the RemoteForwardZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteForwardZones

`func (o *Networkview) SetRemoteForwardZones(v []NetworkviewRemoteForwardZones)`

SetRemoteForwardZones sets RemoteForwardZones field to given value.

### HasRemoteForwardZones

`func (o *Networkview) HasRemoteForwardZones() bool`

HasRemoteForwardZones returns a boolean if a field has been set.

### GetRemoteReverseZones

`func (o *Networkview) GetRemoteReverseZones() []NetworkviewRemoteReverseZones`

GetRemoteReverseZones returns the RemoteReverseZones field if non-nil, zero value otherwise.

### GetRemoteReverseZonesOk

`func (o *Networkview) GetRemoteReverseZonesOk() (*[]NetworkviewRemoteReverseZones, bool)`

GetRemoteReverseZonesOk returns a tuple with the RemoteReverseZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteReverseZones

`func (o *Networkview) SetRemoteReverseZones(v []NetworkviewRemoteReverseZones)`

SetRemoteReverseZones sets RemoteReverseZones field to given value.

### HasRemoteReverseZones

`func (o *Networkview) HasRemoteReverseZones() bool`

HasRemoteReverseZones returns a boolean if a field has been set.

### GetUuid

`func (o *Networkview) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Networkview) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Networkview) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Networkview) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


