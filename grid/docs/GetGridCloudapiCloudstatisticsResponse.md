# GetGridCloudapiCloudstatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllocatedAvailableRatio** | Pointer to **int64** | Ratio of allocated vs. available IPs | [optional] [readonly] 
**AllocatedIpCount** | Pointer to **int64** | Total number of IPs allocated by tenants. | [optional] [readonly] 
**AvailableIpCount** | Pointer to **string** | The total number of IP addresses available to tenants. Only IP addresses in networks that are within a delegation scope are counted. | [optional] [readonly] 
**FixedIpCount** | Pointer to **int64** | The total number of fixed IP addresses currently in use by all tenants in the system. | [optional] [readonly] 
**FloatingIpCount** | Pointer to **int64** | The total number of floating IP addresses currently in use by all tenants in the system. | [optional] [readonly] 
**TenantCount** | Pointer to **int64** | Total number of tenant currently in the system. | [optional] [readonly] 
**TenantIpCount** | Pointer to **int64** | The total number of IP addresses currently in use by all tenants in the system. | [optional] [readonly] 
**TenantVmCount** | Pointer to **int64** | The total number of VMs currently in use by all tenants in the system. | [optional] [readonly] 
**Result** | Pointer to [**GridCloudapiCloudstatistics**](GridCloudapiCloudstatistics.md) |  | [optional] 

## Methods

### NewGetGridCloudapiCloudstatisticsResponse

`func NewGetGridCloudapiCloudstatisticsResponse() *GetGridCloudapiCloudstatisticsResponse`

NewGetGridCloudapiCloudstatisticsResponse instantiates a new GetGridCloudapiCloudstatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridCloudapiCloudstatisticsResponseWithDefaults

`func NewGetGridCloudapiCloudstatisticsResponseWithDefaults() *GetGridCloudapiCloudstatisticsResponse`

NewGetGridCloudapiCloudstatisticsResponseWithDefaults instantiates a new GetGridCloudapiCloudstatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridCloudapiCloudstatisticsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridCloudapiCloudstatisticsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridCloudapiCloudstatisticsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllocatedAvailableRatio

`func (o *GetGridCloudapiCloudstatisticsResponse) GetAllocatedAvailableRatio() int64`

GetAllocatedAvailableRatio returns the AllocatedAvailableRatio field if non-nil, zero value otherwise.

### GetAllocatedAvailableRatioOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetAllocatedAvailableRatioOk() (*int64, bool)`

GetAllocatedAvailableRatioOk returns a tuple with the AllocatedAvailableRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAvailableRatio

`func (o *GetGridCloudapiCloudstatisticsResponse) SetAllocatedAvailableRatio(v int64)`

SetAllocatedAvailableRatio sets AllocatedAvailableRatio field to given value.

### HasAllocatedAvailableRatio

`func (o *GetGridCloudapiCloudstatisticsResponse) HasAllocatedAvailableRatio() bool`

HasAllocatedAvailableRatio returns a boolean if a field has been set.

### GetAllocatedIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetAllocatedIpCount() int64`

GetAllocatedIpCount returns the AllocatedIpCount field if non-nil, zero value otherwise.

### GetAllocatedIpCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetAllocatedIpCountOk() (*int64, bool)`

GetAllocatedIpCountOk returns a tuple with the AllocatedIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetAllocatedIpCount(v int64)`

SetAllocatedIpCount sets AllocatedIpCount field to given value.

### HasAllocatedIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasAllocatedIpCount() bool`

HasAllocatedIpCount returns a boolean if a field has been set.

### GetAvailableIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetAvailableIpCount() string`

GetAvailableIpCount returns the AvailableIpCount field if non-nil, zero value otherwise.

### GetAvailableIpCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetAvailableIpCountOk() (*string, bool)`

GetAvailableIpCountOk returns a tuple with the AvailableIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetAvailableIpCount(v string)`

SetAvailableIpCount sets AvailableIpCount field to given value.

### HasAvailableIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasAvailableIpCount() bool`

HasAvailableIpCount returns a boolean if a field has been set.

### GetFixedIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetFixedIpCount() int64`

GetFixedIpCount returns the FixedIpCount field if non-nil, zero value otherwise.

### GetFixedIpCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetFixedIpCountOk() (*int64, bool)`

GetFixedIpCountOk returns a tuple with the FixedIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetFixedIpCount(v int64)`

SetFixedIpCount sets FixedIpCount field to given value.

### HasFixedIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasFixedIpCount() bool`

HasFixedIpCount returns a boolean if a field has been set.

### GetFloatingIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetFloatingIpCount() int64`

GetFloatingIpCount returns the FloatingIpCount field if non-nil, zero value otherwise.

### GetFloatingIpCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetFloatingIpCountOk() (*int64, bool)`

GetFloatingIpCountOk returns a tuple with the FloatingIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetFloatingIpCount(v int64)`

SetFloatingIpCount sets FloatingIpCount field to given value.

### HasFloatingIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasFloatingIpCount() bool`

HasFloatingIpCount returns a boolean if a field has been set.

### GetTenantCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetTenantCount() int64`

GetTenantCount returns the TenantCount field if non-nil, zero value otherwise.

### GetTenantCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetTenantCountOk() (*int64, bool)`

GetTenantCountOk returns a tuple with the TenantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetTenantCount(v int64)`

SetTenantCount sets TenantCount field to given value.

### HasTenantCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasTenantCount() bool`

HasTenantCount returns a boolean if a field has been set.

### GetTenantIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetTenantIpCount() int64`

GetTenantIpCount returns the TenantIpCount field if non-nil, zero value otherwise.

### GetTenantIpCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetTenantIpCountOk() (*int64, bool)`

GetTenantIpCountOk returns a tuple with the TenantIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetTenantIpCount(v int64)`

SetTenantIpCount sets TenantIpCount field to given value.

### HasTenantIpCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasTenantIpCount() bool`

HasTenantIpCount returns a boolean if a field has been set.

### GetTenantVmCount

`func (o *GetGridCloudapiCloudstatisticsResponse) GetTenantVmCount() int64`

GetTenantVmCount returns the TenantVmCount field if non-nil, zero value otherwise.

### GetTenantVmCountOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetTenantVmCountOk() (*int64, bool)`

GetTenantVmCountOk returns a tuple with the TenantVmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantVmCount

`func (o *GetGridCloudapiCloudstatisticsResponse) SetTenantVmCount(v int64)`

SetTenantVmCount sets TenantVmCount field to given value.

### HasTenantVmCount

`func (o *GetGridCloudapiCloudstatisticsResponse) HasTenantVmCount() bool`

HasTenantVmCount returns a boolean if a field has been set.

### GetResult

`func (o *GetGridCloudapiCloudstatisticsResponse) GetResult() GridCloudapiCloudstatistics`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridCloudapiCloudstatisticsResponse) GetResultOk() (*GridCloudapiCloudstatistics, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridCloudapiCloudstatisticsResponse) SetResult(v GridCloudapiCloudstatistics)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridCloudapiCloudstatisticsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


