# AdmingroupSecuritySetCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetAdp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetApacheHttpsCert** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCcMode** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCertificateAuthAdmins** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCertificateAuthServices** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCheckAuthNs** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCheckSslCertificate** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetDisableHttpsCertRegeneration** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetFipsMode** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetReportingCert** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSecurity** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSessionTimeout** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSubscriberSecureData** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSupportAccess** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSupportInstall** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetAdpDebug** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSupportTimeout** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetUpdateRabbitmqPassword** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupSecuritySetCommands

`func NewAdmingroupSecuritySetCommands() *AdmingroupSecuritySetCommands`

NewAdmingroupSecuritySetCommands instantiates a new AdmingroupSecuritySetCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupSecuritySetCommandsWithDefaults

`func NewAdmingroupSecuritySetCommandsWithDefaults() *AdmingroupSecuritySetCommands`

NewAdmingroupSecuritySetCommandsWithDefaults instantiates a new AdmingroupSecuritySetCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetAdp

`func (o *AdmingroupSecuritySetCommands) GetSetAdp() bool`

GetSetAdp returns the SetAdp field if non-nil, zero value otherwise.

### GetSetAdpOk

`func (o *AdmingroupSecuritySetCommands) GetSetAdpOk() (*bool, bool)`

GetSetAdpOk returns a tuple with the SetAdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAdp

`func (o *AdmingroupSecuritySetCommands) SetSetAdp(v bool)`

SetSetAdp sets SetAdp field to given value.

### HasSetAdp

`func (o *AdmingroupSecuritySetCommands) HasSetAdp() bool`

HasSetAdp returns a boolean if a field has been set.

### GetSetApacheHttpsCert

`func (o *AdmingroupSecuritySetCommands) GetSetApacheHttpsCert() bool`

GetSetApacheHttpsCert returns the SetApacheHttpsCert field if non-nil, zero value otherwise.

### GetSetApacheHttpsCertOk

`func (o *AdmingroupSecuritySetCommands) GetSetApacheHttpsCertOk() (*bool, bool)`

GetSetApacheHttpsCertOk returns a tuple with the SetApacheHttpsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetApacheHttpsCert

`func (o *AdmingroupSecuritySetCommands) SetSetApacheHttpsCert(v bool)`

SetSetApacheHttpsCert sets SetApacheHttpsCert field to given value.

### HasSetApacheHttpsCert

`func (o *AdmingroupSecuritySetCommands) HasSetApacheHttpsCert() bool`

HasSetApacheHttpsCert returns a boolean if a field has been set.

### GetSetCcMode

`func (o *AdmingroupSecuritySetCommands) GetSetCcMode() bool`

GetSetCcMode returns the SetCcMode field if non-nil, zero value otherwise.

### GetSetCcModeOk

`func (o *AdmingroupSecuritySetCommands) GetSetCcModeOk() (*bool, bool)`

GetSetCcModeOk returns a tuple with the SetCcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCcMode

`func (o *AdmingroupSecuritySetCommands) SetSetCcMode(v bool)`

SetSetCcMode sets SetCcMode field to given value.

### HasSetCcMode

`func (o *AdmingroupSecuritySetCommands) HasSetCcMode() bool`

HasSetCcMode returns a boolean if a field has been set.

### GetSetCertificateAuthAdmins

`func (o *AdmingroupSecuritySetCommands) GetSetCertificateAuthAdmins() bool`

GetSetCertificateAuthAdmins returns the SetCertificateAuthAdmins field if non-nil, zero value otherwise.

### GetSetCertificateAuthAdminsOk

`func (o *AdmingroupSecuritySetCommands) GetSetCertificateAuthAdminsOk() (*bool, bool)`

GetSetCertificateAuthAdminsOk returns a tuple with the SetCertificateAuthAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCertificateAuthAdmins

`func (o *AdmingroupSecuritySetCommands) SetSetCertificateAuthAdmins(v bool)`

SetSetCertificateAuthAdmins sets SetCertificateAuthAdmins field to given value.

### HasSetCertificateAuthAdmins

`func (o *AdmingroupSecuritySetCommands) HasSetCertificateAuthAdmins() bool`

HasSetCertificateAuthAdmins returns a boolean if a field has been set.

### GetSetCertificateAuthServices

`func (o *AdmingroupSecuritySetCommands) GetSetCertificateAuthServices() bool`

GetSetCertificateAuthServices returns the SetCertificateAuthServices field if non-nil, zero value otherwise.

### GetSetCertificateAuthServicesOk

`func (o *AdmingroupSecuritySetCommands) GetSetCertificateAuthServicesOk() (*bool, bool)`

GetSetCertificateAuthServicesOk returns a tuple with the SetCertificateAuthServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCertificateAuthServices

`func (o *AdmingroupSecuritySetCommands) SetSetCertificateAuthServices(v bool)`

SetSetCertificateAuthServices sets SetCertificateAuthServices field to given value.

### HasSetCertificateAuthServices

`func (o *AdmingroupSecuritySetCommands) HasSetCertificateAuthServices() bool`

HasSetCertificateAuthServices returns a boolean if a field has been set.

### GetSetCheckAuthNs

`func (o *AdmingroupSecuritySetCommands) GetSetCheckAuthNs() bool`

GetSetCheckAuthNs returns the SetCheckAuthNs field if non-nil, zero value otherwise.

### GetSetCheckAuthNsOk

`func (o *AdmingroupSecuritySetCommands) GetSetCheckAuthNsOk() (*bool, bool)`

GetSetCheckAuthNsOk returns a tuple with the SetCheckAuthNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCheckAuthNs

`func (o *AdmingroupSecuritySetCommands) SetSetCheckAuthNs(v bool)`

SetSetCheckAuthNs sets SetCheckAuthNs field to given value.

### HasSetCheckAuthNs

`func (o *AdmingroupSecuritySetCommands) HasSetCheckAuthNs() bool`

HasSetCheckAuthNs returns a boolean if a field has been set.

### GetSetCheckSslCertificate

`func (o *AdmingroupSecuritySetCommands) GetSetCheckSslCertificate() bool`

GetSetCheckSslCertificate returns the SetCheckSslCertificate field if non-nil, zero value otherwise.

### GetSetCheckSslCertificateOk

`func (o *AdmingroupSecuritySetCommands) GetSetCheckSslCertificateOk() (*bool, bool)`

GetSetCheckSslCertificateOk returns a tuple with the SetCheckSslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCheckSslCertificate

`func (o *AdmingroupSecuritySetCommands) SetSetCheckSslCertificate(v bool)`

SetSetCheckSslCertificate sets SetCheckSslCertificate field to given value.

### HasSetCheckSslCertificate

`func (o *AdmingroupSecuritySetCommands) HasSetCheckSslCertificate() bool`

HasSetCheckSslCertificate returns a boolean if a field has been set.

### GetSetDisableHttpsCertRegeneration

`func (o *AdmingroupSecuritySetCommands) GetSetDisableHttpsCertRegeneration() bool`

GetSetDisableHttpsCertRegeneration returns the SetDisableHttpsCertRegeneration field if non-nil, zero value otherwise.

### GetSetDisableHttpsCertRegenerationOk

`func (o *AdmingroupSecuritySetCommands) GetSetDisableHttpsCertRegenerationOk() (*bool, bool)`

GetSetDisableHttpsCertRegenerationOk returns a tuple with the SetDisableHttpsCertRegeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDisableHttpsCertRegeneration

`func (o *AdmingroupSecuritySetCommands) SetSetDisableHttpsCertRegeneration(v bool)`

SetSetDisableHttpsCertRegeneration sets SetDisableHttpsCertRegeneration field to given value.

### HasSetDisableHttpsCertRegeneration

`func (o *AdmingroupSecuritySetCommands) HasSetDisableHttpsCertRegeneration() bool`

HasSetDisableHttpsCertRegeneration returns a boolean if a field has been set.

### GetSetFipsMode

`func (o *AdmingroupSecuritySetCommands) GetSetFipsMode() bool`

GetSetFipsMode returns the SetFipsMode field if non-nil, zero value otherwise.

### GetSetFipsModeOk

`func (o *AdmingroupSecuritySetCommands) GetSetFipsModeOk() (*bool, bool)`

GetSetFipsModeOk returns a tuple with the SetFipsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetFipsMode

`func (o *AdmingroupSecuritySetCommands) SetSetFipsMode(v bool)`

SetSetFipsMode sets SetFipsMode field to given value.

### HasSetFipsMode

`func (o *AdmingroupSecuritySetCommands) HasSetFipsMode() bool`

HasSetFipsMode returns a boolean if a field has been set.

### GetSetReportingCert

`func (o *AdmingroupSecuritySetCommands) GetSetReportingCert() bool`

GetSetReportingCert returns the SetReportingCert field if non-nil, zero value otherwise.

### GetSetReportingCertOk

`func (o *AdmingroupSecuritySetCommands) GetSetReportingCertOk() (*bool, bool)`

GetSetReportingCertOk returns a tuple with the SetReportingCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetReportingCert

`func (o *AdmingroupSecuritySetCommands) SetSetReportingCert(v bool)`

SetSetReportingCert sets SetReportingCert field to given value.

### HasSetReportingCert

`func (o *AdmingroupSecuritySetCommands) HasSetReportingCert() bool`

HasSetReportingCert returns a boolean if a field has been set.

### GetSetSecurity

`func (o *AdmingroupSecuritySetCommands) GetSetSecurity() bool`

GetSetSecurity returns the SetSecurity field if non-nil, zero value otherwise.

### GetSetSecurityOk

`func (o *AdmingroupSecuritySetCommands) GetSetSecurityOk() (*bool, bool)`

GetSetSecurityOk returns a tuple with the SetSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSecurity

`func (o *AdmingroupSecuritySetCommands) SetSetSecurity(v bool)`

SetSetSecurity sets SetSecurity field to given value.

### HasSetSecurity

`func (o *AdmingroupSecuritySetCommands) HasSetSecurity() bool`

HasSetSecurity returns a boolean if a field has been set.

### GetSetSessionTimeout

`func (o *AdmingroupSecuritySetCommands) GetSetSessionTimeout() bool`

GetSetSessionTimeout returns the SetSessionTimeout field if non-nil, zero value otherwise.

### GetSetSessionTimeoutOk

`func (o *AdmingroupSecuritySetCommands) GetSetSessionTimeoutOk() (*bool, bool)`

GetSetSessionTimeoutOk returns a tuple with the SetSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSessionTimeout

`func (o *AdmingroupSecuritySetCommands) SetSetSessionTimeout(v bool)`

SetSetSessionTimeout sets SetSessionTimeout field to given value.

### HasSetSessionTimeout

`func (o *AdmingroupSecuritySetCommands) HasSetSessionTimeout() bool`

HasSetSessionTimeout returns a boolean if a field has been set.

### GetSetSubscriberSecureData

`func (o *AdmingroupSecuritySetCommands) GetSetSubscriberSecureData() bool`

GetSetSubscriberSecureData returns the SetSubscriberSecureData field if non-nil, zero value otherwise.

### GetSetSubscriberSecureDataOk

`func (o *AdmingroupSecuritySetCommands) GetSetSubscriberSecureDataOk() (*bool, bool)`

GetSetSubscriberSecureDataOk returns a tuple with the SetSubscriberSecureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSubscriberSecureData

`func (o *AdmingroupSecuritySetCommands) SetSetSubscriberSecureData(v bool)`

SetSetSubscriberSecureData sets SetSubscriberSecureData field to given value.

### HasSetSubscriberSecureData

`func (o *AdmingroupSecuritySetCommands) HasSetSubscriberSecureData() bool`

HasSetSubscriberSecureData returns a boolean if a field has been set.

### GetSetSupportAccess

`func (o *AdmingroupSecuritySetCommands) GetSetSupportAccess() bool`

GetSetSupportAccess returns the SetSupportAccess field if non-nil, zero value otherwise.

### GetSetSupportAccessOk

`func (o *AdmingroupSecuritySetCommands) GetSetSupportAccessOk() (*bool, bool)`

GetSetSupportAccessOk returns a tuple with the SetSupportAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSupportAccess

`func (o *AdmingroupSecuritySetCommands) SetSetSupportAccess(v bool)`

SetSetSupportAccess sets SetSupportAccess field to given value.

### HasSetSupportAccess

`func (o *AdmingroupSecuritySetCommands) HasSetSupportAccess() bool`

HasSetSupportAccess returns a boolean if a field has been set.

### GetSetSupportInstall

`func (o *AdmingroupSecuritySetCommands) GetSetSupportInstall() bool`

GetSetSupportInstall returns the SetSupportInstall field if non-nil, zero value otherwise.

### GetSetSupportInstallOk

`func (o *AdmingroupSecuritySetCommands) GetSetSupportInstallOk() (*bool, bool)`

GetSetSupportInstallOk returns a tuple with the SetSupportInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSupportInstall

`func (o *AdmingroupSecuritySetCommands) SetSetSupportInstall(v bool)`

SetSetSupportInstall sets SetSupportInstall field to given value.

### HasSetSupportInstall

`func (o *AdmingroupSecuritySetCommands) HasSetSupportInstall() bool`

HasSetSupportInstall returns a boolean if a field has been set.

### GetSetAdpDebug

`func (o *AdmingroupSecuritySetCommands) GetSetAdpDebug() bool`

GetSetAdpDebug returns the SetAdpDebug field if non-nil, zero value otherwise.

### GetSetAdpDebugOk

`func (o *AdmingroupSecuritySetCommands) GetSetAdpDebugOk() (*bool, bool)`

GetSetAdpDebugOk returns a tuple with the SetAdpDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAdpDebug

`func (o *AdmingroupSecuritySetCommands) SetSetAdpDebug(v bool)`

SetSetAdpDebug sets SetAdpDebug field to given value.

### HasSetAdpDebug

`func (o *AdmingroupSecuritySetCommands) HasSetAdpDebug() bool`

HasSetAdpDebug returns a boolean if a field has been set.

### GetSetSupportTimeout

`func (o *AdmingroupSecuritySetCommands) GetSetSupportTimeout() bool`

GetSetSupportTimeout returns the SetSupportTimeout field if non-nil, zero value otherwise.

### GetSetSupportTimeoutOk

`func (o *AdmingroupSecuritySetCommands) GetSetSupportTimeoutOk() (*bool, bool)`

GetSetSupportTimeoutOk returns a tuple with the SetSupportTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSupportTimeout

`func (o *AdmingroupSecuritySetCommands) SetSetSupportTimeout(v bool)`

SetSetSupportTimeout sets SetSupportTimeout field to given value.

### HasSetSupportTimeout

`func (o *AdmingroupSecuritySetCommands) HasSetSupportTimeout() bool`

HasSetSupportTimeout returns a boolean if a field has been set.

### GetSetUpdateRabbitmqPassword

`func (o *AdmingroupSecuritySetCommands) GetSetUpdateRabbitmqPassword() bool`

GetSetUpdateRabbitmqPassword returns the SetUpdateRabbitmqPassword field if non-nil, zero value otherwise.

### GetSetUpdateRabbitmqPasswordOk

`func (o *AdmingroupSecuritySetCommands) GetSetUpdateRabbitmqPasswordOk() (*bool, bool)`

GetSetUpdateRabbitmqPasswordOk returns a tuple with the SetUpdateRabbitmqPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUpdateRabbitmqPassword

`func (o *AdmingroupSecuritySetCommands) SetSetUpdateRabbitmqPassword(v bool)`

SetSetUpdateRabbitmqPassword sets SetUpdateRabbitmqPassword field to given value.

### HasSetUpdateRabbitmqPassword

`func (o *AdmingroupSecuritySetCommands) HasSetUpdateRabbitmqPassword() bool`

HasSetUpdateRabbitmqPassword returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupSecuritySetCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupSecuritySetCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupSecuritySetCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupSecuritySetCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupSecuritySetCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupSecuritySetCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupSecuritySetCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupSecuritySetCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


