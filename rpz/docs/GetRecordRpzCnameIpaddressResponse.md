# GetRecordRpzCnameIpaddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Canonical** | Pointer to **string** | The canonical name in FQDN format. This value can be in unicode format. | [optional] 
**Comment** | Pointer to **string** | The comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IsIpv4** | Pointer to **bool** | Indicates whether the record is an IPv4 record. If the return value is \&quot;true\&quot;, it is an IPv4 record. Ohterwise, it is an IPv6 record. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for a record in FQDN format. This value cannot be in unicode format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 
**Result** | Pointer to [**RecordRpzCnameIpaddress**](RecordRpzCnameIpaddress.md) |  | [optional] 

## Methods

### NewGetRecordRpzCnameIpaddressResponse

`func NewGetRecordRpzCnameIpaddressResponse() *GetRecordRpzCnameIpaddressResponse`

NewGetRecordRpzCnameIpaddressResponse instantiates a new GetRecordRpzCnameIpaddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordRpzCnameIpaddressResponseWithDefaults

`func NewGetRecordRpzCnameIpaddressResponseWithDefaults() *GetRecordRpzCnameIpaddressResponse`

NewGetRecordRpzCnameIpaddressResponseWithDefaults instantiates a new GetRecordRpzCnameIpaddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordRpzCnameIpaddressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordRpzCnameIpaddressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordRpzCnameIpaddressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCanonical

`func (o *GetRecordRpzCnameIpaddressResponse) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *GetRecordRpzCnameIpaddressResponse) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *GetRecordRpzCnameIpaddressResponse) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordRpzCnameIpaddressResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordRpzCnameIpaddressResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordRpzCnameIpaddressResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetRecordRpzCnameIpaddressResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRecordRpzCnameIpaddressResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRecordRpzCnameIpaddressResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetRecordRpzCnameIpaddressResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetRecordRpzCnameIpaddressResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetRecordRpzCnameIpaddressResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetRecordRpzCnameIpaddressResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetRecordRpzCnameIpaddressResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetRecordRpzCnameIpaddressResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRecordRpzCnameIpaddressResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRecordRpzCnameIpaddressResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRecordRpzCnameIpaddressResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIsIpv4

`func (o *GetRecordRpzCnameIpaddressResponse) GetIsIpv4() bool`

GetIsIpv4 returns the IsIpv4 field if non-nil, zero value otherwise.

### GetIsIpv4Ok

`func (o *GetRecordRpzCnameIpaddressResponse) GetIsIpv4Ok() (*bool, bool)`

GetIsIpv4Ok returns a tuple with the IsIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpv4

`func (o *GetRecordRpzCnameIpaddressResponse) SetIsIpv4(v bool)`

SetIsIpv4 sets IsIpv4 field to given value.

### HasIsIpv4

`func (o *GetRecordRpzCnameIpaddressResponse) HasIsIpv4() bool`

HasIsIpv4 returns a boolean if a field has been set.

### GetName

`func (o *GetRecordRpzCnameIpaddressResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordRpzCnameIpaddressResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordRpzCnameIpaddressResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRpZone

`func (o *GetRecordRpzCnameIpaddressResponse) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *GetRecordRpzCnameIpaddressResponse) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *GetRecordRpzCnameIpaddressResponse) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordRpzCnameIpaddressResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordRpzCnameIpaddressResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordRpzCnameIpaddressResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordRpzCnameIpaddressResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordRpzCnameIpaddressResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordRpzCnameIpaddressResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordRpzCnameIpaddressResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordRpzCnameIpaddressResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordRpzCnameIpaddressResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordRpzCnameIpaddressResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordRpzCnameIpaddressResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordRpzCnameIpaddressResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordRpzCnameIpaddressResponse) GetResult() RecordRpzCnameIpaddress`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordRpzCnameIpaddressResponse) GetResultOk() (*RecordRpzCnameIpaddress, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordRpzCnameIpaddressResponse) SetResult(v RecordRpzCnameIpaddress)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordRpzCnameIpaddressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


