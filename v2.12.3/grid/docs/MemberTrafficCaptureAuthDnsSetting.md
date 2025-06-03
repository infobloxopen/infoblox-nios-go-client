# MemberTrafficCaptureAuthDnsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthDnsLatencyTriggerEnable** | Pointer to **bool** | Enabling trigger automated traffic capture based on authoritative DNS latency. | [optional] 
**AuthDnsLatencyThreshold** | Pointer to **int64** | Authoritative DNS latency below which traffic capture will be triggered. | [optional] 
**AuthDnsLatencyReset** | Pointer to **int64** | Authoritative DNS latency above which traffic capture will stopped. | [optional] 
**AuthDnsLatencyListenOnSource** | Pointer to **string** | The local IP DNS service is listen on (for authoritative DNS latency trigger). | [optional] 
**AuthDnsLatencyListenOnIp** | Pointer to **string** | The DNS listen-on IP address used if auth_dns_latency_on_source is IP. | [optional] 

## Methods

### NewMemberTrafficCaptureAuthDnsSetting

`func NewMemberTrafficCaptureAuthDnsSetting() *MemberTrafficCaptureAuthDnsSetting`

NewMemberTrafficCaptureAuthDnsSetting instantiates a new MemberTrafficCaptureAuthDnsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberTrafficCaptureAuthDnsSettingWithDefaults

`func NewMemberTrafficCaptureAuthDnsSettingWithDefaults() *MemberTrafficCaptureAuthDnsSetting`

NewMemberTrafficCaptureAuthDnsSettingWithDefaults instantiates a new MemberTrafficCaptureAuthDnsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthDnsLatencyTriggerEnable

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyTriggerEnable() bool`

GetAuthDnsLatencyTriggerEnable returns the AuthDnsLatencyTriggerEnable field if non-nil, zero value otherwise.

### GetAuthDnsLatencyTriggerEnableOk

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyTriggerEnableOk() (*bool, bool)`

GetAuthDnsLatencyTriggerEnableOk returns a tuple with the AuthDnsLatencyTriggerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyTriggerEnable

`func (o *MemberTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyTriggerEnable(v bool)`

SetAuthDnsLatencyTriggerEnable sets AuthDnsLatencyTriggerEnable field to given value.

### HasAuthDnsLatencyTriggerEnable

`func (o *MemberTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyTriggerEnable() bool`

HasAuthDnsLatencyTriggerEnable returns a boolean if a field has been set.

### GetAuthDnsLatencyThreshold

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyThreshold() int64`

GetAuthDnsLatencyThreshold returns the AuthDnsLatencyThreshold field if non-nil, zero value otherwise.

### GetAuthDnsLatencyThresholdOk

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyThresholdOk() (*int64, bool)`

GetAuthDnsLatencyThresholdOk returns a tuple with the AuthDnsLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyThreshold

`func (o *MemberTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyThreshold(v int64)`

SetAuthDnsLatencyThreshold sets AuthDnsLatencyThreshold field to given value.

### HasAuthDnsLatencyThreshold

`func (o *MemberTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyThreshold() bool`

HasAuthDnsLatencyThreshold returns a boolean if a field has been set.

### GetAuthDnsLatencyReset

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyReset() int64`

GetAuthDnsLatencyReset returns the AuthDnsLatencyReset field if non-nil, zero value otherwise.

### GetAuthDnsLatencyResetOk

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyResetOk() (*int64, bool)`

GetAuthDnsLatencyResetOk returns a tuple with the AuthDnsLatencyReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyReset

`func (o *MemberTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyReset(v int64)`

SetAuthDnsLatencyReset sets AuthDnsLatencyReset field to given value.

### HasAuthDnsLatencyReset

`func (o *MemberTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyReset() bool`

HasAuthDnsLatencyReset returns a boolean if a field has been set.

### GetAuthDnsLatencyListenOnSource

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyListenOnSource() string`

GetAuthDnsLatencyListenOnSource returns the AuthDnsLatencyListenOnSource field if non-nil, zero value otherwise.

### GetAuthDnsLatencyListenOnSourceOk

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyListenOnSourceOk() (*string, bool)`

GetAuthDnsLatencyListenOnSourceOk returns a tuple with the AuthDnsLatencyListenOnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyListenOnSource

`func (o *MemberTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyListenOnSource(v string)`

SetAuthDnsLatencyListenOnSource sets AuthDnsLatencyListenOnSource field to given value.

### HasAuthDnsLatencyListenOnSource

`func (o *MemberTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyListenOnSource() bool`

HasAuthDnsLatencyListenOnSource returns a boolean if a field has been set.

### GetAuthDnsLatencyListenOnIp

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyListenOnIp() string`

GetAuthDnsLatencyListenOnIp returns the AuthDnsLatencyListenOnIp field if non-nil, zero value otherwise.

### GetAuthDnsLatencyListenOnIpOk

`func (o *MemberTrafficCaptureAuthDnsSetting) GetAuthDnsLatencyListenOnIpOk() (*string, bool)`

GetAuthDnsLatencyListenOnIpOk returns a tuple with the AuthDnsLatencyListenOnIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsLatencyListenOnIp

`func (o *MemberTrafficCaptureAuthDnsSetting) SetAuthDnsLatencyListenOnIp(v string)`

SetAuthDnsLatencyListenOnIp sets AuthDnsLatencyListenOnIp field to given value.

### HasAuthDnsLatencyListenOnIp

`func (o *MemberTrafficCaptureAuthDnsSetting) HasAuthDnsLatencyListenOnIp() bool`

HasAuthDnsLatencyListenOnIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


