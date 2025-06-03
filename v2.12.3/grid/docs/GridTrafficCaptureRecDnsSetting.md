# GridTrafficCaptureRecDnsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecDnsLatencyTriggerEnable** | Pointer to **bool** | Enable triggering automated traffic capture based on recursive DNS latency. | [optional] 
**RecDnsLatencyThreshold** | Pointer to **int64** | Recursive DNS latency below which traffic capture will be triggered. | [optional] 
**RecDnsLatencyReset** | Pointer to **int64** | Recursive DNS latency above which traffic capture will be stopped. | [optional] 
**RecDnsLatencyListenOnSource** | Pointer to **string** | The local IP DNS service is listen on ( for recursive DNS latency trigger). | [optional] 
**KpiMonitoredDomains** | Pointer to [**GridtrafficcapturerecdnssettingKpiMonitoredDomains**](GridtrafficcapturerecdnssettingKpiMonitoredDomains.md) |  | [optional] 

## Methods

### NewGridTrafficCaptureRecDnsSetting

`func NewGridTrafficCaptureRecDnsSetting() *GridTrafficCaptureRecDnsSetting`

NewGridTrafficCaptureRecDnsSetting instantiates a new GridTrafficCaptureRecDnsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridTrafficCaptureRecDnsSettingWithDefaults

`func NewGridTrafficCaptureRecDnsSettingWithDefaults() *GridTrafficCaptureRecDnsSetting`

NewGridTrafficCaptureRecDnsSettingWithDefaults instantiates a new GridTrafficCaptureRecDnsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecDnsLatencyTriggerEnable

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyTriggerEnable() bool`

GetRecDnsLatencyTriggerEnable returns the RecDnsLatencyTriggerEnable field if non-nil, zero value otherwise.

### GetRecDnsLatencyTriggerEnableOk

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyTriggerEnableOk() (*bool, bool)`

GetRecDnsLatencyTriggerEnableOk returns a tuple with the RecDnsLatencyTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyTriggerEnable

`func (o *GridTrafficCaptureRecDnsSetting) SetRecDnsLatencyTriggerEnable(v bool)`

SetRecDnsLatencyTriggerEnable sets RecDnsLatencyTriggerEnable field to given value.

### HasRecDnsLatencyTriggerEnable

`func (o *GridTrafficCaptureRecDnsSetting) HasRecDnsLatencyTriggerEnable() bool`

HasRecDnsLatencyTriggerEnable returns a boolean if a field has been set.

### GetRecDnsLatencyThreshold

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyThreshold() int64`

GetRecDnsLatencyThreshold returns the RecDnsLatencyThreshold field if non-nil, zero value otherwise.

### GetRecDnsLatencyThresholdOk

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyThresholdOk() (*int64, bool)`

GetRecDnsLatencyThresholdOk returns a tuple with the RecDnsLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyThreshold

`func (o *GridTrafficCaptureRecDnsSetting) SetRecDnsLatencyThreshold(v int64)`

SetRecDnsLatencyThreshold sets RecDnsLatencyThreshold field to given value.

### HasRecDnsLatencyThreshold

`func (o *GridTrafficCaptureRecDnsSetting) HasRecDnsLatencyThreshold() bool`

HasRecDnsLatencyThreshold returns a boolean if a field has been set.

### GetRecDnsLatencyReset

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyReset() int64`

GetRecDnsLatencyReset returns the RecDnsLatencyReset field if non-nil, zero value otherwise.

### GetRecDnsLatencyResetOk

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyResetOk() (*int64, bool)`

GetRecDnsLatencyResetOk returns a tuple with the RecDnsLatencyReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyReset

`func (o *GridTrafficCaptureRecDnsSetting) SetRecDnsLatencyReset(v int64)`

SetRecDnsLatencyReset sets RecDnsLatencyReset field to given value.

### HasRecDnsLatencyReset

`func (o *GridTrafficCaptureRecDnsSetting) HasRecDnsLatencyReset() bool`

HasRecDnsLatencyReset returns a boolean if a field has been set.

### GetRecDnsLatencyListenOnSource

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyListenOnSource() string`

GetRecDnsLatencyListenOnSource returns the RecDnsLatencyListenOnSource field if non-nil, zero value otherwise.

### GetRecDnsLatencyListenOnSourceOk

`func (o *GridTrafficCaptureRecDnsSetting) GetRecDnsLatencyListenOnSourceOk() (*string, bool)`

GetRecDnsLatencyListenOnSourceOk returns a tuple with the RecDnsLatencyListenOnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyListenOnSource

`func (o *GridTrafficCaptureRecDnsSetting) SetRecDnsLatencyListenOnSource(v string)`

SetRecDnsLatencyListenOnSource sets RecDnsLatencyListenOnSource field to given value.

### HasRecDnsLatencyListenOnSource

`func (o *GridTrafficCaptureRecDnsSetting) HasRecDnsLatencyListenOnSource() bool`

HasRecDnsLatencyListenOnSource returns a boolean if a field has been set.

### GetKpiMonitoredDomains

`func (o *GridTrafficCaptureRecDnsSetting) GetKpiMonitoredDomains() GridtrafficcapturerecdnssettingKpiMonitoredDomains`

GetKpiMonitoredDomains returns the KpiMonitoredDomains field if non-nil, zero value otherwise.

### GetKpiMonitoredDomainsOk

`func (o *GridTrafficCaptureRecDnsSetting) GetKpiMonitoredDomainsOk() (*GridtrafficcapturerecdnssettingKpiMonitoredDomains, bool)`

GetKpiMonitoredDomainsOk returns a tuple with the KpiMonitoredDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpiMonitoredDomains

`func (o *GridTrafficCaptureRecDnsSetting) SetKpiMonitoredDomains(v GridtrafficcapturerecdnssettingKpiMonitoredDomains)`

SetKpiMonitoredDomains sets KpiMonitoredDomains field to given value.

### HasKpiMonitoredDomains

`func (o *GridTrafficCaptureRecDnsSetting) HasKpiMonitoredDomains() bool`

HasKpiMonitoredDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


