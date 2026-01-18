# RecordCname

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AwsRte53RecordInfo** | Pointer to [**RecordCnameAwsRte53RecordInfo**](RecordCnameAwsRte53RecordInfo.md) |  | [optional] 
**Canonical** | Pointer to **string** | Canonical name in FQDN format. This value can be in unicode format. | [optional] 
**CloudInfo** | Pointer to [**RecordCnameCloudInfo**](RecordCnameCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **int64** | The time of the record creation in Epoch seconds format. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DnsCanonical** | Pointer to **string** | Canonical name in punycode format. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name for the CNAME record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for a CNAME record in FQDN format. This value can be in unicode format. Regular expression search is not supported for unicode values. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordCname

`func NewRecordCname() *RecordCname`

NewRecordCname instantiates a new RecordCname object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCnameWithDefaults

`func NewRecordCnameWithDefaults() *RecordCname`

NewRecordCnameWithDefaults instantiates a new RecordCname object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordCname) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordCname) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordCname) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordCname) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *RecordCname) GetAwsRte53RecordInfo() RecordCnameAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordCname) GetAwsRte53RecordInfoOk() (*RecordCnameAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordCname) SetAwsRte53RecordInfo(v RecordCnameAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordCname) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCanonical

`func (o *RecordCname) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *RecordCname) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *RecordCname) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *RecordCname) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordCname) GetCloudInfo() RecordCnameCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordCname) GetCloudInfoOk() (*RecordCnameCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordCname) SetCloudInfo(v RecordCnameCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordCname) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordCname) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordCname) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordCname) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordCname) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordCname) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordCname) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordCname) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordCname) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordCname) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordCname) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordCname) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordCname) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordCname) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordCname) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordCname) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordCname) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordCname) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordCname) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordCname) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordCname) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordCname) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordCname) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordCname) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordCname) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsCanonical

`func (o *RecordCname) GetDnsCanonical() string`

GetDnsCanonical returns the DnsCanonical field if non-nil, zero value otherwise.

### GetDnsCanonicalOk

`func (o *RecordCname) GetDnsCanonicalOk() (*string, bool)`

GetDnsCanonicalOk returns a tuple with the DnsCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCanonical

`func (o *RecordCname) SetDnsCanonical(v string)`

SetDnsCanonical sets DnsCanonical field to given value.

### HasDnsCanonical

`func (o *RecordCname) HasDnsCanonical() bool`

HasDnsCanonical returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordCname) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordCname) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordCname) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordCname) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordCname) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordCname) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordCname) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordCname) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordCname) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordCname) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordCname) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordCname) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordCname) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordCname) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordCname) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordCname) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordCname) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordCname) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordCname) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordCname) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordCname) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordCname) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordCname) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordCname) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordCname) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordCname) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordCname) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordCname) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordCname) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordCname) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordCname) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordCname) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *RecordCname) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *RecordCname) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *RecordCname) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *RecordCname) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *RecordCname) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordCname) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordCname) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordCname) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordCname) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordCname) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordCname) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordCname) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordCname) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordCname) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordCname) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordCname) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordCname) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordCname) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordCname) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordCname) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


