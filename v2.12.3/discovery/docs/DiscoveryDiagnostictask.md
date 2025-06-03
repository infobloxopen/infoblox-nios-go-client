# DiscoveryDiagnostictask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CommunityString** | Pointer to **string** | The SNMP community string of the discovery diagnostic task. | [optional] 
**DebugSnmp** | Pointer to **bool** | The SNMP debug flag of the discovery diagnostic task. | [optional] 
**ForceTest** | Pointer to **bool** | The force test flag of the discovery diagnostic task. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the discovery diagnostic task. | [optional] 
**NetworkView** | Pointer to **string** | The network view name of the discovery diagnostic task. | [optional] 
**StartTime** | Pointer to **int64** | The time when the discovery diagnostic task was started. | [optional] 
**TaskId** | Pointer to **string** | The ID of the discovery diagnostic task. | [optional] 

## Methods

### NewDiscoveryDiagnostictask

`func NewDiscoveryDiagnostictask() *DiscoveryDiagnostictask`

NewDiscoveryDiagnostictask instantiates a new DiscoveryDiagnostictask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDiagnostictaskWithDefaults

`func NewDiscoveryDiagnostictaskWithDefaults() *DiscoveryDiagnostictask`

NewDiscoveryDiagnostictaskWithDefaults instantiates a new DiscoveryDiagnostictask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryDiagnostictask) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryDiagnostictask) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryDiagnostictask) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryDiagnostictask) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCommunityString

`func (o *DiscoveryDiagnostictask) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *DiscoveryDiagnostictask) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *DiscoveryDiagnostictask) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *DiscoveryDiagnostictask) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetDebugSnmp

`func (o *DiscoveryDiagnostictask) GetDebugSnmp() bool`

GetDebugSnmp returns the DebugSnmp field if non-nil, zero value otherwise.

### GetDebugSnmpOk

`func (o *DiscoveryDiagnostictask) GetDebugSnmpOk() (*bool, bool)`

GetDebugSnmpOk returns a tuple with the DebugSnmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugSnmp

`func (o *DiscoveryDiagnostictask) SetDebugSnmp(v bool)`

SetDebugSnmp sets DebugSnmp field to given value.

### HasDebugSnmp

`func (o *DiscoveryDiagnostictask) HasDebugSnmp() bool`

HasDebugSnmp returns a boolean if a field has been set.

### GetForceTest

`func (o *DiscoveryDiagnostictask) GetForceTest() bool`

GetForceTest returns the ForceTest field if non-nil, zero value otherwise.

### GetForceTestOk

`func (o *DiscoveryDiagnostictask) GetForceTestOk() (*bool, bool)`

GetForceTestOk returns a tuple with the ForceTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceTest

`func (o *DiscoveryDiagnostictask) SetForceTest(v bool)`

SetForceTest sets ForceTest field to given value.

### HasForceTest

`func (o *DiscoveryDiagnostictask) HasForceTest() bool`

HasForceTest returns a boolean if a field has been set.

### GetIpAddress

`func (o *DiscoveryDiagnostictask) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DiscoveryDiagnostictask) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DiscoveryDiagnostictask) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DiscoveryDiagnostictask) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetworkView

`func (o *DiscoveryDiagnostictask) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *DiscoveryDiagnostictask) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *DiscoveryDiagnostictask) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *DiscoveryDiagnostictask) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetStartTime

`func (o *DiscoveryDiagnostictask) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DiscoveryDiagnostictask) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DiscoveryDiagnostictask) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DiscoveryDiagnostictask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTaskId

`func (o *DiscoveryDiagnostictask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *DiscoveryDiagnostictask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *DiscoveryDiagnostictask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *DiscoveryDiagnostictask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


