# GridRestartBannerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If set to True, the restart banner is enabled. | [optional] 
**EnableDoubleConfirmation** | Pointer to **bool** | If set to True, the user is required to input name before restarting the services. | [optional] 

## Methods

### NewGridRestartBannerSetting

`func NewGridRestartBannerSetting() *GridRestartBannerSetting`

NewGridRestartBannerSetting instantiates a new GridRestartBannerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridRestartBannerSettingWithDefaults

`func NewGridRestartBannerSettingWithDefaults() *GridRestartBannerSetting`

NewGridRestartBannerSettingWithDefaults instantiates a new GridRestartBannerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GridRestartBannerSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GridRestartBannerSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GridRestartBannerSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GridRestartBannerSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnableDoubleConfirmation

`func (o *GridRestartBannerSetting) GetEnableDoubleConfirmation() bool`

GetEnableDoubleConfirmation returns the EnableDoubleConfirmation field if non-nil, zero value otherwise.

### GetEnableDoubleConfirmationOk

`func (o *GridRestartBannerSetting) GetEnableDoubleConfirmationOk() (*bool, bool)`

GetEnableDoubleConfirmationOk returns a tuple with the EnableDoubleConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDoubleConfirmation

`func (o *GridRestartBannerSetting) SetEnableDoubleConfirmation(v bool)`

SetEnableDoubleConfirmation sets EnableDoubleConfirmation field to given value.

### HasEnableDoubleConfirmation

`func (o *GridRestartBannerSetting) HasEnableDoubleConfirmation() bool`

HasEnableDoubleConfirmation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


