# FixedaddressCliCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | The CLI user name. | [optional] 
**Password** | Pointer to **string** | The CLI password. | [optional] 
**CredentialType** | Pointer to **string** | The type of the credential. | [optional] 
**Comment** | Pointer to **string** | The commment for the credential. | [optional] 
**Id** | Pointer to **int64** | The Credentials ID. | [optional] [readonly] 
**CredentialGroup** | Pointer to **string** | Group for the CLI credential. | [optional] 

## Methods

### NewFixedaddressCliCredentials

`func NewFixedaddressCliCredentials() *FixedaddressCliCredentials`

NewFixedaddressCliCredentials instantiates a new FixedaddressCliCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedaddressCliCredentialsWithDefaults

`func NewFixedaddressCliCredentialsWithDefaults() *FixedaddressCliCredentials`

NewFixedaddressCliCredentialsWithDefaults instantiates a new FixedaddressCliCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *FixedaddressCliCredentials) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FixedaddressCliCredentials) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FixedaddressCliCredentials) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FixedaddressCliCredentials) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *FixedaddressCliCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FixedaddressCliCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FixedaddressCliCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FixedaddressCliCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCredentialType

`func (o *FixedaddressCliCredentials) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *FixedaddressCliCredentials) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *FixedaddressCliCredentials) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *FixedaddressCliCredentials) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetComment

`func (o *FixedaddressCliCredentials) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FixedaddressCliCredentials) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FixedaddressCliCredentials) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FixedaddressCliCredentials) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *FixedaddressCliCredentials) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedaddressCliCredentials) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedaddressCliCredentials) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FixedaddressCliCredentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *FixedaddressCliCredentials) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *FixedaddressCliCredentials) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *FixedaddressCliCredentials) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *FixedaddressCliCredentials) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


