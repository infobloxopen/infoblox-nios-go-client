# GridTrafficCaptureRecQueriesSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecursiveClientsCountTriggerEnable** | Pointer to **bool** | Enable triggering automated traffic capture based on outgoing recursive queries count. | [optional] 
**RecursiveClientsCountThreshold** | Pointer to **int64** | Concurrent outgoing recursive queries count below which traffic capture will be triggered. | [optional] 
**RecursiveClientsCountReset** | Pointer to **int64** | Concurrent outgoing recursive queries count below which traffic capture will be stopped. | [optional] 

## Methods

### NewGridTrafficCaptureRecQueriesSetting

`func NewGridTrafficCaptureRecQueriesSetting() *GridTrafficCaptureRecQueriesSetting`

NewGridTrafficCaptureRecQueriesSetting instantiates a new GridTrafficCaptureRecQueriesSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridTrafficCaptureRecQueriesSettingWithDefaults

`func NewGridTrafficCaptureRecQueriesSettingWithDefaults() *GridTrafficCaptureRecQueriesSetting`

NewGridTrafficCaptureRecQueriesSettingWithDefaults instantiates a new GridTrafficCaptureRecQueriesSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecursiveClientsCountTriggerEnable

`func (o *GridTrafficCaptureRecQueriesSetting) GetRecursiveClientsCountTriggerEnable() bool`

GetRecursiveClientsCountTriggerEnable returns the RecursiveClientsCountTriggerEnable field if non-nil, zero value otherwise.

### GetRecursiveClientsCountTriggerEnableOk

`func (o *GridTrafficCaptureRecQueriesSetting) GetRecursiveClientsCountTriggerEnableOk() (*bool, bool)`

GetRecursiveClientsCountTriggerEnableOk returns a tuple with the RecursiveClientsCountTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClientsCountTriggerEnable

`func (o *GridTrafficCaptureRecQueriesSetting) SetRecursiveClientsCountTriggerEnable(v bool)`

SetRecursiveClientsCountTriggerEnable sets RecursiveClientsCountTriggerEnable field to given value.

### HasRecursiveClientsCountTriggerEnable

`func (o *GridTrafficCaptureRecQueriesSetting) HasRecursiveClientsCountTriggerEnable() bool`

HasRecursiveClientsCountTriggerEnable returns a boolean if a field has been set.

### GetRecursiveClientsCountThreshold

`func (o *GridTrafficCaptureRecQueriesSetting) GetRecursiveClientsCountThreshold() int64`

GetRecursiveClientsCountThreshold returns the RecursiveClientsCountThreshold field if non-nil, zero value otherwise.

### GetRecursiveClientsCountThresholdOk

`func (o *GridTrafficCaptureRecQueriesSetting) GetRecursiveClientsCountThresholdOk() (*int64, bool)`

GetRecursiveClientsCountThresholdOk returns a tuple with the RecursiveClientsCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClientsCountThreshold

`func (o *GridTrafficCaptureRecQueriesSetting) SetRecursiveClientsCountThreshold(v int64)`

SetRecursiveClientsCountThreshold sets RecursiveClientsCountThreshold field to given value.

### HasRecursiveClientsCountThreshold

`func (o *GridTrafficCaptureRecQueriesSetting) HasRecursiveClientsCountThreshold() bool`

HasRecursiveClientsCountThreshold returns a boolean if a field has been set.

### GetRecursiveClientsCountReset

`func (o *GridTrafficCaptureRecQueriesSetting) GetRecursiveClientsCountReset() int64`

GetRecursiveClientsCountReset returns the RecursiveClientsCountReset field if non-nil, zero value otherwise.

### GetRecursiveClientsCountResetOk

`func (o *GridTrafficCaptureRecQueriesSetting) GetRecursiveClientsCountResetOk() (*int64, bool)`

GetRecursiveClientsCountResetOk returns a tuple with the RecursiveClientsCountReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClientsCountReset

`func (o *GridTrafficCaptureRecQueriesSetting) SetRecursiveClientsCountReset(v int64)`

SetRecursiveClientsCountReset sets RecursiveClientsCountReset field to given value.

### HasRecursiveClientsCountReset

`func (o *GridTrafficCaptureRecQueriesSetting) HasRecursiveClientsCountReset() bool`

HasRecursiveClientsCountReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


