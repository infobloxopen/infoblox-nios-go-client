# GetFilterrelayagentResponse

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
**Result** | Pointer to [**Filterrelayagent**](Filterrelayagent.md) |  | [optional] 

## Methods

### NewGetFilterrelayagentResponse

`func NewGetFilterrelayagentResponse() *GetFilterrelayagentResponse`

NewGetFilterrelayagentResponse instantiates a new GetFilterrelayagentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFilterrelayagentResponseWithDefaults

`func NewGetFilterrelayagentResponseWithDefaults() *GetFilterrelayagentResponse`

NewGetFilterrelayagentResponseWithDefaults instantiates a new GetFilterrelayagentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFilterrelayagentResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFilterrelayagentResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFilterrelayagentResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFilterrelayagentResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCircuitIdName

`func (o *GetFilterrelayagentResponse) GetCircuitIdName() string`

GetCircuitIdName returns the CircuitIdName field if non-nil, zero value otherwise.

### GetCircuitIdNameOk

`func (o *GetFilterrelayagentResponse) GetCircuitIdNameOk() (*string, bool)`

GetCircuitIdNameOk returns a tuple with the CircuitIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitIdName

`func (o *GetFilterrelayagentResponse) SetCircuitIdName(v string)`

SetCircuitIdName sets CircuitIdName field to given value.

### HasCircuitIdName

`func (o *GetFilterrelayagentResponse) HasCircuitIdName() bool`

HasCircuitIdName returns a boolean if a field has been set.

### GetCircuitIdSubstringLength

`func (o *GetFilterrelayagentResponse) GetCircuitIdSubstringLength() int64`

GetCircuitIdSubstringLength returns the CircuitIdSubstringLength field if non-nil, zero value otherwise.

### GetCircuitIdSubstringLengthOk

`func (o *GetFilterrelayagentResponse) GetCircuitIdSubstringLengthOk() (*int64, bool)`

GetCircuitIdSubstringLengthOk returns a tuple with the CircuitIdSubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitIdSubstringLength

`func (o *GetFilterrelayagentResponse) SetCircuitIdSubstringLength(v int64)`

SetCircuitIdSubstringLength sets CircuitIdSubstringLength field to given value.

### HasCircuitIdSubstringLength

`func (o *GetFilterrelayagentResponse) HasCircuitIdSubstringLength() bool`

HasCircuitIdSubstringLength returns a boolean if a field has been set.

### GetCircuitIdSubstringOffset

`func (o *GetFilterrelayagentResponse) GetCircuitIdSubstringOffset() int64`

GetCircuitIdSubstringOffset returns the CircuitIdSubstringOffset field if non-nil, zero value otherwise.

### GetCircuitIdSubstringOffsetOk

`func (o *GetFilterrelayagentResponse) GetCircuitIdSubstringOffsetOk() (*int64, bool)`

GetCircuitIdSubstringOffsetOk returns a tuple with the CircuitIdSubstringOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitIdSubstringOffset

`func (o *GetFilterrelayagentResponse) SetCircuitIdSubstringOffset(v int64)`

SetCircuitIdSubstringOffset sets CircuitIdSubstringOffset field to given value.

### HasCircuitIdSubstringOffset

`func (o *GetFilterrelayagentResponse) HasCircuitIdSubstringOffset() bool`

HasCircuitIdSubstringOffset returns a boolean if a field has been set.

### GetComment

`func (o *GetFilterrelayagentResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFilterrelayagentResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFilterrelayagentResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFilterrelayagentResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetFilterrelayagentResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetFilterrelayagentResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetFilterrelayagentResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetFilterrelayagentResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetFilterrelayagentResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetFilterrelayagentResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetFilterrelayagentResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetFilterrelayagentResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetFilterrelayagentResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetFilterrelayagentResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetFilterrelayagentResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetFilterrelayagentResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIsCircuitId

`func (o *GetFilterrelayagentResponse) GetIsCircuitId() string`

GetIsCircuitId returns the IsCircuitId field if non-nil, zero value otherwise.

### GetIsCircuitIdOk

`func (o *GetFilterrelayagentResponse) GetIsCircuitIdOk() (*string, bool)`

GetIsCircuitIdOk returns a tuple with the IsCircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircuitId

`func (o *GetFilterrelayagentResponse) SetIsCircuitId(v string)`

SetIsCircuitId sets IsCircuitId field to given value.

### HasIsCircuitId

`func (o *GetFilterrelayagentResponse) HasIsCircuitId() bool`

HasIsCircuitId returns a boolean if a field has been set.

### GetIsCircuitIdSubstring

`func (o *GetFilterrelayagentResponse) GetIsCircuitIdSubstring() bool`

GetIsCircuitIdSubstring returns the IsCircuitIdSubstring field if non-nil, zero value otherwise.

### GetIsCircuitIdSubstringOk

`func (o *GetFilterrelayagentResponse) GetIsCircuitIdSubstringOk() (*bool, bool)`

GetIsCircuitIdSubstringOk returns a tuple with the IsCircuitIdSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircuitIdSubstring

`func (o *GetFilterrelayagentResponse) SetIsCircuitIdSubstring(v bool)`

SetIsCircuitIdSubstring sets IsCircuitIdSubstring field to given value.

### HasIsCircuitIdSubstring

`func (o *GetFilterrelayagentResponse) HasIsCircuitIdSubstring() bool`

HasIsCircuitIdSubstring returns a boolean if a field has been set.

### GetIsRemoteId

`func (o *GetFilterrelayagentResponse) GetIsRemoteId() string`

GetIsRemoteId returns the IsRemoteId field if non-nil, zero value otherwise.

### GetIsRemoteIdOk

`func (o *GetFilterrelayagentResponse) GetIsRemoteIdOk() (*string, bool)`

GetIsRemoteIdOk returns a tuple with the IsRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteId

`func (o *GetFilterrelayagentResponse) SetIsRemoteId(v string)`

SetIsRemoteId sets IsRemoteId field to given value.

### HasIsRemoteId

`func (o *GetFilterrelayagentResponse) HasIsRemoteId() bool`

HasIsRemoteId returns a boolean if a field has been set.

### GetIsRemoteIdSubstring

`func (o *GetFilterrelayagentResponse) GetIsRemoteIdSubstring() bool`

GetIsRemoteIdSubstring returns the IsRemoteIdSubstring field if non-nil, zero value otherwise.

### GetIsRemoteIdSubstringOk

`func (o *GetFilterrelayagentResponse) GetIsRemoteIdSubstringOk() (*bool, bool)`

GetIsRemoteIdSubstringOk returns a tuple with the IsRemoteIdSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIdSubstring

`func (o *GetFilterrelayagentResponse) SetIsRemoteIdSubstring(v bool)`

SetIsRemoteIdSubstring sets IsRemoteIdSubstring field to given value.

### HasIsRemoteIdSubstring

`func (o *GetFilterrelayagentResponse) HasIsRemoteIdSubstring() bool`

HasIsRemoteIdSubstring returns a boolean if a field has been set.

### GetName

`func (o *GetFilterrelayagentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFilterrelayagentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFilterrelayagentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFilterrelayagentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteIdName

`func (o *GetFilterrelayagentResponse) GetRemoteIdName() string`

GetRemoteIdName returns the RemoteIdName field if non-nil, zero value otherwise.

### GetRemoteIdNameOk

`func (o *GetFilterrelayagentResponse) GetRemoteIdNameOk() (*string, bool)`

GetRemoteIdNameOk returns a tuple with the RemoteIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdName

`func (o *GetFilterrelayagentResponse) SetRemoteIdName(v string)`

SetRemoteIdName sets RemoteIdName field to given value.

### HasRemoteIdName

`func (o *GetFilterrelayagentResponse) HasRemoteIdName() bool`

HasRemoteIdName returns a boolean if a field has been set.

### GetRemoteIdSubstringLength

`func (o *GetFilterrelayagentResponse) GetRemoteIdSubstringLength() int64`

GetRemoteIdSubstringLength returns the RemoteIdSubstringLength field if non-nil, zero value otherwise.

### GetRemoteIdSubstringLengthOk

`func (o *GetFilterrelayagentResponse) GetRemoteIdSubstringLengthOk() (*int64, bool)`

GetRemoteIdSubstringLengthOk returns a tuple with the RemoteIdSubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdSubstringLength

`func (o *GetFilterrelayagentResponse) SetRemoteIdSubstringLength(v int64)`

SetRemoteIdSubstringLength sets RemoteIdSubstringLength field to given value.

### HasRemoteIdSubstringLength

`func (o *GetFilterrelayagentResponse) HasRemoteIdSubstringLength() bool`

HasRemoteIdSubstringLength returns a boolean if a field has been set.

### GetRemoteIdSubstringOffset

`func (o *GetFilterrelayagentResponse) GetRemoteIdSubstringOffset() int64`

GetRemoteIdSubstringOffset returns the RemoteIdSubstringOffset field if non-nil, zero value otherwise.

### GetRemoteIdSubstringOffsetOk

`func (o *GetFilterrelayagentResponse) GetRemoteIdSubstringOffsetOk() (*int64, bool)`

GetRemoteIdSubstringOffsetOk returns a tuple with the RemoteIdSubstringOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdSubstringOffset

`func (o *GetFilterrelayagentResponse) SetRemoteIdSubstringOffset(v int64)`

SetRemoteIdSubstringOffset sets RemoteIdSubstringOffset field to given value.

### HasRemoteIdSubstringOffset

`func (o *GetFilterrelayagentResponse) HasRemoteIdSubstringOffset() bool`

HasRemoteIdSubstringOffset returns a boolean if a field has been set.

### GetResult

`func (o *GetFilterrelayagentResponse) GetResult() Filterrelayagent`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFilterrelayagentResponse) GetResultOk() (*Filterrelayagent, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFilterrelayagentResponse) SetResult(v Filterrelayagent)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFilterrelayagentResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


