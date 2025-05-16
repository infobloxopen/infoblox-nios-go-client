# DiscoveryMemberpropertiesSnmpv3Credentials

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

### NewDiscoveryMemberpropertiesSnmpv3Credentials

`func NewDiscoveryMemberpropertiesSnmpv3Credentials() *DiscoveryMemberpropertiesSnmpv3Credentials`

NewDiscoveryMemberpropertiesSnmpv3Credentials instantiates a new DiscoveryMemberpropertiesSnmpv3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryMemberpropertiesSnmpv3CredentialsWithDefaults

`func NewDiscoveryMemberpropertiesSnmpv3CredentialsWithDefaults() *DiscoveryMemberpropertiesSnmpv3Credentials`

NewDiscoveryMemberpropertiesSnmpv3CredentialsWithDefaults instantiates a new DiscoveryMemberpropertiesSnmpv3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetComment

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *DiscoveryMemberpropertiesSnmpv3Credentials) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


