# GcpdnstaskgroupTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this task. | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the task is enabled or disabled. | [optional] 
**State** | Pointer to **string** | Current overall state of this task. | [optional] [readonly] 
**StateMsg** | Pointer to **string** | State message of the complete process. | [optional] [readonly] 
**Filter** | Pointer to **string** | Zone filter for this task. | [optional] 
**ScheduleInterval** | Pointer to **int64** | Periodic interval for this task. | [optional] 
**ScheduleUnits** | Pointer to **string** | Units for the schedule interval. | [optional] 
**GcpUser** | Pointer to **string** | Reference to associated Gcp user whose credentials are to be used for this task. | [optional] 
**StatusTimestamp** | Pointer to **int64** | The timestamp when the last state was logged. | [optional] [readonly] 
**LastRun** | Pointer to **int64** | The timestamp when the task was started last. | [optional] [readonly] 
**OverlayState** | Pointer to **string** | Current overlay state of this task. | [optional] [readonly] 
**OverlayStatusTimestamp** | Pointer to **int64** | Timestamp when the last overlay state was updated. | [optional] [readonly] 
**DnsSyncState** | Pointer to **string** | Current state of this task. Not accounting child tasks. | [optional] [readonly] 
**DnsSyncStatusTimestamp** | Pointer to **int64** | Timestamp when the last state was updated. Not accounting child tasks. | [optional] [readonly] 
**ProjectSyncState** | Pointer to **string** | Current state of project sync for this task. | [optional] [readonly] 
**ProjectSyncStatusTimestamp** | Pointer to **int64** | Timestamp when the last state was updated. Not accounting child tasks. | [optional] [readonly] 
**SyncPublicZones** | Pointer to **bool** | Indicates whether public zones are synchronized. | [optional] 
**SyncPrivateZones** | Pointer to **bool** | Indicates whether private zones are synchronized. | [optional] 
**SyncInterval** | Pointer to **int64** | Minimum delay between synchronization with Gcp DNS (in seconds). | [optional] [readonly] 
**ZoneCount** | Pointer to **int64** | Number of overall zones successfully synced by this Gcp DNS sync task and it&#39;s child tasks. | [optional] [readonly] 
**DnsSyncZoneCount** | Pointer to **int64** | Number of zones successfully synced by this Gcp DNS sync task. | [optional] [readonly] 
**CredentialsType** | Pointer to **string** | Credentials type used for connecting to the cloud management platform. | [optional] 

## Methods

### NewGcpdnstaskgroupTaskList

`func NewGcpdnstaskgroupTaskList() *GcpdnstaskgroupTaskList`

NewGcpdnstaskgroupTaskList instantiates a new GcpdnstaskgroupTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpdnstaskgroupTaskListWithDefaults

`func NewGcpdnstaskgroupTaskListWithDefaults() *GcpdnstaskgroupTaskList`

NewGcpdnstaskgroupTaskListWithDefaults instantiates a new GcpdnstaskgroupTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpdnstaskgroupTaskList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpdnstaskgroupTaskList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpdnstaskgroupTaskList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpdnstaskgroupTaskList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisabled

`func (o *GcpdnstaskgroupTaskList) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GcpdnstaskgroupTaskList) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GcpdnstaskgroupTaskList) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GcpdnstaskgroupTaskList) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetState

`func (o *GcpdnstaskgroupTaskList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GcpdnstaskgroupTaskList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GcpdnstaskgroupTaskList) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GcpdnstaskgroupTaskList) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMsg

`func (o *GcpdnstaskgroupTaskList) GetStateMsg() string`

GetStateMsg returns the StateMsg field if non-nil, zero value otherwise.

### GetStateMsgOk

`func (o *GcpdnstaskgroupTaskList) GetStateMsgOk() (*string, bool)`

GetStateMsgOk returns a tuple with the StateMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMsg

`func (o *GcpdnstaskgroupTaskList) SetStateMsg(v string)`

SetStateMsg sets StateMsg field to given value.

### HasStateMsg

`func (o *GcpdnstaskgroupTaskList) HasStateMsg() bool`

HasStateMsg returns a boolean if a field has been set.

### GetFilter

`func (o *GcpdnstaskgroupTaskList) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GcpdnstaskgroupTaskList) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GcpdnstaskgroupTaskList) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GcpdnstaskgroupTaskList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *GcpdnstaskgroupTaskList) GetScheduleInterval() int64`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *GcpdnstaskgroupTaskList) GetScheduleIntervalOk() (*int64, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *GcpdnstaskgroupTaskList) SetScheduleInterval(v int64)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *GcpdnstaskgroupTaskList) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetScheduleUnits

`func (o *GcpdnstaskgroupTaskList) GetScheduleUnits() string`

GetScheduleUnits returns the ScheduleUnits field if non-nil, zero value otherwise.

### GetScheduleUnitsOk

`func (o *GcpdnstaskgroupTaskList) GetScheduleUnitsOk() (*string, bool)`

GetScheduleUnitsOk returns a tuple with the ScheduleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUnits

`func (o *GcpdnstaskgroupTaskList) SetScheduleUnits(v string)`

SetScheduleUnits sets ScheduleUnits field to given value.

### HasScheduleUnits

`func (o *GcpdnstaskgroupTaskList) HasScheduleUnits() bool`

HasScheduleUnits returns a boolean if a field has been set.

### GetGcpUser

`func (o *GcpdnstaskgroupTaskList) GetGcpUser() string`

GetGcpUser returns the GcpUser field if non-nil, zero value otherwise.

### GetGcpUserOk

`func (o *GcpdnstaskgroupTaskList) GetGcpUserOk() (*string, bool)`

GetGcpUserOk returns a tuple with the GcpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpUser

`func (o *GcpdnstaskgroupTaskList) SetGcpUser(v string)`

SetGcpUser sets GcpUser field to given value.

### HasGcpUser

`func (o *GcpdnstaskgroupTaskList) HasGcpUser() bool`

HasGcpUser returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) GetStatusTimestamp() int64`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *GcpdnstaskgroupTaskList) GetStatusTimestampOk() (*int64, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) SetStatusTimestamp(v int64)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.

### GetLastRun

`func (o *GcpdnstaskgroupTaskList) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *GcpdnstaskgroupTaskList) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *GcpdnstaskgroupTaskList) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *GcpdnstaskgroupTaskList) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetOverlayState

`func (o *GcpdnstaskgroupTaskList) GetOverlayState() string`

GetOverlayState returns the OverlayState field if non-nil, zero value otherwise.

### GetOverlayStateOk

`func (o *GcpdnstaskgroupTaskList) GetOverlayStateOk() (*string, bool)`

GetOverlayStateOk returns a tuple with the OverlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayState

`func (o *GcpdnstaskgroupTaskList) SetOverlayState(v string)`

SetOverlayState sets OverlayState field to given value.

### HasOverlayState

`func (o *GcpdnstaskgroupTaskList) HasOverlayState() bool`

HasOverlayState returns a boolean if a field has been set.

### GetOverlayStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) GetOverlayStatusTimestamp() int64`

GetOverlayStatusTimestamp returns the OverlayStatusTimestamp field if non-nil, zero value otherwise.

### GetOverlayStatusTimestampOk

`func (o *GcpdnstaskgroupTaskList) GetOverlayStatusTimestampOk() (*int64, bool)`

GetOverlayStatusTimestampOk returns a tuple with the OverlayStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) SetOverlayStatusTimestamp(v int64)`

SetOverlayStatusTimestamp sets OverlayStatusTimestamp field to given value.

### HasOverlayStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) HasOverlayStatusTimestamp() bool`

HasOverlayStatusTimestamp returns a boolean if a field has been set.

### GetDnsSyncState

`func (o *GcpdnstaskgroupTaskList) GetDnsSyncState() string`

GetDnsSyncState returns the DnsSyncState field if non-nil, zero value otherwise.

### GetDnsSyncStateOk

`func (o *GcpdnstaskgroupTaskList) GetDnsSyncStateOk() (*string, bool)`

GetDnsSyncStateOk returns a tuple with the DnsSyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSyncState

`func (o *GcpdnstaskgroupTaskList) SetDnsSyncState(v string)`

SetDnsSyncState sets DnsSyncState field to given value.

### HasDnsSyncState

`func (o *GcpdnstaskgroupTaskList) HasDnsSyncState() bool`

HasDnsSyncState returns a boolean if a field has been set.

### GetDnsSyncStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) GetDnsSyncStatusTimestamp() int64`

GetDnsSyncStatusTimestamp returns the DnsSyncStatusTimestamp field if non-nil, zero value otherwise.

### GetDnsSyncStatusTimestampOk

`func (o *GcpdnstaskgroupTaskList) GetDnsSyncStatusTimestampOk() (*int64, bool)`

GetDnsSyncStatusTimestampOk returns a tuple with the DnsSyncStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSyncStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) SetDnsSyncStatusTimestamp(v int64)`

SetDnsSyncStatusTimestamp sets DnsSyncStatusTimestamp field to given value.

### HasDnsSyncStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) HasDnsSyncStatusTimestamp() bool`

HasDnsSyncStatusTimestamp returns a boolean if a field has been set.

### GetProjectSyncState

`func (o *GcpdnstaskgroupTaskList) GetProjectSyncState() string`

GetProjectSyncState returns the ProjectSyncState field if non-nil, zero value otherwise.

### GetProjectSyncStateOk

`func (o *GcpdnstaskgroupTaskList) GetProjectSyncStateOk() (*string, bool)`

GetProjectSyncStateOk returns a tuple with the ProjectSyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSyncState

`func (o *GcpdnstaskgroupTaskList) SetProjectSyncState(v string)`

SetProjectSyncState sets ProjectSyncState field to given value.

### HasProjectSyncState

`func (o *GcpdnstaskgroupTaskList) HasProjectSyncState() bool`

HasProjectSyncState returns a boolean if a field has been set.

### GetProjectSyncStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) GetProjectSyncStatusTimestamp() int64`

GetProjectSyncStatusTimestamp returns the ProjectSyncStatusTimestamp field if non-nil, zero value otherwise.

### GetProjectSyncStatusTimestampOk

`func (o *GcpdnstaskgroupTaskList) GetProjectSyncStatusTimestampOk() (*int64, bool)`

GetProjectSyncStatusTimestampOk returns a tuple with the ProjectSyncStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSyncStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) SetProjectSyncStatusTimestamp(v int64)`

SetProjectSyncStatusTimestamp sets ProjectSyncStatusTimestamp field to given value.

### HasProjectSyncStatusTimestamp

`func (o *GcpdnstaskgroupTaskList) HasProjectSyncStatusTimestamp() bool`

HasProjectSyncStatusTimestamp returns a boolean if a field has been set.

### GetSyncPublicZones

`func (o *GcpdnstaskgroupTaskList) GetSyncPublicZones() bool`

GetSyncPublicZones returns the SyncPublicZones field if non-nil, zero value otherwise.

### GetSyncPublicZonesOk

`func (o *GcpdnstaskgroupTaskList) GetSyncPublicZonesOk() (*bool, bool)`

GetSyncPublicZonesOk returns a tuple with the SyncPublicZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPublicZones

`func (o *GcpdnstaskgroupTaskList) SetSyncPublicZones(v bool)`

SetSyncPublicZones sets SyncPublicZones field to given value.

### HasSyncPublicZones

`func (o *GcpdnstaskgroupTaskList) HasSyncPublicZones() bool`

HasSyncPublicZones returns a boolean if a field has been set.

### GetSyncPrivateZones

`func (o *GcpdnstaskgroupTaskList) GetSyncPrivateZones() bool`

GetSyncPrivateZones returns the SyncPrivateZones field if non-nil, zero value otherwise.

### GetSyncPrivateZonesOk

`func (o *GcpdnstaskgroupTaskList) GetSyncPrivateZonesOk() (*bool, bool)`

GetSyncPrivateZonesOk returns a tuple with the SyncPrivateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPrivateZones

`func (o *GcpdnstaskgroupTaskList) SetSyncPrivateZones(v bool)`

SetSyncPrivateZones sets SyncPrivateZones field to given value.

### HasSyncPrivateZones

`func (o *GcpdnstaskgroupTaskList) HasSyncPrivateZones() bool`

HasSyncPrivateZones returns a boolean if a field has been set.

### GetSyncInterval

`func (o *GcpdnstaskgroupTaskList) GetSyncInterval() int64`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *GcpdnstaskgroupTaskList) GetSyncIntervalOk() (*int64, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *GcpdnstaskgroupTaskList) SetSyncInterval(v int64)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *GcpdnstaskgroupTaskList) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### GetZoneCount

`func (o *GcpdnstaskgroupTaskList) GetZoneCount() int64`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *GcpdnstaskgroupTaskList) GetZoneCountOk() (*int64, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *GcpdnstaskgroupTaskList) SetZoneCount(v int64)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *GcpdnstaskgroupTaskList) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.

### GetDnsSyncZoneCount

`func (o *GcpdnstaskgroupTaskList) GetDnsSyncZoneCount() int64`

GetDnsSyncZoneCount returns the DnsSyncZoneCount field if non-nil, zero value otherwise.

### GetDnsSyncZoneCountOk

`func (o *GcpdnstaskgroupTaskList) GetDnsSyncZoneCountOk() (*int64, bool)`

GetDnsSyncZoneCountOk returns a tuple with the DnsSyncZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSyncZoneCount

`func (o *GcpdnstaskgroupTaskList) SetDnsSyncZoneCount(v int64)`

SetDnsSyncZoneCount sets DnsSyncZoneCount field to given value.

### HasDnsSyncZoneCount

`func (o *GcpdnstaskgroupTaskList) HasDnsSyncZoneCount() bool`

HasDnsSyncZoneCount returns a boolean if a field has been set.

### GetCredentialsType

`func (o *GcpdnstaskgroupTaskList) GetCredentialsType() string`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *GcpdnstaskgroupTaskList) GetCredentialsTypeOk() (*string, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *GcpdnstaskgroupTaskList) SetCredentialsType(v string)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *GcpdnstaskgroupTaskList) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


