# GetFtpuserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CreateHomeDir** | Pointer to **bool** | Determines whether to create the home directory with the user name or to use the existing directory as the home directory. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**HomeDir** | Pointer to **string** | The absolute path of the FTP user&#39;s home directory. | [optional] 
**Password** | Pointer to **string** | The FTP user password. | [optional] 
**Permission** | Pointer to **string** | The FTP user permission. | [optional] 
**Username** | Pointer to **string** | The FTP user name. | [optional] 
**Result** | Pointer to [**Ftpuser**](Ftpuser.md) |  | [optional] 

## Methods

### NewGetFtpuserResponse

`func NewGetFtpuserResponse() *GetFtpuserResponse`

NewGetFtpuserResponse instantiates a new GetFtpuserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFtpuserResponseWithDefaults

`func NewGetFtpuserResponseWithDefaults() *GetFtpuserResponse`

NewGetFtpuserResponseWithDefaults instantiates a new GetFtpuserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFtpuserResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFtpuserResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFtpuserResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFtpuserResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCreateHomeDir

`func (o *GetFtpuserResponse) GetCreateHomeDir() bool`

GetCreateHomeDir returns the CreateHomeDir field if non-nil, zero value otherwise.

### GetCreateHomeDirOk

`func (o *GetFtpuserResponse) GetCreateHomeDirOk() (*bool, bool)`

GetCreateHomeDirOk returns a tuple with the CreateHomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHomeDir

`func (o *GetFtpuserResponse) SetCreateHomeDir(v bool)`

SetCreateHomeDir sets CreateHomeDir field to given value.

### HasCreateHomeDir

`func (o *GetFtpuserResponse) HasCreateHomeDir() bool`

HasCreateHomeDir returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetFtpuserResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetFtpuserResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetFtpuserResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetFtpuserResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetHomeDir

`func (o *GetFtpuserResponse) GetHomeDir() string`

GetHomeDir returns the HomeDir field if non-nil, zero value otherwise.

### GetHomeDirOk

`func (o *GetFtpuserResponse) GetHomeDirOk() (*string, bool)`

GetHomeDirOk returns a tuple with the HomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDir

`func (o *GetFtpuserResponse) SetHomeDir(v string)`

SetHomeDir sets HomeDir field to given value.

### HasHomeDir

`func (o *GetFtpuserResponse) HasHomeDir() bool`

HasHomeDir returns a boolean if a field has been set.

### GetPassword

`func (o *GetFtpuserResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetFtpuserResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetFtpuserResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetFtpuserResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPermission

`func (o *GetFtpuserResponse) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *GetFtpuserResponse) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *GetFtpuserResponse) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *GetFtpuserResponse) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUsername

`func (o *GetFtpuserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetFtpuserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetFtpuserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetFtpuserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetResult

`func (o *GetFtpuserResponse) GetResult() Ftpuser`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFtpuserResponse) GetResultOk() (*Ftpuser, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFtpuserResponse) SetResult(v Ftpuser)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFtpuserResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


