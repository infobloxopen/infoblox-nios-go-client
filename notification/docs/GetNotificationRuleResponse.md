# GetNotificationRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllMembers** | Pointer to **bool** | Determines whether the notification rule is applied on all members or not. When this is set to False, the notification rule is applied only on selected_members. | [optional] 
**Comment** | Pointer to **string** | The notification rule descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a notification rule is disabled or not. When this is set to False, the notification rule is enabled. | [optional] 
**EnableEventDeduplication** | Pointer to **bool** | Determines whether the notification rule for event deduplication is enabled. Note that to enable event deduplication, you must set at least one deduplication field. | [optional] 
**EnableEventDeduplicationLog** | Pointer to **bool** | Determines whether the notification rule for the event deduplication syslog is enabled. | [optional] 
**EventDeduplicationFields** | Pointer to **[]string** | The list of fields that must be used in the notification rule for event deduplication. | [optional] 
**EventDeduplicationLookbackPeriod** | Pointer to **int64** | The lookback period for the notification rule for event deduplication. | [optional] 
**EventPriority** | Pointer to **string** | Event priority. | [optional] 
**EventType** | Pointer to **string** | The notification rule event type. | [optional] 
**ExpressionList** | Pointer to [**[]NotificationRuleExpressionList**](NotificationRuleExpressionList.md) | The notification rule expression list. | [optional] 
**Name** | Pointer to **string** | The notification rule name. | [optional] 
**NotificationAction** | Pointer to **string** | The notification rule action is applied if expression list evaluates to True. | [optional] 
**NotificationTarget** | Pointer to **string** | The notification target. | [optional] 
**PublishSettings** | Pointer to [**NotificationRulePublishSettings**](NotificationRulePublishSettings.md) |  | [optional] 
**ScheduledEvent** | Pointer to [**NotificationRuleScheduledEvent**](NotificationRuleScheduledEvent.md) |  | [optional] 
**SelectedMembers** | Pointer to **[]string** | The list of the members on which the notification rule is applied. | [optional] 
**TemplateInstance** | Pointer to [**NotificationRuleTemplateInstance**](NotificationRuleTemplateInstance.md) |  | [optional] 
**UsePublishSettings** | Pointer to **bool** | Use flag for: publish_settings | [optional] 
**Result** | Pointer to [**NotificationRule**](NotificationRule.md) |  | [optional] 

## Methods

### NewGetNotificationRuleResponse

`func NewGetNotificationRuleResponse() *GetNotificationRuleResponse`

NewGetNotificationRuleResponse instantiates a new GetNotificationRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationRuleResponseWithDefaults

`func NewGetNotificationRuleResponseWithDefaults() *GetNotificationRuleResponse`

NewGetNotificationRuleResponseWithDefaults instantiates a new GetNotificationRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNotificationRuleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNotificationRuleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNotificationRuleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNotificationRuleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllMembers

`func (o *GetNotificationRuleResponse) GetAllMembers() bool`

GetAllMembers returns the AllMembers field if non-nil, zero value otherwise.

### GetAllMembersOk

`func (o *GetNotificationRuleResponse) GetAllMembersOk() (*bool, bool)`

GetAllMembersOk returns a tuple with the AllMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMembers

`func (o *GetNotificationRuleResponse) SetAllMembers(v bool)`

SetAllMembers sets AllMembers field to given value.

### HasAllMembers

`func (o *GetNotificationRuleResponse) HasAllMembers() bool`

HasAllMembers returns a boolean if a field has been set.

### GetComment

`func (o *GetNotificationRuleResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNotificationRuleResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNotificationRuleResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNotificationRuleResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetNotificationRuleResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetNotificationRuleResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetNotificationRuleResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetNotificationRuleResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableEventDeduplication

`func (o *GetNotificationRuleResponse) GetEnableEventDeduplication() bool`

GetEnableEventDeduplication returns the EnableEventDeduplication field if non-nil, zero value otherwise.

### GetEnableEventDeduplicationOk

`func (o *GetNotificationRuleResponse) GetEnableEventDeduplicationOk() (*bool, bool)`

GetEnableEventDeduplicationOk returns a tuple with the EnableEventDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventDeduplication

`func (o *GetNotificationRuleResponse) SetEnableEventDeduplication(v bool)`

SetEnableEventDeduplication sets EnableEventDeduplication field to given value.

### HasEnableEventDeduplication

`func (o *GetNotificationRuleResponse) HasEnableEventDeduplication() bool`

HasEnableEventDeduplication returns a boolean if a field has been set.

### GetEnableEventDeduplicationLog

`func (o *GetNotificationRuleResponse) GetEnableEventDeduplicationLog() bool`

GetEnableEventDeduplicationLog returns the EnableEventDeduplicationLog field if non-nil, zero value otherwise.

### GetEnableEventDeduplicationLogOk

`func (o *GetNotificationRuleResponse) GetEnableEventDeduplicationLogOk() (*bool, bool)`

GetEnableEventDeduplicationLogOk returns a tuple with the EnableEventDeduplicationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventDeduplicationLog

`func (o *GetNotificationRuleResponse) SetEnableEventDeduplicationLog(v bool)`

SetEnableEventDeduplicationLog sets EnableEventDeduplicationLog field to given value.

### HasEnableEventDeduplicationLog

`func (o *GetNotificationRuleResponse) HasEnableEventDeduplicationLog() bool`

HasEnableEventDeduplicationLog returns a boolean if a field has been set.

### GetEventDeduplicationFields

`func (o *GetNotificationRuleResponse) GetEventDeduplicationFields() []string`

GetEventDeduplicationFields returns the EventDeduplicationFields field if non-nil, zero value otherwise.

### GetEventDeduplicationFieldsOk

`func (o *GetNotificationRuleResponse) GetEventDeduplicationFieldsOk() (*[]string, bool)`

GetEventDeduplicationFieldsOk returns a tuple with the EventDeduplicationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDeduplicationFields

`func (o *GetNotificationRuleResponse) SetEventDeduplicationFields(v []string)`

SetEventDeduplicationFields sets EventDeduplicationFields field to given value.

### HasEventDeduplicationFields

`func (o *GetNotificationRuleResponse) HasEventDeduplicationFields() bool`

HasEventDeduplicationFields returns a boolean if a field has been set.

### GetEventDeduplicationLookbackPeriod

`func (o *GetNotificationRuleResponse) GetEventDeduplicationLookbackPeriod() int64`

GetEventDeduplicationLookbackPeriod returns the EventDeduplicationLookbackPeriod field if non-nil, zero value otherwise.

### GetEventDeduplicationLookbackPeriodOk

`func (o *GetNotificationRuleResponse) GetEventDeduplicationLookbackPeriodOk() (*int64, bool)`

GetEventDeduplicationLookbackPeriodOk returns a tuple with the EventDeduplicationLookbackPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDeduplicationLookbackPeriod

`func (o *GetNotificationRuleResponse) SetEventDeduplicationLookbackPeriod(v int64)`

SetEventDeduplicationLookbackPeriod sets EventDeduplicationLookbackPeriod field to given value.

### HasEventDeduplicationLookbackPeriod

`func (o *GetNotificationRuleResponse) HasEventDeduplicationLookbackPeriod() bool`

HasEventDeduplicationLookbackPeriod returns a boolean if a field has been set.

### GetEventPriority

`func (o *GetNotificationRuleResponse) GetEventPriority() string`

GetEventPriority returns the EventPriority field if non-nil, zero value otherwise.

### GetEventPriorityOk

`func (o *GetNotificationRuleResponse) GetEventPriorityOk() (*string, bool)`

GetEventPriorityOk returns a tuple with the EventPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPriority

`func (o *GetNotificationRuleResponse) SetEventPriority(v string)`

SetEventPriority sets EventPriority field to given value.

### HasEventPriority

`func (o *GetNotificationRuleResponse) HasEventPriority() bool`

HasEventPriority returns a boolean if a field has been set.

### GetEventType

`func (o *GetNotificationRuleResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetNotificationRuleResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetNotificationRuleResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GetNotificationRuleResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetExpressionList

`func (o *GetNotificationRuleResponse) GetExpressionList() []NotificationRuleExpressionList`

GetExpressionList returns the ExpressionList field if non-nil, zero value otherwise.

### GetExpressionListOk

`func (o *GetNotificationRuleResponse) GetExpressionListOk() (*[]NotificationRuleExpressionList, bool)`

GetExpressionListOk returns a tuple with the ExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionList

`func (o *GetNotificationRuleResponse) SetExpressionList(v []NotificationRuleExpressionList)`

SetExpressionList sets ExpressionList field to given value.

### HasExpressionList

`func (o *GetNotificationRuleResponse) HasExpressionList() bool`

HasExpressionList returns a boolean if a field has been set.

### GetName

`func (o *GetNotificationRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNotificationRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNotificationRuleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNotificationRuleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationAction

`func (o *GetNotificationRuleResponse) GetNotificationAction() string`

GetNotificationAction returns the NotificationAction field if non-nil, zero value otherwise.

### GetNotificationActionOk

`func (o *GetNotificationRuleResponse) GetNotificationActionOk() (*string, bool)`

GetNotificationActionOk returns a tuple with the NotificationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationAction

`func (o *GetNotificationRuleResponse) SetNotificationAction(v string)`

SetNotificationAction sets NotificationAction field to given value.

### HasNotificationAction

`func (o *GetNotificationRuleResponse) HasNotificationAction() bool`

HasNotificationAction returns a boolean if a field has been set.

### GetNotificationTarget

`func (o *GetNotificationRuleResponse) GetNotificationTarget() string`

GetNotificationTarget returns the NotificationTarget field if non-nil, zero value otherwise.

### GetNotificationTargetOk

`func (o *GetNotificationRuleResponse) GetNotificationTargetOk() (*string, bool)`

GetNotificationTargetOk returns a tuple with the NotificationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTarget

`func (o *GetNotificationRuleResponse) SetNotificationTarget(v string)`

SetNotificationTarget sets NotificationTarget field to given value.

### HasNotificationTarget

`func (o *GetNotificationRuleResponse) HasNotificationTarget() bool`

HasNotificationTarget returns a boolean if a field has been set.

### GetPublishSettings

`func (o *GetNotificationRuleResponse) GetPublishSettings() NotificationRulePublishSettings`

GetPublishSettings returns the PublishSettings field if non-nil, zero value otherwise.

### GetPublishSettingsOk

`func (o *GetNotificationRuleResponse) GetPublishSettingsOk() (*NotificationRulePublishSettings, bool)`

GetPublishSettingsOk returns a tuple with the PublishSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishSettings

`func (o *GetNotificationRuleResponse) SetPublishSettings(v NotificationRulePublishSettings)`

SetPublishSettings sets PublishSettings field to given value.

### HasPublishSettings

`func (o *GetNotificationRuleResponse) HasPublishSettings() bool`

HasPublishSettings returns a boolean if a field has been set.

### GetScheduledEvent

`func (o *GetNotificationRuleResponse) GetScheduledEvent() NotificationRuleScheduledEvent`

GetScheduledEvent returns the ScheduledEvent field if non-nil, zero value otherwise.

### GetScheduledEventOk

`func (o *GetNotificationRuleResponse) GetScheduledEventOk() (*NotificationRuleScheduledEvent, bool)`

GetScheduledEventOk returns a tuple with the ScheduledEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEvent

`func (o *GetNotificationRuleResponse) SetScheduledEvent(v NotificationRuleScheduledEvent)`

SetScheduledEvent sets ScheduledEvent field to given value.

### HasScheduledEvent

`func (o *GetNotificationRuleResponse) HasScheduledEvent() bool`

HasScheduledEvent returns a boolean if a field has been set.

### GetSelectedMembers

`func (o *GetNotificationRuleResponse) GetSelectedMembers() []string`

GetSelectedMembers returns the SelectedMembers field if non-nil, zero value otherwise.

### GetSelectedMembersOk

`func (o *GetNotificationRuleResponse) GetSelectedMembersOk() (*[]string, bool)`

GetSelectedMembersOk returns a tuple with the SelectedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedMembers

`func (o *GetNotificationRuleResponse) SetSelectedMembers(v []string)`

SetSelectedMembers sets SelectedMembers field to given value.

### HasSelectedMembers

`func (o *GetNotificationRuleResponse) HasSelectedMembers() bool`

HasSelectedMembers returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *GetNotificationRuleResponse) GetTemplateInstance() NotificationRuleTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *GetNotificationRuleResponse) GetTemplateInstanceOk() (*NotificationRuleTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *GetNotificationRuleResponse) SetTemplateInstance(v NotificationRuleTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *GetNotificationRuleResponse) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetUsePublishSettings

`func (o *GetNotificationRuleResponse) GetUsePublishSettings() bool`

GetUsePublishSettings returns the UsePublishSettings field if non-nil, zero value otherwise.

### GetUsePublishSettingsOk

`func (o *GetNotificationRuleResponse) GetUsePublishSettingsOk() (*bool, bool)`

GetUsePublishSettingsOk returns a tuple with the UsePublishSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePublishSettings

`func (o *GetNotificationRuleResponse) SetUsePublishSettings(v bool)`

SetUsePublishSettings sets UsePublishSettings field to given value.

### HasUsePublishSettings

`func (o *GetNotificationRuleResponse) HasUsePublishSettings() bool`

HasUsePublishSettings returns a boolean if a field has been set.

### GetResult

`func (o *GetNotificationRuleResponse) GetResult() NotificationRule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNotificationRuleResponse) GetResultOk() (*NotificationRule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNotificationRuleResponse) SetResult(v NotificationRule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNotificationRuleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


