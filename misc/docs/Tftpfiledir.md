# Tftpfiledir

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Directory** | Pointer to **string** | The path to the directory that contains file or subdirectory. | [optional] 
**IsSyncedToGm** | Pointer to **bool** | Determines whether the TFTP entity is synchronized to Grid Master. | [optional] [readonly] 
**LastModify** | Pointer to **int64** | The time when the file or directory was last modified. | [optional] [readonly] 
**Name** | Pointer to **string** | The TFTP directory or file name. | [optional] 
**Type** | Pointer to **string** | The type of TFTP file system entity (directory or file). | [optional] 
**VtftpDirMembers** | Pointer to [**[]TftpfiledirVtftpDirMembers**](TftpfiledirVtftpDirMembers.md) | The replication members with TFTP client addresses where this virtual folder is applicable. | [optional] 

## Methods

### NewTftpfiledir

`func NewTftpfiledir() *Tftpfiledir`

NewTftpfiledir instantiates a new Tftpfiledir object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTftpfiledirWithDefaults

`func NewTftpfiledirWithDefaults() *Tftpfiledir`

NewTftpfiledirWithDefaults instantiates a new Tftpfiledir object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Tftpfiledir) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Tftpfiledir) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Tftpfiledir) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Tftpfiledir) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDirectory

`func (o *Tftpfiledir) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *Tftpfiledir) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *Tftpfiledir) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *Tftpfiledir) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetIsSyncedToGm

`func (o *Tftpfiledir) GetIsSyncedToGm() bool`

GetIsSyncedToGm returns the IsSyncedToGm field if non-nil, zero value otherwise.

### GetIsSyncedToGmOk

`func (o *Tftpfiledir) GetIsSyncedToGmOk() (*bool, bool)`

GetIsSyncedToGmOk returns a tuple with the IsSyncedToGm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncedToGm

`func (o *Tftpfiledir) SetIsSyncedToGm(v bool)`

SetIsSyncedToGm sets IsSyncedToGm field to given value.

### HasIsSyncedToGm

`func (o *Tftpfiledir) HasIsSyncedToGm() bool`

HasIsSyncedToGm returns a boolean if a field has been set.

### GetLastModify

`func (o *Tftpfiledir) GetLastModify() int64`

GetLastModify returns the LastModify field if non-nil, zero value otherwise.

### GetLastModifyOk

`func (o *Tftpfiledir) GetLastModifyOk() (*int64, bool)`

GetLastModifyOk returns a tuple with the LastModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModify

`func (o *Tftpfiledir) SetLastModify(v int64)`

SetLastModify sets LastModify field to given value.

### HasLastModify

`func (o *Tftpfiledir) HasLastModify() bool`

HasLastModify returns a boolean if a field has been set.

### GetName

`func (o *Tftpfiledir) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tftpfiledir) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tftpfiledir) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tftpfiledir) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Tftpfiledir) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tftpfiledir) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tftpfiledir) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Tftpfiledir) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVtftpDirMembers

`func (o *Tftpfiledir) GetVtftpDirMembers() []TftpfiledirVtftpDirMembers`

GetVtftpDirMembers returns the VtftpDirMembers field if non-nil, zero value otherwise.

### GetVtftpDirMembersOk

`func (o *Tftpfiledir) GetVtftpDirMembersOk() (*[]TftpfiledirVtftpDirMembers, bool)`

GetVtftpDirMembersOk returns a tuple with the VtftpDirMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtftpDirMembers

`func (o *Tftpfiledir) SetVtftpDirMembers(v []TftpfiledirVtftpDirMembers)`

SetVtftpDirMembers sets VtftpDirMembers field to given value.

### HasVtftpDirMembers

`func (o *Tftpfiledir) HasVtftpDirMembers() bool`

HasVtftpDirMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


