# RecordHostCliCredentials

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

### NewRecordHostCliCredentials

`func NewRecordHostCliCredentials() *RecordHostCliCredentials`

NewRecordHostCliCredentials instantiates a new RecordHostCliCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHostCliCredentialsWithDefaults

`func NewRecordHostCliCredentialsWithDefaults() *RecordHostCliCredentials`

NewRecordHostCliCredentialsWithDefaults instantiates a new RecordHostCliCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *RecordHostCliCredentials) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RecordHostCliCredentials) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RecordHostCliCredentials) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RecordHostCliCredentials) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *RecordHostCliCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RecordHostCliCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RecordHostCliCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RecordHostCliCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCredentialType

`func (o *RecordHostCliCredentials) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *RecordHostCliCredentials) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *RecordHostCliCredentials) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *RecordHostCliCredentials) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetComment

`func (o *RecordHostCliCredentials) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordHostCliCredentials) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordHostCliCredentials) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordHostCliCredentials) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *RecordHostCliCredentials) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordHostCliCredentials) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordHostCliCredentials) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RecordHostCliCredentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *RecordHostCliCredentials) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *RecordHostCliCredentials) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *RecordHostCliCredentials) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *RecordHostCliCredentials) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


