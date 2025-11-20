# Filterrelayagent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CircuitIdName** | Pointer to **string** | The circuit_id_name of a DHCP relay agent filter object. This filter identifies the circuit between the remote host and the relay agent. For example, the identifier can be the ingress interface number of the circuit access unit, perhaps concatenated with the unit ID number and slot number. Also, the circuit ID can be an ATM virtual circuit ID or cable data virtual circuit ID. | [optional] 
**CircuitIdSubstringLength** | Pointer to **int64** | The circuit ID substring length. | [optional] 
**CircuitIdSubstringOffset** | Pointer to **int64** | The circuit ID substring offset. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of a DHCP relay agent filter object. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IsCircuitId** | Pointer to **string** | The circuit ID matching rule of a DHCP relay agent filter object. The circuit_id value takes effect only if the value is \&quot;MATCHES_VALUE\&quot;. | [optional] 
**IsCircuitIdSubstring** | Pointer to **bool** | Determines if the substring of circuit ID, instead of the full circuit ID, is matched. | [optional] 
**IsRemoteId** | Pointer to **string** | The remote ID matching rule of a DHCP relay agent filter object. The remote_id value takes effect only if the value is Matches_Value. | [optional] 
**IsRemoteIdSubstring** | Pointer to **bool** | Determines if the substring of remote ID, instead of the full remote ID, is matched. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP relay agent filter object. | [optional] 
**RemoteIdName** | Pointer to **string** | The remote ID name attribute of a relay agent filter object. This filter identifies the remote host. The remote ID name can represent many different things such as the caller ID telephone number for a dial-up connection, a user name for logging in to the ISP, a modem ID, etc. When the remote ID name is defined on the relay agent, the DHCP server will have a trusted relationship to identify the remote host. The remote ID name is considered as a trusted identifier. | [optional] 
**RemoteIdSubstringLength** | Pointer to **int64** | The remote ID substring length. | [optional] 
**RemoteIdSubstringOffset** | Pointer to **int64** | The remote ID substring offset. | [optional] 

## Methods

### NewFilterrelayagent

`func NewFilterrelayagent() *Filterrelayagent`

NewFilterrelayagent instantiates a new Filterrelayagent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterrelayagentWithDefaults

`func NewFilterrelayagentWithDefaults() *Filterrelayagent`

NewFilterrelayagentWithDefaults instantiates a new Filterrelayagent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Filterrelayagent) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Filterrelayagent) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Filterrelayagent) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Filterrelayagent) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCircuitIdName

`func (o *Filterrelayagent) GetCircuitIdName() string`

GetCircuitIdName returns the CircuitIdName field if non-nil, zero value otherwise.

### GetCircuitIdNameOk

`func (o *Filterrelayagent) GetCircuitIdNameOk() (*string, bool)`

GetCircuitIdNameOk returns a tuple with the CircuitIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitIdName

`func (o *Filterrelayagent) SetCircuitIdName(v string)`

SetCircuitIdName sets CircuitIdName field to given value.

### HasCircuitIdName

`func (o *Filterrelayagent) HasCircuitIdName() bool`

HasCircuitIdName returns a boolean if a field has been set.

### GetCircuitIdSubstringLength

`func (o *Filterrelayagent) GetCircuitIdSubstringLength() int64`

GetCircuitIdSubstringLength returns the CircuitIdSubstringLength field if non-nil, zero value otherwise.

### GetCircuitIdSubstringLengthOk

`func (o *Filterrelayagent) GetCircuitIdSubstringLengthOk() (*int64, bool)`

GetCircuitIdSubstringLengthOk returns a tuple with the CircuitIdSubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitIdSubstringLength

`func (o *Filterrelayagent) SetCircuitIdSubstringLength(v int64)`

SetCircuitIdSubstringLength sets CircuitIdSubstringLength field to given value.

### HasCircuitIdSubstringLength

`func (o *Filterrelayagent) HasCircuitIdSubstringLength() bool`

HasCircuitIdSubstringLength returns a boolean if a field has been set.

### GetCircuitIdSubstringOffset

`func (o *Filterrelayagent) GetCircuitIdSubstringOffset() int64`

GetCircuitIdSubstringOffset returns the CircuitIdSubstringOffset field if non-nil, zero value otherwise.

### GetCircuitIdSubstringOffsetOk

`func (o *Filterrelayagent) GetCircuitIdSubstringOffsetOk() (*int64, bool)`

GetCircuitIdSubstringOffsetOk returns a tuple with the CircuitIdSubstringOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitIdSubstringOffset

`func (o *Filterrelayagent) SetCircuitIdSubstringOffset(v int64)`

SetCircuitIdSubstringOffset sets CircuitIdSubstringOffset field to given value.

### HasCircuitIdSubstringOffset

`func (o *Filterrelayagent) HasCircuitIdSubstringOffset() bool`

HasCircuitIdSubstringOffset returns a boolean if a field has been set.

### GetComment

`func (o *Filterrelayagent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Filterrelayagent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Filterrelayagent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Filterrelayagent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Filterrelayagent) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Filterrelayagent) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Filterrelayagent) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Filterrelayagent) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Filterrelayagent) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Filterrelayagent) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Filterrelayagent) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Filterrelayagent) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Filterrelayagent) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Filterrelayagent) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Filterrelayagent) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Filterrelayagent) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIsCircuitId

`func (o *Filterrelayagent) GetIsCircuitId() string`

GetIsCircuitId returns the IsCircuitId field if non-nil, zero value otherwise.

### GetIsCircuitIdOk

`func (o *Filterrelayagent) GetIsCircuitIdOk() (*string, bool)`

GetIsCircuitIdOk returns a tuple with the IsCircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircuitId

`func (o *Filterrelayagent) SetIsCircuitId(v string)`

SetIsCircuitId sets IsCircuitId field to given value.

### HasIsCircuitId

`func (o *Filterrelayagent) HasIsCircuitId() bool`

HasIsCircuitId returns a boolean if a field has been set.

### GetIsCircuitIdSubstring

`func (o *Filterrelayagent) GetIsCircuitIdSubstring() bool`

GetIsCircuitIdSubstring returns the IsCircuitIdSubstring field if non-nil, zero value otherwise.

### GetIsCircuitIdSubstringOk

`func (o *Filterrelayagent) GetIsCircuitIdSubstringOk() (*bool, bool)`

GetIsCircuitIdSubstringOk returns a tuple with the IsCircuitIdSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircuitIdSubstring

`func (o *Filterrelayagent) SetIsCircuitIdSubstring(v bool)`

SetIsCircuitIdSubstring sets IsCircuitIdSubstring field to given value.

### HasIsCircuitIdSubstring

`func (o *Filterrelayagent) HasIsCircuitIdSubstring() bool`

HasIsCircuitIdSubstring returns a boolean if a field has been set.

### GetIsRemoteId

`func (o *Filterrelayagent) GetIsRemoteId() string`

GetIsRemoteId returns the IsRemoteId field if non-nil, zero value otherwise.

### GetIsRemoteIdOk

`func (o *Filterrelayagent) GetIsRemoteIdOk() (*string, bool)`

GetIsRemoteIdOk returns a tuple with the IsRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteId

`func (o *Filterrelayagent) SetIsRemoteId(v string)`

SetIsRemoteId sets IsRemoteId field to given value.

### HasIsRemoteId

`func (o *Filterrelayagent) HasIsRemoteId() bool`

HasIsRemoteId returns a boolean if a field has been set.

### GetIsRemoteIdSubstring

`func (o *Filterrelayagent) GetIsRemoteIdSubstring() bool`

GetIsRemoteIdSubstring returns the IsRemoteIdSubstring field if non-nil, zero value otherwise.

### GetIsRemoteIdSubstringOk

`func (o *Filterrelayagent) GetIsRemoteIdSubstringOk() (*bool, bool)`

GetIsRemoteIdSubstringOk returns a tuple with the IsRemoteIdSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIdSubstring

`func (o *Filterrelayagent) SetIsRemoteIdSubstring(v bool)`

SetIsRemoteIdSubstring sets IsRemoteIdSubstring field to given value.

### HasIsRemoteIdSubstring

`func (o *Filterrelayagent) HasIsRemoteIdSubstring() bool`

HasIsRemoteIdSubstring returns a boolean if a field has been set.

### GetName

`func (o *Filterrelayagent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filterrelayagent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filterrelayagent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Filterrelayagent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteIdName

`func (o *Filterrelayagent) GetRemoteIdName() string`

GetRemoteIdName returns the RemoteIdName field if non-nil, zero value otherwise.

### GetRemoteIdNameOk

`func (o *Filterrelayagent) GetRemoteIdNameOk() (*string, bool)`

GetRemoteIdNameOk returns a tuple with the RemoteIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdName

`func (o *Filterrelayagent) SetRemoteIdName(v string)`

SetRemoteIdName sets RemoteIdName field to given value.

### HasRemoteIdName

`func (o *Filterrelayagent) HasRemoteIdName() bool`

HasRemoteIdName returns a boolean if a field has been set.

### GetRemoteIdSubstringLength

`func (o *Filterrelayagent) GetRemoteIdSubstringLength() int64`

GetRemoteIdSubstringLength returns the RemoteIdSubstringLength field if non-nil, zero value otherwise.

### GetRemoteIdSubstringLengthOk

`func (o *Filterrelayagent) GetRemoteIdSubstringLengthOk() (*int64, bool)`

GetRemoteIdSubstringLengthOk returns a tuple with the RemoteIdSubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdSubstringLength

`func (o *Filterrelayagent) SetRemoteIdSubstringLength(v int64)`

SetRemoteIdSubstringLength sets RemoteIdSubstringLength field to given value.

### HasRemoteIdSubstringLength

`func (o *Filterrelayagent) HasRemoteIdSubstringLength() bool`

HasRemoteIdSubstringLength returns a boolean if a field has been set.

### GetRemoteIdSubstringOffset

`func (o *Filterrelayagent) GetRemoteIdSubstringOffset() int64`

GetRemoteIdSubstringOffset returns the RemoteIdSubstringOffset field if non-nil, zero value otherwise.

### GetRemoteIdSubstringOffsetOk

`func (o *Filterrelayagent) GetRemoteIdSubstringOffsetOk() (*int64, bool)`

GetRemoteIdSubstringOffsetOk returns a tuple with the RemoteIdSubstringOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdSubstringOffset

`func (o *Filterrelayagent) SetRemoteIdSubstringOffset(v int64)`

SetRemoteIdSubstringOffset sets RemoteIdSubstringOffset field to given value.

### HasRemoteIdSubstringOffset

`func (o *Filterrelayagent) HasRemoteIdSubstringOffset() bool`

HasRemoteIdSubstringOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


