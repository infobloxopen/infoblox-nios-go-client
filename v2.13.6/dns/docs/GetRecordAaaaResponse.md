# GetRecordAaaaResponse

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
**Result** | Pointer to [**RecordAaaa**](RecordAaaa.md) |  | [optional] 

## Methods

### NewGetRecordAaaaResponse

`func NewGetRecordAaaaResponse() *GetRecordAaaaResponse`

NewGetRecordAaaaResponse instantiates a new GetRecordAaaaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordAaaaResponseWithDefaults

`func NewGetRecordAaaaResponseWithDefaults() *GetRecordAaaaResponse`

NewGetRecordAaaaResponseWithDefaults instantiates a new GetRecordAaaaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordAaaaResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordAaaaResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordAaaaResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordAaaaResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *GetRecordAaaaResponse) GetAwsRte53RecordInfo() RecordAaaaAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *GetRecordAaaaResponse) GetAwsRte53RecordInfoOk() (*RecordAaaaAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *GetRecordAaaaResponse) SetAwsRte53RecordInfo(v RecordAaaaAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *GetRecordAaaaResponse) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRecordAaaaResponse) GetCloudInfo() RecordAaaaCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRecordAaaaResponse) GetCloudInfoOk() (*RecordAaaaCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRecordAaaaResponse) SetCloudInfo(v RecordAaaaCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRecordAaaaResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordAaaaResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordAaaaResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordAaaaResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordAaaaResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *GetRecordAaaaResponse) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GetRecordAaaaResponse) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GetRecordAaaaResponse) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GetRecordAaaaResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *GetRecordAaaaResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetRecordAaaaResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetRecordAaaaResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetRecordAaaaResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *GetRecordAaaaResponse) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *GetRecordAaaaResponse) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *GetRecordAaaaResponse) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *GetRecordAaaaResponse) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *GetRecordAaaaResponse) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *GetRecordAaaaResponse) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *GetRecordAaaaResponse) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *GetRecordAaaaResponse) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *GetRecordAaaaResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRecordAaaaResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRecordAaaaResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRecordAaaaResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *GetRecordAaaaResponse) GetDiscoveredData() RecordAaaaDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *GetRecordAaaaResponse) GetDiscoveredDataOk() (*RecordAaaaDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *GetRecordAaaaResponse) SetDiscoveredData(v RecordAaaaDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *GetRecordAaaaResponse) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordAaaaResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordAaaaResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordAaaaResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordAaaaResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRecordAaaaResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRecordAaaaResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRecordAaaaResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRecordAaaaResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *GetRecordAaaaResponse) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *GetRecordAaaaResponse) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *GetRecordAaaaResponse) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *GetRecordAaaaResponse) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv6addr

`func (o *GetRecordAaaaResponse) GetIpv6addr() RecordAaaaIpv6addr`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *GetRecordAaaaResponse) GetIpv6addrOk() (*RecordAaaaIpv6addr, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *GetRecordAaaaResponse) SetIpv6addr(v RecordAaaaIpv6addr)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *GetRecordAaaaResponse) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *GetRecordAaaaResponse) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *GetRecordAaaaResponse) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *GetRecordAaaaResponse) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *GetRecordAaaaResponse) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordAaaaResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordAaaaResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordAaaaResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordAaaaResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetRecordAaaaResponse) GetMsAdUserData() RecordAaaaMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetRecordAaaaResponse) GetMsAdUserDataOk() (*RecordAaaaMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetRecordAaaaResponse) SetMsAdUserData(v RecordAaaaMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetRecordAaaaResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetRecordAaaaResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordAaaaResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordAaaaResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordAaaaResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *GetRecordAaaaResponse) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *GetRecordAaaaResponse) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *GetRecordAaaaResponse) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *GetRecordAaaaResponse) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetRemoveAssociatedPtr

`func (o *GetRecordAaaaResponse) GetRemoveAssociatedPtr() bool`

GetRemoveAssociatedPtr returns the RemoveAssociatedPtr field if non-nil, zero value otherwise.

### GetRemoveAssociatedPtrOk

`func (o *GetRecordAaaaResponse) GetRemoveAssociatedPtrOk() (*bool, bool)`

GetRemoveAssociatedPtrOk returns a tuple with the RemoveAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAssociatedPtr

`func (o *GetRecordAaaaResponse) SetRemoveAssociatedPtr(v bool)`

SetRemoveAssociatedPtr sets RemoveAssociatedPtr field to given value.

### HasRemoveAssociatedPtr

`func (o *GetRecordAaaaResponse) HasRemoveAssociatedPtr() bool`

HasRemoveAssociatedPtr returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *GetRecordAaaaResponse) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *GetRecordAaaaResponse) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *GetRecordAaaaResponse) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *GetRecordAaaaResponse) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordAaaaResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordAaaaResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordAaaaResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordAaaaResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordAaaaResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordAaaaResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordAaaaResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordAaaaResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordAaaaResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordAaaaResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordAaaaResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordAaaaResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordAaaaResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordAaaaResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordAaaaResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordAaaaResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordAaaaResponse) GetResult() RecordAaaa`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordAaaaResponse) GetResultOk() (*RecordAaaa, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordAaaaResponse) SetResult(v RecordAaaa)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordAaaaResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


