# DtcRecordAaaa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AutoCreated** | Pointer to **bool** | Flag that indicates whether this record was automatically created by NIOS. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the DTC record is associated. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the domain name. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 

## Methods

### NewDtcRecordAaaa

`func NewDtcRecordAaaa() *DtcRecordAaaa`

NewDtcRecordAaaa instantiates a new DtcRecordAaaa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcRecordAaaaWithDefaults

`func NewDtcRecordAaaaWithDefaults() *DtcRecordAaaa`

NewDtcRecordAaaaWithDefaults instantiates a new DtcRecordAaaa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcRecordAaaa) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcRecordAaaa) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcRecordAaaa) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcRecordAaaa) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreated

`func (o *DtcRecordAaaa) GetAutoCreated() bool`

GetAutoCreated returns the AutoCreated field if non-nil, zero value otherwise.

### GetAutoCreatedOk

`func (o *DtcRecordAaaa) GetAutoCreatedOk() (*bool, bool)`

GetAutoCreatedOk returns a tuple with the AutoCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreated

`func (o *DtcRecordAaaa) SetAutoCreated(v bool)`

SetAutoCreated sets AutoCreated field to given value.

### HasAutoCreated

`func (o *DtcRecordAaaa) HasAutoCreated() bool`

HasAutoCreated returns a boolean if a field has been set.

### GetComment

`func (o *DtcRecordAaaa) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcRecordAaaa) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcRecordAaaa) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcRecordAaaa) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcRecordAaaa) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcRecordAaaa) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcRecordAaaa) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcRecordAaaa) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *DtcRecordAaaa) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *DtcRecordAaaa) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *DtcRecordAaaa) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *DtcRecordAaaa) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetIpv6addr

`func (o *DtcRecordAaaa) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *DtcRecordAaaa) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *DtcRecordAaaa) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *DtcRecordAaaa) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetTtl

`func (o *DtcRecordAaaa) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcRecordAaaa) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcRecordAaaa) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcRecordAaaa) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *DtcRecordAaaa) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *DtcRecordAaaa) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *DtcRecordAaaa) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *DtcRecordAaaa) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


