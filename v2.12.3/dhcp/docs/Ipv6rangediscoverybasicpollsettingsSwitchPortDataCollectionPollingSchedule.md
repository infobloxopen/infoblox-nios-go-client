# Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weekdays** | Pointer to **[]string** | Days of the week when scheduling is triggered. | [optional] 
**TimeZone** | Pointer to **string** | The time zone for the schedule. | [optional] 
**RecurringTime** | Pointer to **int64** | The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define. | [optional] 
**Frequency** | Pointer to **string** | The frequency for the scheduled task. | [optional] 
**Every** | Pointer to **int64** | The number of frequency to wait before repeating the scheduled task. | [optional] 
**MinutesPastHour** | Pointer to **int64** | The minutes past the hour for the scheduled task. | [optional] 
**HourOfDay** | Pointer to **int64** | The hour of day for the scheduled task. | [optional] 
**Year** | Pointer to **int64** | The year for the scheduled task. | [optional] 
**Month** | Pointer to **int64** | The month for the scheduled task. | [optional] 
**DayOfMonth** | Pointer to **int64** | The day of the month for the scheduled task. | [optional] 
**Repeat** | Pointer to **string** | Indicates if the scheduled task will be repeated or run only once. | [optional] 
**Disable** | Pointer to **bool** | If set to True, the scheduled task is disabled. | [optional] 

## Methods

### NewIpv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule

`func NewIpv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule() *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

NewIpv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule instantiates a new Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleWithDefaults

`func NewIpv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleWithDefaults() *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

NewIpv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleWithDefaults instantiates a new Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeekdays

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetTimeZone

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRecurringTime

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetRecurringTime() int64`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetRecurringTimeOk() (*int64, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetRecurringTime(v int64)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetFrequency

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetEvery

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetEvery() int64`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetEveryOk() (*int64, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetEvery(v int64)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetMinutesPastHour

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetMinutesPastHour() int64`

GetMinutesPastHour returns the MinutesPastHour field if non-nil, zero value otherwise.

### GetMinutesPastHourOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetMinutesPastHourOk() (*int64, bool)`

GetMinutesPastHourOk returns a tuple with the MinutesPastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesPastHour

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetMinutesPastHour(v int64)`

SetMinutesPastHour sets MinutesPastHour field to given value.

### HasMinutesPastHour

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasMinutesPastHour() bool`

HasMinutesPastHour returns a boolean if a field has been set.

### GetHourOfDay

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetHourOfDay() int64`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetHourOfDayOk() (*int64, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetHourOfDay(v int64)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.

### GetYear

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetMonth() int64`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetMonthOk() (*int64, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetMonth(v int64)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetDayOfMonth() int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetDayOfMonthOk() (*int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetDayOfMonth(v int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetRepeat

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetRepeat() string`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetRepeatOk() (*string, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetRepeat(v string)`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### GetDisable

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Ipv6rangediscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) HasDisable() bool`

HasDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


