# GetAzurednstaskgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AzureSubscriptionIdsFileToken** | Pointer to **string** | The Azure Subscription IDs file&#39;s token. | [optional] 
**Comment** | Pointer to **string** | Comment for the task group; maximum 256 characters. | [optional] 
**ConsolidateZones** | Pointer to **bool** | Indicates if all zones need to be saved into a single view. | [optional] 
**ConsolidatedView** | Pointer to **string** | The name of the DNS view for consolidating zones. | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the task group is enabled or disabled. | [optional] 
**GridMember** | Pointer to **string** | Member on which the tasks in this task group will be run. | [optional] 
**MultipleSubscriptionsSyncPolicy** | Pointer to **string** | Discover all child subscriptions or Upload child subscription ids to discover. | [optional] 
**Name** | Pointer to **string** | The name of this Azure DNS sync task group. | [optional] 
**NetworkView** | Pointer to **string** | The name of the tenant&#39;s network view. | [optional] 
**NetworkViewMappingPolicy** | Pointer to **string** | The network view mapping policy. | [optional] 
**SubscriptionsList** | Pointer to **string** | The Azure Subscription IDs list associated with this task. | [optional] [readonly] 
**SyncStatus** | Pointer to **string** | Indicate the overall sync status of this task group. | [optional] [readonly] 
**TaskList** | Pointer to [**[]AzurednstaskgroupTaskList**](AzurednstaskgroupTaskList.md) | List of Azure DNS tasks in this group. | [optional] 
**TenantId** | Pointer to **string** | The Azure tenant ID associated with this task group. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Azurednstaskgroup**](Azurednstaskgroup.md) |  | [optional] 

## Methods

### NewGetAzurednstaskgroupResponse

`func NewGetAzurednstaskgroupResponse() *GetAzurednstaskgroupResponse`

NewGetAzurednstaskgroupResponse instantiates a new GetAzurednstaskgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAzurednstaskgroupResponseWithDefaults

`func NewGetAzurednstaskgroupResponseWithDefaults() *GetAzurednstaskgroupResponse`

NewGetAzurednstaskgroupResponseWithDefaults instantiates a new GetAzurednstaskgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAzurednstaskgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAzurednstaskgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAzurednstaskgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAzurednstaskgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAzureSubscriptionIdsFileToken

`func (o *GetAzurednstaskgroupResponse) GetAzureSubscriptionIdsFileToken() string`

GetAzureSubscriptionIdsFileToken returns the AzureSubscriptionIdsFileToken field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdsFileTokenOk

`func (o *GetAzurednstaskgroupResponse) GetAzureSubscriptionIdsFileTokenOk() (*string, bool)`

GetAzureSubscriptionIdsFileTokenOk returns a tuple with the AzureSubscriptionIdsFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionIdsFileToken

`func (o *GetAzurednstaskgroupResponse) SetAzureSubscriptionIdsFileToken(v string)`

SetAzureSubscriptionIdsFileToken sets AzureSubscriptionIdsFileToken field to given value.

### HasAzureSubscriptionIdsFileToken

`func (o *GetAzurednstaskgroupResponse) HasAzureSubscriptionIdsFileToken() bool`

HasAzureSubscriptionIdsFileToken returns a boolean if a field has been set.

### GetComment

`func (o *GetAzurednstaskgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAzurednstaskgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAzurednstaskgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAzurednstaskgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidateZones

`func (o *GetAzurednstaskgroupResponse) GetConsolidateZones() bool`

GetConsolidateZones returns the ConsolidateZones field if non-nil, zero value otherwise.

### GetConsolidateZonesOk

`func (o *GetAzurednstaskgroupResponse) GetConsolidateZonesOk() (*bool, bool)`

GetConsolidateZonesOk returns a tuple with the ConsolidateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateZones

`func (o *GetAzurednstaskgroupResponse) SetConsolidateZones(v bool)`

SetConsolidateZones sets ConsolidateZones field to given value.

### HasConsolidateZones

`func (o *GetAzurednstaskgroupResponse) HasConsolidateZones() bool`

HasConsolidateZones returns a boolean if a field has been set.

### GetConsolidatedView

`func (o *GetAzurednstaskgroupResponse) GetConsolidatedView() string`

GetConsolidatedView returns the ConsolidatedView field if non-nil, zero value otherwise.

### GetConsolidatedViewOk

`func (o *GetAzurednstaskgroupResponse) GetConsolidatedViewOk() (*string, bool)`

GetConsolidatedViewOk returns a tuple with the ConsolidatedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedView

`func (o *GetAzurednstaskgroupResponse) SetConsolidatedView(v string)`

SetConsolidatedView sets ConsolidatedView field to given value.

### HasConsolidatedView

`func (o *GetAzurednstaskgroupResponse) HasConsolidatedView() bool`

HasConsolidatedView returns a boolean if a field has been set.

### GetDisabled

`func (o *GetAzurednstaskgroupResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetAzurednstaskgroupResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetAzurednstaskgroupResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetAzurednstaskgroupResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGridMember

`func (o *GetAzurednstaskgroupResponse) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *GetAzurednstaskgroupResponse) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *GetAzurednstaskgroupResponse) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *GetAzurednstaskgroupResponse) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetMultipleSubscriptionsSyncPolicy

`func (o *GetAzurednstaskgroupResponse) GetMultipleSubscriptionsSyncPolicy() string`

GetMultipleSubscriptionsSyncPolicy returns the MultipleSubscriptionsSyncPolicy field if non-nil, zero value otherwise.

### GetMultipleSubscriptionsSyncPolicyOk

`func (o *GetAzurednstaskgroupResponse) GetMultipleSubscriptionsSyncPolicyOk() (*string, bool)`

GetMultipleSubscriptionsSyncPolicyOk returns a tuple with the MultipleSubscriptionsSyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSubscriptionsSyncPolicy

`func (o *GetAzurednstaskgroupResponse) SetMultipleSubscriptionsSyncPolicy(v string)`

SetMultipleSubscriptionsSyncPolicy sets MultipleSubscriptionsSyncPolicy field to given value.

### HasMultipleSubscriptionsSyncPolicy

`func (o *GetAzurednstaskgroupResponse) HasMultipleSubscriptionsSyncPolicy() bool`

HasMultipleSubscriptionsSyncPolicy returns a boolean if a field has been set.

### GetName

`func (o *GetAzurednstaskgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAzurednstaskgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAzurednstaskgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAzurednstaskgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetAzurednstaskgroupResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetAzurednstaskgroupResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetAzurednstaskgroupResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetAzurednstaskgroupResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworkViewMappingPolicy

`func (o *GetAzurednstaskgroupResponse) GetNetworkViewMappingPolicy() string`

GetNetworkViewMappingPolicy returns the NetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetNetworkViewMappingPolicyOk

`func (o *GetAzurednstaskgroupResponse) GetNetworkViewMappingPolicyOk() (*string, bool)`

GetNetworkViewMappingPolicyOk returns a tuple with the NetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkViewMappingPolicy

`func (o *GetAzurednstaskgroupResponse) SetNetworkViewMappingPolicy(v string)`

SetNetworkViewMappingPolicy sets NetworkViewMappingPolicy field to given value.

### HasNetworkViewMappingPolicy

`func (o *GetAzurednstaskgroupResponse) HasNetworkViewMappingPolicy() bool`

HasNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetSubscriptionsList

`func (o *GetAzurednstaskgroupResponse) GetSubscriptionsList() string`

GetSubscriptionsList returns the SubscriptionsList field if non-nil, zero value otherwise.

### GetSubscriptionsListOk

`func (o *GetAzurednstaskgroupResponse) GetSubscriptionsListOk() (*string, bool)`

GetSubscriptionsListOk returns a tuple with the SubscriptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsList

`func (o *GetAzurednstaskgroupResponse) SetSubscriptionsList(v string)`

SetSubscriptionsList sets SubscriptionsList field to given value.

### HasSubscriptionsList

`func (o *GetAzurednstaskgroupResponse) HasSubscriptionsList() bool`

HasSubscriptionsList returns a boolean if a field has been set.

### GetSyncStatus

`func (o *GetAzurednstaskgroupResponse) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *GetAzurednstaskgroupResponse) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *GetAzurednstaskgroupResponse) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *GetAzurednstaskgroupResponse) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetTaskList

`func (o *GetAzurednstaskgroupResponse) GetTaskList() []AzurednstaskgroupTaskList`

GetTaskList returns the TaskList field if non-nil, zero value otherwise.

### GetTaskListOk

`func (o *GetAzurednstaskgroupResponse) GetTaskListOk() (*[]AzurednstaskgroupTaskList, bool)`

GetTaskListOk returns a tuple with the TaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskList

`func (o *GetAzurednstaskgroupResponse) SetTaskList(v []AzurednstaskgroupTaskList)`

SetTaskList sets TaskList field to given value.

### HasTaskList

`func (o *GetAzurednstaskgroupResponse) HasTaskList() bool`

HasTaskList returns a boolean if a field has been set.

### GetTenantId

`func (o *GetAzurednstaskgroupResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetAzurednstaskgroupResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetAzurednstaskgroupResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *GetAzurednstaskgroupResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUuid

`func (o *GetAzurednstaskgroupResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetAzurednstaskgroupResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetAzurednstaskgroupResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetAzurednstaskgroupResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetAzurednstaskgroupResponse) GetResult() Azurednstaskgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAzurednstaskgroupResponse) GetResultOk() (*Azurednstaskgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAzurednstaskgroupResponse) SetResult(v Azurednstaskgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAzurednstaskgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


