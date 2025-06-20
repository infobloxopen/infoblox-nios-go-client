# Ftpuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CreateHomeDir** | Pointer to **bool** | Determines whether to create the home directory with the user name or to use the existing directory as the home directory. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**HomeDir** | Pointer to **string** | The absolute path of the FTP user&#39;s home directory. | [optional] 
**Password** | Pointer to **string** | The FTP user password. | [optional] 
**Permission** | Pointer to **string** | The FTP user permission. | [optional] 
**Username** | Pointer to **string** | The FTP user name. | [optional] 

## Methods

### NewFtpuser

`func NewFtpuser() *Ftpuser`

NewFtpuser instantiates a new Ftpuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFtpuserWithDefaults

`func NewFtpuserWithDefaults() *Ftpuser`

NewFtpuserWithDefaults instantiates a new Ftpuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ftpuser) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ftpuser) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ftpuser) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ftpuser) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCreateHomeDir

`func (o *Ftpuser) GetCreateHomeDir() bool`

GetCreateHomeDir returns the CreateHomeDir field if non-nil, zero value otherwise.

### GetCreateHomeDirOk

`func (o *Ftpuser) GetCreateHomeDirOk() (*bool, bool)`

GetCreateHomeDirOk returns a tuple with the CreateHomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHomeDir

`func (o *Ftpuser) SetCreateHomeDir(v bool)`

SetCreateHomeDir sets CreateHomeDir field to given value.

### HasCreateHomeDir

`func (o *Ftpuser) HasCreateHomeDir() bool`

HasCreateHomeDir returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Ftpuser) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Ftpuser) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Ftpuser) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Ftpuser) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHomeDir

`func (o *Ftpuser) GetHomeDir() string`

GetHomeDir returns the HomeDir field if non-nil, zero value otherwise.

### GetHomeDirOk

`func (o *Ftpuser) GetHomeDirOk() (*string, bool)`

GetHomeDirOk returns a tuple with the HomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDir

`func (o *Ftpuser) SetHomeDir(v string)`

SetHomeDir sets HomeDir field to given value.

### HasHomeDir

`func (o *Ftpuser) HasHomeDir() bool`

HasHomeDir returns a boolean if a field has been set.

### GetPassword

`func (o *Ftpuser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Ftpuser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Ftpuser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Ftpuser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPermission

`func (o *Ftpuser) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Ftpuser) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Ftpuser) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Ftpuser) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUsername

`func (o *Ftpuser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Ftpuser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Ftpuser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Ftpuser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


