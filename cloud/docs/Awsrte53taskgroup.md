# Awsrte53taskgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AccountId** | Pointer to **string** | The AWS Account ID associated with this task group. | [optional] [readonly] 
**AccountsList** | Pointer to **string** | The AWS Account IDs list associated with this task group. | [optional] [readonly] 
**AwsAccountIdsFileToken** | Pointer to **string** | The AWS account IDs file&#39;s token. | [optional] 
**Comment** | Pointer to **string** | Comment for the task group; maximum 256 characters. | [optional] 
**ConsolidateZones** | Pointer to **bool** | Indicates if all zones need to be saved into a single view. | [optional] 
**ConsolidatedView** | Pointer to **string** | The name of the DNS view for consolidating zones. | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the task group is enabled or disabled. | [optional] 
**GridMember** | Pointer to **string** | Member on which the tasks in this task group will be run. | [optional] 
**MultipleAccountsSyncPolicy** | Pointer to **string** | Discover all child accounts or Upload child account ids to discover.. | [optional] 
**Name** | Pointer to **string** | The name of this AWS Route53 sync task group. | [optional] 
**NetworkView** | Pointer to **string** | The name of the tenant&#39;s network view. | [optional] 
**NetworkViewMappingPolicy** | Pointer to **string** | The network view mapping policy. | [optional] 
**RoleArn** | Pointer to **string** | Role ARN for syncing child accounts; maximum 128 characters. | [optional] 
**SyncChildAccounts** | Pointer to **bool** | Synchronizing child accounts is enabled or disabled. | [optional] 
**SyncStatus** | Pointer to **string** | Indicate the overall sync status of this task group. | [optional] [readonly] 
**TaskList** | Pointer to [**[]Awsrte53taskgroupTaskList**](Awsrte53taskgroupTaskList.md) | List of AWS Route53 tasks in this group. | [optional] 

## Methods

### NewAwsrte53taskgroup

`func NewAwsrte53taskgroup() *Awsrte53taskgroup`

NewAwsrte53taskgroup instantiates a new Awsrte53taskgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsrte53taskgroupWithDefaults

`func NewAwsrte53taskgroupWithDefaults() *Awsrte53taskgroup`

NewAwsrte53taskgroupWithDefaults instantiates a new Awsrte53taskgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Awsrte53taskgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Awsrte53taskgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Awsrte53taskgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Awsrte53taskgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccountId

`func (o *Awsrte53taskgroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Awsrte53taskgroup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Awsrte53taskgroup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Awsrte53taskgroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountsList

`func (o *Awsrte53taskgroup) GetAccountsList() string`

GetAccountsList returns the AccountsList field if non-nil, zero value otherwise.

### GetAccountsListOk

`func (o *Awsrte53taskgroup) GetAccountsListOk() (*string, bool)`

GetAccountsListOk returns a tuple with the AccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsList

`func (o *Awsrte53taskgroup) SetAccountsList(v string)`

SetAccountsList sets AccountsList field to given value.

### HasAccountsList

`func (o *Awsrte53taskgroup) HasAccountsList() bool`

HasAccountsList returns a boolean if a field has been set.

### GetAwsAccountIdsFileToken

`func (o *Awsrte53taskgroup) GetAwsAccountIdsFileToken() string`

GetAwsAccountIdsFileToken returns the AwsAccountIdsFileToken field if non-nil, zero value otherwise.

### GetAwsAccountIdsFileTokenOk

`func (o *Awsrte53taskgroup) GetAwsAccountIdsFileTokenOk() (*string, bool)`

GetAwsAccountIdsFileTokenOk returns a tuple with the AwsAccountIdsFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountIdsFileToken

`func (o *Awsrte53taskgroup) SetAwsAccountIdsFileToken(v string)`

SetAwsAccountIdsFileToken sets AwsAccountIdsFileToken field to given value.

### HasAwsAccountIdsFileToken

`func (o *Awsrte53taskgroup) HasAwsAccountIdsFileToken() bool`

HasAwsAccountIdsFileToken returns a boolean if a field has been set.

### GetComment

`func (o *Awsrte53taskgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Awsrte53taskgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Awsrte53taskgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Awsrte53taskgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidateZones

`func (o *Awsrte53taskgroup) GetConsolidateZones() bool`

GetConsolidateZones returns the ConsolidateZones field if non-nil, zero value otherwise.

### GetConsolidateZonesOk

`func (o *Awsrte53taskgroup) GetConsolidateZonesOk() (*bool, bool)`

GetConsolidateZonesOk returns a tuple with the ConsolidateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateZones

`func (o *Awsrte53taskgroup) SetConsolidateZones(v bool)`

SetConsolidateZones sets ConsolidateZones field to given value.

### HasConsolidateZones

`func (o *Awsrte53taskgroup) HasConsolidateZones() bool`

HasConsolidateZones returns a boolean if a field has been set.

### GetConsolidatedView

`func (o *Awsrte53taskgroup) GetConsolidatedView() string`

GetConsolidatedView returns the ConsolidatedView field if non-nil, zero value otherwise.

### GetConsolidatedViewOk

`func (o *Awsrte53taskgroup) GetConsolidatedViewOk() (*string, bool)`

GetConsolidatedViewOk returns a tuple with the ConsolidatedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedView

`func (o *Awsrte53taskgroup) SetConsolidatedView(v string)`

SetConsolidatedView sets ConsolidatedView field to given value.

### HasConsolidatedView

`func (o *Awsrte53taskgroup) HasConsolidatedView() bool`

HasConsolidatedView returns a boolean if a field has been set.

### GetDisabled

`func (o *Awsrte53taskgroup) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Awsrte53taskgroup) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Awsrte53taskgroup) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Awsrte53taskgroup) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGridMember

`func (o *Awsrte53taskgroup) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *Awsrte53taskgroup) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *Awsrte53taskgroup) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *Awsrte53taskgroup) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetMultipleAccountsSyncPolicy

`func (o *Awsrte53taskgroup) GetMultipleAccountsSyncPolicy() string`

GetMultipleAccountsSyncPolicy returns the MultipleAccountsSyncPolicy field if non-nil, zero value otherwise.

### GetMultipleAccountsSyncPolicyOk

`func (o *Awsrte53taskgroup) GetMultipleAccountsSyncPolicyOk() (*string, bool)`

GetMultipleAccountsSyncPolicyOk returns a tuple with the MultipleAccountsSyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAccountsSyncPolicy

`func (o *Awsrte53taskgroup) SetMultipleAccountsSyncPolicy(v string)`

SetMultipleAccountsSyncPolicy sets MultipleAccountsSyncPolicy field to given value.

### HasMultipleAccountsSyncPolicy

`func (o *Awsrte53taskgroup) HasMultipleAccountsSyncPolicy() bool`

HasMultipleAccountsSyncPolicy returns a boolean if a field has been set.

### GetName

`func (o *Awsrte53taskgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Awsrte53taskgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Awsrte53taskgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Awsrte53taskgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *Awsrte53taskgroup) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Awsrte53taskgroup) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Awsrte53taskgroup) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Awsrte53taskgroup) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworkViewMappingPolicy

`func (o *Awsrte53taskgroup) GetNetworkViewMappingPolicy() string`

GetNetworkViewMappingPolicy returns the NetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetNetworkViewMappingPolicyOk

`func (o *Awsrte53taskgroup) GetNetworkViewMappingPolicyOk() (*string, bool)`

GetNetworkViewMappingPolicyOk returns a tuple with the NetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkViewMappingPolicy

`func (o *Awsrte53taskgroup) SetNetworkViewMappingPolicy(v string)`

SetNetworkViewMappingPolicy sets NetworkViewMappingPolicy field to given value.

### HasNetworkViewMappingPolicy

`func (o *Awsrte53taskgroup) HasNetworkViewMappingPolicy() bool`

HasNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetRoleArn

`func (o *Awsrte53taskgroup) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *Awsrte53taskgroup) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *Awsrte53taskgroup) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *Awsrte53taskgroup) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSyncChildAccounts

`func (o *Awsrte53taskgroup) GetSyncChildAccounts() bool`

GetSyncChildAccounts returns the SyncChildAccounts field if non-nil, zero value otherwise.

### GetSyncChildAccountsOk

`func (o *Awsrte53taskgroup) GetSyncChildAccountsOk() (*bool, bool)`

GetSyncChildAccountsOk returns a tuple with the SyncChildAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncChildAccounts

`func (o *Awsrte53taskgroup) SetSyncChildAccounts(v bool)`

SetSyncChildAccounts sets SyncChildAccounts field to given value.

### HasSyncChildAccounts

`func (o *Awsrte53taskgroup) HasSyncChildAccounts() bool`

HasSyncChildAccounts returns a boolean if a field has been set.

### GetSyncStatus

`func (o *Awsrte53taskgroup) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *Awsrte53taskgroup) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *Awsrte53taskgroup) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *Awsrte53taskgroup) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetTaskList

`func (o *Awsrte53taskgroup) GetTaskList() []Awsrte53taskgroupTaskList`

GetTaskList returns the TaskList field if non-nil, zero value otherwise.

### GetTaskListOk

`func (o *Awsrte53taskgroup) GetTaskListOk() (*[]Awsrte53taskgroupTaskList, bool)`

GetTaskListOk returns a tuple with the TaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskList

`func (o *Awsrte53taskgroup) SetTaskList(v []Awsrte53taskgroupTaskList)`

SetTaskList sets TaskList field to given value.

### HasTaskList

`func (o *Awsrte53taskgroup) HasTaskList() bool`

HasTaskList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


