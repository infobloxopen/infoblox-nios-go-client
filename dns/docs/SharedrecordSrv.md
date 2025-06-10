# SharedrecordSrv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this shared record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if this shared record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name for this shared record in punycode format. | [optional] [readonly] 
**DnsTarget** | Pointer to **string** | The name for a shared SRV record in punycode format. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Name for this shared record. This value can be in unicode format. | [optional] 
**Port** | Pointer to **int64** | The port of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Priority** | Pointer to **int64** | The priority of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. | [optional] 
**Target** | Pointer to **string** | The target of the shared SRV record in FQDN format. This value can be in unicode format. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Weight** | Pointer to **int64** | The weight of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 

## Methods

### NewSharedrecordSrv

`func NewSharedrecordSrv() *SharedrecordSrv`

NewSharedrecordSrv instantiates a new SharedrecordSrv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedrecordSrvWithDefaults

`func NewSharedrecordSrvWithDefaults() *SharedrecordSrv`

NewSharedrecordSrvWithDefaults instantiates a new SharedrecordSrv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SharedrecordSrv) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SharedrecordSrv) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SharedrecordSrv) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SharedrecordSrv) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SharedrecordSrv) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharedrecordSrv) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharedrecordSrv) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharedrecordSrv) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *SharedrecordSrv) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *SharedrecordSrv) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *SharedrecordSrv) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *SharedrecordSrv) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *SharedrecordSrv) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *SharedrecordSrv) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *SharedrecordSrv) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *SharedrecordSrv) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDnsTarget

`func (o *SharedrecordSrv) GetDnsTarget() string`

GetDnsTarget returns the DnsTarget field if non-nil, zero value otherwise.

### GetDnsTargetOk

`func (o *SharedrecordSrv) GetDnsTargetOk() (*string, bool)`

GetDnsTargetOk returns a tuple with the DnsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTarget

`func (o *SharedrecordSrv) SetDnsTarget(v string)`

SetDnsTarget sets DnsTarget field to given value.

### HasDnsTarget

`func (o *SharedrecordSrv) HasDnsTarget() bool`

HasDnsTarget returns a boolean if a field has been set.

### GetExtattrs

`func (o *SharedrecordSrv) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *SharedrecordSrv) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *SharedrecordSrv) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *SharedrecordSrv) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *SharedrecordSrv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedrecordSrv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedrecordSrv) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedrecordSrv) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *SharedrecordSrv) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SharedrecordSrv) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SharedrecordSrv) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SharedrecordSrv) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *SharedrecordSrv) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SharedrecordSrv) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SharedrecordSrv) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SharedrecordSrv) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *SharedrecordSrv) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *SharedrecordSrv) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *SharedrecordSrv) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *SharedrecordSrv) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTarget

`func (o *SharedrecordSrv) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SharedrecordSrv) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SharedrecordSrv) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SharedrecordSrv) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTtl

`func (o *SharedrecordSrv) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SharedrecordSrv) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SharedrecordSrv) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SharedrecordSrv) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *SharedrecordSrv) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *SharedrecordSrv) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *SharedrecordSrv) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *SharedrecordSrv) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetWeight

`func (o *SharedrecordSrv) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SharedrecordSrv) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SharedrecordSrv) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SharedrecordSrv) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


