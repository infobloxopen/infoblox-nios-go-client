# GetNetworkviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AssociatedDnsViews** | Pointer to **[]string** | The list of DNS views associated with this network view. | [optional] [readonly] 
**AssociatedMembers** | Pointer to [**[]NetworkviewAssociatedMembers**](NetworkviewAssociatedMembers.md) | The list of members associated with a network view. | [optional] [readonly] 
**CloudInfo** | Pointer to [**NetworkviewCloudInfo**](NetworkviewCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the network view; maximum 256 characters. | [optional] 
**DdnsDnsView** | Pointer to **string** | DNS views that will receive the updates if you enable the appliance to send updates to Grid members. | [optional] 
**DdnsZonePrimaries** | Pointer to [**[]NetworkviewDdnsZonePrimaries**](NetworkviewDdnsZonePrimaries.md) | An array of Ddns Zone Primary dhcpddns structs that lists the information of primary zone to wich DDNS updates should be sent. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FederatedRealms** | Pointer to [**[]NetworkviewFederatedRealms**](NetworkviewFederatedRealms.md) | This field contains the federated realms associated to this network view | [optional] 
**InternalForwardZones** | Pointer to **[]string** | The list of linked authoritative DNS zones. | [optional] 
**IsDefault** | Pointer to **bool** | The NIOS appliance provides one default network view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one network view in the appliance. | [optional] [readonly] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MsAdUserData** | Pointer to [**NetworkviewMsAdUserData**](NetworkviewMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the network view. | [optional] 
**RemoteForwardZones** | Pointer to [**[]NetworkviewRemoteForwardZones**](NetworkviewRemoteForwardZones.md) | The list of forward-mapping zones to which the DHCP server sends the updates. | [optional] 
**RemoteReverseZones** | Pointer to [**[]NetworkviewRemoteReverseZones**](NetworkviewRemoteReverseZones.md) | The list of reverse-mapping zones to which the DHCP server sends the updates. | [optional] 
**Result** | Pointer to [**Networkview**](Networkview.md) |  | [optional] 

## Methods

### NewGetNetworkviewResponse

`func NewGetNetworkviewResponse() *GetNetworkviewResponse`

NewGetNetworkviewResponse instantiates a new GetNetworkviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkviewResponseWithDefaults

`func NewGetNetworkviewResponseWithDefaults() *GetNetworkviewResponse`

NewGetNetworkviewResponseWithDefaults instantiates a new GetNetworkviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNetworkviewResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNetworkviewResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNetworkviewResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNetworkviewResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssociatedDnsViews

`func (o *GetNetworkviewResponse) GetAssociatedDnsViews() []string`

GetAssociatedDnsViews returns the AssociatedDnsViews field if non-nil, zero value otherwise.

### GetAssociatedDnsViewsOk

`func (o *GetNetworkviewResponse) GetAssociatedDnsViewsOk() (*[]string, bool)`

GetAssociatedDnsViewsOk returns a tuple with the AssociatedDnsViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDnsViews

`func (o *GetNetworkviewResponse) SetAssociatedDnsViews(v []string)`

SetAssociatedDnsViews sets AssociatedDnsViews field to given value.

### HasAssociatedDnsViews

`func (o *GetNetworkviewResponse) HasAssociatedDnsViews() bool`

HasAssociatedDnsViews returns a boolean if a field has been set.

### GetAssociatedMembers

`func (o *GetNetworkviewResponse) GetAssociatedMembers() []NetworkviewAssociatedMembers`

GetAssociatedMembers returns the AssociatedMembers field if non-nil, zero value otherwise.

### GetAssociatedMembersOk

`func (o *GetNetworkviewResponse) GetAssociatedMembersOk() (*[]NetworkviewAssociatedMembers, bool)`

GetAssociatedMembersOk returns a tuple with the AssociatedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMembers

`func (o *GetNetworkviewResponse) SetAssociatedMembers(v []NetworkviewAssociatedMembers)`

SetAssociatedMembers sets AssociatedMembers field to given value.

### HasAssociatedMembers

`func (o *GetNetworkviewResponse) HasAssociatedMembers() bool`

HasAssociatedMembers returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetNetworkviewResponse) GetCloudInfo() NetworkviewCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetNetworkviewResponse) GetCloudInfoOk() (*NetworkviewCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetNetworkviewResponse) SetCloudInfo(v NetworkviewCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetNetworkviewResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetNetworkviewResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNetworkviewResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNetworkviewResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNetworkviewResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDdnsDnsView

`func (o *GetNetworkviewResponse) GetDdnsDnsView() string`

GetDdnsDnsView returns the DdnsDnsView field if non-nil, zero value otherwise.

### GetDdnsDnsViewOk

`func (o *GetNetworkviewResponse) GetDdnsDnsViewOk() (*string, bool)`

GetDdnsDnsViewOk returns a tuple with the DdnsDnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDnsView

`func (o *GetNetworkviewResponse) SetDdnsDnsView(v string)`

SetDdnsDnsView sets DdnsDnsView field to given value.

### HasDdnsDnsView

`func (o *GetNetworkviewResponse) HasDdnsDnsView() bool`

HasDdnsDnsView returns a boolean if a field has been set.

### GetDdnsZonePrimaries

`func (o *GetNetworkviewResponse) GetDdnsZonePrimaries() []NetworkviewDdnsZonePrimaries`

GetDdnsZonePrimaries returns the DdnsZonePrimaries field if non-nil, zero value otherwise.

### GetDdnsZonePrimariesOk

`func (o *GetNetworkviewResponse) GetDdnsZonePrimariesOk() (*[]NetworkviewDdnsZonePrimaries, bool)`

GetDdnsZonePrimariesOk returns a tuple with the DdnsZonePrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsZonePrimaries

`func (o *GetNetworkviewResponse) SetDdnsZonePrimaries(v []NetworkviewDdnsZonePrimaries)`

SetDdnsZonePrimaries sets DdnsZonePrimaries field to given value.

### HasDdnsZonePrimaries

`func (o *GetNetworkviewResponse) HasDdnsZonePrimaries() bool`

HasDdnsZonePrimaries returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetNetworkviewResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetNetworkviewResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetNetworkviewResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetNetworkviewResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *GetNetworkviewResponse) GetFederatedRealms() []NetworkviewFederatedRealms`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *GetNetworkviewResponse) GetFederatedRealmsOk() (*[]NetworkviewFederatedRealms, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *GetNetworkviewResponse) SetFederatedRealms(v []NetworkviewFederatedRealms)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *GetNetworkviewResponse) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetInternalForwardZones

`func (o *GetNetworkviewResponse) GetInternalForwardZones() []string`

GetInternalForwardZones returns the InternalForwardZones field if non-nil, zero value otherwise.

### GetInternalForwardZonesOk

`func (o *GetNetworkviewResponse) GetInternalForwardZonesOk() (*[]string, bool)`

GetInternalForwardZonesOk returns a tuple with the InternalForwardZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwardZones

`func (o *GetNetworkviewResponse) SetInternalForwardZones(v []string)`

SetInternalForwardZones sets InternalForwardZones field to given value.

### HasInternalForwardZones

`func (o *GetNetworkviewResponse) HasInternalForwardZones() bool`

HasInternalForwardZones returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetNetworkviewResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetNetworkviewResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetNetworkviewResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetNetworkviewResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *GetNetworkviewResponse) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *GetNetworkviewResponse) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *GetNetworkviewResponse) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *GetNetworkviewResponse) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetNetworkviewResponse) GetMsAdUserData() NetworkviewMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetNetworkviewResponse) GetMsAdUserDataOk() (*NetworkviewMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetNetworkviewResponse) SetMsAdUserData(v NetworkviewMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetNetworkviewResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkviewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkviewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkviewResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkviewResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteForwardZones

`func (o *GetNetworkviewResponse) GetRemoteForwardZones() []NetworkviewRemoteForwardZones`

GetRemoteForwardZones returns the RemoteForwardZones field if non-nil, zero value otherwise.

### GetRemoteForwardZonesOk

`func (o *GetNetworkviewResponse) GetRemoteForwardZonesOk() (*[]NetworkviewRemoteForwardZones, bool)`

GetRemoteForwardZonesOk returns a tuple with the RemoteForwardZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteForwardZones

`func (o *GetNetworkviewResponse) SetRemoteForwardZones(v []NetworkviewRemoteForwardZones)`

SetRemoteForwardZones sets RemoteForwardZones field to given value.

### HasRemoteForwardZones

`func (o *GetNetworkviewResponse) HasRemoteForwardZones() bool`

HasRemoteForwardZones returns a boolean if a field has been set.

### GetRemoteReverseZones

`func (o *GetNetworkviewResponse) GetRemoteReverseZones() []NetworkviewRemoteReverseZones`

GetRemoteReverseZones returns the RemoteReverseZones field if non-nil, zero value otherwise.

### GetRemoteReverseZonesOk

`func (o *GetNetworkviewResponse) GetRemoteReverseZonesOk() (*[]NetworkviewRemoteReverseZones, bool)`

GetRemoteReverseZonesOk returns a tuple with the RemoteReverseZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteReverseZones

`func (o *GetNetworkviewResponse) SetRemoteReverseZones(v []NetworkviewRemoteReverseZones)`

SetRemoteReverseZones sets RemoteReverseZones field to given value.

### HasRemoteReverseZones

`func (o *GetNetworkviewResponse) HasRemoteReverseZones() bool`

HasRemoteReverseZones returns a boolean if a field has been set.

### GetResult

`func (o *GetNetworkviewResponse) GetResult() Networkview`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNetworkviewResponse) GetResultOk() (*Networkview, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNetworkviewResponse) SetResult(v Networkview)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNetworkviewResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


