# NetworkcontainerPortControlBlackoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBlackout** | Pointer to **bool** | Determines whether a blackout is enabled or not. | [optional] 
**BlackoutDuration** | Pointer to **int64** | The blackout duration in seconds; minimum value is 1 minute. | [optional] 
**BlackoutSchedule** | Pointer to [**NetworkcontainerportcontrolblackoutsettingBlackoutSchedule**](NetworkcontainerportcontrolblackoutsettingBlackoutSchedule.md) |  | [optional] 

## Methods

### NewNetworkcontainerPortControlBlackoutSetting

`func NewNetworkcontainerPortControlBlackoutSetting() *NetworkcontainerPortControlBlackoutSetting`

NewNetworkcontainerPortControlBlackoutSetting instantiates a new NetworkcontainerPortControlBlackoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkcontainerPortControlBlackoutSettingWithDefaults

`func NewNetworkcontainerPortControlBlackoutSettingWithDefaults() *NetworkcontainerPortControlBlackoutSetting`

NewNetworkcontainerPortControlBlackoutSettingWithDefaults instantiates a new NetworkcontainerPortControlBlackoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableBlackout

`func (o *NetworkcontainerPortControlBlackoutSetting) GetEnableBlackout() bool`

GetEnableBlackout returns the EnableBlackout field if non-nil, zero value otherwise.

### GetEnableBlackoutOk

`func (o *NetworkcontainerPortControlBlackoutSetting) GetEnableBlackoutOk() (*bool, bool)`

GetEnableBlackoutOk returns a tuple with the EnableBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackout

`func (o *NetworkcontainerPortControlBlackoutSetting) SetEnableBlackout(v bool)`

SetEnableBlackout sets EnableBlackout field to given value.

### HasEnableBlackout

`func (o *NetworkcontainerPortControlBlackoutSetting) HasEnableBlackout() bool`

HasEnableBlackout returns a boolean if a field has been set.

### GetBlackoutDuration

`func (o *NetworkcontainerPortControlBlackoutSetting) GetBlackoutDuration() int64`

GetBlackoutDuration returns the BlackoutDuration field if non-nil, zero value otherwise.

### GetBlackoutDurationOk

`func (o *NetworkcontainerPortControlBlackoutSetting) GetBlackoutDurationOk() (*int64, bool)`

GetBlackoutDurationOk returns a tuple with the BlackoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDuration

`func (o *NetworkcontainerPortControlBlackoutSetting) SetBlackoutDuration(v int64)`

SetBlackoutDuration sets BlackoutDuration field to given value.

### HasBlackoutDuration

`func (o *NetworkcontainerPortControlBlackoutSetting) HasBlackoutDuration() bool`

HasBlackoutDuration returns a boolean if a field has been set.

### GetBlackoutSchedule

`func (o *NetworkcontainerPortControlBlackoutSetting) GetBlackoutSchedule() NetworkcontainerportcontrolblackoutsettingBlackoutSchedule`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *NetworkcontainerPortControlBlackoutSetting) GetBlackoutScheduleOk() (*NetworkcontainerportcontrolblackoutsettingBlackoutSchedule, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *NetworkcontainerPortControlBlackoutSetting) SetBlackoutSchedule(v NetworkcontainerportcontrolblackoutsettingBlackoutSchedule)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *NetworkcontainerPortControlBlackoutSetting) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


