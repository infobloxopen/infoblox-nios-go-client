# NotificationRule

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

## Methods

### NewNotificationRule

`func NewNotificationRule() *NotificationRule`

NewNotificationRule instantiates a new NotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleWithDefaults

`func NewNotificationRuleWithDefaults() *NotificationRule`

NewNotificationRuleWithDefaults instantiates a new NotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *NotificationRule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *NotificationRule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *NotificationRule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *NotificationRule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllMembers

`func (o *NotificationRule) GetAllMembers() bool`

GetAllMembers returns the AllMembers field if non-nil, zero value otherwise.

### GetAllMembersOk

`func (o *NotificationRule) GetAllMembersOk() (*bool, bool)`

GetAllMembersOk returns a tuple with the AllMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMembers

`func (o *NotificationRule) SetAllMembers(v bool)`

SetAllMembers sets AllMembers field to given value.

### HasAllMembers

`func (o *NotificationRule) HasAllMembers() bool`

HasAllMembers returns a boolean if a field has been set.

### GetComment

`func (o *NotificationRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NotificationRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NotificationRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NotificationRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *NotificationRule) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *NotificationRule) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *NotificationRule) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *NotificationRule) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableEventDeduplication

`func (o *NotificationRule) GetEnableEventDeduplication() bool`

GetEnableEventDeduplication returns the EnableEventDeduplication field if non-nil, zero value otherwise.

### GetEnableEventDeduplicationOk

`func (o *NotificationRule) GetEnableEventDeduplicationOk() (*bool, bool)`

GetEnableEventDeduplicationOk returns a tuple with the EnableEventDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventDeduplication

`func (o *NotificationRule) SetEnableEventDeduplication(v bool)`

SetEnableEventDeduplication sets EnableEventDeduplication field to given value.

### HasEnableEventDeduplication

`func (o *NotificationRule) HasEnableEventDeduplication() bool`

HasEnableEventDeduplication returns a boolean if a field has been set.

### GetEnableEventDeduplicationLog

`func (o *NotificationRule) GetEnableEventDeduplicationLog() bool`

GetEnableEventDeduplicationLog returns the EnableEventDeduplicationLog field if non-nil, zero value otherwise.

### GetEnableEventDeduplicationLogOk

`func (o *NotificationRule) GetEnableEventDeduplicationLogOk() (*bool, bool)`

GetEnableEventDeduplicationLogOk returns a tuple with the EnableEventDeduplicationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventDeduplicationLog

`func (o *NotificationRule) SetEnableEventDeduplicationLog(v bool)`

SetEnableEventDeduplicationLog sets EnableEventDeduplicationLog field to given value.

### HasEnableEventDeduplicationLog

`func (o *NotificationRule) HasEnableEventDeduplicationLog() bool`

HasEnableEventDeduplicationLog returns a boolean if a field has been set.

### GetEventDeduplicationFields

`func (o *NotificationRule) GetEventDeduplicationFields() []string`

GetEventDeduplicationFields returns the EventDeduplicationFields field if non-nil, zero value otherwise.

### GetEventDeduplicationFieldsOk

`func (o *NotificationRule) GetEventDeduplicationFieldsOk() (*[]string, bool)`

GetEventDeduplicationFieldsOk returns a tuple with the EventDeduplicationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDeduplicationFields

`func (o *NotificationRule) SetEventDeduplicationFields(v []string)`

SetEventDeduplicationFields sets EventDeduplicationFields field to given value.

### HasEventDeduplicationFields

`func (o *NotificationRule) HasEventDeduplicationFields() bool`

HasEventDeduplicationFields returns a boolean if a field has been set.

### GetEventDeduplicationLookbackPeriod

`func (o *NotificationRule) GetEventDeduplicationLookbackPeriod() int64`

GetEventDeduplicationLookbackPeriod returns the EventDeduplicationLookbackPeriod field if non-nil, zero value otherwise.

### GetEventDeduplicationLookbackPeriodOk

`func (o *NotificationRule) GetEventDeduplicationLookbackPeriodOk() (*int64, bool)`

GetEventDeduplicationLookbackPeriodOk returns a tuple with the EventDeduplicationLookbackPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDeduplicationLookbackPeriod

`func (o *NotificationRule) SetEventDeduplicationLookbackPeriod(v int64)`

SetEventDeduplicationLookbackPeriod sets EventDeduplicationLookbackPeriod field to given value.

### HasEventDeduplicationLookbackPeriod

`func (o *NotificationRule) HasEventDeduplicationLookbackPeriod() bool`

HasEventDeduplicationLookbackPeriod returns a boolean if a field has been set.

### GetEventPriority

`func (o *NotificationRule) GetEventPriority() string`

GetEventPriority returns the EventPriority field if non-nil, zero value otherwise.

### GetEventPriorityOk

`func (o *NotificationRule) GetEventPriorityOk() (*string, bool)`

GetEventPriorityOk returns a tuple with the EventPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPriority

`func (o *NotificationRule) SetEventPriority(v string)`

SetEventPriority sets EventPriority field to given value.

### HasEventPriority

`func (o *NotificationRule) HasEventPriority() bool`

HasEventPriority returns a boolean if a field has been set.

### GetEventType

`func (o *NotificationRule) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NotificationRule) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NotificationRule) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *NotificationRule) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetExpressionList

`func (o *NotificationRule) GetExpressionList() []NotificationRuleExpressionList`

GetExpressionList returns the ExpressionList field if non-nil, zero value otherwise.

### GetExpressionListOk

`func (o *NotificationRule) GetExpressionListOk() (*[]NotificationRuleExpressionList, bool)`

GetExpressionListOk returns a tuple with the ExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionList

`func (o *NotificationRule) SetExpressionList(v []NotificationRuleExpressionList)`

SetExpressionList sets ExpressionList field to given value.

### HasExpressionList

`func (o *NotificationRule) HasExpressionList() bool`

HasExpressionList returns a boolean if a field has been set.

### GetName

`func (o *NotificationRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationAction

`func (o *NotificationRule) GetNotificationAction() string`

GetNotificationAction returns the NotificationAction field if non-nil, zero value otherwise.

### GetNotificationActionOk

`func (o *NotificationRule) GetNotificationActionOk() (*string, bool)`

GetNotificationActionOk returns a tuple with the NotificationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationAction

`func (o *NotificationRule) SetNotificationAction(v string)`

SetNotificationAction sets NotificationAction field to given value.

### HasNotificationAction

`func (o *NotificationRule) HasNotificationAction() bool`

HasNotificationAction returns a boolean if a field has been set.

### GetNotificationTarget

`func (o *NotificationRule) GetNotificationTarget() string`

GetNotificationTarget returns the NotificationTarget field if non-nil, zero value otherwise.

### GetNotificationTargetOk

`func (o *NotificationRule) GetNotificationTargetOk() (*string, bool)`

GetNotificationTargetOk returns a tuple with the NotificationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTarget

`func (o *NotificationRule) SetNotificationTarget(v string)`

SetNotificationTarget sets NotificationTarget field to given value.

### HasNotificationTarget

`func (o *NotificationRule) HasNotificationTarget() bool`

HasNotificationTarget returns a boolean if a field has been set.

### GetPublishSettings

`func (o *NotificationRule) GetPublishSettings() NotificationRulePublishSettings`

GetPublishSettings returns the PublishSettings field if non-nil, zero value otherwise.

### GetPublishSettingsOk

`func (o *NotificationRule) GetPublishSettingsOk() (*NotificationRulePublishSettings, bool)`

GetPublishSettingsOk returns a tuple with the PublishSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishSettings

`func (o *NotificationRule) SetPublishSettings(v NotificationRulePublishSettings)`

SetPublishSettings sets PublishSettings field to given value.

### HasPublishSettings

`func (o *NotificationRule) HasPublishSettings() bool`

HasPublishSettings returns a boolean if a field has been set.

### GetScheduledEvent

`func (o *NotificationRule) GetScheduledEvent() NotificationRuleScheduledEvent`

GetScheduledEvent returns the ScheduledEvent field if non-nil, zero value otherwise.

### GetScheduledEventOk

`func (o *NotificationRule) GetScheduledEventOk() (*NotificationRuleScheduledEvent, bool)`

GetScheduledEventOk returns a tuple with the ScheduledEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEvent

`func (o *NotificationRule) SetScheduledEvent(v NotificationRuleScheduledEvent)`

SetScheduledEvent sets ScheduledEvent field to given value.

### HasScheduledEvent

`func (o *NotificationRule) HasScheduledEvent() bool`

HasScheduledEvent returns a boolean if a field has been set.

### GetSelectedMembers

`func (o *NotificationRule) GetSelectedMembers() []string`

GetSelectedMembers returns the SelectedMembers field if non-nil, zero value otherwise.

### GetSelectedMembersOk

`func (o *NotificationRule) GetSelectedMembersOk() (*[]string, bool)`

GetSelectedMembersOk returns a tuple with the SelectedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedMembers

`func (o *NotificationRule) SetSelectedMembers(v []string)`

SetSelectedMembers sets SelectedMembers field to given value.

### HasSelectedMembers

`func (o *NotificationRule) HasSelectedMembers() bool`

HasSelectedMembers returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *NotificationRule) GetTemplateInstance() NotificationRuleTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *NotificationRule) GetTemplateInstanceOk() (*NotificationRuleTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *NotificationRule) SetTemplateInstance(v NotificationRuleTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *NotificationRule) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetUsePublishSettings

`func (o *NotificationRule) GetUsePublishSettings() bool`

GetUsePublishSettings returns the UsePublishSettings field if non-nil, zero value otherwise.

### GetUsePublishSettingsOk

`func (o *NotificationRule) GetUsePublishSettingsOk() (*bool, bool)`

GetUsePublishSettingsOk returns a tuple with the UsePublishSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePublishSettings

`func (o *NotificationRule) SetUsePublishSettings(v bool)`

SetUsePublishSettings sets UsePublishSettings field to given value.

### HasUsePublishSettings

`func (o *NotificationRule) HasUsePublishSettings() bool`

HasUsePublishSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


