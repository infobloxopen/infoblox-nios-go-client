# GridLomUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The LOM user name. | [optional] 
**Password** | Pointer to **string** | The LOM user password. | [optional] 
**Role** | Pointer to **string** | The LOM user role which specifies the list of actions that are allowed for the user. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the LOM user is disabled. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the LOM user. | [optional] 

## Methods

### NewGridLomUsers

`func NewGridLomUsers() *GridLomUsers`

NewGridLomUsers instantiates a new GridLomUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridLomUsersWithDefaults

`func NewGridLomUsersWithDefaults() *GridLomUsers`

NewGridLomUsersWithDefaults instantiates a new GridLomUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GridLomUsers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridLomUsers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridLomUsers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridLomUsers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *GridLomUsers) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GridLomUsers) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GridLomUsers) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GridLomUsers) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *GridLomUsers) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GridLomUsers) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GridLomUsers) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GridLomUsers) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetDisable

`func (o *GridLomUsers) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GridLomUsers) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GridLomUsers) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GridLomUsers) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetComment

`func (o *GridLomUsers) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GridLomUsers) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GridLomUsers) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GridLomUsers) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


