# GetIpamStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Cidr** | Pointer to **int64** | The network CIDR. | [optional] [readonly] 
**ConflictCount** | Pointer to **int64** | The number of conflicts discovered via network discovery. This attribute is only valid for a Network object. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**IpamStatisticsMsAdUserData**](IpamStatisticsMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network address. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The network view. | [optional] [readonly] 
**UnmanagedCount** | Pointer to **int64** | The number of unmanaged IP addresses as discovered by network discovery. This attribute is only valid for a Network object. | [optional] [readonly] 
**Utilization** | Pointer to **int64** | The network utilization in percentage. | [optional] [readonly] 
**UtilizationUpdate** | Pointer to **int64** | The time that the utilization statistics were updated last. This attribute is only valid for a Network object. For a Network Container object, the return value is undefined. | [optional] [readonly] 
**Result** | Pointer to [**IpamStatistics**](IpamStatistics.md) |  | [optional] 

## Methods

### NewGetIpamStatisticsResponse

`func NewGetIpamStatisticsResponse() *GetIpamStatisticsResponse`

NewGetIpamStatisticsResponse instantiates a new GetIpamStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpamStatisticsResponseWithDefaults

`func NewGetIpamStatisticsResponseWithDefaults() *GetIpamStatisticsResponse`

NewGetIpamStatisticsResponseWithDefaults instantiates a new GetIpamStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpamStatisticsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpamStatisticsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpamStatisticsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpamStatisticsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCidr

`func (o *GetIpamStatisticsResponse) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetIpamStatisticsResponse) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetIpamStatisticsResponse) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetIpamStatisticsResponse) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetConflictCount

`func (o *GetIpamStatisticsResponse) GetConflictCount() int64`

GetConflictCount returns the ConflictCount field if non-nil, zero value otherwise.

### GetConflictCountOk

`func (o *GetIpamStatisticsResponse) GetConflictCountOk() (*int64, bool)`

GetConflictCountOk returns a tuple with the ConflictCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCount

`func (o *GetIpamStatisticsResponse) SetConflictCount(v int64)`

SetConflictCount sets ConflictCount field to given value.

### HasConflictCount

`func (o *GetIpamStatisticsResponse) HasConflictCount() bool`

HasConflictCount returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetIpamStatisticsResponse) GetMsAdUserData() IpamStatisticsMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetIpamStatisticsResponse) GetMsAdUserDataOk() (*IpamStatisticsMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetIpamStatisticsResponse) SetMsAdUserData(v IpamStatisticsMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetIpamStatisticsResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *GetIpamStatisticsResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetIpamStatisticsResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetIpamStatisticsResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetIpamStatisticsResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetIpamStatisticsResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetIpamStatisticsResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetIpamStatisticsResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetIpamStatisticsResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetUnmanagedCount

`func (o *GetIpamStatisticsResponse) GetUnmanagedCount() int64`

GetUnmanagedCount returns the UnmanagedCount field if non-nil, zero value otherwise.

### GetUnmanagedCountOk

`func (o *GetIpamStatisticsResponse) GetUnmanagedCountOk() (*int64, bool)`

GetUnmanagedCountOk returns a tuple with the UnmanagedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedCount

`func (o *GetIpamStatisticsResponse) SetUnmanagedCount(v int64)`

SetUnmanagedCount sets UnmanagedCount field to given value.

### HasUnmanagedCount

`func (o *GetIpamStatisticsResponse) HasUnmanagedCount() bool`

HasUnmanagedCount returns a boolean if a field has been set.

### GetUtilization

`func (o *GetIpamStatisticsResponse) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *GetIpamStatisticsResponse) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *GetIpamStatisticsResponse) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *GetIpamStatisticsResponse) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationUpdate

`func (o *GetIpamStatisticsResponse) GetUtilizationUpdate() int64`

GetUtilizationUpdate returns the UtilizationUpdate field if non-nil, zero value otherwise.

### GetUtilizationUpdateOk

`func (o *GetIpamStatisticsResponse) GetUtilizationUpdateOk() (*int64, bool)`

GetUtilizationUpdateOk returns a tuple with the UtilizationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationUpdate

`func (o *GetIpamStatisticsResponse) SetUtilizationUpdate(v int64)`

SetUtilizationUpdate sets UtilizationUpdate field to given value.

### HasUtilizationUpdate

`func (o *GetIpamStatisticsResponse) HasUtilizationUpdate() bool`

HasUtilizationUpdate returns a boolean if a field has been set.

### GetResult

`func (o *GetIpamStatisticsResponse) GetResult() IpamStatistics`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpamStatisticsResponse) GetResultOk() (*IpamStatistics, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpamStatisticsResponse) SetResult(v IpamStatistics)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpamStatisticsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


