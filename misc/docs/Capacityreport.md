# Capacityreport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**HardwareType** | Pointer to **string** | Hardware type of a Grid member. | [optional] [readonly] 
**MaxCapacity** | Pointer to **int64** | The maximum amount of capacity available for the Grid member. | [optional] [readonly] 
**Name** | Pointer to **string** | The Grid member name. | [optional] [readonly] 
**ObjectCounts** | Pointer to [**[]CapacityreportObjectCounts**](CapacityreportObjectCounts.md) | A list of instance counts for object types created on the Grid member. | [optional] [readonly] 
**PercentUsed** | Pointer to **int64** | The percentage of the capacity in use by the Grid member. | [optional] [readonly] 
**Role** | Pointer to **string** | The Grid member role. | [optional] [readonly] 
**TotalObjects** | Pointer to **int64** | The total number of objects created by the Grid member. | [optional] [readonly] 

## Methods

### NewCapacityreport

`func NewCapacityreport() *Capacityreport`

NewCapacityreport instantiates a new Capacityreport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityreportWithDefaults

`func NewCapacityreportWithDefaults() *Capacityreport`

NewCapacityreportWithDefaults instantiates a new Capacityreport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Capacityreport) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Capacityreport) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Capacityreport) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Capacityreport) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetHardwareType

`func (o *Capacityreport) GetHardwareType() string`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *Capacityreport) GetHardwareTypeOk() (*string, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *Capacityreport) SetHardwareType(v string)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *Capacityreport) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *Capacityreport) GetMaxCapacity() int64`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *Capacityreport) GetMaxCapacityOk() (*int64, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *Capacityreport) SetMaxCapacity(v int64)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *Capacityreport) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetName

`func (o *Capacityreport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Capacityreport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Capacityreport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Capacityreport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectCounts

`func (o *Capacityreport) GetObjectCounts() []CapacityreportObjectCounts`

GetObjectCounts returns the ObjectCounts field if non-nil, zero value otherwise.

### GetObjectCountsOk

`func (o *Capacityreport) GetObjectCountsOk() (*[]CapacityreportObjectCounts, bool)`

GetObjectCountsOk returns a tuple with the ObjectCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCounts

`func (o *Capacityreport) SetObjectCounts(v []CapacityreportObjectCounts)`

SetObjectCounts sets ObjectCounts field to given value.

### HasObjectCounts

`func (o *Capacityreport) HasObjectCounts() bool`

HasObjectCounts returns a boolean if a field has been set.

### GetPercentUsed

`func (o *Capacityreport) GetPercentUsed() int64`

GetPercentUsed returns the PercentUsed field if non-nil, zero value otherwise.

### GetPercentUsedOk

`func (o *Capacityreport) GetPercentUsedOk() (*int64, bool)`

GetPercentUsedOk returns a tuple with the PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUsed

`func (o *Capacityreport) SetPercentUsed(v int64)`

SetPercentUsed sets PercentUsed field to given value.

### HasPercentUsed

`func (o *Capacityreport) HasPercentUsed() bool`

HasPercentUsed returns a boolean if a field has been set.

### GetRole

`func (o *Capacityreport) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Capacityreport) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Capacityreport) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Capacityreport) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTotalObjects

`func (o *Capacityreport) GetTotalObjects() int64`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *Capacityreport) GetTotalObjectsOk() (*int64, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *Capacityreport) SetTotalObjects(v int64)`

SetTotalObjects sets TotalObjects field to given value.

### HasTotalObjects

`func (o *Capacityreport) HasTotalObjects() bool`

HasTotalObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


