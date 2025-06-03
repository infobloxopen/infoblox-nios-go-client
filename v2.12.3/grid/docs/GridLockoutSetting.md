# GridLockoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSequentialFailedLoginAttemptsLockout** | Pointer to **bool** | Enable/disable sequential failed login attempts lockout for local users | [optional] 
**SequentialAttempts** | Pointer to **int64** | The number of failed login attempts | [optional] 
**FailedLockoutDuration** | Pointer to **int64** | Time period the account remains locked after sequential failed login attempt lockout. | [optional] 
**NeverUnlockUser** | Pointer to **bool** | Never unlock option is also provided and if set then user account is locked forever and only super user can unlock this account | [optional] 

## Methods

### NewGridLockoutSetting

`func NewGridLockoutSetting() *GridLockoutSetting`

NewGridLockoutSetting instantiates a new GridLockoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridLockoutSettingWithDefaults

`func NewGridLockoutSettingWithDefaults() *GridLockoutSetting`

NewGridLockoutSettingWithDefaults instantiates a new GridLockoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSequentialFailedLoginAttemptsLockout

`func (o *GridLockoutSetting) GetEnableSequentialFailedLoginAttemptsLockout() bool`

GetEnableSequentialFailedLoginAttemptsLockout returns the EnableSequentialFailedLoginAttemptsLockout field if non-nil, zero value otherwise.

### GetEnableSequentialFailedLoginAttemptsLockoutOk

`func (o *GridLockoutSetting) GetEnableSequentialFailedLoginAttemptsLockoutOk() (*bool, bool)`

GetEnableSequentialFailedLoginAttemptsLockoutOk returns a tuple with the EnableSequentialFailedLoginAttemptsLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSequentialFailedLoginAttemptsLockout

`func (o *GridLockoutSetting) SetEnableSequentialFailedLoginAttemptsLockout(v bool)`

SetEnableSequentialFailedLoginAttemptsLockout sets EnableSequentialFailedLoginAttemptsLockout field to given value.

### HasEnableSequentialFailedLoginAttemptsLockout

`func (o *GridLockoutSetting) HasEnableSequentialFailedLoginAttemptsLockout() bool`

HasEnableSequentialFailedLoginAttemptsLockout returns a boolean if a field has been set.

### GetSequentialAttempts

`func (o *GridLockoutSetting) GetSequentialAttempts() int64`

GetSequentialAttempts returns the SequentialAttempts field if non-nil, zero value otherwise.

### GetSequentialAttemptsOk

`func (o *GridLockoutSetting) GetSequentialAttemptsOk() (*int64, bool)`

GetSequentialAttemptsOk returns a tuple with the SequentialAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialAttempts

`func (o *GridLockoutSetting) SetSequentialAttempts(v int64)`

SetSequentialAttempts sets SequentialAttempts field to given value.

### HasSequentialAttempts

`func (o *GridLockoutSetting) HasSequentialAttempts() bool`

HasSequentialAttempts returns a boolean if a field has been set.

### GetFailedLockoutDuration

`func (o *GridLockoutSetting) GetFailedLockoutDuration() int64`

GetFailedLockoutDuration returns the FailedLockoutDuration field if non-nil, zero value otherwise.

### GetFailedLockoutDurationOk

`func (o *GridLockoutSetting) GetFailedLockoutDurationOk() (*int64, bool)`

GetFailedLockoutDurationOk returns a tuple with the FailedLockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLockoutDuration

`func (o *GridLockoutSetting) SetFailedLockoutDuration(v int64)`

SetFailedLockoutDuration sets FailedLockoutDuration field to given value.

### HasFailedLockoutDuration

`func (o *GridLockoutSetting) HasFailedLockoutDuration() bool`

HasFailedLockoutDuration returns a boolean if a field has been set.

### GetNeverUnlockUser

`func (o *GridLockoutSetting) GetNeverUnlockUser() bool`

GetNeverUnlockUser returns the NeverUnlockUser field if non-nil, zero value otherwise.

### GetNeverUnlockUserOk

`func (o *GridLockoutSetting) GetNeverUnlockUserOk() (*bool, bool)`

GetNeverUnlockUserOk returns a tuple with the NeverUnlockUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverUnlockUser

`func (o *GridLockoutSetting) SetNeverUnlockUser(v bool)`

SetNeverUnlockUser sets NeverUnlockUser field to given value.

### HasNeverUnlockUser

`func (o *GridLockoutSetting) HasNeverUnlockUser() bool`

HasNeverUnlockUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


