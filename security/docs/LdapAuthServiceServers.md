# LdapAuthServiceServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IP address or FQDN of the LDAP server. | [optional] 
**AuthenticationType** | Pointer to **string** | The authentication type for the LDAP server. | [optional] 
**BaseDn** | Pointer to **string** | The base DN for the LDAP server. | [optional] 
**BindPassword** | Pointer to **string** | The user password for authentication. | [optional] 
**BindUserDn** | Pointer to **string** | The user DN for authentication. | [optional] 
**Comment** | Pointer to **string** | The LDAP descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines if the LDAP server is disabled. | [optional] 
**Encryption** | Pointer to **string** | The LDAP server encryption type. | [optional] 
**Port** | Pointer to **int64** | The LDAP server port. | [optional] 
**UseMgmtPort** | Pointer to **bool** | Determines if the connection via the MGMT interface is allowed. | [optional] 
**Version** | Pointer to **string** | The LDAP server version. | [optional] 

## Methods

### NewLdapAuthServiceServers

`func NewLdapAuthServiceServers() *LdapAuthServiceServers`

NewLdapAuthServiceServers instantiates a new LdapAuthServiceServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAuthServiceServersWithDefaults

`func NewLdapAuthServiceServersWithDefaults() *LdapAuthServiceServers`

NewLdapAuthServiceServersWithDefaults instantiates a new LdapAuthServiceServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LdapAuthServiceServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LdapAuthServiceServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LdapAuthServiceServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LdapAuthServiceServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *LdapAuthServiceServers) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *LdapAuthServiceServers) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *LdapAuthServiceServers) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *LdapAuthServiceServers) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetBaseDn

`func (o *LdapAuthServiceServers) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapAuthServiceServers) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapAuthServiceServers) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapAuthServiceServers) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindPassword

`func (o *LdapAuthServiceServers) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *LdapAuthServiceServers) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *LdapAuthServiceServers) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *LdapAuthServiceServers) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetBindUserDn

`func (o *LdapAuthServiceServers) GetBindUserDn() string`

GetBindUserDn returns the BindUserDn field if non-nil, zero value otherwise.

### GetBindUserDnOk

`func (o *LdapAuthServiceServers) GetBindUserDnOk() (*string, bool)`

GetBindUserDnOk returns a tuple with the BindUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindUserDn

`func (o *LdapAuthServiceServers) SetBindUserDn(v string)`

SetBindUserDn sets BindUserDn field to given value.

### HasBindUserDn

`func (o *LdapAuthServiceServers) HasBindUserDn() bool`

HasBindUserDn returns a boolean if a field has been set.

### GetComment

`func (o *LdapAuthServiceServers) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LdapAuthServiceServers) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LdapAuthServiceServers) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LdapAuthServiceServers) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *LdapAuthServiceServers) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *LdapAuthServiceServers) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *LdapAuthServiceServers) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *LdapAuthServiceServers) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEncryption

`func (o *LdapAuthServiceServers) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *LdapAuthServiceServers) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *LdapAuthServiceServers) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *LdapAuthServiceServers) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *LdapAuthServiceServers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LdapAuthServiceServers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LdapAuthServiceServers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *LdapAuthServiceServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUseMgmtPort

`func (o *LdapAuthServiceServers) GetUseMgmtPort() bool`

GetUseMgmtPort returns the UseMgmtPort field if non-nil, zero value otherwise.

### GetUseMgmtPortOk

`func (o *LdapAuthServiceServers) GetUseMgmtPortOk() (*bool, bool)`

GetUseMgmtPortOk returns a tuple with the UseMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtPort

`func (o *LdapAuthServiceServers) SetUseMgmtPort(v bool)`

SetUseMgmtPort sets UseMgmtPort field to given value.

### HasUseMgmtPort

`func (o *LdapAuthServiceServers) HasUseMgmtPort() bool`

HasUseMgmtPort returns a boolean if a field has been set.

### GetVersion

`func (o *LdapAuthServiceServers) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LdapAuthServiceServers) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LdapAuthServiceServers) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LdapAuthServiceServers) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


