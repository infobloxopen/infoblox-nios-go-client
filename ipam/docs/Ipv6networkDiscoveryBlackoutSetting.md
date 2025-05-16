# Ipv6networkDiscoveryBlackoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBlackout** | Pointer to **bool** | Determines whether a blackout is enabled or not. | [optional] 
**BlackoutDuration** | Pointer to **int64** | The blackout duration in seconds; minimum value is 1 minute. | [optional] 
**BlackoutSchedule** | Pointer to [**Ipv6networkdiscoveryblackoutsettingBlackoutSchedule**](Ipv6networkdiscoveryblackoutsettingBlackoutSchedule.md) |  | [optional] 

## Methods

### NewIpv6networkDiscoveryBlackoutSetting

`func NewIpv6networkDiscoveryBlackoutSetting() *Ipv6networkDiscoveryBlackoutSetting`

NewIpv6networkDiscoveryBlackoutSetting instantiates a new Ipv6networkDiscoveryBlackoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6networkDiscoveryBlackoutSettingWithDefaults

`func NewIpv6networkDiscoveryBlackoutSettingWithDefaults() *Ipv6networkDiscoveryBlackoutSetting`

NewIpv6networkDiscoveryBlackoutSettingWithDefaults instantiates a new Ipv6networkDiscoveryBlackoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableBlackout

`func (o *Ipv6networkDiscoveryBlackoutSetting) GetEnableBlackout() bool`

GetEnableBlackout returns the EnableBlackout field if non-nil, zero value otherwise.

### GetEnableBlackoutOk

`func (o *Ipv6networkDiscoveryBlackoutSetting) GetEnableBlackoutOk() (*bool, bool)`

GetEnableBlackoutOk returns a tuple with the EnableBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlackout

`func (o *Ipv6networkDiscoveryBlackoutSetting) SetEnableBlackout(v bool)`

SetEnableBlackout sets EnableBlackout field to given value.

### HasEnableBlackout

`func (o *Ipv6networkDiscoveryBlackoutSetting) HasEnableBlackout() bool`

HasEnableBlackout returns a boolean if a field has been set.

### GetBlackoutDuration

`func (o *Ipv6networkDiscoveryBlackoutSetting) GetBlackoutDuration() int64`

GetBlackoutDuration returns the BlackoutDuration field if non-nil, zero value otherwise.

### GetBlackoutDurationOk

`func (o *Ipv6networkDiscoveryBlackoutSetting) GetBlackoutDurationOk() (*int64, bool)`

GetBlackoutDurationOk returns a tuple with the BlackoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDuration

`func (o *Ipv6networkDiscoveryBlackoutSetting) SetBlackoutDuration(v int64)`

SetBlackoutDuration sets BlackoutDuration field to given value.

### HasBlackoutDuration

`func (o *Ipv6networkDiscoveryBlackoutSetting) HasBlackoutDuration() bool`

HasBlackoutDuration returns a boolean if a field has been set.

### GetBlackoutSchedule

`func (o *Ipv6networkDiscoveryBlackoutSetting) GetBlackoutSchedule() Ipv6networkdiscoveryblackoutsettingBlackoutSchedule`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *Ipv6networkDiscoveryBlackoutSetting) GetBlackoutScheduleOk() (*Ipv6networkdiscoveryblackoutsettingBlackoutSchedule, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *Ipv6networkDiscoveryBlackoutSetting) SetBlackoutSchedule(v Ipv6networkdiscoveryblackoutsettingBlackoutSchedule)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *Ipv6networkDiscoveryBlackoutSetting) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


