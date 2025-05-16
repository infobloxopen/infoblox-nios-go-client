# FixedaddressSnmp3Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | The SNMPv3 user name. | [optional] 
**AuthenticationProtocol** | Pointer to **string** | Authentication protocol for the SNMPv3 user. | [optional] 
**AuthenticationPassword** | Pointer to **string** | Authentication password for the SNMPv3 user. | [optional] 
**PrivacyProtocol** | Pointer to **string** | Privacy protocol for the SNMPv3 user. | [optional] 
**PrivacyPassword** | Pointer to **string** | Privacy password for the SNMPv3 user. | [optional] 
**Comment** | Pointer to **string** | Comments for the SNMPv3 user. | [optional] 
**CredentialGroup** | Pointer to **string** | Group for the SNMPv3 credential. | [optional] 

## Methods

### NewFixedaddressSnmp3Credential

`func NewFixedaddressSnmp3Credential() *FixedaddressSnmp3Credential`

NewFixedaddressSnmp3Credential instantiates a new FixedaddressSnmp3Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedaddressSnmp3CredentialWithDefaults

`func NewFixedaddressSnmp3CredentialWithDefaults() *FixedaddressSnmp3Credential`

NewFixedaddressSnmp3CredentialWithDefaults instantiates a new FixedaddressSnmp3Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *FixedaddressSnmp3Credential) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FixedaddressSnmp3Credential) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FixedaddressSnmp3Credential) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FixedaddressSnmp3Credential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *FixedaddressSnmp3Credential) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *FixedaddressSnmp3Credential) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *FixedaddressSnmp3Credential) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *FixedaddressSnmp3Credential) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *FixedaddressSnmp3Credential) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *FixedaddressSnmp3Credential) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *FixedaddressSnmp3Credential) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *FixedaddressSnmp3Credential) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *FixedaddressSnmp3Credential) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *FixedaddressSnmp3Credential) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *FixedaddressSnmp3Credential) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *FixedaddressSnmp3Credential) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *FixedaddressSnmp3Credential) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *FixedaddressSnmp3Credential) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *FixedaddressSnmp3Credential) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *FixedaddressSnmp3Credential) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetComment

`func (o *FixedaddressSnmp3Credential) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FixedaddressSnmp3Credential) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FixedaddressSnmp3Credential) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FixedaddressSnmp3Credential) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *FixedaddressSnmp3Credential) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *FixedaddressSnmp3Credential) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *FixedaddressSnmp3Credential) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *FixedaddressSnmp3Credential) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


