# RecordAaaa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AwsRte53RecordInfo** | Pointer to [**RecordAaaaAwsRte53RecordInfo**](RecordAaaaAwsRte53RecordInfo.md) |  | [optional] 
**CloudInfo** | Pointer to [**RecordAaaaCloudInfo**](RecordAaaaCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **int64** | The time of the record creation in Epoch seconds format. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DiscoveredData** | Pointer to [**RecordAaaaDiscoveredData**](RecordAaaaDiscoveredData.md) |  | [optional] 
**DnsName** | Pointer to **string** | The name for an AAAA record in punycode format. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**Ipv6addr** | Pointer to [**RecordAaaaIpv6addr**](RecordAaaaIpv6addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**RecordAaaaMsAdUserData**](RecordAaaaMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | Name for the AAAA record in FQDN format. This value can be in unicode format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**RemoveAssociatedPtr** | Pointer to **bool** | Delete option that indicates whether the associated PTR records should be removed while deleting the specified A record. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordAaaa

`func NewRecordAaaa() *RecordAaaa`

NewRecordAaaa instantiates a new RecordAaaa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordAaaaWithDefaults

`func NewRecordAaaaWithDefaults() *RecordAaaa`

NewRecordAaaaWithDefaults instantiates a new RecordAaaa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordAaaa) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordAaaa) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordAaaa) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordAaaa) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *RecordAaaa) GetAwsRte53RecordInfo() RecordAaaaAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordAaaa) GetAwsRte53RecordInfoOk() (*RecordAaaaAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordAaaa) SetAwsRte53RecordInfo(v RecordAaaaAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordAaaa) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordAaaa) GetCloudInfo() RecordAaaaCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordAaaa) GetCloudInfoOk() (*RecordAaaaCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordAaaa) SetCloudInfo(v RecordAaaaCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordAaaa) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordAaaa) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordAaaa) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordAaaa) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordAaaa) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordAaaa) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordAaaa) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordAaaa) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordAaaa) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordAaaa) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordAaaa) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordAaaa) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordAaaa) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordAaaa) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordAaaa) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordAaaa) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordAaaa) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordAaaa) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordAaaa) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordAaaa) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordAaaa) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordAaaa) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordAaaa) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordAaaa) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordAaaa) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *RecordAaaa) GetDiscoveredData() RecordAaaaDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *RecordAaaa) GetDiscoveredDataOk() (*RecordAaaaDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *RecordAaaa) SetDiscoveredData(v RecordAaaaDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *RecordAaaa) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordAaaa) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordAaaa) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordAaaa) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordAaaa) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordAaaa) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordAaaa) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordAaaa) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordAaaa) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordAaaa) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordAaaa) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordAaaa) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordAaaa) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv6addr

`func (o *RecordAaaa) GetIpv6addr() RecordAaaaIpv6addr`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *RecordAaaa) GetIpv6addrOk() (*RecordAaaaIpv6addr, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *RecordAaaa) SetIpv6addr(v RecordAaaaIpv6addr)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *RecordAaaa) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *RecordAaaa) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *RecordAaaa) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *RecordAaaa) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *RecordAaaa) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordAaaa) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordAaaa) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordAaaa) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordAaaa) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *RecordAaaa) GetMsAdUserData() RecordAaaaMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *RecordAaaa) GetMsAdUserDataOk() (*RecordAaaaMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *RecordAaaa) SetMsAdUserData(v RecordAaaaMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *RecordAaaa) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *RecordAaaa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordAaaa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordAaaa) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordAaaa) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordAaaa) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordAaaa) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordAaaa) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordAaaa) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetRemoveAssociatedPtr

`func (o *RecordAaaa) GetRemoveAssociatedPtr() bool`

GetRemoveAssociatedPtr returns the RemoveAssociatedPtr field if non-nil, zero value otherwise.

### GetRemoveAssociatedPtrOk

`func (o *RecordAaaa) GetRemoveAssociatedPtrOk() (*bool, bool)`

GetRemoveAssociatedPtrOk returns a tuple with the RemoveAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAssociatedPtr

`func (o *RecordAaaa) SetRemoveAssociatedPtr(v bool)`

SetRemoveAssociatedPtr sets RemoveAssociatedPtr field to given value.

### HasRemoveAssociatedPtr

`func (o *RecordAaaa) HasRemoveAssociatedPtr() bool`

HasRemoveAssociatedPtr returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *RecordAaaa) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *RecordAaaa) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *RecordAaaa) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *RecordAaaa) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *RecordAaaa) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordAaaa) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordAaaa) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordAaaa) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordAaaa) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordAaaa) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordAaaa) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordAaaa) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordAaaa) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordAaaa) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordAaaa) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordAaaa) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordAaaa) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordAaaa) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordAaaa) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordAaaa) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


