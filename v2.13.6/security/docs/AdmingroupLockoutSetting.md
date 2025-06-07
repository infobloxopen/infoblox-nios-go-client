# AdmingroupLockoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSequentialFailedLoginAttemptsLockout** | Pointer to **bool** | Enable/disable sequential failed login attempts lockout for local users | [optional] 
**SequentialAttempts** | Pointer to **int64** | The number of failed login attempts | [optional] 
**FailedLockoutDuration** | Pointer to **int64** | Time period the account remains locked after sequential failed login attempt lockout. | [optional] 
**NeverUnlockUser** | Pointer to **bool** | Never unlock option is also provided and if set then user account is locked forever and only super user can unlock this account | [optional] 

## Methods

### NewAdmingroupLockoutSetting

`func NewAdmingroupLockoutSetting() *AdmingroupLockoutSetting`

NewAdmingroupLockoutSetting instantiates a new AdmingroupLockoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupLockoutSettingWithDefaults

`func NewAdmingroupLockoutSettingWithDefaults() *AdmingroupLockoutSetting`

NewAdmingroupLockoutSettingWithDefaults instantiates a new AdmingroupLockoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSequentialFailedLoginAttemptsLockout

`func (o *AdmingroupLockoutSetting) GetEnableSequentialFailedLoginAttemptsLockout() bool`

GetEnableSequentialFailedLoginAttemptsLockout returns the EnableSequentialFailedLoginAttemptsLockout field if non-nil, zero value otherwise.

### GetEnableSequentialFailedLoginAttemptsLockoutOk

`func (o *AdmingroupLockoutSetting) GetEnableSequentialFailedLoginAttemptsLockoutOk() (*bool, bool)`

GetEnableSequentialFailedLoginAttemptsLockoutOk returns a tuple with the EnableSequentialFailedLoginAttemptsLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSequentialFailedLoginAttemptsLockout

`func (o *AdmingroupLockoutSetting) SetEnableSequentialFailedLoginAttemptsLockout(v bool)`

SetEnableSequentialFailedLoginAttemptsLockout sets EnableSequentialFailedLoginAttemptsLockout field to given value.

### HasEnableSequentialFailedLoginAttemptsLockout

`func (o *AdmingroupLockoutSetting) HasEnableSequentialFailedLoginAttemptsLockout() bool`

HasEnableSequentialFailedLoginAttemptsLockout returns a boolean if a field has been set.

### GetSequentialAttempts

`func (o *AdmingroupLockoutSetting) GetSequentialAttempts() int64`

GetSequentialAttempts returns the SequentialAttempts field if non-nil, zero value otherwise.

### GetSequentialAttemptsOk

`func (o *AdmingroupLockoutSetting) GetSequentialAttemptsOk() (*int64, bool)`

GetSequentialAttemptsOk returns a tuple with the SequentialAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialAttempts

`func (o *AdmingroupLockoutSetting) SetSequentialAttempts(v int64)`

SetSequentialAttempts sets SequentialAttempts field to given value.

### HasSequentialAttempts

`func (o *AdmingroupLockoutSetting) HasSequentialAttempts() bool`

HasSequentialAttempts returns a boolean if a field has been set.

### GetFailedLockoutDuration

`func (o *AdmingroupLockoutSetting) GetFailedLockoutDuration() int64`

GetFailedLockoutDuration returns the FailedLockoutDuration field if non-nil, zero value otherwise.

### GetFailedLockoutDurationOk

`func (o *AdmingroupLockoutSetting) GetFailedLockoutDurationOk() (*int64, bool)`

GetFailedLockoutDurationOk returns a tuple with the FailedLockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLockoutDuration

`func (o *AdmingroupLockoutSetting) SetFailedLockoutDuration(v int64)`

SetFailedLockoutDuration sets FailedLockoutDuration field to given value.

### HasFailedLockoutDuration

`func (o *AdmingroupLockoutSetting) HasFailedLockoutDuration() bool`

HasFailedLockoutDuration returns a boolean if a field has been set.

### GetNeverUnlockUser

`func (o *AdmingroupLockoutSetting) GetNeverUnlockUser() bool`

GetNeverUnlockUser returns the NeverUnlockUser field if non-nil, zero value otherwise.

### GetNeverUnlockUserOk

`func (o *AdmingroupLockoutSetting) GetNeverUnlockUserOk() (*bool, bool)`

GetNeverUnlockUserOk returns a tuple with the NeverUnlockUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverUnlockUser

`func (o *AdmingroupLockoutSetting) SetNeverUnlockUser(v bool)`

SetNeverUnlockUser sets NeverUnlockUser field to given value.

### HasNeverUnlockUser

`func (o *AdmingroupLockoutSetting) HasNeverUnlockUser() bool`

HasNeverUnlockUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


