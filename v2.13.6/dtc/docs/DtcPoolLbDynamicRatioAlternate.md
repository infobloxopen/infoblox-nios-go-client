# DtcPoolLbDynamicRatioAlternate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | The method of the DTC dynamic ratio load balancing. | [optional] 
**Monitor** | Pointer to **string** | The DTC monitor output of which will be used for dynamic ratio load balancing. | [optional] 
**MonitorMetric** | Pointer to **string** | The metric of the DTC SNMP monitor that will be used for dynamic weighing. | [optional] 
**MonitorWeighing** | Pointer to **string** | The DTC monitor weight. &#39;PRIORITY&#39; means that all clients will be forwarded to the least loaded server. &#39;RATIO&#39; means that distribution will be calculated based on dynamic weights. | [optional] 
**InvertMonitorMetric** | Pointer to **bool** | Determines whether the inverted values of the DTC SNMP monitor metric will be used. | [optional] 

## Methods

### NewDtcPoolLbDynamicRatioAlternate

`func NewDtcPoolLbDynamicRatioAlternate() *DtcPoolLbDynamicRatioAlternate`

NewDtcPoolLbDynamicRatioAlternate instantiates a new DtcPoolLbDynamicRatioAlternate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcPoolLbDynamicRatioAlternateWithDefaults

`func NewDtcPoolLbDynamicRatioAlternateWithDefaults() *DtcPoolLbDynamicRatioAlternate`

NewDtcPoolLbDynamicRatioAlternateWithDefaults instantiates a new DtcPoolLbDynamicRatioAlternate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *DtcPoolLbDynamicRatioAlternate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DtcPoolLbDynamicRatioAlternate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DtcPoolLbDynamicRatioAlternate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *DtcPoolLbDynamicRatioAlternate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMonitor

`func (o *DtcPoolLbDynamicRatioAlternate) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *DtcPoolLbDynamicRatioAlternate) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *DtcPoolLbDynamicRatioAlternate) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMonitorMetric

`func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorMetric() string`

GetMonitorMetric returns the MonitorMetric field if non-nil, zero value otherwise.

### GetMonitorMetricOk

`func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorMetricOk() (*string, bool)`

GetMonitorMetricOk returns a tuple with the MonitorMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorMetric

`func (o *DtcPoolLbDynamicRatioAlternate) SetMonitorMetric(v string)`

SetMonitorMetric sets MonitorMetric field to given value.

### HasMonitorMetric

`func (o *DtcPoolLbDynamicRatioAlternate) HasMonitorMetric() bool`

HasMonitorMetric returns a boolean if a field has been set.

### GetMonitorWeighing

`func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorWeighing() string`

GetMonitorWeighing returns the MonitorWeighing field if non-nil, zero value otherwise.

### GetMonitorWeighingOk

`func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorWeighingOk() (*string, bool)`

GetMonitorWeighingOk returns a tuple with the MonitorWeighing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorWeighing

`func (o *DtcPoolLbDynamicRatioAlternate) SetMonitorWeighing(v string)`

SetMonitorWeighing sets MonitorWeighing field to given value.

### HasMonitorWeighing

`func (o *DtcPoolLbDynamicRatioAlternate) HasMonitorWeighing() bool`

HasMonitorWeighing returns a boolean if a field has been set.

### GetInvertMonitorMetric

`func (o *DtcPoolLbDynamicRatioAlternate) GetInvertMonitorMetric() bool`

GetInvertMonitorMetric returns the InvertMonitorMetric field if non-nil, zero value otherwise.

### GetInvertMonitorMetricOk

`func (o *DtcPoolLbDynamicRatioAlternate) GetInvertMonitorMetricOk() (*bool, bool)`

GetInvertMonitorMetricOk returns a tuple with the InvertMonitorMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertMonitorMetric

`func (o *DtcPoolLbDynamicRatioAlternate) SetInvertMonitorMetric(v bool)`

SetInvertMonitorMetric sets InvertMonitorMetric field to given value.

### HasInvertMonitorMetric

`func (o *DtcPoolLbDynamicRatioAlternate) HasInvertMonitorMetric() bool`

HasInvertMonitorMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


