# RangePortControlBlackoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBlackout** | Pointer to **bool** | Determines whether a blackout is enabled or not. | [optional] 
**BlackoutDuration** | Pointer to **int64** | The blackout duration in seconds; minimum value is 1 minute. | [optional] 
**BlackoutSchedule** | Pointer to [**RangeportcontrolblackoutsettingBlackoutSchedule**](RangeportcontrolblackoutsettingBlackoutSchedule.md) |  | [optional] 

## Methods

### NewRangePortControlBlackoutSetting

`func NewRangePortControlBlackoutSetting() *RangePortControlBlackoutSetting`

NewRangePortControlBlackoutSetting instantiates a new RangePortControlBlackoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangePortControlBlackoutSettingWithDefaults

`func NewRangePortControlBlackoutSettingWithDefaults() *RangePortControlBlackoutSetting`

NewRangePortControlBlackoutSettingWithDefaults instantiates a new RangePortControlBlackoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableBlackout

`func (o *RangePortControlBlackoutSetting) GetEnableBlackout() bool`

GetEnableBlackout returns the EnableBlackout field if non-nil, zero value otherwise.

### GetEnableBlackoutOk

`func (o *RangePortControlBlackoutSetting) GetEnableBlackoutOk() (*bool, bool)`

GetEnableBlackoutOk returns a tuple with the EnableBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackout

`func (o *RangePortControlBlackoutSetting) SetEnableBlackout(v bool)`

SetEnableBlackout sets EnableBlackout field to given value.

### HasEnableBlackout

`func (o *RangePortControlBlackoutSetting) HasEnableBlackout() bool`

HasEnableBlackout returns a boolean if a field has been set.

### GetBlackoutDuration

`func (o *RangePortControlBlackoutSetting) GetBlackoutDuration() int64`

GetBlackoutDuration returns the BlackoutDuration field if non-nil, zero value otherwise.

### GetBlackoutDurationOk

`func (o *RangePortControlBlackoutSetting) GetBlackoutDurationOk() (*int64, bool)`

GetBlackoutDurationOk returns a tuple with the BlackoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDuration

`func (o *RangePortControlBlackoutSetting) SetBlackoutDuration(v int64)`

SetBlackoutDuration sets BlackoutDuration field to given value.

### HasBlackoutDuration

`func (o *RangePortControlBlackoutSetting) HasBlackoutDuration() bool`

HasBlackoutDuration returns a boolean if a field has been set.

### GetBlackoutSchedule

`func (o *RangePortControlBlackoutSetting) GetBlackoutSchedule() RangeportcontrolblackoutsettingBlackoutSchedule`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *RangePortControlBlackoutSetting) GetBlackoutScheduleOk() (*RangeportcontrolblackoutsettingBlackoutSchedule, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *RangePortControlBlackoutSetting) SetBlackoutSchedule(v RangeportcontrolblackoutsettingBlackoutSchedule)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *RangePortControlBlackoutSetting) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


