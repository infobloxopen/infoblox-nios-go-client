# GetMastergridResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | IP Address or FQDN to use to contact the SGM | [optional] 
**ConnectionDisabled** | Pointer to **bool** | Indicates whether the sub-grid is currently disabled (true) or not (false). | [optional] 
**ConnectionTimestamp** | Pointer to **int64** | Timestamp of the beginning of the current VPN connection with the SGM. | [optional] [readonly] 
**Detached** | Pointer to **bool** | Indicate detached state of subgrid with MGM. | [optional] [readonly] 
**DiscoverySyncDisabled** | Pointer to **bool** | Enable/Disable discovery service. | [optional] 
**Enable** | Pointer to **bool** | Enable/Disable SuperGrid membership feature | [optional] 
**JoinStatus** | Pointer to **string** | Gives the overall status of the Super Grid connectivity | [optional] [readonly] 
**JoinToken** | Pointer to **string** | Sub grid join token value. | [optional] [readonly] 
**JoinTokenExpiration** | Pointer to **int64** | Timestamp when the sub grid join token expires. | [optional] [readonly] 
**Joined** | Pointer to **bool** | Indicates whether the sub-grid is currently joined or not. | [optional] [readonly] 
**LastEvent** | Pointer to **string** | Last event notified by the Super Grid Master. | [optional] [readonly] 
**LastEventDetails** | Pointer to **string** | additional details on the event. | [optional] [readonly] 
**LastSyncTimestamp** | Pointer to **int64** | Timestamp of the last synchronization with the SGM. | [optional] [readonly] 
**MgmCommunicationMode** | Pointer to **string** | MGM Communication Mode. | [optional] [readonly] 
**Port** | Pointer to **int64** | Alternate port number to use when connecting to the VPN endpoint. | [optional] 
**UseMgmtPort** | Pointer to **bool** | Indicates whether to use the management port to connect to the SGM. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
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

### GetDiscoverySyncDisabled

`func (o *GetMastergridResponse) GetDiscoverySyncDisabled() bool`

GetDiscoverySyncDisabled returns the DiscoverySyncDisabled field if non-nil, zero value otherwise.

### GetDiscoverySyncDisabledOk

`func (o *GetMastergridResponse) GetDiscoverySyncDisabledOk() (*bool, bool)`

GetDiscoverySyncDisabledOk returns a tuple with the DiscoverySyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySyncDisabled

`func (o *GetMastergridResponse) SetDiscoverySyncDisabled(v bool)`

SetDiscoverySyncDisabled sets DiscoverySyncDisabled field to given value.

### HasDiscoverySyncDisabled

`func (o *GetMastergridResponse) HasDiscoverySyncDisabled() bool`

HasDiscoverySyncDisabled returns a boolean if a field has been set.

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

### GetJoinStatus

`func (o *GetMastergridResponse) GetJoinStatus() string`

GetJoinStatus returns the JoinStatus field if non-nil, zero value otherwise.

### GetJoinStatusOk

`func (o *GetMastergridResponse) GetJoinStatusOk() (*string, bool)`

GetJoinStatusOk returns a tuple with the JoinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinStatus

`func (o *GetMastergridResponse) SetJoinStatus(v string)`

SetJoinStatus sets JoinStatus field to given value.

### HasJoinStatus

`func (o *GetMastergridResponse) HasJoinStatus() bool`

HasJoinStatus returns a boolean if a field has been set.

### GetJoinToken

`func (o *GetMastergridResponse) GetJoinToken() string`

GetJoinToken returns the JoinToken field if non-nil, zero value otherwise.

### GetJoinTokenOk

`func (o *GetMastergridResponse) GetJoinTokenOk() (*string, bool)`

GetJoinTokenOk returns a tuple with the JoinToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinToken

`func (o *GetMastergridResponse) SetJoinToken(v string)`

SetJoinToken sets JoinToken field to given value.

### HasJoinToken

`func (o *GetMastergridResponse) HasJoinToken() bool`

HasJoinToken returns a boolean if a field has been set.

### GetJoinTokenExpiration

`func (o *GetMastergridResponse) GetJoinTokenExpiration() int64`

GetJoinTokenExpiration returns the JoinTokenExpiration field if non-nil, zero value otherwise.

### GetJoinTokenExpirationOk

`func (o *GetMastergridResponse) GetJoinTokenExpirationOk() (*int64, bool)`

GetJoinTokenExpirationOk returns a tuple with the JoinTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTokenExpiration

`func (o *GetMastergridResponse) SetJoinTokenExpiration(v int64)`

SetJoinTokenExpiration sets JoinTokenExpiration field to given value.

### HasJoinTokenExpiration

`func (o *GetMastergridResponse) HasJoinTokenExpiration() bool`

HasJoinTokenExpiration returns a boolean if a field has been set.

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

### GetMgmCommunicationMode

`func (o *GetMastergridResponse) GetMgmCommunicationMode() string`

GetMgmCommunicationMode returns the MgmCommunicationMode field if non-nil, zero value otherwise.

### GetMgmCommunicationModeOk

`func (o *GetMastergridResponse) GetMgmCommunicationModeOk() (*string, bool)`

GetMgmCommunicationModeOk returns a tuple with the MgmCommunicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmCommunicationMode

`func (o *GetMastergridResponse) SetMgmCommunicationMode(v string)`

SetMgmCommunicationMode sets MgmCommunicationMode field to given value.

### HasMgmCommunicationMode

`func (o *GetMastergridResponse) HasMgmCommunicationMode() bool`

HasMgmCommunicationMode returns a boolean if a field has been set.

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

### GetUuid

`func (o *GetMastergridResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetMastergridResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetMastergridResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetMastergridResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

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


