# SharedrecordTxt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this shared record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if this shared record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name for this shared record in punycode format. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Name for this shared record. This value can be in unicode format. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. | [optional] 
**Text** | Pointer to **string** | Text associated with the shared record. It can contain up to 255 bytes per substring and up a total of 512 bytes. To enter leading, trailing or embedded spaces in the text, add quotes (\&quot; \&quot;) around the text to preserve the spaces. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 

## Methods

### NewSharedrecordTxt

`func NewSharedrecordTxt() *SharedrecordTxt`

NewSharedrecordTxt instantiates a new SharedrecordTxt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedrecordTxtWithDefaults

`func NewSharedrecordTxtWithDefaults() *SharedrecordTxt`

NewSharedrecordTxtWithDefaults instantiates a new SharedrecordTxt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SharedrecordTxt) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SharedrecordTxt) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SharedrecordTxt) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SharedrecordTxt) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SharedrecordTxt) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharedrecordTxt) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharedrecordTxt) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharedrecordTxt) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *SharedrecordTxt) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *SharedrecordTxt) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *SharedrecordTxt) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *SharedrecordTxt) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *SharedrecordTxt) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *SharedrecordTxt) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *SharedrecordTxt) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *SharedrecordTxt) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrs

`func (o *SharedrecordTxt) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *SharedrecordTxt) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *SharedrecordTxt) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *SharedrecordTxt) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *SharedrecordTxt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedrecordTxt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedrecordTxt) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedrecordTxt) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *SharedrecordTxt) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *SharedrecordTxt) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *SharedrecordTxt) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *SharedrecordTxt) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetText

`func (o *SharedrecordTxt) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SharedrecordTxt) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SharedrecordTxt) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SharedrecordTxt) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTtl

`func (o *SharedrecordTxt) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SharedrecordTxt) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SharedrecordTxt) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SharedrecordTxt) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *SharedrecordTxt) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *SharedrecordTxt) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *SharedrecordTxt) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *SharedrecordTxt) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


