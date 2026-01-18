# RecordRpzPtr

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
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the substitute rule. | [optional] 
**Name** | Pointer to **string** | The name of the RPZ Substitute (PTR Record) Rule object in FQDN format. | [optional] 
**Ptrdname** | Pointer to **string** | The domain name of the RPZ Substitute (PTR Record) Rule object in FQDN format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordRpzPtr

`func NewRecordRpzPtr() *RecordRpzPtr`

NewRecordRpzPtr instantiates a new RecordRpzPtr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRpzPtrWithDefaults

`func NewRecordRpzPtrWithDefaults() *RecordRpzPtr`

NewRecordRpzPtrWithDefaults instantiates a new RecordRpzPtr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordRpzPtr) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordRpzPtr) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordRpzPtr) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordRpzPtr) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *RecordRpzPtr) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordRpzPtr) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordRpzPtr) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordRpzPtr) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RecordRpzPtr) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordRpzPtr) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordRpzPtr) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordRpzPtr) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordRpzPtr) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordRpzPtr) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordRpzPtr) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordRpzPtr) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordRpzPtr) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordRpzPtr) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordRpzPtr) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordRpzPtr) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordRpzPtr) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordRpzPtr) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordRpzPtr) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordRpzPtr) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv4addr

`func (o *RecordRpzPtr) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *RecordRpzPtr) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *RecordRpzPtr) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *RecordRpzPtr) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6addr

`func (o *RecordRpzPtr) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *RecordRpzPtr) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *RecordRpzPtr) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *RecordRpzPtr) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *RecordRpzPtr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordRpzPtr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordRpzPtr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordRpzPtr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPtrdname

`func (o *RecordRpzPtr) GetPtrdname() string`

GetPtrdname returns the Ptrdname field if non-nil, zero value otherwise.

### GetPtrdnameOk

`func (o *RecordRpzPtr) GetPtrdnameOk() (*string, bool)`

GetPtrdnameOk returns a tuple with the Ptrdname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrdname

`func (o *RecordRpzPtr) SetPtrdname(v string)`

SetPtrdname sets Ptrdname field to given value.

### HasPtrdname

`func (o *RecordRpzPtr) HasPtrdname() bool`

HasPtrdname returns a boolean if a field has been set.

### GetRpZone

`func (o *RecordRpzPtr) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *RecordRpzPtr) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *RecordRpzPtr) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *RecordRpzPtr) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetTtl

`func (o *RecordRpzPtr) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordRpzPtr) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordRpzPtr) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordRpzPtr) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordRpzPtr) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordRpzPtr) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordRpzPtr) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordRpzPtr) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordRpzPtr) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordRpzPtr) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordRpzPtr) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordRpzPtr) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordRpzPtr) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordRpzPtr) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordRpzPtr) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordRpzPtr) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


