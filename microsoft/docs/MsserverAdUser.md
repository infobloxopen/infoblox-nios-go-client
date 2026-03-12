# MsserverAdUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginName** | Pointer to **string** | The login name of the Microsoft Server. | [optional] 
**LoginPassword** | Pointer to **string** | The login password of the DHCP Microsoft Server. | [optional] 
**EnableUserSync** | Pointer to **bool** | Determines whether the Active Directory user synchronization is enabled or not. | [optional] 
**SynchronizationInterval** | Pointer to **int64** | The minimum number of minutes between two synchronizations. | [optional] 
**LastSyncTime** | Pointer to **int64** | Timestamp of the last synchronization attempt. | [optional] [readonly] 
**LastSyncStatus** | Pointer to **string** | The status of the last synchronization attempt. | [optional] [readonly] 
**LastSyncDetail** | Pointer to **string** | The detailed status of the last synchronization attempt. | [optional] [readonly] 
**LastSuccessSyncTime** | Pointer to **int64** | Timestamp of the last successful synchronization attempt. | [optional] [readonly] 
**UseLogin** | Pointer to **bool** | Flag to override login name and password from MS server | [optional] 
**UseEnableAdUserSync** | Pointer to **bool** | Flag to override AD User sync from grid level | [optional] 
**UseSynchronizationMinDelay** | Pointer to **bool** | Flag to override synchronization interval from the MS Server | [optional] 
**UseEnableUserSync** | Pointer to **bool** | Use flag for: enable_user_sync | [optional] 
**UseSynchronizationInterval** | Pointer to **bool** | Use flag for: synchronization_interval | [optional] 

## Methods

### NewMsserverAdUser

`func NewMsserverAdUser() *MsserverAdUser`

NewMsserverAdUser instantiates a new MsserverAdUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverAdUserWithDefaults

`func NewMsserverAdUserWithDefaults() *MsserverAdUser`

NewMsserverAdUserWithDefaults instantiates a new MsserverAdUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginName

`func (o *MsserverAdUser) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MsserverAdUser) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MsserverAdUser) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MsserverAdUser) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *MsserverAdUser) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *MsserverAdUser) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *MsserverAdUser) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *MsserverAdUser) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetEnableUserSync

`func (o *MsserverAdUser) GetEnableUserSync() bool`

GetEnableUserSync returns the EnableUserSync field if non-nil, zero value otherwise.

### GetEnableUserSyncOk

`func (o *MsserverAdUser) GetEnableUserSyncOk() (*bool, bool)`

GetEnableUserSyncOk returns a tuple with the EnableUserSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserSync

`func (o *MsserverAdUser) SetEnableUserSync(v bool)`

SetEnableUserSync sets EnableUserSync field to given value.

### HasEnableUserSync

`func (o *MsserverAdUser) HasEnableUserSync() bool`

HasEnableUserSync returns a boolean if a field has been set.

### GetSynchronizationInterval

`func (o *MsserverAdUser) GetSynchronizationInterval() int64`

GetSynchronizationInterval returns the SynchronizationInterval field if non-nil, zero value otherwise.

### GetSynchronizationIntervalOk

`func (o *MsserverAdUser) GetSynchronizationIntervalOk() (*int64, bool)`

GetSynchronizationIntervalOk returns a tuple with the SynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationInterval

`func (o *MsserverAdUser) SetSynchronizationInterval(v int64)`

SetSynchronizationInterval sets SynchronizationInterval field to given value.

### HasSynchronizationInterval

`func (o *MsserverAdUser) HasSynchronizationInterval() bool`

HasSynchronizationInterval returns a boolean if a field has been set.

### GetLastSyncTime

`func (o *MsserverAdUser) GetLastSyncTime() int64`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *MsserverAdUser) GetLastSyncTimeOk() (*int64, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *MsserverAdUser) SetLastSyncTime(v int64)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *MsserverAdUser) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.

### GetLastSyncStatus

`func (o *MsserverAdUser) GetLastSyncStatus() string`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *MsserverAdUser) GetLastSyncStatusOk() (*string, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *MsserverAdUser) SetLastSyncStatus(v string)`

SetLastSyncStatus sets LastSyncStatus field to given value.

### HasLastSyncStatus

`func (o *MsserverAdUser) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.

### GetLastSyncDetail

`func (o *MsserverAdUser) GetLastSyncDetail() string`

GetLastSyncDetail returns the LastSyncDetail field if non-nil, zero value otherwise.

### GetLastSyncDetailOk

`func (o *MsserverAdUser) GetLastSyncDetailOk() (*string, bool)`

GetLastSyncDetailOk returns a tuple with the LastSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDetail

`func (o *MsserverAdUser) SetLastSyncDetail(v string)`

SetLastSyncDetail sets LastSyncDetail field to given value.

### HasLastSyncDetail

`func (o *MsserverAdUser) HasLastSyncDetail() bool`

HasLastSyncDetail returns a boolean if a field has been set.

### GetLastSuccessSyncTime

`func (o *MsserverAdUser) GetLastSuccessSyncTime() int64`

GetLastSuccessSyncTime returns the LastSuccessSyncTime field if non-nil, zero value otherwise.

### GetLastSuccessSyncTimeOk

`func (o *MsserverAdUser) GetLastSuccessSyncTimeOk() (*int64, bool)`

GetLastSuccessSyncTimeOk returns a tuple with the LastSuccessSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessSyncTime

`func (o *MsserverAdUser) SetLastSuccessSyncTime(v int64)`

SetLastSuccessSyncTime sets LastSuccessSyncTime field to given value.

### HasLastSuccessSyncTime

`func (o *MsserverAdUser) HasLastSuccessSyncTime() bool`

HasLastSuccessSyncTime returns a boolean if a field has been set.

### GetUseLogin

`func (o *MsserverAdUser) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *MsserverAdUser) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *MsserverAdUser) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *MsserverAdUser) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetUseEnableAdUserSync

`func (o *MsserverAdUser) GetUseEnableAdUserSync() bool`

GetUseEnableAdUserSync returns the UseEnableAdUserSync field if non-nil, zero value otherwise.

### GetUseEnableAdUserSyncOk

`func (o *MsserverAdUser) GetUseEnableAdUserSyncOk() (*bool, bool)`

GetUseEnableAdUserSyncOk returns a tuple with the UseEnableAdUserSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableAdUserSync

`func (o *MsserverAdUser) SetUseEnableAdUserSync(v bool)`

SetUseEnableAdUserSync sets UseEnableAdUserSync field to given value.

### HasUseEnableAdUserSync

`func (o *MsserverAdUser) HasUseEnableAdUserSync() bool`

HasUseEnableAdUserSync returns a boolean if a field has been set.

### GetUseSynchronizationMinDelay

`func (o *MsserverAdUser) GetUseSynchronizationMinDelay() bool`

GetUseSynchronizationMinDelay returns the UseSynchronizationMinDelay field if non-nil, zero value otherwise.

### GetUseSynchronizationMinDelayOk

`func (o *MsserverAdUser) GetUseSynchronizationMinDelayOk() (*bool, bool)`

GetUseSynchronizationMinDelayOk returns a tuple with the UseSynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationMinDelay

`func (o *MsserverAdUser) SetUseSynchronizationMinDelay(v bool)`

SetUseSynchronizationMinDelay sets UseSynchronizationMinDelay field to given value.

### HasUseSynchronizationMinDelay

`func (o *MsserverAdUser) HasUseSynchronizationMinDelay() bool`

HasUseSynchronizationMinDelay returns a boolean if a field has been set.

### GetUseEnableUserSync

`func (o *MsserverAdUser) GetUseEnableUserSync() bool`

GetUseEnableUserSync returns the UseEnableUserSync field if non-nil, zero value otherwise.

### GetUseEnableUserSyncOk

`func (o *MsserverAdUser) GetUseEnableUserSyncOk() (*bool, bool)`

GetUseEnableUserSyncOk returns a tuple with the UseEnableUserSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableUserSync

`func (o *MsserverAdUser) SetUseEnableUserSync(v bool)`

SetUseEnableUserSync sets UseEnableUserSync field to given value.

### HasUseEnableUserSync

`func (o *MsserverAdUser) HasUseEnableUserSync() bool`

HasUseEnableUserSync returns a boolean if a field has been set.

### GetUseSynchronizationInterval

`func (o *MsserverAdUser) GetUseSynchronizationInterval() bool`

GetUseSynchronizationInterval returns the UseSynchronizationInterval field if non-nil, zero value otherwise.

### GetUseSynchronizationIntervalOk

`func (o *MsserverAdUser) GetUseSynchronizationIntervalOk() (*bool, bool)`

GetUseSynchronizationIntervalOk returns a tuple with the UseSynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationInterval

`func (o *MsserverAdUser) SetUseSynchronizationInterval(v bool)`

SetUseSynchronizationInterval sets UseSynchronizationInterval field to given value.

### HasUseSynchronizationInterval

`func (o *MsserverAdUser) HasUseSynchronizationInterval() bool`

HasUseSynchronizationInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


