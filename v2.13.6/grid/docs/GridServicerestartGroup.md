# GridServicerestartGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Restart Group; maximum 256 characters. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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

## Methods

### NewGridServicerestartGroup

`func NewGridServicerestartGroup() *GridServicerestartGroup`

NewGridServicerestartGroup instantiates a new GridServicerestartGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridServicerestartGroupWithDefaults

`func NewGridServicerestartGroupWithDefaults() *GridServicerestartGroup`

NewGridServicerestartGroupWithDefaults instantiates a new GridServicerestartGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridServicerestartGroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridServicerestartGroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridServicerestartGroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridServicerestartGroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GridServicerestartGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GridServicerestartGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GridServicerestartGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GridServicerestartGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GridServicerestartGroup) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GridServicerestartGroup) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GridServicerestartGroup) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GridServicerestartGroup) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIsDefault

`func (o *GridServicerestartGroup) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GridServicerestartGroup) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GridServicerestartGroup) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GridServicerestartGroup) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *GridServicerestartGroup) GetLastUpdatedTime() int64`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GridServicerestartGroup) GetLastUpdatedTimeOk() (*int64, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GridServicerestartGroup) SetLastUpdatedTime(v int64)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *GridServicerestartGroup) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMembers

`func (o *GridServicerestartGroup) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GridServicerestartGroup) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GridServicerestartGroup) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GridServicerestartGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMode

`func (o *GridServicerestartGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GridServicerestartGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GridServicerestartGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GridServicerestartGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GridServicerestartGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridServicerestartGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridServicerestartGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridServicerestartGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *GridServicerestartGroup) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GridServicerestartGroup) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GridServicerestartGroup) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GridServicerestartGroup) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRecurringSchedule

`func (o *GridServicerestartGroup) GetRecurringSchedule() GridServicerestartGroupRecurringSchedule`

GetRecurringSchedule returns the RecurringSchedule field if non-nil, zero value otherwise.

### GetRecurringScheduleOk

`func (o *GridServicerestartGroup) GetRecurringScheduleOk() (*GridServicerestartGroupRecurringSchedule, bool)`

GetRecurringScheduleOk returns a tuple with the RecurringSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSchedule

`func (o *GridServicerestartGroup) SetRecurringSchedule(v GridServicerestartGroupRecurringSchedule)`

SetRecurringSchedule sets RecurringSchedule field to given value.

### HasRecurringSchedule

`func (o *GridServicerestartGroup) HasRecurringSchedule() bool`

HasRecurringSchedule returns a boolean if a field has been set.

### GetRequests

`func (o *GridServicerestartGroup) GetRequests() []string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GridServicerestartGroup) GetRequestsOk() (*[]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GridServicerestartGroup) SetRequests(v []string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *GridServicerestartGroup) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetService

`func (o *GridServicerestartGroup) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GridServicerestartGroup) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GridServicerestartGroup) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GridServicerestartGroup) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *GridServicerestartGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GridServicerestartGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GridServicerestartGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GridServicerestartGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


