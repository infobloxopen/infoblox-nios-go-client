# GetSharedrecordMxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this shared record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if this shared record is disabled or not. False means that the record is enabled. | [optional] 
**DnsMailExchanger** | Pointer to **string** | The name of the mail exchanger in punycode format. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name for this shared record in punycode format. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**MailExchanger** | Pointer to **string** | The name of the mail exchanger in FQDN format. This value can be in unicode format. | [optional] 
**Name** | Pointer to **string** | Name for this shared record. This value can be in unicode format. | [optional] 
**Preference** | Pointer to **int64** | The preference value. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Result** | Pointer to [**SharedrecordMx**](SharedrecordMx.md) |  | [optional] 

## Methods

### NewGetSharedrecordMxResponse

`func NewGetSharedrecordMxResponse() *GetSharedrecordMxResponse`

NewGetSharedrecordMxResponse instantiates a new GetSharedrecordMxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharedrecordMxResponseWithDefaults

`func NewGetSharedrecordMxResponseWithDefaults() *GetSharedrecordMxResponse`

NewGetSharedrecordMxResponseWithDefaults instantiates a new GetSharedrecordMxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSharedrecordMxResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSharedrecordMxResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSharedrecordMxResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSharedrecordMxResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSharedrecordMxResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSharedrecordMxResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSharedrecordMxResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSharedrecordMxResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetSharedrecordMxResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetSharedrecordMxResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetSharedrecordMxResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetSharedrecordMxResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsMailExchanger

`func (o *GetSharedrecordMxResponse) GetDnsMailExchanger() string`

GetDnsMailExchanger returns the DnsMailExchanger field if non-nil, zero value otherwise.

### GetDnsMailExchangerOk

`func (o *GetSharedrecordMxResponse) GetDnsMailExchangerOk() (*string, bool)`

GetDnsMailExchangerOk returns a tuple with the DnsMailExchanger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMailExchanger

`func (o *GetSharedrecordMxResponse) SetDnsMailExchanger(v string)`

SetDnsMailExchanger sets DnsMailExchanger field to given value.

### HasDnsMailExchanger

`func (o *GetSharedrecordMxResponse) HasDnsMailExchanger() bool`

HasDnsMailExchanger returns a boolean if a field has been set.

### GetDnsName

`func (o *GetSharedrecordMxResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetSharedrecordMxResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetSharedrecordMxResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetSharedrecordMxResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetSharedrecordMxResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetSharedrecordMxResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetSharedrecordMxResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetSharedrecordMxResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetMailExchanger

`func (o *GetSharedrecordMxResponse) GetMailExchanger() string`

GetMailExchanger returns the MailExchanger field if non-nil, zero value otherwise.

### GetMailExchangerOk

`func (o *GetSharedrecordMxResponse) GetMailExchangerOk() (*string, bool)`

GetMailExchangerOk returns a tuple with the MailExchanger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailExchanger

`func (o *GetSharedrecordMxResponse) SetMailExchanger(v string)`

SetMailExchanger sets MailExchanger field to given value.

### HasMailExchanger

`func (o *GetSharedrecordMxResponse) HasMailExchanger() bool`

HasMailExchanger returns a boolean if a field has been set.

### GetName

`func (o *GetSharedrecordMxResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSharedrecordMxResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSharedrecordMxResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSharedrecordMxResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreference

`func (o *GetSharedrecordMxResponse) GetPreference() int64`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *GetSharedrecordMxResponse) GetPreferenceOk() (*int64, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *GetSharedrecordMxResponse) SetPreference(v int64)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *GetSharedrecordMxResponse) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *GetSharedrecordMxResponse) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *GetSharedrecordMxResponse) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *GetSharedrecordMxResponse) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *GetSharedrecordMxResponse) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *GetSharedrecordMxResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetSharedrecordMxResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetSharedrecordMxResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetSharedrecordMxResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetSharedrecordMxResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetSharedrecordMxResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetSharedrecordMxResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetSharedrecordMxResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetResult

`func (o *GetSharedrecordMxResponse) GetResult() SharedrecordMx`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSharedrecordMxResponse) GetResultOk() (*SharedrecordMx, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSharedrecordMxResponse) SetResult(v SharedrecordMx)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSharedrecordMxResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


