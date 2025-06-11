# AzurednstaskgroupTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this task. | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the task is enabled or disabled. | [optional] 
**State** | Pointer to **string** | Indicate the sync status of this task. | [optional] [readonly] 
**StateMsg** | Pointer to **string** | State message for the task. | [optional] [readonly] 
**Filter** | Pointer to **string** | Filter for this task. | [optional] 
**ScheduleInterval** | Pointer to **int64** | Periodic interval for this task. | [optional] 
**ScheduleUnits** | Pointer to **string** | Units for the schedule interval. | [optional] 
**AzureUser** | Pointer to **string** | Reference to associated Azure user whose credentials are to be used for this task. | [optional] 
**StatusTimestamp** | Pointer to **int64** | The timestamp when the last state was logged. | [optional] [readonly] 
**LastRun** | Pointer to **int64** | The timestamp when the task was started last. | [optional] [readonly] 
**SyncPublicZones** | Pointer to **bool** | Indicates whether public zones are synchronized. | [optional] 
**SyncPrivateZones** | Pointer to **bool** | Indicates whether private zones are synchronized. | [optional] 
**ZoneCount** | Pointer to **int64** | The number of zones synchronized by this task. | [optional] [readonly] 
**CredentialsType** | Pointer to **string** | Credentials type used for connecting to the cloud management platform. | [optional] 

## Methods

### NewAzurednstaskgroupTaskList

`func NewAzurednstaskgroupTaskList() *AzurednstaskgroupTaskList`

NewAzurednstaskgroupTaskList instantiates a new AzurednstaskgroupTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurednstaskgroupTaskListWithDefaults

`func NewAzurednstaskgroupTaskListWithDefaults() *AzurednstaskgroupTaskList`

NewAzurednstaskgroupTaskListWithDefaults instantiates a new AzurednstaskgroupTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzurednstaskgroupTaskList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzurednstaskgroupTaskList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzurednstaskgroupTaskList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzurednstaskgroupTaskList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisabled

`func (o *AzurednstaskgroupTaskList) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AzurednstaskgroupTaskList) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AzurednstaskgroupTaskList) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AzurednstaskgroupTaskList) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetState

`func (o *AzurednstaskgroupTaskList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzurednstaskgroupTaskList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzurednstaskgroupTaskList) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AzurednstaskgroupTaskList) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMsg

`func (o *AzurednstaskgroupTaskList) GetStateMsg() string`

GetStateMsg returns the StateMsg field if non-nil, zero value otherwise.

### GetStateMsgOk

`func (o *AzurednstaskgroupTaskList) GetStateMsgOk() (*string, bool)`

GetStateMsgOk returns a tuple with the StateMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMsg

`func (o *AzurednstaskgroupTaskList) SetStateMsg(v string)`

SetStateMsg sets StateMsg field to given value.

### HasStateMsg

`func (o *AzurednstaskgroupTaskList) HasStateMsg() bool`

HasStateMsg returns a boolean if a field has been set.

### GetFilter

`func (o *AzurednstaskgroupTaskList) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AzurednstaskgroupTaskList) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AzurednstaskgroupTaskList) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AzurednstaskgroupTaskList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *AzurednstaskgroupTaskList) GetScheduleInterval() int64`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *AzurednstaskgroupTaskList) GetScheduleIntervalOk() (*int64, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *AzurednstaskgroupTaskList) SetScheduleInterval(v int64)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *AzurednstaskgroupTaskList) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetScheduleUnits

`func (o *AzurednstaskgroupTaskList) GetScheduleUnits() string`

GetScheduleUnits returns the ScheduleUnits field if non-nil, zero value otherwise.

### GetScheduleUnitsOk

`func (o *AzurednstaskgroupTaskList) GetScheduleUnitsOk() (*string, bool)`

GetScheduleUnitsOk returns a tuple with the ScheduleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUnits

`func (o *AzurednstaskgroupTaskList) SetScheduleUnits(v string)`

SetScheduleUnits sets ScheduleUnits field to given value.

### HasScheduleUnits

`func (o *AzurednstaskgroupTaskList) HasScheduleUnits() bool`

HasScheduleUnits returns a boolean if a field has been set.

### GetAzureUser

`func (o *AzurednstaskgroupTaskList) GetAzureUser() string`

GetAzureUser returns the AzureUser field if non-nil, zero value otherwise.

### GetAzureUserOk

`func (o *AzurednstaskgroupTaskList) GetAzureUserOk() (*string, bool)`

GetAzureUserOk returns a tuple with the AzureUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureUser

`func (o *AzurednstaskgroupTaskList) SetAzureUser(v string)`

SetAzureUser sets AzureUser field to given value.

### HasAzureUser

`func (o *AzurednstaskgroupTaskList) HasAzureUser() bool`

HasAzureUser returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *AzurednstaskgroupTaskList) GetStatusTimestamp() int64`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *AzurednstaskgroupTaskList) GetStatusTimestampOk() (*int64, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *AzurednstaskgroupTaskList) SetStatusTimestamp(v int64)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *AzurednstaskgroupTaskList) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.

### GetLastRun

`func (o *AzurednstaskgroupTaskList) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *AzurednstaskgroupTaskList) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *AzurednstaskgroupTaskList) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *AzurednstaskgroupTaskList) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetSyncPublicZones

`func (o *AzurednstaskgroupTaskList) GetSyncPublicZones() bool`

GetSyncPublicZones returns the SyncPublicZones field if non-nil, zero value otherwise.

### GetSyncPublicZonesOk

`func (o *AzurednstaskgroupTaskList) GetSyncPublicZonesOk() (*bool, bool)`

GetSyncPublicZonesOk returns a tuple with the SyncPublicZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPublicZones

`func (o *AzurednstaskgroupTaskList) SetSyncPublicZones(v bool)`

SetSyncPublicZones sets SyncPublicZones field to given value.

### HasSyncPublicZones

`func (o *AzurednstaskgroupTaskList) HasSyncPublicZones() bool`

HasSyncPublicZones returns a boolean if a field has been set.

### GetSyncPrivateZones

`func (o *AzurednstaskgroupTaskList) GetSyncPrivateZones() bool`

GetSyncPrivateZones returns the SyncPrivateZones field if non-nil, zero value otherwise.

### GetSyncPrivateZonesOk

`func (o *AzurednstaskgroupTaskList) GetSyncPrivateZonesOk() (*bool, bool)`

GetSyncPrivateZonesOk returns a tuple with the SyncPrivateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPrivateZones

`func (o *AzurednstaskgroupTaskList) SetSyncPrivateZones(v bool)`

SetSyncPrivateZones sets SyncPrivateZones field to given value.

### HasSyncPrivateZones

`func (o *AzurednstaskgroupTaskList) HasSyncPrivateZones() bool`

HasSyncPrivateZones returns a boolean if a field has been set.

### GetZoneCount

`func (o *AzurednstaskgroupTaskList) GetZoneCount() int64`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *AzurednstaskgroupTaskList) GetZoneCountOk() (*int64, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *AzurednstaskgroupTaskList) SetZoneCount(v int64)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *AzurednstaskgroupTaskList) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.

### GetCredentialsType

`func (o *AzurednstaskgroupTaskList) GetCredentialsType() string`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *AzurednstaskgroupTaskList) GetCredentialsTypeOk() (*string, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *AzurednstaskgroupTaskList) SetCredentialsType(v string)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *AzurednstaskgroupTaskList) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


