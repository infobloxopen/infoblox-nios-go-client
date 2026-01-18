# RecordCaa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CaFlag** | Pointer to **int64** | Flag of CAA record. | [optional] 
**CaTag** | Pointer to **string** | Tag of CAA record. | [optional] 
**CaValue** | Pointer to **string** | Value of CAA record | [optional] 
**CloudInfo** | Pointer to [**RecordCaaCloudInfo**](RecordCaaCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name of the CAA record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The CAA record name in FQDN format. This value can be in unicode format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordCaa

`func NewRecordCaa() *RecordCaa`

NewRecordCaa instantiates a new RecordCaa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCaaWithDefaults

`func NewRecordCaaWithDefaults() *RecordCaa`

NewRecordCaaWithDefaults instantiates a new RecordCaa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordCaa) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordCaa) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordCaa) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordCaa) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCaFlag

`func (o *RecordCaa) GetCaFlag() int64`

GetCaFlag returns the CaFlag field if non-nil, zero value otherwise.

### GetCaFlagOk

`func (o *RecordCaa) GetCaFlagOk() (*int64, bool)`

GetCaFlagOk returns a tuple with the CaFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFlag

`func (o *RecordCaa) SetCaFlag(v int64)`

SetCaFlag sets CaFlag field to given value.

### HasCaFlag

`func (o *RecordCaa) HasCaFlag() bool`

HasCaFlag returns a boolean if a field has been set.

### GetCaTag

`func (o *RecordCaa) GetCaTag() string`

GetCaTag returns the CaTag field if non-nil, zero value otherwise.

### GetCaTagOk

`func (o *RecordCaa) GetCaTagOk() (*string, bool)`

GetCaTagOk returns a tuple with the CaTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaTag

`func (o *RecordCaa) SetCaTag(v string)`

SetCaTag sets CaTag field to given value.

### HasCaTag

`func (o *RecordCaa) HasCaTag() bool`

HasCaTag returns a boolean if a field has been set.

### GetCaValue

`func (o *RecordCaa) GetCaValue() string`

GetCaValue returns the CaValue field if non-nil, zero value otherwise.

### GetCaValueOk

`func (o *RecordCaa) GetCaValueOk() (*string, bool)`

GetCaValueOk returns a tuple with the CaValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaValue

`func (o *RecordCaa) SetCaValue(v string)`

SetCaValue sets CaValue field to given value.

### HasCaValue

`func (o *RecordCaa) HasCaValue() bool`

HasCaValue returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordCaa) GetCloudInfo() RecordCaaCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordCaa) GetCloudInfoOk() (*RecordCaaCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordCaa) SetCloudInfo(v RecordCaaCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordCaa) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordCaa) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordCaa) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordCaa) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordCaa) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordCaa) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordCaa) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordCaa) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordCaa) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordCaa) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordCaa) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordCaa) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordCaa) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordCaa) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordCaa) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordCaa) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordCaa) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordCaa) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordCaa) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordCaa) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordCaa) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordCaa) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordCaa) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordCaa) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordCaa) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordCaa) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordCaa) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordCaa) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordCaa) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordCaa) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordCaa) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordCaa) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordCaa) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordCaa) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordCaa) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordCaa) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordCaa) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordCaa) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordCaa) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordCaa) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordCaa) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordCaa) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordCaa) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordCaa) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordCaa) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordCaa) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordCaa) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordCaa) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordCaa) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordCaa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordCaa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordCaa) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordCaa) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordCaa) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordCaa) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordCaa) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordCaa) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetTtl

`func (o *RecordCaa) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordCaa) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordCaa) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordCaa) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordCaa) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordCaa) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordCaa) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordCaa) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordCaa) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordCaa) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordCaa) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordCaa) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordCaa) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordCaa) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordCaa) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordCaa) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


