# DiscoveryGridpropertiesPortControlBlackoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBlackout** | Pointer to **bool** | Determines whether a blackout is enabled or not. | [optional] 
**BlackoutDuration** | Pointer to **int64** | The blackout duration in seconds; minimum value is 1 minute. | [optional] 
**BlackoutSchedule** | Pointer to [**DiscoverygridpropertiesportcontrolblackoutsettingBlackoutSchedule**](DiscoverygridpropertiesportcontrolblackoutsettingBlackoutSchedule.md) |  | [optional] 

## Methods

### NewDiscoveryGridpropertiesPortControlBlackoutSetting

`func NewDiscoveryGridpropertiesPortControlBlackoutSetting() *DiscoveryGridpropertiesPortControlBlackoutSetting`

NewDiscoveryGridpropertiesPortControlBlackoutSetting instantiates a new DiscoveryGridpropertiesPortControlBlackoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesPortControlBlackoutSettingWithDefaults

`func NewDiscoveryGridpropertiesPortControlBlackoutSettingWithDefaults() *DiscoveryGridpropertiesPortControlBlackoutSetting`

NewDiscoveryGridpropertiesPortControlBlackoutSettingWithDefaults instantiates a new DiscoveryGridpropertiesPortControlBlackoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableBlackout

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) GetEnableBlackout() bool`

GetEnableBlackout returns the EnableBlackout field if non-nil, zero value otherwise.

### GetEnableBlackoutOk

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) GetEnableBlackoutOk() (*bool, bool)`

GetEnableBlackoutOk returns a tuple with the EnableBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackout

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) SetEnableBlackout(v bool)`

SetEnableBlackout sets EnableBlackout field to given value.

### HasEnableBlackout

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) HasEnableBlackout() bool`

HasEnableBlackout returns a boolean if a field has been set.

### GetBlackoutDuration

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) GetBlackoutDuration() int64`

GetBlackoutDuration returns the BlackoutDuration field if non-nil, zero value otherwise.

### GetBlackoutDurationOk

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) GetBlackoutDurationOk() (*int64, bool)`

GetBlackoutDurationOk returns a tuple with the BlackoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDuration

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) SetBlackoutDuration(v int64)`

SetBlackoutDuration sets BlackoutDuration field to given value.

### HasBlackoutDuration

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) HasBlackoutDuration() bool`

HasBlackoutDuration returns a boolean if a field has been set.

### GetBlackoutSchedule

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) GetBlackoutSchedule() DiscoverygridpropertiesportcontrolblackoutsettingBlackoutSchedule`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) GetBlackoutScheduleOk() (*DiscoverygridpropertiesportcontrolblackoutsettingBlackoutSchedule, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) SetBlackoutSchedule(v DiscoverygridpropertiesportcontrolblackoutsettingBlackoutSchedule)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *DiscoveryGridpropertiesPortControlBlackoutSetting) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


