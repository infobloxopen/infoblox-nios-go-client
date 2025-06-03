# GridPasswordSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordMinLength** | Pointer to **int64** | The minimum length of the password. | [optional] 
**NumLowerChar** | Pointer to **int64** | The minimum number of lowercase characters. | [optional] 
**NumUpperChar** | Pointer to **int64** | The minimum number of uppercase characters. | [optional] 
**NumNumericChar** | Pointer to **int64** | The minimum number of numeric characters. | [optional] 
**NumSymbolChar** | Pointer to **int64** | The minimum number of symbol characters. The allowed characters are ! @ # $ % ^ &amp; * ( ). | [optional] 
**CharsToChange** | Pointer to **int64** | The minimum number of characters that must be changed when revising an admin password. | [optional] 
**ExpireDays** | Pointer to **int64** | The number of days of the password expiration period (if enabled). | [optional] 
**ReminderDays** | Pointer to **int64** | The number of days before the password expiration date when the appliance sends a reminder. | [optional] 
**ForceResetEnable** | Pointer to **bool** | If set to True, all new users must change their passwords when they first log in to the system, and existing users must change the passwords that were just reset. | [optional] 
**ExpireEnable** | Pointer to **bool** | If set to True, password expiration is enabled. | [optional] 
**HistoryEnable** | Pointer to **bool** | Enable/disable the password history. | [optional] 
**NumPasswordsSaved** | Pointer to **int64** | Number of saved passwords if password history is enabled. Can be set between 1 to 20. | [optional] 
**MinPasswordAge** | Pointer to **int64** | Minimum password age in days before password can be updated. Can be set between 1 to 9998 days. | [optional] 

## Methods

### NewGridPasswordSetting

`func NewGridPasswordSetting() *GridPasswordSetting`

NewGridPasswordSetting instantiates a new GridPasswordSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridPasswordSettingWithDefaults

`func NewGridPasswordSettingWithDefaults() *GridPasswordSetting`

NewGridPasswordSettingWithDefaults instantiates a new GridPasswordSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordMinLength

`func (o *GridPasswordSetting) GetPasswordMinLength() int64`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *GridPasswordSetting) GetPasswordMinLengthOk() (*int64, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *GridPasswordSetting) SetPasswordMinLength(v int64)`

SetPasswordMinLength sets PasswordMinLength field to given value.

### HasPasswordMinLength

`func (o *GridPasswordSetting) HasPasswordMinLength() bool`

HasPasswordMinLength returns a boolean if a field has been set.

### GetNumLowerChar

`func (o *GridPasswordSetting) GetNumLowerChar() int64`

GetNumLowerChar returns the NumLowerChar field if non-nil, zero value otherwise.

### GetNumLowerCharOk

`func (o *GridPasswordSetting) GetNumLowerCharOk() (*int64, bool)`

GetNumLowerCharOk returns a tuple with the NumLowerChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLowerChar

`func (o *GridPasswordSetting) SetNumLowerChar(v int64)`

SetNumLowerChar sets NumLowerChar field to given value.

### HasNumLowerChar

`func (o *GridPasswordSetting) HasNumLowerChar() bool`

HasNumLowerChar returns a boolean if a field has been set.

### GetNumUpperChar

`func (o *GridPasswordSetting) GetNumUpperChar() int64`

GetNumUpperChar returns the NumUpperChar field if non-nil, zero value otherwise.

### GetNumUpperCharOk

`func (o *GridPasswordSetting) GetNumUpperCharOk() (*int64, bool)`

GetNumUpperCharOk returns a tuple with the NumUpperChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpperChar

`func (o *GridPasswordSetting) SetNumUpperChar(v int64)`

SetNumUpperChar sets NumUpperChar field to given value.

### HasNumUpperChar

`func (o *GridPasswordSetting) HasNumUpperChar() bool`

HasNumUpperChar returns a boolean if a field has been set.

### GetNumNumericChar

`func (o *GridPasswordSetting) GetNumNumericChar() int64`

GetNumNumericChar returns the NumNumericChar field if non-nil, zero value otherwise.

### GetNumNumericCharOk

`func (o *GridPasswordSetting) GetNumNumericCharOk() (*int64, bool)`

GetNumNumericCharOk returns a tuple with the NumNumericChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNumericChar

`func (o *GridPasswordSetting) SetNumNumericChar(v int64)`

SetNumNumericChar sets NumNumericChar field to given value.

### HasNumNumericChar

`func (o *GridPasswordSetting) HasNumNumericChar() bool`

HasNumNumericChar returns a boolean if a field has been set.

### GetNumSymbolChar

`func (o *GridPasswordSetting) GetNumSymbolChar() int64`

GetNumSymbolChar returns the NumSymbolChar field if non-nil, zero value otherwise.

### GetNumSymbolCharOk

`func (o *GridPasswordSetting) GetNumSymbolCharOk() (*int64, bool)`

GetNumSymbolCharOk returns a tuple with the NumSymbolChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSymbolChar

`func (o *GridPasswordSetting) SetNumSymbolChar(v int64)`

SetNumSymbolChar sets NumSymbolChar field to given value.

### HasNumSymbolChar

`func (o *GridPasswordSetting) HasNumSymbolChar() bool`

HasNumSymbolChar returns a boolean if a field has been set.

### GetCharsToChange

`func (o *GridPasswordSetting) GetCharsToChange() int64`

GetCharsToChange returns the CharsToChange field if non-nil, zero value otherwise.

### GetCharsToChangeOk

`func (o *GridPasswordSetting) GetCharsToChangeOk() (*int64, bool)`

GetCharsToChangeOk returns a tuple with the CharsToChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharsToChange

`func (o *GridPasswordSetting) SetCharsToChange(v int64)`

SetCharsToChange sets CharsToChange field to given value.

### HasCharsToChange

`func (o *GridPasswordSetting) HasCharsToChange() bool`

HasCharsToChange returns a boolean if a field has been set.

### GetExpireDays

`func (o *GridPasswordSetting) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *GridPasswordSetting) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *GridPasswordSetting) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *GridPasswordSetting) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetReminderDays

`func (o *GridPasswordSetting) GetReminderDays() int64`

GetReminderDays returns the ReminderDays field if non-nil, zero value otherwise.

### GetReminderDaysOk

`func (o *GridPasswordSetting) GetReminderDaysOk() (*int64, bool)`

GetReminderDaysOk returns a tuple with the ReminderDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDays

`func (o *GridPasswordSetting) SetReminderDays(v int64)`

SetReminderDays sets ReminderDays field to given value.

### HasReminderDays

`func (o *GridPasswordSetting) HasReminderDays() bool`

HasReminderDays returns a boolean if a field has been set.

### GetForceResetEnable

`func (o *GridPasswordSetting) GetForceResetEnable() bool`

GetForceResetEnable returns the ForceResetEnable field if non-nil, zero value otherwise.

### GetForceResetEnableOk

`func (o *GridPasswordSetting) GetForceResetEnableOk() (*bool, bool)`

GetForceResetEnableOk returns a tuple with the ForceResetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceResetEnable

`func (o *GridPasswordSetting) SetForceResetEnable(v bool)`

SetForceResetEnable sets ForceResetEnable field to given value.

### HasForceResetEnable

`func (o *GridPasswordSetting) HasForceResetEnable() bool`

HasForceResetEnable returns a boolean if a field has been set.

### GetExpireEnable

`func (o *GridPasswordSetting) GetExpireEnable() bool`

GetExpireEnable returns the ExpireEnable field if non-nil, zero value otherwise.

### GetExpireEnableOk

`func (o *GridPasswordSetting) GetExpireEnableOk() (*bool, bool)`

GetExpireEnableOk returns a tuple with the ExpireEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireEnable

`func (o *GridPasswordSetting) SetExpireEnable(v bool)`

SetExpireEnable sets ExpireEnable field to given value.

### HasExpireEnable

`func (o *GridPasswordSetting) HasExpireEnable() bool`

HasExpireEnable returns a boolean if a field has been set.

### GetHistoryEnable

`func (o *GridPasswordSetting) GetHistoryEnable() bool`

GetHistoryEnable returns the HistoryEnable field if non-nil, zero value otherwise.

### GetHistoryEnableOk

`func (o *GridPasswordSetting) GetHistoryEnableOk() (*bool, bool)`

GetHistoryEnableOk returns a tuple with the HistoryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryEnable

`func (o *GridPasswordSetting) SetHistoryEnable(v bool)`

SetHistoryEnable sets HistoryEnable field to given value.

### HasHistoryEnable

`func (o *GridPasswordSetting) HasHistoryEnable() bool`

HasHistoryEnable returns a boolean if a field has been set.

### GetNumPasswordsSaved

`func (o *GridPasswordSetting) GetNumPasswordsSaved() int64`

GetNumPasswordsSaved returns the NumPasswordsSaved field if non-nil, zero value otherwise.

### GetNumPasswordsSavedOk

`func (o *GridPasswordSetting) GetNumPasswordsSavedOk() (*int64, bool)`

GetNumPasswordsSavedOk returns a tuple with the NumPasswordsSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPasswordsSaved

`func (o *GridPasswordSetting) SetNumPasswordsSaved(v int64)`

SetNumPasswordsSaved sets NumPasswordsSaved field to given value.

### HasNumPasswordsSaved

`func (o *GridPasswordSetting) HasNumPasswordsSaved() bool`

HasNumPasswordsSaved returns a boolean if a field has been set.

### GetMinPasswordAge

`func (o *GridPasswordSetting) GetMinPasswordAge() int64`

GetMinPasswordAge returns the MinPasswordAge field if non-nil, zero value otherwise.

### GetMinPasswordAgeOk

`func (o *GridPasswordSetting) GetMinPasswordAgeOk() (*int64, bool)`

GetMinPasswordAgeOk returns a tuple with the MinPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPasswordAge

`func (o *GridPasswordSetting) SetMinPasswordAge(v int64)`

SetMinPasswordAge sets MinPasswordAge field to given value.

### HasMinPasswordAge

`func (o *GridPasswordSetting) HasMinPasswordAge() bool`

HasMinPasswordAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


