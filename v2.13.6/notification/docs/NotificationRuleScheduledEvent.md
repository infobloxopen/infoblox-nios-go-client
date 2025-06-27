# NotificationRuleScheduledEvent

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

### NewNotificationRuleScheduledEvent

`func NewNotificationRuleScheduledEvent() *NotificationRuleScheduledEvent`

NewNotificationRuleScheduledEvent instantiates a new NotificationRuleScheduledEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleScheduledEventWithDefaults

`func NewNotificationRuleScheduledEventWithDefaults() *NotificationRuleScheduledEvent`

NewNotificationRuleScheduledEventWithDefaults instantiates a new NotificationRuleScheduledEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeekdays

`func (o *NotificationRuleScheduledEvent) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *NotificationRuleScheduledEvent) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *NotificationRuleScheduledEvent) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *NotificationRuleScheduledEvent) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetTimeZone

`func (o *NotificationRuleScheduledEvent) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *NotificationRuleScheduledEvent) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *NotificationRuleScheduledEvent) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *NotificationRuleScheduledEvent) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRecurringTime

`func (o *NotificationRuleScheduledEvent) GetRecurringTime() int64`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *NotificationRuleScheduledEvent) GetRecurringTimeOk() (*int64, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *NotificationRuleScheduledEvent) SetRecurringTime(v int64)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *NotificationRuleScheduledEvent) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetFrequency

`func (o *NotificationRuleScheduledEvent) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *NotificationRuleScheduledEvent) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *NotificationRuleScheduledEvent) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *NotificationRuleScheduledEvent) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetEvery

`func (o *NotificationRuleScheduledEvent) GetEvery() int64`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *NotificationRuleScheduledEvent) GetEveryOk() (*int64, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *NotificationRuleScheduledEvent) SetEvery(v int64)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *NotificationRuleScheduledEvent) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetMinutesPastHour

`func (o *NotificationRuleScheduledEvent) GetMinutesPastHour() int64`

GetMinutesPastHour returns the MinutesPastHour field if non-nil, zero value otherwise.

### GetMinutesPastHourOk

`func (o *NotificationRuleScheduledEvent) GetMinutesPastHourOk() (*int64, bool)`

GetMinutesPastHourOk returns a tuple with the MinutesPastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesPastHour

`func (o *NotificationRuleScheduledEvent) SetMinutesPastHour(v int64)`

SetMinutesPastHour sets MinutesPastHour field to given value.

### HasMinutesPastHour

`func (o *NotificationRuleScheduledEvent) HasMinutesPastHour() bool`

HasMinutesPastHour returns a boolean if a field has been set.

### GetHourOfDay

`func (o *NotificationRuleScheduledEvent) GetHourOfDay() int64`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *NotificationRuleScheduledEvent) GetHourOfDayOk() (*int64, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *NotificationRuleScheduledEvent) SetHourOfDay(v int64)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *NotificationRuleScheduledEvent) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.

### GetYear

`func (o *NotificationRuleScheduledEvent) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *NotificationRuleScheduledEvent) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *NotificationRuleScheduledEvent) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *NotificationRuleScheduledEvent) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *NotificationRuleScheduledEvent) GetMonth() int64`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *NotificationRuleScheduledEvent) GetMonthOk() (*int64, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *NotificationRuleScheduledEvent) SetMonth(v int64)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *NotificationRuleScheduledEvent) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *NotificationRuleScheduledEvent) GetDayOfMonth() int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *NotificationRuleScheduledEvent) GetDayOfMonthOk() (*int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *NotificationRuleScheduledEvent) SetDayOfMonth(v int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *NotificationRuleScheduledEvent) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetRepeat

`func (o *NotificationRuleScheduledEvent) GetRepeat() string`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *NotificationRuleScheduledEvent) GetRepeatOk() (*string, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *NotificationRuleScheduledEvent) SetRepeat(v string)`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *NotificationRuleScheduledEvent) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### GetDisable

`func (o *NotificationRuleScheduledEvent) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *NotificationRuleScheduledEvent) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *NotificationRuleScheduledEvent) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *NotificationRuleScheduledEvent) HasDisable() bool`

HasDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


