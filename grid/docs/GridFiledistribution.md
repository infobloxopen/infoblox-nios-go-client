# GridFiledistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowUploads** | Pointer to **bool** | Determines whether the uploads to Grid members are allowed. | [optional] 
**BackupStorage** | Pointer to **bool** | Determines whether to include distributed files in the backup. | [optional] 
**CurrentUsage** | Pointer to **int64** | The value is the percentage of the allocated TFTP storage space that is used, expressed in tenth of a percent. Valid values are from 0 to 1000. | [optional] [readonly] 
**EnableAnonymousFtp** | Pointer to **bool** | Determines whether the FTP anonymous login is enabled. | [optional] 
**GlobalStatus** | Pointer to **string** | The Grid file distribution global status. | [optional] [readonly] 
**Name** | Pointer to **string** | The Grid name. | [optional] [readonly] 
**StorageLimit** | Pointer to **int64** | Maximum storage in megabytes allowed on the Grid. The maximum storage space allowed for all file distribution services on a Grid is equal to the storage space allowed to the Grid member with the smallest amount of space allowed. | [optional] 

## Methods

### NewGridFiledistribution

`func NewGridFiledistribution() *GridFiledistribution`

NewGridFiledistribution instantiates a new GridFiledistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridFiledistributionWithDefaults

`func NewGridFiledistributionWithDefaults() *GridFiledistribution`

NewGridFiledistributionWithDefaults instantiates a new GridFiledistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridFiledistribution) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridFiledistribution) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridFiledistribution) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridFiledistribution) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowUploads

`func (o *GridFiledistribution) GetAllowUploads() bool`

GetAllowUploads returns the AllowUploads field if non-nil, zero value otherwise.

### GetAllowUploadsOk

`func (o *GridFiledistribution) GetAllowUploadsOk() (*bool, bool)`

GetAllowUploadsOk returns a tuple with the AllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploads

`func (o *GridFiledistribution) SetAllowUploads(v bool)`

SetAllowUploads sets AllowUploads field to given value.

### HasAllowUploads

`func (o *GridFiledistribution) HasAllowUploads() bool`

HasAllowUploads returns a boolean if a field has been set.

### GetBackupStorage

`func (o *GridFiledistribution) GetBackupStorage() bool`

GetBackupStorage returns the BackupStorage field if non-nil, zero value otherwise.

### GetBackupStorageOk

`func (o *GridFiledistribution) GetBackupStorageOk() (*bool, bool)`

GetBackupStorageOk returns a tuple with the BackupStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorage

`func (o *GridFiledistribution) SetBackupStorage(v bool)`

SetBackupStorage sets BackupStorage field to given value.

### HasBackupStorage

`func (o *GridFiledistribution) HasBackupStorage() bool`

HasBackupStorage returns a boolean if a field has been set.

### GetCurrentUsage

`func (o *GridFiledistribution) GetCurrentUsage() int64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *GridFiledistribution) GetCurrentUsageOk() (*int64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *GridFiledistribution) SetCurrentUsage(v int64)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *GridFiledistribution) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.

### GetEnableAnonymousFtp

`func (o *GridFiledistribution) GetEnableAnonymousFtp() bool`

GetEnableAnonymousFtp returns the EnableAnonymousFtp field if non-nil, zero value otherwise.

### GetEnableAnonymousFtpOk

`func (o *GridFiledistribution) GetEnableAnonymousFtpOk() (*bool, bool)`

GetEnableAnonymousFtpOk returns a tuple with the EnableAnonymousFtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAnonymousFtp

`func (o *GridFiledistribution) SetEnableAnonymousFtp(v bool)`

SetEnableAnonymousFtp sets EnableAnonymousFtp field to given value.

### HasEnableAnonymousFtp

`func (o *GridFiledistribution) HasEnableAnonymousFtp() bool`

HasEnableAnonymousFtp returns a boolean if a field has been set.

### GetGlobalStatus

`func (o *GridFiledistribution) GetGlobalStatus() string`

GetGlobalStatus returns the GlobalStatus field if non-nil, zero value otherwise.

### GetGlobalStatusOk

`func (o *GridFiledistribution) GetGlobalStatusOk() (*string, bool)`

GetGlobalStatusOk returns a tuple with the GlobalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStatus

`func (o *GridFiledistribution) SetGlobalStatus(v string)`

SetGlobalStatus sets GlobalStatus field to given value.

### HasGlobalStatus

`func (o *GridFiledistribution) HasGlobalStatus() bool`

HasGlobalStatus returns a boolean if a field has been set.

### GetName

`func (o *GridFiledistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridFiledistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridFiledistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridFiledistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageLimit

`func (o *GridFiledistribution) GetStorageLimit() int64`

GetStorageLimit returns the StorageLimit field if non-nil, zero value otherwise.

### GetStorageLimitOk

`func (o *GridFiledistribution) GetStorageLimitOk() (*int64, bool)`

GetStorageLimitOk returns a tuple with the StorageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimit

`func (o *GridFiledistribution) SetStorageLimit(v int64)`

SetStorageLimit sets StorageLimit field to given value.

### HasStorageLimit

`func (o *GridFiledistribution) HasStorageLimit() bool`

HasStorageLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


