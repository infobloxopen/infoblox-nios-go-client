# Mastergrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
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

## Methods

### NewMastergrid

`func NewMastergrid() *Mastergrid`

NewMastergrid instantiates a new Mastergrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMastergridWithDefaults

`func NewMastergridWithDefaults() *Mastergrid`

NewMastergridWithDefaults instantiates a new Mastergrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Mastergrid) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Mastergrid) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Mastergrid) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Mastergrid) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *Mastergrid) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Mastergrid) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Mastergrid) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Mastergrid) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConnectionDisabled

`func (o *Mastergrid) GetConnectionDisabled() bool`

GetConnectionDisabled returns the ConnectionDisabled field if non-nil, zero value otherwise.

### GetConnectionDisabledOk

`func (o *Mastergrid) GetConnectionDisabledOk() (*bool, bool)`

GetConnectionDisabledOk returns a tuple with the ConnectionDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDisabled

`func (o *Mastergrid) SetConnectionDisabled(v bool)`

SetConnectionDisabled sets ConnectionDisabled field to given value.

### HasConnectionDisabled

`func (o *Mastergrid) HasConnectionDisabled() bool`

HasConnectionDisabled returns a boolean if a field has been set.

### GetConnectionTimestamp

`func (o *Mastergrid) GetConnectionTimestamp() int64`

GetConnectionTimestamp returns the ConnectionTimestamp field if non-nil, zero value otherwise.

### GetConnectionTimestampOk

`func (o *Mastergrid) GetConnectionTimestampOk() (*int64, bool)`

GetConnectionTimestampOk returns a tuple with the ConnectionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimestamp

`func (o *Mastergrid) SetConnectionTimestamp(v int64)`

SetConnectionTimestamp sets ConnectionTimestamp field to given value.

### HasConnectionTimestamp

`func (o *Mastergrid) HasConnectionTimestamp() bool`

HasConnectionTimestamp returns a boolean if a field has been set.

### GetDetached

`func (o *Mastergrid) GetDetached() bool`

GetDetached returns the Detached field if non-nil, zero value otherwise.

### GetDetachedOk

`func (o *Mastergrid) GetDetachedOk() (*bool, bool)`

GetDetachedOk returns a tuple with the Detached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetached

`func (o *Mastergrid) SetDetached(v bool)`

SetDetached sets Detached field to given value.

### HasDetached

`func (o *Mastergrid) HasDetached() bool`

HasDetached returns a boolean if a field has been set.

### GetEnable

`func (o *Mastergrid) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Mastergrid) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Mastergrid) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Mastergrid) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetJoined

`func (o *Mastergrid) GetJoined() bool`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *Mastergrid) GetJoinedOk() (*bool, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *Mastergrid) SetJoined(v bool)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *Mastergrid) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetLastEvent

`func (o *Mastergrid) GetLastEvent() string`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *Mastergrid) GetLastEventOk() (*string, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *Mastergrid) SetLastEvent(v string)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *Mastergrid) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetLastEventDetails

`func (o *Mastergrid) GetLastEventDetails() string`

GetLastEventDetails returns the LastEventDetails field if non-nil, zero value otherwise.

### GetLastEventDetailsOk

`func (o *Mastergrid) GetLastEventDetailsOk() (*string, bool)`

GetLastEventDetailsOk returns a tuple with the LastEventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventDetails

`func (o *Mastergrid) SetLastEventDetails(v string)`

SetLastEventDetails sets LastEventDetails field to given value.

### HasLastEventDetails

`func (o *Mastergrid) HasLastEventDetails() bool`

HasLastEventDetails returns a boolean if a field has been set.

### GetLastSyncTimestamp

`func (o *Mastergrid) GetLastSyncTimestamp() int64`

GetLastSyncTimestamp returns the LastSyncTimestamp field if non-nil, zero value otherwise.

### GetLastSyncTimestampOk

`func (o *Mastergrid) GetLastSyncTimestampOk() (*int64, bool)`

GetLastSyncTimestampOk returns a tuple with the LastSyncTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTimestamp

`func (o *Mastergrid) SetLastSyncTimestamp(v int64)`

SetLastSyncTimestamp sets LastSyncTimestamp field to given value.

### HasLastSyncTimestamp

`func (o *Mastergrid) HasLastSyncTimestamp() bool`

HasLastSyncTimestamp returns a boolean if a field has been set.

### GetPort

`func (o *Mastergrid) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Mastergrid) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Mastergrid) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *Mastergrid) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *Mastergrid) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Mastergrid) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Mastergrid) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Mastergrid) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *Mastergrid) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *Mastergrid) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *Mastergrid) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *Mastergrid) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


