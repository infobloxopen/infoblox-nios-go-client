# DtcRecordSrv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the DTC record is associated. | [optional] 
**Name** | Pointer to **string** | The name for an SRV record in unicode format. | [optional] 
**Port** | Pointer to **int64** | The port of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Priority** | Pointer to **int64** | The priority of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Target** | Pointer to **string** | The target of the SRV record in FQDN format. This value can be in unicode format. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Weight** | Pointer to **int64** | The weight of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 

## Methods

### NewDtcRecordSrv

`func NewDtcRecordSrv() *DtcRecordSrv`

NewDtcRecordSrv instantiates a new DtcRecordSrv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcRecordSrvWithDefaults

`func NewDtcRecordSrvWithDefaults() *DtcRecordSrv`

NewDtcRecordSrvWithDefaults instantiates a new DtcRecordSrv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcRecordSrv) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcRecordSrv) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcRecordSrv) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcRecordSrv) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcRecordSrv) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcRecordSrv) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcRecordSrv) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcRecordSrv) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcRecordSrv) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcRecordSrv) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcRecordSrv) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcRecordSrv) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *DtcRecordSrv) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *DtcRecordSrv) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *DtcRecordSrv) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *DtcRecordSrv) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetName

`func (o *DtcRecordSrv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcRecordSrv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcRecordSrv) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcRecordSrv) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DtcRecordSrv) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DtcRecordSrv) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DtcRecordSrv) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DtcRecordSrv) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *DtcRecordSrv) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtcRecordSrv) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtcRecordSrv) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtcRecordSrv) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTarget

`func (o *DtcRecordSrv) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DtcRecordSrv) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DtcRecordSrv) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DtcRecordSrv) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTtl

`func (o *DtcRecordSrv) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcRecordSrv) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcRecordSrv) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcRecordSrv) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *DtcRecordSrv) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *DtcRecordSrv) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *DtcRecordSrv) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *DtcRecordSrv) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetUuid

`func (o *DtcRecordSrv) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DtcRecordSrv) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DtcRecordSrv) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DtcRecordSrv) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *DtcRecordSrv) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DtcRecordSrv) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DtcRecordSrv) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DtcRecordSrv) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


