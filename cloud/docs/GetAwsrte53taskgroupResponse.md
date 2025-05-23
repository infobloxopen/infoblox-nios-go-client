# GetAwsrte53taskgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**TaskControl** | Pointer to **map[string]interface{}** |  | [optional] 
**TaskList** | Pointer to [**[]Awsrte53taskgroupTaskList**](Awsrte53taskgroupTaskList.md) | List of AWS Route53 tasks in this group. | [optional] 
**Result** | Pointer to [**Awsrte53taskgroup**](Awsrte53taskgroup.md) |  | [optional] 

## Methods

### NewGetAwsrte53taskgroupResponse

`func NewGetAwsrte53taskgroupResponse() *GetAwsrte53taskgroupResponse`

NewGetAwsrte53taskgroupResponse instantiates a new GetAwsrte53taskgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAwsrte53taskgroupResponseWithDefaults

`func NewGetAwsrte53taskgroupResponseWithDefaults() *GetAwsrte53taskgroupResponse`

NewGetAwsrte53taskgroupResponseWithDefaults instantiates a new GetAwsrte53taskgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAwsrte53taskgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAwsrte53taskgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAwsrte53taskgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAwsrte53taskgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccountId

`func (o *GetAwsrte53taskgroupResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAwsrte53taskgroupResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAwsrte53taskgroupResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetAwsrte53taskgroupResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountsList

`func (o *GetAwsrte53taskgroupResponse) GetAccountsList() string`

GetAccountsList returns the AccountsList field if non-nil, zero value otherwise.

### GetAccountsListOk

`func (o *GetAwsrte53taskgroupResponse) GetAccountsListOk() (*string, bool)`

GetAccountsListOk returns a tuple with the AccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsList

`func (o *GetAwsrte53taskgroupResponse) SetAccountsList(v string)`

SetAccountsList sets AccountsList field to given value.

### HasAccountsList

`func (o *GetAwsrte53taskgroupResponse) HasAccountsList() bool`

HasAccountsList returns a boolean if a field has been set.

### GetAwsAccountIdsFileToken

`func (o *GetAwsrte53taskgroupResponse) GetAwsAccountIdsFileToken() string`

GetAwsAccountIdsFileToken returns the AwsAccountIdsFileToken field if non-nil, zero value otherwise.

### GetAwsAccountIdsFileTokenOk

`func (o *GetAwsrte53taskgroupResponse) GetAwsAccountIdsFileTokenOk() (*string, bool)`

GetAwsAccountIdsFileTokenOk returns a tuple with the AwsAccountIdsFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountIdsFileToken

`func (o *GetAwsrte53taskgroupResponse) SetAwsAccountIdsFileToken(v string)`

SetAwsAccountIdsFileToken sets AwsAccountIdsFileToken field to given value.

### HasAwsAccountIdsFileToken

`func (o *GetAwsrte53taskgroupResponse) HasAwsAccountIdsFileToken() bool`

HasAwsAccountIdsFileToken returns a boolean if a field has been set.

### GetComment

`func (o *GetAwsrte53taskgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAwsrte53taskgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAwsrte53taskgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAwsrte53taskgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidateZones

`func (o *GetAwsrte53taskgroupResponse) GetConsolidateZones() bool`

GetConsolidateZones returns the ConsolidateZones field if non-nil, zero value otherwise.

### GetConsolidateZonesOk

`func (o *GetAwsrte53taskgroupResponse) GetConsolidateZonesOk() (*bool, bool)`

GetConsolidateZonesOk returns a tuple with the ConsolidateZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateZones

`func (o *GetAwsrte53taskgroupResponse) SetConsolidateZones(v bool)`

SetConsolidateZones sets ConsolidateZones field to given value.

### HasConsolidateZones

`func (o *GetAwsrte53taskgroupResponse) HasConsolidateZones() bool`

HasConsolidateZones returns a boolean if a field has been set.

### GetConsolidatedView

`func (o *GetAwsrte53taskgroupResponse) GetConsolidatedView() string`

GetConsolidatedView returns the ConsolidatedView field if non-nil, zero value otherwise.

### GetConsolidatedViewOk

`func (o *GetAwsrte53taskgroupResponse) GetConsolidatedViewOk() (*string, bool)`

GetConsolidatedViewOk returns a tuple with the ConsolidatedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedView

`func (o *GetAwsrte53taskgroupResponse) SetConsolidatedView(v string)`

SetConsolidatedView sets ConsolidatedView field to given value.

### HasConsolidatedView

`func (o *GetAwsrte53taskgroupResponse) HasConsolidatedView() bool`

HasConsolidatedView returns a boolean if a field has been set.

### GetDisabled

`func (o *GetAwsrte53taskgroupResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetAwsrte53taskgroupResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetAwsrte53taskgroupResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetAwsrte53taskgroupResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGridMember

`func (o *GetAwsrte53taskgroupResponse) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *GetAwsrte53taskgroupResponse) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *GetAwsrte53taskgroupResponse) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *GetAwsrte53taskgroupResponse) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetMultipleAccountsSyncPolicy

`func (o *GetAwsrte53taskgroupResponse) GetMultipleAccountsSyncPolicy() string`

GetMultipleAccountsSyncPolicy returns the MultipleAccountsSyncPolicy field if non-nil, zero value otherwise.

### GetMultipleAccountsSyncPolicyOk

`func (o *GetAwsrte53taskgroupResponse) GetMultipleAccountsSyncPolicyOk() (*string, bool)`

GetMultipleAccountsSyncPolicyOk returns a tuple with the MultipleAccountsSyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAccountsSyncPolicy

`func (o *GetAwsrte53taskgroupResponse) SetMultipleAccountsSyncPolicy(v string)`

SetMultipleAccountsSyncPolicy sets MultipleAccountsSyncPolicy field to given value.

### HasMultipleAccountsSyncPolicy

`func (o *GetAwsrte53taskgroupResponse) HasMultipleAccountsSyncPolicy() bool`

HasMultipleAccountsSyncPolicy returns a boolean if a field has been set.

### GetName

`func (o *GetAwsrte53taskgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAwsrte53taskgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAwsrte53taskgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAwsrte53taskgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetAwsrte53taskgroupResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetAwsrte53taskgroupResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetAwsrte53taskgroupResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetAwsrte53taskgroupResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworkViewMappingPolicy

`func (o *GetAwsrte53taskgroupResponse) GetNetworkViewMappingPolicy() string`

GetNetworkViewMappingPolicy returns the NetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetNetworkViewMappingPolicyOk

`func (o *GetAwsrte53taskgroupResponse) GetNetworkViewMappingPolicyOk() (*string, bool)`

GetNetworkViewMappingPolicyOk returns a tuple with the NetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkViewMappingPolicy

`func (o *GetAwsrte53taskgroupResponse) SetNetworkViewMappingPolicy(v string)`

SetNetworkViewMappingPolicy sets NetworkViewMappingPolicy field to given value.

### HasNetworkViewMappingPolicy

`func (o *GetAwsrte53taskgroupResponse) HasNetworkViewMappingPolicy() bool`

HasNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetRoleArn

`func (o *GetAwsrte53taskgroupResponse) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *GetAwsrte53taskgroupResponse) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *GetAwsrte53taskgroupResponse) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *GetAwsrte53taskgroupResponse) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSyncChildAccounts

`func (o *GetAwsrte53taskgroupResponse) GetSyncChildAccounts() bool`

GetSyncChildAccounts returns the SyncChildAccounts field if non-nil, zero value otherwise.

### GetSyncChildAccountsOk

`func (o *GetAwsrte53taskgroupResponse) GetSyncChildAccountsOk() (*bool, bool)`

GetSyncChildAccountsOk returns a tuple with the SyncChildAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncChildAccounts

`func (o *GetAwsrte53taskgroupResponse) SetSyncChildAccounts(v bool)`

SetSyncChildAccounts sets SyncChildAccounts field to given value.

### HasSyncChildAccounts

`func (o *GetAwsrte53taskgroupResponse) HasSyncChildAccounts() bool`

HasSyncChildAccounts returns a boolean if a field has been set.

### GetSyncStatus

`func (o *GetAwsrte53taskgroupResponse) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *GetAwsrte53taskgroupResponse) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *GetAwsrte53taskgroupResponse) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *GetAwsrte53taskgroupResponse) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetTaskControl

`func (o *GetAwsrte53taskgroupResponse) GetTaskControl() map[string]interface{}`

GetTaskControl returns the TaskControl field if non-nil, zero value otherwise.

### GetTaskControlOk

`func (o *GetAwsrte53taskgroupResponse) GetTaskControlOk() (*map[string]interface{}, bool)`

GetTaskControlOk returns a tuple with the TaskControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskControl

`func (o *GetAwsrte53taskgroupResponse) SetTaskControl(v map[string]interface{})`

SetTaskControl sets TaskControl field to given value.

### HasTaskControl

`func (o *GetAwsrte53taskgroupResponse) HasTaskControl() bool`

HasTaskControl returns a boolean if a field has been set.

### GetTaskList

`func (o *GetAwsrte53taskgroupResponse) GetTaskList() []Awsrte53taskgroupTaskList`

GetTaskList returns the TaskList field if non-nil, zero value otherwise.

### GetTaskListOk

`func (o *GetAwsrte53taskgroupResponse) GetTaskListOk() (*[]Awsrte53taskgroupTaskList, bool)`

GetTaskListOk returns a tuple with the TaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskList

`func (o *GetAwsrte53taskgroupResponse) SetTaskList(v []Awsrte53taskgroupTaskList)`

SetTaskList sets TaskList field to given value.

### HasTaskList

`func (o *GetAwsrte53taskgroupResponse) HasTaskList() bool`

HasTaskList returns a boolean if a field has been set.

### GetResult

`func (o *GetAwsrte53taskgroupResponse) GetResult() Awsrte53taskgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAwsrte53taskgroupResponse) GetResultOk() (*Awsrte53taskgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAwsrte53taskgroupResponse) SetResult(v Awsrte53taskgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAwsrte53taskgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


