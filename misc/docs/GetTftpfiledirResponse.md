# GetTftpfiledirResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Directory** | Pointer to **string** | The path to the directory that contains file or subdirectory. | [optional] 
**IsSyncedToGm** | Pointer to **bool** | Determines whether the TFTP entity is synchronized to Grid Master. | [optional] [readonly] 
**LastModify** | Pointer to **int64** | The time when the file or directory was last modified. | [optional] [readonly] 
**Name** | Pointer to **string** | The TFTP directory or file name. | [optional] 
**Type** | Pointer to **string** | The type of TFTP file system entity (directory or file). | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**VtftpDirMembers** | Pointer to [**[]TftpfiledirVtftpDirMembers**](TftpfiledirVtftpDirMembers.md) | The replication members with TFTP client addresses where this virtual folder is applicable. | [optional] 
**Result** | Pointer to [**Tftpfiledir**](Tftpfiledir.md) |  | [optional] 

## Methods

### NewGetTftpfiledirResponse

`func NewGetTftpfiledirResponse() *GetTftpfiledirResponse`

NewGetTftpfiledirResponse instantiates a new GetTftpfiledirResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTftpfiledirResponseWithDefaults

`func NewGetTftpfiledirResponseWithDefaults() *GetTftpfiledirResponse`

NewGetTftpfiledirResponseWithDefaults instantiates a new GetTftpfiledirResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetTftpfiledirResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetTftpfiledirResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetTftpfiledirResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetTftpfiledirResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDirectory

`func (o *GetTftpfiledirResponse) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *GetTftpfiledirResponse) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *GetTftpfiledirResponse) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *GetTftpfiledirResponse) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetIsSyncedToGm

`func (o *GetTftpfiledirResponse) GetIsSyncedToGm() bool`

GetIsSyncedToGm returns the IsSyncedToGm field if non-nil, zero value otherwise.

### GetIsSyncedToGmOk

`func (o *GetTftpfiledirResponse) GetIsSyncedToGmOk() (*bool, bool)`

GetIsSyncedToGmOk returns a tuple with the IsSyncedToGm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncedToGm

`func (o *GetTftpfiledirResponse) SetIsSyncedToGm(v bool)`

SetIsSyncedToGm sets IsSyncedToGm field to given value.

### HasIsSyncedToGm

`func (o *GetTftpfiledirResponse) HasIsSyncedToGm() bool`

HasIsSyncedToGm returns a boolean if a field has been set.

### GetLastModify

`func (o *GetTftpfiledirResponse) GetLastModify() int64`

GetLastModify returns the LastModify field if non-nil, zero value otherwise.

### GetLastModifyOk

`func (o *GetTftpfiledirResponse) GetLastModifyOk() (*int64, bool)`

GetLastModifyOk returns a tuple with the LastModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModify

`func (o *GetTftpfiledirResponse) SetLastModify(v int64)`

SetLastModify sets LastModify field to given value.

### HasLastModify

`func (o *GetTftpfiledirResponse) HasLastModify() bool`

HasLastModify returns a boolean if a field has been set.

### GetName

`func (o *GetTftpfiledirResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTftpfiledirResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTftpfiledirResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTftpfiledirResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetTftpfiledirResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTftpfiledirResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTftpfiledirResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTftpfiledirResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *GetTftpfiledirResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetTftpfiledirResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetTftpfiledirResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetTftpfiledirResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVtftpDirMembers

`func (o *GetTftpfiledirResponse) GetVtftpDirMembers() []TftpfiledirVtftpDirMembers`

GetVtftpDirMembers returns the VtftpDirMembers field if non-nil, zero value otherwise.

### GetVtftpDirMembersOk

`func (o *GetTftpfiledirResponse) GetVtftpDirMembersOk() (*[]TftpfiledirVtftpDirMembers, bool)`

GetVtftpDirMembersOk returns a tuple with the VtftpDirMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtftpDirMembers

`func (o *GetTftpfiledirResponse) SetVtftpDirMembers(v []TftpfiledirVtftpDirMembers)`

SetVtftpDirMembers sets VtftpDirMembers field to given value.

### HasVtftpDirMembers

`func (o *GetTftpfiledirResponse) HasVtftpDirMembers() bool`

HasVtftpDirMembers returns a boolean if a field has been set.

### GetResult

`func (o *GetTftpfiledirResponse) GetResult() Tftpfiledir`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetTftpfiledirResponse) GetResultOk() (*Tftpfiledir, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetTftpfiledirResponse) SetResult(v Tftpfiledir)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetTftpfiledirResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


