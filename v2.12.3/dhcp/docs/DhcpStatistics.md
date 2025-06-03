# DhcpStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP utilization of DHCP objects multiplied by 1000. This is the percentage of the total number of available IP addresses belonging to the object versus the total number of all IP addresses in object. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | A string describing the utilization level of the DHCP object. | [optional] [readonly] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the DHCP object. | [optional] [readonly] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in the DHCP object. | [optional] [readonly] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in the DHCP object. | [optional] [readonly] 

## Methods

### NewDhcpStatistics

`func NewDhcpStatistics() *DhcpStatistics`

NewDhcpStatistics instantiates a new DhcpStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpStatisticsWithDefaults

`func NewDhcpStatisticsWithDefaults() *DhcpStatistics`

NewDhcpStatisticsWithDefaults instantiates a new DhcpStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DhcpStatistics) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DhcpStatistics) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DhcpStatistics) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DhcpStatistics) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *DhcpStatistics) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *DhcpStatistics) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *DhcpStatistics) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *DhcpStatistics) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *DhcpStatistics) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *DhcpStatistics) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *DhcpStatistics) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *DhcpStatistics) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *DhcpStatistics) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *DhcpStatistics) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *DhcpStatistics) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *DhcpStatistics) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetStaticHosts

`func (o *DhcpStatistics) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *DhcpStatistics) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *DhcpStatistics) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *DhcpStatistics) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetTotalHosts

`func (o *DhcpStatistics) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *DhcpStatistics) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *DhcpStatistics) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *DhcpStatistics) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


