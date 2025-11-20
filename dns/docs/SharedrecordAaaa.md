# SharedrecordAaaa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for this shared record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if this shared record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name for this shared record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the shared record. | [optional] 
**Name** | Pointer to **string** | Name for this shared record. This value can be in unicode format. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 

## Methods

### NewSharedrecordAaaa

`func NewSharedrecordAaaa() *SharedrecordAaaa`

NewSharedrecordAaaa instantiates a new SharedrecordAaaa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedrecordAaaaWithDefaults

`func NewSharedrecordAaaaWithDefaults() *SharedrecordAaaa`

NewSharedrecordAaaaWithDefaults instantiates a new SharedrecordAaaa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SharedrecordAaaa) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SharedrecordAaaa) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SharedrecordAaaa) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SharedrecordAaaa) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SharedrecordAaaa) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharedrecordAaaa) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharedrecordAaaa) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharedrecordAaaa) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *SharedrecordAaaa) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *SharedrecordAaaa) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *SharedrecordAaaa) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *SharedrecordAaaa) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *SharedrecordAaaa) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *SharedrecordAaaa) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *SharedrecordAaaa) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *SharedrecordAaaa) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *SharedrecordAaaa) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *SharedrecordAaaa) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *SharedrecordAaaa) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *SharedrecordAaaa) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *SharedrecordAaaa) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *SharedrecordAaaa) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *SharedrecordAaaa) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *SharedrecordAaaa) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *SharedrecordAaaa) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *SharedrecordAaaa) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *SharedrecordAaaa) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *SharedrecordAaaa) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv6addr

`func (o *SharedrecordAaaa) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *SharedrecordAaaa) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *SharedrecordAaaa) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *SharedrecordAaaa) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *SharedrecordAaaa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedrecordAaaa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedrecordAaaa) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedrecordAaaa) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *SharedrecordAaaa) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *SharedrecordAaaa) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *SharedrecordAaaa) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *SharedrecordAaaa) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *SharedrecordAaaa) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SharedrecordAaaa) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SharedrecordAaaa) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SharedrecordAaaa) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *SharedrecordAaaa) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *SharedrecordAaaa) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *SharedrecordAaaa) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *SharedrecordAaaa) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


