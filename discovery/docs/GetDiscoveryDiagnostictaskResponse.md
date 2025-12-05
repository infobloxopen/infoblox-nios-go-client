# GetDiscoveryDiagnostictaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CommunityString** | Pointer to **string** | The SNMP community string of the discovery diagnostic task. | [optional] 
**DebugSnmp** | Pointer to **bool** | The SNMP debug flag of the discovery diagnostic task. | [optional] 
**ForceTest** | Pointer to **bool** | The force test flag of the discovery diagnostic task. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the discovery diagnostic task. | [optional] 
**NetworkView** | Pointer to **string** | The network view name of the discovery diagnostic task. | [optional] 
**StartTime** | Pointer to **int64** | The time when the discovery diagnostic task was started. | [optional] 
**TaskId** | Pointer to **string** | The ID of the discovery diagnostic task. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryDiagnostictask**](DiscoveryDiagnostictask.md) |  | [optional] 

## Methods

### NewGetDiscoveryDiagnostictaskResponse

`func NewGetDiscoveryDiagnostictaskResponse() *GetDiscoveryDiagnostictaskResponse`

NewGetDiscoveryDiagnostictaskResponse instantiates a new GetDiscoveryDiagnostictaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryDiagnostictaskResponseWithDefaults

`func NewGetDiscoveryDiagnostictaskResponseWithDefaults() *GetDiscoveryDiagnostictaskResponse`

NewGetDiscoveryDiagnostictaskResponseWithDefaults instantiates a new GetDiscoveryDiagnostictaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryDiagnostictaskResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryDiagnostictaskResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryDiagnostictaskResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCommunityString

`func (o *GetDiscoveryDiagnostictaskResponse) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *GetDiscoveryDiagnostictaskResponse) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *GetDiscoveryDiagnostictaskResponse) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetDebugSnmp

`func (o *GetDiscoveryDiagnostictaskResponse) GetDebugSnmp() bool`

GetDebugSnmp returns the DebugSnmp field if non-nil, zero value otherwise.

### GetDebugSnmpOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetDebugSnmpOk() (*bool, bool)`

GetDebugSnmpOk returns a tuple with the DebugSnmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugSnmp

`func (o *GetDiscoveryDiagnostictaskResponse) SetDebugSnmp(v bool)`

SetDebugSnmp sets DebugSnmp field to given value.

### HasDebugSnmp

`func (o *GetDiscoveryDiagnostictaskResponse) HasDebugSnmp() bool`

HasDebugSnmp returns a boolean if a field has been set.

### GetForceTest

`func (o *GetDiscoveryDiagnostictaskResponse) GetForceTest() bool`

GetForceTest returns the ForceTest field if non-nil, zero value otherwise.

### GetForceTestOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetForceTestOk() (*bool, bool)`

GetForceTestOk returns a tuple with the ForceTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceTest

`func (o *GetDiscoveryDiagnostictaskResponse) SetForceTest(v bool)`

SetForceTest sets ForceTest field to given value.

### HasForceTest

`func (o *GetDiscoveryDiagnostictaskResponse) HasForceTest() bool`

HasForceTest returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetDiscoveryDiagnostictaskResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetDiscoveryDiagnostictaskResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetDiscoveryDiagnostictaskResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetDiscoveryDiagnostictaskResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetDiscoveryDiagnostictaskResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetDiscoveryDiagnostictaskResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetStartTime

`func (o *GetDiscoveryDiagnostictaskResponse) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetDiscoveryDiagnostictaskResponse) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetDiscoveryDiagnostictaskResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTaskId

`func (o *GetDiscoveryDiagnostictaskResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetDiscoveryDiagnostictaskResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetDiscoveryDiagnostictaskResponse) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetUuid

`func (o *GetDiscoveryDiagnostictaskResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDiscoveryDiagnostictaskResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDiscoveryDiagnostictaskResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryDiagnostictaskResponse) GetResult() DiscoveryDiagnostictask`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryDiagnostictaskResponse) GetResultOk() (*DiscoveryDiagnostictask, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryDiagnostictaskResponse) SetResult(v DiscoveryDiagnostictask)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryDiagnostictaskResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


