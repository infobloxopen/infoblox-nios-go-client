# GetSharedrecordSrvResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for this shared record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if this shared record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name for this shared record in punycode format. | [optional] [readonly] 
**DnsTarget** | Pointer to **string** | The name for a shared SRV record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Name for this shared record. This value can be in unicode format. | [optional] 
**Port** | Pointer to **int64** | The port of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Priority** | Pointer to **int64** | The priority of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. | [optional] 
**Target** | Pointer to **string** | The target of the shared SRV record in FQDN format. This value can be in unicode format. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Weight** | Pointer to **int64** | The weight of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format. | [optional] 
**Result** | Pointer to [**SharedrecordSrv**](SharedrecordSrv.md) |  | [optional] 

## Methods

### NewGetSharedrecordSrvResponse

`func NewGetSharedrecordSrvResponse() *GetSharedrecordSrvResponse`

NewGetSharedrecordSrvResponse instantiates a new GetSharedrecordSrvResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharedrecordSrvResponseWithDefaults

`func NewGetSharedrecordSrvResponseWithDefaults() *GetSharedrecordSrvResponse`

NewGetSharedrecordSrvResponseWithDefaults instantiates a new GetSharedrecordSrvResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSharedrecordSrvResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSharedrecordSrvResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSharedrecordSrvResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSharedrecordSrvResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSharedrecordSrvResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSharedrecordSrvResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSharedrecordSrvResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSharedrecordSrvResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetSharedrecordSrvResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetSharedrecordSrvResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetSharedrecordSrvResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetSharedrecordSrvResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *GetSharedrecordSrvResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetSharedrecordSrvResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetSharedrecordSrvResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetSharedrecordSrvResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDnsTarget

`func (o *GetSharedrecordSrvResponse) GetDnsTarget() string`

GetDnsTarget returns the DnsTarget field if non-nil, zero value otherwise.

### GetDnsTargetOk

`func (o *GetSharedrecordSrvResponse) GetDnsTargetOk() (*string, bool)`

GetDnsTargetOk returns a tuple with the DnsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTarget

`func (o *GetSharedrecordSrvResponse) SetDnsTarget(v string)`

SetDnsTarget sets DnsTarget field to given value.

### HasDnsTarget

`func (o *GetSharedrecordSrvResponse) HasDnsTarget() bool`

HasDnsTarget returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetSharedrecordSrvResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetSharedrecordSrvResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetSharedrecordSrvResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetSharedrecordSrvResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetSharedrecordSrvResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetSharedrecordSrvResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetSharedrecordSrvResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetSharedrecordSrvResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetSharedrecordSrvResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetSharedrecordSrvResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetSharedrecordSrvResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetSharedrecordSrvResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GetSharedrecordSrvResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSharedrecordSrvResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSharedrecordSrvResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSharedrecordSrvResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *GetSharedrecordSrvResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetSharedrecordSrvResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetSharedrecordSrvResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetSharedrecordSrvResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *GetSharedrecordSrvResponse) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetSharedrecordSrvResponse) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetSharedrecordSrvResponse) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetSharedrecordSrvResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *GetSharedrecordSrvResponse) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *GetSharedrecordSrvResponse) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *GetSharedrecordSrvResponse) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *GetSharedrecordSrvResponse) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTarget

`func (o *GetSharedrecordSrvResponse) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetSharedrecordSrvResponse) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetSharedrecordSrvResponse) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GetSharedrecordSrvResponse) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTtl

`func (o *GetSharedrecordSrvResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetSharedrecordSrvResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetSharedrecordSrvResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetSharedrecordSrvResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetSharedrecordSrvResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetSharedrecordSrvResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetSharedrecordSrvResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetSharedrecordSrvResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetWeight

`func (o *GetSharedrecordSrvResponse) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetSharedrecordSrvResponse) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetSharedrecordSrvResponse) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetSharedrecordSrvResponse) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetResult

`func (o *GetSharedrecordSrvResponse) GetResult() SharedrecordSrv`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSharedrecordSrvResponse) GetResultOk() (*SharedrecordSrv, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSharedrecordSrvResponse) SetResult(v SharedrecordSrv)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSharedrecordSrvResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


