# Awsrte53taskgroupTaskList

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
**AwsUser** | Pointer to **string** | Reference to associated AWS user whose credentials are to be used for this task. | [optional] 
**StatusTimestamp** | Pointer to **int64** | The timestamp when the last state was logged. | [optional] [readonly] 
**LastRun** | Pointer to **int64** | The timestamp when the task was started last. | [optional] [readonly] 
**SyncPublicZones** | Pointer to **bool** | Indicates whether public zones are synchronized. | [optional] 
**SyncPrivateZones** | Pointer to **bool** | Indicates whether private zones are synchronized. | [optional] 
**ZoneCount** | Pointer to **int64** | The number of zones synchronized by this task. | [optional] [readonly] 
**CredentialsType** | Pointer to **string** | Credentials type used for connecting to the cloud management platform. | [optional] 

## Methods

### NewAwsrte53taskgroupTaskList

`func NewAwsrte53taskgroupTaskList() *Awsrte53taskgroupTaskList`

NewAwsrte53taskgroupTaskList instantiates a new Awsrte53taskgroupTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsrte53taskgroupTaskListWithDefaults

`func NewAwsrte53taskgroupTaskListWithDefaults() *Awsrte53taskgroupTaskList`

NewAwsrte53taskgroupTaskListWithDefaults instantiates a new Awsrte53taskgroupTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Awsrte53taskgroupTaskList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Awsrte53taskgroupTaskList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Awsrte53taskgroupTaskList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Awsrte53taskgroupTaskList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisabled

`func (o *Awsrte53taskgroupTaskList) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Awsrte53taskgroupTaskList) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Awsrte53taskgroupTaskList) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Awsrte53taskgroupTaskList) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetState

`func (o *Awsrte53taskgroupTaskList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Awsrte53taskgroupTaskList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Awsrte53taskgroupTaskList) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Awsrte53taskgroupTaskList) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMsg

`func (o *Awsrte53taskgroupTaskList) GetStateMsg() string`

GetStateMsg returns the StateMsg field if non-nil, zero value otherwise.

### GetStateMsgOk

`func (o *Awsrte53taskgroupTaskList) GetStateMsgOk() (*string, bool)`

GetStateMsgOk returns a tuple with the StateMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMsg

`func (o *Awsrte53taskgroupTaskList) SetStateMsg(v string)`

SetStateMsg sets StateMsg field to given value.

### HasStateMsg

`func (o *Awsrte53taskgroupTaskList) HasStateMsg() bool`

HasStateMsg returns a boolean if a field has been set.

### GetFilter

`func (o *Awsrte53taskgroupTaskList) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Awsrte53taskgroupTaskList) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Awsrte53taskgroupTaskList) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Awsrte53taskgroupTaskList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *Awsrte53taskgroupTaskList) GetScheduleInterval() int64`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *Awsrte53taskgroupTaskList) GetScheduleIntervalOk() (*int64, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *Awsrte53taskgroupTaskList) SetScheduleInterval(v int64)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *Awsrte53taskgroupTaskList) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetScheduleUnits

`func (o *Awsrte53taskgroupTaskList) GetScheduleUnits() string`

GetScheduleUnits returns the ScheduleUnits field if non-nil, zero value otherwise.

### GetScheduleUnitsOk

`func (o *Awsrte53taskgroupTaskList) GetScheduleUnitsOk() (*string, bool)`

GetScheduleUnitsOk returns a tuple with the ScheduleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUnits

`func (o *Awsrte53taskgroupTaskList) SetScheduleUnits(v string)`

SetScheduleUnits sets ScheduleUnits field to given value.

### HasScheduleUnits

`func (o *Awsrte53taskgroupTaskList) HasScheduleUnits() bool`

HasScheduleUnits returns a boolean if a field has been set.

### GetAwsUser

`func (o *Awsrte53taskgroupTaskList) GetAwsUser() string`

GetAwsUser returns the AwsUser field if non-nil, zero value otherwise.

### GetAwsUserOk

`func (o *Awsrte53taskgroupTaskList) GetAwsUserOk() (*string, bool)`

GetAwsUserOk returns a tuple with the AwsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUser

`func (o *Awsrte53taskgroupTaskList) SetAwsUser(v string)`

SetAwsUser sets AwsUser field to given value.

### HasAwsUser

`func (o *Awsrte53taskgroupTaskList) HasAwsUser() bool`

HasAwsUser returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *Awsrte53taskgroupTaskList) GetStatusTimestamp() int64`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *Awsrte53taskgroupTaskList) GetStatusTimestampOk() (*int64, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *Awsrte53taskgroupTaskList) SetStatusTimestamp(v int64)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *Awsrte53taskgroupTaskList) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.

### GetLastRun

`func (o *Awsrte53taskgroupTaskList) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *Awsrte53taskgroupTaskList) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *Awsrte53taskgroupTaskList) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *Awsrte53taskgroupTaskList) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetSyncPublicZones

`func (o *Awsrte53taskgroupTaskList) GetSyncPublicZones() bool`

GetSyncPublicZones returns the SyncPublicZones field if non-nil, zero value otherwise.

### GetSyncPublicZonesOk

`func (o *Awsrte53taskgroupTaskList) GetSyncPublicZonesOk() (*bool, bool)`

GetSyncPublicZonesOk returns a tuple with the SyncPublicZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPublicZones

`func (o *Awsrte53taskgroupTaskList) SetSyncPublicZones(v bool)`

SetSyncPublicZones sets SyncPublicZones field to given value.

### HasSyncPublicZones

`func (o *Awsrte53taskgroupTaskList) HasSyncPublicZones() bool`

HasSyncPublicZones returns a boolean if a field has been set.

### GetSyncPrivateZones

`func (o *Awsrte53taskgroupTaskList) GetSyncPrivateZones() bool`

GetSyncPrivateZones returns the SyncPrivateZones field if non-nil, zero value otherwise.

### GetSyncPrivateZonesOk

`func (o *Awsrte53taskgroupTaskList) GetSyncPrivateZonesOk() (*bool, bool)`

GetSyncPrivateZonesOk returns a tuple with the SyncPrivateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPrivateZones

`func (o *Awsrte53taskgroupTaskList) SetSyncPrivateZones(v bool)`

SetSyncPrivateZones sets SyncPrivateZones field to given value.

### HasSyncPrivateZones

`func (o *Awsrte53taskgroupTaskList) HasSyncPrivateZones() bool`

HasSyncPrivateZones returns a boolean if a field has been set.

### GetZoneCount

`func (o *Awsrte53taskgroupTaskList) GetZoneCount() int64`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *Awsrte53taskgroupTaskList) GetZoneCountOk() (*int64, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *Awsrte53taskgroupTaskList) SetZoneCount(v int64)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *Awsrte53taskgroupTaskList) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.

### GetCredentialsType

`func (o *Awsrte53taskgroupTaskList) GetCredentialsType() string`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *Awsrte53taskgroupTaskList) GetCredentialsTypeOk() (*string, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *Awsrte53taskgroupTaskList) SetCredentialsType(v string)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *Awsrte53taskgroupTaskList) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


