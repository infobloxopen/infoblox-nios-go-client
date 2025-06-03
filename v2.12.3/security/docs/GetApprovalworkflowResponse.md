# GetApprovalworkflowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ApprovalGroup** | Pointer to **string** | The approval administration group. | [optional] 
**ApprovalNotifyTo** | Pointer to **string** | The destination for approval task notifications. | [optional] 
**ApprovedNotifyTo** | Pointer to **string** | The destination for approved task notifications. | [optional] 
**ApproverComment** | Pointer to **string** | The requirement for the comment when an approver approves a submitted task. | [optional] 
**EnableApprovalNotify** | Pointer to **bool** | Determines whether approval task notifications are enabled. | [optional] 
**EnableApprovedNotify** | Pointer to **bool** | Determines whether approved task notifications are enabled. | [optional] 
**EnableFailedNotify** | Pointer to **bool** | Determines whether failed task notifications are enabled. | [optional] 
**EnableNotifyGroup** | Pointer to **bool** | Determines whether e-mail notifications to admin group&#39;s e-mail address are enabled. | [optional] 
**EnableNotifyUser** | Pointer to **bool** | Determines whether e-mail notifications to an admin member&#39;s e-mail address are enabled. | [optional] 
**EnableRejectedNotify** | Pointer to **bool** | Determines whether rejected task notifications are enabled. | [optional] 
**EnableRescheduledNotify** | Pointer to **bool** | Determines whether rescheduled task notifications are enabled. | [optional] 
**EnableSucceededNotify** | Pointer to **bool** | Determines whether succeeded task notifications are enabled. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FailedNotifyTo** | Pointer to **string** | The destination for failed task notifications. | [optional] 
**RejectedNotifyTo** | Pointer to **string** | The destination for rejected task notifications. | [optional] 
**RescheduledNotifyTo** | Pointer to **string** | The destination for rescheduled task notifications. | [optional] 
**SubmitterComment** | Pointer to **string** | The requirement for the comment when a submitter submits a task for approval. | [optional] 
**SubmitterGroup** | Pointer to **string** | The submitter admininstration group. | [optional] 
**SucceededNotifyTo** | Pointer to **string** | The destination for succeeded task notifications. | [optional] 
**TicketNumber** | Pointer to **string** | The requirement for the ticket number when a submitter submits a task for approval. | [optional] 
**Result** | Pointer to [**Approvalworkflow**](Approvalworkflow.md) |  | [optional] 

## Methods

### NewGetApprovalworkflowResponse

`func NewGetApprovalworkflowResponse() *GetApprovalworkflowResponse`

NewGetApprovalworkflowResponse instantiates a new GetApprovalworkflowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApprovalworkflowResponseWithDefaults

`func NewGetApprovalworkflowResponseWithDefaults() *GetApprovalworkflowResponse`

NewGetApprovalworkflowResponseWithDefaults instantiates a new GetApprovalworkflowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetApprovalworkflowResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetApprovalworkflowResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetApprovalworkflowResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetApprovalworkflowResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetApprovalGroup

`func (o *GetApprovalworkflowResponse) GetApprovalGroup() string`

GetApprovalGroup returns the ApprovalGroup field if non-nil, zero value otherwise.

### GetApprovalGroupOk

`func (o *GetApprovalworkflowResponse) GetApprovalGroupOk() (*string, bool)`

GetApprovalGroupOk returns a tuple with the ApprovalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalGroup

`func (o *GetApprovalworkflowResponse) SetApprovalGroup(v string)`

SetApprovalGroup sets ApprovalGroup field to given value.

### HasApprovalGroup

`func (o *GetApprovalworkflowResponse) HasApprovalGroup() bool`

HasApprovalGroup returns a boolean if a field has been set.

### GetApprovalNotifyTo

`func (o *GetApprovalworkflowResponse) GetApprovalNotifyTo() string`

GetApprovalNotifyTo returns the ApprovalNotifyTo field if non-nil, zero value otherwise.

### GetApprovalNotifyToOk

`func (o *GetApprovalworkflowResponse) GetApprovalNotifyToOk() (*string, bool)`

GetApprovalNotifyToOk returns a tuple with the ApprovalNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalNotifyTo

`func (o *GetApprovalworkflowResponse) SetApprovalNotifyTo(v string)`

SetApprovalNotifyTo sets ApprovalNotifyTo field to given value.

### HasApprovalNotifyTo

`func (o *GetApprovalworkflowResponse) HasApprovalNotifyTo() bool`

HasApprovalNotifyTo returns a boolean if a field has been set.

### GetApprovedNotifyTo

`func (o *GetApprovalworkflowResponse) GetApprovedNotifyTo() string`

GetApprovedNotifyTo returns the ApprovedNotifyTo field if non-nil, zero value otherwise.

### GetApprovedNotifyToOk

`func (o *GetApprovalworkflowResponse) GetApprovedNotifyToOk() (*string, bool)`

GetApprovedNotifyToOk returns a tuple with the ApprovedNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedNotifyTo

`func (o *GetApprovalworkflowResponse) SetApprovedNotifyTo(v string)`

SetApprovedNotifyTo sets ApprovedNotifyTo field to given value.

### HasApprovedNotifyTo

`func (o *GetApprovalworkflowResponse) HasApprovedNotifyTo() bool`

HasApprovedNotifyTo returns a boolean if a field has been set.

### GetApproverComment

`func (o *GetApprovalworkflowResponse) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *GetApprovalworkflowResponse) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *GetApprovalworkflowResponse) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *GetApprovalworkflowResponse) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### GetEnableApprovalNotify

`func (o *GetApprovalworkflowResponse) GetEnableApprovalNotify() bool`

GetEnableApprovalNotify returns the EnableApprovalNotify field if non-nil, zero value otherwise.

### GetEnableApprovalNotifyOk

`func (o *GetApprovalworkflowResponse) GetEnableApprovalNotifyOk() (*bool, bool)`

GetEnableApprovalNotifyOk returns a tuple with the EnableApprovalNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApprovalNotify

`func (o *GetApprovalworkflowResponse) SetEnableApprovalNotify(v bool)`

SetEnableApprovalNotify sets EnableApprovalNotify field to given value.

### HasEnableApprovalNotify

`func (o *GetApprovalworkflowResponse) HasEnableApprovalNotify() bool`

HasEnableApprovalNotify returns a boolean if a field has been set.

### GetEnableApprovedNotify

`func (o *GetApprovalworkflowResponse) GetEnableApprovedNotify() bool`

GetEnableApprovedNotify returns the EnableApprovedNotify field if non-nil, zero value otherwise.

### GetEnableApprovedNotifyOk

`func (o *GetApprovalworkflowResponse) GetEnableApprovedNotifyOk() (*bool, bool)`

GetEnableApprovedNotifyOk returns a tuple with the EnableApprovedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApprovedNotify

`func (o *GetApprovalworkflowResponse) SetEnableApprovedNotify(v bool)`

SetEnableApprovedNotify sets EnableApprovedNotify field to given value.

### HasEnableApprovedNotify

`func (o *GetApprovalworkflowResponse) HasEnableApprovedNotify() bool`

HasEnableApprovedNotify returns a boolean if a field has been set.

### GetEnableFailedNotify

`func (o *GetApprovalworkflowResponse) GetEnableFailedNotify() bool`

GetEnableFailedNotify returns the EnableFailedNotify field if non-nil, zero value otherwise.

### GetEnableFailedNotifyOk

`func (o *GetApprovalworkflowResponse) GetEnableFailedNotifyOk() (*bool, bool)`

GetEnableFailedNotifyOk returns a tuple with the EnableFailedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFailedNotify

`func (o *GetApprovalworkflowResponse) SetEnableFailedNotify(v bool)`

SetEnableFailedNotify sets EnableFailedNotify field to given value.

### HasEnableFailedNotify

`func (o *GetApprovalworkflowResponse) HasEnableFailedNotify() bool`

HasEnableFailedNotify returns a boolean if a field has been set.

### GetEnableNotifyGroup

`func (o *GetApprovalworkflowResponse) GetEnableNotifyGroup() bool`

GetEnableNotifyGroup returns the EnableNotifyGroup field if non-nil, zero value otherwise.

### GetEnableNotifyGroupOk

`func (o *GetApprovalworkflowResponse) GetEnableNotifyGroupOk() (*bool, bool)`

GetEnableNotifyGroupOk returns a tuple with the EnableNotifyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifyGroup

`func (o *GetApprovalworkflowResponse) SetEnableNotifyGroup(v bool)`

SetEnableNotifyGroup sets EnableNotifyGroup field to given value.

### HasEnableNotifyGroup

`func (o *GetApprovalworkflowResponse) HasEnableNotifyGroup() bool`

HasEnableNotifyGroup returns a boolean if a field has been set.

### GetEnableNotifyUser

`func (o *GetApprovalworkflowResponse) GetEnableNotifyUser() bool`

GetEnableNotifyUser returns the EnableNotifyUser field if non-nil, zero value otherwise.

### GetEnableNotifyUserOk

`func (o *GetApprovalworkflowResponse) GetEnableNotifyUserOk() (*bool, bool)`

GetEnableNotifyUserOk returns a tuple with the EnableNotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifyUser

`func (o *GetApprovalworkflowResponse) SetEnableNotifyUser(v bool)`

SetEnableNotifyUser sets EnableNotifyUser field to given value.

### HasEnableNotifyUser

`func (o *GetApprovalworkflowResponse) HasEnableNotifyUser() bool`

HasEnableNotifyUser returns a boolean if a field has been set.

### GetEnableRejectedNotify

`func (o *GetApprovalworkflowResponse) GetEnableRejectedNotify() bool`

GetEnableRejectedNotify returns the EnableRejectedNotify field if non-nil, zero value otherwise.

### GetEnableRejectedNotifyOk

`func (o *GetApprovalworkflowResponse) GetEnableRejectedNotifyOk() (*bool, bool)`

GetEnableRejectedNotifyOk returns a tuple with the EnableRejectedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRejectedNotify

`func (o *GetApprovalworkflowResponse) SetEnableRejectedNotify(v bool)`

SetEnableRejectedNotify sets EnableRejectedNotify field to given value.

### HasEnableRejectedNotify

`func (o *GetApprovalworkflowResponse) HasEnableRejectedNotify() bool`

HasEnableRejectedNotify returns a boolean if a field has been set.

### GetEnableRescheduledNotify

`func (o *GetApprovalworkflowResponse) GetEnableRescheduledNotify() bool`

GetEnableRescheduledNotify returns the EnableRescheduledNotify field if non-nil, zero value otherwise.

### GetEnableRescheduledNotifyOk

`func (o *GetApprovalworkflowResponse) GetEnableRescheduledNotifyOk() (*bool, bool)`

GetEnableRescheduledNotifyOk returns a tuple with the EnableRescheduledNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRescheduledNotify

`func (o *GetApprovalworkflowResponse) SetEnableRescheduledNotify(v bool)`

SetEnableRescheduledNotify sets EnableRescheduledNotify field to given value.

### HasEnableRescheduledNotify

`func (o *GetApprovalworkflowResponse) HasEnableRescheduledNotify() bool`

HasEnableRescheduledNotify returns a boolean if a field has been set.

### GetEnableSucceededNotify

`func (o *GetApprovalworkflowResponse) GetEnableSucceededNotify() bool`

GetEnableSucceededNotify returns the EnableSucceededNotify field if non-nil, zero value otherwise.

### GetEnableSucceededNotifyOk

`func (o *GetApprovalworkflowResponse) GetEnableSucceededNotifyOk() (*bool, bool)`

GetEnableSucceededNotifyOk returns a tuple with the EnableSucceededNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSucceededNotify

`func (o *GetApprovalworkflowResponse) SetEnableSucceededNotify(v bool)`

SetEnableSucceededNotify sets EnableSucceededNotify field to given value.

### HasEnableSucceededNotify

`func (o *GetApprovalworkflowResponse) HasEnableSucceededNotify() bool`

HasEnableSucceededNotify returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetApprovalworkflowResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetApprovalworkflowResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetApprovalworkflowResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetApprovalworkflowResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFailedNotifyTo

`func (o *GetApprovalworkflowResponse) GetFailedNotifyTo() string`

GetFailedNotifyTo returns the FailedNotifyTo field if non-nil, zero value otherwise.

### GetFailedNotifyToOk

`func (o *GetApprovalworkflowResponse) GetFailedNotifyToOk() (*string, bool)`

GetFailedNotifyToOk returns a tuple with the FailedNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNotifyTo

`func (o *GetApprovalworkflowResponse) SetFailedNotifyTo(v string)`

SetFailedNotifyTo sets FailedNotifyTo field to given value.

### HasFailedNotifyTo

`func (o *GetApprovalworkflowResponse) HasFailedNotifyTo() bool`

HasFailedNotifyTo returns a boolean if a field has been set.

### GetRejectedNotifyTo

`func (o *GetApprovalworkflowResponse) GetRejectedNotifyTo() string`

GetRejectedNotifyTo returns the RejectedNotifyTo field if non-nil, zero value otherwise.

### GetRejectedNotifyToOk

`func (o *GetApprovalworkflowResponse) GetRejectedNotifyToOk() (*string, bool)`

GetRejectedNotifyToOk returns a tuple with the RejectedNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNotifyTo

`func (o *GetApprovalworkflowResponse) SetRejectedNotifyTo(v string)`

SetRejectedNotifyTo sets RejectedNotifyTo field to given value.

### HasRejectedNotifyTo

`func (o *GetApprovalworkflowResponse) HasRejectedNotifyTo() bool`

HasRejectedNotifyTo returns a boolean if a field has been set.

### GetRescheduledNotifyTo

`func (o *GetApprovalworkflowResponse) GetRescheduledNotifyTo() string`

GetRescheduledNotifyTo returns the RescheduledNotifyTo field if non-nil, zero value otherwise.

### GetRescheduledNotifyToOk

`func (o *GetApprovalworkflowResponse) GetRescheduledNotifyToOk() (*string, bool)`

GetRescheduledNotifyToOk returns a tuple with the RescheduledNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescheduledNotifyTo

`func (o *GetApprovalworkflowResponse) SetRescheduledNotifyTo(v string)`

SetRescheduledNotifyTo sets RescheduledNotifyTo field to given value.

### HasRescheduledNotifyTo

`func (o *GetApprovalworkflowResponse) HasRescheduledNotifyTo() bool`

HasRescheduledNotifyTo returns a boolean if a field has been set.

### GetSubmitterComment

`func (o *GetApprovalworkflowResponse) GetSubmitterComment() string`

GetSubmitterComment returns the SubmitterComment field if non-nil, zero value otherwise.

### GetSubmitterCommentOk

`func (o *GetApprovalworkflowResponse) GetSubmitterCommentOk() (*string, bool)`

GetSubmitterCommentOk returns a tuple with the SubmitterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterComment

`func (o *GetApprovalworkflowResponse) SetSubmitterComment(v string)`

SetSubmitterComment sets SubmitterComment field to given value.

### HasSubmitterComment

`func (o *GetApprovalworkflowResponse) HasSubmitterComment() bool`

HasSubmitterComment returns a boolean if a field has been set.

### GetSubmitterGroup

`func (o *GetApprovalworkflowResponse) GetSubmitterGroup() string`

GetSubmitterGroup returns the SubmitterGroup field if non-nil, zero value otherwise.

### GetSubmitterGroupOk

`func (o *GetApprovalworkflowResponse) GetSubmitterGroupOk() (*string, bool)`

GetSubmitterGroupOk returns a tuple with the SubmitterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterGroup

`func (o *GetApprovalworkflowResponse) SetSubmitterGroup(v string)`

SetSubmitterGroup sets SubmitterGroup field to given value.

### HasSubmitterGroup

`func (o *GetApprovalworkflowResponse) HasSubmitterGroup() bool`

HasSubmitterGroup returns a boolean if a field has been set.

### GetSucceededNotifyTo

`func (o *GetApprovalworkflowResponse) GetSucceededNotifyTo() string`

GetSucceededNotifyTo returns the SucceededNotifyTo field if non-nil, zero value otherwise.

### GetSucceededNotifyToOk

`func (o *GetApprovalworkflowResponse) GetSucceededNotifyToOk() (*string, bool)`

GetSucceededNotifyToOk returns a tuple with the SucceededNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededNotifyTo

`func (o *GetApprovalworkflowResponse) SetSucceededNotifyTo(v string)`

SetSucceededNotifyTo sets SucceededNotifyTo field to given value.

### HasSucceededNotifyTo

`func (o *GetApprovalworkflowResponse) HasSucceededNotifyTo() bool`

HasSucceededNotifyTo returns a boolean if a field has been set.

### GetTicketNumber

`func (o *GetApprovalworkflowResponse) GetTicketNumber() string`

GetTicketNumber returns the TicketNumber field if non-nil, zero value otherwise.

### GetTicketNumberOk

`func (o *GetApprovalworkflowResponse) GetTicketNumberOk() (*string, bool)`

GetTicketNumberOk returns a tuple with the TicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumber

`func (o *GetApprovalworkflowResponse) SetTicketNumber(v string)`

SetTicketNumber sets TicketNumber field to given value.

### HasTicketNumber

`func (o *GetApprovalworkflowResponse) HasTicketNumber() bool`

HasTicketNumber returns a boolean if a field has been set.

### GetResult

`func (o *GetApprovalworkflowResponse) GetResult() Approvalworkflow`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetApprovalworkflowResponse) GetResultOk() (*Approvalworkflow, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetApprovalworkflowResponse) SetResult(v Approvalworkflow)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetApprovalworkflowResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


