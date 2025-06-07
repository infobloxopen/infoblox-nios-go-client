# Userprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ActiveDashboardType** | Pointer to **string** | Determines the active dashboard type. | [optional] 
**AdminGroup** | Pointer to **string** | The Admin Group object to which the admin belongs. An admin user can belong to only one admin group at a time. | [optional] [readonly] 
**DaysToExpire** | Pointer to **int32** | The number of days left before the admin&#39;s password expires. | [optional] [readonly] 
**Email** | Pointer to **string** | The email address of the admin. | [optional] 
**GlobalSearchOnEa** | Pointer to **bool** | Determines if extensible attribute values will be returned by global search or not. | [optional] 
**GlobalSearchOnNiData** | Pointer to **bool** | Determines if global search will search for network insight devices and interfaces or not. | [optional] 
**GridAdminGroups** | Pointer to **[]string** | List of Admin Group objects that the current user is mapped to. | [optional] [readonly] 
**LastLogin** | Pointer to **int64** | The timestamp when the admin last logged in. | [optional] [readonly] 
**LbTreeNodesAtGenLevel** | Pointer to **int64** | Determines how many nodes are displayed at generation levels. | [optional] 
**LbTreeNodesAtLastLevel** | Pointer to **int64** | Determines how many nodes are displayed at the last level. | [optional] 
**MaxCountWidgets** | Pointer to **int64** | The maximum count of widgets that can be added to one dashboard. | [optional] 
**Name** | Pointer to **string** | The admin name. | [optional] [readonly] 
**OldPassword** | Pointer to **string** | The current password that will be replaced by a new password. To change a password in the database, you must provide both the current and new password values. This is a write-only attribute. | [optional] 
**Password** | Pointer to **string** | The new password of the admin. To change a password in the database, you must provide both the current and new password values. This is a write-only attribute. | [optional] 
**TableSize** | Pointer to **int64** | The number of lines of data a table or a single list view can contain. | [optional] 
**TimeZone** | Pointer to **string** | The time zone of the admin user. | [optional] 
**UseTimeZone** | Pointer to **bool** | Use flag for: time_zone | [optional] 
**UserType** | Pointer to **string** | The admin type. | [optional] [readonly] 

## Methods

### NewUserprofile

`func NewUserprofile() *Userprofile`

NewUserprofile instantiates a new Userprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserprofileWithDefaults

`func NewUserprofileWithDefaults() *Userprofile`

NewUserprofileWithDefaults instantiates a new Userprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Userprofile) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Userprofile) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Userprofile) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Userprofile) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActiveDashboardType

`func (o *Userprofile) GetActiveDashboardType() string`

GetActiveDashboardType returns the ActiveDashboardType field if non-nil, zero value otherwise.

### GetActiveDashboardTypeOk

`func (o *Userprofile) GetActiveDashboardTypeOk() (*string, bool)`

GetActiveDashboardTypeOk returns a tuple with the ActiveDashboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDashboardType

`func (o *Userprofile) SetActiveDashboardType(v string)`

SetActiveDashboardType sets ActiveDashboardType field to given value.

### HasActiveDashboardType

`func (o *Userprofile) HasActiveDashboardType() bool`

HasActiveDashboardType returns a boolean if a field has been set.

### GetAdminGroup

`func (o *Userprofile) GetAdminGroup() string`

GetAdminGroup returns the AdminGroup field if non-nil, zero value otherwise.

### GetAdminGroupOk

`func (o *Userprofile) GetAdminGroupOk() (*string, bool)`

GetAdminGroupOk returns a tuple with the AdminGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroup

`func (o *Userprofile) SetAdminGroup(v string)`

SetAdminGroup sets AdminGroup field to given value.

### HasAdminGroup

`func (o *Userprofile) HasAdminGroup() bool`

HasAdminGroup returns a boolean if a field has been set.

### GetDaysToExpire

`func (o *Userprofile) GetDaysToExpire() int32`

GetDaysToExpire returns the DaysToExpire field if non-nil, zero value otherwise.

### GetDaysToExpireOk

`func (o *Userprofile) GetDaysToExpireOk() (*int32, bool)`

GetDaysToExpireOk returns a tuple with the DaysToExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExpire

`func (o *Userprofile) SetDaysToExpire(v int32)`

SetDaysToExpire sets DaysToExpire field to given value.

### HasDaysToExpire

`func (o *Userprofile) HasDaysToExpire() bool`

HasDaysToExpire returns a boolean if a field has been set.

### GetEmail

`func (o *Userprofile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Userprofile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Userprofile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Userprofile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGlobalSearchOnEa

`func (o *Userprofile) GetGlobalSearchOnEa() bool`

GetGlobalSearchOnEa returns the GlobalSearchOnEa field if non-nil, zero value otherwise.

### GetGlobalSearchOnEaOk

`func (o *Userprofile) GetGlobalSearchOnEaOk() (*bool, bool)`

GetGlobalSearchOnEaOk returns a tuple with the GlobalSearchOnEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchOnEa

`func (o *Userprofile) SetGlobalSearchOnEa(v bool)`

SetGlobalSearchOnEa sets GlobalSearchOnEa field to given value.

### HasGlobalSearchOnEa

`func (o *Userprofile) HasGlobalSearchOnEa() bool`

HasGlobalSearchOnEa returns a boolean if a field has been set.

### GetGlobalSearchOnNiData

`func (o *Userprofile) GetGlobalSearchOnNiData() bool`

GetGlobalSearchOnNiData returns the GlobalSearchOnNiData field if non-nil, zero value otherwise.

### GetGlobalSearchOnNiDataOk

`func (o *Userprofile) GetGlobalSearchOnNiDataOk() (*bool, bool)`

GetGlobalSearchOnNiDataOk returns a tuple with the GlobalSearchOnNiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchOnNiData

`func (o *Userprofile) SetGlobalSearchOnNiData(v bool)`

SetGlobalSearchOnNiData sets GlobalSearchOnNiData field to given value.

### HasGlobalSearchOnNiData

`func (o *Userprofile) HasGlobalSearchOnNiData() bool`

HasGlobalSearchOnNiData returns a boolean if a field has been set.

### GetGridAdminGroups

`func (o *Userprofile) GetGridAdminGroups() []string`

GetGridAdminGroups returns the GridAdminGroups field if non-nil, zero value otherwise.

### GetGridAdminGroupsOk

`func (o *Userprofile) GetGridAdminGroupsOk() (*[]string, bool)`

GetGridAdminGroupsOk returns a tuple with the GridAdminGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridAdminGroups

`func (o *Userprofile) SetGridAdminGroups(v []string)`

SetGridAdminGroups sets GridAdminGroups field to given value.

### HasGridAdminGroups

`func (o *Userprofile) HasGridAdminGroups() bool`

HasGridAdminGroups returns a boolean if a field has been set.

### GetLastLogin

`func (o *Userprofile) GetLastLogin() int64`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Userprofile) GetLastLoginOk() (*int64, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Userprofile) SetLastLogin(v int64)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Userprofile) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLbTreeNodesAtGenLevel

`func (o *Userprofile) GetLbTreeNodesAtGenLevel() int64`

GetLbTreeNodesAtGenLevel returns the LbTreeNodesAtGenLevel field if non-nil, zero value otherwise.

### GetLbTreeNodesAtGenLevelOk

`func (o *Userprofile) GetLbTreeNodesAtGenLevelOk() (*int64, bool)`

GetLbTreeNodesAtGenLevelOk returns a tuple with the LbTreeNodesAtGenLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbTreeNodesAtGenLevel

`func (o *Userprofile) SetLbTreeNodesAtGenLevel(v int64)`

SetLbTreeNodesAtGenLevel sets LbTreeNodesAtGenLevel field to given value.

### HasLbTreeNodesAtGenLevel

`func (o *Userprofile) HasLbTreeNodesAtGenLevel() bool`

HasLbTreeNodesAtGenLevel returns a boolean if a field has been set.

### GetLbTreeNodesAtLastLevel

`func (o *Userprofile) GetLbTreeNodesAtLastLevel() int64`

GetLbTreeNodesAtLastLevel returns the LbTreeNodesAtLastLevel field if non-nil, zero value otherwise.

### GetLbTreeNodesAtLastLevelOk

`func (o *Userprofile) GetLbTreeNodesAtLastLevelOk() (*int64, bool)`

GetLbTreeNodesAtLastLevelOk returns a tuple with the LbTreeNodesAtLastLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbTreeNodesAtLastLevel

`func (o *Userprofile) SetLbTreeNodesAtLastLevel(v int64)`

SetLbTreeNodesAtLastLevel sets LbTreeNodesAtLastLevel field to given value.

### HasLbTreeNodesAtLastLevel

`func (o *Userprofile) HasLbTreeNodesAtLastLevel() bool`

HasLbTreeNodesAtLastLevel returns a boolean if a field has been set.

### GetMaxCountWidgets

`func (o *Userprofile) GetMaxCountWidgets() int64`

GetMaxCountWidgets returns the MaxCountWidgets field if non-nil, zero value otherwise.

### GetMaxCountWidgetsOk

`func (o *Userprofile) GetMaxCountWidgetsOk() (*int64, bool)`

GetMaxCountWidgetsOk returns a tuple with the MaxCountWidgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCountWidgets

`func (o *Userprofile) SetMaxCountWidgets(v int64)`

SetMaxCountWidgets sets MaxCountWidgets field to given value.

### HasMaxCountWidgets

`func (o *Userprofile) HasMaxCountWidgets() bool`

HasMaxCountWidgets returns a boolean if a field has been set.

### GetName

`func (o *Userprofile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Userprofile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Userprofile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Userprofile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOldPassword

`func (o *Userprofile) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *Userprofile) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *Userprofile) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *Userprofile) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetPassword

`func (o *Userprofile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Userprofile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Userprofile) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Userprofile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTableSize

`func (o *Userprofile) GetTableSize() int64`

GetTableSize returns the TableSize field if non-nil, zero value otherwise.

### GetTableSizeOk

`func (o *Userprofile) GetTableSizeOk() (*int64, bool)`

GetTableSizeOk returns a tuple with the TableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableSize

`func (o *Userprofile) SetTableSize(v int64)`

SetTableSize sets TableSize field to given value.

### HasTableSize

`func (o *Userprofile) HasTableSize() bool`

HasTableSize returns a boolean if a field has been set.

### GetTimeZone

`func (o *Userprofile) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Userprofile) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Userprofile) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Userprofile) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUseTimeZone

`func (o *Userprofile) GetUseTimeZone() bool`

GetUseTimeZone returns the UseTimeZone field if non-nil, zero value otherwise.

### GetUseTimeZoneOk

`func (o *Userprofile) GetUseTimeZoneOk() (*bool, bool)`

GetUseTimeZoneOk returns a tuple with the UseTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeZone

`func (o *Userprofile) SetUseTimeZone(v bool)`

SetUseTimeZone sets UseTimeZone field to given value.

### HasUseTimeZone

`func (o *Userprofile) HasUseTimeZone() bool`

HasUseTimeZone returns a boolean if a field has been set.

### GetUserType

`func (o *Userprofile) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *Userprofile) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *Userprofile) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *Userprofile) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


