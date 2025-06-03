# MemberTrafficCaptureRecDnsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecDnsLatencyTriggerEnable** | Pointer to **bool** | Enable triggering automated traffic capture based on recursive DNS latency. | [optional] 
**RecDnsLatencyThreshold** | Pointer to **int64** | Recursive DNS latency below which traffic capture will be triggered. | [optional] 
**RecDnsLatencyReset** | Pointer to **int64** | Recursive DNS latency above which traffic capture will be stopped. | [optional] 
**RecDnsLatencyListenOnSource** | Pointer to **string** | The local IP DNS service is listen on ( for recursive DNS latency trigger). | [optional] 
**RecDnsLatencyListenOnIp** | Pointer to **string** | The DNS listen-on IP address used if rec_dns_latency_listen_on_source is IP. | [optional] 
**KpiMonitoredDomains** | Pointer to [**MembertrafficcapturerecdnssettingKpiMonitoredDomains**](MembertrafficcapturerecdnssettingKpiMonitoredDomains.md) |  | [optional] 

## Methods

### NewMemberTrafficCaptureRecDnsSetting

`func NewMemberTrafficCaptureRecDnsSetting() *MemberTrafficCaptureRecDnsSetting`

NewMemberTrafficCaptureRecDnsSetting instantiates a new MemberTrafficCaptureRecDnsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberTrafficCaptureRecDnsSettingWithDefaults

`func NewMemberTrafficCaptureRecDnsSettingWithDefaults() *MemberTrafficCaptureRecDnsSetting`

NewMemberTrafficCaptureRecDnsSettingWithDefaults instantiates a new MemberTrafficCaptureRecDnsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecDnsLatencyTriggerEnable

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyTriggerEnable() bool`

GetRecDnsLatencyTriggerEnable returns the RecDnsLatencyTriggerEnable field if non-nil, zero value otherwise.

### GetRecDnsLatencyTriggerEnableOk

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyTriggerEnableOk() (*bool, bool)`

GetRecDnsLatencyTriggerEnableOk returns a tuple with the RecDnsLatencyTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyTriggerEnable

`func (o *MemberTrafficCaptureRecDnsSetting) SetRecDnsLatencyTriggerEnable(v bool)`

SetRecDnsLatencyTriggerEnable sets RecDnsLatencyTriggerEnable field to given value.

### HasRecDnsLatencyTriggerEnable

`func (o *MemberTrafficCaptureRecDnsSetting) HasRecDnsLatencyTriggerEnable() bool`

HasRecDnsLatencyTriggerEnable returns a boolean if a field has been set.

### GetRecDnsLatencyThreshold

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyThreshold() int64`

GetRecDnsLatencyThreshold returns the RecDnsLatencyThreshold field if non-nil, zero value otherwise.

### GetRecDnsLatencyThresholdOk

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyThresholdOk() (*int64, bool)`

GetRecDnsLatencyThresholdOk returns a tuple with the RecDnsLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyThreshold

`func (o *MemberTrafficCaptureRecDnsSetting) SetRecDnsLatencyThreshold(v int64)`

SetRecDnsLatencyThreshold sets RecDnsLatencyThreshold field to given value.

### HasRecDnsLatencyThreshold

`func (o *MemberTrafficCaptureRecDnsSetting) HasRecDnsLatencyThreshold() bool`

HasRecDnsLatencyThreshold returns a boolean if a field has been set.

### GetRecDnsLatencyReset

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyReset() int64`

GetRecDnsLatencyReset returns the RecDnsLatencyReset field if non-nil, zero value otherwise.

### GetRecDnsLatencyResetOk

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyResetOk() (*int64, bool)`

GetRecDnsLatencyResetOk returns a tuple with the RecDnsLatencyReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyReset

`func (o *MemberTrafficCaptureRecDnsSetting) SetRecDnsLatencyReset(v int64)`

SetRecDnsLatencyReset sets RecDnsLatencyReset field to given value.

### HasRecDnsLatencyReset

`func (o *MemberTrafficCaptureRecDnsSetting) HasRecDnsLatencyReset() bool`

HasRecDnsLatencyReset returns a boolean if a field has been set.

### GetRecDnsLatencyListenOnSource

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyListenOnSource() string`

GetRecDnsLatencyListenOnSource returns the RecDnsLatencyListenOnSource field if non-nil, zero value otherwise.

### GetRecDnsLatencyListenOnSourceOk

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyListenOnSourceOk() (*string, bool)`

GetRecDnsLatencyListenOnSourceOk returns a tuple with the RecDnsLatencyListenOnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyListenOnSource

`func (o *MemberTrafficCaptureRecDnsSetting) SetRecDnsLatencyListenOnSource(v string)`

SetRecDnsLatencyListenOnSource sets RecDnsLatencyListenOnSource field to given value.

### HasRecDnsLatencyListenOnSource

`func (o *MemberTrafficCaptureRecDnsSetting) HasRecDnsLatencyListenOnSource() bool`

HasRecDnsLatencyListenOnSource returns a boolean if a field has been set.

### GetRecDnsLatencyListenOnIp

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyListenOnIp() string`

GetRecDnsLatencyListenOnIp returns the RecDnsLatencyListenOnIp field if non-nil, zero value otherwise.

### GetRecDnsLatencyListenOnIpOk

`func (o *MemberTrafficCaptureRecDnsSetting) GetRecDnsLatencyListenOnIpOk() (*string, bool)`

GetRecDnsLatencyListenOnIpOk returns a tuple with the RecDnsLatencyListenOnIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecDnsLatencyListenOnIp

`func (o *MemberTrafficCaptureRecDnsSetting) SetRecDnsLatencyListenOnIp(v string)`

SetRecDnsLatencyListenOnIp sets RecDnsLatencyListenOnIp field to given value.

### HasRecDnsLatencyListenOnIp

`func (o *MemberTrafficCaptureRecDnsSetting) HasRecDnsLatencyListenOnIp() bool`

HasRecDnsLatencyListenOnIp returns a boolean if a field has been set.

### GetKpiMonitoredDomains

`func (o *MemberTrafficCaptureRecDnsSetting) GetKpiMonitoredDomains() MembertrafficcapturerecdnssettingKpiMonitoredDomains`

GetKpiMonitoredDomains returns the KpiMonitoredDomains field if non-nil, zero value otherwise.

### GetKpiMonitoredDomainsOk

`func (o *MemberTrafficCaptureRecDnsSetting) GetKpiMonitoredDomainsOk() (*MembertrafficcapturerecdnssettingKpiMonitoredDomains, bool)`

GetKpiMonitoredDomainsOk returns a tuple with the KpiMonitoredDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpiMonitoredDomains

`func (o *MemberTrafficCaptureRecDnsSetting) SetKpiMonitoredDomains(v MembertrafficcapturerecdnssettingKpiMonitoredDomains)`

SetKpiMonitoredDomains sets KpiMonitoredDomains field to given value.

### HasKpiMonitoredDomains

`func (o *MemberTrafficCaptureRecDnsSetting) HasKpiMonitoredDomains() bool`

HasKpiMonitoredDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


