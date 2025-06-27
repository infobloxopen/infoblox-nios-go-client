# DtcMonitorTcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Port** | Pointer to **int64** | The port value for TCP requests. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 

## Methods

### NewDtcMonitorTcp

`func NewDtcMonitorTcp() *DtcMonitorTcp`

NewDtcMonitorTcp instantiates a new DtcMonitorTcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcMonitorTcpWithDefaults

`func NewDtcMonitorTcpWithDefaults() *DtcMonitorTcp`

NewDtcMonitorTcpWithDefaults instantiates a new DtcMonitorTcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcMonitorTcp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcMonitorTcp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcMonitorTcp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcMonitorTcp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcMonitorTcp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcMonitorTcp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcMonitorTcp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcMonitorTcp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DtcMonitorTcp) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DtcMonitorTcp) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DtcMonitorTcp) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DtcMonitorTcp) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterval

`func (o *DtcMonitorTcp) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DtcMonitorTcp) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DtcMonitorTcp) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DtcMonitorTcp) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *DtcMonitorTcp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcMonitorTcp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcMonitorTcp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcMonitorTcp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DtcMonitorTcp) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DtcMonitorTcp) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DtcMonitorTcp) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DtcMonitorTcp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRetryDown

`func (o *DtcMonitorTcp) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *DtcMonitorTcp) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *DtcMonitorTcp) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *DtcMonitorTcp) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *DtcMonitorTcp) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *DtcMonitorTcp) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *DtcMonitorTcp) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *DtcMonitorTcp) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTimeout

`func (o *DtcMonitorTcp) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DtcMonitorTcp) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DtcMonitorTcp) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DtcMonitorTcp) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


