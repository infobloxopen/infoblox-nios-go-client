# RecordPtr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AwsRte53RecordInfo** | Pointer to [**RecordPtrAwsRte53RecordInfo**](RecordPtrAwsRte53RecordInfo.md) |  | [optional] 
**CloudInfo** | Pointer to [**RecordPtrCloudInfo**](RecordPtrCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **int64** | The time of the record creation in Epoch seconds format. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DiscoveredData** | Pointer to [**RecordPtrDiscoveredData**](RecordPtrDiscoveredData.md) |  | [optional] 
**DnsName** | Pointer to **string** | The name for a DNS PTR record in punycode format. | [optional] [readonly] 
**DnsPtrdname** | Pointer to **string** | The domain name of the DNS PTR record in punycode format. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**Ipv4addr** | Pointer to [**RecordPtrIpv4addr**](RecordPtrIpv4addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**Ipv6addr** | Pointer to [**RecordPtrIpv6addr**](RecordPtrIpv6addr.md) |  | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**RecordPtrMsAdUserData**](RecordPtrMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the DNS PTR record in FQDN format. | [optional] 
**Ptrdname** | Pointer to **string** | The domain name of the DNS PTR record in FQDN format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, that the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | Name of the DNS View in which the record resides, for example \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. For example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordPtr

`func NewRecordPtr() *RecordPtr`

NewRecordPtr instantiates a new RecordPtr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordPtrWithDefaults

`func NewRecordPtrWithDefaults() *RecordPtr`

NewRecordPtrWithDefaults instantiates a new RecordPtr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordPtr) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordPtr) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordPtr) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordPtr) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *RecordPtr) GetAwsRte53RecordInfo() RecordPtrAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordPtr) GetAwsRte53RecordInfoOk() (*RecordPtrAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordPtr) SetAwsRte53RecordInfo(v RecordPtrAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordPtr) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordPtr) GetCloudInfo() RecordPtrCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordPtr) GetCloudInfoOk() (*RecordPtrCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordPtr) SetCloudInfo(v RecordPtrCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordPtr) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordPtr) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordPtr) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordPtr) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordPtr) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordPtr) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordPtr) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordPtr) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordPtr) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordPtr) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordPtr) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordPtr) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordPtr) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordPtr) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordPtr) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordPtr) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordPtr) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordPtr) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordPtr) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordPtr) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordPtr) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordPtr) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordPtr) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordPtr) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordPtr) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *RecordPtr) GetDiscoveredData() RecordPtrDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *RecordPtr) GetDiscoveredDataOk() (*RecordPtrDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *RecordPtr) SetDiscoveredData(v RecordPtrDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *RecordPtr) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordPtr) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordPtr) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordPtr) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordPtr) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDnsPtrdname

`func (o *RecordPtr) GetDnsPtrdname() string`

GetDnsPtrdname returns the DnsPtrdname field if non-nil, zero value otherwise.

### GetDnsPtrdnameOk

`func (o *RecordPtr) GetDnsPtrdnameOk() (*string, bool)`

GetDnsPtrdnameOk returns a tuple with the DnsPtrdname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPtrdname

`func (o *RecordPtr) SetDnsPtrdname(v string)`

SetDnsPtrdname sets DnsPtrdname field to given value.

### HasDnsPtrdname

`func (o *RecordPtr) HasDnsPtrdname() bool`

HasDnsPtrdname returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordPtr) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordPtr) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordPtr) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordPtr) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordPtr) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordPtr) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordPtr) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordPtr) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv4addr

`func (o *RecordPtr) GetIpv4addr() RecordPtrIpv4addr`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *RecordPtr) GetIpv4addrOk() (*RecordPtrIpv4addr, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *RecordPtr) SetIpv4addr(v RecordPtrIpv4addr)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *RecordPtr) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *RecordPtr) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *RecordPtr) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *RecordPtr) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *RecordPtr) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetIpv6addr

`func (o *RecordPtr) GetIpv6addr() RecordPtrIpv6addr`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *RecordPtr) GetIpv6addrOk() (*RecordPtrIpv6addr, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *RecordPtr) SetIpv6addr(v RecordPtrIpv6addr)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *RecordPtr) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordPtr) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordPtr) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordPtr) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordPtr) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *RecordPtr) GetMsAdUserData() RecordPtrMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *RecordPtr) GetMsAdUserDataOk() (*RecordPtrMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *RecordPtr) SetMsAdUserData(v RecordPtrMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *RecordPtr) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *RecordPtr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordPtr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordPtr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordPtr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPtrdname

`func (o *RecordPtr) GetPtrdname() string`

GetPtrdname returns the Ptrdname field if non-nil, zero value otherwise.

### GetPtrdnameOk

`func (o *RecordPtr) GetPtrdnameOk() (*string, bool)`

GetPtrdnameOk returns a tuple with the Ptrdname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrdname

`func (o *RecordPtr) SetPtrdname(v string)`

SetPtrdname sets Ptrdname field to given value.

### HasPtrdname

`func (o *RecordPtr) HasPtrdname() bool`

HasPtrdname returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordPtr) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordPtr) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordPtr) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordPtr) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *RecordPtr) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *RecordPtr) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *RecordPtr) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *RecordPtr) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *RecordPtr) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordPtr) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordPtr) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordPtr) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordPtr) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordPtr) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordPtr) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordPtr) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordPtr) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordPtr) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordPtr) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordPtr) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordPtr) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordPtr) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordPtr) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordPtr) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


