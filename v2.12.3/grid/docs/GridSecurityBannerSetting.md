# GridSecurityBannerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | The security level color. | [optional] 
**Level** | Pointer to **string** | The security level. | [optional] 
**Message** | Pointer to **string** | The classification message to be displayed. | [optional] 
**Enable** | Pointer to **bool** | If set to True, the security banner will be displayed on the header and footer of the Grid Manager screen, including the Login screen. | [optional] 

## Methods

### NewGridSecurityBannerSetting

`func NewGridSecurityBannerSetting() *GridSecurityBannerSetting`

NewGridSecurityBannerSetting instantiates a new GridSecurityBannerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridSecurityBannerSettingWithDefaults

`func NewGridSecurityBannerSettingWithDefaults() *GridSecurityBannerSetting`

NewGridSecurityBannerSettingWithDefaults instantiates a new GridSecurityBannerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *GridSecurityBannerSetting) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GridSecurityBannerSetting) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GridSecurityBannerSetting) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *GridSecurityBannerSetting) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLevel

`func (o *GridSecurityBannerSetting) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GridSecurityBannerSetting) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GridSecurityBannerSetting) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GridSecurityBannerSetting) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *GridSecurityBannerSetting) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GridSecurityBannerSetting) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GridSecurityBannerSetting) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GridSecurityBannerSetting) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnable

`func (o *GridSecurityBannerSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GridSecurityBannerSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GridSecurityBannerSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GridSecurityBannerSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


