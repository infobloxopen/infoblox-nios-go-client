# AdmingroupPasswordSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireEnable** | Pointer to **bool** | Whether password expiry enabled or not. | [optional] 
**ExpireDays** | Pointer to **int64** | The days that password must expire | [optional] 
**ReminderDays** | Pointer to **int64** | Days to show up reminder prior to expiration | [optional] 

## Methods

### NewAdmingroupPasswordSetting

`func NewAdmingroupPasswordSetting() *AdmingroupPasswordSetting`

NewAdmingroupPasswordSetting instantiates a new AdmingroupPasswordSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupPasswordSettingWithDefaults

`func NewAdmingroupPasswordSettingWithDefaults() *AdmingroupPasswordSetting`

NewAdmingroupPasswordSettingWithDefaults instantiates a new AdmingroupPasswordSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireEnable

`func (o *AdmingroupPasswordSetting) GetExpireEnable() bool`

GetExpireEnable returns the ExpireEnable field if non-nil, zero value otherwise.

### GetExpireEnableOk

`func (o *AdmingroupPasswordSetting) GetExpireEnableOk() (*bool, bool)`

GetExpireEnableOk returns a tuple with the ExpireEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireEnable

`func (o *AdmingroupPasswordSetting) SetExpireEnable(v bool)`

SetExpireEnable sets ExpireEnable field to given value.

### HasExpireEnable

`func (o *AdmingroupPasswordSetting) HasExpireEnable() bool`

HasExpireEnable returns a boolean if a field has been set.

### GetExpireDays

`func (o *AdmingroupPasswordSetting) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *AdmingroupPasswordSetting) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *AdmingroupPasswordSetting) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *AdmingroupPasswordSetting) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetReminderDays

`func (o *AdmingroupPasswordSetting) GetReminderDays() int64`

GetReminderDays returns the ReminderDays field if non-nil, zero value otherwise.

### GetReminderDaysOk

`func (o *AdmingroupPasswordSetting) GetReminderDaysOk() (*int64, bool)`

GetReminderDaysOk returns a tuple with the ReminderDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDays

`func (o *AdmingroupPasswordSetting) SetReminderDays(v int64)`

SetReminderDays sets ReminderDays field to given value.

### HasReminderDays

`func (o *AdmingroupPasswordSetting) HasReminderDays() bool`

HasReminderDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


