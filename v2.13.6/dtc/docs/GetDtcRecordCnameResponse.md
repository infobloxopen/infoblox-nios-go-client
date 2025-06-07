# GetDtcRecordCnameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AutoCreated** | Pointer to **string** | Flag that indicates whether this record was automatically created by NIOS. | [optional] [readonly] 
**Canonical** | Pointer to **string** | The canonical name of the host. | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DnsCanonical** | Pointer to **string** | The canonical name as server by DNS protocol. | [optional] [readonly] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the DTC record is associated. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Result** | Pointer to [**DtcRecordCname**](DtcRecordCname.md) |  | [optional] 

## Methods

### NewGetDtcRecordCnameResponse

`func NewGetDtcRecordCnameResponse() *GetDtcRecordCnameResponse`

NewGetDtcRecordCnameResponse instantiates a new GetDtcRecordCnameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcRecordCnameResponseWithDefaults

`func NewGetDtcRecordCnameResponseWithDefaults() *GetDtcRecordCnameResponse`

NewGetDtcRecordCnameResponseWithDefaults instantiates a new GetDtcRecordCnameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcRecordCnameResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcRecordCnameResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcRecordCnameResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcRecordCnameResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreated

`func (o *GetDtcRecordCnameResponse) GetAutoCreated() string`

GetAutoCreated returns the AutoCreated field if non-nil, zero value otherwise.

### GetAutoCreatedOk

`func (o *GetDtcRecordCnameResponse) GetAutoCreatedOk() (*string, bool)`

GetAutoCreatedOk returns a tuple with the AutoCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreated

`func (o *GetDtcRecordCnameResponse) SetAutoCreated(v string)`

SetAutoCreated sets AutoCreated field to given value.

### HasAutoCreated

`func (o *GetDtcRecordCnameResponse) HasAutoCreated() bool`

HasAutoCreated returns a boolean if a field has been set.

### GetCanonical

`func (o *GetDtcRecordCnameResponse) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *GetDtcRecordCnameResponse) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *GetDtcRecordCnameResponse) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *GetDtcRecordCnameResponse) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcRecordCnameResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcRecordCnameResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcRecordCnameResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcRecordCnameResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDtcRecordCnameResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDtcRecordCnameResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDtcRecordCnameResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDtcRecordCnameResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsCanonical

`func (o *GetDtcRecordCnameResponse) GetDnsCanonical() string`

GetDnsCanonical returns the DnsCanonical field if non-nil, zero value otherwise.

### GetDnsCanonicalOk

`func (o *GetDtcRecordCnameResponse) GetDnsCanonicalOk() (*string, bool)`

GetDnsCanonicalOk returns a tuple with the DnsCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCanonical

`func (o *GetDtcRecordCnameResponse) SetDnsCanonical(v string)`

SetDnsCanonical sets DnsCanonical field to given value.

### HasDnsCanonical

`func (o *GetDtcRecordCnameResponse) HasDnsCanonical() bool`

HasDnsCanonical returns a boolean if a field has been set.

### GetDtcServer

`func (o *GetDtcRecordCnameResponse) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *GetDtcRecordCnameResponse) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *GetDtcRecordCnameResponse) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *GetDtcRecordCnameResponse) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetTtl

`func (o *GetDtcRecordCnameResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetDtcRecordCnameResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetDtcRecordCnameResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetDtcRecordCnameResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetDtcRecordCnameResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetDtcRecordCnameResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetDtcRecordCnameResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetDtcRecordCnameResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcRecordCnameResponse) GetResult() DtcRecordCname`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcRecordCnameResponse) GetResultOk() (*DtcRecordCname, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcRecordCnameResponse) SetResult(v DtcRecordCname)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcRecordCnameResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


