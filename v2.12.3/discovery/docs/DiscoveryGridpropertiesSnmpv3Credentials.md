# DiscoveryGridpropertiesSnmpv3Credentials

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

### NewDiscoveryGridpropertiesSnmpv3Credentials

`func NewDiscoveryGridpropertiesSnmpv3Credentials() *DiscoveryGridpropertiesSnmpv3Credentials`

NewDiscoveryGridpropertiesSnmpv3Credentials instantiates a new DiscoveryGridpropertiesSnmpv3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesSnmpv3CredentialsWithDefaults

`func NewDiscoveryGridpropertiesSnmpv3CredentialsWithDefaults() *DiscoveryGridpropertiesSnmpv3Credentials`

NewDiscoveryGridpropertiesSnmpv3CredentialsWithDefaults instantiates a new DiscoveryGridpropertiesSnmpv3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetComment

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


