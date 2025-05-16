# GetDhcpStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of DHCP objects multiplied by 1000. This is the percentage of the total number of available IP addresses belonging to the object versus the total number of all IP addresses in object. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of the DHCP object. | [optional] [readonly] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the DHCP object. | [optional] [readonly] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in the DHCP object. | [optional] [readonly] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in the DHCP object. | [optional] [readonly] 
**Result** | Pointer to [**DhcpStatistics**](DhcpStatistics.md) |  | [optional] 

## Methods

### NewGetDhcpStatisticsResponse

`func NewGetDhcpStatisticsResponse() *GetDhcpStatisticsResponse`

NewGetDhcpStatisticsResponse instantiates a new GetDhcpStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDhcpStatisticsResponseWithDefaults

`func NewGetDhcpStatisticsResponseWithDefaults() *GetDhcpStatisticsResponse`

NewGetDhcpStatisticsResponseWithDefaults instantiates a new GetDhcpStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDhcpStatisticsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDhcpStatisticsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDhcpStatisticsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDhcpStatisticsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *GetDhcpStatisticsResponse) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *GetDhcpStatisticsResponse) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *GetDhcpStatisticsResponse) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *GetDhcpStatisticsResponse) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *GetDhcpStatisticsResponse) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *GetDhcpStatisticsResponse) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *GetDhcpStatisticsResponse) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *GetDhcpStatisticsResponse) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *GetDhcpStatisticsResponse) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *GetDhcpStatisticsResponse) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *GetDhcpStatisticsResponse) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *GetDhcpStatisticsResponse) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetStaticHosts

`func (o *GetDhcpStatisticsResponse) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *GetDhcpStatisticsResponse) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *GetDhcpStatisticsResponse) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *GetDhcpStatisticsResponse) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetTotalHosts

`func (o *GetDhcpStatisticsResponse) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *GetDhcpStatisticsResponse) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *GetDhcpStatisticsResponse) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *GetDhcpStatisticsResponse) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetResult

`func (o *GetDhcpStatisticsResponse) GetResult() DhcpStatistics`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDhcpStatisticsResponse) GetResultOk() (*DhcpStatistics, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDhcpStatisticsResponse) SetResult(v DhcpStatistics)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDhcpStatisticsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


