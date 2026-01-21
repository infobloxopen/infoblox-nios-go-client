# GetDtcMonitorSnmpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**Community** | Pointer to **string** | The SNMP community string for SNMP authentication. | [optional] 
**Context** | Pointer to **string** | The SNMPv3 context. | [optional] 
**EngineId** | Pointer to **string** | The SNMPv3 engine identifier. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Oids** | Pointer to [**[]DtcMonitorSnmpOids**](DtcMonitorSnmpOids.md) | A list of OIDs for SNMP monitoring. | [optional] 
**Port** | Pointer to **int64** | The port value for SNMP requests. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 
**User** | Pointer to **string** | The SNMPv3 user setting. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Version** | Pointer to **string** | The SNMP protocol version for the SNMP health check. | [optional] 
**Result** | Pointer to [**DtcMonitorSnmp**](DtcMonitorSnmp.md) |  | [optional] 

## Methods

### NewGetDtcMonitorSnmpResponse

`func NewGetDtcMonitorSnmpResponse() *GetDtcMonitorSnmpResponse`

NewGetDtcMonitorSnmpResponse instantiates a new GetDtcMonitorSnmpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcMonitorSnmpResponseWithDefaults

`func NewGetDtcMonitorSnmpResponseWithDefaults() *GetDtcMonitorSnmpResponse`

NewGetDtcMonitorSnmpResponseWithDefaults instantiates a new GetDtcMonitorSnmpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcMonitorSnmpResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcMonitorSnmpResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcMonitorSnmpResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcMonitorSnmpResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcMonitorSnmpResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcMonitorSnmpResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcMonitorSnmpResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcMonitorSnmpResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommunity

`func (o *GetDtcMonitorSnmpResponse) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *GetDtcMonitorSnmpResponse) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *GetDtcMonitorSnmpResponse) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *GetDtcMonitorSnmpResponse) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetContext

`func (o *GetDtcMonitorSnmpResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetDtcMonitorSnmpResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetDtcMonitorSnmpResponse) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetDtcMonitorSnmpResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEngineId

`func (o *GetDtcMonitorSnmpResponse) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *GetDtcMonitorSnmpResponse) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *GetDtcMonitorSnmpResponse) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *GetDtcMonitorSnmpResponse) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetDtcMonitorSnmpResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetDtcMonitorSnmpResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetDtcMonitorSnmpResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetDtcMonitorSnmpResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetDtcMonitorSnmpResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetDtcMonitorSnmpResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetDtcMonitorSnmpResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetDtcMonitorSnmpResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDtcMonitorSnmpResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDtcMonitorSnmpResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDtcMonitorSnmpResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDtcMonitorSnmpResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterval

`func (o *GetDtcMonitorSnmpResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetDtcMonitorSnmpResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetDtcMonitorSnmpResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetDtcMonitorSnmpResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *GetDtcMonitorSnmpResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcMonitorSnmpResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcMonitorSnmpResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcMonitorSnmpResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOids

`func (o *GetDtcMonitorSnmpResponse) GetOids() []DtcMonitorSnmpOids`

GetOids returns the Oids field if non-nil, zero value otherwise.

### GetOidsOk

`func (o *GetDtcMonitorSnmpResponse) GetOidsOk() (*[]DtcMonitorSnmpOids, bool)`

GetOidsOk returns a tuple with the Oids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOids

`func (o *GetDtcMonitorSnmpResponse) SetOids(v []DtcMonitorSnmpOids)`

SetOids sets Oids field to given value.

### HasOids

`func (o *GetDtcMonitorSnmpResponse) HasOids() bool`

HasOids returns a boolean if a field has been set.

### GetPort

`func (o *GetDtcMonitorSnmpResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetDtcMonitorSnmpResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetDtcMonitorSnmpResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetDtcMonitorSnmpResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRetryDown

`func (o *GetDtcMonitorSnmpResponse) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *GetDtcMonitorSnmpResponse) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *GetDtcMonitorSnmpResponse) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *GetDtcMonitorSnmpResponse) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *GetDtcMonitorSnmpResponse) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *GetDtcMonitorSnmpResponse) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *GetDtcMonitorSnmpResponse) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *GetDtcMonitorSnmpResponse) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTimeout

`func (o *GetDtcMonitorSnmpResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetDtcMonitorSnmpResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetDtcMonitorSnmpResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetDtcMonitorSnmpResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUser

`func (o *GetDtcMonitorSnmpResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetDtcMonitorSnmpResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetDtcMonitorSnmpResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GetDtcMonitorSnmpResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUuid

`func (o *GetDtcMonitorSnmpResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDtcMonitorSnmpResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDtcMonitorSnmpResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDtcMonitorSnmpResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *GetDtcMonitorSnmpResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetDtcMonitorSnmpResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetDtcMonitorSnmpResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetDtcMonitorSnmpResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcMonitorSnmpResponse) GetResult() DtcMonitorSnmp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcMonitorSnmpResponse) GetResultOk() (*DtcMonitorSnmp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcMonitorSnmpResponse) SetResult(v DtcMonitorSnmp)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcMonitorSnmpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


