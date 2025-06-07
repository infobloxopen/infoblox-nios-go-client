# AdmingroupInactivityLockoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInactivityLockoutEnable** | Pointer to **bool** | Enable/disable the account inactivity lockout. | [optional] 
**InactiveDays** | Pointer to **int64** | Number of days after which account gets locked out if user does not login. | [optional] 
**ReminderDays** | Pointer to **int64** | The number of days before the account lockout date when the appliance sends a reminder. | [optional] 
**ReactivateViaSerialConsoleEnable** | Pointer to **bool** | Enable/disable reactivating user account by logging in from serial console. | [optional] 
**ReactivateViaRemoteConsoleEnable** | Pointer to **bool** | Enable/disable reactivating user account by logging in from remote console. | [optional] 

## Methods

### NewAdmingroupInactivityLockoutSetting

`func NewAdmingroupInactivityLockoutSetting() *AdmingroupInactivityLockoutSetting`

NewAdmingroupInactivityLockoutSetting instantiates a new AdmingroupInactivityLockoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupInactivityLockoutSettingWithDefaults

`func NewAdmingroupInactivityLockoutSettingWithDefaults() *AdmingroupInactivityLockoutSetting`

NewAdmingroupInactivityLockoutSettingWithDefaults instantiates a new AdmingroupInactivityLockoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInactivityLockoutEnable

`func (o *AdmingroupInactivityLockoutSetting) GetAccountInactivityLockoutEnable() bool`

GetAccountInactivityLockoutEnable returns the AccountInactivityLockoutEnable field if non-nil, zero value otherwise.

### GetAccountInactivityLockoutEnableOk

`func (o *AdmingroupInactivityLockoutSetting) GetAccountInactivityLockoutEnableOk() (*bool, bool)`

GetAccountInactivityLockoutEnableOk returns a tuple with the AccountInactivityLockoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInactivityLockoutEnable

`func (o *AdmingroupInactivityLockoutSetting) SetAccountInactivityLockoutEnable(v bool)`

SetAccountInactivityLockoutEnable sets AccountInactivityLockoutEnable field to given value.

### HasAccountInactivityLockoutEnable

`func (o *AdmingroupInactivityLockoutSetting) HasAccountInactivityLockoutEnable() bool`

HasAccountInactivityLockoutEnable returns a boolean if a field has been set.

### GetInactiveDays

`func (o *AdmingroupInactivityLockoutSetting) GetInactiveDays() int64`

GetInactiveDays returns the InactiveDays field if non-nil, zero value otherwise.

### GetInactiveDaysOk

`func (o *AdmingroupInactivityLockoutSetting) GetInactiveDaysOk() (*int64, bool)`

GetInactiveDaysOk returns a tuple with the InactiveDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDays

`func (o *AdmingroupInactivityLockoutSetting) SetInactiveDays(v int64)`

SetInactiveDays sets InactiveDays field to given value.

### HasInactiveDays

`func (o *AdmingroupInactivityLockoutSetting) HasInactiveDays() bool`

HasInactiveDays returns a boolean if a field has been set.

### GetReminderDays

`func (o *AdmingroupInactivityLockoutSetting) GetReminderDays() int64`

GetReminderDays returns the ReminderDays field if non-nil, zero value otherwise.

### GetReminderDaysOk

`func (o *AdmingroupInactivityLockoutSetting) GetReminderDaysOk() (*int64, bool)`

GetReminderDaysOk returns a tuple with the ReminderDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDays

`func (o *AdmingroupInactivityLockoutSetting) SetReminderDays(v int64)`

SetReminderDays sets ReminderDays field to given value.

### HasReminderDays

`func (o *AdmingroupInactivityLockoutSetting) HasReminderDays() bool`

HasReminderDays returns a boolean if a field has been set.

### GetReactivateViaSerialConsoleEnable

`func (o *AdmingroupInactivityLockoutSetting) GetReactivateViaSerialConsoleEnable() bool`

GetReactivateViaSerialConsoleEnable returns the ReactivateViaSerialConsoleEnable field if non-nil, zero value otherwise.

### GetReactivateViaSerialConsoleEnableOk

`func (o *AdmingroupInactivityLockoutSetting) GetReactivateViaSerialConsoleEnableOk() (*bool, bool)`

GetReactivateViaSerialConsoleEnableOk returns a tuple with the ReactivateViaSerialConsoleEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivateViaSerialConsoleEnable

`func (o *AdmingroupInactivityLockoutSetting) SetReactivateViaSerialConsoleEnable(v bool)`

SetReactivateViaSerialConsoleEnable sets ReactivateViaSerialConsoleEnable field to given value.

### HasReactivateViaSerialConsoleEnable

`func (o *AdmingroupInactivityLockoutSetting) HasReactivateViaSerialConsoleEnable() bool`

HasReactivateViaSerialConsoleEnable returns a boolean if a field has been set.

### GetReactivateViaRemoteConsoleEnable

`func (o *AdmingroupInactivityLockoutSetting) GetReactivateViaRemoteConsoleEnable() bool`

GetReactivateViaRemoteConsoleEnable returns the ReactivateViaRemoteConsoleEnable field if non-nil, zero value otherwise.

### GetReactivateViaRemoteConsoleEnableOk

`func (o *AdmingroupInactivityLockoutSetting) GetReactivateViaRemoteConsoleEnableOk() (*bool, bool)`

GetReactivateViaRemoteConsoleEnableOk returns a tuple with the ReactivateViaRemoteConsoleEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivateViaRemoteConsoleEnable

`func (o *AdmingroupInactivityLockoutSetting) SetReactivateViaRemoteConsoleEnable(v bool)`

SetReactivateViaRemoteConsoleEnable sets ReactivateViaRemoteConsoleEnable field to given value.

### HasReactivateViaRemoteConsoleEnable

`func (o *AdmingroupInactivityLockoutSetting) HasReactivateViaRemoteConsoleEnable() bool`

HasReactivateViaRemoteConsoleEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


