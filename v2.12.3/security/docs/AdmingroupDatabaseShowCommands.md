# AdmingroupDatabaseShowCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowNamedMaxJournalSize** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowTxnTrace** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDatabaseTransferStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowBackup** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDbPh** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDbsize** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIbdbstat** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupDatabaseShowCommands

`func NewAdmingroupDatabaseShowCommands() *AdmingroupDatabaseShowCommands`

NewAdmingroupDatabaseShowCommands instantiates a new AdmingroupDatabaseShowCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupDatabaseShowCommandsWithDefaults

`func NewAdmingroupDatabaseShowCommandsWithDefaults() *AdmingroupDatabaseShowCommands`

NewAdmingroupDatabaseShowCommandsWithDefaults instantiates a new AdmingroupDatabaseShowCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowNamedMaxJournalSize

`func (o *AdmingroupDatabaseShowCommands) GetShowNamedMaxJournalSize() bool`

GetShowNamedMaxJournalSize returns the ShowNamedMaxJournalSize field if non-nil, zero value otherwise.

### GetShowNamedMaxJournalSizeOk

`func (o *AdmingroupDatabaseShowCommands) GetShowNamedMaxJournalSizeOk() (*bool, bool)`

GetShowNamedMaxJournalSizeOk returns a tuple with the ShowNamedMaxJournalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNamedMaxJournalSize

`func (o *AdmingroupDatabaseShowCommands) SetShowNamedMaxJournalSize(v bool)`

SetShowNamedMaxJournalSize sets ShowNamedMaxJournalSize field to given value.

### HasShowNamedMaxJournalSize

`func (o *AdmingroupDatabaseShowCommands) HasShowNamedMaxJournalSize() bool`

HasShowNamedMaxJournalSize returns a boolean if a field has been set.

### GetShowTxnTrace

`func (o *AdmingroupDatabaseShowCommands) GetShowTxnTrace() bool`

GetShowTxnTrace returns the ShowTxnTrace field if non-nil, zero value otherwise.

### GetShowTxnTraceOk

`func (o *AdmingroupDatabaseShowCommands) GetShowTxnTraceOk() (*bool, bool)`

GetShowTxnTraceOk returns a tuple with the ShowTxnTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTxnTrace

`func (o *AdmingroupDatabaseShowCommands) SetShowTxnTrace(v bool)`

SetShowTxnTrace sets ShowTxnTrace field to given value.

### HasShowTxnTrace

`func (o *AdmingroupDatabaseShowCommands) HasShowTxnTrace() bool`

HasShowTxnTrace returns a boolean if a field has been set.

### GetShowDatabaseTransferStatus

`func (o *AdmingroupDatabaseShowCommands) GetShowDatabaseTransferStatus() bool`

GetShowDatabaseTransferStatus returns the ShowDatabaseTransferStatus field if non-nil, zero value otherwise.

### GetShowDatabaseTransferStatusOk

`func (o *AdmingroupDatabaseShowCommands) GetShowDatabaseTransferStatusOk() (*bool, bool)`

GetShowDatabaseTransferStatusOk returns a tuple with the ShowDatabaseTransferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDatabaseTransferStatus

`func (o *AdmingroupDatabaseShowCommands) SetShowDatabaseTransferStatus(v bool)`

SetShowDatabaseTransferStatus sets ShowDatabaseTransferStatus field to given value.

### HasShowDatabaseTransferStatus

`func (o *AdmingroupDatabaseShowCommands) HasShowDatabaseTransferStatus() bool`

HasShowDatabaseTransferStatus returns a boolean if a field has been set.

### GetShowBackup

`func (o *AdmingroupDatabaseShowCommands) GetShowBackup() bool`

GetShowBackup returns the ShowBackup field if non-nil, zero value otherwise.

### GetShowBackupOk

`func (o *AdmingroupDatabaseShowCommands) GetShowBackupOk() (*bool, bool)`

GetShowBackupOk returns a tuple with the ShowBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBackup

`func (o *AdmingroupDatabaseShowCommands) SetShowBackup(v bool)`

SetShowBackup sets ShowBackup field to given value.

### HasShowBackup

`func (o *AdmingroupDatabaseShowCommands) HasShowBackup() bool`

HasShowBackup returns a boolean if a field has been set.

### GetShowDbPh

`func (o *AdmingroupDatabaseShowCommands) GetShowDbPh() bool`

GetShowDbPh returns the ShowDbPh field if non-nil, zero value otherwise.

### GetShowDbPhOk

`func (o *AdmingroupDatabaseShowCommands) GetShowDbPhOk() (*bool, bool)`

GetShowDbPhOk returns a tuple with the ShowDbPh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDbPh

`func (o *AdmingroupDatabaseShowCommands) SetShowDbPh(v bool)`

SetShowDbPh sets ShowDbPh field to given value.

### HasShowDbPh

`func (o *AdmingroupDatabaseShowCommands) HasShowDbPh() bool`

HasShowDbPh returns a boolean if a field has been set.

### GetShowDbsize

`func (o *AdmingroupDatabaseShowCommands) GetShowDbsize() bool`

GetShowDbsize returns the ShowDbsize field if non-nil, zero value otherwise.

### GetShowDbsizeOk

`func (o *AdmingroupDatabaseShowCommands) GetShowDbsizeOk() (*bool, bool)`

GetShowDbsizeOk returns a tuple with the ShowDbsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDbsize

`func (o *AdmingroupDatabaseShowCommands) SetShowDbsize(v bool)`

SetShowDbsize sets ShowDbsize field to given value.

### HasShowDbsize

`func (o *AdmingroupDatabaseShowCommands) HasShowDbsize() bool`

HasShowDbsize returns a boolean if a field has been set.

### GetShowIbdbstat

`func (o *AdmingroupDatabaseShowCommands) GetShowIbdbstat() bool`

GetShowIbdbstat returns the ShowIbdbstat field if non-nil, zero value otherwise.

### GetShowIbdbstatOk

`func (o *AdmingroupDatabaseShowCommands) GetShowIbdbstatOk() (*bool, bool)`

GetShowIbdbstatOk returns a tuple with the ShowIbdbstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIbdbstat

`func (o *AdmingroupDatabaseShowCommands) SetShowIbdbstat(v bool)`

SetShowIbdbstat sets ShowIbdbstat field to given value.

### HasShowIbdbstat

`func (o *AdmingroupDatabaseShowCommands) HasShowIbdbstat() bool`

HasShowIbdbstat returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupDatabaseShowCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupDatabaseShowCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupDatabaseShowCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupDatabaseShowCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupDatabaseShowCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupDatabaseShowCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupDatabaseShowCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupDatabaseShowCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


