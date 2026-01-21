# Gcpdnstaskgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the task group; maximum 256 characters. | [optional] 
**ConsolidateZones** | Pointer to **bool** | Indicates if all zones need to be saved into a single view. | [optional] 
**ConsolidatedView** | Pointer to **string** | Consolidate all zones into one view. | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the task group is enabled or disabled. | [optional] 
**GcpProjectIdsFileToken** | Pointer to **string** | The Gcp project IDs file&#39;s token. | [optional] 
**GridMember** | Pointer to **string** | Grid member configured to run this task group. | [optional] 
**MultipleProjectsSyncPolicy** | Pointer to **string** | Discover all child projects or Upload child project ids to discover. | [optional] 
**Name** | Pointer to **string** | The name of this Gcp DNS sync task group. | [optional] 
**NetworkView** | Pointer to **string** | The name of the tenant&#39;s network view. | [optional] 
**NetworkViewMappingPolicy** | Pointer to **string** | The network view mapping policy. | [optional] 
**ProjectId** | Pointer to **string** | The Gcp project ID associated with this task group. | [optional] [readonly] 
**ProjectsList** | Pointer to **string** | The Gcp Project IDs list associated with this task. | [optional] [readonly] 
**SyncChildProjects** | Pointer to **bool** | Synchronizing child projects is enabled or disabled. | [optional] 
**SyncStatus** | Pointer to **string** | Sync status for this task. | [optional] [readonly] 
**TaskList** | Pointer to [**[]GcpdnstaskgroupTaskList**](GcpdnstaskgroupTaskList.md) | List of Gcp DNS tasks in this group. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewGcpdnstaskgroup

`func NewGcpdnstaskgroup() *Gcpdnstaskgroup`

NewGcpdnstaskgroup instantiates a new Gcpdnstaskgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpdnstaskgroupWithDefaults

`func NewGcpdnstaskgroupWithDefaults() *Gcpdnstaskgroup`

NewGcpdnstaskgroupWithDefaults instantiates a new Gcpdnstaskgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Gcpdnstaskgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Gcpdnstaskgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Gcpdnstaskgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Gcpdnstaskgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Gcpdnstaskgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Gcpdnstaskgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Gcpdnstaskgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Gcpdnstaskgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidateZones

`func (o *Gcpdnstaskgroup) GetConsolidateZones() bool`

GetConsolidateZones returns the ConsolidateZones field if non-nil, zero value otherwise.

### GetConsolidateZonesOk

`func (o *Gcpdnstaskgroup) GetConsolidateZonesOk() (*bool, bool)`

GetConsolidateZonesOk returns a tuple with the ConsolidateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateZones

`func (o *Gcpdnstaskgroup) SetConsolidateZones(v bool)`

SetConsolidateZones sets ConsolidateZones field to given value.

### HasConsolidateZones

`func (o *Gcpdnstaskgroup) HasConsolidateZones() bool`

HasConsolidateZones returns a boolean if a field has been set.

### GetConsolidatedView

`func (o *Gcpdnstaskgroup) GetConsolidatedView() string`

GetConsolidatedView returns the ConsolidatedView field if non-nil, zero value otherwise.

### GetConsolidatedViewOk

`func (o *Gcpdnstaskgroup) GetConsolidatedViewOk() (*string, bool)`

GetConsolidatedViewOk returns a tuple with the ConsolidatedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedView

`func (o *Gcpdnstaskgroup) SetConsolidatedView(v string)`

SetConsolidatedView sets ConsolidatedView field to given value.

### HasConsolidatedView

`func (o *Gcpdnstaskgroup) HasConsolidatedView() bool`

HasConsolidatedView returns a boolean if a field has been set.

### GetDisabled

`func (o *Gcpdnstaskgroup) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Gcpdnstaskgroup) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Gcpdnstaskgroup) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Gcpdnstaskgroup) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGcpProjectIdsFileToken

`func (o *Gcpdnstaskgroup) GetGcpProjectIdsFileToken() string`

GetGcpProjectIdsFileToken returns the GcpProjectIdsFileToken field if non-nil, zero value otherwise.

### GetGcpProjectIdsFileTokenOk

`func (o *Gcpdnstaskgroup) GetGcpProjectIdsFileTokenOk() (*string, bool)`

GetGcpProjectIdsFileTokenOk returns a tuple with the GcpProjectIdsFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectIdsFileToken

`func (o *Gcpdnstaskgroup) SetGcpProjectIdsFileToken(v string)`

SetGcpProjectIdsFileToken sets GcpProjectIdsFileToken field to given value.

### HasGcpProjectIdsFileToken

`func (o *Gcpdnstaskgroup) HasGcpProjectIdsFileToken() bool`

HasGcpProjectIdsFileToken returns a boolean if a field has been set.

### GetGridMember

`func (o *Gcpdnstaskgroup) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *Gcpdnstaskgroup) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *Gcpdnstaskgroup) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *Gcpdnstaskgroup) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetMultipleProjectsSyncPolicy

`func (o *Gcpdnstaskgroup) GetMultipleProjectsSyncPolicy() string`

GetMultipleProjectsSyncPolicy returns the MultipleProjectsSyncPolicy field if non-nil, zero value otherwise.

### GetMultipleProjectsSyncPolicyOk

`func (o *Gcpdnstaskgroup) GetMultipleProjectsSyncPolicyOk() (*string, bool)`

GetMultipleProjectsSyncPolicyOk returns a tuple with the MultipleProjectsSyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleProjectsSyncPolicy

`func (o *Gcpdnstaskgroup) SetMultipleProjectsSyncPolicy(v string)`

SetMultipleProjectsSyncPolicy sets MultipleProjectsSyncPolicy field to given value.

### HasMultipleProjectsSyncPolicy

`func (o *Gcpdnstaskgroup) HasMultipleProjectsSyncPolicy() bool`

HasMultipleProjectsSyncPolicy returns a boolean if a field has been set.

### GetName

`func (o *Gcpdnstaskgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gcpdnstaskgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gcpdnstaskgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gcpdnstaskgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *Gcpdnstaskgroup) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Gcpdnstaskgroup) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Gcpdnstaskgroup) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Gcpdnstaskgroup) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworkViewMappingPolicy

`func (o *Gcpdnstaskgroup) GetNetworkViewMappingPolicy() string`

GetNetworkViewMappingPolicy returns the NetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetNetworkViewMappingPolicyOk

`func (o *Gcpdnstaskgroup) GetNetworkViewMappingPolicyOk() (*string, bool)`

GetNetworkViewMappingPolicyOk returns a tuple with the NetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkViewMappingPolicy

`func (o *Gcpdnstaskgroup) SetNetworkViewMappingPolicy(v string)`

SetNetworkViewMappingPolicy sets NetworkViewMappingPolicy field to given value.

### HasNetworkViewMappingPolicy

`func (o *Gcpdnstaskgroup) HasNetworkViewMappingPolicy() bool`

HasNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetProjectId

`func (o *Gcpdnstaskgroup) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Gcpdnstaskgroup) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Gcpdnstaskgroup) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Gcpdnstaskgroup) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectsList

`func (o *Gcpdnstaskgroup) GetProjectsList() string`

GetProjectsList returns the ProjectsList field if non-nil, zero value otherwise.

### GetProjectsListOk

`func (o *Gcpdnstaskgroup) GetProjectsListOk() (*string, bool)`

GetProjectsListOk returns a tuple with the ProjectsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsList

`func (o *Gcpdnstaskgroup) SetProjectsList(v string)`

SetProjectsList sets ProjectsList field to given value.

### HasProjectsList

`func (o *Gcpdnstaskgroup) HasProjectsList() bool`

HasProjectsList returns a boolean if a field has been set.

### GetSyncChildProjects

`func (o *Gcpdnstaskgroup) GetSyncChildProjects() bool`

GetSyncChildProjects returns the SyncChildProjects field if non-nil, zero value otherwise.

### GetSyncChildProjectsOk

`func (o *Gcpdnstaskgroup) GetSyncChildProjectsOk() (*bool, bool)`

GetSyncChildProjectsOk returns a tuple with the SyncChildProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncChildProjects

`func (o *Gcpdnstaskgroup) SetSyncChildProjects(v bool)`

SetSyncChildProjects sets SyncChildProjects field to given value.

### HasSyncChildProjects

`func (o *Gcpdnstaskgroup) HasSyncChildProjects() bool`

HasSyncChildProjects returns a boolean if a field has been set.

### GetSyncStatus

`func (o *Gcpdnstaskgroup) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *Gcpdnstaskgroup) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *Gcpdnstaskgroup) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *Gcpdnstaskgroup) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetTaskList

`func (o *Gcpdnstaskgroup) GetTaskList() []GcpdnstaskgroupTaskList`

GetTaskList returns the TaskList field if non-nil, zero value otherwise.

### GetTaskListOk

`func (o *Gcpdnstaskgroup) GetTaskListOk() (*[]GcpdnstaskgroupTaskList, bool)`

GetTaskListOk returns a tuple with the TaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskList

`func (o *Gcpdnstaskgroup) SetTaskList(v []GcpdnstaskgroupTaskList)`

SetTaskList sets TaskList field to given value.

### HasTaskList

`func (o *Gcpdnstaskgroup) HasTaskList() bool`

HasTaskList returns a boolean if a field has been set.

### GetUuid

`func (o *Gcpdnstaskgroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Gcpdnstaskgroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Gcpdnstaskgroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Gcpdnstaskgroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


