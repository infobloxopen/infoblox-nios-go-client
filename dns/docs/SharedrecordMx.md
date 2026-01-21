# SharedrecordMx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for this shared record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if this shared record is disabled or not. False means that the record is enabled. | [optional] 
**DnsMailExchanger** | Pointer to **string** | The name of the mail exchanger in punycode format. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name for this shared record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**MailExchanger** | Pointer to **string** | The name of the mail exchanger in FQDN format. This value can be in unicode format. | [optional] 
**Name** | Pointer to **string** | Name for this shared record. This value can be in unicode format. | [optional] 
**Preference** | Pointer to **int64** | The preference value. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewSharedrecordMx

`func NewSharedrecordMx() *SharedrecordMx`

NewSharedrecordMx instantiates a new SharedrecordMx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedrecordMxWithDefaults

`func NewSharedrecordMxWithDefaults() *SharedrecordMx`

NewSharedrecordMxWithDefaults instantiates a new SharedrecordMx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SharedrecordMx) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SharedrecordMx) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SharedrecordMx) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SharedrecordMx) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SharedrecordMx) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharedrecordMx) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharedrecordMx) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharedrecordMx) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *SharedrecordMx) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *SharedrecordMx) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *SharedrecordMx) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *SharedrecordMx) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsMailExchanger

`func (o *SharedrecordMx) GetDnsMailExchanger() string`

GetDnsMailExchanger returns the DnsMailExchanger field if non-nil, zero value otherwise.

### GetDnsMailExchangerOk

`func (o *SharedrecordMx) GetDnsMailExchangerOk() (*string, bool)`

GetDnsMailExchangerOk returns a tuple with the DnsMailExchanger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMailExchanger

`func (o *SharedrecordMx) SetDnsMailExchanger(v string)`

SetDnsMailExchanger sets DnsMailExchanger field to given value.

### HasDnsMailExchanger

`func (o *SharedrecordMx) HasDnsMailExchanger() bool`

HasDnsMailExchanger returns a boolean if a field has been set.

### GetDnsName

`func (o *SharedrecordMx) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *SharedrecordMx) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *SharedrecordMx) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *SharedrecordMx) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *SharedrecordMx) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *SharedrecordMx) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *SharedrecordMx) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *SharedrecordMx) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *SharedrecordMx) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *SharedrecordMx) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *SharedrecordMx) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *SharedrecordMx) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *SharedrecordMx) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *SharedrecordMx) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *SharedrecordMx) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *SharedrecordMx) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetMailExchanger

`func (o *SharedrecordMx) GetMailExchanger() string`

GetMailExchanger returns the MailExchanger field if non-nil, zero value otherwise.

### GetMailExchangerOk

`func (o *SharedrecordMx) GetMailExchangerOk() (*string, bool)`

GetMailExchangerOk returns a tuple with the MailExchanger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailExchanger

`func (o *SharedrecordMx) SetMailExchanger(v string)`

SetMailExchanger sets MailExchanger field to given value.

### HasMailExchanger

`func (o *SharedrecordMx) HasMailExchanger() bool`

HasMailExchanger returns a boolean if a field has been set.

### GetName

`func (o *SharedrecordMx) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedrecordMx) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedrecordMx) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedrecordMx) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreference

`func (o *SharedrecordMx) GetPreference() int64`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *SharedrecordMx) GetPreferenceOk() (*int64, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *SharedrecordMx) SetPreference(v int64)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *SharedrecordMx) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *SharedrecordMx) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *SharedrecordMx) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *SharedrecordMx) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *SharedrecordMx) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *SharedrecordMx) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SharedrecordMx) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SharedrecordMx) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SharedrecordMx) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *SharedrecordMx) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *SharedrecordMx) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *SharedrecordMx) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *SharedrecordMx) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetUuid

`func (o *SharedrecordMx) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SharedrecordMx) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SharedrecordMx) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SharedrecordMx) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


