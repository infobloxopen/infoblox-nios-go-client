# RecordTxt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AwsRte53RecordInfo** | Pointer to [**RecordTxtAwsRte53RecordInfo**](RecordTxtAwsRte53RecordInfo.md) |  | [optional] 
**CloudInfo** | Pointer to [**RecordTxtCloudInfo**](RecordTxtCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **int64** | The time of the record creation in Epoch seconds format. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name for a TXT record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | Name for the TXT record in FQDN format. This value can be in unicode format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. | [optional] [readonly] 
**Text** | Pointer to **string** | Text associated with the record. It can contain up to 255 bytes per substring, up to a total of 512 bytes. To enter leading, trailing, or embedded spaces in the text, add double quotes (&amp;#92;\&quot; &amp;#92;\&quot;) around the text to preserve the spaces. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordTxt

`func NewRecordTxt() *RecordTxt`

NewRecordTxt instantiates a new RecordTxt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordTxtWithDefaults

`func NewRecordTxtWithDefaults() *RecordTxt`

NewRecordTxtWithDefaults instantiates a new RecordTxt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordTxt) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordTxt) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordTxt) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordTxt) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *RecordTxt) GetAwsRte53RecordInfo() RecordTxtAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordTxt) GetAwsRte53RecordInfoOk() (*RecordTxtAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordTxt) SetAwsRte53RecordInfo(v RecordTxtAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordTxt) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordTxt) GetCloudInfo() RecordTxtCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordTxt) GetCloudInfoOk() (*RecordTxtCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordTxt) SetCloudInfo(v RecordTxtCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordTxt) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordTxt) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordTxt) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordTxt) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordTxt) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordTxt) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordTxt) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordTxt) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordTxt) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordTxt) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordTxt) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordTxt) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordTxt) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordTxt) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordTxt) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordTxt) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordTxt) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordTxt) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordTxt) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordTxt) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordTxt) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordTxt) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordTxt) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordTxt) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordTxt) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordTxt) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordTxt) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordTxt) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordTxt) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordTxt) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordTxt) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordTxt) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordTxt) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordTxt) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordTxt) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordTxt) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordTxt) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordTxt) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordTxt) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordTxt) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordTxt) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordTxt) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordTxt) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordTxt) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordTxt) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordTxt) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordTxt) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordTxt) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordTxt) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordTxt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordTxt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordTxt) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordTxt) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordTxt) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordTxt) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordTxt) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordTxt) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *RecordTxt) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *RecordTxt) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *RecordTxt) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *RecordTxt) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetText

`func (o *RecordTxt) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RecordTxt) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RecordTxt) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RecordTxt) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTtl

`func (o *RecordTxt) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordTxt) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordTxt) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordTxt) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordTxt) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordTxt) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordTxt) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordTxt) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordTxt) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordTxt) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordTxt) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordTxt) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordTxt) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordTxt) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordTxt) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordTxt) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


