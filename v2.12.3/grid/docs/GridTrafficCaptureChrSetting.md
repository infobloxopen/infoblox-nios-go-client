# GridTrafficCaptureChrSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChrTriggerEnable** | Pointer to **bool** | Enable triggering automated traffic capture based on cache hit ratio thresholds. | [optional] 
**ChrThreshold** | Pointer to **int64** | DNS Cache hit ratio threshold(%) below which traffic capture will be triggered. | [optional] 
**ChrReset** | Pointer to **int64** | DNS Cache hit ratio threshold(%) above which traffic capture will be triggered. | [optional] 
**ChrMinCacheUtilization** | Pointer to **int64** | Minimum DNS cache utilization threshold(%) for triggering traffic capture based on DNS cache hit ratio. | [optional] 

## Methods

### NewGridTrafficCaptureChrSetting

`func NewGridTrafficCaptureChrSetting() *GridTrafficCaptureChrSetting`

NewGridTrafficCaptureChrSetting instantiates a new GridTrafficCaptureChrSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridTrafficCaptureChrSettingWithDefaults

`func NewGridTrafficCaptureChrSettingWithDefaults() *GridTrafficCaptureChrSetting`

NewGridTrafficCaptureChrSettingWithDefaults instantiates a new GridTrafficCaptureChrSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChrTriggerEnable

`func (o *GridTrafficCaptureChrSetting) GetChrTriggerEnable() bool`

GetChrTriggerEnable returns the ChrTriggerEnable field if non-nil, zero value otherwise.

### GetChrTriggerEnableOk

`func (o *GridTrafficCaptureChrSetting) GetChrTriggerEnableOk() (*bool, bool)`

GetChrTriggerEnableOk returns a tuple with the ChrTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrTriggerEnable

`func (o *GridTrafficCaptureChrSetting) SetChrTriggerEnable(v bool)`

SetChrTriggerEnable sets ChrTriggerEnable field to given value.

### HasChrTriggerEnable

`func (o *GridTrafficCaptureChrSetting) HasChrTriggerEnable() bool`

HasChrTriggerEnable returns a boolean if a field has been set.

### GetChrThreshold

`func (o *GridTrafficCaptureChrSetting) GetChrThreshold() int64`

GetChrThreshold returns the ChrThreshold field if non-nil, zero value otherwise.

### GetChrThresholdOk

`func (o *GridTrafficCaptureChrSetting) GetChrThresholdOk() (*int64, bool)`

GetChrThresholdOk returns a tuple with the ChrThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrThreshold

`func (o *GridTrafficCaptureChrSetting) SetChrThreshold(v int64)`

SetChrThreshold sets ChrThreshold field to given value.

### HasChrThreshold

`func (o *GridTrafficCaptureChrSetting) HasChrThreshold() bool`

HasChrThreshold returns a boolean if a field has been set.

### GetChrReset

`func (o *GridTrafficCaptureChrSetting) GetChrReset() int64`

GetChrReset returns the ChrReset field if non-nil, zero value otherwise.

### GetChrResetOk

`func (o *GridTrafficCaptureChrSetting) GetChrResetOk() (*int64, bool)`

GetChrResetOk returns a tuple with the ChrReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrReset

`func (o *GridTrafficCaptureChrSetting) SetChrReset(v int64)`

SetChrReset sets ChrReset field to given value.

### HasChrReset

`func (o *GridTrafficCaptureChrSetting) HasChrReset() bool`

HasChrReset returns a boolean if a field has been set.

### GetChrMinCacheUtilization

`func (o *GridTrafficCaptureChrSetting) GetChrMinCacheUtilization() int64`

GetChrMinCacheUtilization returns the ChrMinCacheUtilization field if non-nil, zero value otherwise.

### GetChrMinCacheUtilizationOk

`func (o *GridTrafficCaptureChrSetting) GetChrMinCacheUtilizationOk() (*int64, bool)`

GetChrMinCacheUtilizationOk returns a tuple with the ChrMinCacheUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrMinCacheUtilization

`func (o *GridTrafficCaptureChrSetting) SetChrMinCacheUtilization(v int64)`

SetChrMinCacheUtilization sets ChrMinCacheUtilization field to given value.

### HasChrMinCacheUtilization

`func (o *GridTrafficCaptureChrSetting) HasChrMinCacheUtilization() bool`

HasChrMinCacheUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


