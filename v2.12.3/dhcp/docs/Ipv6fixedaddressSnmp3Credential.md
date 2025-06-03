# Ipv6fixedaddressSnmp3Credential

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

### NewIpv6fixedaddressSnmp3Credential

`func NewIpv6fixedaddressSnmp3Credential() *Ipv6fixedaddressSnmp3Credential`

NewIpv6fixedaddressSnmp3Credential instantiates a new Ipv6fixedaddressSnmp3Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6fixedaddressSnmp3CredentialWithDefaults

`func NewIpv6fixedaddressSnmp3CredentialWithDefaults() *Ipv6fixedaddressSnmp3Credential`

NewIpv6fixedaddressSnmp3CredentialWithDefaults instantiates a new Ipv6fixedaddressSnmp3Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *Ipv6fixedaddressSnmp3Credential) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Ipv6fixedaddressSnmp3Credential) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Ipv6fixedaddressSnmp3Credential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *Ipv6fixedaddressSnmp3Credential) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *Ipv6fixedaddressSnmp3Credential) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *Ipv6fixedaddressSnmp3Credential) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *Ipv6fixedaddressSnmp3Credential) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *Ipv6fixedaddressSnmp3Credential) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *Ipv6fixedaddressSnmp3Credential) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *Ipv6fixedaddressSnmp3Credential) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *Ipv6fixedaddressSnmp3Credential) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *Ipv6fixedaddressSnmp3Credential) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *Ipv6fixedaddressSnmp3Credential) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *Ipv6fixedaddressSnmp3Credential) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *Ipv6fixedaddressSnmp3Credential) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6fixedaddressSnmp3Credential) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6fixedaddressSnmp3Credential) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6fixedaddressSnmp3Credential) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *Ipv6fixedaddressSnmp3Credential) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *Ipv6fixedaddressSnmp3Credential) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *Ipv6fixedaddressSnmp3Credential) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *Ipv6fixedaddressSnmp3Credential) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


