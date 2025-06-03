# RecordRpzCnameIpaddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Canonical** | Pointer to **string** | The canonical name in FQDN format. This value can be in unicode format. | [optional] 
**Comment** | Pointer to **string** | The comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IsIpv4** | Pointer to **bool** | Indicates whether the record is an IPv4 record. If the return value is \&quot;true\&quot;, it is an IPv4 record. Ohterwise, it is an IPv6 record. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for a record in FQDN format. This value cannot be in unicode format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordRpzCnameIpaddress

`func NewRecordRpzCnameIpaddress() *RecordRpzCnameIpaddress`

NewRecordRpzCnameIpaddress instantiates a new RecordRpzCnameIpaddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRpzCnameIpaddressWithDefaults

`func NewRecordRpzCnameIpaddressWithDefaults() *RecordRpzCnameIpaddress`

NewRecordRpzCnameIpaddressWithDefaults instantiates a new RecordRpzCnameIpaddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordRpzCnameIpaddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordRpzCnameIpaddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordRpzCnameIpaddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordRpzCnameIpaddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCanonical

`func (o *RecordRpzCnameIpaddress) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *RecordRpzCnameIpaddress) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *RecordRpzCnameIpaddress) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *RecordRpzCnameIpaddress) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetComment

`func (o *RecordRpzCnameIpaddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordRpzCnameIpaddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordRpzCnameIpaddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordRpzCnameIpaddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RecordRpzCnameIpaddress) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordRpzCnameIpaddress) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordRpzCnameIpaddress) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordRpzCnameIpaddress) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtattrs

`func (o *RecordRpzCnameIpaddress) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *RecordRpzCnameIpaddress) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *RecordRpzCnameIpaddress) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *RecordRpzCnameIpaddress) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIsIpv4

`func (o *RecordRpzCnameIpaddress) GetIsIpv4() bool`

GetIsIpv4 returns the IsIpv4 field if non-nil, zero value otherwise.

### GetIsIpv4Ok

`func (o *RecordRpzCnameIpaddress) GetIsIpv4Ok() (*bool, bool)`

GetIsIpv4Ok returns a tuple with the IsIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpv4

`func (o *RecordRpzCnameIpaddress) SetIsIpv4(v bool)`

SetIsIpv4 sets IsIpv4 field to given value.

### HasIsIpv4

`func (o *RecordRpzCnameIpaddress) HasIsIpv4() bool`

HasIsIpv4 returns a boolean if a field has been set.

### GetName

`func (o *RecordRpzCnameIpaddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordRpzCnameIpaddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordRpzCnameIpaddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordRpzCnameIpaddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRpZone

`func (o *RecordRpzCnameIpaddress) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *RecordRpzCnameIpaddress) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *RecordRpzCnameIpaddress) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *RecordRpzCnameIpaddress) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetTtl

`func (o *RecordRpzCnameIpaddress) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordRpzCnameIpaddress) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordRpzCnameIpaddress) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordRpzCnameIpaddress) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordRpzCnameIpaddress) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordRpzCnameIpaddress) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordRpzCnameIpaddress) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordRpzCnameIpaddress) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordRpzCnameIpaddress) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordRpzCnameIpaddress) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordRpzCnameIpaddress) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordRpzCnameIpaddress) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordRpzCnameIpaddress) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordRpzCnameIpaddress) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordRpzCnameIpaddress) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordRpzCnameIpaddress) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


