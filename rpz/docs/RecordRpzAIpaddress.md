# RecordRpzAIpaddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the substitute rule. | [optional] 
**Name** | Pointer to **string** | The name for a record in FQDN format. This value cannot be in unicode format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordRpzAIpaddress

`func NewRecordRpzAIpaddress() *RecordRpzAIpaddress`

NewRecordRpzAIpaddress instantiates a new RecordRpzAIpaddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRpzAIpaddressWithDefaults

`func NewRecordRpzAIpaddressWithDefaults() *RecordRpzAIpaddress`

NewRecordRpzAIpaddressWithDefaults instantiates a new RecordRpzAIpaddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordRpzAIpaddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordRpzAIpaddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordRpzAIpaddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordRpzAIpaddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *RecordRpzAIpaddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordRpzAIpaddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordRpzAIpaddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordRpzAIpaddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RecordRpzAIpaddress) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordRpzAIpaddress) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordRpzAIpaddress) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordRpzAIpaddress) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordRpzAIpaddress) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordRpzAIpaddress) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordRpzAIpaddress) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordRpzAIpaddress) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordRpzAIpaddress) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordRpzAIpaddress) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordRpzAIpaddress) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordRpzAIpaddress) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordRpzAIpaddress) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordRpzAIpaddress) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordRpzAIpaddress) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordRpzAIpaddress) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv4addr

`func (o *RecordRpzAIpaddress) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *RecordRpzAIpaddress) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *RecordRpzAIpaddress) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *RecordRpzAIpaddress) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetName

`func (o *RecordRpzAIpaddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordRpzAIpaddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordRpzAIpaddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordRpzAIpaddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRpZone

`func (o *RecordRpzAIpaddress) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *RecordRpzAIpaddress) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *RecordRpzAIpaddress) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *RecordRpzAIpaddress) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetTtl

`func (o *RecordRpzAIpaddress) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordRpzAIpaddress) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordRpzAIpaddress) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordRpzAIpaddress) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordRpzAIpaddress) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordRpzAIpaddress) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordRpzAIpaddress) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordRpzAIpaddress) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordRpzAIpaddress) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordRpzAIpaddress) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordRpzAIpaddress) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordRpzAIpaddress) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordRpzAIpaddress) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordRpzAIpaddress) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordRpzAIpaddress) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordRpzAIpaddress) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


