# DtcRecordA

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

## Methods

### NewDtcRecordA

`func NewDtcRecordA() *DtcRecordA`

NewDtcRecordA instantiates a new DtcRecordA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcRecordAWithDefaults

`func NewDtcRecordAWithDefaults() *DtcRecordA`

NewDtcRecordAWithDefaults instantiates a new DtcRecordA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcRecordA) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcRecordA) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcRecordA) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcRecordA) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreated

`func (o *DtcRecordA) GetAutoCreated() bool`

GetAutoCreated returns the AutoCreated field if non-nil, zero value otherwise.

### GetAutoCreatedOk

`func (o *DtcRecordA) GetAutoCreatedOk() (*bool, bool)`

GetAutoCreatedOk returns a tuple with the AutoCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreated

`func (o *DtcRecordA) SetAutoCreated(v bool)`

SetAutoCreated sets AutoCreated field to given value.

### HasAutoCreated

`func (o *DtcRecordA) HasAutoCreated() bool`

HasAutoCreated returns a boolean if a field has been set.

### GetComment

`func (o *DtcRecordA) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcRecordA) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcRecordA) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcRecordA) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcRecordA) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcRecordA) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcRecordA) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcRecordA) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *DtcRecordA) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *DtcRecordA) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *DtcRecordA) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *DtcRecordA) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetIpv4addr

`func (o *DtcRecordA) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *DtcRecordA) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *DtcRecordA) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *DtcRecordA) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetTtl

`func (o *DtcRecordA) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcRecordA) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcRecordA) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcRecordA) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *DtcRecordA) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *DtcRecordA) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *DtcRecordA) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *DtcRecordA) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


