# GridTrafficCaptureQpsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QpsTriggerEnable** | Pointer to **bool** | Enable triggering automated traffic capture based on DNS queries per second threshold. | [optional] 
**QpsThreshold** | Pointer to **int64** | DNS queries per second threshold below which traffic capture will be triggered. | [optional] 
**QpsReset** | Pointer to **int64** | DNS queries per second threshold below which traffic capture will be stopped. | [optional] 

## Methods

### NewGridTrafficCaptureQpsSetting

`func NewGridTrafficCaptureQpsSetting() *GridTrafficCaptureQpsSetting`

NewGridTrafficCaptureQpsSetting instantiates a new GridTrafficCaptureQpsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridTrafficCaptureQpsSettingWithDefaults

`func NewGridTrafficCaptureQpsSettingWithDefaults() *GridTrafficCaptureQpsSetting`

NewGridTrafficCaptureQpsSettingWithDefaults instantiates a new GridTrafficCaptureQpsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQpsTriggerEnable

`func (o *GridTrafficCaptureQpsSetting) GetQpsTriggerEnable() bool`

GetQpsTriggerEnable returns the QpsTriggerEnable field if non-nil, zero value otherwise.

### GetQpsTriggerEnableOk

`func (o *GridTrafficCaptureQpsSetting) GetQpsTriggerEnableOk() (*bool, bool)`

GetQpsTriggerEnableOk returns a tuple with the QpsTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQpsTriggerEnable

`func (o *GridTrafficCaptureQpsSetting) SetQpsTriggerEnable(v bool)`

SetQpsTriggerEnable sets QpsTriggerEnable field to given value.

### HasQpsTriggerEnable

`func (o *GridTrafficCaptureQpsSetting) HasQpsTriggerEnable() bool`

HasQpsTriggerEnable returns a boolean if a field has been set.

### GetQpsThreshold

`func (o *GridTrafficCaptureQpsSetting) GetQpsThreshold() int64`

GetQpsThreshold returns the QpsThreshold field if non-nil, zero value otherwise.

### GetQpsThresholdOk

`func (o *GridTrafficCaptureQpsSetting) GetQpsThresholdOk() (*int64, bool)`

GetQpsThresholdOk returns a tuple with the QpsThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQpsThreshold

`func (o *GridTrafficCaptureQpsSetting) SetQpsThreshold(v int64)`

SetQpsThreshold sets QpsThreshold field to given value.

### HasQpsThreshold

`func (o *GridTrafficCaptureQpsSetting) HasQpsThreshold() bool`

HasQpsThreshold returns a boolean if a field has been set.

### GetQpsReset

`func (o *GridTrafficCaptureQpsSetting) GetQpsReset() int64`

GetQpsReset returns the QpsReset field if non-nil, zero value otherwise.

### GetQpsResetOk

`func (o *GridTrafficCaptureQpsSetting) GetQpsResetOk() (*int64, bool)`

GetQpsResetOk returns a tuple with the QpsReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQpsReset

`func (o *GridTrafficCaptureQpsSetting) SetQpsReset(v int64)`

SetQpsReset sets QpsReset field to given value.

### HasQpsReset

`func (o *GridTrafficCaptureQpsSetting) HasQpsReset() bool`

HasQpsReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


