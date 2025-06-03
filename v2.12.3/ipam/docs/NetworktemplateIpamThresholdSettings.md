# NetworktemplateIpamThresholdSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerValue** | Pointer to **int64** | Indicates the percentage point which triggers the email/SNMP trap sending. | [optional] 
**ResetValue** | Pointer to **int64** | Indicates the percentage point which resets the email/SNMP trap sending. | [optional] 

## Methods

### NewNetworktemplateIpamThresholdSettings

`func NewNetworktemplateIpamThresholdSettings() *NetworktemplateIpamThresholdSettings`

NewNetworktemplateIpamThresholdSettings instantiates a new NetworktemplateIpamThresholdSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworktemplateIpamThresholdSettingsWithDefaults

`func NewNetworktemplateIpamThresholdSettingsWithDefaults() *NetworktemplateIpamThresholdSettings`

NewNetworktemplateIpamThresholdSettingsWithDefaults instantiates a new NetworktemplateIpamThresholdSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerValue

`func (o *NetworktemplateIpamThresholdSettings) GetTriggerValue() int64`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *NetworktemplateIpamThresholdSettings) GetTriggerValueOk() (*int64, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *NetworktemplateIpamThresholdSettings) SetTriggerValue(v int64)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *NetworktemplateIpamThresholdSettings) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetResetValue

`func (o *NetworktemplateIpamThresholdSettings) GetResetValue() int64`

GetResetValue returns the ResetValue field if non-nil, zero value otherwise.

### GetResetValueOk

`func (o *NetworktemplateIpamThresholdSettings) GetResetValueOk() (*int64, bool)`

GetResetValueOk returns a tuple with the ResetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValue

`func (o *NetworktemplateIpamThresholdSettings) SetResetValue(v int64)`

SetResetValue sets ResetValue field to given value.

### HasResetValue

`func (o *NetworktemplateIpamThresholdSettings) HasResetValue() bool`

HasResetValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


