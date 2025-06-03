# DtcMonitorSnmp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**Community** | Pointer to **string** | The SNMP community string for SNMP authentication. | [optional] 
**Context** | Pointer to **string** | The SNMPv3 context. | [optional] 
**EngineId** | Pointer to **string** | The SNMPv3 engine identifier. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Oids** | Pointer to [**[]DtcMonitorSnmpOids**](DtcMonitorSnmpOids.md) | A list of OIDs for SNMP monitoring. | [optional] 
**Port** | Pointer to **int64** | The port value for SNMP requests. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 
**User** | Pointer to **string** | The SNMPv3 user setting. | [optional] 
**Version** | Pointer to **string** | The SNMP protocol version for the SNMP health check. | [optional] 

## Methods

### NewDtcMonitorSnmp

`func NewDtcMonitorSnmp() *DtcMonitorSnmp`

NewDtcMonitorSnmp instantiates a new DtcMonitorSnmp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcMonitorSnmpWithDefaults

`func NewDtcMonitorSnmpWithDefaults() *DtcMonitorSnmp`

NewDtcMonitorSnmpWithDefaults instantiates a new DtcMonitorSnmp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcMonitorSnmp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcMonitorSnmp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcMonitorSnmp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcMonitorSnmp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcMonitorSnmp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcMonitorSnmp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcMonitorSnmp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcMonitorSnmp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommunity

`func (o *DtcMonitorSnmp) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *DtcMonitorSnmp) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *DtcMonitorSnmp) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *DtcMonitorSnmp) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetContext

`func (o *DtcMonitorSnmp) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DtcMonitorSnmp) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DtcMonitorSnmp) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *DtcMonitorSnmp) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEngineId

`func (o *DtcMonitorSnmp) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *DtcMonitorSnmp) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *DtcMonitorSnmp) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *DtcMonitorSnmp) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetExtattrs

`func (o *DtcMonitorSnmp) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *DtcMonitorSnmp) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *DtcMonitorSnmp) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *DtcMonitorSnmp) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetInterval

`func (o *DtcMonitorSnmp) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DtcMonitorSnmp) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DtcMonitorSnmp) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DtcMonitorSnmp) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *DtcMonitorSnmp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcMonitorSnmp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcMonitorSnmp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcMonitorSnmp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOids

`func (o *DtcMonitorSnmp) GetOids() []DtcMonitorSnmpOids`

GetOids returns the Oids field if non-nil, zero value otherwise.

### GetOidsOk

`func (o *DtcMonitorSnmp) GetOidsOk() (*[]DtcMonitorSnmpOids, bool)`

GetOidsOk returns a tuple with the Oids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOids

`func (o *DtcMonitorSnmp) SetOids(v []DtcMonitorSnmpOids)`

SetOids sets Oids field to given value.

### HasOids

`func (o *DtcMonitorSnmp) HasOids() bool`

HasOids returns a boolean if a field has been set.

### GetPort

`func (o *DtcMonitorSnmp) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DtcMonitorSnmp) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DtcMonitorSnmp) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DtcMonitorSnmp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRetryDown

`func (o *DtcMonitorSnmp) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *DtcMonitorSnmp) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *DtcMonitorSnmp) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *DtcMonitorSnmp) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *DtcMonitorSnmp) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *DtcMonitorSnmp) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *DtcMonitorSnmp) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *DtcMonitorSnmp) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTimeout

`func (o *DtcMonitorSnmp) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DtcMonitorSnmp) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DtcMonitorSnmp) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DtcMonitorSnmp) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUser

`func (o *DtcMonitorSnmp) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DtcMonitorSnmp) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DtcMonitorSnmp) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DtcMonitorSnmp) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *DtcMonitorSnmp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DtcMonitorSnmp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DtcMonitorSnmp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DtcMonitorSnmp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


