# NotificationRestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ActionName** | Pointer to **string** | The action name. | [optional] [readonly] 
**AddedOn** | Pointer to **int64** | The time stamp when a template was added. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for this REST API template. | [optional] 
**Content** | Pointer to **string** | The JSON formatted content of a template. The data passed by content creates parameters for a template. | [optional] 
**EventType** | Pointer to **[]string** | The event type. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of a notification REST template. | [optional] 
**OutboundType** | Pointer to **string** | The outbound type for the template. | [optional] [readonly] 
**Parameters** | Pointer to [**[]NotificationRestTemplateParameters**](NotificationRestTemplateParameters.md) | The notification REST template parameters. | [optional] [readonly] 
**TemplateType** | Pointer to **string** | The template type. | [optional] [readonly] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] [readonly] 

## Methods

### NewNotificationRestTemplate

`func NewNotificationRestTemplate() *NotificationRestTemplate`

NewNotificationRestTemplate instantiates a new NotificationRestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRestTemplateWithDefaults

`func NewNotificationRestTemplateWithDefaults() *NotificationRestTemplate`

NewNotificationRestTemplateWithDefaults instantiates a new NotificationRestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *NotificationRestTemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *NotificationRestTemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *NotificationRestTemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *NotificationRestTemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActionName

`func (o *NotificationRestTemplate) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *NotificationRestTemplate) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *NotificationRestTemplate) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *NotificationRestTemplate) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetAddedOn

`func (o *NotificationRestTemplate) GetAddedOn() int64`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *NotificationRestTemplate) GetAddedOnOk() (*int64, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *NotificationRestTemplate) SetAddedOn(v int64)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *NotificationRestTemplate) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetComment

`func (o *NotificationRestTemplate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NotificationRestTemplate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NotificationRestTemplate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NotificationRestTemplate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContent

`func (o *NotificationRestTemplate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationRestTemplate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationRestTemplate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NotificationRestTemplate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEventType

`func (o *NotificationRestTemplate) GetEventType() []string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NotificationRestTemplate) GetEventTypeOk() (*[]string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NotificationRestTemplate) SetEventType(v []string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *NotificationRestTemplate) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *NotificationRestTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRestTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRestTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationRestTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundType

`func (o *NotificationRestTemplate) GetOutboundType() string`

GetOutboundType returns the OutboundType field if non-nil, zero value otherwise.

### GetOutboundTypeOk

`func (o *NotificationRestTemplate) GetOutboundTypeOk() (*string, bool)`

GetOutboundTypeOk returns a tuple with the OutboundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundType

`func (o *NotificationRestTemplate) SetOutboundType(v string)`

SetOutboundType sets OutboundType field to given value.

### HasOutboundType

`func (o *NotificationRestTemplate) HasOutboundType() bool`

HasOutboundType returns a boolean if a field has been set.

### GetParameters

`func (o *NotificationRestTemplate) GetParameters() []NotificationRestTemplateParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *NotificationRestTemplate) GetParametersOk() (*[]NotificationRestTemplateParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *NotificationRestTemplate) SetParameters(v []NotificationRestTemplateParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *NotificationRestTemplate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTemplateType

`func (o *NotificationRestTemplate) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *NotificationRestTemplate) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *NotificationRestTemplate) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *NotificationRestTemplate) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *NotificationRestTemplate) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *NotificationRestTemplate) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *NotificationRestTemplate) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *NotificationRestTemplate) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


