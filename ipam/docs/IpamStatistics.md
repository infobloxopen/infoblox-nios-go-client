# IpamStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Cidr** | Pointer to **int64** | The network CIDR. | [optional] [readonly] 
**ConflictCount** | Pointer to **int64** | The number of conflicts discovered via network discovery. This attribute is only valid for a Network object. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**IpamStatisticsMsAdUserData**](IpamStatisticsMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network address. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The network view. | [optional] [readonly] 
**UnmanagedCount** | Pointer to **int64** | The number of unmanaged IP addresses as discovered by network discovery. This attribute is only valid for a Network object. | [optional] [readonly] 
**Utilization** | Pointer to **int64** | The network utilization in percentage. | [optional] [readonly] 
**UtilizationUpdate** | Pointer to **int64** | The time that the utilization statistics were updated last. This attribute is only valid for a Network object. For a Network Container object, the return value is undefined. | [optional] [readonly] 

## Methods

### NewIpamStatistics

`func NewIpamStatistics() *IpamStatistics`

NewIpamStatistics instantiates a new IpamStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamStatisticsWithDefaults

`func NewIpamStatisticsWithDefaults() *IpamStatistics`

NewIpamStatisticsWithDefaults instantiates a new IpamStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *IpamStatistics) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *IpamStatistics) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *IpamStatistics) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *IpamStatistics) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCidr

`func (o *IpamStatistics) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IpamStatistics) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IpamStatistics) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IpamStatistics) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetConflictCount

`func (o *IpamStatistics) GetConflictCount() int64`

GetConflictCount returns the ConflictCount field if non-nil, zero value otherwise.

### GetConflictCountOk

`func (o *IpamStatistics) GetConflictCountOk() (*int64, bool)`

GetConflictCountOk returns a tuple with the ConflictCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCount

`func (o *IpamStatistics) SetConflictCount(v int64)`

SetConflictCount sets ConflictCount field to given value.

### HasConflictCount

`func (o *IpamStatistics) HasConflictCount() bool`

HasConflictCount returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *IpamStatistics) GetMsAdUserData() IpamStatisticsMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *IpamStatistics) GetMsAdUserDataOk() (*IpamStatisticsMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *IpamStatistics) SetMsAdUserData(v IpamStatisticsMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *IpamStatistics) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *IpamStatistics) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IpamStatistics) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IpamStatistics) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IpamStatistics) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *IpamStatistics) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *IpamStatistics) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *IpamStatistics) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *IpamStatistics) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetUnmanagedCount

`func (o *IpamStatistics) GetUnmanagedCount() int64`

GetUnmanagedCount returns the UnmanagedCount field if non-nil, zero value otherwise.

### GetUnmanagedCountOk

`func (o *IpamStatistics) GetUnmanagedCountOk() (*int64, bool)`

GetUnmanagedCountOk returns a tuple with the UnmanagedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedCount

`func (o *IpamStatistics) SetUnmanagedCount(v int64)`

SetUnmanagedCount sets UnmanagedCount field to given value.

### HasUnmanagedCount

`func (o *IpamStatistics) HasUnmanagedCount() bool`

HasUnmanagedCount returns a boolean if a field has been set.

### GetUtilization

`func (o *IpamStatistics) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamStatistics) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamStatistics) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamStatistics) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationUpdate

`func (o *IpamStatistics) GetUtilizationUpdate() int64`

GetUtilizationUpdate returns the UtilizationUpdate field if non-nil, zero value otherwise.

### GetUtilizationUpdateOk

`func (o *IpamStatistics) GetUtilizationUpdateOk() (*int64, bool)`

GetUtilizationUpdateOk returns a tuple with the UtilizationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationUpdate

`func (o *IpamStatistics) SetUtilizationUpdate(v int64)`

SetUtilizationUpdate sets UtilizationUpdate field to given value.

### HasUtilizationUpdate

`func (o *IpamStatistics) HasUtilizationUpdate() bool`

HasUtilizationUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


