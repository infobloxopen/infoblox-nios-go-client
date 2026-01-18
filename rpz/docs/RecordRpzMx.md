# RecordRpzMx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**MailExchanger** | Pointer to **string** | Mail exchanger name in FQDN format. This value can be in unicode format. | [optional] 
**Name** | Pointer to **string** | The name for a record in FQDN format. This value cannot be in unicode format. | [optional] 
**Preference** | Pointer to **int64** | Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordRpzMx

`func NewRecordRpzMx() *RecordRpzMx`

NewRecordRpzMx instantiates a new RecordRpzMx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRpzMxWithDefaults

`func NewRecordRpzMxWithDefaults() *RecordRpzMx`

NewRecordRpzMxWithDefaults instantiates a new RecordRpzMx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordRpzMx) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordRpzMx) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordRpzMx) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordRpzMx) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *RecordRpzMx) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordRpzMx) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordRpzMx) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordRpzMx) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RecordRpzMx) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordRpzMx) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordRpzMx) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordRpzMx) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordRpzMx) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordRpzMx) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordRpzMx) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordRpzMx) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordRpzMx) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordRpzMx) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordRpzMx) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordRpzMx) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordRpzMx) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordRpzMx) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordRpzMx) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordRpzMx) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetMailExchanger

`func (o *RecordRpzMx) GetMailExchanger() string`

GetMailExchanger returns the MailExchanger field if non-nil, zero value otherwise.

### GetMailExchangerOk

`func (o *RecordRpzMx) GetMailExchangerOk() (*string, bool)`

GetMailExchangerOk returns a tuple with the MailExchanger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailExchanger

`func (o *RecordRpzMx) SetMailExchanger(v string)`

SetMailExchanger sets MailExchanger field to given value.

### HasMailExchanger

`func (o *RecordRpzMx) HasMailExchanger() bool`

HasMailExchanger returns a boolean if a field has been set.

### GetName

`func (o *RecordRpzMx) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordRpzMx) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordRpzMx) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordRpzMx) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreference

`func (o *RecordRpzMx) GetPreference() int64`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *RecordRpzMx) GetPreferenceOk() (*int64, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *RecordRpzMx) SetPreference(v int64)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *RecordRpzMx) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetRpZone

`func (o *RecordRpzMx) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *RecordRpzMx) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *RecordRpzMx) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *RecordRpzMx) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetTtl

`func (o *RecordRpzMx) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordRpzMx) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordRpzMx) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordRpzMx) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordRpzMx) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordRpzMx) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordRpzMx) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordRpzMx) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordRpzMx) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordRpzMx) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordRpzMx) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordRpzMx) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordRpzMx) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordRpzMx) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordRpzMx) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordRpzMx) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


