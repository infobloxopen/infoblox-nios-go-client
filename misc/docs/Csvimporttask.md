# Csvimporttask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Action** | Pointer to **string** | The action to execute. | [optional] 
**AdminName** | Pointer to **string** | The login name of the administrator. | [optional] [readonly] 
**EndTime** | Pointer to **int64** | The end time of this import operation. | [optional] [readonly] 
**FileName** | Pointer to **string** | The name of the file used for the import operation. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | The size of the file used for the import operation. | [optional] [readonly] 
**ImportId** | Pointer to **int64** | The ID of the current import task. | [optional] [readonly] 
**LinesFailed** | Pointer to **int64** | The number of lines that encountered an error. | [optional] [readonly] 
**LinesProcessed** | Pointer to **int64** | The number of lines that have been processed. | [optional] [readonly] 
**LinesWarning** | Pointer to **int64** | The number of lines that encountered a warning. | [optional] [readonly] 
**OnError** | Pointer to **string** | The action to take when an error is encountered. | [optional] 
**Operation** | Pointer to **string** | The operation to execute. | [optional] 
**Separator** | Pointer to **string** | The separator to be used for the data in the CSV file. | [optional] [readonly] 
**StartTime** | Pointer to **int64** | The start time of the import operation. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the import operation | [optional] [readonly] 
**UpdateMethod** | Pointer to **string** | The update method to be used for the operation. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewCsvimporttask

`func NewCsvimporttask() *Csvimporttask`

NewCsvimporttask instantiates a new Csvimporttask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsvimporttaskWithDefaults

`func NewCsvimporttaskWithDefaults() *Csvimporttask`

NewCsvimporttaskWithDefaults instantiates a new Csvimporttask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Csvimporttask) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Csvimporttask) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Csvimporttask) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Csvimporttask) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAction

`func (o *Csvimporttask) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Csvimporttask) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Csvimporttask) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Csvimporttask) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAdminName

`func (o *Csvimporttask) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *Csvimporttask) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *Csvimporttask) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *Csvimporttask) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetEndTime

`func (o *Csvimporttask) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Csvimporttask) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Csvimporttask) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Csvimporttask) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFileName

`func (o *Csvimporttask) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Csvimporttask) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Csvimporttask) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Csvimporttask) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *Csvimporttask) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Csvimporttask) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Csvimporttask) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *Csvimporttask) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetImportId

`func (o *Csvimporttask) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Csvimporttask) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *Csvimporttask) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *Csvimporttask) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetLinesFailed

`func (o *Csvimporttask) GetLinesFailed() int64`

GetLinesFailed returns the LinesFailed field if non-nil, zero value otherwise.

### GetLinesFailedOk

`func (o *Csvimporttask) GetLinesFailedOk() (*int64, bool)`

GetLinesFailedOk returns a tuple with the LinesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesFailed

`func (o *Csvimporttask) SetLinesFailed(v int64)`

SetLinesFailed sets LinesFailed field to given value.

### HasLinesFailed

`func (o *Csvimporttask) HasLinesFailed() bool`

HasLinesFailed returns a boolean if a field has been set.

### GetLinesProcessed

`func (o *Csvimporttask) GetLinesProcessed() int64`

GetLinesProcessed returns the LinesProcessed field if non-nil, zero value otherwise.

### GetLinesProcessedOk

`func (o *Csvimporttask) GetLinesProcessedOk() (*int64, bool)`

GetLinesProcessedOk returns a tuple with the LinesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesProcessed

`func (o *Csvimporttask) SetLinesProcessed(v int64)`

SetLinesProcessed sets LinesProcessed field to given value.

### HasLinesProcessed

`func (o *Csvimporttask) HasLinesProcessed() bool`

HasLinesProcessed returns a boolean if a field has been set.

### GetLinesWarning

`func (o *Csvimporttask) GetLinesWarning() int64`

GetLinesWarning returns the LinesWarning field if non-nil, zero value otherwise.

### GetLinesWarningOk

`func (o *Csvimporttask) GetLinesWarningOk() (*int64, bool)`

GetLinesWarningOk returns a tuple with the LinesWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesWarning

`func (o *Csvimporttask) SetLinesWarning(v int64)`

SetLinesWarning sets LinesWarning field to given value.

### HasLinesWarning

`func (o *Csvimporttask) HasLinesWarning() bool`

HasLinesWarning returns a boolean if a field has been set.

### GetOnError

`func (o *Csvimporttask) GetOnError() string`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *Csvimporttask) GetOnErrorOk() (*string, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *Csvimporttask) SetOnError(v string)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *Csvimporttask) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetOperation

`func (o *Csvimporttask) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Csvimporttask) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Csvimporttask) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Csvimporttask) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSeparator

`func (o *Csvimporttask) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *Csvimporttask) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *Csvimporttask) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *Csvimporttask) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetStartTime

`func (o *Csvimporttask) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Csvimporttask) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Csvimporttask) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Csvimporttask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *Csvimporttask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Csvimporttask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Csvimporttask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Csvimporttask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateMethod

`func (o *Csvimporttask) GetUpdateMethod() string`

GetUpdateMethod returns the UpdateMethod field if non-nil, zero value otherwise.

### GetUpdateMethodOk

`func (o *Csvimporttask) GetUpdateMethodOk() (*string, bool)`

GetUpdateMethodOk returns a tuple with the UpdateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMethod

`func (o *Csvimporttask) SetUpdateMethod(v string)`

SetUpdateMethod sets UpdateMethod field to given value.

### HasUpdateMethod

`func (o *Csvimporttask) HasUpdateMethod() bool`

HasUpdateMethod returns a boolean if a field has been set.

### GetUuid

`func (o *Csvimporttask) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Csvimporttask) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Csvimporttask) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Csvimporttask) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


