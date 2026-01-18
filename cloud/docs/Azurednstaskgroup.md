# Azurednstaskgroup

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

## Methods

### NewAzurednstaskgroup

`func NewAzurednstaskgroup() *Azurednstaskgroup`

NewAzurednstaskgroup instantiates a new Azurednstaskgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurednstaskgroupWithDefaults

`func NewAzurednstaskgroupWithDefaults() *Azurednstaskgroup`

NewAzurednstaskgroupWithDefaults instantiates a new Azurednstaskgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Azurednstaskgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Azurednstaskgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Azurednstaskgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Azurednstaskgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAzureSubscriptionIdsFileToken

`func (o *Azurednstaskgroup) GetAzureSubscriptionIdsFileToken() string`

GetAzureSubscriptionIdsFileToken returns the AzureSubscriptionIdsFileToken field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdsFileTokenOk

`func (o *Azurednstaskgroup) GetAzureSubscriptionIdsFileTokenOk() (*string, bool)`

GetAzureSubscriptionIdsFileTokenOk returns a tuple with the AzureSubscriptionIdsFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionIdsFileToken

`func (o *Azurednstaskgroup) SetAzureSubscriptionIdsFileToken(v string)`

SetAzureSubscriptionIdsFileToken sets AzureSubscriptionIdsFileToken field to given value.

### HasAzureSubscriptionIdsFileToken

`func (o *Azurednstaskgroup) HasAzureSubscriptionIdsFileToken() bool`

HasAzureSubscriptionIdsFileToken returns a boolean if a field has been set.

### GetComment

`func (o *Azurednstaskgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Azurednstaskgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Azurednstaskgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Azurednstaskgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidateZones

`func (o *Azurednstaskgroup) GetConsolidateZones() bool`

GetConsolidateZones returns the ConsolidateZones field if non-nil, zero value otherwise.

### GetConsolidateZonesOk

`func (o *Azurednstaskgroup) GetConsolidateZonesOk() (*bool, bool)`

GetConsolidateZonesOk returns a tuple with the ConsolidateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateZones

`func (o *Azurednstaskgroup) SetConsolidateZones(v bool)`

SetConsolidateZones sets ConsolidateZones field to given value.

### HasConsolidateZones

`func (o *Azurednstaskgroup) HasConsolidateZones() bool`

HasConsolidateZones returns a boolean if a field has been set.

### GetConsolidatedView

`func (o *Azurednstaskgroup) GetConsolidatedView() string`

GetConsolidatedView returns the ConsolidatedView field if non-nil, zero value otherwise.

### GetConsolidatedViewOk

`func (o *Azurednstaskgroup) GetConsolidatedViewOk() (*string, bool)`

GetConsolidatedViewOk returns a tuple with the ConsolidatedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedView

`func (o *Azurednstaskgroup) SetConsolidatedView(v string)`

SetConsolidatedView sets ConsolidatedView field to given value.

### HasConsolidatedView

`func (o *Azurednstaskgroup) HasConsolidatedView() bool`

HasConsolidatedView returns a boolean if a field has been set.

### GetDisabled

`func (o *Azurednstaskgroup) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Azurednstaskgroup) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Azurednstaskgroup) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Azurednstaskgroup) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGridMember

`func (o *Azurednstaskgroup) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *Azurednstaskgroup) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *Azurednstaskgroup) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *Azurednstaskgroup) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetMultipleSubscriptionsSyncPolicy

`func (o *Azurednstaskgroup) GetMultipleSubscriptionsSyncPolicy() string`

GetMultipleSubscriptionsSyncPolicy returns the MultipleSubscriptionsSyncPolicy field if non-nil, zero value otherwise.

### GetMultipleSubscriptionsSyncPolicyOk

`func (o *Azurednstaskgroup) GetMultipleSubscriptionsSyncPolicyOk() (*string, bool)`

GetMultipleSubscriptionsSyncPolicyOk returns a tuple with the MultipleSubscriptionsSyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSubscriptionsSyncPolicy

`func (o *Azurednstaskgroup) SetMultipleSubscriptionsSyncPolicy(v string)`

SetMultipleSubscriptionsSyncPolicy sets MultipleSubscriptionsSyncPolicy field to given value.

### HasMultipleSubscriptionsSyncPolicy

`func (o *Azurednstaskgroup) HasMultipleSubscriptionsSyncPolicy() bool`

HasMultipleSubscriptionsSyncPolicy returns a boolean if a field has been set.

### GetName

`func (o *Azurednstaskgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Azurednstaskgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Azurednstaskgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Azurednstaskgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *Azurednstaskgroup) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Azurednstaskgroup) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Azurednstaskgroup) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Azurednstaskgroup) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworkViewMappingPolicy

`func (o *Azurednstaskgroup) GetNetworkViewMappingPolicy() string`

GetNetworkViewMappingPolicy returns the NetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetNetworkViewMappingPolicyOk

`func (o *Azurednstaskgroup) GetNetworkViewMappingPolicyOk() (*string, bool)`

GetNetworkViewMappingPolicyOk returns a tuple with the NetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkViewMappingPolicy

`func (o *Azurednstaskgroup) SetNetworkViewMappingPolicy(v string)`

SetNetworkViewMappingPolicy sets NetworkViewMappingPolicy field to given value.

### HasNetworkViewMappingPolicy

`func (o *Azurednstaskgroup) HasNetworkViewMappingPolicy() bool`

HasNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetSubscriptionsList

`func (o *Azurednstaskgroup) GetSubscriptionsList() string`

GetSubscriptionsList returns the SubscriptionsList field if non-nil, zero value otherwise.

### GetSubscriptionsListOk

`func (o *Azurednstaskgroup) GetSubscriptionsListOk() (*string, bool)`

GetSubscriptionsListOk returns a tuple with the SubscriptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsList

`func (o *Azurednstaskgroup) SetSubscriptionsList(v string)`

SetSubscriptionsList sets SubscriptionsList field to given value.

### HasSubscriptionsList

`func (o *Azurednstaskgroup) HasSubscriptionsList() bool`

HasSubscriptionsList returns a boolean if a field has been set.

### GetSyncStatus

`func (o *Azurednstaskgroup) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *Azurednstaskgroup) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *Azurednstaskgroup) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *Azurednstaskgroup) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetTaskList

`func (o *Azurednstaskgroup) GetTaskList() []AzurednstaskgroupTaskList`

GetTaskList returns the TaskList field if non-nil, zero value otherwise.

### GetTaskListOk

`func (o *Azurednstaskgroup) GetTaskListOk() (*[]AzurednstaskgroupTaskList, bool)`

GetTaskListOk returns a tuple with the TaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskList

`func (o *Azurednstaskgroup) SetTaskList(v []AzurednstaskgroupTaskList)`

SetTaskList sets TaskList field to given value.

### HasTaskList

`func (o *Azurednstaskgroup) HasTaskList() bool`

HasTaskList returns a boolean if a field has been set.

### GetTenantId

`func (o *Azurednstaskgroup) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Azurednstaskgroup) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Azurednstaskgroup) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Azurednstaskgroup) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


