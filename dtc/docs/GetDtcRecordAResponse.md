# GetDtcRecordAResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AutoCreated** | Pointer to **bool** | Flag that indicates whether this record was automatically created by NIOS. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the DTC record is associated. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the domain name. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Result** | Pointer to [**DtcRecordA**](DtcRecordA.md) |  | [optional] 

## Methods

### NewGetDtcRecordAResponse

`func NewGetDtcRecordAResponse() *GetDtcRecordAResponse`

NewGetDtcRecordAResponse instantiates a new GetDtcRecordAResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcRecordAResponseWithDefaults

`func NewGetDtcRecordAResponseWithDefaults() *GetDtcRecordAResponse`

NewGetDtcRecordAResponseWithDefaults instantiates a new GetDtcRecordAResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcRecordAResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcRecordAResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcRecordAResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcRecordAResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreated

`func (o *GetDtcRecordAResponse) GetAutoCreated() bool`

GetAutoCreated returns the AutoCreated field if non-nil, zero value otherwise.

### GetAutoCreatedOk

`func (o *GetDtcRecordAResponse) GetAutoCreatedOk() (*bool, bool)`

GetAutoCreatedOk returns a tuple with the AutoCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreated

`func (o *GetDtcRecordAResponse) SetAutoCreated(v bool)`

SetAutoCreated sets AutoCreated field to given value.

### HasAutoCreated

`func (o *GetDtcRecordAResponse) HasAutoCreated() bool`

HasAutoCreated returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcRecordAResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcRecordAResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcRecordAResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcRecordAResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDtcRecordAResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDtcRecordAResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDtcRecordAResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDtcRecordAResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *GetDtcRecordAResponse) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *GetDtcRecordAResponse) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *GetDtcRecordAResponse) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *GetDtcRecordAResponse) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetIpv4addr

`func (o *GetDtcRecordAResponse) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *GetDtcRecordAResponse) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *GetDtcRecordAResponse) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *GetDtcRecordAResponse) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetTtl

`func (o *GetDtcRecordAResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetDtcRecordAResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetDtcRecordAResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetDtcRecordAResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetDtcRecordAResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetDtcRecordAResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetDtcRecordAResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetDtcRecordAResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcRecordAResponse) GetResult() DtcRecordA`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcRecordAResponse) GetResultOk() (*DtcRecordA, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcRecordAResponse) SetResult(v DtcRecordA)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcRecordAResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


