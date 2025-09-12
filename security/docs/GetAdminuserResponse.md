# GetAdminuserResponse

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
**Result** | Pointer to [**Adminuser**](Adminuser.md) |  | [optional] 

## Methods

### NewGetAdminuserResponse

`func NewGetAdminuserResponse() *GetAdminuserResponse`

NewGetAdminuserResponse instantiates a new GetAdminuserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdminuserResponseWithDefaults

`func NewGetAdminuserResponseWithDefaults() *GetAdminuserResponse`

NewGetAdminuserResponseWithDefaults instantiates a new GetAdminuserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAdminuserResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAdminuserResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAdminuserResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAdminuserResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdminGroups

`func (o *GetAdminuserResponse) GetAdminGroups() []string`

GetAdminGroups returns the AdminGroups field if non-nil, zero value otherwise.

### GetAdminGroupsOk

`func (o *GetAdminuserResponse) GetAdminGroupsOk() (*[]string, bool)`

GetAdminGroupsOk returns a tuple with the AdminGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroups

`func (o *GetAdminuserResponse) SetAdminGroups(v []string)`

SetAdminGroups sets AdminGroups field to given value.

### HasAdminGroups

`func (o *GetAdminuserResponse) HasAdminGroups() bool`

HasAdminGroups returns a boolean if a field has been set.

### GetAuthMethod

`func (o *GetAdminuserResponse) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *GetAdminuserResponse) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *GetAdminuserResponse) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *GetAdminuserResponse) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetAuthType

`func (o *GetAdminuserResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GetAdminuserResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GetAdminuserResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *GetAdminuserResponse) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCaCertificateIssuer

`func (o *GetAdminuserResponse) GetCaCertificateIssuer() string`

GetCaCertificateIssuer returns the CaCertificateIssuer field if non-nil, zero value otherwise.

### GetCaCertificateIssuerOk

`func (o *GetAdminuserResponse) GetCaCertificateIssuerOk() (*string, bool)`

GetCaCertificateIssuerOk returns a tuple with the CaCertificateIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateIssuer

`func (o *GetAdminuserResponse) SetCaCertificateIssuer(v string)`

SetCaCertificateIssuer sets CaCertificateIssuer field to given value.

### HasCaCertificateIssuer

`func (o *GetAdminuserResponse) HasCaCertificateIssuer() bool`

HasCaCertificateIssuer returns a boolean if a field has been set.

### GetClientCertificateSerialNumber

`func (o *GetAdminuserResponse) GetClientCertificateSerialNumber() string`

GetClientCertificateSerialNumber returns the ClientCertificateSerialNumber field if non-nil, zero value otherwise.

### GetClientCertificateSerialNumberOk

`func (o *GetAdminuserResponse) GetClientCertificateSerialNumberOk() (*string, bool)`

GetClientCertificateSerialNumberOk returns a tuple with the ClientCertificateSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSerialNumber

`func (o *GetAdminuserResponse) SetClientCertificateSerialNumber(v string)`

SetClientCertificateSerialNumber sets ClientCertificateSerialNumber field to given value.

### HasClientCertificateSerialNumber

`func (o *GetAdminuserResponse) HasClientCertificateSerialNumber() bool`

HasClientCertificateSerialNumber returns a boolean if a field has been set.

### GetComment

`func (o *GetAdminuserResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAdminuserResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAdminuserResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAdminuserResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetAdminuserResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetAdminuserResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetAdminuserResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetAdminuserResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEmail

`func (o *GetAdminuserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAdminuserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAdminuserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetAdminuserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableCertificateAuthentication

`func (o *GetAdminuserResponse) GetEnableCertificateAuthentication() bool`

GetEnableCertificateAuthentication returns the EnableCertificateAuthentication field if non-nil, zero value otherwise.

### GetEnableCertificateAuthenticationOk

`func (o *GetAdminuserResponse) GetEnableCertificateAuthenticationOk() (*bool, bool)`

GetEnableCertificateAuthenticationOk returns a tuple with the EnableCertificateAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCertificateAuthentication

`func (o *GetAdminuserResponse) SetEnableCertificateAuthentication(v bool)`

SetEnableCertificateAuthentication sets EnableCertificateAuthentication field to given value.

### HasEnableCertificateAuthentication

`func (o *GetAdminuserResponse) HasEnableCertificateAuthentication() bool`

HasEnableCertificateAuthentication returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetAdminuserResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetAdminuserResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetAdminuserResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetAdminuserResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetAdminuserResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetAdminuserResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetAdminuserResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetAdminuserResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetAdminuserResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetAdminuserResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetAdminuserResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetAdminuserResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GetAdminuserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdminuserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdminuserResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdminuserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *GetAdminuserResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetAdminuserResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetAdminuserResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetAdminuserResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSshKeys

`func (o *GetAdminuserResponse) GetSshKeys() []AdminuserSshKeys`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *GetAdminuserResponse) GetSshKeysOk() (*[]AdminuserSshKeys, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *GetAdminuserResponse) SetSshKeys(v []AdminuserSshKeys)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *GetAdminuserResponse) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetStatus

`func (o *GetAdminuserResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAdminuserResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAdminuserResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAdminuserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetAdminuserResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetAdminuserResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetAdminuserResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetAdminuserResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUseSshKeys

`func (o *GetAdminuserResponse) GetUseSshKeys() bool`

GetUseSshKeys returns the UseSshKeys field if non-nil, zero value otherwise.

### GetUseSshKeysOk

`func (o *GetAdminuserResponse) GetUseSshKeysOk() (*bool, bool)`

GetUseSshKeysOk returns a tuple with the UseSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSshKeys

`func (o *GetAdminuserResponse) SetUseSshKeys(v bool)`

SetUseSshKeys sets UseSshKeys field to given value.

### HasUseSshKeys

`func (o *GetAdminuserResponse) HasUseSshKeys() bool`

HasUseSshKeys returns a boolean if a field has been set.

### GetUseTimeZone

`func (o *GetAdminuserResponse) GetUseTimeZone() bool`

GetUseTimeZone returns the UseTimeZone field if non-nil, zero value otherwise.

### GetUseTimeZoneOk

`func (o *GetAdminuserResponse) GetUseTimeZoneOk() (*bool, bool)`

GetUseTimeZoneOk returns a tuple with the UseTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeZone

`func (o *GetAdminuserResponse) SetUseTimeZone(v bool)`

SetUseTimeZone sets UseTimeZone field to given value.

### HasUseTimeZone

`func (o *GetAdminuserResponse) HasUseTimeZone() bool`

HasUseTimeZone returns a boolean if a field has been set.

### GetResult

`func (o *GetAdminuserResponse) GetResult() Adminuser`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAdminuserResponse) GetResultOk() (*Adminuser, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAdminuserResponse) SetResult(v Adminuser)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAdminuserResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


