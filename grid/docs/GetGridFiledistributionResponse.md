# GetGridFiledistributionResponse

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
**Result** | Pointer to [**GridFiledistribution**](GridFiledistribution.md) |  | [optional] 

## Methods

### NewGetGridFiledistributionResponse

`func NewGetGridFiledistributionResponse() *GetGridFiledistributionResponse`

NewGetGridFiledistributionResponse instantiates a new GetGridFiledistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridFiledistributionResponseWithDefaults

`func NewGetGridFiledistributionResponseWithDefaults() *GetGridFiledistributionResponse`

NewGetGridFiledistributionResponseWithDefaults instantiates a new GetGridFiledistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridFiledistributionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridFiledistributionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridFiledistributionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridFiledistributionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowUploads

`func (o *GetGridFiledistributionResponse) GetAllowUploads() bool`

GetAllowUploads returns the AllowUploads field if non-nil, zero value otherwise.

### GetAllowUploadsOk

`func (o *GetGridFiledistributionResponse) GetAllowUploadsOk() (*bool, bool)`

GetAllowUploadsOk returns a tuple with the AllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploads

`func (o *GetGridFiledistributionResponse) SetAllowUploads(v bool)`

SetAllowUploads sets AllowUploads field to given value.

### HasAllowUploads

`func (o *GetGridFiledistributionResponse) HasAllowUploads() bool`

HasAllowUploads returns a boolean if a field has been set.

### GetBackupStorage

`func (o *GetGridFiledistributionResponse) GetBackupStorage() bool`

GetBackupStorage returns the BackupStorage field if non-nil, zero value otherwise.

### GetBackupStorageOk

`func (o *GetGridFiledistributionResponse) GetBackupStorageOk() (*bool, bool)`

GetBackupStorageOk returns a tuple with the BackupStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorage

`func (o *GetGridFiledistributionResponse) SetBackupStorage(v bool)`

SetBackupStorage sets BackupStorage field to given value.

### HasBackupStorage

`func (o *GetGridFiledistributionResponse) HasBackupStorage() bool`

HasBackupStorage returns a boolean if a field has been set.

### GetCurrentUsage

`func (o *GetGridFiledistributionResponse) GetCurrentUsage() int64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *GetGridFiledistributionResponse) GetCurrentUsageOk() (*int64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *GetGridFiledistributionResponse) SetCurrentUsage(v int64)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *GetGridFiledistributionResponse) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.

### GetEnableAnonymousFtp

`func (o *GetGridFiledistributionResponse) GetEnableAnonymousFtp() bool`

GetEnableAnonymousFtp returns the EnableAnonymousFtp field if non-nil, zero value otherwise.

### GetEnableAnonymousFtpOk

`func (o *GetGridFiledistributionResponse) GetEnableAnonymousFtpOk() (*bool, bool)`

GetEnableAnonymousFtpOk returns a tuple with the EnableAnonymousFtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAnonymousFtp

`func (o *GetGridFiledistributionResponse) SetEnableAnonymousFtp(v bool)`

SetEnableAnonymousFtp sets EnableAnonymousFtp field to given value.

### HasEnableAnonymousFtp

`func (o *GetGridFiledistributionResponse) HasEnableAnonymousFtp() bool`

HasEnableAnonymousFtp returns a boolean if a field has been set.

### GetGlobalStatus

`func (o *GetGridFiledistributionResponse) GetGlobalStatus() string`

GetGlobalStatus returns the GlobalStatus field if non-nil, zero value otherwise.

### GetGlobalStatusOk

`func (o *GetGridFiledistributionResponse) GetGlobalStatusOk() (*string, bool)`

GetGlobalStatusOk returns a tuple with the GlobalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStatus

`func (o *GetGridFiledistributionResponse) SetGlobalStatus(v string)`

SetGlobalStatus sets GlobalStatus field to given value.

### HasGlobalStatus

`func (o *GetGridFiledistributionResponse) HasGlobalStatus() bool`

HasGlobalStatus returns a boolean if a field has been set.

### GetName

`func (o *GetGridFiledistributionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGridFiledistributionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGridFiledistributionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGridFiledistributionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageLimit

`func (o *GetGridFiledistributionResponse) GetStorageLimit() int64`

GetStorageLimit returns the StorageLimit field if non-nil, zero value otherwise.

### GetStorageLimitOk

`func (o *GetGridFiledistributionResponse) GetStorageLimitOk() (*int64, bool)`

GetStorageLimitOk returns a tuple with the StorageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimit

`func (o *GetGridFiledistributionResponse) SetStorageLimit(v int64)`

SetStorageLimit sets StorageLimit field to given value.

### HasStorageLimit

`func (o *GetGridFiledistributionResponse) HasStorageLimit() bool`

HasStorageLimit returns a boolean if a field has been set.

### GetResult

`func (o *GetGridFiledistributionResponse) GetResult() GridFiledistribution`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridFiledistributionResponse) GetResultOk() (*GridFiledistribution, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridFiledistributionResponse) SetResult(v GridFiledistribution)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridFiledistributionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


