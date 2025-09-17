# Adminuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AdminGroups** | Pointer to **[]string** | The names of the Admin Groups to which this Admin User belongs. Currently, this is limited to only one Admin Group. | [optional] 
**AuthMethod** | Pointer to **string** | Determines the way of authentication | [optional] 
**AuthType** | Pointer to **string** | The authentication type for the admin user. | [optional] 
**CaCertificateIssuer** | Pointer to **string** | The CA certificate that is used for user lookup during authentication. | [optional] 
**ClientCertificateSerialNumber** | Pointer to **string** | The serial number of the client certificate. | [optional] 
**Comment** | Pointer to **string** | Comment for the admin user; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the admin user is disabled or not. When this is set to False, the admin user is enabled. | [optional] 
**Email** | Pointer to **string** | The e-mail address for the admin user. | [optional] 
**EnableCertificateAuthentication** | Pointer to **bool** | Determines whether the user is allowed to log in only with the certificate. Regular username/password authentication will be disabled for this user. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the admin user. | [optional] 
**Password** | Pointer to **string** | The password for the administrator to use when logging in. | [optional] 
**SshKeys** | Pointer to [**[]AdminuserSshKeys**](AdminuserSshKeys.md) | List of ssh keys for a particular user. | [optional] 
**Status** | Pointer to **string** | Status of the user account. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | The time zone for this admin user. | [optional] 
**UseSshKeys** | Pointer to **bool** | \\, Enable/disable the ssh keypair authentication. | [optional] 
**UseTimeZone** | Pointer to **bool** | Use flag for: time_zone | [optional] 

## Methods

### NewAdminuser

`func NewAdminuser() *Adminuser`

NewAdminuser instantiates a new Adminuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminuserWithDefaults

`func NewAdminuserWithDefaults() *Adminuser`

NewAdminuserWithDefaults instantiates a new Adminuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Adminuser) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Adminuser) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Adminuser) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Adminuser) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdminGroups

`func (o *Adminuser) GetAdminGroups() []string`

GetAdminGroups returns the AdminGroups field if non-nil, zero value otherwise.

### GetAdminGroupsOk

`func (o *Adminuser) GetAdminGroupsOk() (*[]string, bool)`

GetAdminGroupsOk returns a tuple with the AdminGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroups

`func (o *Adminuser) SetAdminGroups(v []string)`

SetAdminGroups sets AdminGroups field to given value.

### HasAdminGroups

`func (o *Adminuser) HasAdminGroups() bool`

HasAdminGroups returns a boolean if a field has been set.

### GetAuthMethod

`func (o *Adminuser) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *Adminuser) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *Adminuser) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *Adminuser) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetAuthType

`func (o *Adminuser) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *Adminuser) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *Adminuser) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *Adminuser) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCaCertificateIssuer

`func (o *Adminuser) GetCaCertificateIssuer() string`

GetCaCertificateIssuer returns the CaCertificateIssuer field if non-nil, zero value otherwise.

### GetCaCertificateIssuerOk

`func (o *Adminuser) GetCaCertificateIssuerOk() (*string, bool)`

GetCaCertificateIssuerOk returns a tuple with the CaCertificateIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateIssuer

`func (o *Adminuser) SetCaCertificateIssuer(v string)`

SetCaCertificateIssuer sets CaCertificateIssuer field to given value.

### HasCaCertificateIssuer

`func (o *Adminuser) HasCaCertificateIssuer() bool`

HasCaCertificateIssuer returns a boolean if a field has been set.

### GetClientCertificateSerialNumber

`func (o *Adminuser) GetClientCertificateSerialNumber() string`

GetClientCertificateSerialNumber returns the ClientCertificateSerialNumber field if non-nil, zero value otherwise.

### GetClientCertificateSerialNumberOk

`func (o *Adminuser) GetClientCertificateSerialNumberOk() (*string, bool)`

GetClientCertificateSerialNumberOk returns a tuple with the ClientCertificateSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSerialNumber

`func (o *Adminuser) SetClientCertificateSerialNumber(v string)`

SetClientCertificateSerialNumber sets ClientCertificateSerialNumber field to given value.

### HasClientCertificateSerialNumber

`func (o *Adminuser) HasClientCertificateSerialNumber() bool`

HasClientCertificateSerialNumber returns a boolean if a field has been set.

### GetComment

`func (o *Adminuser) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Adminuser) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Adminuser) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Adminuser) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *Adminuser) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Adminuser) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Adminuser) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Adminuser) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEmail

`func (o *Adminuser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Adminuser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Adminuser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Adminuser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableCertificateAuthentication

`func (o *Adminuser) GetEnableCertificateAuthentication() bool`

GetEnableCertificateAuthentication returns the EnableCertificateAuthentication field if non-nil, zero value otherwise.

### GetEnableCertificateAuthenticationOk

`func (o *Adminuser) GetEnableCertificateAuthenticationOk() (*bool, bool)`

GetEnableCertificateAuthenticationOk returns a tuple with the EnableCertificateAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCertificateAuthentication

`func (o *Adminuser) SetEnableCertificateAuthentication(v bool)`

SetEnableCertificateAuthentication sets EnableCertificateAuthentication field to given value.

### HasEnableCertificateAuthentication

`func (o *Adminuser) HasEnableCertificateAuthentication() bool`

HasEnableCertificateAuthentication returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Adminuser) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Adminuser) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Adminuser) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Adminuser) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Adminuser) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Adminuser) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Adminuser) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Adminuser) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Adminuser) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Adminuser) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Adminuser) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Adminuser) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Adminuser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Adminuser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Adminuser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Adminuser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *Adminuser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Adminuser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Adminuser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Adminuser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSshKeys

`func (o *Adminuser) GetSshKeys() []AdminuserSshKeys`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Adminuser) GetSshKeysOk() (*[]AdminuserSshKeys, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Adminuser) SetSshKeys(v []AdminuserSshKeys)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Adminuser) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetStatus

`func (o *Adminuser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Adminuser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Adminuser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Adminuser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *Adminuser) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Adminuser) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Adminuser) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Adminuser) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUseSshKeys

`func (o *Adminuser) GetUseSshKeys() bool`

GetUseSshKeys returns the UseSshKeys field if non-nil, zero value otherwise.

### GetUseSshKeysOk

`func (o *Adminuser) GetUseSshKeysOk() (*bool, bool)`

GetUseSshKeysOk returns a tuple with the UseSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSshKeys

`func (o *Adminuser) SetUseSshKeys(v bool)`

SetUseSshKeys sets UseSshKeys field to given value.

### HasUseSshKeys

`func (o *Adminuser) HasUseSshKeys() bool`

HasUseSshKeys returns a boolean if a field has been set.

### GetUseTimeZone

`func (o *Adminuser) GetUseTimeZone() bool`

GetUseTimeZone returns the UseTimeZone field if non-nil, zero value otherwise.

### GetUseTimeZoneOk

`func (o *Adminuser) GetUseTimeZoneOk() (*bool, bool)`

GetUseTimeZoneOk returns a tuple with the UseTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeZone

`func (o *Adminuser) SetUseTimeZone(v bool)`

SetUseTimeZone sets UseTimeZone field to given value.

### HasUseTimeZone

`func (o *Adminuser) HasUseTimeZone() bool`

HasUseTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


