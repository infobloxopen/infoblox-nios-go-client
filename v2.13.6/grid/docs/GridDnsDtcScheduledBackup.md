# GridDnsDtcScheduledBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the scheduled backup. | [optional] [readonly] 
**Execute** | Pointer to **string** | The state for scheduled backup or restore operation. | [optional] 
**Operation** | Pointer to **string** | The scheduled backup operation. | [optional] 
**BackupType** | Pointer to **string** | The destination of the backup files. | [optional] 
**KeepLocalCopy** | Pointer to **bool** | Determines whether the local backup performed before uploading backup to remote storage. | [optional] 
**BackupFrequency** | Pointer to **string** | The frequency of backups. | [optional] 
**Weekday** | Pointer to **string** | The day of the week when the backup is performed. | [optional] 
**HourOfDay** | Pointer to **int64** | The hour of the day past 12:00 AM the backup is performed. | [optional] 
**MinutesPastHour** | Pointer to **int64** | The minute of the hour when the backup is performed. | [optional] 
**Username** | Pointer to **string** | The user name on the backup server. | [optional] 
**Password** | Pointer to **string** | The user password on the backup server. | [optional] 
**BackupServer** | Pointer to **string** | The IP address of the backup server. | [optional] 
**Path** | Pointer to **string** | The directory path to the backup file stored on the server. | [optional] 
**RestoreType** | Pointer to **string** | The destination of the restore files. | [optional] 
**RestoreServer** | Pointer to **string** | The IP address of the restore server. | [optional] 
**RestoreUsername** | Pointer to **string** | The user name on the restore server. | [optional] 
**RestorePassword** | Pointer to **string** | The password on the restore server. | [optional] 
**RestorePath** | Pointer to **string** | The directory path to the restored file on the server. | [optional] 
**NiosData** | Pointer to **bool** | Determines whether the restore of the NIOS data is enabled. | [optional] 
**DiscoveryData** | Pointer to **bool** | Determines whether the restore the NetMRI data is enabled. | [optional] 
**SplunkAppData** | Pointer to **bool** | Determines whether the restore of the Splunk application data is enabled. | [optional] 
**Enable** | Pointer to **bool** | Determines whether the scheduled backup is enabled. | [optional] 
**UseKeys** | Pointer to **bool** | If set, scp backup support based on keys | [optional] 
**KeyType** | Pointer to **string** | If set, scp backup support based on keys type | [optional] 
**UploadKeys** | Pointer to **bool** | If set, scp backup support to upload keys | [optional] 
**DownloadKeys** | Pointer to **bool** | If set, scp backup support to download keys | [optional] 

## Methods

### NewGridDnsDtcScheduledBackup

`func NewGridDnsDtcScheduledBackup() *GridDnsDtcScheduledBackup`

NewGridDnsDtcScheduledBackup instantiates a new GridDnsDtcScheduledBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsDtcScheduledBackupWithDefaults

`func NewGridDnsDtcScheduledBackupWithDefaults() *GridDnsDtcScheduledBackup`

NewGridDnsDtcScheduledBackupWithDefaults instantiates a new GridDnsDtcScheduledBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GridDnsDtcScheduledBackup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GridDnsDtcScheduledBackup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GridDnsDtcScheduledBackup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GridDnsDtcScheduledBackup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecute

`func (o *GridDnsDtcScheduledBackup) GetExecute() string`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *GridDnsDtcScheduledBackup) GetExecuteOk() (*string, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *GridDnsDtcScheduledBackup) SetExecute(v string)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *GridDnsDtcScheduledBackup) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetOperation

`func (o *GridDnsDtcScheduledBackup) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GridDnsDtcScheduledBackup) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GridDnsDtcScheduledBackup) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GridDnsDtcScheduledBackup) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetBackupType

`func (o *GridDnsDtcScheduledBackup) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *GridDnsDtcScheduledBackup) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *GridDnsDtcScheduledBackup) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *GridDnsDtcScheduledBackup) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetKeepLocalCopy

`func (o *GridDnsDtcScheduledBackup) GetKeepLocalCopy() bool`

GetKeepLocalCopy returns the KeepLocalCopy field if non-nil, zero value otherwise.

### GetKeepLocalCopyOk

`func (o *GridDnsDtcScheduledBackup) GetKeepLocalCopyOk() (*bool, bool)`

GetKeepLocalCopyOk returns a tuple with the KeepLocalCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLocalCopy

`func (o *GridDnsDtcScheduledBackup) SetKeepLocalCopy(v bool)`

SetKeepLocalCopy sets KeepLocalCopy field to given value.

### HasKeepLocalCopy

`func (o *GridDnsDtcScheduledBackup) HasKeepLocalCopy() bool`

HasKeepLocalCopy returns a boolean if a field has been set.

### GetBackupFrequency

`func (o *GridDnsDtcScheduledBackup) GetBackupFrequency() string`

GetBackupFrequency returns the BackupFrequency field if non-nil, zero value otherwise.

### GetBackupFrequencyOk

`func (o *GridDnsDtcScheduledBackup) GetBackupFrequencyOk() (*string, bool)`

GetBackupFrequencyOk returns a tuple with the BackupFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFrequency

`func (o *GridDnsDtcScheduledBackup) SetBackupFrequency(v string)`

SetBackupFrequency sets BackupFrequency field to given value.

### HasBackupFrequency

`func (o *GridDnsDtcScheduledBackup) HasBackupFrequency() bool`

HasBackupFrequency returns a boolean if a field has been set.

### GetWeekday

`func (o *GridDnsDtcScheduledBackup) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *GridDnsDtcScheduledBackup) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *GridDnsDtcScheduledBackup) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *GridDnsDtcScheduledBackup) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### GetHourOfDay

`func (o *GridDnsDtcScheduledBackup) GetHourOfDay() int64`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *GridDnsDtcScheduledBackup) GetHourOfDayOk() (*int64, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *GridDnsDtcScheduledBackup) SetHourOfDay(v int64)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *GridDnsDtcScheduledBackup) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.

### GetMinutesPastHour

`func (o *GridDnsDtcScheduledBackup) GetMinutesPastHour() int64`

GetMinutesPastHour returns the MinutesPastHour field if non-nil, zero value otherwise.

### GetMinutesPastHourOk

`func (o *GridDnsDtcScheduledBackup) GetMinutesPastHourOk() (*int64, bool)`

GetMinutesPastHourOk returns a tuple with the MinutesPastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesPastHour

`func (o *GridDnsDtcScheduledBackup) SetMinutesPastHour(v int64)`

SetMinutesPastHour sets MinutesPastHour field to given value.

### HasMinutesPastHour

`func (o *GridDnsDtcScheduledBackup) HasMinutesPastHour() bool`

HasMinutesPastHour returns a boolean if a field has been set.

### GetUsername

`func (o *GridDnsDtcScheduledBackup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GridDnsDtcScheduledBackup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GridDnsDtcScheduledBackup) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GridDnsDtcScheduledBackup) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GridDnsDtcScheduledBackup) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GridDnsDtcScheduledBackup) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GridDnsDtcScheduledBackup) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GridDnsDtcScheduledBackup) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetBackupServer

`func (o *GridDnsDtcScheduledBackup) GetBackupServer() string`

GetBackupServer returns the BackupServer field if non-nil, zero value otherwise.

### GetBackupServerOk

`func (o *GridDnsDtcScheduledBackup) GetBackupServerOk() (*string, bool)`

GetBackupServerOk returns a tuple with the BackupServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServer

`func (o *GridDnsDtcScheduledBackup) SetBackupServer(v string)`

SetBackupServer sets BackupServer field to given value.

### HasBackupServer

`func (o *GridDnsDtcScheduledBackup) HasBackupServer() bool`

HasBackupServer returns a boolean if a field has been set.

### GetPath

`func (o *GridDnsDtcScheduledBackup) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GridDnsDtcScheduledBackup) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GridDnsDtcScheduledBackup) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GridDnsDtcScheduledBackup) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRestoreType

`func (o *GridDnsDtcScheduledBackup) GetRestoreType() string`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *GridDnsDtcScheduledBackup) GetRestoreTypeOk() (*string, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *GridDnsDtcScheduledBackup) SetRestoreType(v string)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *GridDnsDtcScheduledBackup) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.

### GetRestoreServer

`func (o *GridDnsDtcScheduledBackup) GetRestoreServer() string`

GetRestoreServer returns the RestoreServer field if non-nil, zero value otherwise.

### GetRestoreServerOk

`func (o *GridDnsDtcScheduledBackup) GetRestoreServerOk() (*string, bool)`

GetRestoreServerOk returns a tuple with the RestoreServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServer

`func (o *GridDnsDtcScheduledBackup) SetRestoreServer(v string)`

SetRestoreServer sets RestoreServer field to given value.

### HasRestoreServer

`func (o *GridDnsDtcScheduledBackup) HasRestoreServer() bool`

HasRestoreServer returns a boolean if a field has been set.

### GetRestoreUsername

`func (o *GridDnsDtcScheduledBackup) GetRestoreUsername() string`

GetRestoreUsername returns the RestoreUsername field if non-nil, zero value otherwise.

### GetRestoreUsernameOk

`func (o *GridDnsDtcScheduledBackup) GetRestoreUsernameOk() (*string, bool)`

GetRestoreUsernameOk returns a tuple with the RestoreUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUsername

`func (o *GridDnsDtcScheduledBackup) SetRestoreUsername(v string)`

SetRestoreUsername sets RestoreUsername field to given value.

### HasRestoreUsername

`func (o *GridDnsDtcScheduledBackup) HasRestoreUsername() bool`

HasRestoreUsername returns a boolean if a field has been set.

### GetRestorePassword

`func (o *GridDnsDtcScheduledBackup) GetRestorePassword() string`

GetRestorePassword returns the RestorePassword field if non-nil, zero value otherwise.

### GetRestorePasswordOk

`func (o *GridDnsDtcScheduledBackup) GetRestorePasswordOk() (*string, bool)`

GetRestorePasswordOk returns a tuple with the RestorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePassword

`func (o *GridDnsDtcScheduledBackup) SetRestorePassword(v string)`

SetRestorePassword sets RestorePassword field to given value.

### HasRestorePassword

`func (o *GridDnsDtcScheduledBackup) HasRestorePassword() bool`

HasRestorePassword returns a boolean if a field has been set.

### GetRestorePath

`func (o *GridDnsDtcScheduledBackup) GetRestorePath() string`

GetRestorePath returns the RestorePath field if non-nil, zero value otherwise.

### GetRestorePathOk

`func (o *GridDnsDtcScheduledBackup) GetRestorePathOk() (*string, bool)`

GetRestorePathOk returns a tuple with the RestorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePath

`func (o *GridDnsDtcScheduledBackup) SetRestorePath(v string)`

SetRestorePath sets RestorePath field to given value.

### HasRestorePath

`func (o *GridDnsDtcScheduledBackup) HasRestorePath() bool`

HasRestorePath returns a boolean if a field has been set.

### GetNiosData

`func (o *GridDnsDtcScheduledBackup) GetNiosData() bool`

GetNiosData returns the NiosData field if non-nil, zero value otherwise.

### GetNiosDataOk

`func (o *GridDnsDtcScheduledBackup) GetNiosDataOk() (*bool, bool)`

GetNiosDataOk returns a tuple with the NiosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiosData

`func (o *GridDnsDtcScheduledBackup) SetNiosData(v bool)`

SetNiosData sets NiosData field to given value.

### HasNiosData

`func (o *GridDnsDtcScheduledBackup) HasNiosData() bool`

HasNiosData returns a boolean if a field has been set.

### GetDiscoveryData

`func (o *GridDnsDtcScheduledBackup) GetDiscoveryData() bool`

GetDiscoveryData returns the DiscoveryData field if non-nil, zero value otherwise.

### GetDiscoveryDataOk

`func (o *GridDnsDtcScheduledBackup) GetDiscoveryDataOk() (*bool, bool)`

GetDiscoveryDataOk returns a tuple with the DiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryData

`func (o *GridDnsDtcScheduledBackup) SetDiscoveryData(v bool)`

SetDiscoveryData sets DiscoveryData field to given value.

### HasDiscoveryData

`func (o *GridDnsDtcScheduledBackup) HasDiscoveryData() bool`

HasDiscoveryData returns a boolean if a field has been set.

### GetSplunkAppData

`func (o *GridDnsDtcScheduledBackup) GetSplunkAppData() bool`

GetSplunkAppData returns the SplunkAppData field if non-nil, zero value otherwise.

### GetSplunkAppDataOk

`func (o *GridDnsDtcScheduledBackup) GetSplunkAppDataOk() (*bool, bool)`

GetSplunkAppDataOk returns a tuple with the SplunkAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkAppData

`func (o *GridDnsDtcScheduledBackup) SetSplunkAppData(v bool)`

SetSplunkAppData sets SplunkAppData field to given value.

### HasSplunkAppData

`func (o *GridDnsDtcScheduledBackup) HasSplunkAppData() bool`

HasSplunkAppData returns a boolean if a field has been set.

### GetEnable

`func (o *GridDnsDtcScheduledBackup) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GridDnsDtcScheduledBackup) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GridDnsDtcScheduledBackup) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GridDnsDtcScheduledBackup) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetUseKeys

`func (o *GridDnsDtcScheduledBackup) GetUseKeys() bool`

GetUseKeys returns the UseKeys field if non-nil, zero value otherwise.

### GetUseKeysOk

`func (o *GridDnsDtcScheduledBackup) GetUseKeysOk() (*bool, bool)`

GetUseKeysOk returns a tuple with the UseKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKeys

`func (o *GridDnsDtcScheduledBackup) SetUseKeys(v bool)`

SetUseKeys sets UseKeys field to given value.

### HasUseKeys

`func (o *GridDnsDtcScheduledBackup) HasUseKeys() bool`

HasUseKeys returns a boolean if a field has been set.

### GetKeyType

`func (o *GridDnsDtcScheduledBackup) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GridDnsDtcScheduledBackup) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GridDnsDtcScheduledBackup) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *GridDnsDtcScheduledBackup) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetUploadKeys

`func (o *GridDnsDtcScheduledBackup) GetUploadKeys() bool`

GetUploadKeys returns the UploadKeys field if non-nil, zero value otherwise.

### GetUploadKeysOk

`func (o *GridDnsDtcScheduledBackup) GetUploadKeysOk() (*bool, bool)`

GetUploadKeysOk returns a tuple with the UploadKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadKeys

`func (o *GridDnsDtcScheduledBackup) SetUploadKeys(v bool)`

SetUploadKeys sets UploadKeys field to given value.

### HasUploadKeys

`func (o *GridDnsDtcScheduledBackup) HasUploadKeys() bool`

HasUploadKeys returns a boolean if a field has been set.

### GetDownloadKeys

`func (o *GridDnsDtcScheduledBackup) GetDownloadKeys() bool`

GetDownloadKeys returns the DownloadKeys field if non-nil, zero value otherwise.

### GetDownloadKeysOk

`func (o *GridDnsDtcScheduledBackup) GetDownloadKeysOk() (*bool, bool)`

GetDownloadKeysOk returns a tuple with the DownloadKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadKeys

`func (o *GridDnsDtcScheduledBackup) SetDownloadKeys(v bool)`

SetDownloadKeys sets DownloadKeys field to given value.

### HasDownloadKeys

`func (o *GridDnsDtcScheduledBackup) HasDownloadKeys() bool`

HasDownloadKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


