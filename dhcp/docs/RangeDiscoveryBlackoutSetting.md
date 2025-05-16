# RangeDiscoveryBlackoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBlackout** | Pointer to **bool** | Determines whether a blackout is enabled or not. | [optional] 
**BlackoutDuration** | Pointer to **int64** | The blackout duration in seconds; minimum value is 1 minute. | [optional] 
**BlackoutSchedule** | Pointer to [**RangediscoveryblackoutsettingBlackoutSchedule**](RangediscoveryblackoutsettingBlackoutSchedule.md) |  | [optional] 

## Methods

### NewRangeDiscoveryBlackoutSetting

`func NewRangeDiscoveryBlackoutSetting() *RangeDiscoveryBlackoutSetting`

NewRangeDiscoveryBlackoutSetting instantiates a new RangeDiscoveryBlackoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeDiscoveryBlackoutSettingWithDefaults

`func NewRangeDiscoveryBlackoutSettingWithDefaults() *RangeDiscoveryBlackoutSetting`

NewRangeDiscoveryBlackoutSettingWithDefaults instantiates a new RangeDiscoveryBlackoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableBlackout

`func (o *RangeDiscoveryBlackoutSetting) GetEnableBlackout() bool`

GetEnableBlackout returns the EnableBlackout field if non-nil, zero value otherwise.

### GetEnableBlackoutOk

`func (o *RangeDiscoveryBlackoutSetting) GetEnableBlackoutOk() (*bool, bool)`

GetEnableBlackoutOk returns a tuple with the EnableBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackout

`func (o *RangeDiscoveryBlackoutSetting) SetEnableBlackout(v bool)`

SetEnableBlackout sets EnableBlackout field to given value.

### HasEnableBlackout

`func (o *RangeDiscoveryBlackoutSetting) HasEnableBlackout() bool`

HasEnableBlackout returns a boolean if a field has been set.

### GetBlackoutDuration

`func (o *RangeDiscoveryBlackoutSetting) GetBlackoutDuration() int64`

GetBlackoutDuration returns the BlackoutDuration field if non-nil, zero value otherwise.

### GetBlackoutDurationOk

`func (o *RangeDiscoveryBlackoutSetting) GetBlackoutDurationOk() (*int64, bool)`

GetBlackoutDurationOk returns a tuple with the BlackoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDuration

`func (o *RangeDiscoveryBlackoutSetting) SetBlackoutDuration(v int64)`

SetBlackoutDuration sets BlackoutDuration field to given value.

### HasBlackoutDuration

`func (o *RangeDiscoveryBlackoutSetting) HasBlackoutDuration() bool`

HasBlackoutDuration returns a boolean if a field has been set.

### GetBlackoutSchedule

`func (o *RangeDiscoveryBlackoutSetting) GetBlackoutSchedule() RangediscoveryblackoutsettingBlackoutSchedule`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *RangeDiscoveryBlackoutSetting) GetBlackoutScheduleOk() (*RangediscoveryblackoutsettingBlackoutSchedule, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *RangeDiscoveryBlackoutSetting) SetBlackoutSchedule(v RangediscoveryblackoutsettingBlackoutSchedule)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *RangeDiscoveryBlackoutSetting) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


