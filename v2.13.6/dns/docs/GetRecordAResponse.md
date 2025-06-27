# GetRecordAResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AwsRte53RecordInfo** | Pointer to [**RecordAAwsRte53RecordInfo**](RecordAAwsRte53RecordInfo.md) |  | [optional] 
**CloudInfo** | Pointer to [**RecordACloudInfo**](RecordACloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **int64** | The time of the record creation in Epoch seconds format. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DiscoveredData** | Pointer to [**RecordADiscoveredData**](RecordADiscoveredData.md) |  | [optional] 
**DnsName** | Pointer to **string** | The name for an A record in punycode format. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**Ipv4addr** | Pointer to [**RecordAIpv4addr**](RecordAIpv4addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**RecordAMsAdUserData**](RecordAMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | Name for A record in FQDN format. This value can be in unicode format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**RemoveAssociatedPtr** | Pointer to **bool** | Delete option that indicates whether the associated PTR records should be removed while deleting the specified A record. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 
**Result** | Pointer to [**RecordA**](RecordA.md) |  | [optional] 

## Methods

### NewGetRecordAResponse

`func NewGetRecordAResponse() *GetRecordAResponse`

NewGetRecordAResponse instantiates a new GetRecordAResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordAResponseWithDefaults

`func NewGetRecordAResponseWithDefaults() *GetRecordAResponse`

NewGetRecordAResponseWithDefaults instantiates a new GetRecordAResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordAResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordAResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordAResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordAResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *GetRecordAResponse) GetAwsRte53RecordInfo() RecordAAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *GetRecordAResponse) GetAwsRte53RecordInfoOk() (*RecordAAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *GetRecordAResponse) SetAwsRte53RecordInfo(v RecordAAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *GetRecordAResponse) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRecordAResponse) GetCloudInfo() RecordACloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRecordAResponse) GetCloudInfoOk() (*RecordACloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRecordAResponse) SetCloudInfo(v RecordACloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRecordAResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordAResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordAResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordAResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordAResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *GetRecordAResponse) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GetRecordAResponse) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GetRecordAResponse) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GetRecordAResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *GetRecordAResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetRecordAResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetRecordAResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetRecordAResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *GetRecordAResponse) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *GetRecordAResponse) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *GetRecordAResponse) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *GetRecordAResponse) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *GetRecordAResponse) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *GetRecordAResponse) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *GetRecordAResponse) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *GetRecordAResponse) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *GetRecordAResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRecordAResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRecordAResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRecordAResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *GetRecordAResponse) GetDiscoveredData() RecordADiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *GetRecordAResponse) GetDiscoveredDataOk() (*RecordADiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *GetRecordAResponse) SetDiscoveredData(v RecordADiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *GetRecordAResponse) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordAResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordAResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordAResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordAResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRecordAResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRecordAResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRecordAResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRecordAResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *GetRecordAResponse) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *GetRecordAResponse) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *GetRecordAResponse) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *GetRecordAResponse) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv4addr

`func (o *GetRecordAResponse) GetIpv4addr() RecordAIpv4addr`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *GetRecordAResponse) GetIpv4addrOk() (*RecordAIpv4addr, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *GetRecordAResponse) SetIpv4addr(v RecordAIpv4addr)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *GetRecordAResponse) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *GetRecordAResponse) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *GetRecordAResponse) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *GetRecordAResponse) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *GetRecordAResponse) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordAResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordAResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordAResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordAResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetRecordAResponse) GetMsAdUserData() RecordAMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetRecordAResponse) GetMsAdUserDataOk() (*RecordAMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetRecordAResponse) SetMsAdUserData(v RecordAMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetRecordAResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetRecordAResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordAResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordAResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordAResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *GetRecordAResponse) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *GetRecordAResponse) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *GetRecordAResponse) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *GetRecordAResponse) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetRemoveAssociatedPtr

`func (o *GetRecordAResponse) GetRemoveAssociatedPtr() bool`

GetRemoveAssociatedPtr returns the RemoveAssociatedPtr field if non-nil, zero value otherwise.

### GetRemoveAssociatedPtrOk

`func (o *GetRecordAResponse) GetRemoveAssociatedPtrOk() (*bool, bool)`

GetRemoveAssociatedPtrOk returns a tuple with the RemoveAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAssociatedPtr

`func (o *GetRecordAResponse) SetRemoveAssociatedPtr(v bool)`

SetRemoveAssociatedPtr sets RemoveAssociatedPtr field to given value.

### HasRemoveAssociatedPtr

`func (o *GetRecordAResponse) HasRemoveAssociatedPtr() bool`

HasRemoveAssociatedPtr returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *GetRecordAResponse) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *GetRecordAResponse) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *GetRecordAResponse) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *GetRecordAResponse) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordAResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordAResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordAResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordAResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordAResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordAResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordAResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordAResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordAResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordAResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordAResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordAResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordAResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordAResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordAResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordAResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordAResponse) GetResult() RecordA`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordAResponse) GetResultOk() (*RecordA, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordAResponse) SetResult(v RecordA)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordAResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


