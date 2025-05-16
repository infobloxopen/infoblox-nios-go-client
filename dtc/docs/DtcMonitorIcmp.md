# DtcMonitorIcmp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 

## Methods

### NewDtcMonitorIcmp

`func NewDtcMonitorIcmp() *DtcMonitorIcmp`

NewDtcMonitorIcmp instantiates a new DtcMonitorIcmp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcMonitorIcmpWithDefaults

`func NewDtcMonitorIcmpWithDefaults() *DtcMonitorIcmp`

NewDtcMonitorIcmpWithDefaults instantiates a new DtcMonitorIcmp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcMonitorIcmp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcMonitorIcmp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcMonitorIcmp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcMonitorIcmp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcMonitorIcmp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcMonitorIcmp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcMonitorIcmp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcMonitorIcmp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *DtcMonitorIcmp) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *DtcMonitorIcmp) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *DtcMonitorIcmp) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *DtcMonitorIcmp) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetInterval

`func (o *DtcMonitorIcmp) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DtcMonitorIcmp) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DtcMonitorIcmp) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DtcMonitorIcmp) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *DtcMonitorIcmp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcMonitorIcmp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcMonitorIcmp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcMonitorIcmp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetryDown

`func (o *DtcMonitorIcmp) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *DtcMonitorIcmp) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *DtcMonitorIcmp) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *DtcMonitorIcmp) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *DtcMonitorIcmp) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *DtcMonitorIcmp) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *DtcMonitorIcmp) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *DtcMonitorIcmp) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTimeout

`func (o *DtcMonitorIcmp) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DtcMonitorIcmp) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DtcMonitorIcmp) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DtcMonitorIcmp) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


