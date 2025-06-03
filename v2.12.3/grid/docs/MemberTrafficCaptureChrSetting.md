# MemberTrafficCaptureChrSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChrTriggerEnable** | Pointer to **bool** | Enable triggering automated traffic capture based on cache hit ratio thresholds. | [optional] 
**ChrThreshold** | Pointer to **int64** | DNS Cache hit ratio threshold(%) below which traffic capture will be triggered. | [optional] 
**ChrReset** | Pointer to **int64** | DNS Cache hit ratio threshold(%) above which traffic capture will be triggered. | [optional] 
**ChrMinCacheUtilization** | Pointer to **int64** | Minimum DNS cache utilization threshold(%) for triggering traffic capture based on DNS cache hit ratio. | [optional] 

## Methods

### NewMemberTrafficCaptureChrSetting

`func NewMemberTrafficCaptureChrSetting() *MemberTrafficCaptureChrSetting`

NewMemberTrafficCaptureChrSetting instantiates a new MemberTrafficCaptureChrSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberTrafficCaptureChrSettingWithDefaults

`func NewMemberTrafficCaptureChrSettingWithDefaults() *MemberTrafficCaptureChrSetting`

NewMemberTrafficCaptureChrSettingWithDefaults instantiates a new MemberTrafficCaptureChrSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChrTriggerEnable

`func (o *MemberTrafficCaptureChrSetting) GetChrTriggerEnable() bool`

GetChrTriggerEnable returns the ChrTriggerEnable field if non-nil, zero value otherwise.

### GetChrTriggerEnableOk

`func (o *MemberTrafficCaptureChrSetting) GetChrTriggerEnableOk() (*bool, bool)`

GetChrTriggerEnableOk returns a tuple with the ChrTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrTriggerEnable

`func (o *MemberTrafficCaptureChrSetting) SetChrTriggerEnable(v bool)`

SetChrTriggerEnable sets ChrTriggerEnable field to given value.

### HasChrTriggerEnable

`func (o *MemberTrafficCaptureChrSetting) HasChrTriggerEnable() bool`

HasChrTriggerEnable returns a boolean if a field has been set.

### GetChrThreshold

`func (o *MemberTrafficCaptureChrSetting) GetChrThreshold() int64`

GetChrThreshold returns the ChrThreshold field if non-nil, zero value otherwise.

### GetChrThresholdOk

`func (o *MemberTrafficCaptureChrSetting) GetChrThresholdOk() (*int64, bool)`

GetChrThresholdOk returns a tuple with the ChrThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrThreshold

`func (o *MemberTrafficCaptureChrSetting) SetChrThreshold(v int64)`

SetChrThreshold sets ChrThreshold field to given value.

### HasChrThreshold

`func (o *MemberTrafficCaptureChrSetting) HasChrThreshold() bool`

HasChrThreshold returns a boolean if a field has been set.

### GetChrReset

`func (o *MemberTrafficCaptureChrSetting) GetChrReset() int64`

GetChrReset returns the ChrReset field if non-nil, zero value otherwise.

### GetChrResetOk

`func (o *MemberTrafficCaptureChrSetting) GetChrResetOk() (*int64, bool)`

GetChrResetOk returns a tuple with the ChrReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrReset

`func (o *MemberTrafficCaptureChrSetting) SetChrReset(v int64)`

SetChrReset sets ChrReset field to given value.

### HasChrReset

`func (o *MemberTrafficCaptureChrSetting) HasChrReset() bool`

HasChrReset returns a boolean if a field has been set.

### GetChrMinCacheUtilization

`func (o *MemberTrafficCaptureChrSetting) GetChrMinCacheUtilization() int64`

GetChrMinCacheUtilization returns the ChrMinCacheUtilization field if non-nil, zero value otherwise.

### GetChrMinCacheUtilizationOk

`func (o *MemberTrafficCaptureChrSetting) GetChrMinCacheUtilizationOk() (*int64, bool)`

GetChrMinCacheUtilizationOk returns a tuple with the ChrMinCacheUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrMinCacheUtilization

`func (o *MemberTrafficCaptureChrSetting) SetChrMinCacheUtilization(v int64)`

SetChrMinCacheUtilization sets ChrMinCacheUtilization field to given value.

### HasChrMinCacheUtilization

`func (o *MemberTrafficCaptureChrSetting) HasChrMinCacheUtilization() bool`

HasChrMinCacheUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


