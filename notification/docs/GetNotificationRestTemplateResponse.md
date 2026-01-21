# GetNotificationRestTemplateResponse

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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] [readonly] 
**Result** | Pointer to [**NotificationRestTemplate**](NotificationRestTemplate.md) |  | [optional] 

## Methods

### NewGetNotificationRestTemplateResponse

`func NewGetNotificationRestTemplateResponse() *GetNotificationRestTemplateResponse`

NewGetNotificationRestTemplateResponse instantiates a new GetNotificationRestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationRestTemplateResponseWithDefaults

`func NewGetNotificationRestTemplateResponseWithDefaults() *GetNotificationRestTemplateResponse`

NewGetNotificationRestTemplateResponseWithDefaults instantiates a new GetNotificationRestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNotificationRestTemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNotificationRestTemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNotificationRestTemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNotificationRestTemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActionName

`func (o *GetNotificationRestTemplateResponse) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *GetNotificationRestTemplateResponse) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *GetNotificationRestTemplateResponse) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *GetNotificationRestTemplateResponse) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetAddedOn

`func (o *GetNotificationRestTemplateResponse) GetAddedOn() int64`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *GetNotificationRestTemplateResponse) GetAddedOnOk() (*int64, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *GetNotificationRestTemplateResponse) SetAddedOn(v int64)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *GetNotificationRestTemplateResponse) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetComment

`func (o *GetNotificationRestTemplateResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNotificationRestTemplateResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNotificationRestTemplateResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNotificationRestTemplateResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContent

`func (o *GetNotificationRestTemplateResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetNotificationRestTemplateResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetNotificationRestTemplateResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetNotificationRestTemplateResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEventType

`func (o *GetNotificationRestTemplateResponse) GetEventType() []string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetNotificationRestTemplateResponse) GetEventTypeOk() (*[]string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetNotificationRestTemplateResponse) SetEventType(v []string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GetNotificationRestTemplateResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *GetNotificationRestTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNotificationRestTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNotificationRestTemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNotificationRestTemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundType

`func (o *GetNotificationRestTemplateResponse) GetOutboundType() string`

GetOutboundType returns the OutboundType field if non-nil, zero value otherwise.

### GetOutboundTypeOk

`func (o *GetNotificationRestTemplateResponse) GetOutboundTypeOk() (*string, bool)`

GetOutboundTypeOk returns a tuple with the OutboundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundType

`func (o *GetNotificationRestTemplateResponse) SetOutboundType(v string)`

SetOutboundType sets OutboundType field to given value.

### HasOutboundType

`func (o *GetNotificationRestTemplateResponse) HasOutboundType() bool`

HasOutboundType returns a boolean if a field has been set.

### GetParameters

`func (o *GetNotificationRestTemplateResponse) GetParameters() []NotificationRestTemplateParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetNotificationRestTemplateResponse) GetParametersOk() (*[]NotificationRestTemplateParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetNotificationRestTemplateResponse) SetParameters(v []NotificationRestTemplateParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetNotificationRestTemplateResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTemplateType

`func (o *GetNotificationRestTemplateResponse) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *GetNotificationRestTemplateResponse) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *GetNotificationRestTemplateResponse) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *GetNotificationRestTemplateResponse) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### GetUuid

`func (o *GetNotificationRestTemplateResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetNotificationRestTemplateResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetNotificationRestTemplateResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetNotificationRestTemplateResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GetNotificationRestTemplateResponse) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GetNotificationRestTemplateResponse) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GetNotificationRestTemplateResponse) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *GetNotificationRestTemplateResponse) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetResult

`func (o *GetNotificationRestTemplateResponse) GetResult() NotificationRestTemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNotificationRestTemplateResponse) GetResultOk() (*NotificationRestTemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNotificationRestTemplateResponse) SetResult(v NotificationRestTemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNotificationRestTemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


