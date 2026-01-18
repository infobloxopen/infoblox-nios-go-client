# RecordRpzNaptr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Flags** | Pointer to **string** | The flags used to control the interpretation of the fields for a Substitute (NAPTR Record) Rule object. Supported values for the flags field are \&quot;U\&quot;, \&quot;S\&quot;, \&quot;P\&quot; and \&quot;A\&quot;. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for a record in FQDN format. This value cannot be in unicode format. | [optional] 
**Order** | Pointer to **int64** | The order parameter of the Substitute (NAPTR Record) Rule records. This parameter specifies the order in which the NAPTR rules are applied when multiple rules are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Preference** | Pointer to **int64** | The preference of the Substitute (NAPTR Record) Rule record. The preference field determines the order NAPTR records are processed when multiple records with the same order parameter are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Regexp** | Pointer to **string** | The regular expression-based rewriting rule of the Substitute (NAPTR Record) Rule record. This should be a POSIX compliant regular expression, including the substitution rule and flags. Refer to RFC 2915 for the field syntax details. | [optional] 
**Replacement** | Pointer to **string** | The replacement field of the Substitute (NAPTR Record) Rule object. For nonterminal NAPTR records, this field specifies the next domain name to look up. This value can be in unicode format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Services** | Pointer to **string** | The services field of the Substitute (NAPTR Record) Rule object; maximum 128 characters. The services field contains protocol and service identifiers, such as \&quot;http+E2U\&quot; or \&quot;SIPS+D2T\&quot;. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordRpzNaptr

`func NewRecordRpzNaptr() *RecordRpzNaptr`

NewRecordRpzNaptr instantiates a new RecordRpzNaptr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRpzNaptrWithDefaults

`func NewRecordRpzNaptrWithDefaults() *RecordRpzNaptr`

NewRecordRpzNaptrWithDefaults instantiates a new RecordRpzNaptr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordRpzNaptr) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordRpzNaptr) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordRpzNaptr) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordRpzNaptr) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *RecordRpzNaptr) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordRpzNaptr) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordRpzNaptr) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordRpzNaptr) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *RecordRpzNaptr) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordRpzNaptr) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordRpzNaptr) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordRpzNaptr) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordRpzNaptr) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordRpzNaptr) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordRpzNaptr) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordRpzNaptr) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordRpzNaptr) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordRpzNaptr) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordRpzNaptr) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordRpzNaptr) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordRpzNaptr) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordRpzNaptr) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordRpzNaptr) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordRpzNaptr) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFlags

`func (o *RecordRpzNaptr) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *RecordRpzNaptr) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *RecordRpzNaptr) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *RecordRpzNaptr) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordRpzNaptr) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordRpzNaptr) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordRpzNaptr) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordRpzNaptr) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordRpzNaptr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordRpzNaptr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordRpzNaptr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordRpzNaptr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *RecordRpzNaptr) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RecordRpzNaptr) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RecordRpzNaptr) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RecordRpzNaptr) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPreference

`func (o *RecordRpzNaptr) GetPreference() int64`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *RecordRpzNaptr) GetPreferenceOk() (*int64, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *RecordRpzNaptr) SetPreference(v int64)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *RecordRpzNaptr) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetRegexp

`func (o *RecordRpzNaptr) GetRegexp() string`

GetRegexp returns the Regexp field if non-nil, zero value otherwise.

### GetRegexpOk

`func (o *RecordRpzNaptr) GetRegexpOk() (*string, bool)`

GetRegexpOk returns a tuple with the Regexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexp

`func (o *RecordRpzNaptr) SetRegexp(v string)`

SetRegexp sets Regexp field to given value.

### HasRegexp

`func (o *RecordRpzNaptr) HasRegexp() bool`

HasRegexp returns a boolean if a field has been set.

### GetReplacement

`func (o *RecordRpzNaptr) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *RecordRpzNaptr) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *RecordRpzNaptr) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *RecordRpzNaptr) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetRpZone

`func (o *RecordRpzNaptr) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *RecordRpzNaptr) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *RecordRpzNaptr) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *RecordRpzNaptr) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetServices

`func (o *RecordRpzNaptr) GetServices() string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *RecordRpzNaptr) GetServicesOk() (*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *RecordRpzNaptr) SetServices(v string)`

SetServices sets Services field to given value.

### HasServices

`func (o *RecordRpzNaptr) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTtl

`func (o *RecordRpzNaptr) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordRpzNaptr) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordRpzNaptr) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordRpzNaptr) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordRpzNaptr) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordRpzNaptr) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordRpzNaptr) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordRpzNaptr) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordRpzNaptr) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordRpzNaptr) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordRpzNaptr) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordRpzNaptr) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordRpzNaptr) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordRpzNaptr) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordRpzNaptr) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordRpzNaptr) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


