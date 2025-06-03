# DtcRecordNaptr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the DTC record is associated. | [optional] 
**Flags** | Pointer to **string** | The flags used to control the interpretation of the fields for an NAPTR record object. Supported values for the flags field are \&quot;U\&quot;, \&quot;S\&quot;, \&quot;P\&quot; and \&quot;A\&quot;. | [optional] 
**Order** | Pointer to **int64** | The order parameter of the NAPTR records. This parameter specifies the order in which the NAPTR rules are applied when multiple rules are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Preference** | Pointer to **int64** | The preference of the NAPTR record. The preference field determines the order the NAPTR records are processed when multiple records with the same order parameter are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Regexp** | Pointer to **string** | The regular expression-based rewriting rule of the NAPTR record. This should be a POSIX compliant regular expression, including the substitution rule and flags. Refer to RFC 2915 for the field syntax details. | [optional] 
**Replacement** | Pointer to **string** | The replacement field of the NAPTR record object. For nonterminal NAPTR records, this field specifies the next domain name to look up. This value can be in unicode format. | [optional] 
**Services** | Pointer to **string** | The services field of the NAPTR record object; maximum 128 characters. The services field contains protocol and service identifiers, such as \&quot;http+E2U\&quot; or \&quot;SIPS+D2T\&quot;. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 

## Methods

### NewDtcRecordNaptr

`func NewDtcRecordNaptr() *DtcRecordNaptr`

NewDtcRecordNaptr instantiates a new DtcRecordNaptr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcRecordNaptrWithDefaults

`func NewDtcRecordNaptrWithDefaults() *DtcRecordNaptr`

NewDtcRecordNaptrWithDefaults instantiates a new DtcRecordNaptr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcRecordNaptr) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcRecordNaptr) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcRecordNaptr) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcRecordNaptr) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcRecordNaptr) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcRecordNaptr) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcRecordNaptr) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcRecordNaptr) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcRecordNaptr) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcRecordNaptr) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcRecordNaptr) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcRecordNaptr) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *DtcRecordNaptr) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *DtcRecordNaptr) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *DtcRecordNaptr) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *DtcRecordNaptr) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetFlags

`func (o *DtcRecordNaptr) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *DtcRecordNaptr) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *DtcRecordNaptr) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *DtcRecordNaptr) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetOrder

`func (o *DtcRecordNaptr) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DtcRecordNaptr) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DtcRecordNaptr) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DtcRecordNaptr) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPreference

`func (o *DtcRecordNaptr) GetPreference() int64`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *DtcRecordNaptr) GetPreferenceOk() (*int64, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *DtcRecordNaptr) SetPreference(v int64)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *DtcRecordNaptr) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetRegexp

`func (o *DtcRecordNaptr) GetRegexp() string`

GetRegexp returns the Regexp field if non-nil, zero value otherwise.

### GetRegexpOk

`func (o *DtcRecordNaptr) GetRegexpOk() (*string, bool)`

GetRegexpOk returns a tuple with the Regexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexp

`func (o *DtcRecordNaptr) SetRegexp(v string)`

SetRegexp sets Regexp field to given value.

### HasRegexp

`func (o *DtcRecordNaptr) HasRegexp() bool`

HasRegexp returns a boolean if a field has been set.

### GetReplacement

`func (o *DtcRecordNaptr) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *DtcRecordNaptr) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *DtcRecordNaptr) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *DtcRecordNaptr) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetServices

`func (o *DtcRecordNaptr) GetServices() string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DtcRecordNaptr) GetServicesOk() (*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DtcRecordNaptr) SetServices(v string)`

SetServices sets Services field to given value.

### HasServices

`func (o *DtcRecordNaptr) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTtl

`func (o *DtcRecordNaptr) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcRecordNaptr) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcRecordNaptr) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcRecordNaptr) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *DtcRecordNaptr) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *DtcRecordNaptr) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *DtcRecordNaptr) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *DtcRecordNaptr) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


