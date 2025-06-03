# GridTrafficCaptureAuthDnsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthDnsLatencyTriggerEnable** | Pointer to **bool** | Enabling trigger automated traffic capture based on authoritative DNS latency. | [optional] 
**AuthDnsLatencyThreshold** | Pointer to **int64** | Authoritative DNS latency below which traffic capture will be triggered. | [optional] 
**AuthDnsLatencyReset** | Pointer to **int64** | Authoritative DNS latency above which traffic capture will stopped. | [optional] 
**AuthDnsLatencyListenOnSource** | Pointer to **string** | The local IP DNS service is listen on (for authoritative DNS latency trigger). | [optional] 

## Methods

### NewGridTrafficCaptureAuthDnsSetting

`func NewGridTrafficCaptureAuthDnsSetting() *GridTrafficCaptureAuthDnsSetting`

NewGridTrafficCaptureAuthDnsSetting instantiates a new GridTrafficCaptureAuthDnsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridTrafficCaptureAuthDnsSettingWithDefaults

`func NewGridTrafficCaptureAuthDnsSettingWithDefaults() *GridTrafficCaptureAuthDnsSetting`

NewGridTrafficCaptureAuthDnsSettingWithDefaults instantiates a new GridTrafficCaptureAuthDnsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthDnsLatencyTriggerEnable

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyTriggerEnable() bool`

GetAuthDnsLatencyTriggerEnable returns the AuthDnsLatencyTriggerEnable field if non-nil, zero value otherwise.

### GetAuthDnsLatencyTriggerEnableOk

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyTriggerEnableOk() (*bool, bool)`

GetAuthDnsLatencyTriggerEnableOk returns a tuple with the AuthDnsLatencyTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyTriggerEnable

`func (o *GridTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyTriggerEnable(v bool)`

SetAuthDnsLatencyTriggerEnable sets AuthDnsLatencyTriggerEnable field to given value.

### HasAuthDnsLatencyTriggerEnable

`func (o *GridTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyTriggerEnable() bool`

HasAuthDnsLatencyTriggerEnable returns a boolean if a field has been set.

### GetAuthDnsLatencyThreshold

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyThreshold() int64`

GetAuthDnsLatencyThreshold returns the AuthDnsLatencyThreshold field if non-nil, zero value otherwise.

### GetAuthDnsLatencyThresholdOk

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyThresholdOk() (*int64, bool)`

GetAuthDnsLatencyThresholdOk returns a tuple with the AuthDnsLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyThreshold

`func (o *GridTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyThreshold(v int64)`

SetAuthDnsLatencyThreshold sets AuthDnsLatencyThreshold field to given value.

### HasAuthDnsLatencyThreshold

`func (o *GridTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyThreshold() bool`

HasAuthDnsLatencyThreshold returns a boolean if a field has been set.

### GetAuthDnsLatencyReset

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyReset() int64`

GetAuthDnsLatencyReset returns the AuthDnsLatencyReset field if non-nil, zero value otherwise.

### GetAuthDnsLatencyResetOk

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyResetOk() (*int64, bool)`

GetAuthDnsLatencyResetOk returns a tuple with the AuthDnsLatencyReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyReset

`func (o *GridTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyReset(v int64)`

SetAuthDnsLatencyReset sets AuthDnsLatencyReset field to given value.

### HasAuthDnsLatencyReset

`func (o *GridTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyReset() bool`

HasAuthDnsLatencyReset returns a boolean if a field has been set.

### GetAuthDnsLatencyListenOnSource

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyListenOnSource() string`

GetAuthDnsLatencyListenOnSource returns the AuthDnsLatencyListenOnSource field if non-nil, zero value otherwise.

### GetAuthDnsLatencyListenOnSourceOk

`func (o *GridTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyListenOnSourceOk() (*string, bool)`

GetAuthDnsLatencyListenOnSourceOk returns a tuple with the AuthDnsLatencyListenOnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyListenOnSource

`func (o *GridTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyListenOnSource(v string)`

SetAuthDnsLatencyListenOnSource sets AuthDnsLatencyListenOnSource field to given value.

### HasAuthDnsLatencyListenOnSource

`func (o *GridTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyListenOnSource() bool`

HasAuthDnsLatencyListenOnSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


