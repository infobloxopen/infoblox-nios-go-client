# GridSecuritySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogRollingEnable** | Pointer to **bool** | If set to True, rolling of audit logs is enabled. | [optional] 
**AdminAccessItems** | Pointer to [**[]GridsecuritysettingAdminAccessItems**](GridsecuritysettingAdminAccessItems.md) | A list of access control settings used for security access. | [optional] 
**HttpRedirectEnable** | Pointer to **bool** | If set to True, HTTP connections are redirected to HTTPS. | [optional] 
**LoginBannerEnable** | Pointer to **bool** | If set to True, the login banner is enabled. | [optional] 
**LoginBannerText** | Pointer to **string** | The login banner text. | [optional] 
**RemoteConsoleAccessEnable** | Pointer to **bool** | If set to True, superuser admins can access the Infoblox CLI from a remote location using an SSH (Secure Shell) v2 client. | [optional] 
**SecurityAccessEnable** | Pointer to **bool** | If set to True, HTTP access restrictions are enabled. | [optional] 
**SecurityAccessRemoteConsoleEnable** | Pointer to **bool** | If set to True, remote console access restrictions will be enabled. | [optional] 
**SessionTimeout** | Pointer to **int64** | The session timeout interval in seconds. | [optional] 
**SshPermEnable** | Pointer to **bool** | If set to False, SSH access is permanently disabled. | [optional] [readonly] 
**SupportAccessEnable** | Pointer to **bool** | If set to True, support access for the Grid has been enabled. | [optional] 
**SupportAccessInfo** | Pointer to **string** | Information string to be used for support access requests. | [optional] 
**DisableConcurrentLogin** | Pointer to **bool** | Whether concurrent login allowed gridlevel | [optional] 
**InactivityLockoutSetting** | Pointer to [**GridsecuritysettingInactivityLockoutSetting**](GridsecuritysettingInactivityLockoutSetting.md) |  | [optional] 

## Methods

### NewGridSecuritySetting

`func NewGridSecuritySetting() *GridSecuritySetting`

NewGridSecuritySetting instantiates a new GridSecuritySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridSecuritySettingWithDefaults

`func NewGridSecuritySettingWithDefaults() *GridSecuritySetting`

NewGridSecuritySettingWithDefaults instantiates a new GridSecuritySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogRollingEnable

`func (o *GridSecuritySetting) GetAuditLogRollingEnable() bool`

GetAuditLogRollingEnable returns the AuditLogRollingEnable field if non-nil, zero value otherwise.

### GetAuditLogRollingEnableOk

`func (o *GridSecuritySetting) GetAuditLogRollingEnableOk() (*bool, bool)`

GetAuditLogRollingEnableOk returns a tuple with the AuditLogRollingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogRollingEnable

`func (o *GridSecuritySetting) SetAuditLogRollingEnable(v bool)`

SetAuditLogRollingEnable sets AuditLogRollingEnable field to given value.

### HasAuditLogRollingEnable

`func (o *GridSecuritySetting) HasAuditLogRollingEnable() bool`

HasAuditLogRollingEnable returns a boolean if a field has been set.

### GetAdminAccessItems

`func (o *GridSecuritySetting) GetAdminAccessItems() []GridsecuritysettingAdminAccessItems`

GetAdminAccessItems returns the AdminAccessItems field if non-nil, zero value otherwise.

### GetAdminAccessItemsOk

`func (o *GridSecuritySetting) GetAdminAccessItemsOk() (*[]GridsecuritysettingAdminAccessItems, bool)`

GetAdminAccessItemsOk returns a tuple with the AdminAccessItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccessItems

`func (o *GridSecuritySetting) SetAdminAccessItems(v []GridsecuritysettingAdminAccessItems)`

SetAdminAccessItems sets AdminAccessItems field to given value.

### HasAdminAccessItems

`func (o *GridSecuritySetting) HasAdminAccessItems() bool`

HasAdminAccessItems returns a boolean if a field has been set.

### GetHttpRedirectEnable

`func (o *GridSecuritySetting) GetHttpRedirectEnable() bool`

GetHttpRedirectEnable returns the HttpRedirectEnable field if non-nil, zero value otherwise.

### GetHttpRedirectEnableOk

`func (o *GridSecuritySetting) GetHttpRedirectEnableOk() (*bool, bool)`

GetHttpRedirectEnableOk returns a tuple with the HttpRedirectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRedirectEnable

`func (o *GridSecuritySetting) SetHttpRedirectEnable(v bool)`

SetHttpRedirectEnable sets HttpRedirectEnable field to given value.

### HasHttpRedirectEnable

`func (o *GridSecuritySetting) HasHttpRedirectEnable() bool`

HasHttpRedirectEnable returns a boolean if a field has been set.

### GetLoginBannerEnable

`func (o *GridSecuritySetting) GetLoginBannerEnable() bool`

GetLoginBannerEnable returns the LoginBannerEnable field if non-nil, zero value otherwise.

### GetLoginBannerEnableOk

`func (o *GridSecuritySetting) GetLoginBannerEnableOk() (*bool, bool)`

GetLoginBannerEnableOk returns a tuple with the LoginBannerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBannerEnable

`func (o *GridSecuritySetting) SetLoginBannerEnable(v bool)`

SetLoginBannerEnable sets LoginBannerEnable field to given value.

### HasLoginBannerEnable

`func (o *GridSecuritySetting) HasLoginBannerEnable() bool`

HasLoginBannerEnable returns a boolean if a field has been set.

### GetLoginBannerText

`func (o *GridSecuritySetting) GetLoginBannerText() string`

GetLoginBannerText returns the LoginBannerText field if non-nil, zero value otherwise.

### GetLoginBannerTextOk

`func (o *GridSecuritySetting) GetLoginBannerTextOk() (*string, bool)`

GetLoginBannerTextOk returns a tuple with the LoginBannerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBannerText

`func (o *GridSecuritySetting) SetLoginBannerText(v string)`

SetLoginBannerText sets LoginBannerText field to given value.

### HasLoginBannerText

`func (o *GridSecuritySetting) HasLoginBannerText() bool`

HasLoginBannerText returns a boolean if a field has been set.

### GetRemoteConsoleAccessEnable

`func (o *GridSecuritySetting) GetRemoteConsoleAccessEnable() bool`

GetRemoteConsoleAccessEnable returns the RemoteConsoleAccessEnable field if non-nil, zero value otherwise.

### GetRemoteConsoleAccessEnableOk

`func (o *GridSecuritySetting) GetRemoteConsoleAccessEnableOk() (*bool, bool)`

GetRemoteConsoleAccessEnableOk returns a tuple with the RemoteConsoleAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConsoleAccessEnable

`func (o *GridSecuritySetting) SetRemoteConsoleAccessEnable(v bool)`

SetRemoteConsoleAccessEnable sets RemoteConsoleAccessEnable field to given value.

### HasRemoteConsoleAccessEnable

`func (o *GridSecuritySetting) HasRemoteConsoleAccessEnable() bool`

HasRemoteConsoleAccessEnable returns a boolean if a field has been set.

### GetSecurityAccessEnable

`func (o *GridSecuritySetting) GetSecurityAccessEnable() bool`

GetSecurityAccessEnable returns the SecurityAccessEnable field if non-nil, zero value otherwise.

### GetSecurityAccessEnableOk

`func (o *GridSecuritySetting) GetSecurityAccessEnableOk() (*bool, bool)`

GetSecurityAccessEnableOk returns a tuple with the SecurityAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAccessEnable

`func (o *GridSecuritySetting) SetSecurityAccessEnable(v bool)`

SetSecurityAccessEnable sets SecurityAccessEnable field to given value.

### HasSecurityAccessEnable

`func (o *GridSecuritySetting) HasSecurityAccessEnable() bool`

HasSecurityAccessEnable returns a boolean if a field has been set.

### GetSecurityAccessRemoteConsoleEnable

`func (o *GridSecuritySetting) GetSecurityAccessRemoteConsoleEnable() bool`

GetSecurityAccessRemoteConsoleEnable returns the SecurityAccessRemoteConsoleEnable field if non-nil, zero value otherwise.

### GetSecurityAccessRemoteConsoleEnableOk

`func (o *GridSecuritySetting) GetSecurityAccessRemoteConsoleEnableOk() (*bool, bool)`

GetSecurityAccessRemoteConsoleEnableOk returns a tuple with the SecurityAccessRemoteConsoleEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAccessRemoteConsoleEnable

`func (o *GridSecuritySetting) SetSecurityAccessRemoteConsoleEnable(v bool)`

SetSecurityAccessRemoteConsoleEnable sets SecurityAccessRemoteConsoleEnable field to given value.

### HasSecurityAccessRemoteConsoleEnable

`func (o *GridSecuritySetting) HasSecurityAccessRemoteConsoleEnable() bool`

HasSecurityAccessRemoteConsoleEnable returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *GridSecuritySetting) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *GridSecuritySetting) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *GridSecuritySetting) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *GridSecuritySetting) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetSshPermEnable

`func (o *GridSecuritySetting) GetSshPermEnable() bool`

GetSshPermEnable returns the SshPermEnable field if non-nil, zero value otherwise.

### GetSshPermEnableOk

`func (o *GridSecuritySetting) GetSshPermEnableOk() (*bool, bool)`

GetSshPermEnableOk returns a tuple with the SshPermEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPermEnable

`func (o *GridSecuritySetting) SetSshPermEnable(v bool)`

SetSshPermEnable sets SshPermEnable field to given value.

### HasSshPermEnable

`func (o *GridSecuritySetting) HasSshPermEnable() bool`

HasSshPermEnable returns a boolean if a field has been set.

### GetSupportAccessEnable

`func (o *GridSecuritySetting) GetSupportAccessEnable() bool`

GetSupportAccessEnable returns the SupportAccessEnable field if non-nil, zero value otherwise.

### GetSupportAccessEnableOk

`func (o *GridSecuritySetting) GetSupportAccessEnableOk() (*bool, bool)`

GetSupportAccessEnableOk returns a tuple with the SupportAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccessEnable

`func (o *GridSecuritySetting) SetSupportAccessEnable(v bool)`

SetSupportAccessEnable sets SupportAccessEnable field to given value.

### HasSupportAccessEnable

`func (o *GridSecuritySetting) HasSupportAccessEnable() bool`

HasSupportAccessEnable returns a boolean if a field has been set.

### GetSupportAccessInfo

`func (o *GridSecuritySetting) GetSupportAccessInfo() string`

GetSupportAccessInfo returns the SupportAccessInfo field if non-nil, zero value otherwise.

### GetSupportAccessInfoOk

`func (o *GridSecuritySetting) GetSupportAccessInfoOk() (*string, bool)`

GetSupportAccessInfoOk returns a tuple with the SupportAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccessInfo

`func (o *GridSecuritySetting) SetSupportAccessInfo(v string)`

SetSupportAccessInfo sets SupportAccessInfo field to given value.

### HasSupportAccessInfo

`func (o *GridSecuritySetting) HasSupportAccessInfo() bool`

HasSupportAccessInfo returns a boolean if a field has been set.

### GetDisableConcurrentLogin

`func (o *GridSecuritySetting) GetDisableConcurrentLogin() bool`

GetDisableConcurrentLogin returns the DisableConcurrentLogin field if non-nil, zero value otherwise.

### GetDisableConcurrentLoginOk

`func (o *GridSecuritySetting) GetDisableConcurrentLoginOk() (*bool, bool)`

GetDisableConcurrentLoginOk returns a tuple with the DisableConcurrentLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConcurrentLogin

`func (o *GridSecuritySetting) SetDisableConcurrentLogin(v bool)`

SetDisableConcurrentLogin sets DisableConcurrentLogin field to given value.

### HasDisableConcurrentLogin

`func (o *GridSecuritySetting) HasDisableConcurrentLogin() bool`

HasDisableConcurrentLogin returns a boolean if a field has been set.

### GetInactivityLockoutSetting

`func (o *GridSecuritySetting) GetInactivityLockoutSetting() GridsecuritysettingInactivityLockoutSetting`

GetInactivityLockoutSetting returns the InactivityLockoutSetting field if non-nil, zero value otherwise.

### GetInactivityLockoutSettingOk

`func (o *GridSecuritySetting) GetInactivityLockoutSettingOk() (*GridsecuritysettingInactivityLockoutSetting, bool)`

GetInactivityLockoutSettingOk returns a tuple with the InactivityLockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityLockoutSetting

`func (o *GridSecuritySetting) SetInactivityLockoutSetting(v GridsecuritysettingInactivityLockoutSetting)`

SetInactivityLockoutSetting sets InactivityLockoutSetting field to given value.

### HasInactivityLockoutSetting

`func (o *GridSecuritySetting) HasInactivityLockoutSetting() bool`

HasInactivityLockoutSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


