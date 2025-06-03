# GetCapacityreportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**HardwareType** | Pointer to **string** | Hardware type of a Grid member. | [optional] [readonly] 
**MaxCapacity** | Pointer to **int64** | The maximum amount of capacity available for the Grid member. | [optional] [readonly] 
**Name** | Pointer to **string** | The Grid member name. | [optional] [readonly] 
**ObjectCounts** | Pointer to [**[]CapacityreportObjectCounts**](CapacityreportObjectCounts.md) | A list of instance counts for object types created on the Grid member. | [optional] [readonly] 
**PercentUsed** | Pointer to **int64** | The percentage of the capacity in use by the Grid member. | [optional] [readonly] 
**Role** | Pointer to **string** | The Grid member role. | [optional] [readonly] 
**TotalObjects** | Pointer to **int64** | The total number of objects created by the Grid member. | [optional] [readonly] 
**Result** | Pointer to [**Capacityreport**](Capacityreport.md) |  | [optional] 

## Methods

### NewGetCapacityreportResponse

`func NewGetCapacityreportResponse() *GetCapacityreportResponse`

NewGetCapacityreportResponse instantiates a new GetCapacityreportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCapacityreportResponseWithDefaults

`func NewGetCapacityreportResponseWithDefaults() *GetCapacityreportResponse`

NewGetCapacityreportResponseWithDefaults instantiates a new GetCapacityreportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetCapacityreportResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetCapacityreportResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetCapacityreportResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetCapacityreportResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetHardwareType

`func (o *GetCapacityreportResponse) GetHardwareType() string`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *GetCapacityreportResponse) GetHardwareTypeOk() (*string, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *GetCapacityreportResponse) SetHardwareType(v string)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *GetCapacityreportResponse) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *GetCapacityreportResponse) GetMaxCapacity() int64`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *GetCapacityreportResponse) GetMaxCapacityOk() (*int64, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *GetCapacityreportResponse) SetMaxCapacity(v int64)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *GetCapacityreportResponse) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetName

`func (o *GetCapacityreportResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCapacityreportResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCapacityreportResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCapacityreportResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectCounts

`func (o *GetCapacityreportResponse) GetObjectCounts() []CapacityreportObjectCounts`

GetObjectCounts returns the ObjectCounts field if non-nil, zero value otherwise.

### GetObjectCountsOk

`func (o *GetCapacityreportResponse) GetObjectCountsOk() (*[]CapacityreportObjectCounts, bool)`

GetObjectCountsOk returns a tuple with the ObjectCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCounts

`func (o *GetCapacityreportResponse) SetObjectCounts(v []CapacityreportObjectCounts)`

SetObjectCounts sets ObjectCounts field to given value.

### HasObjectCounts

`func (o *GetCapacityreportResponse) HasObjectCounts() bool`

HasObjectCounts returns a boolean if a field has been set.

### GetPercentUsed

`func (o *GetCapacityreportResponse) GetPercentUsed() int64`

GetPercentUsed returns the PercentUsed field if non-nil, zero value otherwise.

### GetPercentUsedOk

`func (o *GetCapacityreportResponse) GetPercentUsedOk() (*int64, bool)`

GetPercentUsedOk returns a tuple with the PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUsed

`func (o *GetCapacityreportResponse) SetPercentUsed(v int64)`

SetPercentUsed sets PercentUsed field to given value.

### HasPercentUsed

`func (o *GetCapacityreportResponse) HasPercentUsed() bool`

HasPercentUsed returns a boolean if a field has been set.

### GetRole

`func (o *GetCapacityreportResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetCapacityreportResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetCapacityreportResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetCapacityreportResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTotalObjects

`func (o *GetCapacityreportResponse) GetTotalObjects() int64`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *GetCapacityreportResponse) GetTotalObjectsOk() (*int64, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *GetCapacityreportResponse) SetTotalObjects(v int64)`

SetTotalObjects sets TotalObjects field to given value.

### HasTotalObjects

`func (o *GetCapacityreportResponse) HasTotalObjects() bool`

HasTotalObjects returns a boolean if a field has been set.

### GetResult

`func (o *GetCapacityreportResponse) GetResult() Capacityreport`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetCapacityreportResponse) GetResultOk() (*Capacityreport, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetCapacityreportResponse) SetResult(v Capacityreport)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetCapacityreportResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


