# GetGridServicerestartGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Restart Group; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IsDefault** | Pointer to **bool** | Determines if this Restart Group is the default group. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **int64** | The timestamp when the status of the latest request has changed. | [optional] [readonly] 
**Members** | Pointer to **[]string** | The list of members belonging to the group. | [optional] 
**Mode** | Pointer to **string** | The default restart method for this Restart Group. | [optional] 
**Name** | Pointer to **string** | The name of this Restart Group. | [optional] 
**Position** | Pointer to **int64** | The order to restart. | [optional] [readonly] 
**RecurringSchedule** | Pointer to [**GridServicerestartGroupRecurringSchedule**](GridServicerestartGroupRecurringSchedule.md) |  | [optional] 
**Requests** | Pointer to **[]string** | The list of requests associated with a restart group. | [optional] [readonly] 
**Service** | Pointer to **string** | The applicable service for this Restart Group. | [optional] 
**Status** | Pointer to **string** | The restart status for a restart group. | [optional] [readonly] 
**Result** | Pointer to [**GridServicerestartGroup**](GridServicerestartGroup.md) |  | [optional] 

## Methods

### NewGetGridServicerestartGroupResponse

`func NewGetGridServicerestartGroupResponse() *GetGridServicerestartGroupResponse`

NewGetGridServicerestartGroupResponse instantiates a new GetGridServicerestartGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridServicerestartGroupResponseWithDefaults

`func NewGetGridServicerestartGroupResponseWithDefaults() *GetGridServicerestartGroupResponse`

NewGetGridServicerestartGroupResponseWithDefaults instantiates a new GetGridServicerestartGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridServicerestartGroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridServicerestartGroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridServicerestartGroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridServicerestartGroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetGridServicerestartGroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetGridServicerestartGroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetGridServicerestartGroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetGridServicerestartGroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetGridServicerestartGroupResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetGridServicerestartGroupResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetGridServicerestartGroupResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetGridServicerestartGroupResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetGridServicerestartGroupResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetGridServicerestartGroupResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetGridServicerestartGroupResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetGridServicerestartGroupResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *GetGridServicerestartGroupResponse) GetLastUpdatedTime() int64`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GetGridServicerestartGroupResponse) GetLastUpdatedTimeOk() (*int64, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GetGridServicerestartGroupResponse) SetLastUpdatedTime(v int64)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *GetGridServicerestartGroupResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMembers

`func (o *GetGridServicerestartGroupResponse) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetGridServicerestartGroupResponse) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetGridServicerestartGroupResponse) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetGridServicerestartGroupResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMode

`func (o *GetGridServicerestartGroupResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetGridServicerestartGroupResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetGridServicerestartGroupResponse) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetGridServicerestartGroupResponse) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GetGridServicerestartGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGridServicerestartGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGridServicerestartGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGridServicerestartGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *GetGridServicerestartGroupResponse) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetGridServicerestartGroupResponse) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetGridServicerestartGroupResponse) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetGridServicerestartGroupResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRecurringSchedule

`func (o *GetGridServicerestartGroupResponse) GetRecurringSchedule() GridServicerestartGroupRecurringSchedule`

GetRecurringSchedule returns the RecurringSchedule field if non-nil, zero value otherwise.

### GetRecurringScheduleOk

`func (o *GetGridServicerestartGroupResponse) GetRecurringScheduleOk() (*GridServicerestartGroupRecurringSchedule, bool)`

GetRecurringScheduleOk returns a tuple with the RecurringSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSchedule

`func (o *GetGridServicerestartGroupResponse) SetRecurringSchedule(v GridServicerestartGroupRecurringSchedule)`

SetRecurringSchedule sets RecurringSchedule field to given value.

### HasRecurringSchedule

`func (o *GetGridServicerestartGroupResponse) HasRecurringSchedule() bool`

HasRecurringSchedule returns a boolean if a field has been set.

### GetRequests

`func (o *GetGridServicerestartGroupResponse) GetRequests() []string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GetGridServicerestartGroupResponse) GetRequestsOk() (*[]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GetGridServicerestartGroupResponse) SetRequests(v []string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *GetGridServicerestartGroupResponse) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetService

`func (o *GetGridServicerestartGroupResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetGridServicerestartGroupResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetGridServicerestartGroupResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetGridServicerestartGroupResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *GetGridServicerestartGroupResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGridServicerestartGroupResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGridServicerestartGroupResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGridServicerestartGroupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetGridServicerestartGroupResponse) GetResult() GridServicerestartGroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridServicerestartGroupResponse) GetResultOk() (*GridServicerestartGroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridServicerestartGroupResponse) SetResult(v GridServicerestartGroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridServicerestartGroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


