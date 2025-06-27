# GridCloudapiCloudstatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllocatedAvailableRatio** | Pointer to **int64** | Ratio of allocated vs. available IPs | [optional] [readonly] 
**AllocatedIpCount** | Pointer to **int64** | Total number of IPs allocated by tenants. | [optional] [readonly] 
**AvailableIpCount** | Pointer to **string** | The total number of IP addresses available to tenants. Only IP addresses in networks that are within a delegation scope are counted. | [optional] [readonly] 
**FixedIpCount** | Pointer to **int64** | The total number of fixed IP addresses currently in use by all tenants in the system. | [optional] [readonly] 
**FloatingIpCount** | Pointer to **int64** | The total number of floating IP addresses currently in use by all tenants in the system. | [optional] [readonly] 
**TenantCount** | Pointer to **int64** | Total number of tenant currently in the system. | [optional] [readonly] 
**TenantIpCount** | Pointer to **int64** | The total number of IP addresses currently in use by all tenants in the system. | [optional] [readonly] 
**TenantVmCount** | Pointer to **int64** | The total number of VMs currently in use by all tenants in the system. | [optional] [readonly] 

## Methods

### NewGridCloudapiCloudstatistics

`func NewGridCloudapiCloudstatistics() *GridCloudapiCloudstatistics`

NewGridCloudapiCloudstatistics instantiates a new GridCloudapiCloudstatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCloudapiCloudstatisticsWithDefaults

`func NewGridCloudapiCloudstatisticsWithDefaults() *GridCloudapiCloudstatistics`

NewGridCloudapiCloudstatisticsWithDefaults instantiates a new GridCloudapiCloudstatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridCloudapiCloudstatistics) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridCloudapiCloudstatistics) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridCloudapiCloudstatistics) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridCloudapiCloudstatistics) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllocatedAvailableRatio

`func (o *GridCloudapiCloudstatistics) GetAllocatedAvailableRatio() int64`

GetAllocatedAvailableRatio returns the AllocatedAvailableRatio field if non-nil, zero value otherwise.

### GetAllocatedAvailableRatioOk

`func (o *GridCloudapiCloudstatistics) GetAllocatedAvailableRatioOk() (*int64, bool)`

GetAllocatedAvailableRatioOk returns a tuple with the AllocatedAvailableRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAvailableRatio

`func (o *GridCloudapiCloudstatistics) SetAllocatedAvailableRatio(v int64)`

SetAllocatedAvailableRatio sets AllocatedAvailableRatio field to given value.

### HasAllocatedAvailableRatio

`func (o *GridCloudapiCloudstatistics) HasAllocatedAvailableRatio() bool`

HasAllocatedAvailableRatio returns a boolean if a field has been set.

### GetAllocatedIpCount

`func (o *GridCloudapiCloudstatistics) GetAllocatedIpCount() int64`

GetAllocatedIpCount returns the AllocatedIpCount field if non-nil, zero value otherwise.

### GetAllocatedIpCountOk

`func (o *GridCloudapiCloudstatistics) GetAllocatedIpCountOk() (*int64, bool)`

GetAllocatedIpCountOk returns a tuple with the AllocatedIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedIpCount

`func (o *GridCloudapiCloudstatistics) SetAllocatedIpCount(v int64)`

SetAllocatedIpCount sets AllocatedIpCount field to given value.

### HasAllocatedIpCount

`func (o *GridCloudapiCloudstatistics) HasAllocatedIpCount() bool`

HasAllocatedIpCount returns a boolean if a field has been set.

### GetAvailableIpCount

`func (o *GridCloudapiCloudstatistics) GetAvailableIpCount() string`

GetAvailableIpCount returns the AvailableIpCount field if non-nil, zero value otherwise.

### GetAvailableIpCountOk

`func (o *GridCloudapiCloudstatistics) GetAvailableIpCountOk() (*string, bool)`

GetAvailableIpCountOk returns a tuple with the AvailableIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpCount

`func (o *GridCloudapiCloudstatistics) SetAvailableIpCount(v string)`

SetAvailableIpCount sets AvailableIpCount field to given value.

### HasAvailableIpCount

`func (o *GridCloudapiCloudstatistics) HasAvailableIpCount() bool`

HasAvailableIpCount returns a boolean if a field has been set.

### GetFixedIpCount

`func (o *GridCloudapiCloudstatistics) GetFixedIpCount() int64`

GetFixedIpCount returns the FixedIpCount field if non-nil, zero value otherwise.

### GetFixedIpCountOk

`func (o *GridCloudapiCloudstatistics) GetFixedIpCountOk() (*int64, bool)`

GetFixedIpCountOk returns a tuple with the FixedIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpCount

`func (o *GridCloudapiCloudstatistics) SetFixedIpCount(v int64)`

SetFixedIpCount sets FixedIpCount field to given value.

### HasFixedIpCount

`func (o *GridCloudapiCloudstatistics) HasFixedIpCount() bool`

HasFixedIpCount returns a boolean if a field has been set.

### GetFloatingIpCount

`func (o *GridCloudapiCloudstatistics) GetFloatingIpCount() int64`

GetFloatingIpCount returns the FloatingIpCount field if non-nil, zero value otherwise.

### GetFloatingIpCountOk

`func (o *GridCloudapiCloudstatistics) GetFloatingIpCountOk() (*int64, bool)`

GetFloatingIpCountOk returns a tuple with the FloatingIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpCount

`func (o *GridCloudapiCloudstatistics) SetFloatingIpCount(v int64)`

SetFloatingIpCount sets FloatingIpCount field to given value.

### HasFloatingIpCount

`func (o *GridCloudapiCloudstatistics) HasFloatingIpCount() bool`

HasFloatingIpCount returns a boolean if a field has been set.

### GetTenantCount

`func (o *GridCloudapiCloudstatistics) GetTenantCount() int64`

GetTenantCount returns the TenantCount field if non-nil, zero value otherwise.

### GetTenantCountOk

`func (o *GridCloudapiCloudstatistics) GetTenantCountOk() (*int64, bool)`

GetTenantCountOk returns a tuple with the TenantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCount

`func (o *GridCloudapiCloudstatistics) SetTenantCount(v int64)`

SetTenantCount sets TenantCount field to given value.

### HasTenantCount

`func (o *GridCloudapiCloudstatistics) HasTenantCount() bool`

HasTenantCount returns a boolean if a field has been set.

### GetTenantIpCount

`func (o *GridCloudapiCloudstatistics) GetTenantIpCount() int64`

GetTenantIpCount returns the TenantIpCount field if non-nil, zero value otherwise.

### GetTenantIpCountOk

`func (o *GridCloudapiCloudstatistics) GetTenantIpCountOk() (*int64, bool)`

GetTenantIpCountOk returns a tuple with the TenantIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIpCount

`func (o *GridCloudapiCloudstatistics) SetTenantIpCount(v int64)`

SetTenantIpCount sets TenantIpCount field to given value.

### HasTenantIpCount

`func (o *GridCloudapiCloudstatistics) HasTenantIpCount() bool`

HasTenantIpCount returns a boolean if a field has been set.

### GetTenantVmCount

`func (o *GridCloudapiCloudstatistics) GetTenantVmCount() int64`

GetTenantVmCount returns the TenantVmCount field if non-nil, zero value otherwise.

### GetTenantVmCountOk

`func (o *GridCloudapiCloudstatistics) GetTenantVmCountOk() (*int64, bool)`

GetTenantVmCountOk returns a tuple with the TenantVmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantVmCount

`func (o *GridCloudapiCloudstatistics) SetTenantVmCount(v int64)`

SetTenantVmCount sets TenantVmCount field to given value.

### HasTenantVmCount

`func (o *GridCloudapiCloudstatistics) HasTenantVmCount() bool`

HasTenantVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


