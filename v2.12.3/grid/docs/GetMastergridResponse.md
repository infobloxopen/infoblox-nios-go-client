# GetMastergridResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The domain name or IP address for the Master Grid. | [optional] 
**ConnectionDisabled** | Pointer to **bool** | Determines if the sub-grid is currently disabled. | [optional] [readonly] 
**ConnectionTimestamp** | Pointer to **int64** | The timestamp that indicates when the connection to the Master Grid was established. | [optional] [readonly] 
**Detached** | Pointer to **bool** | The detached flag for the Master Grid. | [optional] [readonly] 
**Enable** | Pointer to **bool** | Determines if the Master Grid is enabled. | [optional] 
**Joined** | Pointer to **bool** | The flag shows if the Grid has joined the Master Grid. | [optional] [readonly] 
**LastEvent** | Pointer to **string** | The Master Grid&#39;s last event. | [optional] [readonly] 
**LastEventDetails** | Pointer to **string** | The details of the Master Grid&#39;s last event. | [optional] [readonly] 
**LastSyncTimestamp** | Pointer to **int64** | The timestamp or the last synchronization operation with the Master Grid. | [optional] [readonly] 
**Port** | Pointer to **int64** | The Master Grid port to which the Grid connects. | [optional] 
**Status** | Pointer to **string** | The Master Grid&#39;s status. | [optional] [readonly] 
**UseMgmtPort** | Pointer to **bool** | The flag shows if the MGMT port was used to join the Grid. | [optional] [readonly] 
**Result** | Pointer to [**Mastergrid**](Mastergrid.md) |  | [optional] 

## Methods

### NewGetMastergridResponse

`func NewGetMastergridResponse() *GetMastergridResponse`

NewGetMastergridResponse instantiates a new GetMastergridResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMastergridResponseWithDefaults

`func NewGetMastergridResponseWithDefaults() *GetMastergridResponse`

NewGetMastergridResponseWithDefaults instantiates a new GetMastergridResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMastergridResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMastergridResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMastergridResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMastergridResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetMastergridResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetMastergridResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetMastergridResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetMastergridResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConnectionDisabled

`func (o *GetMastergridResponse) GetConnectionDisabled() bool`

GetConnectionDisabled returns the ConnectionDisabled field if non-nil, zero value otherwise.

### GetConnectionDisabledOk

`func (o *GetMastergridResponse) GetConnectionDisabledOk() (*bool, bool)`

GetConnectionDisabledOk returns a tuple with the ConnectionDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDisabled

`func (o *GetMastergridResponse) SetConnectionDisabled(v bool)`

SetConnectionDisabled sets ConnectionDisabled field to given value.

### HasConnectionDisabled

`func (o *GetMastergridResponse) HasConnectionDisabled() bool`

HasConnectionDisabled returns a boolean if a field has been set.

### GetConnectionTimestamp

`func (o *GetMastergridResponse) GetConnectionTimestamp() int64`

GetConnectionTimestamp returns the ConnectionTimestamp field if non-nil, zero value otherwise.

### GetConnectionTimestampOk

`func (o *GetMastergridResponse) GetConnectionTimestampOk() (*int64, bool)`

GetConnectionTimestampOk returns a tuple with the ConnectionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimestamp

`func (o *GetMastergridResponse) SetConnectionTimestamp(v int64)`

SetConnectionTimestamp sets ConnectionTimestamp field to given value.

### HasConnectionTimestamp

`func (o *GetMastergridResponse) HasConnectionTimestamp() bool`

HasConnectionTimestamp returns a boolean if a field has been set.

### GetDetached

`func (o *GetMastergridResponse) GetDetached() bool`

GetDetached returns the Detached field if non-nil, zero value otherwise.

### GetDetachedOk

`func (o *GetMastergridResponse) GetDetachedOk() (*bool, bool)`

GetDetachedOk returns a tuple with the Detached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetached

`func (o *GetMastergridResponse) SetDetached(v bool)`

SetDetached sets Detached field to given value.

### HasDetached

`func (o *GetMastergridResponse) HasDetached() bool`

HasDetached returns a boolean if a field has been set.

### GetEnable

`func (o *GetMastergridResponse) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GetMastergridResponse) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GetMastergridResponse) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GetMastergridResponse) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetJoined

`func (o *GetMastergridResponse) GetJoined() bool`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *GetMastergridResponse) GetJoinedOk() (*bool, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *GetMastergridResponse) SetJoined(v bool)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *GetMastergridResponse) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetLastEvent

`func (o *GetMastergridResponse) GetLastEvent() string`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *GetMastergridResponse) GetLastEventOk() (*string, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *GetMastergridResponse) SetLastEvent(v string)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *GetMastergridResponse) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetLastEventDetails

`func (o *GetMastergridResponse) GetLastEventDetails() string`

GetLastEventDetails returns the LastEventDetails field if non-nil, zero value otherwise.

### GetLastEventDetailsOk

`func (o *GetMastergridResponse) GetLastEventDetailsOk() (*string, bool)`

GetLastEventDetailsOk returns a tuple with the LastEventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventDetails

`func (o *GetMastergridResponse) SetLastEventDetails(v string)`

SetLastEventDetails sets LastEventDetails field to given value.

### HasLastEventDetails

`func (o *GetMastergridResponse) HasLastEventDetails() bool`

HasLastEventDetails returns a boolean if a field has been set.

### GetLastSyncTimestamp

`func (o *GetMastergridResponse) GetLastSyncTimestamp() int64`

GetLastSyncTimestamp returns the LastSyncTimestamp field if non-nil, zero value otherwise.

### GetLastSyncTimestampOk

`func (o *GetMastergridResponse) GetLastSyncTimestampOk() (*int64, bool)`

GetLastSyncTimestampOk returns a tuple with the LastSyncTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTimestamp

`func (o *GetMastergridResponse) SetLastSyncTimestamp(v int64)`

SetLastSyncTimestamp sets LastSyncTimestamp field to given value.

### HasLastSyncTimestamp

`func (o *GetMastergridResponse) HasLastSyncTimestamp() bool`

HasLastSyncTimestamp returns a boolean if a field has been set.

### GetPort

`func (o *GetMastergridResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetMastergridResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetMastergridResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetMastergridResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *GetMastergridResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMastergridResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMastergridResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMastergridResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *GetMastergridResponse) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *GetMastergridResponse) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *GetMastergridResponse) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *GetMastergridResponse) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.

### GetResult

`func (o *GetMastergridResponse) GetResult() Mastergrid`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMastergridResponse) GetResultOk() (*Mastergrid, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMastergridResponse) SetResult(v Mastergrid)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMastergridResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


